package main

import (
	"RecipeBookApi/database"
	"RecipeBookApi/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env_db")

	database.InitDB()
	defer database.DB.Close()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api/recipes", handlers.GetRecipes)
	router.GET("/api/recipes/:id", handlers.GetRecipeByID)

	log.Fatal(http.ListenAndServe(":8000", router))
}
