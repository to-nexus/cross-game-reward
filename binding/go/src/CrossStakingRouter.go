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

// CrossStakingRouterMetaData contains all meta data concerning the CrossStakingRouter contract.
var CrossStakingRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_crossStaking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"crossStaking\",\"outputs\":[{\"internalType\":\"contractICrossStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStakingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"rewardTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"isNativePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"stakeERC20WithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"stakeNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"unstakeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"unstakeNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wcross\",\"outputs\":[{\"internalType\":\"contractIWCROSS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakedNative\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CSRCanNotZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRInvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRNoStakeFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRNotWCROSSPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CSRTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"7dd96ac4": "crossStaking()",
		"128644b0": "getUserStakingInfo(uint256,address)",
		"7873d5a6": "isNativePool(uint256)",
		"37d9e9cc": "stakeERC20(uint256,uint256)",
		"08466d63": "stakeERC20WithPermit(uint256,uint256,uint256,uint8,bytes32,bytes32)",
		"e63c6bf0": "stakeNative(uint256)",
		"7dfae334": "unstakeERC20(uint256)",
		"ff8ab5f1": "unstakeNative(uint256)",
		"a2db4582": "wcross()",
	},
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161171f38038061171f83398101604081905261002e916100ea565b6001600160a01b03811661005557604051630f59042560e31b815260040160405180910390fd5b6001600160a01b03811660808190526040805163516da2c160e11b8152905163a2db4582916004808201926020929091908290030181865afa15801561009d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100c191906100ea565b6001600160a01b031660a0525061010c565b6001600160a01b03811681146100e7575f5ffd5b50565b5f602082840312156100fa575f5ffd5b8151610105816100d3565b9392505050565b60805160a0516115c86101575f395f81816101c5015281816105bf015281816109020152818161099601528181610bff0152610f0401525f818161014e0152610cd901526115c85ff3fe608060405260043610610093575f3560e01c80637dd96ac411610066578063a2db45821161004c578063a2db4582146101b4578063e63c6bf0146101e7578063ff8ab5f1146101fa575f5ffd5b80637dd96ac41461013d5780637dfae33414610195575f5ffd5b806308466d6314610097578063128644b0146100b857806337d9e9cc146100ef5780637873d5a61461010e575b5f5ffd5b3480156100a2575f5ffd5b506100b66100b136600461120e565b610219565b005b3480156100c3575f5ffd5b506100d76100d2366004611286565b61038a565b6040516100e6939291906112b4565b60405180910390f35b3480156100fa575f5ffd5b506100b6610109366004611352565b6104e9565b348015610119575f5ffd5b5061012d610128366004611372565b6105b0565b60405190151581526020016100e6565b348015610148575f5ffd5b506101707f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e6565b3480156101a0575f5ffd5b506100b66101af366004611372565b61067f565b3480156101bf575f5ffd5b506101707f000000000000000000000000000000000000000000000000000000000000000081565b6100b66101f5366004611372565b6108bb565b348015610205575f5ffd5b506100b6610214366004611372565b610a7c565b5f8511610252576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61025c87610ca9565b90505f8173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102cc9190611389565b6040517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018990526064810188905260ff8716608482015260a4810186905260c4810185905290915073ffffffffffffffffffffffffffffffffffffffff82169063d505accf9060e4015f604051808303815f87803b15801561035e575f5ffd5b505af1158015610370573d5f5f3e3d5ffd5b505050506103808883838a610d5c565b5050505050505050565b5f6060805f61039886610ca9565b6040517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919250908216906327e235e390602401602060405180830381865afa158015610406573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061042a91906113ab565b6040517f31d7a26200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919550908216906331d7a262906024015f60405180830381865afa158015610497573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526104dc91908101906114cb565b9497909650939450505050565b5f8111610522576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61052c83610ca9565b90505f8173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610578573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059c9190611389565b90506105aa84838386610d5c565b50505050565b5f5f6105bb83610ca9565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561063d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106619190611389565b73ffffffffffffffffffffffffffffffffffffffff16149392505050565b5f61068982610ca9565b90505f8173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f99190611389565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8416906327e235e390602401602060405180830381865afa158015610766573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078a91906113ab565b90505f81116107c5576040517f76ca750d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f2cbb261900000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff841690632cbb2619906024015f604051808303815f87803b158015610829575f5ffd5b505af115801561083b573d5f5f3e3d5ffd5b506108619250505073ffffffffffffffffffffffffffffffffffffffff83163383610e70565b6040805173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052859133917f5571be890ca705d540f4e7ec30a60542bd07eedbc800749ac6c5784d0fa7a56691015b60405180910390a350505050565b5f34116108f4576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6108fe82610ef6565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004015f604051808303818588803b158015610966575f5ffd5b505af1158015610978573d5f5f3e3d5ffd5b506109c193505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915083905034610ff8565b6040517f2ee4090800000000000000000000000000000000000000000000000000000000815233600482015234602482015273ffffffffffffffffffffffffffffffffffffffff821690632ee40908906044015f604051808303815f87803b158015610a2b575f5ffd5b505af1158015610a3d573d5f5f3e3d5ffd5b50506040513481528492503391507f94a750b81c0cca844f1d86628787e6f8444b9f979cd708c31ebbbe62984dc2ac9060200160405180910390a35050565b5f610a8682610ef6565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8316906327e235e390602401602060405180830381865afa158015610af3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b1791906113ab565b90505f8111610b52576040517f76ca750d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f2cbb261900000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff831690632cbb2619906024015f604051808303815f87803b158015610bb6575f5ffd5b505af1158015610bc8573d5f5f3e3d5ffd5b50506040517f205c2878000000000000000000000000000000000000000000000000000000008152336004820152602481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16925063205c287891506044015f604051808303815f87803b158015610c57575f5ffd5b505af1158015610c69573d5f5f3e3d5ffd5b50506040518381528592503391507fc3f10c16db5e7a7551ad91c8ed696eefc6a5b65ca66feaeead4a4cf575c1344b9060200160405180910390a3505050565b6040517ea5ae21000000000000000000000000000000000000000000000000000000008152600481018290525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169062a5ae2190602401602060405180830381865afa158015610d32573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d569190611389565b92915050565b610d7e73ffffffffffffffffffffffffffffffffffffffff83163330846110cf565b610d9f73ffffffffffffffffffffffffffffffffffffffff83168483610ff8565b6040517f2ee409080000000000000000000000000000000000000000000000000000000081523360048201526024810182905273ffffffffffffffffffffffffffffffffffffffff841690632ee40908906044015f604051808303815f87803b158015610e0a575f5ffd5b505af1158015610e1c573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018590528793503392507f61857636d15275e09c9adf9b762a4dbb040d0c12f7e000ad7f71bd8639de94ad91016108ad565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052610ef191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611115565b505050565b5f610f0082610ca9565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f82573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fa69190611389565b73ffffffffffffffffffffffffffffffffffffffff1614610ff3576040517fdfe9f23c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b30000000000000000000000000000000000000000000000000000000017905261108484826111b8565b6105aa5760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f60448301526110c591869182169063095ea7b390606401610eaa565b6105aa8482611115565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526105aa9186918216906323b872dd90608401610eaa565b5f5f60205f8451602086015f885af180611134576040513d5f823e3d81fd5b50505f513d9150811561114b578060011415611165565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156105aa576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015611204575081156111e95780600114611204565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f5f5f5f5f60c08789031215611223575f5ffd5b863595506020870135945060408701359350606087013560ff81168114611248575f5ffd5b9598949750929560808101359460a0909101359350915050565b73ffffffffffffffffffffffffffffffffffffffff81168114611283575f5ffd5b50565b5f5f60408385031215611297575f5ffd5b8235915060208301356112a981611262565b809150509250929050565b5f60608201858352606060208401528085518083526080850191506020870192505f5b8181101561130b57835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016112d7565b50508381036040850152845180825260209182019250908501905f5b81811015611345578251845260209384019390920191600101611327565b5091979650505050505050565b5f5f60408385031215611363575f5ffd5b50508035926020909101359150565b5f60208284031215611382575f5ffd5b5035919050565b5f60208284031215611399575f5ffd5b81516113a481611262565b9392505050565b5f602082840312156113bb575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611436576114366113c2565b604052919050565b5f67ffffffffffffffff821115611457576114576113c2565b5060051b60200190565b5f82601f830112611470575f5ffd5b815161148361147e8261143e565b6113ef565b8082825260208201915060208360051b8601019250858311156114a4575f5ffd5b602085015b838110156114c15780518352602092830192016114a9565b5095945050505050565b5f5f604083850312156114dc575f5ffd5b825167ffffffffffffffff8111156114f2575f5ffd5b8301601f81018513611502575f5ffd5b805161151061147e8261143e565b8082825260208201915060208360051b850101925087831115611531575f5ffd5b6020840193505b8284101561155c57835161154b81611262565b825260209384019390910190611538565b80955050505050602083015167ffffffffffffffff81111561157c575f5ffd5b61158885828601611461565b915050925092905056fea26469706673582212206c77043553f41b39d8ce8651353c0ac904a633f909fe89918615dbd3823fbc2564736f6c634300081c0033",
}

// CrossStakingRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossStakingRouterMetaData.ABI instead.
var CrossStakingRouterABI = CrossStakingRouterMetaData.ABI

// CrossStakingRouterBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CrossStakingRouterBinRuntime = "608060405260043610610093575f3560e01c80637dd96ac411610066578063a2db45821161004c578063a2db4582146101b4578063e63c6bf0146101e7578063ff8ab5f1146101fa575f5ffd5b80637dd96ac41461013d5780637dfae33414610195575f5ffd5b806308466d6314610097578063128644b0146100b857806337d9e9cc146100ef5780637873d5a61461010e575b5f5ffd5b3480156100a2575f5ffd5b506100b66100b136600461120e565b610219565b005b3480156100c3575f5ffd5b506100d76100d2366004611286565b61038a565b6040516100e6939291906112b4565b60405180910390f35b3480156100fa575f5ffd5b506100b6610109366004611352565b6104e9565b348015610119575f5ffd5b5061012d610128366004611372565b6105b0565b60405190151581526020016100e6565b348015610148575f5ffd5b506101707f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e6565b3480156101a0575f5ffd5b506100b66101af366004611372565b61067f565b3480156101bf575f5ffd5b506101707f000000000000000000000000000000000000000000000000000000000000000081565b6100b66101f5366004611372565b6108bb565b348015610205575f5ffd5b506100b6610214366004611372565b610a7c565b5f8511610252576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61025c87610ca9565b90505f8173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102a8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102cc9190611389565b6040517fd505accf000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018990526064810188905260ff8716608482015260a4810186905260c4810185905290915073ffffffffffffffffffffffffffffffffffffffff82169063d505accf9060e4015f604051808303815f87803b15801561035e575f5ffd5b505af1158015610370573d5f5f3e3d5ffd5b505050506103808883838a610d5c565b5050505050505050565b5f6060805f61039886610ca9565b6040517f27e235e300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919250908216906327e235e390602401602060405180830381865afa158015610406573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061042a91906113ab565b6040517f31d7a26200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152919550908216906331d7a262906024015f60405180830381865afa158015610497573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526104dc91908101906114cb565b9497909650939450505050565b5f8111610522576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61052c83610ca9565b90505f8173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610578573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059c9190611389565b90506105aa84838386610d5c565b50505050565b5f5f6105bb83610ca9565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561063d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106619190611389565b73ffffffffffffffffffffffffffffffffffffffff16149392505050565b5f61068982610ca9565b90505f8173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f99190611389565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8416906327e235e390602401602060405180830381865afa158015610766573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078a91906113ab565b90505f81116107c5576040517f76ca750d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f2cbb261900000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff841690632cbb2619906024015f604051808303815f87803b158015610829575f5ffd5b505af115801561083b573d5f5f3e3d5ffd5b506108619250505073ffffffffffffffffffffffffffffffffffffffff83163383610e70565b6040805173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052859133917f5571be890ca705d540f4e7ec30a60542bd07eedbc800749ac6c5784d0fa7a56691015b60405180910390a350505050565b5f34116108f4576040517fed54128200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6108fe82610ef6565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004015f604051808303818588803b158015610966575f5ffd5b505af1158015610978573d5f5f3e3d5ffd5b506109c193505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016915083905034610ff8565b6040517f2ee4090800000000000000000000000000000000000000000000000000000000815233600482015234602482015273ffffffffffffffffffffffffffffffffffffffff821690632ee40908906044015f604051808303815f87803b158015610a2b575f5ffd5b505af1158015610a3d573d5f5f3e3d5ffd5b50506040513481528492503391507f94a750b81c0cca844f1d86628787e6f8444b9f979cd708c31ebbbe62984dc2ac9060200160405180910390a35050565b5f610a8682610ef6565b6040517f27e235e30000000000000000000000000000000000000000000000000000000081523360048201529091505f9073ffffffffffffffffffffffffffffffffffffffff8316906327e235e390602401602060405180830381865afa158015610af3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b1791906113ab565b90505f8111610b52576040517f76ca750d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f2cbb261900000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff831690632cbb2619906024015f604051808303815f87803b158015610bb6575f5ffd5b505af1158015610bc8573d5f5f3e3d5ffd5b50506040517f205c2878000000000000000000000000000000000000000000000000000000008152336004820152602481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16925063205c287891506044015f604051808303815f87803b158015610c57575f5ffd5b505af1158015610c69573d5f5f3e3d5ffd5b50506040518381528592503391507fc3f10c16db5e7a7551ad91c8ed696eefc6a5b65ca66feaeead4a4cf575c1344b9060200160405180910390a3505050565b6040517ea5ae21000000000000000000000000000000000000000000000000000000008152600481018290525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169062a5ae2190602401602060405180830381865afa158015610d32573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d569190611389565b92915050565b610d7e73ffffffffffffffffffffffffffffffffffffffff83163330846110cf565b610d9f73ffffffffffffffffffffffffffffffffffffffff83168483610ff8565b6040517f2ee409080000000000000000000000000000000000000000000000000000000081523360048201526024810182905273ffffffffffffffffffffffffffffffffffffffff841690632ee40908906044015f604051808303815f87803b158015610e0a575f5ffd5b505af1158015610e1c573d5f5f3e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018590528793503392507f61857636d15275e09c9adf9b762a4dbb040d0c12f7e000ad7f71bd8639de94ad91016108ad565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052610ef191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611115565b505050565b5f610f0082610ca9565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166372f702f36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f82573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fa69190611389565b73ffffffffffffffffffffffffffffffffffffffff1614610ff3576040517fdfe9f23c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b30000000000000000000000000000000000000000000000000000000017905261108484826111b8565b6105aa5760405173ffffffffffffffffffffffffffffffffffffffff84811660248301525f60448301526110c591869182169063095ea7b390606401610eaa565b6105aa8482611115565b60405173ffffffffffffffffffffffffffffffffffffffff84811660248301528381166044830152606482018390526105aa9186918216906323b872dd90608401610eaa565b5f5f60205f8451602086015f885af180611134576040513d5f823e3d81fd5b50505f513d9150811561114b578060011415611165565b73ffffffffffffffffffffffffffffffffffffffff84163b155b156105aa576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240160405180910390fd5b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015611204575081156111e95780600114611204565b5f8673ffffffffffffffffffffffffffffffffffffffff163b115b9695505050505050565b5f5f5f5f5f5f60c08789031215611223575f5ffd5b863595506020870135945060408701359350606087013560ff81168114611248575f5ffd5b9598949750929560808101359460a0909101359350915050565b73ffffffffffffffffffffffffffffffffffffffff81168114611283575f5ffd5b50565b5f5f60408385031215611297575f5ffd5b8235915060208301356112a981611262565b809150509250929050565b5f60608201858352606060208401528085518083526080850191506020870192505f5b8181101561130b57835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016112d7565b50508381036040850152845180825260209182019250908501905f5b81811015611345578251845260209384019390920191600101611327565b5091979650505050505050565b5f5f60408385031215611363575f5ffd5b50508035926020909101359150565b5f60208284031215611382575f5ffd5b5035919050565b5f60208284031215611399575f5ffd5b81516113a481611262565b9392505050565b5f602082840312156113bb575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611436576114366113c2565b604052919050565b5f67ffffffffffffffff821115611457576114576113c2565b5060051b60200190565b5f82601f830112611470575f5ffd5b815161148361147e8261143e565b6113ef565b8082825260208201915060208360051b8601019250858311156114a4575f5ffd5b602085015b838110156114c15780518352602092830192016114a9565b5095945050505050565b5f5f604083850312156114dc575f5ffd5b825167ffffffffffffffff8111156114f2575f5ffd5b8301601f81018513611502575f5ffd5b805161151061147e8261143e565b8082825260208201915060208360051b850101925087831115611531575f5ffd5b6020840193505b8284101561155c57835161154b81611262565b825260209384019390910190611538565b80955050505050602083015167ffffffffffffffff81111561157c575f5ffd5b61158885828601611461565b915050925092905056fea26469706673582212206c77043553f41b39d8ce8651353c0ac904a633f909fe89918615dbd3823fbc2564736f6c634300081c0033"

