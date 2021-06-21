package settings

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// ConnectDatabase creates a connection with sqlite database
func ConnectDatabase() {
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

// GetDatabase return's the connection instance
func GetDatabase() *gorm.DB {
	return db
}
