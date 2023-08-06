package starknet

import (
	"context"
	"testing"

	"github.com/dontpanicdao/caigo/types"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c, err := NewClient(&ClientConfig{
		HttpCli:     tests.GetConfig().Cli,
		RPCEndpoint: MainnetRPC,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)

	ctx := context.Background()
	types.SNValToBN("0xbf4629c0934bfcfb8e7e3d5de9b3a5b4f93c1f4cc65d7f69b0b1343b47239d").Text(10)
	err = c.DeployAccount(ctx, &Account{
		PublicKey:  tests.GetConfig().StarkNetPuvlic,
		PrivateKey: tests.GetConfig().StarkNetPrivate,
	})
	assert.NoError(t, err)
}
