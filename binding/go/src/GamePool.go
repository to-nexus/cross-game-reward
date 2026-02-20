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

// IGamePoolRound is an auto generated low-level Go binding around an user-defined struct.
type IGamePoolRound struct {
	RoundId           *big.Int
	Creator           common.Address
	TotalReward       *big.Int
	StartBlock        *big.Int
	EndBlock          *big.Int
	RewardPerBlock    *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
	RemainderReward   *big.Int
	IsCancelled       bool
}

// GamePoolMetaData contains all meta data concerning the GamePool contract.
var GamePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPONSOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"cancelRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"cancelRoundToRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimRewardFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimRewardsFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"createRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"createRoundFromReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossGameReward\",\"outputs\":[{\"internalType\":\"contractICrossGameReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRoundCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRoundIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainderReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"}],\"internalType\":\"structIGamePool.Round[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getReclaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemovedRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRemovedTokenRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRewardToken\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reclaimableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRemoved\",\"type\":\"bool\"}],\"internalType\":\"structICrossGameRewardPool.RewardToken\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainderReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"}],\"internalType\":\"structIGamePool.Round\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalAccRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRemovedRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolStatus\",\"outputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reclaimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reclaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removedRewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"rewardTokenAt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRounds\",\"type\":\"uint256\"}],\"name\":\"syncRounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"processed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MinDepositAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"PoolStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"RoundCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"RoundCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"}],\"name\":\"RoundsSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"GPBelowMinimumDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"currentStatus\",\"type\":\"uint8\"}],\"name\":\"GPDepositNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"GPInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPInvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provided\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"GPInvalidRewardToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"GPInvalidStartBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GPNoDepositFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPNoReclaimableAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPNotAllowedInCurrentState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPOnlyRewardRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GPOnlyRoundCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPOnlyRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPPoolStatusUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPRewardIsDepositToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPRewardPerBlockZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"GPRoundAlreadyCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"GPRoundAlreadyStarted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"GPRoundNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPSyncNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
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
		"17bd64eb": "cancelRoundToRecipient(uint256,address)",
		"649a5ec7": "changeDefaultAdminDelay(uint48)",
		"d279c191": "claimReward(address)",
		"35c30fda": "claimRewardFor(address,address)",
		"372500ab": "claimRewards()",
		"1ac6d19d": "claimRewardsFor(address)",
		"1efed5f7": "createRound(uint256,uint256,uint256)",
		"69b65e91": "createRoundFromReserve(address,uint256,uint256,uint256)",
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
		"9df4496b": "syncRounds(uint256)",
		"ff50abdc": "totalDeposited()",
		"84780205": "updateMinDepositAmount(uint256)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"a980356a": "userRewards(address,address)",
		"2e1a7d4d": "withdraw(uint256)",
		"db518db2": "withdrawFor(address,uint256)",
	},
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516155236100f95f395f818161367001528181613699015261386901526155235ff3fe60806040526004361061044f575f3560e01c806379f5ecb711610237578063b5fd73f81161013c578063d547741f116100b7578063f4e2474011610087578063f7c618c11161006d578063f7c618c114610e11578063fd8bdc6814610e30578063ff50abdc14610e45575f5ffd5b8063f4e2474014610dd3578063f665336e14610df2575f5ffd5b8063d547741f14610d5b578063d602b9fd14610d7a578063db518db214610d8e578063f022869214610dad575f5ffd5b8063c89039c51161010c578063cefc1429116100f2578063cefc142914610cbe578063cf6eefb714610cd2578063d279c19114610d3c575f5ffd5b8063c89039c514610c8b578063cc8463c814610caa575f5ffd5b8063b5fd73f814610bf7578063b6b55f2514610c25578063c2d7944414610c44578063c4f59f9b14610c77575f5ffd5b806391d14854116101cc578063a1eda53c1161019c578063a980356a11610182578063a980356a14610b70578063abb06b9514610b8f578063ad3cb1cc14610ba2575f5ffd5b8063a1eda53c14610b2a578063a217fddf14610b5d575f5ffd5b806391d1485414610a4c5780639b80c3f214610aaf5780639ced7e7614610ad75780639df4496b14610af6575f5ffd5b806384ef8ffc1161020757806384ef8ffc146109bc5780638da5cb5b146109f85780638f1327c014610a0c57806391cf6d3e14610a38575f5ffd5b806379f5ecb7146109265780637d984d5f1461095d5780637e07ab091461097e578063847802051461099d575f5ffd5b806335c21d5d116103575780635c975abb116102d257806369b65e91116102a25780636fb7a4e8116102885780636fb7a4e81461087c578063770788721461089d57806378ad8c7d14610912575f5ffd5b806369b65e911461083e5780636d7c49a21461085d575f5ffd5b80635c975abb146107b5578063634e93da146107eb578063645006ca1461080a578063649a5ec71461081f575f5ffd5b80633d509c97116103275780634d1cd0141161030d5780634d1cd0141461076f5780634f1ef2861461078e57806352d1902d146107a1575f5ffd5b80633d509c971461073b5780634002eda61461075a575f5ffd5b806335c21d5d146106ca57806335c30fda146106e957806336568abe14610708578063372500ab14610727575f5ffd5b80631efed5f7116103e75780632e1a7d4d116103b75780632f4f21e21161039d5780632f4f21e21461067a57806331d7a2621461069957806335482379146106b8575f5ffd5b80632e1a7d4d1461063c5780632f2ff15d1461065b575f5ffd5b80631efed5f714610582578063248a9ca3146105af57806327e235e3146105fc5780632dbea37b14610627575f5ffd5b806317bd64eb1161042257806317bd64eb146104e45780631ac6d19d146105035780631af8acec146105225780631c03e6cc14610563575f5ffd5b806301ffc9a714610453578063022d63fb146104875780630aa6220b146104af5780631794bb3c146104c5575b5f5ffd5b34801561045e575f5ffd5b5061047261046d366004614da5565b610e5a565b60405190151581526020015b60405180910390f35b348015610492575f5ffd5b50620697805b60405165ffffffffffff909116815260200161047e565b3480156104ba575f5ffd5b506104c3610eb5565b005b3480156104d0575f5ffd5b506104c36104df366004614df8565b610eca565b3480156104ef575f5ffd5b506104c36104fe366004614e36565b6111d9565b34801561050e575f5ffd5b506104c361051d366004614e64565b6113c0565b34801561052d575f5ffd5b5061055561053c366004614e64565b50604080515f8082526020820190815281830190925291565b60405161047e929190614ef2565b34801561056e575f5ffd5b506104c361057d366004614e64565b61143a565b34801561058d575f5ffd5b506105a161059c366004614f16565b61151b565b60405190815260200161047e565b3480156105ba575f5ffd5b506105a16105c9366004614f3f565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b348015610607575f5ffd5b506105a1610616366004614e64565b60056020525f908152604090205481565b348015610632575f5ffd5b506105a1600c5481565b348015610647575f5ffd5b506104c3610656366004614f3f565b611532565b348015610666575f5ffd5b506104c3610675366004614e36565b61159d565b348015610685575f5ffd5b506104c3610694366004614f56565b6115de565b3480156106a4575f5ffd5b506105556106b3366004614e64565b611662565b3480156106c3575f5ffd5b505f6105a1565b3480156106d5575f5ffd5b506105a16106e4366004614e64565b611712565b3480156106f4575f5ffd5b506104c3610703366004614f80565b611738565b348015610713575f5ffd5b506104c3610722366004614e36565b61180a565b348015610732575f5ffd5b506104c3611957565b348015610746575f5ffd5b506104c3610755366004614e64565b6119ca565b348015610765575f5ffd5b506105a160085481565b34801561077a575f5ffd5b506104c3610789366004614f80565b611a96565b6104c361079c366004614fd9565b611c35565b3480156107ac575f5ffd5b506105a1611c50565b3480156107c0575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610472565b3480156107f6575f5ffd5b506104c3610805366004614e64565b611c7e565b348015610815575f5ffd5b506105a160045481565b34801561082a575f5ffd5b506104c36108393660046150dc565b611c91565b348015610849575f5ffd5b506105a1610858366004615101565b611ca4565b348015610868575f5ffd5b506104c3610877366004615139565b611fc6565b348015610887575f5ffd5b50610890612171565b60405161047e9190615157565b3480156108a8575f5ffd5b506108bc6108b7366004614e64565b612182565b60405161047e91905f60c0820190506001600160a01b0383511682526020830151602083015260408301516040830152606083015160608301526080830151608083015260a0830151151560a083015292915050565b34801561091d575f5ffd5b506105a16122ee565b348015610931575f5ffd5b50610945610940366004614f3f565b6122f9565b6040516001600160a01b03909116815260200161047e565b348015610968575f5ffd5b50610971612373565b60405161047e91906151e5565b348015610989575f5ffd5b506104c3610998366004614f3f565b6124f9565b3480156109a8575f5ffd5b506104c36109b7366004614f3f565b612503565b3480156109c7575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b0316610945565b348015610a03575f5ffd5b506109456125c1565b348015610a17575f5ffd5b50610a2b610a26366004614f3f565b6125ca565b60405161047e9190615233565b348015610a43575f5ffd5b506105a15f5481565b348015610a57575f5ffd5b50610472610a66366004614e36565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610aba575f5ffd5b50604080515f8152602081019091525b60405161047e9190615242565b348015610ae2575f5ffd5b506105a1610af1366004614f80565b6126f9565b348015610b01575f5ffd5b50610b15610b10366004614f3f565b612703565b6040805192835260208301919091520161047e565b348015610b35575f5ffd5b50610b3e612848565b6040805165ffffffffffff93841681529290911660208301520161047e565b348015610b68575f5ffd5b506105a15f81565b348015610b7b575f5ffd5b50610b15610b8a366004614f80565b612905565b348015610b9a575f5ffd5b5060016105a1565b348015610bad575f5ffd5b50610bea6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161047e9190615254565b348015610c02575f5ffd5b50610472610c11366004614e64565b6002546001600160a01b0391821691161490565b348015610c30575f5ffd5b506104c3610c3f366004614f3f565b61296c565b348015610c4f575f5ffd5b506105a17f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a81565b348015610c82575f5ffd5b50610aca6129df565b348015610c96575f5ffd5b50600154610945906001600160a01b031681565b348015610cb5575f5ffd5b50610498612a40565b348015610cc9575f5ffd5b506104c3612b20565b348015610cdd575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161047e565b348015610d47575f5ffd5b506104c3610d56366004614e64565b612b8e565b348015610d66575f5ffd5b506104c3610d75366004614e36565b612c57565b348015610d85575f5ffd5b506104c3612c98565b348015610d99575f5ffd5b506104c3610da8366004614f56565b612caa565b348015610db8575f5ffd5b50600754610dc69060ff1681565b60405161047e919061530d565b348015610dde575f5ffd5b50600354610945906001600160a01b031681565b348015610dfd575f5ffd5b50610472610e0c366004614e64565b505f90565b348015610e1c575f5ffd5b50600254610945906001600160a01b031681565b348015610e3b575f5ffd5b506105a1600f5481565b348015610e50575f5ffd5b506105a160065481565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610eaf5750610eaf82612d1e565b92915050565b5f610ebf81612db4565b610ec7612dbe565b50565b5f610ed3612dc8565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610eff5750825b90505f8267ffffffffffffffff166001148015610f1b5750303b155b905081158015610f29575080155b15610f60576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610fc15784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816611001576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716611041576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316886001600160a01b03160361108c576040517f4fdaf28b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f86116110c5576040517fe3dc980700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016339081179091556110fd905f90612df0565b611105612e02565b61110d612e02565b611115612e02565b435f55600180546001600160a01b038a81167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617835560028054918b169190921617905560048790556007805460ff1916905560085583156111cf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6111e1612e0a565b5f82815260096020526040902080548390611234576040517f50727ff000000000000000000000000000000000000000000000000000000000815260040161122b91815260200190565b60405180910390fd5b506009810154839060ff1615611279576040517f0eb9cf3d00000000000000000000000000000000000000000000000000000000815260040161122b91815260200190565b506001810154839033906001600160a01b03168181146112df576040517f7bd7574700000000000000000000000000000000000000000000000000000000815260048101939093526001600160a01b03918216602484015216604482015260640161122b565b505050806003015443108390611324576040517f6ce0091d00000000000000000000000000000000000000000000000000000000815260040161122b91815260200190565b506001600160a01b038216611337573391505b60098101805460ff19166001179055611351600a84612e90565b50600280820154905461136e906001600160a01b03168483612e9b565b826001600160a01b0316847f5cb7f91cf9cbcf0d6e29da784696d773d49366d9b704d703ccefc9c2f61c88b9836040516113aa91815260200190565b60405180910390a350506113bc612f0f565b5050565b6113c8612e0a565b6113d0612f39565b600260075460ff1660028111156113e9576113e96152a7565b03611420576040517f665134fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61142981612f95565b61143281613093565b610ec7612f0f565b6003546001600160a01b0316331461147e576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b03828116911614610ec7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f43616e6e6f742061646420646966666572656e742072657761726420746f6b6560448201527f6e00000000000000000000000000000000000000000000000000000000000000606482015260840161122b565b5f61152833858585611ca4565b90505b9392505050565b61153a612e0a565b611542612f39565b600260075460ff16600281111561155b5761155b6152a7565b03611592576040517f665134fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611432333383613129565b816115d4576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113bc8282613304565b6115e6612e0a565b6115ee612f39565b5f60075460ff166002811115611606576116066152a7565b60075460ff169114611645576040517f2cd1147f00000000000000000000000000000000000000000000000000000000815260040161122b919061530d565b5061164f82612f95565b61165a33838361334d565b6113bc612f0f565b604080516001808252818301909252606091829190602080830190803683375050604080516001808252818301909252929450905060208083019080368337505060025484519293506001600160a01b0316918491505f906116c6576116c661531b565b60200260200101906001600160a01b031690816001600160a01b0316815250506116ef83613484565b815f815181106117015761170161531b565b602002602001018181525050915091565b6002545f906001600160a01b0383811691161461173057505f919050565b5050600f5490565b611740612e0a565b611748612f39565b600260075460ff166002811115611761576117616152a7565b03611798576040517f665134fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025481906001600160a01b0390811690821681146117f6576040517f661779680000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529116602482015260440161122b565b505061180182612f95565b61165a82613093565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561186557507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b15611948577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16811515806118cb575065ffffffffffff8116155b806118de57504265ffffffffffff821610155b1561191f576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161122b565b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b6119528383613619565b505050565b61195f612e0a565b611967612f39565b600260075460ff166002811115611980576119806152a7565b036119b7576040517f665134fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119c033613093565b6119c8612f0f565b565b6003546001600160a01b03163314611a0e576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f47616d65506f6f6c20646f6573206e6f7420737570706f72742072656d6f766960448201527f6e672072657761726420746f6b656e7300000000000000000000000000000000606482015260840161122b565b6003546001600160a01b03163314611ada576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b03838116911614611b51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420746f6b656e00000000000000000000000000000000000000604482015260640161122b565b5f600f5411611b8c576040517f4f1aa73b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116611bcc576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f80545f909155600254611beb906001600160a01b03168383612e9b565b6002546040518281526001600160a01b038481169216907f6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464906020015b60405180910390a3505050565b611c3d613665565b611c4682613735565b6113bc828261373f565b5f611c5961385e565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f611c8881612db4565b6113bc826138c0565b5f611c9b81612db4565b6113bc82613932565b5f611cad612e0a565b611cb5612f39565b7f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a611cdf81612db4565b5f8511611d18576040517fe3dc980700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8343808211611d5c576040517f1302dfb80000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161122b565b50505f8311611d97576040517f7cf30a0300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6402540be40080611da98689615375565b611db39190615375565b611dbd91906153ad565b90505f8111611df8576040517f6fdfc85a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611e0385836153ad565b611e0d90886153c4565b9050611e1761399a565b60088054905f611e26836153d7565b9091555093505f611e37868861540e565b9050604051806101400160405280868152602001336001600160a01b031681526020018981526020018881526020018281526020018481526020018881526020015f81526020018381526020015f151581525060095f8781526020019081526020015f205f820151815f01556020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155610120820151816009015f6101000a81548160ff021916908315150217905550905050611f4d85600a6139a390919063ffffffff16565b50600254611f66906001600160a01b03168a308b6139ae565b604080518981526020810189905290810182905260608101849052339086907fba1778056db43899781dd248e1cab61b0f655af9ded9396f782b752e73de89899060800160405180910390a350505050611fbe612f0f565b949350505050565b6003546001600160a01b0316331461200a576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460ff16816002811115612022576120226152a7565b816002811115612034576120346152a7565b0361206b576040517f6a79d3d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007805483919060ff1916600183600281111561208a5761208a6152a7565b021790555060028260028111156120a3576120a36152a7565b1480156120d257507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16155b156120e4576120df6139e7565b612134565b60028260028111156120f8576120f86152a7565b1415801561212757507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff165b1561213457612134613a5c565b7fc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb59668183604051612165929190615421565b60405180910390a15050565b606061217d600a613ab4565b905090565b6121c06040518060c001604052805f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b6002546001600160a01b03838116911614612237576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420746f6b656e00000000000000000000000000000000000000604482015260640161122b565b6040805160c0810182526002546001600160a01b0316808252600c54602083015282517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291928301916370a0823190602401602060405180830381865afa1580156122ac573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122d0919061543c565b8152600f5460208201525f6040820181905260609091015292915050565b5f61217d600a613ac0565b5f8115612362576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420696e64657800000000000000000000000000000000000000604482015260640161122b565b50506002546001600160a01b031690565b60605f612380600a613ac0565b90505f8167ffffffffffffffff81111561239c5761239c614fac565b60405190808252806020026020018201604052801561242457816020015b6124116040518061014001604052805f81526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b8152602001906001900390816123ba5790505b5090505f5b828110156124f25760095f61243f600a84613ac9565b815260208082019290925260409081015f208151610140810183528154815260018201546001600160a01b0316938101939093526002810154918301919091526003810154606083015260048101546080830152600581015460a0830152600681015460c0830152600781015460e083015260088101546101008301526009015460ff16151561012082015282518390839081106124df576124df61531b565b6020908102919091010152600101612429565b5092915050565b610ec781336111d9565b6003546001600160a01b03163314612547576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8111612580576040517fe3dc980700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045460408051918252602082018390527f5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329910160405180910390a1600455565b5f61217d613ad4565b6126216040518061014001604052805f81526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b5f82815260096020526040902054829061266a576040517f50727ff000000000000000000000000000000000000000000000000000000000815260040161122b91815260200190565b50505f908152600960208181526040928390208351610140810185528154815260018201546001600160a01b0316928101929092526002810154938201939093526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e0820152600883015461010082015291015460ff16151561012082015290565b5f61152b83613484565b335f9081527f1f0283d10d2580ee30b219d86fcd3b74322ee3f926a3fbc0240167eb1c3e73e86020526040812054819060ff16806127c5575060035f9054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561278c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127b09190615453565b6001600160a01b0316336001600160a01b0316145b6127fb576040517f8e19b9c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61280483613b06565b60408051838152602081018390529294509092507f43cfb7702c1541a40272201ce7f99c71e1f1101be966dbcc218a3da1bbebc6e1910160405180910390a1915091565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840081158015906128ca57504265ffffffffffff831610155b6128d5575f5f6128fc565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b6001600160a01b0382165f908152600560205260408120548190801561294b576001600160a01b0385165f908152600d6020526040902054612948908290615375565b92505b50506001600160a01b039092165f908152600e602052604090205491929050565b612974612e0a565b61297c612f39565b5f60075460ff166002811115612994576129946152a7565b60075460ff1691146129d3576040517f2cd1147f00000000000000000000000000000000000000000000000000000000815260040161122b919061530d565b5061143233338361334d565b6040805160018082528183019092526060915f919060208083019080368337505060025482519293506001600160a01b0316918391505f90612a2357612a2361531b565b6001600160a01b0390921660209283029190910190910152919050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015801590612ac257504265ffffffffffff8216105b612af35781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16612b19565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114612b86576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161122b565b610ec7613d89565b612b96612e0a565b612b9e612f39565b600260075460ff166002811115612bb757612bb76152a7565b03612bee576040517f665134fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025481906001600160a01b039081169082168114612c4c576040517f661779680000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201529116602482015260440161122b565b505061143233613093565b81612c8e576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113bc8282613e9e565b5f612ca281612db4565b610ec7613ee1565b612cb2612e0a565b612cba612f39565b600260075460ff166002811115612cd357612cd36152a7565b03612d0a576040517f665134fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612d1382612f95565b61165a338383613129565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610eaf57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610eaf565b610ec78133613eeb565b6119c85f5f613f77565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610eaf565b612df8614102565b6113bc8282614140565b6119c8614102565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c15612e63576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119c860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b906141fc565b5f61152b8383614203565b6040516001600160a01b0383811660248301526044820183905261195291859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506142e6565b6119c85f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00612e8a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156119c8576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116612fd5576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035f9054906101000a90046001600160a01b03166001600160a01b031663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015613025573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130499190615453565b6001600160a01b0316336001600160a01b031614610ec7576040517fb1a7302300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381165f90815260056020908152604080832054600e90925290912054811515806130c457505f81115b8390613108576040517f500ca0dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161122b565b5081156131205761311761399a565b6131208361436b565b61195283614430565b6001600160a01b0382165f908152600560205260409020548290613185576040517f500ca0dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161122b565b505f811561319357816131ac565b6001600160a01b0383165f908152600560205260409020545b6001600160a01b0384165f90815260056020526040902054909150818181111561320b576040517fb22545ac0000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161122b565b505061321561399a565b61321e8361436b565b61322783614430565b6001600160a01b0383165f908152600560205260408120805483929061324e9084906153c4565b925050819055508060065f82825461326691906153c4565b9091555050600c546001600160a01b0384165f9081526005602052604090205461329091906153ad565b6001600160a01b038085165f908152600d60205260409020919091556001546132bb91168583612e9b565b826001600160a01b03167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5826040516132f691815260200190565b60405180910390a250505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461333d81612db4565b6133478383614517565b50505050565b600454819080821015613395576040517f26a1f59d0000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161122b565b505061339f61399a565b6133a88261436b565b6001546133c0906001600160a01b03168430846139ae565b6001600160a01b0382165f90815260056020526040812080548392906133e790849061540e565b925050819055508060065f8282546133ff919061540e565b9091555050600c546001600160a01b0383165f9081526005602052604090205461342991906153ad565b6001600160a01b0383165f818152600d6020526040908190209290925590517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4906134779084815260200190565b60405180910390a2505050565b6001600160a01b0381165f90815260056020908152604080832054600e9092528220548183036134b5579392505050565b600c545f6134c3600a613ac0565b90505f5b818110156135c1575f6134db600a83613ac9565b5f818152600960208190526040909120908101549192509060ff16806135045750806003015443105b156135105750506135b9565b806006015443116135225750506135b9565b5f8160040154431061353857816004015461353a565b435b90505f82600601548261354d91906153c4565b90505f83600501548261356091906153ad565b90508360040154830361357f57600884015461357c908261540e565b90505b600654156135b35760065461359c670de0b6b3a7640000836153ad565b6135a69190615375565b6135b0908961540e565b97505b50505050505b6001016134c7565b506001600160a01b0386165f908152600d6020526040812054670de0b6b3a7640000906135ee85886153ad565b6135f891906153c4565b6136029190615375565b905061360e818561540e565b979650505050505050565b6001600160a01b038116331461365b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61195282826145f0565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806136fe57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166136f27f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156119c8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6113bc81612db4565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156137b7575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526137b49181019061543c565b60015b6137f8576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161122b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114613854576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161122b565b6119528383614686565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146119c8576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6138c9612a40565b6138d2426146db565b6138dc919061546e565b90506138e8828261472a565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61393c826147d7565b613945426146db565b61394f919061546e565b905061395b8282613f77565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9101612165565b6113bc5f613b06565b5f61152b838361481e565b6040516001600160a01b0384811660248301528381166044830152606482018390526133479186918216906323b872dd90608401612ec8565b6139ef612f39565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b613a6461486a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613a3e565b60605f61152b836148c5565b5f610eaf825490565b5f61152b838361491e565b5f61217d7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b5f5f5f613b13600a613ac0565b9050805f03613b2757505f93849350915050565b5f8415613b4157818510613b3b5781613b43565b84613b43565b815b90505f8167ffffffffffffffff811115613b5f57613b5f614fac565b604051908082528060200260200182016040528015613b88578160200160208202803683370190505b5090505f805b83811015613d3a575f613ba2600a83613ac9565b5f818152600960208190526040909120908101549192509060ff1615613bf357818585613bce816153d7565b965081518110613be057613be061531b565b6020026020010181815250505050613d32565b8060030154431015613c06575050613d32565b80600601544311613c18575050613d32565b5f81600401544310613c2e578160040154613c30565b435b90505f826006015482613c4391906153c4565b90505f836005015482613c5691906153ad565b905083600401548303613c75576008840154613c72908261540e565b90505b6006545f03613c9a5780600f5f828254613c8f919061540e565b90915550613cef9050565b6006545f90613cb1670de0b6b3a7640000846153ad565b613cbb9190615375565b905080856007015f828254613cd0919061540e565b9250508190555080600c5f828254613ce8919061540e565b9091555050505b6006840183905560048401544310613d2c57848888613d0d816153d7565b995081518110613d1f57613d1f61531b565b6020026020010181815250505b50505050505b600101613b8e565b505f5b81811015613d7b57613d72838281518110613d5a57613d5a61531b565b6020026020010151600a612e9090919063ffffffff16565b50600101613d3d565b509196919550909350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16801580613dec57504265ffffffffffff821610155b15613e2d576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161122b565b613e675f613e627feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6145f0565b50613e725f83614517565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154613ed781612db4565b61334783836145f0565b6119c85f5f61472a565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166113bc576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161122b565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015614089574265ffffffffffff82161015614060576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a01000000000000000000000000000000000000000000000000000002178255614089565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b61410a614944565b6119c8576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614148614102565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b0382166141ab576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f600482015260240161122b565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556133475f83614517565b80825d5050565b5f81815260018301602052604081205480156142dd575f6142256001836153c4565b85549091505f90614238906001906153c4565b9050808214614297575f865f0182815481106142565761425661531b565b905f5260205f200154905080875f0184815481106142765761427661531b565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806142a8576142a861548c565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610eaf565b5f915050610eaf565b5f5f60205f8451602086015f885af180614305576040513d5f823e3d81fd5b50505f513d9150811561431c578060011415614329565b6001600160a01b0384163b155b15613347576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161122b565b6001600160a01b0381165f908152600560205260409020548015614404576001600160a01b0382165f908152600d6020526040812054600c54670de0b6b3a764000091906143b990856153ad565b6143c391906153c4565b6143cd9190615375565b90508015614402576001600160a01b0383165f908152600e6020526040812080548392906143fc90849061540e565b90915550505b505b600c5461441190826153ad565b6001600160a01b039092165f908152600d602052604090209190915550565b6001600160a01b0381165f908152600e602052604090205480156113bc576001600160a01b038083165f908152600e60205260408120819055600254909161447a91168484614962565b9050806144d5576001600160a01b038084165f818152600e60205260409081902085905560025490519216917f0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d9190611c289086815260200190565b6002546040518381526001600160a01b03918216918516907f0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b790602001611c28565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400836145e6575f6145707feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b0316146145b0576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b611fbe84846149e4565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561464c57507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b1561467c576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b611fbe8484614aa7565b61468f82614b4b565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156146d3576119528282614bf2565b6113bc614c64565b5f65ffffffffffff821115614726576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161122b565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b038816171784559104168015613347576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f6147e1612a40565b90508065ffffffffffff168365ffffffffffff16116148095761480483826154b9565b61152b565b61152b65ffffffffffff841662069780614c9c565b5f81815260018301602052604081205461486357508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610eaf565b505f610eaf565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166119c8576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060815f0180548060200260200160405190810160405280929190818152602001828054801561491257602002820191905f5260205f20905b8154815260200190600101908083116148fe575b50505050509050919050565b5f825f0182815481106149335761493361531b565b905f5260205f200154905092915050565b5f61494d612dc8565b5468010000000000000000900460ff16919050565b5f61152884856001600160a01b031663a9059cbb868660405160240161499d9291906001600160a01b03929092168252602082015260400190565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050614cab565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166142dd575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055614a5d3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610eaf565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156142dd575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610eaf565b806001600160a01b03163b5f03614b99576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161122b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051614c0e91906154d7565b5f60405180830381855af49150503d805f8114614c46576040519150601f19603f3d011682016040523d82523d5f602084013e614c4b565b606091505b5091509150614c5b858383614cf4565b95945050505050565b34156119c8576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82821882841002821861152b565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015614cea57508115614cdc5780600114614cea565b5f866001600160a01b03163b115b9695505050505050565b606082614d045761480482614d64565b8151158015614d1b57506001600160a01b0384163b155b15614d5d576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161122b565b508061152b565b805115614d7357805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215614db5575f5ffd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461152b575f5ffd5b6001600160a01b0381168114610ec7575f5ffd5b5f5f5f60608486031215614e0a575f5ffd5b8335614e1581614de4565b92506020840135614e2581614de4565b929592945050506040919091013590565b5f5f60408385031215614e47575f5ffd5b823591506020830135614e5981614de4565b809150509250929050565b5f60208284031215614e74575f5ffd5b813561152b81614de4565b5f8151808452602084019350602083015f5b82811015614eb85781516001600160a01b0316865260209586019590910190600101614e91565b5093949350505050565b5f8151808452602084019350602083015f5b82811015614eb8578151865260209586019590910190600101614ed4565b604081525f614f046040830185614e7f565b8281036020840152614c5b8185614ec2565b5f5f5f60608486031215614f28575f5ffd5b505081359360208301359350604090920135919050565b5f60208284031215614f4f575f5ffd5b5035919050565b5f5f60408385031215614f67575f5ffd5b8235614f7281614de4565b946020939093013593505050565b5f5f60408385031215614f91575f5ffd5b8235614f9c81614de4565b91506020830135614e5981614de4565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215614fea575f5ffd5b8235614ff581614de4565b9150602083013567ffffffffffffffff811115615010575f5ffd5b8301601f81018513615020575f5ffd5b803567ffffffffffffffff81111561503a5761503a614fac565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156150a6576150a6614fac565b6040528181528282016020018710156150bd575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f602082840312156150ec575f5ffd5b813565ffffffffffff8116811461152b575f5ffd5b5f5f5f5f60808587031215615114575f5ffd5b843561511f81614de4565b966020860135965060408601359560600135945092505050565b5f60208284031215615149575f5ffd5b81356003811061152b575f5ffd5b602081525f61152b6020830184614ec2565b80518252602081015161518760208401826001600160a01b03169052565b5060408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e083015261010081015161010083015261012081015161195261012084018215159052565b602080825282518282018190525f918401906040840190835b8181101561522857615211838551615169565b6020939093019261014092909201916001016151fe565b509095945050505050565b6101408101610eaf8284615169565b602081525f61152b6020830184614e7f565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110615309577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b60208101610eaf82846152d4565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f826153a8577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8082028115828204841417610eaf57610eaf615348565b81810381811115610eaf57610eaf615348565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361540757615407615348565b5060010190565b80820180821115610eaf57610eaf615348565b6040810161542f82856152d4565b61152b60208301846152d4565b5f6020828403121561544c575f5ffd5b5051919050565b5f60208284031215615463575f5ffd5b815161152b81614de4565b65ffffffffffff8181168382160190811115610eaf57610eaf615348565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b65ffffffffffff8281168282160390811115610eaf57610eaf615348565b5f82518060208501845e5f92019182525091905056fea2646970667358221220760be3ad91b7d6f8d592c9ff58eb0afa52f646d1db5b2a5fb8a27dab9f732c8d64736f6c634300081c0033",
}

// GamePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use GamePoolMetaData.ABI instead.
var GamePoolABI = GamePoolMetaData.ABI

// Deprecated: Use GamePoolMetaData.Sigs instead.
// GamePoolFuncSigs maps the 4-byte function signature to its string representation.
var GamePoolFuncSigs = GamePoolMetaData.Sigs

// GamePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GamePoolMetaData.Bin instead.
var GamePoolBin = GamePoolMetaData.Bin

// DeployGamePool deploys a new Ethereum contract, binding an instance of GamePool to it.
func DeployGamePool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GamePool, error) {
	parsed, err := GamePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GamePoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GamePool{GamePoolCaller: GamePoolCaller{contract: contract}, GamePoolTransactor: GamePoolTransactor{contract: contract}, GamePoolFilterer: GamePoolFilterer{contract: contract}}, nil
}

// GamePool is an auto generated Go binding around an Ethereum contract.
type GamePool struct {
	GamePoolCaller     // Read-only binding to the contract
	GamePoolTransactor // Write-only binding to the contract
	GamePoolFilterer   // Log filterer for contract events
}

// GamePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type GamePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GamePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GamePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GamePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GamePoolSession struct {
	Contract     *GamePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GamePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GamePoolCallerSession struct {
	Contract *GamePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GamePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GamePoolTransactorSession struct {
	Contract     *GamePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GamePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type GamePoolRaw struct {
	Contract *GamePool // Generic contract binding to access the raw methods on
}

// GamePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GamePoolCallerRaw struct {
	Contract *GamePoolCaller // Generic read-only contract binding to access the raw methods on
}

// GamePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GamePoolTransactorRaw struct {
	Contract *GamePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGamePool creates a new instance of GamePool, bound to a specific deployed contract.
func NewGamePool(address common.Address, backend bind.ContractBackend) (*GamePool, error) {
	contract, err := bindGamePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GamePool{GamePoolCaller: GamePoolCaller{contract: contract}, GamePoolTransactor: GamePoolTransactor{contract: contract}, GamePoolFilterer: GamePoolFilterer{contract: contract}}, nil
}

// NewGamePoolCaller creates a new read-only instance of GamePool, bound to a specific deployed contract.
func NewGamePoolCaller(address common.Address, caller bind.ContractCaller) (*GamePoolCaller, error) {
	contract, err := bindGamePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GamePoolCaller{contract: contract}, nil
}

// NewGamePoolTransactor creates a new write-only instance of GamePool, bound to a specific deployed contract.
func NewGamePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*GamePoolTransactor, error) {
	contract, err := bindGamePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GamePoolTransactor{contract: contract}, nil
}

