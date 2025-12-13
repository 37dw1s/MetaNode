package middleware

import (
	"GoLearning/task4/internal/database"
	"GoLearning/task4/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsurePostOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		uidAny, _ := c.Get("userID")
		uid := uint(uidAny.(float64))

		postid := c.Param("post_id")
		var post models.Post
		if err := database.DB.First(&post, postid).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}
		if post.UserID != uid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: not the author"})
			return
		}
		c.Set("post", post)
		c.Next()
	}
}
