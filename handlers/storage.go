package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"your_module_path/services" // Replace with your actual module path
)

// UploadImageHandler handles image uploads to Google Cloud Storage.
// It expects a file in the "image" form field.
func UploadImageHandler(c *gin.Context) {
	// Check if the user is authenticated (middleware should handle this)
	// You can access user info from c.MustGet("user").(your_user_struct) if set by middleware

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file uploaded"})
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer src.Close()

	// Call the storage service to upload the image
	// You might want to pass a specific path or directory based on the image type (deal/review)
	// For now, a generic upload is shown.
	imageURL, err := services.UploadImage(c.Request.Context(), src, file.Filename) // Implement UploadImage in services/storage.go
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}