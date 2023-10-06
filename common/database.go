package common

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./application.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Database Error: (INIT) ", err)
	}

	db.DB()
	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
