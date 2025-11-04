package main

import (
	"fmt"
	"net/http"

	"github.com/rali7196/ecommerce/api_gateway/controllers"
	"github.com/rali7196/ecommerce/api_gateway/middleware"
)

func main() {
	fmt.Println("Starting server on port 3000")
	http.Handle("/", middleware.ApplyMiddleware(func(writer http.ResponseWriter, r *http.Request) {

	}))
	http.Handle("POST /users", middleware.ApplyMiddleware(controllers.CreateUser))
	http.ListenAndServe(":3000", nil)
}
