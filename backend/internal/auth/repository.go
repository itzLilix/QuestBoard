package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id string) (*User, error)
}

type repository struct {
	db *pgx.Conn
}

func (r *repository) GetUserByID(id string) (*User, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT * FROM users WHERE id=$1", id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.CreatedAt, &user.LastLogin, &user.AvatarURL, &user.BannerURL, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) CreateUser(user *User) error {
	row := r.db.QueryRow(context.Background(),
		"INSERT INTO users (username, password_hash, email, role) VALUES ($1, $2, $3, 'user') RETURNING id, created_at",
		user.Username, user.PasswordHash, user.Email)
	err := row.Scan(&user.ID, &user.CreatedAt)
	return err
}

func (r *repository) GetUserByUsername(username string) (*User, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT * FROM users WHERE username=$1", username)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.CreatedAt, &user.LastLogin, &user.AvatarURL, &user.BannerURL, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewRepository(db *pgx.Conn) Repository {
	return &repository{db: db}
}
