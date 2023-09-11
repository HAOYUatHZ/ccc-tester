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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p0\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p1\",\"type\":\"uint256[2]\"}],\"name\":\"ecAdd\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecAdds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ecMul\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecMuls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecPairings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecRecovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a867d0ca": "ecAdd(uint256[2],uint256[2])",
		"7de7133d": "ecAdds(uint256)",
		"71de3724": "ecMul(uint256[2],uint256)",
		"f602cf71": "ecMuls(uint256)",
		"459be292": "ecPairings(uint256)",
		"112355a4": "ecRecovers(uint256)",
	},
	Bin: "0x608060405234801561000f575f80fd5b50610c388061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c8063112355a414610064578063459be2921461007957806371de37241461008c5780637de7133d146100b5578063a867d0ca146100c8578063f602cf71146100db575b5f80fd5b610077610072366004610a68565b6100ee565b005b610077610087366004610a68565b6101cc565b61009f61009a366004610b0c565b6105a8565b6040516100ac9190610b35565b60405180910390f35b6100776100c3366004610a68565b610615565b61009f6100d6366004610b65565b610662565b6100776100e9366004610a68565b6106fa565b7fb94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde97fe742ff452d41413616a5bf43fe15dd88294e983d3d36206c2712f39083d638bd601b7fe0a0fc89be718fbc1033e1d30d78be1c68081562ed2e97af876f286f3453231d5f5b858110156101c457604080515f81526020810180835287905260ff851691810191909152606081018590526080810183905260019060a0016020604051602081039080840390855afa1580156101ad573d5f803e3d5ffd5b508291506101bc905081610bac565b915050610155565b505050505050565b604080516002808252606082019092525f91816020015b604080518082019091525f80825260208201528152602001906001900390816101e3575050604080516002808252606082019092529192505f9190602082015b61022b610a07565b8152602001906001900390816102235790505090506001825f8151811061025457610254610bc4565b60200260200101515f0181815250506002825f8151811061027757610277610bc4565b602002602001015160200181815250507f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed815f815181106102ba576102ba610bc4565b60209081029190910181015151015280517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c29082905f906102fd576102fd610bc4565b6020908102919091010151515280517f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa9082905f9061033e5761033e610bc4565b60200260200101516020015160016002811061035c5761035c610bc4565b602002015280517f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b9082905f9061039557610395610bc4565b6020026020010151602001515f600281106103b2576103b2610bc4565b602002015281517f1aa125a22bd902874034e67868aed40267e5575d5919677987e3bc6dd42a32fe90839060019081106103ee576103ee610bc4565b60200260200101515f0181815250507f1bacc186725464068956d9a191455c2d6f6db282d83645c610510d8d4efbaee08260018151811061043157610431610bc4565b602002602001015160200181815250507f1b7734c80605f71f1e2de61e998ce5854ff2abebb76537c3d67e50d71422a8528160018151811061047557610475610bc4565b60209081029190910181015151015280517f10d5a1e34b2388a5ebe266033a5e0e63c89084203784da0c6bd9b052a78a2cac90829060019081106104bb576104bb610bc4565b6020908102919091010151515280517f275739c5c2cdbc72e37c689e2ab441ea76c1d284b9c46ae8f5c42ead937819e190829060019081106104ff576104ff610bc4565b60200260200101516020015160016002811061051d5761051d610bc4565b602002015280517f018de34c5b7c3d3d75428bbe050f1449ea3d9961d563291f307a1874f7332e65908290600190811061055957610559610bc4565b6020026020010151602001515f6002811061057657610576610bc4565b60200201525f5b838110156105a25761058f838361073a565b508061059a81610bac565b91505061057d565b50505050565b6105b0610a2c565b5f6040518060600160405280855f600281106105ce576105ce610bc4565b60200201518152602001856001600281106105eb576105eb610bc4565b6020020151815260200184815250905060408260608360075f19fa61060e575f80fd5b5092915050565b61061d610a2c565b6001815260026020820152610630610a2c565b60018152600260208201525f5b838110156105a25761064f8383610662565b508061065a81610bac565b91505061063d565b61066a610a2c565b5f6040518060800160405280855f6002811061068857610688610bc4565b60200201518152602001856001600281106106a5576106a5610bc4565b60200201518152602001845f600281106106c1576106c1610bc4565b60200201518152602001846001600281106106de576106de610bc4565b60200201519052905060408260808360065f19fa61060e575f80fd5b610702610a2c565b60018152600260208201525f5b82811015610735576107228260036105a8565b508061072d81610bac565b91505061070f565b505050565b5f808351600661074a9190610bd8565b90505f8167ffffffffffffffff81111561076657610766610a7f565b60405190808252806020026020018201604052801561078f578160200160208202803683370190505b50905061079a610a4a565b5f85518751146107a8575f80fd5b5f5b87518110156109dd578781815181106107c5576107c5610bc4565b60200260200101515f0151848260066107de9190610bd8565b6107e8905f610bef565b815181106107f8576107f8610bc4565b60200260200101818152505087818151811061081657610816610bc4565b602002602001015160200151848260066108309190610bd8565b61083b906001610bef565b8151811061084b5761084b610bc4565b60200260200101818152505086818151811061086957610869610bc4565b6020908102919091010151515184610882836006610bd8565b61088d906002610bef565b8151811061089d5761089d610bc4565b6020026020010181815250508681815181106108bb576108bb610bc4565b602090810291909101810151510151846108d6836006610bd8565b6108e1906003610bef565b815181106108f1576108f1610bc4565b60200260200101818152505086818151811061090f5761090f610bc4565b6020026020010151602001515f6002811061092c5761092c610bc4565b60200201518461093d836006610bd8565b610948906004610bef565b8151811061095857610958610bc4565b60200260200101818152505086818151811061097657610976610bc4565b60200260200101516020015160016002811061099457610994610bc4565b6020020151846109a5836006610bd8565b6109b0906005610bef565b815181106109c0576109c0610bc4565b6020908102919091010152806109d581610bac565b9150506107aa565b50602082602086026020860160085afa9050806109f8575f80fd5b50511515925050505b92915050565b6040518060400160405280610a1a610a2c565b8152602001610a27610a2c565b905290565b60405180604001604052806002906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5f60208284031215610a78575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610aa2575f80fd5b6040516040810181811067ffffffffffffffff82111715610ad157634e487b7160e01b5f52604160045260245ffd5b8060405250806040840185811115610ae7575f80fd5b845b81811015610b01578035835260209283019201610ae9565b509195945050505050565b5f8060608385031215610b1d575f80fd5b610b278484610a93565b946040939093013593505050565b6040810181835f5b6002811015610b5c578151835260209283019290910190600101610b3d565b50505092915050565b5f8060808385031215610b76575f80fd5b610b808484610a93565b9150610b8f8460408501610a93565b90509250929050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201610bbd57610bbd610b98565b5060010190565b634e487b7160e01b5f52603260045260245ffd5b8082028115828204841417610a0157610a01610b98565b80820180821115610a0157610a01610b9856fea26469706673582212203bf0e5c560341bf6728155b4e6a4ee65fb0d4d1b32c0a4d8477f8cf70f03944264736f6c63430008150033",
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

// EcRecovers is a paid mutator transaction binding the contract method 0x112355a4.
//
// Solidity: function ecRecovers(uint256 n) returns()
func (_Ecc *EccTransactor) EcRecovers(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecRecovers", n)
}

// EcRecovers is a paid mutator transaction binding the contract method 0x112355a4.
//
// Solidity: function ecRecovers(uint256 n) returns()
func (_Ecc *EccSession) EcRecovers(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcRecovers(&_Ecc.TransactOpts, n)
}

// EcRecovers is a paid mutator transaction binding the contract method 0x112355a4.
//
// Solidity: function ecRecovers(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcRecovers(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcRecovers(&_Ecc.TransactOpts, n)
}
