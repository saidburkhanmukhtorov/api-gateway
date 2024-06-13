package middleware

import (
	"net/http"

	"github.com/Project_Restaurant/api-gateway/api/token"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := token.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"hello": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected area"})
	c.Next()
}
