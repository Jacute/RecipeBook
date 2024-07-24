package utils

import (
	"math/rand"
	"os"
	"regexp"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

var randomAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var ExtensionRegexp = regexp.MustCompile("^([^/]+)(\\..+)$")

func RandomString(length int) string {
	result := make([]uint8, length)
	for i := 0; i < length; i++ {
		result[i] = randomAlphabet[rand.Intn(len(randomAlphabet))]
	}
	return string(result)
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetExtension(filename string) (string, error) {
	if !ExtensionRegexp.Match([]byte(filename)) {
		return "", &CustomError{
			Message: "Incorrect filename",
		}
	}
	filenameParts := ExtensionRegexp.FindStringSubmatch(filename)
	return filenameParts[len(filenameParts)-1], nil
}
