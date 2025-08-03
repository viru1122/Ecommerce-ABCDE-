package main

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"ecommerce/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Supreet@123 dbname=ecommerce port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate models
	db.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})

	r := gin.Default()

	// ✅ Add CORS middleware here
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userHandler := handlers.UserHandler{DB: db}
	itemHandler := handlers.ItemHandler{DB: db}
	cartHandler := handlers.CartHandler{DB: db}
	orderHandler := handlers.OrderHandler{DB: db}

	// Optional: health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Backend is running"})
	})

	// Routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.ListUsers)
	r.POST("/users/login", userHandler.Login)

	r.POST("/items", itemHandler.CreateItem)
	r.GET("/items", itemHandler.ListItems)

	r.POST("/carts", middleware.AuthMiddleware(db), cartHandler.CreateCart)
	r.GET("/carts", middleware.AuthMiddleware(db), cartHandler.ListCarts)

	r.POST("/orders", middleware.AuthMiddleware(db), orderHandler.CreateOrder)
	r.GET("/orders", middleware.AuthMiddleware(db), orderHandler.ListOrders)

	// Start server
	r.Run(":8080")
}
