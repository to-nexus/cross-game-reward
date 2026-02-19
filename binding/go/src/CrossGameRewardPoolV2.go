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

// ICrossGameRewardPoolRewardToken is an auto generated low-level Go binding around an user-defined struct.

// ICrossGameRewardPoolV2Round is an auto generated low-level Go binding around an user-defined struct.
type ICrossGameRewardPoolV2Round struct {
	RoundId           *big.Int
	TotalReward       *big.Int
	StartBlock        *big.Int
	EndBlock          *big.Int
	RewardPerBlock    *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
	IsCancelled       bool
}

// CrossGameRewardPoolV2MetaData contains all meta data concerning the CrossGameRewardPoolV2 contract.
var CrossGameRewardPoolV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPONSOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"cancelRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimRewardFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimRewardsFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"createRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossGameReward\",\"outputs\":[{\"internalType\":\"contractICrossGameReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRoundCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRoundIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"}],\"internalType\":\"structICrossGameRewardPoolV2.Round[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getReclaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemovedRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRemovedTokenRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRewardToken\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reclaimableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRemoved\",\"type\":\"bool\"}],\"internalType\":\"structICrossGameRewardPool.RewardToken\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"}],\"internalType\":\"structICrossGameRewardPoolV2.Round\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalAccRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRemovedRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolStatus\",\"outputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reclaimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reclaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removedRewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"rewardTokenAt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MinDepositAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"PoolStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"RoundCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"RoundCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"CGRP2BelowMinimumDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2CanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2CanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"currentStatus\",\"type\":\"uint8\"}],\"name\":\"CGRP2DepositNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"CGRP2InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provided\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"CGRP2InvalidRewardToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"CGRP2InvalidStartBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"CGRP2NoDepositFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2NoReclaimableAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2NotAllowedInCurrentState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2OnlyRewardRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2OnlyRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2RewardIsDepositToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CGRP2RewardPerBlockZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"CGRP2RoundAlreadyCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"CGRP2RoundAlreadyStarted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"CGRP2RoundNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"c2d79444": "SPONSOR_ROLE()",
		"ad3cb1cc": "UPGRADE_INTERFACE_VERSION()",
		"cefc1429": "acceptDefaultAdminTransfer()",
		"1c03e6cc": "addRewardToken(address)",
		"27e235e3": "balances(address)",
		"634e93da": "beginDefaultAdminTransfer(address)",
		"d602b9fd": "cancelDefaultAdminTransfer()",
		"7e07ab09": "cancelRound(uint256)",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"d279c191": "claimReward(address)",
		"35c30fda": "claimRewardFor(address,address)",
		"372500ab": "claimRewards()",
		"1ac6d19d": "claimRewardsFor(address)",
		"1efed5f7": "createRound(uint256,uint256,uint256)",
		"f4e24740": "crossGameReward()",
		"84ef8ffc": "defaultAdmin()",
		"cc8463c8": "defaultAdminDelay()",
		"022d63fb": "defaultAdminDelayIncreaseWait()",
		"b6b55f25": "deposit(uint256)",
		"2f4f21e2": "depositFor(address,uint256)",
		"c89039c5": "depositToken()",
		"78ad8c7d": "getActiveRoundCount()",
		"6fb7a4e8": "getActiveRoundIds()",
		"7d984d5f": "getActiveRounds()",
		"35c21d5d": "getReclaimableAmount(address)",
		"9b80c3f2": "getRemovedRewardTokens()",
		"1af8acec": "getRemovedTokenRewards(address)",
		"77078872": "getRewardToken(address)",
		"c4f59f9b": "getRewardTokens()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"8f1327c0": "getRound(uint256)",
		"2dbea37b": "globalAccRewardPerShare()",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"1794bb3c": "initialize(address,address,uint256)",
		"91cf6d3e": "initializedAt()",
		"f665336e": "isRemovedRewardToken(address)",
		"b5fd73f8": "isRewardToken(address)",
		"645006ca": "minDepositAmount()",
		"4002eda6": "nextRoundId()",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"cf6eefb7": "pendingDefaultAdmin()",
		"a1eda53c": "pendingDefaultAdminDelay()",
		"9ced7e76": "pendingReward(address,address)",
		"31d7a262": "pendingRewards(address)",
		"f0228692": "poolStatus()",
		"52d1902d": "proxiableUUID()",
		"4d1cd014": "reclaimTokens(address,address)",
		"fd8bdc68": "reclaimableAmount()",
		"3d509c97": "removeRewardToken(address)",
		"35482379": "removedRewardTokenCount()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"f7c618c1": "rewardToken()",
		"79f5ecb7": "rewardTokenAt(uint256)",
		"abb06b95": "rewardTokenCount()",
		"0aa6220b": "rollbackDefaultAdminDelay()",
		"6d7c49a2": "setPoolStatus(uint8)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"ff50abdc": "totalDeposited()",
		"84780205": "updateMinDepositAmount(uint256)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"a980356a": "userRewards(address,address)",
		"2e1a7d4d": "withdraw(uint256)",
		"db518db2": "withdrawFor(address,uint256)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161518d6100f95f395f81816135b5015281816135de01526137ae015261518d5ff3fe608060405260043610610401575f3560e01c806379f5ecb71161021d578063b6b55f2511610122578063d547741f116100b7578063f4e2474011610087578063f7c618c11161006d578063f7c618c114610d66578063fd8bdc6814610d85578063ff50abdc14610d9a575f5ffd5b8063f4e2474014610d28578063f665336e14610d47575f5ffd5b8063d547741f14610cb0578063d602b9fd14610ccf578063db518db214610ce3578063f022869214610d02575f5ffd5b8063cc8463c8116100f2578063cc8463c814610bff578063cefc142914610c13578063cf6eefb714610c27578063d279c19114610c91575f5ffd5b8063b6b55f2514610b7a578063c2d7944414610b99578063c4f59f9b14610bcc578063c89039c514610be0575f5ffd5b806391d14854116101b2578063a217fddf11610182578063abb06b9511610168578063abb06b9514610ae4578063ad3cb1cc14610af7578063b5fd73f814610b4c575f5ffd5b8063a217fddf14610a9d578063a980356a14610ab0575f5ffd5b806391d14854146109c05780639b80c3f214610a235780639ced7e7614610a4b578063a1eda53c14610a6a575f5ffd5b806384ef8ffc116101ed57806384ef8ffc146109305780638da5cb5b1461096c5780638f1327c01461098057806391cf6d3e146109ac575f5ffd5b806379f5ecb71461089a5780637d984d5f146108d15780637e07ab09146108f25780638478020514610911575f5ffd5b806335c21d5d1161032357806352d1902d116102b8578063649a5ec7116102885780636fb7a4e81161026e5780636fb7a4e8146107f0578063770788721461081157806378ad8c7d14610886575f5ffd5b8063649a5ec7146107b25780636d7c49a2146107d1575f5ffd5b806352d1902d146107345780635c975abb14610748578063634e93da1461077e578063645006ca1461079d575f5ffd5b80633d509c97116102f35780633d509c97146106ce5780634002eda6146106ed5780634d1cd014146107025780634f1ef28614610721575f5ffd5b806335c21d5d1461065d57806335c30fda1461067c57806336568abe1461069b578063372500ab146106ba575f5ffd5b8063248a9ca3116103995780632f2ff15d116103695780632f2ff15d146105ee5780632f4f21e21461060d57806331d7a2621461062c578063354823791461064b575f5ffd5b8063248a9ca31461054257806327e235e31461058f5780632dbea37b146105ba5780632e1a7d4d146105cf575f5ffd5b80631ac6d19d116103d45780631ac6d19d146104965780631af8acec146104b55780631c03e6cc146104f65780631efed5f714610515575f5ffd5b806301ffc9a714610405578063022d63fb146104395780630aa6220b146104615780631794bb3c14610477575b5f5ffd5b348015610410575f5ffd5b5061042461041f366004614a2d565b610daf565b60405190151581526020015b60405180910390f35b348015610444575f5ffd5b50620697805b60405165ffffffffffff9091168152602001610430565b34801561046c575f5ffd5b50610475610e0a565b005b348015610482575f5ffd5b50610475610491366004614a80565b610e1f565b3480156104a1575f5ffd5b506104756104b0366004614abe565b61112e565b3480156104c0575f5ffd5b506104e86104cf366004614abe565b50604080515f8082526020820190815281830190925291565b604051610430929190614b4c565b348015610501575f5ffd5b50610475610510366004614abe565b6111a8565b348015610520575f5ffd5b5061053461052f366004614b70565b61128e565b604051908152602001610430565b34801561054d575f5ffd5b5061053461055c366004614b99565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561059a575f5ffd5b506105346105a9366004614abe565b60056020525f908152604090205481565b3480156105c5575f5ffd5b50610534600e5481565b3480156105da575f5ffd5b506104756105e9366004614b99565b611547565b3480156105f9575f5ffd5b50610475610608366004614bb0565b6115b2565b348015610618575f5ffd5b50610475610627366004614bde565b6115f7565b348015610637575f5ffd5b506104e8610646366004614abe565b61167b565b348015610656575f5ffd5b505f610534565b348015610668575f5ffd5b50610534610677366004614abe565b61172b565b348015610687575f5ffd5b50610475610696366004614c08565b611751565b3480156106a6575f5ffd5b506104756106b5366004614bb0565b611823565b3480156106c5575f5ffd5b50610475611970565b3480156106d9575f5ffd5b506104756106e8366004614abe565b6119e3565b3480156106f8575f5ffd5b5061053460085481565b34801561070d575f5ffd5b5061047561071c366004614c08565b611aaf565b61047561072f366004614c61565b611c4e565b34801561073f575f5ffd5b50610534611c69565b348015610753575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610424565b348015610789575f5ffd5b50610475610798366004614abe565b611c97565b3480156107a8575f5ffd5b5061053460045481565b3480156107bd575f5ffd5b506104756107cc366004614d64565b611caa565b3480156107dc575f5ffd5b506104756107eb366004614d89565b611cbd565b3480156107fb575f5ffd5b50610804611e98565b6040516104309190614da7565b34801561081c575f5ffd5b5061083061082b366004614abe565b611ea9565b60405161043091905f60c0820190506001600160a01b0383511682526020830151602083015260408301516040830152606083015160608301526080830151608083015260a0830151151560a083015292915050565b348015610891575f5ffd5b50610534612015565b3480156108a5575f5ffd5b506108b96108b4366004614b99565b612020565b6040516001600160a01b039091168152602001610430565b3480156108dc575f5ffd5b506108e561209a565b6040516104309190614db9565b3480156108fd575f5ffd5b5061047561090c366004614b99565b6121ec565b34801561091c575f5ffd5b5061047561092b366004614b99565b612375565b34801561093b575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b03166108b9565b348015610977575f5ffd5b506108b9612433565b34801561098b575f5ffd5b5061099f61099a366004614b99565b61243c565b6040516104309190614e52565b3480156109b7575f5ffd5b506105345f5481565b3480156109cb575f5ffd5b506104246109da366004614bb0565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610a2e575f5ffd5b50604080515f8152602081019091525b6040516104309190614eac565b348015610a56575f5ffd5b50610534610a65366004614c08565b61253a565b348015610a75575f5ffd5b50610a7e612544565b6040805165ffffffffffff938416815292909116602083015201610430565b348015610aa8575f5ffd5b506105345f81565b348015610abb575f5ffd5b50610acf610aca366004614c08565b612601565b60408051928352602083019190915201610430565b348015610aef575f5ffd5b506001610534565b348015610b02575f5ffd5b50610b3f6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516104309190614ebe565b348015610b57575f5ffd5b50610424610b66366004614abe565b6002546001600160a01b0391821691161490565b348015610b85575f5ffd5b50610475610b94366004614b99565b612668565b348015610ba4575f5ffd5b506105347f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a81565b348015610bd7575f5ffd5b50610a3e6126db565b348015610beb575f5ffd5b506001546108b9906001600160a01b031681565b348015610c0a575f5ffd5b5061044a61273c565b348015610c1e575f5ffd5b5061047561281c565b348015610c32575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610430565b348015610c9c575f5ffd5b50610475610cab366004614abe565b61288a565b348015610cbb575f5ffd5b50610475610cca366004614bb0565b612953565b348015610cda575f5ffd5b50610475612994565b348015610cee575f5ffd5b50610475610cfd366004614bde565b6129a6565b348015610d0d575f5ffd5b50600754610d1b9060ff1681565b6040516104309190614f77565b348015610d33575f5ffd5b506003546108b9906001600160a01b031681565b348015610d52575f5ffd5b50610424610d61366004614abe565b505f90565b348015610d71575f5ffd5b506002546108b9906001600160a01b031681565b348015610d90575f5ffd5b5061053460115481565b348015610da5575f5ffd5b5061053460065481565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610e045750610e0482612a1a565b92915050565b5f610e1481612ab0565b610e1c612aba565b50565b5f610e28612ac4565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610e545750825b90505f8267ffffffffffffffff166001148015610e705750303b155b905081158015610e7e575080155b15610eb5576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610f165784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816610f56576040517fc5e57bc700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716610f96576040517fc5e57bc700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316886001600160a01b031603610fe1576040517fa1ecd1d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f861161101a576040517f7cd1a25d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001633908117909155611052905f90612aec565b61105a612afe565b611062612afe565b61106a612afe565b435f55600180546001600160a01b038a81167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617835560028054918b169190921617905560048790556007805460ff1916905560085583156111245784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b611136612b06565b61113e612b8c565b600260075460ff16600281111561115757611157614f11565b0361118e576040517f5c8bfea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61119781612be8565b6111a081612ce6565b610e1c612d7c565b6003546001600160a01b031633146111ec576040517f3c6f5f9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b03828116911614610e1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f43616e6e6f742061646420646966666572656e742072657761726420746f6b6560448201527f6e0000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b5f611297612b06565b61129f612b8c565b7f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a6112c981612ab0565b5f8511611302576040517f7cd1a25d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8343808211611346576040517f8f06300600000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611285565b50505f8311611381576040517f31a1cecb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61138c8487614fb2565b90505f81116113c7576040517f4f99577a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6113d28583614fea565b90506113dc612da6565b60088054905f6113eb83615001565b9091555093505f6113fc8688615038565b90506040518061010001604052808681526020018381526020018881526020018281526020018481526020018881526020015f81526020015f151581525060095f8781526020019081526020015f205f820151815f01556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff0219169083151502179055509050506114c585600a612fda90919063ffffffff16565b506114d1600c86612fda565b506002546114ea906001600160a01b0316333085612fe5565b60408051838152602081018990529081018290526060810184905285907ff4c8810c202a3e7371b142615e4811842e135aa5f919299d1cb4050710f7b85e9060800160405180910390a250505050611540612d7c565b9392505050565b61154f612b06565b611557612b8c565b600260075460ff16600281111561157057611570614f11565b036115a7576040517f5c8bfea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111a0333383613061565b816115e9576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115f3828261323c565b5050565b6115ff612b06565b611607612b8c565b5f60075460ff16600281111561161f5761161f614f11565b60075460ff16911461165e576040517fded8e4710000000000000000000000000000000000000000000000000000000081526004016112859190614f77565b5061166882612be8565b61167333838361327f565b6115f3612d7c565b604080516001808252818301909252606091829190602080830190803683375050604080516001808252818301909252929450905060208083019080368337505060025484519293506001600160a01b0316918491505f906116df576116df61504b565b60200260200101906001600160a01b031690816001600160a01b031681525050611708836133b6565b815f8151811061171a5761171a61504b565b602002602001018181525050915091565b6002545f906001600160a01b0383811691161461174957505f919050565b505060115490565b611759612b06565b611761612b8c565b600260075460ff16600281111561177a5761177a614f11565b036117b1576040517f5c8bfea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025481906001600160a01b03908116908216811461180f576040517fc90340c60000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015291166024820152604401611285565b505061181a82612be8565b61167382612ce6565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561187e57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b15611961577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806118e4575065ffffffffffff8116155b806118f757504265ffffffffffff821610155b15611938576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401611285565b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b61196b838361352d565b505050565b611978612b06565b611980612b8c565b600260075460ff16600281111561199957611999614f11565b036119d0576040517f5c8bfea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119d933612ce6565b6119e1612d7c565b565b6003546001600160a01b03163314611a27576040517f3c6f5f9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f563220646f6573206e6f7420737570706f72742072656d6f76696e672072657760448201527f61726420746f6b656e73000000000000000000000000000000000000000000006064820152608401611285565b6003546001600160a01b03163314611af3576040517f3c6f5f9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b03838116911614611b6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420746f6b656e000000000000000000000000000000000000006044820152606401611285565b5f60115411611ba5576040517fcff62bcc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116611be5576040517fc5e57bc700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601180545f909155600254611c04906001600160a01b03168383613579565b6002546040518281526001600160a01b038481169216907f6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464906020015b60405180910390a3505050565b611c566135aa565b611c5f8261367a565b6115f38282613684565b5f611c726137a3565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f611ca181612ab0565b6115f382613805565b5f611cb481612ab0565b6115f382613877565b6003546001600160a01b03163314611d01576040517f3c6f5f9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460ff16816002811115611d1957611d19614f11565b816002811115611d2b57611d2b614f11565b03611d92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f506f6f6c2073746174757320756e6368616e67656400000000000000000000006044820152606401611285565b6007805483919060ff19166001836002811115611db157611db1614f11565b02179055506002826002811115611dca57611dca614f11565b148015611df957507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16155b15611e0b57611e066138df565b611e5b565b6002826002811115611e1f57611e1f614f11565b14158015611e4e57507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff165b15611e5b57611e5b613954565b7fc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb59668183604051611e8c929190615078565b60405180910390a15050565b6060611ea4600a6139ac565b905090565b611ee76040518060c001604052805f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b6002546001600160a01b03838116911614611f5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420746f6b656e000000000000000000000000000000000000006044820152606401611285565b6040805160c0810182526002546001600160a01b0316808252600e54602083015282517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291928301916370a0823190602401602060405180830381865afa158015611fd3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ff79190615093565b815260115460208201525f6040820181905260609091015292915050565b5f611ea4600a6139b8565b5f8115612089576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420696e646578000000000000000000000000000000000000006044820152606401611285565b50506002546001600160a01b031690565b60605f6120a7600a6139b8565b90505f8167ffffffffffffffff8111156120c3576120c3614c34565b60405190808252806020026020018201604052801561213657816020015b6121236040518061010001604052805f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b8152602001906001900390816120e15790505b5090505f5b828110156121e55760095f612151600a846139c1565b815260208082019290925260409081015f20815161010081018352815481526001820154938101939093526002810154918301919091526003810154606083015260048101546080830152600581015460a0830152600681015460c08301526007015460ff16151560e082015282518390839081106121d2576121d261504b565b602090810291909101015260010161213b565b5092915050565b6121f4612b06565b6121fc612b8c565b7f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a61222681612ab0565b5f82815260096020526040902080548390612270576040517fc7957ca900000000000000000000000000000000000000000000000000000000815260040161128591815260200190565b506007810154839060ff16156122b5576040517f2acf2e9700000000000000000000000000000000000000000000000000000000815260040161128591815260200190565b508060020154431083906122f8576040517f7a00baae00000000000000000000000000000000000000000000000000000000815260040161128591815260200190565b5060078101805460ff19166001179055612313600a846139cc565b506001810154600254612330906001600160a01b03163383613579565b837f392fcf1e3627793dc153feb861f66451c925fa12c027044233166cd28f481d858260405161236291815260200190565b60405180910390a2505050610e1c612d7c565b6003546001600160a01b031633146123b9576040517f3c6f5f9a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81116123f2576040517f7cd1a25d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045460408051918252602082018390527f5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329910160405180910390a1600455565b5f611ea46139d7565b61247e6040518061010001604052805f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b5f8281526009602052604090205482906124c7576040517fc7957ca900000000000000000000000000000000000000000000000000000000815260040161128591815260200190565b50505f90815260096020908152604091829020825161010081018452815481526001820154928101929092526002810154928201929092526003820154606082015260048201546080820152600582015460a0820152600682015460c082015260079091015460ff16151560e082015290565b5f611540836133b6565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840081158015906125c657504265ffffffffffff831610155b6125d1575f5f6125f8565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b6001600160a01b0382165f9081526005602052604081205481908015612647576001600160a01b0385165f908152600f6020526040902054612644908290614fb2565b92505b50506001600160a01b039092165f9081526010602052604090205491929050565b612670612b06565b612678612b8c565b5f60075460ff16600281111561269057612690614f11565b60075460ff1691146126cf576040517fded8e4710000000000000000000000000000000000000000000000000000000081526004016112859190614f77565b506111a033338361327f565b6040805160018082528183019092526060915f919060208083019080368337505060025482519293506001600160a01b0316918391505f9061271f5761271f61504b565b6001600160a01b0390921660209283029190910190910152919050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680158015906127be57504265ffffffffffff8216105b6127ef5781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16612815565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114612882576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401611285565b610e1c613a09565b612892612b06565b61289a612b8c565b600260075460ff1660028111156128b3576128b3614f11565b036128ea576040517f5c8bfea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025481906001600160a01b039081169082168114612948576040517fc90340c60000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015291166024820152604401611285565b50506111a033612ce6565b8161298a576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115f38282613b1e565b5f61299e81612ab0565b610e1c613b61565b6129ae612b06565b6129b6612b8c565b600260075460ff1660028111156129cf576129cf614f11565b03612a06576040517f5c8bfea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612a0f82612be8565b611673338383613061565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610e0457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610e04565b610e1c8133613b6b565b6119e15f5f613bf7565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610e04565b612af4613d82565b6115f38282613dc0565b6119e1613d82565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c15612b5f576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119e160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b90613e7c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156119e1576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116612c28576040517fc5e57bc700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035f9054906101000a90046001600160a01b03166001600160a01b031663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c78573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c9c91906150aa565b6001600160a01b0316336001600160a01b031614610e1c576040517f55f5762200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381165f9081526005602090815260408083205460109092529091205481151580612d1757505f81115b8390612d5b576040517fe337309e0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611285565b508115612d7357612d6a612da6565b612d7383613e83565b61196b83613f48565b6119e15f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00612b86565b5f612db1600a6139b8565b9050805f03612dbd5750565b5f8167ffffffffffffffff811115612dd757612dd7614c34565b604051908082528060200260200182016040528015612e00578160200160208202803683370190505b5090505f805b83811015612f93575f612e1a600a836139c1565b5f81815260096020526040902060078101549192509060ff1615612e6957818585612e4481615001565b965081518110612e5657612e5661504b565b6020026020010181815250505050612f8b565b8060020154431015612e7c575050612f8b565b80600501544311612e8e575050612f8b565b5f81600301544310612ea4578160030154612ea6565b435b90505f826005015482612eb991906150c5565b90505f836004015482612ecc9190614fea565b90506006545f03612ef3578060115f828254612ee89190615038565b90915550612f489050565b6006545f90612f0a670de0b6b3a764000084614fea565b612f149190614fb2565b905080856006015f828254612f299190615038565b9250508190555080600e5f828254612f419190615038565b9091555050505b6005840183905560038401544310612f8557848888612f6681615001565b995081518110612f7857612f7861504b565b6020026020010181815250505b50505050505b600101612e06565b505f5b81811015612fd457612fcb838281518110612fb357612fb361504b565b6020026020010151600a6139cc90919063ffffffff16565b50600101612f96565b50505050565b5f611540838361402f565b6040516001600160a01b038481166024830152838116604483015260648201839052612fd49186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061407b565b6001600160a01b0382165f9081526005602052604090205482906130bd576040517fe337309e0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611285565b505f81156130cb57816130e4565b6001600160a01b0383165f908152600560205260409020545b6001600160a01b0384165f908152600560205260409020549091508181811115613143576040517f9eba815000000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611285565b505061314d612da6565b61315683613e83565b61315f83613f48565b6001600160a01b0383165f90815260056020526040812080548392906131869084906150c5565b925050819055508060065f82825461319e91906150c5565b9091555050600e546001600160a01b0384165f908152600560205260409020546131c89190614fea565b6001600160a01b038085165f908152600f60205260409020919091556001546131f391168583613579565b826001600160a01b03167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58260405161322e91815260200190565b60405180910390a250505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461327581612ab0565b612fd48383614100565b6004548190808210156132c7576040517f9090182a00000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611285565b50506132d1612da6565b6132da82613e83565b6001546132f2906001600160a01b0316843084612fe5565b6001600160a01b0382165f9081526005602052604081208054839290613319908490615038565b925050819055508060065f8282546133319190615038565b9091555050600e546001600160a01b0383165f9081526005602052604090205461335b9190614fea565b6001600160a01b0383165f818152600f6020526040908190209290925590517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4906133a99084815260200190565b60405180910390a2505050565b6001600160a01b0381165f9081526005602090815260408083205460109092528220548183036133e7579392505050565b600e545f6133f5600a6139b8565b90505f5b818110156134d5575f61340d600a836139c1565b5f81815260096020526040902060078101549192509060ff16806134345750806002015443105b156134405750506134cd565b806005015443116134525750506134cd565b5f8160030154431061346857816003015461346a565b435b90505f82600501548261347d91906150c5565b90505f8360040154826134909190614fea565b600654909150156134c7576006546134b0670de0b6b3a764000083614fea565b6134ba9190614fb2565b6134c49089615038565b97505b50505050505b6001016133f9565b506001600160a01b0386165f908152600f6020526040812054670de0b6b3a7640000906135028588614fea565b61350c91906150c5565b6135169190614fb2565b90506135228185615038565b979650505050505050565b6001600160a01b038116331461356f576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61196b82826141e1565b6040516001600160a01b0383811660248301526044820183905261196b91859182169063a9059cbb9060640161301a565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061364357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166136377f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156119e1576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6115f381612ab0565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156136fc575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526136f991810190615093565b60015b61373d576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611285565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114613799576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611285565b61196b8383614277565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146119e1576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61380e61273c565b613817426142cc565b61382191906150d8565b905061382d828261431b565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f613881826143c8565b61388a426142cc565b61389491906150d8565b90506138a08282613bf7565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9101611e8c565b6138e7612b8c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b61395c61440f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613936565b60605f6115408361446a565b5f610e04825490565b5f61154083836144c3565b5f61154083836144e9565b5f611ea47feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16801580613a6c57504265ffffffffffff821610155b15613aad576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401611285565b613ae75f613ae27feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6141e1565b50613af25f83614100565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154613b5781612ab0565b612fd483836141e1565b6119e15f5f61431b565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166115f3576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611285565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015613d09574265ffffffffffff82161015613ce0576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255613d09565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b613d8a6145cc565b6119e1576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613dc8613d82565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b038216613e2b576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f6004820152602401611285565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff851602178155612fd45f83614100565b80825d5050565b6001600160a01b0381165f908152600560205260409020548015613f1c576001600160a01b0382165f908152600f6020526040812054600e54670de0b6b3a76400009190613ed19085614fea565b613edb91906150c5565b613ee59190614fb2565b90508015613f1a576001600160a01b0383165f9081526010602052604081208054839290613f14908490615038565b90915550505b505b600e54613f299082614fea565b6001600160a01b039092165f908152600f602052604090209190915550565b6001600160a01b0381165f9081526010602052604090205480156115f3576001600160a01b038083165f9081526010602052604081208190556002549091613f92911684846145ea565b905080613fed576001600160a01b038084165f818152601060205260409081902085905560025490519216917f0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d9190611c419086815260200190565b6002546040518381526001600160a01b03918216918516907f0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b790602001611c41565b5f81815260018301602052604081205461407457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610e04565b505f610e04565b5f5f60205f8451602086015f885af18061409a576040513d5f823e3d81fd5b50505f513d915081156140b15780600114156140be565b6001600160a01b0384163b155b15612fd4576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611285565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400836141cf575f6141597feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b031614614199576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b6141d9848461466c565b949350505050565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561423d57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b1561426d576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6141d9848461472f565b614280826147d3565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156142c45761196b828261487a565b6115f36148ec565b5f65ffffffffffff821115614317576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401611285565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b038816171784559104168015612fd4576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f6143d261273c565b90508065ffffffffffff168365ffffffffffff16116143fa576143f583826150f6565b611540565b61154065ffffffffffff841662069780614924565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166119e1576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060815f018054806020026020016040519081016040528092919081815260200182805480156144b757602002820191905f5260205f20905b8154815260200190600101908083116144a3575b50505050509050919050565b5f825f0182815481106144d8576144d861504b565b905f5260205f200154905092915050565b5f81815260018301602052604081205480156145c3575f61450b6001836150c5565b85549091505f9061451e906001906150c5565b905080821461457d575f865f01828154811061453c5761453c61504b565b905f5260205f200154905080875f01848154811061455c5761455c61504b565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061458e5761458e615114565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610e04565b5f915050610e04565b5f6145d5612ac4565b5468010000000000000000900460ff16919050565b5f6141d984856001600160a01b031663a9059cbb86866040516024016146259291906001600160a01b03929092168252602082015260400190565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050614933565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166145c3575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556146e53390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610e04565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156145c3575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610e04565b806001600160a01b03163b5f03614821576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611285565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516148969190615141565b5f60405180830381855af49150503d805f81146148ce576040519150601f19603f3d011682016040523d82523d5f602084013e6148d3565b606091505b50915091506148e385838361497c565b95945050505050565b34156119e1576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828218828410028218611540565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015614972575081156149645780600114614972565b5f866001600160a01b03163b115b9695505050505050565b60608261498c576143f5826149ec565b81511580156149a357506001600160a01b0384163b155b156149e5576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611285565b5080611540565b8051156149fb57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215614a3d575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611540575f5ffd5b6001600160a01b0381168114610e1c575f5ffd5b5f5f5f60608486031215614a92575f5ffd5b8335614a9d81614a6c565b92506020840135614aad81614a6c565b929592945050506040919091013590565b5f60208284031215614ace575f5ffd5b813561154081614a6c565b5f8151808452602084019350602083015f5b82811015614b125781516001600160a01b0316865260209586019590910190600101614aeb565b5093949350505050565b5f8151808452602084019350602083015f5b82811015614b12578151865260209586019590910190600101614b2e565b604081525f614b5e6040830185614ad9565b82810360208401526148e38185614b1c565b5f5f5f60608486031215614b82575f5ffd5b505081359360208301359350604090920135919050565b5f60208284031215614ba9575f5ffd5b5035919050565b5f5f60408385031215614bc1575f5ffd5b823591506020830135614bd381614a6c565b809150509250929050565b5f5f60408385031215614bef575f5ffd5b8235614bfa81614a6c565b946020939093013593505050565b5f5f60408385031215614c19575f5ffd5b8235614c2481614a6c565b91506020830135614bd381614a6c565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215614c72575f5ffd5b8235614c7d81614a6c565b9150602083013567ffffffffffffffff811115614c98575f5ffd5b8301601f81018513614ca8575f5ffd5b803567ffffffffffffffff811115614cc257614cc2614c34565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715614d2e57614d2e614c34565b604052818152828201602001871015614d45575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215614d74575f5ffd5b813565ffffffffffff81168114611540575f5ffd5b5f60208284031215614d99575f5ffd5b813560038110611540575f5ffd5b602081525f6115406020830184614b1c565b602080825282518282018190525f918401906040840190835b81811015614e4757614e30838551805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e0810151151560e08301525050565b602093909301926101009290920191600101614dd2565b509095945050505050565b6101008101610e048284805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e0810151151560e08301525050565b602081525f6115406020830184614ad9565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110614f73577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b60208101610e048284614f3e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f82614fe5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8082028115828204841417610e0457610e04614f85565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361503157615031614f85565b5060010190565b80820180821115610e0457610e04614f85565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b604081016150868285614f3e565b6115406020830184614f3e565b5f602082840312156150a3575f5ffd5b5051919050565b5f602082840312156150ba575f5ffd5b815161154081614a6c565b81810381811115610e0457610e04614f85565b65ffffffffffff8181168382160190811115610e0457610e04614f85565b65ffffffffffff8281168282160390811115610e0457610e04614f85565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f82518060208501845e5f92019182525091905056fea2646970667358221220bc746ad041df16cedd9c0f20060e705647fe7ec242ec7057a1588e191ee4aef464736f6c634300081c0033",
}

