package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"microservice_types/api_gateway"
	"net/http"
)

type userHandler struct {
}

func getSha256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("password hash %s\n", passwordHash)
	return passwordHash
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var requestBodyJson api_gateway.CreateUserRequest
	_ = json.NewDecoder(r.Body).Decode(&requestBodyJson)
	hashedPassword := getSha256Hash(requestBodyJson.Password)
	fmt.Printf("email: %s password: %s hashedPassword: %s\n", requestBodyJson.Email, requestBodyJson.Password,
		hashedPassword)
}

func main() {
	http.HandleFunc("POST /users", createUser)
	fmt.Println("User service started on port 3001")
	http.ListenAndServe(":3001", nil)
}
