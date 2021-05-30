package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	// Auto Migrate
	err = db.AutoMigrate(&Student{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
