-- posts table format 
CREATE TABLE IF NOT EXISTS posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(30),
  content TEXT,
  tags TEXT,
  user_id INT,
  CONSTRAINT fk_posts_users FOREIGN KEY (user_id) REFERENCES users(id)
)

