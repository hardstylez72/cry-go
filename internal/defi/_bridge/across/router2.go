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

// Router2MetaData contains all meta data concerning the Router2 contract.
var Router2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rootBundleId\",\"type\":\"uint256\"}],\"name\":\"EmergencyDeleteRootBundle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"EnabledDepositRoute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToReturn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"refundAmounts\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l2TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"refundAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ExecutedRelayerRefundRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFilledAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repaymentChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"isSlowRelay\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"payoutAdjustmentPct\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structSpokePool.RelayExecutionInfo\",\"name\":\"updatableRelayData\",\"type\":\"tuple\"}],\"name\":\"FilledRelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quoteTimestamp\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PausedDeposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PausedFills\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fillBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousIdenticalRequests\",\"type\":\"uint256\"}],\"name\":\"RefundRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relayerRefundRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"slowRelayRoot\",\"type\":\"bytes32\"}],\"name\":\"RelayedRootBundle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"newRelayerFeePct\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"updatedMessage\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"depositorSignature\",\"type\":\"bytes\"}],\"name\":\"RequestedSpeedUpDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newBuffer\",\"type\":\"uint32\"}],\"name\":\"SetDepositQuoteTimeBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHubPool\",\"type\":\"address\"}],\"name\":\"SetHubPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"SetXDomainAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Bridge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldErc20Bridge\",\"type\":\"address\"}],\"name\":\"SetZkBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToReturn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"leafId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokensBridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfTokensBridged\",\"type\":\"uint256\"}],\"name\":\"ZkSyncTokensBridged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOW_FILL_MAX_TOKENS_TO_SEND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_DEPOSIT_DETAILS_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_initialDepositId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_crossDomainAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hubPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wrappedNativeTokenAddress\",\"type\":\"address\"}],\"name\":\"__SpokePool_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossDomainAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"quoteTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositQuoteTimeBuffer\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rootBundleId\",\"type\":\"uint256\"}],\"name\":\"emergencyDeleteRootBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"enabledDepositRoutes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountToReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"refundAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32\",\"name\":\"leafId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"l2TokenAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"refundAddresses\",\"type\":\"address[]\"}],\"internalType\":\"structSpokePoolInterface.RelayerRefundLeaf\",\"name\":\"relayerRefundLeaf\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"executeRelayerRefundLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalRelayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rootBundleId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int256\",\"name\":\"payoutAdjustment\",\"type\":\"int256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"executeSlowRelayLeaf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fillCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensToSend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repaymentChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"fillRelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"updatedRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokensToSend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"repaymentChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"updatedRelayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"updatedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositorSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"fillRelayWithUpdatedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_initialDepositId\",\"type\":\"uint32\"},{\"internalType\":\"contractZkBridgeLike\",\"name\":\"_zkErc20Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossDomainAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hubPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wethAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Eth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfDeposits\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"pauseDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"pauseFills\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausedDeposits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausedFills\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"refundsRequested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"relayFills\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"relayerRefundRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slowRelayRoot\",\"type\":\"bytes32\"}],\"name\":\"relayRootBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"refundToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"realizedLpFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fillBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rootBundles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"slowRelayRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"relayerRefundRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCrossDomainAdmin\",\"type\":\"address\"}],\"name\":\"setCrossDomainAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newDepositQuoteTimeBuffer\",\"type\":\"uint32\"}],\"name\":\"setDepositQuoteTimeBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setEnableRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHubPool\",\"type\":\"address\"}],\"name\":\"setHubPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractZkBridgeLike\",\"name\":\"_zkErc20Bridge\",\"type\":\"address\"}],\"name\":\"setZkBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"int64\",\"name\":\"updatedRelayerFeePct\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"depositId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"updatedRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"updatedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositorSignature\",\"type\":\"bytes\"}],\"name\":\"speedUpDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractWETH9Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkErc20Bridge\",\"outputs\":[{\"internalType\":\"contractZkBridgeLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Router2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Router2MetaData.ABI instead.
var Router2ABI = Router2MetaData.ABI

// Router2 is an auto generated Go binding around an Ethereum contract.
type Router2 struct {
	Router2Caller     // Read-only binding to the contract
	Router2Transactor // Write-only binding to the contract
	Router2Filterer   // Log filterer for contract events
}

// Router2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Router2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Router2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Router2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Router2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Router2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Router2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Router2Session struct {
	Contract     *Router2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Router2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Router2CallerSession struct {
	Contract *Router2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Router2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Router2TransactorSession struct {
	Contract     *Router2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Router2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Router2Raw struct {
	Contract *Router2 // Generic contract binding to access the raw methods on
}

// Router2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Router2CallerRaw struct {
	Contract *Router2Caller // Generic read-only contract binding to access the raw methods on
}

// Router2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Router2TransactorRaw struct {
	Contract *Router2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter2 creates a new instance of Router2, bound to a specific deployed contract.
func NewRouter2(address common.Address, backend bind.ContractBackend) (*Router2, error) {
	contract, err := bindRouter2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router2{Router2Caller: Router2Caller{contract: contract}, Router2Transactor: Router2Transactor{contract: contract}, Router2Filterer: Router2Filterer{contract: contract}}, nil
}

// NewRouter2Caller creates a new read-only instance of Router2, bound to a specific deployed contract.
func NewRouter2Caller(address common.Address, caller bind.ContractCaller) (*Router2Caller, error) {
	contract, err := bindRouter2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Router2Caller{contract: contract}, nil
}

// NewRouter2Transactor creates a new write-only instance of Router2, bound to a specific deployed contract.
func NewRouter2Transactor(address common.Address, transactor bind.ContractTransactor) (*Router2Transactor, error) {
	contract, err := bindRouter2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Router2Transactor{contract: contract}, nil
}

// NewRouter2Filterer creates a new log filterer instance of Router2, bound to a specific deployed contract.
func NewRouter2Filterer(address common.Address, filterer bind.ContractFilterer) (*Router2Filterer, error) {
	contract, err := bindRouter2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Router2Filterer{contract: contract}, nil
}

// bindRouter2 binds a generic wrapper to an already deployed contract.
func bindRouter2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Router2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router2 *Router2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router2.Contract.Router2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router2 *Router2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router2.Contract.Router2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router2 *Router2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router2.Contract.Router2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router2 *Router2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router2 *Router2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router2 *Router2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router2.Contract.contract.Transact(opts, method, params...)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_Router2 *Router2Caller) MAXTRANSFERSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "MAX_TRANSFER_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_Router2 *Router2Session) MAXTRANSFERSIZE() (*big.Int, error) {
	return _Router2.Contract.MAXTRANSFERSIZE(&_Router2.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_Router2 *Router2CallerSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _Router2.Contract.MAXTRANSFERSIZE(&_Router2.CallOpts)
}

// SLOWFILLMAXTOKENSTOSEND is a free data retrieval call binding the contract method 0xb5e1bf5f.
//
// Solidity: function SLOW_FILL_MAX_TOKENS_TO_SEND() view returns(uint256)
func (_Router2 *Router2Caller) SLOWFILLMAXTOKENSTOSEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "SLOW_FILL_MAX_TOKENS_TO_SEND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLOWFILLMAXTOKENSTOSEND is a free data retrieval call binding the contract method 0xb5e1bf5f.
//
// Solidity: function SLOW_FILL_MAX_TOKENS_TO_SEND() view returns(uint256)
func (_Router2 *Router2Session) SLOWFILLMAXTOKENSTOSEND() (*big.Int, error) {
	return _Router2.Contract.SLOWFILLMAXTOKENSTOSEND(&_Router2.CallOpts)
}

// SLOWFILLMAXTOKENSTOSEND is a free data retrieval call binding the contract method 0xb5e1bf5f.
//
// Solidity: function SLOW_FILL_MAX_TOKENS_TO_SEND() view returns(uint256)
func (_Router2 *Router2CallerSession) SLOWFILLMAXTOKENSTOSEND() (*big.Int, error) {
	return _Router2.Contract.SLOWFILLMAXTOKENSTOSEND(&_Router2.CallOpts)
}

// UPDATEDEPOSITDETAILSHASH is a free data retrieval call binding the contract method 0xa78e4b60.
//
// Solidity: function UPDATE_DEPOSIT_DETAILS_HASH() view returns(bytes32)
func (_Router2 *Router2Caller) UPDATEDEPOSITDETAILSHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "UPDATE_DEPOSIT_DETAILS_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATEDEPOSITDETAILSHASH is a free data retrieval call binding the contract method 0xa78e4b60.
//
// Solidity: function UPDATE_DEPOSIT_DETAILS_HASH() view returns(bytes32)
func (_Router2 *Router2Session) UPDATEDEPOSITDETAILSHASH() ([32]byte, error) {
	return _Router2.Contract.UPDATEDEPOSITDETAILSHASH(&_Router2.CallOpts)
}

