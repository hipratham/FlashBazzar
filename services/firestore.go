package services

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// Example:
// func InitFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
// 	// Firestore client initialization logic
// 	return nil, nil // Placeholder
// }

// User represents the structure of a user document in Firestore.
type User struct {
	UID          string    `firestore:"uid"`
	Name         string    `firestore:"name"`
	Email        string    `firestore:"email"`
	Phone        string    `firestore:"phone"`
	IsPremium    bool      `firestore:"isPremium"`
	City         string    `firestore:"city"`
	ReferralCode string    `firestore:"referralCode"`
	CreatedAt    time.Time `firestore:"createdAt"`
}

// Vendor represents the structure of a vendor document in Firestore.
type Vendor struct {
	VendorID    string    `firestore:"vendorId"`
	UID         string    `firestore:"uid"`
	BusinessName string    `firestore:"businessName"`
	ContactInfo string    `firestore:"contactInfo"`
	Approved    bool      `firestore:"approved"`
	CreatedAt   time.Time `firestore:"createdAt"`
}

// Deal represents the structure of a deal document in Firestore.
type Deal struct {
	DealID string `firestore:"dealId"`
	VendorID string `firestore:"vendorId"`
	Title string `firestore:"title"`
	Desc string `firestore:"desc"`
	ImageURL string `firestore:"imageURL"`
	OriginalPrice float64 `firestore:"originalPrice"`
	FlashPrice float64 `firestore:"flashPrice"`
	Stock int `firestore:"stock"`
	StartTime time.Time `firestore:"startTime"`
	EndTime time.Time `firestore:"endTime"`
	Boosted bool `firestore:"boosted"`
	CreatedAt time.Time `firestore:"createdAt"`
}

// Order represents the structure of an order document in Firestore.
type Order struct {
	OrderID string `firestore:"orderId"`
	UID string `firestore:"uid"`
	DealID string `firestore:"dealId"`
	Quantity int `firestore:"quantity"`
	Status string `firestore:"status"` // e.g., "pending", "completed"
	AmountPaid float64 `firestore:"amountPaid"`
	OrderedAt time.Time `firestore:"orderedAt"`
	DeliveredAt time.Time `firestore:"deliveredAt"`
}
// Example:
// func GetDealByID(ctx context.Context, client *firestore.Client, dealID string) (*Deal, error) {
// 	// Logic to fetch a deal by ID
// 	return nil, nil // Placeholder
// }

// Add other data access functions here for different collections (users, vendors, etc.)

// GetUser retrieves a user document from the 'users' collection by their UID.
func GetUser(ctx context.Context, client *firestore.Client, uid string) (*User, error) {
	var user User
	docsnap, err := client.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return nil, err
	}
	return &user, docsnap.DataTo(&user)
}

// CreateVendor adds a new vendor document to the 'vendors' collection.
func CreateVendor(ctx context.Context, client *firestore.Client, vendor *Vendor) error {
	_, err := client.Collection("vendors").Doc(vendor.VendorID).Set(ctx, vendor)
	return err
}

// GetVendor retrieves a vendor document from the 'vendors' collection by their VendorID.
func GetVendor(ctx context.Context, client *firestore.Client, vendorID string) (*Vendor, error) {
	var vendor Vendor
	docsnap, err := client.Collection("vendors").Doc(vendorID).Get(ctx)
	if err != nil {
		return nil, err
	}
	return &vendor, docsnap.DataTo(&vendor)
}

// UpdateVendor updates an existing vendor document in the 'vendors' collection.
func UpdateVendor(ctx context.Context, client *firestore.Client, vendorID string, updates map[string]interface{}) error {
	_, err := client.Collection("vendors").Doc(vendorID).Set(ctx, updates, firestore.MergeAll)
	return err
}

// DeleteVendor deletes a vendor document from the 'vendors' collection by their VendorID.
func DeleteVendor(ctx context.Context, client *firestore.Client, vendorID string) error {
	_, err := client.Collection("vendors").Doc(vendorID).Delete(ctx)
	return err
}

// GetVendorByUID retrieves a vendor document from the 'vendors' collection by their UID.
func GetVendorByUID(ctx context.Context, client *firestore.Client, uid string) (*Vendor, error) {
	// This requires a query, which is more complex than getting by document ID directly.
	return nil, nil // Placeholder - implementation needed
}

// CreateDeal adds a new deal document to the 'deals' collection.
func CreateDeal(ctx context.Context, client *firestore.Client, deal *Deal) error {
	_, err := client.Collection("deals").Doc(deal.DealID).Set(ctx, deal)
	return err
}

// GetDeal retrieves a deal document from the 'deals' collection by its DealID.
func GetDeal(ctx context.Context, client *firestore.Client, dealID string) (*Deal, error) {
	var deal Deal
	docsnap, err := client.Collection("deals").Doc(dealID).Get(ctx)
	if err != nil {
		return nil, err
	}
	return &deal, docsnap.DataTo(&deal)
}

// GetAllDeals retrieves all deal documents from the 'deals' collection.
func GetAllDeals(ctx context.Context, client *firestore.Client) ([]Deal, error) {
	var deals []Deal
	iter := client.Collection("deals").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			break // TODO: Handle potential errors other than iterator.Done
		}
		var deal Deal
		if err := doc.DataTo(&deal); err != nil {
			return nil, err
		}
		deals = append(deals, deal)
	}
	return deals, nil
}

// UpdateDeal updates an existing deal document in the 'deals' collection.
func UpdateDeal(ctx context.Context, client *firestore.Client, dealID string, updates map[string]interface{}) error {
	_, err := client.Collection("deals").Doc(dealID).Set(ctx, updates, firestore.MergeAll)
	return err
}

