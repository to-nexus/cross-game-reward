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

// CrossGameRewardRouterMetaData contains all meta data concerning the CrossGameRewardRouter contract.
var CrossGameRewardRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_crossGameReward\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"crossGameReward\",\"outputs\":[{\"internalType\":\"contractICrossGameReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositERC20WithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDepositInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"rewardTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"isNativePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcross\",\"outputs\":[{\"internalType\":\"contractIWCROSS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnNative\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CSRCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRNoDepositFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRNotWCROSSPool\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"f4e24740": "crossGameReward()",
		"411e17a1": "depositERC20(uint256,uint256)",
		"bfda6f8f": "depositERC20WithPermit(uint256,uint256,uint256,uint8,bytes32,bytes32)",
		"608fc37a": "depositNative(uint256)",
		"fac4348b": "getUserDepositInfo(uint256,address)",
		"7873d5a6": "isNativePool(uint256)",
		"a2db4582": "wcross()",
		"d78276c6": "withdrawERC20(uint256)",
		"84276d81": "withdrawNative(uint256)",
	},
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161171b38038061171b83398101604081905261002e916100ea565b6001600160a01b03811661005557604051630f59042560e31b815260040160405180910390fd5b6001600160a01b03811660808190526040805163516da2c160e11b8152905163a2db4582916004808201926020929091908290030181865afa15801561009d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100c191906100ea565b6001600160a01b031660a0525061010c565b6001600160a01b03811681146100e7575f5ffd5b50565b5f602082840312156100fa575f5ffd5b8151610105816100d3565b9392505050565b60805160a0516115c46101575f395f818161012f01528181610323015281816103b7015281816104ac015281816106ef0152610e7a01525f81816101c50152610cd501526115c45ff3fe608060405260043610610093575f3560e01c8063a2db458211610066578063d78276c61161004c578063d78276c614610195578063f4e24740146101b4578063fac4348b146101e7575f5ffd5b8063a2db45821461011e578063bfda6f8f14610176575f5ffd5b8063411e17a114610097578063608fc37a146100b85780637873d5a6146100cb57806384276d81146100ff575b5f5ffd5b3480156100a2575f5ffd5b506100b66100b136600461120a565b610215565b005b6100b66100c636600461122a565b6102dc565b3480156100d6575f5ffd5b506100ea6100e536600461122a565b61049d565b60405190151581526020015b60405180910390f35b34801561010a575f5ffd5b506100b661011936600461122a565b61056c565b348015610129575f5ffd5b506101517f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f6565b348015610181575f5ffd5b506100b6610190366004611241565b610799565b3480156101a0575f5ffd5b506100b66101af36600461122a565b61090a565b3480156101bf575f5ffd5b506101517f000000000000000000000000000000000000000000000000000000000000000081565b3480156101f2575f5ffd5b506102066102013660046112b9565b610b46565b6040516100f6939291906112e7565b5f811161024e576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61025883610ca5565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102a4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102c89190611385565b90506102d684838386610d58565b50505050565b5f3411610315576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61031f82610e6c565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004015f604051808303818588803b158015610387575f5ffd5b505af1158015610399573d5f5f3e3d5ffd5b506103e293505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915083905034610f6e565b6040517f2f4f21e200000000000000000000000000000000000000000000000000000000815233600482015234602482015273ffffffffffffffffffffffffffffffffffffffff821690632f4f21e2906044015f604051808303815f87803b15801561044c575f5ffd5b505af115801561045e573d5f5f3e3d5ffd5b50506040513481528492503391507f5d019a2a9281540cd9eee8188bd4cc82b755df607bc58c2fe35c8e34756f2a5b9060200160405180910390a35050565b5f5f6104a883610ca5565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa15801561052a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061054e9190611385565b73ffffffffffffffffffffffffffffffffffffffff16149392505050565b5f61057682610e6c565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8316906327e235e390602401602060405180830381865afa1580156105e3573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060791906113a7565b90505f8111610642576040517f96148b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f9eca672c00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff831690639eca672c906024015f604051808303815f87803b1580156106a6575f5ffd5b505af11580156106b8573d5f5f3e3d5ffd5b50506040517f205c2878000000000000000000000000000000000000000000000000000000008152336004820152602481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16925063205c287891506044015f604051808303815f87803b158015610747575f5ffd5b505af1158015610759573d5f5f3e3d5ffd5b50506040518381528592503391507f752e8de6c42be507958f1221bc18b53ffbae06d588fdf270baf984151f10ac899060200160405180910390a3505050565b5f85116107d2576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6107dc87610ca5565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610828573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084c9190611385565b6040517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018990526064810188905260ff8716608482015260a4810186905260c4810185905290915073ffffffffffffffffffffffffffffffffffffffff82169063d505accf9060e4015f604051808303815f87803b1580156108de575f5ffd5b505af11580156108f0573d5f5f3e3d5ffd5b505050506109008883838a610d58565b5050505050505050565b5f61091482610ca5565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610960573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109849190611385565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8416906327e235e390602401602060405180830381865afa1580156109f1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a1591906113a7565b90505f8111610a50576040517f96148b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f9eca672c00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff841690639eca672c906024015f604051808303815f87803b158015610ab4575f5ffd5b505af1158015610ac6573d5f5f3e3d5ffd5b50610aec9250505073ffffffffffffffffffffffffffffffffffffffff83163383611088565b6040805173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052859133917f2b720e0f171ae4e7ba1010d6984c73c560d8341aa6c000c18ad796184511dede91015b60405180910390a350505050565b5f6060805f610b5486610ca5565b6040517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919250908216906327e235e390602401602060405180830381865afa158015610bc2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be691906113a7565b6040517f31d7a26200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919550908216906331d7a262906024015f60405180830381865afa158015610c53573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610c9891908101906114c7565b9497909650939450505050565b6040517ea5ae21000000000000000000000000000000000000000000000000000000008152600481018290525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169062a5ae2190602401602060405180830381865afa158015610d2e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d529190611385565b92915050565b610d7a73ffffffffffffffffffffffffffffffffffffffff83163330846110cb565b610d9b73ffffffffffffffffffffffffffffffffffffffff83168483610f6e565b6040517f2f4f21e20000000000000000000000000000000000000000000000000000000081523360048201526024810182905273ffffffffffffffffffffffffffffffffffffffff841690632f4f21e2906044015f604051808303815f87803b158015610e06575f5ffd5b505af1158015610e18573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018590528793503392507f5f369940d0b1718759f89b5a321eea1b2616a8f299ee84ae08963dbdfffd94159101610b38565b5f610e7682610ca5565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ef8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f1c9190611385565b73ffffffffffffffffffffffffffffffffffffffff1614610f69576040517fdfe9f23c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052610ffa8482611111565b6102d65760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f604483015261107e91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611167565b6102d68482611167565b60405173ffffffffffffffffffffffffffffffffffffffff8381166024830152604482018390526110c691859182169063a9059cbb90606401611037565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526102d69186918216906323b872dd90608401611037565b5f5f5f5f60205f8651602088015f8a5af192503d91505f51905082801561115d57508115611142578060011461115d565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180611186576040513d5f823e3d81fd5b50505f513d9150811561119d5780600114156111b7565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156102d6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b5f5f6040838503121561121b575f5ffd5b50508035926020909101359150565b5f6020828403121561123a575f5ffd5b5035919050565b5f5f5f5f5f5f60c08789031215611256575f5ffd5b863595506020870135945060408701359350606087013560ff8116811461127b575f5ffd5b9598949750929560808101359460a0909101359350915050565b73ffffffffffffffffffffffffffffffffffffffff811681146112b6575f5ffd5b50565b5f5f604083850312156112ca575f5ffd5b8235915060208301356112dc81611295565b809150509250929050565b5f60608201858352606060208401528085518083526080850191506020870192505f5b8181101561133e57835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161130a565b50508381036040850152845180825260209182019250908501905f5b8181101561137857825184526020938401939092019160010161135a565b5091979650505050505050565b5f60208284031215611395575f5ffd5b81516113a081611295565b9392505050565b5f602082840312156113b7575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611432576114326113be565b604052919050565b5f67ffffffffffffffff821115611453576114536113be565b5060051b60200190565b5f82601f83011261146c575f5ffd5b815161147f61147a8261143a565b6113eb565b8082825260208201915060208360051b8601019250858311156114a0575f5ffd5b602085015b838110156114bd5780518352602092830192016114a5565b5095945050505050565b5f5f604083850312156114d8575f5ffd5b825167ffffffffffffffff8111156114ee575f5ffd5b8301601f810185136114fe575f5ffd5b805161150c61147a8261143a565b8082825260208201915060208360051b85010192508783111561152d575f5ffd5b6020840193505b8284101561155857835161154781611295565b825260209384019390910190611534565b80955050505050602083015167ffffffffffffffff811115611578575f5ffd5b6115848582860161145d565b915050925092905056fea264697066735822122034f428a153b142c75388017b2abdd3851e1dde3f5b61a3a7ed72dd3e2df1afc864736f6c634300081c0033",
}

// CrossGameRewardRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossGameRewardRouterMetaData.ABI instead.
var CrossGameRewardRouterABI = CrossGameRewardRouterMetaData.ABI

// CrossGameRewardRouterBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossGameRewardRouterBinRuntime = "608060405260043610610093575f3560e01c8063a2db458211610066578063d78276c61161004c578063d78276c614610195578063f4e24740146101b4578063fac4348b146101e7575f5ffd5b8063a2db45821461011e578063bfda6f8f14610176575f5ffd5b8063411e17a114610097578063608fc37a146100b85780637873d5a6146100cb57806384276d81146100ff575b5f5ffd5b3480156100a2575f5ffd5b506100b66100b136600461120a565b610215565b005b6100b66100c636600461122a565b6102dc565b3480156100d6575f5ffd5b506100ea6100e536600461122a565b61049d565b60405190151581526020015b60405180910390f35b34801561010a575f5ffd5b506100b661011936600461122a565b61056c565b348015610129575f5ffd5b506101517f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f6565b348015610181575f5ffd5b506100b6610190366004611241565b610799565b3480156101a0575f5ffd5b506100b66101af36600461122a565b61090a565b3480156101bf575f5ffd5b506101517f000000000000000000000000000000000000000000000000000000000000000081565b3480156101f2575f5ffd5b506102066102013660046112b9565b610b46565b6040516100f6939291906112e7565b5f811161024e576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61025883610ca5565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102a4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102c89190611385565b90506102d684838386610d58565b50505050565b5f3411610315576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61031f82610e6c565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004015f604051808303818588803b158015610387575f5ffd5b505af1158015610399573d5f5f3e3d5ffd5b506103e293505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915083905034610f6e565b6040517f2f4f21e200000000000000000000000000000000000000000000000000000000815233600482015234602482015273ffffffffffffffffffffffffffffffffffffffff821690632f4f21e2906044015f604051808303815f87803b15801561044c575f5ffd5b505af115801561045e573d5f5f3e3d5ffd5b50506040513481528492503391507f5d019a2a9281540cd9eee8188bd4cc82b755df607bc58c2fe35c8e34756f2a5b9060200160405180910390a35050565b5f5f6104a883610ca5565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa15801561052a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061054e9190611385565b73ffffffffffffffffffffffffffffffffffffffff16149392505050565b5f61057682610e6c565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8316906327e235e390602401602060405180830381865afa1580156105e3573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061060791906113a7565b90505f8111610642576040517f96148b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f9eca672c00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff831690639eca672c906024015f604051808303815f87803b1580156106a6575f5ffd5b505af11580156106b8573d5f5f3e3d5ffd5b50506040517f205c2878000000000000000000000000000000000000000000000000000000008152336004820152602481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16925063205c287891506044015f604051808303815f87803b158015610747575f5ffd5b505af1158015610759573d5f5f3e3d5ffd5b50506040518381528592503391507f752e8de6c42be507958f1221bc18b53ffbae06d588fdf270baf984151f10ac899060200160405180910390a3505050565b5f85116107d2576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6107dc87610ca5565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610828573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084c9190611385565b6040517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018990526064810188905260ff8716608482015260a4810186905260c4810185905290915073ffffffffffffffffffffffffffffffffffffffff82169063d505accf9060e4015f604051808303815f87803b1580156108de575f5ffd5b505af11580156108f0573d5f5f3e3d5ffd5b505050506109008883838a610d58565b5050505050505050565b5f61091482610ca5565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610960573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109849190611385565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8416906327e235e390602401602060405180830381865afa1580156109f1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a1591906113a7565b90505f8111610a50576040517f96148b9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f9eca672c00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff841690639eca672c906024015f604051808303815f87803b158015610ab4575f5ffd5b505af1158015610ac6573d5f5f3e3d5ffd5b50610aec9250505073ffffffffffffffffffffffffffffffffffffffff83163383611088565b6040805173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052859133917f2b720e0f171ae4e7ba1010d6984c73c560d8341aa6c000c18ad796184511dede91015b60405180910390a350505050565b5f6060805f610b5486610ca5565b6040517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919250908216906327e235e390602401602060405180830381865afa158015610bc2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be691906113a7565b6040517f31d7a26200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919550908216906331d7a262906024015f60405180830381865afa158015610c53573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610c9891908101906114c7565b9497909650939450505050565b6040517ea5ae21000000000000000000000000000000000000000000000000000000008152600481018290525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169062a5ae2190602401602060405180830381865afa158015610d2e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d529190611385565b92915050565b610d7a73ffffffffffffffffffffffffffffffffffffffff83163330846110cb565b610d9b73ffffffffffffffffffffffffffffffffffffffff83168483610f6e565b6040517f2f4f21e20000000000000000000000000000000000000000000000000000000081523360048201526024810182905273ffffffffffffffffffffffffffffffffffffffff841690632f4f21e2906044015f604051808303815f87803b158015610e06575f5ffd5b505af1158015610e18573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018590528793503392507f5f369940d0b1718759f89b5a321eea1b2616a8f299ee84ae08963dbdfffd94159101610b38565b5f610e7682610ca5565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663c89039c56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ef8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f1c9190611385565b73ffffffffffffffffffffffffffffffffffffffff1614610f69576040517fdfe9f23c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052610ffa8482611111565b6102d65760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f604483015261107e91869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611167565b6102d68482611167565b60405173ffffffffffffffffffffffffffffffffffffffff8381166024830152604482018390526110c691859182169063a9059cbb90606401611037565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526102d69186918216906323b872dd90608401611037565b5f5f5f5f60205f8651602088015f8a5af192503d91505f51905082801561115d57508115611142578060011461115d565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f60205f8451602086015f885af180611186576040513d5f823e3d81fd5b50505f513d9150811561119d5780600114156111b7565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156102d6576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b5f5f6040838503121561121b575f5ffd5b50508035926020909101359150565b5f6020828403121561123a575f5ffd5b5035919050565b5f5f5f5f5f5f60c08789031215611256575f5ffd5b863595506020870135945060408701359350606087013560ff8116811461127b575f5ffd5b9598949750929560808101359460a0909101359350915050565b73ffffffffffffffffffffffffffffffffffffffff811681146112b6575f5ffd5b50565b5f5f604083850312156112ca575f5ffd5b8235915060208301356112dc81611295565b809150509250929050565b5f60608201858352606060208401528085518083526080850191506020870192505f5b8181101561133e57835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161130a565b50508381036040850152845180825260209182019250908501905f5b8181101561137857825184526020938401939092019160010161135a565b5091979650505050505050565b5f60208284031215611395575f5ffd5b81516113a081611295565b9392505050565b5f602082840312156113b7575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611432576114326113be565b604052919050565b5f67ffffffffffffffff821115611453576114536113be565b5060051b60200190565b5f82601f83011261146c575f5ffd5b815161147f61147a8261143a565b6113eb565b8082825260208201915060208360051b8601019250858311156114a0575f5ffd5b602085015b838110156114bd5780518352602092830192016114a5565b5095945050505050565b5f5f604083850312156114d8575f5ffd5b825167ffffffffffffffff8111156114ee575f5ffd5b8301601f810185136114fe575f5ffd5b805161150c61147a8261143a565b8082825260208201915060208360051b85010192508783111561152d575f5ffd5b6020840193505b8284101561155857835161154781611295565b825260209384019390910190611534565b80955050505050602083015167ffffffffffffffff811115611578575f5ffd5b6115848582860161145d565b915050925092905056fea264697066735822122034f428a153b142c75388017b2abdd3851e1dde3f5b61a3a7ed72dd3e2df1afc864736f6c634300081c0033"

