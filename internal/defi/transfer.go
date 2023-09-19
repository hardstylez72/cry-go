package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/contracts/erc_20"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type TransferRes struct {
	Tx    *bozdo.Transaction
	ECost *bozdo.EstimatedGasCost
}

type TransferReq struct {
	Pk       string
	ToAddr   string
	Token    Token
	Amount   *big.Int
	PSubType v1.ProfileSubType

	Gas          *bozdo.Gas
	EstimateOnly bool
}

func (c *EtheriumClient) Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error) {

	if err := r.Validate(c.Cfg.TokenMap); err != nil {
		return nil, err
	}

	wallet, err := newWalletTransactor(r.Pk)
	if err != nil {
		return nil, err
	}

	if r.Token == c.Cfg.MainToken {
		return c.TransferMainToken(ctx, &TransferMainTokenReq{
			Wallet:       wallet,
			ToAddr:       r.ToAddr,
			Amount:       r.Amount,
			Gas:          r.Gas,
			EstimateOnly: r.EstimateOnly,
		})
	}

	tokenAddr := c.Cfg.TokenMap[r.Token]

	trx, err := erc_20.NewStorageTransactor(tokenAddr, c.Cli)
	if err != nil {
		return nil, err
	}

	chainID, err := c.Cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	opt.NoSend = r.EstimateOnly
	if r.Gas.RuleSet() {
		opt.GasLimit = r.Gas.GasLimit.Uint64()
		opt.GasPrice = &r.Gas.GasPrice
	}

	toto := common.HexToAddress(r.ToAddr)

	tx, err := trx.Transfer(opt, toto, r.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "Transfer")
	}

	eCost := &bozdo.EstimatedGasCost{
		GasLimit:    big.NewInt(0).SetUint64(tx.Gas()),
		GasPrice:    tx.GasPrice(),
		TotalGasWei: MinerGasLegacy(tx.GasPrice(), tx.Gas()),
	}

	return &TransferRes{
		Tx:    c.NewTx(tx.Hash(), CodeTransfer, nil),
		ECost: eCost,
	}, nil
}

func (r *TransferReq) Validate(tm map[Token]common.Address) error {
	if r == nil {
		return errors.New("empty request")
	}

	if r.Pk == "" {
		return errors.New("empty wallet")
	}

	if r.Amount.CmpAbs(big.NewInt(0)) == 0 {
		return errors.New("Amount is 0")
	}

	_, ok := tm[r.Token]
	if !ok {
		return errors.New("unknown token: " + r.Token.String())
	}

	return nil
}

type TransferMainTokenReq struct {
	Wallet       *WalletTransactor
	ToAddr       string
	Amount       *big.Int
	Gas          *bozdo.Gas
	EstimateOnly bool
}

func (c *EtheriumClient) TransferMainToken(ctx context.Context, r *TransferMainTokenReq) (*TransferRes, error) {

	nonce, err := c.Cli.NonceAt(ctx, r.Wallet.WalletAddr, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.GetBalance(ctx, &GetBalanceReq{
		WalletAddress: r.Wallet.WalletAddr.String(),
		Token:         c.Cfg.MainToken,
	})
	if err != nil {
		return nil, err
	}

	data := []byte{}

	tot := common.HexToAddress(r.ToAddr)

	header, err := c.Cli.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	gasFeeCap := bozdo.BigIntSum(header.BaseFee, header.BaseFee)
	gas, err := c.Cli.EstimateGas(ctx, ethereum.CallMsg{
		From:       r.Wallet.WalletAddr,
		To:         &tot,
		Gas:        0,
		GasPrice:   nil,
		GasFeeCap:  gasFeeCap,
		GasTipCap:  nil,
		Value:      big.NewInt(1),
		Data:       data,
		AccessList: nil,
	})
	if err != nil {
		return nil, err
	}

	estimate := &bozdo.EstimatedGasCost{
		GasLimit:    big.NewInt(0).SetUint64(gas),
		GasPrice:    gasFeeCap,
		TotalGasWei: MinerGasLegacy(gasFeeCap, gas),
	}

	if r.EstimateOnly {
		return &TransferRes{
			Tx:    nil,
			ECost: estimate,
		}, nil
	}

	if r.Gas.RuleSet() {
		gasFeeCap = &r.Gas.GasPrice
		gas = r.Gas.GasLimit.Uint64()
	}

	am := r.Amount
	if am.Cmp(b.WEI) == 0 {
		am = new(big.Int).Sub(b.WEI, MinerGasLegacy(gasFeeCap, gas))
	}

	to := common.HexToAddress(r.ToAddr)

	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &to,
		Value:     am,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(c.Cfg.NetworkId), r.Wallet.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "types.SignTx")
	}

	err = c.Cli.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, errors.Wrap(err, "Cli.SendTransaction")
	}
	return &TransferRes{
		Tx:    c.NewTx(signedTx.Hash(), CodeTransfer, nil),
		ECost: estimate,
	}, nil
}
