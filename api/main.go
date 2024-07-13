package main

import (
	"RecipeBookApi/database"
	"RecipeBookApi/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api/recipes", handlers.GetRecipes)
	router.GET("/api/recipes/:id", handlers.GetRecipeByID)
	router.GET("/api/images/:filename", handlers.GetImage)

	log.Fatal(http.ListenAndServe(":8000", router))
}
