package traderjoe

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Service struct {
	cli *http.Client
	c   *Config
}

type Config struct {
	Host string
}

func NewService(c *Config) *Service {
	return &Service{
		cli: &http.Client{},
		c:   c,
	}
}

type getSwapData struct {
	FromToken string `json:"fromToken"`
	ToToken   string `json:"toToken"`
	ChainRPC  string `json:"chainRPC"`
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
}

func (s *Service) GetSwapData(ctx context.Context, req *GetSwapDataReq) (*GetSwapDataRes, error) {

	data := &getSwapData{
		FromToken: req.FromToken.String(),
		ToToken:   req.ToToken.String(),
		ChainRPC:  req.ChainRPC,
		Amount:    req.Amount.String(),
		Recipient: req.Recipient.String(),
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, s.c.Host+"/joe/swap-data", bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := s.cli.Do(r)
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

	var ser GetSwapDataRes
	if err := json.Unmarshal(body, &ser); err != nil {
		return nil, err
	}

	return &ser, nil
}
