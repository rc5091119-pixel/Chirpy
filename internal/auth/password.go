package auth

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	str, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "",fmt.Errorf("could not hash password: %w",err)
	}
	return str, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	check, err := argon2id.ComparePasswordAndHash(password, hash)
	return check, err
}
