package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	PROM_LABEL_NETWORK_NAME = "network_name"
	PROM_LABEL_CHAIN_ID     = "chain_id"
	PROM_LABEL_ADDRESS      = "addr"
)

func newBalanceCheckerJob(ep *Endpoint, addressToCheck common.Address) JobFunc {
	gaugeBalance := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "validator_balance",
		ConstLabels: ep.Labels(PROM_LABEL_ADDRESS, addressToCheck.Hex()),
	})

	return func(logger *slog.Logger) error {
		balance, err := ep.Client.BalanceAt(context.Background(), addressToCheck, nil)
		if err != nil {
			return fmt.Errorf("failed to get balance: %w", err)
		} else {
			logger.Debug("got balance", "balance", balance)
			val, _ := balance.Float64()
			gaugeBalance.Set(val)
		}
		return nil
	}
}
