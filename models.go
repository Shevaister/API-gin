package API

import "github.com/pkg/errors"

type Posts struct {
	User  int    `json:"userId" db:"user"`
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Body  string `json:"body" db:"body" binding:"required"`
}

type Comments struct {
	Post  int    `json:"postId"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}

type Users struct {
	Id       int    `json:"-" db:"id"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UpdatePostInput struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Body == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
