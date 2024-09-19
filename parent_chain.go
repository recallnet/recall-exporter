package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	PROM_LABEL_ADDRESS = "address"
)

var (
	parentChainBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "parent_chain_balance",
		Help: "Validator balance on the parent chain",
	}, []string{PROM_LABEL_ADDRESS})
)

func runParentChainBalanceChecker(parentChainAddress string, interval time.Duration) {
	logger := slog.With("comp", "parent-chain-balance-checker")
	parsedParentAddress := common.HexToAddress(parentChainAddress)

	for {
		balance, err := parentChainRpcClient.BalanceAt(context.Background(), parsedParentAddress, nil)
		if err != nil {
			logger.Error("failed to get balance", "error", err)
		} else {
			logger.Debug("got balance", "balance", balance)
			parentChainBalance.WithLabelValues(parentChainAddress).Set(weiToPromEth(balance))
		}

		logger.Debug("sleeping", "duration", interval)
		time.Sleep(interval)
	}
}
