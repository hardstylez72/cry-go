package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/contracts/optimism_fee"
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/routereth"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type StargateBridgeSwapEthReq struct {
	DestChain    v1.Network
	Wallet       *WalletTransactor
	Quantity     *big.Int
	Gas          *Gas
	EstimateOnly bool
	Slippage     SlippagePercent
}

func (r *StargateBridgeSwapEthReq) Validate(srcChain v1.Network) error {
	if srcChain == r.DestChain {
		return errors.New("same chain: " + string(srcChain))
	}

	if r.Wallet == nil {
		return errors.New("wallet is not set")
	}

	if r.Quantity == nil || r.Quantity.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("quantity is not set")
	}
	return nil
}

type StargateBridgeSwapEthRes struct {
	Tx    *types.Transaction
	ECost *EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapEth(ctx context.Context, req *StargateBridgeSwapEthReq) (*StargateBridgeSwapEthRes, error) {

	//var gasLimitPrice = map[v1.Network]uint64{
	//	v1.Network_ARBITRUM:    4_000_000,
	//	v1.Network_OPTIMISM:    1_000_000,
	//	v1.Network_Etherium:    630_000,
	//	v1.Network_POLIGON:     630_000,
	//	v1.Network_BinanaceBNB: 600_000,
	//	v1.Network_AVALANCHE:   700_000,
	//}

	if err := req.Validate(c.Cfg.Network); err != nil {
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

	destChainId := ChainIdMap[req.DestChain]

	l1Gasfee := big.NewInt(0)
	if c.Cfg.Network == v1.Network_OPTIMISM {
		optFeeCaller, err := optimism_fee.NewStorageCaller(common.HexToAddress("0x420000000000000000000000000000000000000F"), c.Cli)
		if err != nil {
			return nil, err
		}

		abi, err := routereth.StorageMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		data, err := abi.Pack("swapETH",
			destChainId,
			req.Wallet.WalletAddr,
			req.Wallet.WalletAddr.Bytes(),
			req.Quantity,
			Slippage(req.Quantity, SlippagePercent05))
		if err != nil {
			return nil, err
		}
		o := &bind.CallOpts{}
		o.Context = ctx
		l1Gasfee, err = optFeeCaller.GetL1Fee(o, data)
		if err != nil {
			return nil, err
		}
	}

	opt.NoSend = req.EstimateOnly

	opt = c.ResoleGas(ctx, req.Gas, opt)

	opt.Value = BigIntSum(req.Quantity, fee.Fee1, Percent(fee.Fee1, 2), l1Gasfee)

	opt.NoSend = req.EstimateOnly

	tr, err := routereth.NewStorageTransactor(c.Cfg.Dict.Stargate.StargateRouterEthAddress, c.Cli)
	if err != nil {
		return nil, err
	}
	tx, err := tr.SwapETH(
		opt,
		destChainId,
		req.Wallet.WalletAddr,
		req.Wallet.WalletAddr.Bytes(),
		req.Quantity,
		Slippage(req.Quantity, req.Slippage),
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SwapETH")
	}

	return &StargateBridgeSwapEthRes{
		Tx:    tx,
		ECost: Estimate(tx, l1Gasfee),
	}, nil
}

func GasPrice(tx *types.Transaction) *big.Int {
	if tx.GasFeeCap() != nil || tx.GasTipCap().Cmp(big.NewInt(0)) > 0 {
		return tx.GasFeeCap()
	}
	return tx.GasPrice()
}

func BigIntSum(values ...*big.Int) *big.Int {

	result := big.NewInt(0)

	for _, v := range values {
		if v == nil {
			continue
		}
		result = new(big.Int).Add(v, result)
	}

	return result

}

type SlippagePercent = string

const (
	SlippagePercent5    SlippagePercent = "5"
	SlippagePercent2    SlippagePercent = "2"
	SlippagePercent1    SlippagePercent = "1"
	SlippagePercent05   SlippagePercent = "0.5"
	SlippagePercent02   SlippagePercent = "0.2"
	SlippagePercent03   SlippagePercent = "0.3"
	SlippagePercent01   SlippagePercent = "0.1"
	SlippagePercent001  SlippagePercent = "0.01"
	SlippagePercentZero SlippagePercent = "0"
)

func Slippage(v *big.Int, slippagePercent SlippagePercent) *big.Int {
	switch slippagePercent {
	case SlippagePercentZero:
		return v
	case SlippagePercent05:
		prec := big.NewInt(0).Div(v, big.NewInt(1000))
		return big.NewInt(0).Mul(prec, big.NewInt(995))
	case SlippagePercent01:
		prec := big.NewInt(0).Div(v, big.NewInt(1000))
		return big.NewInt(0).Mul(prec, big.NewInt(999))
	case SlippagePercent001:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9999))
	case SlippagePercent02:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9998))
	case SlippagePercent03:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9997))
	case SlippagePercent1:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(99))
	case SlippagePercent2:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(98))
	case SlippagePercent5:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(95))
	}
	return nil
}
