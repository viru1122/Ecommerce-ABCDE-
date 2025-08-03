package handlers

import (
	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	h.DB.Create(&user)
	c.JSON(201, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var input struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := h.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid username/password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid username/password"})
		return
	}
	user.Token = uuid.New().String()
	h.DB.Save(&user)
	c.JSON(200, gin.H{"token": user.Token})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	var users []models.User
	h.DB.Find(&users)
	c.JSON(200, users)
}
