package database_client

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func getConnection(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)
	return db, err
}

func CreateUser(email string, hashedPassword string) error {
	db, err := getConnection(DatabaseUrl)

	if err != nil {
		return err
	}

	id := uuid.New()
	_, queryErr := db.ExecContext(context.Background(),
		`INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`,
		id, email, hashedPassword)
	return queryErr
}
