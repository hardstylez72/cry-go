package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	wallet, err := NewWalletTransactor(r.Pk)
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

	chainID := c.Cfg.NetworkId

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

	opt := &TxOpt{
		NoSend:       r.EstimateOnly,
		Pk:           r.Wallet.PK,
		Gas:          r.Gas,
		ExchangeRate: nil,
		Debug:        true,
	}

	d := &bozdo.TxData{
		Code:         bozdo.CodeTransfer,
		Data:         []byte{},
		Value:        r.Amount,
		ContractAddr: common.HexToAddress(r.ToAddr),
		Details:      nil,
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     nil,
	}

	res, err := c.London(ctx, c, opt, d)
	if err != nil {
		return nil, errors.Wrap(err, "LondonReadyTx")
	}
	return &TransferRes{
		Tx:    res.Tx,
		ECost: res.ECost,
	}, nil
}
