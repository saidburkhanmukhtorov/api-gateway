package token

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// claims, ok := token.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return fmt.Errorf("invalid token claims format")
	// }

	// for key, value := range claims {
	// 	log.Printf("%s: %v\n", key, value)
	// }
	return nil
}
