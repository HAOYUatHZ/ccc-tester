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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p0\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p1\",\"type\":\"uint256[2]\"}],\"name\":\"ecAdd\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecAdds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ecMul\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecMuls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecPairings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a867d0ca": "ecAdd(uint256[2],uint256[2])",
		"7de7133d": "ecAdds(uint256)",
		"71de3724": "ecMul(uint256[2],uint256)",
		"f602cf71": "ecMuls(uint256)",
		"459be292": "ecPairings(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610b7f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063459be2921461005c57806371de3724146100715780637de7133d1461009a578063a867d0ca146100ad578063f602cf71146100c0575b600080fd5b61006f61006a36600461098f565b6100d3565b005b61008461007f366004610a3c565b6104bd565b6040516100919190610a67565b60405180910390f35b61006f6100a836600461098f565b61052e565b6100846100bb366004610a98565b61057c565b61006f6100ce36600461098f565b610619565b60408051600280825260608201909252600091816020015b60408051808201909152600080825260208201528152602001906001900390816100eb5750506040805160028082526060820190925291925060009190602082015b61013561092e565b81526020019060019003908161012d57905050905060018260008151811061015f5761015f610acd565b6020026020010151600001818152505060028260008151811061018457610184610acd565b602002602001015160200181815250507f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed816000815181106101c8576101c8610acd565b60209081029190910181015151015280517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c290829060009061020c5761020c610acd565b6020908102919091010151515280517f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa90829060009061024e5761024e610acd565b60200260200101516020015160016002811061026c5761026c610acd565b602002015280517f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b9082906000906102a6576102a6610acd565b6020026020010151602001516000600281106102c4576102c4610acd565b602002015281517f1aa125a22bd902874034e67868aed40267e5575d5919677987e3bc6dd42a32fe908390600190811061030057610300610acd565b602002602001015160000181815250507f1bacc186725464068956d9a191455c2d6f6db282d83645c610510d8d4efbaee08260018151811061034457610344610acd565b602002602001015160200181815250507f1b7734c80605f71f1e2de61e998ce5854ff2abebb76537c3d67e50d71422a8528160018151811061038857610388610acd565b60209081029190910181015151015280517f10d5a1e34b2388a5ebe266033a5e0e63c89084203784da0c6bd9b052a78a2cac90829060019081106103ce576103ce610acd565b6020908102919091010151515280517f275739c5c2cdbc72e37c689e2ab441ea76c1d284b9c46ae8f5c42ead937819e1908290600190811061041257610412610acd565b60200260200101516020015160016002811061043057610430610acd565b602002015280517f018de34c5b7c3d3d75428bbe050f1449ea3d9961d563291f307a1874f7332e65908290600190811061046c5761046c610acd565b60200260200101516020015160006002811061048a5761048a610acd565b602002015260005b838110156104b7576104a4838361065a565b50806104af81610af9565b915050610492565b50505050565b6104c5610953565b60006040518060600160405280856000600281106104e5576104e5610acd565b602002015181526020018560016002811061050257610502610acd565b602002015181526020018481525090506040826060836007600019fa61052757600080fd5b5092915050565b610536610953565b6001815260026020820152610549610953565b600181526002602082015260005b838110156104b757610569838361057c565b508061057481610af9565b915050610557565b610584610953565b60006040518060800160405280856000600281106105a4576105a4610acd565b60200201518152602001856001600281106105c1576105c1610acd565b60200201518152602001846000600281106105de576105de610acd565b60200201518152602001846001600281106105fb576105fb610acd565b6020020151905290506040826080836006600019fa61052757600080fd5b610621610953565b600181526002602082015260005b82811015610655576106428260036104bd565b508061064d81610af9565b91505061062f565b505050565b6000808351600661066b9190610b12565b905060008167ffffffffffffffff811115610688576106886109a8565b6040519080825280602002602001820160405280156106b1578160200160208202803683370190505b5090506106bc610971565b600085518751146106cc57600080fd5b60005b8751811015610905578781815181106106ea576106ea610acd565b602002602001015160000151848260066107049190610b12565b61070f906000610b31565b8151811061071f5761071f610acd565b60200260200101818152505087818151811061073d5761073d610acd565b602002602001015160200151848260066107579190610b12565b610762906001610b31565b8151811061077257610772610acd565b60200260200101818152505086818151811061079057610790610acd565b60209081029190910101515151846107a9836006610b12565b6107b4906002610b31565b815181106107c4576107c4610acd565b6020026020010181815250508681815181106107e2576107e2610acd565b602090810291909101810151510151846107fd836006610b12565b610808906003610b31565b8151811061081857610818610acd565b60200260200101818152505086818151811061083657610836610acd565b60200260200101516020015160006002811061085457610854610acd565b602002015184610865836006610b12565b610870906004610b31565b8151811061088057610880610acd565b60200260200101818152505086818151811061089e5761089e610acd565b6020026020010151602001516001600281106108bc576108bc610acd565b6020020151846108cd836006610b12565b6108d8906005610b31565b815181106108e8576108e8610acd565b6020908102919091010152806108fd81610af9565b9150506106cf565b50602082602086026020860160085afa90508061092157600080fd5b5051151595945050505050565b6040518060400160405280610941610953565b815260200161094e610953565b905290565b60405180604001604052806002906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6000602082840312156109a157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126109cf57600080fd5b6040516040810181811067ffffffffffffffff82111715610a0057634e487b7160e01b600052604160045260246000fd5b8060405250806040840185811115610a1757600080fd5b845b81811015610a31578035835260209283019201610a19565b509195945050505050565b60008060608385031215610a4f57600080fd5b610a5984846109be565b946040939093013593505050565b60408101818360005b6002811015610a8f578151835260209283019290910190600101610a70565b50505092915050565b60008060808385031215610aab57600080fd5b610ab584846109be565b9150610ac484604085016109be565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610b0b57610b0b610ae3565b5060010190565b6000816000190483118215151615610b2c57610b2c610ae3565b500290565b60008219821115610b4457610b44610ae3565b50019056fea26469706673582212207f56d5c0565a17e07acd844dccaed5e13fb2b1da197fcaf6c979393f2cad889264736f6c634300080f0033",
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

// EcPairings is a paid mutator transaction binding the contract method 0x459be292.
//
// Solidity: function ecPairings(uint256 n) returns()
func (_Ecc *EccTransactor) EcPairings(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecPairings", n)
}

// EcPairings is a paid mutator transaction binding the contract method 0x459be292.
//
// Solidity: function ecPairings(uint256 n) returns()
func (_Ecc *EccSession) EcPairings(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcPairings(&_Ecc.TransactOpts, n)
}

// EcPairings is a paid mutator transaction binding the contract method 0x459be292.
//
// Solidity: function ecPairings(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcPairings(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcPairings(&_Ecc.TransactOpts, n)
}
