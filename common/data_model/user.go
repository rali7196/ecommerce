package data_model

import "github.com/google/uuid"

type User struct {
	Email    string
	Password string
	Id       uuid.UUID
}
