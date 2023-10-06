package books

import (
	"go-crud/common"

	"github.com/gin-gonic/gin"
)

type BookModelValidator struct {
	Title       string `form:"title" json:"title" binding:"required,min=4"`
	Description string `form:"description" json:"description" binding:"max=2048"`
	Author      string `form:"author" json:"author" binding:"required,min=4"`

	bookModel BookModel `json:"-"`
}

func NewBookModelValidator() BookModelValidator {
	return BookModelValidator{}
}

func NewBookmodelValidatorFillWith(bookModel BookModel) BookModelValidator {
	bookModelValidator := NewBookModelValidator()
	bookModelValidator.Author = bookModel.Title
	bookModelValidator.Description = bookModel.Description
	bookModelValidator.Author = bookModel.Author

	return bookModelValidator
}

func (s *BookModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)

	if err != nil {
		return err
	}

	s.bookModel.Title = s.Title
	s.bookModel.Description = s.Description
	s.bookModel.Author = s.Author

	return nil
}
