package handlers

import (
	"database/sql"
	"fmt"

	"github.com/fransedor/todo-list-golang/utils"
)

type AddUserInput struct {
	Username string
	Password string
}

func AddUser(db *sql.DB, user AddUserInput) error {
	hashedPassword := utils.Hash(user.Password)
	_, err := db.Exec("INSERT INTO users(username, password) VALUES ($1, $2)", user.Username, hashedPassword)
	if err == nil {
		fmt.Println("Add user success")
	}
	return err
}
