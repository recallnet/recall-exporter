package main

import (
	"fmt"
	"log/slog"
	"maps"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func newBottomupCheckpointChecker(ep *ParentChainEndpoint, addressToCheck common.Address) JobFunc {
	labels := maps.Clone(ep.Labels)
	labels[PROM_LABEL_ADDRESS] = addressToCheck.Hex()

	checkpointHeight := promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "last_bottomup_checkpoint_height",
		ConstLabels: labels,
	})

	return func(logger *slog.Logger) error {
		height, err := ep.SubnetCaller.LastBottomUpCheckpointHeight(nil)
		if err != nil {
			return fmt.Errorf("failed to get last bottomup checkpoint height: %w", err)
		}

		heightF, _ := height.Float64()
		checkpointHeight.Set(heightF)
		logger.Debug("got last bottomup checkpoint", "height", heightF)

		return nil
	}
}
