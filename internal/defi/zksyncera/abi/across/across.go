// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package across

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

// SpokePoolInterfaceRelayerRefundLeaf is an auto generated low-level Go binding around an user-defined struct.
type SpokePoolInterfaceRelayerRefundLeaf struct {
	AmountToReturn  *big.Int
	ChainId         *big.Int
	RefundAmounts   []*big.Int
	LeafId          uint32
	L2TokenAddress  common.Address
	RefundAddresses []common.Address
}

// SpokePoolRelayExecutionInfo is an auto generated low-level Go binding around an user-defined struct.
type SpokePoolRelayExecutionInfo struct {
	Recipient           common.Address
	Message             []byte
	RelayerFeePct       int64
	IsSlowRelay         bool
	PayoutAdjustmentPct *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rootBundleId\",\"type\":\"uint256\"}],\"name\":\"EmergencyDeleteRootBundle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"EnabledDepositRoute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToReturn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"refundAmounts\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l2TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"refundAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ExecutedRelayerRefundRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFilledAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repaymentChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"isSlowRelay\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"payoutAdjustmentPct\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structSpokePool.RelayExecutionInfo\",\"name\":\"updatableRelayData\",\"type\":\"tuple\"}],\"name\":\"FilledRelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quoteTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PausedDeposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PausedFills\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousIdenticalRequests\",\"type\":\"uint256\"}],\"name\":\"RefundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relayerRefundRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"slowRelayRoot\",\"type\":\"bytes32\"}],\"name\":\"RelayedRootBundle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"newRelayerFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"updatedMessage\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"depositorSignature\",\"type\":\"bytes\"}],\"name\":\"RequestedSpeedUpDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newBuffer\",\"type\":\"uint32\"}],\"name\":\"SetDepositQuoteTimeBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHubPool\",\"type\":\"address\"}],\"name\":\"SetHubPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"SetXDomainAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Bridge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldErc20Bridge\",\"type\":\"address\"}],\"name\":\"SetZkBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToReturn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokensBridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfTokensBridged\",\"type\":\"uint256\"}],\"name\":\"ZkSyncTokensBridged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOW_FILL_MAX_TOKENS_TO_SEND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_DEPOSIT_DETAILS_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_initialDepositId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_crossDomainAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hubPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappedNativeTokenAddress\",\"type\":\"address\"}],\"name\":\"__SpokePool_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossDomainAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"quoteTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositQuoteTimeBuffer\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rootBundleId\",\"type\":\"uint256\"}],\"name\":\"emergencyDeleteRootBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"enabledDepositRoutes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountToReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"refundAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32\",\"name\":\"leafId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"l2TokenAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"refundAddresses\",\"type\":\"address[]\"}],\"internalType\":\"structSpokePoolInterface.RelayerRefundLeaf\",\"name\":\"relayerRefundLeaf\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"executeRelayerRefundLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalRelayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int256\",\"name\":\"payoutAdjustment\",\"type\":\"int256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"executeSlowRelayLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fillCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensToSend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repaymentChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"fillRelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"updatedRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensToSend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repaymentChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"updatedRelayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"updatedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositorSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"fillRelayWithUpdatedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_initialDepositId\",\"type\":\"uint32\"},{\"internalType\":\"contractZkBridgeLike\",\"name\":\"_zkErc20Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossDomainAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hubPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wethAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Eth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfDeposits\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"pauseDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"pauseFills\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausedDeposits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausedFills\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"refundsRequested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"relayFills\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"relayerRefundRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slowRelayRoot\",\"type\":\"bytes32\"}],\"name\":\"relayRootBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"refundToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fillBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rootBundles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"slowRelayRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"relayerRefundRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCrossDomainAdmin\",\"type\":\"address\"}],\"name\":\"setCrossDomainAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newDepositQuoteTimeBuffer\",\"type\":\"uint32\"}],\"name\":\"setDepositQuoteTimeBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setEnableRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHubPool\",\"type\":\"address\"}],\"name\":\"setHubPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractZkBridgeLike\",\"name\":\"_zkErc20Bridge\",\"type\":\"address\"}],\"name\":\"setZkBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"updatedRelayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"updatedRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"updatedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositorSignature\",\"type\":\"bytes\"}],\"name\":\"speedUpDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractWETH9Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkErc20Bridge\",\"outputs\":[{\"internalType\":\"contractZkBridgeLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_Storage *StorageCaller) MAXTRANSFERSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAX_TRANSFER_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_Storage *StorageSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _Storage.Contract.MAXTRANSFERSIZE(&_Storage.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_Storage *StorageCallerSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _Storage.Contract.MAXTRANSFERSIZE(&_Storage.CallOpts)
}

// SLOWFILLMAXTOKENSTOSEND is a free data retrieval call binding the contract method 0xb5e1bf5f.
//
// Solidity: function SLOW_FILL_MAX_TOKENS_TO_SEND() view returns(uint256)
func (_Storage *StorageCaller) SLOWFILLMAXTOKENSTOSEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "SLOW_FILL_MAX_TOKENS_TO_SEND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLOWFILLMAXTOKENSTOSEND is a free data retrieval call binding the contract method 0xb5e1bf5f.
//
// Solidity: function SLOW_FILL_MAX_TOKENS_TO_SEND() view returns(uint256)
func (_Storage *StorageSession) SLOWFILLMAXTOKENSTOSEND() (*big.Int, error) {
	return _Storage.Contract.SLOWFILLMAXTOKENSTOSEND(&_Storage.CallOpts)
}

// SLOWFILLMAXTOKENSTOSEND is a free data retrieval call binding the contract method 0xb5e1bf5f.
//
// Solidity: function SLOW_FILL_MAX_TOKENS_TO_SEND() view returns(uint256)
func (_Storage *StorageCallerSession) SLOWFILLMAXTOKENSTOSEND() (*big.Int, error) {
	return _Storage.Contract.SLOWFILLMAXTOKENSTOSEND(&_Storage.CallOpts)
}

// UPDATEDEPOSITDETAILSHASH is a free data retrieval call binding the contract method 0xa78e4b60.
//
// Solidity: function UPDATE_DEPOSIT_DETAILS_HASH() view returns(bytes32)
func (_Storage *StorageCaller) UPDATEDEPOSITDETAILSHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "UPDATE_DEPOSIT_DETAILS_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATEDEPOSITDETAILSHASH is a free data retrieval call binding the contract method 0xa78e4b60.
//
// Solidity: function UPDATE_DEPOSIT_DETAILS_HASH() view returns(bytes32)
func (_Storage *StorageSession) UPDATEDEPOSITDETAILSHASH() ([32]byte, error) {
	return _Storage.Contract.UPDATEDEPOSITDETAILSHASH(&_Storage.CallOpts)
}

// UPDATEDEPOSITDETAILSHASH is a free data retrieval call binding the contract method 0xa78e4b60.
//
// Solidity: function UPDATE_DEPOSIT_DETAILS_HASH() view returns(bytes32)
func (_Storage *StorageCallerSession) UPDATEDEPOSITDETAILSHASH() ([32]byte, error) {
	return _Storage.Contract.UPDATEDEPOSITDETAILSHASH(&_Storage.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Storage *StorageCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Storage *StorageSession) ChainId() (*big.Int, error) {
	return _Storage.Contract.ChainId(&_Storage.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Storage *StorageCallerSession) ChainId() (*big.Int, error) {
	return _Storage.Contract.ChainId(&_Storage.CallOpts)
}

// CrossDomainAdmin is a free data retrieval call binding the contract method 0x5285e058.
//
// Solidity: function crossDomainAdmin() view returns(address)
func (_Storage *StorageCaller) CrossDomainAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "crossDomainAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossDomainAdmin is a free data retrieval call binding the contract method 0x5285e058.
//
// Solidity: function crossDomainAdmin() view returns(address)
func (_Storage *StorageSession) CrossDomainAdmin() (common.Address, error) {
	return _Storage.Contract.CrossDomainAdmin(&_Storage.CallOpts)
}

// CrossDomainAdmin is a free data retrieval call binding the contract method 0x5285e058.
//
// Solidity: function crossDomainAdmin() view returns(address)
func (_Storage *StorageCallerSession) CrossDomainAdmin() (common.Address, error) {
	return _Storage.Contract.CrossDomainAdmin(&_Storage.CallOpts)
}

// DepositCounter is a free data retrieval call binding the contract method 0xc0a8bdd1.
//
// Solidity: function depositCounter(address ) view returns(uint256)
func (_Storage *StorageCaller) DepositCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "depositCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCounter is a free data retrieval call binding the contract method 0xc0a8bdd1.
//
// Solidity: function depositCounter(address ) view returns(uint256)
func (_Storage *StorageSession) DepositCounter(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.DepositCounter(&_Storage.CallOpts, arg0)
}

// DepositCounter is a free data retrieval call binding the contract method 0xc0a8bdd1.
//
// Solidity: function depositCounter(address ) view returns(uint256)
func (_Storage *StorageCallerSession) DepositCounter(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.DepositCounter(&_Storage.CallOpts, arg0)
}

// DepositQuoteTimeBuffer is a free data retrieval call binding the contract method 0x57f6dcb8.
//
// Solidity: function depositQuoteTimeBuffer() view returns(uint32)
func (_Storage *StorageCaller) DepositQuoteTimeBuffer(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "depositQuoteTimeBuffer")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DepositQuoteTimeBuffer is a free data retrieval call binding the contract method 0x57f6dcb8.
//
// Solidity: function depositQuoteTimeBuffer() view returns(uint32)
func (_Storage *StorageSession) DepositQuoteTimeBuffer() (uint32, error) {
	return _Storage.Contract.DepositQuoteTimeBuffer(&_Storage.CallOpts)
}

// DepositQuoteTimeBuffer is a free data retrieval call binding the contract method 0x57f6dcb8.
//
// Solidity: function depositQuoteTimeBuffer() view returns(uint32)
func (_Storage *StorageCallerSession) DepositQuoteTimeBuffer() (uint32, error) {
	return _Storage.Contract.DepositQuoteTimeBuffer(&_Storage.CallOpts)
}

// EnabledDepositRoutes is a free data retrieval call binding the contract method 0x5249fef1.
//
// Solidity: function enabledDepositRoutes(address , uint256 ) view returns(bool)
func (_Storage *StorageCaller) EnabledDepositRoutes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "enabledDepositRoutes", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnabledDepositRoutes is a free data retrieval call binding the contract method 0x5249fef1.
//
// Solidity: function enabledDepositRoutes(address , uint256 ) view returns(bool)
func (_Storage *StorageSession) EnabledDepositRoutes(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Storage.Contract.EnabledDepositRoutes(&_Storage.CallOpts, arg0, arg1)
}

// EnabledDepositRoutes is a free data retrieval call binding the contract method 0x5249fef1.
//
// Solidity: function enabledDepositRoutes(address , uint256 ) view returns(bool)
func (_Storage *StorageCallerSession) EnabledDepositRoutes(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Storage.Contract.EnabledDepositRoutes(&_Storage.CallOpts, arg0, arg1)
}

// FillCounter is a free data retrieval call binding the contract method 0x872af6ea.
//
// Solidity: function fillCounter(address ) view returns(uint256)
func (_Storage *StorageCaller) FillCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "fillCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillCounter is a free data retrieval call binding the contract method 0x872af6ea.
//
// Solidity: function fillCounter(address ) view returns(uint256)
func (_Storage *StorageSession) FillCounter(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.FillCounter(&_Storage.CallOpts, arg0)
}

// FillCounter is a free data retrieval call binding the contract method 0x872af6ea.
//
// Solidity: function fillCounter(address ) view returns(uint256)
func (_Storage *StorageCallerSession) FillCounter(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.FillCounter(&_Storage.CallOpts, arg0)
}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Storage *StorageCaller) GetCurrentTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCurrentTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Storage *StorageSession) GetCurrentTime() (*big.Int, error) {
	return _Storage.Contract.GetCurrentTime(&_Storage.CallOpts)
}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Storage *StorageCallerSession) GetCurrentTime() (*big.Int, error) {
	return _Storage.Contract.GetCurrentTime(&_Storage.CallOpts)
}

