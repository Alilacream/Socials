package store

import (
	"context"
	"database/sql"

	"alilacream/socialx/models"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1,$2,$3,$4) RETURNING id, created_at, updated_at`
	err := s.db.QueryRowContext(
		ctx, query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) Search(ctx context.Context, post *models.Post) error {
	query := `SELECT title, content, tags FROM posts WHERE title=$1 OR content=$2 OR tags=$3`
	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.Tags).Scan(
		&post.ID,
		&post.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) Search_User_Posts(ctx context.Context, user *models.User, post *models.Post) error {
	query := `SELECT * FROM posts 
			LEFT JOIN users ON users.id = posts.user_id
			WHERE username = $1 
	`
	err := s.db.QueryRowContext(ctx, query, user.Username).Scan(
		&post,
	)
	if err != nil {
		return err
	}
	return nil
}
