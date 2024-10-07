// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subnet

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

// Asset is an auto generated low-level Go binding around an user-defined struct.
type Asset struct {
	Kind         uint8
	TokenAddress common.Address
}

// BottomUpCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type BottomUpCheckpoint struct {
	SubnetID                SubnetID
	BlockHeight             *big.Int
	BlockHash               [32]byte
	NextConfigurationNumber uint64
	Msgs                    []IpcEnvelope
}

// FvmAddress is an auto generated low-level Go binding around an user-defined struct.
type FvmAddress struct {
	AddrType uint8
	Payload  []byte
}

// IPCAddress is an auto generated low-level Go binding around an user-defined struct.
type IPCAddress struct {
	SubnetId   SubnetID
	RawAddress FvmAddress
}

// IpcEnvelope is an auto generated low-level Go binding around an user-defined struct.
type IpcEnvelope struct {
	Kind    uint8
	To      IPCAddress
	From    IPCAddress
	Nonce   uint64
	Value   *big.Int
	Message []byte
}

// SubnetID is an auto generated low-level Go binding around an user-defined struct.
type SubnetID struct {
	Root  uint64
	Route []common.Address
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Weight   *big.Int
	Addr     common.Address
	Metadata []byte
}

// ValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidatorInfo struct {
	FederatedPower      *big.Int
	ConfirmedCollateral *big.Int
	TotalCollateral     *big.Int
	Metadata            []byte
}

// SubnetMetaData contains all meta data concerning the Subnet contract.
var SubnetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeValidatorsLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bootstrapped\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpCheckPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpCheckpointAtEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkpoint\",\"type\":\"tuple\",\"internalType\":\"structBottomUpCheckpoint\",\"components\":[{\"name\":\"subnetID\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextConfigurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpCheckpointHashAtEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralSource\",\"inputs\":[],\"outputs\":[{\"name\":\"supply\",\"type\":\"tuple\",\"internalType\":\"structAsset\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumAssetKind\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consensus\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumConsensusType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"crossMsgsHash\",\"inputs\":[{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"genesisBalances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"genesisCircSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"genesisValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structValidator[]\",\"components\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveValidatorsNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBootstrapNodes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfigurationNumbers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPower\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalConfirmedCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalValidatorCollateral\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalValidatorsNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"validator\",\"type\":\"tuple\",\"internalType\":\"structValidatorInfo\",\"components\":[{\"name\":\"federatedPower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"confirmedCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ipcGatewayAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActiveValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWaitingValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastBottomUpCheckpointHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"majorityPercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minActivationCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumPermissionMode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"powerScale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int8\",\"internalType\":\"int8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supplySource\",\"inputs\":[],\"outputs\":[{\"name\":\"supply\",\"type\":\"tuple\",\"internalType\":\"structAsset\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumAssetKind\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"}]",
}

// SubnetABI is the input ABI used to generate the binding from.
// Deprecated: Use SubnetMetaData.ABI instead.
var SubnetABI = SubnetMetaData.ABI

// Subnet is an auto generated Go binding around an Ethereum contract.
type Subnet struct {
	SubnetCaller     // Read-only binding to the contract
	SubnetTransactor // Write-only binding to the contract
	SubnetFilterer   // Log filterer for contract events
}

// SubnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubnetSession struct {
	Contract     *Subnet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubnetCallerSession struct {
	Contract *SubnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SubnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubnetTransactorSession struct {
	Contract     *SubnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubnetRaw struct {
	Contract *Subnet // Generic contract binding to access the raw methods on
}

// SubnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubnetCallerRaw struct {
	Contract *SubnetCaller // Generic read-only contract binding to access the raw methods on
}

// SubnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubnetTransactorRaw struct {
	Contract *SubnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubnet creates a new instance of Subnet, bound to a specific deployed contract.
func NewSubnet(address common.Address, backend bind.ContractBackend) (*Subnet, error) {
	contract, err := bindSubnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subnet{SubnetCaller: SubnetCaller{contract: contract}, SubnetTransactor: SubnetTransactor{contract: contract}, SubnetFilterer: SubnetFilterer{contract: contract}}, nil
}

// NewSubnetCaller creates a new read-only instance of Subnet, bound to a specific deployed contract.
func NewSubnetCaller(address common.Address, caller bind.ContractCaller) (*SubnetCaller, error) {
	contract, err := bindSubnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubnetCaller{contract: contract}, nil
}

// NewSubnetTransactor creates a new write-only instance of Subnet, bound to a specific deployed contract.
func NewSubnetTransactor(address common.Address, transactor bind.ContractTransactor) (*SubnetTransactor, error) {
	contract, err := bindSubnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubnetTransactor{contract: contract}, nil
}

// NewSubnetFilterer creates a new log filterer instance of Subnet, bound to a specific deployed contract.
func NewSubnetFilterer(address common.Address, filterer bind.ContractFilterer) (*SubnetFilterer, error) {
	contract, err := bindSubnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubnetFilterer{contract: contract}, nil
}

// bindSubnet binds a generic wrapper to an already deployed contract.
func bindSubnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SubnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subnet *SubnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subnet.Contract.SubnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subnet *SubnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.Contract.SubnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subnet *SubnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subnet.Contract.SubnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subnet *SubnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subnet *SubnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subnet *SubnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subnet.Contract.contract.Transact(opts, method, params...)
}

