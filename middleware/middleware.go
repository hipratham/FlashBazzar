package middleware

import (
	"context"
	"net/http"
	"strings"

	firebaseauth "firebase.google.com/go/auth"
	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies Firebase ID tokens and sets user info in the context.
func AuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // Corrected extraction of Authorization header
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

		// Extract custom claims, including 'role'
		claims := token.Claims
		role, roleExists := claims["role"].(string)

		if !roleExists {
			// If role is essential and missing, you might want to handle this differently.
			// For now, we proceed but the RoleGuard will likely fail.
			// Log a warning: log.Printf("Role claim missing for user: %s", token.UID)
		}

		c.Set("uid", token.UID)
		c.Set("claims", claims) // Store raw claims
		c.Set("role", role) // Store extracted role
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
		userRole, exists := c.Get("role")
		if !exists {
			// This should not happen if AuthMiddleware runs first
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "User role not found in context."})
			return
		}

		// Compare the extracted role (string type assertion done in AuthMiddleware)
		if ok && userRole == requiredRole {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		}
	}
}