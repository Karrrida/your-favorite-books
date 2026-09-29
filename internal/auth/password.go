package auth

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const cost int = 10

func Hash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return "", fmt.Errorf("password must be less then 72 characters")
		}
		return "", err
	}
	hash := string(passwordHash)
	return hash, nil
} 