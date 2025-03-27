
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware function
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Response time: %v", duration)
	})
}

// Actual handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Middleware!")
}

func main() {
	// Create a new mux router
	mux := http.NewServeMux()

	// Register handler
	mux.Handle("/", loggingMiddleware(http.HandlerFunc(helloHandler)))

	// Start server
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}
