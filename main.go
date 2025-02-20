package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/tint"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const (
	FLAG_METRICS_ADDRESS                                 = "metrics-address"
	FLAG_METRICS_PATH                                    = "metrics-path"
	FLAG_PARENT_CHAIN_RPC_URL                            = "parent-chain-rpc-url"
	FLAG_PARENT_CHAIN_RPC_BEARER_TOKEN                   = "parent-chain-rpc-bearer-token"
	FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL             = "parent-chain-balance-check-interval"
	FLAG_PARENT_CHAIN_NETWORK_NAME                       = "parent-chain-network-name"
	FLAG_PARENT_CHAIN_SUBNET_CONTRACT_ADDRESS            = "parent-chain-subnet-contract-address"
	FLAG_PARENT_CHAIN_BOTTOMUP_CHECKPOINT_CHECK_INTERVAL = "parent-chain-bottomup-checkpoint-check-interval"
	FLAG_VALIDATOR_ADDRESS                               = "validator-address"
	FLAG_RELAYER_ADDRESS                                 = "relayer-address"
	FLAG_FAUCET_ADDRESS                                  = "faucet-address"
	FLAG_SUBNET_RPC_URL                                  = "subnet-rpc-url"
	FLAG_SUBNET_NETWORK_NAME                             = "subnet-network-name"
	FLAG_SUBNET_BALANCE_CHECK_INTERVAL                   = "subnet-balance-check-interval"
	FLAG_SUBNET_GATEWAY_ADDRESS                          = "subnet-gateway-address"
	FLAG_SUBNET_MEMBERSHIP_CHECK_INTERVAL                = "subnet-gateway-check-interval"
	FLAG_SUBNET_BLOB_MANAGER_CONTRACT_ADDRESS            = "subnet-blob-manager-contract-address"
	FLAG_SUBNET_STATS_CHECK_INTERVAL                     = "subnet-stats-check-interval"
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
				Usage:  "Run recall-exporter",
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
					&cli.DurationFlag{
						Name:    FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL,
						Usage:   "How often the balance on the parent chain must be checked",
						Value:   time.Minute,
						EnvVars: []string{"PARENT_CHAIN_BALANCE_CHECK_INTERVAL"},
					},
					&cli.StringFlag{
						Name:    FLAG_PARENT_CHAIN_NETWORK_NAME,
						Usage:   "Parent chain network name",
						EnvVars: []string{"PARENT_CHAIN_NETWORK_NAME"},
						Value:   "parent",
					},
					&cli.StringFlag{
						Name:     FLAG_PARENT_CHAIN_SUBNET_CONTRACT_ADDRESS,
						Usage:    "Parent chain subnet contract address",
						EnvVars:  []string{"PARENT_CHAIN_SUBNET_CONTRACT_ADDRESS"},
						Required: true,
					},
					&cli.DurationFlag{
						Name:    FLAG_PARENT_CHAIN_BOTTOMUP_CHECKPOINT_CHECK_INTERVAL,
						Usage:   "How often the bottomup checkpoint must be checked",
						Value:   time.Minute,
						EnvVars: []string{"PARENT_CHAIN_BOTTOMUP_CHECKPOINT_CHECK_INTERVAL"},
					},
					&cli.StringFlag{
						Name:     FLAG_VALIDATOR_ADDRESS,
						Usage:    "Validator address",
						EnvVars:  []string{"VALIDATOR_ADDRESS"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    FLAG_RELAYER_ADDRESS,
						Usage:   "Subnet relayer address",
						EnvVars: []string{"RELAYER_ADDRESS"},
					},
					&cli.StringFlag{
						Name:    FLAG_FAUCET_ADDRESS,
						Usage:   "Faucet address",
						EnvVars: []string{"FAUCET_ADDRESS"},
					},
					&cli.StringFlag{
						Name:    FLAG_SUBNET_RPC_URL,
						Usage:   "Subnet RPC URL",
						EnvVars: []string{"SUBNET_RPC_URL"},
					},
					&cli.StringFlag{
						Name:    FLAG_SUBNET_NETWORK_NAME,
						Usage:   "Subnet RPC network name",
						EnvVars: []string{"SUBNET_RPC_NETWORK_NAME"},
						Value:   "subnet",
					},
					&cli.DurationFlag{
						Name:    FLAG_SUBNET_BALANCE_CHECK_INTERVAL,
						Usage:   "How often the balance on the subnet must be checked",
						Value:   time.Minute,
						EnvVars: []string{"SUBNET_BALANCE_CHECK_INTERVAL"},
					},
					&cli.StringFlag{
						Name:     FLAG_SUBNET_GATEWAY_ADDRESS,
						Usage:    "Subnet gateway address",
						EnvVars:  []string{"SUBNET_GATEWAY_ADDRESS"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    FLAG_SUBNET_BLOB_MANAGER_CONTRACT_ADDRESS,
						Usage:   "Blob manager contract address",
						EnvVars: []string{"SUBNET_BLOB_MANAGER_CONTRACT_ADDRESS"},
					},
					&cli.DurationFlag{
						Name:    FLAG_SUBNET_STATS_CHECK_INTERVAL,
						Usage:   "Subnet stats check interval",
						EnvVars: []string{"SUBNET_STATS_CHECK_INTERVAL"},
					},
					&cli.DurationFlag{
						Name:    FLAG_SUBNET_MEMBERSHIP_CHECK_INTERVAL,
						Usage:   "Subnet membership check interval",
						EnvVars: []string{"SUBNET_MEMBERSHIP_CHECK_INTERVAL"},
					},
				},
			},
		},
	}

	setupLogging()
	if err := app.Run(os.Args); err != nil {
		slog.Error("Failed to run app", "error", err)
	}
}

