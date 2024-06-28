// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
	// "golang.org/x/time/rate"
)

type Message struct {
	Message string `json:"message"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	msg := Message{"Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db = Connect() // Ensure db is initialized here
	fmt.Println("Connected to the database")
	fmt.Println("Server started on: http://localhost:8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/create-user", CreateUserController)

	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	http.ListenAndServe(":8080", rateLimit(csrfMiddleware(mux)))
}
