package zksyncera

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_nft/merkly"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/pkg/errors"
)

func (c *Client) BridgeNftMerkly(ctx context.Context, req *defi.BridgeNFTReq) (*bozdo.DefaultRes, error) {
	m := &merkly.Maker{
		TokenMap: c.Cfg.TokenMap,
		Cli:      c.ClientL2,
		Network:  c.Cfg.Network,
		CA:       SpecMap["merkly"].Addr,
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

	result := &bozdo.DefaultRes{}

	tx := CreateFunctionCallTransaction(
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

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract, txData.Details)

	return result, nil
}

func (c *Client) GetNFTId(ctx context.Context, txHash common.Hash, targetOwner common.Address) (*big.Int, error) {

	m := &merkly.Maker{
		TokenMap: c.Cfg.TokenMap,
		Cli:      c.ClientL2,
		Network:  c.Cfg.Network,
		CA:       SpecMap["merkly"].Addr,
	}

	//maxId := big.NewInt(1999999)
	minId := big.NewInt(1000000)
	tx, _, err := c.ClientL2.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, err
	}

	logs, err := c.ClientL2.FilterLogsL2(ctx, ethereum.FilterQuery{
		BlockHash: tx.BlockHash,
		Addresses: []common.Address{tx.To},
	})
	if err != nil {
		return nil, err
	}

	for _, log := range logs {
		if log.Address.String() != SpecMap["merkly"].Addr.String() {
			continue
		}
		for _, topic := range log.Topics {
			topicBig := topic.Big()
			if topicBig.Cmp(minId) >= 0 { // &&  maxId.Cmp(topicBig) >= 0

				owner, err := m.GetOwner(ctx, topicBig)
				if err != nil {
					continue
				}

				if owner.String() == targetOwner.String() {
					return topicBig, nil
				}
			}
		}
	}

	return nil, errors.New("nft not found")
}
