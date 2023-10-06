package books

import (
	"go-crud/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BooksRegister(router *gin.RouterGroup) {
	router.POST("/", BookCreate)
	router.GET("/", BookList)
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

func SaveOne(data interface{}) error {

	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func BookList(c *gin.Context) {
	books, err := FindManyBook()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal"})
	}

	serializer := BooksSerializer{c, books}
	c.JSON(http.StatusOK, gin.H{"books": serializer.Response()})
}
