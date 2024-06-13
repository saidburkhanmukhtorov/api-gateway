package middlewares

import (
	"net/http"

	"github.com/Project_Restaurant/api-gateway/api/tokens"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]
	err := tokens.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hello": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}

func Role(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if claims["username"] != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication error"})
		c.Abort()
		return
	}
	c.Next()
}
