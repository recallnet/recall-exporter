package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func runMembershipChecker(ep *Endpoint, sleep time.Duration) {
	logger := ep.Logger.With("checker", "membership")
	gaugeMember := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "subnet_validator",
		ConstLabels: ep.Labels,
	}, []string{PROM_LABEL_ADDRESS})

	knownMembers := map[string]int{}
	run := 0

	for {
		run++

		ms, err := ep.GatewayCaller.GetCurrentMembership(nil)
		if err != nil {
			logger.Error("failed to get current membership", "error", err)
		}

		logger.Debug("got members", "count", len(ms.Validators))
		for _, validator := range ms.Validators {
			addr := validator.Addr.Hex()
			if knownMembers[addr] == 0 {
				logger.Info("add new member", "addr", addr)
			}
			knownMembers[addr] = run
			gaugeMember.WithLabelValues(addr).Set(1)
		}

		// Delete old validators
		for validatorAddress, r := range knownMembers {
			if r < run {
				// This is an old validator that we should remove from the metric set.
				gaugeMember.DeleteLabelValues(validatorAddress)
				delete(knownMembers, validatorAddress)
				logger.Info("remove old member", "addr", validatorAddress)
			}
		}

		logger.Debug("sleeping", "duration", sleep)
		time.Sleep(sleep)
	}
}
