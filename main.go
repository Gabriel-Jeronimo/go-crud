package main

import (
	"go-crud/books"
	"go-crud/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&books.BookModel{})
}

func main() {
	db := common.Init()
	Migrate(db)

	r := gin.Default()

	api := r.Group("/api")
	books.BooksRegister(api.Group("/books"))

	testRoute := r.Group("/api/ping")
	testRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
