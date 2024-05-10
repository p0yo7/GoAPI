// controllers.go

package GoAPI

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"GoAPI/main"
)

func generateToken(id int, user_role string) string {
	secretKey := []byte("secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"user_role": user_role,
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err.Error())
	}
	return tokenString
}

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data loginData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	username := data.Username
	password := data.Password
	db := main.getDB()
	defer db.Close()

	// Retrieve the hashed password from the database
	var id int
	var user_role string
	var hashedPassword string
	err = db.QueryRow(`SELECT id_user, user_role, password FROM users WHERE username = ?`, username).Scan(&id, &user_role, &hashedPassword)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare the hashed password with the plaintext password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Passwords match, generate token
	token := generateToken(id, user_role)

	// Send token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
	db.Close()
}
