package repository

import (
	"API"
	"fmt"
	"gorm.io/gorm"
)

type PostSQL struct {
	db *gorm.DB
}

func NewPostSQL(db *gorm.DB) *PostSQL {
	return &PostSQL{db: db}
}

func (r *PostSQL) Create(userId int, post API.Posts) int {
	/*tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}*/
	var id int
	r.db.Raw(fmt.Sprintf("INSERT INTO `%s` (user, title, body) VALUES (?, ?, ?)", postsTable),
		userId, post.Title, post.Body).Scan(&id)
	/*if err := tx.Raw(fmt.Sprintf("INSERT INTO `%s` (user, title, body) VALUES (?, ?, ?)", postsTable), userId,
		post.Title, post.Body).Scan(&id).Error; err != nil {
		tx.Rollback()
		return 0, err
	}*/

	return id
}
