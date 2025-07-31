package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"gofr.dev/pkg/gofr"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// ValidateJWT validates JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

// AuthMiddleware is a placeholder for JWT authentication
func AuthMiddleware(c *gofr.Context) (interface{}, error) {
	// This is a basic middleware structure for GoFr
	// In actual implementation, this would validate JWT tokens
	return map[string]string{"status": "middleware_ready"}, nil
}