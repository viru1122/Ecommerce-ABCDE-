package tests

import (
	"ecommerce/handlers"
	"ecommerce/models"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ = Describe("Cart API", func() {
	var db *gorm.DB
	var router *gin.Engine
	var cartHandler *handlers.CartHandler

	BeforeEach(func() {
		db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.Cart{}, &models.User{})
		db.Create(&models.User{Username: "test", Password: "pass"})
		router = gin.Default()
		cartHandler = &handlers.CartHandler{DB: db}
		router.GET("/carts", cartHandler.ListCarts)
	})

	It("should list carts", func() {
		db.Create(&models.Cart{UserID: 1})

		req, _ := http.NewRequest("GET", "/carts", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(200))
	})
})
