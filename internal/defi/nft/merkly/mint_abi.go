// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merkly

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MinterMetaData contains all meta data concerning the Minter contract.
var MinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minGasToTransfer\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_layerZeroEndpoint\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startMintId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endMintId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hashedPayload\",\"type\":\"bytes32\"}],\"name\":\"CreditCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hashedPayload\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"CreditStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"MessageFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ReceiveFromChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"RetryMessageSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"SendToChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dstChainIdToBatchLimit\",\"type\":\"uint256\"}],\"name\":\"SetDstChainIdToBatchLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dstChainIdToTransferGas\",\"type\":\"uint256\"}],\"name\":\"SetDstChainIdToTransferGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_type\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minDstGas\",\"type\":\"uint256\"}],\"name\":\"SetMinDstGas\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minGasToTransferAndStore\",\"type\":\"uint256\"}],\"name\":\"SetMinGasToTransferAndStore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"precrime\",\"type\":\"address\"}],\"name\":\"SetPrecrime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"SetTrustedRemoteAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_PAYLOAD_SIZE_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUNCTION_TYPE_SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"bridgeGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"clearCredits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"dstChainIdToBatchLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"dstChainIdToTransferGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateGasBridgeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateSendBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useZro\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateSendFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"forceResumeReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"}],\"name\":\"getTrustedRemoteAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"}],\"name\":\"isTrustedRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzEndpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"minDstGasLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasToTransferAndStore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMintId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonblockingLzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"payloadSizeLimitLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"retryMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"sendBatchFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"sendFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainIdToBatchLimit\",\"type\":\"uint256\"}],\"name\":\"setDstChainIdToBatchLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainIdToTransferGas\",\"type\":\"uint256\"}],\"name\":\"setDstChainIdToTransferGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_packetType\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minGas\",\"type\":\"uint256\"}],\"name\":\"setMinDstGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minGasToTransferAndStore\",\"type\":\"uint256\"}],\"name\":\"setMinGasToTransferAndStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setPayloadSizeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_precrime\",\"type\":\"address\"}],\"name\":\"setPrecrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setReceiveVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_version\",\"type\":\"uint16\"}],\"name\":\"setSendVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_remoteAddress\",\"type\":\"bytes\"}],\"name\":\"setTrustedRemoteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"storedCredits\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"creditsRemain\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemoteLookup\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MinterABI is the input ABI used to generate the binding from.
// Deprecated: Use MinterMetaData.ABI instead.
var MinterABI = MinterMetaData.ABI

// Minter is an auto generated Go binding around an Ethereum contract.
type Minter struct {
	MinterCaller     // Read-only binding to the contract
	MinterTransactor // Write-only binding to the contract
	MinterFilterer   // Log filterer for contract events
}

// MinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterSession struct {
	Contract     *Minter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterCallerSession struct {
	Contract *MinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterTransactorSession struct {
	Contract     *MinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterRaw struct {
	Contract *Minter // Generic contract binding to access the raw methods on
}

// MinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterCallerRaw struct {
	Contract *MinterCaller // Generic read-only contract binding to access the raw methods on
}

// MinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterTransactorRaw struct {
	Contract *MinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinter creates a new instance of Minter, bound to a specific deployed contract.
func NewMinter(address common.Address, backend bind.ContractBackend) (*Minter, error) {
	contract, err := bindMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Minter{MinterCaller: MinterCaller{contract: contract}, MinterTransactor: MinterTransactor{contract: contract}, MinterFilterer: MinterFilterer{contract: contract}}, nil
}

// NewMinterCaller creates a new read-only instance of Minter, bound to a specific deployed contract.
func NewMinterCaller(address common.Address, caller bind.ContractCaller) (*MinterCaller, error) {
	contract, err := bindMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterCaller{contract: contract}, nil
}

// NewMinterTransactor creates a new write-only instance of Minter, bound to a specific deployed contract.
func NewMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterTransactor, error) {
	contract, err := bindMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterTransactor{contract: contract}, nil
}

// NewMinterFilterer creates a new log filterer instance of Minter, bound to a specific deployed contract.
func NewMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterFilterer, error) {
	contract, err := bindMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterFilterer{contract: contract}, nil
}

// bindMinter binds a generic wrapper to an already deployed contract.
func bindMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Minter *MinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Minter.Contract.MinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Minter *MinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.Contract.MinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Minter *MinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Minter.Contract.MinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Minter *MinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Minter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Minter *MinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Minter *MinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Minter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Minter *MinterCaller) DEFAULTPAYLOADSIZELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "DEFAULT_PAYLOAD_SIZE_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Minter *MinterSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Minter.Contract.DEFAULTPAYLOADSIZELIMIT(&_Minter.CallOpts)
}

// DEFAULTPAYLOADSIZELIMIT is a free data retrieval call binding the contract method 0xc4461834.
//
// Solidity: function DEFAULT_PAYLOAD_SIZE_LIMIT() view returns(uint256)
func (_Minter *MinterCallerSession) DEFAULTPAYLOADSIZELIMIT() (*big.Int, error) {
	return _Minter.Contract.DEFAULTPAYLOADSIZELIMIT(&_Minter.CallOpts)
}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Minter *MinterCaller) FUNCTIONTYPESEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "FUNCTION_TYPE_SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Minter *MinterSession) FUNCTIONTYPESEND() (uint16, error) {
	return _Minter.Contract.FUNCTIONTYPESEND(&_Minter.CallOpts)
}

// FUNCTIONTYPESEND is a free data retrieval call binding the contract method 0xaf3fb21c.
//
// Solidity: function FUNCTION_TYPE_SEND() view returns(uint16)
func (_Minter *MinterCallerSession) FUNCTIONTYPESEND() (uint16, error) {
	return _Minter.Contract.FUNCTIONTYPESEND(&_Minter.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Minter *MinterCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Minter *MinterSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Minter.Contract.BalanceOf(&_Minter.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Minter *MinterCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Minter.Contract.BalanceOf(&_Minter.CallOpts, owner)
}

// DstChainIdToBatchLimit is a free data retrieval call binding the contract method 0x4ac3f4ff.
//
// Solidity: function dstChainIdToBatchLimit(uint16 ) view returns(uint256)
func (_Minter *MinterCaller) DstChainIdToBatchLimit(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "dstChainIdToBatchLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DstChainIdToBatchLimit is a free data retrieval call binding the contract method 0x4ac3f4ff.
//
// Solidity: function dstChainIdToBatchLimit(uint16 ) view returns(uint256)
func (_Minter *MinterSession) DstChainIdToBatchLimit(arg0 uint16) (*big.Int, error) {
	return _Minter.Contract.DstChainIdToBatchLimit(&_Minter.CallOpts, arg0)
}

// DstChainIdToBatchLimit is a free data retrieval call binding the contract method 0x4ac3f4ff.
//
// Solidity: function dstChainIdToBatchLimit(uint16 ) view returns(uint256)
func (_Minter *MinterCallerSession) DstChainIdToBatchLimit(arg0 uint16) (*big.Int, error) {
	return _Minter.Contract.DstChainIdToBatchLimit(&_Minter.CallOpts, arg0)
}

// DstChainIdToTransferGas is a free data retrieval call binding the contract method 0xfa25f9b6.
//
// Solidity: function dstChainIdToTransferGas(uint16 ) view returns(uint256)
func (_Minter *MinterCaller) DstChainIdToTransferGas(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "dstChainIdToTransferGas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DstChainIdToTransferGas is a free data retrieval call binding the contract method 0xfa25f9b6.
//
// Solidity: function dstChainIdToTransferGas(uint16 ) view returns(uint256)
func (_Minter *MinterSession) DstChainIdToTransferGas(arg0 uint16) (*big.Int, error) {
	return _Minter.Contract.DstChainIdToTransferGas(&_Minter.CallOpts, arg0)
}

// DstChainIdToTransferGas is a free data retrieval call binding the contract method 0xfa25f9b6.
//
// Solidity: function dstChainIdToTransferGas(uint16 ) view returns(uint256)
func (_Minter *MinterCallerSession) DstChainIdToTransferGas(arg0 uint16) (*big.Int, error) {
	return _Minter.Contract.DstChainIdToTransferGas(&_Minter.CallOpts, arg0)
}

// EstimateGasBridgeFee is a free data retrieval call binding the contract method 0xda09f15d.
//
// Solidity: function estimateGasBridgeFee(uint16 _dstChainId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterCaller) EstimateGasBridgeFee(opts *bind.CallOpts, _dstChainId uint16, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "estimateGasBridgeFee", _dstChainId, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateGasBridgeFee is a free data retrieval call binding the contract method 0xda09f15d.
//
// Solidity: function estimateGasBridgeFee(uint16 _dstChainId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterSession) EstimateGasBridgeFee(_dstChainId uint16, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Minter.Contract.EstimateGasBridgeFee(&_Minter.CallOpts, _dstChainId, _useZro, _adapterParams)
}

// EstimateGasBridgeFee is a free data retrieval call binding the contract method 0xda09f15d.
//
// Solidity: function estimateGasBridgeFee(uint16 _dstChainId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterCallerSession) EstimateGasBridgeFee(_dstChainId uint16, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Minter.Contract.EstimateGasBridgeFee(&_Minter.CallOpts, _dstChainId, _useZro, _adapterParams)
}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xf2353641.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterCaller) EstimateSendBatchFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "estimateSendBatchFee", _dstChainId, _toAddress, _tokenIds, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xf2353641.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterSession) EstimateSendBatchFee(_dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Minter.Contract.EstimateSendBatchFee(&_Minter.CallOpts, _dstChainId, _toAddress, _tokenIds, _useZro, _adapterParams)
}

