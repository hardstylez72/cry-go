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

	ctx := context.Background()

	txHash := common.HexToHash("0x3042e032c90adcc87af3db59f258382e775cfdb060b15025b872ed0f8d6ee9a3")

	r.GetMerklyNFTId(ctx, txHash)
	//tx, err := r.Provider.GetTransaction(txHash)
	//
	//println(tx.To.String())
	//println(tx.From.String())
	////
	//l, err := r.Provider.GetLogs(zkTypes.FilterQuery{
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
	//	Cli:      r.Provider.GetClient(),
	//}
	//
	//m.MakeMintTx(ctx)

}
