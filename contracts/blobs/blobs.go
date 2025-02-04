// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blobs

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

// AddBlobParams is an auto generated low-level Go binding around an user-defined struct.
type AddBlobParams struct {
	Sponsor        common.Address
	Source         string
	BlobHash       string
	MetadataHash   string
	SubscriptionId string
	Size           uint64
	Ttl            uint64
}

// Blob is an auto generated low-level Go binding around an user-defined struct.
type Blob struct {
	Size         uint64
	MetadataHash string
	Subscribers  []Subscriber
	Status       uint8
}

// BlobSourceInfo is an auto generated low-level Go binding around an user-defined struct.
type BlobSourceInfo struct {
	Subscriber     common.Address
	SubscriptionId string
	Source         string
}

// BlobTuple is an auto generated low-level Go binding around an user-defined struct.
type BlobTuple struct {
	BlobHash   string
	SourceInfo []BlobSourceInfo
}

// StorageStats is an auto generated low-level Go binding around an user-defined struct.
type StorageStats struct {
	CapacityFree   uint64
	CapacityUsed   uint64
	NumBlobs       uint64
	NumResolving   uint64
	NumAccounts    uint64
	BytesResolving uint64
	NumAdded       uint64
	BytesAdded     uint64
}

// SubnetStats is an auto generated low-level Go binding around an user-defined struct.
type SubnetStats struct {
	Balance         *big.Int
	CapacityFree    uint64
	CapacityUsed    uint64
	CreditSold      *big.Int
	CreditCommitted *big.Int
	CreditDebited   *big.Int
	TokenCreditRate *big.Int
	NumAccounts     uint64
	NumBlobs        uint64
	NumAdded        uint64
	BytesAdded      uint64
	NumResolving    uint64
	BytesResolving  uint64
}

// Subscriber is an auto generated low-level Go binding around an user-defined struct.
type Subscriber struct {
	Subscriber        common.Address
	SubscriptionGroup []SubscriptionGroup
}

// Subscription is an auto generated low-level Go binding around an user-defined struct.
type Subscription struct {
	Added    uint64
	Expiry   uint64
	Source   string
	Delegate common.Address
	Failed   bool
}

// SubscriptionGroup is an auto generated low-level Go binding around an user-defined struct.
type SubscriptionGroup struct {
	SubscriptionId string
	Subscription   Subscription
}

// BlobsMetaData contains all meta data concerning the Blobs contract.
var BlobsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addBlob\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAddBlobParams\",\"components\":[{\"name\":\"sponsor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteBlob\",\"inputs\":[{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddedBlobs\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"blobs\",\"type\":\"tuple[]\",\"internalType\":\"structBlobTuple[]\",\"components\":[{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceInfo\",\"type\":\"tuple[]\",\"internalType\":\"structBlobSourceInfo[]\",\"components\":[{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlob\",\"inputs\":[{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"blob\",\"type\":\"tuple\",\"internalType\":\"structBlob\",\"components\":[{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"metadataHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"subscribers\",\"type\":\"tuple[]\",\"internalType\":\"structSubscriber[]\",\"components\":[{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subscriptionGroup\",\"type\":\"tuple[]\",\"internalType\":\"structSubscriptionGroup[]\",\"components\":[{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"subscription\",\"type\":\"tuple\",\"internalType\":\"structSubscription\",\"components\":[{\"name\":\"added\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delegate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"failed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBlobStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobStatus\",\"inputs\":[{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumBlobStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingBlobs\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"blobs\",\"type\":\"tuple[]\",\"internalType\":\"structBlobTuple[]\",\"components\":[{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceInfo\",\"type\":\"tuple[]\",\"internalType\":\"structBlobSourceInfo[]\",\"components\":[{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingBlobsCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingBytesCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageStats\",\"inputs\":[],\"outputs\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structStorageStats\",\"components\":[{\"name\":\"capacityFree\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"capacityUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numBlobs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numResolving\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAccounts\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bytesResolving\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAdded\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bytesAdded\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStorageUsage\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubnetStats\",\"inputs\":[],\"outputs\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structSubnetStats\",\"components\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"capacityFree\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"capacityUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"creditSold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditCommitted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditDebited\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenCreditRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numAccounts\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numBlobs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numAdded\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bytesAdded\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numResolving\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"bytesResolving\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"overwriteBlob\",\"inputs\":[{\"name\":\"oldHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAddBlobParams\",\"components\":[{\"name\":\"sponsor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blobHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metadataHash\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"size\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddBlob\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sponsor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blobHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeleteBlob\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"subscriber\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blobHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OverwriteBlob\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newHash\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"subscriptionId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// BlobsABI is the input ABI used to generate the binding from.
// Deprecated: Use BlobsMetaData.ABI instead.
var BlobsABI = BlobsMetaData.ABI

