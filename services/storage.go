package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

// StorageService handles interactions with Google Cloud Storage.
type StorageService struct {
	Client *storage.Client
	Bucket string
}

// NewStorageService creates a new instance of StorageService.
func NewStorageService(ctx context.Context, bucketName string) (*StorageService, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %w", err)
	}
	return &StorageService{
		Client: client,
		Bucket: bucketName,
	}, nil
}

// UploadImage uploads an image file to Google Cloud Storage and returns its public URL.
// It handles basic file type checking and size limits.
func (s *StorageService) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer file.Close()

	// Basic file type check (allow common image types)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	if _, ok := allowedTypes[fileHeader.Header.Get("Content-Type")]; !ok {
		return "", fmt.Errorf("unsupported file type: %s", fileHeader.Header.Get("Content-Type"))
	}

	// Basic file size limit (e.g., 10MB)
	const maxUploadSize = 10 << 20 // 10 MB
	if fileHeader.Size > maxUploadSize {
		return "", fmt.Errorf("file size exceeds limit (%d bytes)", maxUploadSize)
	}

	// Generate a unique filename for the object in the bucket
	objectName := fmt.Sprintf("uploads/%d-%s", time.Now().UnixNano(), fileHeader.Filename)

	// Create a new object handle
	obj := s.Client.Bucket(s.Bucket).Object(objectName)

	// Upload the file data
	wc := obj.NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close storage writer: %w", err)
	}

	// Make the object publicly readable (optional, based on your requirements)
	// If you don't make it public, you'll need to generate signed URLs for access.
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", fmt.Errorf("failed to set object ACL: %w", err)
	}

	// Return the public URL of the uploaded object
	// Adjust the URL format based on your bucket and object name
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", s.Bucket, objectName)

	return publicURL, nil
}