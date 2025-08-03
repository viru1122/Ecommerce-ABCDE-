package handlers

import (
	"ecommerce/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	DB *gorm.DB
}

func (h *CartHandler) CreateCart(c *gin.Context) {
	// Check user authentication
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Bind JSON input with validation
	var input struct {
		ItemID uint `json:"item_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.ItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing 'item_id'"})
		return
	}

	// Fetch or create cart for the user
	var cart models.Cart
	if err := h.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: userID.(uint)}
		if err := h.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
			return
		}
	}

	// Check if the item exists
	var item models.Item
	if err := h.DB.First(&item, input.ItemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Add item to cart
	if err := h.DB.Model(&cart).Association("Items").Append(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	fmt.Printf("Item %d added to cart for user %v\n", input.ItemID, userID)
	c.JSON(http.StatusCreated, cart)
}

func (h *CartHandler) ListCarts(c *gin.Context) {
	var carts []models.Cart
	if err := h.DB.Preload("Items").Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch carts"})
		return
	}
	c.JSON(http.StatusOK, carts)
}
