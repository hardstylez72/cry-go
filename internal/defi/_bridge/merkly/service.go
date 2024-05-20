package merkly

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	"github.com/hardstylez72/cry/internal/lzscan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type Bridge struct {
	Cli *defi.EtheriumClient
	CA  common.Address
}

func NewBridge(cli *defi.EtheriumClient, CA common.Address) *Bridge {
	return &Bridge{
		Cli: cli,
		CA:  CA,
	}
}

func (a *Bridge) Bridge(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {

	ca := a.CA

	destChainId, ok := layerzero.LayerZeroChainMap[req.ToNetwork]
	if !ok {
		return nil, errors.New("lz dest chain not set")
	}

	w, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	caller, err := NewMerklyrefuel(a.CA, a.Cli.Cli)
	if err != nil {
		return nil, err
	}

	abi, err := MerklyrefuelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var _dstChainId uint16    //+
	var _toAddress []byte     //+
	var _adapterParams []byte //+

	_toAddress = w.WalletAddr.Bytes()
	_dstChainId = destChainId

	var NativeForDst *big.Int
	var boba *big.Int

	switch true {
	case req.FromNetwork == v1.Network_ARBITRUM && req.ToNetwork == v1.Network_POLIGON:
		div := pub.Price().MATIC / pub.Price().ETH
		// 1 matic
		boba, _ = big.NewFloat(0).Mul(big.NewFloat(div), big.NewFloat(params.Ether/10)).Int(nil)
		NativeForDst = big.NewInt(params.Ether / 10)
	case req.FromNetwork == v1.Network_POLIGON && req.ToNetwork == v1.Network_Conflux:
		div := pub.Price().CFX / pub.Price().MATIC
		// 1 matic
		boba, _ = big.NewFloat(0).Mul(big.NewFloat(div), big.NewFloat(params.Ether/20)).Int(nil)
		NativeForDst = big.NewInt(50000000000000000)

	case req.FromNetwork == v1.Network_BinanaceBNB && req.ToNetwork == v1.Network_Celo:
		div := pub.Price().CFX / pub.Price().BNB
		// 1 matic
		boba, _ = big.NewFloat(0).Mul(big.NewFloat(div), big.NewFloat(params.Ether/20)).Int(nil)
		NativeForDst = big.NewInt(1000000000000000)
	case req.FromNetwork == v1.Network_BinanaceBNB && req.ToNetwork == v1.Network_DFK:
		div := pub.Price().CFX / pub.Price().BNB
		// 1 matic
		boba, _ = big.NewFloat(0).Mul(big.NewFloat(div), big.NewFloat(params.Ether/20)).Int(nil)
		NativeForDst = big.NewInt(1000000000000000)
	default:
		return nil, errors.New("не поддерживается")
	}

	_adapterParams, err = layerzero.MakeLayerZeroAdapterParamsV2(layerzero.AdapterParamsV2{
		Version:      2,
		GasAmount:    big.NewInt(200_000),
		NativeForDst: NativeForDst,
		AddressOnDst: w.WalletAddr,
	})
	if err != nil {
		return nil, err
	}

	fee, err := caller.EstimateSendFee(&bind.CallOpts{Context: ctx}, destChainId, []byte{}, _adapterParams)
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("bridgeGas", _dstChainId, _toAddress, _adapterParams)
	if err != nil {
		return nil, err
	}

	txData := &bozdo.TxData{
		Data:         pack,
		Value:        bozdo.BigIntSum(fee.NativeFee, boba),
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

func (a *Bridge) WaitForConfirm(ctx context.Context, txId string) error {
	return nil
	s := lzscan.NewService()

	_, err := s.WaitConfirm(ctx, txId)
	return err
}
