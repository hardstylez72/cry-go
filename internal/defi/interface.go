package defi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type Token = v1.Token

type Network = v1.Network

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

type MintAndBridgeNFT interface {
	Networker
	MerklyMintNft(ctx context.Context, req *merkly.MintNFTReq) (*bozdo.DefaultRes, *big.Int, *big.Int, error)
	MerklyBridgeNft(ctx context.Context, req *merkly.BridgeNFTReq) (*bozdo.DefaultRes, error)
	GetMerklyNFTId(ctx context.Context, txHash common.Hash) (*big.Int, error)
}

type StargateSwapper interface {
	Networker
	StargateBridgeSwap(ctx context.Context, req *DefaultBridgeReq) (*bozdo.DefaultRes, error)
}

type TestNetworkBridgeSwapper interface {
	Networker
	TestNetBridgeSwap(ctx context.Context, req *TestNetBridgeSwapReq) (*TestNetBridgeSwapRes, error)
}

type Swapper interface {
	Networker
	Swap(ctx context.Context, req *DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error)
}

type OrbiterSwapper interface {
	Transfer
	OrbiterBridge(ctx context.Context, req *OrbiterBridgeReq) (*OrbiterBridgeRes, error)
}

type DefaultSwapReq struct {
	Network      v1.Network
	Amount       *big.Int
	FromToken    v1.Token
	ToToken      v1.Token
	WalletPK     string
	EstimateOnly bool
	Gas          *bozdo.Gas
	Slippage     SlippagePercent
	Debug        bool
	SubType      v1.ProfileSubType
}

type DefaultBridgeReq struct {
	FromNetwork  v1.Network
	ToNetwork    v1.Network
	WalletPK     string
	Amount       *big.Int
	FromToken    Token
	ToToken      Token
	Gas          *bozdo.Gas
	Slippage     SlippagePercent
	EstimateOnly bool
	Debug        bool
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

type WETHReq struct {
	Amount *big.Int

	Wrap bool

	WalletPK     string
	EstimateOnly bool
	Gas          *bozdo.Gas
}

type WETHRes struct {
	Tx    *bozdo.Transaction
	ECost *bozdo.EstimatedGasCost
}

type WETH interface {
	Networker
	SwapWETH(ctx context.Context, req *WETHReq) (*WETHRes, error)
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

type Minter interface {
	Mint(ctx context.Context, req *SimpleReq) (*bozdo.DefaultRes, error)
	Networker
}

type SimpleReq struct {
	PK      string
	SubType v1.ProfileSubType
	*bozdo.BaseReq
}
