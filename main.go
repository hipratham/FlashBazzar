package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/firestore"
	"flashbazaar/handlers" // Import the handlers package
	"flashbazaar/middleware"
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
			"message": "Hello, world!", // This was the original hello world, can be updated for hosting
 })
 })

	// Add the POST /users route
	router.POST("/users", handlers.CreateUserHandler)

	// Protected routes using middleware
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware)
	{
		// Route requiring authentication
		protected.GET("/profile", func(c *gin.Context) {
			// Access user info from context set by AuthMiddleware
			uid, _ := c.Get("userUID")
			c.JSON(200, gin.H{"message": "Welcome to your profile!", "uid": uid})
		})
		// Hypothetical admin route requiring admin role
		protected.GET("/admin", middleware.RoleGuard("admin"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to the admin panel!"})
		})
	}

	// Vendor routes requiring authentication and vendor role
	vendorRoutes := router.Group("/vendor")
	vendorRoutes.Use(middleware.AuthMiddleware, middleware.RoleGuard("vendor"))
	{
		vendorRoutes.POST("/deal", handlers.CreateDealHandler)
		vendorRoutes.GET("/deal/:dealID", handlers.GetDealHandler)
		vendorRoutes.PUT("/deal/:dealID", handlers.UpdateDealHandler)
		vendorRoutes.DELETE("/deal/:dealID", handlers.DeleteDealHandler)
	}

	router.GET("/deal/today", handlers.GetTodayDealHandler)
	router.Run(":9090")
 }
