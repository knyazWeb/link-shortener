package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"index"`
	Email    string
	Password string
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