// UPDATEDEPOSITDETAILSHASH is a free data retrieval call binding the contract method 0xa78e4b60.
//
// Solidity: function UPDATE_DEPOSIT_DETAILS_HASH() view returns(bytes32)
func (_Router2 *Router2CallerSession) UPDATEDEPOSITDETAILSHASH() ([32]byte, error) {
	return _Router2.Contract.UPDATEDEPOSITDETAILSHASH(&_Router2.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Router2 *Router2Caller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Router2 *Router2Session) ChainId() (*big.Int, error) {
	return _Router2.Contract.ChainId(&_Router2.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Router2 *Router2CallerSession) ChainId() (*big.Int, error) {
	return _Router2.Contract.ChainId(&_Router2.CallOpts)
}

// CrossDomainAdmin is a free data retrieval call binding the contract method 0x5285e058.
//
// Solidity: function crossDomainAdmin() view returns(address)
func (_Router2 *Router2Caller) CrossDomainAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "crossDomainAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossDomainAdmin is a free data retrieval call binding the contract method 0x5285e058.
//
// Solidity: function crossDomainAdmin() view returns(address)
func (_Router2 *Router2Session) CrossDomainAdmin() (common.Address, error) {
	return _Router2.Contract.CrossDomainAdmin(&_Router2.CallOpts)
}

// CrossDomainAdmin is a free data retrieval call binding the contract method 0x5285e058.
//
// Solidity: function crossDomainAdmin() view returns(address)
func (_Router2 *Router2CallerSession) CrossDomainAdmin() (common.Address, error) {
	return _Router2.Contract.CrossDomainAdmin(&_Router2.CallOpts)
}

// DepositCounter is a free data retrieval call binding the contract method 0xc0a8bdd1.
//
// Solidity: function depositCounter(address ) view returns(uint256)
func (_Router2 *Router2Caller) DepositCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "depositCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCounter is a free data retrieval call binding the contract method 0xc0a8bdd1.
//
// Solidity: function depositCounter(address ) view returns(uint256)
func (_Router2 *Router2Session) DepositCounter(arg0 common.Address) (*big.Int, error) {
	return _Router2.Contract.DepositCounter(&_Router2.CallOpts, arg0)
}

// DepositCounter is a free data retrieval call binding the contract method 0xc0a8bdd1.
//
// Solidity: function depositCounter(address ) view returns(uint256)
func (_Router2 *Router2CallerSession) DepositCounter(arg0 common.Address) (*big.Int, error) {
	return _Router2.Contract.DepositCounter(&_Router2.CallOpts, arg0)
}

// DepositQuoteTimeBuffer is a free data retrieval call binding the contract method 0x57f6dcb8.
//
// Solidity: function depositQuoteTimeBuffer() view returns(uint32)
func (_Router2 *Router2Caller) DepositQuoteTimeBuffer(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "depositQuoteTimeBuffer")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DepositQuoteTimeBuffer is a free data retrieval call binding the contract method 0x57f6dcb8.
//
// Solidity: function depositQuoteTimeBuffer() view returns(uint32)
func (_Router2 *Router2Session) DepositQuoteTimeBuffer() (uint32, error) {
	return _Router2.Contract.DepositQuoteTimeBuffer(&_Router2.CallOpts)
}

// DepositQuoteTimeBuffer is a free data retrieval call binding the contract method 0x57f6dcb8.
//
// Solidity: function depositQuoteTimeBuffer() view returns(uint32)
func (_Router2 *Router2CallerSession) DepositQuoteTimeBuffer() (uint32, error) {
	return _Router2.Contract.DepositQuoteTimeBuffer(&_Router2.CallOpts)
}

// EnabledDepositRoutes is a free data retrieval call binding the contract method 0x5249fef1.
//
// Solidity: function enabledDepositRoutes(address , uint256 ) view returns(bool)
func (_Router2 *Router2Caller) EnabledDepositRoutes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "enabledDepositRoutes", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnabledDepositRoutes is a free data retrieval call binding the contract method 0x5249fef1.
//
// Solidity: function enabledDepositRoutes(address , uint256 ) view returns(bool)
func (_Router2 *Router2Session) EnabledDepositRoutes(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Router2.Contract.EnabledDepositRoutes(&_Router2.CallOpts, arg0, arg1)
}

// EnabledDepositRoutes is a free data retrieval call binding the contract method 0x5249fef1.
//
// Solidity: function enabledDepositRoutes(address , uint256 ) view returns(bool)
func (_Router2 *Router2CallerSession) EnabledDepositRoutes(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Router2.Contract.EnabledDepositRoutes(&_Router2.CallOpts, arg0, arg1)
}

// FillCounter is a free data retrieval call binding the contract method 0x872af6ea.
//
// Solidity: function fillCounter(address ) view returns(uint256)
func (_Router2 *Router2Caller) FillCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "fillCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FillCounter is a free data retrieval call binding the contract method 0x872af6ea.
//
// Solidity: function fillCounter(address ) view returns(uint256)
func (_Router2 *Router2Session) FillCounter(arg0 common.Address) (*big.Int, error) {
	return _Router2.Contract.FillCounter(&_Router2.CallOpts, arg0)
}

// FillCounter is a free data retrieval call binding the contract method 0x872af6ea.
//
// Solidity: function fillCounter(address ) view returns(uint256)
func (_Router2 *Router2CallerSession) FillCounter(arg0 common.Address) (*big.Int, error) {
	return _Router2.Contract.FillCounter(&_Router2.CallOpts, arg0)
}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Router2 *Router2Caller) GetCurrentTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "getCurrentTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Router2 *Router2Session) GetCurrentTime() (*big.Int, error) {
	return _Router2.Contract.GetCurrentTime(&_Router2.CallOpts)
}

// GetCurrentTime is a free data retrieval call binding the contract method 0x29cb924d.
//
// Solidity: function getCurrentTime() view returns(uint256)
func (_Router2 *Router2CallerSession) GetCurrentTime() (*big.Int, error) {
	return _Router2.Contract.GetCurrentTime(&_Router2.CallOpts)
}

// HubPool is a free data retrieval call binding the contract method 0xe1904402.
//
// Solidity: function hubPool() view returns(address)
func (_Router2 *Router2Caller) HubPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "hubPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HubPool is a free data retrieval call binding the contract method 0xe1904402.
//
// Solidity: function hubPool() view returns(address)
func (_Router2 *Router2Session) HubPool() (common.Address, error) {
	return _Router2.Contract.HubPool(&_Router2.CallOpts)
}

// HubPool is a free data retrieval call binding the contract method 0xe1904402.
//
// Solidity: function hubPool() view returns(address)
func (_Router2 *Router2CallerSession) HubPool() (common.Address, error) {
	return _Router2.Contract.HubPool(&_Router2.CallOpts)
}

// L2Eth is a free data retrieval call binding the contract method 0xe3229211.
//
// Solidity: function l2Eth() view returns(address)
func (_Router2 *Router2Caller) L2Eth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "l2Eth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Eth is a free data retrieval call binding the contract method 0xe3229211.
//
// Solidity: function l2Eth() view returns(address)
func (_Router2 *Router2Session) L2Eth() (common.Address, error) {
	return _Router2.Contract.L2Eth(&_Router2.CallOpts)
}

// L2Eth is a free data retrieval call binding the contract method 0xe3229211.
//
// Solidity: function l2Eth() view returns(address)
func (_Router2 *Router2CallerSession) L2Eth() (common.Address, error) {
	return _Router2.Contract.L2Eth(&_Router2.CallOpts)
}

// NumberOfDeposits is a free data retrieval call binding the contract method 0xa1244c67.
//
// Solidity: function numberOfDeposits() view returns(uint32)
func (_Router2 *Router2Caller) NumberOfDeposits(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "numberOfDeposits")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumberOfDeposits is a free data retrieval call binding the contract method 0xa1244c67.
//
// Solidity: function numberOfDeposits() view returns(uint32)
func (_Router2 *Router2Session) NumberOfDeposits() (uint32, error) {
	return _Router2.Contract.NumberOfDeposits(&_Router2.CallOpts)
}

// NumberOfDeposits is a free data retrieval call binding the contract method 0xa1244c67.
//
// Solidity: function numberOfDeposits() view returns(uint32)
func (_Router2 *Router2CallerSession) NumberOfDeposits() (uint32, error) {
	return _Router2.Contract.NumberOfDeposits(&_Router2.CallOpts)
}

// PausedDeposits is a free data retrieval call binding the contract method 0x6068d6cb.
//
// Solidity: function pausedDeposits() view returns(bool)
func (_Router2 *Router2Caller) PausedDeposits(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "pausedDeposits")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausedDeposits is a free data retrieval call binding the contract method 0x6068d6cb.
//
// Solidity: function pausedDeposits() view returns(bool)
func (_Router2 *Router2Session) PausedDeposits() (bool, error) {
	return _Router2.Contract.PausedDeposits(&_Router2.CallOpts)
}

// PausedDeposits is a free data retrieval call binding the contract method 0x6068d6cb.
//
// Solidity: function pausedDeposits() view returns(bool)
func (_Router2 *Router2CallerSession) PausedDeposits() (bool, error) {
	return _Router2.Contract.PausedDeposits(&_Router2.CallOpts)
}

// PausedFills is a free data retrieval call binding the contract method 0xdda52113.
//
// Solidity: function pausedFills() view returns(bool)
func (_Router2 *Router2Caller) PausedFills(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "pausedFills")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PausedFills is a free data retrieval call binding the contract method 0xdda52113.
//
// Solidity: function pausedFills() view returns(bool)
func (_Router2 *Router2Session) PausedFills() (bool, error) {
	return _Router2.Contract.PausedFills(&_Router2.CallOpts)
}

