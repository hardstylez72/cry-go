package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/startgatemavrouter"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *EtheriumClient) StargateBridgeSwapMAV(ctx context.Context, req *DefaultBridgeReq) (*DefaultRes, error) {

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
		result.ApproveTx = c.NewTx(limitTx.ApproveTx.Hash(), CodeApprove)
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

	opt.Value = BigIntSum(fee.Fee1, Percent(fee.Fee1, 2))
	opt.NoSend = req.EstimateOnly
	opt = c.ResoleGas(ctx, req.Gas, opt)

	destChainId := ChainIdMap[req.ToNetwork]

	adapterParams, err := MakeLayerZeroAdapterParams(1, big.NewInt(150000))
	if err != nil {
		return nil, err
	}

	tx, err := router.SendOFT(
		opt,
		mavAddr,
		destChainId,
		wt.WalletAddr.Bytes(),
		req.Amount,
		Slippage(req.Amount, SlippagePercent05),
		wt.WalletAddr,
		ZEROADDR,
		adapterParams,
		startgatemavrouter.IOFTWrapperFeeObj{
			CallerBps: big.NewInt(0),
			Caller:    ZEROADDR,
			PartnerId: [2]byte{0, 0},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "router.SendOFT")
	}

	result.ECost = Estimate(tx, nil)
	result.Tx = c.NewTx(tx.Hash(), CodeContract)

	return result, nil
}

func MakeLayerZeroAdapterParams(a uint16, b *big.Int) ([]byte, error) {

	structThing, err := abi.NewType("tuple", "lzAdapterParams", []abi.ArgumentMarshaling{
		{Name: "a", Type: "uint16"},
		{Name: "b", Type: "uint256"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "abi.NewType")
	}

	record := struct {
		A uint16
		B *big.Int
	}{
		A: a,
		B: b,
	}

	args := abi.Arguments{
		{Type: structThing, Name: "lzAdapterParams"},
	}
	adapterParams, err := args.Pack(&record)
	if err != nil {
		return nil, errors.Wrap(err, "args.Pack(&record)")
	}

	// 34 байтов
	// https://github.com/LayerZero-Labs/solidity-examples/blob/main/contracts/lzApp/LzApp.sol#L64
	adapterParams = adapterParams[30:]

	return adapterParams, nil
}
