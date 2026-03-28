package auth

import (
	"context"

	"github.com/itzLilix/QuestBoard/backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User = models.User
type RefreshToken = models.RefreshToken

type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
	SaveRefreshToken(token *RefreshToken) error
	GetRefreshTokenByPrefix(prefix string) (*RefreshToken, error)
	DeleteRefreshToken(prefix string) error
	UpdateLastLogin(user *User) error
}

type repository struct {
	db *pgxpool.Pool
}


func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) scanUser(row pgx.Row, user *User) error{
	return row.Scan(
		&user.ID, 
		&user.Username, 
		&user.PasswordHash, 
		&user.Email, 
		&user.CreatedAt, 
		&user.LastLogin, 
		&user.AvatarURL, 
		&user.BannerURL, 
		&user.Role, 
		&user.DisplayName, 
		&user.IsEmailVerified,
		&user.SessionsPlayed,
		&user.SessionsHosted,
		&user.Rating,
		&user.ReviewsCount)
}

func (r *repository) GetUserByID(id string) (*User, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT * FROM users WHERE id=$1", id)
	user := &User{}
	err := r.scanUser(row, user)
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

func (r *repository) GetUserByEmail(email string) (*User, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT * FROM users WHERE email=$1", email)
	user := &User{}
	err := r.scanUser(row, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) SaveRefreshToken(token *RefreshToken) error {
	row := r.db.QueryRow(context.Background(),
	"INSERT INTO refresh_tokens (user_id, token_prefix, token_hash, expires_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
	token.UserID, token.TokenPrefix, token.TokenHash, token.ExpiresAt)
	err := row.Scan(&token.ID, &token.CreatedAt)
	return err
}

func (r *repository) GetRefreshTokenByPrefix(prefix string) (*RefreshToken, error) {
	row := r.db.QueryRow(context.Background(),
	"SELECT * FROM refresh_tokens WHERE token_prefix=$1", prefix)
	token := &RefreshToken{}
	err := row.Scan(&token.ID, &token.UserID, &token.TokenPrefix, &token.TokenHash, &token.ExpiresAt, &token.CreatedAt)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *repository) DeleteRefreshToken(prefix string) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM refresh_tokens WHERE token_prefix=$1", prefix)
	return err
}

func (r *repository) UpdateLastLogin(user *User) error {
	row := r.db.QueryRow(context.Background(), "UPDATE users SET last_login = NOW() WHERE id = $1 RETURNING last_login", user.ID)
	err := row.Scan(&user.LastLogin)
	return err
}