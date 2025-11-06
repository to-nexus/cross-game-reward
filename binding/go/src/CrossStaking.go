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

// ICrossStakingPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type ICrossStakingPoolInfo struct {
	PoolId       *big.Int
	Pool         common.Address
	StakingToken common.Address
	CreatedAt    *big.Int
}

// CrossStakingMetaData contains all meta data concerning the CrossStaking contract.
var CrossStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minStakeAmount\",\"type\":\"uint256\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossStakingPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActivePoolIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPoolIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolAddress\",\"outputs\":[{\"internalType\":\"contractICrossStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"getPoolCountByStakingToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossStakingPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"getPoolIdsByStakingToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossStakingPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structICrossStaking.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossStakingPool\",\"name\":\"_poolImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"_initialDelay\",\"type\":\"uint48\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolByStakingTokenAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolImplementation\",\"outputs\":[{\"internalType\":\"contractICrossStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractICrossStakingPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossStakingPool\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setPoolImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"enumICrossStakingPool.PoolStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcross\",\"outputs\":[{\"internalType\":\"contractIWCROSS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICrossStakingPool\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"PoolImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFromPool\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSPoolNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
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
		"0f538b40": "getPoolCountByStakingToken(address)",
		"caa9a08d": "getPoolId(address)",
		"eba32d83": "getPoolIdsByStakingToken(address)",
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
		"3b95352b": "poolByStakingTokenAt(address,uint256)",
		"d4175be2": "poolIds(address)",
		"cefa7799": "poolImplementation()",
		"ac4afa38": "pools(uint256)",
		"52d1902d": "proxiableUUID()",
		"35cc9cb4": "removeRewardToken(uint256,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"f887ea40": "router()",
		"d6f74898": "setPoolImplementation(address)",
		"b34c972e": "setPoolStatus(uint256,uint8)",
		"c0d78655": "setRouter(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"1538af09": "updateMinStakeAmount(uint256,uint256)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"a2db4582": "wcross()",
		"b9c17c4d": "withdrawFromPool(uint256,address,address)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516149b26100f95f395f81816121160152818161213f01526122f101526149b25ff3fe608060405260043610610302575f3560e01c8063a1eda53c11610191578063cefc1429116100dc578063e759026811610087578063f19c3d5b11610062578063f19c3d5b14610a5a578063f887ea4014610a6e578063fe96e4ff14610a8d575f5ffd5b8063e7590268146109e7578063eba32d83146109fb578063ec87621c14610a27575f5ffd5b8063d547741f116100b7578063d547741f14610995578063d602b9fd146109b4578063d6f74898146109c8575f5ffd5b8063cefc1429146108ec578063cf6eefb714610900578063d4175be21461096a575f5ffd5b8063b9c17c4d1161013c578063cc8463c811610117578063cc8463c81461089a578063ce24af53146108ae578063cefa7799146108cd575f5ffd5b8063b9c17c4d1461083d578063c0d786551461085c578063caa9a08d1461087b575f5ffd5b8063ac4afa381161016c578063ac4afa381461074c578063ad3cb1cc146107c9578063b34c972e1461081e575f5ffd5b8063a1eda53c146106e7578063a217fddf1461071a578063a2db45821461072d575f5ffd5b806335cc9cb411610251578063649a5ec7116101fc57806391cf6d3e116101d757806391cf6d3e1461065157806391d1485414610665578063a1635945146106c8575f5ffd5b8063649a5ec7146105e257806384ef8ffc146106015780638da5cb5b1461063d575f5ffd5b80634f1ef2861161022c5780634f1ef2861461059c57806352d1902d146105af578063634e93da146105c3575f5ffd5b806335cc9cb41461053f57806336568abe1461055e5780633b95352b1461057d575f5ffd5b80631538af09116102b1578063248a9ca31161028c578063248a9ca31461046b5780632f2ff15d146104b85780632f380b35146104d7575f5ffd5b80631538af0914610418578063155fff621461043757806318e5613114610456575f5ffd5b80630aa6220b116102e15780630aa6220b146103995780630f538b40146103af57806312d36171146103dc575f5ffd5b8062a5ae211461030657806301ffc9a714610342578063022d63fb14610371575b5f5ffd5b348015610311575f5ffd5b506103256103203660046130b9565b610aa1565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561034d575f5ffd5b5061036161035c3660046130d0565b610b0f565b6040519015158152602001610339565b34801561037c575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610339565b3480156103a4575f5ffd5b506103ad610b6a565b005b3480156103ba575f5ffd5b506103ce6103c9366004613123565b610b7f565b604051908152602001610339565b3480156103e7575f5ffd5b506103fb6103f636600461313e565b610b9f565b604080519283526001600160a01b03909116602083015201610339565b348015610423575f5ffd5b506103ad610432366004613168565b610e31565b348015610442575f5ffd5b506103ce6104513660046130b9565b610f32565b348015610461575f5ffd5b506103ce60045481565b348015610476575f5ffd5b506103ce6104853660046130b9565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156104c3575f5ffd5b506103ad6104d2366004613188565b610f3e565b3480156104e2575f5ffd5b506104f66104f13660046130b9565b610f83565b60405161033991905f608082019050825182526001600160a01b0360208401511660208301526001600160a01b0360408401511660408301526060830151606083015292915050565b34801561054a575f5ffd5b506103ad610559366004613188565b611048565b348015610569575f5ffd5b506103ad610578366004613188565b61111c565b348015610588575f5ffd5b506103ce61059736600461313e565b61126e565b6103ad6105aa3660046131e3565b611296565b3480156105ba575f5ffd5b506103ce6112b1565b3480156105ce575f5ffd5b506103ad6105dd366004613123565b6112df565b3480156105ed575f5ffd5b506103ad6105fc3660046132c4565b6112f2565b34801561060c575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b0316610325565b348015610648575f5ffd5b50610325611305565b34801561065c575f5ffd5b506103ce5f5481565b348015610670575f5ffd5b5061036161067f366004613188565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156106d3575f5ffd5b506103ad6106e2366004613188565b61133c565b3480156106f2575f5ffd5b506106fb611410565b6040805165ffffffffffff938416815292909116602083015201610339565b348015610725575f5ffd5b506103ce5f81565b348015610738575f5ffd5b50600154610325906001600160a01b031681565b348015610757575f5ffd5b506107996107663660046130b9565b60056020525f9081526040902080546001820154600283015460039093015491926001600160a01b039182169291169084565b60405161033994939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b3480156107d4575f5ffd5b506108116040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b604051610339919061330b565b348015610829575f5ffd5b506103ad610838366004613329565b6114cd565b348015610848575f5ffd5b506103ad61085736600461334c565b61159f565b348015610867575f5ffd5b506103ad610876366004613123565b61178d565b348015610886575f5ffd5b506103ce610895366004613123565b611839565b3480156108a5575f5ffd5b5061038261188a565b3480156108b9575f5ffd5b506103ad6108c836600461338b565b61196a565b3480156108d8575f5ffd5b50600354610325906001600160a01b031681565b3480156108f7575f5ffd5b506103ad611c20565b34801561090b575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610339565b348015610975575f5ffd5b506103ce610984366004613123565b60066020525f908152604090205481565b3480156109a0575f5ffd5b506103ad6109af366004613188565b611c8e565b3480156109bf575f5ffd5b506103ad611ccf565b3480156109d3575f5ffd5b506103ad6109e2366004613123565b611ce1565b3480156109f2575f5ffd5b506103ce611d8d565b348015610a06575f5ffd5b50610a1a610a15366004613123565b611d98565b60405161033991906133cf565b348015610a32575f5ffd5b506103ce7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610a65575f5ffd5b50610a1a611dbb565b348015610a79575f5ffd5b50600254610325906001600160a01b031681565b348015610a98575f5ffd5b50610a1a611dc7565b5f818152600560205260408120600101546001600160a01b0316610af1576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f908152600560205260409020600101546001600160a01b031690565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610b645750610b6482611fab565b92915050565b5f610b7481612041565b610b7c61204b565b50565b6001600160a01b0381165f908152600760205260408120610b6490612057565b5f5f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610bcb81612041565b6001600160a01b038516610c0b576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8411610c44576040517f5e7c0cae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60048054905f610c538361343e565b909155506040516001600160a01b0387166024820152604481018690529093505f9060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcd6dc6870000000000000000000000000000000000000000000000000000000017905260035490519192505f916001600160a01b03909116908390610cf19061309f565b610cfc929190613475565b604051809103905ff080158015610d15573d5f5f3e3d5ffd5b50604080516080810182528781526001600160a01b0380841660208084018281528d841685870190815242606087019081525f8e8152600585528881209751885592516001880180549188167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055915160028801805491909716921691909117909455925160039094019390935581526006909152208690559350839050610dc3600886612060565b506001600160a01b0387165f908152600760205260409020610de59086612060565b50866001600160a01b0316846001600160a01b0316867f1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a660405160405180910390a45050509250929050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610e5b81612041565b5f838152600560205260409020600101546001600160a01b0316610eab576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517fd8060cd5000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b039091169063d8060cd5906024015b5f604051808303815f87803b158015610f17575f5ffd5b505af1158015610f29573d5f5f3e3d5ffd5b50505050505050565b5f610b6460088361206b565b81610f75576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7f8282612076565b5050565b604080516080810182525f8082526020820181905291810182905260608101919091525f828152600560205260409020600101546001600160a01b0316610ff6576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f9081526005602090815260409182902082516080810184528154815260018201546001600160a01b0390811693820193909352600282015490921692820192909252600390910154606082015290565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861107281612041565b5f838152600560205260409020600101546001600160a01b03166110c2576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f3d509c970000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690633d509c9790602401610f00565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561117757507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b1561125f577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806111dd575065ffffffffffff8116155b806111f057504265ffffffffffff821610155b15611236576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b61126983836120bf565b505050565b6001600160a01b0382165f90815260076020526040812061128f908361206b565b9392505050565b61129e61210b565b6112a7826121db565b610f7f82826121e5565b5f6112ba6122e6565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f6112e981612041565b610f7f82612348565b5f6112fc81612041565b610f7f826123ba565b5f6113377feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b905090565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861136681612041565b5f838152600560205260409020600101546001600160a01b03166113b6576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f1c03e6cc0000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690631c03e6cc90602401610f00565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400811580159061149257504265ffffffffffff831610155b61149d575f5f6114c4565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086114f781612041565b5f838152600560205260409020600101546001600160a01b0316611547576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f6d7c49a20000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636d7c49a290610f009085906004016134c3565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086115c981612041565b5f848152600560205260409020600101546001600160a01b0316611619576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f848152600560205260408082206001015490517ff67771750000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301529091169063f677717590602401602060405180830381865afa158015611688573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116ac9190613502565b5f86815260056020526040908190206001015490517ff940e3850000000000000000000000000000000000000000000000000000000081526001600160a01b038781166004830152868116602483015292935091169063f940e385906044015f604051808303815f87803b158015611722575f5ffd5b505af1158015611734573d5f5f3e3d5ffd5b50505050826001600160a01b0316846001600160a01b0316867f408708e2d99a8cfa59be1466864d74f6ddddc62b5ddc3bb11f1c3cc3ce9ed65f8460405161177e91815260200190565b60405180910390a45050505050565b5f61179781612041565b6001600160a01b0382166117d7576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15905f90a25050565b6001600160a01b0381165f90815260066020526040812054808203610b64576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801580159061190c57504265ffffffffffff8216105b61193d5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611963565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b5f611973612429565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561199f5750825b90505f8267ffffffffffffffff1660011480156119bb5750303b155b9050811580156119c9575080155b15611a00576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611a615784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816611aa1576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716611ae1576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611aeb8688612451565b611af3612463565b435f55600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a16179055604051611b35906130ac565b604051809103905ff080158015611b4e573d5f5f3e3d5ffd5b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091178155600455611bb47f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b088861246b565b508315611c165784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114611c86576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161122d565b610b7c61254c565b81611cc5576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7f8282612661565b5f611cd981612041565b610b7c6126a4565b5f611ceb81612041565b6001600160a01b038216611d2b576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957905f90a25050565b5f6113376008612057565b6001600160a01b0381165f908152600760205260409020606090610b64906126ae565b606061133760086126ae565b60605f611dd46008612057565b90505f8167ffffffffffffffff811115611df057611df06131b6565b604051908082528060200260200182016040528015611e19578160200160208202803683370190505b5090505f805b83811015611f12575f611e3360088361206b565b90505f5f828152600560209081526040918290206001015482517ff022869200000000000000000000000000000000000000000000000000000000815292516001600160a01b039091169263f02286929260048083019391928290030181865afa158015611ea3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ec79190613519565b6002811115611ed857611ed8613496565b03611f095780848481518110611ef057611ef0613534565b602090810291909101015282611f058161343e565b9350505b50600101611e1f565b505f8167ffffffffffffffff811115611f2d57611f2d6131b6565b604051908082528060200260200182016040528015611f56578160200160208202803683370190505b5090505f5b82811015611fa257838181518110611f7557611f75613534565b6020026020010151828281518110611f8f57611f8f613534565b6020908102919091010152600101611f5b565b50949350505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610b6457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610b64565b610b7c81336126ba565b6120555f5f612746565b565b5f610b64825490565b5f61128f83836128d1565b5f61128f838361291d565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120af81612041565b6120b9838361246b565b50505050565b6001600160a01b0381163314612101576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112698282612943565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806121a457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166121987f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610f7f81612041565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561223f575060408051601f3d908101601f1916820190925261223c91810190613502565b60015b612280576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161122d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146122dc576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161122d565b61126983836129d9565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61235161188a565b61235a42612a2e565b6123649190613561565b90506123708282612a7d565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6123c482612b2a565b6123cd42612a2e565b6123d79190613561565b90506123e38282612746565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610b64565b612459612b71565b610f7f8282612baf565b612055612b71565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008361253a575f6124c47feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b031614612504576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b6125448484612c6b565b949350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff168015806125af57504265ffffffffffff821610155b156125f0576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161122d565b61262a5f6126257feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b612943565b506126355f8361246b565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461269a81612041565b6120b98383612943565b6120555f5f612a7d565b60605f61128f83612d55565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f7f576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161122d565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612858574265ffffffffffff8216101561282f576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255612858565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f81815260018301602052604081205461291657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b64565b505f610b64565b5f825f01828154811061293257612932613534565b905f5260205f200154905092915050565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561299f57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b156129cf576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6125448484612dae565b6129e282612e70565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612a26576112698282612f17565b610f7f612f89565b5f65ffffffffffff821115612a79576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161122d565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b0388161717845591041680156120b9576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f612b3461188a565b90508065ffffffffffff168365ffffffffffff1611612b5c57612b57838261357f565b61128f565b61128f65ffffffffffff841662069780612fc1565b612b79612fd0565b612055576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bb7612b71565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216612c1a576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f600482015260240161122d565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556120b95f8361246b565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612d4c575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612d023390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610b64565b5f915050610b64565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612da257602002820191905f5260205f20905b815481526020019060010190808311612d8e575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612d4c575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610b64565b806001600160a01b03163b5f03612ebe576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161122d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051612f33919061359d565b5f60405180830381855af49150503d805f8114612f6b576040519150601f19603f3d011682016040523d82523d5f602084013e612f70565b606091505b5091509150612f80858383612fee565b95945050505050565b3415612055576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82821882841002821861128f565b5f612fd9612429565b5468010000000000000000900460ff16919050565b606082612ffe57612b578261305e565b815115801561301557506001600160a01b0384163b155b15613057576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161122d565b5092915050565b80511561306d57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dc806135b483390190565b610fed8061399083390190565b5f602082840312156130c9575f5ffd5b5035919050565b5f602082840312156130e0575f5ffd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461128f575f5ffd5b6001600160a01b0381168114610b7c575f5ffd5b5f60208284031215613133575f5ffd5b813561128f8161310f565b5f5f6040838503121561314f575f5ffd5b823561315a8161310f565b946020939093013593505050565b5f5f60408385031215613179575f5ffd5b50508035926020909101359150565b5f5f60408385031215613199575f5ffd5b8235915060208301356131ab8161310f565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f604083850312156131f4575f5ffd5b82356131ff8161310f565b9150602083013567ffffffffffffffff81111561321a575f5ffd5b8301601f8101851361322a575f5ffd5b803567ffffffffffffffff811115613244576132446131b6565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715613274576132746131b6565b60405281815282820160200187101561328b575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b803565ffffffffffff811681146132bf575f5ffd5b919050565b5f602082840312156132d4575f5ffd5b61128f826132aa565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61128f60208301846132dd565b60038110610b7c575f5ffd5b5f5f6040838503121561333a575f5ffd5b8235915060208301356131ab8161331d565b5f5f5f6060848603121561335e575f5ffd5b8335925060208401356133708161310f565b915060408401356133808161310f565b809150509250925092565b5f5f5f6060848603121561339d575f5ffd5b83356133a88161310f565b925060208401356133b88161310f565b91506133c6604085016132aa565b90509250925092565b602080825282518282018190525f918401906040840190835b818110156134065783518352602093840193909201916001016133e8565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361346e5761346e613411565b5060010190565b6001600160a01b0383168152604060208201525f61254460408301846132dd565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208101600383106134fc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b5f60208284031215613512575f5ffd5b5051919050565b5f60208284031215613529575f5ffd5b815161128f8161331d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b65ffffffffffff8181168382160190811115610b6457610b64613411565b65ffffffffffff8281168282160390811115610b6457610b64613411565b5f82518060208501845e5f92019182525091905056fe60806040526040516103dc3803806103dc8339810160408190526100229161023b565b61002c8282610033565b5050610320565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030a565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020e57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024c575f5ffd5b82516001600160a01b0381168114610262575f5ffd5b60208401519092506001600160401b0381111561027d575f5ffd5b8301601f8101851361028d575f5ffd5b80516001600160401b038111156102a6576102a6610227565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d4576102d4610227565b6040528181528282016020018710156102eb575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032c5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220927d2839b1d9f37b4bfae40308f9fa5bd1e749b8d8b98c249d9b67e246bf55ba64736f6c634300081c0033608060405234801561000f575f5ffd5b506040518060400160405280600d81526020016c577261707065642043524f535360981b815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100649190610123565b5060046100718282610123565b5050600580546001600160a01b03191633179055506101dd565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100b357607f821691505b6020821081036100d157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561011e57805f5260205f20601f840160051c810160208510156100fc5750805b601f840160051c820191505b8181101561011b575f8155600101610108565b50505b505050565b81516001600160401b0381111561013c5761013c61008b565b6101508161014a845461009f565b846100d7565b6020601f821160018114610182575f831561016b5750848201515b5f19600385901b1c1916600184901b17845561011b565b5f84815260208120601f198516915b828110156101b15787850151825560209485019460019092019101610191565b50848210156101ce57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610e03806101ea5f395ff3fe6080604052600436106100d1575f3560e01c8063313ce5671161007c57806395d89b411161005757806395d89b4114610265578063a9059cbb14610279578063d0e30db014610298578063dd62ed3e146102a0575f5ffd5b8063313ce567146101b85780634cf088d9146101d357806370a0823114610224575f5ffd5b8063205c2878116100ac578063205c28781461015b57806323b872dd1461017a5780632e1a7d4d14610199575f5ffd5b806306fdde03146100e4578063095ea7b31461010e57806318160ddd1461013d575f5ffd5b366100e0576100de6102f1565b005b5f5ffd5b3480156100ef575f5ffd5b506100f8610428565b6040516101059190610bdd565b60405180910390f35b348015610119575f5ffd5b5061012d610128366004610c51565b6104b8565b6040519015158152602001610105565b348015610148575f5ffd5b506002545b604051908152602001610105565b348015610166575f5ffd5b506100de610175366004610c51565b6104d1565b348015610185575f5ffd5b5061012d610194366004610c7b565b610668565b3480156101a4575f5ffd5b506100de6101b3366004610cb9565b61068b565b3480156101c3575f5ffd5b5060405160128152602001610105565b3480156101de575f5ffd5b506005546101ff9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b34801561022f575f5ffd5b5061014d61023e366004610cd0565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b348015610270575f5ffd5b506100f8610698565b348015610284575f5ffd5b5061012d610293366004610c51565b6106a7565b6100de6102f1565b3480156102ab575f5ffd5b5061014d6102ba366004610cf2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561035b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037f9190610d29565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e3576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f341161041c576040517fd8df41ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61042633346106b4565b565b60606003805461043790610d44565b80601f016020809104026020016040519081016040528092919081815260200182805461046390610d44565b80156104ae5780601f10610485576101008083540402835291602001916104ae565b820191905f5260205f20905b81548152906001019060200180831161049157829003601f168201915b5050505050905090565b5f336104c5818585610717565b60019150505b92915050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561053b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055f9190610d29565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105c3576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105cd3382610724565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610623576040519150601f19603f3d011682016040523d82523d5f602084013e610628565b606091505b5050905080610663576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f3361067585828561077e565b61068085858561084c565b506001949350505050565b61069533826104d1565b50565b60606004805461043790610d44565b5f336104c581858561084c565b73ffffffffffffffffffffffffffffffffffffffff8216610708576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6107135f83836108f1565b5050565b6106638383836001610a98565b73ffffffffffffffffffffffffffffffffffffffff8216610773576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b610713825f836108f1565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156108465781811015610838576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016106ff565b61084684848484035f610a98565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661089b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b73ffffffffffffffffffffffffffffffffffffffff82166108ea576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b6106638383835b73ffffffffffffffffffffffffffffffffffffffff8316610928578060025f82825461091d9190610d95565b909155506109d89050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156109ad576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016106ff565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610a0157600280548290039055610a2c565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610a8b91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8416610ae7576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b73ffffffffffffffffffffffffffffffffffffffff8316610b36576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610846578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bcf91815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff81168114610695575f5ffd5b5f5f60408385031215610c62575f5ffd5b8235610c6d81610c30565b946020939093013593505050565b5f5f5f60608486031215610c8d575f5ffd5b8335610c9881610c30565b92506020840135610ca881610c30565b929592945050506040919091013590565b5f60208284031215610cc9575f5ffd5b5035919050565b5f60208284031215610ce0575f5ffd5b8135610ceb81610c30565b9392505050565b5f5f60408385031215610d03575f5ffd5b8235610d0e81610c30565b91506020830135610d1e81610c30565b809150509250929050565b5f60208284031215610d39575f5ffd5b8151610ceb81610c30565b600181811c90821680610d5857607f821691505b602082108103610d8f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156104cb577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea264697066735822122049e14f0b6780996119ea14be68a95b0d6f9fdf536f862c63c8f458bd781fbe9e64736f6c634300081c0033a264697066735822122055e6859dbfea50ef6ad3fd688ca624721a5b6d8358d9811109db2d59a7e67a8764736f6c634300081c0033",
}

// CrossStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossStakingMetaData.ABI instead.
var CrossStakingABI = CrossStakingMetaData.ABI

// CrossStakingBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossStakingBinRuntime = "608060405260043610610302575f3560e01c8063a1eda53c11610191578063cefc1429116100dc578063e759026811610087578063f19c3d5b11610062578063f19c3d5b14610a5a578063f887ea4014610a6e578063fe96e4ff14610a8d575f5ffd5b8063e7590268146109e7578063eba32d83146109fb578063ec87621c14610a27575f5ffd5b8063d547741f116100b7578063d547741f14610995578063d602b9fd146109b4578063d6f74898146109c8575f5ffd5b8063cefc1429146108ec578063cf6eefb714610900578063d4175be21461096a575f5ffd5b8063b9c17c4d1161013c578063cc8463c811610117578063cc8463c81461089a578063ce24af53146108ae578063cefa7799146108cd575f5ffd5b8063b9c17c4d1461083d578063c0d786551461085c578063caa9a08d1461087b575f5ffd5b8063ac4afa381161016c578063ac4afa381461074c578063ad3cb1cc146107c9578063b34c972e1461081e575f5ffd5b8063a1eda53c146106e7578063a217fddf1461071a578063a2db45821461072d575f5ffd5b806335cc9cb411610251578063649a5ec7116101fc57806391cf6d3e116101d757806391cf6d3e1461065157806391d1485414610665578063a1635945146106c8575f5ffd5b8063649a5ec7146105e257806384ef8ffc146106015780638da5cb5b1461063d575f5ffd5b80634f1ef2861161022c5780634f1ef2861461059c57806352d1902d146105af578063634e93da146105c3575f5ffd5b806335cc9cb41461053f57806336568abe1461055e5780633b95352b1461057d575f5ffd5b80631538af09116102b1578063248a9ca31161028c578063248a9ca31461046b5780632f2ff15d146104b85780632f380b35146104d7575f5ffd5b80631538af0914610418578063155fff621461043757806318e5613114610456575f5ffd5b80630aa6220b116102e15780630aa6220b146103995780630f538b40146103af57806312d36171146103dc575f5ffd5b8062a5ae211461030657806301ffc9a714610342578063022d63fb14610371575b5f5ffd5b348015610311575f5ffd5b506103256103203660046130b9565b610aa1565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561034d575f5ffd5b5061036161035c3660046130d0565b610b0f565b6040519015158152602001610339565b34801561037c575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610339565b3480156103a4575f5ffd5b506103ad610b6a565b005b3480156103ba575f5ffd5b506103ce6103c9366004613123565b610b7f565b604051908152602001610339565b3480156103e7575f5ffd5b506103fb6103f636600461313e565b610b9f565b604080519283526001600160a01b03909116602083015201610339565b348015610423575f5ffd5b506103ad610432366004613168565b610e31565b348015610442575f5ffd5b506103ce6104513660046130b9565b610f32565b348015610461575f5ffd5b506103ce60045481565b348015610476575f5ffd5b506103ce6104853660046130b9565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b3480156104c3575f5ffd5b506103ad6104d2366004613188565b610f3e565b3480156104e2575f5ffd5b506104f66104f13660046130b9565b610f83565b60405161033991905f608082019050825182526001600160a01b0360208401511660208301526001600160a01b0360408401511660408301526060830151606083015292915050565b34801561054a575f5ffd5b506103ad610559366004613188565b611048565b348015610569575f5ffd5b506103ad610578366004613188565b61111c565b348015610588575f5ffd5b506103ce61059736600461313e565b61126e565b6103ad6105aa3660046131e3565b611296565b3480156105ba575f5ffd5b506103ce6112b1565b3480156105ce575f5ffd5b506103ad6105dd366004613123565b6112df565b3480156105ed575f5ffd5b506103ad6105fc3660046132c4565b6112f2565b34801561060c575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b0316610325565b348015610648575f5ffd5b50610325611305565b34801561065c575f5ffd5b506103ce5f5481565b348015610670575f5ffd5b5061036161067f366004613188565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156106d3575f5ffd5b506103ad6106e2366004613188565b61133c565b3480156106f2575f5ffd5b506106fb611410565b6040805165ffffffffffff938416815292909116602083015201610339565b348015610725575f5ffd5b506103ce5f81565b348015610738575f5ffd5b50600154610325906001600160a01b031681565b348015610757575f5ffd5b506107996107663660046130b9565b60056020525f9081526040902080546001820154600283015460039093015491926001600160a01b039182169291169084565b60405161033994939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b3480156107d4575f5ffd5b506108116040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b604051610339919061330b565b348015610829575f5ffd5b506103ad610838366004613329565b6114cd565b348015610848575f5ffd5b506103ad61085736600461334c565b61159f565b348015610867575f5ffd5b506103ad610876366004613123565b61178d565b348015610886575f5ffd5b506103ce610895366004613123565b611839565b3480156108a5575f5ffd5b5061038261188a565b3480156108b9575f5ffd5b506103ad6108c836600461338b565b61196a565b3480156108d8575f5ffd5b50600354610325906001600160a01b031681565b3480156108f7575f5ffd5b506103ad611c20565b34801561090b575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610339565b348015610975575f5ffd5b506103ce610984366004613123565b60066020525f908152604090205481565b3480156109a0575f5ffd5b506103ad6109af366004613188565b611c8e565b3480156109bf575f5ffd5b506103ad611ccf565b3480156109d3575f5ffd5b506103ad6109e2366004613123565b611ce1565b3480156109f2575f5ffd5b506103ce611d8d565b348015610a06575f5ffd5b50610a1a610a15366004613123565b611d98565b60405161033991906133cf565b348015610a32575f5ffd5b506103ce7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b348015610a65575f5ffd5b50610a1a611dbb565b348015610a79575f5ffd5b50600254610325906001600160a01b031681565b348015610a98575f5ffd5b50610a1a611dc7565b5f818152600560205260408120600101546001600160a01b0316610af1576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f908152600560205260409020600101546001600160a01b031690565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610b645750610b6482611fab565b92915050565b5f610b7481612041565b610b7c61204b565b50565b6001600160a01b0381165f908152600760205260408120610b6490612057565b5f5f7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610bcb81612041565b6001600160a01b038516610c0b576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8411610c44576040517f5e7c0cae00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60048054905f610c538361343e565b909155506040516001600160a01b0387166024820152604481018690529093505f9060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcd6dc6870000000000000000000000000000000000000000000000000000000017905260035490519192505f916001600160a01b03909116908390610cf19061309f565b610cfc929190613475565b604051809103905ff080158015610d15573d5f5f3e3d5ffd5b50604080516080810182528781526001600160a01b0380841660208084018281528d841685870190815242606087019081525f8e8152600585528881209751885592516001880180549188167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055915160028801805491909716921691909117909455925160039094019390935581526006909152208690559350839050610dc3600886612060565b506001600160a01b0387165f908152600760205260409020610de59086612060565b50866001600160a01b0316846001600160a01b0316867f1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a660405160405180910390a45050509250929050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b08610e5b81612041565b5f838152600560205260409020600101546001600160a01b0316610eab576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517fd8060cd5000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b039091169063d8060cd5906024015b5f604051808303815f87803b158015610f17575f5ffd5b505af1158015610f29573d5f5f3e3d5ffd5b50505050505050565b5f610b6460088361206b565b81610f75576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7f8282612076565b5050565b604080516080810182525f8082526020820181905291810182905260608101919091525f828152600560205260409020600101546001600160a01b0316610ff6576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f9081526005602090815260409182902082516080810184528154815260018201546001600160a01b0390811693820193909352600282015490921692820192909252600390910154606082015290565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861107281612041565b5f838152600560205260409020600101546001600160a01b03166110c2576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f3d509c970000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690633d509c9790602401610f00565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561117757507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b1561125f577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806111dd575065ffffffffffff8116155b806111f057504265ffffffffffff821610155b15611236576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff821660048201526024015b60405180910390fd5b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b61126983836120bf565b505050565b6001600160a01b0382165f90815260076020526040812061128f908361206b565b9392505050565b61129e61210b565b6112a7826121db565b610f7f82826121e5565b5f6112ba6122e6565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f6112e981612041565b610f7f82612348565b5f6112fc81612041565b610f7f826123ba565b5f6113377feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b905090565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0861136681612041565b5f838152600560205260409020600101546001600160a01b03166113b6576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f1c03e6cc0000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015290911690631c03e6cc90602401610f00565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400811580159061149257504265ffffffffffff831610155b61149d575f5f6114c4565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086114f781612041565b5f838152600560205260409020600101546001600160a01b0316611547576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83815260056020526040908190206001015490517f6d7c49a20000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636d7c49a290610f009085906004016134c3565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b086115c981612041565b5f848152600560205260409020600101546001600160a01b0316611619576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f848152600560205260408082206001015490517ff67771750000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301529091169063f677717590602401602060405180830381865afa158015611688573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116ac9190613502565b5f86815260056020526040908190206001015490517ff940e3850000000000000000000000000000000000000000000000000000000081526001600160a01b038781166004830152868116602483015292935091169063f940e385906044015f604051808303815f87803b158015611722575f5ffd5b505af1158015611734573d5f5f3e3d5ffd5b50505050826001600160a01b0316846001600160a01b0316867f408708e2d99a8cfa59be1466864d74f6ddddc62b5ddc3bb11f1c3cc3ce9ed65f8460405161177e91815260200190565b60405180910390a45050505050565b5f61179781612041565b6001600160a01b0382166117d7576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15905f90a25050565b6001600160a01b0381165f90815260066020526040812054808203610b64576040517fc986e59800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801580159061190c57504265ffffffffffff8216105b61193d5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611963565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b5f611973612429565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f8115801561199f5750825b90505f8267ffffffffffffffff1660011480156119bb5750303b155b9050811580156119c9575080155b15611a00576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611a615784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816611aa1576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716611ae1576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611aeb8688612451565b611af3612463565b435f55600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038a16179055604051611b35906130ac565b604051809103905ff080158015611b4e573d5f5f3e3d5ffd5b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091178155600455611bb47f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b088861246b565b508315611c165784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114611c86576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161122d565b610b7c61254c565b81611cc5576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7f8282612661565b5f611cd981612041565b610b7c6126a4565b5f611ceb81612041565b6001600160a01b038216611d2b576040517f21dce27300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040517fdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957905f90a25050565b5f6113376008612057565b6001600160a01b0381165f908152600760205260409020606090610b64906126ae565b606061133760086126ae565b60605f611dd46008612057565b90505f8167ffffffffffffffff811115611df057611df06131b6565b604051908082528060200260200182016040528015611e19578160200160208202803683370190505b5090505f805b83811015611f12575f611e3360088361206b565b90505f5f828152600560209081526040918290206001015482517ff022869200000000000000000000000000000000000000000000000000000000815292516001600160a01b039091169263f02286929260048083019391928290030181865afa158015611ea3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ec79190613519565b6002811115611ed857611ed8613496565b03611f095780848481518110611ef057611ef0613534565b602090810291909101015282611f058161343e565b9350505b50600101611e1f565b505f8167ffffffffffffffff811115611f2d57611f2d6131b6565b604051908082528060200260200182016040528015611f56578160200160208202803683370190505b5090505f5b82811015611fa257838181518110611f7557611f75613534565b6020026020010151828281518110611f8f57611f8f613534565b6020908102919091010152600101611f5b565b50949350505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610b6457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610b64565b610b7c81336126ba565b6120555f5f612746565b565b5f610b64825490565b5f61128f83836128d1565b5f61128f838361291d565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120af81612041565b6120b9838361246b565b50505050565b6001600160a01b0381163314612101576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112698282612943565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806121a457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166121987f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610f7f81612041565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561223f575060408051601f3d908101601f1916820190925261223c91810190613502565b60015b612280576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161122d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146122dc576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161122d565b61126983836129d9565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612055576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61235161188a565b61235a42612a2e565b6123649190613561565b90506123708282612a7d565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6123c482612b2a565b6123cd42612a2e565b6123d79190613561565b90506123e38282612746565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610b64565b612459612b71565b610f7f8282612baf565b612055612b71565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008361253a575f6124c47feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b031614612504576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b6125448484612c6b565b949350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff168015806125af57504265ffffffffffff821610155b156125f0576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161122d565b61262a5f6126257feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b612943565b506126355f8361246b565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461269a81612041565b6120b98383612943565b6120555f5f612a7d565b60605f61128f83612d55565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f7f576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161122d565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612858574265ffffffffffff8216101561282f576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255612858565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f81815260018301602052604081205461291657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610b64565b505f610b64565b5f825f01828154811061293257612932613534565b905f5260205f200154905092915050565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561299f57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b156129cf576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6125448484612dae565b6129e282612e70565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612a26576112698282612f17565b610f7f612f89565b5f65ffffffffffff821115612a79576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161122d565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b0388161717845591041680156120b9576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f612b3461188a565b90508065ffffffffffff168365ffffffffffff1611612b5c57612b57838261357f565b61128f565b61128f65ffffffffffff841662069780612fc1565b612b79612fd0565b612055576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bb7612b71565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216612c1a576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f600482015260240161122d565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556120b95f8361246b565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612d4c575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612d023390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610b64565b5f915050610b64565b6060815f01805480602002602001604051908101604052809291908181526020018280548015612da257602002820191905f5260205f20905b815481526020019060010190808311612d8e575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612d4c575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610b64565b806001600160a01b03163b5f03612ebe576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161122d565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051612f33919061359d565b5f60405180830381855af49150503d805f8114612f6b576040519150601f19603f3d011682016040523d82523d5f602084013e612f70565b606091505b5091509150612f80858383612fee565b95945050505050565b3415612055576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82821882841002821861128f565b5f612fd9612429565b5468010000000000000000900460ff16919050565b606082612ffe57612b578261305e565b815115801561301557506001600160a01b0384163b155b15613057576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161122d565b5092915050565b80511561306d57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103dc806135b483390190565b610fed8061399083390190565b5f602082840312156130c9575f5ffd5b5035919050565b5f602082840312156130e0575f5ffd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461128f575f5ffd5b6001600160a01b0381168114610b7c575f5ffd5b5f60208284031215613133575f5ffd5b813561128f8161310f565b5f5f6040838503121561314f575f5ffd5b823561315a8161310f565b946020939093013593505050565b5f5f60408385031215613179575f5ffd5b50508035926020909101359150565b5f5f60408385031215613199575f5ffd5b8235915060208301356131ab8161310f565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f604083850312156131f4575f5ffd5b82356131ff8161310f565b9150602083013567ffffffffffffffff81111561321a575f5ffd5b8301601f8101851361322a575f5ffd5b803567ffffffffffffffff811115613244576132446131b6565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715613274576132746131b6565b60405281815282820160200187101561328b575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b803565ffffffffffff811681146132bf575f5ffd5b919050565b5f602082840312156132d4575f5ffd5b61128f826132aa565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61128f60208301846132dd565b60038110610b7c575f5ffd5b5f5f6040838503121561333a575f5ffd5b8235915060208301356131ab8161331d565b5f5f5f6060848603121561335e575f5ffd5b8335925060208401356133708161310f565b915060408401356133808161310f565b809150509250925092565b5f5f5f6060848603121561339d575f5ffd5b83356133a88161310f565b925060208401356133b88161310f565b91506133c6604085016132aa565b90509250925092565b602080825282518282018190525f918401906040840190835b818110156134065783518352602093840193909201916001016133e8565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361346e5761346e613411565b5060010190565b6001600160a01b0383168152604060208201525f61254460408301846132dd565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60208101600383106134fc577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b5f60208284031215613512575f5ffd5b5051919050565b5f60208284031215613529575f5ffd5b815161128f8161331d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b65ffffffffffff8181168382160190811115610b6457610b64613411565b65ffffffffffff8281168282160390811115610b6457610b64613411565b5f82518060208501845e5f92019182525091905056fe60806040526040516103dc3803806103dc8339810160408190526100229161023b565b61002c8282610033565b5050610320565b61003c82610091565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561008557610080828261010c565b505050565b61008d61017f565b5050565b806001600160a01b03163b5f036100cb57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610128919061030a565b5f60405180830381855af49150503d805f8114610160576040519150601f19603f3d011682016040523d82523d5f602084013e610165565b606091505b5090925090506101768583836101a0565b95945050505050565b341561019e5760405163b398979f60e01b815260040160405180910390fd5b565b6060826101b5576101b0826101ff565b6101f8565b81511580156101cc57506001600160a01b0384163b155b156101f557604051639996b31560e01b81526001600160a01b03851660048201526024016100c2565b50805b9392505050565b80511561020e57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561024c575f5ffd5b82516001600160a01b0381168114610262575f5ffd5b60208401519092506001600160401b0381111561027d575f5ffd5b8301601f8101851361028d575f5ffd5b80516001600160401b038111156102a6576102a6610227565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102d4576102d4610227565b6040528181528282016020018710156102eb575f5ffd5b8160208401602083015e5f602083830101528093505050509250929050565b5f82518060208501845e5f920191825250919050565b60b08061032c5f395ff3fe6080604052600a600c565b005b60186014601a565b605d565b565b5f60587f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156076573d5ff35b3d5ffdfea2646970667358221220927d2839b1d9f37b4bfae40308f9fa5bd1e749b8d8b98c249d9b67e246bf55ba64736f6c634300081c0033608060405234801561000f575f5ffd5b506040518060400160405280600d81526020016c577261707065642043524f535360981b815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100649190610123565b5060046100718282610123565b5050600580546001600160a01b03191633179055506101dd565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100b357607f821691505b6020821081036100d157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561011e57805f5260205f20601f840160051c810160208510156100fc5750805b601f840160051c820191505b8181101561011b575f8155600101610108565b50505b505050565b81516001600160401b0381111561013c5761013c61008b565b6101508161014a845461009f565b846100d7565b6020601f821160018114610182575f831561016b5750848201515b5f19600385901b1c1916600184901b17845561011b565b5f84815260208120601f198516915b828110156101b15787850151825560209485019460019092019101610191565b50848210156101ce57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610e03806101ea5f395ff3fe6080604052600436106100d1575f3560e01c8063313ce5671161007c57806395d89b411161005757806395d89b4114610265578063a9059cbb14610279578063d0e30db014610298578063dd62ed3e146102a0575f5ffd5b8063313ce567146101b85780634cf088d9146101d357806370a0823114610224575f5ffd5b8063205c2878116100ac578063205c28781461015b57806323b872dd1461017a5780632e1a7d4d14610199575f5ffd5b806306fdde03146100e4578063095ea7b31461010e57806318160ddd1461013d575f5ffd5b366100e0576100de6102f1565b005b5f5ffd5b3480156100ef575f5ffd5b506100f8610428565b6040516101059190610bdd565b60405180910390f35b348015610119575f5ffd5b5061012d610128366004610c51565b6104b8565b6040519015158152602001610105565b348015610148575f5ffd5b506002545b604051908152602001610105565b348015610166575f5ffd5b506100de610175366004610c51565b6104d1565b348015610185575f5ffd5b5061012d610194366004610c7b565b610668565b3480156101a4575f5ffd5b506100de6101b3366004610cb9565b61068b565b3480156101c3575f5ffd5b5060405160128152602001610105565b3480156101de575f5ffd5b506005546101ff9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b34801561022f575f5ffd5b5061014d61023e366004610cd0565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b348015610270575f5ffd5b506100f8610698565b348015610284575f5ffd5b5061012d610293366004610c51565b6106a7565b6100de6102f1565b3480156102ab575f5ffd5b5061014d6102ba366004610cf2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561035b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037f9190610d29565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e3576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f341161041c576040517fd8df41ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61042633346106b4565b565b60606003805461043790610d44565b80601f016020809104026020016040519081016040528092919081815260200182805461046390610d44565b80156104ae5780601f10610485576101008083540402835291602001916104ae565b820191905f5260205f20905b81548152906001019060200180831161049157829003601f168201915b5050505050905090565b5f336104c5818585610717565b60019150505b92915050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561053b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055f9190610d29565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105c3576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105cd3382610724565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610623576040519150601f19603f3d011682016040523d82523d5f602084013e610628565b606091505b5050905080610663576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f3361067585828561077e565b61068085858561084c565b506001949350505050565b61069533826104d1565b50565b60606004805461043790610d44565b5f336104c581858561084c565b73ffffffffffffffffffffffffffffffffffffffff8216610708576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6107135f83836108f1565b5050565b6106638383836001610a98565b73ffffffffffffffffffffffffffffffffffffffff8216610773576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b610713825f836108f1565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156108465781811015610838576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064016106ff565b61084684848484035f610a98565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661089b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b73ffffffffffffffffffffffffffffffffffffffff82166108ea576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b6106638383835b73ffffffffffffffffffffffffffffffffffffffff8316610928578060025f82825461091d9190610d95565b909155506109d89050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156109ad576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016106ff565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610a0157600280548290039055610a2c565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610a8b91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8416610ae7576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b73ffffffffffffffffffffffffffffffffffffffff8316610b36576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016106ff565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610846578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bcf91815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff81168114610695575f5ffd5b5f5f60408385031215610c62575f5ffd5b8235610c6d81610c30565b946020939093013593505050565b5f5f5f60608486031215610c8d575f5ffd5b8335610c9881610c30565b92506020840135610ca881610c30565b929592945050506040919091013590565b5f60208284031215610cc9575f5ffd5b5035919050565b5f60208284031215610ce0575f5ffd5b8135610ceb81610c30565b9392505050565b5f5f60408385031215610d03575f5ffd5b8235610d0e81610c30565b91506020830135610d1e81610c30565b809150509250929050565b5f60208284031215610d39575f5ffd5b8151610ceb81610c30565b600181811c90821680610d5857607f821691505b602082108103610d8f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156104cb577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea264697066735822122049e14f0b6780996119ea14be68a95b0d6f9fdf536f862c63c8f458bd781fbe9e64736f6c634300081c0033a264697066735822122055e6859dbfea50ef6ad3fd688ca624721a5b6d8358d9811109db2d59a7e67a8764736f6c634300081c0033"