// Blobs is an auto generated Go binding around an Ethereum contract.
type Blobs struct {
	BlobsCaller     // Read-only binding to the contract
	BlobsTransactor // Write-only binding to the contract
	BlobsFilterer   // Log filterer for contract events
}

// BlobsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlobsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlobsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlobsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlobsSession struct {
	Contract     *Blobs            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlobsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlobsCallerSession struct {
	Contract *BlobsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BlobsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlobsTransactorSession struct {
	Contract     *BlobsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlobsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlobsRaw struct {
	Contract *Blobs // Generic contract binding to access the raw methods on
}

// BlobsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlobsCallerRaw struct {
	Contract *BlobsCaller // Generic read-only contract binding to access the raw methods on
}

// BlobsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlobsTransactorRaw struct {
	Contract *BlobsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlobs creates a new instance of Blobs, bound to a specific deployed contract.
func NewBlobs(address common.Address, backend bind.ContractBackend) (*Blobs, error) {
	contract, err := bindBlobs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blobs{BlobsCaller: BlobsCaller{contract: contract}, BlobsTransactor: BlobsTransactor{contract: contract}, BlobsFilterer: BlobsFilterer{contract: contract}}, nil
}

// NewBlobsCaller creates a new read-only instance of Blobs, bound to a specific deployed contract.
func NewBlobsCaller(address common.Address, caller bind.ContractCaller) (*BlobsCaller, error) {
	contract, err := bindBlobs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlobsCaller{contract: contract}, nil
}

// NewBlobsTransactor creates a new write-only instance of Blobs, bound to a specific deployed contract.
func NewBlobsTransactor(address common.Address, transactor bind.ContractTransactor) (*BlobsTransactor, error) {
	contract, err := bindBlobs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlobsTransactor{contract: contract}, nil
}

// NewBlobsFilterer creates a new log filterer instance of Blobs, bound to a specific deployed contract.
func NewBlobsFilterer(address common.Address, filterer bind.ContractFilterer) (*BlobsFilterer, error) {
	contract, err := bindBlobs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlobsFilterer{contract: contract}, nil
}

// bindBlobs binds a generic wrapper to an already deployed contract.
func bindBlobs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlobsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blobs *BlobsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blobs.Contract.BlobsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blobs *BlobsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blobs.Contract.BlobsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blobs *BlobsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blobs.Contract.BlobsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blobs *BlobsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blobs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blobs *BlobsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blobs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blobs *BlobsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blobs.Contract.contract.Transact(opts, method, params...)
}

// GetAddedBlobs is a free data retrieval call binding the contract method 0x46877db8.
//
// Solidity: function getAddedBlobs(uint32 size) view returns((string,(address,string,string)[])[] blobs)
func (_Blobs *BlobsCaller) GetAddedBlobs(opts *bind.CallOpts, size uint32) ([]BlobTuple, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getAddedBlobs", size)

	if err != nil {
		return *new([]BlobTuple), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlobTuple)).(*[]BlobTuple)

	return out0, err

}

// GetAddedBlobs is a free data retrieval call binding the contract method 0x46877db8.
//
// Solidity: function getAddedBlobs(uint32 size) view returns((string,(address,string,string)[])[] blobs)
func (_Blobs *BlobsSession) GetAddedBlobs(size uint32) ([]BlobTuple, error) {
	return _Blobs.Contract.GetAddedBlobs(&_Blobs.CallOpts, size)
}

// GetAddedBlobs is a free data retrieval call binding the contract method 0x46877db8.
//
// Solidity: function getAddedBlobs(uint32 size) view returns((string,(address,string,string)[])[] blobs)
func (_Blobs *BlobsCallerSession) GetAddedBlobs(size uint32) ([]BlobTuple, error) {
	return _Blobs.Contract.GetAddedBlobs(&_Blobs.CallOpts, size)
}

