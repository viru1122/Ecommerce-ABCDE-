package tests

import (
	"ecommerce/handlers"
	"ecommerce/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("Order API", func() {
	var db *gorm.DB
	var router *gin.Engine
	var orderHandler *handlers.OrderHandler

	BeforeEach(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		Expect(err).To(BeNil())

		// Migrate all related models
		err = db.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})
		Expect(err).To(BeNil())

		// Create test user, item, cart, and order
		user := models.User{Username: "test", Password: "pass"}
		db.Create(&user)

		item := models.Item{Name: "Test Item", Description: "A test item", Price: 10}
		db.Create(&item)

		cart := models.Cart{UserID: user.ID}
		db.Create(&cart)
		db.Model(&cart).Association("Items").Append(&item)

		order := models.Order{UserID: user.ID, CartID: cart.ID}
		db.Create(&order)

		orderHandler = &handlers.OrderHandler{DB: db}
		router = gin.Default()
		router.GET("/orders", orderHandler.ListOrders)
	})

	It("should list orders", func() {
		req, _ := http.NewRequest("GET", "/orders", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))

		// Optional: check body content
		var response []models.Order
		err := json.Unmarshal(w.Body.Bytes(), &response)
		Expect(err).To(BeNil())
		Expect(len(response)).To(Equal(1))
		Expect(response[0].CartID).NotTo(BeZero())
	})
})