// CrossGameRewardPoolV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossGameRewardPoolV2MetaData.ABI instead.
var CrossGameRewardPoolV2ABI = CrossGameRewardPoolV2MetaData.ABI

// Deprecated: Use CrossGameRewardPoolV2MetaData.Sigs instead.
// CrossGameRewardPoolV2FuncSigs maps the 4-byte function signature to its string representation.
var CrossGameRewardPoolV2FuncSigs = CrossGameRewardPoolV2MetaData.Sigs

// CrossGameRewardPoolV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossGameRewardPoolV2MetaData.Bin instead.
var CrossGameRewardPoolV2Bin = CrossGameRewardPoolV2MetaData.Bin

// DeployCrossGameRewardPoolV2 deploys a new Ethereum contract, binding an instance of CrossGameRewardPoolV2 to it.
func DeployCrossGameRewardPoolV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossGameRewardPoolV2, error) {
	parsed, err := CrossGameRewardPoolV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossGameRewardPoolV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossGameRewardPoolV2{CrossGameRewardPoolV2Caller: CrossGameRewardPoolV2Caller{contract: contract}, CrossGameRewardPoolV2Transactor: CrossGameRewardPoolV2Transactor{contract: contract}, CrossGameRewardPoolV2Filterer: CrossGameRewardPoolV2Filterer{contract: contract}}, nil
}

