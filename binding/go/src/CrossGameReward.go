// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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

// ICrossGameRewardPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type ICrossGameRewardPoolInfo struct {
	PoolId       *big.Int
	Pool         common.Address
	Name         string
	DepositToken common.Address
	CreatedAt    *big.Int
}

// CrossGameRewardMetaData contains all meta data concerning the CrossGameReward contract.
var CrossGameRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"createGamePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gamePoolImplementation\",\"outputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActivePoolIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPoolIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolAddress\",\"outputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"}],\"name\":\"getPoolCountByDepositToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"}],\"name\":\"getPoolIdsByDepositToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"poolType\",\"type\":\"uint8\"}],\"name\":\"getPoolIdsByType\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structICrossGameReward.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolType\",\"outputs\":[{\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"grantSponsorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"_poolImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolByDepositTokenAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolImplementation\",\"outputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolTypes\",\"outputs\":[{\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reclaimFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"revokeSponsorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setGamePoolImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxActiveRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setPoolImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"poolType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradePoolsByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcross\",\"outputs\":[{\"internalType\":\"contractIWCROSS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"GamePoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"GamePoolImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"PoolImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"poolType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"PoolsBatchUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReclaimedFromPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"SponsorRoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"SponsorRoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRGamePoolImplNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"enumICrossGameReward.PoolType\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"CGRInvalidPoolType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRPoolNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ec87621c": "MANAGER_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"a1635945": "addRewardToken(uint256,address)",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"e01085e5": "createGamePool(string,address,address,uint256)",
		"6e13ba6f": "createPool(string,address,uint256)",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"2a21f7ae": "gamePoolImplementation()",
		"fe96e4ff": "getActivePoolIds()",
		"f19c3d5b": "getAllPoolIds()",
		"00a5ae21": "getPoolAddress(uint256)",
		"d4148bcd": "getPoolCountByDepositToken(address)",
		"caa9a08d": "getPoolId(address)",
		"eeea4a79": "getPoolIdsByDepositToken(address)",
		"0bee20d8": "getPoolIdsByType(uint8)",
		"2f380b35": "getPoolInfo(uint256)",
		"cdcf8783": "getPoolType(uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"e7590268": "getTotalPoolCount()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"ed681780": "grantSponsorRole(uint256,address)",
		"91d14854": "hasRole(bytes32,address)",
		"ce24af53": "initialize(address,address,uint48)",
		"91cf6d3e": "initializedAt()",
		"18e56131": "nextPoolId()",
		"8da5cb5b": "owner()",
		"cf6eefb7": "pendingDefaultAdmin()",
		"a1eda53c": "pendingDefaultAdminDelay()",
		"155fff62": "poolAt(uint256)",
		"b5be3221": "poolByDepositTokenAt(address,uint256)",
		"d4175be2": "poolIds(address)",
		"cefa7799": "poolImplementation()",
		"1b95a010": "poolTypes(uint256)",
		"ac4afa38": "pools(uint256)",
		"52d1902d": "proxiableUUID()",
		"c24140b2": "reclaimFromPool(uint256,address,address)",
		"35cc9cb4": "removeRewardToken(uint256,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"b4368ae8": "revokeSponsorRole(uint256,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"f887ea40": "router()",
		"f62d31fd": "setGamePoolImplementation(address)",
		"decbb6e7": "setMaxActiveRounds(uint256,uint256)",
		"d6f74898": "setPoolImplementation(address)",
		"b34c972e": "setPoolStatus(uint256,uint8)",
		"c0d78655": "setRouter(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"61616c46": "updateMinDepositAmount(uint256,uint256)",
		"320533a0": "upgradePoolsByType(uint8,address,bytes)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"a2db4582": "wcross()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615c906100f95f395f81816132040152818161322d01526133df0152615c905ff3fe608060405260043610610370575f3560e01c8063ac4afa38116101c8578063d4148bcd116100fd578063e75902681161009d578063f19c3d5b1161006d578063f19c3d5b14610b72578063f62d31fd14610b86578063f887ea4014610ba5578063fe96e4ff14610bc4575f5ffd5b8063e759026814610aed578063ec87621c14610b01578063ed68178014610b34578063eeea4a7914610b53575f5ffd5b8063d602b9fd116100d8578063d602b9fd14610a7c578063d6f7489814610a90578063decbb6e714610aaf578063e01085e514610ace575f5ffd5b8063d4148bcd14610a13578063d4175be214610a32578063d547741f14610a5d575f5ffd5b8063caa9a08d11610168578063ce24af5311610143578063ce24af5314610957578063cefa779914610976578063cefc142914610995578063cf6eefb7146109a9575f5ffd5b8063caa9a08d14610905578063cc8463c814610924578063cdcf878314610938575f5ffd5b8063b4368ae8116101a3578063b4368ae814610889578063b5be3221146108a8578063c0d78655146108c7578063c24140b2146108e6575f5ffd5b8063ac4afa38146107e5578063ad3cb1cc14610815578063b34c972e1461086a575f5ffd5b806336568abe116102a957806384ef8ffc11610249578063a163594511610219578063a163594514610761578063a1eda53c14610780578063a217fddf146107b3578063a2db4582146107c6575f5ffd5b806384ef8ffc1461069a5780638da5cb5b146106d657806391cf6d3e146106ea57806391d14854146106fe575f5ffd5b806361616c461161028457806361616c4614610601578063634e93da14610620578063649a5ec71461063f5780636e13ba6f1461065e575f5ffd5b806336568abe146105bb5780634f1ef286146105da57806352d1902d146105ed575f5ffd5b80631b95a010116103145780632f2ff15d116102ef5780632f2ff15d146105325780632f380b3514610551578063320533a01461057d57806335cc9cb41461059c575f5ffd5b80631b95a0101461048b578063248a9ca3146104c65780632a21f7ae14610513575f5ffd5b80630aa6220b1161034f5780630aa6220b146104075780630bee20d81461041d578063155fff621461044957806318e5613114610476575f5ffd5b8062a5ae211461037457806301ffc9a7146103b0578063022d63fb146103df575b5f5ffd5b34801561037f575f5ffd5b5061039361038e366004614191565b610bd8565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156103bb575f5ffd5b506103cf6103ca3660046141a8565b610c2d565b60405190151581526020016103a7565b3480156103ea575f5ffd5b50620697805b60405165ffffffffffff90911681526020016103a7565b348015610412575f5ffd5b5061041b610c88565b005b348015610428575f5ffd5b5061043c6104373660046141fa565b610c9d565b6040516103a79190614213565b348015610454575f5ffd5b50610468610463366004614191565b610e13565b6040519081526020016103a7565b348015610481575f5ffd5b5061046860045481565b348015610496575f5ffd5b506104b96104a5366004614191565b600b6020525f908152604090205460ff1681565b6040516103a79190614296565b3480156104d1575f5ffd5b506104686104e0366004614191565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561051e575f5ffd5b50600a54610393906001600160a01b031681565b34801561053d575f5ffd5b5061041b61054c3660046142b8565b610e1f565b34801561055c575f5ffd5b5061057061056b366004614191565b610e64565b6040516103a79190614314565b348015610588575f5ffd5b5061041b6105973660046143ba565b610fc8565b3480156105a7575f5ffd5b5061041b6105b63660046142b8565b611223565b3480156105c6575f5ffd5b5061041b6105d53660046142b8565b61130c565b61041b6105e8366004614446565b61145e565b3480156105f8575f5ffd5b50610468611479565b34801561060c575f5ffd5b5061041b61061b36600461450d565b6114a7565b34801561062b575f5ffd5b5061041b61063a36600461452d565b611561565b34801561064a575f5ffd5b5061041b61065936600461455d565b611574565b348015610669575f5ffd5b5061067d610678366004614576565b611587565b604080519283526001600160a01b039091166020830152016103a7565b3480156106a5575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b0316610393565b3480156106e1575f5ffd5b50610393611906565b3480156106f5575f5ffd5b506104685f5481565b348015610709575f5ffd5b506103cf6107183660046142b8565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561076c575f5ffd5b5061041b61077b3660046142b8565b61193d565b34801561078b575f5ffd5b506107946119f8565b6040805165ffffffffffff9384168152929091166020830152016103a7565b3480156107be575f5ffd5b506104685f81565b3480156107d1575f5ffd5b50600154610393906001600160a01b031681565b3480156107f0575f5ffd5b506108046107ff366004614191565b611ab5565b6040516103a79594939291906145cf565b348015610820575f5ffd5b5061085d6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516103a79190614613565b348015610875575f5ffd5b5061041b610884366004614631565b611b7d565b348015610894575f5ffd5b5061041b6108a33660046142b8565b611c36565b3480156108b3575f5ffd5b506104686108c2366004614654565b611e5a565b3480156108d2575f5ffd5b5061041b6108e136600461452d565b611e82565b3480156108f1575f5ffd5b5061041b61090036600461467e565b611f23565b348015610910575f5ffd5b5061046861091f36600461452d565b6120f8565b34801561092f575f5ffd5b506103f0612130565b348015610943575f5ffd5b506104b9610952366004614191565b612210565b348015610962575f5ffd5b5061041b6109713660046146bd565b61225c565b348015610981575f5ffd5b50600354610393906001600160a01b031681565b3480156109a0575f5ffd5b5061041b6124fc565b3480156109b4575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016103a7565b348015610a1e575f5ffd5b50610468610a2d36600461452d565b61256a565b348015610a3d575f5ffd5b50610468610a4c36600461452d565b60066020525f908152604090205481565b348015610a68575f5ffd5b5061041b610a773660046142b8565b61258a565b348015610a87575f5ffd5b5061041b6125cb565b348015610a9b575f5ffd5b5061041b610aaa36600461452d565b6125dd565b348015610aba575f5ffd5b5061041b610ac936600461450d565b61267e565b348015610ad9575f5ffd5b5061067d610ae8366004614701565b6127b2565b348015610af8575f5ffd5b50610468612bc1565b348015610b0c575f5ffd5b506104687f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610b3f575f5ffd5b5061041b610b4e3660046142b8565b612bcc565b348015610b5e575f5ffd5b5061043c610b6d36600461452d565b612df0565b348015610b7d575f5ffd5b5061043c612e13565b348015610b91575f5ffd5b5061041b610ba036600461452d565b612e1f565b348015610bb0575f5ffd5b50600254610393906001600160a01b031681565b348015610bcf575f5ffd5b5061043c612ec0565b5f818152600560205260408120600101546001600160a01b0316610c0f5760405163c7dfdd2160e01b815260040160405180910390fd5b505f908152600560205260409020600101546001600160a01b031690565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610c825750610c82826130a4565b92915050565b5f610c928161313a565b610c9a613144565b50565b60605f610caa6008613150565b90505f8167ffffffffffffffff811115610cc657610cc6614419565b604051908082528060200260200182016040528015610cef578160200160208202803683370190505b5090505f805b83811015610d79575f610d09600883613159565b9050866001811115610d1d57610d1d614255565b5f828152600b602052604090205460ff166001811115610d3f57610d3f614255565b03610d705780848481518110610d5757610d5761476c565b602090810291909101015282610d6c816147c6565b9350505b50600101610cf5565b505f8167ffffffffffffffff811115610d9457610d94614419565b604051908082528060200260200182016040528015610dbd578160200160208202803683370190505b5090505f5b82811015610e0957838181518110610ddc57610ddc61476c565b6020026020010151828281518110610df657610df661476c565b6020908102919091010152600101610dc2565b5095945050505050565b5f610c82600883613159565b81610e56576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e608282613164565b5050565b610ea46040518060a001604052805f81526020015f6001600160a01b03168152602001606081526020015f6001600160a01b031681526020015f81525090565b5f828152600560205260409020600101546001600160a01b0316610edb5760405163c7dfdd2160e01b815260040160405180910390fd5b5f82815260056020908152604091829020825160a0810184528154815260018201546001600160a01b0316928101929092526002810180549293919291840191610f24906147fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610f50906147fd565b8015610f9b5780601f10610f7257610100808354040283529160200191610f9b565b820191905f5260205f20905b815481529060010190602001808311610f7e57829003601f168201915b505050918352505060038201546001600160a01b0316602082015260049091015460409091015292915050565b5f610fd28161313a565b6001600160a01b038416611012576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61101d6008613150565b90505f805b82811015611112575f611036600883613159565b905088600181111561104a5761104a614255565b5f828152600b602052604090205460ff16600181111561106c5761106c614255565b03611109575f81815260056020526040908190206001015490517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690634f1ef286906110cd908b908b908b90600401614877565b5f604051808303815f87803b1580156110e4575f5ffd5b505af11580156110f6573d5f5f3e3d5ffd5b505050508280611105906147c6565b9350505b50600101611022565b505f87600181111561112657611126614255565b03611183576003805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0388169081179091556040517fdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957905f90a26111d7565b600a805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0388169081179091556040517feb3eb9823880a12aa7872f3d387d85ca1508fcf13f0311e53851853bb044e2a6905f90a25b856001600160a01b03167f1d4313ba8617c7bedafe9b8528dafcacb84a619f94dab23c92c00460dbbe3f668883604051611212929190614899565b60405180910390a250505050505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861124d8161313a565b5f838152600560205260409020600101546001600160a01b03166112845760405163c7dfdd2160e01b815260040160405180910390fd5b5f83815260056020526040908190206001015490517f3d509c970000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690633d509c97906024015b5f604051808303815f87803b1580156112f1575f5ffd5b505af1158015611303573d5f5f3e3d5ffd5b50505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561136757507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b1561144f577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806113cd575065ffffffffffff8116155b806113e057504265ffffffffffff821610155b15611426576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b61145983836131ad565b505050565b6114666131f9565b61146f826132c9565b610e6082826132d3565b5f6114826133d4565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086114d18161313a565b5f838152600560205260409020600101546001600160a01b03166115085760405163c7dfdd2160e01b815260040160405180910390fd5b5f83815260056020526040908190206001015490517f84780205000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b03909116906384780205906024016112da565b5f61156b8161313a565b610e6082613436565b5f61157e8161313a565b610e60826134a8565b5f5f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086115b38161313a565b856115ea576040517f944c928200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03851661162a576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8411611663576040517f944c928200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60048054905f611672836147c6565b909155506040516001600160a01b0387166024820152604481018690529093505f9060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcd6dc6870000000000000000000000000000000000000000000000000000000017905260035490519192505f916001600160a01b0390911690839061171090614177565b61171b9291906148b4565b604051809103905ff080158015611734573d5f5f3e3d5ffd5b5090508093506040518060a00160405280868152602001856001600160a01b031681526020018a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160a01b038a8116602080850191909152426040948501528983526005815291839020845181559184015160018301805473ffffffffffffffffffffffffffffffffffffffff1916919092161790559082015160028201906117f69082614920565b50606082015160038201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392831617905560809092015160049091015584165f908152600660205260409020859055611850600886613517565b506001600160a01b0387165f9081526007602052604090206118729086613517565b505f858152600b6020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001830217905550866001600160a01b0316846001600160a01b0316867f9b64517ebf8d0fab4c3ec4b04d596444ace7293ad69076fe2458bc12be9126e38c8c6040516118f2929190614a19565b60405180910390a450505094509492505050565b5f6119387feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b905090565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086119678161313a565b5f838152600560205260409020600101546001600160a01b031661199e5760405163c7dfdd2160e01b815260040160405180910390fd5b5f83815260056020526040908190206001015490517f1c03e6cc0000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690631c03e6cc906024016112da565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008115801590611a7a57504265ffffffffffff831610155b611a85575f5f611aac565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b60056020525f908152604090208054600182015460028301805492936001600160a01b0390921692611ae6906147fd565b80601f0160208091040260200160405190810160405280929190818152602001828054611b12906147fd565b8015611b5d5780601f10611b3457610100808354040283529160200191611b5d565b820191905f5260205f20905b815481529060010190602001808311611b4057829003601f168201915b50505050600383015460049093015491926001600160a01b031691905085565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611ba78161313a565b5f838152600560205260409020600101546001600160a01b0316611bde5760405163c7dfdd2160e01b815260040160405180910390fd5b5f83815260056020526040908190206001015490517f6d7c49a20000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636d7c49a2906112da908590600401614a2c565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611c608161313a565b5f838152600560205260409020600101546001600160a01b0316611c975760405163c7dfdd2160e01b815260040160405180910390fd5b60015f848152600b602052604090205460ff166001811115611cbb57611cbb614255565b5f858152600b6020526040902054859260019260ff9092169114611d0e576040517f4d3537d600000000000000000000000000000000000000000000000000000000815260040161141d93929190614a46565b5050505f838152600560209081526040918290206001015482517fc2d7944400000000000000000000000000000000000000000000000000000000815292516001600160a01b0390911692839263d547741f92849263c2d794449260048083019391928290030181865afa158015611d88573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611dac9190614a67565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815260048101919091526001600160a01b03861660248201526044015f604051808303815f87803b158015611e09575f5ffd5b505af1158015611e1b573d5f5f3e3d5ffd5b50506040516001600160a01b03861692508691507f5d4cec63c165d8b6c2594e1506a0f45239d7ce890579740653db448bf4f8fccf905f90a350505050565b6001600160a01b0382165f908152600760205260408120611e7b9083613159565b9392505050565b5f611e8c8161313a565b6001600160a01b038216611ecc576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384169081179091556040517fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15905f90a25050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08611f4d8161313a565b5f848152600560205260409020600101546001600160a01b0316611f845760405163c7dfdd2160e01b815260040160405180910390fd5b5f848152600560205260408082206001015490517f35c21d5d0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152909116906335c21d5d90602401602060405180830381865afa158015611ff3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120179190614a67565b5f86815260056020526040908190206001015490517f4d1cd0140000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301528681166024830152929350911690634d1cd014906044015f604051808303815f87803b15801561208d575f5ffd5b505af115801561209f573d5f5f3e3d5ffd5b50505050826001600160a01b0316846001600160a01b0316867fae55fdf2c7467a88ea571a46bc6ecd9b95b7997fa6fed1d1c7f1842b5d603389846040516120e991815260200190565b60405180910390a45050505050565b6001600160a01b0381165f90815260066020526040812054808203610c825760405163c7dfdd2160e01b815260040160405180910390fd5b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680158015906121b257504265ffffffffffff8216105b6121e35781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16612209565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b5f818152600560205260408120600101546001600160a01b03166122475760405163c7dfdd2160e01b815260040160405180910390fd5b505f908152600b602052604090205460ff1690565b5f612265613522565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f811580156122915750825b90505f8267ffffffffffffffff1660011480156122ad5750303b155b9050811580156122bb575080155b156122f2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156123535784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816612393576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0387166123d3576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6123dd868861354a565b6123e561355c565b435f556003805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038a1617905560405161241c90614184565b604051809103905ff080158015612435573d5f5f3e3d5ffd5b506001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03929092169190911781556004556124907f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0888613564565b5083156124f25784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114612562576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161141d565b610c9a61363a565b6001600160a01b0381165f908152600760205260408120610c8290613150565b816125c1576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e60828261374f565b5f6125d58161313a565b610c9a613792565b5f6125e78161313a565b6001600160a01b038216612627576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384169081179091556040517fdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957905f90a25050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086126a88161313a565b5f838152600560205260409020600101546001600160a01b03166126df5760405163c7dfdd2160e01b815260040160405180910390fd5b60015f848152600b602052604090205460ff16600181111561270357612703614255565b5f858152600b6020526040902054859260019260ff9092169114612756576040517f4d3537d600000000000000000000000000000000000000000000000000000000815260040161141d93929190614a46565b5050505f83815260056020526040908190206001015490517f5a6a0153000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690635a6a0153906024016112da565b5f5f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086127de8161313a565b600a546001600160a01b0316612820576040517f6794251100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b86612857576040517f944c928200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038616612897576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0385166128d7576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8411612910576040517f944c928200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60048054905f61291f836147c6565b909155506040516001600160a01b03808916602483015287166044820152606481018690529093505f9060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1794bb3c00000000000000000000000000000000000000000000000000000000179052600a5490519192505f916001600160a01b039091169083906129c590614177565b6129d09291906148b4565b604051809103905ff0801580156129e9573d5f5f3e3d5ffd5b5090508093506040518060a00160405280868152602001856001600160a01b031681526020018b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160a01b038b8116602080850191909152426040948501528983526005815291839020845181559184015160018301805473ffffffffffffffffffffffffffffffffffffffff191691909216179055908201516002820190612aab9082614920565b50606082015160038201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392831617905560809092015160049091015584165f908152600660205260409020859055612b05600886613517565b506001600160a01b0388165f908152600760205260409020612b279086613517565b505f858152600b602052604090208054600191907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682800217905550866001600160a01b0316846001600160a01b0316867fcb6c4a9a1de3ee56c46859db21e5f245245aee3068e7b10d47bc646583cb816d8b8e8e604051612bac93929190614877565b60405180910390a45050509550959350505050565b5f6119386008613150565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08612bf68161313a565b5f838152600560205260409020600101546001600160a01b0316612c2d5760405163c7dfdd2160e01b815260040160405180910390fd5b60015f848152600b602052604090205460ff166001811115612c5157612c51614255565b5f858152600b6020526040902054859260019260ff9092169114612ca4576040517f4d3537d600000000000000000000000000000000000000000000000000000000815260040161141d93929190614a46565b5050505f838152600560209081526040918290206001015482517fc2d7944400000000000000000000000000000000000000000000000000000000815292516001600160a01b03909116928392632f2ff15d92849263c2d794449260048083019391928290030181865afa158015612d1e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d429190614a67565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815260048101919091526001600160a01b03861660248201526044015f604051808303815f87803b158015612d9f575f5ffd5b505af1158015612db1573d5f5f3e3d5ffd5b50506040516001600160a01b03861692508691507fb341de1bcd4424e64431d748873b008dd03d26d1d6e059f321f3555650617a9b905f90a350505050565b6001600160a01b0381165f908152600760205260409020606090610c829061379c565b6060611938600861379c565b5f612e298161313a565b6001600160a01b038216612e69576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384169081179091556040517feb3eb9823880a12aa7872f3d387d85ca1508fcf13f0311e53851853bb044e2a6905f90a25050565b60605f612ecd6008613150565b90505f8167ffffffffffffffff811115612ee957612ee9614419565b604051908082528060200260200182016040528015612f12578160200160208202803683370190505b5090505f805b8381101561300b575f612f2c600883613159565b90505f5f828152600560209081526040918290206001015482517ff022869200000000000000000000000000000000000000000000000000000000815292516001600160a01b039091169263f02286929260048083019391928290030181865afa158015612f9c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fc09190614a7e565b6002811115612fd157612fd1614255565b036130025780848481518110612fe957612fe961476c565b602090810291909101015282612ffe816147c6565b9350505b50600101612f18565b505f8167ffffffffffffffff81111561302657613026614419565b60405190808252806020026020018201604052801561304f578160200160208202803683370190505b5090505f5b8281101561309b5783818151811061306e5761306e61476c565b60200260200101518282815181106130885761308861476c565b6020908102919091010152600101613054565b50949350505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610c8257507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610c82565b610c9a81336137a8565b61314e5f5f613834565b565b5f610c82825490565b5f611e7b83836139bf565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461319d8161313a565b6131a78383613564565b50505050565b6001600160a01b03811633146131ef576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61145982826139e5565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061329257507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166132867f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561314e576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610e608161313a565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561332d575060408051601f3d908101601f1916820190925261332a91810190614a67565b60015b61336e576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161141d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146133ca576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161141d565b6114598383613a70565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461314e576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61343f612130565b61344842613ac5565b6134529190614a99565b905061345e8282613b14565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6134b282613bc1565b6134bb42613ac5565b6134c59190614a99565b90506134d18282613834565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f611e7b8383613c08565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610c82565b613552613c54565b610e608282613c92565b61314e613c54565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083613628575f6135bd7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b0316146135fd576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0385161790555b6136328484613d4e565b949350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff1680158061369d57504265ffffffffffff821610155b156136de576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161141d565b6137185f6137137feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6139e5565b506137235f83613564565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546137888161313a565b6131a783836139e5565b61314e5f5f613b14565b60605f611e7b83613e38565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610e60576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161141d565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015613946574265ffffffffffff8216101561391d576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255613946565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f825f0182815481106139d4576139d461476c565b905f5260205f200154905092915050565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083158015613a4157507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b15613a665760018101805473ffffffffffffffffffffffffffffffffffffffff191690555b6136328484613e91565b613a7982613f53565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613abd576114598282613fef565b610e60614061565b5f65ffffffffffff821115613b10576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161141d565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b0388161717845591041680156131a7576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f613bcb612130565b90508065ffffffffffff168365ffffffffffff1611613bf357613bee8382614ab7565b611e7b565b611e7b65ffffffffffff841662069780614099565b5f818152600183016020526040812054613c4d57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610c82565b505f610c82565b613c5c6140a8565b61314e576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613c9a613c54565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216613cfd576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f600482015260240161141d565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556131a75f83613564565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16613e2f575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055613de53390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610c82565b5f915050610c82565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613e8557602002820191905f5260205f20905b815481526020019060010190808311613e71575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615613e2f575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610c82565b806001600160a01b03163b5f03613fa1576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161141d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161400b9190614ad5565b5f60405180830381855af49150503d805f8114614043576040519150601f19603f3d011682016040523d82523d5f602084013e614048565b606091505b50915091506140588583836140c6565b95945050505050565b341561314e576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828218828410028218611e7b565b5f6140b1613522565b5468010000000000000000900460ff16919050565b6060826140d657613bee82614136565b81511580156140ed57506001600160a01b0384163b155b1561412f576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161141d565b5092915050565b80511561414557805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dc80614aec83390190565b610d9380614ec883390190565b5f602082840312156141a1575f5ffd5b5035919050565b5f602082840312156141b8575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611e7b575f5ffd5b8035600281106141f5575f5ffd5b919050565b5f6020828403121561420a575f5ffd5b611e7b826141e7565b602080825282518282018190525f918401906040840190835b8181101561424a57835183526020938401939092019160010161422c565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6002811061429257614292614255565b9052565b60208101610c828284614282565b6001600160a01b0381168114610c9a575f5ffd5b5f5f604083850312156142c9575f5ffd5b8235915060208301356142db816142a4565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60208152815160208201526001600160a01b0360208301511660408201525f604083015160a0606084015261434c60c08401826142e6565b90506001600160a01b036060850151166080840152608084015160a08401528091505092915050565b5f5f83601f840112614385575f5ffd5b50813567ffffffffffffffff81111561439c575f5ffd5b6020830191508360208285010111156143b3575f5ffd5b9250929050565b5f5f5f5f606085870312156143cd575f5ffd5b6143d6856141e7565b935060208501356143e6816142a4565b9250604085013567ffffffffffffffff811115614401575f5ffd5b61440d87828801614375565b95989497509550505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215614457575f5ffd5b8235614462816142a4565b9150602083013567ffffffffffffffff81111561447d575f5ffd5b8301601f8101851361448d575f5ffd5b803567ffffffffffffffff8111156144a7576144a7614419565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff821117156144d7576144d7614419565b6040528181528282016020018710156144ee575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f6040838503121561451e575f5ffd5b50508035926020909101359150565b5f6020828403121561453d575f5ffd5b8135611e7b816142a4565b803565ffffffffffff811681146141f5575f5ffd5b5f6020828403121561456d575f5ffd5b611e7b82614548565b5f5f5f5f60608587031215614589575f5ffd5b843567ffffffffffffffff81111561459f575f5ffd5b6145ab87828801614375565b90955093505060208501356145bf816142a4565b9396929550929360400135925050565b8581526001600160a01b038516602082015260a060408201525f6145f660a08301866142e6565b6001600160a01b0394909416606083015250608001529392505050565b602081525f611e7b60208301846142e6565b60038110610c9a575f5ffd5b5f5f60408385031215614642575f5ffd5b8235915060208301356142db81614625565b5f5f60408385031215614665575f5ffd5b8235614670816142a4565b946020939093013593505050565b5f5f5f60608486031215614690575f5ffd5b8335925060208401356146a2816142a4565b915060408401356146b2816142a4565b809150509250925092565b5f5f5f606084860312156146cf575f5ffd5b83356146da816142a4565b925060208401356146ea816142a4565b91506146f860408501614548565b90509250925092565b5f5f5f5f5f60808688031215614715575f5ffd5b853567ffffffffffffffff81111561472b575f5ffd5b61473788828901614375565b909650945050602086013561474b816142a4565b9250604086013561475b816142a4565b949793965091946060013592915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036147f6576147f6614799565b5060010190565b600181811c9082168061481157607f821691505b602082108103614848577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b6001600160a01b0384168152604060208201525f61405860408301848661484e565b604081016148a78285614282565b8260208301529392505050565b6001600160a01b0383168152604060208201525f61363260408301846142e6565b601f82111561145957805f5260205f20601f840160051c810160208510156148fa5750805b601f840160051c820191505b81811015614919575f8155600101614906565b5050505050565b815167ffffffffffffffff81111561493a5761493a614419565b61494e8161494884546147fd565b846148d5565b6020601f82116001811461499f575f83156149695750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455614919565b5f84815260208120601f198516915b828110156149ce57878501518255602094850194600190920191016149ae565b5084821015614a0a57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b602081525f61363260208301848661484e565b6020810160038310614a4057614a40614255565b91905290565b83815260608101614a5a6020830185614282565b6136326040830184614282565b5f60208284031215614a77575f5ffd5b5051919050565b5f60208284031215614a8e575f5ffd5b8151611e7b81614625565b65ffffffffffff8181168382160190811115610c8257610c82614799565b65ffffffffffff8281168282160390811115610c8257610c82614799565b5f82518060208501845e5f92019182525091905056fe60806040526040516103dc3803806103dc8339810160408190526100229161023b565b61002c8282610033565b5050610320565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030a565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020e57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024c575f5ffd5b82516001600160a01b0381168114610262575f5ffd5b60208401519092506001600160401b0381111561027d575f5ffd5b8301601f8101851361028d575f5ffd5b80516001600160401b038111156102a6576102a6610227565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d4576102d4610227565b6040528181528282016020018710156102eb575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032c5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220927d2839b1d9f37b4bfae40308f9fa5bd1e749b8d8b98c249d9b67e246bf55ba64736f6c634300081c0033608060405234801561000f575f5ffd5b506040518060400160405280600d81526020016c577261707065642043524f535360981b815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100649190610111565b5060046100718282610111565b5050506101cb565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100a157607f821691505b6020821081036100bf57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561010c57805f5260205f20601f840160051c810160208510156100ea5750805b601f840160051c820191505b81811015610109575f81556001016100f6565b50505b505050565b81516001600160401b0381111561012a5761012a610079565b61013e81610138845461008d565b846100c5565b6020601f821160018114610170575f83156101595750848201515b5f19600385901b1c1916600184901b178455610109565b5f84815260208120601f198516915b8281101561019f578785015182556020948501946001909201910161017f565b50848210156101bc57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610bbb806101d85f395ff3fe6080604052600436106100c6575f3560e01c8063313ce56711610071578063a9059cbb1161004c578063a9059cbb1461021d578063d0e30db01461023c578063dd62ed3e14610244575f5ffd5b8063313ce567146101ad57806370a08231146101c857806395d89b4114610209575f5ffd5b8063205c2878116100a1578063205c28781461015057806323b872dd1461016f5780632e1a7d4d1461018e575f5ffd5b806306fdde03146100d9578063095ea7b31461010357806318160ddd14610132575f5ffd5b366100d5576100d3610295565b005b5f5ffd5b3480156100e4575f5ffd5b506100ed6102a7565b6040516100fa91906109b7565b60405180910390f35b34801561010e575f5ffd5b5061012261011d366004610a32565b610337565b60405190151581526020016100fa565b34801561013d575f5ffd5b506002545b6040519081526020016100fa565b34801561015b575f5ffd5b506100d361016a366004610a32565b610350565b34801561017a575f5ffd5b50610122610189366004610a5a565b610442565b348015610199575f5ffd5b506100d36101a8366004610a94565b610465565b3480156101b8575f5ffd5b50604051601281526020016100fa565b3480156101d3575f5ffd5b506101426101e2366004610aab565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b348015610214575f5ffd5b506100ed610472565b348015610228575f5ffd5b50610122610237366004610a32565b610481565b6100d3610295565b34801561024f575f5ffd5b5061014261025e366004610acb565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b34156102a5576102a5333461048e565b565b6060600380546102b690610afc565b80601f01602080910402602001604051908101604052809291908181526020018280546102e290610afc565b801561032d5780601f106103045761010080835404028352916020019161032d565b820191905f5260205f20905b81548152906001019060200180831161031057829003601f168201915b5050505050905090565b5f336103448185856104f1565b60019150505b92915050565b73ffffffffffffffffffffffffffffffffffffffff821661039d576040517f653345a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103a733826104fe565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f81146103fd576040519150601f19603f3d011682016040523d82523d5f602084013e610402565b606091505b505090508061043d576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f3361044f858285610558565b61045a858585610626565b506001949350505050565b61046f3382610350565b50565b6060600480546102b690610afc565b5f33610344818585610626565b73ffffffffffffffffffffffffffffffffffffffff82166104e2576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6104ed5f83836106cb565b5050565b61043d8383836001610872565b73ffffffffffffffffffffffffffffffffffffffff821661054d576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016104d9565b6104ed825f836106cb565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156106205781811015610612576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016104d9565b61062084848484035f610872565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610675576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016104d9565b73ffffffffffffffffffffffffffffffffffffffff82166106c4576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016104d9565b61043d8383835b73ffffffffffffffffffffffffffffffffffffffff8316610702578060025f8282546106f79190610b4d565b909155506107b29050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610787576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016104d9565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166107db57600280548290039055610806565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161086591815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff84166108c1576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016104d9565b73ffffffffffffffffffffffffffffffffffffffff8316610910576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016104d9565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610620578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516109a991815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610a2d575f5ffd5b919050565b5f5f60408385031215610a43575f5ffd5b610a4c83610a0a565b946020939093013593505050565b5f5f5f60608486031215610a6c575f5ffd5b610a7584610a0a565b9250610a8360208501610a0a565b929592945050506040919091013590565b5f60208284031215610aa4575f5ffd5b5035919050565b5f60208284031215610abb575f5ffd5b610ac482610a0a565b9392505050565b5f5f60408385031215610adc575f5ffd5b610ae583610a0a565b9150610af360208401610a0a565b90509250929050565b600181811c90821680610b1057607f821691505b602082108103610b47577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8082018082111561034a577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220575452f56f036be07f08a88f74a0472beab5a9974ca317cbb044dd21be9192c864736f6c634300081c0033a26469706673582212209e72833559582117e1303a9ad3b8112ae4730d4ff0a56f4980e801dfca54f69a64736f6c634300081c0033",
}

// CrossGameRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossGameRewardMetaData.ABI instead.
var CrossGameRewardABI = CrossGameRewardMetaData.ABI

// Deprecated: Use CrossGameRewardMetaData.Sigs instead.
// CrossGameRewardFuncSigs maps the 4-byte function signature to its string representation.
var CrossGameRewardFuncSigs = CrossGameRewardMetaData.Sigs

// CrossGameRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossGameRewardMetaData.Bin instead.
var CrossGameRewardBin = CrossGameRewardMetaData.Bin

// DeployCrossGameReward deploys a new Ethereum contract, binding an instance of CrossGameReward to it.
func DeployCrossGameReward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossGameReward, error) {
	parsed, err := CrossGameRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossGameRewardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossGameReward{CrossGameRewardCaller: CrossGameRewardCaller{contract: contract}, CrossGameRewardTransactor: CrossGameRewardTransactor{contract: contract}, CrossGameRewardFilterer: CrossGameRewardFilterer{contract: contract}}, nil
}

// CrossGameReward is an auto generated Go binding around an Ethereum contract.
type CrossGameReward struct {
	CrossGameRewardCaller     // Read-only binding to the contract
	CrossGameRewardTransactor // Write-only binding to the contract
	CrossGameRewardFilterer   // Log filterer for contract events
}

// CrossGameRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossGameRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossGameRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossGameRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossGameRewardSession struct {
	Contract     *CrossGameReward  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossGameRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossGameRewardCallerSession struct {
	Contract *CrossGameRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CrossGameRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossGameRewardTransactorSession struct {
	Contract     *CrossGameRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CrossGameRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossGameRewardRaw struct {
	Contract *CrossGameReward // Generic contract binding to access the raw methods on
}

// CrossGameRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossGameRewardCallerRaw struct {
	Contract *CrossGameRewardCaller // Generic read-only contract binding to access the raw methods on
}

// CrossGameRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossGameRewardTransactorRaw struct {
	Contract *CrossGameRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossGameReward creates a new instance of CrossGameReward, bound to a specific deployed contract.
func NewCrossGameReward(address common.Address, backend bind.ContractBackend) (*CrossGameReward, error) {
	contract, err := bindCrossGameReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossGameReward{CrossGameRewardCaller: CrossGameRewardCaller{contract: contract}, CrossGameRewardTransactor: CrossGameRewardTransactor{contract: contract}, CrossGameRewardFilterer: CrossGameRewardFilterer{contract: contract}}, nil
}

// NewCrossGameRewardCaller creates a new read-only instance of CrossGameReward, bound to a specific deployed contract.
func NewCrossGameRewardCaller(address common.Address, caller bind.ContractCaller) (*CrossGameRewardCaller, error) {
	contract, err := bindCrossGameReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardCaller{contract: contract}, nil
}

// NewCrossGameRewardTransactor creates a new write-only instance of CrossGameReward, bound to a specific deployed contract.
func NewCrossGameRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossGameRewardTransactor, error) {
	contract, err := bindCrossGameReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardTransactor{contract: contract}, nil
}

