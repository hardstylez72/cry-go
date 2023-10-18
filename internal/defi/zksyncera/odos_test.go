package zksyncera

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/odos"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestOdos(t *testing.T) {

	cli, err := NewMainNetClient(&ClientConfig{
		HttpCli:     tests.GetConfig().Cli,
		RPCEndpoint: MainNetURL,
	})
	assert.NoError(t, err)

	m := odos.OdosMaker{
		CA:       common.HexToAddress("0x4bba932e9792a2b917d47830c93a9bc79320e4f7"),
		TokenMap: cli.Cfg.TokenMap,
		CliHttp:  &http.Client{},
		ChainId:  cli.NetworkId,
		Addr:     tests.GetConfig().Wallet,
	}

	req := &defi.DefaultSwapReq{
		Network:      v1.Network_ZKSYNCERA,
		Amount:       big.NewInt(1e6),
		FromToken:    v1.Token_USDT,
		ToToken:      v1.Token_USDC,
		WalletPK:     tests.GetConfig().PK,
		EstimateOnly: true,
		Gas:          nil,
		Slippage:     "0.1",
		Debug:        false,
		SubType:      v1.ProfileSubType_Metamask,
		ExchangeRate: nil,
	}
	res, err := m.Quote(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	res2, err := m.Assemble(context.Background(), req, res)
	assert.NoError(t, err)
	assert.NotNil(t, res2)
}
