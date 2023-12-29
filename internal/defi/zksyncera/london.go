package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/types"
)

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
