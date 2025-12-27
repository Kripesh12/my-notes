package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(byte), err
}

func ComparePassword(hash, password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

