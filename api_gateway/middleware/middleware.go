package middleware

import (
	"fmt"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received Request: %s %s %s\n", r.URL, r.Method, r.Body)
		next.ServeHTTP(w, r)
	})
}

func ApplyMiddleware(handler http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{loggingMiddleware, corsMiddleware}

	finalHandler := handler
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler)
	}

	return finalHandler
}
