package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BooksRegister(router *gin.RouterGroup) {
	router.POST("/", BookCreate)
	router.GET("/", BookList)
}

func BookCreate(c *gin.Context) {

}

func BookList(c *gin.Context) {
	books, err := FindManyBook()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
	}

	serializer := BooksSerializer{c, books}
	c.JSON(http.StatusOK, gin.H{"books": serializer.Response()})
}
