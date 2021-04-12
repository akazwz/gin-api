package response

import "github.com/akaedison/go-gin-demo/model"

type UserResponse struct {
	User model.User `json:"user"`
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expires_at"`
}
