package handlers

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       string
	Username string
	Password string
}

func GetUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		fmt.Printf("Error query users: %v", err)
	}
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			fmt.Printf("Error adding user to struct: %v", err)
			rows.Close()
			log.Fatal("Exiting rows scan")
		}
		users = append(users, user)
	}
	return users
}
