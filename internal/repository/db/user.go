package db

import (
	"context"
	"database/sql"
	"my-crud-app/internal/domain"
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(ctx context.Context, user domain.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, password, registered_at) values ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.RegisteredAt)

	return err
}

func (u *User) GetUser(ctx context.Context, email, password string) (domain.User, error) {
	var user domain.User
	err := u.db.QueryRow("SELECT id, name, email, registered_at FROM users WHERE email=$1 AND password=$2", email, password).
		Scan(&user.Id, &user.Name, &user.Email, &user.RegisteredAt)

	return user, err
}
