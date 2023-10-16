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
	"github.com/hardstylez72/cry/internal/defi/contracts/stargate/routereth"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type StargateBridgeSwapEthReq struct {
	DestChain    v1.Network
	Wallet       *WalletTransactor
	Quantity     *big.Int
	Gas          *bozdo.Gas
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
	ECost *bozdo.EstimatedGasCost
}

func (c *EtheriumClient) StargateBridgeSwapEth(ctx context.Context, req *StargateBridgeSwapEthReq) (*StargateBridgeSwapEthRes, error) {
	details := []bozdo.TxDetail{}
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

	fee.Fee1 = bozdo.BigIntSum(fee.Fee1, bozdo.Percent(fee.Fee1, layerzero.LayerZeroBoostPercent))
	destChainId := layerzero.LayerZeroChainMap[req.DestChain]

	l1Gasfee := big.NewInt(0)
	if c.Cfg.Network == v1.Network_OPTIMISM {
		amSlip, err := Slippage(req.Quantity, req.Slippage)
		if err != nil {
			return nil, err
		}

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
			amSlip)
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

	opt.NoSend = req.EstimateOnly

	opt = c.ResoleGas(ctx, req.Gas, opt)

	opt.Value = bozdo.BigIntSum(req.Quantity, fee.Fee1, l1Gasfee)

	opt.NoSend = req.EstimateOnly
	amSlip, err := Slippage(req.Quantity, req.Slippage)
	if err != nil {
		return nil, err
	}
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
		amSlip,
	)
	if err != nil {
		return nil, errors.Wrap(err, "tr.SwapETH")
	}

	return &StargateBridgeSwapEthRes{
		Tx:    tx,
		ECost: Estimate(tx, bozdo.BigIntSum(l1Gasfee, fee.Fee1), "bridge", details),
	}, nil
}

func GasPrice(tx *types.Transaction) *big.Int {
	if tx.GasFeeCap() != nil || tx.GasTipCap().Cmp(big.NewInt(0)) > 0 {
		return tx.GasFeeCap()
	}
	return tx.GasPrice()
}
