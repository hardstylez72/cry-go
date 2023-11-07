package across

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Bridge struct {
	*defi.EtheriumClient
	HttpCli *http.Client
	Debug   bool
}

func NewAcrossBridge(cli *defi.EtheriumClient) *Bridge {
	return &Bridge{
		EtheriumClient: cli,
		HttpCli:        &http.Client{},
		Debug:          false,
	}
}

func (a *Bridge) Bridge(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {

	var data *bozdo.TxData
	var err error
	switch true {
	case req.FromToken == v1.Token_ETH:
		data, err = a.MakeBridgeTx(ctx, req)
	default:
		data, err = a.MakeBridgeTx2(ctx, req)
	}
	if err != nil {
		return nil, err
	}

	wt, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	approve, err := a.TokenLimitChecker(ctx, &defi.TokenLimitCheckerReq{
		Token:       req.FromToken,
		Wallet:      wt,
		Amount:      req.Amount,
		SpenderAddr: data.ContractAddr,
	})
	if err != nil {
		return nil, err
	}

	opt := &defi.TxOpt{
		NoSend:       req.EstimateOnly,
		Pk:           req.PK,
		Gas:          req.Gas,
		ExchangeRate: nil,
		Debug:        req.Debug,
		TaskType:     v1.TaskType_AcrossBridge,
	}
	if req.Debug {
		log.Log.Debug("across bridge tx: ", zap.String("to", data.ContractAddr.String()))
	}
	tx, err := a.LondonReadyTx(ctx, opt, data)
	if err != nil {
		return nil, err
	}

	if approve.ApproveTx != nil {
		tx.ApproveTx = a.NewTx(approve.ApproveTx.Hash(), bozdo.CodeApprove, nil)
	}

	return tx, nil
}

func (a *Bridge) MakeBridgeTx2(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
	}

	ca := a.Cfg.Dict.Across.RouterToken

	wt, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	_, supported := a.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	fee, err := a.SuggestFee(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "SuggestFee")
	}

	if a.Debug {
		log.Log.Debug("Bridge.SuggestFee", fmt.Sprintf("%v", fee))
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

	maxCount, _ = big.NewInt(0).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)

	recipient = wt.WalletAddr
	fromTokenResolved, err := a.ResolveFromToken(req.FromNetwork, req.FromToken)
	if err != nil {
		return nil, err
	}
	originToken = *fromTokenResolved

	amount = req.Amount
	destinationChainId = bozdo.ChainMap[req.ToNetwork]

	contract, err := Router2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := contract.Pack("deposit", recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(data)))
	}

	extraFee, ok := big.NewInt(0).SetString(fee.RelayFeeTotal, 10)
	if !ok {
		extraFee = big.NewInt(0)
	}

	return &bozdo.TxData{
		Data:         data,
		Value:        value,
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewAcrossDetails(fee.RelayFeeTotal, req.FromToken, req.FromNetwork)},
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     extraFee,
		Code:         bozdo.CodeBridge,
	}, nil
}