// PausedFills is a free data retrieval call binding the contract method 0xdda52113.
//
// Solidity: function pausedFills() view returns(bool)
func (_Router2 *Router2CallerSession) PausedFills() (bool, error) {
	return _Router2.Contract.PausedFills(&_Router2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Router2 *Router2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Router2 *Router2Session) ProxiableUUID() ([32]byte, error) {
	return _Router2.Contract.ProxiableUUID(&_Router2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Router2 *Router2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _Router2.Contract.ProxiableUUID(&_Router2.CallOpts)
}

// RefundsRequested is a free data retrieval call binding the contract method 0x23cd9a47.
//
// Solidity: function refundsRequested(bytes32 ) view returns(uint256)
func (_Router2 *Router2Caller) RefundsRequested(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "refundsRequested", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundsRequested is a free data retrieval call binding the contract method 0x23cd9a47.
//
// Solidity: function refundsRequested(bytes32 ) view returns(uint256)
func (_Router2 *Router2Session) RefundsRequested(arg0 [32]byte) (*big.Int, error) {
	return _Router2.Contract.RefundsRequested(&_Router2.CallOpts, arg0)
}

// RefundsRequested is a free data retrieval call binding the contract method 0x23cd9a47.
//
// Solidity: function refundsRequested(bytes32 ) view returns(uint256)
func (_Router2 *Router2CallerSession) RefundsRequested(arg0 [32]byte) (*big.Int, error) {
	return _Router2.Contract.RefundsRequested(&_Router2.CallOpts, arg0)
}

// RelayFills is a free data retrieval call binding the contract method 0xf06850f6.
//
// Solidity: function relayFills(bytes32 ) view returns(uint256)
func (_Router2 *Router2Caller) RelayFills(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "relayFills", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayFills is a free data retrieval call binding the contract method 0xf06850f6.
//
// Solidity: function relayFills(bytes32 ) view returns(uint256)
func (_Router2 *Router2Session) RelayFills(arg0 [32]byte) (*big.Int, error) {
	return _Router2.Contract.RelayFills(&_Router2.CallOpts, arg0)
}

// RelayFills is a free data retrieval call binding the contract method 0xf06850f6.
//
// Solidity: function relayFills(bytes32 ) view returns(uint256)
func (_Router2 *Router2CallerSession) RelayFills(arg0 [32]byte) (*big.Int, error) {
	return _Router2.Contract.RelayFills(&_Router2.CallOpts, arg0)
}

// RootBundles is a free data retrieval call binding the contract method 0xee2a53f8.
//
// Solidity: function rootBundles(uint256 ) view returns(bytes32 slowRelayRoot, bytes32 relayerRefundRoot)
func (_Router2 *Router2Caller) RootBundles(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SlowRelayRoot     [32]byte
	RelayerRefundRoot [32]byte
}, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "rootBundles", arg0)

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
func (_Router2 *Router2Session) RootBundles(arg0 *big.Int) (struct {
	SlowRelayRoot     [32]byte
	RelayerRefundRoot [32]byte
}, error) {
	return _Router2.Contract.RootBundles(&_Router2.CallOpts, arg0)
}

// RootBundles is a free data retrieval call binding the contract method 0xee2a53f8.
//
// Solidity: function rootBundles(uint256 ) view returns(bytes32 slowRelayRoot, bytes32 relayerRefundRoot)
func (_Router2 *Router2CallerSession) RootBundles(arg0 *big.Int) (struct {
	SlowRelayRoot     [32]byte
	RelayerRefundRoot [32]byte
}, error) {
	return _Router2.Contract.RootBundles(&_Router2.CallOpts, arg0)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Router2 *Router2Caller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Router2 *Router2Session) WrappedNativeToken() (common.Address, error) {
	return _Router2.Contract.WrappedNativeToken(&_Router2.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Router2 *Router2CallerSession) WrappedNativeToken() (common.Address, error) {
	return _Router2.Contract.WrappedNativeToken(&_Router2.CallOpts)
}

// ZkErc20Bridge is a free data retrieval call binding the contract method 0xbb3e04b5.
//
// Solidity: function zkErc20Bridge() view returns(address)
func (_Router2 *Router2Caller) ZkErc20Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router2.contract.Call(opts, &out, "zkErc20Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkErc20Bridge is a free data retrieval call binding the contract method 0xbb3e04b5.
//
// Solidity: function zkErc20Bridge() view returns(address)
func (_Router2 *Router2Session) ZkErc20Bridge() (common.Address, error) {
	return _Router2.Contract.ZkErc20Bridge(&_Router2.CallOpts)
}

// ZkErc20Bridge is a free data retrieval call binding the contract method 0xbb3e04b5.
//
// Solidity: function zkErc20Bridge() view returns(address)
func (_Router2 *Router2CallerSession) ZkErc20Bridge() (common.Address, error) {
	return _Router2.Contract.ZkErc20Bridge(&_Router2.CallOpts)
}

// SpokePoolInit is a paid mutator transaction binding the contract method 0x3ddf5059.
//
// Solidity: function __SpokePool_init(uint32 _initialDepositId, address _crossDomainAdmin, address _hubPool, address _wrappedNativeTokenAddress) returns()
func (_Router2 *Router2Transactor) SpokePoolInit(opts *bind.TransactOpts, _initialDepositId uint32, _crossDomainAdmin common.Address, _hubPool common.Address, _wrappedNativeTokenAddress common.Address) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "__SpokePool_init", _initialDepositId, _crossDomainAdmin, _hubPool, _wrappedNativeTokenAddress)
}

// SpokePoolInit is a paid mutator transaction binding the contract method 0x3ddf5059.
//
// Solidity: function __SpokePool_init(uint32 _initialDepositId, address _crossDomainAdmin, address _hubPool, address _wrappedNativeTokenAddress) returns()
func (_Router2 *Router2Session) SpokePoolInit(_initialDepositId uint32, _crossDomainAdmin common.Address, _hubPool common.Address, _wrappedNativeTokenAddress common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SpokePoolInit(&_Router2.TransactOpts, _initialDepositId, _crossDomainAdmin, _hubPool, _wrappedNativeTokenAddress)
}

// SpokePoolInit is a paid mutator transaction binding the contract method 0x3ddf5059.
//
// Solidity: function __SpokePool_init(uint32 _initialDepositId, address _crossDomainAdmin, address _hubPool, address _wrappedNativeTokenAddress) returns()
func (_Router2 *Router2TransactorSession) SpokePoolInit(_initialDepositId uint32, _crossDomainAdmin common.Address, _hubPool common.Address, _wrappedNativeTokenAddress common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SpokePoolInit(&_Router2.TransactOpts, _initialDepositId, _crossDomainAdmin, _hubPool, _wrappedNativeTokenAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x1186ec33.
//
// Solidity: function deposit(address recipient, address originToken, uint256 amount, uint256 destinationChainId, int64 relayerFeePct, uint32 quoteTimestamp, bytes message, uint256 maxCount) payable returns()
func (_Router2 *Router2Transactor) Deposit(opts *bind.TransactOpts, recipient common.Address, originToken common.Address, amount *big.Int, destinationChainId *big.Int, relayerFeePct int64, quoteTimestamp uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "deposit", recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1186ec33.
//
// Solidity: function deposit(address recipient, address originToken, uint256 amount, uint256 destinationChainId, int64 relayerFeePct, uint32 quoteTimestamp, bytes message, uint256 maxCount) payable returns()
func (_Router2 *Router2Session) Deposit(recipient common.Address, originToken common.Address, amount *big.Int, destinationChainId *big.Int, relayerFeePct int64, quoteTimestamp uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.Deposit(&_Router2.TransactOpts, recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1186ec33.
//
// Solidity: function deposit(address recipient, address originToken, uint256 amount, uint256 destinationChainId, int64 relayerFeePct, uint32 quoteTimestamp, bytes message, uint256 maxCount) payable returns()
func (_Router2 *Router2TransactorSession) Deposit(recipient common.Address, originToken common.Address, amount *big.Int, destinationChainId *big.Int, relayerFeePct int64, quoteTimestamp uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.Deposit(&_Router2.TransactOpts, recipient, originToken, amount, destinationChainId, relayerFeePct, quoteTimestamp, message, maxCount)
}

// EmergencyDeleteRootBundle is a paid mutator transaction binding the contract method 0x8a7860ce.
//
// Solidity: function emergencyDeleteRootBundle(uint256 rootBundleId) returns()
func (_Router2 *Router2Transactor) EmergencyDeleteRootBundle(opts *bind.TransactOpts, rootBundleId *big.Int) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "emergencyDeleteRootBundle", rootBundleId)
}

// EmergencyDeleteRootBundle is a paid mutator transaction binding the contract method 0x8a7860ce.
//
// Solidity: function emergencyDeleteRootBundle(uint256 rootBundleId) returns()
func (_Router2 *Router2Session) EmergencyDeleteRootBundle(rootBundleId *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.EmergencyDeleteRootBundle(&_Router2.TransactOpts, rootBundleId)
}

// EmergencyDeleteRootBundle is a paid mutator transaction binding the contract method 0x8a7860ce.
//
// Solidity: function emergencyDeleteRootBundle(uint256 rootBundleId) returns()
func (_Router2 *Router2TransactorSession) EmergencyDeleteRootBundle(rootBundleId *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.EmergencyDeleteRootBundle(&_Router2.TransactOpts, rootBundleId)
}

// ExecuteRelayerRefundLeaf is a paid mutator transaction binding the contract method 0x1b3d5559.
//
// Solidity: function executeRelayerRefundLeaf(uint32 rootBundleId, (uint256,uint256,uint256[],uint32,address,address[]) relayerRefundLeaf, bytes32[] proof) returns()
func (_Router2 *Router2Transactor) ExecuteRelayerRefundLeaf(opts *bind.TransactOpts, rootBundleId uint32, relayerRefundLeaf SpokePoolInterfaceRelayerRefundLeaf, proof [][32]byte) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "executeRelayerRefundLeaf", rootBundleId, relayerRefundLeaf, proof)
}

// ExecuteRelayerRefundLeaf is a paid mutator transaction binding the contract method 0x1b3d5559.
//
// Solidity: function executeRelayerRefundLeaf(uint32 rootBundleId, (uint256,uint256,uint256[],uint32,address,address[]) relayerRefundLeaf, bytes32[] proof) returns()
func (_Router2 *Router2Session) ExecuteRelayerRefundLeaf(rootBundleId uint32, relayerRefundLeaf SpokePoolInterfaceRelayerRefundLeaf, proof [][32]byte) (*types.Transaction, error) {
	return _Router2.Contract.ExecuteRelayerRefundLeaf(&_Router2.TransactOpts, rootBundleId, relayerRefundLeaf, proof)
}

// ExecuteRelayerRefundLeaf is a paid mutator transaction binding the contract method 0x1b3d5559.
//
// Solidity: function executeRelayerRefundLeaf(uint32 rootBundleId, (uint256,uint256,uint256[],uint32,address,address[]) relayerRefundLeaf, bytes32[] proof) returns()
func (_Router2 *Router2TransactorSession) ExecuteRelayerRefundLeaf(rootBundleId uint32, relayerRefundLeaf SpokePoolInterfaceRelayerRefundLeaf, proof [][32]byte) (*types.Transaction, error) {
	return _Router2.Contract.ExecuteRelayerRefundLeaf(&_Router2.TransactOpts, rootBundleId, relayerRefundLeaf, proof)
}

// ExecuteSlowRelayLeaf is a paid mutator transaction binding the contract method 0x0c2097f7.
//
// Solidity: function executeSlowRelayLeaf(address depositor, address recipient, address destinationToken, uint256 totalRelayAmount, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, uint32 rootBundleId, bytes message, int256 payoutAdjustment, bytes32[] proof) returns()
func (_Router2 *Router2Transactor) ExecuteSlowRelayLeaf(opts *bind.TransactOpts, depositor common.Address, recipient common.Address, destinationToken common.Address, totalRelayAmount *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, rootBundleId uint32, message []byte, payoutAdjustment *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "executeSlowRelayLeaf", depositor, recipient, destinationToken, totalRelayAmount, originChainId, realizedLpFeePct, relayerFeePct, depositId, rootBundleId, message, payoutAdjustment, proof)
}

// ExecuteSlowRelayLeaf is a paid mutator transaction binding the contract method 0x0c2097f7.
//
// Solidity: function executeSlowRelayLeaf(address depositor, address recipient, address destinationToken, uint256 totalRelayAmount, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, uint32 rootBundleId, bytes message, int256 payoutAdjustment, bytes32[] proof) returns()
func (_Router2 *Router2Session) ExecuteSlowRelayLeaf(depositor common.Address, recipient common.Address, destinationToken common.Address, totalRelayAmount *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, rootBundleId uint32, message []byte, payoutAdjustment *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Router2.Contract.ExecuteSlowRelayLeaf(&_Router2.TransactOpts, depositor, recipient, destinationToken, totalRelayAmount, originChainId, realizedLpFeePct, relayerFeePct, depositId, rootBundleId, message, payoutAdjustment, proof)
}

// ExecuteSlowRelayLeaf is a paid mutator transaction binding the contract method 0x0c2097f7.
//
// Solidity: function executeSlowRelayLeaf(address depositor, address recipient, address destinationToken, uint256 totalRelayAmount, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, uint32 rootBundleId, bytes message, int256 payoutAdjustment, bytes32[] proof) returns()
func (_Router2 *Router2TransactorSession) ExecuteSlowRelayLeaf(depositor common.Address, recipient common.Address, destinationToken common.Address, totalRelayAmount *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, rootBundleId uint32, message []byte, payoutAdjustment *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Router2.Contract.ExecuteSlowRelayLeaf(&_Router2.TransactOpts, depositor, recipient, destinationToken, totalRelayAmount, originChainId, realizedLpFeePct, relayerFeePct, depositId, rootBundleId, message, payoutAdjustment, proof)
}

// FillRelay is a paid mutator transaction binding the contract method 0x44b8be68.
//
// Solidity: function fillRelay(address depositor, address recipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, bytes message, uint256 maxCount) returns()
func (_Router2 *Router2Transactor) FillRelay(opts *bind.TransactOpts, depositor common.Address, recipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "fillRelay", depositor, recipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, depositId, message, maxCount)
}

// FillRelay is a paid mutator transaction binding the contract method 0x44b8be68.
//
// Solidity: function fillRelay(address depositor, address recipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, bytes message, uint256 maxCount) returns()
func (_Router2 *Router2Session) FillRelay(depositor common.Address, recipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.FillRelay(&_Router2.TransactOpts, depositor, recipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, depositId, message, maxCount)
}

// FillRelay is a paid mutator transaction binding the contract method 0x44b8be68.
//
// Solidity: function fillRelay(address depositor, address recipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, uint32 depositId, bytes message, uint256 maxCount) returns()
func (_Router2 *Router2TransactorSession) FillRelay(depositor common.Address, recipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, depositId uint32, message []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.FillRelay(&_Router2.TransactOpts, depositor, recipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, depositId, message, maxCount)
}

// FillRelayWithUpdatedDeposit is a paid mutator transaction binding the contract method 0x5ceaec32.
//
// Solidity: function fillRelayWithUpdatedDeposit(address depositor, address recipient, address updatedRecipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, int64 updatedRelayerFeePct, uint32 depositId, bytes message, bytes updatedMessage, bytes depositorSignature, uint256 maxCount) returns()
func (_Router2 *Router2Transactor) FillRelayWithUpdatedDeposit(opts *bind.TransactOpts, depositor common.Address, recipient common.Address, updatedRecipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, updatedRelayerFeePct int64, depositId uint32, message []byte, updatedMessage []byte, depositorSignature []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "fillRelayWithUpdatedDeposit", depositor, recipient, updatedRecipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, updatedRelayerFeePct, depositId, message, updatedMessage, depositorSignature, maxCount)
}

// FillRelayWithUpdatedDeposit is a paid mutator transaction binding the contract method 0x5ceaec32.
//
// Solidity: function fillRelayWithUpdatedDeposit(address depositor, address recipient, address updatedRecipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, int64 updatedRelayerFeePct, uint32 depositId, bytes message, bytes updatedMessage, bytes depositorSignature, uint256 maxCount) returns()
func (_Router2 *Router2Session) FillRelayWithUpdatedDeposit(depositor common.Address, recipient common.Address, updatedRecipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, updatedRelayerFeePct int64, depositId uint32, message []byte, updatedMessage []byte, depositorSignature []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.FillRelayWithUpdatedDeposit(&_Router2.TransactOpts, depositor, recipient, updatedRecipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, updatedRelayerFeePct, depositId, message, updatedMessage, depositorSignature, maxCount)
}

// FillRelayWithUpdatedDeposit is a paid mutator transaction binding the contract method 0x5ceaec32.
//
// Solidity: function fillRelayWithUpdatedDeposit(address depositor, address recipient, address updatedRecipient, address destinationToken, uint256 amount, uint256 maxTokensToSend, uint256 repaymentChainId, uint256 originChainId, int64 realizedLpFeePct, int64 relayerFeePct, int64 updatedRelayerFeePct, uint32 depositId, bytes message, bytes updatedMessage, bytes depositorSignature, uint256 maxCount) returns()
func (_Router2 *Router2TransactorSession) FillRelayWithUpdatedDeposit(depositor common.Address, recipient common.Address, updatedRecipient common.Address, destinationToken common.Address, amount *big.Int, maxTokensToSend *big.Int, repaymentChainId *big.Int, originChainId *big.Int, realizedLpFeePct int64, relayerFeePct int64, updatedRelayerFeePct int64, depositId uint32, message []byte, updatedMessage []byte, depositorSignature []byte, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.FillRelayWithUpdatedDeposit(&_Router2.TransactOpts, depositor, recipient, updatedRecipient, destinationToken, amount, maxTokensToSend, repaymentChainId, originChainId, realizedLpFeePct, relayerFeePct, updatedRelayerFeePct, depositId, message, updatedMessage, depositorSignature, maxCount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf2c20924.
//
// Solidity: function initialize(uint32 _initialDepositId, address _zkErc20Bridge, address _crossDomainAdmin, address _hubPool, address _wethAddress) returns()
func (_Router2 *Router2Transactor) Initialize(opts *bind.TransactOpts, _initialDepositId uint32, _zkErc20Bridge common.Address, _crossDomainAdmin common.Address, _hubPool common.Address, _wethAddress common.Address) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "initialize", _initialDepositId, _zkErc20Bridge, _crossDomainAdmin, _hubPool, _wethAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf2c20924.
//
// Solidity: function initialize(uint32 _initialDepositId, address _zkErc20Bridge, address _crossDomainAdmin, address _hubPool, address _wethAddress) returns()
func (_Router2 *Router2Session) Initialize(_initialDepositId uint32, _zkErc20Bridge common.Address, _crossDomainAdmin common.Address, _hubPool common.Address, _wethAddress common.Address) (*types.Transaction, error) {
	return _Router2.Contract.Initialize(&_Router2.TransactOpts, _initialDepositId, _zkErc20Bridge, _crossDomainAdmin, _hubPool, _wethAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf2c20924.
//
// Solidity: function initialize(uint32 _initialDepositId, address _zkErc20Bridge, address _crossDomainAdmin, address _hubPool, address _wethAddress) returns()
func (_Router2 *Router2TransactorSession) Initialize(_initialDepositId uint32, _zkErc20Bridge common.Address, _crossDomainAdmin common.Address, _hubPool common.Address, _wethAddress common.Address) (*types.Transaction, error) {
	return _Router2.Contract.Initialize(&_Router2.TransactOpts, _initialDepositId, _zkErc20Bridge, _crossDomainAdmin, _hubPool, _wethAddress)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Router2 *Router2Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Router2 *Router2Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router2.Contract.Multicall(&_Router2.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Router2 *Router2TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router2.Contract.Multicall(&_Router2.TransactOpts, data)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x738b62e5.
//
// Solidity: function pauseDeposits(bool pause) returns()
func (_Router2 *Router2Transactor) PauseDeposits(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "pauseDeposits", pause)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x738b62e5.
//
// Solidity: function pauseDeposits(bool pause) returns()
func (_Router2 *Router2Session) PauseDeposits(pause bool) (*types.Transaction, error) {
	return _Router2.Contract.PauseDeposits(&_Router2.TransactOpts, pause)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x738b62e5.
//
// Solidity: function pauseDeposits(bool pause) returns()
func (_Router2 *Router2TransactorSession) PauseDeposits(pause bool) (*types.Transaction, error) {
	return _Router2.Contract.PauseDeposits(&_Router2.TransactOpts, pause)
}

// PauseFills is a paid mutator transaction binding the contract method 0x99cc2968.
//
// Solidity: function pauseFills(bool pause) returns()
func (_Router2 *Router2Transactor) PauseFills(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "pauseFills", pause)
}

// PauseFills is a paid mutator transaction binding the contract method 0x99cc2968.
//
// Solidity: function pauseFills(bool pause) returns()
func (_Router2 *Router2Session) PauseFills(pause bool) (*types.Transaction, error) {
	return _Router2.Contract.PauseFills(&_Router2.TransactOpts, pause)
}

// PauseFills is a paid mutator transaction binding the contract method 0x99cc2968.
//
// Solidity: function pauseFills(bool pause) returns()
func (_Router2 *Router2TransactorSession) PauseFills(pause bool) (*types.Transaction, error) {
	return _Router2.Contract.PauseFills(&_Router2.TransactOpts, pause)
}

// RelayRootBundle is a paid mutator transaction binding the contract method 0x493a4f84.
//
// Solidity: function relayRootBundle(bytes32 relayerRefundRoot, bytes32 slowRelayRoot) returns()
func (_Router2 *Router2Transactor) RelayRootBundle(opts *bind.TransactOpts, relayerRefundRoot [32]byte, slowRelayRoot [32]byte) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "relayRootBundle", relayerRefundRoot, slowRelayRoot)
}

// RelayRootBundle is a paid mutator transaction binding the contract method 0x493a4f84.
//
// Solidity: function relayRootBundle(bytes32 relayerRefundRoot, bytes32 slowRelayRoot) returns()
func (_Router2 *Router2Session) RelayRootBundle(relayerRefundRoot [32]byte, slowRelayRoot [32]byte) (*types.Transaction, error) {
	return _Router2.Contract.RelayRootBundle(&_Router2.TransactOpts, relayerRefundRoot, slowRelayRoot)
}

// RelayRootBundle is a paid mutator transaction binding the contract method 0x493a4f84.
//
// Solidity: function relayRootBundle(bytes32 relayerRefundRoot, bytes32 slowRelayRoot) returns()
func (_Router2 *Router2TransactorSession) RelayRootBundle(relayerRefundRoot [32]byte, slowRelayRoot [32]byte) (*types.Transaction, error) {
	return _Router2.Contract.RelayRootBundle(&_Router2.TransactOpts, relayerRefundRoot, slowRelayRoot)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa634893f.
//
// Solidity: function requestRefund(address refundToken, uint256 amount, uint256 originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 depositId, uint256 fillBlock, uint256 maxCount) returns()
func (_Router2 *Router2Transactor) RequestRefund(opts *bind.TransactOpts, refundToken common.Address, amount *big.Int, originChainId *big.Int, destinationChainId *big.Int, realizedLpFeePct int64, depositId uint32, fillBlock *big.Int, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "requestRefund", refundToken, amount, originChainId, destinationChainId, realizedLpFeePct, depositId, fillBlock, maxCount)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa634893f.
//
// Solidity: function requestRefund(address refundToken, uint256 amount, uint256 originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 depositId, uint256 fillBlock, uint256 maxCount) returns()
func (_Router2 *Router2Session) RequestRefund(refundToken common.Address, amount *big.Int, originChainId *big.Int, destinationChainId *big.Int, realizedLpFeePct int64, depositId uint32, fillBlock *big.Int, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.RequestRefund(&_Router2.TransactOpts, refundToken, amount, originChainId, destinationChainId, realizedLpFeePct, depositId, fillBlock, maxCount)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xa634893f.
//
// Solidity: function requestRefund(address refundToken, uint256 amount, uint256 originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 depositId, uint256 fillBlock, uint256 maxCount) returns()
func (_Router2 *Router2TransactorSession) RequestRefund(refundToken common.Address, amount *big.Int, originChainId *big.Int, destinationChainId *big.Int, realizedLpFeePct int64, depositId uint32, fillBlock *big.Int, maxCount *big.Int) (*types.Transaction, error) {
	return _Router2.Contract.RequestRefund(&_Router2.TransactOpts, refundToken, amount, originChainId, destinationChainId, realizedLpFeePct, depositId, fillBlock, maxCount)
}

// SetCrossDomainAdmin is a paid mutator transaction binding the contract method 0xde7eba78.
//
// Solidity: function setCrossDomainAdmin(address newCrossDomainAdmin) returns()
func (_Router2 *Router2Transactor) SetCrossDomainAdmin(opts *bind.TransactOpts, newCrossDomainAdmin common.Address) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "setCrossDomainAdmin", newCrossDomainAdmin)
}

// SetCrossDomainAdmin is a paid mutator transaction binding the contract method 0xde7eba78.
//
// Solidity: function setCrossDomainAdmin(address newCrossDomainAdmin) returns()
func (_Router2 *Router2Session) SetCrossDomainAdmin(newCrossDomainAdmin common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SetCrossDomainAdmin(&_Router2.TransactOpts, newCrossDomainAdmin)
}

// SetCrossDomainAdmin is a paid mutator transaction binding the contract method 0xde7eba78.
//
// Solidity: function setCrossDomainAdmin(address newCrossDomainAdmin) returns()
func (_Router2 *Router2TransactorSession) SetCrossDomainAdmin(newCrossDomainAdmin common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SetCrossDomainAdmin(&_Router2.TransactOpts, newCrossDomainAdmin)
}

// SetDepositQuoteTimeBuffer is a paid mutator transaction binding the contract method 0x2752042e.
//
// Solidity: function setDepositQuoteTimeBuffer(uint32 newDepositQuoteTimeBuffer) returns()
func (_Router2 *Router2Transactor) SetDepositQuoteTimeBuffer(opts *bind.TransactOpts, newDepositQuoteTimeBuffer uint32) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "setDepositQuoteTimeBuffer", newDepositQuoteTimeBuffer)
}

// SetDepositQuoteTimeBuffer is a paid mutator transaction binding the contract method 0x2752042e.
//
// Solidity: function setDepositQuoteTimeBuffer(uint32 newDepositQuoteTimeBuffer) returns()
func (_Router2 *Router2Session) SetDepositQuoteTimeBuffer(newDepositQuoteTimeBuffer uint32) (*types.Transaction, error) {
	return _Router2.Contract.SetDepositQuoteTimeBuffer(&_Router2.TransactOpts, newDepositQuoteTimeBuffer)
}

// SetDepositQuoteTimeBuffer is a paid mutator transaction binding the contract method 0x2752042e.
//
// Solidity: function setDepositQuoteTimeBuffer(uint32 newDepositQuoteTimeBuffer) returns()
func (_Router2 *Router2TransactorSession) SetDepositQuoteTimeBuffer(newDepositQuoteTimeBuffer uint32) (*types.Transaction, error) {
	return _Router2.Contract.SetDepositQuoteTimeBuffer(&_Router2.TransactOpts, newDepositQuoteTimeBuffer)
}

// SetEnableRoute is a paid mutator transaction binding the contract method 0x272751c7.
//
// Solidity: function setEnableRoute(address originToken, uint256 destinationChainId, bool enabled) returns()
func (_Router2 *Router2Transactor) SetEnableRoute(opts *bind.TransactOpts, originToken common.Address, destinationChainId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "setEnableRoute", originToken, destinationChainId, enabled)
}

// SetEnableRoute is a paid mutator transaction binding the contract method 0x272751c7.
//
// Solidity: function setEnableRoute(address originToken, uint256 destinationChainId, bool enabled) returns()
func (_Router2 *Router2Session) SetEnableRoute(originToken common.Address, destinationChainId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Router2.Contract.SetEnableRoute(&_Router2.TransactOpts, originToken, destinationChainId, enabled)
}

// SetEnableRoute is a paid mutator transaction binding the contract method 0x272751c7.
//
// Solidity: function setEnableRoute(address originToken, uint256 destinationChainId, bool enabled) returns()
func (_Router2 *Router2TransactorSession) SetEnableRoute(originToken common.Address, destinationChainId *big.Int, enabled bool) (*types.Transaction, error) {
	return _Router2.Contract.SetEnableRoute(&_Router2.TransactOpts, originToken, destinationChainId, enabled)
}

// SetHubPool is a paid mutator transaction binding the contract method 0x1dfb2d02.
//
// Solidity: function setHubPool(address newHubPool) returns()
func (_Router2 *Router2Transactor) SetHubPool(opts *bind.TransactOpts, newHubPool common.Address) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "setHubPool", newHubPool)
}

