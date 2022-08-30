package db

import (
	"context"
	"database/sql"
	"log"
	"rest-go/models"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (p *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (p *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)", post.Id, post.PostContent, post.UserId)
	return err
}

func (p *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, post_content, created_at, user_id FROM posts WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var post = models.Post{}
	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UserId); err == nil {
			return &post, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresRepository) Close() error {
	return p.db.Close()
}

func (p *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := p.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2 and user_id = $3", post.PostContent, post.Id, post.UserId)
	return err
}

func (p *PostgresRepository) DeletePost(ctx context.Context, id string, userid string) error {
	_, err := p.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 and user_id = $2", id, userid)
	return err
}