// CrossGameRewardPoolV2 is an auto generated Go binding around an Ethereum contract.
type CrossGameRewardPoolV2 struct {
	CrossGameRewardPoolV2Caller     // Read-only binding to the contract
	CrossGameRewardPoolV2Transactor // Write-only binding to the contract
	CrossGameRewardPoolV2Filterer   // Log filterer for contract events
}

// CrossGameRewardPoolV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrossGameRewardPoolV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardPoolV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossGameRewardPoolV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardPoolV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossGameRewardPoolV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardPoolV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossGameRewardPoolV2Session struct {
	Contract     *CrossGameRewardPoolV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossGameRewardPoolV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossGameRewardPoolV2CallerSession struct {
	Contract *CrossGameRewardPoolV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CrossGameRewardPoolV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossGameRewardPoolV2TransactorSession struct {
	Contract     *CrossGameRewardPoolV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CrossGameRewardPoolV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrossGameRewardPoolV2Raw struct {
	Contract *CrossGameRewardPoolV2 // Generic contract binding to access the raw methods on
}

// CrossGameRewardPoolV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossGameRewardPoolV2CallerRaw struct {
	Contract *CrossGameRewardPoolV2Caller // Generic read-only contract binding to access the raw methods on
}

// CrossGameRewardPoolV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossGameRewardPoolV2TransactorRaw struct {
	Contract *CrossGameRewardPoolV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossGameRewardPoolV2 creates a new instance of CrossGameRewardPoolV2, bound to a specific deployed contract.
func NewCrossGameRewardPoolV2(address common.Address, backend bind.ContractBackend) (*CrossGameRewardPoolV2, error) {
	contract, err := bindCrossGameRewardPoolV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2{CrossGameRewardPoolV2Caller: CrossGameRewardPoolV2Caller{contract: contract}, CrossGameRewardPoolV2Transactor: CrossGameRewardPoolV2Transactor{contract: contract}, CrossGameRewardPoolV2Filterer: CrossGameRewardPoolV2Filterer{contract: contract}}, nil
}

// NewCrossGameRewardPoolV2Caller creates a new read-only instance of CrossGameRewardPoolV2, bound to a specific deployed contract.
func NewCrossGameRewardPoolV2Caller(address common.Address, caller bind.ContractCaller) (*CrossGameRewardPoolV2Caller, error) {
	contract, err := bindCrossGameRewardPoolV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2Caller{contract: contract}, nil
}

// NewCrossGameRewardPoolV2Transactor creates a new write-only instance of CrossGameRewardPoolV2, bound to a specific deployed contract.
func NewCrossGameRewardPoolV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CrossGameRewardPoolV2Transactor, error) {
	contract, err := bindCrossGameRewardPoolV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2Transactor{contract: contract}, nil
}

