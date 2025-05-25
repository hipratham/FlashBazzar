package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// AuthMiddleware is a Gin middleware to verify Firebase ID tokens.
func AuthMiddleware() gin.HandlerFunc {
	// Initialize Firebase Admin SDK (consider moving this to a shared initialization)
	// Replace "path/to/your/serviceAccountKey.json" with your actual service account key path or use GOOGLE_APPLICATION_CREDENTIALS env var
	opt := option.WithCredentialsFile("serviceAccountKey.json") // Assuming the key file is named serviceAccountKey.json
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// In a real application, handle this error more gracefully, e.g., log and exit
		panic("Error initializing Firebase app: " + err.Error())
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		// Log the error and potentially panic
		panic("error getting firebase auth client: " + err.Error())
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			c.Abort()
			return
		}

		// Attach the user's UID to the request context
		c.Set("uid", token.UID)
		c.Next()
	}
}