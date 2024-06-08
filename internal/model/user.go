package model

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Email    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func AddUser(db *sql.DB, user User) (User, error) {
	passwordHash, errPassword := HashPassword(user.Password)
	if errPassword != nil {
		fmt.Println("failed to insert user: %v", errPassword)
	}
	_, err := db.Exec("INSERT INTO users (username,email,password_hash) VALUES  ($1,$2,$3)", user.Username, user.Email, passwordHash)
	if err != nil {
		return User{}, fmt.Errorf("failed to insert user: %v", err)
	}
	return user, nil
}
