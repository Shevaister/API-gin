package API

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
