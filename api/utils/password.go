package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	return err == nil
}
