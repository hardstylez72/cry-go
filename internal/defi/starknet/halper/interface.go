package halper

import (
	"context"
)

type IsAccountDeployedReq struct {
	ChainRPC   string `json:"chainRPC"`
	PrivateKey string `json:"privateKey"`
	Proxy      string `json:"proxy"`
	SubType    string `json:"account"`
}

type IsAccountDeployedRes struct {
	SubType  string `json:"account"`
	Deployed bool   `json:"deployed"`
}

type DeployAccountReq struct {
	ChainRPC     string `json:"chainRPC"`
	PrivateKey   string `json:"privateKey"`
	Proxy        string `json:"proxy"`
	EstimateOnly bool   `json:"estimateOnly"`
	SubType      string `json:"account"`
}

type DeployAccountRes struct {
	EstimatedMaxFee string  `json:"estimatedMaxFee"`
	ContractAddr    *string `json:"contractAddr"`
	TxHash          *string `json:"txHash"`
}

type AccountPubKeyReq struct {
	SubType    string `json:"account"`
	PrivateKey string `json:"privateKey"`
}
type AccountPubKeyRes struct {
	PublicKey string `json:"publicKey"`
}

type BaseTx struct {
	ChainRPC     string `json:"chainRPC"`
	PrivateKey   string `json:"privateKey"`
	Proxy        string `json:"proxy"`
	SubType      string `json:"account"`
	EstimateOnly bool   `json:"estimateOnly"`
	MaxFee       string `json:"maxFee"`
}

type ApproveReq struct {
	BaseTx

	Spender string `json:"spender"`
	Amount  string `json:"amount"`
	Token   string `json:"token"`
	SubType string `json:"account"`
}

type ApproveRes struct {
	Error *string `json:"error"`
	TxId  *string `json:"txId"`
}

type GenerateReq struct {
	SubType string `json:"account"`
	Count   int    `json:"count"`
}

type AccountGenerated struct {
	Pk   string `json:"pk"`
	Seed string `json:"seed"`
	Pub  string `json:"pub"`
}

type LPRes struct {
	DefaultRes
}

type LiquidityBridgeReq struct {
	Proxy        string `json:"proxy"`
	PKEth        string `json:"pkEth"`
	PKStark      string `json:"pkStark"`
	Percent      string `json:"percent"`
	EstimateOnly bool   `json:"estimateOnly"`
	SubType      string `json:"account"`
	FromNetwork  string `json:"fromNetwork"`
	ToNetwork    string `json:"toNetwork"`
}

type LiquidityBridgeRes struct {
	Gas struct {
		Limit string `json:"limit"`
		Price string `json:"price"`
		Total string `json:"total"`
	}

	ContractAddr string
	TxHash       *string
}

type BalanceReq struct {
	ChainRPC string `json:"chainRPC"`
	Proxy    string `json:"proxy"`
	Pub      string `json:"pub"`
	Token    string `json:"token"`
}

type BalanceRes struct {
	WEI string `json:"wei"`
}

type HalperService interface {
	IsAccountDeployed(ctx context.Context, req *IsAccountDeployedReq) (*IsAccountDeployedRes, error)
	DeployAccount(ctx context.Context, req *DeployAccountReq) (*DeployAccountRes, error)
	AccountPubKey(ctx context.Context, req *AccountPubKeyReq) (*AccountPubKeyRes, error)
	Swap(ctx context.Context, req *DefaultSwapReq) (*DefaultSwapRes, error)
	Approve(ctx context.Context, req *ApproveReq) (*ApproveRes, error)
	Generate(ctx context.Context, req *GenerateReq) (*[]AccountGenerated, error)
	LiquidityBridge(ctx context.Context, req *LiquidityBridgeReq) (*LiquidityBridgeRes, error)
	SendDmail(ctx context.Context, req *SendDmailReq) (*SendDmailRes, error)
	Mint(ctx context.Context, req *MintReq) (*MintRes, error)
	Transfer(ctx context.Context, req *TransferReq) (*TransferRes, error)
	ZkLend(ctx context.Context, req *LPRreq) (*LPRes, error)
	Balance(ctx context.Context, req *BalanceReq) (*BalanceRes, error)
	Nostra(ctx context.Context, req *LPRreq) (*LPRes, error)
}
