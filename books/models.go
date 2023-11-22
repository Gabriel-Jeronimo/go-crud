package books

import (
	"go-crud/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookModel struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string `gorm:"size:2048"`
	Author      string
}

func (bookModel *BookModel) BeforeCreate(tx *gorm.DB) (err error) {
	bookModel.ID = uuid.NewString()
	return
}

func UpdateBook(id string, book *BookModel) error {
	db := common.GetDB()

	err := db.Model(&BookModel{ID: id}).Updates(book).Error
	return err
}

func FindByIdBook(id string) (BookModel, error) {
	db := common.GetDB()

	var model BookModel

	err := db.Find(&model, BookModel{ID: id}).Error

	return model, err
}

func FindManyBook() ([]BookModel, error) {
	db := common.GetDB()

	var models []BookModel

	db.Find(&models)

	return models, nil
}

func DeleteBook(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(BookModel{}).Error
	return err
}