// GetBlob is a free data retrieval call binding the contract method 0xf4a89bd0.
//
// Solidity: function getBlob(string blobHash) view returns((uint64,string,(address,(string,(uint64,uint64,string,address,bool))[])[],uint8) blob)
func (_Blobs *BlobsCaller) GetBlob(opts *bind.CallOpts, blobHash string) (Blob, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getBlob", blobHash)

	if err != nil {
		return *new(Blob), err
	}

	out0 := *abi.ConvertType(out[0], new(Blob)).(*Blob)

	return out0, err

}

// GetBlob is a free data retrieval call binding the contract method 0xf4a89bd0.
//
// Solidity: function getBlob(string blobHash) view returns((uint64,string,(address,(string,(uint64,uint64,string,address,bool))[])[],uint8) blob)
func (_Blobs *BlobsSession) GetBlob(blobHash string) (Blob, error) {
	return _Blobs.Contract.GetBlob(&_Blobs.CallOpts, blobHash)
}

// GetBlob is a free data retrieval call binding the contract method 0xf4a89bd0.
//
// Solidity: function getBlob(string blobHash) view returns((uint64,string,(address,(string,(uint64,uint64,string,address,bool))[])[],uint8) blob)
func (_Blobs *BlobsCallerSession) GetBlob(blobHash string) (Blob, error) {
	return _Blobs.Contract.GetBlob(&_Blobs.CallOpts, blobHash)
}

// GetBlobStatus is a free data retrieval call binding the contract method 0x554df0f8.
//
// Solidity: function getBlobStatus(address subscriber, string blobHash, string subscriptionId) view returns(uint8 status)
func (_Blobs *BlobsCaller) GetBlobStatus(opts *bind.CallOpts, subscriber common.Address, blobHash string, subscriptionId string) (uint8, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getBlobStatus", subscriber, blobHash, subscriptionId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetBlobStatus is a free data retrieval call binding the contract method 0x554df0f8.
//
// Solidity: function getBlobStatus(address subscriber, string blobHash, string subscriptionId) view returns(uint8 status)
func (_Blobs *BlobsSession) GetBlobStatus(subscriber common.Address, blobHash string, subscriptionId string) (uint8, error) {
	return _Blobs.Contract.GetBlobStatus(&_Blobs.CallOpts, subscriber, blobHash, subscriptionId)
}

// GetBlobStatus is a free data retrieval call binding the contract method 0x554df0f8.
//
// Solidity: function getBlobStatus(address subscriber, string blobHash, string subscriptionId) view returns(uint8 status)
func (_Blobs *BlobsCallerSession) GetBlobStatus(subscriber common.Address, blobHash string, subscriptionId string) (uint8, error) {
	return _Blobs.Contract.GetBlobStatus(&_Blobs.CallOpts, subscriber, blobHash, subscriptionId)
}

// GetPendingBlobs is a free data retrieval call binding the contract method 0x9d69985d.
//
// Solidity: function getPendingBlobs(uint32 size) view returns((string,(address,string,string)[])[] blobs)
func (_Blobs *BlobsCaller) GetPendingBlobs(opts *bind.CallOpts, size uint32) ([]BlobTuple, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getPendingBlobs", size)

	if err != nil {
		return *new([]BlobTuple), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlobTuple)).(*[]BlobTuple)

	return out0, err

}

// GetPendingBlobs is a free data retrieval call binding the contract method 0x9d69985d.
//
// Solidity: function getPendingBlobs(uint32 size) view returns((string,(address,string,string)[])[] blobs)
func (_Blobs *BlobsSession) GetPendingBlobs(size uint32) ([]BlobTuple, error) {
	return _Blobs.Contract.GetPendingBlobs(&_Blobs.CallOpts, size)
}

// GetPendingBlobs is a free data retrieval call binding the contract method 0x9d69985d.
//
// Solidity: function getPendingBlobs(uint32 size) view returns((string,(address,string,string)[])[] blobs)
func (_Blobs *BlobsCallerSession) GetPendingBlobs(size uint32) ([]BlobTuple, error) {
	return _Blobs.Contract.GetPendingBlobs(&_Blobs.CallOpts, size)
}

