// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package credit

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

// Account is an auto generated low-level Go binding around an user-defined struct.
type Account struct {
	CapacityUsed    *big.Int
	CreditFree      *big.Int
	CreditCommitted *big.Int
	LastDebitEpoch  uint64
	Approvals       []Approvals
}

// Approval is an auto generated low-level Go binding around an user-defined struct.
type Approval struct {
	RequiredCaller common.Address
	Approval       CreditApproval
}

// Approvals is an auto generated low-level Go binding around an user-defined struct.
type Approvals struct {
	Receiver common.Address
	Approval []Approval
}

// Balance is an auto generated low-level Go binding around an user-defined struct.
type Balance struct {
	CreditFree      *big.Int
	CreditCommitted *big.Int
	LastDebitEpoch  uint64
}

// CreditApproval is an auto generated low-level Go binding around an user-defined struct.
type CreditApproval struct {
	Limit     *big.Int
	Expiry    uint64
	Committed *big.Int
}

// CreditStats is an auto generated low-level Go binding around an user-defined struct.
type CreditStats struct {
	Balance         *big.Int
	CreditSold      *big.Int
	CreditCommitted *big.Int
	CreditDebited   *big.Int
	CreditDebitRate uint64
	NumAccounts     uint64
}

// StorageStats is an auto generated low-level Go binding around an user-defined struct.
type StorageStats struct {
	CapacityFree *big.Int
	CapacityUsed *big.Int
	NumBlobs     uint64
	NumResolving uint64
}

// SubnetStats is an auto generated low-level Go binding around an user-defined struct.
type SubnetStats struct {
	Balance         *big.Int
	CapacityFree    *big.Int
	CapacityUsed    *big.Int
	CreditSold      *big.Int
	CreditCommitted *big.Int
	CreditDebited   *big.Int
	CreditDebitRate uint64
	NumAccounts     uint64
	NumBlobs        uint64
	NumResolving    uint64
}

// Usage is an auto generated low-level Go binding around an user-defined struct.
type Usage struct {
	CapacityUsed *big.Int
}

// CreditMetaData contains all meta data concerning the Credit contract.
var CreditMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"approveCredit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCaller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCaller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCaller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buyCredit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"buyCredit\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getAccount\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"tuple\",\"internalType\":\"structAccount\",\"components\":[{\"name\":\"capacityUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditFree\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditCommitted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastDebitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"approvals\",\"type\":\"tuple[]\",\"internalType\":\"structApprovals[]\",\"components\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approval\",\"type\":\"tuple[]\",\"internalType\":\"structApproval[]\",\"components\":[{\"name\":\"requiredCaller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approval\",\"type\":\"tuple\",\"internalType\":\"structCreditApproval\",\"components\":[{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"committed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditBalance\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"tuple\",\"internalType\":\"structBalance\",\"components\":[{\"name\":\"creditFree\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditCommitted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastDebitEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditStats\",\"inputs\":[],\"outputs\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structCreditStats\",\"components\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditSold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditCommitted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditDebited\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditDebitRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageStats\",\"inputs\":[],\"outputs\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structStorageStats\",\"components\":[{\"name\":\"capacityFree\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numBlobs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numResolving\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageUsage\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"usage\",\"type\":\"tuple\",\"internalType\":\"structUsage\",\"components\":[{\"name\":\"capacityUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubnetStats\",\"inputs\":[],\"outputs\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structSubnetStats\",\"components\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityFree\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditSold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditCommitted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditDebited\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditDebitRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numBlobs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numResolving\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeCredit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCaller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApproveCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requiredCaller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"limit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BuyCredit\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevokeCredit\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requiredCaller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// CreditABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditMetaData.ABI instead.
var CreditABI = CreditMetaData.ABI