// NewGamePoolFilterer creates a new log filterer instance of GamePool, bound to a specific deployed contract.
func NewGamePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*GamePoolFilterer, error) {
	contract, err := bindGamePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GamePoolFilterer{contract: contract}, nil
}

// bindGamePool binds a generic wrapper to an already deployed contract.
func bindGamePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GamePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GamePool *GamePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GamePool.Contract.GamePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GamePool *GamePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.Contract.GamePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GamePool *GamePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GamePool.Contract.GamePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GamePool *GamePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GamePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GamePool *GamePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GamePool *GamePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GamePool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GamePool *GamePoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GamePool *GamePoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GamePool.Contract.DEFAULTADMINROLE(&_GamePool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GamePool *GamePoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GamePool.Contract.DEFAULTADMINROLE(&_GamePool.CallOpts)
}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_GamePool *GamePoolCaller) SPONSORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "SPONSOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_GamePool *GamePoolSession) SPONSORROLE() ([32]byte, error) {
	return _GamePool.Contract.SPONSORROLE(&_GamePool.CallOpts)
}

// SPONSORROLE is a free data retrieval call binding the contract method 0xc2d79444.
//
// Solidity: function SPONSOR_ROLE() view returns(bytes32)
func (_GamePool *GamePoolCallerSession) SPONSORROLE() ([32]byte, error) {
	return _GamePool.Contract.SPONSORROLE(&_GamePool.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GamePool *GamePoolCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GamePool *GamePoolSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GamePool.Contract.UPGRADEINTERFACEVERSION(&_GamePool.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GamePool *GamePoolCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GamePool.Contract.UPGRADEINTERFACEVERSION(&_GamePool.CallOpts)
}

// AddRewardToken is a free data retrieval call binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address token) view returns()
func (_GamePool *GamePoolCaller) AddRewardToken(opts *bind.CallOpts, token common.Address) error {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "addRewardToken", token)

	if err != nil {
		return err
	}

	return err

}

// AddRewardToken is a free data retrieval call binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address token) view returns()
func (_GamePool *GamePoolSession) AddRewardToken(token common.Address) error {
	return _GamePool.Contract.AddRewardToken(&_GamePool.CallOpts, token)
}

// AddRewardToken is a free data retrieval call binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address token) view returns()
func (_GamePool *GamePoolCallerSession) AddRewardToken(token common.Address) error {
	return _GamePool.Contract.AddRewardToken(&_GamePool.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GamePool *GamePoolCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GamePool *GamePoolSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GamePool.Contract.Balances(&_GamePool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GamePool *GamePoolCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GamePool.Contract.Balances(&_GamePool.CallOpts, arg0)
}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_GamePool *GamePoolCaller) CrossGameReward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "crossGameReward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_GamePool *GamePoolSession) CrossGameReward() (common.Address, error) {
	return _GamePool.Contract.CrossGameReward(&_GamePool.CallOpts)
}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_GamePool *GamePoolCallerSession) CrossGameReward() (common.Address, error) {
	return _GamePool.Contract.CrossGameReward(&_GamePool.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_GamePool *GamePoolCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_GamePool *GamePoolSession) DefaultAdmin() (common.Address, error) {
	return _GamePool.Contract.DefaultAdmin(&_GamePool.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_GamePool *GamePoolCallerSession) DefaultAdmin() (common.Address, error) {
	return _GamePool.Contract.DefaultAdmin(&_GamePool.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_GamePool *GamePoolCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_GamePool *GamePoolSession) DefaultAdminDelay() (*big.Int, error) {
	return _GamePool.Contract.DefaultAdminDelay(&_GamePool.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_GamePool *GamePoolCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _GamePool.Contract.DefaultAdminDelay(&_GamePool.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_GamePool *GamePoolCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_GamePool *GamePoolSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _GamePool.Contract.DefaultAdminDelayIncreaseWait(&_GamePool.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_GamePool *GamePoolCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _GamePool.Contract.DefaultAdminDelayIncreaseWait(&_GamePool.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_GamePool *GamePoolCaller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_GamePool *GamePoolSession) DepositToken() (common.Address, error) {
	return _GamePool.Contract.DepositToken(&_GamePool.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_GamePool *GamePoolCallerSession) DepositToken() (common.Address, error) {
	return _GamePool.Contract.DepositToken(&_GamePool.CallOpts)
}

// GetActiveRoundCount is a free data retrieval call binding the contract method 0x78ad8c7d.
//
// Solidity: function getActiveRoundCount() view returns(uint256)
func (_GamePool *GamePoolCaller) GetActiveRoundCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getActiveRoundCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveRoundCount is a free data retrieval call binding the contract method 0x78ad8c7d.
//
// Solidity: function getActiveRoundCount() view returns(uint256)
func (_GamePool *GamePoolSession) GetActiveRoundCount() (*big.Int, error) {
	return _GamePool.Contract.GetActiveRoundCount(&_GamePool.CallOpts)
}

// GetActiveRoundCount is a free data retrieval call binding the contract method 0x78ad8c7d.
//
// Solidity: function getActiveRoundCount() view returns(uint256)
func (_GamePool *GamePoolCallerSession) GetActiveRoundCount() (*big.Int, error) {
	return _GamePool.Contract.GetActiveRoundCount(&_GamePool.CallOpts)
}

// GetActiveRoundIds is a free data retrieval call binding the contract method 0x6fb7a4e8.
//
// Solidity: function getActiveRoundIds() view returns(uint256[])
func (_GamePool *GamePoolCaller) GetActiveRoundIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getActiveRoundIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveRoundIds is a free data retrieval call binding the contract method 0x6fb7a4e8.
//
// Solidity: function getActiveRoundIds() view returns(uint256[])
func (_GamePool *GamePoolSession) GetActiveRoundIds() ([]*big.Int, error) {
	return _GamePool.Contract.GetActiveRoundIds(&_GamePool.CallOpts)
}

// GetActiveRoundIds is a free data retrieval call binding the contract method 0x6fb7a4e8.
//
// Solidity: function getActiveRoundIds() view returns(uint256[])
func (_GamePool *GamePoolCallerSession) GetActiveRoundIds() ([]*big.Int, error) {
	return _GamePool.Contract.GetActiveRoundIds(&_GamePool.CallOpts)
}

// GetActiveRounds is a free data retrieval call binding the contract method 0x7d984d5f.
//
// Solidity: function getActiveRounds() view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_GamePool *GamePoolCaller) GetActiveRounds(opts *bind.CallOpts) ([]IGamePoolRound, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getActiveRounds")

	if err != nil {
		return *new([]IGamePoolRound), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGamePoolRound)).(*[]IGamePoolRound)

	return out0, err

}

// GetActiveRounds is a free data retrieval call binding the contract method 0x7d984d5f.
//
// Solidity: function getActiveRounds() view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_GamePool *GamePoolSession) GetActiveRounds() ([]IGamePoolRound, error) {
	return _GamePool.Contract.GetActiveRounds(&_GamePool.CallOpts)
}

// GetActiveRounds is a free data retrieval call binding the contract method 0x7d984d5f.
//
// Solidity: function getActiveRounds() view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool)[])
func (_GamePool *GamePoolCallerSession) GetActiveRounds() ([]IGamePoolRound, error) {
	return _GamePool.Contract.GetActiveRounds(&_GamePool.CallOpts)
}

// GetReclaimableAmount is a free data retrieval call binding the contract method 0x35c21d5d.
//
// Solidity: function getReclaimableAmount(address token) view returns(uint256)
func (_GamePool *GamePoolCaller) GetReclaimableAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getReclaimableAmount", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReclaimableAmount is a free data retrieval call binding the contract method 0x35c21d5d.
//
// Solidity: function getReclaimableAmount(address token) view returns(uint256)
func (_GamePool *GamePoolSession) GetReclaimableAmount(token common.Address) (*big.Int, error) {
	return _GamePool.Contract.GetReclaimableAmount(&_GamePool.CallOpts, token)
}

// GetReclaimableAmount is a free data retrieval call binding the contract method 0x35c21d5d.
//
// Solidity: function getReclaimableAmount(address token) view returns(uint256)
func (_GamePool *GamePoolCallerSession) GetReclaimableAmount(token common.Address) (*big.Int, error) {
	return _GamePool.Contract.GetReclaimableAmount(&_GamePool.CallOpts, token)
}

// GetRemovedRewardTokens is a free data retrieval call binding the contract method 0x9b80c3f2.
//
// Solidity: function getRemovedRewardTokens() pure returns(address[])
func (_GamePool *GamePoolCaller) GetRemovedRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getRemovedRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRemovedRewardTokens is a free data retrieval call binding the contract method 0x9b80c3f2.
//
// Solidity: function getRemovedRewardTokens() pure returns(address[])
func (_GamePool *GamePoolSession) GetRemovedRewardTokens() ([]common.Address, error) {
	return _GamePool.Contract.GetRemovedRewardTokens(&_GamePool.CallOpts)
}

// GetRemovedRewardTokens is a free data retrieval call binding the contract method 0x9b80c3f2.
//
// Solidity: function getRemovedRewardTokens() pure returns(address[])
func (_GamePool *GamePoolCallerSession) GetRemovedRewardTokens() ([]common.Address, error) {
	return _GamePool.Contract.GetRemovedRewardTokens(&_GamePool.CallOpts)
}

// GetRemovedTokenRewards is a free data retrieval call binding the contract method 0x1af8acec.
//
// Solidity: function getRemovedTokenRewards(address ) pure returns(address[] tokens, uint256[] rewards)
func (_GamePool *GamePoolCaller) GetRemovedTokenRewards(opts *bind.CallOpts, arg0 common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getRemovedTokenRewards", arg0)

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
func (_GamePool *GamePoolSession) GetRemovedTokenRewards(arg0 common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _GamePool.Contract.GetRemovedTokenRewards(&_GamePool.CallOpts, arg0)
}

// GetRemovedTokenRewards is a free data retrieval call binding the contract method 0x1af8acec.
//
// Solidity: function getRemovedTokenRewards(address ) pure returns(address[] tokens, uint256[] rewards)
func (_GamePool *GamePoolCallerSession) GetRemovedTokenRewards(arg0 common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _GamePool.Contract.GetRemovedTokenRewards(&_GamePool.CallOpts, arg0)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address token) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_GamePool *GamePoolCaller) GetRewardToken(opts *bind.CallOpts, token common.Address) (ICrossGameRewardPoolRewardToken, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getRewardToken", token)

	if err != nil {
		return *new(ICrossGameRewardPoolRewardToken), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossGameRewardPoolRewardToken)).(*ICrossGameRewardPoolRewardToken)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address token) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_GamePool *GamePoolSession) GetRewardToken(token common.Address) (ICrossGameRewardPoolRewardToken, error) {
	return _GamePool.Contract.GetRewardToken(&_GamePool.CallOpts, token)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address token) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_GamePool *GamePoolCallerSession) GetRewardToken(token common.Address) (ICrossGameRewardPoolRewardToken, error) {
	return _GamePool.Contract.GetRewardToken(&_GamePool.CallOpts, token)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_GamePool *GamePoolCaller) GetRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_GamePool *GamePoolSession) GetRewardTokens() ([]common.Address, error) {
	return _GamePool.Contract.GetRewardTokens(&_GamePool.CallOpts)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_GamePool *GamePoolCallerSession) GetRewardTokens() ([]common.Address, error) {
	return _GamePool.Contract.GetRewardTokens(&_GamePool.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GamePool *GamePoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GamePool *GamePoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GamePool.Contract.GetRoleAdmin(&_GamePool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GamePool *GamePoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GamePool.Contract.GetRoleAdmin(&_GamePool.CallOpts, role)
}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 roundId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_GamePool *GamePoolCaller) GetRound(opts *bind.CallOpts, roundId *big.Int) (IGamePoolRound, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "getRound", roundId)

	if err != nil {
		return *new(IGamePoolRound), err
	}

	out0 := *abi.ConvertType(out[0], new(IGamePoolRound)).(*IGamePoolRound)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 roundId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_GamePool *GamePoolSession) GetRound(roundId *big.Int) (IGamePoolRound, error) {
	return _GamePool.Contract.GetRound(&_GamePool.CallOpts, roundId)
}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 roundId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_GamePool *GamePoolCallerSession) GetRound(roundId *big.Int) (IGamePoolRound, error) {
	return _GamePool.Contract.GetRound(&_GamePool.CallOpts, roundId)
}

