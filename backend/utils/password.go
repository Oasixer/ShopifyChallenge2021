package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPass: hashes a password with a salt
// params: password string
// returns: string or err, the hashed password,
func HashPass(pass string) (string, error) {
	bytePass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Hash error: %s", err)
		return "", errors.New("coulndt hash password")
	}

	return string(hash), nil
}

// CompareHashPass: compares a hash and a password
// params: password string, hash string
// returns: bool, true if they match, false otherwise
func CompareHashPass(pass string, hash string) bool {
	byteHash := []byte(hash)
	bytePass := []byte(pass)

	// returns an err if false
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)

	if err != nil {
		return false
	}

	return true
}
