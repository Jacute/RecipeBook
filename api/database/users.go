package database

import (
	"RecipeBookApi/models"
	"RecipeBookApi/utils"
)

func AddUser(username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := DB.QueryRow("SELECT username, email, password FROM users WHERE username = ?", username)
	err := query.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := DB.QueryRow("SELECT username, email, password FROM users WHERE email = ?", email)
	err := query.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
