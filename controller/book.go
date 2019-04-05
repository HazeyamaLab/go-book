package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mytheta/gin_api/model"
	"github.com/mytheta/gin_api/service"
)

type BookController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(s service.BookService) BookController {
	return &bookController{bookService: s}
}

func (b *bookController) Create(ctx *gin.Context) {

	var book model.Book
	err := ctx.BindJSON(&book)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	err = b.bookService.Create(book)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "本登録成功",
	})
}

func (b *bookController) FindAll(ctx *gin.Context) {
	books, err := b.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b *bookController) Update(ctx *gin.Context) {
	var book model.Book
	err := ctx.BindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	err = b.bookService.Update(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "編集できました,"})

}

func (b *bookController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ok, err := b.bookService.IsExistByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
	}

	err = b.bookService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "削除成功"})
}
