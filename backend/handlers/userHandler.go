package handlers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"ecommerce-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Register a new user
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.CreatedAt = time.Now()
	query := "INSERT INTO users (name, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password, user.Role, user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

// Login user and return JWT token
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedPassword string
	query := "SELECT password FROM users WHERE email = $1"
	err := config.DB.QueryRow(query, user.Email).Scan(&storedPassword)
	if err != nil || storedPassword != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