// Deprecated: Use CrossGameRewardRouterMetaData.Sigs instead.
// CrossGameRewardRouterFuncSigs maps the 4-byte function signature to its string representation.
var CrossGameRewardRouterFuncSigs = CrossGameRewardRouterMetaData.Sigs

// CrossGameRewardRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossGameRewardRouterMetaData.Bin instead.
var CrossGameRewardRouterBin = CrossGameRewardRouterMetaData.Bin

// DeployCrossGameRewardRouter deploys a new Ethereum contract, binding an instance of CrossGameRewardRouter to it.
func DeployCrossGameRewardRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _crossGameReward common.Address) (common.Address, *types.Transaction, *CrossGameRewardRouter, error) {
	parsed, err := CrossGameRewardRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossGameRewardRouterBin), backend, _crossGameReward)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossGameRewardRouter{CrossGameRewardRouterCaller: CrossGameRewardRouterCaller{contract: contract}, CrossGameRewardRouterTransactor: CrossGameRewardRouterTransactor{contract: contract}, CrossGameRewardRouterFilterer: CrossGameRewardRouterFilterer{contract: contract}}, nil
}

// CrossGameRewardRouter is an auto generated Go binding around an Ethereum contract.
type CrossGameRewardRouter struct {
	CrossGameRewardRouterCaller     // Read-only binding to the contract
	CrossGameRewardRouterTransactor // Write-only binding to the contract
	CrossGameRewardRouterFilterer   // Log filterer for contract events
}

// CrossGameRewardRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossGameRewardRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossGameRewardRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossGameRewardRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossGameRewardRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossGameRewardRouterSession struct {
	Contract     *CrossGameRewardRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrossGameRewardRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossGameRewardRouterCallerSession struct {
	Contract *CrossGameRewardRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CrossGameRewardRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossGameRewardRouterTransactorSession struct {
	Contract     *CrossGameRewardRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CrossGameRewardRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossGameRewardRouterRaw struct {
	Contract *CrossGameRewardRouter // Generic contract binding to access the raw methods on
}

// CrossGameRewardRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossGameRewardRouterCallerRaw struct {
	Contract *CrossGameRewardRouterCaller // Generic read-only contract binding to access the raw methods on
}

// CrossGameRewardRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossGameRewardRouterTransactorRaw struct {
	Contract *CrossGameRewardRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossGameRewardRouter creates a new instance of CrossGameRewardRouter, bound to a specific deployed contract.
func NewCrossGameRewardRouter(address common.Address, backend bind.ContractBackend) (*CrossGameRewardRouter, error) {
	contract, err := bindCrossGameRewardRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouter{CrossGameRewardRouterCaller: CrossGameRewardRouterCaller{contract: contract}, CrossGameRewardRouterTransactor: CrossGameRewardRouterTransactor{contract: contract}, CrossGameRewardRouterFilterer: CrossGameRewardRouterFilterer{contract: contract}}, nil
}

// NewCrossGameRewardRouterCaller creates a new read-only instance of CrossGameRewardRouter, bound to a specific deployed contract.
func NewCrossGameRewardRouterCaller(address common.Address, caller bind.ContractCaller) (*CrossGameRewardRouterCaller, error) {
	contract, err := bindCrossGameRewardRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterCaller{contract: contract}, nil
}

// NewCrossGameRewardRouterTransactor creates a new write-only instance of CrossGameRewardRouter, bound to a specific deployed contract.
func NewCrossGameRewardRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossGameRewardRouterTransactor, error) {
	contract, err := bindCrossGameRewardRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterTransactor{contract: contract}, nil
}