// Deprecated: Use CrossStakingRouterMetaData.Sigs instead.
// CrossStakingRouterFuncSigs maps the 4-byte function signature to its string representation.
var CrossStakingRouterFuncSigs = CrossStakingRouterMetaData.Sigs

// CrossStakingRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossStakingRouterMetaData.Bin instead.
var CrossStakingRouterBin = CrossStakingRouterMetaData.Bin

// DeployCrossStakingRouter deploys a new Ethereum contract, binding an instance of CrossStakingRouter to it.
func DeployCrossStakingRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _crossStaking common.Address) (common.Address, *types.Transaction, *CrossStakingRouter, error) {
	parsed, err := CrossStakingRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossStakingRouterBin), backend, _crossStaking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossStakingRouter{CrossStakingRouterCaller: CrossStakingRouterCaller{contract: contract}, CrossStakingRouterTransactor: CrossStakingRouterTransactor{contract: contract}, CrossStakingRouterFilterer: CrossStakingRouterFilterer{contract: contract}}, nil
}

// CrossStakingRouter is an auto generated Go binding around an Ethereum contract.
type CrossStakingRouter struct {
	CrossStakingRouterCaller     // Read-only binding to the contract
	CrossStakingRouterTransactor // Write-only binding to the contract
	CrossStakingRouterFilterer   // Log filterer for contract events
}

// CrossStakingRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossStakingRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossStakingRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossStakingRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossStakingRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossStakingRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossStakingRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossStakingRouterSession struct {
	Contract     *CrossStakingRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CrossStakingRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossStakingRouterCallerSession struct {
	Contract *CrossStakingRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CrossStakingRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossStakingRouterTransactorSession struct {
	Contract     *CrossStakingRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CrossStakingRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossStakingRouterRaw struct {
	Contract *CrossStakingRouter // Generic contract binding to access the raw methods on
}

// CrossStakingRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossStakingRouterCallerRaw struct {
	Contract *CrossStakingRouterCaller // Generic read-only contract binding to access the raw methods on
}

// CrossStakingRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossStakingRouterTransactorRaw struct {
	Contract *CrossStakingRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossStakingRouter creates a new instance of CrossStakingRouter, bound to a specific deployed contract.
func NewCrossStakingRouter(address common.Address, backend bind.ContractBackend) (*CrossStakingRouter, error) {
	contract, err := bindCrossStakingRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouter{CrossStakingRouterCaller: CrossStakingRouterCaller{contract: contract}, CrossStakingRouterTransactor: CrossStakingRouterTransactor{contract: contract}, CrossStakingRouterFilterer: CrossStakingRouterFilterer{contract: contract}}, nil
}

// NewCrossStakingRouterCaller creates a new read-only instance of CrossStakingRouter, bound to a specific deployed contract.
func NewCrossStakingRouterCaller(address common.Address, caller bind.ContractCaller) (*CrossStakingRouterCaller, error) {
	contract, err := bindCrossStakingRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterCaller{contract: contract}, nil
}

// NewCrossStakingRouterTransactor creates a new write-only instance of CrossStakingRouter, bound to a specific deployed contract.
func NewCrossStakingRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossStakingRouterTransactor, error) {
	contract, err := bindCrossStakingRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterTransactor{contract: contract}, nil
}

