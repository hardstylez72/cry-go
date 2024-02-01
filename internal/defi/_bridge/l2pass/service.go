package l2pass

import (
	"context"
	"encoding/json"
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

var cli = http.Client{}

func (a *Bridge) Prices() (map[string]map[string]float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=canto,ethereum,moonbeam,moonriver,coredaoorg,klay-token,conflux-token,celo,metis-token,ethereum,kava,ethereum,fantom,ethereum,matic-network,ethereum,ethereum,xdai,ethereum,fuse-network-token,mantle,ethereum,ethereum,binancecoin,binancecoin,harmony,avalanche-2,ethereum&vs_currencies=usd"

	get, err := cli.Get(url)
	if err != nil {
		return nil, err
	}

	m := map[string]map[string]float64{}
	err = json.NewDecoder(get.Body).Decode(&m)
	if err != nil {
		return nil, err
	}

	return m, nil
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

	caller, err := NewRefuelCaller(a.CA, a.Cli.Cli)
	if err != nil {
		return nil, err
	}

	abi, err := RefuelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var _dstChainId uint16               //+
	var zroPaymentAddress common.Address //+
	var nativeForDst *big.Int            //+
	var addressOnDst common.Address      // +

	_dstChainId = destChainId
	zroPaymentAddress = defi.ZeroAddress
	addressOnDst = w.WalletAddr

	switch true {
	case req.FromNetwork == v1.Network_POLIGON && req.ToNetwork == v1.Network_Fuse:
		nativeForDst = big.NewInt(50000000000000000)
	case req.FromNetwork == v1.Network_POLIGON && req.ToNetwork == v1.Network_Loot:
		nativeForDst = big.NewInt(50000000000000000)
	case req.FromNetwork == v1.Network_POLIGON && req.ToNetwork == v1.Network_Klaytn:
		nativeForDst = big.NewInt(50000000000000000)
	default:
		return nil, errors.New("не поддерживается")
	}

	fee, err := caller.EstimateGasRefuelFee(&bind.CallOpts{Context: ctx}, destChainId, nativeForDst, addressOnDst, false)
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("gasRefuel", _dstChainId, zroPaymentAddress, nativeForDst, addressOnDst)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug("l2pass refuel data: 0x" + common.Bytes2Hex(pack))
	}

	txData := &bozdo.TxData{
		Data:         pack,
		Value:        bozdo.BigIntSum(fee.NativeFee),
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee.NativeFee, req.FromNetwork, req.FromToken)},
		Code:         bozdo.CodeBridge,
	}

	if req.Debug {
		log.Log.Debug("value: ", txData.Value.String())
	}
	opt := &defi.TxOpt{
		NoSend:   req.EstimateOnly,
		Pk:       req.PK,
		Gas:      req.Gas,
		Debug:    req.Debug,
		TaskType: v1.TaskType_L2PassRefuel,
	}

	return a.Cli.London(ctx, a.Cli, opt, txData)
}

func (a *Bridge) WaitForConfirm(ctx context.Context, txId string) error {
	s := lzscan.NewService()

	_, err := s.WaitConfirm(ctx, txId)
	return err
}
