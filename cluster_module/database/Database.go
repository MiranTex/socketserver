package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqliteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("socket_project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
