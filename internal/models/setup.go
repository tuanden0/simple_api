package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	// Auto Migrate
	db.AutoMigrate(&Department{}, &Course{}, &Group{}, &Student{})

	DB = db
}
