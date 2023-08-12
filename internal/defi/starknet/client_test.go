package starknet

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c, err := NewClient(&ClientConfig{
		HttpCli:     tests.GetConfig().Cli,
		RPCEndpoint: MainnetRPC,
		Proxy:       "",
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)

	ctx := context.Background()
	res, err := c.GetBalance(ctx, &defi.GetBalanceReq{
		WalletAddress: tests.GetConfig().StarkNetPuvlic,
		Token:         v1.Token_ETH,
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGen(t *testing.T) {
	pub, err := GetPublicKeyV2(tests.GetConfig().StarkNetPrivate)
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
