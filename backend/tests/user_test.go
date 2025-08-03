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

var _ = Describe("User API", func() {
	var db *gorm.DB
	var router *gin.Engine
	var userHandler *handlers.UserHandler

	BeforeEach(func() {
		db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.User{})
		router = gin.Default()
		userHandler = &handlers.UserHandler{DB: db}
		router.POST("/users", userHandler.CreateUser)
	})

	It("should create a user", func() {
		user := map[string]string{"username": "testuser", "password": "testpass"}
		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(201))
		var response models.User
		json.Unmarshal(w.Body.Bytes(), &response)
		Expect(response.Username).To(Equal("testuser"))
	})
})
