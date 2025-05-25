package handlers

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// firebaseApp is the Firebase application instance.
var firebaseApp *firebase.App

// InitFirebaseAuth initializes the Firebase application.
func InitFirebaseAuth() error {
	opt := option.WithCredentialsFile("path/to/your/firebase/admin_sdk.json") // Replace with your service account key path
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	firebaseApp = app
	return nil
}

// AuthMiddleware is a Gin middleware to verify Firebase ID tokens.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		idToken := strings.ReplaceAll(authHeader, "Bearer ", "")
		if idToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			c.Abort()
			return
		}

		client, err := firebaseApp.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firebase Auth client"})
			c.Abort()
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Firebase ID token"})
			c.Abort()
			return
		}

		// Store the authenticated user's UID in the context
		c.Set("user_id", token.UID)
		c.Next()
	}
}

// UserData is a struct to hold user information from the Firebase ID token.
type UserData struct {
	UID   string `json:"uid"`
	Email string `json:"email,omitempty"`
}

// GetAuthenticatedUserHandler handles requests to get authenticated user information.
// It extracts the user ID from the context set by the AuthMiddleware
// and can optionally fetch more details from Firebase Auth.
func GetAuthenticatedUserHandler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		// This should not happen if AuthMiddleware is applied correctly,
		// but it's good practice to handle the case.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// In a real application, you might fetch more user details from Firestore
	// or Firebase Auth here if needed.
	c.JSON(http.StatusOK, gin.H{"user": UserData{UID: userID.(string)}})
}

// AuthenticatedRouteExample is a handler for a route that requires authentication.
func AuthenticatedRouteExample(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Authenticated route accessed successfully", "user_id": userID})
}

// RegisterAuthRoutes registers authentication related handlers.
func RegisterAuthRoutes(router *gin.Engine) {
	// This is where you would define public authentication routes like sign-up and sign-in
	// For example:
	// router.POST("/signup", SignUpHandler)
	// router.POST("/signin", SignInHandler)

	authenticated := router.Group("/")
	authenticated.Use(AuthMiddleware())
	{
		authenticated.GET("/auth/user", GetAuthenticatedUserHandler)
		// Add other authenticated routes here
		// authenticated.GET("/profile", GetProfileHandler)
		// authenticated.POST("/deals/:id/buy", BuyDealHandler)
	}
}