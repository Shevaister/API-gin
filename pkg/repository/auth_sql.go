package repository

import (
	"API"
	"fmt"
	"gorm.io/gorm"
)

type AuthSQL struct {
	db *gorm.DB
}

func NewAuthSQL(db *gorm.DB) *AuthSQL {
	return &AuthSQL{db: db}
}

func (r *AuthSQL) CreateUser(user API.Users) (int, error) {
	field := API.Users{Email: user.Email, Password: user.Password}
	err := r.db.Create(&field).Error
	r.db.Select(field, "email", "password")
	r.db.Last(&field)

	return field.Id, err
}

func (r *AuthSQL) GetUser(email, password string) API.Users {
	var user API.Users
	r.db.Raw(fmt.Sprintf("SELECT id FROM `%s` WHERE email = ? AND password = ?", usersTable), email, password).Scan(&user)

	return user
}