func commandRun(ctx *cli.Context) error {
	slog.Info("running recall-exporter", "git-commit", GitCommit, "build-time", BuildTime)

	subnetEP, err := newSubnetEndpoint(ctx)
	if err != nil {
		return err
	}

	parentChainEP, err := newParentChainEndpoint(ctx)
	if err != nil {
		return err
	}

	startSubnetJobs(subnetEP, ctx)
	startParentChainJobs(parentChainEP, ctx)

	metricsAddress := ctx.String(FLAG_METRICS_ADDRESS)
	metricsPath := ctx.String(FLAG_METRICS_PATH)
	metricsUrl := "http://" + metricsAddress + metricsPath
	slog.Info("Exporting metrics", "url", metricsUrl)

	http.Handle(metricsPath, promhttp.Handler())
	return http.ListenAndServe(metricsAddress, nil)
}

func setupLogging() {
	val, valSet := os.LookupEnv("GO_LOG")
	logLevel := slog.LevelInfo
	var err error
	if valSet {
		err = logLevel.UnmarshalText([]byte(val))
	}
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      logLevel,
		TimeFormat: time.RFC3339,
	})))
	if err != nil {
		slog.Warn("invalid GO_LOG value, ", "GO_LOG", val, "error", err)
	}
}

func validatorAddress(ctx *cli.Context) common.Address {
	return common.HexToAddress(ctx.String(FLAG_VALIDATOR_ADDRESS))
}

const (
	addressOwnerValidator     = "validator"
	addressOwnerSubnetRelayer = "relayer"
	addressOwnerFaucet        = "faucet"
)

func startSubnetJobs(ep *SubnetEndpoint, ctx *cli.Context) {
	network := ctx.String(FLAG_SUBNET_NETWORK_NAME)
	StartJob("balance-validator", network, newBalanceCheckerJob(ep.Endpoint, addressOwnerValidator, validatorAddress(ctx)), ctx.Duration(FLAG_SUBNET_BALANCE_CHECK_INTERVAL))
	StartJob("membership", network, newMembershipChecker(ep), ctx.Duration(FLAG_SUBNET_MEMBERSHIP_CHECK_INTERVAL))

	faucetAddress := ctx.String(FLAG_FAUCET_ADDRESS)
	if faucetAddress != "" {
		StartJob("balance-faucet", network,
			newBalanceCheckerJob(ep.Endpoint, addressOwnerFaucet, common.HexToAddress(faucetAddress)),
			ctx.Duration(FLAG_SUBNET_BALANCE_CHECK_INTERVAL))
	}

	if ep.BlobsCaller != nil {
		StartJob("subnet-stats", network, newSubnetStatsChecker(ep), ctx.Duration(FLAG_SUBNET_STATS_CHECK_INTERVAL))
	}
}

func startParentChainJobs(ep *ParentChainEndpoint, ctx *cli.Context) {
	network := ctx.String(FLAG_PARENT_CHAIN_NETWORK_NAME)
	StartJob("balance-validator", network, newBalanceCheckerJob(ep.Endpoint, addressOwnerValidator, validatorAddress(ctx)), ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))
	StartJob("balance-erc20", network, newErc20TokenBalanceCheckerJob(ep, addressOwnerValidator, validatorAddress(ctx)), ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))

	relayerAddress := ctx.String(FLAG_RELAYER_ADDRESS)
	if relayerAddress != "" {
		StartJob("balance-relayer", network,
			newBalanceCheckerJob(ep.Endpoint, addressOwnerSubnetRelayer, common.HexToAddress(relayerAddress)),
			ctx.Duration(FLAG_PARENT_CHAIN_BALANCE_CHECK_INTERVAL))
	}

	StartJob("bottomup-checkpoint", network, newBottomupCheckpointChecker(ep), ctx.Duration(FLAG_PARENT_CHAIN_BOTTOMUP_CHECKPOINT_CHECK_INTERVAL))

	collateralChecker = NewCollateralChecker(ep)
	StartJob("collateral", network, collateralChecker.updateAllCollateralMetrics, ctx.Duration(FLAG_SUBNET_MEMBERSHIP_CHECK_INTERVAL))
}
