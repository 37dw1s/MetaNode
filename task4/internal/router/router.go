package router

import (
	"GoLearning/task4/internal/handlers"
	"GoLearning/task4/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	//读取文章
	r.GET("/posts", handlers.ListPosts)
	r.GET("/posts/:post_id", handlers.GetPost)

	//读取评论
	r.GET("/posts/:post_id/comments", handlers.ListComments)

	//登录状态
	authed := r.Group("/authed", middleware.JWTAuth())
	{
		//创建文章
		authed.POST("/posts", handlers.CreatePost)

		//作者才能改删
		owner := authed.Group("/posts/:post_id", middleware.EnsurePostOwner())
		{
			//修改文章
			owner.PATCH("", handlers.UpdatePost)
			//删除文章
			owner.DELETE("", handlers.DeletePost)
		}

		//创建评论
		authed.POST("/posts/:post_id/comments", handlers.CreateComment)
	}

	return r
}