// Credit is an auto generated Go binding around an Ethereum contract.
type Credit struct {
	CreditCaller     // Read-only binding to the contract
	CreditTransactor // Write-only binding to the contract
	CreditFilterer   // Log filterer for contract events
}

// CreditCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditSession struct {
	Contract     *Credit           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditCallerSession struct {
	Contract *CreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CreditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditTransactorSession struct {
	Contract     *CreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditRaw struct {
	Contract *Credit // Generic contract binding to access the raw methods on
}

// CreditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditCallerRaw struct {
	Contract *CreditCaller // Generic read-only contract binding to access the raw methods on
}

// CreditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditTransactorRaw struct {
	Contract *CreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredit creates a new instance of Credit, bound to a specific deployed contract.
func NewCredit(address common.Address, backend bind.ContractBackend) (*Credit, error) {
	contract, err := bindCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Credit{CreditCaller: CreditCaller{contract: contract}, CreditTransactor: CreditTransactor{contract: contract}, CreditFilterer: CreditFilterer{contract: contract}}, nil
}

// NewCreditCaller creates a new read-only instance of Credit, bound to a specific deployed contract.
func NewCreditCaller(address common.Address, caller bind.ContractCaller) (*CreditCaller, error) {
	contract, err := bindCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditCaller{contract: contract}, nil
}

// NewCreditTransactor creates a new write-only instance of Credit, bound to a specific deployed contract.
func NewCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditTransactor, error) {
	contract, err := bindCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditTransactor{contract: contract}, nil
}

// NewCreditFilterer creates a new log filterer instance of Credit, bound to a specific deployed contract.
func NewCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFilterer, error) {
	contract, err := bindCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFilterer{contract: contract}, nil
}

// bindCredit binds a generic wrapper to an already deployed contract.
func bindCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credit *CreditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credit.Contract.CreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credit *CreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.Contract.CreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credit *CreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credit.Contract.CreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credit *CreditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credit *CreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credit *CreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credit.Contract.contract.Transact(opts, method, params...)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address addr) view returns((uint256,uint256,uint256,uint64,(address,(address,(uint256,uint64,uint256))[])[]) account)
func (_Credit *CreditCaller) GetAccount(opts *bind.CallOpts, addr common.Address) (Account, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "getAccount", addr)

	if err != nil {
		return *new(Account), err
	}

	out0 := *abi.ConvertType(out[0], new(Account)).(*Account)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address addr) view returns((uint256,uint256,uint256,uint64,(address,(address,(uint256,uint64,uint256))[])[]) account)
func (_Credit *CreditSession) GetAccount(addr common.Address) (Account, error) {
	return _Credit.Contract.GetAccount(&_Credit.CallOpts, addr)
}

// GetAccount is a free data retrieval call binding the contract method 0xfbcbc0f1.
//
// Solidity: function getAccount(address addr) view returns((uint256,uint256,uint256,uint64,(address,(address,(uint256,uint64,uint256))[])[]) account)
func (_Credit *CreditCallerSession) GetAccount(addr common.Address) (Account, error) {
	return _Credit.Contract.GetAccount(&_Credit.CallOpts, addr)
}

// GetCreditBalance is a free data retrieval call binding the contract method 0x6add1b3f.
//
// Solidity: function getCreditBalance(address addr) view returns((uint256,uint256,uint64) balance)
func (_Credit *CreditCaller) GetCreditBalance(opts *bind.CallOpts, addr common.Address) (Balance, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "getCreditBalance", addr)

	if err != nil {
		return *new(Balance), err
	}

	out0 := *abi.ConvertType(out[0], new(Balance)).(*Balance)

	return out0, err

}

// GetCreditBalance is a free data retrieval call binding the contract method 0x6add1b3f.
//
// Solidity: function getCreditBalance(address addr) view returns((uint256,uint256,uint64) balance)
func (_Credit *CreditSession) GetCreditBalance(addr common.Address) (Balance, error) {
	return _Credit.Contract.GetCreditBalance(&_Credit.CallOpts, addr)
}

