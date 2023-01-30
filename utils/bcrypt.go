package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Generate Hash by doing 14 Rounds.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Compare Given Password with Database Hashed Password.
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}