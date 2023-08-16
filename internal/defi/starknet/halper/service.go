package halper

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

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
	Proxy        string `json:"proxy"`
	ChainRPC     string `json:"chainRPC"`
	PrivateKey   string `json:"privateKey"`
	FromToken    string `json:"fromToken"`
	ToToken      string `json:"toToken"`
	Amount       string `json:"amount"`
	EstimateOnly bool   `json:"estimateOnly"`
	Fee          string `json:"fee"`
}

type DefaultSwapRes struct {
	ApproveTx *string `json:"approveTxId"`
	SwapTx    *string `json:"swapTxId"`
	Fee       *string `json:"maxFee"`
	Error     *string `json:"error"`
}

func (s *Service) Swap10k(ctx context.Context, req *DefaultSwapReq) (*DefaultSwapRes, error) {
	res, err := Request[DefaultSwapReq, DefaultSwapRes](ctx, s.cli, s.c.Host+"/starknet/10k_swap", req)
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
