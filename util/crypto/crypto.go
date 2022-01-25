package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	return HashString(password)
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func HashString(input string) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		return nil
	}

	return bytes
}
