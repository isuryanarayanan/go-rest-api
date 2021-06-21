package models

import (
	"github.com/isuryanarayanan/go-rest-api/settings"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Title  string
	Author User `gorm:"foreignKey:ID"`
}

func init() {
	db := settings.GetDatabase()
	db.AutoMigrate(&Book{})
}
