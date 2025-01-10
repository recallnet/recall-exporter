package main

import (
	"fmt"
	"log/slog"
	"math/big"

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

		collateralChecker.setValidatorAddresses(validatorAddresses, logger)

		return nil
	}
}

const (
	PROM_LABEL_TYPE = "type"
)

func bigIntToFloat(b *big.Int) float64 {
	v, _ := b.Float64()
	return v
}

func newSubnetStatsChecker(ep *SubnetEndpoint) JobFunc {
	gaugeCapacity := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_capacity",
		ConstLabels: ep.ConstLabels(),
	}, []string{PROM_LABEL_TYPE})
	gaugeCredit := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_credit",
		ConstLabels: ep.ConstLabels(),
	}, []string{PROM_LABEL_TYPE})
	gaugeBalance := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_balance",
		ConstLabels: ep.ConstLabels(),
	})
	gaugeAccounts := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_accounts",
		ConstLabels: ep.ConstLabels(),
	})
	gaugeBlobs := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_blobs",
		ConstLabels: ep.ConstLabels(),
	})
	gaugeResolving := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_resolving",
		ConstLabels: ep.ConstLabels(),
	})
	gaugeTokenCreditRate := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_HOKU,
		Name:        "subnet_stats_token_credit_rate",
		ConstLabels: ep.ConstLabels(),
	})

	return func(logger *slog.Logger) error {
		stats, err := ep.BlobsCaller.GetSubnetStats(nil)
		if err != nil {
			return fmt.Errorf("failed to get subnet stats: %w", err)
		}

		gaugeBalance.Set(bigIntToFloat(stats.Balance))
		gaugeBlobs.Set(float64(stats.NumBlobs))
		gaugeResolving.Set(float64(stats.NumResolving))
		gaugeAccounts.Set(float64(stats.NumAccounts))
		gaugeTokenCreditRate.Set(bigIntToFloat(stats.TokenCreditRate))

		gaugeCapacity.WithLabelValues("free").Set(float64(stats.CapacityFree))
		gaugeCapacity.WithLabelValues("used").Set(float64(stats.CapacityUsed))

		gaugeCredit.WithLabelValues("sold").Set(bigIntToFloat(stats.CreditSold))
		gaugeCredit.WithLabelValues("committed").Set(bigIntToFloat(stats.CreditCommitted))
		gaugeCredit.WithLabelValues("debited").Set(bigIntToFloat(stats.CreditDebited))

		logger.Debug("updated stats")
		return nil
	}
}
