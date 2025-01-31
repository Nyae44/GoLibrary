package models

import (
	"gorm.io/gorm"
	"time"
)

type Member struct {
	gorm.Model
	UserID      uint    `gorm:"not null" json:"user_id"`
	User        User    `gorm:"foreignkey:UserID" json:"user"`
	PhoneNumber string  `gorm:"unique" json:"phone_number"`
	Email       string  `gorm:"unique" json:"email"`
	RentalDebt  float64 `gorm:"default:0" json:"rental_debt"`
	Balance     float64 `gorm:"default:0" json:"balance"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (m *Member) AddToRentalDebt(db *gorm.DB, amount float64) error {
	m.RentalDebt += amount
	return db.Save(m).Error
}
