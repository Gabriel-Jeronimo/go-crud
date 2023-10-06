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