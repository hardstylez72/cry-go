package bnb

import (
	"context"
	"math/big"
	"testing"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	req := &defi.DefaultBridgeReq{
		FromNetwork:  v1.Network_BinanaceBNB,
		ToNetwork:    v1.Network_Core,
		PK:           tests.GetConfig().PK,
		Amount:       big.NewInt(994000590000000000),
		FromToken:    v1.Token_USDT,
		ToToken:      v1.Token_USDT,
		Gas:          nil,
		Slippage:     "0.1",
		EstimateOnly: true,
		Debug:        true,
		SubType:      v1.ProfileSubType_Metamask,
	}

	nft, err := r.Bridge(context.Background(), req, v1.TaskType_CoreDaoBridge)
	assert.NoError(t, err)
	assert.NotNil(t, nft)

	//tx, penging, err := r.defi.CliL1.TransactionByHash(context.Background(), common.HexToHash("0x02433465f2e9a56484dea517b4ab932ca1d5970451ea8cef90499713eb447021"))
	//assert.NoError(t, err)
	//assert.NotNil(t, penging)
	//assert.NotNil(t, tx)
	//rec, err := bind.WaitMined(context.Background(), r.defi.CliL1, tx)
	//assert.NotNil(t, rec)
	//assert.NotNil(t, err)
	//t.Run("balance", func(t *testing.T) {
	//	b, err := r.GetFundingBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         v1.Token_USDT,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, b)
	//})
}
