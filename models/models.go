package models

import (
	"time"

	"cloud.google.com/go/firestore"
)

// User represents a user in the system.
type User struct {
	UserID string `firestore:"userId"`
	Email  string `firestore:"email"`
	Role   string `firestore:"role"` // e.g., "customer", "vendor", "admin"
}

// Vendor represents a vendor associated with a user.
type Vendor struct {
	VendorID  string `firestore:"vendorId"`
	UserID    string `firestore:"userId"`
	Name      string `firestore:"name"`
	Description string `firestore:"description"`
}

// Deal represents a flash deal.
type Deal struct {
	DealID      string    `firestore:"dealId"`
	VendorID    string    `firestore:"vendorId"`
	Name        string    `firestore:"name"`
	Description string    `firestore:"description"`
	Price       float64   `firestore:"price"`
	Stock       int       `firestore:"stock"`
	StartTime   time.Time `firestore:"startTime"`
	EndTime     time.Time `firestore:"endTime"`
}

// Order represents a customer order for a deal.
type Order struct {
	OrderID   string    `firestore:"orderId"`
	DealID    string    `firestore:"dealId"`
	UserID    string    `firestore:"userId"`
	Quantity  int       `firestore:"quantity"`
	Price     float64   `firestore:"price"` // Price at the time of order
	Status    string    `firestore:"status"` // e.g., "pending", "completed", "cancelled"
	CreatedAt time.Time `firestore:"createdAt"`
}

// Review represents a review for a deal.
type Review struct {
	ReviewID  string    `firestore:"reviewId"`
	DealID    string    `firestore:"dealId"`
	UserID    string    `firestore:"userId"`
	Rating    int       `firestore:"rating"` // e.g., 1-5
	Comment   string    `firestore:"comment"`
	CreatedAt time.Time `firestore:"createdAt"`
}

// Subscription represents a user subscription to notifications for future deals.
type Subscription struct {
	SubscriptionID string    `firestore:"subscriptionId"`
	UserID         string    `firestore:"userId"`
	Method         string    `firestore:"method"` // e.g., "email", "sms"
	ContactInfo    string    `firestore:"contactInfo"` // Email address or phone number
	CreatedAt      time.Time `firestore:"createdAt"`
}