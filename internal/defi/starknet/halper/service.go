package halper

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/pkg/errors"
)

type Service struct {
	cli *http.Client
	c   *Config
}

type Config struct {
	Host string
	RPC  string
}

func NewService(c *Config) *Service {
	return &Service{
		cli: &http.Client{Timeout: time.Minute * 2},
		c:   c,
	}
}

type SendDmailReq struct {
	BaseTx
}

type SendDmailRes struct {
	DefaultRes
}

type MintReq struct {
	BaseTx
}

type MintRes struct {
	DefaultRes
}

type TransferReq struct {
	BaseTx
	ToAddr string `json:"toAddr"`
	Token  string `json:"token"`
	Amount string `json:"amount"`
}

type TransferRes struct {
	DefaultRes
}

func (s *Service) Transfer(ctx context.Context, req *TransferReq) (*TransferRes, error) {
	return Request[TransferReq, TransferRes](ctx, s.cli, s.c.Host+"/starknet/transfer", req)
}
func (s *Service) Mint(ctx context.Context, req *MintReq) (*MintRes, error) {
	return Request[MintReq, MintRes](ctx, s.cli, s.c.Host+"/starknet/mint", req)
}

func (s *Service) SendDmail(ctx context.Context, req *SendDmailReq) (*SendDmailRes, error) {
	return Request[SendDmailReq, SendDmailRes](ctx, s.cli, s.c.Host+"/starknet/dmail", req)
}
func (s *Service) LiquidityBridge(ctx context.Context, req *LiquidityBridgeReq) (*LiquidityBridgeRes, error) {
	return Request[LiquidityBridgeReq, LiquidityBridgeRes](ctx, s.cli, s.c.Host+"/starknet/liquidity-bridge", req)
}

func (s *Service) Generate(ctx context.Context, req *GenerateReq) (*[]AccountGenerated, error) {
	return Request[GenerateReq, []AccountGenerated](ctx, s.cli, s.c.Host+"/starknet/generate", req)
}

func (s *Service) Approve(ctx context.Context, req *ApproveReq) (*ApproveRes, error) {
	return Request[ApproveReq, ApproveRes](ctx, s.cli, s.c.Host+"/starknet/approve", req)
}

func (s *Service) AccountPubKey(ctx context.Context, req *AccountPubKeyReq) (*AccountPubKeyRes, error) {
	return Request[AccountPubKeyReq, AccountPubKeyRes](ctx, s.cli, s.c.Host+"/starknet/account_pub", req)
}

func (s *Service) IsAccountDeployed(ctx context.Context, req *IsAccountDeployedReq) (*IsAccountDeployedRes, error) {
	return Request[IsAccountDeployedReq, IsAccountDeployedRes](ctx, s.cli, s.c.Host+"/starknet/is_account_deployed", req)
}
func (s *Service) DeployAccount(ctx context.Context, req *DeployAccountReq) (*DeployAccountRes, error) {
	return Request[DeployAccountReq, DeployAccountRes](ctx, s.cli, s.c.Host+"/starknet/deploy_account", req)
}

type DefaultSwapReq struct {
	Proxy        string               `json:"proxy"`
	ChainRPC     string               `json:"chainRPC"`
	PrivateKey   string               `json:"privateKey"`
	FromToken    string               `json:"fromToken"`
	ToToken      string               `json:"toToken"`
	Amount       string               `json:"amount"`
	EstimateOnly bool                 `json:"estimateOnly"`
	Fee          string               `json:"fee"`
	Slippage     defi.SlippagePercent `json:"slippage"`
	Platform     string               `json:"platform"`
	SubType      string               `json:"account"`
}

type DefaultSwapRes struct {
	ApproveTx *string `json:"approveTxId"`
	SwapTx    *string `json:"swapTxId"`
	Fee       *string `json:"maxFee"`
	Error     *string `json:"error"`
}

type DefaultRes struct {
	EstimatedMaxFee string  `json:"estimatedMaxFee"`
	ContractAddr    *string `json:"contractAddr"`
	TxHash          *string `json:"txHash"`
}

func (s *Service) Swap(ctx context.Context, req *DefaultSwapReq) (*DefaultSwapRes, error) {
	res, err := Request[DefaultSwapReq, DefaultSwapRes](ctx, s.cli, s.c.Host+"/starknet/swap", req)
	if err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, errors.New(*res.Error)
	}

	return res, nil
}

func Request[REQ any, RES any](ctx context.Context, c *http.Client, url string, req *REQ) (*RES, error) {

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ser RES
	if err := json.Unmarshal(body, &ser); err != nil {
		return nil, err
	}

	return &ser, nil
}
