package dto

type NameAndEmailFields struct {
	Name  string `json:"name" binding:"required,max=191"`
	Email string `json:"email" binding:"required,email,max=191"`
}

type PasswordField struct {
	Password string `json:"password" binding:"required,min=8,max=191"`
}

type RegisterRequest struct {
	NameAndEmailFields
	PasswordField
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
