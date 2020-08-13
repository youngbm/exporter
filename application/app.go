package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	timeCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "myTimeCounter",
			Help: "This is my time Counter.",
		})
	randomGague = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "myRandomGague",
			Help: "This is my Randon Gauge.",
		})
)

func myMetrics() {
	go func() {
		for {
			timeCounter.Add(5)
			randomGague.Set(rand.Float64() * float64(rand.Int()))
			time.Sleep(time.Microsecond * 100)
		}
	}()
}

func main() {
	myMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
