package controllers

import (
	"go-blog-app/database"
	"go-blog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		PostID uint
		Body   string
	}
	c.BindJSON(&input)

	comment := models.Comment{
		PostID: input.PostID,
		UserID: userID.(uint),
		Body:   input.Body,
	}

	result := database.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func UpdateComment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	commentID := c.Param("id")

	var comment models.Comment
	result := database.DB.First(&comment, commentID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not your comment"})
		return
	}

	var input struct {
		Body string
	}
	c.BindJSON(&input)

	database.DB.Model(&comment).Update("body", input.Body)

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func DeleteComment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	commentID := c.Param("id")

	var comment models.Comment
	result := database.DB.First(&comment, commentID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not your comment"})
		return
	}

	database.DB.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
