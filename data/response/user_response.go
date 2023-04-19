package response

import (
	"time"

	"mygram-api/models"
)

type UserResponse struct {
	Id        string         `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Age       int            `json:"age"`
	Photos    []models.Photo `json:"photos,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
