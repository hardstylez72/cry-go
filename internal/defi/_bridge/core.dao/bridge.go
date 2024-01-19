package core_dao

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/log"
	"github.com/hardstylez72/cry/internal/lzscan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Bridge struct {
	Cli     *defi.EtheriumClient
	HttpCli *http.Client
}

func NewBridge(cli *defi.EtheriumClient) *Bridge {
	return &Bridge{
		Cli:     cli,
		HttpCli: &http.Client{},
	}
}

func (a *Bridge) Bridge(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {
	if req.FromNetwork == v1.Network_BinanaceBNB {
		return a.bridgeFromBNB(ctx, req)
	}

	if req.FromNetwork == v1.Network_Core {
		return a.bridgeFromCore(ctx, req)
	}

	return nil, errors.New("unsupported network")
}

func (a *Bridge) bridgeFromBNB(ctx context.Context, req *defi.DefaultBridgeReq) (res *bozdo.DefaultRes, err error) {
	ca := common.HexToAddress("0x52e75D318cFB31f9A2EdFa2DFee26B161255B233")

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	limitTx, err := a.Cli.TokenLimitChecker(ctx, &defi.TokenLimitCheckerReq{
		Token:       req.FromToken,
		Wallet:      w,
		Amount:      req.Amount,
		SpenderAddr: ca,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}

	defer func() {
		if limitTx != nil && limitTx.LimitExtended {
			res.ApproveTx = a.Cli.NewTx(limitTx.ApproveTx.Hash(), defi.CodeApprove, nil)
		}
	}()

	tr, err := NewBridgebnbCaller(ca, a.Cli.Cli)
	if err != nil {
		return nil, err
	}

	fee, err := tr.EstimateBridgeFee(&bind.CallOpts{Context: ctx}, false, []byte{})
	if err != nil {
		return nil, err
	}

	abi, err := BridgebnbMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var token common.Address       //+
	var amountLD *big.Int          //+
	var to common.Address          //+
	var callParams LzLibCallParams //+
	var adapterParams []byte

	token = a.Cli.Cfg.TokenMap[req.FromToken]

	pub, err := a.Cli.GetPublicKey(req.PK, req.SubType)
	if err != nil {
		return nil, err
	}
	to = common.HexToAddress(pub)

	callParams = LzLibCallParams{
		RefundAddress:     to,
		ZroPaymentAddress: defi.ZeroAddress,
	}

	amountLD, err = defi.Slippage(req.Amount, req.Slippage)
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("bridge", token, amountLD, to, callParams, adapterParams)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(pack)))
	}

	txData := &bozdo.TxData{
		Data:         pack,
		Value:        fee.NativeFee,
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee.NativeFee, req.FromNetwork, req.FromToken)},
		Code:         bozdo.CodeBridge,
	}

	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: v1.TaskType_CoreDaoBridge,
	}

	return a.Cli.London(ctx, a.Cli, opt, txData)
}

func (a *Bridge) bridgeFromCore(ctx context.Context, req *defi.DefaultBridgeReq) (res *bozdo.DefaultRes, _ error) {
	ca := common.HexToAddress("0xA4218e1F39DA4AaDaC971066458Db56e901bcbdE")

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	limitTx, err := a.Cli.TokenLimitChecker(ctx, &defi.TokenLimitCheckerReq{
		Token:       req.FromToken,
		Wallet:      w,
		Amount:      req.Amount,
		SpenderAddr: ca,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}

	defer func() {
		if limitTx != nil && limitTx.LimitExtended {
			res.ApproveTx = a.Cli.NewTx(limitTx.ApproveTx.Hash(), defi.CodeApprove, nil)
		}
	}()

	tr, err := NewBridgecoreCaller(ca, a.Cli.Cli)
	if err != nil {
		return nil, err
	}

	destChainId, ok := layerzero.LayerZeroChainMap[req.ToNetwork]
	if !ok {
		return nil, errors.New("lz dest chain not set")
	}

	fee, err := tr.EstimateBridgeFee(&bind.CallOpts{Context: ctx}, destChainId, false, []byte{})
	if err != nil {
		return nil, err
	}

	abi, err := BridgecoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var localToken common.Address  //+
	var remoteChainId uint16       //+
	var amount *big.Int            //+
	var to common.Address          //+
	var unwrapWeth bool            //+
	var callParams LzLibCallParams //+
	var adapterParams []byte       //+

	localToken = a.Cli.Cfg.TokenMap[req.FromToken]

	remoteChainId = layerzero.LayerZeroChainMap[req.ToNetwork]

	to = w.WalletAddr

	callParams = LzLibCallParams{
		RefundAddress:     to,
		ZroPaymentAddress: defi.ZeroAddress,
	}

	amount, err = defi.Slippage(req.Amount, req.Slippage)
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("bridge", localToken, remoteChainId, amount, to, unwrapWeth, callParams, adapterParams)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(pack)))
	}

	txData := &bozdo.TxData{
		Data:         pack,
		Value:        fee.NativeFee,
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee.NativeFee, req.FromNetwork, req.FromToken)},
		Code:         bozdo.CodeBridge,
	}

	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: v1.TaskType_CoreDaoBridge,
	}

	return a.Cli.London(ctx, a.Cli, opt, txData)
}

func (a *Bridge) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {
	s := lzscan.NewService()

	_, err := s.WaitConfirm(ctx, txId)
	return err
}
