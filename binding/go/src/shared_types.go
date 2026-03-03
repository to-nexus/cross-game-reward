// Code generated - DO NOT EDIT.
// This file contains shared type definitions.

package binding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ICrossGameRewardPoolRewardToken struct {
	Token                common.Address
	RewardPerTokenStored *big.Int
	LastBalance          *big.Int
	ReclaimableAmount    *big.Int
	DistributedAmount    *big.Int
	IsRemoved            bool
}


