package models

import (
	"github.com/isuryanarayanan/go-rest-api/settings"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name  string
	Email string
}

func init() {
	db := settings.GetDatabase()
	db.AutoMigrate(&User{})
}
