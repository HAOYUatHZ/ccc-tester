// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modexp

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

// ModExpMetaData contains all meta data concerning the ModExp contract.
var ModExpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"modexps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7eec8dda": "modexps(uint256)",
	},
	Bin: "0x608060405234801561000f575f80fd5b5061015d8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80637eec8dda1461002d575b5f80fd5b61004061003b3660046100ec565b610042565b005b60635f5b8281101561006a576100578261006f565b508061006281610103565b915050610046565b505050565b5f60405160208152602080820152602060408201528260608201527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808201527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760a082015260208160c08360056107d05a03fa505192915050565b5f602082840312156100fc575f80fd5b5035919050565b5f6001820161012057634e487b7160e01b5f52601160045260245ffd5b506001019056fea2646970667358221220f7e4718f9112196e18133a4cab60616a524d5580d26808cc58d8843538e27b3164736f6c63430008150033",
}

// ModExpABI is the input ABI used to generate the binding from.
// Deprecated: Use ModExpMetaData.ABI instead.
var ModExpABI = ModExpMetaData.ABI

// Deprecated: Use ModExpMetaData.Sigs instead.
// ModExpFuncSigs maps the 4-byte function signature to its string representation.
var ModExpFuncSigs = ModExpMetaData.Sigs

// ModExpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModExpMetaData.Bin instead.
var ModExpBin = ModExpMetaData.Bin

// DeployModExp deploys a new Ethereum contract, binding an instance of ModExp to it.
func DeployModExp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModExp, error) {
	parsed, err := ModExpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModExpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModExp{ModExpCaller: ModExpCaller{contract: contract}, ModExpTransactor: ModExpTransactor{contract: contract}, ModExpFilterer: ModExpFilterer{contract: contract}}, nil
}

// ModExp is an auto generated Go binding around an Ethereum contract.
type ModExp struct {
	ModExpCaller     // Read-only binding to the contract
	ModExpTransactor // Write-only binding to the contract
	ModExpFilterer   // Log filterer for contract events
}

// ModExpCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModExpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModExpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModExpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModExpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModExpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModExpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModExpSession struct {
	Contract     *ModExp           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModExpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModExpCallerSession struct {
	Contract *ModExpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModExpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModExpTransactorSession struct {
	Contract     *ModExpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModExpRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModExpRaw struct {
	Contract *ModExp // Generic contract binding to access the raw methods on
}

// ModExpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModExpCallerRaw struct {
	Contract *ModExpCaller // Generic read-only contract binding to access the raw methods on
}

// ModExpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModExpTransactorRaw struct {
	Contract *ModExpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModExp creates a new instance of ModExp, bound to a specific deployed contract.
func NewModExp(address common.Address, backend bind.ContractBackend) (*ModExp, error) {
	contract, err := bindModExp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModExp{ModExpCaller: ModExpCaller{contract: contract}, ModExpTransactor: ModExpTransactor{contract: contract}, ModExpFilterer: ModExpFilterer{contract: contract}}, nil
}

// NewModExpCaller creates a new read-only instance of ModExp, bound to a specific deployed contract.
func NewModExpCaller(address common.Address, caller bind.ContractCaller) (*ModExpCaller, error) {
	contract, err := bindModExp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModExpCaller{contract: contract}, nil
}

// NewModExpTransactor creates a new write-only instance of ModExp, bound to a specific deployed contract.
func NewModExpTransactor(address common.Address, transactor bind.ContractTransactor) (*ModExpTransactor, error) {
	contract, err := bindModExp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModExpTransactor{contract: contract}, nil
}

// NewModExpFilterer creates a new log filterer instance of ModExp, bound to a specific deployed contract.
func NewModExpFilterer(address common.Address, filterer bind.ContractFilterer) (*ModExpFilterer, error) {
	contract, err := bindModExp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModExpFilterer{contract: contract}, nil
}

// bindModExp binds a generic wrapper to an already deployed contract.
func bindModExp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModExpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModExp *ModExpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModExp.Contract.ModExpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModExp *ModExpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModExp.Contract.ModExpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModExp *ModExpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModExp.Contract.ModExpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModExp *ModExpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModExp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModExp *ModExpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModExp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModExp *ModExpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModExp.Contract.contract.Transact(opts, method, params...)
}

// Modexps is a paid mutator transaction binding the contract method 0x7eec8dda.
//
// Solidity: function modexps(uint256 n) returns()
func (_ModExp *ModExpTransactor) Modexps(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _ModExp.contract.Transact(opts, "modexps", n)
}

// Modexps is a paid mutator transaction binding the contract method 0x7eec8dda.
//
// Solidity: function modexps(uint256 n) returns()
func (_ModExp *ModExpSession) Modexps(n *big.Int) (*types.Transaction, error) {
	return _ModExp.Contract.Modexps(&_ModExp.TransactOpts, n)
}

// Modexps is a paid mutator transaction binding the contract method 0x7eec8dda.
//
// Solidity: function modexps(uint256 n) returns()
func (_ModExp *ModExpTransactorSession) Modexps(n *big.Int) (*types.Transaction, error) {
	return _ModExp.Contract.Modexps(&_ModExp.TransactOpts, n)
}
