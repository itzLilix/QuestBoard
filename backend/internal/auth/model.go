package auth

import "time"

type User struct {
	ID           string   	`json:"id"`
	Username     string   	`json:"username"`
	PasswordHash string   	`json:"-"`
	Email        string   	`json:"email"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastLogin    *time.Time  `json:"lastLogin,omitempty"`
	AvatarURL    *string   	`json:"avatarUrl,omitempty"`
	BannerURL    *string   	`json:"bannerUrl,omitempty"`
	Role         string   	`json:"role"`
	DisplayName  *string	`json:"displayName,omitempty"`
	IsEmailVerified bool `json:"isEmailVerified"`
}

type RefreshToken struct {
	ID string `json:"id"`
	UserID string `json:"userId"`
	TokenPrefix string `json:"tokenPrefix"`
	TokenHash string `json:"tokenHash"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}