// GlobalAccRewardPerShare is a free data retrieval call binding the contract method 0x2dbea37b.
//
// Solidity: function globalAccRewardPerShare() view returns(uint256)
func (_GamePool *GamePoolCaller) GlobalAccRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "globalAccRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalAccRewardPerShare is a free data retrieval call binding the contract method 0x2dbea37b.
//
// Solidity: function globalAccRewardPerShare() view returns(uint256)
func (_GamePool *GamePoolSession) GlobalAccRewardPerShare() (*big.Int, error) {
	return _GamePool.Contract.GlobalAccRewardPerShare(&_GamePool.CallOpts)
}

// GlobalAccRewardPerShare is a free data retrieval call binding the contract method 0x2dbea37b.
//
// Solidity: function globalAccRewardPerShare() view returns(uint256)
func (_GamePool *GamePoolCallerSession) GlobalAccRewardPerShare() (*big.Int, error) {
	return _GamePool.Contract.GlobalAccRewardPerShare(&_GamePool.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GamePool *GamePoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GamePool *GamePoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GamePool.Contract.HasRole(&_GamePool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GamePool *GamePoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GamePool.Contract.HasRole(&_GamePool.CallOpts, role, account)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_GamePool *GamePoolCaller) InitializedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "initializedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_GamePool *GamePoolSession) InitializedAt() (*big.Int, error) {
	return _GamePool.Contract.InitializedAt(&_GamePool.CallOpts)
}

// InitializedAt is a free data retrieval call binding the contract method 0x91cf6d3e.
//
// Solidity: function initializedAt() view returns(uint256)
func (_GamePool *GamePoolCallerSession) InitializedAt() (*big.Int, error) {
	return _GamePool.Contract.InitializedAt(&_GamePool.CallOpts)
}

// IsRemovedRewardToken is a free data retrieval call binding the contract method 0xf665336e.
//
// Solidity: function isRemovedRewardToken(address ) pure returns(bool)
func (_GamePool *GamePoolCaller) IsRemovedRewardToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "isRemovedRewardToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRemovedRewardToken is a free data retrieval call binding the contract method 0xf665336e.
//
// Solidity: function isRemovedRewardToken(address ) pure returns(bool)
func (_GamePool *GamePoolSession) IsRemovedRewardToken(arg0 common.Address) (bool, error) {
	return _GamePool.Contract.IsRemovedRewardToken(&_GamePool.CallOpts, arg0)
}

// IsRemovedRewardToken is a free data retrieval call binding the contract method 0xf665336e.
//
// Solidity: function isRemovedRewardToken(address ) pure returns(bool)
func (_GamePool *GamePoolCallerSession) IsRemovedRewardToken(arg0 common.Address) (bool, error) {
	return _GamePool.Contract.IsRemovedRewardToken(&_GamePool.CallOpts, arg0)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address token) view returns(bool)
func (_GamePool *GamePoolCaller) IsRewardToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "isRewardToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address token) view returns(bool)
func (_GamePool *GamePoolSession) IsRewardToken(token common.Address) (bool, error) {
	return _GamePool.Contract.IsRewardToken(&_GamePool.CallOpts, token)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address token) view returns(bool)
func (_GamePool *GamePoolCallerSession) IsRewardToken(token common.Address) (bool, error) {
	return _GamePool.Contract.IsRewardToken(&_GamePool.CallOpts, token)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() view returns(uint256)
func (_GamePool *GamePoolCaller) MinDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "minDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() view returns(uint256)
func (_GamePool *GamePoolSession) MinDepositAmount() (*big.Int, error) {
	return _GamePool.Contract.MinDepositAmount(&_GamePool.CallOpts)
}

// MinDepositAmount is a free data retrieval call binding the contract method 0x645006ca.
//
// Solidity: function minDepositAmount() view returns(uint256)
func (_GamePool *GamePoolCallerSession) MinDepositAmount() (*big.Int, error) {
	return _GamePool.Contract.MinDepositAmount(&_GamePool.CallOpts)
}

// NextRoundId is a free data retrieval call binding the contract method 0x4002eda6.
//
// Solidity: function nextRoundId() view returns(uint256)
func (_GamePool *GamePoolCaller) NextRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "nextRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRoundId is a free data retrieval call binding the contract method 0x4002eda6.
//
// Solidity: function nextRoundId() view returns(uint256)
func (_GamePool *GamePoolSession) NextRoundId() (*big.Int, error) {
	return _GamePool.Contract.NextRoundId(&_GamePool.CallOpts)
}

// NextRoundId is a free data retrieval call binding the contract method 0x4002eda6.
//
// Solidity: function nextRoundId() view returns(uint256)
func (_GamePool *GamePoolCallerSession) NextRoundId() (*big.Int, error) {
	return _GamePool.Contract.NextRoundId(&_GamePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GamePool *GamePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GamePool *GamePoolSession) Owner() (common.Address, error) {
	return _GamePool.Contract.Owner(&_GamePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GamePool *GamePoolCallerSession) Owner() (common.Address, error) {
	return _GamePool.Contract.Owner(&_GamePool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GamePool *GamePoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GamePool *GamePoolSession) Paused() (bool, error) {
	return _GamePool.Contract.Paused(&_GamePool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GamePool *GamePoolCallerSession) Paused() (bool, error) {
	return _GamePool.Contract.Paused(&_GamePool.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_GamePool *GamePoolCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_GamePool *GamePoolSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _GamePool.Contract.PendingDefaultAdmin(&_GamePool.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_GamePool *GamePoolCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _GamePool.Contract.PendingDefaultAdmin(&_GamePool.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_GamePool *GamePoolCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_GamePool *GamePoolSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _GamePool.Contract.PendingDefaultAdminDelay(&_GamePool.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_GamePool *GamePoolCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _GamePool.Contract.PendingDefaultAdminDelay(&_GamePool.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address ) view returns(uint256 amount)
func (_GamePool *GamePoolCaller) PendingReward(opts *bind.CallOpts, user common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "pendingReward", user, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address ) view returns(uint256 amount)
func (_GamePool *GamePoolSession) PendingReward(user common.Address, arg1 common.Address) (*big.Int, error) {
	return _GamePool.Contract.PendingReward(&_GamePool.CallOpts, user, arg1)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address ) view returns(uint256 amount)
func (_GamePool *GamePoolCallerSession) PendingReward(user common.Address, arg1 common.Address) (*big.Int, error) {
	return _GamePool.Contract.PendingReward(&_GamePool.CallOpts, user, arg1)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address user) view returns(address[] tokens, uint256[] rewards)
func (_GamePool *GamePoolCaller) PendingRewards(opts *bind.CallOpts, user common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "pendingRewards", user)

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
func (_GamePool *GamePoolSession) PendingRewards(user common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _GamePool.Contract.PendingRewards(&_GamePool.CallOpts, user)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address user) view returns(address[] tokens, uint256[] rewards)
func (_GamePool *GamePoolCallerSession) PendingRewards(user common.Address) (struct {
	Tokens  []common.Address
	Rewards []*big.Int
}, error) {
	return _GamePool.Contract.PendingRewards(&_GamePool.CallOpts, user)
}

// PoolStatus is a free data retrieval call binding the contract method 0xf0228692.
//
// Solidity: function poolStatus() view returns(uint8)
func (_GamePool *GamePoolCaller) PoolStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "poolStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolStatus is a free data retrieval call binding the contract method 0xf0228692.
//
// Solidity: function poolStatus() view returns(uint8)
func (_GamePool *GamePoolSession) PoolStatus() (uint8, error) {
	return _GamePool.Contract.PoolStatus(&_GamePool.CallOpts)
}

// PoolStatus is a free data retrieval call binding the contract method 0xf0228692.
//
// Solidity: function poolStatus() view returns(uint8)
func (_GamePool *GamePoolCallerSession) PoolStatus() (uint8, error) {
	return _GamePool.Contract.PoolStatus(&_GamePool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GamePool *GamePoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GamePool *GamePoolSession) ProxiableUUID() ([32]byte, error) {
	return _GamePool.Contract.ProxiableUUID(&_GamePool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GamePool *GamePoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GamePool.Contract.ProxiableUUID(&_GamePool.CallOpts)
}

// ReclaimableAmount is a free data retrieval call binding the contract method 0xfd8bdc68.
//
// Solidity: function reclaimableAmount() view returns(uint256)
func (_GamePool *GamePoolCaller) ReclaimableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "reclaimableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReclaimableAmount is a free data retrieval call binding the contract method 0xfd8bdc68.
//
// Solidity: function reclaimableAmount() view returns(uint256)
func (_GamePool *GamePoolSession) ReclaimableAmount() (*big.Int, error) {
	return _GamePool.Contract.ReclaimableAmount(&_GamePool.CallOpts)
}

// ReclaimableAmount is a free data retrieval call binding the contract method 0xfd8bdc68.
//
// Solidity: function reclaimableAmount() view returns(uint256)
func (_GamePool *GamePoolCallerSession) ReclaimableAmount() (*big.Int, error) {
	return _GamePool.Contract.ReclaimableAmount(&_GamePool.CallOpts)
}

// RemoveRewardToken is a free data retrieval call binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address ) view returns()
func (_GamePool *GamePoolCaller) RemoveRewardToken(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "removeRewardToken", arg0)

	if err != nil {
		return err
	}

	return err

}

// RemoveRewardToken is a free data retrieval call binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address ) view returns()
func (_GamePool *GamePoolSession) RemoveRewardToken(arg0 common.Address) error {
	return _GamePool.Contract.RemoveRewardToken(&_GamePool.CallOpts, arg0)
}

// RemoveRewardToken is a free data retrieval call binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address ) view returns()
func (_GamePool *GamePoolCallerSession) RemoveRewardToken(arg0 common.Address) error {
	return _GamePool.Contract.RemoveRewardToken(&_GamePool.CallOpts, arg0)
}

// RemovedRewardTokenCount is a free data retrieval call binding the contract method 0x35482379.
//
// Solidity: function removedRewardTokenCount() pure returns(uint256)
func (_GamePool *GamePoolCaller) RemovedRewardTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "removedRewardTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemovedRewardTokenCount is a free data retrieval call binding the contract method 0x35482379.
//
// Solidity: function removedRewardTokenCount() pure returns(uint256)
func (_GamePool *GamePoolSession) RemovedRewardTokenCount() (*big.Int, error) {
	return _GamePool.Contract.RemovedRewardTokenCount(&_GamePool.CallOpts)
}

// RemovedRewardTokenCount is a free data retrieval call binding the contract method 0x35482379.
//
// Solidity: function removedRewardTokenCount() pure returns(uint256)
func (_GamePool *GamePoolCallerSession) RemovedRewardTokenCount() (*big.Int, error) {
	return _GamePool.Contract.RemovedRewardTokenCount(&_GamePool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_GamePool *GamePoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_GamePool *GamePoolSession) RewardToken() (common.Address, error) {
	return _GamePool.Contract.RewardToken(&_GamePool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_GamePool *GamePoolCallerSession) RewardToken() (common.Address, error) {
	return _GamePool.Contract.RewardToken(&_GamePool.CallOpts)
}

// RewardTokenAt is a free data retrieval call binding the contract method 0x79f5ecb7.
//
// Solidity: function rewardTokenAt(uint256 index) view returns(address)
func (_GamePool *GamePoolCaller) RewardTokenAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "rewardTokenAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokenAt is a free data retrieval call binding the contract method 0x79f5ecb7.
//
// Solidity: function rewardTokenAt(uint256 index) view returns(address)
func (_GamePool *GamePoolSession) RewardTokenAt(index *big.Int) (common.Address, error) {
	return _GamePool.Contract.RewardTokenAt(&_GamePool.CallOpts, index)
}

// RewardTokenAt is a free data retrieval call binding the contract method 0x79f5ecb7.
//
// Solidity: function rewardTokenAt(uint256 index) view returns(address)
func (_GamePool *GamePoolCallerSession) RewardTokenAt(index *big.Int) (common.Address, error) {
	return _GamePool.Contract.RewardTokenAt(&_GamePool.CallOpts, index)
}

// RewardTokenCount is a free data retrieval call binding the contract method 0xabb06b95.
//
// Solidity: function rewardTokenCount() pure returns(uint256)
func (_GamePool *GamePoolCaller) RewardTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "rewardTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenCount is a free data retrieval call binding the contract method 0xabb06b95.
//
// Solidity: function rewardTokenCount() pure returns(uint256)
func (_GamePool *GamePoolSession) RewardTokenCount() (*big.Int, error) {
	return _GamePool.Contract.RewardTokenCount(&_GamePool.CallOpts)
}

// RewardTokenCount is a free data retrieval call binding the contract method 0xabb06b95.
//
// Solidity: function rewardTokenCount() pure returns(uint256)
func (_GamePool *GamePoolCallerSession) RewardTokenCount() (*big.Int, error) {
	return _GamePool.Contract.RewardTokenCount(&_GamePool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GamePool *GamePoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GamePool *GamePoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GamePool.Contract.SupportsInterface(&_GamePool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GamePool *GamePoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GamePool.Contract.SupportsInterface(&_GamePool.CallOpts, interfaceId)
}

// TotalDeposited is a free data retrieval call binding the contract method 0xff50abdc.
//
// Solidity: function totalDeposited() view returns(uint256)
func (_GamePool *GamePoolCaller) TotalDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "totalDeposited")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposited is a free data retrieval call binding the contract method 0xff50abdc.
//
// Solidity: function totalDeposited() view returns(uint256)
func (_GamePool *GamePoolSession) TotalDeposited() (*big.Int, error) {
	return _GamePool.Contract.TotalDeposited(&_GamePool.CallOpts)
}

// TotalDeposited is a free data retrieval call binding the contract method 0xff50abdc.
//
// Solidity: function totalDeposited() view returns(uint256)
func (_GamePool *GamePoolCallerSession) TotalDeposited() (*big.Int, error) {
	return _GamePool.Contract.TotalDeposited(&_GamePool.CallOpts)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address account, address ) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_GamePool *GamePoolCaller) UserRewards(opts *bind.CallOpts, account common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "userRewards", account, arg1)

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
func (_GamePool *GamePoolSession) UserRewards(account common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _GamePool.Contract.UserRewards(&_GamePool.CallOpts, account, arg1)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address account, address ) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_GamePool *GamePoolCallerSession) UserRewards(account common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _GamePool.Contract.UserRewards(&_GamePool.CallOpts, account, arg1)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_GamePool *GamePoolTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_GamePool *GamePoolSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _GamePool.Contract.AcceptDefaultAdminTransfer(&_GamePool.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_GamePool *GamePoolTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _GamePool.Contract.AcceptDefaultAdminTransfer(&_GamePool.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_GamePool *GamePoolTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_GamePool *GamePoolSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.BeginDefaultAdminTransfer(&_GamePool.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_GamePool *GamePoolTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.BeginDefaultAdminTransfer(&_GamePool.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_GamePool *GamePoolTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_GamePool *GamePoolSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _GamePool.Contract.CancelDefaultAdminTransfer(&_GamePool.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_GamePool *GamePoolTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _GamePool.Contract.CancelDefaultAdminTransfer(&_GamePool.TransactOpts)
}

// CancelRound is a paid mutator transaction binding the contract method 0x7e07ab09.
//
// Solidity: function cancelRound(uint256 roundId) returns()
func (_GamePool *GamePoolTransactor) CancelRound(opts *bind.TransactOpts, roundId *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "cancelRound", roundId)
}

// CancelRound is a paid mutator transaction binding the contract method 0x7e07ab09.
//
// Solidity: function cancelRound(uint256 roundId) returns()
func (_GamePool *GamePoolSession) CancelRound(roundId *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.CancelRound(&_GamePool.TransactOpts, roundId)
}

// CancelRound is a paid mutator transaction binding the contract method 0x7e07ab09.
//
// Solidity: function cancelRound(uint256 roundId) returns()
func (_GamePool *GamePoolTransactorSession) CancelRound(roundId *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.CancelRound(&_GamePool.TransactOpts, roundId)
}

// CancelRoundToRecipient is a paid mutator transaction binding the contract method 0x17bd64eb.
//
// Solidity: function cancelRoundToRecipient(uint256 roundId, address recipient) returns()
func (_GamePool *GamePoolTransactor) CancelRoundToRecipient(opts *bind.TransactOpts, roundId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "cancelRoundToRecipient", roundId, recipient)
}

// CancelRoundToRecipient is a paid mutator transaction binding the contract method 0x17bd64eb.
//
// Solidity: function cancelRoundToRecipient(uint256 roundId, address recipient) returns()
func (_GamePool *GamePoolSession) CancelRoundToRecipient(roundId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.CancelRoundToRecipient(&_GamePool.TransactOpts, roundId, recipient)
}

// CancelRoundToRecipient is a paid mutator transaction binding the contract method 0x17bd64eb.
//
// Solidity: function cancelRoundToRecipient(uint256 roundId, address recipient) returns()
func (_GamePool *GamePoolTransactorSession) CancelRoundToRecipient(roundId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.CancelRoundToRecipient(&_GamePool.TransactOpts, roundId, recipient)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_GamePool *GamePoolTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_GamePool *GamePoolSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.ChangeDefaultAdminDelay(&_GamePool.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_GamePool *GamePoolTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.ChangeDefaultAdminDelay(&_GamePool.TransactOpts, newDelay)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address token) returns()
func (_GamePool *GamePoolTransactor) ClaimReward(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "claimReward", token)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address token) returns()
func (_GamePool *GamePoolSession) ClaimReward(token common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ClaimReward(&_GamePool.TransactOpts, token)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address token) returns()
func (_GamePool *GamePoolTransactorSession) ClaimReward(token common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ClaimReward(&_GamePool.TransactOpts, token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0x35c30fda.
//
// Solidity: function claimRewardFor(address account, address token) returns()
func (_GamePool *GamePoolTransactor) ClaimRewardFor(opts *bind.TransactOpts, account common.Address, token common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "claimRewardFor", account, token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0x35c30fda.
//
// Solidity: function claimRewardFor(address account, address token) returns()
func (_GamePool *GamePoolSession) ClaimRewardFor(account common.Address, token common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ClaimRewardFor(&_GamePool.TransactOpts, account, token)
}

// ClaimRewardFor is a paid mutator transaction binding the contract method 0x35c30fda.
//
// Solidity: function claimRewardFor(address account, address token) returns()
func (_GamePool *GamePoolTransactorSession) ClaimRewardFor(account common.Address, token common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ClaimRewardFor(&_GamePool.TransactOpts, account, token)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_GamePool *GamePoolTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_GamePool *GamePoolSession) ClaimRewards() (*types.Transaction, error) {
	return _GamePool.Contract.ClaimRewards(&_GamePool.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_GamePool *GamePoolTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _GamePool.Contract.ClaimRewards(&_GamePool.TransactOpts)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_GamePool *GamePoolTransactor) ClaimRewardsFor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "claimRewardsFor", account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_GamePool *GamePoolSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ClaimRewardsFor(&_GamePool.TransactOpts, account)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0x1ac6d19d.
//
// Solidity: function claimRewardsFor(address account) returns()
func (_GamePool *GamePoolTransactorSession) ClaimRewardsFor(account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ClaimRewardsFor(&_GamePool.TransactOpts, account)
}

// CreateRound is a paid mutator transaction binding the contract method 0x1efed5f7.
//
// Solidity: function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_GamePool *GamePoolTransactor) CreateRound(opts *bind.TransactOpts, amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "createRound", amount, startBlock, durationBlocks)
}

// CreateRound is a paid mutator transaction binding the contract method 0x1efed5f7.
//
// Solidity: function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_GamePool *GamePoolSession) CreateRound(amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.CreateRound(&_GamePool.TransactOpts, amount, startBlock, durationBlocks)
}

// CreateRound is a paid mutator transaction binding the contract method 0x1efed5f7.
//
// Solidity: function createRound(uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_GamePool *GamePoolTransactorSession) CreateRound(amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.CreateRound(&_GamePool.TransactOpts, amount, startBlock, durationBlocks)
}

// CreateRoundFromReserve is a paid mutator transaction binding the contract method 0x69b65e91.
//
// Solidity: function createRoundFromReserve(address reserve, uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_GamePool *GamePoolTransactor) CreateRoundFromReserve(opts *bind.TransactOpts, reserve common.Address, amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "createRoundFromReserve", reserve, amount, startBlock, durationBlocks)
}

// CreateRoundFromReserve is a paid mutator transaction binding the contract method 0x69b65e91.
//
// Solidity: function createRoundFromReserve(address reserve, uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_GamePool *GamePoolSession) CreateRoundFromReserve(reserve common.Address, amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.CreateRoundFromReserve(&_GamePool.TransactOpts, reserve, amount, startBlock, durationBlocks)
}

// CreateRoundFromReserve is a paid mutator transaction binding the contract method 0x69b65e91.
//
// Solidity: function createRoundFromReserve(address reserve, uint256 amount, uint256 startBlock, uint256 durationBlocks) returns(uint256 roundId)
func (_GamePool *GamePoolTransactorSession) CreateRoundFromReserve(reserve common.Address, amount *big.Int, startBlock *big.Int, durationBlocks *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.CreateRoundFromReserve(&_GamePool.TransactOpts, reserve, amount, startBlock, durationBlocks)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_GamePool *GamePoolTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_GamePool *GamePoolSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.Deposit(&_GamePool.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_GamePool *GamePoolTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.Deposit(&_GamePool.TransactOpts, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns()
func (_GamePool *GamePoolTransactor) DepositFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "depositFor", account, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns()
func (_GamePool *GamePoolSession) DepositFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.DepositFor(&_GamePool.TransactOpts, account, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns()
func (_GamePool *GamePoolTransactorSession) DepositFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.DepositFor(&_GamePool.TransactOpts, account, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.GrantRole(&_GamePool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.GrantRole(&_GamePool.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _depositToken, address _rewardToken, uint256 _minDepositAmount) returns()
func (_GamePool *GamePoolTransactor) Initialize(opts *bind.TransactOpts, _depositToken common.Address, _rewardToken common.Address, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "initialize", _depositToken, _rewardToken, _minDepositAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _depositToken, address _rewardToken, uint256 _minDepositAmount) returns()
func (_GamePool *GamePoolSession) Initialize(_depositToken common.Address, _rewardToken common.Address, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.Initialize(&_GamePool.TransactOpts, _depositToken, _rewardToken, _minDepositAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _depositToken, address _rewardToken, uint256 _minDepositAmount) returns()
func (_GamePool *GamePoolTransactorSession) Initialize(_depositToken common.Address, _rewardToken common.Address, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.Initialize(&_GamePool.TransactOpts, _depositToken, _rewardToken, _minDepositAmount)
}

// ReclaimTokens is a paid mutator transaction binding the contract method 0x4d1cd014.
//
// Solidity: function reclaimTokens(address token, address to) returns()
func (_GamePool *GamePoolTransactor) ReclaimTokens(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "reclaimTokens", token, to)
}

// ReclaimTokens is a paid mutator transaction binding the contract method 0x4d1cd014.
//
// Solidity: function reclaimTokens(address token, address to) returns()
func (_GamePool *GamePoolSession) ReclaimTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ReclaimTokens(&_GamePool.TransactOpts, token, to)
}

// ReclaimTokens is a paid mutator transaction binding the contract method 0x4d1cd014.
//
// Solidity: function reclaimTokens(address token, address to) returns()
func (_GamePool *GamePoolTransactorSession) ReclaimTokens(token common.Address, to common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.ReclaimTokens(&_GamePool.TransactOpts, token, to)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.RenounceRole(&_GamePool.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.RenounceRole(&_GamePool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.RevokeRole(&_GamePool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GamePool *GamePoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GamePool.Contract.RevokeRole(&_GamePool.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_GamePool *GamePoolTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_GamePool *GamePoolSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _GamePool.Contract.RollbackDefaultAdminDelay(&_GamePool.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_GamePool *GamePoolTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _GamePool.Contract.RollbackDefaultAdminDelay(&_GamePool.TransactOpts)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x6d7c49a2.
//
// Solidity: function setPoolStatus(uint8 newStatus) returns()
func (_GamePool *GamePoolTransactor) SetPoolStatus(opts *bind.TransactOpts, newStatus uint8) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "setPoolStatus", newStatus)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x6d7c49a2.
//
// Solidity: function setPoolStatus(uint8 newStatus) returns()
func (_GamePool *GamePoolSession) SetPoolStatus(newStatus uint8) (*types.Transaction, error) {
	return _GamePool.Contract.SetPoolStatus(&_GamePool.TransactOpts, newStatus)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x6d7c49a2.
//
// Solidity: function setPoolStatus(uint8 newStatus) returns()
func (_GamePool *GamePoolTransactorSession) SetPoolStatus(newStatus uint8) (*types.Transaction, error) {
	return _GamePool.Contract.SetPoolStatus(&_GamePool.TransactOpts, newStatus)
}

// SyncRounds is a paid mutator transaction binding the contract method 0x9df4496b.
//
// Solidity: function syncRounds(uint256 maxRounds) returns(uint256 processed, uint256 removed)
func (_GamePool *GamePoolTransactor) SyncRounds(opts *bind.TransactOpts, maxRounds *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "syncRounds", maxRounds)
}

// SyncRounds is a paid mutator transaction binding the contract method 0x9df4496b.
//
// Solidity: function syncRounds(uint256 maxRounds) returns(uint256 processed, uint256 removed)
func (_GamePool *GamePoolSession) SyncRounds(maxRounds *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.SyncRounds(&_GamePool.TransactOpts, maxRounds)
}

// SyncRounds is a paid mutator transaction binding the contract method 0x9df4496b.
//
// Solidity: function syncRounds(uint256 maxRounds) returns(uint256 processed, uint256 removed)
func (_GamePool *GamePoolTransactorSession) SyncRounds(maxRounds *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.SyncRounds(&_GamePool.TransactOpts, maxRounds)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 amount) returns()
func (_GamePool *GamePoolTransactor) UpdateMinDepositAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "updateMinDepositAmount", amount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 amount) returns()
func (_GamePool *GamePoolSession) UpdateMinDepositAmount(amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.UpdateMinDepositAmount(&_GamePool.TransactOpts, amount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 amount) returns()
func (_GamePool *GamePoolTransactorSession) UpdateMinDepositAmount(amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.UpdateMinDepositAmount(&_GamePool.TransactOpts, amount)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GamePool *GamePoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GamePool *GamePoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GamePool.Contract.UpgradeToAndCall(&_GamePool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GamePool *GamePoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GamePool.Contract.UpgradeToAndCall(&_GamePool.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_GamePool *GamePoolTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_GamePool *GamePoolSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.Withdraw(&_GamePool.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_GamePool *GamePoolTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.Withdraw(&_GamePool.TransactOpts, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address account, uint256 amount) returns()
func (_GamePool *GamePoolTransactor) WithdrawFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "withdrawFor", account, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address account, uint256 amount) returns()
func (_GamePool *GamePoolSession) WithdrawFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.WithdrawFor(&_GamePool.TransactOpts, account, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address account, uint256 amount) returns()
func (_GamePool *GamePoolTransactorSession) WithdrawFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.WithdrawFor(&_GamePool.TransactOpts, account, amount)
}

// GamePoolDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the GamePool contract.
type GamePoolDefaultAdminDelayChangeCanceledIterator struct {
	Event *GamePoolDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *GamePoolDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolDefaultAdminDelayChangeCanceled)
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
		it.Event = new(GamePoolDefaultAdminDelayChangeCanceled)
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
func (it *GamePoolDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the GamePool contract.
type GamePoolDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_GamePool *GamePoolFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*GamePoolDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &GamePoolDefaultAdminDelayChangeCanceledIterator{contract: _GamePool.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_GamePool *GamePoolFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *GamePoolDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolDefaultAdminDelayChangeCanceled)
				if err := _GamePool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*GamePoolDefaultAdminDelayChangeCanceled, error) {
	event := new(GamePoolDefaultAdminDelayChangeCanceled)
	if err := _GamePool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the GamePool contract.
type GamePoolDefaultAdminDelayChangeScheduledIterator struct {
	Event *GamePoolDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *GamePoolDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolDefaultAdminDelayChangeScheduled)
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
		it.Event = new(GamePoolDefaultAdminDelayChangeScheduled)
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
func (it *GamePoolDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the GamePool contract.
type GamePoolDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_GamePool *GamePoolFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*GamePoolDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &GamePoolDefaultAdminDelayChangeScheduledIterator{contract: _GamePool.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_GamePool *GamePoolFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *GamePoolDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolDefaultAdminDelayChangeScheduled)
				if err := _GamePool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*GamePoolDefaultAdminDelayChangeScheduled, error) {
	event := new(GamePoolDefaultAdminDelayChangeScheduled)
	if err := _GamePool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the GamePool contract.
type GamePoolDefaultAdminTransferCanceledIterator struct {
	Event *GamePoolDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *GamePoolDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolDefaultAdminTransferCanceled)
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
		it.Event = new(GamePoolDefaultAdminTransferCanceled)
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
func (it *GamePoolDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the GamePool contract.
type GamePoolDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_GamePool *GamePoolFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*GamePoolDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &GamePoolDefaultAdminTransferCanceledIterator{contract: _GamePool.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_GamePool *GamePoolFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *GamePoolDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolDefaultAdminTransferCanceled)
				if err := _GamePool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*GamePoolDefaultAdminTransferCanceled, error) {
	event := new(GamePoolDefaultAdminTransferCanceled)
	if err := _GamePool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the GamePool contract.
type GamePoolDefaultAdminTransferScheduledIterator struct {
	Event *GamePoolDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *GamePoolDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolDefaultAdminTransferScheduled)
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
		it.Event = new(GamePoolDefaultAdminTransferScheduled)
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
func (it *GamePoolDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the GamePool contract.
type GamePoolDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_GamePool *GamePoolFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*GamePoolDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolDefaultAdminTransferScheduledIterator{contract: _GamePool.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_GamePool *GamePoolFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *GamePoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolDefaultAdminTransferScheduled)
				if err := _GamePool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*GamePoolDefaultAdminTransferScheduled, error) {
	event := new(GamePoolDefaultAdminTransferScheduled)
	if err := _GamePool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GamePool contract.
type GamePoolDepositedIterator struct {
	Event *GamePoolDeposited // Event containing the contract specifics and raw log

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
func (it *GamePoolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolDeposited)
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
		it.Event = new(GamePoolDeposited)
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
func (it *GamePoolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolDeposited represents a Deposited event raised by the GamePool contract.
type GamePoolDeposited struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_GamePool *GamePoolFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*GamePoolDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolDepositedIterator{contract: _GamePool.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 amount)
func (_GamePool *GamePoolFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GamePoolDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolDeposited)
				if err := _GamePool.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseDeposited(log types.Log) (*GamePoolDeposited, error) {
	event := new(GamePoolDeposited)
	if err := _GamePool.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GamePool contract.
type GamePoolInitializedIterator struct {
	Event *GamePoolInitialized // Event containing the contract specifics and raw log

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
func (it *GamePoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolInitialized)
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
		it.Event = new(GamePoolInitialized)
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
func (it *GamePoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolInitialized represents a Initialized event raised by the GamePool contract.
type GamePoolInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GamePool *GamePoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*GamePoolInitializedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GamePoolInitializedIterator{contract: _GamePool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GamePool *GamePoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GamePoolInitialized) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolInitialized)
				if err := _GamePool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseInitialized(log types.Log) (*GamePoolInitialized, error) {
	event := new(GamePoolInitialized)
	if err := _GamePool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolMinDepositAmountUpdatedIterator is returned from FilterMinDepositAmountUpdated and is used to iterate over the raw logs and unpacked data for MinDepositAmountUpdated events raised by the GamePool contract.
type GamePoolMinDepositAmountUpdatedIterator struct {
	Event *GamePoolMinDepositAmountUpdated // Event containing the contract specifics and raw log

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
func (it *GamePoolMinDepositAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolMinDepositAmountUpdated)
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
		it.Event = new(GamePoolMinDepositAmountUpdated)
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
func (it *GamePoolMinDepositAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolMinDepositAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolMinDepositAmountUpdated represents a MinDepositAmountUpdated event raised by the GamePool contract.
type GamePoolMinDepositAmountUpdated struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinDepositAmountUpdated is a free log retrieval operation binding the contract event 0x5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329.
//
// Solidity: event MinDepositAmountUpdated(uint256 oldAmount, uint256 newAmount)
func (_GamePool *GamePoolFilterer) FilterMinDepositAmountUpdated(opts *bind.FilterOpts) (*GamePoolMinDepositAmountUpdatedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "MinDepositAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &GamePoolMinDepositAmountUpdatedIterator{contract: _GamePool.contract, event: "MinDepositAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchMinDepositAmountUpdated is a free log subscription operation binding the contract event 0x5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329.
//
// Solidity: event MinDepositAmountUpdated(uint256 oldAmount, uint256 newAmount)
func (_GamePool *GamePoolFilterer) WatchMinDepositAmountUpdated(opts *bind.WatchOpts, sink chan<- *GamePoolMinDepositAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "MinDepositAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolMinDepositAmountUpdated)
				if err := _GamePool.contract.UnpackLog(event, "MinDepositAmountUpdated", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseMinDepositAmountUpdated(log types.Log) (*GamePoolMinDepositAmountUpdated, error) {
	event := new(GamePoolMinDepositAmountUpdated)
	if err := _GamePool.contract.UnpackLog(event, "MinDepositAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GamePool contract.
type GamePoolPausedIterator struct {
	Event *GamePoolPaused // Event containing the contract specifics and raw log

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
func (it *GamePoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolPaused)
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
		it.Event = new(GamePoolPaused)
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
func (it *GamePoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolPaused represents a Paused event raised by the GamePool contract.
type GamePoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GamePool *GamePoolFilterer) FilterPaused(opts *bind.FilterOpts) (*GamePoolPausedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GamePoolPausedIterator{contract: _GamePool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GamePool *GamePoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GamePoolPaused) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolPaused)
				if err := _GamePool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParsePaused(log types.Log) (*GamePoolPaused, error) {
	event := new(GamePoolPaused)
	if err := _GamePool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolPoolStatusChangedIterator is returned from FilterPoolStatusChanged and is used to iterate over the raw logs and unpacked data for PoolStatusChanged events raised by the GamePool contract.
type GamePoolPoolStatusChangedIterator struct {
	Event *GamePoolPoolStatusChanged // Event containing the contract specifics and raw log

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
func (it *GamePoolPoolStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolPoolStatusChanged)
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
		it.Event = new(GamePoolPoolStatusChanged)
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
func (it *GamePoolPoolStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolPoolStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolPoolStatusChanged represents a PoolStatusChanged event raised by the GamePool contract.
type GamePoolPoolStatusChanged struct {
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolStatusChanged is a free log retrieval operation binding the contract event 0xc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb5966.
//
// Solidity: event PoolStatusChanged(uint8 oldStatus, uint8 newStatus)
func (_GamePool *GamePoolFilterer) FilterPoolStatusChanged(opts *bind.FilterOpts) (*GamePoolPoolStatusChangedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "PoolStatusChanged")
	if err != nil {
		return nil, err
	}
	return &GamePoolPoolStatusChangedIterator{contract: _GamePool.contract, event: "PoolStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPoolStatusChanged is a free log subscription operation binding the contract event 0xc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb5966.
//
// Solidity: event PoolStatusChanged(uint8 oldStatus, uint8 newStatus)
func (_GamePool *GamePoolFilterer) WatchPoolStatusChanged(opts *bind.WatchOpts, sink chan<- *GamePoolPoolStatusChanged) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "PoolStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolPoolStatusChanged)
				if err := _GamePool.contract.UnpackLog(event, "PoolStatusChanged", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParsePoolStatusChanged(log types.Log) (*GamePoolPoolStatusChanged, error) {
	event := new(GamePoolPoolStatusChanged)
	if err := _GamePool.contract.UnpackLog(event, "PoolStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRewardClaimFailedIterator is returned from FilterRewardClaimFailed and is used to iterate over the raw logs and unpacked data for RewardClaimFailed events raised by the GamePool contract.
type GamePoolRewardClaimFailedIterator struct {
	Event *GamePoolRewardClaimFailed // Event containing the contract specifics and raw log

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
func (it *GamePoolRewardClaimFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRewardClaimFailed)
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
		it.Event = new(GamePoolRewardClaimFailed)
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
func (it *GamePoolRewardClaimFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRewardClaimFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRewardClaimFailed represents a RewardClaimFailed event raised by the GamePool contract.
type GamePoolRewardClaimFailed struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimFailed is a free log retrieval operation binding the contract event 0x0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d91.
//
// Solidity: event RewardClaimFailed(address indexed account, address indexed token, uint256 amount)
func (_GamePool *GamePoolFilterer) FilterRewardClaimFailed(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*GamePoolRewardClaimFailedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RewardClaimFailed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRewardClaimFailedIterator{contract: _GamePool.contract, event: "RewardClaimFailed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimFailed is a free log subscription operation binding the contract event 0x0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d91.
//
// Solidity: event RewardClaimFailed(address indexed account, address indexed token, uint256 amount)
func (_GamePool *GamePoolFilterer) WatchRewardClaimFailed(opts *bind.WatchOpts, sink chan<- *GamePoolRewardClaimFailed, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RewardClaimFailed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRewardClaimFailed)
				if err := _GamePool.contract.UnpackLog(event, "RewardClaimFailed", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseRewardClaimFailed(log types.Log) (*GamePoolRewardClaimFailed, error) {
	event := new(GamePoolRewardClaimFailed)
	if err := _GamePool.contract.UnpackLog(event, "RewardClaimFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the GamePool contract.
type GamePoolRewardClaimedIterator struct {
	Event *GamePoolRewardClaimed // Event containing the contract specifics and raw log

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
func (it *GamePoolRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRewardClaimed)
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
		it.Event = new(GamePoolRewardClaimed)
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
func (it *GamePoolRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRewardClaimed represents a RewardClaimed event raised by the GamePool contract.
type GamePoolRewardClaimed struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed account, address indexed token, uint256 amount)
func (_GamePool *GamePoolFilterer) FilterRewardClaimed(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*GamePoolRewardClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RewardClaimed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRewardClaimedIterator{contract: _GamePool.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed account, address indexed token, uint256 amount)
func (_GamePool *GamePoolFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *GamePoolRewardClaimed, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RewardClaimed", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRewardClaimed)
				if err := _GamePool.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseRewardClaimed(log types.Log) (*GamePoolRewardClaimed, error) {
	event := new(GamePoolRewardClaimed)
	if err := _GamePool.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GamePool contract.
type GamePoolRoleAdminChangedIterator struct {
	Event *GamePoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GamePoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRoleAdminChanged)
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
		it.Event = new(GamePoolRoleAdminChanged)
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
func (it *GamePoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRoleAdminChanged represents a RoleAdminChanged event raised by the GamePool contract.
type GamePoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GamePool *GamePoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GamePoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRoleAdminChangedIterator{contract: _GamePool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GamePool *GamePoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GamePoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRoleAdminChanged)
				if err := _GamePool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseRoleAdminChanged(log types.Log) (*GamePoolRoleAdminChanged, error) {
	event := new(GamePoolRoleAdminChanged)
	if err := _GamePool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GamePool contract.
type GamePoolRoleGrantedIterator struct {
	Event *GamePoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *GamePoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRoleGranted)
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
		it.Event = new(GamePoolRoleGranted)
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
func (it *GamePoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRoleGranted represents a RoleGranted event raised by the GamePool contract.
type GamePoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GamePool *GamePoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GamePoolRoleGrantedIterator, error) {

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

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRoleGrantedIterator{contract: _GamePool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GamePool *GamePoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GamePoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRoleGranted)
				if err := _GamePool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseRoleGranted(log types.Log) (*GamePoolRoleGranted, error) {
	event := new(GamePoolRoleGranted)
	if err := _GamePool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GamePool contract.
type GamePoolRoleRevokedIterator struct {
	Event *GamePoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GamePoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRoleRevoked)
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
		it.Event = new(GamePoolRoleRevoked)
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
func (it *GamePoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRoleRevoked represents a RoleRevoked event raised by the GamePool contract.
type GamePoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GamePool *GamePoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GamePoolRoleRevokedIterator, error) {

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

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRoleRevokedIterator{contract: _GamePool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GamePool *GamePoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GamePoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRoleRevoked)
				if err := _GamePool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseRoleRevoked(log types.Log) (*GamePoolRoleRevoked, error) {
	event := new(GamePoolRoleRevoked)
	if err := _GamePool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRoundCancelledIterator is returned from FilterRoundCancelled and is used to iterate over the raw logs and unpacked data for RoundCancelled events raised by the GamePool contract.
type GamePoolRoundCancelledIterator struct {
	Event *GamePoolRoundCancelled // Event containing the contract specifics and raw log

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
func (it *GamePoolRoundCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRoundCancelled)
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
		it.Event = new(GamePoolRoundCancelled)
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
func (it *GamePoolRoundCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRoundCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRoundCancelled represents a RoundCancelled event raised by the GamePool contract.
type GamePoolRoundCancelled struct {
	RoundId      *big.Int
	Recipient    common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRoundCancelled is a free log retrieval operation binding the contract event 0x5cb7f91cf9cbcf0d6e29da784696d773d49366d9b704d703ccefc9c2f61c88b9.
//
// Solidity: event RoundCancelled(uint256 indexed roundId, address indexed recipient, uint256 refundAmount)
func (_GamePool *GamePoolFilterer) FilterRoundCancelled(opts *bind.FilterOpts, roundId []*big.Int, recipient []common.Address) (*GamePoolRoundCancelledIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RoundCancelled", roundIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRoundCancelledIterator{contract: _GamePool.contract, event: "RoundCancelled", logs: logs, sub: sub}, nil
}

// WatchRoundCancelled is a free log subscription operation binding the contract event 0x5cb7f91cf9cbcf0d6e29da784696d773d49366d9b704d703ccefc9c2f61c88b9.
//
// Solidity: event RoundCancelled(uint256 indexed roundId, address indexed recipient, uint256 refundAmount)
func (_GamePool *GamePoolFilterer) WatchRoundCancelled(opts *bind.WatchOpts, sink chan<- *GamePoolRoundCancelled, roundId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RoundCancelled", roundIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRoundCancelled)
				if err := _GamePool.contract.UnpackLog(event, "RoundCancelled", log); err != nil {
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

// ParseRoundCancelled is a log parse operation binding the contract event 0x5cb7f91cf9cbcf0d6e29da784696d773d49366d9b704d703ccefc9c2f61c88b9.
//
// Solidity: event RoundCancelled(uint256 indexed roundId, address indexed recipient, uint256 refundAmount)
func (_GamePool *GamePoolFilterer) ParseRoundCancelled(log types.Log) (*GamePoolRoundCancelled, error) {
	event := new(GamePoolRoundCancelled)
	if err := _GamePool.contract.UnpackLog(event, "RoundCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRoundCreatedIterator is returned from FilterRoundCreated and is used to iterate over the raw logs and unpacked data for RoundCreated events raised by the GamePool contract.
type GamePoolRoundCreatedIterator struct {
	Event *GamePoolRoundCreated // Event containing the contract specifics and raw log

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
func (it *GamePoolRoundCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRoundCreated)
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
		it.Event = new(GamePoolRoundCreated)
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
func (it *GamePoolRoundCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRoundCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRoundCreated represents a RoundCreated event raised by the GamePool contract.
type GamePoolRoundCreated struct {
	RoundId        *big.Int
	Creator        common.Address
	TotalReward    *big.Int
	StartBlock     *big.Int
	EndBlock       *big.Int
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRoundCreated is a free log retrieval operation binding the contract event 0xba1778056db43899781dd248e1cab61b0f655af9ded9396f782b752e73de8989.
//
// Solidity: event RoundCreated(uint256 indexed roundId, address indexed creator, uint256 totalReward, uint256 startBlock, uint256 endBlock, uint256 rewardPerBlock)
func (_GamePool *GamePoolFilterer) FilterRoundCreated(opts *bind.FilterOpts, roundId []*big.Int, creator []common.Address) (*GamePoolRoundCreatedIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RoundCreated", roundIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolRoundCreatedIterator{contract: _GamePool.contract, event: "RoundCreated", logs: logs, sub: sub}, nil
}

// WatchRoundCreated is a free log subscription operation binding the contract event 0xba1778056db43899781dd248e1cab61b0f655af9ded9396f782b752e73de8989.
//
// Solidity: event RoundCreated(uint256 indexed roundId, address indexed creator, uint256 totalReward, uint256 startBlock, uint256 endBlock, uint256 rewardPerBlock)
func (_GamePool *GamePoolFilterer) WatchRoundCreated(opts *bind.WatchOpts, sink chan<- *GamePoolRoundCreated, roundId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RoundCreated", roundIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRoundCreated)
				if err := _GamePool.contract.UnpackLog(event, "RoundCreated", log); err != nil {
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

// ParseRoundCreated is a log parse operation binding the contract event 0xba1778056db43899781dd248e1cab61b0f655af9ded9396f782b752e73de8989.
//
// Solidity: event RoundCreated(uint256 indexed roundId, address indexed creator, uint256 totalReward, uint256 startBlock, uint256 endBlock, uint256 rewardPerBlock)
func (_GamePool *GamePoolFilterer) ParseRoundCreated(log types.Log) (*GamePoolRoundCreated, error) {
	event := new(GamePoolRoundCreated)
	if err := _GamePool.contract.UnpackLog(event, "RoundCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolRoundsSyncedIterator is returned from FilterRoundsSynced and is used to iterate over the raw logs and unpacked data for RoundsSynced events raised by the GamePool contract.
type GamePoolRoundsSyncedIterator struct {
	Event *GamePoolRoundsSynced // Event containing the contract specifics and raw log

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
func (it *GamePoolRoundsSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolRoundsSynced)
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
		it.Event = new(GamePoolRoundsSynced)
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
func (it *GamePoolRoundsSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolRoundsSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolRoundsSynced represents a RoundsSynced event raised by the GamePool contract.
type GamePoolRoundsSynced struct {
	Processed *big.Int
	Removed   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRoundsSynced is a free log retrieval operation binding the contract event 0x43cfb7702c1541a40272201ce7f99c71e1f1101be966dbcc218a3da1bbebc6e1.
//
// Solidity: event RoundsSynced(uint256 processed, uint256 removed)
func (_GamePool *GamePoolFilterer) FilterRoundsSynced(opts *bind.FilterOpts) (*GamePoolRoundsSyncedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "RoundsSynced")
	if err != nil {
		return nil, err
	}
	return &GamePoolRoundsSyncedIterator{contract: _GamePool.contract, event: "RoundsSynced", logs: logs, sub: sub}, nil
}

// WatchRoundsSynced is a free log subscription operation binding the contract event 0x43cfb7702c1541a40272201ce7f99c71e1f1101be966dbcc218a3da1bbebc6e1.
//
// Solidity: event RoundsSynced(uint256 processed, uint256 removed)
func (_GamePool *GamePoolFilterer) WatchRoundsSynced(opts *bind.WatchOpts, sink chan<- *GamePoolRoundsSynced) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "RoundsSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolRoundsSynced)
				if err := _GamePool.contract.UnpackLog(event, "RoundsSynced", log); err != nil {
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

// ParseRoundsSynced is a log parse operation binding the contract event 0x43cfb7702c1541a40272201ce7f99c71e1f1101be966dbcc218a3da1bbebc6e1.
//
// Solidity: event RoundsSynced(uint256 processed, uint256 removed)
func (_GamePool *GamePoolFilterer) ParseRoundsSynced(log types.Log) (*GamePoolRoundsSynced, error) {
	event := new(GamePoolRoundsSynced)
	if err := _GamePool.contract.UnpackLog(event, "RoundsSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolTokensReclaimedIterator is returned from FilterTokensReclaimed and is used to iterate over the raw logs and unpacked data for TokensReclaimed events raised by the GamePool contract.
type GamePoolTokensReclaimedIterator struct {
	Event *GamePoolTokensReclaimed // Event containing the contract specifics and raw log

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
func (it *GamePoolTokensReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolTokensReclaimed)
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
		it.Event = new(GamePoolTokensReclaimed)
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
func (it *GamePoolTokensReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolTokensReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolTokensReclaimed represents a TokensReclaimed event raised by the GamePool contract.
type GamePoolTokensReclaimed struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensReclaimed is a free log retrieval operation binding the contract event 0x6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464.
//
// Solidity: event TokensReclaimed(address indexed token, address indexed to, uint256 amount)
func (_GamePool *GamePoolFilterer) FilterTokensReclaimed(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GamePoolTokensReclaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "TokensReclaimed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolTokensReclaimedIterator{contract: _GamePool.contract, event: "TokensReclaimed", logs: logs, sub: sub}, nil
}

// WatchTokensReclaimed is a free log subscription operation binding the contract event 0x6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464.
//
// Solidity: event TokensReclaimed(address indexed token, address indexed to, uint256 amount)
func (_GamePool *GamePoolFilterer) WatchTokensReclaimed(opts *bind.WatchOpts, sink chan<- *GamePoolTokensReclaimed, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "TokensReclaimed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolTokensReclaimed)
				if err := _GamePool.contract.UnpackLog(event, "TokensReclaimed", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseTokensReclaimed(log types.Log) (*GamePoolTokensReclaimed, error) {
	event := new(GamePoolTokensReclaimed)
	if err := _GamePool.contract.UnpackLog(event, "TokensReclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GamePool contract.
type GamePoolUnpausedIterator struct {
	Event *GamePoolUnpaused // Event containing the contract specifics and raw log

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
func (it *GamePoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolUnpaused)
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
		it.Event = new(GamePoolUnpaused)
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
func (it *GamePoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolUnpaused represents a Unpaused event raised by the GamePool contract.
type GamePoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GamePool *GamePoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GamePoolUnpausedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GamePoolUnpausedIterator{contract: _GamePool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GamePool *GamePoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GamePoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolUnpaused)
				if err := _GamePool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseUnpaused(log types.Log) (*GamePoolUnpaused, error) {
	event := new(GamePoolUnpaused)
	if err := _GamePool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GamePool contract.
type GamePoolUpgradedIterator struct {
	Event *GamePoolUpgraded // Event containing the contract specifics and raw log

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
func (it *GamePoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolUpgraded)
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
		it.Event = new(GamePoolUpgraded)
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
func (it *GamePoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolUpgraded represents a Upgraded event raised by the GamePool contract.
type GamePoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GamePool *GamePoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GamePoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolUpgradedIterator{contract: _GamePool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GamePool *GamePoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GamePoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolUpgraded)
				if err := _GamePool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseUpgraded(log types.Log) (*GamePoolUpgraded, error) {
	event := new(GamePoolUpgraded)
	if err := _GamePool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GamePoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GamePool contract.
type GamePoolWithdrawnIterator struct {
	Event *GamePoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *GamePoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolWithdrawn)
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
		it.Event = new(GamePoolWithdrawn)
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
func (it *GamePoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolWithdrawn represents a Withdrawn event raised by the GamePool contract.
type GamePoolWithdrawn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_GamePool *GamePoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*GamePoolWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &GamePoolWithdrawnIterator{contract: _GamePool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed account, uint256 amount)
func (_GamePool *GamePoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GamePoolWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolWithdrawn)
				if err := _GamePool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_GamePool *GamePoolFilterer) ParseWithdrawn(log types.Log) (*GamePoolWithdrawn, error) {
	event := new(GamePoolWithdrawn)
	if err := _GamePool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
