package controllers

import (
	"encoding/json"
	"fmt"
	"microservice_types/api_gateway"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser called")
	response := api_gateway.CreateUserResponse{Status: http.StatusOK}

	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}