// Deprecated: Use CrossStakingMetaData.Sigs instead.
// CrossStakingFuncSigs maps the 4-byte function signature to its string representation.
var CrossStakingFuncSigs = CrossStakingMetaData.Sigs

// CrossStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossStakingMetaData.Bin instead.
var CrossStakingBin = CrossStakingMetaData.Bin

// DeployCrossStaking deploys a new Ethereum contract, binding an instance of CrossStaking to it.
func DeployCrossStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossStaking, error) {
	parsed, err := CrossStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossStakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossStaking{CrossStakingCaller: CrossStakingCaller{contract: contract}, CrossStakingTransactor: CrossStakingTransactor{contract: contract}, CrossStakingFilterer: CrossStakingFilterer{contract: contract}}, nil
}

// CrossStaking is an auto generated Go binding around an Ethereum contract.
type CrossStaking struct {
	CrossStakingCaller     // Read-only binding to the contract
	CrossStakingTransactor // Write-only binding to the contract
	CrossStakingFilterer   // Log filterer for contract events
}

// CrossStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossStakingSession struct {
	Contract     *CrossStaking     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossStakingCallerSession struct {
	Contract *CrossStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrossStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossStakingTransactorSession struct {
	Contract     *CrossStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrossStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossStakingRaw struct {
	Contract *CrossStaking // Generic contract binding to access the raw methods on
}

// CrossStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossStakingCallerRaw struct {
	Contract *CrossStakingCaller // Generic read-only contract binding to access the raw methods on
}

// CrossStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossStakingTransactorRaw struct {
	Contract *CrossStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossStaking creates a new instance of CrossStaking, bound to a specific deployed contract.
func NewCrossStaking(address common.Address, backend bind.ContractBackend) (*CrossStaking, error) {
	contract, err := bindCrossStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossStaking{CrossStakingCaller: CrossStakingCaller{contract: contract}, CrossStakingTransactor: CrossStakingTransactor{contract: contract}, CrossStakingFilterer: CrossStakingFilterer{contract: contract}}, nil
}

// NewCrossStakingCaller creates a new read-only instance of CrossStaking, bound to a specific deployed contract.
func NewCrossStakingCaller(address common.Address, caller bind.ContractCaller) (*CrossStakingCaller, error) {
	contract, err := bindCrossStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossStakingCaller{contract: contract}, nil
}

// NewCrossStakingTransactor creates a new write-only instance of CrossStaking, bound to a specific deployed contract.
func NewCrossStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossStakingTransactor, error) {
	contract, err := bindCrossStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossStakingTransactor{contract: contract}, nil
}

