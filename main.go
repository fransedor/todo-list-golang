package main

import (
	"fmt"
	"log"

	"github.com/fransedor/todo-list-golang/database"
	handlers "github.com/fransedor/todo-list-golang/handlers/users"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()
	users := handlers.GetUsers(db)

	fmt.Printf("Found %d users\n", len(users))
	fmt.Println("Enter new username: ")
	var newUsername string
	fmt.Scanln(&newUsername)
	fmt.Println("Enter new password: ")
	var newPassword string
	fmt.Scanln(&newPassword)
	err = handlers.AddUser(db, handlers.AddUserInput{Username: newUsername, Password: newPassword})
	if err != nil {
		fmt.Printf("Error adding user: %v", err)
	}
	db.Close()
}
