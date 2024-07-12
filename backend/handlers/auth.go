package handlers

import (
	"RecipeBookApi/database"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostRegister(c *gin.Context) {
	session := sessions.Default(c)

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.String(400, "Missing username or password")
		return
	}

	_, err := database.GetUserByUsername(username)
	if err == nil {
		c.String(409, "Username already exists")
		return
	}

	_, err = database.GetUserByEmail(email)
	if err == nil {
		c.String(409, "Email already exists")
		return
	}

	err = database.AddUser(username, email, password)
	if err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}
	session.Set("username", username)
	session.Save()
	c.Status(200)
}

func PostLogin(c *gin.Context) {
	session := sessions.Default(c)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.String(400, "Missing username or password")
		return
	}

	user, err := database.GetUserByUsername(username)
	if err != nil {
		c.String(401, "Invalid username or password")
		return
	}
	session.Set("username", user.Name)
	session.Save()
	c.Status(200)
}

func GetLogout(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("User logged out:", session.Get("username"))
	session.Delete("username")
	c.Redirect(301, "/")
}

func GetWhoami(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	usernameStr, _ := username.(string)
	c.String(200, usernameStr)
}