// SetHubPool is a paid mutator transaction binding the contract method 0x1dfb2d02.
//
// Solidity: function setHubPool(address newHubPool) returns()
func (_Router2 *Router2Session) SetHubPool(newHubPool common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SetHubPool(&_Router2.TransactOpts, newHubPool)
}

// SetHubPool is a paid mutator transaction binding the contract method 0x1dfb2d02.
//
// Solidity: function setHubPool(address newHubPool) returns()
func (_Router2 *Router2TransactorSession) SetHubPool(newHubPool common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SetHubPool(&_Router2.TransactOpts, newHubPool)
}

// SetZkBridge is a paid mutator transaction binding the contract method 0xcd44141f.
//
// Solidity: function setZkBridge(address _zkErc20Bridge) returns()
func (_Router2 *Router2Transactor) SetZkBridge(opts *bind.TransactOpts, _zkErc20Bridge common.Address) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "setZkBridge", _zkErc20Bridge)
}

// SetZkBridge is a paid mutator transaction binding the contract method 0xcd44141f.
//
// Solidity: function setZkBridge(address _zkErc20Bridge) returns()
func (_Router2 *Router2Session) SetZkBridge(_zkErc20Bridge common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SetZkBridge(&_Router2.TransactOpts, _zkErc20Bridge)
}

