package tests

import (
	"bytes"
	"encoding/json"
	"go-crud/books"
	"go-crud/common"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	books.BooksRegister(v1.Group("/books"))

	testRoute := r.Group("/api/ping")
	testRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func MakeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)

	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func setup() {
	db := common.Init()
	db.AutoMigrate(&books.BookModel{})

}
