package main

import (
	"context"
	"maps"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	PROM_LABEL_NETWORK_NAME = "network_name"
	PROM_LABEL_CHAIN_ID     = "chain_id"
	PROM_LABEL_ADDRESS      = "addr"
)

func runBalanceChecker(ep *Endpoint, addressToCheck common.Address, interval time.Duration) {
	labels := maps.Clone(ep.Labels)
	labels[PROM_LABEL_ADDRESS] = addressToCheck.Hex()

	logger := ep.Logger.With("checker", "balance")

	gaugeBalance := promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "validator_balance",
		ConstLabels: labels,
	})
	for {
		balance, err := ep.Client.BalanceAt(context.Background(), addressToCheck, nil)
		if err != nil {
			logger.Error("failed to get balance", "error", err)
		} else {
			logger.Debug("got balance", "balance", balance)
			val, _ := balance.Float64()
			gaugeBalance.Set(val)
		}

		logger.Debug("sleeping", "duration", interval)
		time.Sleep(interval)
	}
}
