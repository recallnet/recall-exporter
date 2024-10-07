package main

import (
	"maps"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func runBottomupCheckpointChecker(ep *ParentChainEndpoint, addressToCheck common.Address, interval time.Duration) {
	logger := ep.Logger.With("checker", "bottomup-checkpoint")

	labels := maps.Clone(ep.Labels)
	labels[PROM_LABEL_ADDRESS] = addressToCheck.Hex()

	checkpointHeight := promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "last_bottomup_checkpoint_height",
		ConstLabels: labels,
	})

	for {
		height, err := ep.SubnetCaller.LastBottomUpCheckpointHeight(nil)
		if err != nil {
			logger.Error("failed to get last bottomup checkpoint height", "error", err)
		}

		heightF, _ := height.Float64()
		checkpointHeight.Set(heightF)
		logger.Debug("got last bottomup checkpoint", "height", heightF)

		logger.Debug("sleeping", "duration", interval)
		time.Sleep(interval)
	}

}
