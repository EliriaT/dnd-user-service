package server

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash the password")
	}
	return string(hashedPassword), nil
}

func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