// NewCrossGameRewardPoolV2Filterer creates a new log filterer instance of CrossGameRewardPoolV2, bound to a specific deployed contract.
func NewCrossGameRewardPoolV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CrossGameRewardPoolV2Filterer, error) {
	contract, err := bindCrossGameRewardPoolV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2Filterer{contract: contract}, nil
}

// bindCrossGameRewardPoolV2 binds a generic wrapper to an already deployed contract.
func bindCrossGameRewardPoolV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossGameRewardPoolV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossGameRewardPoolV2.Contract.CrossGameRewardPoolV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CrossGameRewardPoolV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CrossGameRewardPoolV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossGameRewardPoolV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.DEFAULTADMINROLE(&_CrossGameRewardPoolV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.DEFAULTADMINROLE(&_CrossGameRewardPoolV2.CallOpts)
}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) SPONSORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "SPONSOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) SPONSORROLE() ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.SPONSORROLE(&_CrossGameRewardPoolV2.CallOpts)
}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) SPONSORROLE() ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.SPONSORROLE(&_CrossGameRewardPoolV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossGameRewardPoolV2.Contract.UPGRADEINTERFACEVERSION(&_CrossGameRewardPoolV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CrossGameRewardPoolV2.Contract.UPGRADEINTERFACEVERSION(&_CrossGameRewardPoolV2.CallOpts)
}

// AddRewardToken is a free data retrieval call binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address token) view returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) AddRewardToken(opts *bind.CallOpts, token common.Address) error {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "addRewardToken", token)

	if err != nil {
		return err
	}

	return err

}

// AddRewardToken is a free data retrieval call binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address token) view returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) AddRewardToken(token common.Address) error {
	return _CrossGameRewardPoolV2.Contract.AddRewardToken(&_CrossGameRewardPoolV2.CallOpts, token)
}

// AddRewardToken is a free data retrieval call binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address token) view returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) AddRewardToken(token common.Address) error {
	return _CrossGameRewardPoolV2.Contract.AddRewardToken(&_CrossGameRewardPoolV2.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) Balances(arg0 common.Address) (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.Balances(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.Balances(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) CrossGameReward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "crossGameReward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) CrossGameReward() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.CrossGameReward(&_CrossGameRewardPoolV2.CallOpts)
}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) CrossGameReward() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.CrossGameReward(&_CrossGameRewardPoolV2.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) DefaultAdmin() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.DefaultAdmin(&_CrossGameRewardPoolV2.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) DefaultAdmin() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.DefaultAdmin(&_CrossGameRewardPoolV2.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) DefaultAdminDelay() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.DefaultAdminDelay(&_CrossGameRewardPoolV2.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.DefaultAdminDelay(&_CrossGameRewardPoolV2.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.DefaultAdminDelayIncreaseWait(&_CrossGameRewardPoolV2.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.DefaultAdminDelayIncreaseWait(&_CrossGameRewardPoolV2.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) DepositToken() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.DepositToken(&_CrossGameRewardPoolV2.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) DepositToken() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.DepositToken(&_CrossGameRewardPoolV2.CallOpts)
}

// GetActiveRoundCount is a free data retrieval call binding the contract method 0x78ad8c7d.
//
// Solidity: function getActiveRoundCount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetActiveRoundCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getActiveRoundCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveRoundCount is a free data retrieval call binding the contract method 0x78ad8c7d.
//
// Solidity: function getActiveRoundCount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetActiveRoundCount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GetActiveRoundCount(&_CrossGameRewardPoolV2.CallOpts)
}

// GetActiveRoundCount is a free data retrieval call binding the contract method 0x78ad8c7d.
//
// Solidity: function getActiveRoundCount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetActiveRoundCount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GetActiveRoundCount(&_CrossGameRewardPoolV2.CallOpts)
}

// GetActiveRoundIds is a free data retrieval call binding the contract method 0x6fb7a4e8.
//
// Solidity: function getActiveRoundIds() view returns(uint256[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetActiveRoundIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getActiveRoundIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveRoundIds is a free data retrieval call binding the contract method 0x6fb7a4e8.
//
// Solidity: function getActiveRoundIds() view returns(uint256[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetActiveRoundIds() ([]*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GetActiveRoundIds(&_CrossGameRewardPoolV2.CallOpts)
}

// GetActiveRoundIds is a free data retrieval call binding the contract method 0x6fb7a4e8.
//
// Solidity: function getActiveRoundIds() view returns(uint256[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetActiveRoundIds() ([]*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GetActiveRoundIds(&_CrossGameRewardPoolV2.CallOpts)
}

// GetActiveRounds is a free data retrieval call binding the contract method 0x7d984d5f.
//
// Solidity: function getActiveRounds() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetActiveRounds(opts *bind.CallOpts) ([]ICrossGameRewardPoolV2Round, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getActiveRounds")

	if err != nil {
		return *new([]ICrossGameRewardPoolV2Round), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICrossGameRewardPoolV2Round)).(*[]ICrossGameRewardPoolV2Round)

	return out0, err

}

// GetActiveRounds is a free data retrieval call binding the contract method 0x7d984d5f.
//
// Solidity: function getActiveRounds() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetActiveRounds() ([]ICrossGameRewardPoolV2Round, error) {
	return _CrossGameRewardPoolV2.Contract.GetActiveRounds(&_CrossGameRewardPoolV2.CallOpts)
}

