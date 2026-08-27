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
	// pq.Array is a builtin function in postgres lib of pq. the name is self explanotory fn.
	query := `SELECT title, content, tags, created_at  FROM posts WHERE id = $1`
	err := s.db.QueryRowContext(ctx, query, post.ID).Scan(
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) Search_User_Posts(ctx context.Context, username string, posts *[]models.Post) error {
	query := `SELECT * FROM posts 
			LEFT JOIN users ON users.id = posts.user_id
			WHERE username = $1 
	`
	// what happens when there are multiple posts
	rows, err := s.db.QueryContext(ctx, query, username)
	if err != nil {
		return err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, pq.Array(&post.Tags), &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return err
		}
		*posts = append(*posts, post)
	}
	return nil
}

func (s *PostStore) AllPosts(ctx context.Context, posts *[]models.Post) error {
	// pq.Array is a builtin function in postgres lib of pq. the name is self explanotory fn.
	query := `SELECT id, title, content, tags, created_at, updated_at  FROM posts`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return err
		}
		*posts = append(*posts, post)
	}
	return nil
}