// GetPendingBlobsCount is a free data retrieval call binding the contract method 0xe7923ebb.
//
// Solidity: function getPendingBlobsCount() view returns(uint64)
func (_Blobs *BlobsCaller) GetPendingBlobsCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getPendingBlobsCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPendingBlobsCount is a free data retrieval call binding the contract method 0xe7923ebb.
//
// Solidity: function getPendingBlobsCount() view returns(uint64)
func (_Blobs *BlobsSession) GetPendingBlobsCount() (uint64, error) {
	return _Blobs.Contract.GetPendingBlobsCount(&_Blobs.CallOpts)
}

// GetPendingBlobsCount is a free data retrieval call binding the contract method 0xe7923ebb.
//
// Solidity: function getPendingBlobsCount() view returns(uint64)
func (_Blobs *BlobsCallerSession) GetPendingBlobsCount() (uint64, error) {
	return _Blobs.Contract.GetPendingBlobsCount(&_Blobs.CallOpts)
}

// GetPendingBytesCount is a free data retrieval call binding the contract method 0x7525f96d.
//
// Solidity: function getPendingBytesCount() view returns(uint64)
func (_Blobs *BlobsCaller) GetPendingBytesCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getPendingBytesCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPendingBytesCount is a free data retrieval call binding the contract method 0x7525f96d.
//
// Solidity: function getPendingBytesCount() view returns(uint64)
func (_Blobs *BlobsSession) GetPendingBytesCount() (uint64, error) {
	return _Blobs.Contract.GetPendingBytesCount(&_Blobs.CallOpts)
}

// GetPendingBytesCount is a free data retrieval call binding the contract method 0x7525f96d.
//
// Solidity: function getPendingBytesCount() view returns(uint64)
func (_Blobs *BlobsCallerSession) GetPendingBytesCount() (uint64, error) {
	return _Blobs.Contract.GetPendingBytesCount(&_Blobs.CallOpts)
}

// GetStorageStats is a free data retrieval call binding the contract method 0x537b633a.
//
// Solidity: function getStorageStats() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) stats)
func (_Blobs *BlobsCaller) GetStorageStats(opts *bind.CallOpts) (StorageStats, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getStorageStats")

	if err != nil {
		return *new(StorageStats), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageStats)).(*StorageStats)

	return out0, err

}

// GetStorageStats is a free data retrieval call binding the contract method 0x537b633a.
//
// Solidity: function getStorageStats() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) stats)
func (_Blobs *BlobsSession) GetStorageStats() (StorageStats, error) {
	return _Blobs.Contract.GetStorageStats(&_Blobs.CallOpts)
}

// GetStorageStats is a free data retrieval call binding the contract method 0x537b633a.
//
// Solidity: function getStorageStats() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) stats)
func (_Blobs *BlobsCallerSession) GetStorageStats() (StorageStats, error) {
	return _Blobs.Contract.GetStorageStats(&_Blobs.CallOpts)
}

