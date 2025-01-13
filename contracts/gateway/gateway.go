// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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

// BottomUpCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type BottomUpCheckpoint struct {
	SubnetID                SubnetID
	BlockHeight             *big.Int
	BlockHash               [32]byte
	NextConfigurationNumber uint64
	Msgs                    []IpcEnvelope
	Activity                CompressedActivityRollup
}

// BottomUpMsgBatch is an auto generated low-level Go binding around an user-defined struct.
type BottomUpMsgBatch struct {
	SubnetID    SubnetID
	BlockHeight *big.Int
	Msgs        []IpcEnvelope
}

// CompressedActivityRollup is an auto generated low-level Go binding around an user-defined struct.
type CompressedActivityRollup struct {
	Consensus ConsensusCompressedSummary
}

// ConsensusAggregatedStats is an auto generated low-level Go binding around an user-defined struct.
type ConsensusAggregatedStats struct {
	TotalActiveValidators   uint64
	TotalNumBlocksCommitted uint64
}

// ConsensusCompressedSummary is an auto generated low-level Go binding around an user-defined struct.
type ConsensusCompressedSummary struct {
	Stats              ConsensusAggregatedStats
	DataRootCommitment [32]byte
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

// Membership is an auto generated low-level Go binding around an user-defined struct.
type Membership struct {
	Validators          []Validator
	ConfigurationNumber uint64
}

// ParentFinality is an auto generated low-level Go binding around an user-defined struct.
type ParentFinality struct {
	Height    *big.Int
	BlockHash [32]byte
}

// QuorumInfo is an auto generated low-level Go binding around an user-defined struct.
type QuorumInfo struct {
	Hash          [32]byte
	RootHash      [32]byte
	Threshold     *big.Int
	CurrentWeight *big.Int
	Reached       bool
}

// Subnet is an auto generated low-level Go binding around an user-defined struct.
type Subnet struct {
	Stake                *big.Int
	GenesisEpoch         *big.Int
	CircSupply           *big.Int
	TopDownNonce         uint64
	AppliedBottomUpNonce uint64
	Id                   SubnetID
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

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"appliedTopDownNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpCheckPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpCheckpoint\",\"inputs\":[{\"name\":\"e\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBottomUpCheckpoint\",\"components\":[{\"name\":\"subnetID\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextConfigurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"activity\",\"type\":\"tuple\",\"internalType\":\"structCompressedActivityRollup\",\"components\":[{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structConsensus.CompressedSummary\",\"components\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structConsensus.AggregatedStats\",\"components\":[{\"name\":\"totalActiveValidators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalNumBlocksCommitted\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataRootCommitment\",\"type\":\"bytes32\",\"internalType\":\"Consensus.MerkleHash\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpMsgBatch\",\"inputs\":[{\"name\":\"e\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBottomUpMsgBatch\",\"components\":[{\"name\":\"subnetID\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bottomUpNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppliedBottomUpNonce\",\"inputs\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckpointCurrentWeight\",\"inputs\":[{\"name\":\"h\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckpointInfo\",\"inputs\":[{\"name\":\"h\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structQuorumInfo\",\"components\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reached\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckpointRetentionHeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckpointSignatureBundle\",\"inputs\":[{\"name\":\"h\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ch\",\"type\":\"tuple\",\"internalType\":\"structBottomUpCheckpoint\",\"components\":[{\"name\":\"subnetID\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextConfigurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"activity\",\"type\":\"tuple\",\"internalType\":\"structCompressedActivityRollup\",\"components\":[{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structConsensus.CompressedSummary\",\"components\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structConsensus.AggregatedStats\",\"components\":[{\"name\":\"totalActiveValidators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalNumBlocksCommitted\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataRootCommitment\",\"type\":\"bytes32\",\"internalType\":\"Consensus.MerkleHash\"}]}]}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structQuorumInfo\",\"components\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reached\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"signatories\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCommitSha\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentBottomUpCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"checkpoint\",\"type\":\"tuple\",\"internalType\":\"structBottomUpCheckpoint\",\"components\":[{\"name\":\"subnetID\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextConfigurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"activity\",\"type\":\"tuple\",\"internalType\":\"structCompressedActivityRollup\",\"components\":[{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structConsensus.CompressedSummary\",\"components\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structConsensus.AggregatedStats\",\"components\":[{\"name\":\"totalActiveValidators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalNumBlocksCommitted\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataRootCommitment\",\"type\":\"bytes32\",\"internalType\":\"Consensus.MerkleHash\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentConfigurationNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentMembership\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMembership\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structValidator[]\",\"components\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"configurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIncompleteCheckpointHeights\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIncompleteCheckpoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structBottomUpCheckpoint[]\",\"components\":[{\"name\":\"subnetID\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"blockHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextConfigurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structIpcEnvelope[]\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"activity\",\"type\":\"tuple\",\"internalType\":\"structCompressedActivityRollup\",\"components\":[{\"name\":\"consensus\",\"type\":\"tuple\",\"internalType\":\"structConsensus.CompressedSummary\",\"components\":[{\"name\":\"stats\",\"type\":\"tuple\",\"internalType\":\"structConsensus.AggregatedStats\",\"components\":[{\"name\":\"totalActiveValidators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalNumBlocksCommitted\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataRootCommitment\",\"type\":\"bytes32\",\"internalType\":\"Consensus.MerkleHash\"}]}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastConfigurationNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastMembership\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMembership\",\"components\":[{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structValidator[]\",\"components\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"configurationNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestParentFinality\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structParentFinality\",\"components\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNetworkName\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParentFinality\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structParentFinality\",\"components\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumThreshold\",\"inputs\":[{\"name\":\"totalWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubnet\",\"inputs\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSubnet\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"genesisEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"circSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"topDownNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appliedBottomUpNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"id\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubnetKeys\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSubnetTopDownMsgsLength\",\"inputs\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTopDownNonce\",\"inputs\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorConfigurationNumbers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listSubnets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structSubnet[]\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"genesisEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"circSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"topDownNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appliedBottomUpNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"id\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"majorityPercentage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMsgsPerBottomUpBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postbox\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"storableMsg\",\"type\":\"tuple\",\"internalType\":\"structIpcEnvelope\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumIpcMsgKind\"},{\"name\":\"to\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"from\",\"type\":\"tuple\",\"internalType\":\"structIPCAddress\",\"components\":[{\"name\":\"subnetId\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"rawAddress\",\"type\":\"tuple\",\"internalType\":\"structFvmAddress\",\"components\":[{\"name\":\"addrType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"subnets\",\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"subnet\",\"type\":\"tuple\",\"internalType\":\"structSubnet\",\"components\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"genesisEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"circSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"topDownNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"appliedBottomUpNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"id\",\"type\":\"tuple\",\"internalType\":\"structSubnetID\",\"components\":[{\"name\":\"root\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"route\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSubnets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"}]",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// AppliedTopDownNonce is a free data retrieval call binding the contract method 0x8789f83b.
//
// Solidity: function appliedTopDownNonce() view returns(uint64)
func (_Gateway *GatewayCaller) AppliedTopDownNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "appliedTopDownNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AppliedTopDownNonce is a free data retrieval call binding the contract method 0x8789f83b.
//
// Solidity: function appliedTopDownNonce() view returns(uint64)
func (_Gateway *GatewaySession) AppliedTopDownNonce() (uint64, error) {
	return _Gateway.Contract.AppliedTopDownNonce(&_Gateway.CallOpts)
}

// AppliedTopDownNonce is a free data retrieval call binding the contract method 0x8789f83b.
//
// Solidity: function appliedTopDownNonce() view returns(uint64)
func (_Gateway *GatewayCallerSession) AppliedTopDownNonce() (uint64, error) {
	return _Gateway.Contract.AppliedTopDownNonce(&_Gateway.CallOpts)
}

// BottomUpCheckPeriod is a free data retrieval call binding the contract method 0x06c46853.
//
// Solidity: function bottomUpCheckPeriod() view returns(uint256)
func (_Gateway *GatewayCaller) BottomUpCheckPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "bottomUpCheckPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BottomUpCheckPeriod is a free data retrieval call binding the contract method 0x06c46853.
//
// Solidity: function bottomUpCheckPeriod() view returns(uint256)
func (_Gateway *GatewaySession) BottomUpCheckPeriod() (*big.Int, error) {
	return _Gateway.Contract.BottomUpCheckPeriod(&_Gateway.CallOpts)
}

// BottomUpCheckPeriod is a free data retrieval call binding the contract method 0x06c46853.
//
// Solidity: function bottomUpCheckPeriod() view returns(uint256)
func (_Gateway *GatewayCallerSession) BottomUpCheckPeriod() (*big.Int, error) {
	return _Gateway.Contract.BottomUpCheckPeriod(&_Gateway.CallOpts)
}

// BottomUpCheckpoint is a free data retrieval call binding the contract method 0x2da5794a.
//
// Solidity: function bottomUpCheckpoint(uint256 e) view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))))
func (_Gateway *GatewayCaller) BottomUpCheckpoint(opts *bind.CallOpts, e *big.Int) (BottomUpCheckpoint, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "bottomUpCheckpoint", e)

	if err != nil {
		return *new(BottomUpCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(BottomUpCheckpoint)).(*BottomUpCheckpoint)

	return out0, err

}

// BottomUpCheckpoint is a free data retrieval call binding the contract method 0x2da5794a.
//
// Solidity: function bottomUpCheckpoint(uint256 e) view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))))
func (_Gateway *GatewaySession) BottomUpCheckpoint(e *big.Int) (BottomUpCheckpoint, error) {
	return _Gateway.Contract.BottomUpCheckpoint(&_Gateway.CallOpts, e)
}

// BottomUpCheckpoint is a free data retrieval call binding the contract method 0x2da5794a.
//
// Solidity: function bottomUpCheckpoint(uint256 e) view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))))
func (_Gateway *GatewayCallerSession) BottomUpCheckpoint(e *big.Int) (BottomUpCheckpoint, error) {
	return _Gateway.Contract.BottomUpCheckpoint(&_Gateway.CallOpts, e)
}

// BottomUpMsgBatch is a free data retrieval call binding the contract method 0xdd81b5cf.
//
// Solidity: function bottomUpMsgBatch(uint256 e) view returns(((uint64,address[]),uint256,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[]))
func (_Gateway *GatewayCaller) BottomUpMsgBatch(opts *bind.CallOpts, e *big.Int) (BottomUpMsgBatch, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "bottomUpMsgBatch", e)

	if err != nil {
		return *new(BottomUpMsgBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(BottomUpMsgBatch)).(*BottomUpMsgBatch)

	return out0, err

}

// BottomUpMsgBatch is a free data retrieval call binding the contract method 0xdd81b5cf.
//
// Solidity: function bottomUpMsgBatch(uint256 e) view returns(((uint64,address[]),uint256,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[]))
func (_Gateway *GatewaySession) BottomUpMsgBatch(e *big.Int) (BottomUpMsgBatch, error) {
	return _Gateway.Contract.BottomUpMsgBatch(&_Gateway.CallOpts, e)
}

// BottomUpMsgBatch is a free data retrieval call binding the contract method 0xdd81b5cf.
//
// Solidity: function bottomUpMsgBatch(uint256 e) view returns(((uint64,address[]),uint256,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[]))
func (_Gateway *GatewayCallerSession) BottomUpMsgBatch(e *big.Int) (BottomUpMsgBatch, error) {
	return _Gateway.Contract.BottomUpMsgBatch(&_Gateway.CallOpts, e)
}

// BottomUpNonce is a free data retrieval call binding the contract method 0x41b6a2e8.
//
// Solidity: function bottomUpNonce() view returns(uint64)
func (_Gateway *GatewayCaller) BottomUpNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "bottomUpNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BottomUpNonce is a free data retrieval call binding the contract method 0x41b6a2e8.
//
// Solidity: function bottomUpNonce() view returns(uint64)
func (_Gateway *GatewaySession) BottomUpNonce() (uint64, error) {
	return _Gateway.Contract.BottomUpNonce(&_Gateway.CallOpts)
}

// BottomUpNonce is a free data retrieval call binding the contract method 0x41b6a2e8.
//
// Solidity: function bottomUpNonce() view returns(uint64)
func (_Gateway *GatewayCallerSession) BottomUpNonce() (uint64, error) {
	return _Gateway.Contract.BottomUpNonce(&_Gateway.CallOpts)
}

// GetAppliedBottomUpNonce is a free data retrieval call binding the contract method 0x38d66932.
//
// Solidity: function getAppliedBottomUpNonce((uint64,address[]) subnetId) view returns(bool, uint64)
func (_Gateway *GatewayCaller) GetAppliedBottomUpNonce(opts *bind.CallOpts, subnetId SubnetID) (bool, uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getAppliedBottomUpNonce", subnetId)

	if err != nil {
		return *new(bool), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetAppliedBottomUpNonce is a free data retrieval call binding the contract method 0x38d66932.
//
// Solidity: function getAppliedBottomUpNonce((uint64,address[]) subnetId) view returns(bool, uint64)
func (_Gateway *GatewaySession) GetAppliedBottomUpNonce(subnetId SubnetID) (bool, uint64, error) {
	return _Gateway.Contract.GetAppliedBottomUpNonce(&_Gateway.CallOpts, subnetId)
}

// GetAppliedBottomUpNonce is a free data retrieval call binding the contract method 0x38d66932.
//
// Solidity: function getAppliedBottomUpNonce((uint64,address[]) subnetId) view returns(bool, uint64)
func (_Gateway *GatewayCallerSession) GetAppliedBottomUpNonce(subnetId SubnetID) (bool, uint64, error) {
	return _Gateway.Contract.GetAppliedBottomUpNonce(&_Gateway.CallOpts, subnetId)
}

// GetCheckpointCurrentWeight is a free data retrieval call binding the contract method 0xb3ab3f74.
//
// Solidity: function getCheckpointCurrentWeight(uint256 h) view returns(uint256)
func (_Gateway *GatewayCaller) GetCheckpointCurrentWeight(opts *bind.CallOpts, h *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCheckpointCurrentWeight", h)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCheckpointCurrentWeight is a free data retrieval call binding the contract method 0xb3ab3f74.
//
// Solidity: function getCheckpointCurrentWeight(uint256 h) view returns(uint256)
func (_Gateway *GatewaySession) GetCheckpointCurrentWeight(h *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetCheckpointCurrentWeight(&_Gateway.CallOpts, h)
}

// GetCheckpointCurrentWeight is a free data retrieval call binding the contract method 0xb3ab3f74.
//
// Solidity: function getCheckpointCurrentWeight(uint256 h) view returns(uint256)
func (_Gateway *GatewayCallerSession) GetCheckpointCurrentWeight(h *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetCheckpointCurrentWeight(&_Gateway.CallOpts, h)
}

// GetCheckpointInfo is a free data retrieval call binding the contract method 0xac12d763.
//
// Solidity: function getCheckpointInfo(uint256 h) view returns((bytes32,bytes32,uint256,uint256,bool))
func (_Gateway *GatewayCaller) GetCheckpointInfo(opts *bind.CallOpts, h *big.Int) (QuorumInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCheckpointInfo", h)

	if err != nil {
		return *new(QuorumInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(QuorumInfo)).(*QuorumInfo)

	return out0, err

}

// GetCheckpointInfo is a free data retrieval call binding the contract method 0xac12d763.
//
// Solidity: function getCheckpointInfo(uint256 h) view returns((bytes32,bytes32,uint256,uint256,bool))
func (_Gateway *GatewaySession) GetCheckpointInfo(h *big.Int) (QuorumInfo, error) {
	return _Gateway.Contract.GetCheckpointInfo(&_Gateway.CallOpts, h)
}

// GetCheckpointInfo is a free data retrieval call binding the contract method 0xac12d763.
//
// Solidity: function getCheckpointInfo(uint256 h) view returns((bytes32,bytes32,uint256,uint256,bool))
func (_Gateway *GatewayCallerSession) GetCheckpointInfo(h *big.Int) (QuorumInfo, error) {
	return _Gateway.Contract.GetCheckpointInfo(&_Gateway.CallOpts, h)
}

// GetCheckpointRetentionHeight is a free data retrieval call binding the contract method 0x4aa8f8a5.
//
// Solidity: function getCheckpointRetentionHeight() view returns(uint256)
func (_Gateway *GatewayCaller) GetCheckpointRetentionHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCheckpointRetentionHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCheckpointRetentionHeight is a free data retrieval call binding the contract method 0x4aa8f8a5.
//
// Solidity: function getCheckpointRetentionHeight() view returns(uint256)
func (_Gateway *GatewaySession) GetCheckpointRetentionHeight() (*big.Int, error) {
	return _Gateway.Contract.GetCheckpointRetentionHeight(&_Gateway.CallOpts)
}

// GetCheckpointRetentionHeight is a free data retrieval call binding the contract method 0x4aa8f8a5.
//
// Solidity: function getCheckpointRetentionHeight() view returns(uint256)
func (_Gateway *GatewayCallerSession) GetCheckpointRetentionHeight() (*big.Int, error) {
	return _Gateway.Contract.GetCheckpointRetentionHeight(&_Gateway.CallOpts)
}

// GetCheckpointSignatureBundle is a free data retrieval call binding the contract method 0xca41d5ce.
//
// Solidity: function getCheckpointSignatureBundle(uint256 h) view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))) ch, (bytes32,bytes32,uint256,uint256,bool) info, address[] signatories, bytes[] signatures)
func (_Gateway *GatewayCaller) GetCheckpointSignatureBundle(opts *bind.CallOpts, h *big.Int) (struct {
	Ch          BottomUpCheckpoint
	Info        QuorumInfo
	Signatories []common.Address
	Signatures  [][]byte
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCheckpointSignatureBundle", h)

	outstruct := new(struct {
		Ch          BottomUpCheckpoint
		Info        QuorumInfo
		Signatories []common.Address
		Signatures  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ch = *abi.ConvertType(out[0], new(BottomUpCheckpoint)).(*BottomUpCheckpoint)
	outstruct.Info = *abi.ConvertType(out[1], new(QuorumInfo)).(*QuorumInfo)
	outstruct.Signatories = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.Signatures = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetCheckpointSignatureBundle is a free data retrieval call binding the contract method 0xca41d5ce.
//
// Solidity: function getCheckpointSignatureBundle(uint256 h) view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))) ch, (bytes32,bytes32,uint256,uint256,bool) info, address[] signatories, bytes[] signatures)
func (_Gateway *GatewaySession) GetCheckpointSignatureBundle(h *big.Int) (struct {
	Ch          BottomUpCheckpoint
	Info        QuorumInfo
	Signatories []common.Address
	Signatures  [][]byte
}, error) {
	return _Gateway.Contract.GetCheckpointSignatureBundle(&_Gateway.CallOpts, h)
}

// GetCheckpointSignatureBundle is a free data retrieval call binding the contract method 0xca41d5ce.
//
// Solidity: function getCheckpointSignatureBundle(uint256 h) view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))) ch, (bytes32,bytes32,uint256,uint256,bool) info, address[] signatories, bytes[] signatures)
func (_Gateway *GatewayCallerSession) GetCheckpointSignatureBundle(h *big.Int) (struct {
	Ch          BottomUpCheckpoint
	Info        QuorumInfo
	Signatories []common.Address
	Signatures  [][]byte
}, error) {
	return _Gateway.Contract.GetCheckpointSignatureBundle(&_Gateway.CallOpts, h)
}

// GetCommitSha is a free data retrieval call binding the contract method 0x444ead51.
//
// Solidity: function getCommitSha() view returns(bytes32)
func (_Gateway *GatewayCaller) GetCommitSha(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCommitSha")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitSha is a free data retrieval call binding the contract method 0x444ead51.
//
// Solidity: function getCommitSha() view returns(bytes32)
func (_Gateway *GatewaySession) GetCommitSha() ([32]byte, error) {
	return _Gateway.Contract.GetCommitSha(&_Gateway.CallOpts)
}

// GetCommitSha is a free data retrieval call binding the contract method 0x444ead51.
//
// Solidity: function getCommitSha() view returns(bytes32)
func (_Gateway *GatewayCallerSession) GetCommitSha() ([32]byte, error) {
	return _Gateway.Contract.GetCommitSha(&_Gateway.CallOpts)
}

// GetCurrentBottomUpCheckpoint is a free data retrieval call binding the contract method 0xd6c5c397.
//
// Solidity: function getCurrentBottomUpCheckpoint() view returns(bool exists, uint256 epoch, ((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))) checkpoint)
func (_Gateway *GatewayCaller) GetCurrentBottomUpCheckpoint(opts *bind.CallOpts) (struct {
	Exists     bool
	Epoch      *big.Int
	Checkpoint BottomUpCheckpoint
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCurrentBottomUpCheckpoint")

	outstruct := new(struct {
		Exists     bool
		Epoch      *big.Int
		Checkpoint BottomUpCheckpoint
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Epoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Checkpoint = *abi.ConvertType(out[2], new(BottomUpCheckpoint)).(*BottomUpCheckpoint)

	return *outstruct, err

}

// GetCurrentBottomUpCheckpoint is a free data retrieval call binding the contract method 0xd6c5c397.
//
// Solidity: function getCurrentBottomUpCheckpoint() view returns(bool exists, uint256 epoch, ((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))) checkpoint)
func (_Gateway *GatewaySession) GetCurrentBottomUpCheckpoint() (struct {
	Exists     bool
	Epoch      *big.Int
	Checkpoint BottomUpCheckpoint
}, error) {
	return _Gateway.Contract.GetCurrentBottomUpCheckpoint(&_Gateway.CallOpts)
}

// GetCurrentBottomUpCheckpoint is a free data retrieval call binding the contract method 0xd6c5c397.
//
// Solidity: function getCurrentBottomUpCheckpoint() view returns(bool exists, uint256 epoch, ((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32))) checkpoint)
func (_Gateway *GatewayCallerSession) GetCurrentBottomUpCheckpoint() (struct {
	Exists     bool
	Epoch      *big.Int
	Checkpoint BottomUpCheckpoint
}, error) {
	return _Gateway.Contract.GetCurrentBottomUpCheckpoint(&_Gateway.CallOpts)
}

// GetCurrentConfigurationNumber is a free data retrieval call binding the contract method 0x544dddff.
//
// Solidity: function getCurrentConfigurationNumber() view returns(uint64)
func (_Gateway *GatewayCaller) GetCurrentConfigurationNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCurrentConfigurationNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurrentConfigurationNumber is a free data retrieval call binding the contract method 0x544dddff.
//
// Solidity: function getCurrentConfigurationNumber() view returns(uint64)
func (_Gateway *GatewaySession) GetCurrentConfigurationNumber() (uint64, error) {
	return _Gateway.Contract.GetCurrentConfigurationNumber(&_Gateway.CallOpts)
}

// GetCurrentConfigurationNumber is a free data retrieval call binding the contract method 0x544dddff.
//
// Solidity: function getCurrentConfigurationNumber() view returns(uint64)
func (_Gateway *GatewayCallerSession) GetCurrentConfigurationNumber() (uint64, error) {
	return _Gateway.Contract.GetCurrentConfigurationNumber(&_Gateway.CallOpts)
}

// GetCurrentMembership is a free data retrieval call binding the contract method 0x6ad21bb0.
//
// Solidity: function getCurrentMembership() view returns(((uint256,address,bytes)[],uint64))
func (_Gateway *GatewayCaller) GetCurrentMembership(opts *bind.CallOpts) (Membership, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getCurrentMembership")

	if err != nil {
		return *new(Membership), err
	}

	out0 := *abi.ConvertType(out[0], new(Membership)).(*Membership)

	return out0, err

}

// GetCurrentMembership is a free data retrieval call binding the contract method 0x6ad21bb0.
//
// Solidity: function getCurrentMembership() view returns(((uint256,address,bytes)[],uint64))
func (_Gateway *GatewaySession) GetCurrentMembership() (Membership, error) {
	return _Gateway.Contract.GetCurrentMembership(&_Gateway.CallOpts)
}

// GetCurrentMembership is a free data retrieval call binding the contract method 0x6ad21bb0.
//
// Solidity: function getCurrentMembership() view returns(((uint256,address,bytes)[],uint64))
func (_Gateway *GatewayCallerSession) GetCurrentMembership() (Membership, error) {
	return _Gateway.Contract.GetCurrentMembership(&_Gateway.CallOpts)
}

// GetIncompleteCheckpointHeights is a free data retrieval call binding the contract method 0xa517218f.
//
// Solidity: function getIncompleteCheckpointHeights() view returns(uint256[])
func (_Gateway *GatewayCaller) GetIncompleteCheckpointHeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getIncompleteCheckpointHeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetIncompleteCheckpointHeights is a free data retrieval call binding the contract method 0xa517218f.
//
// Solidity: function getIncompleteCheckpointHeights() view returns(uint256[])
func (_Gateway *GatewaySession) GetIncompleteCheckpointHeights() ([]*big.Int, error) {
	return _Gateway.Contract.GetIncompleteCheckpointHeights(&_Gateway.CallOpts)
}

// GetIncompleteCheckpointHeights is a free data retrieval call binding the contract method 0xa517218f.
//
// Solidity: function getIncompleteCheckpointHeights() view returns(uint256[])
func (_Gateway *GatewayCallerSession) GetIncompleteCheckpointHeights() ([]*big.Int, error) {
	return _Gateway.Contract.GetIncompleteCheckpointHeights(&_Gateway.CallOpts)
}

// GetIncompleteCheckpoints is a free data retrieval call binding the contract method 0x97042766.
//
// Solidity: function getIncompleteCheckpoints() view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32)))[])
func (_Gateway *GatewayCaller) GetIncompleteCheckpoints(opts *bind.CallOpts) ([]BottomUpCheckpoint, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getIncompleteCheckpoints")

	if err != nil {
		return *new([]BottomUpCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new([]BottomUpCheckpoint)).(*[]BottomUpCheckpoint)

	return out0, err

}

// GetIncompleteCheckpoints is a free data retrieval call binding the contract method 0x97042766.
//
// Solidity: function getIncompleteCheckpoints() view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32)))[])
func (_Gateway *GatewaySession) GetIncompleteCheckpoints() ([]BottomUpCheckpoint, error) {
	return _Gateway.Contract.GetIncompleteCheckpoints(&_Gateway.CallOpts)
}

// GetIncompleteCheckpoints is a free data retrieval call binding the contract method 0x97042766.
//
// Solidity: function getIncompleteCheckpoints() view returns(((uint64,address[]),uint256,bytes32,uint64,(uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes)[],(((uint64,uint64),bytes32)))[])
func (_Gateway *GatewayCallerSession) GetIncompleteCheckpoints() ([]BottomUpCheckpoint, error) {
	return _Gateway.Contract.GetIncompleteCheckpoints(&_Gateway.CallOpts)
}

// GetLastConfigurationNumber is a free data retrieval call binding the contract method 0xb1ba49b0.
//
// Solidity: function getLastConfigurationNumber() view returns(uint64)
func (_Gateway *GatewayCaller) GetLastConfigurationNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getLastConfigurationNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastConfigurationNumber is a free data retrieval call binding the contract method 0xb1ba49b0.
//
// Solidity: function getLastConfigurationNumber() view returns(uint64)
func (_Gateway *GatewaySession) GetLastConfigurationNumber() (uint64, error) {
	return _Gateway.Contract.GetLastConfigurationNumber(&_Gateway.CallOpts)
}

// GetLastConfigurationNumber is a free data retrieval call binding the contract method 0xb1ba49b0.
//
// Solidity: function getLastConfigurationNumber() view returns(uint64)
func (_Gateway *GatewayCallerSession) GetLastConfigurationNumber() (uint64, error) {
	return _Gateway.Contract.GetLastConfigurationNumber(&_Gateway.CallOpts)
}

// GetLastMembership is a free data retrieval call binding the contract method 0xf3229131.
//
// Solidity: function getLastMembership() view returns(((uint256,address,bytes)[],uint64))
func (_Gateway *GatewayCaller) GetLastMembership(opts *bind.CallOpts) (Membership, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getLastMembership")

	if err != nil {
		return *new(Membership), err
	}

	out0 := *abi.ConvertType(out[0], new(Membership)).(*Membership)

	return out0, err

}

// GetLastMembership is a free data retrieval call binding the contract method 0xf3229131.
//
// Solidity: function getLastMembership() view returns(((uint256,address,bytes)[],uint64))
func (_Gateway *GatewaySession) GetLastMembership() (Membership, error) {
	return _Gateway.Contract.GetLastMembership(&_Gateway.CallOpts)
}

// GetLastMembership is a free data retrieval call binding the contract method 0xf3229131.
//
// Solidity: function getLastMembership() view returns(((uint256,address,bytes)[],uint64))
func (_Gateway *GatewayCallerSession) GetLastMembership() (Membership, error) {
	return _Gateway.Contract.GetLastMembership(&_Gateway.CallOpts)
}

// GetLatestParentFinality is a free data retrieval call binding the contract method 0x0338150f.
//
// Solidity: function getLatestParentFinality() view returns((uint256,bytes32))
func (_Gateway *GatewayCaller) GetLatestParentFinality(opts *bind.CallOpts) (ParentFinality, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getLatestParentFinality")

	if err != nil {
		return *new(ParentFinality), err
	}

	out0 := *abi.ConvertType(out[0], new(ParentFinality)).(*ParentFinality)

	return out0, err

}

// GetLatestParentFinality is a free data retrieval call binding the contract method 0x0338150f.
//
// Solidity: function getLatestParentFinality() view returns((uint256,bytes32))
func (_Gateway *GatewaySession) GetLatestParentFinality() (ParentFinality, error) {
	return _Gateway.Contract.GetLatestParentFinality(&_Gateway.CallOpts)
}

// GetLatestParentFinality is a free data retrieval call binding the contract method 0x0338150f.
//
// Solidity: function getLatestParentFinality() view returns((uint256,bytes32))
func (_Gateway *GatewayCallerSession) GetLatestParentFinality() (ParentFinality, error) {
	return _Gateway.Contract.GetLatestParentFinality(&_Gateway.CallOpts)
}

// GetNetworkName is a free data retrieval call binding the contract method 0x94074b03.
//
// Solidity: function getNetworkName() view returns((uint64,address[]))
func (_Gateway *GatewayCaller) GetNetworkName(opts *bind.CallOpts) (SubnetID, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getNetworkName")

	if err != nil {
		return *new(SubnetID), err
	}

	out0 := *abi.ConvertType(out[0], new(SubnetID)).(*SubnetID)

	return out0, err

}

// GetNetworkName is a free data retrieval call binding the contract method 0x94074b03.
//
// Solidity: function getNetworkName() view returns((uint64,address[]))
func (_Gateway *GatewaySession) GetNetworkName() (SubnetID, error) {
	return _Gateway.Contract.GetNetworkName(&_Gateway.CallOpts)
}

// GetNetworkName is a free data retrieval call binding the contract method 0x94074b03.
//
// Solidity: function getNetworkName() view returns((uint64,address[]))
func (_Gateway *GatewayCallerSession) GetNetworkName() (SubnetID, error) {
	return _Gateway.Contract.GetNetworkName(&_Gateway.CallOpts)
}

// GetParentFinality is a free data retrieval call binding the contract method 0x7edeac92.
//
// Solidity: function getParentFinality(uint256 blockNumber) view returns((uint256,bytes32))
func (_Gateway *GatewayCaller) GetParentFinality(opts *bind.CallOpts, blockNumber *big.Int) (ParentFinality, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getParentFinality", blockNumber)

	if err != nil {
		return *new(ParentFinality), err
	}

	out0 := *abi.ConvertType(out[0], new(ParentFinality)).(*ParentFinality)

	return out0, err

}

// GetParentFinality is a free data retrieval call binding the contract method 0x7edeac92.
//
// Solidity: function getParentFinality(uint256 blockNumber) view returns((uint256,bytes32))
func (_Gateway *GatewaySession) GetParentFinality(blockNumber *big.Int) (ParentFinality, error) {
	return _Gateway.Contract.GetParentFinality(&_Gateway.CallOpts, blockNumber)
}

// GetParentFinality is a free data retrieval call binding the contract method 0x7edeac92.
//
// Solidity: function getParentFinality(uint256 blockNumber) view returns((uint256,bytes32))
func (_Gateway *GatewayCallerSession) GetParentFinality(blockNumber *big.Int) (ParentFinality, error) {
	return _Gateway.Contract.GetParentFinality(&_Gateway.CallOpts, blockNumber)
}

// GetQuorumThreshold is a free data retrieval call binding the contract method 0x06572c1a.
//
// Solidity: function getQuorumThreshold(uint256 totalWeight) view returns(uint256)
func (_Gateway *GatewayCaller) GetQuorumThreshold(opts *bind.CallOpts, totalWeight *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getQuorumThreshold", totalWeight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumThreshold is a free data retrieval call binding the contract method 0x06572c1a.
//
// Solidity: function getQuorumThreshold(uint256 totalWeight) view returns(uint256)
func (_Gateway *GatewaySession) GetQuorumThreshold(totalWeight *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetQuorumThreshold(&_Gateway.CallOpts, totalWeight)
}

// GetQuorumThreshold is a free data retrieval call binding the contract method 0x06572c1a.
//
// Solidity: function getQuorumThreshold(uint256 totalWeight) view returns(uint256)
func (_Gateway *GatewayCallerSession) GetQuorumThreshold(totalWeight *big.Int) (*big.Int, error) {
	return _Gateway.Contract.GetQuorumThreshold(&_Gateway.CallOpts, totalWeight)
}

// GetSubnet is a free data retrieval call binding the contract method 0xc66c66a1.
//
// Solidity: function getSubnet((uint64,address[]) subnetId) view returns(bool, (uint256,uint256,uint256,uint64,uint64,(uint64,address[])))
func (_Gateway *GatewayCaller) GetSubnet(opts *bind.CallOpts, subnetId SubnetID) (bool, Subnet, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getSubnet", subnetId)

	if err != nil {
		return *new(bool), *new(Subnet), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(Subnet)).(*Subnet)

	return out0, out1, err

}

// GetSubnet is a free data retrieval call binding the contract method 0xc66c66a1.
//
// Solidity: function getSubnet((uint64,address[]) subnetId) view returns(bool, (uint256,uint256,uint256,uint64,uint64,(uint64,address[])))
func (_Gateway *GatewaySession) GetSubnet(subnetId SubnetID) (bool, Subnet, error) {
	return _Gateway.Contract.GetSubnet(&_Gateway.CallOpts, subnetId)
}

// GetSubnet is a free data retrieval call binding the contract method 0xc66c66a1.
//
// Solidity: function getSubnet((uint64,address[]) subnetId) view returns(bool, (uint256,uint256,uint256,uint64,uint64,(uint64,address[])))
func (_Gateway *GatewayCallerSession) GetSubnet(subnetId SubnetID) (bool, Subnet, error) {
	return _Gateway.Contract.GetSubnet(&_Gateway.CallOpts, subnetId)
}

// GetSubnetKeys is a free data retrieval call binding the contract method 0x3594c3c1.
//
// Solidity: function getSubnetKeys() view returns(bytes32[])
func (_Gateway *GatewayCaller) GetSubnetKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getSubnetKeys")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSubnetKeys is a free data retrieval call binding the contract method 0x3594c3c1.
//
// Solidity: function getSubnetKeys() view returns(bytes32[])
func (_Gateway *GatewaySession) GetSubnetKeys() ([][32]byte, error) {
	return _Gateway.Contract.GetSubnetKeys(&_Gateway.CallOpts)
}

// GetSubnetKeys is a free data retrieval call binding the contract method 0x3594c3c1.
//
// Solidity: function getSubnetKeys() view returns(bytes32[])
func (_Gateway *GatewayCallerSession) GetSubnetKeys() ([][32]byte, error) {
	return _Gateway.Contract.GetSubnetKeys(&_Gateway.CallOpts)
}

// GetSubnetTopDownMsgsLength is a free data retrieval call binding the contract method 0x9d3070b5.
//
// Solidity: function getSubnetTopDownMsgsLength((uint64,address[]) subnetId) view returns(uint256)
func (_Gateway *GatewayCaller) GetSubnetTopDownMsgsLength(opts *bind.CallOpts, subnetId SubnetID) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getSubnetTopDownMsgsLength", subnetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubnetTopDownMsgsLength is a free data retrieval call binding the contract method 0x9d3070b5.
//
// Solidity: function getSubnetTopDownMsgsLength((uint64,address[]) subnetId) view returns(uint256)
func (_Gateway *GatewaySession) GetSubnetTopDownMsgsLength(subnetId SubnetID) (*big.Int, error) {
	return _Gateway.Contract.GetSubnetTopDownMsgsLength(&_Gateway.CallOpts, subnetId)
}

// GetSubnetTopDownMsgsLength is a free data retrieval call binding the contract method 0x9d3070b5.
//
// Solidity: function getSubnetTopDownMsgsLength((uint64,address[]) subnetId) view returns(uint256)
func (_Gateway *GatewayCallerSession) GetSubnetTopDownMsgsLength(subnetId SubnetID) (*big.Int, error) {
	return _Gateway.Contract.GetSubnetTopDownMsgsLength(&_Gateway.CallOpts, subnetId)
}

// GetTopDownNonce is a free data retrieval call binding the contract method 0x42398a9a.
//
// Solidity: function getTopDownNonce((uint64,address[]) subnetId) view returns(bool, uint64)
func (_Gateway *GatewayCaller) GetTopDownNonce(opts *bind.CallOpts, subnetId SubnetID) (bool, uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getTopDownNonce", subnetId)

	if err != nil {
		return *new(bool), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetTopDownNonce is a free data retrieval call binding the contract method 0x42398a9a.
//
// Solidity: function getTopDownNonce((uint64,address[]) subnetId) view returns(bool, uint64)
func (_Gateway *GatewaySession) GetTopDownNonce(subnetId SubnetID) (bool, uint64, error) {
	return _Gateway.Contract.GetTopDownNonce(&_Gateway.CallOpts, subnetId)
}

// GetTopDownNonce is a free data retrieval call binding the contract method 0x42398a9a.
//
// Solidity: function getTopDownNonce((uint64,address[]) subnetId) view returns(bool, uint64)
func (_Gateway *GatewayCallerSession) GetTopDownNonce(subnetId SubnetID) (bool, uint64, error) {
	return _Gateway.Contract.GetTopDownNonce(&_Gateway.CallOpts, subnetId)
}

// GetValidatorConfigurationNumbers is a free data retrieval call binding the contract method 0xfa34a400.
//
// Solidity: function getValidatorConfigurationNumbers() view returns(uint64, uint64)
func (_Gateway *GatewayCaller) GetValidatorConfigurationNumbers(opts *bind.CallOpts) (uint64, uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getValidatorConfigurationNumbers")

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetValidatorConfigurationNumbers is a free data retrieval call binding the contract method 0xfa34a400.
//
// Solidity: function getValidatorConfigurationNumbers() view returns(uint64, uint64)
func (_Gateway *GatewaySession) GetValidatorConfigurationNumbers() (uint64, uint64, error) {
	return _Gateway.Contract.GetValidatorConfigurationNumbers(&_Gateway.CallOpts)
}

// GetValidatorConfigurationNumbers is a free data retrieval call binding the contract method 0xfa34a400.
//
// Solidity: function getValidatorConfigurationNumbers() view returns(uint64, uint64)
func (_Gateway *GatewayCallerSession) GetValidatorConfigurationNumbers() (uint64, uint64, error) {
	return _Gateway.Contract.GetValidatorConfigurationNumbers(&_Gateway.CallOpts)
}

// ListSubnets is a free data retrieval call binding the contract method 0x5d029685.
//
// Solidity: function listSubnets() view returns((uint256,uint256,uint256,uint64,uint64,(uint64,address[]))[])
func (_Gateway *GatewayCaller) ListSubnets(opts *bind.CallOpts) ([]Subnet, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "listSubnets")

	if err != nil {
		return *new([]Subnet), err
	}

	out0 := *abi.ConvertType(out[0], new([]Subnet)).(*[]Subnet)

	return out0, err

}

// ListSubnets is a free data retrieval call binding the contract method 0x5d029685.
//
// Solidity: function listSubnets() view returns((uint256,uint256,uint256,uint64,uint64,(uint64,address[]))[])
func (_Gateway *GatewaySession) ListSubnets() ([]Subnet, error) {
	return _Gateway.Contract.ListSubnets(&_Gateway.CallOpts)
}

// ListSubnets is a free data retrieval call binding the contract method 0x5d029685.
//
// Solidity: function listSubnets() view returns((uint256,uint256,uint256,uint64,uint64,(uint64,address[]))[])
func (_Gateway *GatewayCallerSession) ListSubnets() ([]Subnet, error) {
	return _Gateway.Contract.ListSubnets(&_Gateway.CallOpts)
}

// MajorityPercentage is a free data retrieval call binding the contract method 0x599c7bd1.
//
// Solidity: function majorityPercentage() view returns(uint64)
func (_Gateway *GatewayCaller) MajorityPercentage(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "majorityPercentage")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MajorityPercentage is a free data retrieval call binding the contract method 0x599c7bd1.
//
// Solidity: function majorityPercentage() view returns(uint64)
func (_Gateway *GatewaySession) MajorityPercentage() (uint64, error) {
	return _Gateway.Contract.MajorityPercentage(&_Gateway.CallOpts)
}

// MajorityPercentage is a free data retrieval call binding the contract method 0x599c7bd1.
//
// Solidity: function majorityPercentage() view returns(uint64)
func (_Gateway *GatewayCallerSession) MajorityPercentage() (uint64, error) {
	return _Gateway.Contract.MajorityPercentage(&_Gateway.CallOpts)
}

// MaxMsgsPerBottomUpBatch is a free data retrieval call binding the contract method 0x05aff0b3.
//
// Solidity: function maxMsgsPerBottomUpBatch() view returns(uint64)
func (_Gateway *GatewayCaller) MaxMsgsPerBottomUpBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "maxMsgsPerBottomUpBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxMsgsPerBottomUpBatch is a free data retrieval call binding the contract method 0x05aff0b3.
//
// Solidity: function maxMsgsPerBottomUpBatch() view returns(uint64)
func (_Gateway *GatewaySession) MaxMsgsPerBottomUpBatch() (uint64, error) {
	return _Gateway.Contract.MaxMsgsPerBottomUpBatch(&_Gateway.CallOpts)
}

// MaxMsgsPerBottomUpBatch is a free data retrieval call binding the contract method 0x05aff0b3.
//
// Solidity: function maxMsgsPerBottomUpBatch() view returns(uint64)
func (_Gateway *GatewayCallerSession) MaxMsgsPerBottomUpBatch() (uint64, error) {
	return _Gateway.Contract.MaxMsgsPerBottomUpBatch(&_Gateway.CallOpts)
}

// Postbox is a free data retrieval call binding the contract method 0x8cfd78e7.
//
// Solidity: function postbox(bytes32 id) view returns((uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes) storableMsg)
func (_Gateway *GatewayCaller) Postbox(opts *bind.CallOpts, id [32]byte) (IpcEnvelope, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "postbox", id)

	if err != nil {
		return *new(IpcEnvelope), err
	}

	out0 := *abi.ConvertType(out[0], new(IpcEnvelope)).(*IpcEnvelope)

	return out0, err

}

// Postbox is a free data retrieval call binding the contract method 0x8cfd78e7.
//
// Solidity: function postbox(bytes32 id) view returns((uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes) storableMsg)
func (_Gateway *GatewaySession) Postbox(id [32]byte) (IpcEnvelope, error) {
	return _Gateway.Contract.Postbox(&_Gateway.CallOpts, id)
}

// Postbox is a free data retrieval call binding the contract method 0x8cfd78e7.
//
// Solidity: function postbox(bytes32 id) view returns((uint8,((uint64,address[]),(uint8,bytes)),((uint64,address[]),(uint8,bytes)),uint64,uint256,bytes) storableMsg)
func (_Gateway *GatewayCallerSession) Postbox(id [32]byte) (IpcEnvelope, error) {
	return _Gateway.Contract.Postbox(&_Gateway.CallOpts, id)
}

// Subnets is a free data retrieval call binding the contract method 0x02e30f9a.
//
// Solidity: function subnets(bytes32 h) view returns((uint256,uint256,uint256,uint64,uint64,(uint64,address[])) subnet)
func (_Gateway *GatewayCaller) Subnets(opts *bind.CallOpts, h [32]byte) (Subnet, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "subnets", h)

	if err != nil {
		return *new(Subnet), err
	}

	out0 := *abi.ConvertType(out[0], new(Subnet)).(*Subnet)

	return out0, err

}

// Subnets is a free data retrieval call binding the contract method 0x02e30f9a.
//
// Solidity: function subnets(bytes32 h) view returns((uint256,uint256,uint256,uint64,uint64,(uint64,address[])) subnet)
func (_Gateway *GatewaySession) Subnets(h [32]byte) (Subnet, error) {
	return _Gateway.Contract.Subnets(&_Gateway.CallOpts, h)
}

// Subnets is a free data retrieval call binding the contract method 0x02e30f9a.
//
// Solidity: function subnets(bytes32 h) view returns((uint256,uint256,uint256,uint64,uint64,(uint64,address[])) subnet)
func (_Gateway *GatewayCallerSession) Subnets(h [32]byte) (Subnet, error) {
	return _Gateway.Contract.Subnets(&_Gateway.CallOpts, h)
}

// TotalSubnets is a free data retrieval call binding the contract method 0xa2b67158.
//
// Solidity: function totalSubnets() view returns(uint64)
func (_Gateway *GatewayCaller) TotalSubnets(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "totalSubnets")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSubnets is a free data retrieval call binding the contract method 0xa2b67158.
//
// Solidity: function totalSubnets() view returns(uint64)
func (_Gateway *GatewaySession) TotalSubnets() (uint64, error) {
	return _Gateway.Contract.TotalSubnets(&_Gateway.CallOpts)
}

// TotalSubnets is a free data retrieval call binding the contract method 0xa2b67158.
//
// Solidity: function totalSubnets() view returns(uint64)
func (_Gateway *GatewayCallerSession) TotalSubnets() (uint64, error) {
	return _Gateway.Contract.TotalSubnets(&_Gateway.CallOpts)
}
