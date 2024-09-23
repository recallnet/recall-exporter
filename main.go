package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/tint"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const (
	FLAG_METRICS_ADDRESS                     = "metrics-address"
	FLAG_METRICS_PATH                        = "metrics-path"
	FLAG_PARENT_CHAIN_RPC_URL                = "parent-chain-rpc-url"
	FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN       = "parent-chain-rpc-bearer-token"
	FLAG_VALIDATOR_ADDRESS                   = "validator-address"
	FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL = "parent-chain-balance-check-interval"
	FLAG_SUBNET_RPC_URL                      = "subnet-rpc-url"
	FLAG_SUBNET_BALANCE_CHECK_INTERVAL       = "subnet-balance-check-interval"
)

var (
	parentChainRpcClient *ethclient.Client
	subnetRpcClient      *ethclient.Client
)

var (
	GitCommit string
	BuildTime string
)

func main() {
	app := cli.App{
		Version: GitCommit,
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run hoku-exporter",
				Action: commandRun,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    FLAG_METRICS_ADDRESS,
						Usage:   "Address to expose metrics on",
						Value:   "localhost:9010",
						EnvVars: []string{"METRICS_ADDRESS"},
					},
					&cli.StringFlag{
						Name:  FLAG_METRICS_PATH,
						Usage: "Metrics path",
						Value: "/metrics",
					},
					&cli.StringFlag{
						Name:     FLAG_PARENT_CHAIN_RPC_URL,
						Usage:    "Parent chain RPC URL",
						EnvVars:  []string{"PARENT_CHAIN_RPC_URL"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN,
						Usage:   "Parent chain RPC bearer token",
						EnvVars: []string{"PARENT_CHAIN_RPC_BEARER_TOKEN"},
					},
					&cli.StringFlag{
						Name:     FLAG_VALIDATOR_ADDRESS,
						Usage:    "Validator address",
						EnvVars:  []string{"VALIDATOR_ADDRESS"},
						Required: true,
					},
					&cli.DurationFlag{
						Name:    FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL,
						Usage:   "How often the balance on the parent chain must be checked",
						Value:   time.Minute,
						EnvVars: []string{"PARENT_CHAIN_BALANCE_CHECK_INTERVAL"},
					},
					&cli.StringFlag{
						Name:    FLAG_SUBNET_RPC_URL,
						Usage:   "Subnet RPC URL",
						EnvVars: []string{"SUBNET_RPC_URL"},
					},
					&cli.DurationFlag{
						Name:    FLAG_SUBNET_BALANCE_CHECK_INTERVAL,
						Usage:   "How often the balance on the subnet must be checked",
						Value:   time.Minute,
						EnvVars: []string{"SUBNET_BALANCE_CHECK_INTERVAL"},
					},
				},
			},
		},
	}

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.RFC3339,
	})))

	if err := app.Run(os.Args); err != nil {
		slog.Error("Failed to run app", "error", err)
	}
}

func commandRun(ctx *cli.Context) error {
	slog.Info("running hoku-exporter", "git-commit", GitCommit, "build-time", BuildTime)
	validatorAddress := common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))

	parentChainRpcClient, err := connectToParentChainRpcEndpoint(ctx, ctx.String(FLAG_PARENT_CHAIN_RPC_URL))
	if err != nil {
		return err
	}
	go runBalanceChecker(parentChainRpcClient, validatorAddress, "parent", ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))

	subnetRpcUrl := ctx.String(FLAG_SUBNET_RPC_URL)
	if subnetRpcUrl == "" {
		slog.Warn("Subnet RPC URL is not configured.")
	} else {
		subnetRpcClient, err := connectToSubnetRpcEndpoint(subnetRpcUrl)
		if err != nil {
			return err
		}
		go runBalanceChecker(subnetRpcClient, validatorAddress, "subnet", ctx.Duration(FLAG_SUBNET_BALANCE_CHECK_INTERVAL))
	}

	metricsAddress := ctx.String(FLAG_METRICS_ADDRESS)
	metricsPath := ctx.String(FLAG_METRICS_PATH)
	metricsUrl := "http://" + metricsAddress + metricsPath
	slog.Info("Exporting metrics", "url", metricsUrl)

	http.Handle(metricsPath, promhttp.Handler())
	return http.ListenAndServe(metricsAddress, nil)
}

func connectToParentChainRpcEndpoint(ctx *cli.Context, rpcUrl string) (*ethclient.Client, error) {
	rpcOptions := []rpc.ClientOption{}
	rpcToken := ctx.String(FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN)
	if rpcToken != "" {
		slog.Debug("Setting bearer token for the parent chain RPC endpoint.")
		rpcOptions = append(rpcOptions, rpc.WithHeader("Authorization", "Bearer "+rpcToken))
	}
	rpcClient, err := rpc.DialOptions(context.Background(), rpcUrl, rpcOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC URL %s: %w", rpcUrl, err)
	}
	client := ethclient.NewClient(rpcClient)
	slog.Info("connected to parent chain RPC URL", "url", rpcUrl)
	return client, nil
}

func connectToSubnetRpcEndpoint(rpcUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC URL %s: %w", rpcUrl, err)
	}
	slog.Info("connected to subnet RPC URL", "url", rpcUrl)
	return client, nil
}
