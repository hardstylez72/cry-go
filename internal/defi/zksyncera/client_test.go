package zksyncera

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	r, err := NewMainNetClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	r.syncSwapPoolRates(context.Background(), common.HexToAddress("0x80115c708E12eDd42E504c1cD52Aea96C547c05c"))
	//d := AcrossBridge{
	//	cli:     r,
	//	HttpCli: &http.Client{},
	//	Debug:   true,
	//}
	//
	//
	//d.WaitTxComplete(context.Background(), "0x34cd2eddfc129d7040fa87f292e82cd0433f37b71b4204ccc3e2342925059ec0")

	//d.MakeBridgeTx(context.Background(), &defi.DefaultBridgeReq{
	//	FromNetwork:  v1.Network_ZKSYNCERA,
	//	ToNetwork:    v1.Network_ARBITRUM,
	//	PK:           tests.GetConfig().PK,
	//	Amount:       big.NewInt(100_000_0),
	//	FromToken:    v1.Token_USDC,
	//	ToToken:      v1.Token_USDC,
	//	Gas:          nil,
	//	Slippage:     "0.01",
	//	EstimateOnly: true,
	//	Debug:        false,
	//})

	//r.WaitTxComplete(context.Background(), "0xf0ce2ba93e5d068cd47a64b79fbc04df42876d9c29a2b7b5cb909d4f7efb148e")

	//ctx := context.Background()
	//
	//res, err := r.SendDmailMessage(ctx, &SendDmailMessageReq{
	//	PK: tests.GetConfig().PK,
	//	BaseReq: &bozdo.BaseReq{
	//		EstimateOnly: true,
	//		Gas:          nil,
	//		Debug:        true,
	//	},
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, res)

	//tx, err := r.rpcL2.GetTransaction(txHash)
	//
	//println(tx.To.String())
	//println(tx.From.String())
	////
	//l, err := r.rpcL2.GetLogs(zkTypes.FilterQuery{
	//	BlockHash: &tx.BlockHash,
	//	FromBlock: nil,
	//	ToBlock:   nil,
	//	Addresses: []common.Address{tx.To},
	//	Topics:    nil,
	//})
	//println(l == nil)
	//println(l[0].Topics[0].Big().String())
	//println(l[0].Address.String())
	//println(l[0].Topics[1].Big().String())
	//println(l[0].Topics[2].Big().String())
	//println(l[0].Topics[3].Big().String())
	//m := &merkly.Maker{
	//	TokenMap: r.Cfg.TokenMap,
	//	CliL1:      r.rpcL2.GetClient(),
	//}
	//
	//m.MakeMintTx(ctx)

}
