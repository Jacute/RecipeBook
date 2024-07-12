package database

import "log"

func CreateRecipe(name string, description string, ingredients string, creatorUsername string, isPrivate bool) error {
	stmt, err := DB.Prepare("INSERT INTO recipes (name, description, ingredients, creator_id, is_private) VALUES (?, ?, ?, (SELECT id FROM users WHERE username = ?), ?)")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(name, description, ingredients, creatorUsername, isPrivate)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetRecipeByID(id int, creatorUsername string) (*Recipe, error) {
	recipe := &Recipe{}

	query := DB.QueryRow("SELECT id, name, description, ingredients FROM recipes WHERE id = ? AND (is_private = false OR creator_id = (SELECT id FROM users WHERE username = ?))", id, creatorUsername)
	err := query.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Ingredients)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

func GetRecipes() ([]*Recipe, error) {
	recipes := make([]*Recipe, 0)

	rows, err := DB.Query("SELECT id, name, description, ingredients FROM recipes WHERE is_private = false;")
	if err != nil {
		return recipes, err
	}
	defer rows.Close()

	for rows.Next() {
		recipe := &Recipe{}

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
