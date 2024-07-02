package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash create a hash from a string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare password hash with a string
func VerifyPassword(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))

}
