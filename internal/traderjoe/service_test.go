package traderjoe

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	s := NewService(&Config{
		Host: "http://localhost:3244",
	})

	am, _ := big.NewInt(0).SetString("23948228429657588650", 10)

	_, err := s.GetSwapData(context.Background(), &GetSwapDataReq{
		FromToken: v1.Token_STG,
		ToToken:   v1.Token_ETH,
		ChainRPC:  "https://arb1.arbitrum.io/rpc",
		Amount:    am,
		Recipient: common.HexToAddress("0x4A6e7c137a6691D55693CA3Bc7E5C698d9d43815"),
	})
	assert.NoError(t, err)
}
