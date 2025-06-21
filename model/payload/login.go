// Package payload ...
package payload

// LoginCredential ...
type LoginCredential struct {
	Username string `json:"username" validate:"required" example:"user1"`
	Password string `json:"password" validate:"required" example:"password1"`
}

// Token ...
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
