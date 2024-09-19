package main

import (
	"math/big"
)

// weiToPromEth scales Wei to Eth and converts it to float to be usable as prometheus metric value.
func weiToPromEth(tokenValue *big.Int) float64 {
	ethValue := new(big.Float).SetInt(tokenValue)
	ethValue.Quo(ethValue, big.NewFloat(1e18))
	result, _ := ethValue.Float64()
	return result
}
