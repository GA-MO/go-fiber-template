package models

import "time"

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdatePassword struct {
	Username    string `json:"username" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
