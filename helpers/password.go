package helpers

import (
	"user_management_service/config"
	
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}

func IsValidPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
