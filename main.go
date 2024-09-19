package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gochain/web3"
	"github.com/lmittmann/tint"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const (
	FLAG_METRICS_ADDRESS                     = "metrics-address"
	FLAG_METRICS_PATH                        = "metrics-path"
	FLAG_PARENT_CHAIN_RPC_URL                = "parent-chain-rpc-url"
	FLAG_PARENT_CHAIN_ADDRESS                = "parent-chain-address"
	FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL = "parent-chain-balance-check-interval"
	FLAG_SUBNET_ADDRESS                      = "subnet-address"
	FLAG_SUBNET_BALANCE_CHECK_INTERVAL       = "subnet-balance-check-interval"
)

var parentChainRpcClient web3.Client

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
						Name:     FLAG_PARENT_CHAIN_ADDRESS,
						Usage:    "Parent chain address",
						EnvVars:  []string{"PARENT_CHAIN_ADDRESS"},
						Required: true,
					},
					&cli.DurationFlag{
						Name:  FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL,
						Usage: "How often the balance on the parent chain must be checked",
						Value: time.Minute,
					},
					&cli.StringFlag{
						Name:     FLAG_SUBNET_ADDRESS,
						Usage:    "Subnet address",
						EnvVars:  []string{"SUBNET_ADDRESS"},
						Required: true,
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
	client, err := web3.Dial(ctx.String(FLAG_PARENT_CHAIN_RPC_URL))
	if err != nil {
		return fmt.Errorf("failed to dial RPC URL: %w", err)
	}
	parentChainRpcClient = client

	go runParentChainBalanceChecker(ctx.String(FLAG_PARENT_CHAIN_ADDRESS), ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))

	metricsAddress := ctx.String(FLAG_METRICS_ADDRESS)
	metricsPath := ctx.String(FLAG_METRICS_PATH)
	metricsUrl := "http://" + metricsAddress + metricsPath
	slog.Info("Exporting metrics", "url", metricsUrl)

	http.Handle(metricsPath, promhttp.Handler())
	return http.ListenAndServe(metricsAddress, nil)
}