// NewCrossStakingFilterer creates a new log filterer instance of CrossStaking, bound to a specific deployed contract.
func NewCrossStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossStakingFilterer, error) {
	contract, err := bindCrossStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossStakingFilterer{contract: contract}, nil
}

// bindCrossStaking binds a generic wrapper to an already deployed contract.
func bindCrossStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossStaking *CrossStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossStaking.Contract.CrossStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossStaking *CrossStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStaking.Contract.CrossStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossStaking *CrossStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossStaking.Contract.CrossStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossStaking *CrossStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossStaking *CrossStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossStaking *CrossStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossStaking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossStaking *CrossStakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossStaking *CrossStakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossStaking.Contract.DEFAULTADMINROLE(&_CrossStaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossStaking *CrossStakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossStaking.Contract.DEFAULTADMINROLE(&_CrossStaking.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CrossStaking *CrossStakingCaller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CrossStaking *CrossStakingSession) MANAGERROLE() ([32]byte, error) {
	return _CrossStaking.Contract.MANAGERROLE(&_CrossStaking.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_CrossStaking *CrossStakingCallerSession) MANAGERROLE() ([32]byte, error) {
	return _CrossStaking.Contract.MANAGERROLE(&_CrossStaking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossStaking *CrossStakingCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossStaking *CrossStakingSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossStaking.Contract.UPGRADEINTERFACEVERSION(&_CrossStaking.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossStaking *CrossStakingCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossStaking.Contract.UPGRADEINTERFACEVERSION(&_CrossStaking.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossStaking *CrossStakingCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossStaking *CrossStakingSession) DefaultAdmin() (common.Address, error) {
	return _CrossStaking.Contract.DefaultAdmin(&_CrossStaking.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossStaking *CrossStakingCallerSession) DefaultAdmin() (common.Address, error) {
	return _CrossStaking.Contract.DefaultAdmin(&_CrossStaking.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossStaking *CrossStakingCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossStaking *CrossStakingSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossStaking.Contract.DefaultAdminDelay(&_CrossStaking.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossStaking *CrossStakingCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossStaking.Contract.DefaultAdminDelay(&_CrossStaking.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossStaking *CrossStakingCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossStaking *CrossStakingSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossStaking.Contract.DefaultAdminDelayIncreaseWait(&_CrossStaking.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossStaking *CrossStakingCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossStaking.Contract.DefaultAdminDelayIncreaseWait(&_CrossStaking.CallOpts)
}

// GetActivePoolIds is a free data retrieval call binding the contract method 0xfe96e4ff.
//
// Solidity: function getActivePoolIds() view returns(uint256[])
func (_CrossStaking *CrossStakingCaller) GetActivePoolIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getActivePoolIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActivePoolIds is a free data retrieval call binding the contract method 0xfe96e4ff.
//
// Solidity: function getActivePoolIds() view returns(uint256[])
func (_CrossStaking *CrossStakingSession) GetActivePoolIds() ([]*big.Int, error) {
	return _CrossStaking.Contract.GetActivePoolIds(&_CrossStaking.CallOpts)
}

// GetActivePoolIds is a free data retrieval call binding the contract method 0xfe96e4ff.
//
// Solidity: function getActivePoolIds() view returns(uint256[])
func (_CrossStaking *CrossStakingCallerSession) GetActivePoolIds() ([]*big.Int, error) {
	return _CrossStaking.Contract.GetActivePoolIds(&_CrossStaking.CallOpts)
}

// GetAllPoolIds is a free data retrieval call binding the contract method 0xf19c3d5b.
//
// Solidity: function getAllPoolIds() view returns(uint256[])
func (_CrossStaking *CrossStakingCaller) GetAllPoolIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getAllPoolIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllPoolIds is a free data retrieval call binding the contract method 0xf19c3d5b.
//
// Solidity: function getAllPoolIds() view returns(uint256[])
func (_CrossStaking *CrossStakingSession) GetAllPoolIds() ([]*big.Int, error) {
	return _CrossStaking.Contract.GetAllPoolIds(&_CrossStaking.CallOpts)
}

// GetAllPoolIds is a free data retrieval call binding the contract method 0xf19c3d5b.
//
// Solidity: function getAllPoolIds() view returns(uint256[])
func (_CrossStaking *CrossStakingCallerSession) GetAllPoolIds() ([]*big.Int, error) {
	return _CrossStaking.Contract.GetAllPoolIds(&_CrossStaking.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x00a5ae21.
//
// Solidity: function getPoolAddress(uint256 poolId) view returns(address)
func (_CrossStaking *CrossStakingCaller) GetPoolAddress(opts *bind.CallOpts, poolId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getPoolAddress", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAddress is a free data retrieval call binding the contract method 0x00a5ae21.
//
// Solidity: function getPoolAddress(uint256 poolId) view returns(address)
func (_CrossStaking *CrossStakingSession) GetPoolAddress(poolId *big.Int) (common.Address, error) {
	return _CrossStaking.Contract.GetPoolAddress(&_CrossStaking.CallOpts, poolId)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x00a5ae21.
//
// Solidity: function getPoolAddress(uint256 poolId) view returns(address)
func (_CrossStaking *CrossStakingCallerSession) GetPoolAddress(poolId *big.Int) (common.Address, error) {
	return _CrossStaking.Contract.GetPoolAddress(&_CrossStaking.CallOpts, poolId)
}

// GetPoolCountByStakingToken is a free data retrieval call binding the contract method 0x0f538b40.
//
// Solidity: function getPoolCountByStakingToken(address stakingToken) view returns(uint256)
func (_CrossStaking *CrossStakingCaller) GetPoolCountByStakingToken(opts *bind.CallOpts, stakingToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getPoolCountByStakingToken", stakingToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolCountByStakingToken is a free data retrieval call binding the contract method 0x0f538b40.
//
// Solidity: function getPoolCountByStakingToken(address stakingToken) view returns(uint256)
func (_CrossStaking *CrossStakingSession) GetPoolCountByStakingToken(stakingToken common.Address) (*big.Int, error) {
	return _CrossStaking.Contract.GetPoolCountByStakingToken(&_CrossStaking.CallOpts, stakingToken)
}

// GetPoolCountByStakingToken is a free data retrieval call binding the contract method 0x0f538b40.
//
// Solidity: function getPoolCountByStakingToken(address stakingToken) view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) GetPoolCountByStakingToken(stakingToken common.Address) (*big.Int, error) {
	return _CrossStaking.Contract.GetPoolCountByStakingToken(&_CrossStaking.CallOpts, stakingToken)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address pool) view returns(uint256)
func (_CrossStaking *CrossStakingCaller) GetPoolId(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getPoolId", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address pool) view returns(uint256)
func (_CrossStaking *CrossStakingSession) GetPoolId(pool common.Address) (*big.Int, error) {
	return _CrossStaking.Contract.GetPoolId(&_CrossStaking.CallOpts, pool)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address pool) view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) GetPoolId(pool common.Address) (*big.Int, error) {
	return _CrossStaking.Contract.GetPoolId(&_CrossStaking.CallOpts, pool)
}

// GetPoolIdsByStakingToken is a free data retrieval call binding the contract method 0xeba32d83.
//
// Solidity: function getPoolIdsByStakingToken(address stakingToken) view returns(uint256[])
func (_CrossStaking *CrossStakingCaller) GetPoolIdsByStakingToken(opts *bind.CallOpts, stakingToken common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getPoolIdsByStakingToken", stakingToken)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPoolIdsByStakingToken is a free data retrieval call binding the contract method 0xeba32d83.
//
// Solidity: function getPoolIdsByStakingToken(address stakingToken) view returns(uint256[])
func (_CrossStaking *CrossStakingSession) GetPoolIdsByStakingToken(stakingToken common.Address) ([]*big.Int, error) {
	return _CrossStaking.Contract.GetPoolIdsByStakingToken(&_CrossStaking.CallOpts, stakingToken)
}

// GetPoolIdsByStakingToken is a free data retrieval call binding the contract method 0xeba32d83.
//
// Solidity: function getPoolIdsByStakingToken(address stakingToken) view returns(uint256[])
func (_CrossStaking *CrossStakingCallerSession) GetPoolIdsByStakingToken(stakingToken common.Address) ([]*big.Int, error) {
	return _CrossStaking.Contract.GetPoolIdsByStakingToken(&_CrossStaking.CallOpts, stakingToken)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,address,uint256))
func (_CrossStaking *CrossStakingCaller) GetPoolInfo(opts *bind.CallOpts, poolId *big.Int) (ICrossStakingPoolInfo, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getPoolInfo", poolId)

	if err != nil {
		return *new(ICrossStakingPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossStakingPoolInfo)).(*ICrossStakingPoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,address,uint256))
func (_CrossStaking *CrossStakingSession) GetPoolInfo(poolId *big.Int) (ICrossStakingPoolInfo, error) {
	return _CrossStaking.Contract.GetPoolInfo(&_CrossStaking.CallOpts, poolId)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 poolId) view returns((uint256,address,address,uint256))
func (_CrossStaking *CrossStakingCallerSession) GetPoolInfo(poolId *big.Int) (ICrossStakingPoolInfo, error) {
	return _CrossStaking.Contract.GetPoolInfo(&_CrossStaking.CallOpts, poolId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossStaking *CrossStakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossStaking *CrossStakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossStaking.Contract.GetRoleAdmin(&_CrossStaking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossStaking *CrossStakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossStaking.Contract.GetRoleAdmin(&_CrossStaking.CallOpts, role)
}

// GetTotalPoolCount is a free data retrieval call binding the contract method 0xe7590268.
//
// Solidity: function getTotalPoolCount() view returns(uint256)
func (_CrossStaking *CrossStakingCaller) GetTotalPoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "getTotalPoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPoolCount is a free data retrieval call binding the contract method 0xe7590268.
//
// Solidity: function getTotalPoolCount() view returns(uint256)
func (_CrossStaking *CrossStakingSession) GetTotalPoolCount() (*big.Int, error) {
	return _CrossStaking.Contract.GetTotalPoolCount(&_CrossStaking.CallOpts)
}

// GetTotalPoolCount is a free data retrieval call binding the contract method 0xe7590268.
//
// Solidity: function getTotalPoolCount() view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) GetTotalPoolCount() (*big.Int, error) {
	return _CrossStaking.Contract.GetTotalPoolCount(&_CrossStaking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossStaking *CrossStakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossStaking *CrossStakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossStaking.Contract.HasRole(&_CrossStaking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossStaking *CrossStakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossStaking.Contract.HasRole(&_CrossStaking.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossStaking *CrossStakingCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossStaking *CrossStakingSession) InitializedAt() (*big.Int, error) {
	return _CrossStaking.Contract.InitializedAt(&_CrossStaking.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) InitializedAt() (*big.Int, error) {
	return _CrossStaking.Contract.InitializedAt(&_CrossStaking.CallOpts)
}

// NextPoolId is a free data retrieval call binding the contract method 0x18e56131.
//
// Solidity: function nextPoolId() view returns(uint256)
func (_CrossStaking *CrossStakingCaller) NextPoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "nextPoolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPoolId is a free data retrieval call binding the contract method 0x18e56131.
//
// Solidity: function nextPoolId() view returns(uint256)
func (_CrossStaking *CrossStakingSession) NextPoolId() (*big.Int, error) {
	return _CrossStaking.Contract.NextPoolId(&_CrossStaking.CallOpts)
}

// NextPoolId is a free data retrieval call binding the contract method 0x18e56131.
//
// Solidity: function nextPoolId() view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) NextPoolId() (*big.Int, error) {
	return _CrossStaking.Contract.NextPoolId(&_CrossStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossStaking *CrossStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossStaking *CrossStakingSession) Owner() (common.Address, error) {
	return _CrossStaking.Contract.Owner(&_CrossStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossStaking *CrossStakingCallerSession) Owner() (common.Address, error) {
	return _CrossStaking.Contract.Owner(&_CrossStaking.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossStaking *CrossStakingCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_CrossStaking *CrossStakingSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossStaking.Contract.PendingDefaultAdmin(&_CrossStaking.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossStaking *CrossStakingCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossStaking.Contract.PendingDefaultAdmin(&_CrossStaking.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossStaking *CrossStakingCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_CrossStaking *CrossStakingSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossStaking.Contract.PendingDefaultAdminDelay(&_CrossStaking.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossStaking *CrossStakingCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossStaking.Contract.PendingDefaultAdminDelay(&_CrossStaking.CallOpts)
}

// PoolAt is a free data retrieval call binding the contract method 0x155fff62.
//
// Solidity: function poolAt(uint256 index) view returns(uint256)
func (_CrossStaking *CrossStakingCaller) PoolAt(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "poolAt", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAt is a free data retrieval call binding the contract method 0x155fff62.
//
// Solidity: function poolAt(uint256 index) view returns(uint256)
func (_CrossStaking *CrossStakingSession) PoolAt(index *big.Int) (*big.Int, error) {
	return _CrossStaking.Contract.PoolAt(&_CrossStaking.CallOpts, index)
}

// PoolAt is a free data retrieval call binding the contract method 0x155fff62.
//
// Solidity: function poolAt(uint256 index) view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) PoolAt(index *big.Int) (*big.Int, error) {
	return _CrossStaking.Contract.PoolAt(&_CrossStaking.CallOpts, index)
}

// PoolByStakingTokenAt is a free data retrieval call binding the contract method 0x3b95352b.
//
// Solidity: function poolByStakingTokenAt(address stakingToken, uint256 index) view returns(uint256)
func (_CrossStaking *CrossStakingCaller) PoolByStakingTokenAt(opts *bind.CallOpts, stakingToken common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "poolByStakingTokenAt", stakingToken, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolByStakingTokenAt is a free data retrieval call binding the contract method 0x3b95352b.
//
// Solidity: function poolByStakingTokenAt(address stakingToken, uint256 index) view returns(uint256)
func (_CrossStaking *CrossStakingSession) PoolByStakingTokenAt(stakingToken common.Address, index *big.Int) (*big.Int, error) {
	return _CrossStaking.Contract.PoolByStakingTokenAt(&_CrossStaking.CallOpts, stakingToken, index)
}

// PoolByStakingTokenAt is a free data retrieval call binding the contract method 0x3b95352b.
//
// Solidity: function poolByStakingTokenAt(address stakingToken, uint256 index) view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) PoolByStakingTokenAt(stakingToken common.Address, index *big.Int) (*big.Int, error) {
	return _CrossStaking.Contract.PoolByStakingTokenAt(&_CrossStaking.CallOpts, stakingToken, index)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint256)
func (_CrossStaking *CrossStakingCaller) PoolIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "poolIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint256)
func (_CrossStaking *CrossStakingSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _CrossStaking.Contract.PoolIds(&_CrossStaking.CallOpts, arg0)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint256)
func (_CrossStaking *CrossStakingCallerSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _CrossStaking.Contract.PoolIds(&_CrossStaking.CallOpts, arg0)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_CrossStaking *CrossStakingCaller) PoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "poolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_CrossStaking *CrossStakingSession) PoolImplementation() (common.Address, error) {
	return _CrossStaking.Contract.PoolImplementation(&_CrossStaking.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_CrossStaking *CrossStakingCallerSession) PoolImplementation() (common.Address, error) {
	return _CrossStaking.Contract.PoolImplementation(&_CrossStaking.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, address stakingToken, uint256 createdAt)
func (_CrossStaking *CrossStakingCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	StakingToken common.Address
	CreatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "pools", arg0)

	outstruct := new(struct {
		PoolId       *big.Int
		Pool         common.Address
		StakingToken common.Address
		CreatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakingToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, address stakingToken, uint256 createdAt)
func (_CrossStaking *CrossStakingSession) Pools(arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	StakingToken common.Address
	CreatedAt    *big.Int
}, error) {
	return _CrossStaking.Contract.Pools(&_CrossStaking.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 poolId, address pool, address stakingToken, uint256 createdAt)
func (_CrossStaking *CrossStakingCallerSession) Pools(arg0 *big.Int) (struct {
	PoolId       *big.Int
	Pool         common.Address
	StakingToken common.Address
	CreatedAt    *big.Int
}, error) {
	return _CrossStaking.Contract.Pools(&_CrossStaking.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossStaking *CrossStakingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossStaking *CrossStakingSession) ProxiableUUID() ([32]byte, error) {
	return _CrossStaking.Contract.ProxiableUUID(&_CrossStaking.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossStaking *CrossStakingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossStaking.Contract.ProxiableUUID(&_CrossStaking.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_CrossStaking *CrossStakingCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_CrossStaking *CrossStakingSession) Router() (common.Address, error) {
	return _CrossStaking.Contract.Router(&_CrossStaking.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_CrossStaking *CrossStakingCallerSession) Router() (common.Address, error) {
	return _CrossStaking.Contract.Router(&_CrossStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossStaking *CrossStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossStaking *CrossStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossStaking.Contract.SupportsInterface(&_CrossStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossStaking *CrossStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossStaking.Contract.SupportsInterface(&_CrossStaking.CallOpts, interfaceId)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossStaking *CrossStakingCaller) Wcross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStaking.contract.Call(opts, &out, "wcross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossStaking *CrossStakingSession) Wcross() (common.Address, error) {
	return _CrossStaking.Contract.Wcross(&_CrossStaking.CallOpts)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossStaking *CrossStakingCallerSession) Wcross() (common.Address, error) {
	return _CrossStaking.Contract.Wcross(&_CrossStaking.CallOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossStaking *CrossStakingTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossStaking *CrossStakingSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossStaking.Contract.AcceptDefaultAdminTransfer(&_CrossStaking.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossStaking *CrossStakingTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossStaking.Contract.AcceptDefaultAdminTransfer(&_CrossStaking.TransactOpts)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa1635945.
//
// Solidity: function addRewardToken(uint256 poolId, address token) returns()
func (_CrossStaking *CrossStakingTransactor) AddRewardToken(opts *bind.TransactOpts, poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "addRewardToken", poolId, token)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa1635945.
//
// Solidity: function addRewardToken(uint256 poolId, address token) returns()
func (_CrossStaking *CrossStakingSession) AddRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.AddRewardToken(&_CrossStaking.TransactOpts, poolId, token)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0xa1635945.
//
// Solidity: function addRewardToken(uint256 poolId, address token) returns()
func (_CrossStaking *CrossStakingTransactorSession) AddRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.AddRewardToken(&_CrossStaking.TransactOpts, poolId, token)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossStaking *CrossStakingTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossStaking *CrossStakingSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.BeginDefaultAdminTransfer(&_CrossStaking.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossStaking *CrossStakingTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.BeginDefaultAdminTransfer(&_CrossStaking.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossStaking *CrossStakingTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossStaking *CrossStakingSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossStaking.Contract.CancelDefaultAdminTransfer(&_CrossStaking.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossStaking *CrossStakingTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossStaking.Contract.CancelDefaultAdminTransfer(&_CrossStaking.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossStaking *CrossStakingTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossStaking *CrossStakingSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.ChangeDefaultAdminDelay(&_CrossStaking.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossStaking *CrossStakingTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.ChangeDefaultAdminDelay(&_CrossStaking.TransactOpts, newDelay)
}

// CreatePool is a paid mutator transaction binding the contract method 0x12d36171.
//
// Solidity: function createPool(address stakingToken, uint256 minStakeAmount) returns(uint256 poolId, address pool)
func (_CrossStaking *CrossStakingTransactor) CreatePool(opts *bind.TransactOpts, stakingToken common.Address, minStakeAmount *big.Int) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "createPool", stakingToken, minStakeAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x12d36171.
//
// Solidity: function createPool(address stakingToken, uint256 minStakeAmount) returns(uint256 poolId, address pool)
func (_CrossStaking *CrossStakingSession) CreatePool(stakingToken common.Address, minStakeAmount *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.CreatePool(&_CrossStaking.TransactOpts, stakingToken, minStakeAmount)
}

// CreatePool is a paid mutator transaction binding the contract method 0x12d36171.
//
// Solidity: function createPool(address stakingToken, uint256 minStakeAmount) returns(uint256 poolId, address pool)
func (_CrossStaking *CrossStakingTransactorSession) CreatePool(stakingToken common.Address, minStakeAmount *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.CreatePool(&_CrossStaking.TransactOpts, stakingToken, minStakeAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.GrantRole(&_CrossStaking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.GrantRole(&_CrossStaking.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xce24af53.
//
// Solidity: function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) returns()
func (_CrossStaking *CrossStakingTransactor) Initialize(opts *bind.TransactOpts, _poolImplementation common.Address, _admin common.Address, _initialDelay *big.Int) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "initialize", _poolImplementation, _admin, _initialDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0xce24af53.
//
// Solidity: function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) returns()
func (_CrossStaking *CrossStakingSession) Initialize(_poolImplementation common.Address, _admin common.Address, _initialDelay *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.Initialize(&_CrossStaking.TransactOpts, _poolImplementation, _admin, _initialDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0xce24af53.
//
// Solidity: function initialize(address _poolImplementation, address _admin, uint48 _initialDelay) returns()
func (_CrossStaking *CrossStakingTransactorSession) Initialize(_poolImplementation common.Address, _admin common.Address, _initialDelay *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.Initialize(&_CrossStaking.TransactOpts, _poolImplementation, _admin, _initialDelay)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x35cc9cb4.
//
// Solidity: function removeRewardToken(uint256 poolId, address token) returns()
func (_CrossStaking *CrossStakingTransactor) RemoveRewardToken(opts *bind.TransactOpts, poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "removeRewardToken", poolId, token)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x35cc9cb4.
//
// Solidity: function removeRewardToken(uint256 poolId, address token) returns()
func (_CrossStaking *CrossStakingSession) RemoveRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.RemoveRewardToken(&_CrossStaking.TransactOpts, poolId, token)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x35cc9cb4.
//
// Solidity: function removeRewardToken(uint256 poolId, address token) returns()
func (_CrossStaking *CrossStakingTransactorSession) RemoveRewardToken(poolId *big.Int, token common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.RemoveRewardToken(&_CrossStaking.TransactOpts, poolId, token)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.RenounceRole(&_CrossStaking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.RenounceRole(&_CrossStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.RevokeRole(&_CrossStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossStaking *CrossStakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.RevokeRole(&_CrossStaking.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossStaking *CrossStakingTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossStaking *CrossStakingSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossStaking.Contract.RollbackDefaultAdminDelay(&_CrossStaking.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossStaking *CrossStakingTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossStaking.Contract.RollbackDefaultAdminDelay(&_CrossStaking.TransactOpts)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0xd6f74898.
//
// Solidity: function setPoolImplementation(address newImplementation) returns()
func (_CrossStaking *CrossStakingTransactor) SetPoolImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "setPoolImplementation", newImplementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0xd6f74898.
//
// Solidity: function setPoolImplementation(address newImplementation) returns()
func (_CrossStaking *CrossStakingSession) SetPoolImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.SetPoolImplementation(&_CrossStaking.TransactOpts, newImplementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0xd6f74898.
//
// Solidity: function setPoolImplementation(address newImplementation) returns()
func (_CrossStaking *CrossStakingTransactorSession) SetPoolImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.SetPoolImplementation(&_CrossStaking.TransactOpts, newImplementation)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0xb34c972e.
//
// Solidity: function setPoolStatus(uint256 poolId, uint8 status) returns()
func (_CrossStaking *CrossStakingTransactor) SetPoolStatus(opts *bind.TransactOpts, poolId *big.Int, status uint8) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "setPoolStatus", poolId, status)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0xb34c972e.
//
// Solidity: function setPoolStatus(uint256 poolId, uint8 status) returns()
func (_CrossStaking *CrossStakingSession) SetPoolStatus(poolId *big.Int, status uint8) (*types.Transaction, error) {
	return _CrossStaking.Contract.SetPoolStatus(&_CrossStaking.TransactOpts, poolId, status)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0xb34c972e.
//
// Solidity: function setPoolStatus(uint256 poolId, uint8 status) returns()
func (_CrossStaking *CrossStakingTransactorSession) SetPoolStatus(poolId *big.Int, status uint8) (*types.Transaction, error) {
	return _CrossStaking.Contract.SetPoolStatus(&_CrossStaking.TransactOpts, poolId, status)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_CrossStaking *CrossStakingTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_CrossStaking *CrossStakingSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.SetRouter(&_CrossStaking.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_CrossStaking *CrossStakingTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.SetRouter(&_CrossStaking.TransactOpts, _router)
}

// UpdateMinStakeAmount is a paid mutator transaction binding the contract method 0x1538af09.
//
// Solidity: function updateMinStakeAmount(uint256 poolId, uint256 amount) returns()
func (_CrossStaking *CrossStakingTransactor) UpdateMinStakeAmount(opts *bind.TransactOpts, poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "updateMinStakeAmount", poolId, amount)
}

// UpdateMinStakeAmount is a paid mutator transaction binding the contract method 0x1538af09.
//
// Solidity: function updateMinStakeAmount(uint256 poolId, uint256 amount) returns()
func (_CrossStaking *CrossStakingSession) UpdateMinStakeAmount(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.UpdateMinStakeAmount(&_CrossStaking.TransactOpts, poolId, amount)
}

// UpdateMinStakeAmount is a paid mutator transaction binding the contract method 0x1538af09.
//
// Solidity: function updateMinStakeAmount(uint256 poolId, uint256 amount) returns()
func (_CrossStaking *CrossStakingTransactorSession) UpdateMinStakeAmount(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossStaking.Contract.UpdateMinStakeAmount(&_CrossStaking.TransactOpts, poolId, amount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossStaking *CrossStakingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossStaking *CrossStakingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossStaking.Contract.UpgradeToAndCall(&_CrossStaking.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossStaking *CrossStakingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossStaking.Contract.UpgradeToAndCall(&_CrossStaking.TransactOpts, newImplementation, data)
}

// WithdrawFromPool is a paid mutator transaction binding the contract method 0xb9c17c4d.
//
// Solidity: function withdrawFromPool(uint256 poolId, address token, address to) returns()
func (_CrossStaking *CrossStakingTransactor) WithdrawFromPool(opts *bind.TransactOpts, poolId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossStaking.contract.Transact(opts, "withdrawFromPool", poolId, token, to)
}

// WithdrawFromPool is a paid mutator transaction binding the contract method 0xb9c17c4d.
//
// Solidity: function withdrawFromPool(uint256 poolId, address token, address to) returns()
func (_CrossStaking *CrossStakingSession) WithdrawFromPool(poolId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.WithdrawFromPool(&_CrossStaking.TransactOpts, poolId, token, to)
}

// WithdrawFromPool is a paid mutator transaction binding the contract method 0xb9c17c4d.
//
// Solidity: function withdrawFromPool(uint256 poolId, address token, address to) returns()
func (_CrossStaking *CrossStakingTransactorSession) WithdrawFromPool(poolId *big.Int, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossStaking.Contract.WithdrawFromPool(&_CrossStaking.TransactOpts, poolId, token, to)
}

// CrossStakingDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the CrossStaking contract.
type CrossStakingDefaultAdminDelayChangeCanceledIterator struct {
	Event *CrossStakingDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *CrossStakingDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingDefaultAdminDelayChangeCanceled)
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
		it.Event = new(CrossStakingDefaultAdminDelayChangeCanceled)
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
func (it *CrossStakingDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the CrossStaking contract.
type CrossStakingDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossStaking *CrossStakingFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*CrossStakingDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossStakingDefaultAdminDelayChangeCanceledIterator{contract: _CrossStaking.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossStaking *CrossStakingFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *CrossStakingDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingDefaultAdminDelayChangeCanceled)
				if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*CrossStakingDefaultAdminDelayChangeCanceled, error) {
	event := new(CrossStakingDefaultAdminDelayChangeCanceled)
	if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the CrossStaking contract.
type CrossStakingDefaultAdminDelayChangeScheduledIterator struct {
	Event *CrossStakingDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *CrossStakingDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingDefaultAdminDelayChangeScheduled)
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
		it.Event = new(CrossStakingDefaultAdminDelayChangeScheduled)
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
func (it *CrossStakingDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the CrossStaking contract.
type CrossStakingDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossStaking *CrossStakingFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*CrossStakingDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &CrossStakingDefaultAdminDelayChangeScheduledIterator{contract: _CrossStaking.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossStaking *CrossStakingFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *CrossStakingDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingDefaultAdminDelayChangeScheduled)
				if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*CrossStakingDefaultAdminDelayChangeScheduled, error) {
	event := new(CrossStakingDefaultAdminDelayChangeScheduled)
	if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the CrossStaking contract.
type CrossStakingDefaultAdminTransferCanceledIterator struct {
	Event *CrossStakingDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *CrossStakingDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingDefaultAdminTransferCanceled)
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
		it.Event = new(CrossStakingDefaultAdminTransferCanceled)
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
func (it *CrossStakingDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the CrossStaking contract.
type CrossStakingDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossStaking *CrossStakingFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*CrossStakingDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossStakingDefaultAdminTransferCanceledIterator{contract: _CrossStaking.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossStaking *CrossStakingFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *CrossStakingDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingDefaultAdminTransferCanceled)
				if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*CrossStakingDefaultAdminTransferCanceled, error) {
	event := new(CrossStakingDefaultAdminTransferCanceled)
	if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the CrossStaking contract.
type CrossStakingDefaultAdminTransferScheduledIterator struct {
	Event *CrossStakingDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *CrossStakingDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingDefaultAdminTransferScheduled)
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
		it.Event = new(CrossStakingDefaultAdminTransferScheduled)
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
func (it *CrossStakingDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the CrossStaking contract.
type CrossStakingDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossStaking *CrossStakingFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*CrossStakingDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingDefaultAdminTransferScheduledIterator{contract: _CrossStaking.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossStaking *CrossStakingFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *CrossStakingDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingDefaultAdminTransferScheduled)
				if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*CrossStakingDefaultAdminTransferScheduled, error) {
	event := new(CrossStakingDefaultAdminTransferScheduled)
	if err := _CrossStaking.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossStaking contract.
type CrossStakingInitializedIterator struct {
	Event *CrossStakingInitialized // Event containing the contract specifics and raw log

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
func (it *CrossStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingInitialized)
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
		it.Event = new(CrossStakingInitialized)
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
func (it *CrossStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingInitialized represents a Initialized event raised by the CrossStaking contract.
type CrossStakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossStaking *CrossStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrossStakingInitializedIterator, error) {

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossStakingInitializedIterator{contract: _CrossStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossStaking *CrossStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingInitialized)
				if err := _CrossStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseInitialized(log types.Log) (*CrossStakingInitialized, error) {
	event := new(CrossStakingInitialized)
	if err := _CrossStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the CrossStaking contract.
type CrossStakingPoolCreatedIterator struct {
	Event *CrossStakingPoolCreated // Event containing the contract specifics and raw log

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
func (it *CrossStakingPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingPoolCreated)
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
		it.Event = new(CrossStakingPoolCreated)
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
func (it *CrossStakingPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingPoolCreated represents a PoolCreated event raised by the CrossStaking contract.
type CrossStakingPoolCreated struct {
	PoolId       *big.Int
	PoolAddress  common.Address
	StakingToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a6.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed stakingToken)
func (_CrossStaking *CrossStakingFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int, poolAddress []common.Address, stakingToken []common.Address) (*CrossStakingPoolCreatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var stakingTokenRule []interface{}
	for _, stakingTokenItem := range stakingToken {
		stakingTokenRule = append(stakingTokenRule, stakingTokenItem)
	}

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "PoolCreated", poolIdRule, poolAddressRule, stakingTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingPoolCreatedIterator{contract: _CrossStaking.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x1a7a1d16c3ec9167827a7be3534be26288720a0bdd3de56d290f415db3d3e0a6.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed stakingToken)
func (_CrossStaking *CrossStakingFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *CrossStakingPoolCreated, poolId []*big.Int, poolAddress []common.Address, stakingToken []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var stakingTokenRule []interface{}
	for _, stakingTokenItem := range stakingToken {
		stakingTokenRule = append(stakingTokenRule, stakingTokenItem)
	}

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "PoolCreated", poolIdRule, poolAddressRule, stakingTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingPoolCreated)
				if err := _CrossStaking.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed poolAddress, address indexed stakingToken)
func (_CrossStaking *CrossStakingFilterer) ParsePoolCreated(log types.Log) (*CrossStakingPoolCreated, error) {
	event := new(CrossStakingPoolCreated)
	if err := _CrossStaking.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingPoolImplementationSetIterator is returned from FilterPoolImplementationSet and is used to iterate over the raw logs and unpacked data for PoolImplementationSet events raised by the CrossStaking contract.
type CrossStakingPoolImplementationSetIterator struct {
	Event *CrossStakingPoolImplementationSet // Event containing the contract specifics and raw log

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
func (it *CrossStakingPoolImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingPoolImplementationSet)
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
		it.Event = new(CrossStakingPoolImplementationSet)
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
func (it *CrossStakingPoolImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingPoolImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingPoolImplementationSet represents a PoolImplementationSet event raised by the CrossStaking contract.
type CrossStakingPoolImplementationSet struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolImplementationSet is a free log retrieval operation binding the contract event 0xdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957.
//
// Solidity: event PoolImplementationSet(address indexed implementation)
func (_CrossStaking *CrossStakingFilterer) FilterPoolImplementationSet(opts *bind.FilterOpts, implementation []common.Address) (*CrossStakingPoolImplementationSetIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "PoolImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingPoolImplementationSetIterator{contract: _CrossStaking.contract, event: "PoolImplementationSet", logs: logs, sub: sub}, nil
}

// WatchPoolImplementationSet is a free log subscription operation binding the contract event 0xdd6f7e9de2078ecfceba0b29adf9a7f2d9a97cc573945494fddbdf223dde8957.
//
// Solidity: event PoolImplementationSet(address indexed implementation)
func (_CrossStaking *CrossStakingFilterer) WatchPoolImplementationSet(opts *bind.WatchOpts, sink chan<- *CrossStakingPoolImplementationSet, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "PoolImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingPoolImplementationSet)
				if err := _CrossStaking.contract.UnpackLog(event, "PoolImplementationSet", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParsePoolImplementationSet(log types.Log) (*CrossStakingPoolImplementationSet, error) {
	event := new(CrossStakingPoolImplementationSet)
	if err := _CrossStaking.contract.UnpackLog(event, "PoolImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossStaking contract.
type CrossStakingRoleAdminChangedIterator struct {
	Event *CrossStakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossStakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRoleAdminChanged)
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
		it.Event = new(CrossStakingRoleAdminChanged)
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
func (it *CrossStakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRoleAdminChanged represents a RoleAdminChanged event raised by the CrossStaking contract.
type CrossStakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossStaking *CrossStakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossStakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRoleAdminChangedIterator{contract: _CrossStaking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossStaking *CrossStakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossStakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRoleAdminChanged)
				if err := _CrossStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseRoleAdminChanged(log types.Log) (*CrossStakingRoleAdminChanged, error) {
	event := new(CrossStakingRoleAdminChanged)
	if err := _CrossStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossStaking contract.
type CrossStakingRoleGrantedIterator struct {
	Event *CrossStakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossStakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRoleGranted)
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
		it.Event = new(CrossStakingRoleGranted)
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
func (it *CrossStakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRoleGranted represents a RoleGranted event raised by the CrossStaking contract.
type CrossStakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossStaking *CrossStakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossStakingRoleGrantedIterator, error) {

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

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRoleGrantedIterator{contract: _CrossStaking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossStaking *CrossStakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossStakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRoleGranted)
				if err := _CrossStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseRoleGranted(log types.Log) (*CrossStakingRoleGranted, error) {
	event := new(CrossStakingRoleGranted)
	if err := _CrossStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossStaking contract.
type CrossStakingRoleRevokedIterator struct {
	Event *CrossStakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossStakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRoleRevoked)
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
		it.Event = new(CrossStakingRoleRevoked)
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
func (it *CrossStakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRoleRevoked represents a RoleRevoked event raised by the CrossStaking contract.
type CrossStakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossStaking *CrossStakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossStakingRoleRevokedIterator, error) {

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

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRoleRevokedIterator{contract: _CrossStaking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossStaking *CrossStakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossStakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRoleRevoked)
				if err := _CrossStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseRoleRevoked(log types.Log) (*CrossStakingRoleRevoked, error) {
	event := new(CrossStakingRoleRevoked)
	if err := _CrossStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRouterSetIterator is returned from FilterRouterSet and is used to iterate over the raw logs and unpacked data for RouterSet events raised by the CrossStaking contract.
type CrossStakingRouterSetIterator struct {
	Event *CrossStakingRouterSet // Event containing the contract specifics and raw log

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
func (it *CrossStakingRouterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRouterSet)
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
		it.Event = new(CrossStakingRouterSet)
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
func (it *CrossStakingRouterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRouterSet represents a RouterSet event raised by the CrossStaking contract.
type CrossStakingRouterSet struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterSet is a free log retrieval operation binding the contract event 0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15.
//
// Solidity: event RouterSet(address indexed router)
func (_CrossStaking *CrossStakingFilterer) FilterRouterSet(opts *bind.FilterOpts, router []common.Address) (*CrossStakingRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "RouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterSetIterator{contract: _CrossStaking.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

// WatchRouterSet is a free log subscription operation binding the contract event 0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15.
//
// Solidity: event RouterSet(address indexed router)
func (_CrossStaking *CrossStakingFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *CrossStakingRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "RouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRouterSet)
				if err := _CrossStaking.contract.UnpackLog(event, "RouterSet", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseRouterSet(log types.Log) (*CrossStakingRouterSet, error) {
	event := new(CrossStakingRouterSet)
	if err := _CrossStaking.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossStaking contract.
type CrossStakingUpgradedIterator struct {
	Event *CrossStakingUpgraded // Event containing the contract specifics and raw log

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
func (it *CrossStakingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingUpgraded)
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
		it.Event = new(CrossStakingUpgraded)
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
func (it *CrossStakingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingUpgraded represents a Upgraded event raised by the CrossStaking contract.
type CrossStakingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossStaking *CrossStakingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossStakingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingUpgradedIterator{contract: _CrossStaking.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossStaking *CrossStakingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossStakingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingUpgraded)
				if err := _CrossStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossStaking *CrossStakingFilterer) ParseUpgraded(log types.Log) (*CrossStakingUpgraded, error) {
	event := new(CrossStakingUpgraded)
	if err := _CrossStaking.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingWithdrawnFromPoolIterator is returned from FilterWithdrawnFromPool and is used to iterate over the raw logs and unpacked data for WithdrawnFromPool events raised by the CrossStaking contract.
type CrossStakingWithdrawnFromPoolIterator struct {
	Event *CrossStakingWithdrawnFromPool // Event containing the contract specifics and raw log

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
func (it *CrossStakingWithdrawnFromPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingWithdrawnFromPool)
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
		it.Event = new(CrossStakingWithdrawnFromPool)
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
func (it *CrossStakingWithdrawnFromPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingWithdrawnFromPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingWithdrawnFromPool represents a WithdrawnFromPool event raised by the CrossStaking contract.
type CrossStakingWithdrawnFromPool struct {
	PoolId *big.Int
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFromPool is a free log retrieval operation binding the contract event 0x408708e2d99a8cfa59be1466864d74f6ddddc62b5ddc3bb11f1c3cc3ce9ed65f.
//
// Solidity: event WithdrawnFromPool(uint256 indexed poolId, address indexed token, address indexed to, uint256 amount)
func (_CrossStaking *CrossStakingFilterer) FilterWithdrawnFromPool(opts *bind.FilterOpts, poolId []*big.Int, token []common.Address, to []common.Address) (*CrossStakingWithdrawnFromPoolIterator, error) {

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

	logs, sub, err := _CrossStaking.contract.FilterLogs(opts, "WithdrawnFromPool", poolIdRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingWithdrawnFromPoolIterator{contract: _CrossStaking.contract, event: "WithdrawnFromPool", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFromPool is a free log subscription operation binding the contract event 0x408708e2d99a8cfa59be1466864d74f6ddddc62b5ddc3bb11f1c3cc3ce9ed65f.
//
// Solidity: event WithdrawnFromPool(uint256 indexed poolId, address indexed token, address indexed to, uint256 amount)
func (_CrossStaking *CrossStakingFilterer) WatchWithdrawnFromPool(opts *bind.WatchOpts, sink chan<- *CrossStakingWithdrawnFromPool, poolId []*big.Int, token []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossStaking.contract.WatchLogs(opts, "WithdrawnFromPool", poolIdRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingWithdrawnFromPool)
				if err := _CrossStaking.contract.UnpackLog(event, "WithdrawnFromPool", log); err != nil {
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

// ParseWithdrawnFromPool is a log parse operation binding the contract event 0x408708e2d99a8cfa59be1466864d74f6ddddc62b5ddc3bb11f1c3cc3ce9ed65f.
//
// Solidity: event WithdrawnFromPool(uint256 indexed poolId, address indexed token, address indexed to, uint256 amount)
func (_CrossStaking *CrossStakingFilterer) ParseWithdrawnFromPool(log types.Log) (*CrossStakingWithdrawnFromPool, error) {
	event := new(CrossStakingWithdrawnFromPool)
	if err := _CrossStaking.contract.UnpackLog(event, "WithdrawnFromPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
