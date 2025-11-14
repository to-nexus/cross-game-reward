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

// WCROSSMetaData contains all meta data concerning the WCROSS contract.
var WCROSSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameReward\",\"outputs\":[{\"internalType\":\"contractCrossGameReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WCROSSInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WCROSSInvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WCROSSTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WCROSSUnauthorized\",\"type\":\"error\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"d0e30db0": "deposit()",
		"bf3308de": "gameReward()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
		"205c2878": "withdrawTo(address,uint256)",
	},
	Bin: "0x60a060405234801561000f575f5ffd5b506040518060400160405280600d81526020016c577261707065642043524f535360981b815250604051806040016040528060068152602001655743524f535360d01b81525081600390816100649190610115565b5060046100718282610115565b505033608052506101cf565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100a557607f821691505b6020821081036100c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561011057805f5260205f20601f840160051c810160208510156100ee5750805b601f840160051c820191505b8181101561010d575f81556001016100fa565b50505b505050565b81516001600160401b0381111561012e5761012e61007d565b6101428161013c8454610091565b846100c9565b6020601f821160018114610174575f831561015d5750848201515b5f19600385901b1c1916600184901b17845561010d565b5f84815260208120601f198516915b828110156101a35787850151825560209485019460019092019101610183565b50848210156101c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b608051610e556101f55f395f8181610258015281816102fa01526104d90152610e555ff3fe6080604052600436106100d1575f3560e01c8063313ce5671161007c578063a9059cbb11610057578063a9059cbb14610228578063bf3308de14610247578063d0e30db01461029f578063dd62ed3e146102a7575f5ffd5b8063313ce567146101b857806370a08231146101d357806395d89b4114610214575f5ffd5b8063205c2878116100ac578063205c28781461015b57806323b872dd1461017a5780632e1a7d4d14610199575f5ffd5b806306fdde03146100e4578063095ea7b31461010e57806318160ddd1461013d575f5ffd5b366100e0576100de6102f8565b005b5f5ffd5b3480156100ef575f5ffd5b506100f861042e565b6040516101059190610c2f565b60405180910390f35b348015610119575f5ffd5b5061012d610128366004610ca3565b6104be565b6040519015158152602001610105565b348015610148575f5ffd5b506002545b604051908152602001610105565b348015610166575f5ffd5b506100de610175366004610ca3565b6104d7565b348015610185575f5ffd5b5061012d610194366004610ccd565b6106ba565b3480156101a4575f5ffd5b506100de6101b3366004610d0b565b6106dd565b3480156101c3575f5ffd5b5060405160128152602001610105565b3480156101de575f5ffd5b5061014d6101ed366004610d22565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b34801561021f575f5ffd5b506100f86106ea565b348015610233575f5ffd5b5061012d610242366004610ca3565b6106f9565b348015610252575f5ffd5b5061027a7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b6100de6102f8565b3480156102b2575f5ffd5b5061014d6102c1366004610d44565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610361573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103859190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f3411610422576040517fd8df41ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61042c3334610706565b565b60606003805461043d90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461046990610d96565b80156104b45780601f1061048b576101008083540402835291602001916104b4565b820191905f5260205f20905b81548152906001019060200180831161049757829003601f168201915b5050505050905090565b5f336104cb818585610769565b60019150505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610540573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105649190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105c8576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610615576040517f653345a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61061f3382610776565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610675576040519150601f19603f3d011682016040523d82523d5f602084013e61067a565b606091505b50509050806106b5576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f336106c78582856107d0565b6106d285858561089e565b506001949350505050565b6106e733826104d7565b50565b60606004805461043d90610d96565b5f336104cb81858561089e565b73ffffffffffffffffffffffffffffffffffffffff821661075a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6107655f8383610943565b5050565b6106b58383836001610aea565b73ffffffffffffffffffffffffffffffffffffffff82166107c5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b610765825f83610943565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610898578181101561088a576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610751565b61089884848484035f610aea565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166108ed576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff821661093c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b6106b58383835b73ffffffffffffffffffffffffffffffffffffffff831661097a578060025f82825461096f9190610de7565b90915550610a2a9050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156109ff576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610751565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610a5357600280548290039055610a7e565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610add91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8416610b39576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8316610b88576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610898578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610c2191815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff811681146106e7575f5ffd5b5f5f60408385031215610cb4575f5ffd5b8235610cbf81610c82565b946020939093013593505050565b5f5f5f60608486031215610cdf575f5ffd5b8335610cea81610c82565b92506020840135610cfa81610c82565b929592945050506040919091013590565b5f60208284031215610d1b575f5ffd5b5035919050565b5f60208284031215610d32575f5ffd5b8135610d3d81610c82565b9392505050565b5f5f60408385031215610d55575f5ffd5b8235610d6081610c82565b91506020830135610d7081610c82565b809150509250929050565b5f60208284031215610d8b575f5ffd5b8151610d3d81610c82565b600181811c90821680610daa57607f821691505b602082108103610de1577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156104d1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212209ff99d1df7425b5a2d579ebe50cbad7ebcb40604a5fdfb17acf3da6fd6c753ef64736f6c634300081c0033",
}

