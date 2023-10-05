package books

import (
	"go-crud/common"

	"github.com/gin-gonic/gin"
)

type BookModelValidator struct {
	Book struct {
		Title       string `form:"title" json:"title" binding:"required,min=4"`
		Description string `form:"description" json:"description" binding:"max=2048"`
		Author      string `form:"author" json:"author" binding:"required,min=4"`
	} `json:"book"`
	bookModel BookModel `json:"-"`
}

func NewBookModelValidator() BookModelValidator {
	return BookModelValidator{}
}

func NewBookmodelValidatorFillWith(bookModel BookModel) BookModelValidator {
	bookModelValidator := NewBookModelValidator()
	bookModelValidator.Book.Author = bookModel.Title
	bookModelValidator.Book.Description = bookModel.Description
	bookModelValidator.Book.Author = bookModel.Author

	return bookModelValidator
}

func (s *BookModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)

	if err != nil {
		return err
	}

	s.bookModel.Title = s.Book.Title
	s.bookModel.Description = s.Book.Description
	s.bookModel.Author = s.Book.Author

	return nil
}
