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
	DepositToken common.Address
	CreatedAt    *big.Int
}

// CrossGameRewardMetaData contains all meta data concerning the CrossGameReward contract.
var CrossGameRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActivePoolIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPoolIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolAddress\",\"outputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"}],\"name\":\"getPoolCountByDepositToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"}],\"name\":\"getPoolIdsByDepositToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structICrossGameReward.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"_poolImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolByDepositTokenAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolImplementation\",\"outputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reclaimFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setPoolImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcross\",\"outputs\":[{\"internalType\":\"contractIWCROSS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossGameRewardPool\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"PoolImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReclaimedFromPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRCSCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRPoolNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"ec87621c": "MANAGER_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"a1635945": "addRewardToken(uint256,address)",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"12d36171": "createPool(address,uint256)",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"fe96e4ff": "getActivePoolIds()",
		"f19c3d5b": "getAllPoolIds()",
		"00a5ae21": "getPoolAddress(uint256)",
		"d4148bcd": "getPoolCountByDepositToken(address)",
		"caa9a08d": "getPoolId(address)",
		"eeea4a79": "getPoolIdsByDepositToken(address)",
		"2f380b35": "getPoolInfo(uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"e7590268": "getTotalPoolCount()",
		"2f2ff15d": "grantRole(bytes32,address)",
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
		"ac4afa38": "pools(uint256)",
		"52d1902d": "proxiableUUID()",
		"c24140b2": "reclaimFromPool(uint256,address,address)",
		"35cc9cb4": "removeRewardToken(uint256,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"f887ea40": "router()",
		"d6f74898": "setPoolImplementation(address)",
		"b34c972e": "setPoolStatus(uint256,uint8)",
		"c0d78655": "setRouter(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"61616c46": "updateMinDepositAmount(uint256,uint256)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"a2db4582": "wcross()",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614a0f6100f95f395f818161210d0152818161213601526122e80152614a0f5ff3fe608060405260043610610302575f3560e01c8063a2db458211610191578063cf6eefb7116100dc578063e759026811610087578063f19c3d5b11610062578063f19c3d5b14610a5a578063f887ea4014610a6e578063fe96e4ff14610a8d575f5ffd5b8063e7590268146109e7578063ec87621c146109fb578063eeea4a7914610a2e575f5ffd5b8063d547741f116100b7578063d547741f14610995578063d602b9fd146109b4578063d6f74898146109c8575f5ffd5b8063cf6eefb7146108e1578063d4148bcd1461094b578063d4175be21461096a575f5ffd5b8063c24140b21161013c578063ce24af5311610117578063ce24af531461088f578063cefa7799146108ae578063cefc1429146108cd575f5ffd5b8063c24140b21461083d578063caa9a08d1461085c578063cc8463c81461087b575f5ffd5b8063b34c972e1161016c578063b34c972e146107e0578063b5be3221146107ff578063c0d786551461081e575f5ffd5b8063a2db4582146106ef578063ac4afa381461070e578063ad3cb1cc1461078b575f5ffd5b80634f1ef286116102515780638da5cb5b116101fc578063a1635945116101d7578063a16359451461068a578063a1eda53c146106a9578063a217fddf146106dc575f5ffd5b80638da5cb5b146105ff57806391cf6d3e1461061357806391d1485414610627575f5ffd5b8063634e93da1161022c578063634e93da14610585578063649a5ec7146105a457806384ef8ffc146105c3575f5ffd5b80634f1ef2861461053f57806352d1902d1461055257806361616c4614610566575f5ffd5b806318e56131116102b15780632f380b351161028c5780632f380b351461049957806335cc9cb41461050157806336568abe14610520575f5ffd5b806318e5613114610418578063248a9ca31461042d5780632f2ff15d1461047a575f5ffd5b80630aa6220b116102e15780630aa6220b1461039957806312d36171146103af578063155fff62146103eb575f5ffd5b8062a5ae211461030657806301ffc9a714610342578063022d63fb14610371575b5f5ffd5b348015610311575f5ffd5b506103256103203660046130b9565b610aa1565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561034d575f5ffd5b5061036161035c3660046130d0565b610b0f565b6040519015158152602001610339565b34801561037c575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610339565b3480156103a4575f5ffd5b506103ad610b6a565b005b3480156103ba575f5ffd5b506103ce6103c9366004613123565b610b7f565b604080519283526001600160a01b03909116602083015201610339565b3480156103f6575f5ffd5b5061040a6104053660046130b9565b610e11565b604051908152602001610339565b348015610423575f5ffd5b5061040a60045481565b348015610438575f5ffd5b5061040a6104473660046130b9565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b348015610485575f5ffd5b506103ad61049436600461314d565b610e1d565b3480156104a4575f5ffd5b506104b86104b33660046130b9565b610e62565b60405161033991905f608082019050825182526001600160a01b0360208401511660208301526001600160a01b0360408401511660408301526060830151606083015292915050565b34801561050c575f5ffd5b506103ad61051b36600461314d565b610f27565b34801561052b575f5ffd5b506103ad61053a36600461314d565b611029565b6103ad61054d3660046131a8565b61117b565b34801561055d575f5ffd5b5061040a611196565b348015610571575f5ffd5b506103ad61058036600461326f565b6111c4565b348015610590575f5ffd5b506103ad61059f36600461328f565b611297565b3480156105af575f5ffd5b506103ad6105be3660046132c4565b6112aa565b3480156105ce575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b0316610325565b34801561060a575f5ffd5b506103256112bd565b34801561061e575f5ffd5b5061040a5f5481565b348015610632575f5ffd5b5061036161064136600461314d565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610695575f5ffd5b506103ad6106a436600461314d565b6112f4565b3480156106b4575f5ffd5b506106bd6113c8565b6040805165ffffffffffff938416815292909116602083015201610339565b3480156106e7575f5ffd5b5061040a5f81565b3480156106fa575f5ffd5b50600154610325906001600160a01b031681565b348015610719575f5ffd5b5061075b6107283660046130b9565b60056020525f9081526040902080546001820154600283015460039093015491926001600160a01b039182169291169084565b60405161033994939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b348015610796575f5ffd5b506107d36040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b604051610339919061330b565b3480156107eb575f5ffd5b506103ad6107fa366004613329565b611485565b34801561080a575f5ffd5b5061040a610819366004613123565b611557565b348015610829575f5ffd5b506103ad61083836600461328f565b61157f565b348015610848575f5ffd5b506103ad61085736600461334c565b61162b565b348015610867575f5ffd5b5061040a61087636600461328f565b611819565b348015610886575f5ffd5b5061038261186a565b34801561089a575f5ffd5b506103ad6108a936600461338b565b61194a565b3480156108b9575f5ffd5b50600354610325906001600160a01b031681565b3480156108d8575f5ffd5b506103ad611c00565b3480156108ec575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610339565b348015610956575f5ffd5b5061040a61096536600461328f565b611c6e565b348015610975575f5ffd5b5061040a61098436600461328f565b60066020525f908152604090205481565b3480156109a0575f5ffd5b506103ad6109af36600461314d565b611c8e565b3480156109bf575f5ffd5b506103ad611ccf565b3480156109d3575f5ffd5b506103ad6109e236600461328f565b611ce1565b3480156109f2575f5ffd5b5061040a611d8d565b348015610a06575f5ffd5b5061040a7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610a39575f5ffd5b50610a4d610a4836600461328f565b611d98565b60405161033991906133cf565b348015610a65575f5ffd5b50610a4d611dbb565b348015610a79575f5ffd5b50600254610325906001600160a01b031681565b348015610a98575f5ffd5b50610a4d611dc7565b5f818152600560205260408120600101546001600160a01b0316610af1576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f908152600560205260409020600101546001600160a01b031690565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610b645750610b6482611fab565b92915050565b5f610b7481612041565b610b7c61204b565b50565b5f5f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610bab81612041565b6001600160a01b038516610beb576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8411610c24576040517f559ee9f200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60048054905f610c338361343e565b909155506040516001600160a01b0387166024820152604481018690529093505f9060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcd6dc6870000000000000000000000000000000000000000000000000000000017905260035490519192505f916001600160a01b03909116908390610cd19061309f565b610cdc929190613475565b604051809103905ff080158015610cf5573d5f5f3e3d5ffd5b50604080516080810182528781526001600160a01b0380841660208084018281528d841685870190815242606087019081525f8e8152600585528881209751885592516001880180549188167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055915160028801805491909716921691909117909455925160039094019390935581526006909152208690559350839050610da3600886612057565b506001600160a01b0387165f908152600760205260409020610dc59086612057565b50866001600160a01b0316846001600160a01b0316867f1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a660405160405180910390a45050509250929050565b5f610b64600883612062565b81610e54576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e5e828261206d565b5050565b604080516080810182525f8082526020820181905291810182905260608101919091525f828152600560205260409020600101546001600160a01b0316610ed5576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f9081526005602090815260409182902082516080810184528154815260018201546001600160a01b0390811693820193909352600282015490921692820192909252600390910154606082015290565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610f5181612041565b5f838152600560205260409020600101546001600160a01b0316610fa1576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f3d509c970000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690633d509c97906024015b5f604051808303815f87803b15801561100e575f5ffd5b505af1158015611020573d5f5f3e3d5ffd5b50505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561108457507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b1561116c577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806110ea575065ffffffffffff8116155b806110fd57504265ffffffffffff821610155b15611143576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b61117683836120b6565b505050565b611183612102565b61118c826121d2565b610e5e82826121dc565b5f61119f6122dd565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086111ee81612041565b5f838152600560205260409020600101546001600160a01b031661123e576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f84780205000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690638478020590602401610ff7565b5f6112a181612041565b610e5e8261233f565b5f6112b481612041565b610e5e826123b1565b5f6112ef7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b905090565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861131e81612041565b5f838152600560205260409020600101546001600160a01b031661136e576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f1c03e6cc0000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690631c03e6cc90602401610ff7565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400811580159061144a57504265ffffffffffff831610155b611455575f5f61147c565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086114af81612041565b5f838152600560205260409020600101546001600160a01b03166114ff576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f6d7c49a20000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636d7c49a290610ff79085906004016134c3565b6001600160a01b0382165f9081526007602052604081206115789083612062565b9392505050565b5f61158981612041565b6001600160a01b0382166115c9576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15905f90a25050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861165581612041565b5f848152600560205260409020600101546001600160a01b03166116a5576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f848152600560205260408082206001015490517f35c21d5d0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152909116906335c21d5d90602401602060405180830381865afa158015611714573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117389190613502565b5f86815260056020526040908190206001015490517f4d1cd0140000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301528681166024830152929350911690634d1cd014906044015f604051808303815f87803b1580156117ae575f5ffd5b505af11580156117c0573d5f5f3e3d5ffd5b50505050826001600160a01b0316846001600160a01b0316867fae55fdf2c7467a88ea571a46bc6ecd9b95b7997fa6fed1d1c7f1842b5d6033898460405161180a91815260200190565b60405180910390a45050505050565b6001600160a01b0381165f90815260066020526040812054808203610b64576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680158015906118ec57504265ffffffffffff8216105b61191d5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611943565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b5f611953612420565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561197f5750825b90505f8267ffffffffffffffff16600114801561199b5750303b155b9050811580156119a9575080155b156119e0576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611a415784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816611a81576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716611ac1576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611acb8688612448565b611ad361245a565b435f55600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a16179055604051611b15906130ac565b604051809103905ff080158015611b2e573d5f5f3e3d5ffd5b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091178155600455611b947f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0888612462565b508315611bf65784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114611c66576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161113a565b610b7c612543565b6001600160a01b0381165f908152600760205260408120610b6490612658565b81611cc5576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e5e8282612661565b5f611cd981612041565b610b7c6126a4565b5f611ceb81612041565b6001600160a01b038216611d2b576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957905f90a25050565b5f6112ef6008612658565b6001600160a01b0381165f908152600760205260409020606090610b64906126ae565b60606112ef60086126ae565b60605f611dd46008612658565b90505f8167ffffffffffffffff811115611df057611df061317b565b604051908082528060200260200182016040528015611e19578160200160208202803683370190505b5090505f805b83811015611f12575f611e33600883612062565b90505f5f828152600560209081526040918290206001015482517ff022869200000000000000000000000000000000000000000000000000000000815292516001600160a01b039091169263f02286929260048083019391928290030181865afa158015611ea3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ec79190613519565b6002811115611ed857611ed8613496565b03611f095780848481518110611ef057611ef0613534565b602090810291909101015282611f058161343e565b9350505b50600101611e1f565b505f8167ffffffffffffffff811115611f2d57611f2d61317b565b604051908082528060200260200182016040528015611f56578160200160208202803683370190505b5090505f5b82811015611fa257838181518110611f7557611f75613534565b6020026020010151828281518110611f8f57611f8f613534565b6020908102919091010152600101611f5b565b50949350505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610b6457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610b64565b610b7c81336126ba565b6120555f5f612746565b565b5f61157883836128d1565b5f611578838361291d565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120a681612041565b6120b08383612462565b50505050565b6001600160a01b03811633146120f8576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111768282612943565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061219b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661218f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610e5e81612041565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612236575060408051601f3d908101601f1916820190925261223391810190613502565b60015b612277576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161113a565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146122d3576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161113a565b61117683836129d9565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61234861186a565b61235142612a2e565b61235b9190613561565b90506123678282612a7d565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6123bb82612b2a565b6123c442612a2e565b6123ce9190613561565b90506123da8282612746565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610b64565b612450612b71565b610e5e8282612baf565b612055612b71565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083612531575f6124bb7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b0316146124fb576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b61253b8484612c6b565b949350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff168015806125a657504265ffffffffffff821610155b156125e7576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161113a565b6126215f61261c7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b612943565b5061262c5f83612462565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f610b64825490565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461269a81612041565b6120b08383612943565b6120555f5f612a7d565b60605f61157883612d55565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610e5e576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161113a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612858574265ffffffffffff8216101561282f576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255612858565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f81815260018301602052604081205461291657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b64565b505f610b64565b5f825f01828154811061293257612932613534565b905f5260205f200154905092915050565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561299f57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b156129cf576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61253b8484612dae565b6129e282612e70565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612a26576111768282612f17565b610e5e612f89565b5f65ffffffffffff821115612a79576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161113a565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b0388161717845591041680156120b0576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f612b3461186a565b90508065ffffffffffff168365ffffffffffff1611612b5c57612b57838261357f565b611578565b61157865ffffffffffff841662069780612fc1565b612b79612fd0565b612055576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bb7612b71565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216612c1a576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f600482015260240161113a565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556120b05f83612462565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612d4c575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612d023390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610b64565b5f915050610b64565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612da257602002820191905f5260205f20905b815481526020019060010190808311612d8e575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612d4c575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610b64565b806001600160a01b03163b5f03612ebe576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161113a565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051612f33919061359d565b5f60405180830381855af49150503d805f8114612f6b576040519150601f19603f3d011682016040523d82523d5f602084013e612f70565b606091505b5091509150612f80858383612fee565b95945050505050565b3415612055576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828218828410028218611578565b5f612fd9612420565b5468010000000000000000900460ff16919050565b606082612ffe57612b578261305e565b815115801561301557506001600160a01b0384163b155b15613057576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161113a565b5092915050565b80511561306d57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dc806135b483390190565b61104a8061399083390190565b5f602082840312156130c9575f5ffd5b5035919050565b5f602082840312156130e0575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611578575f5ffd5b6001600160a01b0381168114610b7c575f5ffd5b5f5f60408385031215613134575f5ffd5b823561313f8161310f565b946020939093013593505050565b5f5f6040838503121561315e575f5ffd5b8235915060208301356131708161310f565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f604083850312156131b9575f5ffd5b82356131c48161310f565b9150602083013567ffffffffffffffff8111156131df575f5ffd5b8301601f810185136131ef575f5ffd5b803567ffffffffffffffff8111156132095761320961317b565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff821117156132395761323961317b565b604052818152828201602001871015613250575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f60408385031215613280575f5ffd5b50508035926020909101359150565b5f6020828403121561329f575f5ffd5b81356115788161310f565b803565ffffffffffff811681146132bf575f5ffd5b919050565b5f602082840312156132d4575f5ffd5b611578826132aa565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61157860208301846132dd565b60038110610b7c575f5ffd5b5f5f6040838503121561333a575f5ffd5b8235915060208301356131708161331d565b5f5f5f6060848603121561335e575f5ffd5b8335925060208401356133708161310f565b915060408401356133808161310f565b809150509250925092565b5f5f5f6060848603121561339d575f5ffd5b83356133a88161310f565b925060208401356133b88161310f565b91506133c6604085016132aa565b90509250925092565b602080825282518282018190525f918401906040840190835b818110156134065783518352602093840193909201916001016133e8565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361346e5761346e613411565b5060010190565b6001600160a01b0383168152604060208201525f61253b60408301846132dd565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208101600383106134fc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b5f60208284031215613512575f5ffd5b5051919050565b5f60208284031215613529575f5ffd5b81516115788161331d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b65ffffffffffff8181168382160190811115610b6457610b64613411565b65ffffffffffff8281168282160390811115610b6457610b64613411565b5f82518060208501845e5f92019182525091905056fe60806040526040516103dc3803806103dc8339810160408190526100229161023b565b61002c8282610033565b5050610320565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030a565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020e57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024c575f5ffd5b82516001600160a01b0381168114610262575f5ffd5b60208401519092506001600160401b0381111561027d575f5ffd5b8301601f8101851361028d575f5ffd5b80516001600160401b038111156102a6576102a6610227565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d4576102d4610227565b6040528181528282016020018710156102eb575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032c5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220927d2839b1d9f37b4bfae40308f9fa5bd1e749b8d8b98c249d9b67e246bf55ba64736f6c634300081c003360a060405234801561000f575f5ffd5b506040518060400160405280600d81526020016c577261707065642043524f535360981b815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100649190610115565b5060046100718282610115565b505033608052506101cf565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100a557607f821691505b6020821081036100c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561011057805f5260205f20601f840160051c810160208510156100ee5750805b601f840160051c820191505b8181101561010d575f81556001016100fa565b50505b505050565b81516001600160401b0381111561012e5761012e61007d565b6101428161013c8454610091565b846100c9565b6020601f821160018114610174575f831561015d5750848201515b5f19600385901b1c1916600184901b17845561010d565b5f84815260208120601f198516915b828110156101a35787850151825560209485019460019092019101610183565b50848210156101c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b608051610e556101f55f395f8181610258015281816102fa01526104d90152610e555ff3fe6080604052600436106100d1575f3560e01c8063313ce5671161007c578063a9059cbb11610057578063a9059cbb14610228578063bf3308de14610247578063d0e30db01461029f578063dd62ed3e146102a7575f5ffd5b8063313ce567146101b857806370a08231146101d357806395d89b4114610214575f5ffd5b8063205c2878116100ac578063205c28781461015b57806323b872dd1461017a5780632e1a7d4d14610199575f5ffd5b806306fdde03146100e4578063095ea7b31461010e57806318160ddd1461013d575f5ffd5b366100e0576100de6102f8565b005b5f5ffd5b3480156100ef575f5ffd5b506100f861042e565b6040516101059190610c2f565b60405180910390f35b348015610119575f5ffd5b5061012d610128366004610ca3565b6104be565b6040519015158152602001610105565b348015610148575f5ffd5b506002545b604051908152602001610105565b348015610166575f5ffd5b506100de610175366004610ca3565b6104d7565b348015610185575f5ffd5b5061012d610194366004610ccd565b6106ba565b3480156101a4575f5ffd5b506100de6101b3366004610d0b565b6106dd565b3480156101c3575f5ffd5b5060405160128152602001610105565b3480156101de575f5ffd5b5061014d6101ed366004610d22565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b34801561021f575f5ffd5b506100f86106ea565b348015610233575f5ffd5b5061012d610242366004610ca3565b6106f9565b348015610252575f5ffd5b5061027a7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b6100de6102f8565b3480156102b2575f5ffd5b5061014d6102c1366004610d44565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610361573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103859190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f3411610422576040517fd8df41ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61042c3334610706565b565b60606003805461043d90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461046990610d96565b80156104b45780601f1061048b576101008083540402835291602001916104b4565b820191905f5260205f20905b81548152906001019060200180831161049757829003601f168201915b5050505050905090565b5f336104cb818585610769565b60019150505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610540573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105649190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105c8576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610615576040517f653345a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61061f3382610776565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610675576040519150601f19603f3d011682016040523d82523d5f602084013e61067a565b606091505b50509050806106b5576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f336106c78582856107d0565b6106d285858561089e565b506001949350505050565b6106e733826104d7565b50565b60606004805461043d90610d96565b5f336104cb81858561089e565b73ffffffffffffffffffffffffffffffffffffffff821661075a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6107655f8383610943565b5050565b6106b58383836001610aea565b73ffffffffffffffffffffffffffffffffffffffff82166107c5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b610765825f83610943565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610898578181101561088a576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610751565b61089884848484035f610aea565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166108ed576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff821661093c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b6106b58383835b73ffffffffffffffffffffffffffffffffffffffff831661097a578060025f82825461096f9190610de7565b90915550610a2a9050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156109ff576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610751565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610a5357600280548290039055610a7e565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610add91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8416610b39576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8316610b88576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610898578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610c2191815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff811681146106e7575f5ffd5b5f5f60408385031215610cb4575f5ffd5b8235610cbf81610c82565b946020939093013593505050565b5f5f5f60608486031215610cdf575f5ffd5b8335610cea81610c82565b92506020840135610cfa81610c82565b929592945050506040919091013590565b5f60208284031215610d1b575f5ffd5b5035919050565b5f60208284031215610d32575f5ffd5b8135610d3d81610c82565b9392505050565b5f5f60408385031215610d55575f5ffd5b8235610d6081610c82565b91506020830135610d7081610c82565b809150509250929050565b5f60208284031215610d8b575f5ffd5b8151610d3d81610c82565b600181811c90821680610daa57607f821691505b602082108103610de1577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156104d1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212209ff99d1df7425b5a2d579ebe50cbad7ebcb40604a5fdfb17acf3da6fd6c753ef64736f6c634300081c0033a26469706673582212209705a97d7c3e1413a287142721cf430295ab52ef6d82c083d5b32b805c04ef3c64736f6c634300081c0033",
}

// CrossGameRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossGameRewardMetaData.ABI instead.
var CrossGameRewardABI = CrossGameRewardMetaData.ABI

// CrossGameRewardBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossGameRewardBinRuntime = "608060405260043610610302575f3560e01c8063a2db458211610191578063cf6eefb7116100dc578063e759026811610087578063f19c3d5b11610062578063f19c3d5b14610a5a578063f887ea4014610a6e578063fe96e4ff14610a8d575f5ffd5b8063e7590268146109e7578063ec87621c146109fb578063eeea4a7914610a2e575f5ffd5b8063d547741f116100b7578063d547741f14610995578063d602b9fd146109b4578063d6f74898146109c8575f5ffd5b8063cf6eefb7146108e1578063d4148bcd1461094b578063d4175be21461096a575f5ffd5b8063c24140b21161013c578063ce24af5311610117578063ce24af531461088f578063cefa7799146108ae578063cefc1429146108cd575f5ffd5b8063c24140b21461083d578063caa9a08d1461085c578063cc8463c81461087b575f5ffd5b8063b34c972e1161016c578063b34c972e146107e0578063b5be3221146107ff578063c0d786551461081e575f5ffd5b8063a2db4582146106ef578063ac4afa381461070e578063ad3cb1cc1461078b575f5ffd5b80634f1ef286116102515780638da5cb5b116101fc578063a1635945116101d7578063a16359451461068a578063a1eda53c146106a9578063a217fddf146106dc575f5ffd5b80638da5cb5b146105ff57806391cf6d3e1461061357806391d1485414610627575f5ffd5b8063634e93da1161022c578063634e93da14610585578063649a5ec7146105a457806384ef8ffc146105c3575f5ffd5b80634f1ef2861461053f57806352d1902d1461055257806361616c4614610566575f5ffd5b806318e56131116102b15780632f380b351161028c5780632f380b351461049957806335cc9cb41461050157806336568abe14610520575f5ffd5b806318e5613114610418578063248a9ca31461042d5780632f2ff15d1461047a575f5ffd5b80630aa6220b116102e15780630aa6220b1461039957806312d36171146103af578063155fff62146103eb575f5ffd5b8062a5ae211461030657806301ffc9a714610342578063022d63fb14610371575b5f5ffd5b348015610311575f5ffd5b506103256103203660046130b9565b610aa1565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561034d575f5ffd5b5061036161035c3660046130d0565b610b0f565b6040519015158152602001610339565b34801561037c575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610339565b3480156103a4575f5ffd5b506103ad610b6a565b005b3480156103ba575f5ffd5b506103ce6103c9366004613123565b610b7f565b604080519283526001600160a01b03909116602083015201610339565b3480156103f6575f5ffd5b5061040a6104053660046130b9565b610e11565b604051908152602001610339565b348015610423575f5ffd5b5061040a60045481565b348015610438575f5ffd5b5061040a6104473660046130b9565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b348015610485575f5ffd5b506103ad61049436600461314d565b610e1d565b3480156104a4575f5ffd5b506104b86104b33660046130b9565b610e62565b60405161033991905f608082019050825182526001600160a01b0360208401511660208301526001600160a01b0360408401511660408301526060830151606083015292915050565b34801561050c575f5ffd5b506103ad61051b36600461314d565b610f27565b34801561052b575f5ffd5b506103ad61053a36600461314d565b611029565b6103ad61054d3660046131a8565b61117b565b34801561055d575f5ffd5b5061040a611196565b348015610571575f5ffd5b506103ad61058036600461326f565b6111c4565b348015610590575f5ffd5b506103ad61059f36600461328f565b611297565b3480156105af575f5ffd5b506103ad6105be3660046132c4565b6112aa565b3480156105ce575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b0316610325565b34801561060a575f5ffd5b506103256112bd565b34801561061e575f5ffd5b5061040a5f5481565b348015610632575f5ffd5b5061036161064136600461314d565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610695575f5ffd5b506103ad6106a436600461314d565b6112f4565b3480156106b4575f5ffd5b506106bd6113c8565b6040805165ffffffffffff938416815292909116602083015201610339565b3480156106e7575f5ffd5b5061040a5f81565b3480156106fa575f5ffd5b50600154610325906001600160a01b031681565b348015610719575f5ffd5b5061075b6107283660046130b9565b60056020525f9081526040902080546001820154600283015460039093015491926001600160a01b039182169291169084565b60405161033994939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b348015610796575f5ffd5b506107d36040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b604051610339919061330b565b3480156107eb575f5ffd5b506103ad6107fa366004613329565b611485565b34801561080a575f5ffd5b5061040a610819366004613123565b611557565b348015610829575f5ffd5b506103ad61083836600461328f565b61157f565b348015610848575f5ffd5b506103ad61085736600461334c565b61162b565b348015610867575f5ffd5b5061040a61087636600461328f565b611819565b348015610886575f5ffd5b5061038261186a565b34801561089a575f5ffd5b506103ad6108a936600461338b565b61194a565b3480156108b9575f5ffd5b50600354610325906001600160a01b031681565b3480156108d8575f5ffd5b506103ad611c00565b3480156108ec575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610339565b348015610956575f5ffd5b5061040a61096536600461328f565b611c6e565b348015610975575f5ffd5b5061040a61098436600461328f565b60066020525f908152604090205481565b3480156109a0575f5ffd5b506103ad6109af36600461314d565b611c8e565b3480156109bf575f5ffd5b506103ad611ccf565b3480156109d3575f5ffd5b506103ad6109e236600461328f565b611ce1565b3480156109f2575f5ffd5b5061040a611d8d565b348015610a06575f5ffd5b5061040a7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610a39575f5ffd5b50610a4d610a4836600461328f565b611d98565b60405161033991906133cf565b348015610a65575f5ffd5b50610a4d611dbb565b348015610a79575f5ffd5b50600254610325906001600160a01b031681565b348015610a98575f5ffd5b50610a4d611dc7565b5f818152600560205260408120600101546001600160a01b0316610af1576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f908152600560205260409020600101546001600160a01b031690565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610b645750610b6482611fab565b92915050565b5f610b7481612041565b610b7c61204b565b50565b5f5f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610bab81612041565b6001600160a01b038516610beb576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8411610c24576040517f559ee9f200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60048054905f610c338361343e565b909155506040516001600160a01b0387166024820152604481018690529093505f9060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcd6dc6870000000000000000000000000000000000000000000000000000000017905260035490519192505f916001600160a01b03909116908390610cd19061309f565b610cdc929190613475565b604051809103905ff080158015610cf5573d5f5f3e3d5ffd5b50604080516080810182528781526001600160a01b0380841660208084018281528d841685870190815242606087019081525f8e8152600585528881209751885592516001880180549188167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055915160028801805491909716921691909117909455925160039094019390935581526006909152208690559350839050610da3600886612057565b506001600160a01b0387165f908152600760205260409020610dc59086612057565b50866001600160a01b0316846001600160a01b0316867f1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a660405160405180910390a45050509250929050565b5f610b64600883612062565b81610e54576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e5e828261206d565b5050565b604080516080810182525f8082526020820181905291810182905260608101919091525f828152600560205260409020600101546001600160a01b0316610ed5576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f9081526005602090815260409182902082516080810184528154815260018201546001600160a01b0390811693820193909352600282015490921692820192909252600390910154606082015290565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610f5181612041565b5f838152600560205260409020600101546001600160a01b0316610fa1576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f3d509c970000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690633d509c97906024015b5f604051808303815f87803b15801561100e575f5ffd5b505af1158015611020573d5f5f3e3d5ffd5b50505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561108457507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b1561116c577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806110ea575065ffffffffffff8116155b806110fd57504265ffffffffffff821610155b15611143576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b61117683836120b6565b505050565b611183612102565b61118c826121d2565b610e5e82826121dc565b5f61119f6122dd565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086111ee81612041565b5f838152600560205260409020600101546001600160a01b031661123e576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f84780205000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690638478020590602401610ff7565b5f6112a181612041565b610e5e8261233f565b5f6112b481612041565b610e5e826123b1565b5f6112ef7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b905090565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861131e81612041565b5f838152600560205260409020600101546001600160a01b031661136e576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f1c03e6cc0000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690631c03e6cc90602401610ff7565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400811580159061144a57504265ffffffffffff831610155b611455575f5f61147c565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086114af81612041565b5f838152600560205260409020600101546001600160a01b03166114ff576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f6d7c49a20000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636d7c49a290610ff79085906004016134c3565b6001600160a01b0382165f9081526007602052604081206115789083612062565b9392505050565b5f61158981612041565b6001600160a01b0382166115c9576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15905f90a25050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861165581612041565b5f848152600560205260409020600101546001600160a01b03166116a5576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f848152600560205260408082206001015490517f35c21d5d0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152909116906335c21d5d90602401602060405180830381865afa158015611714573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117389190613502565b5f86815260056020526040908190206001015490517f4d1cd0140000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301528681166024830152929350911690634d1cd014906044015f604051808303815f87803b1580156117ae575f5ffd5b505af11580156117c0573d5f5f3e3d5ffd5b50505050826001600160a01b0316846001600160a01b0316867fae55fdf2c7467a88ea571a46bc6ecd9b95b7997fa6fed1d1c7f1842b5d6033898460405161180a91815260200190565b60405180910390a45050505050565b6001600160a01b0381165f90815260066020526040812054808203610b64576040517fc7dfdd2100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680158015906118ec57504265ffffffffffff8216105b61191d5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611943565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b5f611953612420565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561197f5750825b90505f8267ffffffffffffffff16600114801561199b5750303b155b9050811580156119a9575080155b156119e0576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611a415784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816611a81576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716611ac1576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611acb8688612448565b611ad361245a565b435f55600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a16179055604051611b15906130ac565b604051809103905ff080158015611b2e573d5f5f3e3d5ffd5b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091178155600455611b947f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0888612462565b508315611bf65784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114611c66576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161113a565b610b7c612543565b6001600160a01b0381165f908152600760205260408120610b6490612658565b81611cc5576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e5e8282612661565b5f611cd981612041565b610b7c6126a4565b5f611ceb81612041565b6001600160a01b038216611d2b576040517f113a909800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957905f90a25050565b5f6112ef6008612658565b6001600160a01b0381165f908152600760205260409020606090610b64906126ae565b60606112ef60086126ae565b60605f611dd46008612658565b90505f8167ffffffffffffffff811115611df057611df061317b565b604051908082528060200260200182016040528015611e19578160200160208202803683370190505b5090505f805b83811015611f12575f611e33600883612062565b90505f5f828152600560209081526040918290206001015482517ff022869200000000000000000000000000000000000000000000000000000000815292516001600160a01b039091169263f02286929260048083019391928290030181865afa158015611ea3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ec79190613519565b6002811115611ed857611ed8613496565b03611f095780848481518110611ef057611ef0613534565b602090810291909101015282611f058161343e565b9350505b50600101611e1f565b505f8167ffffffffffffffff811115611f2d57611f2d61317b565b604051908082528060200260200182016040528015611f56578160200160208202803683370190505b5090505f5b82811015611fa257838181518110611f7557611f75613534565b6020026020010151828281518110611f8f57611f8f613534565b6020908102919091010152600101611f5b565b50949350505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610b6457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610b64565b610b7c81336126ba565b6120555f5f612746565b565b5f61157883836128d1565b5f611578838361291d565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120a681612041565b6120b08383612462565b50505050565b6001600160a01b03811633146120f8576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111768282612943565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061219b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661218f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610e5e81612041565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612236575060408051601f3d908101601f1916820190925261223391810190613502565b60015b612277576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161113a565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146122d3576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161113a565b61117683836129d9565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61234861186a565b61235142612a2e565b61235b9190613561565b90506123678282612a7d565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6123bb82612b2a565b6123c442612a2e565b6123ce9190613561565b90506123da8282612746565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610b64565b612450612b71565b610e5e8282612baf565b612055612b71565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840083612531575f6124bb7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b0316146124fb576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b61253b8484612c6b565b949350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff168015806125a657504265ffffffffffff821610155b156125e7576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161113a565b6126215f61261c7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b612943565b5061262c5f83612462565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f610b64825490565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461269a81612041565b6120b08383612943565b6120555f5f612a7d565b60605f61157883612d55565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610e5e576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161113a565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612858574265ffffffffffff8216101561282f576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255612858565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f81815260018301602052604081205461291657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b64565b505f610b64565b5f825f01828154811061293257612932613534565b905f5260205f200154905092915050565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561299f57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b156129cf576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61253b8484612dae565b6129e282612e70565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612a26576111768282612f17565b610e5e612f89565b5f65ffffffffffff821115612a79576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161113a565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b0388161717845591041680156120b0576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f612b3461186a565b90508065ffffffffffff168365ffffffffffff1611612b5c57612b57838261357f565b611578565b61157865ffffffffffff841662069780612fc1565b612b79612fd0565b612055576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bb7612b71565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216612c1a576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f600482015260240161113a565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556120b05f83612462565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612d4c575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612d023390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610b64565b5f915050610b64565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612da257602002820191905f5260205f20905b815481526020019060010190808311612d8e575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612d4c575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610b64565b806001600160a01b03163b5f03612ebe576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161113a565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051612f33919061359d565b5f60405180830381855af49150503d805f8114612f6b576040519150601f19603f3d011682016040523d82523d5f602084013e612f70565b606091505b5091509150612f80858383612fee565b95945050505050565b3415612055576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828218828410028218611578565b5f612fd9612420565b5468010000000000000000900460ff16919050565b606082612ffe57612b578261305e565b815115801561301557506001600160a01b0384163b155b15613057576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161113a565b5092915050565b80511561306d57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dc806135b483390190565b61104a8061399083390190565b5f602082840312156130c9575f5ffd5b5035919050565b5f602082840312156130e0575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611578575f5ffd5b6001600160a01b0381168114610b7c575f5ffd5b5f5f60408385031215613134575f5ffd5b823561313f8161310f565b946020939093013593505050565b5f5f6040838503121561315e575f5ffd5b8235915060208301356131708161310f565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f604083850312156131b9575f5ffd5b82356131c48161310f565b9150602083013567ffffffffffffffff8111156131df575f5ffd5b8301601f810185136131ef575f5ffd5b803567ffffffffffffffff8111156132095761320961317b565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff821117156132395761323961317b565b604052818152828201602001871015613250575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f60408385031215613280575f5ffd5b50508035926020909101359150565b5f6020828403121561329f575f5ffd5b81356115788161310f565b803565ffffffffffff811681146132bf575f5ffd5b919050565b5f602082840312156132d4575f5ffd5b611578826132aa565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61157860208301846132dd565b60038110610b7c575f5ffd5b5f5f6040838503121561333a575f5ffd5b8235915060208301356131708161331d565b5f5f5f6060848603121561335e575f5ffd5b8335925060208401356133708161310f565b915060408401356133808161310f565b809150509250925092565b5f5f5f6060848603121561339d575f5ffd5b83356133a88161310f565b925060208401356133b88161310f565b91506133c6604085016132aa565b90509250925092565b602080825282518282018190525f918401906040840190835b818110156134065783518352602093840193909201916001016133e8565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361346e5761346e613411565b5060010190565b6001600160a01b0383168152604060208201525f61253b60408301846132dd565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208101600383106134fc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b5f60208284031215613512575f5ffd5b5051919050565b5f60208284031215613529575f5ffd5b81516115788161331d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b65ffffffffffff8181168382160190811115610b6457610b64613411565b65ffffffffffff8281168282160390811115610b6457610b64613411565b5f82518060208501845e5f92019182525091905056fe60806040526040516103dc3803806103dc8339810160408190526100229161023b565b61002c8282610033565b5050610320565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030a565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020e57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024c575f5ffd5b82516001600160a01b0381168114610262575f5ffd5b60208401519092506001600160401b0381111561027d575f5ffd5b8301601f8101851361028d575f5ffd5b80516001600160401b038111156102a6576102a6610227565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d4576102d4610227565b6040528181528282016020018710156102eb575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032c5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220927d2839b1d9f37b4bfae40308f9fa5bd1e749b8d8b98c249d9b67e246bf55ba64736f6c634300081c003360a060405234801561000f575f5ffd5b506040518060400160405280600d81526020016c577261707065642043524f535360981b815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100649190610115565b5060046100718282610115565b505033608052506101cf565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100a557607f821691505b6020821081036100c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561011057805f5260205f20601f840160051c810160208510156100ee5750805b601f840160051c820191505b8181101561010d575f81556001016100fa565b50505b505050565b81516001600160401b0381111561012e5761012e61007d565b6101428161013c8454610091565b846100c9565b6020601f821160018114610174575f831561015d5750848201515b5f19600385901b1c1916600184901b17845561010d565b5f84815260208120601f198516915b828110156101a35787850151825560209485019460019092019101610183565b50848210156101c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b608051610e556101f55f395f8181610258015281816102fa01526104d90152610e555ff3fe6080604052600436106100d1575f3560e01c8063313ce5671161007c578063a9059cbb11610057578063a9059cbb14610228578063bf3308de14610247578063d0e30db01461029f578063dd62ed3e146102a7575f5ffd5b8063313ce567146101b857806370a08231146101d357806395d89b4114610214575f5ffd5b8063205c2878116100ac578063205c28781461015b57806323b872dd1461017a5780632e1a7d4d14610199575f5ffd5b806306fdde03146100e4578063095ea7b31461010e57806318160ddd1461013d575f5ffd5b366100e0576100de6102f8565b005b5f5ffd5b3480156100ef575f5ffd5b506100f861042e565b6040516101059190610c2f565b60405180910390f35b348015610119575f5ffd5b5061012d610128366004610ca3565b6104be565b6040519015158152602001610105565b348015610148575f5ffd5b506002545b604051908152602001610105565b348015610166575f5ffd5b506100de610175366004610ca3565b6104d7565b348015610185575f5ffd5b5061012d610194366004610ccd565b6106ba565b3480156101a4575f5ffd5b506100de6101b3366004610d0b565b6106dd565b3480156101c3575f5ffd5b5060405160128152602001610105565b3480156101de575f5ffd5b5061014d6101ed366004610d22565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b34801561021f575f5ffd5b506100f86106ea565b348015610233575f5ffd5b5061012d610242366004610ca3565b6106f9565b348015610252575f5ffd5b5061027a7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b6100de6102f8565b3480156102b2575f5ffd5b5061014d6102c1366004610d44565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610361573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103859190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f3411610422576040517fd8df41ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61042c3334610706565b565b60606003805461043d90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461046990610d96565b80156104b45780601f1061048b576101008083540402835291602001916104b4565b820191905f5260205f20905b81548152906001019060200180831161049757829003601f168201915b5050505050905090565b5f336104cb818585610769565b60019150505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610540573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105649190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105c8576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610615576040517f653345a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61061f3382610776565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610675576040519150601f19603f3d011682016040523d82523d5f602084013e61067a565b606091505b50509050806106b5576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f336106c78582856107d0565b6106d285858561089e565b506001949350505050565b6106e733826104d7565b50565b60606004805461043d90610d96565b5f336104cb81858561089e565b73ffffffffffffffffffffffffffffffffffffffff821661075a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6107655f8383610943565b5050565b6106b58383836001610aea565b73ffffffffffffffffffffffffffffffffffffffff82166107c5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b610765825f83610943565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610898578181101561088a576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610751565b61089884848484035f610aea565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166108ed576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff821661093c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b6106b58383835b73ffffffffffffffffffffffffffffffffffffffff831661097a578060025f82825461096f9190610de7565b90915550610a2a9050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156109ff576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610751565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610a5357600280548290039055610a7e565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610add91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8416610b39576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8316610b88576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610898578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610c2191815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff811681146106e7575f5ffd5b5f5f60408385031215610cb4575f5ffd5b8235610cbf81610c82565b946020939093013593505050565b5f5f5f60608486031215610cdf575f5ffd5b8335610cea81610c82565b92506020840135610cfa81610c82565b929592945050506040919091013590565b5f60208284031215610d1b575f5ffd5b5035919050565b5f60208284031215610d32575f5ffd5b8135610d3d81610c82565b9392505050565b5f5f60408385031215610d55575f5ffd5b8235610d6081610c82565b91506020830135610d7081610c82565b809150509250929050565b5f60208284031215610d8b575f5ffd5b8151610d3d81610c82565b600181811c90821680610daa57607f821691505b602082108103610de1577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156104d1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212209ff99d1df7425b5a2d579ebe50cbad7ebcb40604a5fdfb17acf3da6fd6c753ef64736f6c634300081c0033a26469706673582212209705a97d7c3e1413a287142721cf430295ab52ef6d82c083d5b32b805c04ef3c64736f6c634300081c0033"

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

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,address,uint256))
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
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,address,uint256))
func (_CrossGameReward *CrossGameRewardSession) GetPoolInfo(poolId *big.Int) (ICrossGameRewardPoolInfo, error) {
	return _CrossGameReward.Contract.GetPoolInfo(&_CrossGameReward.CallOpts, poolId)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,address,uint256))