// NewCrossStakingRouterFilterer creates a new log filterer instance of CrossStakingRouter, bound to a specific deployed contract.
func NewCrossStakingRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossStakingRouterFilterer, error) {
	contract, err := bindCrossStakingRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterFilterer{contract: contract}, nil
}

// bindCrossStakingRouter binds a generic wrapper to an already deployed contract.
func bindCrossStakingRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossStakingRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossStakingRouter *CrossStakingRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossStakingRouter.Contract.CrossStakingRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossStakingRouter *CrossStakingRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.CrossStakingRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossStakingRouter *CrossStakingRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.CrossStakingRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossStakingRouter *CrossStakingRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossStakingRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossStakingRouter *CrossStakingRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossStakingRouter *CrossStakingRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.contract.Transact(opts, method, params...)
}

// CrossStaking is a free data retrieval call binding the contract method 0x7dd96ac4.
//
// Solidity: function crossStaking() view returns(address)
func (_CrossStakingRouter *CrossStakingRouterCaller) CrossStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStakingRouter.contract.Call(opts, &out, "crossStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossStaking is a free data retrieval call binding the contract method 0x7dd96ac4.
//
// Solidity: function crossStaking() view returns(address)
func (_CrossStakingRouter *CrossStakingRouterSession) CrossStaking() (common.Address, error) {
	return _CrossStakingRouter.Contract.CrossStaking(&_CrossStakingRouter.CallOpts)
}

// CrossStaking is a free data retrieval call binding the contract method 0x7dd96ac4.
//
// Solidity: function crossStaking() view returns(address)
func (_CrossStakingRouter *CrossStakingRouterCallerSession) CrossStaking() (common.Address, error) {
	return _CrossStakingRouter.Contract.CrossStaking(&_CrossStakingRouter.CallOpts)
}

// GetUserStakingInfo is a free data retrieval call binding the contract method 0x128644b0.
//
// Solidity: function getUserStakingInfo(uint256 poolId, address user) view returns(uint256 stakedAmount, address[] rewardTokens, uint256[] pendingRewards)
func (_CrossStakingRouter *CrossStakingRouterCaller) GetUserStakingInfo(opts *bind.CallOpts, poolId *big.Int, user common.Address) (struct {
	StakedAmount   *big.Int
	RewardTokens   []common.Address
	PendingRewards []*big.Int
}, error) {
	var out []interface{}
	err := _CrossStakingRouter.contract.Call(opts, &out, "getUserStakingInfo", poolId, user)

	outstruct := new(struct {
		StakedAmount   *big.Int
		RewardTokens   []common.Address
		PendingRewards []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardTokens = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.PendingRewards = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetUserStakingInfo is a free data retrieval call binding the contract method 0x128644b0.
//
// Solidity: function getUserStakingInfo(uint256 poolId, address user) view returns(uint256 stakedAmount, address[] rewardTokens, uint256[] pendingRewards)
func (_CrossStakingRouter *CrossStakingRouterSession) GetUserStakingInfo(poolId *big.Int, user common.Address) (struct {
	StakedAmount   *big.Int
	RewardTokens   []common.Address
	PendingRewards []*big.Int
}, error) {
	return _CrossStakingRouter.Contract.GetUserStakingInfo(&_CrossStakingRouter.CallOpts, poolId, user)
}

// GetUserStakingInfo is a free data retrieval call binding the contract method 0x128644b0.
//
// Solidity: function getUserStakingInfo(uint256 poolId, address user) view returns(uint256 stakedAmount, address[] rewardTokens, uint256[] pendingRewards)
func (_CrossStakingRouter *CrossStakingRouterCallerSession) GetUserStakingInfo(poolId *big.Int, user common.Address) (struct {
	StakedAmount   *big.Int
	RewardTokens   []common.Address
	PendingRewards []*big.Int
}, error) {
	return _CrossStakingRouter.Contract.GetUserStakingInfo(&_CrossStakingRouter.CallOpts, poolId, user)
}

// IsNativePool is a free data retrieval call binding the contract method 0x7873d5a6.
//
// Solidity: function isNativePool(uint256 poolId) view returns(bool)
func (_CrossStakingRouter *CrossStakingRouterCaller) IsNativePool(opts *bind.CallOpts, poolId *big.Int) (bool, error) {
	var out []interface{}
	err := _CrossStakingRouter.contract.Call(opts, &out, "isNativePool", poolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNativePool is a free data retrieval call binding the contract method 0x7873d5a6.
//
// Solidity: function isNativePool(uint256 poolId) view returns(bool)
func (_CrossStakingRouter *CrossStakingRouterSession) IsNativePool(poolId *big.Int) (bool, error) {
	return _CrossStakingRouter.Contract.IsNativePool(&_CrossStakingRouter.CallOpts, poolId)
}

// IsNativePool is a free data retrieval call binding the contract method 0x7873d5a6.
//
// Solidity: function isNativePool(uint256 poolId) view returns(bool)
func (_CrossStakingRouter *CrossStakingRouterCallerSession) IsNativePool(poolId *big.Int) (bool, error) {
	return _CrossStakingRouter.Contract.IsNativePool(&_CrossStakingRouter.CallOpts, poolId)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossStakingRouter *CrossStakingRouterCaller) Wcross(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossStakingRouter.contract.Call(opts, &out, "wcross")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossStakingRouter *CrossStakingRouterSession) Wcross() (common.Address, error) {
	return _CrossStakingRouter.Contract.Wcross(&_CrossStakingRouter.CallOpts)
}

// Wcross is a free data retrieval call binding the contract method 0xa2db4582.
//
// Solidity: function wcross() view returns(address)
func (_CrossStakingRouter *CrossStakingRouterCallerSession) Wcross() (common.Address, error) {
	return _CrossStakingRouter.Contract.Wcross(&_CrossStakingRouter.CallOpts)
}

// StakeERC20 is a paid mutator transaction binding the contract method 0x37d9e9cc.
//
// Solidity: function stakeERC20(uint256 poolId, uint256 amount) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactor) StakeERC20(opts *bind.TransactOpts, poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.contract.Transact(opts, "stakeERC20", poolId, amount)
}

// StakeERC20 is a paid mutator transaction binding the contract method 0x37d9e9cc.
//
// Solidity: function stakeERC20(uint256 poolId, uint256 amount) returns()
func (_CrossStakingRouter *CrossStakingRouterSession) StakeERC20(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.StakeERC20(&_CrossStakingRouter.TransactOpts, poolId, amount)
}

// StakeERC20 is a paid mutator transaction binding the contract method 0x37d9e9cc.
//
// Solidity: function stakeERC20(uint256 poolId, uint256 amount) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactorSession) StakeERC20(poolId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.StakeERC20(&_CrossStakingRouter.TransactOpts, poolId, amount)
}

// StakeERC20WithPermit is a paid mutator transaction binding the contract method 0x08466d63.
//
// Solidity: function stakeERC20WithPermit(uint256 poolId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactor) StakeERC20WithPermit(opts *bind.TransactOpts, poolId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossStakingRouter.contract.Transact(opts, "stakeERC20WithPermit", poolId, amount, deadline, v, r, s)
}

// StakeERC20WithPermit is a paid mutator transaction binding the contract method 0x08466d63.
//
// Solidity: function stakeERC20WithPermit(uint256 poolId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossStakingRouter *CrossStakingRouterSession) StakeERC20WithPermit(poolId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.StakeERC20WithPermit(&_CrossStakingRouter.TransactOpts, poolId, amount, deadline, v, r, s)
}

// StakeERC20WithPermit is a paid mutator transaction binding the contract method 0x08466d63.
//
// Solidity: function stakeERC20WithPermit(uint256 poolId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactorSession) StakeERC20WithPermit(poolId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.StakeERC20WithPermit(&_CrossStakingRouter.TransactOpts, poolId, amount, deadline, v, r, s)
}

// StakeNative is a paid mutator transaction binding the contract method 0xe63c6bf0.
//
// Solidity: function stakeNative(uint256 poolId) payable returns()
func (_CrossStakingRouter *CrossStakingRouterTransactor) StakeNative(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.contract.Transact(opts, "stakeNative", poolId)
}

// StakeNative is a paid mutator transaction binding the contract method 0xe63c6bf0.
//
// Solidity: function stakeNative(uint256 poolId) payable returns()
func (_CrossStakingRouter *CrossStakingRouterSession) StakeNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.StakeNative(&_CrossStakingRouter.TransactOpts, poolId)
}