// WCROSSABI is the input ABI used to generate the binding from.
// Deprecated: Use WCROSSMetaData.ABI instead.
var WCROSSABI = WCROSSMetaData.ABI

// WCROSSBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const WCROSSBinRuntime = "6080604052600436106100d1575f3560e01c8063313ce5671161007c578063a9059cbb11610057578063a9059cbb14610228578063bf3308de14610247578063d0e30db01461029f578063dd62ed3e146102a7575f5ffd5b8063313ce567146101b857806370a08231146101d357806395d89b4114610214575f5ffd5b8063205c2878116100ac578063205c28781461015b57806323b872dd1461017a5780632e1a7d4d14610199575f5ffd5b806306fdde03146100e4578063095ea7b31461010e57806318160ddd1461013d575f5ffd5b366100e0576100de6102f8565b005b5f5ffd5b3480156100ef575f5ffd5b506100f861042e565b6040516101059190610c2f565b60405180910390f35b348015610119575f5ffd5b5061012d610128366004610ca3565b6104be565b6040519015158152602001610105565b348015610148575f5ffd5b506002545b604051908152602001610105565b348015610166575f5ffd5b506100de610175366004610ca3565b6104d7565b348015610185575f5ffd5b5061012d610194366004610ccd565b6106ba565b3480156101a4575f5ffd5b506100de6101b3366004610d0b565b6106dd565b3480156101c3575f5ffd5b5060405160128152602001610105565b3480156101de575f5ffd5b5061014d6101ed366004610d22565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b34801561021f575f5ffd5b506100f86106ea565b348015610233575f5ffd5b5061012d610242366004610ca3565b6106f9565b348015610252575f5ffd5b5061027a7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610105565b6100de6102f8565b3480156102b2575f5ffd5b5061014d6102c1366004610d44565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610361573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103859190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f3411610422576040517fd8df41ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61042c3334610706565b565b60606003805461043d90610d96565b80601f016020809104026020016040519081016040528092919081815260200182805461046990610d96565b80156104b45780601f1061048b576101008083540402835291602001916104b4565b820191905f5260205f20905b81548152906001019060200180831161049757829003601f168201915b5050505050905090565b5f336104cb818585610769565b60019150505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f887ea406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610540573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105649190610d7b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105c8576040517f6bd46d6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610615576040517f653345a600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61061f3382610776565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114610675576040519150601f19603f3d011682016040523d82523d5f602084013e61067a565b606091505b50509050806106b5576040517f2997048700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f336106c78582856107d0565b6106d285858561089e565b506001949350505050565b6106e733826104d7565b50565b60606004805461043d90610d96565b5f336104cb81858561089e565b73ffffffffffffffffffffffffffffffffffffffff821661075a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024015b60405180910390fd5b6107655f8383610943565b5050565b6106b58383836001610aea565b73ffffffffffffffffffffffffffffffffffffffff82166107c5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b610765825f83610943565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610898578181101561088a576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610751565b61089884848484035f610aea565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166108ed576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff821661093c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b6106b58383835b73ffffffffffffffffffffffffffffffffffffffff831661097a578060025f82825461096f9190610de7565b90915550610a2a9050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156109ff576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610751565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610a5357600280548290039055610a7e565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610add91815260200190565b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8416610b39576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8316610b88576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610751565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610898578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610c2191815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff811681146106e7575f5ffd5b5f5f60408385031215610cb4575f5ffd5b8235610cbf81610c82565b946020939093013593505050565b5f5f5f60608486031215610cdf575f5ffd5b8335610cea81610c82565b92506020840135610cfa81610c82565b929592945050506040919091013590565b5f60208284031215610d1b575f5ffd5b5035919050565b5f60208284031215610d32575f5ffd5b8135610d3d81610c82565b9392505050565b5f5f60408385031215610d55575f5ffd5b8235610d6081610c82565b91506020830135610d7081610c82565b809150509250929050565b5f60208284031215610d8b575f5ffd5b8151610d3d81610c82565b600181811c90821680610daa57607f821691505b602082108103610de1577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156104d1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212209ff99d1df7425b5a2d579ebe50cbad7ebcb40604a5fdfb17acf3da6fd6c753ef64736f6c634300081c0033"

// Deprecated: Use WCROSSMetaData.Sigs instead.
// WCROSSFuncSigs maps the 4-byte function signature to its string representation.
var WCROSSFuncSigs = WCROSSMetaData.Sigs

// WCROSSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WCROSSMetaData.Bin instead.
var WCROSSBin = WCROSSMetaData.Bin

// DeployWCROSS deploys a new Ethereum contract, binding an instance of WCROSS to it.
func DeployWCROSS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WCROSS, error) {
	parsed, err := WCROSSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WCROSSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WCROSS{WCROSSCaller: WCROSSCaller{contract: contract}, WCROSSTransactor: WCROSSTransactor{contract: contract}, WCROSSFilterer: WCROSSFilterer{contract: contract}}, nil
}

// WCROSS is an auto generated Go binding around an Ethereum contract.
type WCROSS struct {
	WCROSSCaller     // Read-only binding to the contract
	WCROSSTransactor // Write-only binding to the contract
	WCROSSFilterer   // Log filterer for contract events
}

// WCROSSCaller is an auto generated read-only Go binding around an Ethereum contract.
type WCROSSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WCROSSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WCROSSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WCROSSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WCROSSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WCROSSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WCROSSSession struct {
	Contract     *WCROSS           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WCROSSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WCROSSCallerSession struct {
	Contract *WCROSSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WCROSSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WCROSSTransactorSession struct {
	Contract     *WCROSSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WCROSSRaw is an auto generated low-level Go binding around an Ethereum contract.
type WCROSSRaw struct {
	Contract *WCROSS // Generic contract binding to access the raw methods on
}

// WCROSSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WCROSSCallerRaw struct {
	Contract *WCROSSCaller // Generic read-only contract binding to access the raw methods on
}

// WCROSSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WCROSSTransactorRaw struct {
	Contract *WCROSSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWCROSS creates a new instance of WCROSS, bound to a specific deployed contract.
func NewWCROSS(address common.Address, backend bind.ContractBackend) (*WCROSS, error) {
	contract, err := bindWCROSS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WCROSS{WCROSSCaller: WCROSSCaller{contract: contract}, WCROSSTransactor: WCROSSTransactor{contract: contract}, WCROSSFilterer: WCROSSFilterer{contract: contract}}, nil
}

// NewWCROSSCaller creates a new read-only instance of WCROSS, bound to a specific deployed contract.
func NewWCROSSCaller(address common.Address, caller bind.ContractCaller) (*WCROSSCaller, error) {
	contract, err := bindWCROSS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WCROSSCaller{contract: contract}, nil
}

// NewWCROSSTransactor creates a new write-only instance of WCROSS, bound to a specific deployed contract.
func NewWCROSSTransactor(address common.Address, transactor bind.ContractTransactor) (*WCROSSTransactor, error) {
	contract, err := bindWCROSS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WCROSSTransactor{contract: contract}, nil
}

// NewWCROSSFilterer creates a new log filterer instance of WCROSS, bound to a specific deployed contract.
func NewWCROSSFilterer(address common.Address, filterer bind.ContractFilterer) (*WCROSSFilterer, error) {
	contract, err := bindWCROSS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WCROSSFilterer{contract: contract}, nil
}

// bindWCROSS binds a generic wrapper to an already deployed contract.
func bindWCROSS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WCROSSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WCROSS *WCROSSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WCROSS.Contract.WCROSSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WCROSS *WCROSSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WCROSS.Contract.WCROSSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WCROSS *WCROSSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WCROSS.Contract.WCROSSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WCROSS *WCROSSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WCROSS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WCROSS *WCROSSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WCROSS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WCROSS *WCROSSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WCROSS.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WCROSS *WCROSSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WCROSS *WCROSSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WCROSS.Contract.Allowance(&_WCROSS.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WCROSS *WCROSSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WCROSS.Contract.Allowance(&_WCROSS.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WCROSS *WCROSSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WCROSS *WCROSSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WCROSS.Contract.BalanceOf(&_WCROSS.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WCROSS *WCROSSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WCROSS.Contract.BalanceOf(&_WCROSS.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WCROSS *WCROSSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WCROSS *WCROSSSession) Decimals() (uint8, error) {
	return _WCROSS.Contract.Decimals(&_WCROSS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WCROSS *WCROSSCallerSession) Decimals() (uint8, error) {
	return _WCROSS.Contract.Decimals(&_WCROSS.CallOpts)
}

// GameReward is a free data retrieval call binding the contract method 0xbf3308de.
//
// Solidity: function gameReward() view returns(address)
func (_WCROSS *WCROSSCaller) GameReward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "gameReward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameReward is a free data retrieval call binding the contract method 0xbf3308de.
//
// Solidity: function gameReward() view returns(address)
func (_WCROSS *WCROSSSession) GameReward() (common.Address, error) {
	return _WCROSS.Contract.GameReward(&_WCROSS.CallOpts)
}

// GameReward is a free data retrieval call binding the contract method 0xbf3308de.
//
// Solidity: function gameReward() view returns(address)
func (_WCROSS *WCROSSCallerSession) GameReward() (common.Address, error) {
	return _WCROSS.Contract.GameReward(&_WCROSS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WCROSS *WCROSSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WCROSS *WCROSSSession) Name() (string, error) {
	return _WCROSS.Contract.Name(&_WCROSS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WCROSS *WCROSSCallerSession) Name() (string, error) {
	return _WCROSS.Contract.Name(&_WCROSS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WCROSS *WCROSSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WCROSS *WCROSSSession) Symbol() (string, error) {
	return _WCROSS.Contract.Symbol(&_WCROSS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WCROSS *WCROSSCallerSession) Symbol() (string, error) {
	return _WCROSS.Contract.Symbol(&_WCROSS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WCROSS *WCROSSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WCROSS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WCROSS *WCROSSSession) TotalSupply() (*big.Int, error) {
	return _WCROSS.Contract.TotalSupply(&_WCROSS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WCROSS *WCROSSCallerSession) TotalSupply() (*big.Int, error) {
	return _WCROSS.Contract.TotalSupply(&_WCROSS.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WCROSS *WCROSSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WCROSS *WCROSSSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.Approve(&_WCROSS.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WCROSS *WCROSSTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.Approve(&_WCROSS.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WCROSS *WCROSSTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WCROSS.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WCROSS *WCROSSSession) Deposit() (*types.Transaction, error) {
	return _WCROSS.Contract.Deposit(&_WCROSS.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WCROSS *WCROSSTransactorSession) Deposit() (*types.Transaction, error) {
	return _WCROSS.Contract.Deposit(&_WCROSS.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WCROSS *WCROSSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WCROSS *WCROSSSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.Transfer(&_WCROSS.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WCROSS *WCROSSTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.Transfer(&_WCROSS.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WCROSS *WCROSSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WCROSS *WCROSSSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.TransferFrom(&_WCROSS.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WCROSS *WCROSSTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.TransferFrom(&_WCROSS.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WCROSS *WCROSSTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WCROSS.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WCROSS *WCROSSSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.Withdraw(&_WCROSS.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WCROSS *WCROSSTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.Withdraw(&_WCROSS.TransactOpts, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_WCROSS *WCROSSTransactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WCROSS.contract.Transact(opts, "withdrawTo", to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_WCROSS *WCROSSSession) WithdrawTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.WithdrawTo(&_WCROSS.TransactOpts, to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address to, uint256 amount) returns()
func (_WCROSS *WCROSSTransactorSession) WithdrawTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WCROSS.Contract.WithdrawTo(&_WCROSS.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WCROSS *WCROSSTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WCROSS.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WCROSS *WCROSSSession) Receive() (*types.Transaction, error) {
	return _WCROSS.Contract.Receive(&_WCROSS.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WCROSS *WCROSSTransactorSession) Receive() (*types.Transaction, error) {
	return _WCROSS.Contract.Receive(&_WCROSS.TransactOpts)
}

// WCROSSApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WCROSS contract.
type WCROSSApprovalIterator struct {
	Event *WCROSSApproval // Event containing the contract specifics and raw log

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
func (it *WCROSSApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WCROSSApproval)
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
		it.Event = new(WCROSSApproval)
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
func (it *WCROSSApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WCROSSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WCROSSApproval represents a Approval event raised by the WCROSS contract.
type WCROSSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WCROSS *WCROSSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WCROSSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WCROSS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WCROSSApprovalIterator{contract: _WCROSS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WCROSS *WCROSSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WCROSSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WCROSS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WCROSSApproval)
				if err := _WCROSS.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WCROSS *WCROSSFilterer) ParseApproval(log types.Log) (*WCROSSApproval, error) {
	event := new(WCROSSApproval)
	if err := _WCROSS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WCROSSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WCROSS contract.
type WCROSSTransferIterator struct {
	Event *WCROSSTransfer // Event containing the contract specifics and raw log

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
func (it *WCROSSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WCROSSTransfer)
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
		it.Event = new(WCROSSTransfer)
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
func (it *WCROSSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WCROSSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WCROSSTransfer represents a Transfer event raised by the WCROSS contract.
type WCROSSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WCROSS *WCROSSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WCROSSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WCROSS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WCROSSTransferIterator{contract: _WCROSS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WCROSS *WCROSSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WCROSSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WCROSS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WCROSSTransfer)
				if err := _WCROSS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WCROSS *WCROSSFilterer) ParseTransfer(log types.Log) (*WCROSSTransfer, error) {
	event := new(WCROSSTransfer)
	if err := _WCROSS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
