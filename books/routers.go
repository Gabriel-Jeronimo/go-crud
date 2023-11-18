package books

import (
	"errors"
	"go-crud/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BooksRegister(router *gin.RouterGroup) {
	router.POST("/", BookCreate)
	router.GET("/", BookList)
	router.GET("/:id", BookById)
	router.DELETE("/:id", BookDelete)
}

func BookById(c *gin.Context) {
	id := c.Param("id")

	book, err := FindByIdBook(id)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("books", errors.New("book not found")))
	}

	serializer := BookSerializer{c, book}

	c.JSON(http.StatusOK, gin.H{"book": serializer.Response()})
}

func BookCreate(c *gin.Context) {
	bookModelValidator := NewBookModelValidator()
	if err := bookModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&bookModelValidator.bookModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}

	serializer := BookSerializer{c, bookModelValidator.bookModel}
	c.JSON(http.StatusCreated, gin.H{"data": serializer.Response()})
}

func BookList(c *gin.Context) {
	books, err := FindManyBook()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	serializer := BooksSerializer{c, books}
	c.JSON(http.StatusOK, gin.H{"books": serializer.Response()})
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	err := DeleteBook(&BookModel{ID: id})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("books", errors.New("book not found")))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func SaveOne(data interface{}) error {

	db := common.GetDB()
	err := db.Save(data).Error
	return err
}
