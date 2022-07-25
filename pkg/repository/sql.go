package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	postsTable    = "posts"
	commentsTable = "comments"
	usersTable    = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	//SSLMode  string
}

func NewSQLDB(cfg Config) (*gorm.DB, error) {
	//dsn := "root:@tcp(127.0.0.1:3306)/commentsdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
