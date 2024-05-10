package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"GoAPI/controllers" // Importing controllers from the same directory
	"GoAPI/models"      // Importing models from the same directory
)

var db *gorm.DB

func getDB() *gorm.DB {
	return db
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	msg := Message{"Hello World"} // Assuming Message struct is defined
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	user := os.Getenv("user")
	password := os.Getenv("password")
	ip := os.Getenv("ip")
	port := os.Getenv("port")
	dbName := os.Getenv("db")
	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbName
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.User{}, &models.ExercisesPlan{}, &models.Exercise{}, &models.RoutineExercise{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to the database")
	fmt.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", controllers.Login)
	http.ListenAndServe(":8080", nil)
}
