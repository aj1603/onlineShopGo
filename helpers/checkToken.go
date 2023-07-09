package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckToken(c *gin.Context) (jwt.MapClaims, map[string]any) {
	tokenString := c.Request.Header.Get("Authorization")

	if tokenString == "" {
		return nil, gin.H{"error": "No authorization header provided"}
	}

	splitToken := strings.Split(tokenString, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	} else {
		return nil, gin.H{"error": "Invalid Token"}
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, gin.H{"error": "Invalid token signature"}
		} else {
			return nil, gin.H{"error": "Invalid token"}
		}

	}

	if !token.Valid {
		return nil, gin.H{"error": "Invalid token"}
	}

	return claims, nil
}
