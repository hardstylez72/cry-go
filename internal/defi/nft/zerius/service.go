package zerius

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Client struct {
	cli      *defi.EtheriumClient
	tokenMap map[v1.Token]common.Address
}

func NewClient() {
	
}
