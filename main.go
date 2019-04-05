package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/gin_api/controller"
	"github.com/mytheta/gin_api/repository"
	"github.com/mytheta/gin_api/service"
)

func main() {

	//set up
	bookController := bookInjector()

	// server 起動
	r := gin.Default()

	//routing
	r.POST("/book", bookController.Create)
	r.GET("/book", bookController.FindAll)
	r.PUT("/book/@:id", bookController.Update)
	r.DELETE("/book/@:id", bookController.Delete)

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
