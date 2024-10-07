package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func runMembershipChecker(ep *SubnetEndpoint, sleep time.Duration) {
	logger := ep.Logger.With("checker", "membership")
	gaugeMember := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "subnet_validator_weight",
		ConstLabels: ep.Labels,
	}, []string{PROM_LABEL_ADDRESS})

	knownMembers := map[string]int{}
	currentRunIteration := 0

	for {
		currentRunIteration++

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
			knownMembers[addr] = currentRunIteration
			weight, _ := validator.Weight.Float64()
			gaugeMember.WithLabelValues(addr).Set(weight)
		}

		// Delete old validators
		for validatorAddress, lastSeenAtRunIteration := range knownMembers {
			if lastSeenAtRunIteration < currentRunIteration {
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
