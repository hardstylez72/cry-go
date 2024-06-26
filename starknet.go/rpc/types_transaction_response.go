package rpc

import "github.com/hardstylez72/cry/starknet.go/felt"

// AddDeclareTransactionResponse provides the output for AddDeclareTransaction.
type AddDeclareTransactionResponse struct {
	TransactionHash *felt.Felt `json:"transaction_hash"`
	ClassHash       *felt.Felt `json:"class_hash"`
}

// AddDeployTransactionResponse provides the output for AddDeployTransaction.
type AddDeployTransactionResponse struct {
	TransactionHash *felt.Felt `json:"transaction_hash"`
	ContractAddress *felt.Felt `json:"contract_address"`
}

type AddInvokeTransactionResponse struct {
	TransactionHash *felt.Felt `json:"transaction_hash"`
}