// StakeNative is a paid mutator transaction binding the contract method 0xe63c6bf0.
//
// Solidity: function stakeNative(uint256 poolId) payable returns()
func (_CrossStakingRouter *CrossStakingRouterTransactorSession) StakeNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.StakeNative(&_CrossStakingRouter.TransactOpts, poolId)
}

// UnstakeERC20 is a paid mutator transaction binding the contract method 0x7dfae334.
//
// Solidity: function unstakeERC20(uint256 poolId) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactor) UnstakeERC20(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.contract.Transact(opts, "unstakeERC20", poolId)
}

// UnstakeERC20 is a paid mutator transaction binding the contract method 0x7dfae334.
//
// Solidity: function unstakeERC20(uint256 poolId) returns()
func (_CrossStakingRouter *CrossStakingRouterSession) UnstakeERC20(poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.UnstakeERC20(&_CrossStakingRouter.TransactOpts, poolId)
}

// UnstakeERC20 is a paid mutator transaction binding the contract method 0x7dfae334.
//
// Solidity: function unstakeERC20(uint256 poolId) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactorSession) UnstakeERC20(poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.UnstakeERC20(&_CrossStakingRouter.TransactOpts, poolId)
}

// UnstakeNative is a paid mutator transaction binding the contract method 0xff8ab5f1.
//
// Solidity: function unstakeNative(uint256 poolId) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactor) UnstakeNative(opts *bind.TransactOpts, poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.contract.Transact(opts, "unstakeNative", poolId)
}

