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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sha256\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sha256Yul\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"sha256s\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bebc76dd": "sha256(bytes)",
		"fe0ed4c1": "sha256Yul(bytes)",
		"61ef0ccb": "sha256s(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610330806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806361ef0ccb14610046578063bebc76dd1461005b578063fe0ed4c114610084575b600080fd5b610059610054366004610174565b610097565b005b61006e6100693660046101a3565b6100e4565b60405161007b9190610284565b60405180910390f35b61006e6100923660046101a3565b610158565b604080516103e7602082015260009101604051602081830303815290604052905060005b828110156100df576100cc826100e4565b50806100d7816102b7565b9150506100bb565b505050565b606060008060026001600160a01b03168460405161010291906102de565b600060405180830381855afa9150503d806000811461013d576040519150601f19603f3d011682016040523d82523d6000602084013e610142565b606091505b50915091508161015157600080fd5b5050919050565b6060602060006020600060025afa61016f57600080fd5b919050565b60006020828403121561018657600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156101b557600080fd5b813567ffffffffffffffff808211156101cd57600080fd5b818401915084601f8301126101e157600080fd5b8135818111156101f3576101f361018d565b604051601f8201601f19908116603f0116810190838211818310171561021b5761021b61018d565b8160405282815287602084870101111561023457600080fd5b826020860160208301376000928101602001929092525095945050505050565b60005b8381101561026f578181015183820152602001610257565b8381111561027e576000848401525b50505050565b60208152600082518060208401526102a3816040850160208701610254565b601f01601f19169190910160400192915050565b6000600182016102d757634e487b7160e01b600052601160045260246000fd5b5060010190565b600082516102f0818460208701610254565b919091019291505056fea26469706673582212202e64a09e1bd8cec6c549ffd283275ef8c522d55429699f3a71a72e1eb4b76ebb64736f6c634300080f0033",
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
