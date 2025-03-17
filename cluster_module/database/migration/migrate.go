package main

import (
	"fmt"
	models "socket_project/cluster_module/Models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSqliteConnectionForMigration() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("socket_project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {

	dbConnection := GetSqliteConnectionForMigration()

	if dbConnection == nil {
		fmt.Println("Expected db connection to be not nil, got nil")
		return
	}

	err := dbConnection.AutoMigrate(&models.Cluster{})

	if err != nil {
		fmt.Println("Error while migrating the table")
		return
	}

	fmt.Println("Table migrated successfully")

}
