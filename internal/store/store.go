package store

import (
	"context"
	"database/sql"

	"alilacream/socialx/internal/env"
	"alilacream/socialx/models"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *models.Post) error
		Search(ctx context.Context, post *models.Post) error
		Search_User_Posts(ctx context.Context, username string, post []models.Post) error
	}
	Users interface {
		Create(ctx context.Context, user *models.User) error
		Search(ctx context.Context, user *models.User) error
		// NOTE: this one is utilized for loginv1 handler
		Check_User_Exist(ctx context.Context, user *models.User) error
	}
	JWTSecret string // since we'll be passing jwt secret a lot let just set the call of getVar() in main only
}

func NewPQStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		JWTSecret: env.GetVar("SECRET_KEY"),
	}
}
