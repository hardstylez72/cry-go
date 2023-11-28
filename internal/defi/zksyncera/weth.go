package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/weth"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) SwapWETH(ctx context.Context, req *defi.WETHReq) (*defi.WETHRes, error) {
	if req.Wrap {
		return c.WrapETH(ctx, req)
	} else {
		return c.UnWrapETH(ctx, req)
	}
}

func (c *Client) WrapETH(ctx context.Context, req *defi.WETHReq) (*defi.WETHRes, error) {

	w, wtx, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, err
	}

	wethabi, err := weth.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := wethabi.Pack("deposit")
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	tx := &types.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:       w.Address(),
			To:         &c.Cfg.Weth,
			Gas:        0,
			GasPrice:   big.NewInt(0),
			GasFeeCap:  nil,
			GasTipCap:  nil,
			Value:      req.Amount,
			Data:       data,
			AccessList: nil,
		},
		Meta: &types.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
		},
	}

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &defi.WETHRes{}
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract, nil)

	return result, nil
}

func (c *Client) UnWrapETH(ctx context.Context, req *defi.WETHReq) (*defi.WETHRes, error) {

	w, wtx, err := c.Wallet(req.WalletPK)
	if err != nil {
		return nil, err
	}

	wethabi, err := weth.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := wethabi.Pack("withdraw", req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw function: %w", err)
	}

	tx := &types.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:       w.Address(),
			To:         &c.Cfg.Weth,
			Gas:        0,
			GasPrice:   big.NewInt(0),
			GasFeeCap:  nil,
			GasTipCap:  nil,
			Value:      nil,
			Data:       data,
			AccessList: nil,
		},
		Meta: &types.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
		},
	}

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, fmt.Errorf("Make712Tx: %w", err)
	}
	result := &defi.WETHRes{}
	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}
	result.Tx = c.NewTx(hash, defi.CodeContract, nil)

	return result, nil
}

func (c *Client) Make712Tx(ctx context.Context, tx *types.CallMsg, gasOpt *bozdo.Gas, signer *accounts.BaseSigner) (_ []byte, _ *bozdo.EstimatedGasCost, err error) {
	
	nonce, err := c.ClientL2.NonceAt(ctx, tx.From, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to GetGasPrice: %w", err)
	}

	tip := big.NewInt(100_000_000)

	var gas, gasPrice *big.Int
	if gasOpt.RuleSet() {
		gas = &gasOpt.GasLimit
		gasPrice = &gasOpt.GasPrice
		gasPrice = defi.ResolveGasPriceZksync(&gasOpt.MaxGas, gas, gasPrice)
	} else {

		header, err := c.ClientL2.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "ClientL2.HeaderByNumber")
		}

		gasPrice = bozdo.BigIntSum(header.BaseFee, tip)
		tx.GasFeeCap = gasPrice
		//tx.GasTipCap = tip
		if tx.Gas == 0 {
			gasLimit, err := c.ClientL2.EstimateGasL2(ctx, *tx)
			if err != nil {
				return nil, nil, errors.Wrap(err, "ClientL2.EstimateGasL2")
			}
			gas = big.NewInt(0).SetUint64(gasLimit)
		} else {
			gas = big.NewInt(0).SetUint64(tx.Gas)
		}

		if err != nil {
			return nil, nil, fmt.Errorf("failed to EstimateGas: %w", err)
		}

	}

	prepared := &types.Transaction712{
		Nonce:      new(big.Int).SetUint64(nonce),
		GasTipCap:  tip, // TODO: Estimate correct one
		GasFeeCap:  gasPrice,
		Gas:        gas,
		To:         tx.To,
		Value:      tx.Value,
		Data:       tx.Data,
		AccessList: nil,
		ChainID:    c.NetworkId,
		From:       &tx.From,
		Meta:       tx.Meta,
	}

	signature, err := signer.SignTypedData(signer.Domain(), prepared)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Signer.SignTypedData")
	}
	rawTx, err := prepared.RLPValues(signature)
	if err != nil {
		return nil, nil, errors.Wrap(err, "prepared.RLPValues")
	}

	return rawTx, &bozdo.EstimatedGasCost{
		GasLimit:    gas,
		GasPrice:    gasPrice,
		TotalGasWei: defi.MinerGasLegacy(gasPrice, gas.Uint64()),
	}, nil

}