// NewCrossGameRewardRouterFilterer creates a new log filterer instance of CrossGameRewardRouter, bound to a specific deployed contract.
func NewCrossGameRewardRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossGameRewardRouterFilterer, error) {
	contract, err := bindCrossGameRewardRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterFilterer{contract: contract}, nil
}

// bindCrossGameRewardRouter binds a generic wrapper to an already deployed contract.
func bindCrossGameRewardRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossGameRewardRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossGameRewardRouter *CrossGameRewardRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossGameRewardRouter.Contract.CrossGameRewardRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossGameRewardRouter *CrossGameRewardRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.CrossGameRewardRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossGameRewardRouter *CrossGameRewardRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.CrossGameRewardRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossGameRewardRouter *CrossGameRewardRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossGameRewardRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.contract.Transact(opts, method, params...)
}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_CrossGameRewardRouter *CrossGameRewardRouterCaller) CrossGameReward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardRouter.contract.Call(opts, &out, "crossGameReward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) CrossGameReward() (common.Address, error) {
	return _CrossGameRewardRouter.Contract.CrossGameReward(&_CrossGameRewardRouter.CallOpts)
}

// CrossGameReward is a free data retrieval call binding the contract method 0xf4e24740.
//
// Solidity: function crossGameReward() view returns(address)
func (_CrossGameRewardRouter *CrossGameRewardRouterCallerSession) CrossGameReward() (common.Address, error) {
	return _CrossGameRewardRouter.Contract.CrossGameReward(&_CrossGameRewardRouter.CallOpts)
}

