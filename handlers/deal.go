package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"flashbazaar/models" // Assuming models are defined in models package
	"flashbazaar/services" // Assuming Firestore service is in services package
)

// DealHandler handles requests related to deals.
type DealHandler struct {
	FirestoreService *services.FirestoreService
}

// NewDealHandler creates a new instance of DealHandler.
func NewDealHandler() *DealHandler {
	return &DealHandler{
		// TODO: Initialize Firestore service
	}
}

// CreateDealHandler handles the creation of a new deal.
// Requires vendor authentication and approval.
func (h *DealHandler) CreateDealHandler(c *gin.Context) {
	var deal models.Deal
	if err := c.ShouldBindJSON(&deal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Basic validation (add more as needed)
	if deal.Name == "" || deal.Price <= 0 || deal.Stock <= 0 || deal.StartTime.IsZero() || deal.EndTime.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deal data"})
		return
	}

	// Ensure end time is after start time
	if !deal.EndTime.After(deal.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End time must be after start time"})
		return
	}

	// - Get authenticated vendor ID from context
	vendorID, exists := c.Get("userID") // Assuming AuthMiddleware sets "userID"
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Vendor not authenticated"})
		return
	}
	deal.VendorID = vendorID.(string)

	// TODO: Check if vendor is approved - Requires fetching vendor document from Firestore
	// vendor, err := h.FirestoreService.GetVendorByID(deal.VendorID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get vendor information"})
	// 	return
	// }
	// if !vendor.IsApproved {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "Vendor not approved"})
	// 	return
	// }

	// Save deal to Firestore
	ref, _, err := h.FirestoreService.Client.Collection("deals").Add(c, deal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create deal", "details": err.Error()})
		return
	}

	deal.DealID = ref.ID // Assign the Firestore generated ID
	c.JSON(http.StatusCreated, deal)
}

// GetDealHandler handles retrieving a specific deal by ID.
func (h *DealHandler) GetDealHandler(c *gin.Context) {
	dealID := c.Param("id")
	if dealID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deal ID is required"})
		return
	}

	doc, err := h.FirestoreService.Client.Collection("deals").Doc(dealID).Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch deal", "details": err.Error()})
		return
	}

	var deal models.Deal
	if err := doc.DataTo(&deal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse deal data", "details": err.Error()})
		return
	}
	deal.DealID = doc.Ref.ID // Assign the Firestore ID

	c.JSON(http.StatusOK, deal)
}

// UpdateDealHandler handles updating an existing deal.
// Requires vendor authentication and ownership of the deal.
func (h *DealHandler) UpdateDealHandler(c *gin.Context) {
	// TODO: Implement update deal logic
	// - Get deal ID from request parameters (c.Param("id"))
	// - Parse request body for updated deal data
	// - Validate input data
	// - Get authenticated vendor ID from context
	// - Fetch deal from Firestore and check ownership (requires Firestore service)
	// - Update deal in Firestore (requires Firestore service)
	// - Return success response
	c.JSON(http.StatusOK, gin.H{"message": "UpdateDealHandler placeholder"})
}

// DeleteDealHandler handles deleting a deal.
// Requires vendor authentication and ownership of the deal.
func (h *DealHandler) DeleteDealHandler(c *gin.Context) {
	// TODO: Implement delete deal logic
	// - Get deal ID from request parameters (c.Param("id"))
	// - Get authenticated vendor ID from context
	// - Fetch deal from Firestore and check ownership (requires Firestore service)
	// - Delete deal from Firestore (requires Firestore service)
	// - Return success response
	c.JSON(http.StatusOK, gin.H{"message": "DeleteDealHandler placeholder"})
}