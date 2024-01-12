package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			log.Println("Sending metrics to Prometheus")
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	// Simulate a many goroutines processing
	// go_goroutines in metrics data
	timer := time.NewTimer(1 * time.Minute)

	for {
		select {
		case <-timer.C:
			// if timer expired, stop generate goroutines
			log.Println("Timer expired")
			return
		default:
			go func() {
				log.Println("Process with goroutine")
				time.Sleep(30 * time.Second)
			}()
		}
		time.Sleep(1 * time.Second)
	}
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of processed events",
	})
)

func main() {
	prometheus.Register(opsProcessed)

	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
