// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// EccMetaData contains all meta data concerning the Ecc contract.
var EccMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p0\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p1\",\"type\":\"uint256[2]\"}],\"name\":\"ecAdd\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecAdds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ecMul\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecMuls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a867d0ca": "ecAdd(uint256[2],uint256[2])",
		"7de7133d": "ecAdds(uint256)",
		"71de3724": "ecMul(uint256[2],uint256)",
		"f602cf71": "ecMuls(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610411806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806371de3724146100515780637de7133d1461007a578063a867d0ca1461008f578063f602cf71146100a2575b600080fd5b61006461005f3660046102f4565b6100b5565b604051610071919061031f565b60405180910390f35b61008d610088366004610350565b610126565b005b61006461009d366004610369565b61017a565b61008d6100b0366004610350565b610217565b6100bd610258565b60006040518060600160405280856000600281106100dd576100dd61039e565b60200201518152602001856001600281106100fa576100fa61039e565b602002015181526020018481525090506040826060836007600019fa61011f57600080fd5b5092915050565b61012e610258565b6001815260026020820152610141610258565b600181526002602082015260005b8381101561017457610161838361017a565b508061016c816103b4565b91505061014f565b50505050565b610182610258565b60006040518060800160405280856000600281106101a2576101a261039e565b60200201518152602001856001600281106101bf576101bf61039e565b60200201518152602001846000600281106101dc576101dc61039e565b60200201518152602001846001600281106101f9576101f961039e565b6020020151905290506040826080836006600019fa61011f57600080fd5b61021f610258565b600181526002602082015260005b82811015610253576102408260036100b5565b508061024b816103b4565b91505061022d565b505050565b60405180604001604052806002906020820280368337509192915050565b600082601f83011261028757600080fd5b6040516040810181811067ffffffffffffffff821117156102b857634e487b7160e01b600052604160045260246000fd5b80604052508060408401858111156102cf57600080fd5b845b818110156102e95780358352602092830192016102d1565b509195945050505050565b6000806060838503121561030757600080fd5b6103118484610276565b946040939093013593505050565b60408101818360005b6002811015610347578151835260209283019290910190600101610328565b50505092915050565b60006020828403121561036257600080fd5b5035919050565b6000806080838503121561037c57600080fd5b6103868484610276565b91506103958460408501610276565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b6000600182016103d457634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122055137f6c6300e199954b29bd03922370828d013bb32ba1d1741b7add1075c4a164736f6c634300080f0033",
}

// EccABI is the input ABI used to generate the binding from.
// Deprecated: Use EccMetaData.ABI instead.
var EccABI = EccMetaData.ABI

// Deprecated: Use EccMetaData.Sigs instead.
// EccFuncSigs maps the 4-byte function signature to its string representation.
var EccFuncSigs = EccMetaData.Sigs

// EccBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EccMetaData.Bin instead.
var EccBin = EccMetaData.Bin

// DeployEcc deploys a new Ethereum contract, binding an instance of Ecc to it.
func DeployEcc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ecc, error) {
	parsed, err := EccMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EccBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ecc{EccCaller: EccCaller{contract: contract}, EccTransactor: EccTransactor{contract: contract}, EccFilterer: EccFilterer{contract: contract}}, nil
}

// Ecc is an auto generated Go binding around an Ethereum contract.
type Ecc struct {
	EccCaller     // Read-only binding to the contract
	EccTransactor // Write-only binding to the contract
	EccFilterer   // Log filterer for contract events
}

