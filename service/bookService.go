package service

import (
	"github.com/mytheta/gin_api/model"
	"github.com/mytheta/gin_api/repository"
)

type BookService interface {
	Create(book model.Book) error
	FindAll() ([]model.Book, error)
	Update(book model.Book) error
	Delete(id uint) error
	IsExistByID(id uint) (bool, error)
}

type bookService struct {
	repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{r}
}

func (b *bookService) Create(book model.Book) error {
	return b.BookRepository.Create(book)
}

func (b *bookService) FindAll() ([]model.Book, error) {
	return b.BookRepository.FindAll()
}

func (b *bookService) Update(book model.Book) error {
	return b.BookRepository.Update(book)
}

func (b *bookService) Delete(id uint) error {
	return b.BookRepository.Delete(id)
}

func (b *bookService) IsExistByID(id uint) (bool, error) {
	return b.BookRepository.IsExistByID(id)
}
