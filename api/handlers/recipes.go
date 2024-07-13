package handlers

import (
	"RecipeBookApi/database"
	"RecipeBookApi/utils"
	"path/filepath"
	"strconv"

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
	recipeIDStr := c.Param("id")
	recipeID, err := strconv.Atoi(recipeIDStr)
	if err != nil {
		c.String(400, "Bad ID")
		return
	}

	recipe, err := database.GetRecipeByID(recipeID)
	if err != nil {
		c.String(404, "Recipe not found")
		return
	}
	c.JSON(200, recipe)
}

func GetImage(c *gin.Context) {
	filename := c.Param("filename")
	if filename == "" {
		c.String(400, "Bad filename")
		return
	}
	imagePath := filepath.Join("images", filename)
	if !utils.FileExists(imagePath) {
		c.String(404, "Image not found")
		return
	}
	c.File(imagePath)
}
