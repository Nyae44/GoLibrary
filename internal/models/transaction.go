package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID           uint      `gorm:"not null" json:"user_id"`
	User             User      `gorm:"foreignkey:UserID" json:"user"`
	MemberID         uint      `gorm:"not null" json:"member_id"`
	Member           Member    `gorm:"foreignkey:MemberID" json:"member"`
	BookID           uint      `gorm:"not null" json:"book_id"`
	Book             Book      `gorm:"foreignkey:BookID" json:"book"`
	IssueDate        time.Time `gorm:"not null" json:"issue_date"`
	ReturnDate       time.Time `gorm:"not null" json:"return_date"`
	ActualReturnDate time.Time `gorm:"not null" json:"actual_return_date"`
	FeesCharged      float64   `gorm:"not null" json:"fees_charged"`
	Penalty          float64   `gorm:"not null" json:"penalty"`
	Type             string    `gorm:"not null" json:"type"`
}
