package database

import (
	"RecipeBookApi/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func InitDB() {
	var err error

	DB, err = sql.Open("mysql", config.ConnectionString)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Println("Successfully connected to MySQL")
}