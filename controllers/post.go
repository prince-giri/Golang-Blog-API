package controllers

import (
	"go-blog-app/database"
	"go-blog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		Title  string
		Body   string
		Status string // draft, published, archived
	}
	c.BindJSON(&input)

	post := models.Post{
		Title:  input.Title,
		Body:   input.Body,
		Status: input.Status,
		UserID: userID.(uint),
	}

	result := database.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func ListPublishedPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Where("status = ?", "published").Preload("Comments").Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func UpdatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Param("id")

	var post models.Post
	result := database.DB.First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not your post"})
		return
	}

	var input struct {
		Title  string
		Body   string
		Status string
	}
	c.BindJSON(&input)

	database.DB.Model(&post).Updates(models.Post{
		Title:  input.Title,
		Body:   input.Body,
		Status: input.Status,
	})

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Param("id")

	var post models.Post
	result := database.DB.First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not your post"})
		return
	}

	database.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
