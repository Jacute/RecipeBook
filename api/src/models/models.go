package models

type Recipe struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Ingredients     string `json:"ingredients"`
	Steps           string `json:"steps"`
	ImagePath       string `json:"image_path"`
	CreatorUsername string `json:"creator_username"`
}

type RecipeGet struct {
	Recipe
}

type RecipeCreate struct {
	Recipe
	IsPrivate bool `json:"is_private"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
