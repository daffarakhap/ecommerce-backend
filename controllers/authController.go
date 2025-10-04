package controllers

import (
	"ecommerce-backend/dto"
	"ecommerce-backend/middlewares"
	"ecommerce-backend/models"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// cek email sudah ada
		var exists models.User
		if err := db.Where("email = ?", input.Email).First(&exists).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
			return
		}

		hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		user := models.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hashed),
			Role:     "customer",
		}

		db.Create(&user)

		resp := dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered", "user": resp})
	}
}

// Login
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		token, _ := middlewares.GenerateJWT(user.ID, user.Role)

		resp := dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token, "user": resp})
	}
}

