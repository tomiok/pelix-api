package users

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// Combine password and salt.
func hashPassword(pwd string) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil

}

// Check if two passwords match
func doPasswordsMatch(hashedPassword, plainPassword string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string, so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword))
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}

	return true
}
