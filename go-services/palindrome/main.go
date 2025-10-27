package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync/atomic"

	"github.com/gorilla/mux"
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
	Value    bool   `json:"value"`
	CacheHit bool   `json:"cache_hit"`
}

type HealthResponse struct {
	OK bool `json:"ok"`
}

// Metrics
var requestsTotal int64

// Regex for extracting alphanumeric characters
var alphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

// Check if a string is a palindrome (alphanumeric only)
func isPalindrome(text string) bool {
	if text == "" {
		return true // Empty string is considered a palindrome
	}
	
	// Convert to lowercase and extract only alphanumeric characters
	cleaned := strings.ToLower(alphanumericRegex.ReplaceAllString(text, ""))
	
	if len(cleaned) == 0 {
		return true
	}
	
	// Check if cleaned string equals its reverse
	runes := []rune(cleaned)
	length := len(runes)
	
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	
	return true
}

// Handle POST /op endpoint
func handleOp(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&requestsTotal, 1)
	
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req OpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Handle invalid JSON gracefully
		response := OpResponse{
			Key:      "palindrome",
			Value:    false,
			CacheHit: false,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Use normalized text if available, otherwise use original text
	textToCheck := req.Deps.Normalized
	if textToCheck == "" {
		textToCheck = req.Text
	}
	
	isPalin := isPalindrome(textToCheck)
	
	response := OpResponse{
		Key:      "palindrome",
		Value:    isPalin,
		CacheHit: false,
	}
	
	json.NewEncoder(w).Encode(response)
}

// Handle GET /healthz endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := HealthResponse{OK: true}
	json.NewEncoder(w).Encode(response)
}

// Handle GET /metrics endpoint
func handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "# HELP palindrome_requests_total Total requests processed\\n")
	fmt.Fprintf(w, "# TYPE palindrome_requests_total counter\\n")
	fmt.Fprintf(w, "palindrome_requests_total %d\\n", atomic.LoadInt64(&requestsTotal))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8011"
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
	
	log.Printf("Palindrome service starting on port %s", port)
	log.Printf("Endpoints: POST /op, GET /healthz, GET /metrics")
	
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}