// GetCreditBalance is a free data retrieval call binding the contract method 0x6add1b3f.
//
// Solidity: function getCreditBalance(address addr) view returns((uint256,uint256,uint64) balance)
func (_Credit *CreditCallerSession) GetCreditBalance(addr common.Address) (Balance, error) {
	return _Credit.Contract.GetCreditBalance(&_Credit.CallOpts, addr)
}

// GetCreditStats is a free data retrieval call binding the contract method 0xc7d0b56c.
//
// Solidity: function getCreditStats() view returns((uint256,uint256,uint256,uint256,uint64,uint64) stats)
func (_Credit *CreditCaller) GetCreditStats(opts *bind.CallOpts) (CreditStats, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "getCreditStats")

	if err != nil {
		return *new(CreditStats), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditStats)).(*CreditStats)

	return out0, err

}

// GetCreditStats is a free data retrieval call binding the contract method 0xc7d0b56c.
//
// Solidity: function getCreditStats() view returns((uint256,uint256,uint256,uint256,uint64,uint64) stats)
func (_Credit *CreditSession) GetCreditStats() (CreditStats, error) {
	return _Credit.Contract.GetCreditStats(&_Credit.CallOpts)
}

// GetCreditStats is a free data retrieval call binding the contract method 0xc7d0b56c.
//
// Solidity: function getCreditStats() view returns((uint256,uint256,uint256,uint256,uint64,uint64) stats)
func (_Credit *CreditCallerSession) GetCreditStats() (CreditStats, error) {
	return _Credit.Contract.GetCreditStats(&_Credit.CallOpts)
}

// GetStorageStats is a free data retrieval call binding the contract method 0x537b633a.
//
// Solidity: function getStorageStats() view returns((uint256,uint256,uint64,uint64) stats)
func (_Credit *CreditCaller) GetStorageStats(opts *bind.CallOpts) (StorageStats, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "getStorageStats")

	if err != nil {
		return *new(StorageStats), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageStats)).(*StorageStats)

	return out0, err

}

// GetStorageStats is a free data retrieval call binding the contract method 0x537b633a.
//
// Solidity: function getStorageStats() view returns((uint256,uint256,uint64,uint64) stats)
func (_Credit *CreditSession) GetStorageStats() (StorageStats, error) {
	return _Credit.Contract.GetStorageStats(&_Credit.CallOpts)
}

// GetStorageStats is a free data retrieval call binding the contract method 0x537b633a.
//
// Solidity: function getStorageStats() view returns((uint256,uint256,uint64,uint64) stats)
func (_Credit *CreditCallerSession) GetStorageStats() (StorageStats, error) {
	return _Credit.Contract.GetStorageStats(&_Credit.CallOpts)
}

// GetStorageUsage is a free data retrieval call binding the contract method 0xbf1a5af7.
//
// Solidity: function getStorageUsage(address addr) view returns((uint256) usage)
func (_Credit *CreditCaller) GetStorageUsage(opts *bind.CallOpts, addr common.Address) (Usage, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "getStorageUsage", addr)

	if err != nil {
		return *new(Usage), err
	}

	out0 := *abi.ConvertType(out[0], new(Usage)).(*Usage)

	return out0, err

}

// GetStorageUsage is a free data retrieval call binding the contract method 0xbf1a5af7.
//
// Solidity: function getStorageUsage(address addr) view returns((uint256) usage)
func (_Credit *CreditSession) GetStorageUsage(addr common.Address) (Usage, error) {
	return _Credit.Contract.GetStorageUsage(&_Credit.CallOpts, addr)
}

// GetStorageUsage is a free data retrieval call binding the contract method 0xbf1a5af7.
//
// Solidity: function getStorageUsage(address addr) view returns((uint256) usage)
func (_Credit *CreditCallerSession) GetStorageUsage(addr common.Address) (Usage, error) {
	return _Credit.Contract.GetStorageUsage(&_Credit.CallOpts, addr)
}

