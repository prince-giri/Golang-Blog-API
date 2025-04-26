package controllers

import (
	"go-blog-app/database"
	"go-blog-app/models"
	"go-blog-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string
		Email    string
		Password string
	}
	c.BindJSON(&input)

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.User{Name: input.Name, Email: input.Email, Password: string(hash)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string
		Password string
	}
	c.BindJSON(&input)

	var user models.User
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