// GetUserDepositInfo is a free data retrieval call binding the contract method 0xfac4348b.
//
// Solidity: function getUserDepositInfo(uint256 poolId, address user) view returns(uint256 depositedAmount, address[] rewardTokens, uint256[] pendingRewards)
func (_CrossGameRewardRouter *CrossGameRewardRouterCaller) GetUserDepositInfo(opts *bind.CallOpts, poolId *big.Int, user common.Address) (struct {
	DepositedAmount *big.Int
	RewardTokens    []common.Address
	PendingRewards  []*big.Int
}, error) {
	var out []interface{}
	err := _CrossGameRewardRouter.contract.Call(opts, &out, "getUserDepositInfo", poolId, user)

	outstruct := new(struct {
		DepositedAmount *big.Int
		RewardTokens    []common.Address
		PendingRewards  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardTokens = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.PendingRewards = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetUserDepositInfo is a free data retrieval call binding the contract method 0xfac4348b.
//
// Solidity: function getUserDepositInfo(uint256 poolId, address user) view returns(uint256 depositedAmount, address[] rewardTokens, uint256[] pendingRewards)
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) GetUserDepositInfo(poolId *big.Int, user common.Address) (struct {
	DepositedAmount *big.Int
	RewardTokens    []common.Address
	PendingRewards  []*big.Int
}, error) {
	return _CrossGameRewardRouter.Contract.GetUserDepositInfo(&_CrossGameRewardRouter.CallOpts, poolId, user)
}

// GetUserDepositInfo is a free data retrieval call binding the contract method 0xfac4348b.
//
// Solidity: function getUserDepositInfo(uint256 poolId, address user) view returns(uint256 depositedAmount, address[] rewardTokens, uint256[] pendingRewards)
func (_CrossGameRewardRouter *CrossGameRewardRouterCallerSession) GetUserDepositInfo(poolId *big.Int, user common.Address) (struct {
	DepositedAmount *big.Int
	RewardTokens    []common.Address
	PendingRewards  []*big.Int
}, error) {
	return _CrossGameRewardRouter.Contract.GetUserDepositInfo(&_CrossGameRewardRouter.CallOpts, poolId, user)
}

// IsNativePool is a free data retrieval call binding the contract method 0x7873d5a6.
//
// Solidity: function isNativePool(uint256 poolId) view returns(bool)
func (_CrossGameRewardRouter *CrossGameRewardRouterCaller) IsNativePool(opts *bind.CallOpts, poolId *big.Int) (bool, error) {
	var out []interface{}
	err := _CrossGameRewardRouter.contract.Call(opts, &out, "isNativePool", poolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNativePool is a free data retrieval call binding the contract method 0x7873d5a6.
//
// Solidity: function isNativePool(uint256 poolId) view returns(bool)
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) IsNativePool(poolId *big.Int) (bool, error) {
	return _CrossGameRewardRouter.Contract.IsNativePool(&_CrossGameRewardRouter.CallOpts, poolId)
}

// IsNativePool is a free data retrieval call binding the contract method 0x7873d5a6.
//
// Solidity: function isNativePool(uint256 poolId) view returns(bool)
func (_CrossGameRewardRouter *CrossGameRewardRouterCallerSession) IsNativePool(poolId *big.Int) (bool, error) {
	return _CrossGameRewardRouter.Contract.IsNativePool(&_CrossGameRewardRouter.CallOpts, poolId)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossGameRewardRouter *CrossGameRewardRouterCaller) Wcross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossGameRewardRouter.contract.Call(opts, &out, "wcross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) Wcross() (common.Address, error) {
	return _CrossGameRewardRouter.Contract.Wcross(&_CrossGameRewardRouter.CallOpts)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossGameRewardRouter *CrossGameRewardRouterCallerSession) Wcross() (common.Address, error) {
	return _CrossGameRewardRouter.Contract.Wcross(&_CrossGameRewardRouter.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x411e17a1.
//
// Solidity: function depositERC20(uint256 poolId, uint256 amount) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactor) DepositERC20(opts *bind.TransactOpts, poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.contract.Transact(opts, "depositERC20", poolId, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x411e17a1.
//
// Solidity: function depositERC20(uint256 poolId, uint256 amount) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) DepositERC20(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.DepositERC20(&_CrossGameRewardRouter.TransactOpts, poolId, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x411e17a1.
//
// Solidity: function depositERC20(uint256 poolId, uint256 amount) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorSession) DepositERC20(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.DepositERC20(&_CrossGameRewardRouter.TransactOpts, poolId, amount)
}

// DepositERC20WithPermit is a paid mutator transaction binding the contract method 0xbfda6f8f.
//
// Solidity: function depositERC20WithPermit(uint256 poolId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactor) DepositERC20WithPermit(opts *bind.TransactOpts, poolId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossGameRewardRouter.contract.Transact(opts, "depositERC20WithPermit", poolId, amount, deadline, v, r, s)
}

// DepositERC20WithPermit is a paid mutator transaction binding the contract method 0xbfda6f8f.
//
// Solidity: function depositERC20WithPermit(uint256 poolId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) DepositERC20WithPermit(poolId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.DepositERC20WithPermit(&_CrossGameRewardRouter.TransactOpts, poolId, amount, deadline, v, r, s)
}

// DepositERC20WithPermit is a paid mutator transaction binding the contract method 0xbfda6f8f.
//
// Solidity: function depositERC20WithPermit(uint256 poolId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorSession) DepositERC20WithPermit(poolId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.DepositERC20WithPermit(&_CrossGameRewardRouter.TransactOpts, poolId, amount, deadline, v, r, s)
}

// DepositNative is a paid mutator transaction binding the contract method 0x608fc37a.
//
// Solidity: function depositNative(uint256 poolId) payable returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactor) DepositNative(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.contract.Transact(opts, "depositNative", poolId)
}

