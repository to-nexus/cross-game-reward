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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPONSOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"cancelRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"cancelRoundToRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimRewardFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimRewardsFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"createRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationBlocks\",\"type\":\"uint256\"}],\"name\":\"createRoundFromReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossGameReward\",\"outputs\":[{\"internalType\":\"contractICrossGameReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRoundCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRoundIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainderReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"}],\"internalType\":\"structIGamePool.Round[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getReclaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemovedRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRemovedTokenRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRewardToken\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reclaimableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRemoved\",\"type\":\"bool\"}],\"internalType\":\"structICrossGameRewardPool.RewardToken\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainderReward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"}],\"internalType\":\"structIGamePool.Round\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalAccRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRemovedRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxActiveRounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolStatus\",\"outputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reclaimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reclaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removedRewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"rewardTokenAt\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxActiveRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setPoolStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRounds\",\"type\":\"uint256\"}],\"name\":\"syncRounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"processed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"MaxActiveRoundsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MinDepositAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"PoolStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"RoundCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"RoundCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"}],\"name\":\"RoundsSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensReclaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"GPBelowMinimumDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPCanNotZeroValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumICrossGameRewardPool.PoolStatus\",\"name\":\"currentStatus\",\"type\":\"uint8\"}],\"name\":\"GPDepositNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"GPInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPInvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPInvalidMaxActiveRounds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provided\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"GPInvalidRewardToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"GPInvalidStartBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"GPMaxActiveRoundsReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GPNoDepositFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPNoReclaimableAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPNotAllowedInCurrentState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPOnlyRewardRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GPOnlyRoundCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPOnlyRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPPoolStatusUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPRewardIsDepositToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPRewardPerBlockZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"GPRoundAlreadyCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"GPRoundAlreadyStarted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"GPRoundNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GPSyncNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"}]",
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
		"14cfd856": "maxActiveRounds()",
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
		"5a6a0153": "setMaxActiveRounds(uint256)",
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
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161563c6100f95f395f81816137160152818161373f015261390f015261563c5ff3fe608060405260043610610483575f3560e01c806378ad8c7d11610251578063b5fd73f81161013c578063d547741f116100b7578063f4e2474011610087578063f7c618c11161006d578063f7c618c114610e79578063fd8bdc6814610e98578063ff50abdc14610ead575f5ffd5b8063f4e2474014610e3b578063f665336e14610e5a575f5ffd5b8063d547741f14610dc3578063d602b9fd14610de2578063db518db214610df6578063f022869214610e15575f5ffd5b8063c89039c51161010c578063cefc1429116100f2578063cefc142914610d26578063cf6eefb714610d3a578063d279c19114610da4575f5ffd5b8063c89039c514610cf3578063cc8463c814610d12575f5ffd5b8063b5fd73f814610c5f578063b6b55f2514610c8d578063c2d7944414610cac578063c4f59f9b14610cdf575f5ffd5b806391d14854116101cc578063a1eda53c1161019c578063a980356a11610182578063a980356a14610bd8578063abb06b9514610bf7578063ad3cb1cc14610c0a575f5ffd5b8063a1eda53c14610b92578063a217fddf14610bc5575f5ffd5b806391d1485414610ab45780639b80c3f214610b175780639ced7e7614610b3f5780639df4496b14610b5e575f5ffd5b806384780205116102215780638da5cb5b116102075780638da5cb5b14610a605780638f1327c014610a7457806391cf6d3e14610aa0575f5ffd5b80638478020514610a0557806384ef8ffc14610a24575f5ffd5b806378ad8c7d1461097a57806379f5ecb71461098e5780637d984d5f146109c55780637e07ab09146109e6575f5ffd5b806335c21d5d116103715780635a6a0153116102ec578063649a5ec7116102bc5780636d7c49a2116102a25780636d7c49a2146108c55780636fb7a4e8146108e45780637707887214610905575f5ffd5b8063649a5ec71461088757806369b65e91146108a6575f5ffd5b80635a6a0153146107fe5780635c975abb1461081d578063634e93da14610853578063645006ca14610872575f5ffd5b80633d509c97116103415780634d1cd014116103275780634d1cd014146107b85780634f1ef286146107d757806352d1902d146107ea575f5ffd5b80633d509c97146107845780634002eda6146107a3575f5ffd5b806335c21d5d1461071357806335c30fda1461073257806336568abe14610751578063372500ab14610770575f5ffd5b80631efed5f7116104015780632e1a7d4d116103d15780632f4f21e2116103b75780632f4f21e2146106c357806331d7a262146106e25780633548237914610701575f5ffd5b80632e1a7d4d146106855780632f2ff15d146106a4575f5ffd5b80631efed5f7146105d9578063248a9ca3146105f857806327e235e3146106455780632dbea37b14610670575f5ffd5b80631794bb3c116104565780631ac6d19d1161043c5780631ac6d19d1461055a5780631af8acec146105795780631c03e6cc146105ba575f5ffd5b80631794bb3c1461051c57806317bd64eb1461053b575f5ffd5b806301ffc9a714610487578063022d63fb146104bb5780630aa6220b146104e357806314cfd856146104f9575b5f5ffd5b348015610492575f5ffd5b506104a66104a1366004614ebe565b610ec2565b60405190151581526020015b60405180910390f35b3480156104c6575f5ffd5b50620697805b60405165ffffffffffff90911681526020016104b2565b3480156104ee575f5ffd5b506104f7610f1d565b005b348015610504575f5ffd5b5061050e60105481565b6040519081526020016104b2565b348015610527575f5ffd5b506104f7610536366004614f11565b610f32565b348015610546575f5ffd5b506104f7610555366004614f4f565b611246565b348015610565575f5ffd5b506104f7610574366004614f7d565b61145b565b348015610584575f5ffd5b506105ac610593366004614f7d565b50604080515f8082526020820190815281830190925291565b6040516104b292919061500b565b3480156105c5575f5ffd5b506104f76105d4366004614f7d565b611485565b3480156105e4575f5ffd5b5061050e6105f336600461502f565b611566565b348015610603575f5ffd5b5061050e610612366004615058565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b348015610650575f5ffd5b5061050e61065f366004614f7d565b60056020525f908152604090205481565b34801561067b575f5ffd5b5061050e600c5481565b348015610690575f5ffd5b506104f761069f366004615058565b61157d565b3480156106af575f5ffd5b506104f76106be366004614f4f565b611598565b3480156106ce575f5ffd5b506104f76106dd36600461506f565b6115d9565b3480156106ed575f5ffd5b506105ac6106fc366004614f7d565b61165d565b34801561070c575f5ffd5b505f61050e565b34801561071e575f5ffd5b5061050e61072d366004614f7d565b61170d565b34801561073d575f5ffd5b506104f761074c366004615099565b611733565b34801561075c575f5ffd5b506104f761076b366004614f4f565b6117b5565b34801561077b575f5ffd5b506104f7611902565b34801561078f575f5ffd5b506104f761079e366004614f7d565b611925565b3480156107ae575f5ffd5b5061050e60085481565b3480156107c3575f5ffd5b506104f76107d2366004615099565b6119f1565b6104f76107e53660046150f2565b611b90565b3480156107f5575f5ffd5b5061050e611bab565b348015610809575f5ffd5b506104f7610818366004615058565b611bd9565b348015610828575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166104a6565b34801561085e575f5ffd5b506104f761086d366004614f7d565b611c97565b34801561087d575f5ffd5b5061050e60045481565b348015610892575f5ffd5b506104f76108a13660046151f5565b611caa565b3480156108b1575f5ffd5b5061050e6108c036600461521a565b611cbd565b3480156108d0575f5ffd5b506104f76108df366004615252565b612049565b3480156108ef575f5ffd5b506108f86121f4565b6040516104b29190615270565b348015610910575f5ffd5b5061092461091f366004614f7d565b612205565b6040516104b291905f60c0820190506001600160a01b0383511682526020830151602083015260408301516040830152606083015160608301526080830151608083015260a0830151151560a083015292915050565b348015610985575f5ffd5b5061050e612371565b348015610999575f5ffd5b506109ad6109a8366004615058565b61237c565b6040516001600160a01b0390911681526020016104b2565b3480156109d0575f5ffd5b506109d96123f6565b6040516104b291906152fe565b3480156109f1575f5ffd5b506104f7610a00366004615058565b61257c565b348015610a10575f5ffd5b506104f7610a1f366004615058565b612586565b348015610a2f575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b03166109ad565b348015610a6b575f5ffd5b506109ad612644565b348015610a7f575f5ffd5b50610a93610a8e366004615058565b61264d565b6040516104b2919061534c565b348015610aab575f5ffd5b5061050e5f5481565b348015610abf575f5ffd5b506104a6610ace366004614f4f565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610b22575f5ffd5b50604080515f8152602081019091525b6040516104b2919061535b565b348015610b4a575f5ffd5b5061050e610b59366004615099565b61277c565b348015610b69575f5ffd5b50610b7d610b78366004615058565b6127e7565b604080519283526020830191909152016104b2565b348015610b9d575f5ffd5b50610ba661292c565b6040805165ffffffffffff9384168152929091166020830152016104b2565b348015610bd0575f5ffd5b5061050e5f81565b348015610be3575f5ffd5b50610b7d610bf2366004615099565b6129e9565b348015610c02575f5ffd5b50600161050e565b348015610c15575f5ffd5b50610c526040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516104b2919061536d565b348015610c6a575f5ffd5b506104a6610c79366004614f7d565b6002546001600160a01b0391821691161490565b348015610c98575f5ffd5b506104f7610ca7366004615058565b612ab2565b348015610cb7575f5ffd5b5061050e7f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a81565b348015610cea575f5ffd5b50610b32612b25565b348015610cfe575f5ffd5b506001546109ad906001600160a01b031681565b348015610d1d575f5ffd5b506104cc612b86565b348015610d31575f5ffd5b506104f7612c66565b348015610d45575f5ffd5b507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840054604080516001600160a01b03831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016104b2565b348015610daf575f5ffd5b506104f7610dbe366004614f7d565b612cd4565b348015610dce575f5ffd5b506104f7610ddd366004614f4f565b612d4d565b348015610ded575f5ffd5b506104f7612d8e565b348015610e01575f5ffd5b506104f7610e1036600461506f565b612da0565b348015610e20575f5ffd5b50600754610e2e9060ff1681565b6040516104b29190615426565b348015610e46575f5ffd5b506003546109ad906001600160a01b031681565b348015610e65575f5ffd5b506104a6610e74366004614f7d565b505f90565b348015610e84575f5ffd5b506002546109ad906001600160a01b031681565b348015610ea3575f5ffd5b5061050e600f5481565b348015610eb8575f5ffd5b5061050e60065481565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610f175750610f1782612dc4565b92915050565b5f610f2781612e5a565b610f2f612e64565b50565b5f610f3b612e6e565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015610f675750825b90505f8267ffffffffffffffff166001148015610f835750303b155b905081158015610f91575080155b15610fc8576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156110295784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816611069576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0387166110a9576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316886001600160a01b0316036110f4576040517f4fdaf28b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f861161112d576040517fe3dc980700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001633908117909155611165905f90612e96565b61116d612ea8565b611175612ea8565b61117d612ea8565b435f55600180546001600160a01b038a81167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617835560028054918b169190921617905560048790556007805460ff191690556008556032601055831561123c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b61124e612eb0565b5f828152600960205260409020805483906112a1576040517f50727ff000000000000000000000000000000000000000000000000000000000815260040161129891815260200190565b60405180910390fd5b506009810154839060ff16156112e6576040517f0eb9cf3d00000000000000000000000000000000000000000000000000000000815260040161129891815260200190565b506001810154839033906001600160a01b031681811461134c576040517f7bd7574700000000000000000000000000000000000000000000000000000000815260048101939093526001600160a01b039182166024840152166044820152606401611298565b5050506001600160a01b03821661138f576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060030154431083906113d1576040517f6ce0091d00000000000000000000000000000000000000000000000000000000815260040161129891815260200190565b5060098101805460ff191660011790556113ec600a84612f36565b506002808201549054611409906001600160a01b03168483612f41565b826001600160a01b0316847f5cb7f91cf9cbcf0d6e29da784696d773d49366d9b704d703ccefc9c2f61c88b98360405161144591815260200190565b60405180910390a35050611457612fb5565b5050565b611463612eb0565b61146b612fdf565b6114748161303b565b61147d81613139565b610f2f612fb5565b6003546001600160a01b031633146114c9576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b03828116911614610f2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f43616e6e6f742061646420646966666572656e742072657761726420746f6b6560448201527f6e000000000000000000000000000000000000000000000000000000000000006064820152608401611298565b5f61157333858585611cbd565b90505b9392505050565b611585612eb0565b61158d612fdf565b61147d3333836131cf565b816115cf576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61145782826133aa565b6115e1612eb0565b6115e9612fdf565b5f60075460ff166002811115611601576116016153c0565b60075460ff169114611640576040517f2cd1147f0000000000000000000000000000000000000000000000000000000081526004016112989190615426565b5061164a8261303b565b6116553383836133f3565b611457612fb5565b604080516001808252818301909252606091829190602080830190803683375050604080516001808252818301909252929450905060208083019080368337505060025484519293506001600160a01b0316918491505f906116c1576116c1615434565b60200260200101906001600160a01b031690816001600160a01b0316815250506116ea8361352a565b815f815181106116fc576116fc615434565b602002602001018181525050915091565b6002545f906001600160a01b0383811691161461172b57505f919050565b5050600f5490565b61173b612eb0565b611743612fdf565b60025481906001600160a01b0390811690821681146117a1576040517f661779680000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015291166024820152604401611298565b50506117ac8261303b565b61165582613139565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008215801561181057507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038381169116145b156118f3577feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611876575065ffffffffffff8116155b8061188957504265ffffffffffff821610155b156118ca576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401611298565b505080547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1681555b6118fd83836136bf565b505050565b61190a612eb0565b611912612fdf565b61191b33613139565b611923612fb5565b565b6003546001600160a01b03163314611969576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f47616d65506f6f6c20646f6573206e6f7420737570706f72742072656d6f766960448201527f6e672072657761726420746f6b656e73000000000000000000000000000000006064820152608401611298565b6003546001600160a01b03163314611a35576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b03838116911614611aac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420746f6b656e000000000000000000000000000000000000006044820152606401611298565b5f600f5411611ae7576040517f4f1aa73b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116611b27576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f80545f909155600254611b46906001600160a01b03168383612f41565b6002546040518281526001600160a01b038481169216907f6a5e278fe27e73fb0093ca72181eb6eaff00da814a3dd5e4ca3d618e23951464906020015b60405180910390a3505050565b611b9861370b565b611ba1826137db565b61145782826137e5565b5f611bb4613904565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6003546001600160a01b03163314611c1d576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8111611c56576040517f0bab4bd700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60105460408051918252602082018390527f0528d77b9faf6330e42880ad6b8266236ed5d1bd23a917f5f6c21eecc7ed00c6910160405180910390a1601055565b5f611ca181612e5a565b61145782613966565b5f611cb481612e5a565b611457826139d8565b5f611cc6612eb0565b611cce612fdf565b7f1597bc5e34ff090612f53164e4e642d2ab4fc78bffe19ed1b602a0d12559561a611cf881612e5a565b5f8511611d31576040517fe3dc980700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8343808211611d75576040517f1302dfb800000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611298565b50505f8311611db0576040517f7cf30a0300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6010541580611dc95750601054611dc7600a613a40565b105b611dd3600a613a40565b6010549091611e17576040517fef85100800000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611298565b505f90506402540be40080611e2c868961548e565b611e36919061548e565b611e4091906154c6565b90505f8111611e7b576040517f6fdfc85a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611e8685836154c6565b611e9090886154dd565b9050611e9a613a49565b60088054905f611ea9836154f0565b9091555093505f611eba8688615527565b9050604051806101400160405280868152602001336001600160a01b031681526020018981526020018881526020018281526020018481526020018881526020015f81526020018381526020015f151581525060095f8781526020019081526020015f205f820151815f01556020820151816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155610120820151816009015f6101000a81548160ff021916908315150217905550905050611fd085600a613a5290919063ffffffff16565b50600254611fe9906001600160a01b03168a308b613a5d565b604080518981526020810189905290810182905260608101849052339086907fba1778056db43899781dd248e1cab61b0f655af9ded9396f782b752e73de89899060800160405180910390a350505050612041612fb5565b949350505050565b6003546001600160a01b0316331461208d576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460ff168160028111156120a5576120a56153c0565b8160028111156120b7576120b76153c0565b036120ee576040517f6a79d3d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007805483919060ff1916600183600281111561210d5761210d6153c0565b02179055506002826002811115612126576121266153c0565b14801561215557507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16155b1561216757612162613a96565b6121b7565b600282600281111561217b5761217b6153c0565b141580156121aa57507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff165b156121b7576121b7613b0b565b7fc86dbb487587a3e53cea849629dfead70a66cdb484b712da924a034e11cb596681836040516121e892919061553a565b60405180910390a15050565b6060612200600a613b63565b905090565b6122436040518060c001604052805f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b6002546001600160a01b038381169116146122ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420746f6b656e000000000000000000000000000000000000006044820152606401611298565b6040805160c0810182526002546001600160a01b0316808252600c54602083015282517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015291928301916370a0823190602401602060405180830381865afa15801561232f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123539190615555565b8152600f5460208201525f6040820181905260609091015292915050565b5f612200600a613a40565b5f81156123e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c696420696e646578000000000000000000000000000000000000006044820152606401611298565b50506002546001600160a01b031690565b60605f612403600a613a40565b90505f8167ffffffffffffffff81111561241f5761241f6150c5565b6040519080825280602002602001820160405280156124a757816020015b6124946040518061014001604052805f81526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b81526020019060019003908161243d5790505b5090505f5b828110156125755760095f6124c2600a84613b6f565b815260208082019290925260409081015f208151610140810183528154815260018201546001600160a01b0316938101939093526002810154918301919091526003810154606083015260048101546080830152600581015460a0830152600681015460c0830152600781015460e083015260088101546101008301526009015460ff161515610120820152825183908390811061256257612562615434565b60209081029190910101526001016124ac565b5092915050565b610f2f8133611246565b6003546001600160a01b031633146125ca576040517fd89d80bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8111612603576040517fe3dc980700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045460408051918252602082018390527f5fb4589fcdfab8bd40d9776abc10876bb1cb02c0edab28d05cc42869b40e0329910160405180910390a1600455565b5f612200613b7a565b6126a46040518061014001604052805f81526020015f6001600160a01b031681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f151581525090565b5f8281526009602052604090205482906126ed576040517f50727ff000000000000000000000000000000000000000000000000000000000815260040161129891815260200190565b50505f908152600960208181526040928390208351610140810185528154815260018201546001600160a01b0316928101929092526002810154938201939093526003830154606082015260048301546080820152600583015460a0820152600683015460c0820152600783015460e0820152600883015461010082015291015460ff16151561012082015290565b6002545f9082906001600160a01b0390811690821681146127dc576040517f661779680000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015291166024820152604401611298565b50506115768361352a565b335f9081527f1f0283d10d2580ee30b219d86fcd3b74322ee3f926a3fbc0240167eb1c3e73e86020526040812054819060ff16806128a9575060035f9054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612870573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612894919061556c565b6001600160a01b0316336001600160a01b0316145b6128df576040517f8e19b9c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128e883613bac565b60408051838152602081018390529294509092507f43cfb7702c1541a40272201ce7f99c71e1f1101be966dbcc218a3da1bbebc6e1910160405180910390a1915091565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff167feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840081158015906129ae57504265ffffffffffff831610155b6129b9575f5f6129e0565b600181015474010000000000000000000000000000000000000000900465ffffffffffff16825b92509250509091565b6002545f90819083906001600160a01b039081169082168114612a4b576040517f661779680000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015291166024820152604401611298565b50506001600160a01b0384165f908152600560205260409020548015612a91576001600160a01b0385165f908152600d6020526040902054612a8e90829061548e565b92505b50506001600160a01b039092165f908152600e602052604090205491929050565b612aba612eb0565b612ac2612fdf565b5f60075460ff166002811115612ada57612ada6153c0565b60075460ff169114612b19576040517f2cd1147f0000000000000000000000000000000000000000000000000000000081526004016112989190615426565b5061147d3333836133f3565b6040805160018082528183019092526060915f919060208083019080368337505060025482519293506001600160a01b0316918391505f90612b6957612b69615434565b6001600160a01b0390921660209283029190910190910152919050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401545f907feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015801590612c0857504265ffffffffffff8216105b612c395781547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16612c5f565b600182015474010000000000000000000000000000000000000000900465ffffffffffff165b9250505090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400546001600160a01b0316338114612ccc576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401611298565b610f2f613ea2565b612cdc612eb0565b612ce4612fdf565b60025481906001600160a01b039081169082168114612d42576040517f661779680000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015291166024820152604401611298565b505061147d33613139565b81612d84576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114578282613fb7565b5f612d9881612e5a565b610f2f613ffa565b612da8612eb0565b612db0612fdf565b612db98261303b565b6116553383836131cf565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610f1757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610f17565b610f2f8133614004565b6119235f5f614090565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610f17565b612e9e61421b565b6114578282614259565b61192361421b565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c15612f09576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61192360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b90614315565b5f611576838361431c565b6040516001600160a01b038381166024830152604482018390526118fd91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506143ff565b6119235f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00612f30565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611923576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811661307b576040517f5b1b7b4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035f9054906101000a90046001600160a01b03166001600160a01b031663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130cb573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130ef919061556c565b6001600160a01b0316336001600160a01b031614610f2f576040517fb1a7302300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381165f90815260056020908152604080832054600e909252909120548115158061316a57505f81115b83906131ae576040517f500ca0dc0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611298565b5081156131c6576131bd613a49565b6131c683614484565b6118fd83614549565b6001600160a01b0382165f90815260056020526040902054829061322b576040517f500ca0dc0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611298565b505f81156132395781613252565b6001600160a01b0383165f908152600560205260409020545b6001600160a01b0384165f9081526005602052604090205490915081818111156132b1576040517fb22545ac00000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611298565b50506132bb613a49565b6132c483614484565b6132cd83614549565b6001600160a01b0383165f90815260056020526040812080548392906132f49084906154dd565b925050819055508060065f82825461330c91906154dd565b9091555050600c546001600160a01b0384165f9081526005602052604090205461333691906154c6565b6001600160a01b038085165f908152600d602052604090209190915560015461336191168583612f41565b826001600160a01b03167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d58260405161339c91815260200190565b60405180910390a250505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546133e381612e5a565b6133ed8383614630565b50505050565b60045481908082101561343b576040517f26a1f59d00000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611298565b5050613445613a49565b61344e82614484565b600154613466906001600160a01b0316843084613a5d565b6001600160a01b0382165f908152600560205260408120805483929061348d908490615527565b925050819055508060065f8282546134a59190615527565b9091555050600c546001600160a01b0383165f908152600560205260409020546134cf91906154c6565b6001600160a01b0383165f818152600d6020526040908190209290925590517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c49061351d9084815260200190565b60405180910390a2505050565b6001600160a01b0381165f90815260056020908152604080832054600e90925282205481830361355b579392505050565b600c545f613569600a613a40565b90505f5b81811015613667575f613581600a83613b6f565b5f818152600960208190526040909120908101549192509060ff16806135aa5750806003015443105b156135b657505061365f565b806006015443116135c857505061365f565b5f816004015443106135de5781600401546135e0565b435b90505f8260060154826135f391906154dd565b90505f83600501548261360691906154c6565b9050836004015483036136255760088401546136229082615527565b90505b6006541561365957600654613642670de0b6b3a7640000836154c6565b61364c919061548e565b6136569089615527565b97505b50505050505b60010161356d565b506001600160a01b0386165f908152600d6020526040812054670de0b6b3a76400009061369485886154c6565b61369e91906154dd565b6136a8919061548e565b90506136b48185615527565b979650505050505050565b6001600160a01b0381163314613701576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118fd8282614709565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806137a457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166137987f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15611923576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61145781612e5a565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561385d575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261385a91810190615555565b60015b61389e576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611298565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146138fa576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611298565b6118fd838361479f565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611923576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61396f612b86565b613978426147f4565b6139829190615587565b905061398e8282614843565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6139e2826148f0565b6139eb426147f4565b6139f59190615587565b9050613a018282614090565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b91016121e8565b5f610f17825490565b6114575f613bac565b5f6115768383614937565b6040516001600160a01b0384811660248301528381166044830152606482018390526133ed9186918216906323b872dd90608401612f6e565b613a9e612fdf565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b613b13614983565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613aed565b60605f611576836149de565b5f6115768383614a37565b5f6122007feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b5f5f5f613bb9600a613a40565b9050805f03613bcd57505f93849350915050565b5f8415613be757818510613be15781613be9565b84613be9565b815b90505f8167ffffffffffffffff811115613c0557613c056150c5565b604051908082528060200260200182016040528015613c2e578160200160208202803683370190505b5090505f805b83811015613e53575f613c48600a83613b6f565b5f818152600960208190526040909120908101549192509060ff1615613c9957818585613c74816154f0565b965081518110613c8657613c86615434565b6020026020010181815250505050613e4b565b8060030154431015613cac575050613e4b565b80600601544311613cbe575050613e4b565b5f81600401544310613cd4578160040154613cd6565b435b90505f826006015482613ce991906154dd565b90505f836005015482613cfc91906154c6565b905083600401548303613d1b576008840154613d189082615527565b90505b6006545f03613d405780600f5f828254613d359190615527565b90915550613e089050565b6006545f90613d57670de0b6b3a7640000846154c6565b613d61919061548e565b90508015613def5780856007015f828254613d7c9190615527565b9250508190555080600c5f828254613d949190615527565b90915550506006545f90670de0b6b3a764000090613db290846154c6565b613dbc919061548e565b90505f613dc982856154dd565b90508015613de85780600f5f828254613de29190615527565b90915550505b5050613e06565b81600f5f828254613e009190615527565b90915550505b505b6006840183905560048401544310613e4557848888613e26816154f0565b995081518110613e3857613e38615434565b6020026020010181815250505b50505050505b600101613c34565b505f5b81811015613e9457613e8b838281518110613e7357613e73615434565b6020026020010151600a612f3690919063ffffffff16565b50600101613e56565b509196919550909350505050565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080546001600160a01b0381169074010000000000000000000000000000000000000000900465ffffffffffff16801580613f0557504265ffffffffffff821610155b15613f46576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401611298565b613f805f613f7b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b614709565b50613f8b5f83614630565b505081547fffffffffffff00000000000000000000000000000000000000000000000000001690915550565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154613ff081612e5a565b6133ed8383614709565b6119235f5f614843565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611457576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611298565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401547feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680156141a2574265ffffffffffff82161015614179576001820154825479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090910465ffffffffffff167a010000000000000000000000000000000000000000000000000000021782556141a2565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b5060010180546001600160a01b03167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b614223614a5d565b611923576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61426161421b565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984006001600160a01b0382166142c4576040517fc22c80220000000000000000000000000000000000000000000000000000000081525f6004820152602401611298565b805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167a01000000000000000000000000000000000000000000000000000065ffffffffffff8516021781556133ed5f83614630565b80825d5050565b5f81815260018301602052604081205480156143f6575f61433e6001836154dd565b85549091505f90614351906001906154dd565b90508082146143b0575f865f01828154811061436f5761436f615434565b905f5260205f200154905080875f01848154811061438f5761438f615434565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806143c1576143c16155a5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610f17565b5f915050610f17565b5f5f60205f8451602086015f885af18061441e576040513d5f823e3d81fd5b50505f513d91508115614435578060011415614442565b6001600160a01b0384163b155b156133ed576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611298565b6001600160a01b0381165f90815260056020526040902054801561451d576001600160a01b0382165f908152600d6020526040812054600c54670de0b6b3a764000091906144d290856154c6565b6144dc91906154dd565b6144e6919061548e565b9050801561451b576001600160a01b0383165f908152600e602052604081208054839290614515908490615527565b90915550505b505b600c5461452a90826154c6565b6001600160a01b039092165f908152600d602052604090209190915550565b6001600160a01b0381165f908152600e60205260409020548015611457576001600160a01b038083165f908152600e60205260408120819055600254909161459391168484614a7b565b9050806145ee576001600160a01b038084165f818152600e60205260409081902085905560025490519216917f0a2d568d757e153f191d3031fa904fd43c649279614e983e3de9141dfc082d9190611b839086815260200190565b6002546040518381526001600160a01b03918216918516907f0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b790602001611b83565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698400836146ff575f6146897feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b031690565b6001600160a01b0316146146c9576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0385161790555b6120418484614afd565b5f7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d86984008315801561476557507feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d8698401546001600160a01b038481169116145b15614795576001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6120418484614bc0565b6147a882614c64565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156147ec576118fd8282614d0b565b611457614d7d565b5f65ffffffffffff82111561483f576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401611298565b5090565b7feef3dac4538c82c8ace4063ab0acd2d15cdb5883aa1dff7c2673abb3d869840080547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff000000000000000000000000000000000000000000000000000084166001600160a01b0388161717845591041680156133ed576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a150505050565b5f5f6148fa612b86565b90508065ffffffffffff168365ffffffffffff16116149225761491d83826155d2565b611576565b61157665ffffffffffff841662069780614db5565b5f81815260018301602052604081205461497c57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610f17565b505f610f17565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611923576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060815f01805480602002602001604051908101604052809291908181526020018280548015614a2b57602002820191905f5260205f20905b815481526020019060010190808311614a17575b50505050509050919050565b5f825f018281548110614a4c57614a4c615434565b905f5260205f200154905092915050565b5f614a66612e6e565b5468010000000000000000900460ff16919050565b5f61157384856001600160a01b031663a9059cbb8686604051602401614ab69291906001600160a01b03929092168252602082015260400190565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050614dc4565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166143f6575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055614b763390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610f17565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156143f6575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610f17565b806001600160a01b03163b5f03614cb2576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611298565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051614d2791906155f0565b5f60405180830381855af49150503d805f8114614d5f576040519150601f19603f3d011682016040523d82523d5f602084013e614d64565b606091505b5091509150614d74858383614e0d565b95945050505050565b3415611923576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828218828410028218611576565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015614e0357508115614df55780600114614e03565b5f866001600160a01b03163b115b9695505050505050565b606082614e1d5761491d82614e7d565b8151158015614e3457506001600160a01b0384163b155b15614e76576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611298565b5080611576565b805115614e8c57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215614ece575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611576575f5ffd5b6001600160a01b0381168114610f2f575f5ffd5b5f5f5f60608486031215614f23575f5ffd5b8335614f2e81614efd565b92506020840135614f3e81614efd565b929592945050506040919091013590565b5f5f60408385031215614f60575f5ffd5b823591506020830135614f7281614efd565b809150509250929050565b5f60208284031215614f8d575f5ffd5b813561157681614efd565b5f8151808452602084019350602083015f5b82811015614fd15781516001600160a01b0316865260209586019590910190600101614faa565b5093949350505050565b5f8151808452602084019350602083015f5b82811015614fd1578151865260209586019590910190600101614fed565b604081525f61501d6040830185614f98565b8281036020840152614d748185614fdb565b5f5f5f60608486031215615041575f5ffd5b505081359360208301359350604090920135919050565b5f60208284031215615068575f5ffd5b5035919050565b5f5f60408385031215615080575f5ffd5b823561508b81614efd565b946020939093013593505050565b5f5f604083850312156150aa575f5ffd5b82356150b581614efd565b91506020830135614f7281614efd565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215615103575f5ffd5b823561510e81614efd565b9150602083013567ffffffffffffffff811115615129575f5ffd5b8301601f81018513615139575f5ffd5b803567ffffffffffffffff811115615153576151536150c5565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156151bf576151bf6150c5565b6040528181528282016020018710156151d6575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f60208284031215615205575f5ffd5b813565ffffffffffff81168114611576575f5ffd5b5f5f5f5f6080858703121561522d575f5ffd5b843561523881614efd565b966020860135965060408601359560600135945092505050565b5f60208284031215615262575f5ffd5b813560038110611576575f5ffd5b602081525f6115766020830184614fdb565b8051825260208101516152a060208401826001600160a01b03169052565b5060408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301526101008101516101008301526101208101516118fd61012084018215159052565b602080825282518282018190525f918401906040840190835b818110156153415761532a838551615282565b602093909301926101409290920191600101615317565b509095945050505050565b6101408101610f178284615282565b602081525f6115766020830184614f98565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110615422577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b60208101610f1782846153ed565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f826154c1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8082028115828204841417610f1757610f17615461565b81810381811115610f1757610f17615461565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361552057615520615461565b5060010190565b80820180821115610f1757610f17615461565b6040810161554882856153ed565b61157660208301846153ed565b5f60208284031215615565575f5ffd5b5051919050565b5f6020828403121561557c575f5ffd5b815161157681614efd565b65ffffffffffff8181168382160190811115610f1757610f17615461565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b65ffffffffffff8281168282160390811115610f1757610f17615461565b5f82518060208501845e5f92019182525091905056fea264697066735822122098715f3050000ddd8453e72c01b435008f5dc176122cc1574cf771682d76e12b64736f6c634300081c0033",
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

