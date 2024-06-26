package etherium

import (
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	r.defi.Cli.FilterLogs()
	//
	//t.Run("balance", func(t *testing.T) {
	//	b, err := r.GetFundingBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         defi.TokenETH,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, b)
	//})
	//
	//t.Run("stg stake", func(t *testing.T) {
	//
	//	wallet, err := defi.NewWalletTransactor(tests.GetConfig().PK)
	//	assert.NoError(t, err)
	//
	//	am := defi.TokenAmountToWEI(6)
	//
	//	res, err := r.StakeSTG(context.Background(), &defi.StakeSTGReq{
	//		Amount: am,
	//		Wallet: wallet,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, res)
	//})
}