// UnstakeNative is a paid mutator transaction binding the contract method 0xff8ab5f1.
//
// Solidity: function unstakeNative(uint256 poolId) returns()
func (_CrossStakingRouter *CrossStakingRouterSession) UnstakeNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.UnstakeNative(&_CrossStakingRouter.TransactOpts, poolId)
}

// UnstakeNative is a paid mutator transaction binding the contract method 0xff8ab5f1.
//
// Solidity: function unstakeNative(uint256 poolId) returns()
func (_CrossStakingRouter *CrossStakingRouterTransactorSession) UnstakeNative(poolId *big.Int) (*types.Transaction, error) {
	return _CrossStakingRouter.Contract.UnstakeNative(&_CrossStakingRouter.TransactOpts, poolId)
}

// CrossStakingRouterStakedERC20Iterator is returned from FilterStakedERC20 and is used to iterate over the raw logs and unpacked data for StakedERC20 events raised by the CrossStakingRouter contract.
type CrossStakingRouterStakedERC20Iterator struct {
	Event *CrossStakingRouterStakedERC20 // Event containing the contract specifics and raw log

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
func (it *CrossStakingRouterStakedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRouterStakedERC20)
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
		it.Event = new(CrossStakingRouterStakedERC20)
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
func (it *CrossStakingRouterStakedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRouterStakedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRouterStakedERC20 represents a StakedERC20 event raised by the CrossStakingRouter contract.
type CrossStakingRouterStakedERC20 struct {
	User   common.Address
	PoolId *big.Int
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakedERC20 is a free log retrieval operation binding the contract event 0x61857636d15275e09c9adf9b762a4dbb040d0c12f7e000ad7f71bd8639de94ad.
//
// Solidity: event StakedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) FilterStakedERC20(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossStakingRouterStakedERC20Iterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.FilterLogs(opts, "StakedERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterStakedERC20Iterator{contract: _CrossStakingRouter.contract, event: "StakedERC20", logs: logs, sub: sub}, nil
}

// WatchStakedERC20 is a free log subscription operation binding the contract event 0x61857636d15275e09c9adf9b762a4dbb040d0c12f7e000ad7f71bd8639de94ad.
//
// Solidity: event StakedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) WatchStakedERC20(opts *bind.WatchOpts, sink chan<- *CrossStakingRouterStakedERC20, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.WatchLogs(opts, "StakedERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRouterStakedERC20)
				if err := _CrossStakingRouter.contract.UnpackLog(event, "StakedERC20", log); err != nil {
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

// ParseStakedERC20 is a log parse operation binding the contract event 0x61857636d15275e09c9adf9b762a4dbb040d0c12f7e000ad7f71bd8639de94ad.
//
// Solidity: event StakedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) ParseStakedERC20(log types.Log) (*CrossStakingRouterStakedERC20, error) {
	event := new(CrossStakingRouterStakedERC20)
	if err := _CrossStakingRouter.contract.UnpackLog(event, "StakedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRouterStakedNativeIterator is returned from FilterStakedNative and is used to iterate over the raw logs and unpacked data for StakedNative events raised by the CrossStakingRouter contract.
type CrossStakingRouterStakedNativeIterator struct {
	Event *CrossStakingRouterStakedNative // Event containing the contract specifics and raw log

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
func (it *CrossStakingRouterStakedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRouterStakedNative)
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
		it.Event = new(CrossStakingRouterStakedNative)
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
func (it *CrossStakingRouterStakedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRouterStakedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRouterStakedNative represents a StakedNative event raised by the CrossStakingRouter contract.
type CrossStakingRouterStakedNative struct {
	User   common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakedNative is a free log retrieval operation binding the contract event 0x94a750b81c0cca844f1d86628787e6f8444b9f979cd708c31ebbbe62984dc2ac.
//
// Solidity: event StakedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) FilterStakedNative(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossStakingRouterStakedNativeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.FilterLogs(opts, "StakedNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterStakedNativeIterator{contract: _CrossStakingRouter.contract, event: "StakedNative", logs: logs, sub: sub}, nil
}

// WatchStakedNative is a free log subscription operation binding the contract event 0x94a750b81c0cca844f1d86628787e6f8444b9f979cd708c31ebbbe62984dc2ac.
//
// Solidity: event StakedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) WatchStakedNative(opts *bind.WatchOpts, sink chan<- *CrossStakingRouterStakedNative, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.WatchLogs(opts, "StakedNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRouterStakedNative)
				if err := _CrossStakingRouter.contract.UnpackLog(event, "StakedNative", log); err != nil {
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

// ParseStakedNative is a log parse operation binding the contract event 0x94a750b81c0cca844f1d86628787e6f8444b9f979cd708c31ebbbe62984dc2ac.
//
// Solidity: event StakedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) ParseStakedNative(log types.Log) (*CrossStakingRouterStakedNative, error) {
	event := new(CrossStakingRouterStakedNative)
	if err := _CrossStakingRouter.contract.UnpackLog(event, "StakedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRouterUnstakedERC20Iterator is returned from FilterUnstakedERC20 and is used to iterate over the raw logs and unpacked data for UnstakedERC20 events raised by the CrossStakingRouter contract.
type CrossStakingRouterUnstakedERC20Iterator struct {
	Event *CrossStakingRouterUnstakedERC20 // Event containing the contract specifics and raw log

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
func (it *CrossStakingRouterUnstakedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRouterUnstakedERC20)
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
		it.Event = new(CrossStakingRouterUnstakedERC20)
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
func (it *CrossStakingRouterUnstakedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRouterUnstakedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRouterUnstakedERC20 represents a UnstakedERC20 event raised by the CrossStakingRouter contract.
type CrossStakingRouterUnstakedERC20 struct {
	User   common.Address
	PoolId *big.Int
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstakedERC20 is a free log retrieval operation binding the contract event 0x5571be890ca705d540f4e7ec30a60542bd07eedbc800749ac6c5784d0fa7a566.
//
// Solidity: event UnstakedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) FilterUnstakedERC20(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossStakingRouterUnstakedERC20Iterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.FilterLogs(opts, "UnstakedERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterUnstakedERC20Iterator{contract: _CrossStakingRouter.contract, event: "UnstakedERC20", logs: logs, sub: sub}, nil
}

// WatchUnstakedERC20 is a free log subscription operation binding the contract event 0x5571be890ca705d540f4e7ec30a60542bd07eedbc800749ac6c5784d0fa7a566.
//
// Solidity: event UnstakedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) WatchUnstakedERC20(opts *bind.WatchOpts, sink chan<- *CrossStakingRouterUnstakedERC20, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.WatchLogs(opts, "UnstakedERC20", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRouterUnstakedERC20)
				if err := _CrossStakingRouter.contract.UnpackLog(event, "UnstakedERC20", log); err != nil {
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

// ParseUnstakedERC20 is a log parse operation binding the contract event 0x5571be890ca705d540f4e7ec30a60542bd07eedbc800749ac6c5784d0fa7a566.
//
// Solidity: event UnstakedERC20(address indexed user, uint256 indexed poolId, address token, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) ParseUnstakedERC20(log types.Log) (*CrossStakingRouterUnstakedERC20, error) {
	event := new(CrossStakingRouterUnstakedERC20)
	if err := _CrossStakingRouter.contract.UnpackLog(event, "UnstakedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossStakingRouterUnstakedNativeIterator is returned from FilterUnstakedNative and is used to iterate over the raw logs and unpacked data for UnstakedNative events raised by the CrossStakingRouter contract.
type CrossStakingRouterUnstakedNativeIterator struct {
	Event *CrossStakingRouterUnstakedNative // Event containing the contract specifics and raw log

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
func (it *CrossStakingRouterUnstakedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossStakingRouterUnstakedNative)
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
		it.Event = new(CrossStakingRouterUnstakedNative)
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
func (it *CrossStakingRouterUnstakedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossStakingRouterUnstakedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossStakingRouterUnstakedNative represents a UnstakedNative event raised by the CrossStakingRouter contract.
type CrossStakingRouterUnstakedNative struct {
	User   common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstakedNative is a free log retrieval operation binding the contract event 0xc3f10c16db5e7a7551ad91c8ed696eefc6a5b65ca66feaeead4a4cf575c1344b.
//
// Solidity: event UnstakedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) FilterUnstakedNative(opts *bind.FilterOpts, user []common.Address, poolId []*big.Int) (*CrossStakingRouterUnstakedNativeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.FilterLogs(opts, "UnstakedNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossStakingRouterUnstakedNativeIterator{contract: _CrossStakingRouter.contract, event: "UnstakedNative", logs: logs, sub: sub}, nil
}

// WatchUnstakedNative is a free log subscription operation binding the contract event 0xc3f10c16db5e7a7551ad91c8ed696eefc6a5b65ca66feaeead4a4cf575c1344b.
//
// Solidity: event UnstakedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) WatchUnstakedNative(opts *bind.WatchOpts, sink chan<- *CrossStakingRouterUnstakedNative, user []common.Address, poolId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _CrossStakingRouter.contract.WatchLogs(opts, "UnstakedNative", userRule, poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossStakingRouterUnstakedNative)
				if err := _CrossStakingRouter.contract.UnpackLog(event, "UnstakedNative", log); err != nil {
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

// ParseUnstakedNative is a log parse operation binding the contract event 0xc3f10c16db5e7a7551ad91c8ed696eefc6a5b65ca66feaeead4a4cf575c1344b.
//
// Solidity: event UnstakedNative(address indexed user, uint256 indexed poolId, uint256 amount)
func (_CrossStakingRouter *CrossStakingRouterFilterer) ParseUnstakedNative(log types.Log) (*CrossStakingRouterUnstakedNative, error) {
	event := new(CrossStakingRouterUnstakedNative)
	if err := _CrossStakingRouter.contract.UnpackLog(event, "UnstakedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
