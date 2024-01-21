package handlers

import (
	"database/sql"
	"fmt"
)

type AddUserInput struct {
	Username string
	Password string
}

func AddUser(db *sql.DB, user AddUserInput) error {
	_, err := db.Exec("INSERT INTO users(username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err == nil {
		fmt.Println("Add user success")
	}
	return err
}
