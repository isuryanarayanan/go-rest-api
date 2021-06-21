package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Book struct {
	gorm.Model
	Title  string
	Author User `gorm:"foreignKey:ID"`
}

func initDatabaseConnection() {
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func migrateModels() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Book{})
}

func main() {
	initDatabaseConnection()
	migrateModels()
}
