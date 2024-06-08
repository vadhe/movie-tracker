package main

import (
	"fmt"
	"movie-tracker/internal/database"
	"movie-tracker/internal/model"
)

func User() {
	user := model.User{
		Username: "admin",
		Email:    "admin@gmail.com",
		Password: "admin123",
	}
	dbService := database.New()
	_, err := model.AddUser(dbService.DB(), user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully added user")
	}
}
