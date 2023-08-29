// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sha256

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

// Sha256MetaData contains all meta data concerning the Sha256 contract.
var Sha256MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"sha256\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"out\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"sha256s\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bebc76dd": "sha256(bytes)",
		"61ef0ccb": "sha256s(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506102f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806361ef0ccb1461003b578063bebc76dd14610050575b600080fd5b61004e61004936600461013a565b610079565b005b61006361005e366004610169565b6100c6565b604051610070919061024a565b60405180910390f35b604080516103e7602082015260009101604051602081830303815290604052905060005b828110156100c1576100ae826100c6565b50806100b98161027d565b91505061009d565b505050565b606060008060026001600160a01b0316846040516100e491906102a4565b600060405180830381855afa9150503d806000811461011f576040519150601f19603f3d011682016040523d82523d6000602084013e610124565b606091505b50915091508161013357600080fd5b5050919050565b60006020828403121561014c57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561017b57600080fd5b813567ffffffffffffffff8082111561019357600080fd5b818401915084601f8301126101a757600080fd5b8135818111156101b9576101b9610153565b604051601f8201601f19908116603f011681019083821181831017156101e1576101e1610153565b816040528281528760208487010111156101fa57600080fd5b826020860160208301376000928101602001929092525095945050505050565b60005b8381101561023557818101518382015260200161021d565b83811115610244576000848401525b50505050565b602081526000825180602084015261026981604085016020870161021a565b601f01601f19169190910160400192915050565b60006001820161029d57634e487b7160e01b600052601160045260246000fd5b5060010190565b600082516102b681846020870161021a565b919091019291505056fea2646970667358221220758c8bf3f286aaad54172d182f54f138a9fe62c2b6c697b75f78b57b24ff3db764736f6c634300080f0033",
}

// Sha256ABI is the input ABI used to generate the binding from.
// Deprecated: Use Sha256MetaData.ABI instead.
var Sha256ABI = Sha256MetaData.ABI

// Deprecated: Use Sha256MetaData.Sigs instead.
// Sha256FuncSigs maps the 4-byte function signature to its string representation.
var Sha256FuncSigs = Sha256MetaData.Sigs

// Sha256Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Sha256MetaData.Bin instead.
var Sha256Bin = Sha256MetaData.Bin

// DeploySha256 deploys a new Ethereum contract, binding an instance of Sha256 to it.
func DeploySha256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sha256, error) {
	parsed, err := Sha256MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Sha256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sha256{Sha256Caller: Sha256Caller{contract: contract}, Sha256Transactor: Sha256Transactor{contract: contract}, Sha256Filterer: Sha256Filterer{contract: contract}}, nil
}

// Sha256 is an auto generated Go binding around an Ethereum contract.
type Sha256 struct {
	Sha256Caller     // Read-only binding to the contract
	Sha256Transactor // Write-only binding to the contract
	Sha256Filterer   // Log filterer for contract events
}

// Sha256Caller is an auto generated read-only Go binding around an Ethereum contract.
type Sha256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sha256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Sha256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sha256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sha256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sha256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sha256Session struct {
	Contract     *Sha256           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Sha256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sha256CallerSession struct {
	Contract *Sha256Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Sha256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sha256TransactorSession struct {
	Contract     *Sha256Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Sha256Raw is an auto generated low-level Go binding around an Ethereum contract.
type Sha256Raw struct {
	Contract *Sha256 // Generic contract binding to access the raw methods on
}

// Sha256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sha256CallerRaw struct {
	Contract *Sha256Caller // Generic read-only contract binding to access the raw methods on
}

// Sha256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sha256TransactorRaw struct {
	Contract *Sha256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSha256 creates a new instance of Sha256, bound to a specific deployed contract.
func NewSha256(address common.Address, backend bind.ContractBackend) (*Sha256, error) {
	contract, err := bindSha256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sha256{Sha256Caller: Sha256Caller{contract: contract}, Sha256Transactor: Sha256Transactor{contract: contract}, Sha256Filterer: Sha256Filterer{contract: contract}}, nil
}

// NewSha256Caller creates a new read-only instance of Sha256, bound to a specific deployed contract.
func NewSha256Caller(address common.Address, caller bind.ContractCaller) (*Sha256Caller, error) {
	contract, err := bindSha256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sha256Caller{contract: contract}, nil
}

// NewSha256Transactor creates a new write-only instance of Sha256, bound to a specific deployed contract.
func NewSha256Transactor(address common.Address, transactor bind.ContractTransactor) (*Sha256Transactor, error) {
	contract, err := bindSha256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sha256Transactor{contract: contract}, nil
}

// NewSha256Filterer creates a new log filterer instance of Sha256, bound to a specific deployed contract.
func NewSha256Filterer(address common.Address, filterer bind.ContractFilterer) (*Sha256Filterer, error) {
	contract, err := bindSha256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sha256Filterer{contract: contract}, nil
}

// bindSha256 binds a generic wrapper to an already deployed contract.
func bindSha256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Sha256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sha256 *Sha256Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sha256.Contract.Sha256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sha256 *Sha256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sha256.Contract.Sha256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sha256 *Sha256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sha256.Contract.Sha256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sha256 *Sha256CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sha256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sha256 *Sha256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sha256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sha256 *Sha256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sha256.Contract.contract.Transact(opts, method, params...)
}

// Sha256 is a free data retrieval call binding the contract method 0xbebc76dd.
//
// Solidity: function sha256(bytes input) view returns(bytes out)
func (_Sha256 *Sha256Caller) Sha256(opts *bind.CallOpts, input []byte) ([]byte, error) {
	var out []interface{}
	err := _Sha256.contract.Call(opts, &out, "sha256", input)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Sha256 is a free data retrieval call binding the contract method 0xbebc76dd.
//
// Solidity: function sha256(bytes input) view returns(bytes out)
func (_Sha256 *Sha256Session) Sha256(input []byte) ([]byte, error) {
	return _Sha256.Contract.Sha256(&_Sha256.CallOpts, input)
}

// Sha256 is a free data retrieval call binding the contract method 0xbebc76dd.
//
// Solidity: function sha256(bytes input) view returns(bytes out)
func (_Sha256 *Sha256CallerSession) Sha256(input []byte) ([]byte, error) {
	return _Sha256.Contract.Sha256(&_Sha256.CallOpts, input)
}

// Sha256s is a paid mutator transaction binding the contract method 0x61ef0ccb.
//
// Solidity: function sha256s(uint256 n) returns()
func (_Sha256 *Sha256Transactor) Sha256s(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Sha256.contract.Transact(opts, "sha256s", n)
}

// Sha256s is a paid mutator transaction binding the contract method 0x61ef0ccb.
//
// Solidity: function sha256s(uint256 n) returns()
func (_Sha256 *Sha256Session) Sha256s(n *big.Int) (*types.Transaction, error) {
	return _Sha256.Contract.Sha256s(&_Sha256.TransactOpts, n)
}

// Sha256s is a paid mutator transaction binding the contract method 0x61ef0ccb.
//
// Solidity: function sha256s(uint256 n) returns()
func (_Sha256 *Sha256TransactorSession) Sha256s(n *big.Int) (*types.Transaction, error) {
	return _Sha256.Contract.Sha256s(&_Sha256.TransactOpts, n)
}