// GetSubnetStats is a free data retrieval call binding the contract method 0x67f6c710.
//
// Solidity: function getSubnetStats() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64) stats)
func (_Credit *CreditCaller) GetSubnetStats(opts *bind.CallOpts) (SubnetStats, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "getSubnetStats")

	if err != nil {
		return *new(SubnetStats), err
	}

	out0 := *abi.ConvertType(out[0], new(SubnetStats)).(*SubnetStats)

	return out0, err

}

// GetSubnetStats is a free data retrieval call binding the contract method 0x67f6c710.
//
// Solidity: function getSubnetStats() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64) stats)
func (_Credit *CreditSession) GetSubnetStats() (SubnetStats, error) {
	return _Credit.Contract.GetSubnetStats(&_Credit.CallOpts)
}

// GetSubnetStats is a free data retrieval call binding the contract method 0x67f6c710.
//
// Solidity: function getSubnetStats() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64) stats)
func (_Credit *CreditCallerSession) GetSubnetStats() (SubnetStats, error) {
	return _Credit.Contract.GetSubnetStats(&_Credit.CallOpts)
}

// ApproveCredit is a paid mutator transaction binding the contract method 0x01e98bfa.
//
// Solidity: function approveCredit(address receiver) returns()
func (_Credit *CreditTransactor) ApproveCredit(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approveCredit", receiver)
}

// ApproveCredit is a paid mutator transaction binding the contract method 0x01e98bfa.
//
// Solidity: function approveCredit(address receiver) returns()
func (_Credit *CreditSession) ApproveCredit(receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit(&_Credit.TransactOpts, receiver)
}

// ApproveCredit is a paid mutator transaction binding the contract method 0x01e98bfa.
//
// Solidity: function approveCredit(address receiver) returns()
func (_Credit *CreditTransactorSession) ApproveCredit(receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit(&_Credit.TransactOpts, receiver)
}

// ApproveCredit0 is a paid mutator transaction binding the contract method 0x5e1087ac.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller) returns()
func (_Credit *CreditTransactor) ApproveCredit0(opts *bind.TransactOpts, from common.Address, receiver common.Address, requiredCaller common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approveCredit0", from, receiver, requiredCaller)
}

// ApproveCredit0 is a paid mutator transaction binding the contract method 0x5e1087ac.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller) returns()
func (_Credit *CreditSession) ApproveCredit0(from common.Address, receiver common.Address, requiredCaller common.Address) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit0(&_Credit.TransactOpts, from, receiver, requiredCaller)
}

// ApproveCredit0 is a paid mutator transaction binding the contract method 0x5e1087ac.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller) returns()
func (_Credit *CreditTransactorSession) ApproveCredit0(from common.Address, receiver common.Address, requiredCaller common.Address) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit0(&_Credit.TransactOpts, from, receiver, requiredCaller)
}

// ApproveCredit1 is a paid mutator transaction binding the contract method 0x7efdf0fe.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller, uint256 limit, uint64 ttl) returns()
func (_Credit *CreditTransactor) ApproveCredit1(opts *bind.TransactOpts, from common.Address, receiver common.Address, requiredCaller common.Address, limit *big.Int, ttl uint64) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approveCredit1", from, receiver, requiredCaller, limit, ttl)
}

// ApproveCredit1 is a paid mutator transaction binding the contract method 0x7efdf0fe.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller, uint256 limit, uint64 ttl) returns()
func (_Credit *CreditSession) ApproveCredit1(from common.Address, receiver common.Address, requiredCaller common.Address, limit *big.Int, ttl uint64) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit1(&_Credit.TransactOpts, from, receiver, requiredCaller, limit, ttl)
}

