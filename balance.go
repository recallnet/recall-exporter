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
	PROM_LABEL_CHAIN_NAME = "chain_name"
	PROM_LABEL_ADDRESS    = "addr"
)

var (
	parentChainBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "validator_balance",
	}, []string{PROM_LABEL_CHAIN_NAME, PROM_LABEL_ADDRESS})
)

func runBalanceChecker(client *ethclient.Client, addressToCheck common.Address, chainName string, interval time.Duration) {
	logger := slog.With("chain", chainName)
	for {
		balance, err := client.BalanceAt(context.Background(), addressToCheck, nil)
		if err != nil {
			logger.Error("failed to get balance", "error", err)
		} else {
			logger.Debug("got balance", "balance", balance)
			parentChainBalance.WithLabelValues(chainName, addressToCheck.Hex()).Set(weiToPromEth(balance))
		}

		logger.Debug("sleeping", "duration", interval)
		time.Sleep(interval)
	}
}
