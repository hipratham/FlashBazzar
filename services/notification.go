package services

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010" // Use v2010 for sending messages
)

// SendGrid config
var sendgridClient *sendgrid.Client
var sendgridAPIKey string // Loaded from environment

// Twilio config
var twilioClient *twilio.RestClient
var twilioAccountSID string // Loaded from environment
var twilioAuthToken string   // Loaded from environment
var twilioPhoneNumber string // Loaded from environment

// InitNotificationClients initializes the SendGrid and Twilio clients.
// Call this function during application startup.
func InitNotificationClients(sgAPIKey, twilioSID, twilioAuth, twilioPhone string) {
	sendgridAPIKey = sgAPIKey
	twilioAccountSID = twilioSID
	twilioAuthToken = twilioAuth
	twilioPhoneNumber = twilioPhone

	sendgridClient = sendgrid.NewSendClient(sendgridAPIKey)
	twilioClient = twilio.NewRestClient() // Uses environment variables by default, but we'll set explicitly later if needed
}

// SendEmail sends an email using SendGrid.
func SendEmail(toEmail, subject, plainTextContent, htmlContent string) error {
	if sendgridClient == nil {
		log.Println("SendGrid client not initialized.")
		return fmt.Errorf("SendGrid client not initialized")
	}

	from := mail.NewEmail("FlashBazzar", "no-reply@flashbazzar.com") // Replace with your sender email
	to := mail.NewEmail("", toEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := sendgridClient.Send(message)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	log.Printf("Email sent: Status Code %d", response.StatusCode)
	// You can check response.Body and response.Headers for more details
	return nil
}

// SendSMS sends an SMS message using Twilio.
func SendSMS(toPhoneNumber, messageBody string) error {
	if twilioClient == nil {
		log.Println("Twilio client not initialized.")
		return fmt.Errorf("Twilio client not initialized")
	}

	// Ensure Twilio client is configured with SID and AuthToken explicitly if not relying on env vars
	// This is an example; in a real app, manage Twilio client configuration properly
	// twilioClient.SetAccountSid(twilioAccountSID)
	// twilioClient.SetAuthToken(twilioAuthToken)

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toPhoneNumber)
	params.SetFrom(twilioPhoneNumber)
	params.SetBody(messageBody)

	resp, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return err
	}
	log.Printf("SMS sent: SID %s, Status %s", resp.Sid, resp.Status)
	return nil
}

// NotificationRequest struct for potential API requests
type NotificationRequest struct {
	Type    string `json:"type" binding:"required,oneof=email sms push"` // push for Phase 2
	To      string `json:"to" binding:"required"`
	Subject string `json:"subject,omitempty"` // for email
	Body    string `json:"body" binding:"required"`
}

// Placeholder for Nepali SMS Gateway integration if needed
// func SendNepaliSMS(toPhoneNumber, messageBody string) error {
// 	// Implement logic for Nepali SMS Gateway API call
// 	log.Println("Sending SMS via Nepali Gateway (placeholder)")
// 	return nil // Or return error
// }