// ApproveCredit1 is a paid mutator transaction binding the contract method 0x7efdf0fe.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller, uint256 limit, uint64 ttl) returns()
func (_Credit *CreditTransactorSession) ApproveCredit1(from common.Address, receiver common.Address, requiredCaller common.Address, limit *big.Int, ttl uint64) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit1(&_Credit.TransactOpts, from, receiver, requiredCaller, limit, ttl)
}

// ApproveCredit2 is a paid mutator transaction binding the contract method 0xa4f3809f.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller, uint256 limit) returns()
func (_Credit *CreditTransactor) ApproveCredit2(opts *bind.TransactOpts, from common.Address, receiver common.Address, requiredCaller common.Address, limit *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approveCredit2", from, receiver, requiredCaller, limit)
}

// ApproveCredit2 is a paid mutator transaction binding the contract method 0xa4f3809f.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller, uint256 limit) returns()
func (_Credit *CreditSession) ApproveCredit2(from common.Address, receiver common.Address, requiredCaller common.Address, limit *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit2(&_Credit.TransactOpts, from, receiver, requiredCaller, limit)
}

// ApproveCredit2 is a paid mutator transaction binding the contract method 0xa4f3809f.
//
// Solidity: function approveCredit(address from, address receiver, address requiredCaller, uint256 limit) returns()
func (_Credit *CreditTransactorSession) ApproveCredit2(from common.Address, receiver common.Address, requiredCaller common.Address, limit *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit2(&_Credit.TransactOpts, from, receiver, requiredCaller, limit)
}

// ApproveCredit3 is a paid mutator transaction binding the contract method 0xf1d6ed64.
//
// Solidity: function approveCredit(address from, address receiver) returns()
func (_Credit *CreditTransactor) ApproveCredit3(opts *bind.TransactOpts, from common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approveCredit3", from, receiver)
}

// ApproveCredit3 is a paid mutator transaction binding the contract method 0xf1d6ed64.
//
// Solidity: function approveCredit(address from, address receiver) returns()
func (_Credit *CreditSession) ApproveCredit3(from common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit3(&_Credit.TransactOpts, from, receiver)
}

// ApproveCredit3 is a paid mutator transaction binding the contract method 0xf1d6ed64.
//
// Solidity: function approveCredit(address from, address receiver) returns()
func (_Credit *CreditTransactorSession) ApproveCredit3(from common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.ApproveCredit3(&_Credit.TransactOpts, from, receiver)
}

// BuyCredit is a paid mutator transaction binding the contract method 0x8e4e6f06.
//
// Solidity: function buyCredit() payable returns()
func (_Credit *CreditTransactor) BuyCredit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "buyCredit")
}

// BuyCredit is a paid mutator transaction binding the contract method 0x8e4e6f06.
//
// Solidity: function buyCredit() payable returns()
func (_Credit *CreditSession) BuyCredit() (*types.Transaction, error) {
	return _Credit.Contract.BuyCredit(&_Credit.TransactOpts)
}

// BuyCredit is a paid mutator transaction binding the contract method 0x8e4e6f06.
//
// Solidity: function buyCredit() payable returns()
func (_Credit *CreditTransactorSession) BuyCredit() (*types.Transaction, error) {
	return _Credit.Contract.BuyCredit(&_Credit.TransactOpts)
}

// BuyCredit0 is a paid mutator transaction binding the contract method 0xa38eae9f.
//
// Solidity: function buyCredit(address recipient) payable returns()
func (_Credit *CreditTransactor) BuyCredit0(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "buyCredit0", recipient)
}

// BuyCredit0 is a paid mutator transaction binding the contract method 0xa38eae9f.
//
// Solidity: function buyCredit(address recipient) payable returns()
func (_Credit *CreditSession) BuyCredit0(recipient common.Address) (*types.Transaction, error) {
	return _Credit.Contract.BuyCredit0(&_Credit.TransactOpts, recipient)
}

