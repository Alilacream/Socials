package store

import (
	"context"
	"database/sql"

	"alilacream/socialx/models"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *models.Post) error
		Search(ctx context.Context, post *models.Post) error
	}
	Users interface {
		Create(ctx context.Context, user *models.User) error
		Check_User_Exist(ctx context.Context, user *models.User) error
	}
}

func NewPQStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
