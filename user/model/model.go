package model

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email`
	Password string `json:"password" binding:"required,min=12`
}
