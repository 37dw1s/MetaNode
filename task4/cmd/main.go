package main

import (
	"GoLearning/task4/internal/database"
	"GoLearning/task4/internal/handlers"
	"GoLearning/task4/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	err := database.Init()
	if err != nil {
		return
	}

	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	//post读取
	r.GET("/posts", handlers.ListPosts)
	r.GET("/posts/:id", handlers.GetPost)

	//登录状态
	authed := r.Group("/authed", middleware.JWTAuth())
	{
		//创建
		authed.POST("/posts", handlers.CreatePost)

		//作者才能改删
		owner := authed.Group("/posts/:id", middleware.EnsurePostOwner())
		{
			//修改
			owner.PATCH("", handlers.UpdatePost)
			//删除
			owner.DELETE("", handlers.DeletePost)
		}

		//创建评论
		authed.POST("/posts/:id/comments", handlers.CreateComment)
	}
	//读评论
	r.GET("/posts/:id/comments", handlers.ListComments)

	//jwtmw.AuthRequired()

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
