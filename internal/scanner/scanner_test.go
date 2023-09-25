package scanner

import (
	"context"
	"net/http"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestNewScannerZkSync(t *testing.T) {

	scanner := NewScannerZkSync(&http.Client{})
	stat, err := Stat(context.Background(), scanner, tests.GetConfig().Wallet.String())
	assert.NoError(t, err)
	assert.NotNil(t, stat)

}