// MaxActiveRounds is a free data retrieval call binding the contract method 0x14cfd856.
//
// Solidity: function maxActiveRounds() view returns(uint256)
func (_GamePool *GamePoolCaller) MaxActiveRounds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "maxActiveRounds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxActiveRounds is a free data retrieval call binding the contract method 0x14cfd856.
//
// Solidity: function maxActiveRounds() view returns(uint256)
func (_GamePool *GamePoolSession) MaxActiveRounds() (*big.Int, error) {
	return _GamePool.Contract.MaxActiveRounds(&_GamePool.CallOpts)
}

// MaxActiveRounds is a free data retrieval call binding the contract method 0x14cfd856.
//
// Solidity: function maxActiveRounds() view returns(uint256)
func (_GamePool *GamePoolCallerSession) MaxActiveRounds() (*big.Int, error) {
	return _GamePool.Contract.MaxActiveRounds(&_GamePool.CallOpts)
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
// Solidity: function pendingReward(address user, address token) view returns(uint256 amount)
func (_GamePool *GamePoolCaller) PendingReward(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "pendingReward", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address token) view returns(uint256 amount)
func (_GamePool *GamePoolSession) PendingReward(user common.Address, token common.Address) (*big.Int, error) {
	return _GamePool.Contract.PendingReward(&_GamePool.CallOpts, user, token)
}

// PendingReward is a free data retrieval call binding the contract method 0x9ced7e76.
//
// Solidity: function pendingReward(address user, address token) view returns(uint256 amount)
func (_GamePool *GamePoolCallerSession) PendingReward(user common.Address, token common.Address) (*big.Int, error) {
	return _GamePool.Contract.PendingReward(&_GamePool.CallOpts, user, token)
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
// Solidity: function userRewards(address account, address token) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_GamePool *GamePoolCaller) UserRewards(opts *bind.CallOpts, account common.Address, token common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	var out []interface{}
	err := _GamePool.contract.Call(opts, &out, "userRewards", account, token)

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
// Solidity: function userRewards(address account, address token) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_GamePool *GamePoolSession) UserRewards(account common.Address, token common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _GamePool.Contract.UserRewards(&_GamePool.CallOpts, account, token)
}

// UserRewards is a free data retrieval call binding the contract method 0xa980356a.
//
// Solidity: function userRewards(address account, address token) view returns(uint256 rewardPerTokenPaid, uint256 rewards)
func (_GamePool *GamePoolCallerSession) UserRewards(account common.Address, token common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _GamePool.Contract.UserRewards(&_GamePool.CallOpts, account, token)
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

// SetMaxActiveRounds is a paid mutator transaction binding the contract method 0x5a6a0153.
//
// Solidity: function setMaxActiveRounds(uint256 newMax) returns()
func (_GamePool *GamePoolTransactor) SetMaxActiveRounds(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _GamePool.contract.Transact(opts, "setMaxActiveRounds", newMax)
}

// SetMaxActiveRounds is a paid mutator transaction binding the contract method 0x5a6a0153.
//
// Solidity: function setMaxActiveRounds(uint256 newMax) returns()
func (_GamePool *GamePoolSession) SetMaxActiveRounds(newMax *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.SetMaxActiveRounds(&_GamePool.TransactOpts, newMax)
}

// SetMaxActiveRounds is a paid mutator transaction binding the contract method 0x5a6a0153.
//
// Solidity: function setMaxActiveRounds(uint256 newMax) returns()
func (_GamePool *GamePoolTransactorSession) SetMaxActiveRounds(newMax *big.Int) (*types.Transaction, error) {
	return _GamePool.Contract.SetMaxActiveRounds(&_GamePool.TransactOpts, newMax)
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

// GamePoolMaxActiveRoundsUpdatedIterator is returned from FilterMaxActiveRoundsUpdated and is used to iterate over the raw logs and unpacked data for MaxActiveRoundsUpdated events raised by the GamePool contract.
type GamePoolMaxActiveRoundsUpdatedIterator struct {
	Event *GamePoolMaxActiveRoundsUpdated // Event containing the contract specifics and raw log

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
func (it *GamePoolMaxActiveRoundsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GamePoolMaxActiveRoundsUpdated)
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
		it.Event = new(GamePoolMaxActiveRoundsUpdated)
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
func (it *GamePoolMaxActiveRoundsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GamePoolMaxActiveRoundsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GamePoolMaxActiveRoundsUpdated represents a MaxActiveRoundsUpdated event raised by the GamePool contract.
type GamePoolMaxActiveRoundsUpdated struct {
	OldMax *big.Int
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxActiveRoundsUpdated is a free log retrieval operation binding the contract event 0x0528d77b9faf6330e42880ad6b8266236ed5d1bd23a917f5f6c21eecc7ed00c6.
//
// Solidity: event MaxActiveRoundsUpdated(uint256 oldMax, uint256 newMax)
func (_GamePool *GamePoolFilterer) FilterMaxActiveRoundsUpdated(opts *bind.FilterOpts) (*GamePoolMaxActiveRoundsUpdatedIterator, error) {

	logs, sub, err := _GamePool.contract.FilterLogs(opts, "MaxActiveRoundsUpdated")
	if err != nil {
		return nil, err
	}
	return &GamePoolMaxActiveRoundsUpdatedIterator{contract: _GamePool.contract, event: "MaxActiveRoundsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxActiveRoundsUpdated is a free log subscription operation binding the contract event 0x0528d77b9faf6330e42880ad6b8266236ed5d1bd23a917f5f6c21eecc7ed00c6.
//
// Solidity: event MaxActiveRoundsUpdated(uint256 oldMax, uint256 newMax)
func (_GamePool *GamePoolFilterer) WatchMaxActiveRoundsUpdated(opts *bind.WatchOpts, sink chan<- *GamePoolMaxActiveRoundsUpdated) (event.Subscription, error) {

	logs, sub, err := _GamePool.contract.WatchLogs(opts, "MaxActiveRoundsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GamePoolMaxActiveRoundsUpdated)
				if err := _GamePool.contract.UnpackLog(event, "MaxActiveRoundsUpdated", log); err != nil {
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

// ParseMaxActiveRoundsUpdated is a log parse operation binding the contract event 0x0528d77b9faf6330e42880ad6b8266236ed5d1bd23a917f5f6c21eecc7ed00c6.
//
// Solidity: event MaxActiveRoundsUpdated(uint256 oldMax, uint256 newMax)
func (_GamePool *GamePoolFilterer) ParseMaxActiveRoundsUpdated(log types.Log) (*GamePoolMaxActiveRoundsUpdated, error) {
	event := new(GamePoolMaxActiveRoundsUpdated)
	if err := _GamePool.contract.UnpackLog(event, "MaxActiveRoundsUpdated", log); err != nil {
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