func (a *Bridge) MakeBridgeTx(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.TxData, error) {

	value := big.NewInt(0)
	ca := a.Cfg.Dict.Across.RouterToken
	if req.FromToken == v1.Token_ETH {
		value = req.Amount
		ca = a.Cfg.Dict.Across.RouterETH
	}

	wt, err := defi.NewWalletTransactor(req.PK)
	if err != nil {
		return nil, err
	}

	_, supported := a.Cfg.TokenMap[req.FromToken]
	if !supported {
		return nil, defi.ErrTokenNotSupportedFn(req.FromToken)
	}

	fee, err := a.SuggestFee(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "SuggestFee")
	}

	if a.Debug {
		log.Log.Debug("Bridge.SuggestFee", fmt.Sprintf("%v", fee))
	}

	var spokePool common.Address
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

	spokePool = common.HexToAddress(fee.SpokePoolAddress)

	quoteTimestamp = uint32(ts)

	message = []byte{}

	maxCount, _ = big.NewInt(0).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)

	recipient = wt.WalletAddr
	fromTokenResolved, err := a.ResolveFromToken(req.FromNetwork, req.FromToken)
	if err != nil {
		return nil, err
	}
	originToken = *fromTokenResolved

	amount = req.Amount
	destinationChainId = bozdo.ChainMap[req.ToNetwork]

	contract, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	data, err := contract.Pack("deposit", spokePool, recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
	if err != nil {
		return nil, err
	}

	if req.Debug {
		log.Log.Debug(zap.String("data", "0x"+common.Bytes2Hex(data)))
	}

	extraFee, ok := big.NewInt(0).SetString(fee.RelayFeeTotal, 10)
	if !ok {
		extraFee = big.NewInt(0)
	}

	//lpFee, ok := big.NewInt(0).SetString(fee.LpFeePct, 10)
	//if !ok {
	//	lpFee = big.NewInt(0)
	//}

	//lpFee = big.NewInt(0).Quo(lpFee, big.NewInt(1000))

	return &bozdo.TxData{
		Data:         data,
		Value:        bozdo.BigIntSum(value),
		ContractAddr: ca,
		Details:      []bozdo.TxDetail{bozdo.NewAcrossDetails(fee.RelayFeeTotal, req.FromToken, req.FromNetwork)},
		Rate:         nil,
		Gas:          nil,
		ExtraFee:     extraFee,
		Code:         bozdo.CodeBridge,
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

func (a *Bridge) ResolveFromToken(network v1.Network, token v1.Token) (*common.Address, error) {
	switch true {
	case network == v1.Network_ARBITRUM && token == v1.Token_ETH:
		addr := common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")
		return &addr, nil
	case network == v1.Network_OPTIMISM && token == v1.Token_ETH:
		addr := common.HexToAddress("0x4200000000000000000000000000000000000006")
		return &addr, nil
	default:
		fromToken, supported := a.Cfg.TokenMap[token]
		if !supported {
			return nil, defi.ErrTokenNotSupportedFn(token)
		}
		return &fromToken, nil
	}
}

func (a *Bridge) SuggestFee(ctx context.Context, arg *defi.DefaultBridgeReq) (*SuggestFeeRes, error) {

	token, err := a.ResolveFromToken(arg.FromNetwork, arg.FromToken)
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{
		"https://across.to/api/suggested-fees?token=", token.String(),
		"&destinationChainId=", bozdo.ChainMap[arg.ToNetwork].String(),
		"&originChainId=", bozdo.ChainMap[arg.FromNetwork].String(),
		"&amount=", arg.Amount.String(),
		"&skipAmountLimit=false",
	}, "")

	//https://across.to/api/suggested-fees?token=0x82aF49447D8a07e3bd95BD0d56f35241523fBab1
	//https://across.to/api/suggested-fees?token=0x0000000000000000000000000000000000000000
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

func (a *Bridge) WaitForConfirm(ctx context.Context, txId string, receiver string) error {

	wa := common.HexToAddress(receiver)

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		if a.Debug {
			log.Log.Debug("Bridge.WaitTxComplete.GetTx")
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

func (a *Bridge) GetTx(ctx context.Context, wa common.Address, txId string) (*string, error) {
	limit := 10
	offset := 0

	for {

		if a.Debug {
			log.Log.Debug("Bridge.GetTx", fmt.Sprintf("limit %d offset: %d", limit, offset))
		}

		res, err := a.GetTxOffset(ctx, wa, limit, offset)
		if err != nil {
			return nil, errors.Wrap(err, "a.GetTx")
		}

		if a.Debug {
			log.Log.Debug("Bridge.GetTx", fmt.Sprintf("total: %d", res.Pagination.Total))
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
func (a *Bridge) GetTxOffset(ctx context.Context, wa common.Address, limit, offset int) (*GetTxRes, error) {
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