// SetZkBridge is a paid mutator transaction binding the contract method 0xcd44141f.
//
// Solidity: function setZkBridge(address _zkErc20Bridge) returns()
func (_Router2 *Router2TransactorSession) SetZkBridge(_zkErc20Bridge common.Address) (*types.Transaction, error) {
	return _Router2.Contract.SetZkBridge(&_Router2.TransactOpts, _zkErc20Bridge)
}

// SpeedUpDeposit is a paid mutator transaction binding the contract method 0x7e688bba.
//
// Solidity: function speedUpDeposit(address depositor, int64 updatedRelayerFeePct, uint32 depositId, address updatedRecipient, bytes updatedMessage, bytes depositorSignature) returns()
func (_Router2 *Router2Transactor) SpeedUpDeposit(opts *bind.TransactOpts, depositor common.Address, updatedRelayerFeePct int64, depositId uint32, updatedRecipient common.Address, updatedMessage []byte, depositorSignature []byte) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "speedUpDeposit", depositor, updatedRelayerFeePct, depositId, updatedRecipient, updatedMessage, depositorSignature)
}

// SpeedUpDeposit is a paid mutator transaction binding the contract method 0x7e688bba.
//
// Solidity: function speedUpDeposit(address depositor, int64 updatedRelayerFeePct, uint32 depositId, address updatedRecipient, bytes updatedMessage, bytes depositorSignature) returns()
func (_Router2 *Router2Session) SpeedUpDeposit(depositor common.Address, updatedRelayerFeePct int64, depositId uint32, updatedRecipient common.Address, updatedMessage []byte, depositorSignature []byte) (*types.Transaction, error) {
	return _Router2.Contract.SpeedUpDeposit(&_Router2.TransactOpts, depositor, updatedRelayerFeePct, depositId, updatedRecipient, updatedMessage, depositorSignature)
}

