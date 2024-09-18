package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

var (
	dummyCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dummy_counter",
		Help: "Dummy counter",
	})
)

var logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

const (
	FLAG_METRICS_ADDRESS = "metrics-address"
	FLAG_METRICS_PATH    = "metrics-path"
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run hoku-exporter",
				Action: commandRun,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  FLAG_METRICS_ADDRESS,
						Usage: "Address to expose metrics on",
						Value: "localhost:9010",
					},
					&cli.StringFlag{
						Name:  FLAG_METRICS_PATH,
						Usage: "Metrics path",
						Value: "/metrics",
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logger.Error("Failed to run app", "error", err)
	}
}

func commandRun(ctx *cli.Context) error {
	go runDummyCounter()

	metricsAddress := ctx.String(FLAG_METRICS_ADDRESS)
	metricsPath := ctx.String(FLAG_METRICS_PATH)

	logger.Info("Exporting metrics", "url", "http://"+metricsAddress+metricsPath)

	http.Handle(metricsPath, promhttp.Handler())
	return http.ListenAndServe(metricsAddress, nil)
}

func runDummyCounter() {
	for {
		dummyCounter.Inc()
		time.Sleep(time.Second)
	}
}