func (_CrossGameReward *CrossGameRewardCallerSession) GetPoolInfo(poolId *big.Int) (ICrossGameRewardPoolInfo, error) {
	return _CrossGameReward.Contract.GetPoolInfo(&_CrossGameReward.CallOpts, poolId)
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

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, address depositToken, uint256 createdAt)
func (_CrossGameReward *CrossGameRewardCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	DepositToken common.Address
	CreatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameReward.contract.Call(opts, &out, "pools", arg0)

	outstruct := new(struct {
		PoolId       *big.Int
		Pool         common.Address
		DepositToken common.Address
		CreatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DepositToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, address depositToken, uint256 createdAt)
func (_CrossGameReward *CrossGameRewardSession) Pools(arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	DepositToken common.Address
	CreatedAt    *big.Int
}, error) {
	return _CrossGameReward.Contract.Pools(&_CrossGameReward.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, address depositToken, uint256 createdAt)
func (_CrossGameReward *CrossGameRewardCallerSession) Pools(arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
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

// CreatePool is a paid mutator transaction binding the contract method 0x12d36171.
//
// Solidity: function createPool(address depositToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardTransactor) CreatePool(opts *bind.TransactOpts, depositToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.contract.Transact(opts, "createPool", depositToken, minDepositAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x12d36171.
//
// Solidity: function createPool(address depositToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardSession) CreatePool(depositToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CreatePool(&_CrossGameReward.TransactOpts, depositToken, minDepositAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x12d36171.
//
// Solidity: function createPool(address depositToken, uint256 minDepositAmount) returns(uint256 poolId, address pool)
func (_CrossGameReward *CrossGameRewardTransactorSession) CreatePool(depositToken common.Address, minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameReward.Contract.CreatePool(&_CrossGameReward.TransactOpts, depositToken, minDepositAmount)
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
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a6.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed depositToken)
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

// WatchPoolCreated is a free log subscription operation binding the contract event 0x1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a6.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed depositToken)
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

// ParsePoolCreated is a log parse operation binding the contract event 0x1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a6.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed depositToken)
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