// HubPool is a free data retrieval call binding the contract method 0xe1904402.
//
// Solidity: function hubPool() view returns(address)
func (_Storage *StorageCaller) HubPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hubPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HubPool is a free data retrieval call binding the contract method 0xe1904402.
//
// Solidity: function hubPool() view returns(address)
func (_Storage *StorageSession) HubPool() (common.Address, error) {
	return _Storage.Contract.HubPool(&_Storage.CallOpts)
}

// HubPool is a free data retrieval call binding the contract method 0xe1904402.
//
// Solidity: function hubPool() view returns(address)
func (_Storage *StorageCallerSession) HubPool() (common.Address, error) {
	return _Storage.Contract.HubPool(&_Storage.CallOpts)
}

// L2Eth is a free data retrieval call binding the contract method 0xe3229211.
//
// Solidity: function l2Eth() view returns(address)
func (_Storage *StorageCaller) L2Eth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "l2Eth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Eth is a free data retrieval call binding the contract method 0xe3229211.
//
// Solidity: function l2Eth() view returns(address)
func (_Storage *StorageSession) L2Eth() (common.Address, error) {
	return _Storage.Contract.L2Eth(&_Storage.CallOpts)
}

// L2Eth is a free data retrieval call binding the contract method 0xe3229211.
//
// Solidity: function l2Eth() view returns(address)
func (_Storage *StorageCallerSession) L2Eth() (common.Address, error) {
	return _Storage.Contract.L2Eth(&_Storage.CallOpts)
}

// NumberOfDeposits is a free data retrieval call binding the contract method 0xa1244c67.
//
// Solidity: function numberOfDeposits() view returns(uint32)
func (_Storage *StorageCaller) NumberOfDeposits(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "numberOfDeposits")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumberOfDeposits is a free data retrieval call binding the contract method 0xa1244c67.
//
// Solidity: function numberOfDeposits() view returns(uint32)
func (_Storage *StorageSession) NumberOfDeposits() (uint32, error) {
	return _Storage.Contract.NumberOfDeposits(&_Storage.CallOpts)
}

// NumberOfDeposits is a free data retrieval call binding the contract method 0xa1244c67.
//
// Solidity: function numberOfDeposits() view returns(uint32)
func (_Storage *StorageCallerSession) NumberOfDeposits() (uint32, error) {
	return _Storage.Contract.NumberOfDeposits(&_Storage.CallOpts)
}

