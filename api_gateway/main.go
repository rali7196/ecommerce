package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rali7196/ecommerce/api_gateway/authentication"
)

type structResponse struct {
	response string
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	response := structResponse{"Hello World"}
	responseJson, err := json.Marshal(response)

	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func CorsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next.ServeHTTP(w, r)
	}
}

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received Request")
		next.ServeHTTP(w, r)
	}
}

func ApplyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	middlewares := []func(http.Handler) http.HandlerFunc{
		CorsMiddleware, LoggingMiddleware}

	finalHandler := handler
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler)
	}

	return finalHandler
}

func main() {
	authentication.Test()
	fmt.Println("Starting server on port 3000")
	http.HandleFunc("/", ApplyMiddleware(testHandler))
	http.ListenAndServe(":3000", nil)
}
