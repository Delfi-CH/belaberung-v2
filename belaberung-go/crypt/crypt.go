package crypt

import (
    "golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hashString, password string) bool {
	hash := []byte(hashString)
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}