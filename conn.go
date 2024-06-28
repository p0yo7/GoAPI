// conn.go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func Connect() *gorm.DB {
	user := os.Getenv("user")
	password := os.Getenv("password")
	ip := os.Getenv("ip")
	port := os.Getenv("port")
	dbname := os.Getenv("db")

	// Validate environment variables
	if user == "" || password == "" || ip == "" || port == "" || dbname == "" {
		panic("Missing one or more environment variables: user, password, ip, port, db")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbname)
	fmt.Println("DSN:", dsn) // Add this line for debugging
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getDB() *gorm.DB {
	return db
}
