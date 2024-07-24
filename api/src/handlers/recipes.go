package handlers

import (
	"RecipeBookApi/config"
	"RecipeBookApi/database"
	"RecipeBookApi/models"
	"RecipeBookApi/utils"
	"encoding/json"
	"path"
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

func CreateRecipe(c *gin.Context) {
	recipeJson := c.Request.FormValue("recipe")
	if recipeJson == "" {
		c.String(400, "Recipe is required")
		return
	}

	var recipe models.RecipeCreate
	if err := json.Unmarshal([]byte(recipeJson), &recipe); err != nil {
		c.String(400, "Invalid request body")
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.String(400, "Photo is required")
		return
	}

	if file.Size > config.FileSizeUploadLimit {
		c.String(400, "File is too large")
		return
	}

	mimeType := file.Header.Get("Content-Type")
	extension, err := utils.GetExtension(file.Filename)
	if err != nil || !utils.IsValidImage(mimeType, extension) {
		c.String(400, "File should be image")
		return
	}

	savePath := path.Join("images", utils.RandomString(32)) + extension

	recipe.ImagePath = savePath
	recipe.CreatorUsername = c.GetString("username")
	err = database.CreateRecipe(recipe)
	if alreadyExistsErr, ok := err.(*database.AlreadyExistsError); ok {
		c.String(400, alreadyExistsErr.Message)
		return
	} else if err != nil {
		c.String(500, "Error creating recipe")
		return
	}

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(500, "Error saving uploaded file")
		return
	}

	c.String(201, "Recipe created successfully")
}
