package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	"github.com/pkg/errors"
)

func (c *EtheriumClient) newTxOpt(ctx context.Context, pk string) (*bind.TransactOpts, error) {
	w, err := newWalletTransactor(pk)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(w.PrivateKey, c.Cfg.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx
	return opt, nil
}

func (c *EtheriumClient) MerklyMintNft(ctx context.Context, req *merkly.MintNFTReq) (*bozdo.DefaultRes, error) {
	t, err := merkly.NewMinterTransactor(c.Cfg.Dict.Merkly.NFT, c.Cli)
	if err != nil {
		return nil, err
	}

	opt, err := c.newTxOpt(ctx, req.WalletPK)
	if err != nil {
		return nil, err
	}

	opt.NoSend = req.EstimateOnly

	caller, err := merkly.NewMinterCaller(c.Cfg.Dict.Merkly.NFT, c.Cli)
	if err != nil {
		return nil, err
	}
	fee, err := caller.Fee(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	opt.Value = fee

	opt = c.ResoleGas(ctx, req.Gas, opt)

	tx, err := t.Mint(opt)
	if err != nil {
		return nil, err
	}

	res := &bozdo.DefaultRes{}

	details := []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee, c.Cfg.Network, c.Cfg.MainToken)}

	res.ECost = Estimate(tx, nil, "mint", details)

	if !req.EstimateOnly {
		res.Tx = c.NewTx(tx.Hash(), bozdo.CodeMint, details)
	}

	return res, nil
}
func (c *EtheriumClient) MerklyBridgeNft(ctx context.Context, req *merkly.BridgeNFTReq) (*bozdo.DefaultRes, error) {

	var _from common.Address              //+
	var _dstChainId uint16                //+
	var _toAddress []byte                 //+
	var _tokenId *big.Int                 //+
	var _refundAddress common.Address     //+
	var _zroPaymentAddress common.Address //+
	var _adapterParams []byte             //+

	tr, err := newWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, err
	}
	_from = tr.WalletAddr

	dstchainId, ok := layerzero.LayerZeroChainMap[req.ToNetwork]
	if !ok {
		return nil, errors.New("network is not supported: " + req.ToNetwork.String())
	}
	_dstChainId = dstchainId
	_toAddress = tr.WalletAddr.Bytes()
	_tokenId = req.NFTId
	_refundAddress = tr.WalletAddr
	_zroPaymentAddress = bozdo.ZEROADDR

	adapterParams, err := layerzero.MakeLayerZeroAdapterParams(1, big.NewInt(400000))
	if err != nil {
		return nil, err
	}
	_adapterParams = adapterParams

	caller, err := merkly.NewMinterCaller(c.Cfg.Dict.Merkly.NFT, c.Cli)
	if err != nil {
		return nil, err
	}
	f, err := caller.EstimateGasBridgeFee(&bind.CallOpts{Context: ctx}, _dstChainId, false, adapterParams)
	if err != nil {
		return nil, err
	}
	f.NativeFee = bozdo.BigIntSum(f.NativeFee, bozdo.Percent(f.NativeFee, layerzero.LayerZeroBoostPercent))

	opt, err := c.newTxOpt(ctx, req.WalletPK)
	if err != nil {
		return nil, err
	}
	opt.Value = f.NativeFee

	t, err := merkly.NewMinterTransactor(c.Cfg.Dict.Merkly.NFT, c.Cli)
	if err != nil {
		return nil, err
	}

	opt.NoSend = req.EstimateOnly

	tx, err := t.SendFrom(opt, _from,
		_dstChainId,
		_toAddress,
		_tokenId,
		_refundAddress,
		_zroPaymentAddress,
		_adapterParams)
	if err != nil {
		return nil, err
	}

	res := &bozdo.DefaultRes{}

	details := []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(f.NativeFee, c.Cfg.Network, c.Cfg.MainToken)}

	res.ECost = Estimate(tx, nil, "bridge", details)

	if !req.EstimateOnly {
		res.Tx = c.NewTx(tx.Hash(), bozdo.CodeBridge, details)
	}

	return res, nil
}

func (c *EtheriumClient) GetMerklyNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error) {
	minId := big.NewInt(1000000)

	receipt, err := c.Cli.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	tx, _, err := c.Cli.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}

	logs, err := c.Cli.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &receipt.BlockHash,
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []common.Address{*tx.To()},
		Topics:    nil,
	})
	if err != nil {
		return nil, err
	}

	caller, err := merkly.NewMinterCaller(c.Cfg.Dict.Merkly.NFT, c.Cli)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		//if log.Address.String() != c.Cfg.Dict.Merkly.NFT.String() {
		//	continue
		//}
		for _, topic := range log.Topics {

			topicBig := topic.Big()
			if topicBig.Cmp(minId) >= 0 { // &&  maxId.Cmp(topicBig) >= 0

				owner1, err := caller.OwnerOf(&bind.CallOpts{Context: ctx}, topicBig)
				if err != nil {
					continue
				}
				//24134178
				if owner.String() == owner1.String() {
					return topicBig, nil
				}
			}

		}
	}

	return nil, errors.New("nft not found")
}
