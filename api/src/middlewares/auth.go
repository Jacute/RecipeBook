package middlewares

import (
	"RecipeBookApi/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenCookie, err := c.Cookie("auth-token")
		if tokenCookie == "" || err != nil {
			c.String(401, "Unauthorized")
			c.Abort()
			return
		}
		token, err := utils.VerifyToken(tokenCookie)
		if err != nil {
			c.String(401, err.Error())
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
		} else {
			c.String(401, "Invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