// BuyCredit0 is a paid mutator transaction binding the contract method 0xa38eae9f.
//
// Solidity: function buyCredit(address recipient) payable returns()
func (_Credit *CreditTransactorSession) BuyCredit0(recipient common.Address) (*types.Transaction, error) {
	return _Credit.Contract.BuyCredit0(&_Credit.TransactOpts, recipient)
}

// RevokeCredit is a paid mutator transaction binding the contract method 0xa84a1535.
//
// Solidity: function revokeCredit(address from, address receiver) returns()
func (_Credit *CreditTransactor) RevokeCredit(opts *bind.TransactOpts, from common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "revokeCredit", from, receiver)
}

// RevokeCredit is a paid mutator transaction binding the contract method 0xa84a1535.
//
// Solidity: function revokeCredit(address from, address receiver) returns()
func (_Credit *CreditSession) RevokeCredit(from common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.RevokeCredit(&_Credit.TransactOpts, from, receiver)
}

// RevokeCredit is a paid mutator transaction binding the contract method 0xa84a1535.
//
// Solidity: function revokeCredit(address from, address receiver) returns()
func (_Credit *CreditTransactorSession) RevokeCredit(from common.Address, receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.RevokeCredit(&_Credit.TransactOpts, from, receiver)
}

// RevokeCredit0 is a paid mutator transaction binding the contract method 0xa8ef8caf.
//
// Solidity: function revokeCredit(address receiver) returns()
func (_Credit *CreditTransactor) RevokeCredit0(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "revokeCredit0", receiver)
}

// RevokeCredit0 is a paid mutator transaction binding the contract method 0xa8ef8caf.
//
// Solidity: function revokeCredit(address receiver) returns()
func (_Credit *CreditSession) RevokeCredit0(receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.RevokeCredit0(&_Credit.TransactOpts, receiver)
}

// RevokeCredit0 is a paid mutator transaction binding the contract method 0xa8ef8caf.
//
// Solidity: function revokeCredit(address receiver) returns()
func (_Credit *CreditTransactorSession) RevokeCredit0(receiver common.Address) (*types.Transaction, error) {
	return _Credit.Contract.RevokeCredit0(&_Credit.TransactOpts, receiver)
}

// RevokeCredit1 is a paid mutator transaction binding the contract method 0xd188931d.
//
// Solidity: function revokeCredit(address from, address receiver, address requiredCaller) returns()
func (_Credit *CreditTransactor) RevokeCredit1(opts *bind.TransactOpts, from common.Address, receiver common.Address, requiredCaller common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "revokeCredit1", from, receiver, requiredCaller)
}

// RevokeCredit1 is a paid mutator transaction binding the contract method 0xd188931d.
//
// Solidity: function revokeCredit(address from, address receiver, address requiredCaller) returns()
func (_Credit *CreditSession) RevokeCredit1(from common.Address, receiver common.Address, requiredCaller common.Address) (*types.Transaction, error) {
	return _Credit.Contract.RevokeCredit1(&_Credit.TransactOpts, from, receiver, requiredCaller)
}

// RevokeCredit1 is a paid mutator transaction binding the contract method 0xd188931d.
//
// Solidity: function revokeCredit(address from, address receiver, address requiredCaller) returns()
func (_Credit *CreditTransactorSession) RevokeCredit1(from common.Address, receiver common.Address, requiredCaller common.Address) (*types.Transaction, error) {
	return _Credit.Contract.RevokeCredit1(&_Credit.TransactOpts, from, receiver, requiredCaller)
}