// EstimateSendBatchFee is a free data retrieval call binding the contract method 0xf2353641.
//
// Solidity: function estimateSendBatchFee(uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterCallerSession) EstimateSendBatchFee(_dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Minter.Contract.EstimateSendBatchFee(&_Minter.CallOpts, _dstChainId, _toAddress, _tokenIds, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterCaller) EstimateSendFee(opts *bind.CallOpts, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "estimateSendFee", _dstChainId, _toAddress, _tokenId, _useZro, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Minter.Contract.EstimateSendFee(&_Minter.CallOpts, _dstChainId, _toAddress, _tokenId, _useZro, _adapterParams)
}

// EstimateSendFee is a free data retrieval call binding the contract method 0x2a205e3d.
//
// Solidity: function estimateSendFee(uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, bool _useZro, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_Minter *MinterCallerSession) EstimateSendFee(_dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _useZro bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _Minter.Contract.EstimateSendFee(&_Minter.CallOpts, _dstChainId, _toAddress, _tokenId, _useZro, _adapterParams)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Minter *MinterCaller) FailedMessages(opts *bind.CallOpts, arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "failedMessages", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Minter *MinterSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Minter.Contract.FailedMessages(&_Minter.CallOpts, arg0, arg1, arg2)
}

// FailedMessages is a free data retrieval call binding the contract method 0x5b8c41e6.
//
// Solidity: function failedMessages(uint16 , bytes , uint64 ) view returns(bytes32)
func (_Minter *MinterCallerSession) FailedMessages(arg0 uint16, arg1 []byte, arg2 uint64) ([32]byte, error) {
	return _Minter.Contract.FailedMessages(&_Minter.CallOpts, arg0, arg1, arg2)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Minter *MinterCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Minter *MinterSession) Fee() (*big.Int, error) {
	return _Minter.Contract.Fee(&_Minter.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Minter *MinterCallerSession) Fee() (*big.Int, error) {
	return _Minter.Contract.Fee(&_Minter.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Minter *MinterCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Minter *MinterSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Minter.Contract.GetApproved(&_Minter.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Minter *MinterCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Minter.Contract.GetApproved(&_Minter.CallOpts, tokenId)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Minter *MinterCaller) GetConfig(opts *bind.CallOpts, _version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "getConfig", _version, _chainId, arg2, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Minter *MinterSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Minter.Contract.GetConfig(&_Minter.CallOpts, _version, _chainId, arg2, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0xf5ecbdbc.
//
// Solidity: function getConfig(uint16 _version, uint16 _chainId, address , uint256 _configType) view returns(bytes)
func (_Minter *MinterCallerSession) GetConfig(_version uint16, _chainId uint16, arg2 common.Address, _configType *big.Int) ([]byte, error) {
	return _Minter.Contract.GetConfig(&_Minter.CallOpts, _version, _chainId, arg2, _configType)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Minter *MinterCaller) GetTrustedRemoteAddress(opts *bind.CallOpts, _remoteChainId uint16) ([]byte, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "getTrustedRemoteAddress", _remoteChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Minter *MinterSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Minter.Contract.GetTrustedRemoteAddress(&_Minter.CallOpts, _remoteChainId)
}

// GetTrustedRemoteAddress is a free data retrieval call binding the contract method 0x9f38369a.
//
// Solidity: function getTrustedRemoteAddress(uint16 _remoteChainId) view returns(bytes)
func (_Minter *MinterCallerSession) GetTrustedRemoteAddress(_remoteChainId uint16) ([]byte, error) {
	return _Minter.Contract.GetTrustedRemoteAddress(&_Minter.CallOpts, _remoteChainId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Minter *MinterCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Minter *MinterSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Minter.Contract.IsApprovedForAll(&_Minter.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Minter *MinterCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Minter.Contract.IsApprovedForAll(&_Minter.CallOpts, owner, operator)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Minter *MinterCaller) IsTrustedRemote(opts *bind.CallOpts, _srcChainId uint16, _srcAddress []byte) (bool, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "isTrustedRemote", _srcChainId, _srcAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Minter *MinterSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Minter.Contract.IsTrustedRemote(&_Minter.CallOpts, _srcChainId, _srcAddress)
}

// IsTrustedRemote is a free data retrieval call binding the contract method 0x3d8b38f6.
//
// Solidity: function isTrustedRemote(uint16 _srcChainId, bytes _srcAddress) view returns(bool)
func (_Minter *MinterCallerSession) IsTrustedRemote(_srcChainId uint16, _srcAddress []byte) (bool, error) {
	return _Minter.Contract.IsTrustedRemote(&_Minter.CallOpts, _srcChainId, _srcAddress)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Minter *MinterCaller) LzEndpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "lzEndpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Minter *MinterSession) LzEndpoint() (common.Address, error) {
	return _Minter.Contract.LzEndpoint(&_Minter.CallOpts)
}

// LzEndpoint is a free data retrieval call binding the contract method 0xb353aaa7.
//
// Solidity: function lzEndpoint() view returns(address)
func (_Minter *MinterCallerSession) LzEndpoint() (common.Address, error) {
	return _Minter.Contract.LzEndpoint(&_Minter.CallOpts)
}

// MaxMintId is a free data retrieval call binding the contract method 0xe1d4c870.
//
// Solidity: function maxMintId() view returns(uint256)
func (_Minter *MinterCaller) MaxMintId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "maxMintId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMintId is a free data retrieval call binding the contract method 0xe1d4c870.
//
// Solidity: function maxMintId() view returns(uint256)
func (_Minter *MinterSession) MaxMintId() (*big.Int, error) {
	return _Minter.Contract.MaxMintId(&_Minter.CallOpts)
}

// MaxMintId is a free data retrieval call binding the contract method 0xe1d4c870.
//
// Solidity: function maxMintId() view returns(uint256)
func (_Minter *MinterCallerSession) MaxMintId() (*big.Int, error) {
	return _Minter.Contract.MaxMintId(&_Minter.CallOpts)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Minter *MinterCaller) MinDstGasLookup(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "minDstGasLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Minter *MinterSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Minter.Contract.MinDstGasLookup(&_Minter.CallOpts, arg0, arg1)
}

// MinDstGasLookup is a free data retrieval call binding the contract method 0x8cfd8f5c.
//
// Solidity: function minDstGasLookup(uint16 , uint16 ) view returns(uint256)
func (_Minter *MinterCallerSession) MinDstGasLookup(arg0 uint16, arg1 uint16) (*big.Int, error) {
	return _Minter.Contract.MinDstGasLookup(&_Minter.CallOpts, arg0, arg1)
}

// MinGasToTransferAndStore is a free data retrieval call binding the contract method 0x48288190.
//
// Solidity: function minGasToTransferAndStore() view returns(uint256)
func (_Minter *MinterCaller) MinGasToTransferAndStore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "minGasToTransferAndStore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGasToTransferAndStore is a free data retrieval call binding the contract method 0x48288190.
//
// Solidity: function minGasToTransferAndStore() view returns(uint256)
func (_Minter *MinterSession) MinGasToTransferAndStore() (*big.Int, error) {
	return _Minter.Contract.MinGasToTransferAndStore(&_Minter.CallOpts)
}

// MinGasToTransferAndStore is a free data retrieval call binding the contract method 0x48288190.
//
// Solidity: function minGasToTransferAndStore() view returns(uint256)
func (_Minter *MinterCallerSession) MinGasToTransferAndStore() (*big.Int, error) {
	return _Minter.Contract.MinGasToTransferAndStore(&_Minter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Minter *MinterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Minter *MinterSession) Name() (string, error) {
	return _Minter.Contract.Name(&_Minter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Minter *MinterCallerSession) Name() (string, error) {
	return _Minter.Contract.Name(&_Minter.CallOpts)
}

// NextMintId is a free data retrieval call binding the contract method 0x6aa99da3.
//
// Solidity: function nextMintId() view returns(uint256)
func (_Minter *MinterCaller) NextMintId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "nextMintId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMintId is a free data retrieval call binding the contract method 0x6aa99da3.
//
// Solidity: function nextMintId() view returns(uint256)
func (_Minter *MinterSession) NextMintId() (*big.Int, error) {
	return _Minter.Contract.NextMintId(&_Minter.CallOpts)
}

// NextMintId is a free data retrieval call binding the contract method 0x6aa99da3.
//
// Solidity: function nextMintId() view returns(uint256)
func (_Minter *MinterCallerSession) NextMintId() (*big.Int, error) {
	return _Minter.Contract.NextMintId(&_Minter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Minter *MinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Minter *MinterSession) Owner() (common.Address, error) {
	return _Minter.Contract.Owner(&_Minter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Minter *MinterCallerSession) Owner() (common.Address, error) {
	return _Minter.Contract.Owner(&_Minter.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Minter *MinterCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Minter *MinterSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Minter.Contract.OwnerOf(&_Minter.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Minter *MinterCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Minter.Contract.OwnerOf(&_Minter.CallOpts, tokenId)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Minter *MinterCaller) PayloadSizeLimitLookup(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "payloadSizeLimitLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Minter *MinterSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Minter.Contract.PayloadSizeLimitLookup(&_Minter.CallOpts, arg0)
}

// PayloadSizeLimitLookup is a free data retrieval call binding the contract method 0x3f1f4fa4.
//
// Solidity: function payloadSizeLimitLookup(uint16 ) view returns(uint256)
func (_Minter *MinterCallerSession) PayloadSizeLimitLookup(arg0 uint16) (*big.Int, error) {
	return _Minter.Contract.PayloadSizeLimitLookup(&_Minter.CallOpts, arg0)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Minter *MinterCaller) Precrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "precrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Minter *MinterSession) Precrime() (common.Address, error) {
	return _Minter.Contract.Precrime(&_Minter.CallOpts)
}

// Precrime is a free data retrieval call binding the contract method 0x950c8a74.
//
// Solidity: function precrime() view returns(address)
func (_Minter *MinterCallerSession) Precrime() (common.Address, error) {
	return _Minter.Contract.Precrime(&_Minter.CallOpts)
}

// StoredCredits is a free data retrieval call binding the contract method 0x22a3ecf9.
//
// Solidity: function storedCredits(bytes32 ) view returns(uint16 srcChainId, address toAddress, uint256 index, bool creditsRemain)
func (_Minter *MinterCaller) StoredCredits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SrcChainId    uint16
	ToAddress     common.Address
	Index         *big.Int
	CreditsRemain bool
}, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "storedCredits", arg0)

	outstruct := new(struct {
		SrcChainId    uint16
		ToAddress     common.Address
		Index         *big.Int
		CreditsRemain bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcChainId = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.ToAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreditsRemain = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// StoredCredits is a free data retrieval call binding the contract method 0x22a3ecf9.
//
// Solidity: function storedCredits(bytes32 ) view returns(uint16 srcChainId, address toAddress, uint256 index, bool creditsRemain)
func (_Minter *MinterSession) StoredCredits(arg0 [32]byte) (struct {
	SrcChainId    uint16
	ToAddress     common.Address
	Index         *big.Int
	CreditsRemain bool
}, error) {
	return _Minter.Contract.StoredCredits(&_Minter.CallOpts, arg0)
}

// StoredCredits is a free data retrieval call binding the contract method 0x22a3ecf9.
//
// Solidity: function storedCredits(bytes32 ) view returns(uint16 srcChainId, address toAddress, uint256 index, bool creditsRemain)
func (_Minter *MinterCallerSession) StoredCredits(arg0 [32]byte) (struct {
	SrcChainId    uint16
	ToAddress     common.Address
	Index         *big.Int
	CreditsRemain bool
}, error) {
	return _Minter.Contract.StoredCredits(&_Minter.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Minter *MinterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Minter *MinterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Minter.Contract.SupportsInterface(&_Minter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Minter *MinterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Minter.Contract.SupportsInterface(&_Minter.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Minter *MinterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Minter *MinterSession) Symbol() (string, error) {
	return _Minter.Contract.Symbol(&_Minter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Minter *MinterCallerSession) Symbol() (string, error) {
	return _Minter.Contract.Symbol(&_Minter.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Minter *MinterCaller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Minter *MinterSession) TokenURI(id *big.Int) (string, error) {
	return _Minter.Contract.TokenURI(&_Minter.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Minter *MinterCallerSession) TokenURI(id *big.Int) (string, error) {
	return _Minter.Contract.TokenURI(&_Minter.CallOpts, id)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Minter *MinterCaller) TrustedRemoteLookup(opts *bind.CallOpts, arg0 uint16) ([]byte, error) {
	var out []interface{}
	err := _Minter.contract.Call(opts, &out, "trustedRemoteLookup", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Minter *MinterSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Minter.Contract.TrustedRemoteLookup(&_Minter.CallOpts, arg0)
}

// TrustedRemoteLookup is a free data retrieval call binding the contract method 0x7533d788.
//
// Solidity: function trustedRemoteLookup(uint16 ) view returns(bytes)
func (_Minter *MinterCallerSession) TrustedRemoteLookup(arg0 uint16) ([]byte, error) {
	return _Minter.Contract.TrustedRemoteLookup(&_Minter.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Minter *MinterTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Minter *MinterSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.Approve(&_Minter.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Minter *MinterTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.Approve(&_Minter.TransactOpts, to, tokenId)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x5154c0b4.
//
// Solidity: function bridgeGas(uint16 _dstChainId, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterTransactor) BridgeGas(opts *bind.TransactOpts, _dstChainId uint16, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "bridgeGas", _dstChainId, _zroPaymentAddress, _adapterParams)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x5154c0b4.
//
// Solidity: function bridgeGas(uint16 _dstChainId, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterSession) BridgeGas(_dstChainId uint16, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.Contract.BridgeGas(&_Minter.TransactOpts, _dstChainId, _zroPaymentAddress, _adapterParams)
}

// BridgeGas is a paid mutator transaction binding the contract method 0x5154c0b4.
//
// Solidity: function bridgeGas(uint16 _dstChainId, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterTransactorSession) BridgeGas(_dstChainId uint16, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.Contract.BridgeGas(&_Minter.TransactOpts, _dstChainId, _zroPaymentAddress, _adapterParams)
}

// ClearCredits is a paid mutator transaction binding the contract method 0x8ffa1f2a.
//
// Solidity: function clearCredits(bytes _payload) returns()
func (_Minter *MinterTransactor) ClearCredits(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "clearCredits", _payload)
}

// ClearCredits is a paid mutator transaction binding the contract method 0x8ffa1f2a.
//
// Solidity: function clearCredits(bytes _payload) returns()
func (_Minter *MinterSession) ClearCredits(_payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.ClearCredits(&_Minter.TransactOpts, _payload)
}

// ClearCredits is a paid mutator transaction binding the contract method 0x8ffa1f2a.
//
// Solidity: function clearCredits(bytes _payload) returns()
func (_Minter *MinterTransactorSession) ClearCredits(_payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.ClearCredits(&_Minter.TransactOpts, _payload)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Minter *MinterTransactor) ForceResumeReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "forceResumeReceive", _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Minter *MinterSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Minter.Contract.ForceResumeReceive(&_Minter.TransactOpts, _srcChainId, _srcAddress)
}

// ForceResumeReceive is a paid mutator transaction binding the contract method 0x42d65a8d.
//
// Solidity: function forceResumeReceive(uint16 _srcChainId, bytes _srcAddress) returns()
func (_Minter *MinterTransactorSession) ForceResumeReceive(_srcChainId uint16, _srcAddress []byte) (*types.Transaction, error) {
	return _Minter.Contract.ForceResumeReceive(&_Minter.TransactOpts, _srcChainId, _srcAddress)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Minter *MinterTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Minter *MinterSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.LzReceive(&_Minter.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Minter *MinterTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.LzReceive(&_Minter.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Minter *MinterTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Minter *MinterSession) Mint() (*types.Transaction, error) {
	return _Minter.Contract.Mint(&_Minter.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Minter *MinterTransactorSession) Mint() (*types.Transaction, error) {
	return _Minter.Contract.Mint(&_Minter.TransactOpts)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Minter *MinterTransactor) NonblockingLzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "nonblockingLzReceive", _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Minter *MinterSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.NonblockingLzReceive(&_Minter.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// NonblockingLzReceive is a paid mutator transaction binding the contract method 0x66ad5c8a.
//
// Solidity: function nonblockingLzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) returns()
func (_Minter *MinterTransactorSession) NonblockingLzReceive(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.NonblockingLzReceive(&_Minter.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Minter *MinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Minter *MinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Minter.Contract.RenounceOwnership(&_Minter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Minter *MinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Minter.Contract.RenounceOwnership(&_Minter.TransactOpts)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Minter *MinterTransactor) RetryMessage(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "retryMessage", _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Minter *MinterSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.RetryMessage(&_Minter.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// RetryMessage is a paid mutator transaction binding the contract method 0xd1deba1f.
//
// Solidity: function retryMessage(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload) payable returns()
func (_Minter *MinterTransactorSession) RetryMessage(_srcChainId uint16, _srcAddress []byte, _nonce uint64, _payload []byte) (*types.Transaction, error) {
	return _Minter.Contract.RetryMessage(&_Minter.TransactOpts, _srcChainId, _srcAddress, _nonce, _payload)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Minter *MinterTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Minter *MinterSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SafeTransferFrom(&_Minter.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Minter *MinterTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SafeTransferFrom(&_Minter.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Minter *MinterTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Minter *MinterSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Minter.Contract.SafeTransferFrom0(&_Minter.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Minter *MinterTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Minter.Contract.SafeTransferFrom0(&_Minter.TransactOpts, from, to, tokenId, data)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0xab3ffb93.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterTransactor) SendBatchFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "sendBatchFrom", _from, _dstChainId, _toAddress, _tokenIds, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0xab3ffb93.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterSession) SendBatchFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.Contract.SendBatchFrom(&_Minter.TransactOpts, _from, _dstChainId, _toAddress, _tokenIds, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendBatchFrom is a paid mutator transaction binding the contract method 0xab3ffb93.
//
// Solidity: function sendBatchFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256[] _tokenIds, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterTransactorSession) SendBatchFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenIds []*big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.Contract.SendBatchFrom(&_Minter.TransactOpts, _from, _dstChainId, _toAddress, _tokenIds, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterTransactor) SendFrom(opts *bind.TransactOpts, _from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "sendFrom", _from, _dstChainId, _toAddress, _tokenId, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.Contract.SendFrom(&_Minter.TransactOpts, _from, _dstChainId, _toAddress, _tokenId, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SendFrom is a paid mutator transaction binding the contract method 0x51905636.
//
// Solidity: function sendFrom(address _from, uint16 _dstChainId, bytes _toAddress, uint256 _tokenId, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_Minter *MinterTransactorSession) SendFrom(_from common.Address, _dstChainId uint16, _toAddress []byte, _tokenId *big.Int, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _Minter.Contract.SendFrom(&_Minter.TransactOpts, _from, _dstChainId, _toAddress, _tokenId, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Minter *MinterTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Minter *MinterSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Minter.Contract.SetApprovalForAll(&_Minter.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Minter *MinterTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Minter.Contract.SetApprovalForAll(&_Minter.TransactOpts, operator, approved)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Minter *MinterTransactor) SetConfig(opts *bind.TransactOpts, _version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setConfig", _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Minter *MinterSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Minter.Contract.SetConfig(&_Minter.TransactOpts, _version, _chainId, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xcbed8b9c.
//
// Solidity: function setConfig(uint16 _version, uint16 _chainId, uint256 _configType, bytes _config) returns()
func (_Minter *MinterTransactorSession) SetConfig(_version uint16, _chainId uint16, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _Minter.Contract.SetConfig(&_Minter.TransactOpts, _version, _chainId, _configType, _config)
}

// SetDstChainIdToBatchLimit is a paid mutator transaction binding the contract method 0x9ea5d6b1.
//
// Solidity: function setDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit) returns()
func (_Minter *MinterTransactor) SetDstChainIdToBatchLimit(opts *bind.TransactOpts, _dstChainId uint16, _dstChainIdToBatchLimit *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setDstChainIdToBatchLimit", _dstChainId, _dstChainIdToBatchLimit)
}

// SetDstChainIdToBatchLimit is a paid mutator transaction binding the contract method 0x9ea5d6b1.
//
// Solidity: function setDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit) returns()
func (_Minter *MinterSession) SetDstChainIdToBatchLimit(_dstChainId uint16, _dstChainIdToBatchLimit *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetDstChainIdToBatchLimit(&_Minter.TransactOpts, _dstChainId, _dstChainIdToBatchLimit)
}

// SetDstChainIdToBatchLimit is a paid mutator transaction binding the contract method 0x9ea5d6b1.
//
// Solidity: function setDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit) returns()
func (_Minter *MinterTransactorSession) SetDstChainIdToBatchLimit(_dstChainId uint16, _dstChainIdToBatchLimit *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetDstChainIdToBatchLimit(&_Minter.TransactOpts, _dstChainId, _dstChainIdToBatchLimit)
}

// SetDstChainIdToTransferGas is a paid mutator transaction binding the contract method 0xd12473a5.
//
// Solidity: function setDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas) returns()
func (_Minter *MinterTransactor) SetDstChainIdToTransferGas(opts *bind.TransactOpts, _dstChainId uint16, _dstChainIdToTransferGas *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setDstChainIdToTransferGas", _dstChainId, _dstChainIdToTransferGas)
}

// SetDstChainIdToTransferGas is a paid mutator transaction binding the contract method 0xd12473a5.
//
// Solidity: function setDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas) returns()
func (_Minter *MinterSession) SetDstChainIdToTransferGas(_dstChainId uint16, _dstChainIdToTransferGas *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetDstChainIdToTransferGas(&_Minter.TransactOpts, _dstChainId, _dstChainIdToTransferGas)
}

// SetDstChainIdToTransferGas is a paid mutator transaction binding the contract method 0xd12473a5.
//
// Solidity: function setDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas) returns()
func (_Minter *MinterTransactorSession) SetDstChainIdToTransferGas(_dstChainId uint16, _dstChainIdToTransferGas *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetDstChainIdToTransferGas(&_Minter.TransactOpts, _dstChainId, _dstChainIdToTransferGas)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Minter *MinterTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Minter *MinterSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetFee(&_Minter.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Minter *MinterTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetFee(&_Minter.TransactOpts, _fee)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Minter *MinterTransactor) SetMinDstGas(opts *bind.TransactOpts, _dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setMinDstGas", _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Minter *MinterSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetMinDstGas(&_Minter.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinDstGas is a paid mutator transaction binding the contract method 0xdf2a5b3b.
//
// Solidity: function setMinDstGas(uint16 _dstChainId, uint16 _packetType, uint256 _minGas) returns()
func (_Minter *MinterTransactorSession) SetMinDstGas(_dstChainId uint16, _packetType uint16, _minGas *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetMinDstGas(&_Minter.TransactOpts, _dstChainId, _packetType, _minGas)
}

// SetMinGasToTransferAndStore is a paid mutator transaction binding the contract method 0x0b4cad4c.
//
// Solidity: function setMinGasToTransferAndStore(uint256 _minGasToTransferAndStore) returns()
func (_Minter *MinterTransactor) SetMinGasToTransferAndStore(opts *bind.TransactOpts, _minGasToTransferAndStore *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setMinGasToTransferAndStore", _minGasToTransferAndStore)
}

// SetMinGasToTransferAndStore is a paid mutator transaction binding the contract method 0x0b4cad4c.
//
// Solidity: function setMinGasToTransferAndStore(uint256 _minGasToTransferAndStore) returns()
func (_Minter *MinterSession) SetMinGasToTransferAndStore(_minGasToTransferAndStore *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetMinGasToTransferAndStore(&_Minter.TransactOpts, _minGasToTransferAndStore)
}

// SetMinGasToTransferAndStore is a paid mutator transaction binding the contract method 0x0b4cad4c.
//
// Solidity: function setMinGasToTransferAndStore(uint256 _minGasToTransferAndStore) returns()
func (_Minter *MinterTransactorSession) SetMinGasToTransferAndStore(_minGasToTransferAndStore *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetMinGasToTransferAndStore(&_Minter.TransactOpts, _minGasToTransferAndStore)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Minter *MinterTransactor) SetPayloadSizeLimit(opts *bind.TransactOpts, _dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setPayloadSizeLimit", _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Minter *MinterSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetPayloadSizeLimit(&_Minter.TransactOpts, _dstChainId, _size)
}

// SetPayloadSizeLimit is a paid mutator transaction binding the contract method 0x0df37483.
//
// Solidity: function setPayloadSizeLimit(uint16 _dstChainId, uint256 _size) returns()
func (_Minter *MinterTransactorSession) SetPayloadSizeLimit(_dstChainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.SetPayloadSizeLimit(&_Minter.TransactOpts, _dstChainId, _size)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Minter *MinterTransactor) SetPrecrime(opts *bind.TransactOpts, _precrime common.Address) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setPrecrime", _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Minter *MinterSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Minter.Contract.SetPrecrime(&_Minter.TransactOpts, _precrime)
}

// SetPrecrime is a paid mutator transaction binding the contract method 0xbaf3292d.
//
// Solidity: function setPrecrime(address _precrime) returns()
func (_Minter *MinterTransactorSession) SetPrecrime(_precrime common.Address) (*types.Transaction, error) {
	return _Minter.Contract.SetPrecrime(&_Minter.TransactOpts, _precrime)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Minter *MinterTransactor) SetReceiveVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setReceiveVersion", _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Minter *MinterSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Minter.Contract.SetReceiveVersion(&_Minter.TransactOpts, _version)
}

// SetReceiveVersion is a paid mutator transaction binding the contract method 0x10ddb137.
//
// Solidity: function setReceiveVersion(uint16 _version) returns()
func (_Minter *MinterTransactorSession) SetReceiveVersion(_version uint16) (*types.Transaction, error) {
	return _Minter.Contract.SetReceiveVersion(&_Minter.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Minter *MinterTransactor) SetSendVersion(opts *bind.TransactOpts, _version uint16) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setSendVersion", _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Minter *MinterSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Minter.Contract.SetSendVersion(&_Minter.TransactOpts, _version)
}

// SetSendVersion is a paid mutator transaction binding the contract method 0x07e0db17.
//
// Solidity: function setSendVersion(uint16 _version) returns()
func (_Minter *MinterTransactorSession) SetSendVersion(_version uint16) (*types.Transaction, error) {
	return _Minter.Contract.SetSendVersion(&_Minter.TransactOpts, _version)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Minter *MinterTransactor) SetTrustedRemote(opts *bind.TransactOpts, _remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setTrustedRemote", _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Minter *MinterSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Minter.Contract.SetTrustedRemote(&_Minter.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xeb8d72b7.
//
// Solidity: function setTrustedRemote(uint16 _remoteChainId, bytes _path) returns()
func (_Minter *MinterTransactorSession) SetTrustedRemote(_remoteChainId uint16, _path []byte) (*types.Transaction, error) {
	return _Minter.Contract.SetTrustedRemote(&_Minter.TransactOpts, _remoteChainId, _path)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Minter *MinterTransactor) SetTrustedRemoteAddress(opts *bind.TransactOpts, _remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "setTrustedRemoteAddress", _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Minter *MinterSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Minter.Contract.SetTrustedRemoteAddress(&_Minter.TransactOpts, _remoteChainId, _remoteAddress)
}

// SetTrustedRemoteAddress is a paid mutator transaction binding the contract method 0xa6c3d165.
//
// Solidity: function setTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress) returns()
func (_Minter *MinterTransactorSession) SetTrustedRemoteAddress(_remoteChainId uint16, _remoteAddress []byte) (*types.Transaction, error) {
	return _Minter.Contract.SetTrustedRemoteAddress(&_Minter.TransactOpts, _remoteChainId, _remoteAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Minter *MinterTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Minter *MinterSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.TransferFrom(&_Minter.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Minter *MinterTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Minter.Contract.TransferFrom(&_Minter.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Minter *MinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Minter *MinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Minter.Contract.TransferOwnership(&_Minter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Minter *MinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Minter.Contract.TransferOwnership(&_Minter.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Minter *MinterTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Minter.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Minter *MinterSession) Withdraw() (*types.Transaction, error) {
	return _Minter.Contract.Withdraw(&_Minter.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Minter *MinterTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Minter.Contract.Withdraw(&_Minter.TransactOpts)
}

// MinterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Minter contract.
type MinterApprovalIterator struct {
	Event *MinterApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterApproval represents a Approval event raised by the Minter contract.
type MinterApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Minter *MinterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MinterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MinterApprovalIterator{contract: _Minter.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Minter *MinterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MinterApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterApproval)
				if err := _Minter.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Minter *MinterFilterer) ParseApproval(log types.Log) (*MinterApproval, error) {
	event := new(MinterApproval)
	if err := _Minter.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Minter contract.
type MinterApprovalForAllIterator struct {
	Event *MinterApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterApprovalForAll represents a ApprovalForAll event raised by the Minter contract.
type MinterApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Minter *MinterFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MinterApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MinterApprovalForAllIterator{contract: _Minter.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Minter *MinterFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MinterApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterApprovalForAll)
				if err := _Minter.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Minter *MinterFilterer) ParseApprovalForAll(log types.Log) (*MinterApprovalForAll, error) {
	event := new(MinterApprovalForAll)
	if err := _Minter.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterCreditClearedIterator is returned from FilterCreditCleared and is used to iterate over the raw logs and unpacked data for CreditCleared events raised by the Minter contract.
type MinterCreditClearedIterator struct {
	Event *MinterCreditCleared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterCreditClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterCreditCleared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterCreditCleared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterCreditClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterCreditClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterCreditCleared represents a CreditCleared event raised by the Minter contract.
type MinterCreditCleared struct {
	HashedPayload [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreditCleared is a free log retrieval operation binding the contract event 0xd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba233.
//
// Solidity: event CreditCleared(bytes32 _hashedPayload)
func (_Minter *MinterFilterer) FilterCreditCleared(opts *bind.FilterOpts) (*MinterCreditClearedIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "CreditCleared")
	if err != nil {
		return nil, err
	}
	return &MinterCreditClearedIterator{contract: _Minter.contract, event: "CreditCleared", logs: logs, sub: sub}, nil
}

// WatchCreditCleared is a free log subscription operation binding the contract event 0xd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba233.
//
// Solidity: event CreditCleared(bytes32 _hashedPayload)
func (_Minter *MinterFilterer) WatchCreditCleared(opts *bind.WatchOpts, sink chan<- *MinterCreditCleared) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "CreditCleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterCreditCleared)
				if err := _Minter.contract.UnpackLog(event, "CreditCleared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreditCleared is a log parse operation binding the contract event 0xd7be02b8dd0d27bd0517a9cb4d7469ce27df4313821ae5ec1ff69acc594ba233.
//
// Solidity: event CreditCleared(bytes32 _hashedPayload)
func (_Minter *MinterFilterer) ParseCreditCleared(log types.Log) (*MinterCreditCleared, error) {
	event := new(MinterCreditCleared)
	if err := _Minter.contract.UnpackLog(event, "CreditCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterCreditStoredIterator is returned from FilterCreditStored and is used to iterate over the raw logs and unpacked data for CreditStored events raised by the Minter contract.
type MinterCreditStoredIterator struct {
	Event *MinterCreditStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterCreditStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterCreditStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterCreditStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterCreditStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterCreditStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterCreditStored represents a CreditStored event raised by the Minter contract.
type MinterCreditStored struct {
	HashedPayload [32]byte
	Payload       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreditStored is a free log retrieval operation binding the contract event 0x10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7.
//
// Solidity: event CreditStored(bytes32 _hashedPayload, bytes _payload)
func (_Minter *MinterFilterer) FilterCreditStored(opts *bind.FilterOpts) (*MinterCreditStoredIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "CreditStored")
	if err != nil {
		return nil, err
	}
	return &MinterCreditStoredIterator{contract: _Minter.contract, event: "CreditStored", logs: logs, sub: sub}, nil
}

// WatchCreditStored is a free log subscription operation binding the contract event 0x10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7.
//
// Solidity: event CreditStored(bytes32 _hashedPayload, bytes _payload)
func (_Minter *MinterFilterer) WatchCreditStored(opts *bind.WatchOpts, sink chan<- *MinterCreditStored) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "CreditStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterCreditStored)
				if err := _Minter.contract.UnpackLog(event, "CreditStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreditStored is a log parse operation binding the contract event 0x10e0b70d256bccc84b7027506978bd8b68984a870788b93b479def144c839ad7.
//
// Solidity: event CreditStored(bytes32 _hashedPayload, bytes _payload)
func (_Minter *MinterFilterer) ParseCreditStored(log types.Log) (*MinterCreditStored, error) {
	event := new(MinterCreditStored)
	if err := _Minter.contract.UnpackLog(event, "CreditStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterMessageFailedIterator is returned from FilterMessageFailed and is used to iterate over the raw logs and unpacked data for MessageFailed events raised by the Minter contract.
type MinterMessageFailedIterator struct {
	Event *MinterMessageFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterMessageFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterMessageFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterMessageFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterMessageFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterMessageFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterMessageFailed represents a MessageFailed event raised by the Minter contract.
type MinterMessageFailed struct {
	SrcChainId uint16
	SrcAddress []byte
	Nonce      uint64
	Payload    []byte
	Reason     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageFailed is a free log retrieval operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Minter *MinterFilterer) FilterMessageFailed(opts *bind.FilterOpts) (*MinterMessageFailedIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return &MinterMessageFailedIterator{contract: _Minter.contract, event: "MessageFailed", logs: logs, sub: sub}, nil
}

// WatchMessageFailed is a free log subscription operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Minter *MinterFilterer) WatchMessageFailed(opts *bind.WatchOpts, sink chan<- *MinterMessageFailed) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "MessageFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterMessageFailed)
				if err := _Minter.contract.UnpackLog(event, "MessageFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageFailed is a log parse operation binding the contract event 0xe183f33de2837795525b4792ca4cd60535bd77c53b7e7030060bfcf5734d6b0c.
//
// Solidity: event MessageFailed(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes _payload, bytes _reason)
func (_Minter *MinterFilterer) ParseMessageFailed(log types.Log) (*MinterMessageFailed, error) {
	event := new(MinterMessageFailed)
	if err := _Minter.contract.UnpackLog(event, "MessageFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Minter contract.
type MinterOwnershipTransferredIterator struct {
	Event *MinterOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterOwnershipTransferred represents a OwnershipTransferred event raised by the Minter contract.
type MinterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Minter *MinterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MinterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MinterOwnershipTransferredIterator{contract: _Minter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Minter *MinterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MinterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterOwnershipTransferred)
				if err := _Minter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Minter *MinterFilterer) ParseOwnershipTransferred(log types.Log) (*MinterOwnershipTransferred, error) {
	event := new(MinterOwnershipTransferred)
	if err := _Minter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterReceiveFromChainIterator is returned from FilterReceiveFromChain and is used to iterate over the raw logs and unpacked data for ReceiveFromChain events raised by the Minter contract.
type MinterReceiveFromChainIterator struct {
	Event *MinterReceiveFromChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterReceiveFromChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterReceiveFromChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterReceiveFromChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterReceiveFromChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterReceiveFromChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterReceiveFromChain represents a ReceiveFromChain event raised by the Minter contract.
type MinterReceiveFromChain struct {
	SrcChainId uint16
	SrcAddress common.Hash
	ToAddress  common.Address
	TokenIds   []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveFromChain is a free log retrieval operation binding the contract event 0x5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds)
func (_Minter *MinterFilterer) FilterReceiveFromChain(opts *bind.FilterOpts, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (*MinterReceiveFromChainIterator, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &MinterReceiveFromChainIterator{contract: _Minter.contract, event: "ReceiveFromChain", logs: logs, sub: sub}, nil
}

// WatchReceiveFromChain is a free log subscription operation binding the contract event 0x5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds)
func (_Minter *MinterFilterer) WatchReceiveFromChain(opts *bind.WatchOpts, sink chan<- *MinterReceiveFromChain, _srcChainId []uint16, _srcAddress [][]byte, _toAddress []common.Address) (event.Subscription, error) {

	var _srcChainIdRule []interface{}
	for _, _srcChainIdItem := range _srcChainId {
		_srcChainIdRule = append(_srcChainIdRule, _srcChainIdItem)
	}
	var _srcAddressRule []interface{}
	for _, _srcAddressItem := range _srcAddress {
		_srcAddressRule = append(_srcAddressRule, _srcAddressItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "ReceiveFromChain", _srcChainIdRule, _srcAddressRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterReceiveFromChain)
				if err := _Minter.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveFromChain is a log parse operation binding the contract event 0x5b821db8a46f8ecbe1941ba2f51cfeea9643268b56631f70d45e2a745d990265.
//
// Solidity: event ReceiveFromChain(uint16 indexed _srcChainId, bytes indexed _srcAddress, address indexed _toAddress, uint256[] _tokenIds)
func (_Minter *MinterFilterer) ParseReceiveFromChain(log types.Log) (*MinterReceiveFromChain, error) {
	event := new(MinterReceiveFromChain)
	if err := _Minter.contract.UnpackLog(event, "ReceiveFromChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterRetryMessageSuccessIterator is returned from FilterRetryMessageSuccess and is used to iterate over the raw logs and unpacked data for RetryMessageSuccess events raised by the Minter contract.
type MinterRetryMessageSuccessIterator struct {
	Event *MinterRetryMessageSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterRetryMessageSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRetryMessageSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterRetryMessageSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterRetryMessageSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRetryMessageSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRetryMessageSuccess represents a RetryMessageSuccess event raised by the Minter contract.
type MinterRetryMessageSuccess struct {
	SrcChainId  uint16
	SrcAddress  []byte
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRetryMessageSuccess is a free log retrieval operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Minter *MinterFilterer) FilterRetryMessageSuccess(opts *bind.FilterOpts) (*MinterRetryMessageSuccessIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return &MinterRetryMessageSuccessIterator{contract: _Minter.contract, event: "RetryMessageSuccess", logs: logs, sub: sub}, nil
}

// WatchRetryMessageSuccess is a free log subscription operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Minter *MinterFilterer) WatchRetryMessageSuccess(opts *bind.WatchOpts, sink chan<- *MinterRetryMessageSuccess) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "RetryMessageSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRetryMessageSuccess)
				if err := _Minter.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRetryMessageSuccess is a log parse operation binding the contract event 0xc264d91f3adc5588250e1551f547752ca0cfa8f6b530d243b9f9f4cab10ea8e5.
//
// Solidity: event RetryMessageSuccess(uint16 _srcChainId, bytes _srcAddress, uint64 _nonce, bytes32 _payloadHash)
func (_Minter *MinterFilterer) ParseRetryMessageSuccess(log types.Log) (*MinterRetryMessageSuccess, error) {
	event := new(MinterRetryMessageSuccess)
	if err := _Minter.contract.UnpackLog(event, "RetryMessageSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSendToChainIterator is returned from FilterSendToChain and is used to iterate over the raw logs and unpacked data for SendToChain events raised by the Minter contract.
type MinterSendToChainIterator struct {
	Event *MinterSendToChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSendToChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSendToChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSendToChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSendToChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSendToChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSendToChain represents a SendToChain event raised by the Minter contract.
type MinterSendToChain struct {
	DstChainId uint16
	From       common.Address
	ToAddress  common.Hash
	TokenIds   []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendToChain is a free log retrieval operation binding the contract event 0xe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds)
func (_Minter *MinterFilterer) FilterSendToChain(opts *bind.FilterOpts, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (*MinterSendToChainIterator, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SendToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return &MinterSendToChainIterator{contract: _Minter.contract, event: "SendToChain", logs: logs, sub: sub}, nil
}

// WatchSendToChain is a free log subscription operation binding the contract event 0xe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds)
func (_Minter *MinterFilterer) WatchSendToChain(opts *bind.WatchOpts, sink chan<- *MinterSendToChain, _dstChainId []uint16, _from []common.Address, _toAddress [][]byte) (event.Subscription, error) {

	var _dstChainIdRule []interface{}
	for _, _dstChainIdItem := range _dstChainId {
		_dstChainIdRule = append(_dstChainIdRule, _dstChainIdItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toAddressRule []interface{}
	for _, _toAddressItem := range _toAddress {
		_toAddressRule = append(_toAddressRule, _toAddressItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SendToChain", _dstChainIdRule, _fromRule, _toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSendToChain)
				if err := _Minter.contract.UnpackLog(event, "SendToChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendToChain is a log parse operation binding the contract event 0xe1b87c47fdeb4f9cbadbca9df3af7aba453bb6e501075d0440d88125b711522a.
//
// Solidity: event SendToChain(uint16 indexed _dstChainId, address indexed _from, bytes indexed _toAddress, uint256[] _tokenIds)
func (_Minter *MinterFilterer) ParseSendToChain(log types.Log) (*MinterSendToChain, error) {
	event := new(MinterSendToChain)
	if err := _Minter.contract.UnpackLog(event, "SendToChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetDstChainIdToBatchLimitIterator is returned from FilterSetDstChainIdToBatchLimit and is used to iterate over the raw logs and unpacked data for SetDstChainIdToBatchLimit events raised by the Minter contract.
type MinterSetDstChainIdToBatchLimitIterator struct {
	Event *MinterSetDstChainIdToBatchLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetDstChainIdToBatchLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetDstChainIdToBatchLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetDstChainIdToBatchLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetDstChainIdToBatchLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetDstChainIdToBatchLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetDstChainIdToBatchLimit represents a SetDstChainIdToBatchLimit event raised by the Minter contract.
type MinterSetDstChainIdToBatchLimit struct {
	DstChainId             uint16
	DstChainIdToBatchLimit *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDstChainIdToBatchLimit is a free log retrieval operation binding the contract event 0x7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d.
//
// Solidity: event SetDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit)
func (_Minter *MinterFilterer) FilterSetDstChainIdToBatchLimit(opts *bind.FilterOpts) (*MinterSetDstChainIdToBatchLimitIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetDstChainIdToBatchLimit")
	if err != nil {
		return nil, err
	}
	return &MinterSetDstChainIdToBatchLimitIterator{contract: _Minter.contract, event: "SetDstChainIdToBatchLimit", logs: logs, sub: sub}, nil
}

// WatchSetDstChainIdToBatchLimit is a free log subscription operation binding the contract event 0x7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d.
//
// Solidity: event SetDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit)
func (_Minter *MinterFilterer) WatchSetDstChainIdToBatchLimit(opts *bind.WatchOpts, sink chan<- *MinterSetDstChainIdToBatchLimit) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetDstChainIdToBatchLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetDstChainIdToBatchLimit)
				if err := _Minter.contract.UnpackLog(event, "SetDstChainIdToBatchLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDstChainIdToBatchLimit is a log parse operation binding the contract event 0x7315f7654d594ead24a30160ed9ba2d23247f543016b918343591e93d7afdb6d.
//
// Solidity: event SetDstChainIdToBatchLimit(uint16 _dstChainId, uint256 _dstChainIdToBatchLimit)
func (_Minter *MinterFilterer) ParseSetDstChainIdToBatchLimit(log types.Log) (*MinterSetDstChainIdToBatchLimit, error) {
	event := new(MinterSetDstChainIdToBatchLimit)
	if err := _Minter.contract.UnpackLog(event, "SetDstChainIdToBatchLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetDstChainIdToTransferGasIterator is returned from FilterSetDstChainIdToTransferGas and is used to iterate over the raw logs and unpacked data for SetDstChainIdToTransferGas events raised by the Minter contract.
type MinterSetDstChainIdToTransferGasIterator struct {
	Event *MinterSetDstChainIdToTransferGas // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetDstChainIdToTransferGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetDstChainIdToTransferGas)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetDstChainIdToTransferGas)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetDstChainIdToTransferGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetDstChainIdToTransferGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetDstChainIdToTransferGas represents a SetDstChainIdToTransferGas event raised by the Minter contract.
type MinterSetDstChainIdToTransferGas struct {
	DstChainId              uint16
	DstChainIdToTransferGas *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSetDstChainIdToTransferGas is a free log retrieval operation binding the contract event 0xc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb.
//
// Solidity: event SetDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas)
func (_Minter *MinterFilterer) FilterSetDstChainIdToTransferGas(opts *bind.FilterOpts) (*MinterSetDstChainIdToTransferGasIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetDstChainIdToTransferGas")
	if err != nil {
		return nil, err
	}
	return &MinterSetDstChainIdToTransferGasIterator{contract: _Minter.contract, event: "SetDstChainIdToTransferGas", logs: logs, sub: sub}, nil
}

// WatchSetDstChainIdToTransferGas is a free log subscription operation binding the contract event 0xc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb.
//
// Solidity: event SetDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas)
func (_Minter *MinterFilterer) WatchSetDstChainIdToTransferGas(opts *bind.WatchOpts, sink chan<- *MinterSetDstChainIdToTransferGas) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetDstChainIdToTransferGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetDstChainIdToTransferGas)
				if err := _Minter.contract.UnpackLog(event, "SetDstChainIdToTransferGas", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDstChainIdToTransferGas is a log parse operation binding the contract event 0xc46df2983228ac2d9754e94a0d565e6671665dc8ad38602bc8e544f0685a29fb.
//
// Solidity: event SetDstChainIdToTransferGas(uint16 _dstChainId, uint256 _dstChainIdToTransferGas)
func (_Minter *MinterFilterer) ParseSetDstChainIdToTransferGas(log types.Log) (*MinterSetDstChainIdToTransferGas, error) {
	event := new(MinterSetDstChainIdToTransferGas)
	if err := _Minter.contract.UnpackLog(event, "SetDstChainIdToTransferGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetMinDstGasIterator is returned from FilterSetMinDstGas and is used to iterate over the raw logs and unpacked data for SetMinDstGas events raised by the Minter contract.
type MinterSetMinDstGasIterator struct {
	Event *MinterSetMinDstGas // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetMinDstGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetMinDstGas)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetMinDstGas)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetMinDstGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetMinDstGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetMinDstGas represents a SetMinDstGas event raised by the Minter contract.
type MinterSetMinDstGas struct {
	DstChainId uint16
	Type       uint16
	MinDstGas  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMinDstGas is a free log retrieval operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Minter *MinterFilterer) FilterSetMinDstGas(opts *bind.FilterOpts) (*MinterSetMinDstGasIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return &MinterSetMinDstGasIterator{contract: _Minter.contract, event: "SetMinDstGas", logs: logs, sub: sub}, nil
}

// WatchSetMinDstGas is a free log subscription operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Minter *MinterFilterer) WatchSetMinDstGas(opts *bind.WatchOpts, sink chan<- *MinterSetMinDstGas) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetMinDstGas")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetMinDstGas)
				if err := _Minter.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMinDstGas is a log parse operation binding the contract event 0x9d5c7c0b934da8fefa9c7760c98383778a12dfbfc0c3b3106518f43fb9508ac0.
//
// Solidity: event SetMinDstGas(uint16 _dstChainId, uint16 _type, uint256 _minDstGas)
func (_Minter *MinterFilterer) ParseSetMinDstGas(log types.Log) (*MinterSetMinDstGas, error) {
	event := new(MinterSetMinDstGas)
	if err := _Minter.contract.UnpackLog(event, "SetMinDstGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetMinGasToTransferAndStoreIterator is returned from FilterSetMinGasToTransferAndStore and is used to iterate over the raw logs and unpacked data for SetMinGasToTransferAndStore events raised by the Minter contract.
type MinterSetMinGasToTransferAndStoreIterator struct {
	Event *MinterSetMinGasToTransferAndStore // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetMinGasToTransferAndStoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetMinGasToTransferAndStore)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetMinGasToTransferAndStore)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetMinGasToTransferAndStoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetMinGasToTransferAndStoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetMinGasToTransferAndStore represents a SetMinGasToTransferAndStore event raised by the Minter contract.
type MinterSetMinGasToTransferAndStore struct {
	MinGasToTransferAndStore *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetMinGasToTransferAndStore is a free log retrieval operation binding the contract event 0xfebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d.
//
// Solidity: event SetMinGasToTransferAndStore(uint256 _minGasToTransferAndStore)
func (_Minter *MinterFilterer) FilterSetMinGasToTransferAndStore(opts *bind.FilterOpts) (*MinterSetMinGasToTransferAndStoreIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetMinGasToTransferAndStore")
	if err != nil {
		return nil, err
	}
	return &MinterSetMinGasToTransferAndStoreIterator{contract: _Minter.contract, event: "SetMinGasToTransferAndStore", logs: logs, sub: sub}, nil
}

// WatchSetMinGasToTransferAndStore is a free log subscription operation binding the contract event 0xfebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d.
//
// Solidity: event SetMinGasToTransferAndStore(uint256 _minGasToTransferAndStore)
func (_Minter *MinterFilterer) WatchSetMinGasToTransferAndStore(opts *bind.WatchOpts, sink chan<- *MinterSetMinGasToTransferAndStore) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetMinGasToTransferAndStore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetMinGasToTransferAndStore)
				if err := _Minter.contract.UnpackLog(event, "SetMinGasToTransferAndStore", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMinGasToTransferAndStore is a log parse operation binding the contract event 0xfebbc4f8bb9ec2313950c718d43123124b15778efda4c1f1d529de2995b4f34d.
//
// Solidity: event SetMinGasToTransferAndStore(uint256 _minGasToTransferAndStore)
func (_Minter *MinterFilterer) ParseSetMinGasToTransferAndStore(log types.Log) (*MinterSetMinGasToTransferAndStore, error) {
	event := new(MinterSetMinGasToTransferAndStore)
	if err := _Minter.contract.UnpackLog(event, "SetMinGasToTransferAndStore", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetPrecrimeIterator is returned from FilterSetPrecrime and is used to iterate over the raw logs and unpacked data for SetPrecrime events raised by the Minter contract.
type MinterSetPrecrimeIterator struct {
	Event *MinterSetPrecrime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetPrecrimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetPrecrime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetPrecrime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetPrecrimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetPrecrimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetPrecrime represents a SetPrecrime event raised by the Minter contract.
type MinterSetPrecrime struct {
	Precrime common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPrecrime is a free log retrieval operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Minter *MinterFilterer) FilterSetPrecrime(opts *bind.FilterOpts) (*MinterSetPrecrimeIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return &MinterSetPrecrimeIterator{contract: _Minter.contract, event: "SetPrecrime", logs: logs, sub: sub}, nil
}

// WatchSetPrecrime is a free log subscription operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Minter *MinterFilterer) WatchSetPrecrime(opts *bind.WatchOpts, sink chan<- *MinterSetPrecrime) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetPrecrime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetPrecrime)
				if err := _Minter.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPrecrime is a log parse operation binding the contract event 0x5db758e995a17ec1ad84bdef7e8c3293a0bd6179bcce400dff5d4c3d87db726b.
//
// Solidity: event SetPrecrime(address precrime)
func (_Minter *MinterFilterer) ParseSetPrecrime(log types.Log) (*MinterSetPrecrime, error) {
	event := new(MinterSetPrecrime)
	if err := _Minter.contract.UnpackLog(event, "SetPrecrime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the Minter contract.
type MinterSetTrustedRemoteIterator struct {
	Event *MinterSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetTrustedRemote represents a SetTrustedRemote event raised by the Minter contract.
type MinterSetTrustedRemote struct {
	RemoteChainId uint16
	Path          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Minter *MinterFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*MinterSetTrustedRemoteIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &MinterSetTrustedRemoteIterator{contract: _Minter.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Minter *MinterFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *MinterSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetTrustedRemote)
				if err := _Minter.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0xfa41487ad5d6728f0b19276fa1eddc16558578f5109fc39d2dc33c3230470dab.
//
// Solidity: event SetTrustedRemote(uint16 _remoteChainId, bytes _path)
func (_Minter *MinterFilterer) ParseSetTrustedRemote(log types.Log) (*MinterSetTrustedRemote, error) {
	event := new(MinterSetTrustedRemote)
	if err := _Minter.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterSetTrustedRemoteAddressIterator is returned from FilterSetTrustedRemoteAddress and is used to iterate over the raw logs and unpacked data for SetTrustedRemoteAddress events raised by the Minter contract.
type MinterSetTrustedRemoteAddressIterator struct {
	Event *MinterSetTrustedRemoteAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterSetTrustedRemoteAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterSetTrustedRemoteAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterSetTrustedRemoteAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterSetTrustedRemoteAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterSetTrustedRemoteAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterSetTrustedRemoteAddress represents a SetTrustedRemoteAddress event raised by the Minter contract.
type MinterSetTrustedRemoteAddress struct {
	RemoteChainId uint16
	RemoteAddress []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemoteAddress is a free log retrieval operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Minter *MinterFilterer) FilterSetTrustedRemoteAddress(opts *bind.FilterOpts) (*MinterSetTrustedRemoteAddressIterator, error) {

	logs, sub, err := _Minter.contract.FilterLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return &MinterSetTrustedRemoteAddressIterator{contract: _Minter.contract, event: "SetTrustedRemoteAddress", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemoteAddress is a free log subscription operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Minter *MinterFilterer) WatchSetTrustedRemoteAddress(opts *bind.WatchOpts, sink chan<- *MinterSetTrustedRemoteAddress) (event.Subscription, error) {

	logs, sub, err := _Minter.contract.WatchLogs(opts, "SetTrustedRemoteAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterSetTrustedRemoteAddress)
				if err := _Minter.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemoteAddress is a log parse operation binding the contract event 0x8c0400cfe2d1199b1a725c78960bcc2a344d869b80590d0f2bd005db15a572ce.
//
// Solidity: event SetTrustedRemoteAddress(uint16 _remoteChainId, bytes _remoteAddress)
func (_Minter *MinterFilterer) ParseSetTrustedRemoteAddress(log types.Log) (*MinterSetTrustedRemoteAddress, error) {
	event := new(MinterSetTrustedRemoteAddress)
	if err := _Minter.contract.UnpackLog(event, "SetTrustedRemoteAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Minter contract.
type MinterTransferIterator struct {
	Event *MinterTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MinterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MinterTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MinterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterTransfer represents a Transfer event raised by the Minter contract.
type MinterTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Minter *MinterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MinterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Minter.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MinterTransferIterator{contract: _Minter.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Minter *MinterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MinterTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Minter.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterTransfer)
				if err := _Minter.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Minter *MinterFilterer) ParseTransfer(log types.Log) (*MinterTransfer, error) {
	event := new(MinterTransfer)
	if err := _Minter.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
