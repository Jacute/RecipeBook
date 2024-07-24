package database

import (
	"RecipeBookApi/models"
	"fmt"
	"log"
)

func GetRecipes() ([]*models.RecipeGet, error) {
	recipes := make([]*models.RecipeGet, 0)

	rows, err := DB.Query(`
		SELECT r.id, r.name, r.description, r.ingredients, r.steps, r.image_path, u.username
		FROM recipes r
		JOIN users u ON r.creator_id = u.id
		WHERE r.is_private = false;
	`)
	if err != nil {
		return recipes, err
	}
	defer rows.Close()

	for rows.Next() {
		recipe := &models.RecipeGet{}

		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.Steps, &recipe.ImagePath, &recipe.CreatorUsername); err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}

func GetRecipeByID(id int) (*models.RecipeGet, error) {
	recipe := &models.RecipeGet{}

	query := DB.QueryRow(`
		SELECT r.id, r.name, r.description, r.ingredients, r.steps, r.image_path, u.username
		FROM recipes r
		JOIN users u ON r.creator_id = u.id
		WHERE r.id = ? AND r.is_private = false
	`, id)
	err := query.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.Steps, &recipe.ImagePath, &recipe.CreatorUsername)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

func GetRecipeByName(name string) (*models.RecipeGet, error) {
	recipe := &models.RecipeGet{}

	query := DB.QueryRow(`
		SELECT r.id, r.name, r.description, r.ingredients, r.steps, r.image_path, u.username
		FROM recipes r
		JOIN users u ON r.creator_id = u.id
		WHERE r.name = ? AND r.is_private = false
	`, name)
	err := query.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.Steps, &recipe.ImagePath, &recipe.CreatorUsername)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

func CreateRecipe(recipe models.RecipeCreate) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = GetRecipeByName(recipe.Name)
	if err == nil {
		return &AlreadyExistsError{
			Message: fmt.Sprintf("%s already exists", recipe.Name),
		}
	}

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

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
