package starknet

import (
	"context"
	"fmt"
	"strings"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/starknet.go/gateway"
	"github.com/hardstylez72/cry/starknet.go/types"
)

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {
	acceptedOnL2 := false
	var receipt *gateway.TransactionReceipt
	var err error
	fmt.Println("Polling until transaction is accepted on L2...")
	for !acceptedOnL2 {
		_, receipt, err = c.GW.WaitForTransaction(ctx, tx, 1, 60)
		if err != nil {
			if strings.Contains(err.Error(), "tx not finalized") {
				return defi.ErrTxNotFound
			} else {
				return fmt.Errorf("Transaction Failure (%s): can't poll to desired status: %s", tx, err.Error())
			}

		}
		if receipt.Status == types.TransactionRejected {
			return defi.ErrTxStatusFailed
		}

		if receipt.Status == types.TransactionAcceptedOnL2 {
			acceptedOnL2 = true
		}
	}
	return nil
}
