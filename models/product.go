package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string          `json:"name" gorm:"unique;not null" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Image       string          `json:"image"`
	Price       decimal.Decimal `json:"price" gorm:"not null;type:numeric" binding:"required"`
}