// ActiveValidatorsLimit is a free data retrieval call binding the contract method 0x3354c3e1.
//
// Solidity: function activeValidatorsLimit() view returns(uint16)
func (_Subnet *SubnetCaller) ActiveValidatorsLimit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "activeValidatorsLimit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ActiveValidatorsLimit is a free data retrieval call binding the contract method 0x3354c3e1.
//
// Solidity: function activeValidatorsLimit() view returns(uint16)
func (_Subnet *SubnetSession) ActiveValidatorsLimit() (uint16, error) {
	return _Subnet.Contract.ActiveValidatorsLimit(&_Subnet.CallOpts)
}

// ActiveValidatorsLimit is a free data retrieval call binding the contract method 0x3354c3e1.
//
// Solidity: function activeValidatorsLimit() view returns(uint16)
func (_Subnet *SubnetCallerSession) ActiveValidatorsLimit() (uint16, error) {
	return _Subnet.Contract.ActiveValidatorsLimit(&_Subnet.CallOpts)
}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_Subnet *SubnetCaller) Bootstrapped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "bootstrapped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_Subnet *SubnetSession) Bootstrapped() (bool, error) {
	return _Subnet.Contract.Bootstrapped(&_Subnet.CallOpts)
}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_Subnet *SubnetCallerSession) Bootstrapped() (bool, error) {
	return _Subnet.Contract.Bootstrapped(&_Subnet.CallOpts)
}

// BottomUpCheckPeriod is a free data retrieval call binding the contract method 0x06c46853.
//
// Solidity: function bottomUpCheckPeriod() view returns(uint256)
func (_Subnet *SubnetCaller) BottomUpCheckPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "bottomUpCheckPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BottomUpCheckPeriod is a free data retrieval call binding the contract method 0x06c46853.
//
// Solidity: function bottomUpCheckPeriod() view returns(uint256)
func (_Subnet *SubnetSession) BottomUpCheckPeriod() (*big.Int, error) {
	return _Subnet.Contract.BottomUpCheckPeriod(&_Subnet.CallOpts)
}

// BottomUpCheckPeriod is a free data retrieval call binding the contract method 0x06c46853.
//
// Solidity: function bottomUpCheckPeriod() view returns(uint256)
func (_Subnet *SubnetCallerSession) BottomUpCheckPeriod() (*big.Int, error) {
	return _Subnet.Contract.BottomUpCheckPeriod(&_Subnet.CallOpts)
}

