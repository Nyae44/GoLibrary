package models

import (
	"gorm.io/gorm"
	"time"
)

// Book model
type Book struct {
	gorm.Model
	BookID            int            `json:"id" gorm:"primary_key"`
	Title             string         `json:"title" gorm:"not null"`
	Author            string         `json:"author" gorm:"not null"`
	ISBN              string         `json:"isbn" gorm:"unique"`
	Quantity          int            `json:"quantity"`
	RemainingQuantity int            `json:"remaining_quantity"`
	RentalFee         int            `json:"rental_fee"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
