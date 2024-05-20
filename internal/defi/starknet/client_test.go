package starknet

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestSD(t *testing.T) {
	ss := nr("0x317862b9885ec102b45e3126c8a0f1e4dd95124d3d90bd545b2cd74cadc6ee")
	println(ss)
}

func TestClient(t *testing.T) {
	c, err := NewClient(&ClientConfig{
		HttpCli:     tests.GetConfig().Cli,
		RPCEndpoint: MainnetRPC,
		Proxy:       "",
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)

	ctx := context.Background()

	c.WaitTxComplete(ctx, "0x23845f7a03ef2f6a2ac992288666885ca4fbdceaed607e2c1a27d91d28c3af6")

	//r, err := c.Approve(ctx, &ApproveReq{
	//	Token:       v1.Token_ETH,
	//	Amount:      big.NewInt(1000000000),
	//	SpenderAddr: "0x7a6f98c03379b9513ca84cca1373ff452a7462a3b61598f0af5bb27ad7f76d1",
	//	PK:          tests.GetConfig().StarkNetPrivate,
	//})
	//assert.NotNil(t, r)

	//c.GetFundingBalance(ctx, &defi.GetBalanceReq{
	//	WalletAddress: tests.GetConfig().StarkNetPuvlic,
	//	Token:         v1.Token_ETH,
	//})

	//assert.NoError(t, err)
	//assert.NotNil(t, r)
	//res, err := c.Allowed(ctx, &AllowedReq{
	//	Token:       v1.Token_ETH,
	//	WalletAddr:  tests.GetConfig().StarkNetPuvlic,
	//	SpenderAddr: "0x7a6f98c03379b9513ca84cca1373ff452a7462a3b61598f0af5bb27ad7f76d1",
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, res)
}

func TestGen(t *testing.T) {
	pub, err := GetPublicKeyHash(tests.GetConfig().StarkNetPrivate)
	assert.NoError(t, err)
	assert.Equal(t, pub, tests.GetConfig().StarkNetPuvlic)
}

//println(err)
//println(res == nil)

//err = c.DeployAccount(ctx, &Account{
//	PublicKey:  tests.GetConfig().StarkNetPuvlic,
//	PrivateKey: tests.GetConfig().StarkNetPrivate,
//})
//assert.NoError(t, err)
//}
