package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// open connection with the darabase

func OpenDB() *gorm.DB {
	if DB == nil {
		DB, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		return DB
	}
	return DB
}