// GetStorageUsage is a free data retrieval call binding the contract method 0xbf1a5af7.
//
// Solidity: function getStorageUsage(address addr) view returns(uint256)
func (_Blobs *BlobsCaller) GetStorageUsage(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getStorageUsage", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStorageUsage is a free data retrieval call binding the contract method 0xbf1a5af7.
//
// Solidity: function getStorageUsage(address addr) view returns(uint256)
func (_Blobs *BlobsSession) GetStorageUsage(addr common.Address) (*big.Int, error) {
	return _Blobs.Contract.GetStorageUsage(&_Blobs.CallOpts, addr)
}

// GetStorageUsage is a free data retrieval call binding the contract method 0xbf1a5af7.
//
// Solidity: function getStorageUsage(address addr) view returns(uint256)
func (_Blobs *BlobsCallerSession) GetStorageUsage(addr common.Address) (*big.Int, error) {
	return _Blobs.Contract.GetStorageUsage(&_Blobs.CallOpts, addr)
}

// GetSubnetStats is a free data retrieval call binding the contract method 0x67f6c710.
//
// Solidity: function getSubnetStats() view returns((uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64) stats)
func (_Blobs *BlobsCaller) GetSubnetStats(opts *bind.CallOpts) (SubnetStats, error) {
	var out []interface{}
	err := _Blobs.contract.Call(opts, &out, "getSubnetStats")

	if err != nil {
		return *new(SubnetStats), err
	}

	out0 := *abi.ConvertType(out[0], new(SubnetStats)).(*SubnetStats)

	return out0, err

}

// GetSubnetStats is a free data retrieval call binding the contract method 0x67f6c710.
//
// Solidity: function getSubnetStats() view returns((uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64) stats)
func (_Blobs *BlobsSession) GetSubnetStats() (SubnetStats, error) {
	return _Blobs.Contract.GetSubnetStats(&_Blobs.CallOpts)
}

// GetSubnetStats is a free data retrieval call binding the contract method 0x67f6c710.
//
// Solidity: function getSubnetStats() view returns((uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64) stats)
func (_Blobs *BlobsCallerSession) GetSubnetStats() (SubnetStats, error) {
	return _Blobs.Contract.GetSubnetStats(&_Blobs.CallOpts)
}

// AddBlob is a paid mutator transaction binding the contract method 0x6312deb8.
//
// Solidity: function addBlob((address,string,string,string,string,uint64,uint64) params) returns()
func (_Blobs *BlobsTransactor) AddBlob(opts *bind.TransactOpts, params AddBlobParams) (*types.Transaction, error) {
	return _Blobs.contract.Transact(opts, "addBlob", params)
}

// AddBlob is a paid mutator transaction binding the contract method 0x6312deb8.
//
// Solidity: function addBlob((address,string,string,string,string,uint64,uint64) params) returns()
func (_Blobs *BlobsSession) AddBlob(params AddBlobParams) (*types.Transaction, error) {
	return _Blobs.Contract.AddBlob(&_Blobs.TransactOpts, params)
}

// AddBlob is a paid mutator transaction binding the contract method 0x6312deb8.
//
// Solidity: function addBlob((address,string,string,string,string,uint64,uint64) params) returns()
func (_Blobs *BlobsTransactorSession) AddBlob(params AddBlobParams) (*types.Transaction, error) {
	return _Blobs.Contract.AddBlob(&_Blobs.TransactOpts, params)
}

// DeleteBlob is a paid mutator transaction binding the contract method 0xb9d64bde.
//
// Solidity: function deleteBlob(address subscriber, string blobHash, string subscriptionId) returns()
func (_Blobs *BlobsTransactor) DeleteBlob(opts *bind.TransactOpts, subscriber common.Address, blobHash string, subscriptionId string) (*types.Transaction, error) {
	return _Blobs.contract.Transact(opts, "deleteBlob", subscriber, blobHash, subscriptionId)
}

// DeleteBlob is a paid mutator transaction binding the contract method 0xb9d64bde.
//
// Solidity: function deleteBlob(address subscriber, string blobHash, string subscriptionId) returns()
func (_Blobs *BlobsSession) DeleteBlob(subscriber common.Address, blobHash string, subscriptionId string) (*types.Transaction, error) {
	return _Blobs.Contract.DeleteBlob(&_Blobs.TransactOpts, subscriber, blobHash, subscriptionId)
}

// DeleteBlob is a paid mutator transaction binding the contract method 0xb9d64bde.
//
// Solidity: function deleteBlob(address subscriber, string blobHash, string subscriptionId) returns()
func (_Blobs *BlobsTransactorSession) DeleteBlob(subscriber common.Address, blobHash string, subscriptionId string) (*types.Transaction, error) {
	return _Blobs.Contract.DeleteBlob(&_Blobs.TransactOpts, subscriber, blobHash, subscriptionId)
}

// OverwriteBlob is a paid mutator transaction binding the contract method 0x3024b064.
//
// Solidity: function overwriteBlob(string oldHash, (address,string,string,string,string,uint64,uint64) params) returns()
func (_Blobs *BlobsTransactor) OverwriteBlob(opts *bind.TransactOpts, oldHash string, params AddBlobParams) (*types.Transaction, error) {
	return _Blobs.contract.Transact(opts, "overwriteBlob", oldHash, params)
}

// OverwriteBlob is a paid mutator transaction binding the contract method 0x3024b064.
//
// Solidity: function overwriteBlob(string oldHash, (address,string,string,string,string,uint64,uint64) params) returns()
func (_Blobs *BlobsSession) OverwriteBlob(oldHash string, params AddBlobParams) (*types.Transaction, error) {
	return _Blobs.Contract.OverwriteBlob(&_Blobs.TransactOpts, oldHash, params)
}

// OverwriteBlob is a paid mutator transaction binding the contract method 0x3024b064.
//
// Solidity: function overwriteBlob(string oldHash, (address,string,string,string,string,uint64,uint64) params) returns()
func (_Blobs *BlobsTransactorSession) OverwriteBlob(oldHash string, params AddBlobParams) (*types.Transaction, error) {
	return _Blobs.Contract.OverwriteBlob(&_Blobs.TransactOpts, oldHash, params)
}

// BlobsAddBlobIterator is returned from FilterAddBlob and is used to iterate over the raw logs and unpacked data for AddBlob events raised by the Blobs contract.
type BlobsAddBlobIterator struct {
	Event *BlobsAddBlob // Event containing the contract specifics and raw log

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
func (it *BlobsAddBlobIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobsAddBlob)
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
		it.Event = new(BlobsAddBlob)
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
func (it *BlobsAddBlobIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlobsAddBlobIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlobsAddBlob represents a AddBlob event raised by the Blobs contract.
type BlobsAddBlob struct {
	Caller         common.Address
	Sponsor        common.Address
	BlobHash       string
	SubscriptionId string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddBlob is a free log retrieval operation binding the contract event 0xc6d59e590b1d011b1c8c1d8dc181321d5a6a46d233e4ccfbb53242c3926cf6d7.
//
// Solidity: event AddBlob(address indexed caller, address indexed sponsor, string blobHash, string subscriptionId)
func (_Blobs *BlobsFilterer) FilterAddBlob(opts *bind.FilterOpts, caller []common.Address, sponsor []common.Address) (*BlobsAddBlobIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Blobs.contract.FilterLogs(opts, "AddBlob", callerRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return &BlobsAddBlobIterator{contract: _Blobs.contract, event: "AddBlob", logs: logs, sub: sub}, nil
}

// WatchAddBlob is a free log subscription operation binding the contract event 0xc6d59e590b1d011b1c8c1d8dc181321d5a6a46d233e4ccfbb53242c3926cf6d7.
//
// Solidity: event AddBlob(address indexed caller, address indexed sponsor, string blobHash, string subscriptionId)
func (_Blobs *BlobsFilterer) WatchAddBlob(opts *bind.WatchOpts, sink chan<- *BlobsAddBlob, caller []common.Address, sponsor []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}

	logs, sub, err := _Blobs.contract.WatchLogs(opts, "AddBlob", callerRule, sponsorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlobsAddBlob)
				if err := _Blobs.contract.UnpackLog(event, "AddBlob", log); err != nil {
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

// ParseAddBlob is a log parse operation binding the contract event 0xc6d59e590b1d011b1c8c1d8dc181321d5a6a46d233e4ccfbb53242c3926cf6d7.
//
// Solidity: event AddBlob(address indexed caller, address indexed sponsor, string blobHash, string subscriptionId)
func (_Blobs *BlobsFilterer) ParseAddBlob(log types.Log) (*BlobsAddBlob, error) {
	event := new(BlobsAddBlob)
	if err := _Blobs.contract.UnpackLog(event, "AddBlob", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlobsDeleteBlobIterator is returned from FilterDeleteBlob and is used to iterate over the raw logs and unpacked data for DeleteBlob events raised by the Blobs contract.
type BlobsDeleteBlobIterator struct {
	Event *BlobsDeleteBlob // Event containing the contract specifics and raw log

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
func (it *BlobsDeleteBlobIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobsDeleteBlob)
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
		it.Event = new(BlobsDeleteBlob)
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
func (it *BlobsDeleteBlobIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlobsDeleteBlobIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlobsDeleteBlob represents a DeleteBlob event raised by the Blobs contract.
type BlobsDeleteBlob struct {
	Caller         common.Address
	Subscriber     common.Address
	BlobHash       string
	SubscriptionId string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeleteBlob is a free log retrieval operation binding the contract event 0xfc0b88b10f3dacbc1ef1a05b1b2f07f791b079f92de0d1e464b51c394125208f.
//
// Solidity: event DeleteBlob(address indexed caller, address indexed subscriber, string blobHash, string subscriptionId)
func (_Blobs *BlobsFilterer) FilterDeleteBlob(opts *bind.FilterOpts, caller []common.Address, subscriber []common.Address) (*BlobsDeleteBlobIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var subscriberRule []interface{}
	for _, subscriberItem := range subscriber {
		subscriberRule = append(subscriberRule, subscriberItem)
	}

	logs, sub, err := _Blobs.contract.FilterLogs(opts, "DeleteBlob", callerRule, subscriberRule)
	if err != nil {
		return nil, err
	}
	return &BlobsDeleteBlobIterator{contract: _Blobs.contract, event: "DeleteBlob", logs: logs, sub: sub}, nil
}

// WatchDeleteBlob is a free log subscription operation binding the contract event 0xfc0b88b10f3dacbc1ef1a05b1b2f07f791b079f92de0d1e464b51c394125208f.
//
// Solidity: event DeleteBlob(address indexed caller, address indexed subscriber, string blobHash, string subscriptionId)
func (_Blobs *BlobsFilterer) WatchDeleteBlob(opts *bind.WatchOpts, sink chan<- *BlobsDeleteBlob, caller []common.Address, subscriber []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var subscriberRule []interface{}
	for _, subscriberItem := range subscriber {
		subscriberRule = append(subscriberRule, subscriberItem)
	}

	logs, sub, err := _Blobs.contract.WatchLogs(opts, "DeleteBlob", callerRule, subscriberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlobsDeleteBlob)
				if err := _Blobs.contract.UnpackLog(event, "DeleteBlob", log); err != nil {
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

// ParseDeleteBlob is a log parse operation binding the contract event 0xfc0b88b10f3dacbc1ef1a05b1b2f07f791b079f92de0d1e464b51c394125208f.
//
// Solidity: event DeleteBlob(address indexed caller, address indexed subscriber, string blobHash, string subscriptionId)
func (_Blobs *BlobsFilterer) ParseDeleteBlob(log types.Log) (*BlobsDeleteBlob, error) {
	event := new(BlobsDeleteBlob)
	if err := _Blobs.contract.UnpackLog(event, "DeleteBlob", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlobsOverwriteBlobIterator is returned from FilterOverwriteBlob and is used to iterate over the raw logs and unpacked data for OverwriteBlob events raised by the Blobs contract.
type BlobsOverwriteBlobIterator struct {
	Event *BlobsOverwriteBlob // Event containing the contract specifics and raw log

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
func (it *BlobsOverwriteBlobIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobsOverwriteBlob)
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
		it.Event = new(BlobsOverwriteBlob)
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
func (it *BlobsOverwriteBlobIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlobsOverwriteBlobIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlobsOverwriteBlob represents a OverwriteBlob event raised by the Blobs contract.
type BlobsOverwriteBlob struct {
	Caller         common.Address
	OldHash        string
	NewHash        string
	SubscriptionId string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOverwriteBlob is a free log retrieval operation binding the contract event 0xa0926267b9d3ecf6bfeca3c26eba50c219cc76721ec8f94692034b6de9d14e0a.
//
// Solidity: event OverwriteBlob(address indexed caller, string oldHash, string newHash, string subscriptionId)
func (_Blobs *BlobsFilterer) FilterOverwriteBlob(opts *bind.FilterOpts, caller []common.Address) (*BlobsOverwriteBlobIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Blobs.contract.FilterLogs(opts, "OverwriteBlob", callerRule)
	if err != nil {
		return nil, err
	}
	return &BlobsOverwriteBlobIterator{contract: _Blobs.contract, event: "OverwriteBlob", logs: logs, sub: sub}, nil
}

// WatchOverwriteBlob is a free log subscription operation binding the contract event 0xa0926267b9d3ecf6bfeca3c26eba50c219cc76721ec8f94692034b6de9d14e0a.
//
// Solidity: event OverwriteBlob(address indexed caller, string oldHash, string newHash, string subscriptionId)
func (_Blobs *BlobsFilterer) WatchOverwriteBlob(opts *bind.WatchOpts, sink chan<- *BlobsOverwriteBlob, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Blobs.contract.WatchLogs(opts, "OverwriteBlob", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlobsOverwriteBlob)
				if err := _Blobs.contract.UnpackLog(event, "OverwriteBlob", log); err != nil {
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

// ParseOverwriteBlob is a log parse operation binding the contract event 0xa0926267b9d3ecf6bfeca3c26eba50c219cc76721ec8f94692034b6de9d14e0a.
//
// Solidity: event OverwriteBlob(address indexed caller, string oldHash, string newHash, string subscriptionId)
func (_Blobs *BlobsFilterer) ParseOverwriteBlob(log types.Log) (*BlobsOverwriteBlob, error) {
	event := new(BlobsOverwriteBlob)
	if err := _Blobs.contract.UnpackLog(event, "OverwriteBlob", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