// BottomUpCheckpointAtEpoch is a free data retrieval call binding the contract method 0x4b27aa72.
//
// Solidity: function bottomUpCheckpointAtEpoch(uint256 epoch) view returns(bool exists, ((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[]) checkpoint)
func (_Subnet *SubnetCaller) BottomUpCheckpointAtEpoch(opts *bind.CallOpts, epoch *big.Int) (struct {
	Exists     bool
	Checkpoint BottomUpCheckpoint
}, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "bottomUpCheckpointAtEpoch", epoch)

	outstruct := new(struct {
		Exists     bool
		Checkpoint BottomUpCheckpoint
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Checkpoint = *abi.ConvertType(out[1], new(BottomUpCheckpoint)).(*BottomUpCheckpoint)

	return *outstruct, err

}

// BottomUpCheckpointAtEpoch is a free data retrieval call binding the contract method 0x4b27aa72.
//
// Solidity: function bottomUpCheckpointAtEpoch(uint256 epoch) view returns(bool exists, ((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[]) checkpoint)
func (_Subnet *SubnetSession) BottomUpCheckpointAtEpoch(epoch *big.Int) (struct {
	Exists     bool
	Checkpoint BottomUpCheckpoint
}, error) {
	return _Subnet.Contract.BottomUpCheckpointAtEpoch(&_Subnet.CallOpts, epoch)
}

// BottomUpCheckpointAtEpoch is a free data retrieval call binding the contract method 0x4b27aa72.
//
// Solidity: function bottomUpCheckpointAtEpoch(uint256 epoch) view returns(bool exists, ((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[]) checkpoint)
func (_Subnet *SubnetCallerSession) BottomUpCheckpointAtEpoch(epoch *big.Int) (struct {
	Exists     bool
	Checkpoint BottomUpCheckpoint
}, error) {
	return _Subnet.Contract.BottomUpCheckpointAtEpoch(&_Subnet.CallOpts, epoch)
}

// BottomUpCheckpointHashAtEpoch is a free data retrieval call binding the contract method 0x4b0694e2.
//
// Solidity: function bottomUpCheckpointHashAtEpoch(uint256 epoch) view returns(bool, bytes32)
func (_Subnet *SubnetCaller) BottomUpCheckpointHashAtEpoch(opts *bind.CallOpts, epoch *big.Int) (bool, [32]byte, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "bottomUpCheckpointHashAtEpoch", epoch)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// BottomUpCheckpointHashAtEpoch is a free data retrieval call binding the contract method 0x4b0694e2.
//
// Solidity: function bottomUpCheckpointHashAtEpoch(uint256 epoch) view returns(bool, bytes32)
func (_Subnet *SubnetSession) BottomUpCheckpointHashAtEpoch(epoch *big.Int) (bool, [32]byte, error) {
	return _Subnet.Contract.BottomUpCheckpointHashAtEpoch(&_Subnet.CallOpts, epoch)
}

// BottomUpCheckpointHashAtEpoch is a free data retrieval call binding the contract method 0x4b0694e2.
//
// Solidity: function bottomUpCheckpointHashAtEpoch(uint256 epoch) view returns(bool, bytes32)
func (_Subnet *SubnetCallerSession) BottomUpCheckpointHashAtEpoch(epoch *big.Int) (bool, [32]byte, error) {
	return _Subnet.Contract.BottomUpCheckpointHashAtEpoch(&_Subnet.CallOpts, epoch)
}

// CollateralSource is a free data retrieval call binding the contract method 0xb6797d3c.
//
// Solidity: function collateralSource() view returns((uint8,address) supply)
func (_Subnet *SubnetCaller) CollateralSource(opts *bind.CallOpts) (Asset, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "collateralSource")

	if err != nil {
		return *new(Asset), err
	}

	out0 := *abi.ConvertType(out[0], new(Asset)).(*Asset)

	return out0, err

}

// CollateralSource is a free data retrieval call binding the contract method 0xb6797d3c.
//
// Solidity: function collateralSource() view returns((uint8,address) supply)
func (_Subnet *SubnetSession) CollateralSource() (Asset, error) {
	return _Subnet.Contract.CollateralSource(&_Subnet.CallOpts)
}

// CollateralSource is a free data retrieval call binding the contract method 0xb6797d3c.
//
// Solidity: function collateralSource() view returns((uint8,address) supply)
func (_Subnet *SubnetCallerSession) CollateralSource() (Asset, error) {
	return _Subnet.Contract.CollateralSource(&_Subnet.CallOpts)
}

// Consensus is a free data retrieval call binding the contract method 0x8ef3f761.
//
// Solidity: function consensus() view returns(uint8)
func (_Subnet *SubnetCaller) Consensus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "consensus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Consensus is a free data retrieval call binding the contract method 0x8ef3f761.
//
// Solidity: function consensus() view returns(uint8)
func (_Subnet *SubnetSession) Consensus() (uint8, error) {
	return _Subnet.Contract.Consensus(&_Subnet.CallOpts)
}

// Consensus is a free data retrieval call binding the contract method 0x8ef3f761.
//
// Solidity: function consensus() view returns(uint8)
func (_Subnet *SubnetCallerSession) Consensus() (uint8, error) {
	return _Subnet.Contract.Consensus(&_Subnet.CallOpts)
}

// CrossMsgsHash is a free data retrieval call binding the contract method 0xe02d971b.
//
// Solidity: function crossMsgsHash((uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[] messages) pure returns(bytes32)
func (_Subnet *SubnetCaller) CrossMsgsHash(opts *bind.CallOpts, messages []IpcEnvelope) ([32]byte, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "crossMsgsHash", messages)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CrossMsgsHash is a free data retrieval call binding the contract method 0xe02d971b.
//
// Solidity: function crossMsgsHash((uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[] messages) pure returns(bytes32)
func (_Subnet *SubnetSession) CrossMsgsHash(messages []IpcEnvelope) ([32]byte, error) {
	return _Subnet.Contract.CrossMsgsHash(&_Subnet.CallOpts, messages)
}

