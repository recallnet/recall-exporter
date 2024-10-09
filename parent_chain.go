package main

import (
	"fmt"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func newBottomupCheckpointChecker(ep *ParentChainEndpoint) JobFunc {
	checkpointHeight := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "last_bottomup_checkpoint_height",
		ConstLabels: ep.ConstLabels(),
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
