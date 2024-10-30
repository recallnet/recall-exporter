package main

import (
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func newMembershipChecker(ep *SubnetEndpoint) JobFunc {
	gaugeMember := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "validator_weight",
		ConstLabels: ep.ConstLabels(),
	}, []string{PROM_LABEL_ADDRESS})

	knownMembers := map[string]int{}
	currentRunIteration := 0

	return func(logger *slog.Logger) error {
		currentRunIteration++

		ms, err := ep.GatewayCaller.GetCurrentMembership(nil)
		if err != nil {
			return fmt.Errorf("failed to get current membership: %w", err)
		}

		logger.Debug("got members", "count", len(ms.Validators))
		validatorAddresses := []common.Address{}
		for _, validator := range ms.Validators {
			validatorAddresses = append(validatorAddresses, validator.Addr)
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

		validatorInfoCollector.setSubnetMembership(validatorAddresses, logger)

		return nil
	}
}
