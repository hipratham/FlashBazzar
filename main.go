package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/firestore"
	"flashbazaar/handlers" // Import the handlers package
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

 func main() {
	ctx := context.Background()
	conf := &firebase.Config{}
	// For local development with the emulator
	opt := option.WithoutAuthentication()
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("firebase.NewApp error: %v", err)
	}

	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore error: %v", err)
	}

	defer firestoreClient.Close()

	// Initialize Gin router
 router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
 })
 })

	// Add the POST /users route
	router.POST("/users", handlers.CreateUserHandler)

	router.Run(":9090")
 }
