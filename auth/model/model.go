package model

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=12"`
}

type LoginResponse struct {
	Data interface{} `json:"data"`
}