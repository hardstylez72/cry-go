package merkly

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

// https://explorer.zksync.io/address/0x6dd28C2c5B91DD63b4d4E78EcAC7139878371768#contract
//go:generate abigen --abi mint_abi.json --pkg merkly --type Minter --out mint_abi.go

type Maker struct {
	TokenMap map[v1.Token]common.Address
	Cli      *ethclient.Client
	Network  v1.Network
	CA       common.Address
}

type MintNFTReq struct {
	WalletPK string

	bozdo.BaseReq
}

type BridgeNFTReq struct {
	WalletPK    string
	FromNetwork v1.Network
	ToNetwork   v1.Network
	NFTId       *big.Int
	bozdo.BaseReq
}

// site https://explorer.zksync.io/tx/0xa762372ccb04ca4495a9d90915b27fa44b2c402c63455f5a81e6312506abddb8
func (m *Maker) MakeMintTx(ctx context.Context) (_ *bozdo.TxData, nftId *big.Int, _ error) {
	c, err := NewMinterCaller(m.CA, m.Cli)
	if err != nil {
		return nil, nil, err
	}

	opt := &bind.CallOpts{Context: ctx}

	fee, err := c.Fee(opt)
	if err != nil {
		return nil, nil, err
	}

	mintId, err := c.NextMintId(opt)
	if err != nil {
		return nil, nil, err
	}

	a, err := MinterMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	data, err := a.Pack("mint")
	if err != nil {
		return nil, nil, err
	}

	tx := &bozdo.TxData{
		Data:         data,
		Value:        fee,
		ContractAddr: m.CA,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee, m.Network, v1.Token_ETH)},
	}

	return tx, mintId, nil
}

type MakeBridgeTxReq struct {
	W           common.Address
	FromNetwork v1.Network
	ToNetwork   v1.Network
	NFTId       *big.Int
}

// site https://explorer.zksync.io/tx/0x7c14129f2d862fa7e50cfddb390ebe95c60a205c70dc00d4f67d1466e9b9b479
func (m *Maker) MakeBridgeTx(ctx context.Context, req *MakeBridgeTxReq) (*bozdo.TxData, error) {

	var _from common.Address              //+
	var _dstChainId uint16                //+
	var _toAddress []byte                 //+
	var _tokenId *big.Int                 //+
	var _refundAddress common.Address     //+
	var _zroPaymentAddress common.Address //+
	var _adapterParams []byte             //+

	_from = req.W

	dstchainId, ok := layerzero.LayerZeroChainMap[req.ToNetwork]
	if !ok {
		return nil, errors.New("network is not supported: " + req.ToNetwork.String())
	}
	_dstChainId = dstchainId
	_toAddress = req.W.Bytes()
	_tokenId = req.NFTId
	_refundAddress = req.W
	_zroPaymentAddress = bozdo.ZEROADDR

	adapterParams, err := layerzero.MakeLayerZeroAdapterParams(1, big.NewInt(400000))
	if err != nil {
		return nil, err
	}
	_adapterParams = adapterParams

	c, err := NewMinterCaller(m.CA, m.Cli)
	if err != nil {
		return nil, err
	}
	opt := &bind.CallOpts{Context: ctx}
	f, err := c.EstimateGasBridgeFee(opt, _dstChainId, false, adapterParams)
	if err != nil {
		return nil, err
	}
	f.NativeFee = bozdo.BigIntSum(f.NativeFee, bozdo.Percent(f.NativeFee, 20))
	a, err := MinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := a.Pack(
		"sendFrom",
		_from,
		_dstChainId,
		_toAddress,
		_tokenId,
		_refundAddress,
		_zroPaymentAddress,
		_adapterParams,
	)
	if err != nil {
		return nil, err
	}

	tx := &bozdo.TxData{
		Data:         data,
		Value:        f.NativeFee,
		ContractAddr: m.CA,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(f.NativeFee, m.Network, v1.Token_ETH)},
	}

	return tx, nil
}
