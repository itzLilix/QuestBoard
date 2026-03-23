package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(username, email, password string) (*User, string, error)
	Login(username, password string) (*User, string, error)
	ValidateToken(tokenString string) (*User, error)
}

type claims struct {
	ID string
	Username string
	Role string
	jwt.RegisteredClaims
}

func (s *service) ValidateToken(tokenString string) (*User, error){
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	var user *User

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*claims); ok {
		user, err = s.repo.GetUserByID(claims.ID)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unknown claims type")
	}
	return user, nil
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Register(username, email, password string) (*User, string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: string(passwordHash),
	}
	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, "", err
	}

	token, err := s.generateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, err
}

func (s *service) Login(username, password string) (*User, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", err
	}

	token, err := s.generateAccessToken(user)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (s *service) generateAccessToken(user *User) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})

	return token.SignedString(secretKey)
}

func (s *service) generateRefreshToken(user *User) (string, error) {
	return "", nil
}
