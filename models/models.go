package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Posts    []Post
	Comments []Comment
}

type Post struct {
	gorm.Model
	Title    string
	Body     string
	Status   string // draft, published, archived
	UserID   uint
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Body   string
	UserID uint
	PostID uint
}
