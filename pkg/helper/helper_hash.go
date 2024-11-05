package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string,error){

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("error Hashing Password")
	}

	return string(hashedPassword),nil

}

func CheckPasswordHash(password string, hash string) bool {

  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  

	return err == nil

}
