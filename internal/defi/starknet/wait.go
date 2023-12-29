package starknet

import (
	"context"
	"fmt"
	"strings"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/starknet.go/felt"
	"github.com/hardstylez72/cry/starknet.go/types"
)

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {
	acceptedOnL2 := false
	fmt.Println("Polling until transaction is accepted on L2...")
	for !acceptedOnL2 {
		ff := &felt.Felt{}
		ff.SetString(tx)
		receipt, err := c.GWP.WaitForTransaction(ctx, ff, 5)
		if err != nil {
			if strings.Contains(err.Error(), "tx not finalized") {
				return defi.ErrTxNotFound
			} else if strings.Contains(err.Error(), "REVERTED") {
				return defi.ErrTxStatusFailed
			} else {
				return fmt.Errorf("Transaction Failure (%s): can't poll to desired status: %s", tx, err.Error())
			}

		}
		if receipt.String() == types.TransactionRejected.String() {
			return defi.ErrTxStatusFailed
		}

		if receipt.String() == types.TransactionAcceptedOnL2.String() || receipt.String() == "ACCEPTED_ON_L2" {
			acceptedOnL2 = true
		}
	}
	return nil
}
