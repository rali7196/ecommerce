package main

import (
	"fmt"
	"net/http"

	"github.com/rali7196/ecommerce/api_gateway/controllers"
	"github.com/rali7196/ecommerce/api_gateway/middleware"
)

func main() {
	fmt.Println("Starting server on port 3000")

	mux := http.NewServeMux()
	var uc controllers.UserController
	uc.RegisterRoutes(mux)

	http.ListenAndServe(":3000", middleware.ApplyMiddleware(mux))
}
