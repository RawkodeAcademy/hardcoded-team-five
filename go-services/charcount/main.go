package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"unicode/utf8"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Request and Response structures
type OpRequest struct {
	Text string `json:"text,omitempty"`
	Deps struct {
		Normalized string `json:"normalized,omitempty"`
	} `json:"deps,omitempty"`
}

type OpResponse struct {
	Key      string `json:"key"`
	Value    int    `json:"value"`
	CacheHit bool   `json:"cache_hit"`
}

type HealthResponse struct {
	OK bool `json:"ok"`
}

// Metrics
var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "charcount_requests_total",
			Help: "Total number of requests processed by charcount service",
		},
		[]string{"endpoint", "method"},
	)
	
	requestsTotal int64
)

func init() {
	prometheus.MustRegister(requestCounter)
}

// Count characters in text using UTF-8 rune counting
func countChars(text string) int {
	if text == "" {
		return 0
	}
	return utf8.RuneCountInString(text)
}

// Handle POST /op endpoint
func handleOp(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&requestsTotal, 1)
	requestCounter.WithLabelValues("op", "POST").Inc()
	
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req OpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Handle invalid JSON gracefully
		response := OpResponse{
			Key:      "char_count",
			Value:    0,
			CacheHit: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Use normalized text if available, otherwise use original text
	textToCount := req.Deps.Normalized
	if textToCount == "" {
		textToCount = req.Text
	}
	
	charCount := countChars(textToCount)
	
	response := OpResponse{
		Key:      "char_count",
		Value:    charCount,
		CacheHit: false,
	}
	
	json.NewEncoder(w).Encode(response)
}

// Handle GET /healthz endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	requestCounter.WithLabelValues("healthz", "GET").Inc()
	
	w.Header().Set("Content-Type", "application/json")
	
	response := HealthResponse{OK: true}
	json.NewEncoder(w).Encode(response)
}

// Handle GET /metrics endpoint
func handleMetrics(w http.ResponseWriter, r *http.Request) {
	// Custom metrics response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "# HELP charcount_requests_total Total requests processed\n")
	fmt.Fprintf(w, "# TYPE charcount_requests_total counter\n")
	fmt.Fprintf(w, "charcount_requests_total %d\n", atomic.LoadInt64(&requestsTotal))
	
	// Also serve Prometheus metrics
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8007"
	}
	
	r := mux.NewRouter()
	
	// API endpoints
	r.HandleFunc("/op", handleOp).Methods("POST")
	r.HandleFunc("/healthz", handleHealth).Methods("GET")
	r.HandleFunc("/metrics", handleMetrics).Methods("GET")
	
	// Add CORS headers for development
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	})
	
	log.Printf("CharCount service starting on port %s", port)
	log.Printf("Endpoints: POST /op, GET /healthz, GET /metrics")
	
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}