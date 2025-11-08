package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"microservice_types/api_gateway"
)

type UserController struct{}

var userServiceUrl = "http://localhost:3001"

func (u UserController) routeCreateUserRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("routeCreateUserRequest called")

	_, err := http.Post(userServiceUrl+"/users", "application/json", r.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		responseJson, _ := json.Marshal(api_gateway.CreateUserResponse{Status: http.StatusInternalServerError})
		w.Write(responseJson)
		return
	}

	response := api_gateway.CreateUserResponse{Status: http.StatusOK}
	responseJson, err := json.Marshal(response)

	w.Write(responseJson)
}

func (u UserController) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /users", http.HandlerFunc(u.routeCreateUserRequest))
}
