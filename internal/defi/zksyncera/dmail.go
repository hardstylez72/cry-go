package zksyncera

import (
	"context"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/abi/dmail"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/pkg/errors"
)

func (c *Client) SendDmailMessage(ctx context.Context, req *defi.SimpleReq) (*bozdo.DefaultRes, error) {

	wtx, err := NewWalletTransactor(req.PK, c.NetworkId)
	if err != nil {
		return nil, err
	}
	txData, err := DmailMakeTxData()
	if err != nil {
		return nil, err
	}
	result := &bozdo.DefaultRes{}

	tx := CreateFunctionCallTransaction(
		wtx.WalletAddr,
		txData.ContractAddr,
		big.NewInt(0),
		big.NewInt(0),
		txData.Value,
		txData.Data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, req.Gas, wtx.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	result.ECost = estimate

	if req.EstimateOnly {
		return result, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "ClientL2.SendRawTransaction")
	}

	result.Tx = c.NewTx(hash, defi.CodeContract, txData.Details)

	return result, nil
}

func DmailMakeTxData() (*bozdo.TxData, error) {
	a, err := dmail.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	to := lib.RandEmail(rand.New(lib.RandSource).Intn(20))
	subj := lib.RandString(rand.New(lib.RandSource).Intn(100))

	data, err := a.Pack("send_mail", to, subj)
	if err != nil {
		return nil, err
	}
	return &bozdo.TxData{
		Data:         data,
		Value:        big.NewInt(0),
		ContractAddr: common.HexToAddress("0x981F198286E40F9979274E0876636E9144B8FB8E"),
		Details:      nil,
	}, nil
}
