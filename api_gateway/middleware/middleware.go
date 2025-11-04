package middleware

import (
	"fmt"
	"net/http"
)

func corsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next.ServeHTTP(w, r)
	}
}

func loggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received Request: %s %s %s\n", r.URL, r.Method, r.Body)
		next.ServeHTTP(w, r)
	}
}

func ApplyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	middlewares := []func(http.Handler) http.HandlerFunc{
		corsMiddleware, loggingMiddleware}

	finalHandler := handler
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler)
	}

	return finalHandler
}
