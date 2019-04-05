package main

import (
	"github.com/HazeyamaLab/go-book/controller"
	"github.com/HazeyamaLab/go-book/repository"
	"github.com/HazeyamaLab/go-book/service"
	"github.com/gin-gonic/gin"
)

func main() {

	//set up
	bookController := bookInjector()

	// server 起動
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")

	//routing
	r.GET("/", bookController.Index)
	r.POST("/create", bookController.Create)
	r.GET("/update/@:id", bookController.UpdateConfirm)
	r.POST("/update/@:id", bookController.Update)
	r.POST("/delete/@:id", bookController.Delete)
	r.GET("/delete/confirm/@:id", bookController.DeleteConfirm)

	//Listening and serving HTTP
	r.Run(":8000")
}

// bookの依存を解決します
func bookInjector() controller.BookController {
	db := repository.NewDBConn()
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)
	return bookController
}
