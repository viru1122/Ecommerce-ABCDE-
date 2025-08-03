package tests

import (
	"bytes"
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

var _ = Describe("Item API", func() {
	var db *gorm.DB
	var router *gin.Engine
	var itemHandler *handlers.ItemHandler

	BeforeEach(func() {
		db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.Item{})
		router = gin.Default()
		itemHandler = &handlers.ItemHandler{DB: db}
		router.POST("/items", itemHandler.CreateItem)
		router.GET("/items", itemHandler.ListItems)
	})

	It("should create an item", func() {
		item := map[string]interface{}{"name": "Laptop", "price": 999.99}
		body, _ := json.Marshal(item)
		req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(201))
	})

	It("should list items", func() {
		db.Create(&models.Item{Name: "Phone", Price: 499.99})

		req, _ := http.NewRequest("GET", "/items", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(200))
		Expect(w.Body.String()).To(ContainSubstring("Phone"))
	})
})
