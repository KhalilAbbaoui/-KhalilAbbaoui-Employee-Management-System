package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
    "employee-management-system/utils" // Import your JWT utility functions
)

// JWTAuthMiddleware is a middleware for protecting routes with JWT
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the token from the Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        // Extract token from the header (assuming format: "Bearer <token>")
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is missing"})
            c.Abort()
            return
        }

        // Validate the token
        _, _, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Token is valid, proceed to the next handler
        c.Next()
    }
}