// CreditApproveCreditIterator is returned from FilterApproveCredit and is used to iterate over the raw logs and unpacked data for ApproveCredit events raised by the Credit contract.
type CreditApproveCreditIterator struct {
	Event *CreditApproveCredit // Event containing the contract specifics and raw log

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
func (it *CreditApproveCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditApproveCredit)
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
		it.Event = new(CreditApproveCredit)
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
func (it *CreditApproveCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditApproveCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditApproveCredit represents a ApproveCredit event raised by the Credit contract.
type CreditApproveCredit struct {
	From           common.Address
	Receiver       common.Address
	RequiredCaller common.Address
	Limit          *big.Int
	Ttl            uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproveCredit is a free log retrieval operation binding the contract event 0x45040e1f4609e73b821b459ec1930c24f93d6eaeff44a6f3afd62b1f6564257f.
//
// Solidity: event ApproveCredit(address indexed from, address indexed receiver, address indexed requiredCaller, uint256 limit, uint64 ttl)
func (_Credit *CreditFilterer) FilterApproveCredit(opts *bind.FilterOpts, from []common.Address, receiver []common.Address, requiredCaller []common.Address) (*CreditApproveCreditIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requiredCallerRule []interface{}
	for _, requiredCallerItem := range requiredCaller {
		requiredCallerRule = append(requiredCallerRule, requiredCallerItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "ApproveCredit", fromRule, receiverRule, requiredCallerRule)
	if err != nil {
		return nil, err
	}
	return &CreditApproveCreditIterator{contract: _Credit.contract, event: "ApproveCredit", logs: logs, sub: sub}, nil
}

// WatchApproveCredit is a free log subscription operation binding the contract event 0x45040e1f4609e73b821b459ec1930c24f93d6eaeff44a6f3afd62b1f6564257f.
//
// Solidity: event ApproveCredit(address indexed from, address indexed receiver, address indexed requiredCaller, uint256 limit, uint64 ttl)
func (_Credit *CreditFilterer) WatchApproveCredit(opts *bind.WatchOpts, sink chan<- *CreditApproveCredit, from []common.Address, receiver []common.Address, requiredCaller []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requiredCallerRule []interface{}
	for _, requiredCallerItem := range requiredCaller {
		requiredCallerRule = append(requiredCallerRule, requiredCallerItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "ApproveCredit", fromRule, receiverRule, requiredCallerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditApproveCredit)
				if err := _Credit.contract.UnpackLog(event, "ApproveCredit", log); err != nil {
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

// ParseApproveCredit is a log parse operation binding the contract event 0x45040e1f4609e73b821b459ec1930c24f93d6eaeff44a6f3afd62b1f6564257f.
//
// Solidity: event ApproveCredit(address indexed from, address indexed receiver, address indexed requiredCaller, uint256 limit, uint64 ttl)
func (_Credit *CreditFilterer) ParseApproveCredit(log types.Log) (*CreditApproveCredit, error) {
	event := new(CreditApproveCredit)
	if err := _Credit.contract.UnpackLog(event, "ApproveCredit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditBuyCreditIterator is returned from FilterBuyCredit and is used to iterate over the raw logs and unpacked data for BuyCredit events raised by the Credit contract.
type CreditBuyCreditIterator struct {
	Event *CreditBuyCredit // Event containing the contract specifics and raw log

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
func (it *CreditBuyCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditBuyCredit)
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
		it.Event = new(CreditBuyCredit)
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
func (it *CreditBuyCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditBuyCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditBuyCredit represents a BuyCredit event raised by the Credit contract.
type CreditBuyCredit struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyCredit is a free log retrieval operation binding the contract event 0x62bc306e27ea992b4fe17bf6ae7d2618241afa1a5cc4dd29590e0613dc77f47e.
//
// Solidity: event BuyCredit(address indexed addr, uint256 amount)
func (_Credit *CreditFilterer) FilterBuyCredit(opts *bind.FilterOpts, addr []common.Address) (*CreditBuyCreditIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "BuyCredit", addrRule)
	if err != nil {
		return nil, err
	}
	return &CreditBuyCreditIterator{contract: _Credit.contract, event: "BuyCredit", logs: logs, sub: sub}, nil
}

// WatchBuyCredit is a free log subscription operation binding the contract event 0x62bc306e27ea992b4fe17bf6ae7d2618241afa1a5cc4dd29590e0613dc77f47e.
//
// Solidity: event BuyCredit(address indexed addr, uint256 amount)
func (_Credit *CreditFilterer) WatchBuyCredit(opts *bind.WatchOpts, sink chan<- *CreditBuyCredit, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "BuyCredit", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditBuyCredit)
				if err := _Credit.contract.UnpackLog(event, "BuyCredit", log); err != nil {
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

// ParseBuyCredit is a log parse operation binding the contract event 0x62bc306e27ea992b4fe17bf6ae7d2618241afa1a5cc4dd29590e0613dc77f47e.
//
// Solidity: event BuyCredit(address indexed addr, uint256 amount)
func (_Credit *CreditFilterer) ParseBuyCredit(log types.Log) (*CreditBuyCredit, error) {
	event := new(CreditBuyCredit)
	if err := _Credit.contract.UnpackLog(event, "BuyCredit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditRevokeCreditIterator is returned from FilterRevokeCredit and is used to iterate over the raw logs and unpacked data for RevokeCredit events raised by the Credit contract.
type CreditRevokeCreditIterator struct {
	Event *CreditRevokeCredit // Event containing the contract specifics and raw log

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
func (it *CreditRevokeCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditRevokeCredit)
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
		it.Event = new(CreditRevokeCredit)
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
func (it *CreditRevokeCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditRevokeCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditRevokeCredit represents a RevokeCredit event raised by the Credit contract.
type CreditRevokeCredit struct {
	From           common.Address
	Receiver       common.Address
	RequiredCaller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevokeCredit is a free log retrieval operation binding the contract event 0x1094af57a7dd84b3b21a2aa03a74001ace2e9bbd3ed4d17a884c4fea3dd8da04.
//
// Solidity: event RevokeCredit(address indexed from, address indexed receiver, address indexed requiredCaller)
func (_Credit *CreditFilterer) FilterRevokeCredit(opts *bind.FilterOpts, from []common.Address, receiver []common.Address, requiredCaller []common.Address) (*CreditRevokeCreditIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requiredCallerRule []interface{}
	for _, requiredCallerItem := range requiredCaller {
		requiredCallerRule = append(requiredCallerRule, requiredCallerItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "RevokeCredit", fromRule, receiverRule, requiredCallerRule)
	if err != nil {
		return nil, err
	}
	return &CreditRevokeCreditIterator{contract: _Credit.contract, event: "RevokeCredit", logs: logs, sub: sub}, nil
}

// WatchRevokeCredit is a free log subscription operation binding the contract event 0x1094af57a7dd84b3b21a2aa03a74001ace2e9bbd3ed4d17a884c4fea3dd8da04.
//
// Solidity: event RevokeCredit(address indexed from, address indexed receiver, address indexed requiredCaller)
func (_Credit *CreditFilterer) WatchRevokeCredit(opts *bind.WatchOpts, sink chan<- *CreditRevokeCredit, from []common.Address, receiver []common.Address, requiredCaller []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requiredCallerRule []interface{}
	for _, requiredCallerItem := range requiredCaller {
		requiredCallerRule = append(requiredCallerRule, requiredCallerItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "RevokeCredit", fromRule, receiverRule, requiredCallerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditRevokeCredit)
				if err := _Credit.contract.UnpackLog(event, "RevokeCredit", log); err != nil {
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

// ParseRevokeCredit is a log parse operation binding the contract event 0x1094af57a7dd84b3b21a2aa03a74001ace2e9bbd3ed4d17a884c4fea3dd8da04.
//
// Solidity: event RevokeCredit(address indexed from, address indexed receiver, address indexed requiredCaller)
func (_Credit *CreditFilterer) ParseRevokeCredit(log types.Log) (*CreditRevokeCredit, error) {
	event := new(CreditRevokeCredit)
	if err := _Credit.contract.UnpackLog(event, "RevokeCredit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
