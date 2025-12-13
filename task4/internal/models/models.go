package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post
	Comments []Comment
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title"`
	Content  string `gorm:"not null" json:"content"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	User     User   `json:"author"`
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	//User    User   `json:"author"`
	PostID uint `json:"post_id"`
	Post   Post `json:"-"`
}
