package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/startgatemavrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *EtheriumClient) StargateBridgeSwapMAV(ctx context.Context, req *DefaultBridgeReq) (*DefaultRes, error) {
	details := []bozdo.TxDetail{}
	result := &DefaultRes{}

	// https://bscscan.com/address/0x86355F02119bdBC28ED6A4D5E0cA327Ca7730fFF
	constractAddr := common.HexToAddress("0x86355F02119bdBC28ED6A4D5E0cA327Ca7730fFF")

	mavAddr, ok := c.Cfg.TokenMap[v1.Token_MAV]
	if !ok {
		return nil, ErrTokenNotSupportedFn(req.FromToken)
	}

	router, err := startgatemavrouter.NewStorageTransactor(constractAddr, c.Cli)
	if err != nil {
		return nil, err
	}

	wt, err := NewWalletTransactor(req.WalletPK)
	if err != nil {
		return nil, err
	}

	limitTx, err := c.TokenLimitChecker(ctx, &TokenLimitCheckerReq{
		Token:       req.FromToken,
		Wallet:      wt,
		Amount:      req.Amount,
		SpenderAddr: constractAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "TokenLimitChecker")
	}
	if limitTx.LimitExtended {
		result.ApproveTx = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove, nil)
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(wt.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.ToNetwork,
		Wallet:  wt.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	details = append(details, bozdo.NewLZFeeDetails(fee.Fee1, c.Cfg.Network, v1.Token_ETH))

	fee.Fee1 = bozdo.BigIntSum(fee.Fee1, bozdo.Percent(fee.Fee1, 2))
	opt.Value = fee.Fee1
	opt.NoSend = req.EstimateOnly
	opt = c.ResoleGas(ctx, req.Gas, opt)

	destChainId := layerzero.LayerZeroChainMap[req.ToNetwork]

	adapterParams, err := layerzero.MakeLayerZeroAdapterParams(1, big.NewInt(150000))
	if err != nil {
		return nil, err
	}

	amSlip, err := Slippage(req.Amount, req.Slippage)
	if err != nil {
		return nil, err
	}

	tx, err := router.SendOFT(
		opt,
		mavAddr,
		destChainId,
		wt.WalletAddr.Bytes(),
		req.Amount,
		amSlip,
		wt.WalletAddr,
		bozdo.ZEROADDR,
		adapterParams,
		startgatemavrouter.IOFTWrapperFeeObj{
			CallerBps: big.NewInt(0),
			Caller:    bozdo.ZEROADDR,
			PartnerId: [2]byte{0, 0},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "router.SendOFT")
	}

	result.ECost = Estimate(tx, nil, "bridge", details)
	result.Tx = c.NewTx(tx.Hash(), CodeContract, details)

	return result, nil
}
