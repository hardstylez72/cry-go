package zksyncera

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/orbiter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
)

type GetBalanceReq struct {
	WalletAddress common.Address
	Token         v1.Token
}

type GetBalanceRes struct {
	Req           *GetBalanceReq
	WEI           *big.Int
	ETHER         *big.Float
	HumanReadable string
	Float         float64
}

func (c *Client) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	w, err := NewWalletTransactor(pk, big.NewInt(324))
	if err != nil {
		return "", err
	}
	return w.WalletAddrHR, nil
}

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {

	if req.Token == c.Cfg.MainToken {
		b, err := c.ClientL2.BalanceAt(ctx, common.HexToAddress(req.WalletAddress), nil)
		if err != nil {
			return nil, err
		}

		f, _ := defi.WEIToEther(b).Float64()
		return &defi.GetBalanceRes{
			Req:           req,
			WEI:           b,
			ETHER:         defi.WEIToEther(b),
			HumanReadable: defi.WEIToEther(b).String(),
			Float:         f,
		}, nil
	}

	ta, ok := c.Cfg.TokenMap[req.Token]
	if !ok {
		return nil, defi.ErrTokenNotSupportedFn(req.Token)
	}

	token, err := erc20.NewIERC20(ta, c.ClientL2)
	if err != nil {
		return nil, err
	}

	b, err := token.BalanceOf(nil, common.HexToAddress(req.WalletAddress))
	if err != nil {
		return nil, err
	}

	f, _ := defi.WEIToEther(b).Float64()
	res := &defi.GetBalanceRes{
		Req:           req,
		WEI:           b,
		ETHER:         defi.WeiToToken(b, req.Token),
		HumanReadable: defi.WeiToToken(b, req.Token).String(),
		Float:         f,
	}

	return res, nil
}

func (c *Client) TxViewFn(id string) string {
	return c.Cfg.TxViewFn(id)
}

func (c *Client) GetNetworkToken() v1.Token {
	return c.Cfg.MainToken
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {

	txHash := common.HexToHash(tx)
	count := 0
	max := 60

	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		count++
		receipt, err := c.ClientL2.TransactionReceipt(ctx, txHash)
		if errors.Is(err, ethereum.NotFound) {
			if count > max {
				return defi.ErrTxNotFound
			}
			continue
		}

		if err != nil {
			return err
		}

		if receipt.Status == 0 && receipt.BlockNumber != nil {
			return defi.ErrTxStatusFailed
		}

		if receipt.Status == 1 && receipt.BlockNumber != nil {
			return nil
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-queryTicker.C:
		}
	}
}

type L1L2BridgeReq struct {
	Amount       *big.Int
	WalletPK     string
	EstimateOnly bool
	Gas          *bozdo.Gas
}

type L1L2BridgeRes struct {
	TxHash           *bozdo.Transaction
	EstimatedGasCost *bozdo.EstimatedGasCost
}

func (c *Client) OrbiterBridge(ctx context.Context, req *defi.OrbiterBridgeReq) (*defi.OrbiterBridgeRes, error) {

	r := &defi.OrbiterBridgeRes{}

	opt, err := req.OrbiterService.MakeTx(&orbiter.MakeTxReq{
		FromNetwork: req.FromNetwork,
		ToNetwork:   req.ToNetwork,
		FromToken:   req.FromToken,
		ToToken:     req.ToToken,
		Amount:      req.Amount,
	})
	if err != nil {
		return nil, err
	}
	req.Amount = orbiter.WrapValueWei(req.Amount, opt.Value4DigitCode)

	res, err := c.Transfer(ctx, &defi.TransferReq{
		Pk:           req.WalletPk,
		ToAddr:       (opt.MakerReceiverAddr),
		Token:        req.FromToken,
		Amount:       req.Amount,
		Gas:          req.Gas,
		EstimateOnly: req.EstimateOnly,
	})
	if err != nil {
		return nil, err
	}
	r.Tx = res.Tx
	r.ECost = res.ECost

	return r, nil
}

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	if r.Token == c.Cfg.MainToken {
		return c.TransferMainToken(ctx, r)
	}
	return c.TransferToken(ctx, r)
}

func (c *Client) TransferToken(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	res := &defi.TransferRes{}

	to, supported := c.Cfg.TokenMap[r.Token]
	if !supported {
		return nil, defi.ErrTokenNotSupported
	}

	wtx, err := NewWalletTransactor(r.Pk, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "newWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.PKb, &c.ClientL2, c.ClientL1)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	addr := common.HexToAddress(r.ToAddr)

	_, err = c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       r.Token,
		WalletPK:    r.Pk,
		Amount:      r.Amount,
		SpenderAddr: to,
		Gas:         nil,
	})
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.TokenLimitChecker")
	}

	tokenAbi, err := abi.JSON(strings.NewReader(erc20.IERC20MetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to load ethTokenAbi: %w", err)
	}

	data, err := tokenAbi.Pack("transfer", addr, r.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack transfer function: %w", err)
	}
	tx := CreateFunctionCallTransaction(
		w.Address(),
		to,
		big.NewInt(0),
		big.NewInt(0),
		r.Amount,
		data,
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, r.Gas, wtx.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	res.ECost = estimate

	if r.EstimateOnly {
		return res, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	res.Tx = c.NewTx(hash, defi.CodeTransfer, nil)

	return res, nil

}
func (c *Client) TransferMainToken(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {

	res := &defi.TransferRes{}

	wtx, err := NewWalletTransactor(r.Pk, c.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "newWalletTransactor")
	}

	w, err := accounts.NewWallet(wtx.PKb, &c.ClientL2, c.ClientL1)
	if err != nil {
		return nil, errors.Wrap(err, "zksync2.NewWallet")
	}

	to := common.HexToAddress(r.ToAddr)

	tx := CreateFunctionCallTransaction(
		w.Address(),
		to,
		big.NewInt(0),
		big.NewInt(0),
		r.Amount,
		[]byte{},
		nil, nil,
	)

	raw, estimate, err := c.Make712Tx(ctx, tx, r.Gas, wtx.Signer)
	if err != nil {
		return nil, errors.Wrap(err, "Make712Tx")
	}

	res.ECost = estimate

	if r.EstimateOnly {
		return res, nil
	}

	hash, err := c.ClientL2.SendRawTransaction(ctx, raw)
	if err != nil {
		return nil, errors.Wrap(err, "rpcL2.SendRawTransaction")
	}

	res.Tx = c.NewTx(hash, defi.CodeTransfer, nil)

	return res, nil
}

func (c *Client) StargateBridgeSwap(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {
	return c.StargateBridge(ctx, req)
}

func (c *Client) Network() v1.Network {
	return c.Cfg.Network
}
