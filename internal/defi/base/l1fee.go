package base

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/optimism_fee"
)

func (c *Client) EstimateL1GasFee(ctx context.Context, data []byte) (*big.Int, error) {
	optFeeCaller, err := optimism_fee.NewStorageCaller(common.HexToAddress("0x420000000000000000000000000000000000000F"), c.defi.Cli)
	if err != nil {
		return nil, err
	}

	o := &bind.CallOpts{}
	o.Context = ctx
	l1Gasfee, err := optFeeCaller.GetL1Fee(o, data)
	if err != nil {
		return nil, err
	}
	return l1Gasfee, nil
}

func TxGas(tx *types.Transaction) string {

	s := fmt.Sprintf("gas:%d gasPrice:%s GWEI maxFee:%s GWEI tipCap:%s GWEI",
		tx.Gas(), Gwei(tx.GasPrice()).String(), Gwei(tx.GasFeeCap()).String(), Gwei(tx.GasTipCap()).String())
	return s
}
