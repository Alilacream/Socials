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

// TODO:  need to migrate to string UUID
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"` // was not being read in the postman, now it is
	CreatedAt string
}

type DBConfig struct {
	DSN                string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}
