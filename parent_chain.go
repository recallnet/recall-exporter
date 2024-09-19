package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	parentChainBalance = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "paraent_chain_balance",
		Help: "Validator balance on the parent chain",
	})
)

func runParentChainBalanceChecker(parentChainAddress string, interval time.Duration) {
	logger := slog.With("comp", "parent-balance")
	for {
		balance, err := parentChainRpcClient.GetBalance(context.Background(), parentChainAddress, nil)
		if err != nil {
			logger.Error("failed to get balance", "error", err)
		} else {
			logger.Debug("got balance", "balance", balance)
			parentChainBalance.Set(float64(balance.Int64()))
		}

		logger.Debug("sleeping", "duration", interval)
		time.Sleep(interval)
	}
}
