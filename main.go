package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var (
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
	)
)

func Endpoint(w http.ResponseWriter, r *http.Request) {
	// Increment the counter for each HTTP request.
	httpRequestsTotal.Inc()
	log.Println("Querying Endpoint, httpRequestsTotal: ", httpRequestsTotal)
	log.Println(r.Header)
	// Handle the HTTP request as normal.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

func main() {
	// Register the metric with the Prometheus client.
	prometheus.MustRegister(httpRequestsTotal)

	mux := http.NewServeMux()

	mux.HandleFunc("/", Endpoint)
	port := os.Getenv("PORT")
	log.Println("Listening app on port: ", port)

	// Expose the Prometheus metrics on an HTTP endpoint.
	mux.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server.
	http.ListenAndServe(":"+port, mux)
}
