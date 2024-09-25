package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	PROM_LABEL_NETWORK_NAME = "network_name"
	PROM_LABEL_CHAIN_ID     = "chain_id"
	PROM_LABEL_ADDRESS      = "addr"
)

var (
	parentChainBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "validator_balance",
	}, []string{PROM_LABEL_NETWORK_NAME, PROM_LABEL_CHAIN_ID, PROM_LABEL_ADDRESS})
)

func runBalanceChecker(client *ethclient.Client, addressToCheck common.Address, chainName string, interval time.Duration) {
	logger := slog.With("chain", chainName)
	chainId, err := client.ChainID(context.Background())
	var chainIdStr string
	if err == nil {
		chainIdStr = chainId.String()
	} else {
		logger.Error("failed to get chain ID", "error", err)
	}
	for {
		balance, err := client.BalanceAt(context.Background(), addressToCheck, nil)
		if err != nil {
			logger.Error("failed to get balance", "error", err)
		} else {
			logger.Debug("got balance", "balance", balance)
			val, _ := balance.Float64()
			parentChainBalance.WithLabelValues(chainName, chainIdStr, addressToCheck.Hex()).Set(val)
		}

		logger.Debug("sleeping", "duration", interval)
		time.Sleep(interval)
	}
}
