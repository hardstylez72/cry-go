package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/traderjoe"
	"github.com/pkg/errors"
)

func (c *EtheriumClient) Swapper(ctx context.Context, req *DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_TraderJoeSwap:
		return c.TraderJoeSwap(ctx, req)
	default:
		return nil, errors.New("unsupported task type: " + taskType.String())
	}
}

func (c *EtheriumClient) TraderJoeSwap(ctx context.Context, req *DefaultSwapReq) (*bozdo.DefaultRes, error) {

	result := &bozdo.DefaultRes{}

	tr, err := NewWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, err
	}
	res, err := c.traderJoeService.GetSwapData(ctx, &traderjoe.GetSwapDataReq{
		FromToken: req.FromToken,
		ToToken:   req.ToToken,
		ChainRPC:  c.Cfg.MainNet,
		Amount:    req.Amount,
		Recipient: tr.WalletAddr,
	})
	if err != nil {
		return nil, err
	}

	limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		Wallet:      tr,
		Amount:      req.Amount,
		SpenderAddr: common.HexToAddress(res.ContractAddr),
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	if limitTx.LimitExtended {
		result.ApproveTx = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove, nil)
	}

	to := common.HexToAddress(res.ContractAddr)

	gasPrice, err := c.Cli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	value, err := hexutil.DecodeBig(res.Value)
	if err != nil {
		return nil, err
	}

	data := common.Hex2Bytes(res.Data[2:])
	println(common.Bytes2Hex(data))

	estimate, err := c.Cli.EstimateGas(ctx, ethereum.CallMsg{
		From:      tr.WalletAddr,
		To:        &to,
		Gas:       0,
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		Value:     value,
		Data:      data,
	})
	if err != nil {
		return nil, err
	}

	e := &bozdo.EstimatedGasCost{
		GasLimit:    new(big.Int).SetUint64(estimate),
		GasPrice:    gasPrice,
		TotalGasWei: MinerGasLegacy(gasPrice, estimate),
	}

	result.ECost = e

	if req.EstimateOnly {
		return result, nil
	}

	var tx *types.Transaction

	nonce, err := c.Cli.NonceAt(ctx, tr.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	head, errHead := c.Cli.HeaderByNumber(ctx, nil)
	if errHead == nil && head.BaseFee != nil {
		txx := &types.DynamicFeeTx{
			ChainID:   c.Cfg.NetworkId,
			Nonce:     nonce,
			GasTipCap: nil,
			GasFeeCap: head.BaseFee,
			Gas:       estimate,
			To:        &to,
			Value:     value,
			Data:      data,
		}

		tx = types.NewTx(txx)
	} else {
		txx := &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      estimate,
			To:       &to,
			Value:    value,
			Data:     data,
		}
		tx = types.NewTx(txx)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(tr.PrivateKey, c.Cfg.NetworkId)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}

	tx, err = opt.Signer(tr.WalletAddr, tx)
	if err != nil {
		return nil, err
	}

	if err := c.Cli.SendTransaction(ctx, tx); err != nil {
		return nil, err
	}

	result.Tx = c.NewTx(tx.Hash(), CodeContract, nil)

	return result, nil
}