// DeleteDeal deletes a deal document from the 'deals' collection by its DealID.
func DeleteDeal(ctx context.Context, client *firestore.Client, dealID string) error {
	_, err := client.Collection("deals").Doc(dealID).Delete(ctx)
	return err
}

// CreateOrder adds a new order document to the 'orders' collection.
func CreateOrder(ctx context.Context, client *firestore.Client, order *Order) error {
	_, err := client.Collection("orders").Doc(order.OrderID).Set(ctx, order)
	return err
}

// GetOrder retrieves an order document from the 'orders' collection by its OrderID.
func GetOrder(ctx context.Context, client *firestore.Client, orderID string) (*Order, error) {
	var order Order
	docsnap, err := client.Collection("orders").Doc(orderID).Get(ctx)
	if err != nil {
		return nil, err
	}
	return &order, docsnap.DataTo(&order)
}

// GetUserOrders retrieves all order documents for a specific user from the 'orders' collection.
func GetUserOrders(ctx context.Context, client *firestore.Client, uid string) ([]Order, error) {
	var orders []Order
	iter := client.Collection("orders").Where("uid", "==", uid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			// TODO: Handle potential errors other than iterator.Done
			if err.Error() == "iterator.Done" {
				break
			}
			return nil, err
		}
		var order Order
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// GetVendorOrders retrieves all order documents for deals associated with a specific vendor.
// This requires joining or querying across collections which is complex in Firestore without denormalization.
// A more efficient approach for a dashboard might involve a denormalized field in the order or
// running a Cloud Function to aggregate vendor orders.
func GetVendorOrders(ctx context.Context, client *firestore.Client, vendorID string) ([]Order, error) {
	// Placeholder - efficient implementation needed
	// This would likely involve:
	// 1. Getting all deals for the vendor.
	// 2. Querying orders collection for each deal ID.
	// This can be inefficient for many deals.
	// Alternatively, store vendorId on the order document or use a collection group query if feasible.

	return nil, nil // Placeholder
}

// Review represents the structure of a review document in Firestore.
type Review struct {
	ReviewID string `firestore:"reviewId"`
	OrderID string `firestore:"orderId"`
	UID string `firestore:"uid"`
	Rating int `firestore:"rating"` // e.g., 1-5
	Comment string `firestore:"comment"`
	ImageURLs []string `firestore:"imageURLs"` // URLs to images stored in Firebase Storage
	SubmittedAt time.Time `firestore:"submittedAt"`
}

// CreateReview adds a new review document to the 'reviews' collection.
func CreateReview(ctx context.Context, client *firestore.Client, review *Review) error {
	_, err := client.Collection("reviews").Doc(review.ReviewID).Set(ctx, review)
	return err
}

// GetDealReviews retrieves all review documents for a specific deal from the 'reviews' collection.
func GetDealReviews(ctx context.Context, client *firestore.Client, dealID string) ([]Review, error) {
	var reviews []Review
	// To get reviews for a deal, we'd typically query by orderId, and then look up the dealId from the order.
	// A more direct approach would be to store dealId on the review document, which is a common denormalization strategy.
	// Assuming dealId is stored on the review document for efficiency.
	iter := client.Collection("reviews").Where("dealId", "==", dealID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			// TODO: Handle potential errors other than iterator.Done
			if err.Error() == "iterator.Done" {
				break
			}
			return nil, err
		}
		var review Review
		if err := doc.DataTo(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// Subscription represents the structure of a subscription document in Firestore.
type Subscription struct {
	UID string `firestore:"uid"`
	Type string `firestore:"type"` // e.g., "email", "SMS", "push"
	Endpoint string `firestore:"endpoint,omitempty"` // For push notifications
	CreatedAt time.Time `firestore:"createdAt"`
}

// CreateSubscription adds a new subscription document to the 'subscriptions' collection.
// The document ID can be a combination of UID and Type to ensure uniqueness per user and type.
func CreateSubscription(ctx context.Context, client *firestore.Client, subscription *Subscription) error {
	docID := subscription.UID + "_" + subscription.Type
	_, err := client.Collection("subscriptions").Doc(docID).Set(ctx, subscription)
	return err
}

// DeleteSubscription deletes a subscription document from the 'subscriptions' collection.
// The document ID is a combination of UID and Type.
func DeleteSubscription(ctx context.Context, client *firestore.Client, uid, subType string) error {
	docID := uid + "_" + subType
	_, err := client.Collection("subscriptions").Doc(docID).Delete(ctx)
	return err
}

// GetUserSubscriptions retrieves all subscription documents for a specific user from the 'subscriptions' collection.
func GetUserSubscriptions(ctx context.Context, client *firestore.Client, uid string) ([]Subscription, error) {
	var subscriptions []Subscription
	iter := client.Collection("subscriptions").Where("uid", "==", uid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			// TODO: Handle potential errors other than iterator.Done
			if err.Error() == "iterator.Done" {
				break
			}
			return nil, err
		}
		var subscription Subscription
		if err := doc.DataTo(&subscription); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}


// GetUserReviews retrieves all review documents submitted by a specific user from the 'reviews' collection.
func GetUserReviews(ctx context.Context, client *firestore.Client, uid string) ([]Review, error) {
	var reviews []Review
	iter := client.Collection("reviews").Where("uid", "==", uid).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			if err.Error() == "iterator.Done" {
				break
			}
			return nil, err
		}
		var review Review
		if err := doc.DataTo(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}