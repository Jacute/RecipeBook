package database

import (
	"RecipeBookApi/config"
	"RecipeBookApi/models"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func InitDB() {
	connectMySQL()
	createDB()
	fillRecipes()
}

func connectMySQL() {
	var err error

	connectionString := config.GetMySQL()

	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Println("Successfully connected to MySQL")
}

func createDB() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(64) UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		panic(err.Error())
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS recipes (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) UNIQUE NOT NULL,
        description TEXT NOT NULL,
        ingredients TEXT NOT NULL,
		steps TEXT NOT NULL,
		image_path TEXT NOT NULL,
		creator_id INTEGER,
		is_private BOOLEAN DEFAULT TRUE,

		FOREIGN KEY (creator_id) REFERENCES users(id)
	);`)
	if err != nil {
		panic(err.Error())
	}
}

func fillRecipes() {
	bytes, err := os.ReadFile("recipes.json")
	if err != nil {
		panic(err)
	}

	var recipes []models.RecipeCreate
	err = json.Unmarshal(bytes, &recipes)
	if err != nil {
		panic(err)
	}

	for _, recipe := range recipes {
		recipe.CreatorUsername = "admin"
		recipe.IsPrivate = false

		err = CreateRecipe(recipe)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
