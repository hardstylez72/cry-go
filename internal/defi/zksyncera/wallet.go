package zksyncera

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

type WalletTransactor struct {
	WalletAddr   common.Address
	WalletAddrHR string
	PK           string
	PKb          []byte
	Signer       *accounts.BaseSigner
}

func NewWalletTransactor(privateKey string, networkId *big.Int) (*WalletTransactor, error) {

	pkb, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "hex.DecodeString")
	}
	signer, err := accounts.NewBaseSignerFromRawPrivateKey(pkb, networkId.Int64())
	if err != nil {
		return nil, err
	}

	return &WalletTransactor{
		WalletAddr:   signer.Address(),
		WalletAddrHR: signer.Address().String(),
		PK:           privateKey,
		PKb:          pkb,
		Signer:       signer,
	}, nil
}

func (c *Client) Wallet(pk string) (*accounts.Wallet, *WalletTransactor, error) {
	wtx, err := NewWalletTransactor(pk, c.NetworkId)
	if err != nil {
		return nil, nil, errors.Wrap(err, "newWalletTransactor")
	}
	w, err := accounts.NewWallet(wtx.PKb, &c.ClientL2, c.ClientL1)
	if err != nil {
		return nil, nil, errors.Wrap(err, "accounts.NewWallet")
	}
	return w, wtx, nil
}

func CreateFunctionCallTransaction(from, to common.Address, gasPrice, gasLimit, value *big.Int, data hexutil.Bytes,
	customSignature hexutil.Bytes, paymasterParams *types.PaymasterParams) *types.CallMsg {

	return &types.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:       from,
			To:         &to,
			Gas:        gasLimit.Uint64(),
			GasPrice:   gasPrice,
			GasFeeCap:  nil,
			GasTipCap:  nil,
			Value:      value,
			Data:       data,
			AccessList: nil,
		},
		Meta: &types.Eip712Meta{
			GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			CustomSignature: customSignature,
			PaymasterParams: paymasterParams,
		},
	}
}
