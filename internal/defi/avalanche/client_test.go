package avalanche

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{
		HttpCli:     tests.GetConfig().Cli,
		RPCEndpoint: MainNetURL,
	})

	h, err := r.GetStargateBridgeFee(context.Background(), &defi.GetStargateBridgeFeeReq{
		ToChain: v1.Network_OPTIMISM,
		Wallet:  tests.GetConfig().Wallet,
	})
	assert.NoError(t, err)
	assert.NotNil(t, h)

	println(defi.WeiToToken(h.Fee1, v1.Token_AVAX).String())

	//assert.NoError(t, err)
	//assert.NotNil(t, r)
	//
	//t.Run("balance", func(t *testing.T) {
	//	b, err := r.GetBalance(context.Background(), &defi.GetBalanceReq{
	//		WalletAddress: tests.GetConfig().Wallet,
	//		Token:         v1.Token_MATIC,
	//	})
	//	assert.NoError(t, err)
	//	assert.NotNil(t, b)
	//})

}
