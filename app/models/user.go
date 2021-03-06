package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	//gorm.Model
	Email    string `json:"username" gorm:"unique;" validate:"required,email,min=6,max=32"`
	Password string `json:"password" gorm:"type:text;" validate:"required,min=6"`
	Fullname string `json:"fullname"`
	Role     int    `json:"role"`
}

func NewUser(email string, pass string, fullname string, role int) User {
	user := User{email, pass, fullname, role}
	return user
}

func (u User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}
