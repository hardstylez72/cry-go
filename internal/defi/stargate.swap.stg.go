package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/layerzero"
	"github.com/hardstylez72/cry/internal/defi/contracts/optimism_fee"
	"github.com/hardstylez72/cry/internal/defi/contracts/stg"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type StargateBridgeSwapSTGReq struct {
	DestChain    v1.Network
	Wallet       *WalletTransactor
	Quantity     *big.Int
	Gas          *bozdo.Gas
	EstimateOnly bool
}

func (r *StargateBridgeSwapSTGReq) Validate() error {
	if r == nil {
		return errors.New("empty request")
	}

	if r.Wallet == nil {
		return errors.New("empty wallet")
	}

	if r.Quantity.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("zero quantity")
	}

	if r.Quantity == nil {
		return errors.New("quantity or fee empty")
	}

	return nil
}

type StargateBridgeSwapSTGRes struct {
	Tx    *types.Transaction
	ECost *bozdo.EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapSTG(ctx context.Context, req *StargateBridgeSwapSTGReq) (*StargateBridgeSwapSTGRes, error) {
	details := []bozdo.TxDetail{}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	tr, err := stg.NewStgTransactor(c.Cfg.TokenMap[v1.Token_STG], c.Cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	fee, err := c.GetStargateBridgeFee(ctx, &GetStargateBridgeFeeReq{
		ToChain: req.DestChain,
		Wallet:  req.Wallet.WalletAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetStargateBridgeFee")
	}

	fee.Fee1 = bozdo.BigIntSum(bozdo.Percent(fee.Fee1, layerzero.LayerZeroBoostPercent), fee.Fee1)

	destChainId := layerzero.LayerZeroChainMap[req.DestChain]

	l1Gasfee := big.NewInt(0)
	if c.Cfg.Network == v1.Network_OPTIMISM {
		optFeeCaller, err := optimism_fee.NewStorageCaller(common.HexToAddress("0x420000000000000000000000000000000000000F"), c.Cli)
		if err != nil {
			return nil, err
		}

		abi, err := stg.StgMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		data, err := abi.Pack("sendTokens",
			destChainId,
			req.Wallet.WalletAddr.Bytes(),
			req.Quantity,
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
			[]byte{},
		)
		if err != nil {
			return nil, err
		}
		o := &bind.CallOpts{}
		o.Context = ctx
		l1Gasfee, err = optFeeCaller.GetL1Fee(o, data)
		if err != nil {
			return nil, err
		}
		details = append(details, bozdo.NewL1FeeDetails(l1Gasfee, c.Cfg.Network, v1.Token_ETH))
	}
	details = append(details, bozdo.NewLZFeeDetails(fee.Fee1, c.Cfg.Network, v1.Token_ETH))

	opt.Value = bozdo.BigIntSum(fee.Fee1, l1Gasfee)
	opt.NoSend = req.EstimateOnly

	opt = c.ResoleGas(ctx, req.Gas, opt)

	tx, err := tr.SendTokens(
		opt,
		destChainId,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		[]byte{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SendTokens")
	}

	return &StargateBridgeSwapSTGRes{
		Tx:    tx,
		ECost: Estimate(tx, nil, "bridge", details),
	}, nil
}

func Estimate(tx *types.Transaction, extraFee *big.Int, name string, details []bozdo.TxDetail) *bozdo.EstimatedGasCost {
	gasLimit := new(big.Int).SetUint64(tx.Gas())

	if tx.Type() == types.DynamicFeeTxType {
		fee := bozdo.BigIntSum(tx.GasFeeCap())

		return &bozdo.EstimatedGasCost{
			Name:        "bridge",
			GasLimit:    gasLimit,
			GasPrice:    fee,
			TotalGasWei: new(big.Int).Mul(bozdo.BigIntSum(fee), gasLimit),
			Details:     details,
			ExtraFee:    extraFee,
		}
	}

	return &bozdo.EstimatedGasCost{
		Type:        bozdo.TxTypeLegacy,
		Name:        name,
		GasLimit:    gasLimit,
		GasPrice:    tx.GasPrice(),
		TotalGasWei: new(big.Int).Mul(gasLimit, tx.GasPrice()),
		Details:     details,
		ExtraFee:    extraFee,
	}
}
