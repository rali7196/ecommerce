package database_client

import (
	"context"
	"database/sql"
	"database_client/data_model"
	"testing"
)

func setup()    {}
func tearDown() {}

func TestCreateUser(t *testing.T) {
	email := "billy123@gmail.com"
	password := "zbc123"
	err := CreateUser(email, password)
	if err != nil {
		t.Fatal("Create user query failed", err)
	}

	db, err := sql.Open("pgx", DatabaseUrl)
	if err != nil {
		t.Fatal("opening database connection failed", err)
	}
	rows, err := db.QueryContext(context.Background(), "SELECT id, email, password FROM users")

	var users []data_model.User
	for rows.Next() {
		var user data_model.User
		err = rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			t.Error("scan row failed", err)
		}
		users = append(users, user)
	}

	if len(users) != 1 {
		t.Fatal("expected 1 user, got ", len(users))
	}

	if users[0].Email != email {
		t.Fatal("expected email to be ", email)
	}
}
