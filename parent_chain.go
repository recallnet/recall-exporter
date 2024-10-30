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

var validatorInfoCollector *ValidatorInfoCollector

type ValidatorInfoCollector struct {
	endpoint *ParentChainEndpoint

	// Address -> check iteration
	currentSubnetValidatorAddresses map[common.Address]int
	addressesMux                    sync.Mutex
	validatorAddressesIteration     int

	gaugeValidatorCollateralConfirmed *prometheus.GaugeVec
	gaugeValidatorCollateralTotal     *prometheus.GaugeVec
}

func NewValidatorInfoCollector(ep *ParentChainEndpoint) *ValidatorInfoCollector {
	return &ValidatorInfoCollector{
		endpoint:                        ep,
		currentSubnetValidatorAddresses: map[common.Address]int{},

		gaugeValidatorCollateralConfirmed: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   PROM_NAMESPACE_HOKU,
			Name:        "validator_collateral_confirmed",
			ConstLabels: ep.ConstLabels(),
		}, []string{PROM_LABEL_ADDRESS}),
		gaugeValidatorCollateralTotal: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   PROM_NAMESPACE_HOKU,
			Name:        "validator_collateral_total",
			ConstLabels: ep.ConstLabels(),
		}, []string{PROM_LABEL_ADDRESS}),
	}
}

func (c *ValidatorInfoCollector) setSubnetMembership(addresses []common.Address, logger *slog.Logger) {
	c.addressesMux.Lock()
	defer c.addressesMux.Unlock()

	c.validatorAddressesIteration++

	for _, addr := range addresses {
		if c.currentSubnetValidatorAddresses[addr] == 0 {
			logger.Info("adding new validator", "addr", addr.Hex())
		}
		c.currentSubnetValidatorAddresses[addr] = c.validatorAddressesIteration
	}

	// delete old
	for addr, iteration := range c.currentSubnetValidatorAddresses {
		if iteration < c.validatorAddressesIteration {
			addrHex := addr.Hex()
			logger.Info("removing validator", "addr", addrHex)
			c.gaugeValidatorCollateralConfirmed.DeleteLabelValues(addrHex)
			c.gaugeValidatorCollateralTotal.DeleteLabelValues(addrHex)
			delete(c.currentSubnetValidatorAddresses, addr)
		}
	}
}

func (c *ValidatorInfoCollector) collectValidatorsInfos(logger *slog.Logger) error {
	c.addressesMux.Lock()
	defer c.addressesMux.Unlock()

	logger.Debug("collecting infos for all validators", "validator-count", len(c.currentSubnetValidatorAddresses))
	for addr := range c.currentSubnetValidatorAddresses {
		err := c.publishValidatorMetrics(addr, logger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ValidatorInfoCollector) publishValidatorMetrics(address common.Address, logger *slog.Logger) error {
	addrHex := address.Hex()

	info, err := c.endpoint.SubnetCaller.GetValidator(nil, address)
	if err != nil {
		return fmt.Errorf("failed to get validator info: %w", err)
	}

	value, _ := info.ConfirmedCollateral.Float64()
	c.gaugeValidatorCollateralConfirmed.WithLabelValues(addrHex).Set(value)
	logger.Debug("got collateral confirmed", "address", addrHex, "collateral", value)

	value, _ = info.TotalCollateral.Float64()
	c.gaugeValidatorCollateralTotal.WithLabelValues(addrHex).Set(value)
	logger.Debug("got collateral total", "address", addrHex, "collateral", value)

	return nil
}