// PausedDeposits is a free data retrieval call binding the contract method 0x6068d6cb.
//
// Solidity: function pausedDeposits() view returns(bool)
func (_Storage *StorageCaller) PausedDeposits(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pausedDeposits")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausedDeposits is a free data retrieval call binding the contract method 0x6068d6cb.
//
// Solidity: function pausedDeposits() view returns(bool)
func (_Storage *StorageSession) PausedDeposits() (bool, error) {
	return _Storage.Contract.PausedDeposits(&_Storage.CallOpts)
}

// PausedDeposits is a free data retrieval call binding the contract method 0x6068d6cb.
//
// Solidity: function pausedDeposits() view returns(bool)
func (_Storage *StorageCallerSession) PausedDeposits() (bool, error) {
	return _Storage.Contract.PausedDeposits(&_Storage.CallOpts)
}

// PausedFills is a free data retrieval call binding the contract method 0xdda52113.
//
// Solidity: function pausedFills() view returns(bool)
func (_Storage *StorageCaller) PausedFills(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pausedFills")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausedFills is a free data retrieval call binding the contract method 0xdda52113.
//
// Solidity: function pausedFills() view returns(bool)
func (_Storage *StorageSession) PausedFills() (bool, error) {
	return _Storage.Contract.PausedFills(&_Storage.CallOpts)
}

// PausedFills is a free data retrieval call binding the contract method 0xdda52113.
//
// Solidity: function pausedFills() view returns(bool)
func (_Storage *StorageCallerSession) PausedFills() (bool, error) {
	return _Storage.Contract.PausedFills(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// RefundsRequested is a free data retrieval call binding the contract method 0x23cd9a47.
//
// Solidity: function refundsRequested(bytes32 ) view returns(uint256)
func (_Storage *StorageCaller) RefundsRequested(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "refundsRequested", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundsRequested is a free data retrieval call binding the contract method 0x23cd9a47.
//
// Solidity: function refundsRequested(bytes32 ) view returns(uint256)
func (_Storage *StorageSession) RefundsRequested(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.RefundsRequested(&_Storage.CallOpts, arg0)
}

// RefundsRequested is a free data retrieval call binding the contract method 0x23cd9a47.
//
// Solidity: function refundsRequested(bytes32 ) view returns(uint256)
func (_Storage *StorageCallerSession) RefundsRequested(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.RefundsRequested(&_Storage.CallOpts, arg0)
}

// RelayFills is a free data retrieval call binding the contract method 0xf06850f6.
//
// Solidity: function relayFills(bytes32 ) view returns(uint256)
func (_Storage *StorageCaller) RelayFills(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "relayFills", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayFills is a free data retrieval call binding the contract method 0xf06850f6.
//
// Solidity: function relayFills(bytes32 ) view returns(uint256)
func (_Storage *StorageSession) RelayFills(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.RelayFills(&_Storage.CallOpts, arg0)
}

// RelayFills is a free data retrieval call binding the contract method 0xf06850f6.
//
// Solidity: function relayFills(bytes32 ) view returns(uint256)
func (_Storage *StorageCallerSession) RelayFills(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.RelayFills(&_Storage.CallOpts, arg0)
}

// RootBundles is a free data retrieval call binding the contract method 0xee2a53f8.
//
// Solidity: function rootBundles(uint256 ) view returns(bytes32 slowRelayRoot, bytes32 relayerRefundRoot)
func (_Storage *StorageCaller) RootBundles(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SlowRelayRoot     [32]byte
	RelayerRefundRoot [32]byte
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "rootBundles", arg0)

	outstruct := new(struct {
		SlowRelayRoot     [32]byte
		RelayerRefundRoot [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlowRelayRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RelayerRefundRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RootBundles is a free data retrieval call binding the contract method 0xee2a53f8.
//
// Solidity: function rootBundles(uint256 ) view returns(bytes32 slowRelayRoot, bytes32 relayerRefundRoot)
func (_Storage *StorageSession) RootBundles(arg0 *big.Int) (struct {
	SlowRelayRoot     [32]byte
	RelayerRefundRoot [32]byte
}, error) {
	return _Storage.Contract.RootBundles(&_Storage.CallOpts, arg0)
}

// RootBundles is a free data retrieval call binding the contract method 0xee2a53f8.
//
// Solidity: function rootBundles(uint256 ) view returns(bytes32 slowRelayRoot, bytes32 relayerRefundRoot)
func (_Storage *StorageCallerSession) RootBundles(arg0 *big.Int) (struct {
	SlowRelayRoot     [32]byte
	RelayerRefundRoot [32]byte
}, error) {
	return _Storage.Contract.RootBundles(&_Storage.CallOpts, arg0)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Storage *StorageCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Storage *StorageSession) WrappedNativeToken() (common.Address, error) {
	return _Storage.Contract.WrappedNativeToken(&_Storage.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Storage *StorageCallerSession) WrappedNativeToken() (common.Address, error) {
	return _Storage.Contract.WrappedNativeToken(&_Storage.CallOpts)
}

// ZkErc20Bridge is a free data retrieval call binding the contract method 0xbb3e04b5.
//
// Solidity: function zkErc20Bridge() view returns(address)
func (_Storage *StorageCaller) ZkErc20Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "zkErc20Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkErc20Bridge is a free data retrieval call binding the contract method 0xbb3e04b5.
//
// Solidity: function zkErc20Bridge() view returns(address)
func (_Storage *StorageSession) ZkErc20Bridge() (common.Address, error) {
	return _Storage.Contract.ZkErc20Bridge(&_Storage.CallOpts)
}

// ZkErc20Bridge is a free data retrieval call binding the contract method 0xbb3e04b5.
//
// Solidity: function zkErc20Bridge() view returns(address)
func (_Storage *StorageCallerSession) ZkErc20Bridge() (common.Address, error) {
	return _Storage.Contract.ZkErc20Bridge(&_Storage.CallOpts)
}

// SpokePoolInit is a paid mutator transaction binding the contract method 0x3ddf5059.
//
// Solidity: function __SpokePool_init(uint32 _initialDepositId, address _crossDomainAdmin, address _hubPool, address _wrappedNativeTokenAddress) returns()
func (_Storage *StorageTransactor) SpokePoolInit(opts *bind.TransactOpts, _initialDepositId uint32, _crossDomainAdmin common.Address, _hubPool common.Address, _wrappedNativeTokenAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "__SpokePool_init", _initialDepositId, _crossDomainAdmin, _hubPool, _wrappedNativeTokenAddress)
}

// SpokePoolInit is a paid mutator transaction binding the contract method 0x3ddf5059.
//
// Solidity: function __SpokePool_init(uint32 _initialDepositId, address _crossDomainAdmin, address _hubPool, address _wrappedNativeTokenAddress) returns()
func (_Storage *StorageSession) SpokePoolInit(_initialDepositId uint32, _crossDomainAdmin common.Address, _hubPool common.Address, _wrappedNativeTokenAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SpokePoolInit(&_Storage.TransactOpts, _initialDepositId, _crossDomainAdmin, _hubPool, _wrappedNativeTokenAddress)
}

// SpokePoolInit is a paid mutator transaction binding the contract method 0x3ddf5059.
//
// Solidity: function __SpokePool_init(uint32 _initialDepositId, address _crossDomainAdmin, address _hubPool, address _wrappedNativeTokenAddress) returns()
func (_Storage *StorageTransactorSession) SpokePoolInit(_initialDepositId uint32, _crossDomainAdmin common.Address, _hubPool common.Address, _wrappedNativeTokenAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SpokePoolInit(&_Storage.TransactOpts, _initialDepositId, _crossDomainAdmin, _hubPool, _wrappedNativeTokenAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x1186ec33.
//
// Solidity: function deposit(address recipient, address originToken, uint256 amount, uint256 destinationChainId, int64 relayerFeePct, uint32 quoteTimestamp, bytes message, uint256 maxCount) payable returns()
func (_Storage *StorageTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, originToken common.Address, amount *big.Int, destinationChainId *big.Int, relayerFeePct int64, quoteTimestamp uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deposit", recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1186ec33.
//
// Solidity: function deposit(address recipient, address originToken, uint256 amount, uint256 destinationChainId, int64 relayerFeePct, uint32 quoteTimestamp, bytes message, uint256 maxCount) payable returns()
func (_Storage *StorageSession) Deposit(recipient common.Address, originToken common.Address, amount *big.Int, destinationChainId *big.Int, relayerFeePct int64, quoteTimestamp uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Deposit(&_Storage.TransactOpts, recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1186ec33.
//
// Solidity: function deposit(address recipient, address originToken, uint256 amount, uint256 destinationChainId, int64 relayerFeePct, uint32 quoteTimestamp, bytes message, uint256 maxCount) payable returns()
func (_Storage *StorageTransactorSession) Deposit(recipient common.Address, originToken common.Address, amount *big.Int, destinationChainId *big.Int, relayerFeePct int64, quoteTimestamp uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Deposit(&_Storage.TransactOpts, recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
}

// EmergencyDeleteRootBundle is a paid mutator transaction binding the contract method 0x8a7860ce.
//
// Solidity: function emergencyDeleteRootBundle(uint256 rootBundleId) returns()
func (_Storage *StorageTransactor) EmergencyDeleteRootBundle(opts *bind.TransactOpts, rootBundleId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "emergencyDeleteRootBundle", rootBundleId)
}

// EmergencyDeleteRootBundle is a paid mutator transaction binding the contract method 0x8a7860ce.
//
// Solidity: function emergencyDeleteRootBundle(uint256 rootBundleId) returns()
func (_Storage *StorageSession) EmergencyDeleteRootBundle(rootBundleId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.EmergencyDeleteRootBundle(&_Storage.TransactOpts, rootBundleId)
}

// EmergencyDeleteRootBundle is a paid mutator transaction binding the contract method 0x8a7860ce.
//
// Solidity: function emergencyDeleteRootBundle(uint256 rootBundleId) returns()
func (_Storage *StorageTransactorSession) EmergencyDeleteRootBundle(rootBundleId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.EmergencyDeleteRootBundle(&_Storage.TransactOpts, rootBundleId)
}

// ExecuteRelayerRefundLeaf is a paid mutator transaction binding the contract method 0x1b3d5559.
//
// Solidity: function executeRelayerRefundLeaf(uint32 rootBundleId, (uint256,uint256,uint256[],uint32,address,address[]) relayerRefundLeaf, bytes32[] proof) returns()
func (_Storage *StorageTransactor) ExecuteRelayerRefundLeaf(opts *bind.TransactOpts, rootBundleId uint32, relayerRefundLeaf SpokePoolInterfaceRelayerRefundLeaf, proof [][32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "executeRelayerRefundLeaf", rootBundleId, relayerRefundLeaf, proof)
}

// ExecuteRelayerRefundLeaf is a paid mutator transaction binding the contract method 0x1b3d5559.
//
// Solidity: function executeRelayerRefundLeaf(uint32 rootBundleId, (uint256,uint256,uint256[],uint32,address,address[]) relayerRefundLeaf, bytes32[] proof) returns()
func (_Storage *StorageSession) ExecuteRelayerRefundLeaf(rootBundleId uint32, relayerRefundLeaf SpokePoolInterfaceRelayerRefundLeaf, proof [][32]byte) (*types.Transaction, error) {
	return _Storage.Contract.ExecuteRelayerRefundLeaf(&_Storage.TransactOpts, rootBundleId, relayerRefundLeaf, proof)
}

// ExecuteRelayerRefundLeaf is a paid mutator transaction binding the contract method 0x1b3d5559.
//
// Solidity: function executeRelayerRefundLeaf(uint32 rootBundleId, (uint256,uint256,uint256[],uint32,address,address[]) relayerRefundLeaf, bytes32[] proof) returns()
func (_Storage *StorageTransactorSession) ExecuteRelayerRefundLeaf(rootBundleId uint32, relayerRefundLeaf SpokePoolInterfaceRelayerRefundLeaf, proof [][32]byte) (*types.Transaction, error) {
	return _Storage.Contract.ExecuteRelayerRefundLeaf(&_Storage.TransactOpts, rootBundleId, relayerRefundLeaf, proof)
}

// ExecuteSlowRelayLeaf is a paid mutator transaction binding the contract method 0x0c2097f7.
//
// Solidity: function executeSlowRelayLeaf(address depositor, address recipient, address destinationToken, uint256 totalRelayAmount, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, uint32 rootBundleId, bytes message, int256 payoutAdjustment, bytes32[] proof) returns()
func (_Storage *StorageTransactor) ExecuteSlowRelayLeaf(opts *bind.TransactOpts, depositor common.Address, recipient common.Address, destinationToken common.Address, totalRelayAmount *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, rootBundleId uint32, message []byte, payoutAdjustment *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "executeSlowRelayLeaf", depositor, recipient, destinationToken, totalRelayAmount, originChainId, realizedLpFeePct, relayerFeePct, depositId, rootBundleId, message, payoutAdjustment, proof)
}

// ExecuteSlowRelayLeaf is a paid mutator transaction binding the contract method 0x0c2097f7.
//
// Solidity: function executeSlowRelayLeaf(address depositor, address recipient, address destinationToken, uint256 totalRelayAmount, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, uint32 rootBundleId, bytes message, int256 payoutAdjustment, bytes32[] proof) returns()
func (_Storage *StorageSession) ExecuteSlowRelayLeaf(depositor common.Address, recipient common.Address, destinationToken common.Address, totalRelayAmount *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, rootBundleId uint32, message []byte, payoutAdjustment *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Storage.Contract.ExecuteSlowRelayLeaf(&_Storage.TransactOpts, depositor, recipient, destinationToken, totalRelayAmount, originChainId, realizedLpFeePct, relayerFeePct, depositId, rootBundleId, message, payoutAdjustment, proof)
}

// ExecuteSlowRelayLeaf is a paid mutator transaction binding the contract method 0x0c2097f7.
//
// Solidity: function executeSlowRelayLeaf(address depositor, address recipient, address destinationToken, uint256 totalRelayAmount, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, uint32 rootBundleId, bytes message, int256 payoutAdjustment, bytes32[] proof) returns()
func (_Storage *StorageTransactorSession) ExecuteSlowRelayLeaf(depositor common.Address, recipient common.Address, destinationToken common.Address, totalRelayAmount *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, rootBundleId uint32, message []byte, payoutAdjustment *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Storage.Contract.ExecuteSlowRelayLeaf(&_Storage.TransactOpts, depositor, recipient, destinationToken, totalRelayAmount, originChainId, realizedLpFeePct, relayerFeePct, depositId, rootBundleId, message, payoutAdjustment, proof)
}

// FillRelay is a paid mutator transaction binding the contract method 0x44b8be68.
//
// Solidity: function fillRelay(address depositor, address recipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, bytes message, uint256 maxCount) returns()
func (_Storage *StorageTransactor) FillRelay(opts *bind.TransactOpts, depositor common.Address, recipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillRelay", depositor, recipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, depositId, message, maxCount)
}

// FillRelay is a paid mutator transaction binding the contract method 0x44b8be68.
//
// Solidity: function fillRelay(address depositor, address recipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, bytes message, uint256 maxCount) returns()
func (_Storage *StorageSession) FillRelay(depositor common.Address, recipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillRelay(&_Storage.TransactOpts, depositor, recipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, depositId, message, maxCount)
}

// FillRelay is a paid mutator transaction binding the contract method 0x44b8be68.
//
// Solidity: function fillRelay(address depositor, address recipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, bytes message, uint256 maxCount) returns()
func (_Storage *StorageTransactorSession) FillRelay(depositor common.Address, recipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillRelay(&_Storage.TransactOpts, depositor, recipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, depositId, message, maxCount)
}

// FillRelayWithUpdatedDeposit is a paid mutator transaction binding the contract method 0x5ceaec32.
//
// Solidity: function fillRelayWithUpdatedDeposit(address depositor, address recipient, address updatedRecipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, int64 updatedRelayerFeePct, uint32 depositId, bytes message, bytes updatedMessage, bytes depositorSignature, uint256 maxCount) returns()
func (_Storage *StorageTransactor) FillRelayWithUpdatedDeposit(opts *bind.TransactOpts, depositor common.Address, recipient common.Address, updatedRecipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, updatedRelayerFeePct int64, depositId uint32, message []byte, updatedMessage []byte, depositorSignature []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillRelayWithUpdatedDeposit", depositor, recipient, updatedRecipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, updatedRelayerFeePct, depositId, message, updatedMessage, depositorSignature, maxCount)
}

// FillRelayWithUpdatedDeposit is a paid mutator transaction binding the contract method 0x5ceaec32.
//
// Solidity: function fillRelayWithUpdatedDeposit(address depositor, address recipient, address updatedRecipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, int64 updatedRelayerFeePct, uint32 depositId, bytes message, bytes updatedMessage, bytes depositorSignature, uint256 maxCount) returns()
func (_Storage *StorageSession) FillRelayWithUpdatedDeposit(depositor common.Address, recipient common.Address, updatedRecipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, updatedRelayerFeePct int64, depositId uint32, message []byte, updatedMessage []byte, depositorSignature []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillRelayWithUpdatedDeposit(&_Storage.TransactOpts, depositor, recipient, updatedRecipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, updatedRelayerFeePct, depositId, message, updatedMessage, depositorSignature, maxCount)
}

// FillRelayWithUpdatedDeposit is a paid mutator transaction binding the contract method 0x5ceaec32.
//
// Solidity: function fillRelayWithUpdatedDeposit(address depositor, address recipient, address updatedRecipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, int64 updatedRelayerFeePct, uint32 depositId, bytes message, bytes updatedMessage, bytes depositorSignature, uint256 maxCount) returns()
func (_Storage *StorageTransactorSession) FillRelayWithUpdatedDeposit(depositor common.Address, recipient common.Address, updatedRecipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, updatedRelayerFeePct int64, depositId uint32, message []byte, updatedMessage []byte, depositorSignature []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillRelayWithUpdatedDeposit(&_Storage.TransactOpts, depositor, recipient, updatedRecipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, updatedRelayerFeePct, depositId, message, updatedMessage, depositorSignature, maxCount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf2c20924.
//
// Solidity: function initialize(uint32 _initialDepositId, address _zkErc20Bridge, address _crossDomainAdmin, address _hubPool, address _wethAddress) returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts, _initialDepositId uint32, _zkErc20Bridge common.Address, _crossDomainAdmin common.Address, _hubPool common.Address, _wethAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize", _initialDepositId, _zkErc20Bridge, _crossDomainAdmin, _hubPool, _wethAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf2c20924.
//
// Solidity: function initialize(uint32 _initialDepositId, address _zkErc20Bridge, address _crossDomainAdmin, address _hubPool, address _wethAddress) returns()
func (_Storage *StorageSession) Initialize(_initialDepositId uint32, _zkErc20Bridge common.Address, _crossDomainAdmin common.Address, _hubPool common.Address, _wethAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, _initialDepositId, _zkErc20Bridge, _crossDomainAdmin, _hubPool, _wethAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf2c20924.
//
// Solidity: function initialize(uint32 _initialDepositId, address _zkErc20Bridge, address _crossDomainAdmin, address _hubPool, address _wethAddress) returns()
func (_Storage *StorageTransactorSession) Initialize(_initialDepositId uint32, _zkErc20Bridge common.Address, _crossDomainAdmin common.Address, _hubPool common.Address, _wethAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, _initialDepositId, _zkErc20Bridge, _crossDomainAdmin, _hubPool, _wethAddress)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Storage *StorageTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Storage *StorageSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Storage.Contract.Multicall(&_Storage.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Storage *StorageTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Storage.Contract.Multicall(&_Storage.TransactOpts, data)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x738b62e5.
//
// Solidity: function pauseDeposits(bool pause) returns()
func (_Storage *StorageTransactor) PauseDeposits(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "pauseDeposits", pause)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x738b62e5.
//
// Solidity: function pauseDeposits(bool pause) returns()
func (_Storage *StorageSession) PauseDeposits(pause bool) (*types.Transaction, error) {
	return _Storage.Contract.PauseDeposits(&_Storage.TransactOpts, pause)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x738b62e5.
//
// Solidity: function pauseDeposits(bool pause) returns()
func (_Storage *StorageTransactorSession) PauseDeposits(pause bool) (*types.Transaction, error) {
	return _Storage.Contract.PauseDeposits(&_Storage.TransactOpts, pause)
}

// PauseFills is a paid mutator transaction binding the contract method 0x99cc2968.
//
// Solidity: function pauseFills(bool pause) returns()
func (_Storage *StorageTransactor) PauseFills(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "pauseFills", pause)
}

// PauseFills is a paid mutator transaction binding the contract method 0x99cc2968.
//
// Solidity: function pauseFills(bool pause) returns()
func (_Storage *StorageSession) PauseFills(pause bool) (*types.Transaction, error) {
	return _Storage.Contract.PauseFills(&_Storage.TransactOpts, pause)
}

// PauseFills is a paid mutator transaction binding the contract method 0x99cc2968.
//
// Solidity: function pauseFills(bool pause) returns()
func (_Storage *StorageTransactorSession) PauseFills(pause bool) (*types.Transaction, error) {
	return _Storage.Contract.PauseFills(&_Storage.TransactOpts, pause)
}

// RelayRootBundle is a paid mutator transaction binding the contract method 0x493a4f84.
//
// Solidity: function relayRootBundle(bytes32 relayerRefundRoot, bytes32 slowRelayRoot) returns()
func (_Storage *StorageTransactor) RelayRootBundle(opts *bind.TransactOpts, relayerRefundRoot [32]byte, slowRelayRoot [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "relayRootBundle", relayerRefundRoot, slowRelayRoot)
}

// RelayRootBundle is a paid mutator transaction binding the contract method 0x493a4f84.
//
// Solidity: function relayRootBundle(bytes32 relayerRefundRoot, bytes32 slowRelayRoot) returns()
func (_Storage *StorageSession) RelayRootBundle(relayerRefundRoot [32]byte, slowRelayRoot [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.RelayRootBundle(&_Storage.TransactOpts, relayerRefundRoot, slowRelayRoot)
}

// RelayRootBundle is a paid mutator transaction binding the contract method 0x493a4f84.
//
// Solidity: function relayRootBundle(bytes32 relayerRefundRoot, bytes32 slowRelayRoot) returns()
func (_Storage *StorageTransactorSession) RelayRootBundle(relayerRefundRoot [32]byte, slowRelayRoot [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.RelayRootBundle(&_Storage.TransactOpts, relayerRefundRoot, slowRelayRoot)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa634893f.
//
// Solidity: function requestRefund(address refundToken, uint256 amount, uint256 originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 depositId, uint256 fillBlock, uint256 maxCount) returns()
func (_Storage *StorageTransactor) RequestRefund(opts *bind.TransactOpts, refundToken common.Address, amount *big.Int, originChainId *big.Int, destinationChainId *big.Int, realizedLpFeePct int64, depositId uint32, fillBlock *big.Int, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "requestRefund", refundToken, amount, originChainId, destinationChainId, realizedLpFeePct, depositId, fillBlock, maxCount)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa634893f.
//
// Solidity: function requestRefund(address refundToken, uint256 amount, uint256 originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 depositId, uint256 fillBlock, uint256 maxCount) returns()
func (_Storage *StorageSession) RequestRefund(refundToken common.Address, amount *big.Int, originChainId *big.Int, destinationChainId *big.Int, realizedLpFeePct int64, depositId uint32, fillBlock *big.Int, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RequestRefund(&_Storage.TransactOpts, refundToken, amount, originChainId, destinationChainId, realizedLpFeePct, depositId, fillBlock, maxCount)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa634893f.
//
// Solidity: function requestRefund(address refundToken, uint256 amount, uint256 originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 depositId, uint256 fillBlock, uint256 maxCount) returns()
func (_Storage *StorageTransactorSession) RequestRefund(refundToken common.Address, amount *big.Int, originChainId *big.Int, destinationChainId *big.Int, realizedLpFeePct int64, depositId uint32, fillBlock *big.Int, maxCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RequestRefund(&_Storage.TransactOpts, refundToken, amount, originChainId, destinationChainId, realizedLpFeePct, depositId, fillBlock, maxCount)
}

// SetCrossDomainAdmin is a paid mutator transaction binding the contract method 0xde7eba78.
//
// Solidity: function setCrossDomainAdmin(address newCrossDomainAdmin) returns()
func (_Storage *StorageTransactor) SetCrossDomainAdmin(opts *bind.TransactOpts, newCrossDomainAdmin common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setCrossDomainAdmin", newCrossDomainAdmin)
}

// SetCrossDomainAdmin is a paid mutator transaction binding the contract method 0xde7eba78.
//
// Solidity: function setCrossDomainAdmin(address newCrossDomainAdmin) returns()
func (_Storage *StorageSession) SetCrossDomainAdmin(newCrossDomainAdmin common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetCrossDomainAdmin(&_Storage.TransactOpts, newCrossDomainAdmin)
}

// SetCrossDomainAdmin is a paid mutator transaction binding the contract method 0xde7eba78.
//
// Solidity: function setCrossDomainAdmin(address newCrossDomainAdmin) returns()
func (_Storage *StorageTransactorSession) SetCrossDomainAdmin(newCrossDomainAdmin common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetCrossDomainAdmin(&_Storage.TransactOpts, newCrossDomainAdmin)
}

// SetDepositQuoteTimeBuffer is a paid mutator transaction binding the contract method 0x2752042e.
//
// Solidity: function setDepositQuoteTimeBuffer(uint32 newDepositQuoteTimeBuffer) returns()
func (_Storage *StorageTransactor) SetDepositQuoteTimeBuffer(opts *bind.TransactOpts, newDepositQuoteTimeBuffer uint32) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setDepositQuoteTimeBuffer", newDepositQuoteTimeBuffer)
}

// SetDepositQuoteTimeBuffer is a paid mutator transaction binding the contract method 0x2752042e.
//
// Solidity: function setDepositQuoteTimeBuffer(uint32 newDepositQuoteTimeBuffer) returns()
func (_Storage *StorageSession) SetDepositQuoteTimeBuffer(newDepositQuoteTimeBuffer uint32) (*types.Transaction, error) {
	return _Storage.Contract.SetDepositQuoteTimeBuffer(&_Storage.TransactOpts, newDepositQuoteTimeBuffer)
}

// SetDepositQuoteTimeBuffer is a paid mutator transaction binding the contract method 0x2752042e.
//
// Solidity: function setDepositQuoteTimeBuffer(uint32 newDepositQuoteTimeBuffer) returns()
func (_Storage *StorageTransactorSession) SetDepositQuoteTimeBuffer(newDepositQuoteTimeBuffer uint32) (*types.Transaction, error) {
	return _Storage.Contract.SetDepositQuoteTimeBuffer(&_Storage.TransactOpts, newDepositQuoteTimeBuffer)
}

// SetEnableRoute is a paid mutator transaction binding the contract method 0x272751c7.
//
// Solidity: function setEnableRoute(address originToken, uint256 destinationChainId, bool enabled) returns()
func (_Storage *StorageTransactor) SetEnableRoute(opts *bind.TransactOpts, originToken common.Address, destinationChainId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setEnableRoute", originToken, destinationChainId, enabled)
}

// SetEnableRoute is a paid mutator transaction binding the contract method 0x272751c7.
//
// Solidity: function setEnableRoute(address originToken, uint256 destinationChainId, bool enabled) returns()
func (_Storage *StorageSession) SetEnableRoute(originToken common.Address, destinationChainId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Storage.Contract.SetEnableRoute(&_Storage.TransactOpts, originToken, destinationChainId, enabled)
}

// SetEnableRoute is a paid mutator transaction binding the contract method 0x272751c7.
//
// Solidity: function setEnableRoute(address originToken, uint256 destinationChainId, bool enabled) returns()
func (_Storage *StorageTransactorSession) SetEnableRoute(originToken common.Address, destinationChainId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Storage.Contract.SetEnableRoute(&_Storage.TransactOpts, originToken, destinationChainId, enabled)
}

// SetHubPool is a paid mutator transaction binding the contract method 0x1dfb2d02.
//
// Solidity: function setHubPool(address newHubPool) returns()
func (_Storage *StorageTransactor) SetHubPool(opts *bind.TransactOpts, newHubPool common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setHubPool", newHubPool)
}

// SetHubPool is a paid mutator transaction binding the contract method 0x1dfb2d02.
//
// Solidity: function setHubPool(address newHubPool) returns()
func (_Storage *StorageSession) SetHubPool(newHubPool common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetHubPool(&_Storage.TransactOpts, newHubPool)
}

// SetHubPool is a paid mutator transaction binding the contract method 0x1dfb2d02.
//
// Solidity: function setHubPool(address newHubPool) returns()
func (_Storage *StorageTransactorSession) SetHubPool(newHubPool common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetHubPool(&_Storage.TransactOpts, newHubPool)
}

// SetZkBridge is a paid mutator transaction binding the contract method 0xcd44141f.
//
// Solidity: function setZkBridge(address _zkErc20Bridge) returns()
func (_Storage *StorageTransactor) SetZkBridge(opts *bind.TransactOpts, _zkErc20Bridge common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setZkBridge", _zkErc20Bridge)
}

// SetZkBridge is a paid mutator transaction binding the contract method 0xcd44141f.
//
// Solidity: function setZkBridge(address _zkErc20Bridge) returns()
func (_Storage *StorageSession) SetZkBridge(_zkErc20Bridge common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetZkBridge(&_Storage.TransactOpts, _zkErc20Bridge)
}

// SetZkBridge is a paid mutator transaction binding the contract method 0xcd44141f.
//
// Solidity: function setZkBridge(address _zkErc20Bridge) returns()
func (_Storage *StorageTransactorSession) SetZkBridge(_zkErc20Bridge common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetZkBridge(&_Storage.TransactOpts, _zkErc20Bridge)
}

// SpeedUpDeposit is a paid mutator transaction binding the contract method 0x7e688bba.
//
// Solidity: function speedUpDeposit(address depositor, int64 updatedRelayerFeePct, uint32 depositId, address updatedRecipient, bytes updatedMessage, bytes depositorSignature) returns()
func (_Storage *StorageTransactor) SpeedUpDeposit(opts *bind.TransactOpts, depositor common.Address, updatedRelayerFeePct int64, depositId uint32, updatedRecipient common.Address, updatedMessage []byte, depositorSignature []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "speedUpDeposit", depositor, updatedRelayerFeePct, depositId, updatedRecipient, updatedMessage, depositorSignature)
}

// SpeedUpDeposit is a paid mutator transaction binding the contract method 0x7e688bba.
//
// Solidity: function speedUpDeposit(address depositor, int64 updatedRelayerFeePct, uint32 depositId, address updatedRecipient, bytes updatedMessage, bytes depositorSignature) returns()
func (_Storage *StorageSession) SpeedUpDeposit(depositor common.Address, updatedRelayerFeePct int64, depositId uint32, updatedRecipient common.Address, updatedMessage []byte, depositorSignature []byte) (*types.Transaction, error) {
	return _Storage.Contract.SpeedUpDeposit(&_Storage.TransactOpts, depositor, updatedRelayerFeePct, depositId, updatedRecipient, updatedMessage, depositorSignature)
}

// SpeedUpDeposit is a paid mutator transaction binding the contract method 0x7e688bba.
//
// Solidity: function speedUpDeposit(address depositor, int64 updatedRelayerFeePct, uint32 depositId, address updatedRecipient, bytes updatedMessage, bytes depositorSignature) returns()
func (_Storage *StorageTransactorSession) SpeedUpDeposit(depositor common.Address, updatedRelayerFeePct int64, depositId uint32, updatedRecipient common.Address, updatedMessage []byte, depositorSignature []byte) (*types.Transaction, error) {
	return _Storage.Contract.SpeedUpDeposit(&_Storage.TransactOpts, depositor, updatedRelayerFeePct, depositId, updatedRecipient, updatedMessage, depositorSignature)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeTo(&_Storage.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Storage *StorageTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeTo(&_Storage.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactorSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// StorageAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Storage contract.
type StorageAdminChangedIterator struct {
	Event *StorageAdminChanged // Event containing the contract specifics and raw log

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
func (it *StorageAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAdminChanged)
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
		it.Event = new(StorageAdminChanged)
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
func (it *StorageAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAdminChanged represents a AdminChanged event raised by the Storage contract.
type StorageAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StorageAdminChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StorageAdminChangedIterator{contract: _Storage.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StorageAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAdminChanged)
				if err := _Storage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Storage *StorageFilterer) ParseAdminChanged(log types.Log) (*StorageAdminChanged, error) {
	event := new(StorageAdminChanged)
	if err := _Storage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Storage contract.
type StorageBeaconUpgradedIterator struct {
	Event *StorageBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageBeaconUpgraded)
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
		it.Event = new(StorageBeaconUpgraded)
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
func (it *StorageBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageBeaconUpgraded represents a BeaconUpgraded event raised by the Storage contract.
type StorageBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StorageBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StorageBeaconUpgradedIterator{contract: _Storage.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StorageBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageBeaconUpgraded)
				if err := _Storage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Storage *StorageFilterer) ParseBeaconUpgraded(log types.Log) (*StorageBeaconUpgraded, error) {
	event := new(StorageBeaconUpgraded)
	if err := _Storage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageEmergencyDeleteRootBundleIterator is returned from FilterEmergencyDeleteRootBundle and is used to iterate over the raw logs and unpacked data for EmergencyDeleteRootBundle events raised by the Storage contract.
type StorageEmergencyDeleteRootBundleIterator struct {
	Event *StorageEmergencyDeleteRootBundle // Event containing the contract specifics and raw log

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
func (it *StorageEmergencyDeleteRootBundleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageEmergencyDeleteRootBundle)
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
		it.Event = new(StorageEmergencyDeleteRootBundle)
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
func (it *StorageEmergencyDeleteRootBundleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageEmergencyDeleteRootBundleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageEmergencyDeleteRootBundle represents a EmergencyDeleteRootBundle event raised by the Storage contract.
type StorageEmergencyDeleteRootBundle struct {
	RootBundleId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyDeleteRootBundle is a free log retrieval operation binding the contract event 0x3569b846531b754c99cb80df3f49cd72fa6fe106aaee5ab8e0caf35a9d7ce88d.
//
// Solidity: event EmergencyDeleteRootBundle(uint256 indexed rootBundleId)
func (_Storage *StorageFilterer) FilterEmergencyDeleteRootBundle(opts *bind.FilterOpts, rootBundleId []*big.Int) (*StorageEmergencyDeleteRootBundleIterator, error) {

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "EmergencyDeleteRootBundle", rootBundleIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageEmergencyDeleteRootBundleIterator{contract: _Storage.contract, event: "EmergencyDeleteRootBundle", logs: logs, sub: sub}, nil
}

// WatchEmergencyDeleteRootBundle is a free log subscription operation binding the contract event 0x3569b846531b754c99cb80df3f49cd72fa6fe106aaee5ab8e0caf35a9d7ce88d.
//
// Solidity: event EmergencyDeleteRootBundle(uint256 indexed rootBundleId)
func (_Storage *StorageFilterer) WatchEmergencyDeleteRootBundle(opts *bind.WatchOpts, sink chan<- *StorageEmergencyDeleteRootBundle, rootBundleId []*big.Int) (event.Subscription, error) {

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "EmergencyDeleteRootBundle", rootBundleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageEmergencyDeleteRootBundle)
				if err := _Storage.contract.UnpackLog(event, "EmergencyDeleteRootBundle", log); err != nil {
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

// ParseEmergencyDeleteRootBundle is a log parse operation binding the contract event 0x3569b846531b754c99cb80df3f49cd72fa6fe106aaee5ab8e0caf35a9d7ce88d.
//
// Solidity: event EmergencyDeleteRootBundle(uint256 indexed rootBundleId)
func (_Storage *StorageFilterer) ParseEmergencyDeleteRootBundle(log types.Log) (*StorageEmergencyDeleteRootBundle, error) {
	event := new(StorageEmergencyDeleteRootBundle)
	if err := _Storage.contract.UnpackLog(event, "EmergencyDeleteRootBundle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageEnabledDepositRouteIterator is returned from FilterEnabledDepositRoute and is used to iterate over the raw logs and unpacked data for EnabledDepositRoute events raised by the Storage contract.
type StorageEnabledDepositRouteIterator struct {
	Event *StorageEnabledDepositRoute // Event containing the contract specifics and raw log

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
func (it *StorageEnabledDepositRouteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageEnabledDepositRoute)
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
		it.Event = new(StorageEnabledDepositRoute)
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
func (it *StorageEnabledDepositRouteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageEnabledDepositRouteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageEnabledDepositRoute represents a EnabledDepositRoute event raised by the Storage contract.
type StorageEnabledDepositRoute struct {
	OriginToken        common.Address
	DestinationChainId *big.Int
	Enabled            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEnabledDepositRoute is a free log retrieval operation binding the contract event 0x0a21fdd43d0ad0c62689ee7230a47309a050755bcc52eba00310add65297692a.
//
// Solidity: event EnabledDepositRoute(address indexed originToken, uint256 indexed destinationChainId, bool enabled)
func (_Storage *StorageFilterer) FilterEnabledDepositRoute(opts *bind.FilterOpts, originToken []common.Address, destinationChainId []*big.Int) (*StorageEnabledDepositRouteIterator, error) {

	var originTokenRule []interface{}
	for _, originTokenItem := range originToken {
		originTokenRule = append(originTokenRule, originTokenItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "EnabledDepositRoute", originTokenRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageEnabledDepositRouteIterator{contract: _Storage.contract, event: "EnabledDepositRoute", logs: logs, sub: sub}, nil
}

// WatchEnabledDepositRoute is a free log subscription operation binding the contract event 0x0a21fdd43d0ad0c62689ee7230a47309a050755bcc52eba00310add65297692a.
//
// Solidity: event EnabledDepositRoute(address indexed originToken, uint256 indexed destinationChainId, bool enabled)
func (_Storage *StorageFilterer) WatchEnabledDepositRoute(opts *bind.WatchOpts, sink chan<- *StorageEnabledDepositRoute, originToken []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var originTokenRule []interface{}
	for _, originTokenItem := range originToken {
		originTokenRule = append(originTokenRule, originTokenItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "EnabledDepositRoute", originTokenRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageEnabledDepositRoute)
				if err := _Storage.contract.UnpackLog(event, "EnabledDepositRoute", log); err != nil {
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

// ParseEnabledDepositRoute is a log parse operation binding the contract event 0x0a21fdd43d0ad0c62689ee7230a47309a050755bcc52eba00310add65297692a.
//
// Solidity: event EnabledDepositRoute(address indexed originToken, uint256 indexed destinationChainId, bool enabled)
func (_Storage *StorageFilterer) ParseEnabledDepositRoute(log types.Log) (*StorageEnabledDepositRoute, error) {
	event := new(StorageEnabledDepositRoute)
	if err := _Storage.contract.UnpackLog(event, "EnabledDepositRoute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageExecutedRelayerRefundRootIterator is returned from FilterExecutedRelayerRefundRoot and is used to iterate over the raw logs and unpacked data for ExecutedRelayerRefundRoot events raised by the Storage contract.
type StorageExecutedRelayerRefundRootIterator struct {
	Event *StorageExecutedRelayerRefundRoot // Event containing the contract specifics and raw log

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
func (it *StorageExecutedRelayerRefundRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageExecutedRelayerRefundRoot)
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
		it.Event = new(StorageExecutedRelayerRefundRoot)
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
func (it *StorageExecutedRelayerRefundRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageExecutedRelayerRefundRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageExecutedRelayerRefundRoot represents a ExecutedRelayerRefundRoot event raised by the Storage contract.
type StorageExecutedRelayerRefundRoot struct {
	AmountToReturn  *big.Int
	ChainId         *big.Int
	RefundAmounts   []*big.Int
	RootBundleId    uint32
	LeafId          uint32
	L2TokenAddress  common.Address
	RefundAddresses []common.Address
	Caller          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExecutedRelayerRefundRoot is a free log retrieval operation binding the contract event 0xf8bd640004bcec1b89657020f561d0b070cbdf662d0b158db9dccb0a8301bfab.
//
// Solidity: event ExecutedRelayerRefundRoot(uint256 amountToReturn, uint256 indexed chainId, uint256[] refundAmounts, uint32 indexed rootBundleId, uint32 indexed leafId, address l2TokenAddress, address[] refundAddresses, address caller)
func (_Storage *StorageFilterer) FilterExecutedRelayerRefundRoot(opts *bind.FilterOpts, chainId []*big.Int, rootBundleId []uint32, leafId []uint32) (*StorageExecutedRelayerRefundRootIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}
	var leafIdRule []interface{}
	for _, leafIdItem := range leafId {
		leafIdRule = append(leafIdRule, leafIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ExecutedRelayerRefundRoot", chainIdRule, rootBundleIdRule, leafIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageExecutedRelayerRefundRootIterator{contract: _Storage.contract, event: "ExecutedRelayerRefundRoot", logs: logs, sub: sub}, nil
}

// WatchExecutedRelayerRefundRoot is a free log subscription operation binding the contract event 0xf8bd640004bcec1b89657020f561d0b070cbdf662d0b158db9dccb0a8301bfab.
//
// Solidity: event ExecutedRelayerRefundRoot(uint256 amountToReturn, uint256 indexed chainId, uint256[] refundAmounts, uint32 indexed rootBundleId, uint32 indexed leafId, address l2TokenAddress, address[] refundAddresses, address caller)
func (_Storage *StorageFilterer) WatchExecutedRelayerRefundRoot(opts *bind.WatchOpts, sink chan<- *StorageExecutedRelayerRefundRoot, chainId []*big.Int, rootBundleId []uint32, leafId []uint32) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}
	var leafIdRule []interface{}
	for _, leafIdItem := range leafId {
		leafIdRule = append(leafIdRule, leafIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ExecutedRelayerRefundRoot", chainIdRule, rootBundleIdRule, leafIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageExecutedRelayerRefundRoot)
				if err := _Storage.contract.UnpackLog(event, "ExecutedRelayerRefundRoot", log); err != nil {
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

// ParseExecutedRelayerRefundRoot is a log parse operation binding the contract event 0xf8bd640004bcec1b89657020f561d0b070cbdf662d0b158db9dccb0a8301bfab.
//
// Solidity: event ExecutedRelayerRefundRoot(uint256 amountToReturn, uint256 indexed chainId, uint256[] refundAmounts, uint32 indexed rootBundleId, uint32 indexed leafId, address l2TokenAddress, address[] refundAddresses, address caller)
func (_Storage *StorageFilterer) ParseExecutedRelayerRefundRoot(log types.Log) (*StorageExecutedRelayerRefundRoot, error) {
	event := new(StorageExecutedRelayerRefundRoot)
	if err := _Storage.contract.UnpackLog(event, "ExecutedRelayerRefundRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageFilledRelayIterator is returned from FilterFilledRelay and is used to iterate over the raw logs and unpacked data for FilledRelay events raised by the Storage contract.
type StorageFilledRelayIterator struct {
	Event *StorageFilledRelay // Event containing the contract specifics and raw log

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
func (it *StorageFilledRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFilledRelay)
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
		it.Event = new(StorageFilledRelay)
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
func (it *StorageFilledRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFilledRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFilledRelay represents a FilledRelay event raised by the Storage contract.
type StorageFilledRelay struct {
	Amount             *big.Int
	TotalFilledAmount  *big.Int
	FillAmount         *big.Int
	RepaymentChainId   *big.Int
	OriginChainId      *big.Int
	DestinationChainId *big.Int
	RelayerFeePct      int64
	RealizedLpFeePct   int64
	DepositId          uint32
	DestinationToken   common.Address
	Relayer            common.Address
	Depositor          common.Address
	Recipient          common.Address
	Message            []byte
	UpdatableRelayData SpokePoolRelayExecutionInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFilledRelay is a free log retrieval operation binding the contract event 0x8ab9dc6c19fe88e69bc70221b339c84332752fdd49591b7c51e66bae3947b73c.
//
// Solidity: event FilledRelay(uint256 amount, uint256 totalFilledAmount, uint256 fillAmount, uint256 repaymentChainId, uint256 indexed originChainId, uint256 destinationChainId, int64 relayerFeePct, int64 realizedLpFeePct, uint32 indexed depositId, address destinationToken, address relayer, address indexed depositor, address recipient, bytes message, (address,bytes,int64,bool,int256) updatableRelayData)
func (_Storage *StorageFilterer) FilterFilledRelay(opts *bind.FilterOpts, originChainId []*big.Int, depositId []uint32, depositor []common.Address) (*StorageFilledRelayIterator, error) {

	var originChainIdRule []interface{}
	for _, originChainIdItem := range originChainId {
		originChainIdRule = append(originChainIdRule, originChainIdItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FilledRelay", originChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &StorageFilledRelayIterator{contract: _Storage.contract, event: "FilledRelay", logs: logs, sub: sub}, nil
}

// WatchFilledRelay is a free log subscription operation binding the contract event 0x8ab9dc6c19fe88e69bc70221b339c84332752fdd49591b7c51e66bae3947b73c.
//
// Solidity: event FilledRelay(uint256 amount, uint256 totalFilledAmount, uint256 fillAmount, uint256 repaymentChainId, uint256 indexed originChainId, uint256 destinationChainId, int64 relayerFeePct, int64 realizedLpFeePct, uint32 indexed depositId, address destinationToken, address relayer, address indexed depositor, address recipient, bytes message, (address,bytes,int64,bool,int256) updatableRelayData)
func (_Storage *StorageFilterer) WatchFilledRelay(opts *bind.WatchOpts, sink chan<- *StorageFilledRelay, originChainId []*big.Int, depositId []uint32, depositor []common.Address) (event.Subscription, error) {

	var originChainIdRule []interface{}
	for _, originChainIdItem := range originChainId {
		originChainIdRule = append(originChainIdRule, originChainIdItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FilledRelay", originChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFilledRelay)
				if err := _Storage.contract.UnpackLog(event, "FilledRelay", log); err != nil {
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

// ParseFilledRelay is a log parse operation binding the contract event 0x8ab9dc6c19fe88e69bc70221b339c84332752fdd49591b7c51e66bae3947b73c.
//
// Solidity: event FilledRelay(uint256 amount, uint256 totalFilledAmount, uint256 fillAmount, uint256 repaymentChainId, uint256 indexed originChainId, uint256 destinationChainId, int64 relayerFeePct, int64 realizedLpFeePct, uint32 indexed depositId, address destinationToken, address relayer, address indexed depositor, address recipient, bytes message, (address,bytes,int64,bool,int256) updatableRelayData)
func (_Storage *StorageFilterer) ParseFilledRelay(log types.Log) (*StorageFilledRelay, error) {
	event := new(StorageFilledRelay)
	if err := _Storage.contract.UnpackLog(event, "FilledRelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the Storage contract.
type StorageFundsDepositedIterator struct {
	Event *StorageFundsDeposited // Event containing the contract specifics and raw log

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
func (it *StorageFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFundsDeposited)
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
		it.Event = new(StorageFundsDeposited)
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
func (it *StorageFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFundsDeposited represents a FundsDeposited event raised by the Storage contract.
type StorageFundsDeposited struct {
	Amount             *big.Int
	OriginChainId      *big.Int
	DestinationChainId *big.Int
	RelayerFeePct      int64
	DepositId          uint32
	QuoteTimestamp     uint32
	OriginToken        common.Address
	Recipient          common.Address
	Depositor          common.Address
	Message            []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0xafc4df6845a4ab948b492800d3d8a25d538a102a2bc07cd01f1cfa097fddcff6.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, uint256 indexed destinationChainId, int64 relayerFeePct, uint32 indexed depositId, uint32 quoteTimestamp, address originToken, address recipient, address indexed depositor, bytes message)
func (_Storage *StorageFilterer) FilterFundsDeposited(opts *bind.FilterOpts, destinationChainId []*big.Int, depositId []uint32, depositor []common.Address) (*StorageFundsDepositedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FundsDeposited", destinationChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &StorageFundsDepositedIterator{contract: _Storage.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0xafc4df6845a4ab948b492800d3d8a25d538a102a2bc07cd01f1cfa097fddcff6.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, uint256 indexed destinationChainId, int64 relayerFeePct, uint32 indexed depositId, uint32 quoteTimestamp, address originToken, address recipient, address indexed depositor, bytes message)
func (_Storage *StorageFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *StorageFundsDeposited, destinationChainId []*big.Int, depositId []uint32, depositor []common.Address) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FundsDeposited", destinationChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFundsDeposited)
				if err := _Storage.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
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

// ParseFundsDeposited is a log parse operation binding the contract event 0xafc4df6845a4ab948b492800d3d8a25d538a102a2bc07cd01f1cfa097fddcff6.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, uint256 indexed destinationChainId, int64 relayerFeePct, uint32 indexed depositId, uint32 quoteTimestamp, address originToken, address recipient, address indexed depositor, bytes message)
func (_Storage *StorageFilterer) ParseFundsDeposited(log types.Log) (*StorageFundsDeposited, error) {
	event := new(StorageFundsDeposited)
	if err := _Storage.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoragePausedDepositsIterator is returned from FilterPausedDeposits and is used to iterate over the raw logs and unpacked data for PausedDeposits events raised by the Storage contract.
type StoragePausedDepositsIterator struct {
	Event *StoragePausedDeposits // Event containing the contract specifics and raw log

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
func (it *StoragePausedDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoragePausedDeposits)
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
		it.Event = new(StoragePausedDeposits)
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
func (it *StoragePausedDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoragePausedDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoragePausedDeposits represents a PausedDeposits event raised by the Storage contract.
type StoragePausedDeposits struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPausedDeposits is a free log retrieval operation binding the contract event 0xe88463c2f254e2b070013a2dc7ee1e099f9bc00534cbdf03af551dc26ae49219.
//
// Solidity: event PausedDeposits(bool isPaused)
func (_Storage *StorageFilterer) FilterPausedDeposits(opts *bind.FilterOpts) (*StoragePausedDepositsIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "PausedDeposits")
	if err != nil {
		return nil, err
	}
	return &StoragePausedDepositsIterator{contract: _Storage.contract, event: "PausedDeposits", logs: logs, sub: sub}, nil
}

// WatchPausedDeposits is a free log subscription operation binding the contract event 0xe88463c2f254e2b070013a2dc7ee1e099f9bc00534cbdf03af551dc26ae49219.
//
// Solidity: event PausedDeposits(bool isPaused)
func (_Storage *StorageFilterer) WatchPausedDeposits(opts *bind.WatchOpts, sink chan<- *StoragePausedDeposits) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "PausedDeposits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoragePausedDeposits)
				if err := _Storage.contract.UnpackLog(event, "PausedDeposits", log); err != nil {
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

// ParsePausedDeposits is a log parse operation binding the contract event 0xe88463c2f254e2b070013a2dc7ee1e099f9bc00534cbdf03af551dc26ae49219.
//
// Solidity: event PausedDeposits(bool isPaused)
func (_Storage *StorageFilterer) ParsePausedDeposits(log types.Log) (*StoragePausedDeposits, error) {
	event := new(StoragePausedDeposits)
	if err := _Storage.contract.UnpackLog(event, "PausedDeposits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoragePausedFillsIterator is returned from FilterPausedFills and is used to iterate over the raw logs and unpacked data for PausedFills events raised by the Storage contract.
type StoragePausedFillsIterator struct {
	Event *StoragePausedFills // Event containing the contract specifics and raw log

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
func (it *StoragePausedFillsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoragePausedFills)
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
		it.Event = new(StoragePausedFills)
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
func (it *StoragePausedFillsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoragePausedFillsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoragePausedFills represents a PausedFills event raised by the Storage contract.
type StoragePausedFills struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPausedFills is a free log retrieval operation binding the contract event 0x2d5b62420992e5a4afce0e77742636ca2608ef58289fd2e1baa5161ef6e7e41e.
//
// Solidity: event PausedFills(bool isPaused)
func (_Storage *StorageFilterer) FilterPausedFills(opts *bind.FilterOpts) (*StoragePausedFillsIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "PausedFills")
	if err != nil {
		return nil, err
	}
	return &StoragePausedFillsIterator{contract: _Storage.contract, event: "PausedFills", logs: logs, sub: sub}, nil
}

// WatchPausedFills is a free log subscription operation binding the contract event 0x2d5b62420992e5a4afce0e77742636ca2608ef58289fd2e1baa5161ef6e7e41e.
//
// Solidity: event PausedFills(bool isPaused)
func (_Storage *StorageFilterer) WatchPausedFills(opts *bind.WatchOpts, sink chan<- *StoragePausedFills) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "PausedFills")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoragePausedFills)
				if err := _Storage.contract.UnpackLog(event, "PausedFills", log); err != nil {
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

// ParsePausedFills is a log parse operation binding the contract event 0x2d5b62420992e5a4afce0e77742636ca2608ef58289fd2e1baa5161ef6e7e41e.
//
// Solidity: event PausedFills(bool isPaused)
func (_Storage *StorageFilterer) ParsePausedFills(log types.Log) (*StoragePausedFills, error) {
	event := new(StoragePausedFills)
	if err := _Storage.contract.UnpackLog(event, "PausedFills", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRefundRequestedIterator is returned from FilterRefundRequested and is used to iterate over the raw logs and unpacked data for RefundRequested events raised by the Storage contract.
type StorageRefundRequestedIterator struct {
	Event *StorageRefundRequested // Event containing the contract specifics and raw log

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
func (it *StorageRefundRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRefundRequested)
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
		it.Event = new(StorageRefundRequested)
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
func (it *StorageRefundRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRefundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRefundRequested represents a RefundRequested event raised by the Storage contract.
type StorageRefundRequested struct {
	Relayer                   common.Address
	RefundToken               common.Address
	Amount                    *big.Int
	OriginChainId             *big.Int
	DestinationChainId        *big.Int
	RealizedLpFeePct          int64
	DepositId                 uint32
	FillBlock                 *big.Int
	PreviousIdenticalRequests *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterRefundRequested is a free log retrieval operation binding the contract event 0xace81ce0f8b8d27a1aed0c4df5f6b2743e46fc50fe3e0183dd7cd6a7b9db22fb.
//
// Solidity: event RefundRequested(address indexed relayer, address refundToken, uint256 amount, uint256 indexed originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 indexed depositId, uint256 fillBlock, uint256 previousIdenticalRequests)
func (_Storage *StorageFilterer) FilterRefundRequested(opts *bind.FilterOpts, relayer []common.Address, originChainId []*big.Int, depositId []uint32) (*StorageRefundRequestedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	var originChainIdRule []interface{}
	for _, originChainIdItem := range originChainId {
		originChainIdRule = append(originChainIdRule, originChainIdItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RefundRequested", relayerRule, originChainIdRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageRefundRequestedIterator{contract: _Storage.contract, event: "RefundRequested", logs: logs, sub: sub}, nil
}

// WatchRefundRequested is a free log subscription operation binding the contract event 0xace81ce0f8b8d27a1aed0c4df5f6b2743e46fc50fe3e0183dd7cd6a7b9db22fb.
//
// Solidity: event RefundRequested(address indexed relayer, address refundToken, uint256 amount, uint256 indexed originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 indexed depositId, uint256 fillBlock, uint256 previousIdenticalRequests)
func (_Storage *StorageFilterer) WatchRefundRequested(opts *bind.WatchOpts, sink chan<- *StorageRefundRequested, relayer []common.Address, originChainId []*big.Int, depositId []uint32) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	var originChainIdRule []interface{}
	for _, originChainIdItem := range originChainId {
		originChainIdRule = append(originChainIdRule, originChainIdItem)
	}

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RefundRequested", relayerRule, originChainIdRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRefundRequested)
				if err := _Storage.contract.UnpackLog(event, "RefundRequested", log); err != nil {
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

// ParseRefundRequested is a log parse operation binding the contract event 0xace81ce0f8b8d27a1aed0c4df5f6b2743e46fc50fe3e0183dd7cd6a7b9db22fb.
//
// Solidity: event RefundRequested(address indexed relayer, address refundToken, uint256 amount, uint256 indexed originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 indexed depositId, uint256 fillBlock, uint256 previousIdenticalRequests)
func (_Storage *StorageFilterer) ParseRefundRequested(log types.Log) (*StorageRefundRequested, error) {
	event := new(StorageRefundRequested)
	if err := _Storage.contract.UnpackLog(event, "RefundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRelayedRootBundleIterator is returned from FilterRelayedRootBundle and is used to iterate over the raw logs and unpacked data for RelayedRootBundle events raised by the Storage contract.
type StorageRelayedRootBundleIterator struct {
	Event *StorageRelayedRootBundle // Event containing the contract specifics and raw log

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
func (it *StorageRelayedRootBundleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRelayedRootBundle)
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
		it.Event = new(StorageRelayedRootBundle)
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
func (it *StorageRelayedRootBundleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRelayedRootBundleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRelayedRootBundle represents a RelayedRootBundle event raised by the Storage contract.
type StorageRelayedRootBundle struct {
	RootBundleId      uint32
	RelayerRefundRoot [32]byte
	SlowRelayRoot     [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRelayedRootBundle is a free log retrieval operation binding the contract event 0xc86ba04c55bc5eb2f2876b91c438849a296dbec7b08751c3074d92e04f0a77af.
//
// Solidity: event RelayedRootBundle(uint32 indexed rootBundleId, bytes32 indexed relayerRefundRoot, bytes32 indexed slowRelayRoot)
func (_Storage *StorageFilterer) FilterRelayedRootBundle(opts *bind.FilterOpts, rootBundleId []uint32, relayerRefundRoot [][32]byte, slowRelayRoot [][32]byte) (*StorageRelayedRootBundleIterator, error) {

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}
	var relayerRefundRootRule []interface{}
	for _, relayerRefundRootItem := range relayerRefundRoot {
		relayerRefundRootRule = append(relayerRefundRootRule, relayerRefundRootItem)
	}
	var slowRelayRootRule []interface{}
	for _, slowRelayRootItem := range slowRelayRoot {
		slowRelayRootRule = append(slowRelayRootRule, slowRelayRootItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RelayedRootBundle", rootBundleIdRule, relayerRefundRootRule, slowRelayRootRule)
	if err != nil {
		return nil, err
	}
	return &StorageRelayedRootBundleIterator{contract: _Storage.contract, event: "RelayedRootBundle", logs: logs, sub: sub}, nil
}

// WatchRelayedRootBundle is a free log subscription operation binding the contract event 0xc86ba04c55bc5eb2f2876b91c438849a296dbec7b08751c3074d92e04f0a77af.
//
// Solidity: event RelayedRootBundle(uint32 indexed rootBundleId, bytes32 indexed relayerRefundRoot, bytes32 indexed slowRelayRoot)
func (_Storage *StorageFilterer) WatchRelayedRootBundle(opts *bind.WatchOpts, sink chan<- *StorageRelayedRootBundle, rootBundleId []uint32, relayerRefundRoot [][32]byte, slowRelayRoot [][32]byte) (event.Subscription, error) {

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}
	var relayerRefundRootRule []interface{}
	for _, relayerRefundRootItem := range relayerRefundRoot {
		relayerRefundRootRule = append(relayerRefundRootRule, relayerRefundRootItem)
	}
	var slowRelayRootRule []interface{}
	for _, slowRelayRootItem := range slowRelayRoot {
		slowRelayRootRule = append(slowRelayRootRule, slowRelayRootItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RelayedRootBundle", rootBundleIdRule, relayerRefundRootRule, slowRelayRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRelayedRootBundle)
				if err := _Storage.contract.UnpackLog(event, "RelayedRootBundle", log); err != nil {
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

// ParseRelayedRootBundle is a log parse operation binding the contract event 0xc86ba04c55bc5eb2f2876b91c438849a296dbec7b08751c3074d92e04f0a77af.
//
// Solidity: event RelayedRootBundle(uint32 indexed rootBundleId, bytes32 indexed relayerRefundRoot, bytes32 indexed slowRelayRoot)
func (_Storage *StorageFilterer) ParseRelayedRootBundle(log types.Log) (*StorageRelayedRootBundle, error) {
	event := new(StorageRelayedRootBundle)
	if err := _Storage.contract.UnpackLog(event, "RelayedRootBundle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRequestedSpeedUpDepositIterator is returned from FilterRequestedSpeedUpDeposit and is used to iterate over the raw logs and unpacked data for RequestedSpeedUpDeposit events raised by the Storage contract.
type StorageRequestedSpeedUpDepositIterator struct {
	Event *StorageRequestedSpeedUpDeposit // Event containing the contract specifics and raw log

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
func (it *StorageRequestedSpeedUpDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRequestedSpeedUpDeposit)
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
		it.Event = new(StorageRequestedSpeedUpDeposit)
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
func (it *StorageRequestedSpeedUpDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRequestedSpeedUpDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRequestedSpeedUpDeposit represents a RequestedSpeedUpDeposit event raised by the Storage contract.
type StorageRequestedSpeedUpDeposit struct {
	NewRelayerFeePct   int64
	DepositId          uint32
	Depositor          common.Address
	UpdatedRecipient   common.Address
	UpdatedMessage     []byte
	DepositorSignature []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRequestedSpeedUpDeposit is a free log retrieval operation binding the contract event 0xa6aa57bd282fc82378458de27be4eadfa791a0c7321c49562eeba8b2643dd566.
//
// Solidity: event RequestedSpeedUpDeposit(int64 newRelayerFeePct, uint32 indexed depositId, address indexed depositor, address updatedRecipient, bytes updatedMessage, bytes depositorSignature)
func (_Storage *StorageFilterer) FilterRequestedSpeedUpDeposit(opts *bind.FilterOpts, depositId []uint32, depositor []common.Address) (*StorageRequestedSpeedUpDepositIterator, error) {

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RequestedSpeedUpDeposit", depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &StorageRequestedSpeedUpDepositIterator{contract: _Storage.contract, event: "RequestedSpeedUpDeposit", logs: logs, sub: sub}, nil
}

// WatchRequestedSpeedUpDeposit is a free log subscription operation binding the contract event 0xa6aa57bd282fc82378458de27be4eadfa791a0c7321c49562eeba8b2643dd566.
//
// Solidity: event RequestedSpeedUpDeposit(int64 newRelayerFeePct, uint32 indexed depositId, address indexed depositor, address updatedRecipient, bytes updatedMessage, bytes depositorSignature)
func (_Storage *StorageFilterer) WatchRequestedSpeedUpDeposit(opts *bind.WatchOpts, sink chan<- *StorageRequestedSpeedUpDeposit, depositId []uint32, depositor []common.Address) (event.Subscription, error) {

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RequestedSpeedUpDeposit", depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRequestedSpeedUpDeposit)
				if err := _Storage.contract.UnpackLog(event, "RequestedSpeedUpDeposit", log); err != nil {
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

// ParseRequestedSpeedUpDeposit is a log parse operation binding the contract event 0xa6aa57bd282fc82378458de27be4eadfa791a0c7321c49562eeba8b2643dd566.
//
// Solidity: event RequestedSpeedUpDeposit(int64 newRelayerFeePct, uint32 indexed depositId, address indexed depositor, address updatedRecipient, bytes updatedMessage, bytes depositorSignature)
func (_Storage *StorageFilterer) ParseRequestedSpeedUpDeposit(log types.Log) (*StorageRequestedSpeedUpDeposit, error) {
	event := new(StorageRequestedSpeedUpDeposit)
	if err := _Storage.contract.UnpackLog(event, "RequestedSpeedUpDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetDepositQuoteTimeBufferIterator is returned from FilterSetDepositQuoteTimeBuffer and is used to iterate over the raw logs and unpacked data for SetDepositQuoteTimeBuffer events raised by the Storage contract.
type StorageSetDepositQuoteTimeBufferIterator struct {
	Event *StorageSetDepositQuoteTimeBuffer // Event containing the contract specifics and raw log

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
func (it *StorageSetDepositQuoteTimeBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetDepositQuoteTimeBuffer)
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
		it.Event = new(StorageSetDepositQuoteTimeBuffer)
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
func (it *StorageSetDepositQuoteTimeBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetDepositQuoteTimeBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetDepositQuoteTimeBuffer represents a SetDepositQuoteTimeBuffer event raised by the Storage contract.
type StorageSetDepositQuoteTimeBuffer struct {
	NewBuffer uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetDepositQuoteTimeBuffer is a free log retrieval operation binding the contract event 0x0e55dd180fa793d9036c804d0a116e6a7617a48e72cee1f83d92793a793fcc03.
//
// Solidity: event SetDepositQuoteTimeBuffer(uint32 newBuffer)
func (_Storage *StorageFilterer) FilterSetDepositQuoteTimeBuffer(opts *bind.FilterOpts) (*StorageSetDepositQuoteTimeBufferIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetDepositQuoteTimeBuffer")
	if err != nil {
		return nil, err
	}
	return &StorageSetDepositQuoteTimeBufferIterator{contract: _Storage.contract, event: "SetDepositQuoteTimeBuffer", logs: logs, sub: sub}, nil
}

// WatchSetDepositQuoteTimeBuffer is a free log subscription operation binding the contract event 0x0e55dd180fa793d9036c804d0a116e6a7617a48e72cee1f83d92793a793fcc03.
//
// Solidity: event SetDepositQuoteTimeBuffer(uint32 newBuffer)
func (_Storage *StorageFilterer) WatchSetDepositQuoteTimeBuffer(opts *bind.WatchOpts, sink chan<- *StorageSetDepositQuoteTimeBuffer) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetDepositQuoteTimeBuffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetDepositQuoteTimeBuffer)
				if err := _Storage.contract.UnpackLog(event, "SetDepositQuoteTimeBuffer", log); err != nil {
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

// ParseSetDepositQuoteTimeBuffer is a log parse operation binding the contract event 0x0e55dd180fa793d9036c804d0a116e6a7617a48e72cee1f83d92793a793fcc03.
//
// Solidity: event SetDepositQuoteTimeBuffer(uint32 newBuffer)
func (_Storage *StorageFilterer) ParseSetDepositQuoteTimeBuffer(log types.Log) (*StorageSetDepositQuoteTimeBuffer, error) {
	event := new(StorageSetDepositQuoteTimeBuffer)
	if err := _Storage.contract.UnpackLog(event, "SetDepositQuoteTimeBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetHubPoolIterator is returned from FilterSetHubPool and is used to iterate over the raw logs and unpacked data for SetHubPool events raised by the Storage contract.
type StorageSetHubPoolIterator struct {
	Event *StorageSetHubPool // Event containing the contract specifics and raw log

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
func (it *StorageSetHubPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetHubPool)
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
		it.Event = new(StorageSetHubPool)
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
func (it *StorageSetHubPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetHubPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetHubPool represents a SetHubPool event raised by the Storage contract.
type StorageSetHubPool struct {
	NewHubPool common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetHubPool is a free log retrieval operation binding the contract event 0x1f17a88f67b0f49060a34bec1a4723a563620e6aa265eb640b5046dcee0759a0.
//
// Solidity: event SetHubPool(address indexed newHubPool)
func (_Storage *StorageFilterer) FilterSetHubPool(opts *bind.FilterOpts, newHubPool []common.Address) (*StorageSetHubPoolIterator, error) {

	var newHubPoolRule []interface{}
	for _, newHubPoolItem := range newHubPool {
		newHubPoolRule = append(newHubPoolRule, newHubPoolItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetHubPool", newHubPoolRule)
	if err != nil {
		return nil, err
	}
	return &StorageSetHubPoolIterator{contract: _Storage.contract, event: "SetHubPool", logs: logs, sub: sub}, nil
}

// WatchSetHubPool is a free log subscription operation binding the contract event 0x1f17a88f67b0f49060a34bec1a4723a563620e6aa265eb640b5046dcee0759a0.
//
// Solidity: event SetHubPool(address indexed newHubPool)
func (_Storage *StorageFilterer) WatchSetHubPool(opts *bind.WatchOpts, sink chan<- *StorageSetHubPool, newHubPool []common.Address) (event.Subscription, error) {

	var newHubPoolRule []interface{}
	for _, newHubPoolItem := range newHubPool {
		newHubPoolRule = append(newHubPoolRule, newHubPoolItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetHubPool", newHubPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetHubPool)
				if err := _Storage.contract.UnpackLog(event, "SetHubPool", log); err != nil {
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

// ParseSetHubPool is a log parse operation binding the contract event 0x1f17a88f67b0f49060a34bec1a4723a563620e6aa265eb640b5046dcee0759a0.
//
// Solidity: event SetHubPool(address indexed newHubPool)
func (_Storage *StorageFilterer) ParseSetHubPool(log types.Log) (*StorageSetHubPool, error) {
	event := new(StorageSetHubPool)
	if err := _Storage.contract.UnpackLog(event, "SetHubPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetXDomainAdminIterator is returned from FilterSetXDomainAdmin and is used to iterate over the raw logs and unpacked data for SetXDomainAdmin events raised by the Storage contract.
type StorageSetXDomainAdminIterator struct {
	Event *StorageSetXDomainAdmin // Event containing the contract specifics and raw log

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
func (it *StorageSetXDomainAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetXDomainAdmin)
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
		it.Event = new(StorageSetXDomainAdmin)
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
func (it *StorageSetXDomainAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetXDomainAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetXDomainAdmin represents a SetXDomainAdmin event raised by the Storage contract.
type StorageSetXDomainAdmin struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetXDomainAdmin is a free log retrieval operation binding the contract event 0xa9e8c42c9e7fca7f62755189a16b2f5314d43d8fb24e91ba54e6d65f9314e849.
//
// Solidity: event SetXDomainAdmin(address indexed newAdmin)
func (_Storage *StorageFilterer) FilterSetXDomainAdmin(opts *bind.FilterOpts, newAdmin []common.Address) (*StorageSetXDomainAdminIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetXDomainAdmin", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &StorageSetXDomainAdminIterator{contract: _Storage.contract, event: "SetXDomainAdmin", logs: logs, sub: sub}, nil
}

// WatchSetXDomainAdmin is a free log subscription operation binding the contract event 0xa9e8c42c9e7fca7f62755189a16b2f5314d43d8fb24e91ba54e6d65f9314e849.
//
// Solidity: event SetXDomainAdmin(address indexed newAdmin)
func (_Storage *StorageFilterer) WatchSetXDomainAdmin(opts *bind.WatchOpts, sink chan<- *StorageSetXDomainAdmin, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetXDomainAdmin", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetXDomainAdmin)
				if err := _Storage.contract.UnpackLog(event, "SetXDomainAdmin", log); err != nil {
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

// ParseSetXDomainAdmin is a log parse operation binding the contract event 0xa9e8c42c9e7fca7f62755189a16b2f5314d43d8fb24e91ba54e6d65f9314e849.
//
// Solidity: event SetXDomainAdmin(address indexed newAdmin)
func (_Storage *StorageFilterer) ParseSetXDomainAdmin(log types.Log) (*StorageSetXDomainAdmin, error) {
	event := new(StorageSetXDomainAdmin)
	if err := _Storage.contract.UnpackLog(event, "SetXDomainAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSetZkBridgeIterator is returned from FilterSetZkBridge and is used to iterate over the raw logs and unpacked data for SetZkBridge events raised by the Storage contract.
type StorageSetZkBridgeIterator struct {
	Event *StorageSetZkBridge // Event containing the contract specifics and raw log

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
func (it *StorageSetZkBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSetZkBridge)
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
		it.Event = new(StorageSetZkBridge)
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
func (it *StorageSetZkBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSetZkBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSetZkBridge represents a SetZkBridge event raised by the Storage contract.
type StorageSetZkBridge struct {
	Erc20Bridge    common.Address
	OldErc20Bridge common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetZkBridge is a free log retrieval operation binding the contract event 0xab9987862d7dab8c2bf2f09eb3ad7b6c1961c86fe3e9e7d3ef50dc98995d72d4.
//
// Solidity: event SetZkBridge(address indexed erc20Bridge, address indexed oldErc20Bridge)
func (_Storage *StorageFilterer) FilterSetZkBridge(opts *bind.FilterOpts, erc20Bridge []common.Address, oldErc20Bridge []common.Address) (*StorageSetZkBridgeIterator, error) {

	var erc20BridgeRule []interface{}
	for _, erc20BridgeItem := range erc20Bridge {
		erc20BridgeRule = append(erc20BridgeRule, erc20BridgeItem)
	}
	var oldErc20BridgeRule []interface{}
	for _, oldErc20BridgeItem := range oldErc20Bridge {
		oldErc20BridgeRule = append(oldErc20BridgeRule, oldErc20BridgeItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SetZkBridge", erc20BridgeRule, oldErc20BridgeRule)
	if err != nil {
		return nil, err
	}
	return &StorageSetZkBridgeIterator{contract: _Storage.contract, event: "SetZkBridge", logs: logs, sub: sub}, nil
}

// WatchSetZkBridge is a free log subscription operation binding the contract event 0xab9987862d7dab8c2bf2f09eb3ad7b6c1961c86fe3e9e7d3ef50dc98995d72d4.
//
// Solidity: event SetZkBridge(address indexed erc20Bridge, address indexed oldErc20Bridge)
func (_Storage *StorageFilterer) WatchSetZkBridge(opts *bind.WatchOpts, sink chan<- *StorageSetZkBridge, erc20Bridge []common.Address, oldErc20Bridge []common.Address) (event.Subscription, error) {

	var erc20BridgeRule []interface{}
	for _, erc20BridgeItem := range erc20Bridge {
		erc20BridgeRule = append(erc20BridgeRule, erc20BridgeItem)
	}
	var oldErc20BridgeRule []interface{}
	for _, oldErc20BridgeItem := range oldErc20Bridge {
		oldErc20BridgeRule = append(oldErc20BridgeRule, oldErc20BridgeItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SetZkBridge", erc20BridgeRule, oldErc20BridgeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSetZkBridge)
				if err := _Storage.contract.UnpackLog(event, "SetZkBridge", log); err != nil {
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

// ParseSetZkBridge is a log parse operation binding the contract event 0xab9987862d7dab8c2bf2f09eb3ad7b6c1961c86fe3e9e7d3ef50dc98995d72d4.
//
// Solidity: event SetZkBridge(address indexed erc20Bridge, address indexed oldErc20Bridge)
func (_Storage *StorageFilterer) ParseSetZkBridge(log types.Log) (*StorageSetZkBridge, error) {
	event := new(StorageSetZkBridge)
	if err := _Storage.contract.UnpackLog(event, "SetZkBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageTokensBridgedIterator is returned from FilterTokensBridged and is used to iterate over the raw logs and unpacked data for TokensBridged events raised by the Storage contract.
type StorageTokensBridgedIterator struct {
	Event *StorageTokensBridged // Event containing the contract specifics and raw log

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
func (it *StorageTokensBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageTokensBridged)
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
		it.Event = new(StorageTokensBridged)
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
func (it *StorageTokensBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageTokensBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageTokensBridged represents a TokensBridged event raised by the Storage contract.
type StorageTokensBridged struct {
	AmountToReturn *big.Int
	ChainId        *big.Int
	LeafId         uint32
	L2TokenAddress common.Address
	Caller         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokensBridged is a free log retrieval operation binding the contract event 0x828fc203220356df8f072a91681caee7d5c75095e2a95e80ed5a14b384697f71.
//
// Solidity: event TokensBridged(uint256 amountToReturn, uint256 indexed chainId, uint32 indexed leafId, address indexed l2TokenAddress, address caller)
func (_Storage *StorageFilterer) FilterTokensBridged(opts *bind.FilterOpts, chainId []*big.Int, leafId []uint32, l2TokenAddress []common.Address) (*StorageTokensBridgedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var leafIdRule []interface{}
	for _, leafIdItem := range leafId {
		leafIdRule = append(leafIdRule, leafIdItem)
	}
	var l2TokenAddressRule []interface{}
	for _, l2TokenAddressItem := range l2TokenAddress {
		l2TokenAddressRule = append(l2TokenAddressRule, l2TokenAddressItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "TokensBridged", chainIdRule, leafIdRule, l2TokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &StorageTokensBridgedIterator{contract: _Storage.contract, event: "TokensBridged", logs: logs, sub: sub}, nil
}

// WatchTokensBridged is a free log subscription operation binding the contract event 0x828fc203220356df8f072a91681caee7d5c75095e2a95e80ed5a14b384697f71.
//
// Solidity: event TokensBridged(uint256 amountToReturn, uint256 indexed chainId, uint32 indexed leafId, address indexed l2TokenAddress, address caller)
func (_Storage *StorageFilterer) WatchTokensBridged(opts *bind.WatchOpts, sink chan<- *StorageTokensBridged, chainId []*big.Int, leafId []uint32, l2TokenAddress []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var leafIdRule []interface{}
	for _, leafIdItem := range leafId {
		leafIdRule = append(leafIdRule, leafIdItem)
	}
	var l2TokenAddressRule []interface{}
	for _, l2TokenAddressItem := range l2TokenAddress {
		l2TokenAddressRule = append(l2TokenAddressRule, l2TokenAddressItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "TokensBridged", chainIdRule, leafIdRule, l2TokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageTokensBridged)
				if err := _Storage.contract.UnpackLog(event, "TokensBridged", log); err != nil {
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

// ParseTokensBridged is a log parse operation binding the contract event 0x828fc203220356df8f072a91681caee7d5c75095e2a95e80ed5a14b384697f71.
//
// Solidity: event TokensBridged(uint256 amountToReturn, uint256 indexed chainId, uint32 indexed leafId, address indexed l2TokenAddress, address caller)
func (_Storage *StorageFilterer) ParseTokensBridged(log types.Log) (*StorageTokensBridged, error) {
	event := new(StorageTokensBridged)
	if err := _Storage.contract.UnpackLog(event, "TokensBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Storage contract.
type StorageUpgradedIterator struct {
	Event *StorageUpgraded // Event containing the contract specifics and raw log

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
func (it *StorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpgraded)
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
		it.Event = new(StorageUpgraded)
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
func (it *StorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpgraded represents a Upgraded event raised by the Storage contract.
type StorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StorageUpgradedIterator{contract: _Storage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpgraded)
				if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) ParseUpgraded(log types.Log) (*StorageUpgraded, error) {
	event := new(StorageUpgraded)
	if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageZkSyncTokensBridgedIterator is returned from FilterZkSyncTokensBridged and is used to iterate over the raw logs and unpacked data for ZkSyncTokensBridged events raised by the Storage contract.
type StorageZkSyncTokensBridgedIterator struct {
	Event *StorageZkSyncTokensBridged // Event containing the contract specifics and raw log

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
func (it *StorageZkSyncTokensBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageZkSyncTokensBridged)
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
		it.Event = new(StorageZkSyncTokensBridged)
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
func (it *StorageZkSyncTokensBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageZkSyncTokensBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageZkSyncTokensBridged represents a ZkSyncTokensBridged event raised by the Storage contract.
type StorageZkSyncTokensBridged struct {
	L2Token               common.Address
	Target                common.Address
	NumberOfTokensBridged *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZkSyncTokensBridged is a free log retrieval operation binding the contract event 0x9d523e6fab68ea063b948ed2d9e4e9091522a3bdbb63f1ad70c744b0edb9cdbf.
//
// Solidity: event ZkSyncTokensBridged(address indexed l2Token, address target, uint256 numberOfTokensBridged)
func (_Storage *StorageFilterer) FilterZkSyncTokensBridged(opts *bind.FilterOpts, l2Token []common.Address) (*StorageZkSyncTokensBridgedIterator, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ZkSyncTokensBridged", l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &StorageZkSyncTokensBridgedIterator{contract: _Storage.contract, event: "ZkSyncTokensBridged", logs: logs, sub: sub}, nil
}

// WatchZkSyncTokensBridged is a free log subscription operation binding the contract event 0x9d523e6fab68ea063b948ed2d9e4e9091522a3bdbb63f1ad70c744b0edb9cdbf.
//
// Solidity: event ZkSyncTokensBridged(address indexed l2Token, address target, uint256 numberOfTokensBridged)
func (_Storage *StorageFilterer) WatchZkSyncTokensBridged(opts *bind.WatchOpts, sink chan<- *StorageZkSyncTokensBridged, l2Token []common.Address) (event.Subscription, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ZkSyncTokensBridged", l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageZkSyncTokensBridged)
				if err := _Storage.contract.UnpackLog(event, "ZkSyncTokensBridged", log); err != nil {
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

// ParseZkSyncTokensBridged is a log parse operation binding the contract event 0x9d523e6fab68ea063b948ed2d9e4e9091522a3bdbb63f1ad70c744b0edb9cdbf.
//
// Solidity: event ZkSyncTokensBridged(address indexed l2Token, address target, uint256 numberOfTokensBridged)
func (_Storage *StorageFilterer) ParseZkSyncTokensBridged(log types.Log) (*StorageZkSyncTokensBridged, error) {
	event := new(StorageZkSyncTokensBridged)
	if err := _Storage.contract.UnpackLog(event, "ZkSyncTokensBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
