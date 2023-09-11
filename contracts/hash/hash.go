// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hash

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

// HashMetaData contains all meta data concerning the Hash contract.
var HashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"keccak256s\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sha256\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sha256Yul\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"sha256s\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"61fe787d": "keccak256s(uint256)",
		"bebc76dd": "sha256(bytes)",
		"fe0ed4c1": "sha256Yul(bytes)",
		"61ef0ccb": "sha256s(uint256)",
	},
	Bin: "0x608060405234801561000f575f80fd5b506103608061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806361ef0ccb1461004e57806361fe787d14610063578063bebc76dd14610076578063fe0ed4c11461009f575b5f80fd5b61006161005c3660046101c1565b6100b2565b005b6100616100713660046101c1565b6100fd565b6100896100843660046101ec565b610139565b60405161009691906102b9565b60405180910390f35b6100896100ad3660046101ec565b6101a8565b604080516103e760208201525f910160405160208183030381529060405290505f5b828110156100f8576100e582610139565b50806100f0816102eb565b9150506100d4565b505050565b604080516103e760208201525f910160405160208183030381529060405290505f5b828110156100f85780610131816102eb565b91505061011f565b60605f8060026001600160a01b031684604051610156919061030f565b5f60405180830381855afa9150503d805f811461018e576040519150601f19603f3d011682016040523d82523d5f602084013e610193565b606091505b5091509150816101a1575f80fd5b5050919050565b606060205f60205f60025afa6101bc575f80fd5b919050565b5f602082840312156101d1575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f602082840312156101fc575f80fd5b813567ffffffffffffffff80821115610213575f80fd5b818401915084601f830112610226575f80fd5b813581811115610238576102386101d8565b604051601f8201601f19908116603f01168101908382118183101715610260576102606101d8565b81604052828152876020848701011115610278575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f5b838110156102b1578181015183820152602001610299565b50505f910152565b602081525f82518060208401526102d7816040850160208701610297565b601f01601f19169190910160400192915050565b5f6001820161030857634e487b7160e01b5f52601160045260245ffd5b5060010190565b5f8251610320818460208701610297565b919091019291505056fea2646970667358221220c9ee2457bd94cb508c3cb69d1a161cd28c3024f45986e1a0c1d776b72e9013d164736f6c63430008150033",
}

// HashABI is the input ABI used to generate the binding from.
// Deprecated: Use HashMetaData.ABI instead.
var HashABI = HashMetaData.ABI

// Deprecated: Use HashMetaData.Sigs instead.
// HashFuncSigs maps the 4-byte function signature to its string representation.
var HashFuncSigs = HashMetaData.Sigs

// HashBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashMetaData.Bin instead.
var HashBin = HashMetaData.Bin

// DeployHash deploys a new Ethereum contract, binding an instance of Hash to it.
func DeployHash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hash, error) {
	parsed, err := HashMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hash{HashCaller: HashCaller{contract: contract}, HashTransactor: HashTransactor{contract: contract}, HashFilterer: HashFilterer{contract: contract}}, nil
}

// Hash is an auto generated Go binding around an Ethereum contract.
type Hash struct {
	HashCaller     // Read-only binding to the contract
	HashTransactor // Write-only binding to the contract
	HashFilterer   // Log filterer for contract events
}

// HashCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashSession struct {
	Contract     *Hash             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashCallerSession struct {
	Contract *HashCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashTransactorSession struct {
	Contract     *HashTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashRaw struct {
	Contract *Hash // Generic contract binding to access the raw methods on
}

// HashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashCallerRaw struct {
	Contract *HashCaller // Generic read-only contract binding to access the raw methods on
}

// HashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashTransactorRaw struct {
	Contract *HashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHash creates a new instance of Hash, bound to a specific deployed contract.
func NewHash(address common.Address, backend bind.ContractBackend) (*Hash, error) {
	contract, err := bindHash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hash{HashCaller: HashCaller{contract: contract}, HashTransactor: HashTransactor{contract: contract}, HashFilterer: HashFilterer{contract: contract}}, nil
}

// NewHashCaller creates a new read-only instance of Hash, bound to a specific deployed contract.
func NewHashCaller(address common.Address, caller bind.ContractCaller) (*HashCaller, error) {
	contract, err := bindHash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashCaller{contract: contract}, nil
}

// NewHashTransactor creates a new write-only instance of Hash, bound to a specific deployed contract.
func NewHashTransactor(address common.Address, transactor bind.ContractTransactor) (*HashTransactor, error) {
	contract, err := bindHash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashTransactor{contract: contract}, nil
}

// NewHashFilterer creates a new log filterer instance of Hash, bound to a specific deployed contract.
func NewHashFilterer(address common.Address, filterer bind.ContractFilterer) (*HashFilterer, error) {
	contract, err := bindHash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashFilterer{contract: contract}, nil
}

// bindHash binds a generic wrapper to an already deployed contract.
func bindHash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hash *HashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hash.Contract.HashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hash *HashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hash.Contract.HashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hash *HashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hash.Contract.HashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hash *HashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hash *HashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hash *HashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hash.Contract.contract.Transact(opts, method, params...)
}