// SpeedUpDeposit is a paid mutator transaction binding the contract method 0x7e688bba.
//
// Solidity: function speedUpDeposit(address depositor, int64 updatedRelayerFeePct, uint32 depositId, address updatedRecipient, bytes updatedMessage, bytes depositorSignature) returns()
func (_Router2 *Router2TransactorSession) SpeedUpDeposit(depositor common.Address, updatedRelayerFeePct int64, depositId uint32, updatedRecipient common.Address, updatedMessage []byte, depositorSignature []byte) (*types.Transaction, error) {
	return _Router2.Contract.SpeedUpDeposit(&_Router2.TransactOpts, depositor, updatedRelayerFeePct, depositId, updatedRecipient, updatedMessage, depositorSignature)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Router2 *Router2Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Router2 *Router2Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Router2.Contract.UpgradeTo(&_Router2.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Router2 *Router2TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Router2.Contract.UpgradeTo(&_Router2.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Router2 *Router2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Router2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Router2 *Router2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Router2.Contract.UpgradeToAndCall(&_Router2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Router2 *Router2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Router2.Contract.UpgradeToAndCall(&_Router2.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router2 *Router2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router2 *Router2Session) Receive() (*types.Transaction, error) {
	return _Router2.Contract.Receive(&_Router2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router2 *Router2TransactorSession) Receive() (*types.Transaction, error) {
	return _Router2.Contract.Receive(&_Router2.TransactOpts)
}

// Router2AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Router2 contract.
type Router2AdminChangedIterator struct {
	Event *Router2AdminChanged // Event containing the contract specifics and raw log

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
func (it *Router2AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2AdminChanged)
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
		it.Event = new(Router2AdminChanged)
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
func (it *Router2AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2AdminChanged represents a AdminChanged event raised by the Router2 contract.
type Router2AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Router2 *Router2Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*Router2AdminChangedIterator, error) {

	logs, sub, err := _Router2.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &Router2AdminChangedIterator{contract: _Router2.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Router2 *Router2Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *Router2AdminChanged) (event.Subscription, error) {

	logs, sub, err := _Router2.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2AdminChanged)
				if err := _Router2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseAdminChanged(log types.Log) (*Router2AdminChanged, error) {
	event := new(Router2AdminChanged)
	if err := _Router2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Router2 contract.
type Router2BeaconUpgradedIterator struct {
	Event *Router2BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *Router2BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2BeaconUpgraded)
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
		it.Event = new(Router2BeaconUpgraded)
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
func (it *Router2BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2BeaconUpgraded represents a BeaconUpgraded event raised by the Router2 contract.
type Router2BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Router2 *Router2Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*Router2BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &Router2BeaconUpgradedIterator{contract: _Router2.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Router2 *Router2Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *Router2BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2BeaconUpgraded)
				if err := _Router2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseBeaconUpgraded(log types.Log) (*Router2BeaconUpgraded, error) {
	event := new(Router2BeaconUpgraded)
	if err := _Router2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2EmergencyDeleteRootBundleIterator is returned from FilterEmergencyDeleteRootBundle and is used to iterate over the raw logs and unpacked data for EmergencyDeleteRootBundle events raised by the Router2 contract.
type Router2EmergencyDeleteRootBundleIterator struct {
	Event *Router2EmergencyDeleteRootBundle // Event containing the contract specifics and raw log

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
func (it *Router2EmergencyDeleteRootBundleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2EmergencyDeleteRootBundle)
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
		it.Event = new(Router2EmergencyDeleteRootBundle)
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
func (it *Router2EmergencyDeleteRootBundleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2EmergencyDeleteRootBundleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2EmergencyDeleteRootBundle represents a EmergencyDeleteRootBundle event raised by the Router2 contract.
type Router2EmergencyDeleteRootBundle struct {
	RootBundleId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyDeleteRootBundle is a free log retrieval operation binding the contract event 0x3569b846531b754c99cb80df3f49cd72fa6fe106aaee5ab8e0caf35a9d7ce88d.
//
// Solidity: event EmergencyDeleteRootBundle(uint256 indexed rootBundleId)
func (_Router2 *Router2Filterer) FilterEmergencyDeleteRootBundle(opts *bind.FilterOpts, rootBundleId []*big.Int) (*Router2EmergencyDeleteRootBundleIterator, error) {

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "EmergencyDeleteRootBundle", rootBundleIdRule)
	if err != nil {
		return nil, err
	}
	return &Router2EmergencyDeleteRootBundleIterator{contract: _Router2.contract, event: "EmergencyDeleteRootBundle", logs: logs, sub: sub}, nil
}

// WatchEmergencyDeleteRootBundle is a free log subscription operation binding the contract event 0x3569b846531b754c99cb80df3f49cd72fa6fe106aaee5ab8e0caf35a9d7ce88d.
//
// Solidity: event EmergencyDeleteRootBundle(uint256 indexed rootBundleId)
func (_Router2 *Router2Filterer) WatchEmergencyDeleteRootBundle(opts *bind.WatchOpts, sink chan<- *Router2EmergencyDeleteRootBundle, rootBundleId []*big.Int) (event.Subscription, error) {

	var rootBundleIdRule []interface{}
	for _, rootBundleIdItem := range rootBundleId {
		rootBundleIdRule = append(rootBundleIdRule, rootBundleIdItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "EmergencyDeleteRootBundle", rootBundleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2EmergencyDeleteRootBundle)
				if err := _Router2.contract.UnpackLog(event, "EmergencyDeleteRootBundle", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseEmergencyDeleteRootBundle(log types.Log) (*Router2EmergencyDeleteRootBundle, error) {
	event := new(Router2EmergencyDeleteRootBundle)
	if err := _Router2.contract.UnpackLog(event, "EmergencyDeleteRootBundle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2EnabledDepositRouteIterator is returned from FilterEnabledDepositRoute and is used to iterate over the raw logs and unpacked data for EnabledDepositRoute events raised by the Router2 contract.
type Router2EnabledDepositRouteIterator struct {
	Event *Router2EnabledDepositRoute // Event containing the contract specifics and raw log

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
func (it *Router2EnabledDepositRouteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2EnabledDepositRoute)
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
		it.Event = new(Router2EnabledDepositRoute)
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
func (it *Router2EnabledDepositRouteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2EnabledDepositRouteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2EnabledDepositRoute represents a EnabledDepositRoute event raised by the Router2 contract.
type Router2EnabledDepositRoute struct {
	OriginToken        common.Address
	DestinationChainId *big.Int
	Enabled            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEnabledDepositRoute is a free log retrieval operation binding the contract event 0x0a21fdd43d0ad0c62689ee7230a47309a050755bcc52eba00310add65297692a.
//
// Solidity: event EnabledDepositRoute(address indexed originToken, uint256 indexed destinationChainId, bool enabled)
func (_Router2 *Router2Filterer) FilterEnabledDepositRoute(opts *bind.FilterOpts, originToken []common.Address, destinationChainId []*big.Int) (*Router2EnabledDepositRouteIterator, error) {

	var originTokenRule []interface{}
	for _, originTokenItem := range originToken {
		originTokenRule = append(originTokenRule, originTokenItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "EnabledDepositRoute", originTokenRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &Router2EnabledDepositRouteIterator{contract: _Router2.contract, event: "EnabledDepositRoute", logs: logs, sub: sub}, nil
}

// WatchEnabledDepositRoute is a free log subscription operation binding the contract event 0x0a21fdd43d0ad0c62689ee7230a47309a050755bcc52eba00310add65297692a.
//
// Solidity: event EnabledDepositRoute(address indexed originToken, uint256 indexed destinationChainId, bool enabled)
func (_Router2 *Router2Filterer) WatchEnabledDepositRoute(opts *bind.WatchOpts, sink chan<- *Router2EnabledDepositRoute, originToken []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var originTokenRule []interface{}
	for _, originTokenItem := range originToken {
		originTokenRule = append(originTokenRule, originTokenItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "EnabledDepositRoute", originTokenRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2EnabledDepositRoute)
				if err := _Router2.contract.UnpackLog(event, "EnabledDepositRoute", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseEnabledDepositRoute(log types.Log) (*Router2EnabledDepositRoute, error) {
	event := new(Router2EnabledDepositRoute)
	if err := _Router2.contract.UnpackLog(event, "EnabledDepositRoute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2ExecutedRelayerRefundRootIterator is returned from FilterExecutedRelayerRefundRoot and is used to iterate over the raw logs and unpacked data for ExecutedRelayerRefundRoot events raised by the Router2 contract.
type Router2ExecutedRelayerRefundRootIterator struct {
	Event *Router2ExecutedRelayerRefundRoot // Event containing the contract specifics and raw log

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
func (it *Router2ExecutedRelayerRefundRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2ExecutedRelayerRefundRoot)
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
		it.Event = new(Router2ExecutedRelayerRefundRoot)
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
func (it *Router2ExecutedRelayerRefundRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2ExecutedRelayerRefundRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2ExecutedRelayerRefundRoot represents a ExecutedRelayerRefundRoot event raised by the Router2 contract.
type Router2ExecutedRelayerRefundRoot struct {
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
func (_Router2 *Router2Filterer) FilterExecutedRelayerRefundRoot(opts *bind.FilterOpts, chainId []*big.Int, rootBundleId []uint32, leafId []uint32) (*Router2ExecutedRelayerRefundRootIterator, error) {

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

	logs, sub, err := _Router2.contract.FilterLogs(opts, "ExecutedRelayerRefundRoot", chainIdRule, rootBundleIdRule, leafIdRule)
	if err != nil {
		return nil, err
	}
	return &Router2ExecutedRelayerRefundRootIterator{contract: _Router2.contract, event: "ExecutedRelayerRefundRoot", logs: logs, sub: sub}, nil
}

// WatchExecutedRelayerRefundRoot is a free log subscription operation binding the contract event 0xf8bd640004bcec1b89657020f561d0b070cbdf662d0b158db9dccb0a8301bfab.
//
// Solidity: event ExecutedRelayerRefundRoot(uint256 amountToReturn, uint256 indexed chainId, uint256[] refundAmounts, uint32 indexed rootBundleId, uint32 indexed leafId, address l2TokenAddress, address[] refundAddresses, address caller)
func (_Router2 *Router2Filterer) WatchExecutedRelayerRefundRoot(opts *bind.WatchOpts, sink chan<- *Router2ExecutedRelayerRefundRoot, chainId []*big.Int, rootBundleId []uint32, leafId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _Router2.contract.WatchLogs(opts, "ExecutedRelayerRefundRoot", chainIdRule, rootBundleIdRule, leafIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2ExecutedRelayerRefundRoot)
				if err := _Router2.contract.UnpackLog(event, "ExecutedRelayerRefundRoot", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseExecutedRelayerRefundRoot(log types.Log) (*Router2ExecutedRelayerRefundRoot, error) {
	event := new(Router2ExecutedRelayerRefundRoot)
	if err := _Router2.contract.UnpackLog(event, "ExecutedRelayerRefundRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2FilledRelayIterator is returned from FilterFilledRelay and is used to iterate over the raw logs and unpacked data for FilledRelay events raised by the Router2 contract.
type Router2FilledRelayIterator struct {
	Event *Router2FilledRelay // Event containing the contract specifics and raw log

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
func (it *Router2FilledRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2FilledRelay)
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
		it.Event = new(Router2FilledRelay)
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
func (it *Router2FilledRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2FilledRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2FilledRelay represents a FilledRelay event raised by the Router2 contract.
type Router2FilledRelay struct {
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
func (_Router2 *Router2Filterer) FilterFilledRelay(opts *bind.FilterOpts, originChainId []*big.Int, depositId []uint32, depositor []common.Address) (*Router2FilledRelayIterator, error) {

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

	logs, sub, err := _Router2.contract.FilterLogs(opts, "FilledRelay", originChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &Router2FilledRelayIterator{contract: _Router2.contract, event: "FilledRelay", logs: logs, sub: sub}, nil
}

// WatchFilledRelay is a free log subscription operation binding the contract event 0x8ab9dc6c19fe88e69bc70221b339c84332752fdd49591b7c51e66bae3947b73c.
//
// Solidity: event FilledRelay(uint256 amount, uint256 totalFilledAmount, uint256 fillAmount, uint256 repaymentChainId, uint256 indexed originChainId, uint256 destinationChainId, int64 relayerFeePct, int64 realizedLpFeePct, uint32 indexed depositId, address destinationToken, address relayer, address indexed depositor, address recipient, bytes message, (address,bytes,int64,bool,int256) updatableRelayData)
func (_Router2 *Router2Filterer) WatchFilledRelay(opts *bind.WatchOpts, sink chan<- *Router2FilledRelay, originChainId []*big.Int, depositId []uint32, depositor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Router2.contract.WatchLogs(opts, "FilledRelay", originChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2FilledRelay)
				if err := _Router2.contract.UnpackLog(event, "FilledRelay", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseFilledRelay(log types.Log) (*Router2FilledRelay, error) {
	event := new(Router2FilledRelay)
	if err := _Router2.contract.UnpackLog(event, "FilledRelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2FundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the Router2 contract.
type Router2FundsDepositedIterator struct {
	Event *Router2FundsDeposited // Event containing the contract specifics and raw log

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
func (it *Router2FundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2FundsDeposited)
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
		it.Event = new(Router2FundsDeposited)
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
func (it *Router2FundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2FundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2FundsDeposited represents a FundsDeposited event raised by the Router2 contract.
type Router2FundsDeposited struct {
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
func (_Router2 *Router2Filterer) FilterFundsDeposited(opts *bind.FilterOpts, destinationChainId []*big.Int, depositId []uint32, depositor []common.Address) (*Router2FundsDepositedIterator, error) {

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

	logs, sub, err := _Router2.contract.FilterLogs(opts, "FundsDeposited", destinationChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &Router2FundsDepositedIterator{contract: _Router2.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0xafc4df6845a4ab948b492800d3d8a25d538a102a2bc07cd01f1cfa097fddcff6.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, uint256 indexed destinationChainId, int64 relayerFeePct, uint32 indexed depositId, uint32 quoteTimestamp, address originToken, address recipient, address indexed depositor, bytes message)
func (_Router2 *Router2Filterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *Router2FundsDeposited, destinationChainId []*big.Int, depositId []uint32, depositor []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Router2.contract.WatchLogs(opts, "FundsDeposited", destinationChainIdRule, depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2FundsDeposited)
				if err := _Router2.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseFundsDeposited(log types.Log) (*Router2FundsDeposited, error) {
	event := new(Router2FundsDeposited)
	if err := _Router2.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Router2 contract.
type Router2InitializedIterator struct {
	Event *Router2Initialized // Event containing the contract specifics and raw log

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
func (it *Router2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2Initialized)
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
		it.Event = new(Router2Initialized)
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
func (it *Router2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2Initialized represents a Initialized event raised by the Router2 contract.
type Router2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router2 *Router2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Router2InitializedIterator, error) {

	logs, sub, err := _Router2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Router2InitializedIterator{contract: _Router2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Router2 *Router2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Router2Initialized) (event.Subscription, error) {

	logs, sub, err := _Router2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2Initialized)
				if err := _Router2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseInitialized(log types.Log) (*Router2Initialized, error) {
	event := new(Router2Initialized)
	if err := _Router2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2PausedDepositsIterator is returned from FilterPausedDeposits and is used to iterate over the raw logs and unpacked data for PausedDeposits events raised by the Router2 contract.
type Router2PausedDepositsIterator struct {
	Event *Router2PausedDeposits // Event containing the contract specifics and raw log

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
func (it *Router2PausedDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2PausedDeposits)
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
		it.Event = new(Router2PausedDeposits)
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
func (it *Router2PausedDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2PausedDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2PausedDeposits represents a PausedDeposits event raised by the Router2 contract.
type Router2PausedDeposits struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPausedDeposits is a free log retrieval operation binding the contract event 0xe88463c2f254e2b070013a2dc7ee1e099f9bc00534cbdf03af551dc26ae49219.
//
// Solidity: event PausedDeposits(bool isPaused)
func (_Router2 *Router2Filterer) FilterPausedDeposits(opts *bind.FilterOpts) (*Router2PausedDepositsIterator, error) {

	logs, sub, err := _Router2.contract.FilterLogs(opts, "PausedDeposits")
	if err != nil {
		return nil, err
	}
	return &Router2PausedDepositsIterator{contract: _Router2.contract, event: "PausedDeposits", logs: logs, sub: sub}, nil
}

// WatchPausedDeposits is a free log subscription operation binding the contract event 0xe88463c2f254e2b070013a2dc7ee1e099f9bc00534cbdf03af551dc26ae49219.
//
// Solidity: event PausedDeposits(bool isPaused)
func (_Router2 *Router2Filterer) WatchPausedDeposits(opts *bind.WatchOpts, sink chan<- *Router2PausedDeposits) (event.Subscription, error) {

	logs, sub, err := _Router2.contract.WatchLogs(opts, "PausedDeposits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2PausedDeposits)
				if err := _Router2.contract.UnpackLog(event, "PausedDeposits", log); err != nil {
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
func (_Router2 *Router2Filterer) ParsePausedDeposits(log types.Log) (*Router2PausedDeposits, error) {
	event := new(Router2PausedDeposits)
	if err := _Router2.contract.UnpackLog(event, "PausedDeposits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2PausedFillsIterator is returned from FilterPausedFills and is used to iterate over the raw logs and unpacked data for PausedFills events raised by the Router2 contract.
type Router2PausedFillsIterator struct {
	Event *Router2PausedFills // Event containing the contract specifics and raw log

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
func (it *Router2PausedFillsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2PausedFills)
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
		it.Event = new(Router2PausedFills)
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
func (it *Router2PausedFillsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2PausedFillsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2PausedFills represents a PausedFills event raised by the Router2 contract.
type Router2PausedFills struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPausedFills is a free log retrieval operation binding the contract event 0x2d5b62420992e5a4afce0e77742636ca2608ef58289fd2e1baa5161ef6e7e41e.
//
// Solidity: event PausedFills(bool isPaused)
func (_Router2 *Router2Filterer) FilterPausedFills(opts *bind.FilterOpts) (*Router2PausedFillsIterator, error) {

	logs, sub, err := _Router2.contract.FilterLogs(opts, "PausedFills")
	if err != nil {
		return nil, err
	}
	return &Router2PausedFillsIterator{contract: _Router2.contract, event: "PausedFills", logs: logs, sub: sub}, nil
}

// WatchPausedFills is a free log subscription operation binding the contract event 0x2d5b62420992e5a4afce0e77742636ca2608ef58289fd2e1baa5161ef6e7e41e.
//
// Solidity: event PausedFills(bool isPaused)
func (_Router2 *Router2Filterer) WatchPausedFills(opts *bind.WatchOpts, sink chan<- *Router2PausedFills) (event.Subscription, error) {

	logs, sub, err := _Router2.contract.WatchLogs(opts, "PausedFills")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2PausedFills)
				if err := _Router2.contract.UnpackLog(event, "PausedFills", log); err != nil {
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
func (_Router2 *Router2Filterer) ParsePausedFills(log types.Log) (*Router2PausedFills, error) {
	event := new(Router2PausedFills)
	if err := _Router2.contract.UnpackLog(event, "PausedFills", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2RefundRequestedIterator is returned from FilterRefundRequested and is used to iterate over the raw logs and unpacked data for RefundRequested events raised by the Router2 contract.
type Router2RefundRequestedIterator struct {
	Event *Router2RefundRequested // Event containing the contract specifics and raw log

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
func (it *Router2RefundRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2RefundRequested)
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
		it.Event = new(Router2RefundRequested)
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
func (it *Router2RefundRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2RefundRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2RefundRequested represents a RefundRequested event raised by the Router2 contract.
type Router2RefundRequested struct {
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
func (_Router2 *Router2Filterer) FilterRefundRequested(opts *bind.FilterOpts, relayer []common.Address, originChainId []*big.Int, depositId []uint32) (*Router2RefundRequestedIterator, error) {

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

	logs, sub, err := _Router2.contract.FilterLogs(opts, "RefundRequested", relayerRule, originChainIdRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return &Router2RefundRequestedIterator{contract: _Router2.contract, event: "RefundRequested", logs: logs, sub: sub}, nil
}

// WatchRefundRequested is a free log subscription operation binding the contract event 0xace81ce0f8b8d27a1aed0c4df5f6b2743e46fc50fe3e0183dd7cd6a7b9db22fb.
//
// Solidity: event RefundRequested(address indexed relayer, address refundToken, uint256 amount, uint256 indexed originChainId, uint256 destinationChainId, int64 realizedLpFeePct, uint32 indexed depositId, uint256 fillBlock, uint256 previousIdenticalRequests)
func (_Router2 *Router2Filterer) WatchRefundRequested(opts *bind.WatchOpts, sink chan<- *Router2RefundRequested, relayer []common.Address, originChainId []*big.Int, depositId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _Router2.contract.WatchLogs(opts, "RefundRequested", relayerRule, originChainIdRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2RefundRequested)
				if err := _Router2.contract.UnpackLog(event, "RefundRequested", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseRefundRequested(log types.Log) (*Router2RefundRequested, error) {
	event := new(Router2RefundRequested)
	if err := _Router2.contract.UnpackLog(event, "RefundRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2RelayedRootBundleIterator is returned from FilterRelayedRootBundle and is used to iterate over the raw logs and unpacked data for RelayedRootBundle events raised by the Router2 contract.
type Router2RelayedRootBundleIterator struct {
	Event *Router2RelayedRootBundle // Event containing the contract specifics and raw log

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
func (it *Router2RelayedRootBundleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2RelayedRootBundle)
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
		it.Event = new(Router2RelayedRootBundle)
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
func (it *Router2RelayedRootBundleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2RelayedRootBundleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2RelayedRootBundle represents a RelayedRootBundle event raised by the Router2 contract.
type Router2RelayedRootBundle struct {
	RootBundleId      uint32
	RelayerRefundRoot [32]byte
	SlowRelayRoot     [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRelayedRootBundle is a free log retrieval operation binding the contract event 0xc86ba04c55bc5eb2f2876b91c438849a296dbec7b08751c3074d92e04f0a77af.
//
// Solidity: event RelayedRootBundle(uint32 indexed rootBundleId, bytes32 indexed relayerRefundRoot, bytes32 indexed slowRelayRoot)
func (_Router2 *Router2Filterer) FilterRelayedRootBundle(opts *bind.FilterOpts, rootBundleId []uint32, relayerRefundRoot [][32]byte, slowRelayRoot [][32]byte) (*Router2RelayedRootBundleIterator, error) {

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

	logs, sub, err := _Router2.contract.FilterLogs(opts, "RelayedRootBundle", rootBundleIdRule, relayerRefundRootRule, slowRelayRootRule)
	if err != nil {
		return nil, err
	}
	return &Router2RelayedRootBundleIterator{contract: _Router2.contract, event: "RelayedRootBundle", logs: logs, sub: sub}, nil
}

// WatchRelayedRootBundle is a free log subscription operation binding the contract event 0xc86ba04c55bc5eb2f2876b91c438849a296dbec7b08751c3074d92e04f0a77af.
//
// Solidity: event RelayedRootBundle(uint32 indexed rootBundleId, bytes32 indexed relayerRefundRoot, bytes32 indexed slowRelayRoot)
func (_Router2 *Router2Filterer) WatchRelayedRootBundle(opts *bind.WatchOpts, sink chan<- *Router2RelayedRootBundle, rootBundleId []uint32, relayerRefundRoot [][32]byte, slowRelayRoot [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Router2.contract.WatchLogs(opts, "RelayedRootBundle", rootBundleIdRule, relayerRefundRootRule, slowRelayRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2RelayedRootBundle)
				if err := _Router2.contract.UnpackLog(event, "RelayedRootBundle", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseRelayedRootBundle(log types.Log) (*Router2RelayedRootBundle, error) {
	event := new(Router2RelayedRootBundle)
	if err := _Router2.contract.UnpackLog(event, "RelayedRootBundle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2RequestedSpeedUpDepositIterator is returned from FilterRequestedSpeedUpDeposit and is used to iterate over the raw logs and unpacked data for RequestedSpeedUpDeposit events raised by the Router2 contract.
type Router2RequestedSpeedUpDepositIterator struct {
	Event *Router2RequestedSpeedUpDeposit // Event containing the contract specifics and raw log

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
func (it *Router2RequestedSpeedUpDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2RequestedSpeedUpDeposit)
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
		it.Event = new(Router2RequestedSpeedUpDeposit)
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
func (it *Router2RequestedSpeedUpDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2RequestedSpeedUpDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2RequestedSpeedUpDeposit represents a RequestedSpeedUpDeposit event raised by the Router2 contract.
type Router2RequestedSpeedUpDeposit struct {
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
func (_Router2 *Router2Filterer) FilterRequestedSpeedUpDeposit(opts *bind.FilterOpts, depositId []uint32, depositor []common.Address) (*Router2RequestedSpeedUpDepositIterator, error) {

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "RequestedSpeedUpDeposit", depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &Router2RequestedSpeedUpDepositIterator{contract: _Router2.contract, event: "RequestedSpeedUpDeposit", logs: logs, sub: sub}, nil
}

// WatchRequestedSpeedUpDeposit is a free log subscription operation binding the contract event 0xa6aa57bd282fc82378458de27be4eadfa791a0c7321c49562eeba8b2643dd566.
//
// Solidity: event RequestedSpeedUpDeposit(int64 newRelayerFeePct, uint32 indexed depositId, address indexed depositor, address updatedRecipient, bytes updatedMessage, bytes depositorSignature)
func (_Router2 *Router2Filterer) WatchRequestedSpeedUpDeposit(opts *bind.WatchOpts, sink chan<- *Router2RequestedSpeedUpDeposit, depositId []uint32, depositor []common.Address) (event.Subscription, error) {

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "RequestedSpeedUpDeposit", depositIdRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2RequestedSpeedUpDeposit)
				if err := _Router2.contract.UnpackLog(event, "RequestedSpeedUpDeposit", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseRequestedSpeedUpDeposit(log types.Log) (*Router2RequestedSpeedUpDeposit, error) {
	event := new(Router2RequestedSpeedUpDeposit)
	if err := _Router2.contract.UnpackLog(event, "RequestedSpeedUpDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2SetDepositQuoteTimeBufferIterator is returned from FilterSetDepositQuoteTimeBuffer and is used to iterate over the raw logs and unpacked data for SetDepositQuoteTimeBuffer events raised by the Router2 contract.
type Router2SetDepositQuoteTimeBufferIterator struct {
	Event *Router2SetDepositQuoteTimeBuffer // Event containing the contract specifics and raw log

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
func (it *Router2SetDepositQuoteTimeBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2SetDepositQuoteTimeBuffer)
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
		it.Event = new(Router2SetDepositQuoteTimeBuffer)
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
func (it *Router2SetDepositQuoteTimeBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2SetDepositQuoteTimeBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2SetDepositQuoteTimeBuffer represents a SetDepositQuoteTimeBuffer event raised by the Router2 contract.
type Router2SetDepositQuoteTimeBuffer struct {
	NewBuffer uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetDepositQuoteTimeBuffer is a free log retrieval operation binding the contract event 0x0e55dd180fa793d9036c804d0a116e6a7617a48e72cee1f83d92793a793fcc03.
//
// Solidity: event SetDepositQuoteTimeBuffer(uint32 newBuffer)
func (_Router2 *Router2Filterer) FilterSetDepositQuoteTimeBuffer(opts *bind.FilterOpts) (*Router2SetDepositQuoteTimeBufferIterator, error) {

	logs, sub, err := _Router2.contract.FilterLogs(opts, "SetDepositQuoteTimeBuffer")
	if err != nil {
		return nil, err
	}
	return &Router2SetDepositQuoteTimeBufferIterator{contract: _Router2.contract, event: "SetDepositQuoteTimeBuffer", logs: logs, sub: sub}, nil
}

// WatchSetDepositQuoteTimeBuffer is a free log subscription operation binding the contract event 0x0e55dd180fa793d9036c804d0a116e6a7617a48e72cee1f83d92793a793fcc03.
//
// Solidity: event SetDepositQuoteTimeBuffer(uint32 newBuffer)
func (_Router2 *Router2Filterer) WatchSetDepositQuoteTimeBuffer(opts *bind.WatchOpts, sink chan<- *Router2SetDepositQuoteTimeBuffer) (event.Subscription, error) {

	logs, sub, err := _Router2.contract.WatchLogs(opts, "SetDepositQuoteTimeBuffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2SetDepositQuoteTimeBuffer)
				if err := _Router2.contract.UnpackLog(event, "SetDepositQuoteTimeBuffer", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseSetDepositQuoteTimeBuffer(log types.Log) (*Router2SetDepositQuoteTimeBuffer, error) {
	event := new(Router2SetDepositQuoteTimeBuffer)
	if err := _Router2.contract.UnpackLog(event, "SetDepositQuoteTimeBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2SetHubPoolIterator is returned from FilterSetHubPool and is used to iterate over the raw logs and unpacked data for SetHubPool events raised by the Router2 contract.
type Router2SetHubPoolIterator struct {
	Event *Router2SetHubPool // Event containing the contract specifics and raw log

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
func (it *Router2SetHubPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2SetHubPool)
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
		it.Event = new(Router2SetHubPool)
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
func (it *Router2SetHubPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2SetHubPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2SetHubPool represents a SetHubPool event raised by the Router2 contract.
type Router2SetHubPool struct {
	NewHubPool common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetHubPool is a free log retrieval operation binding the contract event 0x1f17a88f67b0f49060a34bec1a4723a563620e6aa265eb640b5046dcee0759a0.
//
// Solidity: event SetHubPool(address indexed newHubPool)
func (_Router2 *Router2Filterer) FilterSetHubPool(opts *bind.FilterOpts, newHubPool []common.Address) (*Router2SetHubPoolIterator, error) {

	var newHubPoolRule []interface{}
	for _, newHubPoolItem := range newHubPool {
		newHubPoolRule = append(newHubPoolRule, newHubPoolItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "SetHubPool", newHubPoolRule)
	if err != nil {
		return nil, err
	}
	return &Router2SetHubPoolIterator{contract: _Router2.contract, event: "SetHubPool", logs: logs, sub: sub}, nil
}

// WatchSetHubPool is a free log subscription operation binding the contract event 0x1f17a88f67b0f49060a34bec1a4723a563620e6aa265eb640b5046dcee0759a0.
//
// Solidity: event SetHubPool(address indexed newHubPool)
func (_Router2 *Router2Filterer) WatchSetHubPool(opts *bind.WatchOpts, sink chan<- *Router2SetHubPool, newHubPool []common.Address) (event.Subscription, error) {

	var newHubPoolRule []interface{}
	for _, newHubPoolItem := range newHubPool {
		newHubPoolRule = append(newHubPoolRule, newHubPoolItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "SetHubPool", newHubPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2SetHubPool)
				if err := _Router2.contract.UnpackLog(event, "SetHubPool", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseSetHubPool(log types.Log) (*Router2SetHubPool, error) {
	event := new(Router2SetHubPool)
	if err := _Router2.contract.UnpackLog(event, "SetHubPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2SetXDomainAdminIterator is returned from FilterSetXDomainAdmin and is used to iterate over the raw logs and unpacked data for SetXDomainAdmin events raised by the Router2 contract.
type Router2SetXDomainAdminIterator struct {
	Event *Router2SetXDomainAdmin // Event containing the contract specifics and raw log

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
func (it *Router2SetXDomainAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2SetXDomainAdmin)
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
		it.Event = new(Router2SetXDomainAdmin)
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
func (it *Router2SetXDomainAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2SetXDomainAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2SetXDomainAdmin represents a SetXDomainAdmin event raised by the Router2 contract.
type Router2SetXDomainAdmin struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetXDomainAdmin is a free log retrieval operation binding the contract event 0xa9e8c42c9e7fca7f62755189a16b2f5314d43d8fb24e91ba54e6d65f9314e849.
//
// Solidity: event SetXDomainAdmin(address indexed newAdmin)
func (_Router2 *Router2Filterer) FilterSetXDomainAdmin(opts *bind.FilterOpts, newAdmin []common.Address) (*Router2SetXDomainAdminIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "SetXDomainAdmin", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &Router2SetXDomainAdminIterator{contract: _Router2.contract, event: "SetXDomainAdmin", logs: logs, sub: sub}, nil
}

// WatchSetXDomainAdmin is a free log subscription operation binding the contract event 0xa9e8c42c9e7fca7f62755189a16b2f5314d43d8fb24e91ba54e6d65f9314e849.
//
// Solidity: event SetXDomainAdmin(address indexed newAdmin)
func (_Router2 *Router2Filterer) WatchSetXDomainAdmin(opts *bind.WatchOpts, sink chan<- *Router2SetXDomainAdmin, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "SetXDomainAdmin", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2SetXDomainAdmin)
				if err := _Router2.contract.UnpackLog(event, "SetXDomainAdmin", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseSetXDomainAdmin(log types.Log) (*Router2SetXDomainAdmin, error) {
	event := new(Router2SetXDomainAdmin)
	if err := _Router2.contract.UnpackLog(event, "SetXDomainAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2SetZkBridgeIterator is returned from FilterSetZkBridge and is used to iterate over the raw logs and unpacked data for SetZkBridge events raised by the Router2 contract.
type Router2SetZkBridgeIterator struct {
	Event *Router2SetZkBridge // Event containing the contract specifics and raw log

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
func (it *Router2SetZkBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2SetZkBridge)
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
		it.Event = new(Router2SetZkBridge)
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
func (it *Router2SetZkBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2SetZkBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2SetZkBridge represents a SetZkBridge event raised by the Router2 contract.
type Router2SetZkBridge struct {
	Erc20Bridge    common.Address
	OldErc20Bridge common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetZkBridge is a free log retrieval operation binding the contract event 0xab9987862d7dab8c2bf2f09eb3ad7b6c1961c86fe3e9e7d3ef50dc98995d72d4.
//
// Solidity: event SetZkBridge(address indexed erc20Bridge, address indexed oldErc20Bridge)
func (_Router2 *Router2Filterer) FilterSetZkBridge(opts *bind.FilterOpts, erc20Bridge []common.Address, oldErc20Bridge []common.Address) (*Router2SetZkBridgeIterator, error) {

	var erc20BridgeRule []interface{}
	for _, erc20BridgeItem := range erc20Bridge {
		erc20BridgeRule = append(erc20BridgeRule, erc20BridgeItem)
	}
	var oldErc20BridgeRule []interface{}
	for _, oldErc20BridgeItem := range oldErc20Bridge {
		oldErc20BridgeRule = append(oldErc20BridgeRule, oldErc20BridgeItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "SetZkBridge", erc20BridgeRule, oldErc20BridgeRule)
	if err != nil {
		return nil, err
	}
	return &Router2SetZkBridgeIterator{contract: _Router2.contract, event: "SetZkBridge", logs: logs, sub: sub}, nil
}

// WatchSetZkBridge is a free log subscription operation binding the contract event 0xab9987862d7dab8c2bf2f09eb3ad7b6c1961c86fe3e9e7d3ef50dc98995d72d4.
//
// Solidity: event SetZkBridge(address indexed erc20Bridge, address indexed oldErc20Bridge)
func (_Router2 *Router2Filterer) WatchSetZkBridge(opts *bind.WatchOpts, sink chan<- *Router2SetZkBridge, erc20Bridge []common.Address, oldErc20Bridge []common.Address) (event.Subscription, error) {

	var erc20BridgeRule []interface{}
	for _, erc20BridgeItem := range erc20Bridge {
		erc20BridgeRule = append(erc20BridgeRule, erc20BridgeItem)
	}
	var oldErc20BridgeRule []interface{}
	for _, oldErc20BridgeItem := range oldErc20Bridge {
		oldErc20BridgeRule = append(oldErc20BridgeRule, oldErc20BridgeItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "SetZkBridge", erc20BridgeRule, oldErc20BridgeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2SetZkBridge)
				if err := _Router2.contract.UnpackLog(event, "SetZkBridge", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseSetZkBridge(log types.Log) (*Router2SetZkBridge, error) {
	event := new(Router2SetZkBridge)
	if err := _Router2.contract.UnpackLog(event, "SetZkBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2TokensBridgedIterator is returned from FilterTokensBridged and is used to iterate over the raw logs and unpacked data for TokensBridged events raised by the Router2 contract.
type Router2TokensBridgedIterator struct {
	Event *Router2TokensBridged // Event containing the contract specifics and raw log

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
func (it *Router2TokensBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2TokensBridged)
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
		it.Event = new(Router2TokensBridged)
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
func (it *Router2TokensBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2TokensBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2TokensBridged represents a TokensBridged event raised by the Router2 contract.
type Router2TokensBridged struct {
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
func (_Router2 *Router2Filterer) FilterTokensBridged(opts *bind.FilterOpts, chainId []*big.Int, leafId []uint32, l2TokenAddress []common.Address) (*Router2TokensBridgedIterator, error) {

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

	logs, sub, err := _Router2.contract.FilterLogs(opts, "TokensBridged", chainIdRule, leafIdRule, l2TokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &Router2TokensBridgedIterator{contract: _Router2.contract, event: "TokensBridged", logs: logs, sub: sub}, nil
}

// WatchTokensBridged is a free log subscription operation binding the contract event 0x828fc203220356df8f072a91681caee7d5c75095e2a95e80ed5a14b384697f71.
//
// Solidity: event TokensBridged(uint256 amountToReturn, uint256 indexed chainId, uint32 indexed leafId, address indexed l2TokenAddress, address caller)
func (_Router2 *Router2Filterer) WatchTokensBridged(opts *bind.WatchOpts, sink chan<- *Router2TokensBridged, chainId []*big.Int, leafId []uint32, l2TokenAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Router2.contract.WatchLogs(opts, "TokensBridged", chainIdRule, leafIdRule, l2TokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2TokensBridged)
				if err := _Router2.contract.UnpackLog(event, "TokensBridged", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseTokensBridged(log types.Log) (*Router2TokensBridged, error) {
	event := new(Router2TokensBridged)
	if err := _Router2.contract.UnpackLog(event, "TokensBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Router2 contract.
type Router2UpgradedIterator struct {
	Event *Router2Upgraded // Event containing the contract specifics and raw log

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
func (it *Router2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2Upgraded)
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
		it.Event = new(Router2Upgraded)
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
func (it *Router2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2Upgraded represents a Upgraded event raised by the Router2 contract.
type Router2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Router2 *Router2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Router2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Router2UpgradedIterator{contract: _Router2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Router2 *Router2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Router2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2Upgraded)
				if err := _Router2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseUpgraded(log types.Log) (*Router2Upgraded, error) {
	event := new(Router2Upgraded)
	if err := _Router2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Router2ZkSyncTokensBridgedIterator is returned from FilterZkSyncTokensBridged and is used to iterate over the raw logs and unpacked data for ZkSyncTokensBridged events raised by the Router2 contract.
type Router2ZkSyncTokensBridgedIterator struct {
	Event *Router2ZkSyncTokensBridged // Event containing the contract specifics and raw log

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
func (it *Router2ZkSyncTokensBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Router2ZkSyncTokensBridged)
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
		it.Event = new(Router2ZkSyncTokensBridged)
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
func (it *Router2ZkSyncTokensBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Router2ZkSyncTokensBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Router2ZkSyncTokensBridged represents a ZkSyncTokensBridged event raised by the Router2 contract.
type Router2ZkSyncTokensBridged struct {
	L2Token               common.Address
	Target                common.Address
	NumberOfTokensBridged *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZkSyncTokensBridged is a free log retrieval operation binding the contract event 0x9d523e6fab68ea063b948ed2d9e4e9091522a3bdbb63f1ad70c744b0edb9cdbf.
//
// Solidity: event ZkSyncTokensBridged(address indexed l2Token, address target, uint256 numberOfTokensBridged)
func (_Router2 *Router2Filterer) FilterZkSyncTokensBridged(opts *bind.FilterOpts, l2Token []common.Address) (*Router2ZkSyncTokensBridgedIterator, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _Router2.contract.FilterLogs(opts, "ZkSyncTokensBridged", l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &Router2ZkSyncTokensBridgedIterator{contract: _Router2.contract, event: "ZkSyncTokensBridged", logs: logs, sub: sub}, nil
}

// WatchZkSyncTokensBridged is a free log subscription operation binding the contract event 0x9d523e6fab68ea063b948ed2d9e4e9091522a3bdbb63f1ad70c744b0edb9cdbf.
//
// Solidity: event ZkSyncTokensBridged(address indexed l2Token, address target, uint256 numberOfTokensBridged)
func (_Router2 *Router2Filterer) WatchZkSyncTokensBridged(opts *bind.WatchOpts, sink chan<- *Router2ZkSyncTokensBridged, l2Token []common.Address) (event.Subscription, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _Router2.contract.WatchLogs(opts, "ZkSyncTokensBridged", l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Router2ZkSyncTokensBridged)
				if err := _Router2.contract.UnpackLog(event, "ZkSyncTokensBridged", log); err != nil {
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
func (_Router2 *Router2Filterer) ParseZkSyncTokensBridged(log types.Log) (*Router2ZkSyncTokensBridged, error) {
	event := new(Router2ZkSyncTokensBridged)
	if err := _Router2.contract.UnpackLog(event, "ZkSyncTokensBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