// EccCaller is an auto generated read-only Go binding around an Ethereum contract.
type EccCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EccTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EccTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EccFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EccFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EccSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EccSession struct {
	Contract     *Ecc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EccCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EccCallerSession struct {
	Contract *EccCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EccTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EccTransactorSession struct {
	Contract     *EccTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EccRaw is an auto generated low-level Go binding around an Ethereum contract.
type EccRaw struct {
	Contract *Ecc // Generic contract binding to access the raw methods on
}

// EccCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EccCallerRaw struct {
	Contract *EccCaller // Generic read-only contract binding to access the raw methods on
}

// EccTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EccTransactorRaw struct {
	Contract *EccTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcc creates a new instance of Ecc, bound to a specific deployed contract.
func NewEcc(address common.Address, backend bind.ContractBackend) (*Ecc, error) {
	contract, err := bindEcc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecc{EccCaller: EccCaller{contract: contract}, EccTransactor: EccTransactor{contract: contract}, EccFilterer: EccFilterer{contract: contract}}, nil
}

// NewEccCaller creates a new read-only instance of Ecc, bound to a specific deployed contract.
func NewEccCaller(address common.Address, caller bind.ContractCaller) (*EccCaller, error) {
	contract, err := bindEcc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EccCaller{contract: contract}, nil
}

// NewEccTransactor creates a new write-only instance of Ecc, bound to a specific deployed contract.
func NewEccTransactor(address common.Address, transactor bind.ContractTransactor) (*EccTransactor, error) {
	contract, err := bindEcc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EccTransactor{contract: contract}, nil
}

// NewEccFilterer creates a new log filterer instance of Ecc, bound to a specific deployed contract.
func NewEccFilterer(address common.Address, filterer bind.ContractFilterer) (*EccFilterer, error) {
	contract, err := bindEcc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EccFilterer{contract: contract}, nil
}

// bindEcc binds a generic wrapper to an already deployed contract.
func bindEcc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EccABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecc *EccRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecc.Contract.EccCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecc *EccRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecc.Contract.EccTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecc *EccRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecc.Contract.EccTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecc *EccCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecc *EccTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecc *EccTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecc.Contract.contract.Transact(opts, method, params...)
}

// EcAdd is a free data retrieval call binding the contract method 0xa867d0ca.
//
// Solidity: function ecAdd(uint256[2] p0, uint256[2] p1) view returns(uint256[2] retP)
func (_Ecc *EccCaller) EcAdd(opts *bind.CallOpts, p0 [2]*big.Int, p1 [2]*big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Ecc.contract.Call(opts, &out, "ecAdd", p0, p1)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// EcAdd is a free data retrieval call binding the contract method 0xa867d0ca.
//
// Solidity: function ecAdd(uint256[2] p0, uint256[2] p1) view returns(uint256[2] retP)
func (_Ecc *EccSession) EcAdd(p0 [2]*big.Int, p1 [2]*big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcAdd(&_Ecc.CallOpts, p0, p1)
}

// EcAdd is a free data retrieval call binding the contract method 0xa867d0ca.
//
// Solidity: function ecAdd(uint256[2] p0, uint256[2] p1) view returns(uint256[2] retP)
func (_Ecc *EccCallerSession) EcAdd(p0 [2]*big.Int, p1 [2]*big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcAdd(&_Ecc.CallOpts, p0, p1)
}

// EcMul is a free data retrieval call binding the contract method 0x71de3724.
//
// Solidity: function ecMul(uint256[2] p, uint256 s) view returns(uint256[2] retP)
func (_Ecc *EccCaller) EcMul(opts *bind.CallOpts, p [2]*big.Int, s *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Ecc.contract.Call(opts, &out, "ecMul", p, s)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// EcMul is a free data retrieval call binding the contract method 0x71de3724.
//
// Solidity: function ecMul(uint256[2] p, uint256 s) view returns(uint256[2] retP)
func (_Ecc *EccSession) EcMul(p [2]*big.Int, s *big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcMul(&_Ecc.CallOpts, p, s)
}

// EcMul is a free data retrieval call binding the contract method 0x71de3724.
//
// Solidity: function ecMul(uint256[2] p, uint256 s) view returns(uint256[2] retP)
func (_Ecc *EccCallerSession) EcMul(p [2]*big.Int, s *big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcMul(&_Ecc.CallOpts, p, s)
}

// EcAdds is a paid mutator transaction binding the contract method 0x7de7133d.
//
// Solidity: function ecAdds(uint256 n) returns()
func (_Ecc *EccTransactor) EcAdds(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecAdds", n)
}

// EcAdds is a paid mutator transaction binding the contract method 0x7de7133d.
//
// Solidity: function ecAdds(uint256 n) returns()
func (_Ecc *EccSession) EcAdds(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcAdds(&_Ecc.TransactOpts, n)
}

// EcAdds is a paid mutator transaction binding the contract method 0x7de7133d.
//
// Solidity: function ecAdds(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcAdds(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcAdds(&_Ecc.TransactOpts, n)
}

// EcMuls is a paid mutator transaction binding the contract method 0xf602cf71.
//
// Solidity: function ecMuls(uint256 n) returns()
func (_Ecc *EccTransactor) EcMuls(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecMuls", n)
}

// EcMuls is a paid mutator transaction binding the contract method 0xf602cf71.
//
// Solidity: function ecMuls(uint256 n) returns()
func (_Ecc *EccSession) EcMuls(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcMuls(&_Ecc.TransactOpts, n)
}

// EcMuls is a paid mutator transaction binding the contract method 0xf602cf71.
//
// Solidity: function ecMuls(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcMuls(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcMuls(&_Ecc.TransactOpts, n)
}
