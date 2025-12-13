package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
type User struct {
	gorm.Model
	Name      string
	Posts     []Post
	PostCount int `gorm:"not null;default:0"`
}

type Post struct {
	gorm.Model
	UserID        uint `gorm:"not null"`
	User          User
	Content       string `gorm:"type:longtext"`
	Comments      []Comment
	CommentStatus string `gorm:"not null;default:'无评论'"`
}

type Comment struct {
	gorm.Model
	PostID uint `gorm:"not null"`
	Post   Post
	Body   string `gorm:"type:text;not null"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_status", "有评论").Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var cnt int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

func main() {
	db, err := gorm.Open(mysql.Open("root:tt@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		panic(err)
	}

	u1 := User{Name: "Alice"}
	db.Create(&u1)

	p1 := Post{
		UserID:  u1.ID,
		Content: "Hello, this is Alice's first post",
	}
	db.Create(&p1)

	c1 := Comment{
		PostID: p1.ID,
		Body:   "Nice post, Alice!",
	}
	db.Create(&c1)

	//u2 := User{
	//	Name: "Bob",
	//	Posts: []Post{
	//		{
	//			Content: "Bob writes something interesting.",
	//			Comments: []Comment{
	//				{Body: "Great thoughts, Bob!"},
	//				{Body: "Love your content!"},
	//			},
	//		},
	//	},
	//}
	//
	//db.Create(&u2)

	//u2 := User{Name: "Bob"}
	//db.Create(&u2)
	//
	//p2 := Post{Content: "Bob's post"}
	//u2.Posts = append(u2.Posts, p2)
	//db.Session(&gorm.Session{FullSaveAssociations: true}).
	//	Omit("Posts.User").
	//	Save(&u2)

	u2 := User{Name: "Bob"}
	db.Create(&u2)

	p2 := Post{Content: "Bob writes something interesting."}
	db.Model(&u2).
		Omit("User").
		Association("Posts").
		Append(&p2)

	c2 := Comment{Body: "Good!"}
	db.Model(&p2).
		Omit("Post").
		Association("Comments").
		Append(&c2)

	var u User
	if err := db.
		Preload("Posts").
		Preload("Posts.Comments").
		Take(&u).Error; err != nil {
		panic(err)
	}

	if err := db.Delete(&u.Posts[0].Comments[0]).Error; err != nil {
		log.Fatal(err)
	}

}
