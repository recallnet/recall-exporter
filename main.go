package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	dummyCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dummy_counter",
		Help: "Dummy counter",
	})
)

func main() {
	go runDummyCounter()

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe("localhost:9010", nil); err != nil {
		panic(err)
	}
}

func runDummyCounter() {
	for {
		dummyCounter.Inc()
		time.Sleep(time.Second)
	}
}
