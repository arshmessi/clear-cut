package auth

import (
	"time"

	"clear-cut/config"

	"github.com/golang-jwt/jwt/v4"
)

// Define the JWT key from the configuration
var jwtKey = []byte(config.AppConfig.JWTSecret)

// Claims defines the structure of the JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT creates a new JWT token with a 24-hour expiration
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT parses and validates the JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