// CrossMsgsHash is a free data retrieval call binding the contract method 0xe02d971b.
//
// Solidity: function crossMsgsHash((uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[] messages) pure returns(bytes32)
func (_Subnet *SubnetCallerSession) CrossMsgsHash(messages []IpcEnvelope) ([32]byte, error) {
	return _Subnet.Contract.CrossMsgsHash(&_Subnet.CallOpts, messages)
}

// GenesisBalances is a free data retrieval call binding the contract method 0x903e6930.
//
// Solidity: function genesisBalances() view returns(address[], uint256[])
func (_Subnet *SubnetCaller) GenesisBalances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "genesisBalances")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GenesisBalances is a free data retrieval call binding the contract method 0x903e6930.
//
// Solidity: function genesisBalances() view returns(address[], uint256[])
func (_Subnet *SubnetSession) GenesisBalances() ([]common.Address, []*big.Int, error) {
	return _Subnet.Contract.GenesisBalances(&_Subnet.CallOpts)
}

// GenesisBalances is a free data retrieval call binding the contract method 0x903e6930.
//
// Solidity: function genesisBalances() view returns(address[], uint256[])
func (_Subnet *SubnetCallerSession) GenesisBalances() ([]common.Address, []*big.Int, error) {
	return _Subnet.Contract.GenesisBalances(&_Subnet.CallOpts)
}

// GenesisCircSupply is a free data retrieval call binding the contract method 0x948628a9.
//
// Solidity: function genesisCircSupply() view returns(uint256)
func (_Subnet *SubnetCaller) GenesisCircSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "genesisCircSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisCircSupply is a free data retrieval call binding the contract method 0x948628a9.
//
// Solidity: function genesisCircSupply() view returns(uint256)
func (_Subnet *SubnetSession) GenesisCircSupply() (*big.Int, error) {
	return _Subnet.Contract.GenesisCircSupply(&_Subnet.CallOpts)
}

// GenesisCircSupply is a free data retrieval call binding the contract method 0x948628a9.
//
// Solidity: function genesisCircSupply() view returns(uint256)
func (_Subnet *SubnetCallerSession) GenesisCircSupply() (*big.Int, error) {
	return _Subnet.Contract.GenesisCircSupply(&_Subnet.CallOpts)
}

// GenesisValidators is a free data retrieval call binding the contract method 0xd92e8f12.
//
// Solidity: function genesisValidators() view returns((uint256,address,bytes)[])
func (_Subnet *SubnetCaller) GenesisValidators(opts *bind.CallOpts) ([]Validator, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "genesisValidators")

	if err != nil {
		return *new([]Validator), err
	}

	out0 := *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)

	return out0, err

}

// GenesisValidators is a free data retrieval call binding the contract method 0xd92e8f12.
//
// Solidity: function genesisValidators() view returns((uint256,address,bytes)[])
func (_Subnet *SubnetSession) GenesisValidators() ([]Validator, error) {
	return _Subnet.Contract.GenesisValidators(&_Subnet.CallOpts)
}

// GenesisValidators is a free data retrieval call binding the contract method 0xd92e8f12.
//
// Solidity: function genesisValidators() view returns((uint256,address,bytes)[])
func (_Subnet *SubnetCallerSession) GenesisValidators() ([]Validator, error) {
	return _Subnet.Contract.GenesisValidators(&_Subnet.CallOpts)
}

