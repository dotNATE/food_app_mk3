package dto

type LoginRequest struct {
	EmailField
	PasswordField
}

type AuthIdentity struct {
	ID       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Password string `json:"-"`
}
