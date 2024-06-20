package zksyncera

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/zksyncera/abi/across"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type AcrossBridge struct {
	*Client
	HttpCli *http.Client
	Debug   bool
}

func (a *AcrossBridge) Bridge(ctx context.Context, req *defi.DefaultBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {

	if taskType != v1.TaskType_AcrossBridge {
		return nil, errors.New("unsupported task type: " + taskType.String())
	}

	return a.GenericBridge(ctx, a, req)
}

func (a *AcrossBridge) MakeBridgeTx(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	ca := common.HexToAddress("0xE0B015E54d54fc84a6cB9B666099c46adE9335FF")

	wt, err := NewWalletTransactor(req.PK, a.NetworkId)
	if err != nil {
		return nil, err
	}

	fromToken, supported := a.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	fee, err := a.SuggestFee(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "SuggestFee")
	}

	if a.Debug {
		log.Log.Debug("AcrossBridge.SuggestFee", fmt.Sprintf("%v", fee))
	}

	var recipient common.Address    //+
	var originToken common.Address  //+
	var amount *big.Int             //+
	var destinationChainId *big.Int //++
	var relayerFeePct int64         //+
	var quoteTimestamp uint32       //+
	var message []byte              //+
	var maxCount *big.Int

	relayerFeePct, err = lib.StringToInt(fee.RelayFeePct)
	if err != nil {
		return nil, err
	}

	ts, err := lib.StringToInt(fee.Timestamp)
	if err != nil {
		return nil, err
	}

	quoteTimestamp = uint32(ts)

	message = []byte{}

	maxCount = big.NewInt(math.MaxInt64)

	recipient = wt.WalletAddr
	originToken = fromToken
	amount = req.Amount
	destinationChainId = bozdo.ChainMap[req.ToNetwork]

	contract, err := across.StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := contract.Pack("deposit", recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
	if err != nil {
		return nil, err
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: ca,
		Details:      nil,
		Rate:         nil,
		Gas:          nil,
	}, nil
}

type SuggestFeeRes struct {
	CapitalFeePct    string `json:"capitalFeePct"`
	CapitalFeeTotal  string `json:"capitalFeeTotal"`
	RelayGasFeePct   string `json:"relayGasFeePct"`
	RelayGasFeeTotal string `json:"relayGasFeeTotal"`
	RelayFeePct      string `json:"relayFeePct"`
	RelayFeeTotal    string `json:"relayFeeTotal"`
	LpFeePct         string `json:"lpFeePct"`
	Timestamp        string `json:"timestamp"`
	IsAmountTooLow   bool   `json:"isAmountTooLow"`
	QuoteBlock       string `json:"quoteBlock"`
	SpokePoolAddress string `json:"spokePoolAddress"`
}

func (a *AcrossBridge) SuggestFee(ctx context.Context, arg *defi.DefaultBridgeReq) (*SuggestFeeRes, error) {

	fromToken, supported := a.Cfg.TokenMap[arg.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(arg.FromToken)
	}

	url := strings.Join([]string{
		"https://across.to/api/suggested-fees?token=", fromToken.String(),
		"&destinationChainId=", bozdo.ChainMap[arg.ToNetwork].String(),
		"&originChainId=", bozdo.ChainMap[arg.FromNetwork].String(),
		"&amount=", arg.Amount.String(),
		"&skipAmountLimit=false",
	}, "")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := a.HttpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(b))
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r SuggestFeeRes
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *AcrossBridge) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {

	time.Sleep(time.Second * 30)

	return nil

	tx, _, err := a.ClientL2.TransactionByHash(ctx, common.HexToHash(txId))
	if err != nil {
		return errors.Wrap(err, "cli.ClientL2.TransactionByHash")
	}

	wa := tx.From

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		if a.Debug {
			log.Log.Debug("AcrossBridge.WaitTxComplete.GetTx")
		}

		status, err := a.GetTx(ctx, wa, txId)
		if err != nil {
			if errors.Is(err, defi.ErrTxNotFound) {
				time.Sleep(time.Second * 10)
				continue
			} else {
				return err
			}
		}

		switch *status {
		case "pending":
			continue
		case "filled":
			return nil
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

type GetTxRes struct {
	Deposits   []AcrossDeposit `json:"deposits"`
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"pagination"`
}

type AcrossDeposit struct {
	DepositId              int           `json:"depositId"`
	DepositTime            int           `json:"depositTime"`
	Status                 string        `json:"status"`
	Filled                 string        `json:"filled"`
	SourceChainId          int           `json:"sourceChainId"`
	DestinationChainId     int           `json:"destinationChainId"`
	AssetAddr              string        `json:"assetAddr"`
	DepositorAddr          string        `json:"depositorAddr"`
	RecipientAddr          string        `json:"recipientAddr"`
	Message                string        `json:"message"`
	Amount                 string        `json:"amount"`
	DepositTxHash          string        `json:"depositTxHash"`
	FillTxs                []string      `json:"fillTxs"`
	SpeedUps               []interface{} `json:"speedUps"`
	DepositRelayerFeePct   string        `json:"depositRelayerFeePct"`
	InitialRelayerFeePct   string        `json:"initialRelayerFeePct"`
	SuggestedRelayerFeePct string        `json:"suggestedRelayerFeePct"`
}

func (a *AcrossBridge) GetTx(ctx context.Context, wa common.Address, txId string) (*string, error) {
	limit := 10
	offset := 0

	for {

		if a.Debug {
			log.Log.Debug("AcrossBridge.GetTx", fmt.Sprintf("limit %d offset: %d", limit, offset))
		}

		res, err := a.GetTxOffset(ctx, wa, limit, offset)
		if err != nil {
			return nil, errors.Wrap(err, "a.GetTx")
		}

		if a.Debug {
			log.Log.Debug("AcrossBridge.GetTx", fmt.Sprintf("total: %d", res.Pagination.Total))
		}

		for _, d := range res.Deposits {
			if d.DepositTxHash == txId {
				return &d.Status, nil
			}
		}

		if offset+limit < res.Pagination.Total {
			offset += limit
		} else {
			return nil, defi.ErrTxNotFound
		}
	}
}

// https://across.to/api/deposits?address=0x3C7519E33dD2a659A2E8544Ac3E57E00d7a93fd3&limit=10&offset=0
// https://api.across.to/deposits?address=0x4A6e7c137a6691D55693CA3Bc7E5C698d9d43815&limit=10&offset=0
func (a *AcrossBridge) GetTxOffset(ctx context.Context, wa common.Address, limit, offset int) (*GetTxRes, error) {
	url := strings.Join([]string{
		"https://api.across.to/deposits?address=", wa.String(),
		"&limit=", strconv.Itoa(limit),
		"&offset=", strconv.Itoa(offset),
	}, "")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := a.HttpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data GetTxRes
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
