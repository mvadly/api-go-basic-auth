package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	Username := os.Getenv("DB_USERNAME")
	Password := os.Getenv("DB_PASSWORD")
	Hostname := os.Getenv("DB_HOSTNAME")
	DbName := os.Getenv("DB_DATABASE")

	// connection := [4]string{Username, Password, Hostname, DbName}
	fmt.Print("user : ", Username)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Hostname, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database not connected")
	}

	return db

}
