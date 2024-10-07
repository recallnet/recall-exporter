package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hokunet/hoku-exporter/gateway"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli/v2"
)

type Endpoint struct {
	Client        *ethclient.Client
	Logger        *slog.Logger
	Labels        prometheus.Labels
	GatewayCaller *gateway.GatewayCaller
}

func connectToRpcEndpoint(rpcUrl, token, networkName string) (*Endpoint, error) {
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
	slog.Info("connected to RPC endpoint", "url", rpcUrl)

	chainId, err := client.ChainID(context.Background())
	var chainIdStr string
	if err == nil {
		chainIdStr = chainId.String()
	} else {
		slog.Error("failed to get chain ID", "error", err)
	}

	labels := prometheus.Labels{}
	labels[PROM_LABEL_NETWORK_NAME] = networkName
	labels[PROM_LABEL_CHAIN_ID] = chainIdStr

	return &Endpoint{
		Client: client,
		Logger: slog.With("network", networkName),
		Labels: labels,
	}, nil
}

func startParentChainJobs(ctx *cli.Context) error {
	validatorAddress := common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))
	ep, err := connectToRpcEndpoint(
		ctx.String(FLAG_PARENT_CHAIN_RPC_URL),
		ctx.String(FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN),
		ctx.String(FLAG_PARENT_CHAIN_NETWORK_NAME))
	if err != nil {
		return err
	}

	go runBalanceChecker(
		ep,
		validatorAddress,
		ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))

	return nil
}

func startSubnetJobs(ctx *cli.Context) error {
	validatorAddress := common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))
	subnetRpcUrl := ctx.String(FLAG_SUBNET_RPC_URL)
	if subnetRpcUrl == "" {
		slog.Warn("Subnet RPC URL is not configured.")
		return nil
	}

	ep, err := connectToRpcEndpoint(subnetRpcUrl, "", ctx.String(FLAG_SUBNET_NETWORK_NAME))
	if err != nil {
		return err
	}

	gwAddress := common.HexToAddress(ctx.String(FLAG_SUBNET_GATEWAY_ADDRESS))
	ep.GatewayCaller, err = gateway.NewGatewayCaller(gwAddress, ep.Client)
	if err != nil {
		return fmt.Errorf("failed to create subnet gateway caller: %w", err)
	}

	go runBalanceChecker(
		ep,
		validatorAddress,
		ctx.Duration(FLAG_SUBNET_BALANCE_CHECK_INTERVAL))
	go runMembershipChecker(ep, ctx.Duration(FLAG_SUBNET_MEMBERSHIP_CHECK_INTERVAL))
	return nil
}
