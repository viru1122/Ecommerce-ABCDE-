package handlers

import (
	"ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemHandler struct {
	DB *gorm.DB
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&item)
	c.JSON(201, item)
}

func (h *ItemHandler) ListItems(c *gin.Context) {
	var items []models.Item
	h.DB.Find(&items)
	c.JSON(200, items)
}
