package models

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	UserID    int64    `json:"user_id"` // forgein key to the user
	CreatedAt string   `json:"created"`
	UpdatedAt string   `json:"updated"`
}

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
type DBConfig struct {
	DSN                string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}