// GetActiveRounds is a free data retrieval call binding the contract method 0x7d984d5f.
//
// Solidity: function getActiveRounds() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetActiveRounds() ([]ICrossGameRewardPoolV2Round, error) {
	return _CrossGameRewardPoolV2.Contract.GetActiveRounds(&_CrossGameRewardPoolV2.CallOpts)
}

// GetReclaimableAmount is a free data retrieval call binding the contract method 0x35c21d5d.
//
// Solidity: function getReclaimableAmount(address token) view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetReclaimableAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getReclaimableAmount", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReclaimableAmount is a free data retrieval call binding the contract method 0x35c21d5d.
//
// Solidity: function getReclaimableAmount(address token) view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetReclaimableAmount(token common.Address) (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GetReclaimableAmount(&_CrossGameRewardPoolV2.CallOpts, token)
}

// GetReclaimableAmount is a free data retrieval call binding the contract method 0x35c21d5d.
//
// Solidity: function getReclaimableAmount(address token) view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetReclaimableAmount(token common.Address) (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GetReclaimableAmount(&_CrossGameRewardPoolV2.CallOpts, token)
}

// GetRemovedRewardTokens is a free data retrieval call binding the contract method 0x9b80c3f2.
//
// Solidity: function getRemovedRewardTokens() pure returns(address[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetRemovedRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getRemovedRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRemovedRewardTokens is a free data retrieval call binding the contract method 0x9b80c3f2.
//
// Solidity: function getRemovedRewardTokens() pure returns(address[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetRemovedRewardTokens() ([]common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.GetRemovedRewardTokens(&_CrossGameRewardPoolV2.CallOpts)
}

// GetRemovedRewardTokens is a free data retrieval call binding the contract method 0x9b80c3f2.
//
// Solidity: function getRemovedRewardTokens() pure returns(address[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetRemovedRewardTokens() ([]common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.GetRemovedRewardTokens(&_CrossGameRewardPoolV2.CallOpts)
}

// GetRemovedTokenRewards is a free data retrieval call binding the contract method 0x1af8acec.
//
// Solidity: function getRemovedTokenRewards(address ) pure returns(address[] tokens, uint256[] rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetRemovedTokenRewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getRemovedTokenRewards", arg0)

	outstruct := new(struct {
		Tokens  []common.Address
		Rewards []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Rewards = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRemovedTokenRewards is a free data retrieval call binding the contract method 0x1af8acec.
//
// Solidity: function getRemovedTokenRewards(address ) pure returns(address[] tokens, uint256[] rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetRemovedTokenRewards(arg0 common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.GetRemovedTokenRewards(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// GetRemovedTokenRewards is a free data retrieval call binding the contract method 0x1af8acec.
//
// Solidity: function getRemovedTokenRewards(address ) pure returns(address[] tokens, uint256[] rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetRemovedTokenRewards(arg0 common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.GetRemovedTokenRewards(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address token) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetRewardToken(opts *bind.CallOpts, token common.Address) (ICrossGameRewardPoolRewardToken, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getRewardToken", token)

	if err != nil {
		return *new(ICrossGameRewardPoolRewardToken), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossGameRewardPoolRewardToken)).(*ICrossGameRewardPoolRewardToken)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address token) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetRewardToken(token common.Address) (ICrossGameRewardPoolRewardToken, error) {
	return _CrossGameRewardPoolV2.Contract.GetRewardToken(&_CrossGameRewardPoolV2.CallOpts, token)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address token) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetRewardToken(token common.Address) (ICrossGameRewardPoolRewardToken, error) {
	return _CrossGameRewardPoolV2.Contract.GetRewardToken(&_CrossGameRewardPoolV2.CallOpts, token)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetRewardTokens() ([]common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.GetRewardTokens(&_CrossGameRewardPoolV2.CallOpts)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetRewardTokens() ([]common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.GetRewardTokens(&_CrossGameRewardPoolV2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.GetRoleAdmin(&_CrossGameRewardPoolV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.GetRoleAdmin(&_CrossGameRewardPoolV2.CallOpts, role)
}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 roundId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GetRound(opts *bind.CallOpts, roundId *big.Int) (ICrossGameRewardPoolV2Round, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "getRound", roundId)

	if err != nil {
		return *new(ICrossGameRewardPoolV2Round), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossGameRewardPoolV2Round)).(*ICrossGameRewardPoolV2Round)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 roundId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GetRound(roundId *big.Int) (ICrossGameRewardPoolV2Round, error) {
	return _CrossGameRewardPoolV2.Contract.GetRound(&_CrossGameRewardPoolV2.CallOpts, roundId)
}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 roundId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GetRound(roundId *big.Int) (ICrossGameRewardPoolV2Round, error) {
	return _CrossGameRewardPoolV2.Contract.GetRound(&_CrossGameRewardPoolV2.CallOpts, roundId)
}

// GlobalAccRewardPerShare is a free data retrieval call binding the contract method 0x2dbea37b.
//
// Solidity: function globalAccRewardPerShare() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) GlobalAccRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "globalAccRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalAccRewardPerShare is a free data retrieval call binding the contract method 0x2dbea37b.
//
// Solidity: function globalAccRewardPerShare() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GlobalAccRewardPerShare() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GlobalAccRewardPerShare(&_CrossGameRewardPoolV2.CallOpts)
}

// GlobalAccRewardPerShare is a free data retrieval call binding the contract method 0x2dbea37b.
//
// Solidity: function globalAccRewardPerShare() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) GlobalAccRewardPerShare() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.GlobalAccRewardPerShare(&_CrossGameRewardPoolV2.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.HasRole(&_CrossGameRewardPoolV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.HasRole(&_CrossGameRewardPoolV2.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) InitializedAt() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.InitializedAt(&_CrossGameRewardPoolV2.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) InitializedAt() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.InitializedAt(&_CrossGameRewardPoolV2.CallOpts)
}

// IsRemovedRewardToken is a free data retrieval call binding the contract method 0xf665336e.
//
// Solidity: function isRemovedRewardToken(address ) pure returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) IsRemovedRewardToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "isRemovedRewardToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRemovedRewardToken is a free data retrieval call binding the contract method 0xf665336e.
//
// Solidity: function isRemovedRewardToken(address ) pure returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) IsRemovedRewardToken(arg0 common.Address) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.IsRemovedRewardToken(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// IsRemovedRewardToken is a free data retrieval call binding the contract method 0xf665336e.
//
// Solidity: function isRemovedRewardToken(address ) pure returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) IsRemovedRewardToken(arg0 common.Address) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.IsRemovedRewardToken(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address token) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) IsRewardToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "isRewardToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address token) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) IsRewardToken(token common.Address) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.IsRewardToken(&_CrossGameRewardPoolV2.CallOpts, token)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address token) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) IsRewardToken(token common.Address) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.IsRewardToken(&_CrossGameRewardPoolV2.CallOpts, token)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) MinDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "minDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) MinDepositAmount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.MinDepositAmount(&_CrossGameRewardPoolV2.CallOpts)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) MinDepositAmount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.MinDepositAmount(&_CrossGameRewardPoolV2.CallOpts)
}

// NextRoundId is a free data retrieval call binding the contract method 0x4002eda6.
//
// Solidity: function nextRoundId() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) NextRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "nextRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRoundId is a free data retrieval call binding the contract method 0x4002eda6.
//
// Solidity: function nextRoundId() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) NextRoundId() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.NextRoundId(&_CrossGameRewardPoolV2.CallOpts)
}

// NextRoundId is a free data retrieval call binding the contract method 0x4002eda6.
//
// Solidity: function nextRoundId() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) NextRoundId() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.NextRoundId(&_CrossGameRewardPoolV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) Owner() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.Owner(&_CrossGameRewardPoolV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) Owner() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.Owner(&_CrossGameRewardPoolV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) Paused() (bool, error) {
	return _CrossGameRewardPoolV2.Contract.Paused(&_CrossGameRewardPoolV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) Paused() (bool, error) {
	return _CrossGameRewardPoolV2.Contract.Paused(&_CrossGameRewardPoolV2.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.PendingDefaultAdmin(&_CrossGameRewardPoolV2.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.PendingDefaultAdmin(&_CrossGameRewardPoolV2.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.PendingDefaultAdminDelay(&_CrossGameRewardPoolV2.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.PendingDefaultAdminDelay(&_CrossGameRewardPoolV2.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address ) view returns(uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) PendingReward(opts *bind.CallOpts, user common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "pendingReward", user, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address ) view returns(uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) PendingReward(user common.Address, arg1 common.Address) (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.PendingReward(&_CrossGameRewardPoolV2.CallOpts, user, arg1)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address ) view returns(uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) PendingReward(user common.Address, arg1 common.Address) (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.PendingReward(&_CrossGameRewardPoolV2.CallOpts, user, arg1)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address user) view returns(address[] tokens, uint256[] rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) PendingRewards(opts *bind.CallOpts, user common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "pendingRewards", user)

	outstruct := new(struct {
		Tokens  []common.Address
		Rewards []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Rewards = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address user) view returns(address[] tokens, uint256[] rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) PendingRewards(user common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.PendingRewards(&_CrossGameRewardPoolV2.CallOpts, user)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address user) view returns(address[] tokens, uint256[] rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) PendingRewards(user common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.PendingRewards(&_CrossGameRewardPoolV2.CallOpts, user)
}

// PoolStatus is a free data retrieval call binding the contract method 0xf0228692.
//
// Solidity: function poolStatus() view returns(uint8)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) PoolStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "poolStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolStatus is a free data retrieval call binding the contract method 0xf0228692.
//
// Solidity: function poolStatus() view returns(uint8)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) PoolStatus() (uint8, error) {
	return _CrossGameRewardPoolV2.Contract.PoolStatus(&_CrossGameRewardPoolV2.CallOpts)
}

// PoolStatus is a free data retrieval call binding the contract method 0xf0228692.
//
// Solidity: function poolStatus() view returns(uint8)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) PoolStatus() (uint8, error) {
	return _CrossGameRewardPoolV2.Contract.PoolStatus(&_CrossGameRewardPoolV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ProxiableUUID() ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.ProxiableUUID(&_CrossGameRewardPoolV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _CrossGameRewardPoolV2.Contract.ProxiableUUID(&_CrossGameRewardPoolV2.CallOpts)
}

// ReclaimableAmount is a free data retrieval call binding the contract method 0xfd8bdc68.
//
// Solidity: function reclaimableAmount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) ReclaimableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "reclaimableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReclaimableAmount is a free data retrieval call binding the contract method 0xfd8bdc68.
//
// Solidity: function reclaimableAmount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ReclaimableAmount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.ReclaimableAmount(&_CrossGameRewardPoolV2.CallOpts)
}

// ReclaimableAmount is a free data retrieval call binding the contract method 0xfd8bdc68.
//
// Solidity: function reclaimableAmount() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) ReclaimableAmount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.ReclaimableAmount(&_CrossGameRewardPoolV2.CallOpts)
}

// RemoveRewardToken is a free data retrieval call binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address ) view returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) RemoveRewardToken(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "removeRewardToken", arg0)

	if err != nil {
		return err
	}

	return err

}

// RemoveRewardToken is a free data retrieval call binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address ) view returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RemoveRewardToken(arg0 common.Address) error {
	return _CrossGameRewardPoolV2.Contract.RemoveRewardToken(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// RemoveRewardToken is a free data retrieval call binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address ) view returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) RemoveRewardToken(arg0 common.Address) error {
	return _CrossGameRewardPoolV2.Contract.RemoveRewardToken(&_CrossGameRewardPoolV2.CallOpts, arg0)
}

// RemovedRewardTokenCount is a free data retrieval call binding the contract method 0x35482379.
//
// Solidity: function removedRewardTokenCount() pure returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) RemovedRewardTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "removedRewardTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemovedRewardTokenCount is a free data retrieval call binding the contract method 0x35482379.
//
// Solidity: function removedRewardTokenCount() pure returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RemovedRewardTokenCount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.RemovedRewardTokenCount(&_CrossGameRewardPoolV2.CallOpts)
}

// RemovedRewardTokenCount is a free data retrieval call binding the contract method 0x35482379.
//
// Solidity: function removedRewardTokenCount() pure returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) RemovedRewardTokenCount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.RemovedRewardTokenCount(&_CrossGameRewardPoolV2.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RewardToken() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.RewardToken(&_CrossGameRewardPoolV2.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) RewardToken() (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.RewardToken(&_CrossGameRewardPoolV2.CallOpts)
}

// RewardTokenAt is a free data retrieval call binding the contract method 0x79f5ecb7.
//
// Solidity: function rewardTokenAt(uint256 index) view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) RewardTokenAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "rewardTokenAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokenAt is a free data retrieval call binding the contract method 0x79f5ecb7.
//
// Solidity: function rewardTokenAt(uint256 index) view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RewardTokenAt(index *big.Int) (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.RewardTokenAt(&_CrossGameRewardPoolV2.CallOpts, index)
}

