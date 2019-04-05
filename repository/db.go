package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDBConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:book@tcp(book_db:3306)/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//マイグレーション
	//db.AutoMigrate(&Book{})

	return db
}
