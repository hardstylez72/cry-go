package halper

import (
	"context"
)

type IsAccountDeployedReq struct {
	ChainRPC   string `json:"chainRPC"`
	PrivateKey string `json:"privateKey"`
	Proxy      string `json:"proxy"`
}

type IsAccountDeployedRes struct {
	Deployed bool `json:"deployed"`
}

type DeployAccountReq struct {
	ChainRPC     string `json:"chainRPC"`
	PrivateKey   string `json:"privateKey"`
	Proxy        string `json:"proxy"`
	EstimateOnly bool   `json:"estimateOnly"`
}

type DeployAccountRes struct {
	EstimatedMaxFee string  `json:"estimatedMaxFee"`
	ContractAddr    *string `json:"contractAddr"`
	TxHash          *string `json:"txHash"`
}

type AccountPubKeyReq struct {
	PrivateKey string `json:"privateKey"`
}
type AccountPubKeyRes struct {
	PublicKey string `json:"publicKey"`
}

type HalperService interface {
	IsAccountDeployed(ctx context.Context, req *IsAccountDeployedReq) (*IsAccountDeployedRes, error)
	DeployAccount(ctx context.Context, req *DeployAccountReq) (*DeployAccountRes, error)
	AccountPubKey(ctx context.Context, req *AccountPubKeyReq) (*AccountPubKeyRes, error)
	Swap10k(ctx context.Context, req *DefaultSwapReq) (*DefaultSwapRes, error)
}
