package function

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(secret string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func CheckHash(hash string, secret string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err
}
