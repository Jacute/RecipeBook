package handlers

import (
	"RecipeBookApi/database"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetRecipes(c *gin.Context) {
	recipes, err := database.GetRecipes()
	if err != nil {
		c.String(500, "Error fetching recipes")
		return
	}
	c.JSON(200, recipes)
}

func GetRecipeByID(c *gin.Context) {
	session := sessions.Default(c)
	username, ok := session.Get("username").(string)
	if !ok {
		username = ""
	}

	recipeIDStr := c.Param("id")
	recipeID, err := strconv.Atoi(recipeIDStr)
	if err != nil {
		c.String(400, "Bad ID")
		return
	}

	recipe, err := database.GetRecipeByID(recipeID, username)
	if err != nil {
		c.String(404, "Recipe not found")
		return
	}
	c.JSON(200, recipe)
}
