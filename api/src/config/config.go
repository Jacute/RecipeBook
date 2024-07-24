package config

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

type MySQLConfig struct {
	database string
	username string
	password string
	hostname string
}

var databaseConfig = MySQLConfig{
	database: os.Getenv("MYSQL_DATABASE"),
	username: os.Getenv("MYSQL_USER"),
	password: os.Getenv("MYSQL_PASSWORD"),
	hostname: os.Getenv("DATABASE_HOSTNAME"),
}

var ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", databaseConfig.username, databaseConfig.password, databaseConfig.hostname, databaseConfig.database)
var FileSizeUploadLimit = int64(8 << 20) // 8 MB
var MaxMultipartMemory = int64(8 << 20)  // 8 MB

var JWT_PUBLIC_KEY_PATH = os.Getenv("JWT_PUBLIC_KEY_PATH")
var JWT_PRIVATE_KEY_PATH = os.Getenv("JWT_PRIVATE_KEY_PATH") // private key for tests
var JWT_PUBLIC_KEY = readJwtPublicKey()
var JWT_PRIVATE_KEY = readJwtPrivateKey()

func readJwtPublicKey() *rsa.PublicKey {
	keyData, err := os.ReadFile(JWT_PUBLIC_KEY_PATH)
	if err != nil {
		panic(err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		panic(err)
	}
	return key
}

func readJwtPrivateKey() *rsa.PrivateKey {
	keyData, err := os.ReadFile(JWT_PRIVATE_KEY_PATH)
	if err != nil {
		log.Printf("Failed to read private key file: %v\n", err)
		return nil
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Printf("Failed to parse private key file")
	}

	return key
}
