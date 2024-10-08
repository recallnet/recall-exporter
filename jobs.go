package main

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hokunet/hoku-exporter/contracts/gateway"
	"github.com/hokunet/hoku-exporter/contracts/subnet"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/urfave/cli/v2"
)

type Endpoint struct {
	Client *ethclient.Client
	Labels prometheus.Labels
}

type SubnetEndpoint struct {
	*Endpoint
	GatewayCaller *gateway.GatewayCaller
}

type ParentChainEndpoint struct {
	*Endpoint
	SubnetCaller *subnet.SubnetCaller
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

	// TODO: retry if it is not possible to get the chain ID.
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
		Labels: labels,
	}, nil
}

func startParentChainJobs(ctx *cli.Context) error {
	validatorAddress := common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))
	networkName := ctx.String(FLAG_PARENT_CHAIN_NETWORK_NAME)
	ep, err := connectToRpcEndpoint(
		ctx.String(FLAG_PARENT_CHAIN_RPC_URL),
		ctx.String(FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN),
		networkName)
	if err != nil {
		return err
	}

	parentChainEp := &ParentChainEndpoint{Endpoint: ep}

	subnetContractAddress := common.HexToAddress(ctx.String(FLAG_PARENT_CHAIN_SUBNET_CONTRACT_ADDRESS))
	parentChainEp.SubnetCaller, err = subnet.NewSubnetCaller(subnetContractAddress, ep.Client)
	if err != nil {
		return fmt.Errorf("failed to create SubnetCaller: %w", err)
	}

	logger := slog.With("network", networkName)

	StartJob("balance", newBalanceCheckerJob(ep, validatorAddress), ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL), logger)
	StartJob("bottomup-checkpoint", newBottomupCheckpointChecker(parentChainEp, validatorAddress), ctx.Duration(FLAG_PARENT_CHAIN_BOTTOMUP_CHECKPOINT_CHECK_INTERVAL), logger)

	return nil
}

func startSubnetJobs(ctx *cli.Context) error {
	validatorAddress := common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))
	subnetRpcUrl := ctx.String(FLAG_SUBNET_RPC_URL)
	if subnetRpcUrl == "" {
		slog.Warn("Subnet RPC URL is not configured.")
		return nil
	}

	networkName := ctx.String(FLAG_SUBNET_NETWORK_NAME)

	ep, err := connectToRpcEndpoint(subnetRpcUrl, "", networkName)
	if err != nil {
		return err
	}

	subnetEp := &SubnetEndpoint{Endpoint: ep}
	logger := slog.With("network", networkName)

	gwAddress := common.HexToAddress(ctx.String(FLAG_SUBNET_GATEWAY_ADDRESS))
	subnetEp.GatewayCaller, err = gateway.NewGatewayCaller(gwAddress, ep.Client)
	if err != nil {
		return fmt.Errorf("failed to create subnet gateway caller: %w", err)
	}

	StartJob("balance", newBalanceCheckerJob(ep, validatorAddress), ctx.Duration(FLAG_SUBNET_BALANCE_CHECK_INTERVAL), logger)
	StartJob("membership", newMembershipChecker(subnetEp), ctx.Duration(FLAG_SUBNET_MEMBERSHIP_CHECK_INTERVAL), logger)

	return nil
}

const (
	PROM_LABEL_JOB_NAME   = "job"
	PROM_LABEL_JOB_STATUS = "status"
)

var (
	counterJobRun = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "job_run",
	}, []string{PROM_LABEL_JOB_NAME, PROM_LABEL_JOB_STATUS})
)

type JobFunc func(*slog.Logger) error

func StartJob(name string, job JobFunc, interval time.Duration, parentLogger *slog.Logger) {
	logger := parentLogger.With("job", name)

	jobWrapper := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic", "reason", r, "stack", debug.Stack())
				err = fmt.Errorf("panic: %v", r)
			}
		}()
		return job(logger)
	}
	go func() {
		logger.Info("starting job")
		for {
			status := "success"
			err := jobWrapper()
			if err != nil {
				logger.Error("failed to run job", "error", err)
				status = "failure"
			}
			counterJobRun.WithLabelValues(name, status).Inc()

			logger.Debug("sleeping", "duration", interval)
			time.Sleep(interval)
		}
	}()
}
