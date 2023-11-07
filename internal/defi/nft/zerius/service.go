package zerius

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/pkg/errors"
)

// https://basescan.org/address/0x178608fFe2Cca5d36f3Fc6e69426c4D3A5A74A41#code
//go:generate abigen --abi abi.json --pkg zerius --type Minter --out abi.go

type Client struct {
	cli *defi.EtheriumClient
}

func NewClient(cli *defi.EtheriumClient) *Client {
	return &Client{
		cli: cli,
	}
}

func (c *Client) MakeMintTx(ctx context.Context, req *defi.SimpleReq) (*bozdo.TxData, error) {

	ca := c.cli.Cfg.Dict.Zerius.Minter

	caller, err := NewMinterCaller(ca, c.cli.Cli)
	if err != nil {
		return nil, errors.Wrap(err, "NewMinterCaller")
	}

	fee, err := caller.MintFee(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "caller.MintFee")
	}

	abi, err := MinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	pack, err := abi.Pack("mint")
	if err != nil {
		return nil, err
	}

	return &bozdo.TxData{
		Data:         pack,
		Value:        fee,
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewProtocolFeeDetails(fee, c.cli.Network(), c.cli.Cfg.MainToken)},
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     nil,
		Code:         bozdo.CodeMint,
	}, nil
}

//func (c *Client) BridgeNft(ctx context.Context, req *defi.BridgeNFTReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
//
//}
//func (c *Client)  GetNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error) {
//
//}