// DepositNative is a paid mutator transaction binding the contract method 0x608fc37a.
//
// Solidity: function depositNative(uint256 poolId) payable returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) DepositNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.DepositNative(&_CrossGameRewardRouter.TransactOpts, poolId)
}

// DepositNative is a paid mutator transaction binding the contract method 0x608fc37a.
//
// Solidity: function depositNative(uint256 poolId) payable returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorSession) DepositNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.DepositNative(&_CrossGameRewardRouter.TransactOpts, poolId)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xd78276c6.
//
// Solidity: function withdrawERC20(uint256 poolId) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactor) WithdrawERC20(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.contract.Transact(opts, "withdrawERC20", poolId)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xd78276c6.
//
// Solidity: function withdrawERC20(uint256 poolId) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) WithdrawERC20(poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.WithdrawERC20(&_CrossGameRewardRouter.TransactOpts, poolId)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xd78276c6.
//
// Solidity: function withdrawERC20(uint256 poolId) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorSession) WithdrawERC20(poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.WithdrawERC20(&_CrossGameRewardRouter.TransactOpts, poolId)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x84276d81.
//
// Solidity: function withdrawNative(uint256 poolId) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactor) WithdrawNative(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.contract.Transact(opts, "withdrawNative", poolId)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x84276d81.
//
// Solidity: function withdrawNative(uint256 poolId) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterSession) WithdrawNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.WithdrawNative(&_CrossGameRewardRouter.TransactOpts, poolId)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x84276d81.
//
// Solidity: function withdrawNative(uint256 poolId) returns()
func (_CrossGameRewardRouter *CrossGameRewardRouterTransactorSession) WithdrawNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossGameRewardRouter.Contract.WithdrawNative(&_CrossGameRewardRouter.TransactOpts, poolId)
}

// CrossGameRewardRouterDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterDepositedERC20Iterator struct {
	Event *CrossGameRewardRouterDepositedERC20 // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRouterDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRouterDepositedERC20)
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
		it.Event = new(CrossGameRewardRouterDepositedERC20)
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
func (it *CrossGameRewardRouterDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRouterDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRouterDepositedERC20 represents a DepositedERC20 event raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterDepositedERC20 struct {
	User   common.Address
	PoolId *big.Int
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC20 is a free log retrieval operation binding the contract event 0x5f369940d0b1718759f89b5a321eea1b2616a8f299ee84ae08963dbdfffd9415.
//
// Solidity: event DepositedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) FilterDepositedERC20(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossGameRewardRouterDepositedERC20Iterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.FilterLogs(opts, "DepositedERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterDepositedERC20Iterator{contract: _CrossGameRewardRouter.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0x5f369940d0b1718759f89b5a321eea1b2616a8f299ee84ae08963dbdfffd9415.
//
// Solidity: event DepositedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRouterDepositedERC20, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.WatchLogs(opts, "DepositedERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRouterDepositedERC20)
				if err := _CrossGameRewardRouter.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
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

// ParseDepositedERC20 is a log parse operation binding the contract event 0x5f369940d0b1718759f89b5a321eea1b2616a8f299ee84ae08963dbdfffd9415.
//
// Solidity: event DepositedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) ParseDepositedERC20(log types.Log) (*CrossGameRewardRouterDepositedERC20, error) {
	event := new(CrossGameRewardRouterDepositedERC20)
	if err := _CrossGameRewardRouter.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRouterDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterDepositedNativeIterator struct {
	Event *CrossGameRewardRouterDepositedNative // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRouterDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRouterDepositedNative)
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
		it.Event = new(CrossGameRewardRouterDepositedNative)
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
func (it *CrossGameRewardRouterDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRouterDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRouterDepositedNative represents a DepositedNative event raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterDepositedNative struct {
	User   common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0x5d019a2a9281540cd9eee8188bd4cc82b755df607bc58c2fe35c8e34756f2a5b.
//
// Solidity: event DepositedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) FilterDepositedNative(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossGameRewardRouterDepositedNativeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.FilterLogs(opts, "DepositedNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterDepositedNativeIterator{contract: _CrossGameRewardRouter.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x5d019a2a9281540cd9eee8188bd4cc82b755df607bc58c2fe35c8e34756f2a5b.
//
// Solidity: event DepositedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRouterDepositedNative, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.WatchLogs(opts, "DepositedNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRouterDepositedNative)
				if err := _CrossGameRewardRouter.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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

// ParseDepositedNative is a log parse operation binding the contract event 0x5d019a2a9281540cd9eee8188bd4cc82b755df607bc58c2fe35c8e34756f2a5b.
//
// Solidity: event DepositedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) ParseDepositedNative(log types.Log) (*CrossGameRewardRouterDepositedNative, error) {
	event := new(CrossGameRewardRouterDepositedNative)
	if err := _CrossGameRewardRouter.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRouterWithdrawnERC20Iterator is returned from FilterWithdrawnERC20 and is used to iterate over the raw logs and unpacked data for WithdrawnERC20 events raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterWithdrawnERC20Iterator struct {
	Event *CrossGameRewardRouterWithdrawnERC20 // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRouterWithdrawnERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRouterWithdrawnERC20)
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
		it.Event = new(CrossGameRewardRouterWithdrawnERC20)
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
func (it *CrossGameRewardRouterWithdrawnERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRouterWithdrawnERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRouterWithdrawnERC20 represents a WithdrawnERC20 event raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterWithdrawnERC20 struct {
	User   common.Address
	PoolId *big.Int
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC20 is a free log retrieval operation binding the contract event 0x2b720e0f171ae4e7ba1010d6984c73c560d8341aa6c000c18ad796184511dede.
//
// Solidity: event WithdrawnERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) FilterWithdrawnERC20(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossGameRewardRouterWithdrawnERC20Iterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.FilterLogs(opts, "WithdrawnERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterWithdrawnERC20Iterator{contract: _CrossGameRewardRouter.contract, event: "WithdrawnERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC20 is a free log subscription operation binding the contract event 0x2b720e0f171ae4e7ba1010d6984c73c560d8341aa6c000c18ad796184511dede.
//
// Solidity: event WithdrawnERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) WatchWithdrawnERC20(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRouterWithdrawnERC20, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.WatchLogs(opts, "WithdrawnERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRouterWithdrawnERC20)
				if err := _CrossGameRewardRouter.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
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

// ParseWithdrawnERC20 is a log parse operation binding the contract event 0x2b720e0f171ae4e7ba1010d6984c73c560d8341aa6c000c18ad796184511dede.
//
// Solidity: event WithdrawnERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) ParseWithdrawnERC20(log types.Log) (*CrossGameRewardRouterWithdrawnERC20, error) {
	event := new(CrossGameRewardRouterWithdrawnERC20)
	if err := _CrossGameRewardRouter.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossGameRewardRouterWithdrawnNativeIterator is returned from FilterWithdrawnNative and is used to iterate over the raw logs and unpacked data for WithdrawnNative events raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterWithdrawnNativeIterator struct {
	Event *CrossGameRewardRouterWithdrawnNative // Event containing the contract specifics and raw log

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
func (it *CrossGameRewardRouterWithdrawnNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossGameRewardRouterWithdrawnNative)
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
		it.Event = new(CrossGameRewardRouterWithdrawnNative)
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
func (it *CrossGameRewardRouterWithdrawnNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossGameRewardRouterWithdrawnNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossGameRewardRouterWithdrawnNative represents a WithdrawnNative event raised by the CrossGameRewardRouter contract.
type CrossGameRewardRouterWithdrawnNative struct {
	User   common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnNative is a free log retrieval operation binding the contract event 0x752e8de6c42be507958f1221bc18b53ffbae06d588fdf270baf984151f10ac89.
//
// Solidity: event WithdrawnNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) FilterWithdrawnNative(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossGameRewardRouterWithdrawnNativeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.FilterLogs(opts, "WithdrawnNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossGameRewardRouterWithdrawnNativeIterator{contract: _CrossGameRewardRouter.contract, event: "WithdrawnNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNative is a free log subscription operation binding the contract event 0x752e8de6c42be507958f1221bc18b53ffbae06d588fdf270baf984151f10ac89.
//
// Solidity: event WithdrawnNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) WatchWithdrawnNative(opts *bind.WatchOpts, sink chan<- *CrossGameRewardRouterWithdrawnNative, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossGameRewardRouter.contract.WatchLogs(opts, "WithdrawnNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossGameRewardRouterWithdrawnNative)
				if err := _CrossGameRewardRouter.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
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

// ParseWithdrawnNative is a log parse operation binding the contract event 0x752e8de6c42be507958f1221bc18b53ffbae06d588fdf270baf984151f10ac89.
//
// Solidity: event WithdrawnNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossGameRewardRouter *CrossGameRewardRouterFilterer) ParseWithdrawnNative(log types.Log) (*CrossGameRewardRouterWithdrawnNative, error) {
	event := new(CrossGameRewardRouterWithdrawnNative)
	if err := _CrossGameRewardRouter.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
