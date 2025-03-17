package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqliteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("socket_project.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}