// NewCrossGameRewardFilterer creates a new log filterer instance of CrossGameReward, bound to a specific deployed contract.
func NewCrossGameRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossGameRewardFilterer, error) {
	contract, err := bindCrossGameReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardFilterer{contract: contract}, nil
}

// bindCrossGameReward binds a generic wrapper to an already deployed contract.
func bindCrossGameReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossGameRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossGameReward *CrossGameRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossGameReward.Contract.CrossGameRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossGameReward *CrossGameRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CrossGameRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossGameReward *CrossGameRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CrossGameRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossGameReward *CrossGameRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossGameReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossGameReward *CrossGameRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossGameReward *CrossGameRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossGameReward.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossGameReward.Contract.DEFAULTADMINROLE(&_CrossGameReward.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossGameReward.Contract.DEFAULTADMINROLE(&_CrossGameReward.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardSession) MANAGERROLE() ([32]byte, error) {
	return _CrossGameReward.Contract.MANAGERROLE(&_CrossGameReward.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCallerSession) MANAGERROLE() ([32]byte, error) {
	return _CrossGameReward.Contract.MANAGERROLE(&_CrossGameReward.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossGameReward *CrossGameRewardCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossGameReward *CrossGameRewardSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossGameReward.Contract.UPGRADEINTERFACEVERSION(&_CrossGameReward.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossGameReward *CrossGameRewardCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossGameReward.Contract.UPGRADEINTERFACEVERSION(&_CrossGameReward.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossGameReward *CrossGameRewardSession) DefaultAdmin() (common.Address, error) {
	return _CrossGameReward.Contract.DefaultAdmin(&_CrossGameReward.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) DefaultAdmin() (common.Address, error) {
	return _CrossGameReward.Contract.DefaultAdmin(&_CrossGameReward.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossGameReward *CrossGameRewardCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossGameReward *CrossGameRewardSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossGameReward.Contract.DefaultAdminDelay(&_CrossGameReward.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossGameReward *CrossGameRewardCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossGameReward.Contract.DefaultAdminDelay(&_CrossGameReward.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossGameReward *CrossGameRewardCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossGameReward *CrossGameRewardSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossGameReward.Contract.DefaultAdminDelayIncreaseWait(&_CrossGameReward.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossGameReward *CrossGameRewardCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossGameReward.Contract.DefaultAdminDelayIncreaseWait(&_CrossGameReward.CallOpts)
}

// GamePoolImplementation is a free data retrieval call binding the contract method 0x2a21f7ae.
//
// Solidity: function gamePoolImplementation() view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) GamePoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "gamePoolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GamePoolImplementation is a free data retrieval call binding the contract method 0x2a21f7ae.
//
// Solidity: function gamePoolImplementation() view returns(address)
func (_CrossGameReward *CrossGameRewardSession) GamePoolImplementation() (common.Address, error) {
	return _CrossGameReward.Contract.GamePoolImplementation(&_CrossGameReward.CallOpts)
}

// GamePoolImplementation is a free data retrieval call binding the contract method 0x2a21f7ae.
//
// Solidity: function gamePoolImplementation() view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) GamePoolImplementation() (common.Address, error) {
	return _CrossGameReward.Contract.GamePoolImplementation(&_CrossGameReward.CallOpts)
}

// GetActivePoolIds is a free data retrieval call binding the contract method 0xfe96e4ff.
//
// Solidity: function getActivePoolIds() view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCaller) GetActivePoolIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getActivePoolIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActivePoolIds is a free data retrieval call binding the contract method 0xfe96e4ff.
//
// Solidity: function getActivePoolIds() view returns(uint256[])
func (_CrossGameReward *CrossGameRewardSession) GetActivePoolIds() ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetActivePoolIds(&_CrossGameReward.CallOpts)
}

// GetActivePoolIds is a free data retrieval call binding the contract method 0xfe96e4ff.
//
// Solidity: function getActivePoolIds() view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCallerSession) GetActivePoolIds() ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetActivePoolIds(&_CrossGameReward.CallOpts)
}

// GetAllPoolIds is a free data retrieval call binding the contract method 0xf19c3d5b.
//
// Solidity: function getAllPoolIds() view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCaller) GetAllPoolIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getAllPoolIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllPoolIds is a free data retrieval call binding the contract method 0xf19c3d5b.
//
// Solidity: function getAllPoolIds() view returns(uint256[])
func (_CrossGameReward *CrossGameRewardSession) GetAllPoolIds() ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetAllPoolIds(&_CrossGameReward.CallOpts)
}

// GetAllPoolIds is a free data retrieval call binding the contract method 0xf19c3d5b.
//
// Solidity: function getAllPoolIds() view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCallerSession) GetAllPoolIds() ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetAllPoolIds(&_CrossGameReward.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x00a5ae21.
//
// Solidity: function getPoolAddress(uint256 poolId) view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) GetPoolAddress(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolAddress", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAddress is a free data retrieval call binding the contract method 0x00a5ae21.
//
// Solidity: function getPoolAddress(uint256 poolId) view returns(address)
func (_CrossGameReward *CrossGameRewardSession) GetPoolAddress(poolId *big.Int) (common.Address, error) {
	return _CrossGameReward.Contract.GetPoolAddress(&_CrossGameReward.CallOpts, poolId)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x00a5ae21.
//
// Solidity: function getPoolAddress(uint256 poolId) view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolAddress(poolId *big.Int) (common.Address, error) {
	return _CrossGameReward.Contract.GetPoolAddress(&_CrossGameReward.CallOpts, poolId)
}

// GetPoolCountByDepositToken is a free data retrieval call binding the contract method 0xd4148bcd.
//
// Solidity: function getPoolCountByDepositToken(address depositToken) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) GetPoolCountByDepositToken(opts *bind.CallOpts, depositToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolCountByDepositToken", depositToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolCountByDepositToken is a free data retrieval call binding the contract method 0xd4148bcd.
//
// Solidity: function getPoolCountByDepositToken(address depositToken) view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) GetPoolCountByDepositToken(depositToken common.Address) (*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolCountByDepositToken(&_CrossGameReward.CallOpts, depositToken)
}

// GetPoolCountByDepositToken is a free data retrieval call binding the contract method 0xd4148bcd.
//
// Solidity: function getPoolCountByDepositToken(address depositToken) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolCountByDepositToken(depositToken common.Address) (*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolCountByDepositToken(&_CrossGameReward.CallOpts, depositToken)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address pool) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) GetPoolId(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolId", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address pool) view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) GetPoolId(pool common.Address) (*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolId(&_CrossGameReward.CallOpts, pool)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address pool) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolId(pool common.Address) (*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolId(&_CrossGameReward.CallOpts, pool)
}

// GetPoolIdsByDepositToken is a free data retrieval call binding the contract method 0xeeea4a79.
//
// Solidity: function getPoolIdsByDepositToken(address depositToken) view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCaller) GetPoolIdsByDepositToken(opts *bind.CallOpts, depositToken common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolIdsByDepositToken", depositToken)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPoolIdsByDepositToken is a free data retrieval call binding the contract method 0xeeea4a79.
//
// Solidity: function getPoolIdsByDepositToken(address depositToken) view returns(uint256[])
func (_CrossGameReward *CrossGameRewardSession) GetPoolIdsByDepositToken(depositToken common.Address) ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolIdsByDepositToken(&_CrossGameReward.CallOpts, depositToken)
}

// GetPoolIdsByDepositToken is a free data retrieval call binding the contract method 0xeeea4a79.
//
// Solidity: function getPoolIdsByDepositToken(address depositToken) view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolIdsByDepositToken(depositToken common.Address) ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolIdsByDepositToken(&_CrossGameReward.CallOpts, depositToken)
}

// GetPoolIdsByType is a free data retrieval call binding the contract method 0x0bee20d8.
//
// Solidity: function getPoolIdsByType(uint8 poolType) view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCaller) GetPoolIdsByType(opts *bind.CallOpts, poolType uint8) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolIdsByType", poolType)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPoolIdsByType is a free data retrieval call binding the contract method 0x0bee20d8.
//
// Solidity: function getPoolIdsByType(uint8 poolType) view returns(uint256[])
func (_CrossGameReward *CrossGameRewardSession) GetPoolIdsByType(poolType uint8) ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolIdsByType(&_CrossGameReward.CallOpts, poolType)
}

// GetPoolIdsByType is a free data retrieval call binding the contract method 0x0bee20d8.
//
// Solidity: function getPoolIdsByType(uint8 poolType) view returns(uint256[])
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolIdsByType(poolType uint8) ([]*big.Int, error) {
	return _CrossGameReward.Contract.GetPoolIdsByType(&_CrossGameReward.CallOpts, poolType)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,string,address,uint256))
func (_CrossGameReward *CrossGameRewardCaller) GetPoolInfo(opts *bind.CallOpts, poolId *big.Int) (ICrossGameRewardPoolInfo, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolInfo", poolId)

	if err != nil {
		return *new(ICrossGameRewardPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossGameRewardPoolInfo)).(*ICrossGameRewardPoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,string,address,uint256))
func (_CrossGameReward *CrossGameRewardSession) GetPoolInfo(poolId *big.Int) (ICrossGameRewardPoolInfo, error) {
	return _CrossGameReward.Contract.GetPoolInfo(&_CrossGameReward.CallOpts, poolId)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,string,address,uint256))
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolInfo(poolId *big.Int) (ICrossGameRewardPoolInfo, error) {
	return _CrossGameReward.Contract.GetPoolInfo(&_CrossGameReward.CallOpts, poolId)
}

