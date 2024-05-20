package linea

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/_bridge/across"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {
	return c.defi.GetBalance(ctx, req)
}

func (c *Client) TxViewFn(id string) string {
	return c.defi.TxViewFn(id)
}

func (c *Client) GetNetworkToken() defi.Token {
	return c.defi.GetNetworkToken()
}

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	return c.defi.Transfer(ctx, r)
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) WaitTxComplete(ctx context.Context, s string) error {

	txId := common.HexToHash(s)

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {

		tx, err := c.defi.Cli.TransactionReceipt(ctx, txId)
		if err == nil {

			if tx.Status == types.ReceiptStatusFailed {
				return errors.Wrap(defi.ErrTxStatusFailed, "invalid status")
			}
			if tx.Status == types.ReceiptStatusSuccessful {
				return nil
			}
		}

		if err != nil {
			if err.Error() == "not found" {
				return errors.Wrap(defi.ErrTxNotFound, "blockchain error")
			}
		}

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:

			return nil
		}
	}
}

func (c *Client) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	return c.defi.GetPublicKey(pk, subType)
}

func (c *Client) Network() v1.Network {
	return c.defi.Cfg.Network
}

func (c *Client) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {
	switch taskType {
	case v1.TaskType_AcrossBridge:
		b := across.NewAcrossBridge(c.defi)
		return b.WaitForConfirm(ctx, txId, receiver)
	default:
		return errors.New("bridge unsupported")
	}
}

func (c *Client) OrbiterBridge(ctx context.Context, req *defi.OrbiterBridgeReq) (*defi.OrbiterBridgeRes, error) {
	return c.defi.OrbiterBridge(ctx, req)
}

func (c *Client) GetNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error) {
	return c.defi.GetNFTId(ctx, txHash, owner)
}

type Fee struct {
	MaxFeeCap      *big.Int
	MaxPriorityFee *big.Int
}

// https://docs.linea.build/build-on-linea/gas-fees
func GetFee(ctx context.Context, c *defi.EtheriumClient) (*Fee, error) {

	history, err := c.Cli.FeeHistory(ctx, 5, nil, []float64{20})
	if err != nil {
		return nil, err
	}

	averagePriorityFee := big.NewInt(0)
	count := int64(0)
	for _, rewardos := range history.Reward {
		for _, rewardo := range rewardos {
			count++
			averagePriorityFee = bozdo.BigIntSum(averagePriorityFee, rewardo)
		}
	}

	averagePriorityFee = new(big.Int).Div(averagePriorityFee, big.NewInt(count))
	baseFee := history.BaseFee[len(history.BaseFee)-1]
	MaxFeeCap := bozdo.BigIntSum(baseFee, baseFee, averagePriorityFee)

	return &Fee{
		MaxFeeCap:      MaxFeeCap,
		MaxPriorityFee: averagePriorityFee,
	}, nil
}
func LondonReadyTx(ctx context.Context, c *defi.EtheriumClient, opt *defi.TxOpt, data *bozdo.TxData) (*bozdo.DefaultRes, error) {

	wt, err := defi.NewWalletTransactor(opt.Pk)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Cli.NonceAt(ctx, wt.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	dynamic := types.DynamicFeeTx{
		ChainID: c.Cfg.NetworkId,
		Nonce:   nonce,
		To:      &data.ContractAddr,
		Value:   data.Value,
		Data:    data.Data,
	}

	dynamic.Value = data.Value

	if opt.NoSend || !opt.Gas.RuleSet() {

		fee, err := GetFee(ctx, c)
		if err != nil {
			return nil, err
		}

		dynamic.GasTipCap = fee.MaxPriorityFee
		dynamic.GasFeeCap = fee.MaxFeeCap

		gas, err := c.Cli.EstimateGas(ctx, defi.TxToCallMsg(wt.WalletAddr, types.NewTx(&dynamic)))
		if err != nil {
			return nil, err
		}
		dynamic.Gas = gas
	}

	if opt.Gas.RuleSet() {
		dynamic.Gas = opt.Gas.GasLimit.Uint64()
		dynamic.GasFeeCap = &opt.Gas.GasPrice
		dynamic.GasTipCap = big.NewInt(0).Sub(&opt.Gas.GasPrice, big.NewInt(10)) //todo: может быть придумать что поумнее
	}

	r := &bozdo.DefaultRes{
		ECost: defi.Estimate(types.NewTx(&dynamic), bozdo.BigIntSum(data.ExtraFee), "", nil),
	}

	r.ECost.Details = data.Details

	if data.Rate != nil && opt.ExchangeRate != nil {
		r.ECost.Details = append(r.ECost.Details, bozdo.NewSwapRateRatio(*opt.ExchangeRate, *data.Rate))
	}

	if !opt.NoSend {

		tx := types.NewTx(&dynamic)

		if opt.Debug {
			log.Log.Debug("swap tx gas "+opt.TaskType.String(), zap.String("gas", defi.TxGas(tx)))
		}

		signer := types.NewLondonSigner(c.Cfg.NetworkId)

		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), wt.PrivateKey)
		if err != nil {
			return nil, err
		}

		tx, err = tx.WithSignature(signer, signature)
		if err != nil {
			return nil, err
		}

		if err := c.Cli.SendTransaction(ctx, tx); err != nil {
			return nil, err
		}

		r.Tx = c.NewTx(tx.Hash(), data.Code, r.ECost.Details)
	}

	return r, nil
}
