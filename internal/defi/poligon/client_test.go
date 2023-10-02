package poligon

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	r, err := NewClient(&ClientConfig{RPCEndpoint: MainNetURL, HttpCli: tests.GetConfig().Cli})
	assert.NoError(t, err)
	assert.NotNil(t, r)

	nft, err := r.GetMerklyNFTId(context.Background(), common.HexToHash("0x5bb5e941e171d06d682e7c28eb1e8e38fd906e18371e9beb7935df7eab98d0c3"), common.HexToAddress("0x33c7242f2a272ca19aa3a2ad9e836be35f68f130"))
	assert.NoError(t, err)
	assert.NotNil(t, nft)

	println(nft.String())

}