// GetActiveValidatorsNumber is a free data retrieval call binding the contract method 0xc7cda762.
//
// Solidity: function getActiveValidatorsNumber() view returns(uint16)
func (_Subnet *SubnetCaller) GetActiveValidatorsNumber(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getActiveValidatorsNumber")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetActiveValidatorsNumber is a free data retrieval call binding the contract method 0xc7cda762.
//
// Solidity: function getActiveValidatorsNumber() view returns(uint16)
func (_Subnet *SubnetSession) GetActiveValidatorsNumber() (uint16, error) {
	return _Subnet.Contract.GetActiveValidatorsNumber(&_Subnet.CallOpts)
}

// GetActiveValidatorsNumber is a free data retrieval call binding the contract method 0xc7cda762.
//
// Solidity: function getActiveValidatorsNumber() view returns(uint16)
func (_Subnet *SubnetCallerSession) GetActiveValidatorsNumber() (uint16, error) {
	return _Subnet.Contract.GetActiveValidatorsNumber(&_Subnet.CallOpts)
}

// GetBootstrapNodes is a free data retrieval call binding the contract method 0x9754b29e.
//
// Solidity: function getBootstrapNodes() view returns(string[])
func (_Subnet *SubnetCaller) GetBootstrapNodes(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getBootstrapNodes")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetBootstrapNodes is a free data retrieval call binding the contract method 0x9754b29e.
//
// Solidity: function getBootstrapNodes() view returns(string[])
func (_Subnet *SubnetSession) GetBootstrapNodes() ([]string, error) {
	return _Subnet.Contract.GetBootstrapNodes(&_Subnet.CallOpts)
}

// GetBootstrapNodes is a free data retrieval call binding the contract method 0x9754b29e.
//
// Solidity: function getBootstrapNodes() view returns(string[])
func (_Subnet *SubnetCallerSession) GetBootstrapNodes() ([]string, error) {
	return _Subnet.Contract.GetBootstrapNodes(&_Subnet.CallOpts)
}

// GetConfigurationNumbers is a free data retrieval call binding the contract method 0x38a210b3.
//
// Solidity: function getConfigurationNumbers() view returns(uint64, uint64)
func (_Subnet *SubnetCaller) GetConfigurationNumbers(opts *bind.CallOpts) (uint64, uint64, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getConfigurationNumbers")

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetConfigurationNumbers is a free data retrieval call binding the contract method 0x38a210b3.
//
// Solidity: function getConfigurationNumbers() view returns(uint64, uint64)
func (_Subnet *SubnetSession) GetConfigurationNumbers() (uint64, uint64, error) {
	return _Subnet.Contract.GetConfigurationNumbers(&_Subnet.CallOpts)
}

// GetConfigurationNumbers is a free data retrieval call binding the contract method 0x38a210b3.
//
// Solidity: function getConfigurationNumbers() view returns(uint64, uint64)
func (_Subnet *SubnetCallerSession) GetConfigurationNumbers() (uint64, uint64, error) {
	return _Subnet.Contract.GetConfigurationNumbers(&_Subnet.CallOpts)
}

// GetParent is a free data retrieval call binding the contract method 0x80f76021.
//
// Solidity: function getParent() view returns((uint64,address[]))
func (_Subnet *SubnetCaller) GetParent(opts *bind.CallOpts) (SubnetID, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParent")

	if err != nil {
		return *new(SubnetID), err
	}

	out0 := *abi.ConvertType(out[0], new(SubnetID)).(*SubnetID)

	return out0, err

}

// GetParent is a free data retrieval call binding the contract method 0x80f76021.
//
// Solidity: function getParent() view returns((uint64,address[]))
func (_Subnet *SubnetSession) GetParent() (SubnetID, error) {
	return _Subnet.Contract.GetParent(&_Subnet.CallOpts)
}

// GetParent is a free data retrieval call binding the contract method 0x80f76021.
//
// Solidity: function getParent() view returns((uint64,address[]))
func (_Subnet *SubnetCallerSession) GetParent() (SubnetID, error) {
	return _Subnet.Contract.GetParent(&_Subnet.CallOpts)
}

// GetPower is a free data retrieval call binding the contract method 0x5dd9147c.
//
// Solidity: function getPower(address validator) view returns(uint256)
func (_Subnet *SubnetCaller) GetPower(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getPower", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPower is a free data retrieval call binding the contract method 0x5dd9147c.
//
// Solidity: function getPower(address validator) view returns(uint256)
func (_Subnet *SubnetSession) GetPower(validator common.Address) (*big.Int, error) {
	return _Subnet.Contract.GetPower(&_Subnet.CallOpts, validator)
}

// GetPower is a free data retrieval call binding the contract method 0x5dd9147c.
//
// Solidity: function getPower(address validator) view returns(uint256)
func (_Subnet *SubnetCallerSession) GetPower(validator common.Address) (*big.Int, error) {
	return _Subnet.Contract.GetPower(&_Subnet.CallOpts, validator)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_Subnet *SubnetCaller) GetTotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getTotalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_Subnet *SubnetSession) GetTotalCollateral() (*big.Int, error) {
	return _Subnet.Contract.GetTotalCollateral(&_Subnet.CallOpts)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_Subnet *SubnetCallerSession) GetTotalCollateral() (*big.Int, error) {
	return _Subnet.Contract.GetTotalCollateral(&_Subnet.CallOpts)
}

// GetTotalConfirmedCollateral is a free data retrieval call binding the contract method 0x332a5ac9.
//
// Solidity: function getTotalConfirmedCollateral() view returns(uint256)
func (_Subnet *SubnetCaller) GetTotalConfirmedCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getTotalConfirmedCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalConfirmedCollateral is a free data retrieval call binding the contract method 0x332a5ac9.
//
// Solidity: function getTotalConfirmedCollateral() view returns(uint256)
func (_Subnet *SubnetSession) GetTotalConfirmedCollateral() (*big.Int, error) {
	return _Subnet.Contract.GetTotalConfirmedCollateral(&_Subnet.CallOpts)
}

// GetTotalConfirmedCollateral is a free data retrieval call binding the contract method 0x332a5ac9.
//
// Solidity: function getTotalConfirmedCollateral() view returns(uint256)
func (_Subnet *SubnetCallerSession) GetTotalConfirmedCollateral() (*big.Int, error) {
	return _Subnet.Contract.GetTotalConfirmedCollateral(&_Subnet.CallOpts)
}

// GetTotalValidatorCollateral is a free data retrieval call binding the contract method 0x1597bf7e.
//
// Solidity: function getTotalValidatorCollateral(address validator) view returns(uint256)
func (_Subnet *SubnetCaller) GetTotalValidatorCollateral(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getTotalValidatorCollateral", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValidatorCollateral is a free data retrieval call binding the contract method 0x1597bf7e.
//
// Solidity: function getTotalValidatorCollateral(address validator) view returns(uint256)
func (_Subnet *SubnetSession) GetTotalValidatorCollateral(validator common.Address) (*big.Int, error) {
	return _Subnet.Contract.GetTotalValidatorCollateral(&_Subnet.CallOpts, validator)
}

// GetTotalValidatorCollateral is a free data retrieval call binding the contract method 0x1597bf7e.
//
// Solidity: function getTotalValidatorCollateral(address validator) view returns(uint256)
func (_Subnet *SubnetCallerSession) GetTotalValidatorCollateral(validator common.Address) (*big.Int, error) {
	return _Subnet.Contract.GetTotalValidatorCollateral(&_Subnet.CallOpts, validator)
}

// GetTotalValidatorsNumber is a free data retrieval call binding the contract method 0x52d182d1.
//
// Solidity: function getTotalValidatorsNumber() view returns(uint16)
func (_Subnet *SubnetCaller) GetTotalValidatorsNumber(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getTotalValidatorsNumber")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTotalValidatorsNumber is a free data retrieval call binding the contract method 0x52d182d1.
//
// Solidity: function getTotalValidatorsNumber() view returns(uint16)
func (_Subnet *SubnetSession) GetTotalValidatorsNumber() (uint16, error) {
	return _Subnet.Contract.GetTotalValidatorsNumber(&_Subnet.CallOpts)
}

// GetTotalValidatorsNumber is a free data retrieval call binding the contract method 0x52d182d1.
//
// Solidity: function getTotalValidatorsNumber() view returns(uint16)
func (_Subnet *SubnetCallerSession) GetTotalValidatorsNumber() (uint16, error) {
	return _Subnet.Contract.GetTotalValidatorsNumber(&_Subnet.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validatorAddress) view returns((uint256,uint256,uint256,bytes) validator)
func (_Subnet *SubnetCaller) GetValidator(opts *bind.CallOpts, validatorAddress common.Address) (ValidatorInfo, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getValidator", validatorAddress)

	if err != nil {
		return *new(ValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorInfo)).(*ValidatorInfo)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validatorAddress) view returns((uint256,uint256,uint256,bytes) validator)
func (_Subnet *SubnetSession) GetValidator(validatorAddress common.Address) (ValidatorInfo, error) {
	return _Subnet.Contract.GetValidator(&_Subnet.CallOpts, validatorAddress)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validatorAddress) view returns((uint256,uint256,uint256,bytes) validator)
func (_Subnet *SubnetCallerSession) GetValidator(validatorAddress common.Address) (ValidatorInfo, error) {
	return _Subnet.Contract.GetValidator(&_Subnet.CallOpts, validatorAddress)
}

// IpcGatewayAddr is a free data retrieval call binding the contract method 0xcfca2824.
//
// Solidity: function ipcGatewayAddr() view returns(address)
func (_Subnet *SubnetCaller) IpcGatewayAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "ipcGatewayAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IpcGatewayAddr is a free data retrieval call binding the contract method 0xcfca2824.
//
// Solidity: function ipcGatewayAddr() view returns(address)
func (_Subnet *SubnetSession) IpcGatewayAddr() (common.Address, error) {
	return _Subnet.Contract.IpcGatewayAddr(&_Subnet.CallOpts)
}

// IpcGatewayAddr is a free data retrieval call binding the contract method 0xcfca2824.
//
// Solidity: function ipcGatewayAddr() view returns(address)
func (_Subnet *SubnetCallerSession) IpcGatewayAddr() (common.Address, error) {
	return _Subnet.Contract.IpcGatewayAddr(&_Subnet.CallOpts)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address validator) view returns(bool)
func (_Subnet *SubnetCaller) IsActiveValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "isActiveValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address validator) view returns(bool)
func (_Subnet *SubnetSession) IsActiveValidator(validator common.Address) (bool, error) {
	return _Subnet.Contract.IsActiveValidator(&_Subnet.CallOpts, validator)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address validator) view returns(bool)
func (_Subnet *SubnetCallerSession) IsActiveValidator(validator common.Address) (bool, error) {
	return _Subnet.Contract.IsActiveValidator(&_Subnet.CallOpts, validator)
}

// IsWaitingValidator is a free data retrieval call binding the contract method 0xd081be03.
//
// Solidity: function isWaitingValidator(address validator) view returns(bool)
func (_Subnet *SubnetCaller) IsWaitingValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "isWaitingValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWaitingValidator is a free data retrieval call binding the contract method 0xd081be03.
//
// Solidity: function isWaitingValidator(address validator) view returns(bool)
func (_Subnet *SubnetSession) IsWaitingValidator(validator common.Address) (bool, error) {
	return _Subnet.Contract.IsWaitingValidator(&_Subnet.CallOpts, validator)
}

// IsWaitingValidator is a free data retrieval call binding the contract method 0xd081be03.
//
// Solidity: function isWaitingValidator(address validator) view returns(bool)
func (_Subnet *SubnetCallerSession) IsWaitingValidator(validator common.Address) (bool, error) {
	return _Subnet.Contract.IsWaitingValidator(&_Subnet.CallOpts, validator)
}

// Killed is a free data retrieval call binding the contract method 0x1f3a0e41.
//
// Solidity: function killed() view returns(bool)
func (_Subnet *SubnetCaller) Killed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Killed is a free data retrieval call binding the contract method 0x1f3a0e41.
//
// Solidity: function killed() view returns(bool)
func (_Subnet *SubnetSession) Killed() (bool, error) {
	return _Subnet.Contract.Killed(&_Subnet.CallOpts)
}

// Killed is a free data retrieval call binding the contract method 0x1f3a0e41.
//
// Solidity: function killed() view returns(bool)
func (_Subnet *SubnetCallerSession) Killed() (bool, error) {
	return _Subnet.Contract.Killed(&_Subnet.CallOpts)
}

// LastBottomUpCheckpointHeight is a free data retrieval call binding the contract method 0x72d0a0e0.
//
// Solidity: function lastBottomUpCheckpointHeight() view returns(uint256)
func (_Subnet *SubnetCaller) LastBottomUpCheckpointHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "lastBottomUpCheckpointHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBottomUpCheckpointHeight is a free data retrieval call binding the contract method 0x72d0a0e0.
//
// Solidity: function lastBottomUpCheckpointHeight() view returns(uint256)
func (_Subnet *SubnetSession) LastBottomUpCheckpointHeight() (*big.Int, error) {
	return _Subnet.Contract.LastBottomUpCheckpointHeight(&_Subnet.CallOpts)
}

// LastBottomUpCheckpointHeight is a free data retrieval call binding the contract method 0x72d0a0e0.
//
// Solidity: function lastBottomUpCheckpointHeight() view returns(uint256)
func (_Subnet *SubnetCallerSession) LastBottomUpCheckpointHeight() (*big.Int, error) {
	return _Subnet.Contract.LastBottomUpCheckpointHeight(&_Subnet.CallOpts)
}

// MajorityPercentage is a free data retrieval call binding the contract method 0x599c7bd1.
//
// Solidity: function majorityPercentage() view returns(uint8)
func (_Subnet *SubnetCaller) MajorityPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "majorityPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MajorityPercentage is a free data retrieval call binding the contract method 0x599c7bd1.
//
// Solidity: function majorityPercentage() view returns(uint8)
func (_Subnet *SubnetSession) MajorityPercentage() (uint8, error) {
	return _Subnet.Contract.MajorityPercentage(&_Subnet.CallOpts)
}

// MajorityPercentage is a free data retrieval call binding the contract method 0x599c7bd1.
//
// Solidity: function majorityPercentage() view returns(uint8)
func (_Subnet *SubnetCallerSession) MajorityPercentage() (uint8, error) {
	return _Subnet.Contract.MajorityPercentage(&_Subnet.CallOpts)
}

// MinActivationCollateral is a free data retrieval call binding the contract method 0x9e33bd02.
//
// Solidity: function minActivationCollateral() view returns(uint256)
func (_Subnet *SubnetCaller) MinActivationCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "minActivationCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinActivationCollateral is a free data retrieval call binding the contract method 0x9e33bd02.
//
// Solidity: function minActivationCollateral() view returns(uint256)
func (_Subnet *SubnetSession) MinActivationCollateral() (*big.Int, error) {
	return _Subnet.Contract.MinActivationCollateral(&_Subnet.CallOpts)
}

// MinActivationCollateral is a free data retrieval call binding the contract method 0x9e33bd02.
//
// Solidity: function minActivationCollateral() view returns(uint256)
func (_Subnet *SubnetCallerSession) MinActivationCollateral() (*big.Int, error) {
	return _Subnet.Contract.MinActivationCollateral(&_Subnet.CallOpts)
}

// MinValidators is a free data retrieval call binding the contract method 0xc5ab2241.
//
// Solidity: function minValidators() view returns(uint64)
func (_Subnet *SubnetCaller) MinValidators(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "minValidators")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinValidators is a free data retrieval call binding the contract method 0xc5ab2241.
//
// Solidity: function minValidators() view returns(uint64)
func (_Subnet *SubnetSession) MinValidators() (uint64, error) {
	return _Subnet.Contract.MinValidators(&_Subnet.CallOpts)
}

// MinValidators is a free data retrieval call binding the contract method 0xc5ab2241.
//
// Solidity: function minValidators() view returns(uint64)
func (_Subnet *SubnetCallerSession) MinValidators() (uint64, error) {
	return _Subnet.Contract.MinValidators(&_Subnet.CallOpts)
}

// PermissionMode is a free data retrieval call binding the contract method 0xf0cf6c96.
//
// Solidity: function permissionMode() view returns(uint8)
func (_Subnet *SubnetCaller) PermissionMode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "permissionMode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PermissionMode is a free data retrieval call binding the contract method 0xf0cf6c96.
//
// Solidity: function permissionMode() view returns(uint8)
func (_Subnet *SubnetSession) PermissionMode() (uint8, error) {
	return _Subnet.Contract.PermissionMode(&_Subnet.CallOpts)
}

// PermissionMode is a free data retrieval call binding the contract method 0xf0cf6c96.
//
// Solidity: function permissionMode() view returns(uint8)
func (_Subnet *SubnetCallerSession) PermissionMode() (uint8, error) {
	return _Subnet.Contract.PermissionMode(&_Subnet.CallOpts)
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() view returns(int8)
func (_Subnet *SubnetCaller) PowerScale(opts *bind.CallOpts) (int8, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "powerScale")

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() view returns(int8)
func (_Subnet *SubnetSession) PowerScale() (int8, error) {
	return _Subnet.Contract.PowerScale(&_Subnet.CallOpts)
}

// PowerScale is a free data retrieval call binding the contract method 0xad81e4d6.
//
// Solidity: function powerScale() view returns(int8)
func (_Subnet *SubnetCallerSession) PowerScale() (int8, error) {
	return _Subnet.Contract.PowerScale(&_Subnet.CallOpts)
}

// SupplySource is a free data retrieval call binding the contract method 0x80875df7.
//
// Solidity: function supplySource() view returns((uint8,address) supply)
func (_Subnet *SubnetCaller) SupplySource(opts *bind.CallOpts) (Asset, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "supplySource")

	if err != nil {
		return *new(Asset), err
	}

	out0 := *abi.ConvertType(out[0], new(Asset)).(*Asset)

	return out0, err

}

// SupplySource is a free data retrieval call binding the contract method 0x80875df7.
//
// Solidity: function supplySource() view returns((uint8,address) supply)
func (_Subnet *SubnetSession) SupplySource() (Asset, error) {
	return _Subnet.Contract.SupplySource(&_Subnet.CallOpts)
}

// SupplySource is a free data retrieval call binding the contract method 0x80875df7.
//
// Solidity: function supplySource() view returns((uint8,address) supply)
func (_Subnet *SubnetCallerSession) SupplySource() (Asset, error) {
	return _Subnet.Contract.SupplySource(&_Subnet.CallOpts)
}
