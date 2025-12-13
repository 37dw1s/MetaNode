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
		uid := uidAny.(float64)

		id := c.Param("id")
		var post models.Post
		if err := database.DB.First(&post, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}
		if post.UserID != uint(uid) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: not the author"})
			return
		}
		c.Set("post", post)
		c.Next()
	}
}