// RewardTokenAt is a free data retrieval call binding the contract method 0x79f5ecb7.
//
// Solidity: function rewardTokenAt(uint256 index) view returns(address)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) RewardTokenAt(index *big.Int) (common.Address, error) {
	return _CrossGameRewardPoolV2.Contract.RewardTokenAt(&_CrossGameRewardPoolV2.CallOpts, index)
}

// RewardTokenCount is a free data retrieval call binding the contract method 0xabb06b95.
//
// Solidity: function rewardTokenCount() pure returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) RewardTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "rewardTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenCount is a free data retrieval call binding the contract method 0xabb06b95.
//
// Solidity: function rewardTokenCount() pure returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RewardTokenCount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.RewardTokenCount(&_CrossGameRewardPoolV2.CallOpts)
}

// RewardTokenCount is a free data retrieval call binding the contract method 0xabb06b95.
//
// Solidity: function rewardTokenCount() pure returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) RewardTokenCount() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.RewardTokenCount(&_CrossGameRewardPoolV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.SupportsInterface(&_CrossGameRewardPoolV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossGameRewardPoolV2.Contract.SupportsInterface(&_CrossGameRewardPoolV2.CallOpts, interfaceId)
}

// TotalDeposited is a free data retrieval call binding the contract method 0xff50abdc.
//
// Solidity: function totalDeposited() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) TotalDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "totalDeposited")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposited is a free data retrieval call binding the contract method 0xff50abdc.
//
// Solidity: function totalDeposited() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) TotalDeposited() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.TotalDeposited(&_CrossGameRewardPoolV2.CallOpts)
}

// TotalDeposited is a free data retrieval call binding the contract method 0xff50abdc.
//
// Solidity: function totalDeposited() view returns(uint256)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) TotalDeposited() (*big.Int, error) {
	return _CrossGameRewardPoolV2.Contract.TotalDeposited(&_CrossGameRewardPoolV2.CallOpts)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address account, address ) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Caller) UserRewards(opts *bind.CallOpts, account common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	var out []interface{}
	err := _CrossGameRewardPoolV2.contract.Call(opts, &out, "userRewards", account, arg1)

	outstruct := new(struct {
		RewardPerTokenPaid *big.Int
		Rewards            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardPerTokenPaid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address account, address ) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) UserRewards(account common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.UserRewards(&_CrossGameRewardPoolV2.CallOpts, account, arg1)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address account, address ) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2CallerSession) UserRewards(account common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _CrossGameRewardPoolV2.Contract.UserRewards(&_CrossGameRewardPoolV2.CallOpts, account, arg1)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.AcceptDefaultAdminTransfer(&_CrossGameRewardPoolV2.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.AcceptDefaultAdminTransfer(&_CrossGameRewardPoolV2.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.BeginDefaultAdminTransfer(&_CrossGameRewardPoolV2.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.BeginDefaultAdminTransfer(&_CrossGameRewardPoolV2.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CancelDefaultAdminTransfer(&_CrossGameRewardPoolV2.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CancelDefaultAdminTransfer(&_CrossGameRewardPoolV2.TransactOpts)
}

// CancelRound is a paid mutator transaction binding the contract method 0x7e07ab09.
//
// Solidity: function cancelRound(uint256 roundId) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) CancelRound(opts *bind.TransactOpts, roundId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "cancelRound", roundId)
}

// CancelRound is a paid mutator transaction binding the contract method 0x7e07ab09.
//
// Solidity: function cancelRound(uint256 roundId) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) CancelRound(roundId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CancelRound(&_CrossGameRewardPoolV2.TransactOpts, roundId)
}

// CancelRound is a paid mutator transaction binding the contract method 0x7e07ab09.
//
// Solidity: function cancelRound(uint256 roundId) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) CancelRound(roundId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CancelRound(&_CrossGameRewardPoolV2.TransactOpts, roundId)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ChangeDefaultAdminDelay(&_CrossGameRewardPoolV2.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ChangeDefaultAdminDelay(&_CrossGameRewardPoolV2.TransactOpts, newDelay)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address token) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) ClaimReward(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "claimReward", token)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address token) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ClaimReward(token common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimReward(&_CrossGameRewardPoolV2.TransactOpts, token)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address token) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) ClaimReward(token common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimReward(&_CrossGameRewardPoolV2.TransactOpts, token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0x35c30fda.
//
// Solidity: function claimRewardFor(address account, address token) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) ClaimRewardFor(opts *bind.TransactOpts, account common.Address, token common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "claimRewardFor", account, token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0x35c30fda.
//
// Solidity: function claimRewardFor(address account, address token) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ClaimRewardFor(account common.Address, token common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimRewardFor(&_CrossGameRewardPoolV2.TransactOpts, account, token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0x35c30fda.
//
// Solidity: function claimRewardFor(address account, address token) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) ClaimRewardFor(account common.Address, token common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimRewardFor(&_CrossGameRewardPoolV2.TransactOpts, account, token)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ClaimRewards() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimRewards(&_CrossGameRewardPoolV2.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimRewards(&_CrossGameRewardPoolV2.TransactOpts)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) ClaimRewardsFor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "claimRewardsFor", account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimRewardsFor(&_CrossGameRewardPoolV2.TransactOpts, account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ClaimRewardsFor(&_CrossGameRewardPoolV2.TransactOpts, account)
}

// CreateRound is a paid mutator transaction binding the contract method 0x1efed5f7.
//
// Solidity: function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) CreateRound(opts *bind.TransactOpts, amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "createRound", amount, startBlock, durationBlocks)
}

// CreateRound is a paid mutator transaction binding the contract method 0x1efed5f7.
//
// Solidity: function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) CreateRound(amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CreateRound(&_CrossGameRewardPoolV2.TransactOpts, amount, startBlock, durationBlocks)
}

// CreateRound is a paid mutator transaction binding the contract method 0x1efed5f7.
//
// Solidity: function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) CreateRound(amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.CreateRound(&_CrossGameRewardPoolV2.TransactOpts, amount, startBlock, durationBlocks)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.Deposit(&_CrossGameRewardPoolV2.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.Deposit(&_CrossGameRewardPoolV2.TransactOpts, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) DepositFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "depositFor", account, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) DepositFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.DepositFor(&_CrossGameRewardPoolV2.TransactOpts, account, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) DepositFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.DepositFor(&_CrossGameRewardPoolV2.TransactOpts, account, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.GrantRole(&_CrossGameRewardPoolV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.GrantRole(&_CrossGameRewardPoolV2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _depositToken, address _rewardToken, uint256 _minDepositAmount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) Initialize(opts *bind.TransactOpts, _depositToken common.Address, _rewardToken common.Address, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "initialize", _depositToken, _rewardToken, _minDepositAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _depositToken, address _rewardToken, uint256 _minDepositAmount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) Initialize(_depositToken common.Address, _rewardToken common.Address, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.Initialize(&_CrossGameRewardPoolV2.TransactOpts, _depositToken, _rewardToken, _minDepositAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _depositToken, address _rewardToken, uint256 _minDepositAmount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) Initialize(_depositToken common.Address, _rewardToken common.Address, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.Initialize(&_CrossGameRewardPoolV2.TransactOpts, _depositToken, _rewardToken, _minDepositAmount)
}

// ReclaimTokens is a paid mutator transaction binding the contract method 0x4d1cd014.
//
// Solidity: function reclaimTokens(address token, address to) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) ReclaimTokens(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "reclaimTokens", token, to)
}

// ReclaimTokens is a paid mutator transaction binding the contract method 0x4d1cd014.
//
// Solidity: function reclaimTokens(address token, address to) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) ReclaimTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ReclaimTokens(&_CrossGameRewardPoolV2.TransactOpts, token, to)
}

