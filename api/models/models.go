package models

type RecipeGet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Steps       string `json:"steps"`
	ImagePath   string `json:"image_path"`
}

type RecipeCreate struct {
	RecipeGet
	CreatorUsername string `json:"-"`
	IsPrivate       bool   `json:"is_private"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
