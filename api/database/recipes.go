package database

import (
	"RecipeBookApi/models"
	"log"
)

func CreateRecipe(recipe models.RecipeCreate) error {
	stmt, err := DB.Prepare("INSERT INTO recipes (name, description, ingredients, steps, image_path, creator_id, is_private) VALUES (?, ?, ?, ?, ?, (SELECT id FROM users WHERE username = ?), ?)")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(recipe.Name, recipe.Description, recipe.Ingredients, recipe.Steps, recipe.ImagePath, recipe.CreatorUsername, recipe.IsPrivate)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetRecipeByID(id int) (*models.RecipeGet, error) {
	recipe := &models.RecipeGet{}

	query := DB.QueryRow("SELECT id, name, description, ingredients, steps, image_path FROM recipes WHERE id = ? AND is_private = false", id)
	err := query.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.Steps, &recipe.ImagePath)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

func GetRecipes() ([]*models.RecipeGet, error) {
	recipes := make([]*models.RecipeGet, 0)

	rows, err := DB.Query("SELECT id, name, description, ingredients FROM recipes WHERE is_private = false;")
	if err != nil {
		return recipes, err
	}
	defer rows.Close()

	for rows.Next() {
		recipe := &models.RecipeGet{}

		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients); err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}
