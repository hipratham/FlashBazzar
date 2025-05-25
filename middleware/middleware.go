package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies Firebase ID tokens and sets user info in the context.
func AuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		client, err := app.Auth(context.Background())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firebase Auth client"})
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Get user record to potentially access custom claims like 'role'
		user, err := client.GetUser(context.Background(), token.UID)
		if err != nil {
			// Log this error but don't necessarily abort, token is verified
			// Depending on requirements, you might want to abort if user record is essential
		}

		c.Set("uid", token.UID)
		c.Set("claims", token.Claims)
		c.Set("user", user) // Optionally set the full user record
		c.Next()
	}
}
package middleware

// This file will contain middleware functions for the application.
// Examples: authentication, logging, rate limiting.

// RoleGuard checks if the authenticated user has the required role.
func RoleGuard(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User claims not found in context. Ensure AuthMiddleware is used."})
			return
		}

		userClaims, ok := claims.(map[string]interface{})
		userRole, ok := userClaims["role"].(string)

		if ok && userRole == requiredRole {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		}
	}
}