package store

import (
	"context"
	"database/sql"

	"alilacream/socialx/models"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (Name, Email, Password)
	VALUES ($1, $2, $3) RETURNING id, created_at 
	`
	err := s.db.QueryRowContext(ctx, query,
		user.Name,
		user.Email,
		user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