// GetPoolType is a free data retrieval call binding the contract method 0xcdcf8783.
//
// Solidity: function getPoolType(uint256 poolId) view returns(uint8)
func (_CrossGameReward *CrossGameRewardCaller) GetPoolType(opts *bind.CallOpts, poolId *big.Int) (uint8, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getPoolType", poolId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPoolType is a free data retrieval call binding the contract method 0xcdcf8783.
//
// Solidity: function getPoolType(uint256 poolId) view returns(uint8)
func (_CrossGameReward *CrossGameRewardSession) GetPoolType(poolId *big.Int) (uint8, error) {
	return _CrossGameReward.Contract.GetPoolType(&_CrossGameReward.CallOpts, poolId)
}

// GetPoolType is a free data retrieval call binding the contract method 0xcdcf8783.
//
// Solidity: function getPoolType(uint256 poolId) view returns(uint8)
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolType(poolId *big.Int) (uint8, error) {
	return _CrossGameReward.Contract.GetPoolType(&_CrossGameReward.CallOpts, poolId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossGameReward *CrossGameRewardSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossGameReward.Contract.GetRoleAdmin(&_CrossGameReward.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossGameReward.Contract.GetRoleAdmin(&_CrossGameReward.CallOpts, role)
}

// GetTotalPoolCount is a free data retrieval call binding the contract method 0xe7590268.
//
// Solidity: function getTotalPoolCount() view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) GetTotalPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "getTotalPoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPoolCount is a free data retrieval call binding the contract method 0xe7590268.
//
// Solidity: function getTotalPoolCount() view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) GetTotalPoolCount() (*big.Int, error) {
	return _CrossGameReward.Contract.GetTotalPoolCount(&_CrossGameReward.CallOpts)
}

// GetTotalPoolCount is a free data retrieval call binding the contract method 0xe7590268.
//
// Solidity: function getTotalPoolCount() view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) GetTotalPoolCount() (*big.Int, error) {
	return _CrossGameReward.Contract.GetTotalPoolCount(&_CrossGameReward.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossGameReward *CrossGameRewardCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossGameReward *CrossGameRewardSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossGameReward.Contract.HasRole(&_CrossGameReward.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossGameReward *CrossGameRewardCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossGameReward.Contract.HasRole(&_CrossGameReward.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) InitializedAt() (*big.Int, error) {
	return _CrossGameReward.Contract.InitializedAt(&_CrossGameReward.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) InitializedAt() (*big.Int, error) {
	return _CrossGameReward.Contract.InitializedAt(&_CrossGameReward.CallOpts)
}

// NextPoolId is a free data retrieval call binding the contract method 0x18e56131.
//
// Solidity: function nextPoolId() view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) NextPoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "nextPoolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPoolId is a free data retrieval call binding the contract method 0x18e56131.
//
// Solidity: function nextPoolId() view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) NextPoolId() (*big.Int, error) {
	return _CrossGameReward.Contract.NextPoolId(&_CrossGameReward.CallOpts)
}

// NextPoolId is a free data retrieval call binding the contract method 0x18e56131.
//
// Solidity: function nextPoolId() view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) NextPoolId() (*big.Int, error) {
	return _CrossGameReward.Contract.NextPoolId(&_CrossGameReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossGameReward *CrossGameRewardSession) Owner() (common.Address, error) {
	return _CrossGameReward.Contract.Owner(&_CrossGameReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) Owner() (common.Address, error) {
	return _CrossGameReward.Contract.Owner(&_CrossGameReward.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossGameReward *CrossGameRewardCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossGameReward *CrossGameRewardSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossGameReward.Contract.PendingDefaultAdmin(&_CrossGameReward.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossGameReward *CrossGameRewardCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossGameReward.Contract.PendingDefaultAdmin(&_CrossGameReward.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossGameReward *CrossGameRewardCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossGameReward *CrossGameRewardSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossGameReward.Contract.PendingDefaultAdminDelay(&_CrossGameReward.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossGameReward *CrossGameRewardCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossGameReward.Contract.PendingDefaultAdminDelay(&_CrossGameReward.CallOpts)
}

// PoolAt is a free data retrieval call binding the contract method 0x155fff62.
//
// Solidity: function poolAt(uint256 index) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) PoolAt(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "poolAt", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAt is a free data retrieval call binding the contract method 0x155fff62.
//
// Solidity: function poolAt(uint256 index) view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) PoolAt(index *big.Int) (*big.Int, error) {
	return _CrossGameReward.Contract.PoolAt(&_CrossGameReward.CallOpts, index)
}

// PoolAt is a free data retrieval call binding the contract method 0x155fff62.
//
// Solidity: function poolAt(uint256 index) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) PoolAt(index *big.Int) (*big.Int, error) {
	return _CrossGameReward.Contract.PoolAt(&_CrossGameReward.CallOpts, index)
}

// PoolByDepositTokenAt is a free data retrieval call binding the contract method 0xb5be3221.
//
// Solidity: function poolByDepositTokenAt(address depositToken, uint256 index) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) PoolByDepositTokenAt(opts *bind.CallOpts, depositToken common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "poolByDepositTokenAt", depositToken, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolByDepositTokenAt is a free data retrieval call binding the contract method 0xb5be3221.
//
// Solidity: function poolByDepositTokenAt(address depositToken, uint256 index) view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) PoolByDepositTokenAt(depositToken common.Address, index *big.Int) (*big.Int, error) {
	return _CrossGameReward.Contract.PoolByDepositTokenAt(&_CrossGameReward.CallOpts, depositToken, index)
}

// PoolByDepositTokenAt is a free data retrieval call binding the contract method 0xb5be3221.
//
// Solidity: function poolByDepositTokenAt(address depositToken, uint256 index) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) PoolByDepositTokenAt(depositToken common.Address, index *big.Int) (*big.Int, error) {
	return _CrossGameReward.Contract.PoolByDepositTokenAt(&_CrossGameReward.CallOpts, depositToken, index)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCaller) PoolIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "poolIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint256)
func (_CrossGameReward *CrossGameRewardSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _CrossGameReward.Contract.PoolIds(&_CrossGameReward.CallOpts, arg0)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint256)
func (_CrossGameReward *CrossGameRewardCallerSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _CrossGameReward.Contract.PoolIds(&_CrossGameReward.CallOpts, arg0)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) PoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "poolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_CrossGameReward *CrossGameRewardSession) PoolImplementation() (common.Address, error) {
	return _CrossGameReward.Contract.PoolImplementation(&_CrossGameReward.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) PoolImplementation() (common.Address, error) {
	return _CrossGameReward.Contract.PoolImplementation(&_CrossGameReward.CallOpts)
}

// PoolTypes is a free data retrieval call binding the contract method 0x1b95a010.
//
// Solidity: function poolTypes(uint256 ) view returns(uint8)
func (_CrossGameReward *CrossGameRewardCaller) PoolTypes(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "poolTypes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolTypes is a free data retrieval call binding the contract method 0x1b95a010.
//
// Solidity: function poolTypes(uint256 ) view returns(uint8)
func (_CrossGameReward *CrossGameRewardSession) PoolTypes(arg0 *big.Int) (uint8, error) {
	return _CrossGameReward.Contract.PoolTypes(&_CrossGameReward.CallOpts, arg0)
}

// PoolTypes is a free data retrieval call binding the contract method 0x1b95a010.
//
// Solidity: function poolTypes(uint256 ) view returns(uint8)
func (_CrossGameReward *CrossGameRewardCallerSession) PoolTypes(arg0 *big.Int) (uint8, error) {
	return _CrossGameReward.Contract.PoolTypes(&_CrossGameReward.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, string name, address depositToken, uint256 createdAt)
func (_CrossGameReward *CrossGameRewardCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	Name         string
	DepositToken common.Address
	CreatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "pools", arg0)

	outstruct := new(struct {
		PoolId       *big.Int
		Pool         common.Address
		Name         string
		DepositToken common.Address
		CreatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.DepositToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, string name, address depositToken, uint256 createdAt)
func (_CrossGameReward *CrossGameRewardSession) Pools(arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	Name         string
	DepositToken common.Address
	CreatedAt    *big.Int
}, error) {
	return _CrossGameReward.Contract.Pools(&_CrossGameReward.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, string name, address depositToken, uint256 createdAt)
func (_CrossGameReward *CrossGameRewardCallerSession) Pools(arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	Name         string
	DepositToken common.Address
	CreatedAt    *big.Int
}, error) {
	return _CrossGameReward.Contract.Pools(&_CrossGameReward.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardSession) ProxiableUUID() ([32]byte, error) {
	return _CrossGameReward.Contract.ProxiableUUID(&_CrossGameReward.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossGameReward *CrossGameRewardCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossGameReward.Contract.ProxiableUUID(&_CrossGameReward.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_CrossGameReward *CrossGameRewardSession) Router() (common.Address, error) {
	return _CrossGameReward.Contract.Router(&_CrossGameReward.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) Router() (common.Address, error) {
	return _CrossGameReward.Contract.Router(&_CrossGameReward.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossGameReward *CrossGameRewardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossGameReward *CrossGameRewardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossGameReward.Contract.SupportsInterface(&_CrossGameReward.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossGameReward *CrossGameRewardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossGameReward.Contract.SupportsInterface(&_CrossGameReward.CallOpts, interfaceId)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossGameReward *CrossGameRewardCaller) Wcross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "wcross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossGameReward *CrossGameRewardSession) Wcross() (common.Address, error) {
	return _CrossGameReward.Contract.Wcross(&_CrossGameReward.CallOpts)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossGameReward *CrossGameRewardCallerSession) Wcross() (common.Address, error) {
	return _CrossGameReward.Contract.Wcross(&_CrossGameReward.CallOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossGameReward *CrossGameRewardTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossGameReward *CrossGameRewardSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameReward.Contract.AcceptDefaultAdminTransfer(&_CrossGameReward.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameReward.Contract.AcceptDefaultAdminTransfer(&_CrossGameReward.TransactOpts)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa1635945.
//
// Solidity: function addRewardToken(uint256 poolId, address token) returns()
func (_CrossGameReward *CrossGameRewardTransactor) AddRewardToken(opts *bind.TransactOpts, poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "addRewardToken", poolId, token)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa1635945.
//
// Solidity: function addRewardToken(uint256 poolId, address token) returns()
func (_CrossGameReward *CrossGameRewardSession) AddRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.AddRewardToken(&_CrossGameReward.TransactOpts, poolId, token)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa1635945.
//
// Solidity: function addRewardToken(uint256 poolId, address token) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) AddRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.AddRewardToken(&_CrossGameReward.TransactOpts, poolId, token)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossGameReward *CrossGameRewardTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossGameReward *CrossGameRewardSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.BeginDefaultAdminTransfer(&_CrossGameReward.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.BeginDefaultAdminTransfer(&_CrossGameReward.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossGameReward *CrossGameRewardTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossGameReward *CrossGameRewardSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameReward.Contract.CancelDefaultAdminTransfer(&_CrossGameReward.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameReward.Contract.CancelDefaultAdminTransfer(&_CrossGameReward.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossGameReward *CrossGameRewardTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossGameReward *CrossGameRewardSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.ChangeDefaultAdminDelay(&_CrossGameReward.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.ChangeDefaultAdminDelay(&_CrossGameReward.TransactOpts, newDelay)
}

// CreateGamePool is a paid mutator transaction binding the contract method 0xe01085e5.
//
// Solidity: function createGamePool(string name, address depositToken, address rewardToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardTransactor) CreateGamePool(opts *bind.TransactOpts, name string, depositToken common.Address, rewardToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "createGamePool", name, depositToken, rewardToken, minDepositAmount)
}

// CreateGamePool is a paid mutator transaction binding the contract method 0xe01085e5.
//
// Solidity: function createGamePool(string name, address depositToken, address rewardToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardSession) CreateGamePool(name string, depositToken common.Address, rewardToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CreateGamePool(&_CrossGameReward.TransactOpts, name, depositToken, rewardToken, minDepositAmount)
}

// CreateGamePool is a paid mutator transaction binding the contract method 0xe01085e5.
//
// Solidity: function createGamePool(string name, address depositToken, address rewardToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardTransactorSession) CreateGamePool(name string, depositToken common.Address, rewardToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CreateGamePool(&_CrossGameReward.TransactOpts, name, depositToken, rewardToken, minDepositAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x6e13ba6f.
//
// Solidity: function createPool(string name, address depositToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardTransactor) CreatePool(opts *bind.TransactOpts, name string, depositToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "createPool", name, depositToken, minDepositAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x6e13ba6f.
//
// Solidity: function createPool(string name, address depositToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardSession) CreatePool(name string, depositToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CreatePool(&_CrossGameReward.TransactOpts, name, depositToken, minDepositAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x6e13ba6f.
//
// Solidity: function createPool(string name, address depositToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardTransactorSession) CreatePool(name string, depositToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CreatePool(&_CrossGameReward.TransactOpts, name, depositToken, minDepositAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.GrantRole(&_CrossGameReward.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.GrantRole(&_CrossGameReward.TransactOpts, role, account)
}

// GrantSponsorRole is a paid mutator transaction binding the contract method 0xed681780.
//
// Solidity: function grantSponsorRole(uint256 poolId, address sponsor) returns()
func (_CrossGameReward *CrossGameRewardTransactor) GrantSponsorRole(opts *bind.TransactOpts, poolId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "grantSponsorRole", poolId, sponsor)
}

// GrantSponsorRole is a paid mutator transaction binding the contract method 0xed681780.
//
// Solidity: function grantSponsorRole(uint256 poolId, address sponsor) returns()
func (_CrossGameReward *CrossGameRewardSession) GrantSponsorRole(poolId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.GrantSponsorRole(&_CrossGameReward.TransactOpts, poolId, sponsor)
}

// GrantSponsorRole is a paid mutator transaction binding the contract method 0xed681780.
//
// Solidity: function grantSponsorRole(uint256 poolId, address sponsor) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) GrantSponsorRole(poolId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.GrantSponsorRole(&_CrossGameReward.TransactOpts, poolId, sponsor)
}

// Initialize is a paid mutator transaction binding the contract method 0xce24af53.
//
// Solidity: function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) returns()
func (_CrossGameReward *CrossGameRewardTransactor) Initialize(opts *bind.TransactOpts, _poolImplementation common.Address, _admin common.Address, _initialDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "initialize", _poolImplementation, _admin, _initialDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0xce24af53.
//
// Solidity: function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) returns()
func (_CrossGameReward *CrossGameRewardSession) Initialize(_poolImplementation common.Address, _admin common.Address, _initialDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.Initialize(&_CrossGameReward.TransactOpts, _poolImplementation, _admin, _initialDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0xce24af53.
//
// Solidity: function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) Initialize(_poolImplementation common.Address, _admin common.Address, _initialDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.Initialize(&_CrossGameReward.TransactOpts, _poolImplementation, _admin, _initialDelay)
}

// ReclaimFromPool is a paid mutator transaction binding the contract method 0xc24140b2.
//
// Solidity: function reclaimFromPool(uint256 poolId, address token, address to) returns()
func (_CrossGameReward *CrossGameRewardTransactor) ReclaimFromPool(opts *bind.TransactOpts, poolId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "reclaimFromPool", poolId, token, to)
}

// ReclaimFromPool is a paid mutator transaction binding the contract method 0xc24140b2.
//
// Solidity: function reclaimFromPool(uint256 poolId, address token, address to) returns()
func (_CrossGameReward *CrossGameRewardSession) ReclaimFromPool(poolId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.ReclaimFromPool(&_CrossGameReward.TransactOpts, poolId, token, to)
}

// ReclaimFromPool is a paid mutator transaction binding the contract method 0xc24140b2.
//
// Solidity: function reclaimFromPool(uint256 poolId, address token, address to) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) ReclaimFromPool(poolId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.ReclaimFromPool(&_CrossGameReward.TransactOpts, poolId, token, to)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x35cc9cb4.
//
// Solidity: function removeRewardToken(uint256 poolId, address token) returns()
func (_CrossGameReward *CrossGameRewardTransactor) RemoveRewardToken(opts *bind.TransactOpts, poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "removeRewardToken", poolId, token)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x35cc9cb4.
//
// Solidity: function removeRewardToken(uint256 poolId, address token) returns()
func (_CrossGameReward *CrossGameRewardSession) RemoveRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RemoveRewardToken(&_CrossGameReward.TransactOpts, poolId, token)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x35cc9cb4.
//
// Solidity: function removeRewardToken(uint256 poolId, address token) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) RemoveRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RemoveRewardToken(&_CrossGameReward.TransactOpts, poolId, token)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RenounceRole(&_CrossGameReward.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RenounceRole(&_CrossGameReward.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RevokeRole(&_CrossGameReward.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RevokeRole(&_CrossGameReward.TransactOpts, role, account)
}

// RevokeSponsorRole is a paid mutator transaction binding the contract method 0xb4368ae8.
//
// Solidity: function revokeSponsorRole(uint256 poolId, address sponsor) returns()
func (_CrossGameReward *CrossGameRewardTransactor) RevokeSponsorRole(opts *bind.TransactOpts, poolId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "revokeSponsorRole", poolId, sponsor)
}

// RevokeSponsorRole is a paid mutator transaction binding the contract method 0xb4368ae8.
//
// Solidity: function revokeSponsorRole(uint256 poolId, address sponsor) returns()
func (_CrossGameReward *CrossGameRewardSession) RevokeSponsorRole(poolId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RevokeSponsorRole(&_CrossGameReward.TransactOpts, poolId, sponsor)
}

// RevokeSponsorRole is a paid mutator transaction binding the contract method 0xb4368ae8.
//
// Solidity: function revokeSponsorRole(uint256 poolId, address sponsor) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) RevokeSponsorRole(poolId *big.Int, sponsor common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.RevokeSponsorRole(&_CrossGameReward.TransactOpts, poolId, sponsor)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossGameReward *CrossGameRewardTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossGameReward *CrossGameRewardSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossGameReward.Contract.RollbackDefaultAdminDelay(&_CrossGameReward.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossGameReward.Contract.RollbackDefaultAdminDelay(&_CrossGameReward.TransactOpts)
}

// SetGamePoolImplementation is a paid mutator transaction binding the contract method 0xf62d31fd.
//
// Solidity: function setGamePoolImplementation(address newImplementation) returns()
func (_CrossGameReward *CrossGameRewardTransactor) SetGamePoolImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "setGamePoolImplementation", newImplementation)
}

// SetGamePoolImplementation is a paid mutator transaction binding the contract method 0xf62d31fd.
//
// Solidity: function setGamePoolImplementation(address newImplementation) returns()
func (_CrossGameReward *CrossGameRewardSession) SetGamePoolImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetGamePoolImplementation(&_CrossGameReward.TransactOpts, newImplementation)
}

// SetGamePoolImplementation is a paid mutator transaction binding the contract method 0xf62d31fd.
//
// Solidity: function setGamePoolImplementation(address newImplementation) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) SetGamePoolImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetGamePoolImplementation(&_CrossGameReward.TransactOpts, newImplementation)
}

// SetMaxActiveRounds is a paid mutator transaction binding the contract method 0xdecbb6e7.
//
// Solidity: function setMaxActiveRounds(uint256 poolId, uint256 newMax) returns()
func (_CrossGameReward *CrossGameRewardTransactor) SetMaxActiveRounds(opts *bind.TransactOpts, poolId *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "setMaxActiveRounds", poolId, newMax)
}

// SetMaxActiveRounds is a paid mutator transaction binding the contract method 0xdecbb6e7.
//
// Solidity: function setMaxActiveRounds(uint256 poolId, uint256 newMax) returns()
func (_CrossGameReward *CrossGameRewardSession) SetMaxActiveRounds(poolId *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetMaxActiveRounds(&_CrossGameReward.TransactOpts, poolId, newMax)
}

// SetMaxActiveRounds is a paid mutator transaction binding the contract method 0xdecbb6e7.
//
// Solidity: function setMaxActiveRounds(uint256 poolId, uint256 newMax) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) SetMaxActiveRounds(poolId *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetMaxActiveRounds(&_CrossGameReward.TransactOpts, poolId, newMax)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0xd6f74898.
//
// Solidity: function setPoolImplementation(address newImplementation) returns()
func (_CrossGameReward *CrossGameRewardTransactor) SetPoolImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "setPoolImplementation", newImplementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0xd6f74898.
//
// Solidity: function setPoolImplementation(address newImplementation) returns()
func (_CrossGameReward *CrossGameRewardSession) SetPoolImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetPoolImplementation(&_CrossGameReward.TransactOpts, newImplementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0xd6f74898.
//
// Solidity: function setPoolImplementation(address newImplementation) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) SetPoolImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetPoolImplementation(&_CrossGameReward.TransactOpts, newImplementation)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0xb34c972e.
//
// Solidity: function setPoolStatus(uint256 poolId, uint8 status) returns()
func (_CrossGameReward *CrossGameRewardTransactor) SetPoolStatus(opts *bind.TransactOpts, poolId *big.Int, status uint8) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "setPoolStatus", poolId, status)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0xb34c972e.
//
// Solidity: function setPoolStatus(uint256 poolId, uint8 status) returns()
func (_CrossGameReward *CrossGameRewardSession) SetPoolStatus(poolId *big.Int, status uint8) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetPoolStatus(&_CrossGameReward.TransactOpts, poolId, status)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0xb34c972e.
//
// Solidity: function setPoolStatus(uint256 poolId, uint8 status) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) SetPoolStatus(poolId *big.Int, status uint8) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetPoolStatus(&_CrossGameReward.TransactOpts, poolId, status)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_CrossGameReward *CrossGameRewardTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_CrossGameReward *CrossGameRewardSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetRouter(&_CrossGameReward.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _CrossGameReward.Contract.SetRouter(&_CrossGameReward.TransactOpts, _router)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x61616c46.
//
// Solidity: function updateMinDepositAmount(uint256 poolId, uint256 amount) returns()
func (_CrossGameReward *CrossGameRewardTransactor) UpdateMinDepositAmount(opts *bind.TransactOpts, poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "updateMinDepositAmount", poolId, amount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x61616c46.
//
// Solidity: function updateMinDepositAmount(uint256 poolId, uint256 amount) returns()
func (_CrossGameReward *CrossGameRewardSession) UpdateMinDepositAmount(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.UpdateMinDepositAmount(&_CrossGameReward.TransactOpts, poolId, amount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x61616c46.
//
// Solidity: function updateMinDepositAmount(uint256 poolId, uint256 amount) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) UpdateMinDepositAmount(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.UpdateMinDepositAmount(&_CrossGameReward.TransactOpts, poolId, amount)
}

// UpgradePoolsByType is a paid mutator transaction binding the contract method 0x320533a0.
//
// Solidity: function upgradePoolsByType(uint8 poolType, address newImplementation, bytes data) returns()
func (_CrossGameReward *CrossGameRewardTransactor) UpgradePoolsByType(opts *bind.TransactOpts, poolType uint8, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "upgradePoolsByType", poolType, newImplementation, data)
}

// UpgradePoolsByType is a paid mutator transaction binding the contract method 0x320533a0.
//
// Solidity: function upgradePoolsByType(uint8 poolType, address newImplementation, bytes data) returns()
func (_CrossGameReward *CrossGameRewardSession) UpgradePoolsByType(poolType uint8, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameReward.Contract.UpgradePoolsByType(&_CrossGameReward.TransactOpts, poolType, newImplementation, data)
}

// UpgradePoolsByType is a paid mutator transaction binding the contract method 0x320533a0.
//
// Solidity: function upgradePoolsByType(uint8 poolType, address newImplementation, bytes data) returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) UpgradePoolsByType(poolType uint8, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameReward.Contract.UpgradePoolsByType(&_CrossGameReward.TransactOpts, poolType, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossGameReward *CrossGameRewardTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossGameReward *CrossGameRewardSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameReward.Contract.UpgradeToAndCall(&_CrossGameReward.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossGameReward *CrossGameRewardTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameReward.Contract.UpgradeToAndCall(&_CrossGameReward.TransactOpts, newImplementation, data)
}

// CrossGameRewardDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminDelayChangeCanceledIterator struct {
	Event *CrossGameRewardDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardDefaultAdminDelayChangeCanceled)
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
		it.Event = new(CrossGameRewardDefaultAdminDelayChangeCanceled)
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
func (it *CrossGameRewardDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossGameReward *CrossGameRewardFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*CrossGameRewardDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardDefaultAdminDelayChangeCanceledIterator{contract: _CrossGameReward.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossGameReward *CrossGameRewardFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardDefaultAdminDelayChangeCanceled)
				if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossGameReward *CrossGameRewardFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*CrossGameRewardDefaultAdminDelayChangeCanceled, error) {
	event := new(CrossGameRewardDefaultAdminDelayChangeCanceled)
	if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminDelayChangeScheduledIterator struct {
	Event *CrossGameRewardDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardDefaultAdminDelayChangeScheduled)
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
		it.Event = new(CrossGameRewardDefaultAdminDelayChangeScheduled)
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
func (it *CrossGameRewardDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossGameReward *CrossGameRewardFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*CrossGameRewardDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardDefaultAdminDelayChangeScheduledIterator{contract: _CrossGameReward.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossGameReward *CrossGameRewardFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardDefaultAdminDelayChangeScheduled)
				if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossGameReward *CrossGameRewardFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*CrossGameRewardDefaultAdminDelayChangeScheduled, error) {
	event := new(CrossGameRewardDefaultAdminDelayChangeScheduled)
	if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminTransferCanceledIterator struct {
	Event *CrossGameRewardDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardDefaultAdminTransferCanceled)
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
		it.Event = new(CrossGameRewardDefaultAdminTransferCanceled)
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
func (it *CrossGameRewardDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossGameReward *CrossGameRewardFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*CrossGameRewardDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardDefaultAdminTransferCanceledIterator{contract: _CrossGameReward.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossGameReward *CrossGameRewardFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardDefaultAdminTransferCanceled)
				if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossGameReward *CrossGameRewardFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*CrossGameRewardDefaultAdminTransferCanceled, error) {
	event := new(CrossGameRewardDefaultAdminTransferCanceled)
	if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminTransferScheduledIterator struct {
	Event *CrossGameRewardDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardDefaultAdminTransferScheduled)
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
		it.Event = new(CrossGameRewardDefaultAdminTransferScheduled)
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
func (it *CrossGameRewardDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the CrossGameReward contract.
type CrossGameRewardDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossGameReward *CrossGameRewardFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*CrossGameRewardDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardDefaultAdminTransferScheduledIterator{contract: _CrossGameReward.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossGameReward *CrossGameRewardFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardDefaultAdminTransferScheduled)
				if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossGameReward *CrossGameRewardFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*CrossGameRewardDefaultAdminTransferScheduled, error) {
	event := new(CrossGameRewardDefaultAdminTransferScheduled)
	if err := _CrossGameReward.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardGamePoolCreatedIterator is returned from FilterGamePoolCreated and is used to iterate over the raw logs and unpacked data for GamePoolCreated events raised by the CrossGameReward contract.
type CrossGameRewardGamePoolCreatedIterator struct {
	Event *CrossGameRewardGamePoolCreated // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardGamePoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardGamePoolCreated)
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
		it.Event = new(CrossGameRewardGamePoolCreated)
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
func (it *CrossGameRewardGamePoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardGamePoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardGamePoolCreated represents a GamePoolCreated event raised by the CrossGameReward contract.
type CrossGameRewardGamePoolCreated struct {
	PoolId       *big.Int
	PoolAddress  common.Address
	DepositToken common.Address
	RewardToken  common.Address
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGamePoolCreated is a free log retrieval operation binding the contract event 0xcb6c4a9a1de3ee56c46859db21e5f245245aee3068e7b10d47bc646583cb816d.
//
// Solidity: event GamePoolCreated(uint256 indexed poolId, address indexed poolAddress, address depositToken, address indexed rewardToken, string name)
func (_CrossGameReward *CrossGameRewardFilterer) FilterGamePoolCreated(opts *bind.FilterOpts, poolId []*big.Int, poolAddress []common.Address, rewardToken []common.Address) (*CrossGameRewardGamePoolCreatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "GamePoolCreated", poolIdRule, poolAddressRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardGamePoolCreatedIterator{contract: _CrossGameReward.contract, event: "GamePoolCreated", logs: logs, sub: sub}, nil
}

// WatchGamePoolCreated is a free log subscription operation binding the contract event 0xcb6c4a9a1de3ee56c46859db21e5f245245aee3068e7b10d47bc646583cb816d.
//
// Solidity: event GamePoolCreated(uint256 indexed poolId, address indexed poolAddress, address depositToken, address indexed rewardToken, string name)
func (_CrossGameReward *CrossGameRewardFilterer) WatchGamePoolCreated(opts *bind.WatchOpts, sink chan<- *CrossGameRewardGamePoolCreated, poolId []*big.Int, poolAddress []common.Address, rewardToken []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "GamePoolCreated", poolIdRule, poolAddressRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardGamePoolCreated)
				if err := _CrossGameReward.contract.UnpackLog(event, "GamePoolCreated", log); err != nil {
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

// ParseGamePoolCreated is a log parse operation binding the contract event 0xcb6c4a9a1de3ee56c46859db21e5f245245aee3068e7b10d47bc646583cb816d.
//
// Solidity: event GamePoolCreated(uint256 indexed poolId, address indexed poolAddress, address depositToken, address indexed rewardToken, string name)
func (_CrossGameReward *CrossGameRewardFilterer) ParseGamePoolCreated(log types.Log) (*CrossGameRewardGamePoolCreated, error) {
	event := new(CrossGameRewardGamePoolCreated)
	if err := _CrossGameReward.contract.UnpackLog(event, "GamePoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardGamePoolImplementationSetIterator is returned from FilterGamePoolImplementationSet and is used to iterate over the raw logs and unpacked data for GamePoolImplementationSet events raised by the CrossGameReward contract.
type CrossGameRewardGamePoolImplementationSetIterator struct {
	Event *CrossGameRewardGamePoolImplementationSet // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardGamePoolImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardGamePoolImplementationSet)
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
		it.Event = new(CrossGameRewardGamePoolImplementationSet)
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
func (it *CrossGameRewardGamePoolImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardGamePoolImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardGamePoolImplementationSet represents a GamePoolImplementationSet event raised by the CrossGameReward contract.
type CrossGameRewardGamePoolImplementationSet struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGamePoolImplementationSet is a free log retrieval operation binding the contract event 0xeb3eb9823880a12aa7872f3d387d85ca1508fcf13f0311e53851853bb044e2a6.
//
// Solidity: event GamePoolImplementationSet(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) FilterGamePoolImplementationSet(opts *bind.FilterOpts, implementation []common.Address) (*CrossGameRewardGamePoolImplementationSetIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "GamePoolImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardGamePoolImplementationSetIterator{contract: _CrossGameReward.contract, event: "GamePoolImplementationSet", logs: logs, sub: sub}, nil
}

// WatchGamePoolImplementationSet is a free log subscription operation binding the contract event 0xeb3eb9823880a12aa7872f3d387d85ca1508fcf13f0311e53851853bb044e2a6.
//
// Solidity: event GamePoolImplementationSet(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) WatchGamePoolImplementationSet(opts *bind.WatchOpts, sink chan<- *CrossGameRewardGamePoolImplementationSet, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "GamePoolImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardGamePoolImplementationSet)
				if err := _CrossGameReward.contract.UnpackLog(event, "GamePoolImplementationSet", log); err != nil {
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

// ParseGamePoolImplementationSet is a log parse operation binding the contract event 0xeb3eb9823880a12aa7872f3d387d85ca1508fcf13f0311e53851853bb044e2a6.
//
// Solidity: event GamePoolImplementationSet(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) ParseGamePoolImplementationSet(log types.Log) (*CrossGameRewardGamePoolImplementationSet, error) {
	event := new(CrossGameRewardGamePoolImplementationSet)
	if err := _CrossGameReward.contract.UnpackLog(event, "GamePoolImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossGameReward contract.
type CrossGameRewardInitializedIterator struct {
	Event *CrossGameRewardInitialized // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardInitialized)
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
		it.Event = new(CrossGameRewardInitialized)
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
func (it *CrossGameRewardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardInitialized represents a Initialized event raised by the CrossGameReward contract.
type CrossGameRewardInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossGameReward *CrossGameRewardFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossGameRewardInitializedIterator, error) {

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardInitializedIterator{contract: _CrossGameReward.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossGameReward *CrossGameRewardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossGameRewardInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardInitialized)
				if err := _CrossGameReward.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossGameReward *CrossGameRewardFilterer) ParseInitialized(log types.Log) (*CrossGameRewardInitialized, error) {
	event := new(CrossGameRewardInitialized)
	if err := _CrossGameReward.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the CrossGameReward contract.
type CrossGameRewardPoolCreatedIterator struct {
	Event *CrossGameRewardPoolCreated // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolCreated)
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
		it.Event = new(CrossGameRewardPoolCreated)
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
func (it *CrossGameRewardPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolCreated represents a PoolCreated event raised by the CrossGameReward contract.
type CrossGameRewardPoolCreated struct {
	PoolId       *big.Int
	PoolAddress  common.Address
	DepositToken common.Address
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x9b64517ebf8d0fab4c3ec4b04d596444ace7293ad69076fe2458bc12be9126e3.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed depositToken, string name)
func (_CrossGameReward *CrossGameRewardFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int, poolAddress []common.Address, depositToken []common.Address) (*CrossGameRewardPoolCreatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var depositTokenRule []interface{}
	for _, depositTokenItem := range depositToken {
		depositTokenRule = append(depositTokenRule, depositTokenItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "PoolCreated", poolIdRule, poolAddressRule, depositTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolCreatedIterator{contract: _CrossGameReward.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x9b64517ebf8d0fab4c3ec4b04d596444ace7293ad69076fe2458bc12be9126e3.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed depositToken, string name)
func (_CrossGameReward *CrossGameRewardFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolCreated, poolId []*big.Int, poolAddress []common.Address, depositToken []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var depositTokenRule []interface{}
	for _, depositTokenItem := range depositToken {
		depositTokenRule = append(depositTokenRule, depositTokenItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "PoolCreated", poolIdRule, poolAddressRule, depositTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolCreated)
				if err := _CrossGameReward.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x9b64517ebf8d0fab4c3ec4b04d596444ace7293ad69076fe2458bc12be9126e3.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed depositToken, string name)
func (_CrossGameReward *CrossGameRewardFilterer) ParsePoolCreated(log types.Log) (*CrossGameRewardPoolCreated, error) {
	event := new(CrossGameRewardPoolCreated)
	if err := _CrossGameReward.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolImplementationSetIterator is returned from FilterPoolImplementationSet and is used to iterate over the raw logs and unpacked data for PoolImplementationSet events raised by the CrossGameReward contract.
type CrossGameRewardPoolImplementationSetIterator struct {
	Event *CrossGameRewardPoolImplementationSet // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolImplementationSet)
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
		it.Event = new(CrossGameRewardPoolImplementationSet)
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
func (it *CrossGameRewardPoolImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolImplementationSet represents a PoolImplementationSet event raised by the CrossGameReward contract.
type CrossGameRewardPoolImplementationSet struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolImplementationSet is a free log retrieval operation binding the contract event 0xdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957.
//
// Solidity: event PoolImplementationSet(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) FilterPoolImplementationSet(opts *bind.FilterOpts, implementation []common.Address) (*CrossGameRewardPoolImplementationSetIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "PoolImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolImplementationSetIterator{contract: _CrossGameReward.contract, event: "PoolImplementationSet", logs: logs, sub: sub}, nil
}

// WatchPoolImplementationSet is a free log subscription operation binding the contract event 0xdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957.
//
// Solidity: event PoolImplementationSet(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) WatchPoolImplementationSet(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolImplementationSet, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "PoolImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolImplementationSet)
				if err := _CrossGameReward.contract.UnpackLog(event, "PoolImplementationSet", log); err != nil {
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

// ParsePoolImplementationSet is a log parse operation binding the contract event 0xdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957.
//
// Solidity: event PoolImplementationSet(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) ParsePoolImplementationSet(log types.Log) (*CrossGameRewardPoolImplementationSet, error) {
	event := new(CrossGameRewardPoolImplementationSet)
	if err := _CrossGameReward.contract.UnpackLog(event, "PoolImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolsBatchUpgradedIterator is returned from FilterPoolsBatchUpgraded and is used to iterate over the raw logs and unpacked data for PoolsBatchUpgraded events raised by the CrossGameReward contract.
type CrossGameRewardPoolsBatchUpgradedIterator struct {
	Event *CrossGameRewardPoolsBatchUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolsBatchUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolsBatchUpgraded)
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
		it.Event = new(CrossGameRewardPoolsBatchUpgraded)
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
func (it *CrossGameRewardPoolsBatchUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolsBatchUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolsBatchUpgraded represents a PoolsBatchUpgraded event raised by the CrossGameReward contract.
type CrossGameRewardPoolsBatchUpgraded struct {
	PoolType          uint8
	NewImplementation common.Address
	Count             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPoolsBatchUpgraded is a free log retrieval operation binding the contract event 0x1d4313ba8617c7bedafe9b8528dafcacb84a619f94dab23c92c00460dbbe3f66.
//
// Solidity: event PoolsBatchUpgraded(uint8 poolType, address indexed newImplementation, uint256 count)
func (_CrossGameReward *CrossGameRewardFilterer) FilterPoolsBatchUpgraded(opts *bind.FilterOpts, newImplementation []common.Address) (*CrossGameRewardPoolsBatchUpgradedIterator, error) {

	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "PoolsBatchUpgraded", newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolsBatchUpgradedIterator{contract: _CrossGameReward.contract, event: "PoolsBatchUpgraded", logs: logs, sub: sub}, nil
}

// WatchPoolsBatchUpgraded is a free log subscription operation binding the contract event 0x1d4313ba8617c7bedafe9b8528dafcacb84a619f94dab23c92c00460dbbe3f66.
//
// Solidity: event PoolsBatchUpgraded(uint8 poolType, address indexed newImplementation, uint256 count)
func (_CrossGameReward *CrossGameRewardFilterer) WatchPoolsBatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolsBatchUpgraded, newImplementation []common.Address) (event.Subscription, error) {

	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "PoolsBatchUpgraded", newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolsBatchUpgraded)
				if err := _CrossGameReward.contract.UnpackLog(event, "PoolsBatchUpgraded", log); err != nil {
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

// ParsePoolsBatchUpgraded is a log parse operation binding the contract event 0x1d4313ba8617c7bedafe9b8528dafcacb84a619f94dab23c92c00460dbbe3f66.
//
// Solidity: event PoolsBatchUpgraded(uint8 poolType, address indexed newImplementation, uint256 count)
func (_CrossGameReward *CrossGameRewardFilterer) ParsePoolsBatchUpgraded(log types.Log) (*CrossGameRewardPoolsBatchUpgraded, error) {
	event := new(CrossGameRewardPoolsBatchUpgraded)
	if err := _CrossGameReward.contract.UnpackLog(event, "PoolsBatchUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardReclaimedFromPoolIterator is returned from FilterReclaimedFromPool and is used to iterate over the raw logs and unpacked data for ReclaimedFromPool events raised by the CrossGameReward contract.
type CrossGameRewardReclaimedFromPoolIterator struct {
	Event *CrossGameRewardReclaimedFromPool // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardReclaimedFromPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardReclaimedFromPool)
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
		it.Event = new(CrossGameRewardReclaimedFromPool)
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
func (it *CrossGameRewardReclaimedFromPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardReclaimedFromPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardReclaimedFromPool represents a ReclaimedFromPool event raised by the CrossGameReward contract.
type CrossGameRewardReclaimedFromPool struct {
	PoolId *big.Int
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReclaimedFromPool is a free log retrieval operation binding the contract event 0xae55fdf2c7467a88ea571a46bc6ecd9b95b7997fa6fed1d1c7f1842b5d603389.
//
// Solidity: event ReclaimedFromPool(uint256 indexed poolId, address indexed token, address indexed to, uint256 amount)
func (_CrossGameReward *CrossGameRewardFilterer) FilterReclaimedFromPool(opts *bind.FilterOpts, poolId []*big.Int, token []common.Address, to []common.Address) (*CrossGameRewardReclaimedFromPoolIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "ReclaimedFromPool", poolIdRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardReclaimedFromPoolIterator{contract: _CrossGameReward.contract, event: "ReclaimedFromPool", logs: logs, sub: sub}, nil
}

// WatchReclaimedFromPool is a free log subscription operation binding the contract event 0xae55fdf2c7467a88ea571a46bc6ecd9b95b7997fa6fed1d1c7f1842b5d603389.
//
// Solidity: event ReclaimedFromPool(uint256 indexed poolId, address indexed token, address indexed to, uint256 amount)
func (_CrossGameReward *CrossGameRewardFilterer) WatchReclaimedFromPool(opts *bind.WatchOpts, sink chan<- *CrossGameRewardReclaimedFromPool, poolId []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "ReclaimedFromPool", poolIdRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardReclaimedFromPool)
				if err := _CrossGameReward.contract.UnpackLog(event, "ReclaimedFromPool", log); err != nil {
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

// ParseReclaimedFromPool is a log parse operation binding the contract event 0xae55fdf2c7467a88ea571a46bc6ecd9b95b7997fa6fed1d1c7f1842b5d603389.
//
// Solidity: event ReclaimedFromPool(uint256 indexed poolId, address indexed token, address indexed to, uint256 amount)
func (_CrossGameReward *CrossGameRewardFilterer) ParseReclaimedFromPool(log types.Log) (*CrossGameRewardReclaimedFromPool, error) {
	event := new(CrossGameRewardReclaimedFromPool)
	if err := _CrossGameReward.contract.UnpackLog(event, "ReclaimedFromPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossGameReward contract.
type CrossGameRewardRoleAdminChangedIterator struct {
	Event *CrossGameRewardRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRoleAdminChanged)
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
		it.Event = new(CrossGameRewardRoleAdminChanged)
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
func (it *CrossGameRewardRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRoleAdminChanged represents a RoleAdminChanged event raised by the CrossGameReward contract.
type CrossGameRewardRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossGameReward *CrossGameRewardFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossGameRewardRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRoleAdminChangedIterator{contract: _CrossGameReward.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossGameReward *CrossGameRewardFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRoleAdminChanged)
				if err := _CrossGameReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossGameReward *CrossGameRewardFilterer) ParseRoleAdminChanged(log types.Log) (*CrossGameRewardRoleAdminChanged, error) {
	event := new(CrossGameRewardRoleAdminChanged)
	if err := _CrossGameReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossGameReward contract.
type CrossGameRewardRoleGrantedIterator struct {
	Event *CrossGameRewardRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRoleGranted)
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
		it.Event = new(CrossGameRewardRoleGranted)
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
func (it *CrossGameRewardRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRoleGranted represents a RoleGranted event raised by the CrossGameReward contract.
type CrossGameRewardRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameReward *CrossGameRewardFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossGameRewardRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRoleGrantedIterator{contract: _CrossGameReward.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameReward *CrossGameRewardFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRoleGranted)
				if err := _CrossGameReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameReward *CrossGameRewardFilterer) ParseRoleGranted(log types.Log) (*CrossGameRewardRoleGranted, error) {
	event := new(CrossGameRewardRoleGranted)
	if err := _CrossGameReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossGameReward contract.
type CrossGameRewardRoleRevokedIterator struct {
	Event *CrossGameRewardRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRoleRevoked)
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
		it.Event = new(CrossGameRewardRoleRevoked)
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
func (it *CrossGameRewardRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRoleRevoked represents a RoleRevoked event raised by the CrossGameReward contract.
type CrossGameRewardRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameReward *CrossGameRewardFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossGameRewardRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRoleRevokedIterator{contract: _CrossGameReward.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameReward *CrossGameRewardFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRoleRevoked)
				if err := _CrossGameReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameReward *CrossGameRewardFilterer) ParseRoleRevoked(log types.Log) (*CrossGameRewardRoleRevoked, error) {
	event := new(CrossGameRewardRoleRevoked)
	if err := _CrossGameReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRouterSetIterator is returned from FilterRouterSet and is used to iterate over the raw logs and unpacked data for RouterSet events raised by the CrossGameReward contract.
type CrossGameRewardRouterSetIterator struct {
	Event *CrossGameRewardRouterSet // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRouterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRouterSet)
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
		it.Event = new(CrossGameRewardRouterSet)
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
func (it *CrossGameRewardRouterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRouterSet represents a RouterSet event raised by the CrossGameReward contract.
type CrossGameRewardRouterSet struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterSet is a free log retrieval operation binding the contract event 0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15.
//
// Solidity: event RouterSet(address indexed router)
func (_CrossGameReward *CrossGameRewardFilterer) FilterRouterSet(opts *bind.FilterOpts, router []common.Address) (*CrossGameRewardRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "RouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterSetIterator{contract: _CrossGameReward.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

// WatchRouterSet is a free log subscription operation binding the contract event 0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15.
//
// Solidity: event RouterSet(address indexed router)
func (_CrossGameReward *CrossGameRewardFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "RouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRouterSet)
				if err := _CrossGameReward.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

// ParseRouterSet is a log parse operation binding the contract event 0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15.
//
// Solidity: event RouterSet(address indexed router)
func (_CrossGameReward *CrossGameRewardFilterer) ParseRouterSet(log types.Log) (*CrossGameRewardRouterSet, error) {
	event := new(CrossGameRewardRouterSet)
	if err := _CrossGameReward.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardSponsorRoleGrantedIterator is returned from FilterSponsorRoleGranted and is used to iterate over the raw logs and unpacked data for SponsorRoleGranted events raised by the CrossGameReward contract.
type CrossGameRewardSponsorRoleGrantedIterator struct {
	Event *CrossGameRewardSponsorRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardSponsorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardSponsorRoleGranted)
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
		it.Event = new(CrossGameRewardSponsorRoleGranted)
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
func (it *CrossGameRewardSponsorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardSponsorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardSponsorRoleGranted represents a SponsorRoleGranted event raised by the CrossGameReward contract.
type CrossGameRewardSponsorRoleGranted struct {
	PoolId  *big.Int
	Sponsor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSponsorRoleGranted is a free log retrieval operation binding the contract event 0xb341de1bcd4424e64431d748873b008dd03d26d1d6e059f321f3555650617a9b.
//
// Solidity: event SponsorRoleGranted(uint256 indexed poolId, address indexed sponsor)
func (_CrossGameReward *CrossGameRewardFilterer) FilterSponsorRoleGranted(opts *bind.FilterOpts, poolId []*big.Int, sponsor []common.Address) (*CrossGameRewardSponsorRoleGrantedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "SponsorRoleGranted", poolIdRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardSponsorRoleGrantedIterator{contract: _CrossGameReward.contract, event: "SponsorRoleGranted", logs: logs, sub: sub}, nil
}

// WatchSponsorRoleGranted is a free log subscription operation binding the contract event 0xb341de1bcd4424e64431d748873b008dd03d26d1d6e059f321f3555650617a9b.
//
// Solidity: event SponsorRoleGranted(uint256 indexed poolId, address indexed sponsor)
func (_CrossGameReward *CrossGameRewardFilterer) WatchSponsorRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossGameRewardSponsorRoleGranted, poolId []*big.Int, sponsor []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "SponsorRoleGranted", poolIdRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardSponsorRoleGranted)
				if err := _CrossGameReward.contract.UnpackLog(event, "SponsorRoleGranted", log); err != nil {
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

// ParseSponsorRoleGranted is a log parse operation binding the contract event 0xb341de1bcd4424e64431d748873b008dd03d26d1d6e059f321f3555650617a9b.
//
// Solidity: event SponsorRoleGranted(uint256 indexed poolId, address indexed sponsor)
func (_CrossGameReward *CrossGameRewardFilterer) ParseSponsorRoleGranted(log types.Log) (*CrossGameRewardSponsorRoleGranted, error) {
	event := new(CrossGameRewardSponsorRoleGranted)
	if err := _CrossGameReward.contract.UnpackLog(event, "SponsorRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardSponsorRoleRevokedIterator is returned from FilterSponsorRoleRevoked and is used to iterate over the raw logs and unpacked data for SponsorRoleRevoked events raised by the CrossGameReward contract.
type CrossGameRewardSponsorRoleRevokedIterator struct {
	Event *CrossGameRewardSponsorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardSponsorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardSponsorRoleRevoked)
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
		it.Event = new(CrossGameRewardSponsorRoleRevoked)
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
func (it *CrossGameRewardSponsorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardSponsorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardSponsorRoleRevoked represents a SponsorRoleRevoked event raised by the CrossGameReward contract.
type CrossGameRewardSponsorRoleRevoked struct {
	PoolId  *big.Int
	Sponsor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSponsorRoleRevoked is a free log retrieval operation binding the contract event 0x5d4cec63c165d8b6c2594e1506a0f45239d7ce890579740653db448bf4f8fccf.
//
// Solidity: event SponsorRoleRevoked(uint256 indexed poolId, address indexed sponsor)
func (_CrossGameReward *CrossGameRewardFilterer) FilterSponsorRoleRevoked(opts *bind.FilterOpts, poolId []*big.Int, sponsor []common.Address) (*CrossGameRewardSponsorRoleRevokedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "SponsorRoleRevoked", poolIdRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardSponsorRoleRevokedIterator{contract: _CrossGameReward.contract, event: "SponsorRoleRevoked", logs: logs, sub: sub}, nil
}

// WatchSponsorRoleRevoked is a free log subscription operation binding the contract event 0x5d4cec63c165d8b6c2594e1506a0f45239d7ce890579740653db448bf4f8fccf.
//
// Solidity: event SponsorRoleRevoked(uint256 indexed poolId, address indexed sponsor)
func (_CrossGameReward *CrossGameRewardFilterer) WatchSponsorRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossGameRewardSponsorRoleRevoked, poolId []*big.Int, sponsor []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "SponsorRoleRevoked", poolIdRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardSponsorRoleRevoked)
				if err := _CrossGameReward.contract.UnpackLog(event, "SponsorRoleRevoked", log); err != nil {
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

// ParseSponsorRoleRevoked is a log parse operation binding the contract event 0x5d4cec63c165d8b6c2594e1506a0f45239d7ce890579740653db448bf4f8fccf.
//
// Solidity: event SponsorRoleRevoked(uint256 indexed poolId, address indexed sponsor)
func (_CrossGameReward *CrossGameRewardFilterer) ParseSponsorRoleRevoked(log types.Log) (*CrossGameRewardSponsorRoleRevoked, error) {
	event := new(CrossGameRewardSponsorRoleRevoked)
	if err := _CrossGameReward.contract.UnpackLog(event, "SponsorRoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossGameReward contract.
type CrossGameRewardUpgradedIterator struct {
	Event *CrossGameRewardUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardUpgraded)
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
		it.Event = new(CrossGameRewardUpgraded)
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
func (it *CrossGameRewardUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardUpgraded represents a Upgraded event raised by the CrossGameReward contract.
type CrossGameRewardUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossGameRewardUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardUpgradedIterator{contract: _CrossGameReward.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossGameReward *CrossGameRewardFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossGameRewardUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameReward.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardUpgraded)
				if err := _CrossGameReward.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossGameReward *CrossGameRewardFilterer) ParseUpgraded(log types.Log) (*CrossGameRewardUpgraded, error) {
	event := new(CrossGameRewardUpgraded)
	if err := _CrossGameReward.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
