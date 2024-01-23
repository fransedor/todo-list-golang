package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) []byte {
	inputAsByteArray := []byte(input)
	hashedInput, err := bcrypt.GenerateFromPassword(inputAsByteArray, 5)
	if err != nil {
		fmt.Println("Error hashing input")
	}
	return hashedInput
}
