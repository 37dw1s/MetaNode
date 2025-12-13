package handlers

import (
	"GoLearning/task4/internal/database"
	"GoLearning/task4/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	uid := uint(c.MustGet("userID").(float64))
	var commentReq models.Comment
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postID := c.Param("post_id")
	//var post models.Post
	//if err := database.DB.First(&post, postID).Error; err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	//	return
	//}
	id64, _ := strconv.ParseUint(postID, 10, 64)

	cm := models.Comment{Content: commentReq.Content, UserID: uid, PostID: uint(id64)}
	if err := database.DB.Create(&cm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cm)
}

func ListComments(c *gin.Context) {
	var cmts []models.Comment
	if err := database.DB.Where("post_id = ?", c.Param("post_id")).Find(&cmts).Order("id DES").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cmts)
}
