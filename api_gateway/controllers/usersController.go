package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"microservice_types/api_gateway"
)

type UserController struct{}

func (u UserController) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createUser called")

	var requestBodyJson api_gateway.CreateUserRequest
	json.NewDecoder(r.Body).Decode(&requestBodyJson)

	//var RequestBodyCopy api_gateway.CreateUserRequest
	//err = json.NewDecoder(r.Body).Decode(&RequestBodyCopy)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	response := api_gateway.CreateUserResponse{Status: http.StatusOK}

	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func (u UserController) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /users", http.HandlerFunc(u.createUser))
}
