package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) MerklyMintNft(ctx context.Context, req *merkly.MintNFTReq) (*defi.DefaultRes, *big.Int, *big.Int, error) {
	m := &merkly.Maker{
		TokenMap: c.Cfg.TokenMap,
		Cli:      c.Provider.GetClient(),
		Network:  c.Cfg.Network,
		CA:       common.HexToAddress("0x6dd28C2c5B91DD63b4d4E78EcAC7139878371768"),
	}

	txData, mintId, err := m.MakeMintTx(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	result := &defi.DefaultRes{}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, nil, nil, err
	}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		txData.ContractAddr,
		big.NewInt(0),
		big.NewInt(0),
		txData.Value,
		txData.Data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Make712Tx")
	}

	estimate.Name = "Mint"
	estimate.Details = txData.Details
	result.ECost = estimate

	if req.EstimateOnly {
		return result, mintId, txData.Value, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract, txData.Details)

	return result, mintId, txData.Value, nil
}
func (c *Client) MerklyBridgeNft(ctx context.Context, req *merkly.BridgeNFTReq) (*defi.DefaultRes, error) {
	m := &merkly.Maker{
		TokenMap: c.Cfg.TokenMap,
		Cli:      c.Provider.GetClient(),
		Network:  c.Cfg.Network,
		CA:       common.HexToAddress("0x6dd28C2c5B91DD63b4d4E78EcAC7139878371768"),
	}

	transactor, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	txData, err := m.MakeBridgeTx(ctx, &merkly.MakeBridgeTxReq{
		W:           transactor.WalletAddr,
		FromNetwork: req.FromNetwork,
		ToNetwork:   req.ToNetwork,
		NFTId:       req.NFTId,
	})
	if err != nil {
		return nil, err
	}

	result := &defi.DefaultRes{}

	tx := utils.CreateFunctionCallTransaction(
		transactor.WalletAddr,
		txData.ContractAddr,
		big.NewInt(0),
		big.NewInt(0),
		txData.Value,
		txData.Data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, transactor.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	estimate.Name = "Bridge"
	estimate.Details = txData.Details
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.Provider.SendRawTransaction(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Provider.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract, txData.Details)

	return result, nil
}
