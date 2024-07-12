package main

import (
	"RecipeBookApi/database"
	"RecipeBookApi/handlers"
	"RecipeBookApi/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	sessionSecret, err := utils.GenerateRandomBytes(64)
	if err != nil {
		panic(err)
	}
	store := cookie.NewStore(sessionSecret)
	router.Use(sessions.Sessions("session-name", store))

	router.GET("/api/recipes", handlers.GetRecipes)
	router.GET("/api/recipes/:id", handlers.GetRecipeByID)

	router.GET("/api/auth/whoami", handlers.GetWhoami)
	router.POST("/api/auth/login", handlers.PostLogin)
	router.POST("/api/auth/register", handlers.PostRegister)
	router.GET("/api/auth/logout", handlers.GetLogout)

	log.Fatal(http.ListenAndServe(":8000", router))
}
