package store

import (
	"context"
	"database/sql"

	"alilacream/socialx/lib"
	"alilacream/socialx/models"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name,last_name, username, email, password)
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at 
	`
	hashPass, err := lib.HashPassword(user.Password)
	if err != nil {
		return err
	}
	err = s.db.QueryRowContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		hashPass).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
