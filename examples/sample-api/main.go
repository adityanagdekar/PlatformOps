package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	startedAt      = time.Now()
	healthRequests uint64
	readyRequests  uint64
	metricRequests uint64
	totalRequests  uint64
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", track("health", healthHandler))
	mux.HandleFunc("GET /ready", track("ready", readyHandler))
	mux.HandleFunc("GET /metrics", track("metrics", metricsHandler))

	addr := ":" + envOrDefault("PORT", "8080")
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("sample-api listening on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeText(w, http.StatusOK, "ok\n")
}

func readyHandler(w http.ResponseWriter, _ *http.Request) {
	if envOrDefault("READY", "true") != "true" {
		writeText(w, http.StatusServiceUnavailable, "not ready\n")
		return
	}

	writeText(w, http.StatusOK, "ready\n")
}

func metricsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	uptimeSeconds := time.Since(startedAt).Seconds()
	fmt.Fprintf(w, "# HELP sample_api_uptime_seconds Seconds since the sample API started.\n")
	fmt.Fprintf(w, "# TYPE sample_api_uptime_seconds gauge\n")
	fmt.Fprintf(w, "sample_api_uptime_seconds %.0f\n", uptimeSeconds)
	fmt.Fprintf(w, "# HELP sample_api_http_requests_total Total HTTP requests handled by endpoint.\n")
	fmt.Fprintf(w, "# TYPE sample_api_http_requests_total counter\n")
	fmt.Fprintf(w, "sample_api_http_requests_total{path=\"/health\"} %d\n", atomic.LoadUint64(&healthRequests))
	fmt.Fprintf(w, "sample_api_http_requests_total{path=\"/ready\"} %d\n", atomic.LoadUint64(&readyRequests))
	fmt.Fprintf(w, "sample_api_http_requests_total{path=\"/metrics\"} %d\n", atomic.LoadUint64(&metricRequests))
	fmt.Fprintf(w, "sample_api_http_requests_total{path=\"all\"} %d\n", atomic.LoadUint64(&totalRequests))
}

func track(endpoint string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&totalRequests, 1)

		switch endpoint {
		case "health":
			atomic.AddUint64(&healthRequests, 1)
		case "ready":
			atomic.AddUint64(&readyRequests, 1)
		case "metrics":
			atomic.AddUint64(&metricRequests, 1)
		}

		next(w, r)
	}
}

func writeText(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(status)

	if _, err := w.Write([]byte(body)); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
