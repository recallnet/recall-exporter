package main

import (
	"time"
)

func runMembershipChecker(ep *Endpoint, sleep time.Duration) {
	logger := ep.Logger.With("checker", "membership")
	for {
		ms, err := ep.GatewayCaller.GetCurrentMembership(nil)
		if err != nil {
			logger.Error("failed to get current membership", "error", err)
		}

		logger.Debug("got members", "count", len(ms.Validators))

		logger.Debug("sleeping", "duration", sleep)
		time.Sleep(sleep)
	}
}
