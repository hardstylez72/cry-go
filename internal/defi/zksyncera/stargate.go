package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/router"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/startgatemavrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type stargateBridgeMaker struct {
	source *Client
}

func (c *Client) StargateBridge(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {
	return c.GenericBridge(ctx, &stargateBridgeMaker{c}, req)
}

func (m stargateBridgeMaker) MakeBridgeTx(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.TxData, error) {

	c := m.source

	contractAddr := common.HexToAddress("0xDAc7479e5F7c01CC59bbF7c1C4EDF5604ADA1FF2")

	wt, err := NewWalletTransactor(req.PK, c.Cfg.networkId)
	if err != nil {
		return nil, err
	}

	fromToken, ok := c.Cfg.TokenMap[req.FromToken]
	if err != nil {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	bps := big.NewInt(0)

	var _oft common.Address                          //+
	var _dstChainId uint16                           //+
	var _toAddress []byte                            //+
	var _amount *big.Int                             //+
	var _minAmount *big.Int                          //+
	var _refundAddress common.Address                //+
	var _zroPaymentAddress common.Address            //+
	var _adapterParams []byte                        //+
	var _feeObj startgatemavrouter.IOFTWrapperFeeObj // +

	abis, err := startgatemavrouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	dist, ok := layerzero.LayerZeroChainMap[req.ToNetwork]
	if !ok {
		return nil, errors.New("network is not supported: " + req.ToNetwork.String())
	}
	_oft = fromToken
	_dstChainId = dist
	_toAddress = wt.WalletAddr.Bytes()
	_amount = req.Amount

	amSlip, err := defi.Slippage(req.Amount, req.Slippage)
	if err != nil {
		return nil, err
	}

	_minAmount = amSlip
	_refundAddress = wt.WalletAddr
	_zroPaymentAddress = ZEROADDR
	_feeObj = startgatemavrouter.IOFTWrapperFeeObj{
		CallerBps: bps,
		Caller:    ZEROADDR,
		PartnerId: [2]byte{0, 0},
	}

	adapterParams, err := layerzero.MakeLayerZeroAdapterParams(1, big.NewInt(100000))
	if err != nil {
		return nil, err
	}
	//0x00010000000000000000000000000000000000000000000000000000000000061a80
	//_adapterParams = common.Hex2Bytes("0x000100000000000000000000000000000000000000000000000000000000000186a0")
	_adapterParams = adapterParams

	fee, err := m.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain:  req.ToNetwork,
		WalletPK: req.PK,
	})
	if err != nil {
		return nil, err
	}

	value := fee.Fee1

	data, err := abis.Pack(
		"sendOFT",
		_oft,
		_dstChainId,
		_toAddress,
		_amount,
		_minAmount,
		_refundAddress,
		_zroPaymentAddress,
		_adapterParams,
		_feeObj,
	)
	if err != nil {
		return nil, err
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: contractAddr,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee.Fee1, c.Cfg.Network, v1.Token_ETH)},
	}, nil
}

type GetStargateBridgeFeeReq struct {
	ToChain  v1.Network
	WalletPK string
}

type GetStargateBridgeFeeRes struct {
	Fee1 *big.Int
	Fee2 *big.Int
}

func (m stargateBridgeMaker) GetStargateBridgeFee(ctx context.Context, req *GetStargateBridgeFeeReq) (*GetStargateBridgeFeeRes, error) {

	c := m.source

	w, _, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	provider := c.ClientL1
	if err != nil {
		return nil, errors.Wrap(err, "CreateEthereumProvider")
	}

	trx, err := router.NewRouterCaller(common.HexToAddress("0x8731d54E9D02c286767d56ac03e8037C07e01e98"), provider)
	if err != nil {
		return nil, err
	}

	toAddress := w.Address().Bytes()
	payload := common.HexToAddress("0").Bytes()
	distChain, ok := layerzero.LayerZeroChainMap[req.ToChain]
	if !ok {
		return nil, errors.New("invalid chain: " + string(req.ToChain))
	}

	opt := &bind.CallOpts{
		Context: ctx,
	}

	fee1, fee2, err := trx.QuoteLayerZeroFee(opt, distChain, layerzero.TypeFuncSwap, toAddress, payload, router.IStargateRouterlzTxObj{
		DstGasForCall:   big.NewInt(0),
		DstNativeAmount: big.NewInt(0),
		DstNativeAddr:   []byte{},
	})
	if err != nil {
		return nil, err
	}

	return &GetStargateBridgeFeeRes{
		Fee1: fee1,
		Fee2: fee2,
	}, nil
}
