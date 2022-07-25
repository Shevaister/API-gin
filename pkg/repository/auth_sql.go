package repository

import (
	"API"
	"gorm.io/gorm"
)

type AuthSQL struct {
	db *gorm.DB
}

func NewAuthSQL(db *gorm.DB) *AuthSQL {
	return &AuthSQL{db: db}
}

func (r *AuthSQL) CreateUser(user API.Users) (int, error) {
	var id int

	field := API.Users{Password: user.Email, Email: user.Password}
	r.db.Create(&field)
	r.db.Select(field, "email", "password")
	r.db.Last(&id)

	return id, nil
}
