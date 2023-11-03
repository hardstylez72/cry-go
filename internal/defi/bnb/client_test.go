package bnb

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	nft, err := r.GetNFTId(context.Background(), common.HexToHash("0xe43fa82a0121f893b74de3328f78e9a7eebf90f6ddcfb10bdf8a5434ab45bfb3"))
	assert.NoError(t, err)
	assert.NotNil(t, nft)

	println(nft.String())

	res, err := r.BridgeNftMerkly(context.Background(), &merkly.BridgeNFTReq{
		WalletPK:    tests.GetConfig().PK,
		FromNetwork: v1.Network_BinanaceBNB,
		ToNetwork:   v1.Network_ZKSYNCERA,
		NFTId:       nft,
		BaseReq: bozdo.BaseReq{
			EstimateOnly: false,
			Gas:          nil,
			Debug:        false,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
	println(res.Tx.Url)

	h, err := r.defi.Cli.HeaderByNumber(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, h)

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
