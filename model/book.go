package model

import "time"

type Book struct {
	ID                 uint      `gorm:"primary_key" json:"id" `
	Title              string    `json:"title" binding:"required,max=127"`
	Author             string    `json:"author"`
	Price              int       `json:"price"`
	DateOfPublication  time.Time `json:"date_of_publication"`
}
