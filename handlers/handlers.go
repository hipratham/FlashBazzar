package handlers

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"flashbazaar/models" // Assuming models package is named flashbazaar/models
)

// This package will contain the API handlers for the application.
// Handlers for public routes, user authentication, and vendor panel
// functionalities will be added here.

var FirestoreClient *firestore.Client

// CreateUserHandler handles the creation of a new user in Firestore.
func CreateUserHandler(c *gin.Context) {
	var newUser models.User

	// Bind the JSON request body to the User struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the user to the "users" collection in Firestore
	_, _, err := FirestoreClient.Collection("users").Add(context.Background(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in Firestore", "details": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
}


package handlers

// This package will contain the API handlers for the application.
// Handlers for public routes, user authentication, and vendor panel
// functionalities will be added here.