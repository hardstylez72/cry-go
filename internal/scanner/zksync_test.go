package scanner

import (
	"context"
	"net/http"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestZkSync_Transactions(t *testing.T) {

	s := NewScannerZkSync(&http.Client{})

	txs, err := s.Transfers(context.Background(), tests.GetConfig().Wallet.String())
	assert.NoError(t, err)
	assert.NotNil(t, txs)
}
