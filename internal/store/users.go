package store

import (
	"context"
	"database/sql"
	"errors"

	"alilacream/socialx/lib"
	"alilacream/socialx/models"
)

type UserStore struct {
	db *sql.DB
}

// Create: UserStore method to create the new user provided in the params
func (s *UserStore) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name,last_name, username, email, password)
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at `
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

// Checking the user if he exists aswell as checking the passwword within,make it valuable for our login handler
func (s *UserStore) Check(ctx context.Context, user *models.User) (*models.User, error) {
	query := `SELECT id ,username, password FROM users 
	WHERE username = $1`
	var UserCheck models.User
	err := s.db.QueryRowContext(ctx, query, user.Username).Scan(&UserCheck.ID, &UserCheck.Username, &UserCheck.Password)
	// logically the row should return either a Invalid Query OR nothing
	if err != nil {
		return nil, err
	}
	// checking if the password is valid or not (NOTE:it takes the HASHED one, and the user inputed one )
	if check := lib.CheckPassword(UserCheck.Password, user.Password); !check {
		return nil, errors.New("invalid password")
	}
	return &UserCheck, nil
}
