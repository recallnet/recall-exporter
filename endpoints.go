package main

import (
	"context"
	"fmt"
	"log/slog"
	"maps"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hokunet/hoku-exporter/contracts/gateway"
	"github.com/hokunet/hoku-exporter/contracts/subnet"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli/v2"
)

type Endpoint struct {
	Client *ethclient.Client
	labels prometheus.Labels
}

func newEndpoint(rpcUrl, token, networkName string) (*Endpoint, error) {
	rpcOptions := []rpc.ClientOption{}
	if token != "" {
		slog.Debug("Setting bearer token for the parent chain RPC endpoint.")
		rpcOptions = append(rpcOptions, rpc.WithHeader("Authorization", "Bearer "+token))
	}
	rpcClient, err := rpc.DialOptions(context.Background(), rpcUrl, rpcOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC URL %s: %w", rpcUrl, err)
	}
	client := ethclient.NewClient(rpcClient)

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID for network %s: %w", networkName, err)
	}
	slog.Info("connected to RPC endpoint", "url", rpcUrl)

	labels := prometheus.Labels{}
	labels[PROM_LABEL_NETWORK_NAME] = networkName
	labels[PROM_LABEL_CHAIN_ID] = chainId.String()

	return &Endpoint{
		Client: client,
		labels: labels,
	}, nil
}

// ConstLabels takes keyVal pairs and returns prometheus ConstLabels
func (e *Endpoint) ConstLabels(keyVal ...string) prometheus.Labels {
	labels := maps.Clone(e.labels)
	for i := 0; i < len(keyVal); i += 2 {
		labels[keyVal[i]] = keyVal[i+1]
	}
	return labels
}

type SubnetEndpoint struct {
	*Endpoint
	GatewayCaller *gateway.GatewayCaller
}

func newSubnetEndpoint(ctx *cli.Context) (*SubnetEndpoint, error) {
	subnetRpcUrl := ctx.String(FLAG_SUBNET_RPC_URL)
	networkName := ctx.String(FLAG_SUBNET_NETWORK_NAME)

	ep, err := newEndpoint(subnetRpcUrl, "", networkName)
	if err != nil {
		return nil, err
	}

	subnetEp := &SubnetEndpoint{Endpoint: ep}

	gwAddress := common.HexToAddress(ctx.String(FLAG_SUBNET_GATEWAY_ADDRESS))
	subnetEp.GatewayCaller, err = gateway.NewGatewayCaller(gwAddress, ep.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create subnet gateway caller: %w", err)
	}

	return subnetEp, nil
}

type ParentChainEndpoint struct {
	*Endpoint
	SubnetCaller *subnet.SubnetCaller
}

func newParentChainEndpoint(ctx *cli.Context) (*ParentChainEndpoint, error) {
	networkName := ctx.String(FLAG_PARENT_CHAIN_NETWORK_NAME)
	ep, err := newEndpoint(
		ctx.String(FLAG_PARENT_CHAIN_RPC_URL),
		ctx.String(FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN),
		networkName)
	if err != nil {
		return nil, err
	}

	parentChainEp := &ParentChainEndpoint{Endpoint: ep}

	subnetContractAddress := common.HexToAddress(ctx.String(FLAG_PARENT_CHAIN_SUBNET_CONTRACT_ADDRESS))
	parentChainEp.SubnetCaller, err = subnet.NewSubnetCaller(subnetContractAddress, ep.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create SubnetCaller: %w", err)
	}

	return parentChainEp, nil
}
