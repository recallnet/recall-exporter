package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lmittmann/tint"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const (
	FLAG_METRICS_ADDRESS                     = "metrics-address"
	FLAG_METRICS_PATH                        = "metrics-path"
	FLAG_PARENT_CHAIN_RPC_URL                = "parent-chain-rpc-url"
	FLAG_VALIDATOR_ADDRESS                   = "validator-address"
	FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL = "parent-chain-balance-check-interval"
	FLAG_SUBNET_RPC_URL                      = "subnet-rpc-url"
	FLAG_SUBNET_BALANCE_CHECK_INTERVAL       = "subnet-balance-check-interval"
)

var (
	parentChainRpcClient *ethclient.Client
	subnetRpcClient      *ethclient.Client
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run hoku-exporter",
				Action: commandRun,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  FLAG_METRICS_ADDRESS,
						Usage: "Address to expose metrics on",
						Value: "localhost:9010",
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
						Name:     FLAG_VALIDATOR_ADDRESS,
						Usage:    "Validator address",
						EnvVars:  []string{"VALIDATOR_ADDRESS"},
						Required: true,
					},
					&cli.DurationFlag{
						Name:  FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL,
						Usage: "How often the balance on the parent chain must be checked",
						Value: time.Minute,
					},
					&cli.StringFlag{
						Name:    FLAG_SUBNET_RPC_URL,
						Usage:   "Subnet RPC URL",
						EnvVars: []string{"SUBNET_RPC_URL"},
					},
					&cli.DurationFlag{
						Name:  FLAG_SUBNET_BALANCE_CHECK_INTERVAL,
						Usage: "How often the balance on the subnet must be checked",
						Value: time.Minute,
					},
				},
			},
		},
	}

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level: slog.LevelDebug,
	})))

	if err := app.Run(os.Args); err != nil {
		slog.Error("Failed to run app", "error", err)
	}
}

func commandRun(ctx *cli.Context) error {
	if err := setupRpcClients(ctx); err != nil {
		return err
	}

	validatorAddress := common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))
	go runBalanceChecker(parentChainRpcClient, validatorAddress, "parent", ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))
	go runBalanceChecker(subnetRpcClient, validatorAddress, "subnet", ctx.Duration(FLAG_SUBNET_BALANCE_CHECK_INTERVAL))

	metricsAddress := ctx.String(FLAG_METRICS_ADDRESS)
	metricsPath := ctx.String(FLAG_METRICS_PATH)
	metricsUrl := "http://" + metricsAddress + metricsPath
	slog.Info("Exporting metrics", "url", metricsUrl)

	http.Handle(metricsPath, promhttp.Handler())
	return http.ListenAndServe(metricsAddress, nil)
}

func setupRpcClients(ctx *cli.Context) error {
	rpcUrl := ctx.String(FLAG_PARENT_CHAIN_RPC_URL)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return fmt.Errorf("failed to dial RPC URL %s: %w", rpcUrl, err)
	}
	parentChainRpcClient = client

	rpcUrl = ctx.String(FLAG_SUBNET_RPC_URL)
	if rpcUrl != "" {
		client, err = ethclient.Dial(rpcUrl)
		if err != nil {
			return fmt.Errorf("failed to dial RPC URL %s: %w", rpcUrl, err)
		}
		subnetRpcClient = client
	}

	return nil
}
