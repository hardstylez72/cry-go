package zksyncera

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/izumiquoter"
	"github.com/hardstylez72/cry/internal/defi/contracts/zksyncera/izumirouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type izumiMaker struct {
	*Client
}

// eth -> usdc
//

// site ETH -> USDC
// me USDC -> ETH
// me ETH -> USDC https://explorer.zksync.io/tx/0x36731ad49420976e96d3a8889a78d5ca3661aeae1cb0f73a28fdaa4794941716

func (c *Client) IzumiSwap(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.DefaultRes, error) {
	return c.GenericSwap(ctx, &izumiMaker{c}, req)
}

func (c *izumiMaker) MakeSwapTx(ctx context.Context, req *defi.DefaultSwapReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	w, err := NewWalletTransactor(req.WalletPK, c.NetworkId)
	if err != nil {
		return nil, err
	}

	fromTokenAddr, supported := c.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	toTokenAddr, supported := c.Cfg.TokenMap[req.ToToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.ToToken)
	}

	izumiquotertr, err := izumiquoter.NewStorageCaller(c.Cfg.IZUMI.Quoter, c.ClientL2)
	if err != nil {
		return nil, err
	}

	opt := &bind.CallOpts{Context: ctx}

	recipient := w.WalletAddr
	if req.ToToken == v1.Token_ETH {
		recipient = ZEROADDR
	}

	fee := big.NewInt(2000)

	path := makeIzumiPath(fromTokenAddr, toTokenAddr, fee)

	acq, err := izumiquotertr.SwapAmount(opt, req.Amount, path)
	if err != nil {
		return nil, err
	}

	params := izumirouter.SwapSwapAmountParams{
		Path:        path,
		Recipient:   recipient,
		Amount:      req.Amount,
		MinAcquired: acq.Acquire,
		Deadline:    DefaultDeadLine(),
	}

	constractabi, err := izumirouter.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	p1, err := constractabi.Pack("swapAmount", params)
	if err != nil {
		return nil, err
	}

	if req.FromToken == v1.Token_WETH || req.ToToken == v1.Token_WETH {
		return &bozdo.TxData{
			Data:         p1,
			Value:        value,
			ContractAddr: c.Cfg.IZUMI.Router,
			Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, acq.Acquire),
		}, nil
	}

	var p2 []byte
	if req.ToToken == v1.Token_ETH {
		p2, err = constractabi.Pack("unwrapWETH9", big.NewInt(0), w.WalletAddr)
		if err != nil {
			return nil, err
		}
	}

	if req.FromToken == v1.Token_ETH {
		p2, err = constractabi.Pack("refundETH")
		if err != nil {
			return nil, err
		}
	}

	if len(p2) == 0 {
		data, err := constractabi.Pack("multicall", [][]byte{p1})
		if err != nil {
			return nil, err
		}

		return &bozdo.TxData{
			Data:         data,
			Value:        value,
			ContractAddr: c.Cfg.IZUMI.Router,
			Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, acq.Acquire),
		}, nil
	}

	data, err := constractabi.Pack("multicall", [][]byte{p1, p2})
	if err != nil {
		return nil, err
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: c.Cfg.IZUMI.Router,
		Rate:         CalcRate(req.FromToken, req.ToToken, req.Amount, acq.Acquire),
	}, nil
}

// site 5aea5775959fbc2557cc8789bc1bf90a239d9a91 0007d0 3355df6d4c9c3035724fd0e3914de96a5a83aaf4
// me	5aea5775959fbc2557cc8789bc1bf90a239d9a91 0007d0 3355df6d4c9c3035724fd0e3914de96a5a83aaf4
// 5aea5775959fbc2557cc8789bc1bf90a239d9a91 0007d0 16a9494e257703797d747540f01683952547ee5b
func makeIzumiPath(from, to common.Address, fee *big.Int) []byte {

	feeHex := fmt.Sprintf("%06X", fee.Int64())
	feeB := common.Hex2Bytes(feeHex)
	pathb := make([]byte, 0, 400)

	pathb = append(pathb, from.Bytes()...)
	pathb = append(pathb, feeB...)
	pathb = append(pathb, to.Bytes()...)

	return pathb
}
