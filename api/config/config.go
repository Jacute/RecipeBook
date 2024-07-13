package config

import (
	"fmt"
	"os"
)

type MySQLConfig struct {
	database string
	username string
	password string
	hostname string
}

func GetMySQL() string {
	databaseConfig := MySQLConfig{
		database: os.Getenv("MYSQL_DATABASE"),
		username: os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		hostname: os.Getenv("DATABASE_HOSTNAME"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", databaseConfig.username, databaseConfig.password, databaseConfig.hostname, databaseConfig.database)
	// connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", databaseConfig.username, databaseConfig.password, databaseConfig.database)
	return connectionString
}
