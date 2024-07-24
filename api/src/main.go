package main

import (
	"RecipeBookApi/config"
	"RecipeBookApi/database"
	"RecipeBookApi/handlers"
	"RecipeBookApi/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.New()

	router.MaxMultipartMemory = config.MaxMultipartMemory

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api/recipes", handlers.GetRecipes)
	router.GET("/api/recipes/:id", handlers.GetRecipeByID)
	router.GET("/api/images/:filename", handlers.GetImage)
	router.POST("/api/recipes", middlewares.AuthMiddleware(), handlers.CreateRecipe) // TODO: route for authorized

	return router
}

func main() {
	database.InitDB()
	defer database.DB.Close()

	router := setupRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
