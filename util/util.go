package util

import "github.com/HazeyamaLab/go-book/model"

func TotalPrice(books []model.Book) (int, error) {
	var totalPrice int
	for _, book := range books {
		totalPrice += book.Price
	}

	return totalPrice, nil
}