// ReclaimTokens is a paid mutator transaction binding the contract method 0x4d1cd014.
//
// Solidity: function reclaimTokens(address token, address to) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) ReclaimTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.ReclaimTokens(&_CrossGameRewardPoolV2.TransactOpts, token, to)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.RenounceRole(&_CrossGameRewardPoolV2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.RenounceRole(&_CrossGameRewardPoolV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.RevokeRole(&_CrossGameRewardPoolV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.RevokeRole(&_CrossGameRewardPoolV2.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.RollbackDefaultAdminDelay(&_CrossGameRewardPoolV2.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.RollbackDefaultAdminDelay(&_CrossGameRewardPoolV2.TransactOpts)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x6d7c49a2.
//
// Solidity: function setPoolStatus(uint8 newStatus) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) SetPoolStatus(opts *bind.TransactOpts, newStatus uint8) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "setPoolStatus", newStatus)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x6d7c49a2.
//
// Solidity: function setPoolStatus(uint8 newStatus) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) SetPoolStatus(newStatus uint8) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.SetPoolStatus(&_CrossGameRewardPoolV2.TransactOpts, newStatus)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x6d7c49a2.
//
// Solidity: function setPoolStatus(uint8 newStatus) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) SetPoolStatus(newStatus uint8) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.SetPoolStatus(&_CrossGameRewardPoolV2.TransactOpts, newStatus)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) UpdateMinDepositAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "updateMinDepositAmount", amount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) UpdateMinDepositAmount(amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.UpdateMinDepositAmount(&_CrossGameRewardPoolV2.TransactOpts, amount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) UpdateMinDepositAmount(amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.UpdateMinDepositAmount(&_CrossGameRewardPoolV2.TransactOpts, amount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.UpgradeToAndCall(&_CrossGameRewardPoolV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.UpgradeToAndCall(&_CrossGameRewardPoolV2.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.Withdraw(&_CrossGameRewardPoolV2.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.Withdraw(&_CrossGameRewardPoolV2.TransactOpts, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address account, uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Transactor) WithdrawFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.contract.Transact(opts, "withdrawFor", account, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address account, uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Session) WithdrawFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.WithdrawFor(&_CrossGameRewardPoolV2.TransactOpts, account, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address account, uint256 amount) returns()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2TransactorSession) WithdrawFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardPoolV2.Contract.WithdrawFor(&_CrossGameRewardPoolV2.TransactOpts, account, amount)
}

// CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator struct {
	Event *CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled)
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
		it.Event = new(CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled)
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
func (it *CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2DefaultAdminDelayChangeCanceledIterator{contract: _CrossGameRewardPoolV2.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled, error) {
	event := new(CrossGameRewardPoolV2DefaultAdminDelayChangeCanceled)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator struct {
	Event *CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled)
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
		it.Event = new(CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled)
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
func (it *CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2DefaultAdminDelayChangeScheduledIterator{contract: _CrossGameRewardPoolV2.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled, error) {
	event := new(CrossGameRewardPoolV2DefaultAdminDelayChangeScheduled)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator struct {
	Event *CrossGameRewardPoolV2DefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2DefaultAdminTransferCanceled)
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
		it.Event = new(CrossGameRewardPoolV2DefaultAdminTransferCanceled)
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
func (it *CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2DefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2DefaultAdminTransferCanceledIterator{contract: _CrossGameRewardPoolV2.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2DefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2DefaultAdminTransferCanceled)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseDefaultAdminTransferCanceled(log types.Log) (*CrossGameRewardPoolV2DefaultAdminTransferCanceled, error) {
	event := new(CrossGameRewardPoolV2DefaultAdminTransferCanceled)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator struct {
	Event *CrossGameRewardPoolV2DefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2DefaultAdminTransferScheduled)
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
		it.Event = new(CrossGameRewardPoolV2DefaultAdminTransferScheduled)
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
func (it *CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2DefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2DefaultAdminTransferScheduledIterator{contract: _CrossGameRewardPoolV2.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2DefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2DefaultAdminTransferScheduled)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseDefaultAdminTransferScheduled(log types.Log) (*CrossGameRewardPoolV2DefaultAdminTransferScheduled, error) {
	event := new(CrossGameRewardPoolV2DefaultAdminTransferScheduled)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2DepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2DepositedIterator struct {
	Event *CrossGameRewardPoolV2Deposited // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2Deposited)
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
		it.Event = new(CrossGameRewardPoolV2Deposited)
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
func (it *CrossGameRewardPoolV2DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2Deposited represents a Deposited event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2Deposited struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*CrossGameRewardPoolV2DepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2DepositedIterator{contract: _CrossGameRewardPoolV2.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2Deposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2Deposited)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseDeposited(log types.Log) (*CrossGameRewardPoolV2Deposited, error) {
	event := new(CrossGameRewardPoolV2Deposited)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2InitializedIterator struct {
	Event *CrossGameRewardPoolV2Initialized // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2Initialized)
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
		it.Event = new(CrossGameRewardPoolV2Initialized)
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
func (it *CrossGameRewardPoolV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2Initialized represents a Initialized event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*CrossGameRewardPoolV2InitializedIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2InitializedIterator{contract: _CrossGameRewardPoolV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2Initialized) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2Initialized)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseInitialized(log types.Log) (*CrossGameRewardPoolV2Initialized, error) {
	event := new(CrossGameRewardPoolV2Initialized)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2MinDepositAmountUpdatedIterator is returned from FilterMinDepositAmountUpdated and is used to iterate over the raw logs and unpacked data for MinDepositAmountUpdated events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2MinDepositAmountUpdatedIterator struct {
	Event *CrossGameRewardPoolV2MinDepositAmountUpdated // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2MinDepositAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2MinDepositAmountUpdated)
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
		it.Event = new(CrossGameRewardPoolV2MinDepositAmountUpdated)
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
func (it *CrossGameRewardPoolV2MinDepositAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2MinDepositAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2MinDepositAmountUpdated represents a MinDepositAmountUpdated event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2MinDepositAmountUpdated struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinDepositAmountUpdated is a free log retrieval operation binding the contract event 0x5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329.
//
// Solidity: event MinDepositAmountUpdated(uint256 oldAmount, uint256 newAmount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterMinDepositAmountUpdated(opts *bind.FilterOpts) (*CrossGameRewardPoolV2MinDepositAmountUpdatedIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "MinDepositAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2MinDepositAmountUpdatedIterator{contract: _CrossGameRewardPoolV2.contract, event: "MinDepositAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchMinDepositAmountUpdated is a free log subscription operation binding the contract event 0x5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329.
//
// Solidity: event MinDepositAmountUpdated(uint256 oldAmount, uint256 newAmount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchMinDepositAmountUpdated(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2MinDepositAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "MinDepositAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2MinDepositAmountUpdated)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "MinDepositAmountUpdated", log); err != nil {
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

// ParseMinDepositAmountUpdated is a log parse operation binding the contract event 0x5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329.
//
// Solidity: event MinDepositAmountUpdated(uint256 oldAmount, uint256 newAmount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseMinDepositAmountUpdated(log types.Log) (*CrossGameRewardPoolV2MinDepositAmountUpdated, error) {
	event := new(CrossGameRewardPoolV2MinDepositAmountUpdated)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "MinDepositAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2PausedIterator struct {
	Event *CrossGameRewardPoolV2Paused // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2Paused)
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
		it.Event = new(CrossGameRewardPoolV2Paused)
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
func (it *CrossGameRewardPoolV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2Paused represents a Paused event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterPaused(opts *bind.FilterOpts) (*CrossGameRewardPoolV2PausedIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2PausedIterator{contract: _CrossGameRewardPoolV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2Paused) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2Paused)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParsePaused(log types.Log) (*CrossGameRewardPoolV2Paused, error) {
	event := new(CrossGameRewardPoolV2Paused)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2PoolStatusChangedIterator is returned from FilterPoolStatusChanged and is used to iterate over the raw logs and unpacked data for PoolStatusChanged events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2PoolStatusChangedIterator struct {
	Event *CrossGameRewardPoolV2PoolStatusChanged // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2PoolStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2PoolStatusChanged)
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
		it.Event = new(CrossGameRewardPoolV2PoolStatusChanged)
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
func (it *CrossGameRewardPoolV2PoolStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2PoolStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2PoolStatusChanged represents a PoolStatusChanged event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2PoolStatusChanged struct {
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolStatusChanged is a free log retrieval operation binding the contract event 0xc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb5966.
//
// Solidity: event PoolStatusChanged(uint8 oldStatus, uint8 newStatus)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterPoolStatusChanged(opts *bind.FilterOpts) (*CrossGameRewardPoolV2PoolStatusChangedIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "PoolStatusChanged")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2PoolStatusChangedIterator{contract: _CrossGameRewardPoolV2.contract, event: "PoolStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPoolStatusChanged is a free log subscription operation binding the contract event 0xc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb5966.
//
// Solidity: event PoolStatusChanged(uint8 oldStatus, uint8 newStatus)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchPoolStatusChanged(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2PoolStatusChanged) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "PoolStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2PoolStatusChanged)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "PoolStatusChanged", log); err != nil {
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

// ParsePoolStatusChanged is a log parse operation binding the contract event 0xc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb5966.
//
// Solidity: event PoolStatusChanged(uint8 oldStatus, uint8 newStatus)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParsePoolStatusChanged(log types.Log) (*CrossGameRewardPoolV2PoolStatusChanged, error) {
	event := new(CrossGameRewardPoolV2PoolStatusChanged)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "PoolStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RewardClaimFailedIterator is returned from FilterRewardClaimFailed and is used to iterate over the raw logs and unpacked data for RewardClaimFailed events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RewardClaimFailedIterator struct {
	Event *CrossGameRewardPoolV2RewardClaimFailed // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RewardClaimFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RewardClaimFailed)
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
		it.Event = new(CrossGameRewardPoolV2RewardClaimFailed)
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
func (it *CrossGameRewardPoolV2RewardClaimFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RewardClaimFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RewardClaimFailed represents a RewardClaimFailed event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RewardClaimFailed struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimFailed is a free log retrieval operation binding the contract event 0x0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d91.
//
// Solidity: event RewardClaimFailed(address indexed account, address indexed token, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRewardClaimFailed(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*CrossGameRewardPoolV2RewardClaimFailedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RewardClaimFailed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RewardClaimFailedIterator{contract: _CrossGameRewardPoolV2.contract, event: "RewardClaimFailed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimFailed is a free log subscription operation binding the contract event 0x0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d91.
//
// Solidity: event RewardClaimFailed(address indexed account, address indexed token, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRewardClaimFailed(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RewardClaimFailed, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RewardClaimFailed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RewardClaimFailed)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RewardClaimFailed", log); err != nil {
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

// ParseRewardClaimFailed is a log parse operation binding the contract event 0x0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d91.
//
// Solidity: event RewardClaimFailed(address indexed account, address indexed token, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRewardClaimFailed(log types.Log) (*CrossGameRewardPoolV2RewardClaimFailed, error) {
	event := new(CrossGameRewardPoolV2RewardClaimFailed)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RewardClaimFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RewardClaimedIterator struct {
	Event *CrossGameRewardPoolV2RewardClaimed // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RewardClaimed)
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
		it.Event = new(CrossGameRewardPoolV2RewardClaimed)
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
func (it *CrossGameRewardPoolV2RewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RewardClaimed represents a RewardClaimed event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RewardClaimed struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed account, address indexed token, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRewardClaimed(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*CrossGameRewardPoolV2RewardClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RewardClaimed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RewardClaimedIterator{contract: _CrossGameRewardPoolV2.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed account, address indexed token, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RewardClaimed, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RewardClaimed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RewardClaimed)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed account, address indexed token, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRewardClaimed(log types.Log) (*CrossGameRewardPoolV2RewardClaimed, error) {
	event := new(CrossGameRewardPoolV2RewardClaimed)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoleAdminChangedIterator struct {
	Event *CrossGameRewardPoolV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RoleAdminChanged)
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
		it.Event = new(CrossGameRewardPoolV2RoleAdminChanged)
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
func (it *CrossGameRewardPoolV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RoleAdminChanged represents a RoleAdminChanged event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossGameRewardPoolV2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RoleAdminChangedIterator{contract: _CrossGameRewardPoolV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RoleAdminChanged)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRoleAdminChanged(log types.Log) (*CrossGameRewardPoolV2RoleAdminChanged, error) {
	event := new(CrossGameRewardPoolV2RoleAdminChanged)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoleGrantedIterator struct {
	Event *CrossGameRewardPoolV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RoleGranted)
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
		it.Event = new(CrossGameRewardPoolV2RoleGranted)
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
func (it *CrossGameRewardPoolV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RoleGranted represents a RoleGranted event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossGameRewardPoolV2RoleGrantedIterator, error) {

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

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RoleGrantedIterator{contract: _CrossGameRewardPoolV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RoleGranted)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRoleGranted(log types.Log) (*CrossGameRewardPoolV2RoleGranted, error) {
	event := new(CrossGameRewardPoolV2RoleGranted)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoleRevokedIterator struct {
	Event *CrossGameRewardPoolV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RoleRevoked)
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
		it.Event = new(CrossGameRewardPoolV2RoleRevoked)
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
func (it *CrossGameRewardPoolV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RoleRevoked represents a RoleRevoked event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossGameRewardPoolV2RoleRevokedIterator, error) {

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

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RoleRevokedIterator{contract: _CrossGameRewardPoolV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RoleRevoked)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRoleRevoked(log types.Log) (*CrossGameRewardPoolV2RoleRevoked, error) {
	event := new(CrossGameRewardPoolV2RoleRevoked)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RoundCancelledIterator is returned from FilterRoundCancelled and is used to iterate over the raw logs and unpacked data for RoundCancelled events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoundCancelledIterator struct {
	Event *CrossGameRewardPoolV2RoundCancelled // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RoundCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RoundCancelled)
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
		it.Event = new(CrossGameRewardPoolV2RoundCancelled)
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
func (it *CrossGameRewardPoolV2RoundCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RoundCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RoundCancelled represents a RoundCancelled event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoundCancelled struct {
	RoundId      *big.Int
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRoundCancelled is a free log retrieval operation binding the contract event 0x392fcf1e3627793dc153feb861f66451c925fa12c027044233166cd28f481d85.
//
// Solidity: event RoundCancelled(uint256 indexed roundId, uint256 refundAmount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRoundCancelled(opts *bind.FilterOpts, roundId []*big.Int) (*CrossGameRewardPoolV2RoundCancelledIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RoundCancelled", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RoundCancelledIterator{contract: _CrossGameRewardPoolV2.contract, event: "RoundCancelled", logs: logs, sub: sub}, nil
}

// WatchRoundCancelled is a free log subscription operation binding the contract event 0x392fcf1e3627793dc153feb861f66451c925fa12c027044233166cd28f481d85.
//
// Solidity: event RoundCancelled(uint256 indexed roundId, uint256 refundAmount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRoundCancelled(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RoundCancelled, roundId []*big.Int) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RoundCancelled", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RoundCancelled)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoundCancelled", log); err != nil {
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

// ParseRoundCancelled is a log parse operation binding the contract event 0x392fcf1e3627793dc153feb861f66451c925fa12c027044233166cd28f481d85.
//
// Solidity: event RoundCancelled(uint256 indexed roundId, uint256 refundAmount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRoundCancelled(log types.Log) (*CrossGameRewardPoolV2RoundCancelled, error) {
	event := new(CrossGameRewardPoolV2RoundCancelled)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoundCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2RoundCreatedIterator is returned from FilterRoundCreated and is used to iterate over the raw logs and unpacked data for RoundCreated events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoundCreatedIterator struct {
	Event *CrossGameRewardPoolV2RoundCreated // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2RoundCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2RoundCreated)
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
		it.Event = new(CrossGameRewardPoolV2RoundCreated)
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
func (it *CrossGameRewardPoolV2RoundCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2RoundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2RoundCreated represents a RoundCreated event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2RoundCreated struct {
	RoundId        *big.Int
	TotalReward    *big.Int
	StartBlock     *big.Int
	EndBlock       *big.Int
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRoundCreated is a free log retrieval operation binding the contract event 0xf4c8810c202a3e7371b142615e4811842e135aa5f919299d1cb4050710f7b85e.
//
// Solidity: event RoundCreated(uint256 indexed roundId, uint256 totalReward, uint256 startBlock, uint256 endBlock, uint256 rewardPerBlock)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterRoundCreated(opts *bind.FilterOpts, roundId []*big.Int) (*CrossGameRewardPoolV2RoundCreatedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "RoundCreated", roundIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2RoundCreatedIterator{contract: _CrossGameRewardPoolV2.contract, event: "RoundCreated", logs: logs, sub: sub}, nil
}

// WatchRoundCreated is a free log subscription operation binding the contract event 0xf4c8810c202a3e7371b142615e4811842e135aa5f919299d1cb4050710f7b85e.
//
// Solidity: event RoundCreated(uint256 indexed roundId, uint256 totalReward, uint256 startBlock, uint256 endBlock, uint256 rewardPerBlock)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchRoundCreated(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2RoundCreated, roundId []*big.Int) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "RoundCreated", roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2RoundCreated)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoundCreated", log); err != nil {
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

// ParseRoundCreated is a log parse operation binding the contract event 0xf4c8810c202a3e7371b142615e4811842e135aa5f919299d1cb4050710f7b85e.
//
// Solidity: event RoundCreated(uint256 indexed roundId, uint256 totalReward, uint256 startBlock, uint256 endBlock, uint256 rewardPerBlock)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseRoundCreated(log types.Log) (*CrossGameRewardPoolV2RoundCreated, error) {
	event := new(CrossGameRewardPoolV2RoundCreated)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "RoundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2TokensReclaimedIterator is returned from FilterTokensReclaimed and is used to iterate over the raw logs and unpacked data for TokensReclaimed events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2TokensReclaimedIterator struct {
	Event *CrossGameRewardPoolV2TokensReclaimed // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2TokensReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2TokensReclaimed)
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
		it.Event = new(CrossGameRewardPoolV2TokensReclaimed)
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
func (it *CrossGameRewardPoolV2TokensReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2TokensReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2TokensReclaimed represents a TokensReclaimed event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2TokensReclaimed struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensReclaimed is a free log retrieval operation binding the contract event 0x6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464.
//
// Solidity: event TokensReclaimed(address indexed token, address indexed to, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterTokensReclaimed(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*CrossGameRewardPoolV2TokensReclaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "TokensReclaimed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2TokensReclaimedIterator{contract: _CrossGameRewardPoolV2.contract, event: "TokensReclaimed", logs: logs, sub: sub}, nil
}

// WatchTokensReclaimed is a free log subscription operation binding the contract event 0x6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464.
//
// Solidity: event TokensReclaimed(address indexed token, address indexed to, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchTokensReclaimed(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2TokensReclaimed, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "TokensReclaimed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2TokensReclaimed)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "TokensReclaimed", log); err != nil {
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

// ParseTokensReclaimed is a log parse operation binding the contract event 0x6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464.
//
// Solidity: event TokensReclaimed(address indexed token, address indexed to, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseTokensReclaimed(log types.Log) (*CrossGameRewardPoolV2TokensReclaimed, error) {
	event := new(CrossGameRewardPoolV2TokensReclaimed)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "TokensReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2UnpausedIterator struct {
	Event *CrossGameRewardPoolV2Unpaused // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2Unpaused)
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
		it.Event = new(CrossGameRewardPoolV2Unpaused)
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
func (it *CrossGameRewardPoolV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2Unpaused represents a Unpaused event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CrossGameRewardPoolV2UnpausedIterator, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2UnpausedIterator{contract: _CrossGameRewardPoolV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2Unpaused)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseUnpaused(log types.Log) (*CrossGameRewardPoolV2Unpaused, error) {
	event := new(CrossGameRewardPoolV2Unpaused)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2UpgradedIterator struct {
	Event *CrossGameRewardPoolV2Upgraded // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2Upgraded)
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
		it.Event = new(CrossGameRewardPoolV2Upgraded)
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
func (it *CrossGameRewardPoolV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2Upgraded represents a Upgraded event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CrossGameRewardPoolV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2UpgradedIterator{contract: _CrossGameRewardPoolV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2Upgraded)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseUpgraded(log types.Log) (*CrossGameRewardPoolV2Upgraded, error) {
	event := new(CrossGameRewardPoolV2Upgraded)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardPoolV2WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2WithdrawnIterator struct {
	Event *CrossGameRewardPoolV2Withdrawn // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardPoolV2WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardPoolV2Withdrawn)
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
		it.Event = new(CrossGameRewardPoolV2Withdrawn)
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
func (it *CrossGameRewardPoolV2WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardPoolV2WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardPoolV2Withdrawn represents a Withdrawn event raised by the CrossGameRewardPoolV2 contract.
type CrossGameRewardPoolV2Withdrawn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*CrossGameRewardPoolV2WithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardPoolV2WithdrawnIterator{contract: _CrossGameRewardPoolV2.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CrossGameRewardPoolV2Withdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CrossGameRewardPoolV2.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardPoolV2Withdrawn)
				if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_CrossGameRewardPoolV2 *CrossGameRewardPoolV2Filterer) ParseWithdrawn(log types.Log) (*CrossGameRewardPoolV2Withdrawn, error) {
	event := new(CrossGameRewardPoolV2Withdrawn)
	if err := _CrossGameRewardPoolV2.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
