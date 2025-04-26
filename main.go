package main

import (
	"go-blog-app/controllers"
	"go-blog-app/database"
	"go-blog-app/middleware"
	"go-blog-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	router := gin.Default()

	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/posts", controllers.ListPublishedPosts)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/posts", controllers.CreatePost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)

		protected.POST("/comments", controllers.AddComment)
		protected.PUT("/comments/:id", controllers.UpdateComment)
		protected.DELETE("/comments/:id", controllers.DeleteComment)
	}

	router.Run(":8080")
}
