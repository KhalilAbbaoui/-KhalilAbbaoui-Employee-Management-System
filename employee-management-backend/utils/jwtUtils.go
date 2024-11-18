package utils

import (
	"fmt"
	//"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)



var secretKey = []byte("your-secret-key") // Use an environment variable or more secure method in production

// GenerateToken generates a JWT for a given user ID
func GenerateToken(userID string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(24 * time.Hour) // token expires in 24 hours

	// Create a new JWT token with claims
	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	// Create a token using the HMAC signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken verifies the JWT token and returns the claims if valid.
func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	// Parse the token and extract claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key for validating the token
		return secretKey, nil
	})

	// Handle errors during token parsing
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing token: %v", err)
	}

	// Ensure the token is valid
	if !token.Valid {
		return nil, nil, fmt.Errorf("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, nil, fmt.Errorf("invalid claims type")
	}

	// Return the token and claims
	return token, *claims, nil
}
