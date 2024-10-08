package main

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	PROM_LABEL_JOB_NAME   = "job"
	PROM_LABEL_JOB_STATUS = "status"
)

var (
	counterJobRun = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "job_run",
	}, []string{PROM_LABEL_JOB_NAME, PROM_LABEL_JOB_STATUS})
)

type JobFunc func(*slog.Logger) error

func StartJob(name string, job JobFunc, interval time.Duration, parentLogger *slog.Logger) {
	logger := parentLogger.With("job", name)

	jobWrapper := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic", "reason", r, "stack", debug.Stack())
				err = fmt.Errorf("panic: %v", r)
			}
		}()
		return job(logger)
	}
	go func() {
		logger.Info("starting job")
		for {
			status := "success"
			err := jobWrapper()
			if err != nil {
				logger.Error("failed to run job", "error", err)
				status = "failure"
			}
			counterJobRun.WithLabelValues(name, status).Inc()

			logger.Debug("sleeping", "duration", interval)
			time.Sleep(interval)
		}
	}()
}
