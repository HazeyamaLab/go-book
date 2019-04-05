package controller

import (
	"log"
	"strconv"

	"github.com/HazeyamaLab/go-book/model"
	"github.com/HazeyamaLab/go-book/service"
	"github.com/HazeyamaLab/go-book/util"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	Index(ctx *gin.Context)
	Create(c *gin.Context)
	UpdateConfirm(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	DeleteConfirm(c *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(s service.BookService) BookController {
	return &bookController{bookService: s}
}

func (b *bookController) Index(ctx *gin.Context) {

	books, err := b.bookService.FindAll()
	if err != nil {
		ctx.HTML(500, "500.html", nil)
		return
	}

	totalPrice, err := util.TotalPrice(books)
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ctx.HTML(200, "view/index.html", gin.H{"books": books, "num": len(books), "sumPrice": totalPrice})
}

func (b *bookController) Create(ctx *gin.Context) {
	title := ctx.PostForm("title")

	//intに変換
	price, err := strconv.Atoi(ctx.PostForm("price"))
	if err != nil {
		panic(err)
	}

	book := model.Book{Title: title, Price: price}

	err = b.bookService.Create(book)
	if err != nil {
		log.Println(err)
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ctx.Redirect(302, "/")
}

func (b *bookController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		panic(err)
	}

	title := ctx.PostForm("title")

	//intに変換
	price, err := strconv.Atoi(ctx.PostForm("price"))
	if err != nil {
		panic(err)
	}

	book := model.Book{ID: uint(id), Title: title, Price: price}

	err = b.bookService.Update(book)
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ctx.Redirect(302, "/")
}

func (b *bookController) UpdateConfirm(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	book, err := b.bookService.FindOne(uint(id))
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ctx.HTML(200, "view/edit/index.html", gin.H{"book": book})
}

func (b *bookController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ok, err := b.bookService.IsExistByID(uint(id))
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}
	if !ok {
		ctx.HTML(500, "view/500.html", nil)
	}

	err = b.bookService.Delete(uint(id))
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ctx.Redirect(302, "/")
}

func (b *bookController) DeleteConfirm(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	book, err := b.bookService.FindOne(uint(id))
	if err != nil {
		ctx.HTML(500, "view/500.html", nil)
		return
	}

	ctx.HTML(200, "view/delete/index.html", gin.H{"book": book})
}
