package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/webhook"
)

// Define necessary structs for Stripe requests and responses

// CreateCheckoutSessionRequest represents the request body for creating a Stripe checkout session.
type CreateCheckoutSessionRequest struct {
	DealID   string `json:"deal_id" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
}

// CheckoutSessionResponse represents the response containing the Stripe checkout session URL.
type CheckoutSessionResponse struct {
	URL string `json:"url"`
}

// CreateCheckoutSession creates a new Stripe checkout session.
func CreateCheckoutSession(dealID string, quantity int64) (*stripe.CheckoutSession, error) {
	// TODO: Fetch deal information (price, name) from Firestore using dealID

	// TODO: Calculate the total amount based on deal price and quantity

	// Example: Create a line item for the checkout session
	// lineItems := []*stripe.CheckoutSessionLineItemParams{
	// 	{
	// 		PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
	// 			Currency: stripe.String("usd"), // Or your local currency
	// 			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
	// 				Name: stripe.String("Deal Title"), // Replace with actual deal title
	// 			},
	// 			UnitAmount: stripe.Int64(1000), // Replace with actual unit amount in cents
	// 		},
	// 		Quantity: stripe.Int64(quantity),
	// 	},
	// }

	params := &stripe.CheckoutSessionParams{
		// LineItems: lineItems, // Uncomment when line items are ready
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("YOUR_SUCCESS_URL?session_id={CHECKOUT_SESSION_ID}"), // TODO: Replace with your success URL
		CancelURL:  stripe.String("YOUR_CANCEL_URL"),                               // TODO: Replace with your cancel URL
		// TODO: Add metadata to the session to link it to your internal order/user
	}

	// TODO: Use your Stripe secret key from environment variables
	// stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	sess, err := session.New(params)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

// HandleStripeWebhook processes incoming Stripe webhook events.
func HandleStripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error reading request body"})
		return
	}

	// TODO: Use your Stripe webhook secret from environment variables
	// webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	// event, err := webhook.ConstructEvent(payload, c.Request.Header.Get("Stripe-Signature"), webhookSecret)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Error verifying webhook signature"})
	// 	return
	// }

	// // Handle the event
	// switch event.Type {
	// case "checkout.session.completed":
	// 	var completedSession stripe.CheckoutSession
	// 	err := json.Unmarshal(event.Data.Raw, &completedSession)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing webhook JSON"})
	// 		return
	// 	}
	// 	// TODO: Handle successful payment (e.g., update order status, decrement stock)
	// 	// You can access metadata from completedSession to find your internal order ID
	// 	// fmt.Printf("Checkout session completed: %s\n", completedSession.ID)
	// case "payment_intent.succeeded":
	// 	var paymentIntent stripe.PaymentIntent
	// 	err := json.Unmarshal(event.Data.Raw, &paymentIntent)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing webhook JSON"})
	// 		return
	// 	}
	// 	// TODO: Handle payment intent succeeded (if you are not using checkout sessions)
	// 	// fmt.Printf("PaymentIntent was successful: %s\n", paymentIntent.ID)
	// default:
	// 	// fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	// }

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Helper function to retrieve deal information (placeholder)
// TODO: Implement this function to fetch deal data from Firestore
func getDealPriceAndTitle(dealID string) (unitAmount int64, title string, err error) {
	// Placeholder implementation
	// In a real application, you would fetch this from your database
	return 1000, "Placeholder Deal", nil
}