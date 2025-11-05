package api_gateway

type CreateUserResponse struct {
	Status int
}

type CreateUserRequest struct {
	Email    string
	Password string
}
