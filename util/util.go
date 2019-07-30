package util

import "github.com/HazeyamaLab/go-book/model"

func TotalPrice(books []model.Book) (int, error) {
	var totalPrice int
	for _, book := range books {
		totalPrice += book.Price
	}

	return totalPrice, nil
}

//func CarTotalPrice(cars []model.Car) (int, error) {
//	var carTotalPrice int
//	for _, car := range cars {
//		carTotalPrice += car.Value
//	}
//
//	return carTotalPrice, nil
//}
