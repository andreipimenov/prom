package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var method = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Request method count",
	},
	[]string{"method"},
)

var duration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration",
		Help: "Request duration",
	},
	[]string{"method"},
)

func init() {
	rand.Seed(time.Now().UnixNano())

	prometheus.MustRegister(method)
	prometheus.MustRegister(duration)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		duration.WithLabelValues(r.Method).Observe(time.Since(start).Seconds())
		method.WithLabelValues(r.Method).Inc()
	}()

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.Method))
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	log.Printf("Server is listening on :%v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