// Sha256 is a free data retrieval call binding the contract method 0xbebc76dd.
//
// Solidity: function sha256(bytes input) view returns(bytes out)
func (_Hash *HashCaller) Sha256(opts *bind.CallOpts, input []byte) ([]byte, error) {
	var out []interface{}
	err := _Hash.contract.Call(opts, &out, "sha256", input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Sha256 is a free data retrieval call binding the contract method 0xbebc76dd.
//
// Solidity: function sha256(bytes input) view returns(bytes out)
func (_Hash *HashSession) Sha256(input []byte) ([]byte, error) {
	return _Hash.Contract.Sha256(&_Hash.CallOpts, input)
}

// Sha256 is a free data retrieval call binding the contract method 0xbebc76dd.
//
// Solidity: function sha256(bytes input) view returns(bytes out)
func (_Hash *HashCallerSession) Sha256(input []byte) ([]byte, error) {
	return _Hash.Contract.Sha256(&_Hash.CallOpts, input)
}

// Sha256Yul is a free data retrieval call binding the contract method 0xfe0ed4c1.
//
// Solidity: function sha256Yul(bytes input) view returns(bytes out)
func (_Hash *HashCaller) Sha256Yul(opts *bind.CallOpts, input []byte) ([]byte, error) {
	var out []interface{}
	err := _Hash.contract.Call(opts, &out, "sha256Yul", input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Sha256Yul is a free data retrieval call binding the contract method 0xfe0ed4c1.
//
// Solidity: function sha256Yul(bytes input) view returns(bytes out)
func (_Hash *HashSession) Sha256Yul(input []byte) ([]byte, error) {
	return _Hash.Contract.Sha256Yul(&_Hash.CallOpts, input)
}

// Sha256Yul is a free data retrieval call binding the contract method 0xfe0ed4c1.
//
// Solidity: function sha256Yul(bytes input) view returns(bytes out)
func (_Hash *HashCallerSession) Sha256Yul(input []byte) ([]byte, error) {
	return _Hash.Contract.Sha256Yul(&_Hash.CallOpts, input)
}

// Keccak256s is a paid mutator transaction binding the contract method 0x61fe787d.
//
// Solidity: function keccak256s(uint256 n) returns()
func (_Hash *HashTransactor) Keccak256s(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Hash.contract.Transact(opts, "keccak256s", n)
}

// Keccak256s is a paid mutator transaction binding the contract method 0x61fe787d.
//
// Solidity: function keccak256s(uint256 n) returns()
func (_Hash *HashSession) Keccak256s(n *big.Int) (*types.Transaction, error) {
	return _Hash.Contract.Keccak256s(&_Hash.TransactOpts, n)
}

// Keccak256s is a paid mutator transaction binding the contract method 0x61fe787d.
//
// Solidity: function keccak256s(uint256 n) returns()
func (_Hash *HashTransactorSession) Keccak256s(n *big.Int) (*types.Transaction, error) {
	return _Hash.Contract.Keccak256s(&_Hash.TransactOpts, n)
}

// Sha256s is a paid mutator transaction binding the contract method 0x61ef0ccb.
//
// Solidity: function sha256s(uint256 n) returns()
func (_Hash *HashTransactor) Sha256s(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Hash.contract.Transact(opts, "sha256s", n)
}

// Sha256s is a paid mutator transaction binding the contract method 0x61ef0ccb.
//
// Solidity: function sha256s(uint256 n) returns()
func (_Hash *HashSession) Sha256s(n *big.Int) (*types.Transaction, error) {
	return _Hash.Contract.Sha256s(&_Hash.TransactOpts, n)
}

// Sha256s is a paid mutator transaction binding the contract method 0x61ef0ccb.
//
// Solidity: function sha256s(uint256 n) returns()
func (_Hash *HashTransactorSession) Sha256s(n *big.Int) (*types.Transaction, error) {
	return _Hash.Contract.Sha256s(&_Hash.TransactOpts, n)
}
