package zksyncera

import (
	"context"
	"testing"

	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	r, err := NewMainNetClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	ctx := context.Background()

	m := &merkly.Maker{
		TokenMap: r.Cfg.TokenMap,
		Cli:      r.Provider.GetClient(),
	}

	m.MakeMintTx(ctx)

}
