package main

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func newBottomupCheckpointChecker(ep *ParentChainEndpoint) JobFunc {
	checkpointHeight := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   PROM_NAMESPACE_RECALL,
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

var collateralChecker *CollateralChecker

type CollateralChecker struct {
	endpoint *ParentChainEndpoint

	// currentSubnetValidatorAddresses contains currently known validator addresses pointing to the iteration counter.
	// This info allows to determine what addresses were removed from the subnet when updating the subnet membership.
	currentSubnetValidatorAddresses map[common.Address]int
	validatorAddressesMux           sync.Mutex
	validatorAddressesIteration     int

	gaugeValidatorCollateral *prometheus.GaugeVec
	gaugeTotalCollateral     prometheus.Gauge
}

func NewCollateralChecker(ep *ParentChainEndpoint) *CollateralChecker {
	return &CollateralChecker{
		endpoint:                        ep,
		currentSubnetValidatorAddresses: map[common.Address]int{},

		gaugeValidatorCollateral: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   PROM_NAMESPACE_RECALL,
			Name:        "validator_collateral",
			ConstLabels: ep.ConstLabels(),
		}, []string{PROM_LABEL_ADDRESS}),
		gaugeTotalCollateral: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace:   PROM_NAMESPACE_RECALL,
			Name:        "collateral_confirmed_total",
			Help:        "Total amount of confirmed collateral across all validators.",
			ConstLabels: ep.ConstLabels(),
		}),
	}
}

// setValidatorAddresses sets addresses of all known validators in the subnet.
func (c *CollateralChecker) setValidatorAddresses(addresses []common.Address, logger *slog.Logger) {
	c.validatorAddressesMux.Lock()
	defer c.validatorAddressesMux.Unlock()

	initialValidatorCount := len(c.currentSubnetValidatorAddresses)
	c.validatorAddressesIteration++

	logger.Debug("setting validator addresses in collateral checker")
	for _, addr := range addresses {
		if c.currentSubnetValidatorAddresses[addr] == 0 {
			logger.Info("adding new validator", "addr", addr.Hex())
		}
		c.currentSubnetValidatorAddresses[addr] = c.validatorAddressesIteration
	}

	// Delete old validator addresses and their metrics.
	for addr, iteration := range c.currentSubnetValidatorAddresses {
		if iteration < c.validatorAddressesIteration {
			addrHex := addr.Hex()
			logger.Info("removing validator", "addr", addrHex)
			c.gaugeValidatorCollateral.DeleteLabelValues(addrHex)
			delete(c.currentSubnetValidatorAddresses, addr)
		}
	}

	if initialValidatorCount == 0 {
		logger.Info("initial validator count was 0 -> triggering collateral check")
		go c.updateAllCollateralMetrics(logger.With("context", "initial-collateral-check"))
	}
}

// updateAllCollateralMetrics updates collateral metrics for every known validator and the total subnet collateral.
func (c *CollateralChecker) updateAllCollateralMetrics(logger *slog.Logger) error {
	c.validatorAddressesMux.Lock()
	defer c.validatorAddressesMux.Unlock()

	logger.Debug("collecting collateral for all validators", "validator-count", len(c.currentSubnetValidatorAddresses))
	c.updateTotalCollateralMetric(logger)
	for addr := range c.currentSubnetValidatorAddresses {
		err := c.updateValidatorCollateralMetric(addr, logger)
		if err != nil {
			return err
		}
	}
	return nil
}

// updateValidatorCollateralMetric updates the collateral metric for the given validator address.
func (c *CollateralChecker) updateValidatorCollateralMetric(address common.Address, logger *slog.Logger) error {
	addrHex := address.Hex()

	collateral, err := c.endpoint.SubnetCaller.GetTotalValidatorCollateral(nil, address)
	if err != nil {
		return fmt.Errorf("failed to get validator collateral: %w", err)
	}

	value, _ := collateral.Float64()
	c.gaugeValidatorCollateral.WithLabelValues(addrHex).Set(value)
	logger.Debug("got collateral", "address", addrHex, "collateral", value)

	return nil
}

func (c *CollateralChecker) updateTotalCollateralMetric(logger *slog.Logger) error {
	collateral, err := c.endpoint.SubnetCaller.GetTotalConfirmedCollateral(nil)
	if err != nil {
		return fmt.Errorf("failed to get total collateral: %w", err)
	}

	value, _ := collateral.Float64()
	c.gaugeTotalCollateral.Set(value)
	logger.Debug("got total collateral", "collateral", value)

	return nil
}
