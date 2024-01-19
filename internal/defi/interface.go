package defi

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Token = v1.Token

type Network = v1.Network

type Spec struct {
	Addr     common.Address
	TaskType v1.TaskType
}

var DeadLine = func() *big.Int {
	return new(big.Int).SetInt64(time.Now().Add(time.Second * 20).Unix())
}

var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
var NativeTokenAddress = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")

type Networker interface {
	TxViewFn(id string) string
	GetNetworkId() *big.Int
	WaitTxComplete(ctx context.Context, tx string) error
	Balancer
	GetPublicKey(pk string, subType v1.ProfileSubType) (string, error)
	Network() v1.Network
}

type Balancer interface {
	GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error)
	GetNetworkToken() Token
}

type TestNetworkBridgeSwapper interface {
	Networker
	TestNetBridgeSwap(ctx context.Context, req *TestNetBridgeSwapReq) (*TestNetBridgeSwapRes, error)
}

type OrbiterSwapper interface {
	Transfer
	OrbiterBridge(ctx context.Context, req *OrbiterBridgeReq) (*OrbiterBridgeRes, error)
}

type TxCode = string

const (
	CodeApprove  TxCode = "approve"
	CodeContract TxCode = "contract"
	CodeTransfer TxCode = "transfer"
)

func (c *EtheriumClient) NewTx(id common.Hash, code TxCode, details []bozdo.TxDetail) *bozdo.Transaction {
	return &bozdo.Transaction{
		Code:    code,
		Network: c.Cfg.Network,
		Id:      id.String(),
		Url:     c.TxViewFn(id.String()),
		Details: details,
	}
}

type ZkSyncOfficialWithdrawalEtherium interface {
	Networker
}

type Transfer interface {
	Networker
	Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error)
}

type LiquidityBridger interface {
	Networker
	LiquidityBridge(ctx context.Context, req *LiquidityBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error)
}

type LiquidityBridgeReq struct {
	From         v1.Network
	To           v1.Network
	Percent      string
	Token        v1.Token
	PkFrom       string
	PkTo         string
	Proxy        string
	EstimateOnly bool
	FromSubType  v1.ProfileSubType
	ToSubType    v1.ProfileSubType
}

type DmailSender interface {
	SendDmailMessage(ctx context.Context, req *SimpleReq) (*bozdo.DefaultRes, error)
	Networker
}

type SimpleReq struct {
	PK      string
	SubType v1.ProfileSubType
	*bozdo.BaseReq
}

func CalcRate(from, to v1.Token, amIn, amOut *big.Int) *float64 {
	fam := WeiToToken(amIn, from)
	tam := WeiToToken(amOut, to)

	ratio, _ := big.NewFloat(0).Quo(tam, fam).Float64()

	if ratio < 1 {
		ratio = 1 / ratio
	}
	return &ratio
}
