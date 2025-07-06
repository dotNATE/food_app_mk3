package dto

type NameField struct {
	Name string `json:"name" binding:"required,max=191"`
}

type EmailField struct {
	Email string `json:"email" binding:"required,email,max=191"`
}

type PasswordField struct {
	Password string `json:"password" binding:"required,min=8,max=191"`
}

type RegisterRequest struct {
	NameField
	EmailField
	PasswordField
}

func NewRegisterRequest(name string, email string, password string) *RegisterRequest {
	return &RegisterRequest{
		NameField:     NameField{Name: name},
		EmailField:    EmailField{Email: email},
		PasswordField: PasswordField{Password: password},
	}
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
