package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/scan"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Token = v1.Token

type Network = v1.Network

type Networker interface {
	TxViewFn(id string) string
	GetNetworkId() *big.Int
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	WaitTxComplete(ctx context.Context, tx common.Hash) error
	Balancer
}

type Balancer interface {
	GetBalance(ctx context.Context, req *GetBalanceReq) (*GetBalanceRes, error)
	GetNetworkToken() Token
}

type StargateSwapper interface {
	Networker
	StargateBridgeSwap(ctx context.Context, req *DefaultBridgeReq) (*DefaultRes, error)
}
type TestNetworkBridgeSwapper interface {
	Networker
	TestNetBridgeSwap(ctx context.Context, req *TestNetBridgeSwapReq) (*TestNetBridgeSwapRes, error)
}

type TraderJoeSwap interface {
	Networker
	TraderJoeSwap(ctx context.Context, req *DefaultSwapReq) (*DefaultRes, error)
}

type OrbiterSwapper interface {
	Transfer
	OrbiterBridge(ctx context.Context, req *OrbiterBridgeReq) (*OrbiterBridgeRes, error)
}

type Gas struct {
	Network       v1.Network
	GasMultiplier float64
	MaxGas        big.Int

	GasBeforeMultiplier big.Int
	GasLimit            big.Int
	GasPrice            big.Int
	TotalGas            big.Int
}

type TxData struct {
	Data         []byte
	Value        *big.Int
	ContractAddr common.Address
}

func (g *Gas) RuleSet() bool {
	return g != nil
}

type SyncSwapReq struct {
	Network      v1.Network
	Amount       *big.Int
	FromToken    v1.Token
	ToToken      v1.Token
	WalletPK     string
	EstimateOnly bool
	Gas          *Gas
}

type DefaultSwapReq struct {
	Network      v1.Network
	Amount       *big.Int
	FromToken    v1.Token
	ToToken      v1.Token
	WalletPK     string
	EstimateOnly bool
	Gas          *Gas
	Debug        bool
}

type DefaultRes struct {
	Tx        *Transaction
	ApproveTx *Transaction
	ECost     *EstimatedGasCost
}

type SyncSwapRes = DefaultRes

type DefaultBridgeReq struct {
	FromNetwork  v1.Network
	ToNetwork    v1.Network
	WalletPK     string
	Amount       *big.Int
	FromToken    Token
	ToToken      Token
	Gas          *Gas
	EstimateOnly bool
	Debug        bool
}

type TxCode = string

const (
	CodeApprove  TxCode = "approve"
	CodeContract TxCode = "contract"
	CodeTransfer TxCode = "transfer"
)

type Transaction struct {
	Code    TxCode
	Network v1.Network
	Id      string
	Url     string
}

func (c *EtheriumClient) NewTx(id common.Hash, code TxCode) *Transaction {
	return &Transaction{
		Code:    code,
		Network: c.Cfg.Network,
		Id:      id.String(),
		Url:     c.TxViewFn(id.String()),
	}
}

type EstimatedGasCost struct {
	GasLimit    *big.Int
	GasPrice    *big.Int
	TotalGasWei *big.Int
	L2Fee       *big.Int
}

type SyncSwapper interface {
	Networker
	SyncSwap(ctx context.Context, req *SyncSwapReq) (*SyncSwapRes, error)
	WaitSwapFinish(ctx context.Context, txId common.Hash) (*scan.TxStatus, error)
}

type ZkSyncOfficialWithdrawalEtherium interface {
	Networker
}

type Transfer interface {
	Networker
	Transfer(ctx context.Context, r *TransferReq) (*TransferRes, error)
}

type WETHReq struct {
	Amount *big.Int

	Wrap bool

	WalletPK     string
	EstimateOnly bool
	Gas          *Gas
}

type WETHRes struct {
	Tx    *Transaction
	ECost *EstimatedGasCost
}

type WETH interface {
	Networker
	SwapWETH(ctx context.Context, req *WETHReq) (*WETHRes, error)
}
