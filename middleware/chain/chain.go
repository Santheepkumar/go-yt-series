
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware type
type Middleware func(http.Handler) http.Handler

// Function to chain multiple middlewares
func chainMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Response time: %v", duration)
	})
}

// Authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "mysecretkey" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Example handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Middleware!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/chain", chainMiddlewares(http.HandlerFunc(helloHandler), loggingMiddleware, authMiddleware))

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}
