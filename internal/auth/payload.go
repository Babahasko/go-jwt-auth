package auth

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginResponse struct {
	Token string `json:"token" validate:"required"`
}

type RegisterRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterResponse struct {
	Token string `json:"token" validate:"required"`
}