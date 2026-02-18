package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/email"
)

func TestSendEmail(t *testing.T) {
	infobipClient, auth := newClientAndAuth()

	apiResponse, httpResponse, err := infobipClient.
		EmailAPI.
		SendEmail(auth).
		To([]string{"<DESTINATION1>", "<DESTINATION2>"}).
		From("<FROM>").
		Subject("<SUBJECT>").
		Text("<TEXT>").
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to send Email: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestSendEmailWithTrackingAndSchedule(t *testing.T) {
	client, auth := newClientAndAuth()

	sendAtTime := time.Now().Add(2 * time.Hour)
	ibSendAt := infobip.Time{T: sendAtTime}

	apiResponse, httpResponse, err := client.
		EmailAPI.
		SendEmail(auth).
		To([]string{"<DESTINATION1>", "<DESTINATION2>"}).
		Cc([]string{"<CC_RECIPIENT>"}).
		Bcc([]string{"<BCC_RECIPIENT>"}).
		From("<FROM>").
		ReplyTo("<REPLY_TO>").
		Subject("Seasonal offers").
		Html("<h1>Hello!</h1><p>Enjoy our seasonal offer with tracked opens and clicks.</p>").
		Track(true).
		TrackClicks(true).
		TrackOpens(true).
		TrackingUrl("https://example.com/track-email").
		BulkId("bulk-2026-02").
		MessageId("msg-2026-02-11").
		SendAt(ibSendAt).
		IntermediateReport(true).
		NotifyUrl("https://example.com/email/dlr").
		CallbackData("promo-callback").
		Execute()

	if err != nil {
		t.Fatalf("Failed to send advanced Email: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestValidateEmailAddress(t *testing.T) {
	client, auth := newClientAndAuth()

	req := email.NewValidationRequest("validations@example.com")

	if _, httpResp, err := client.
		EmailAPI.
		ValidateEmailAddresses(auth).
		ValidationRequest(*req).
		Execute(); err != nil {
		t.Fatalf("Failed to validate email address: %v (http=%+v)", err, httpResp)
	}
}

func TestRequestEmailValidations(t *testing.T) {
	client, auth := newClientAndAuth()

	destinations := []email.ValidationDestination{
		{Destination: "bulk1@example.com"},
		{Destination: "bulk2@example.com"},
	}
	bulkReq := email.NewApiBulkRequest(destinations)
	bulkReq.SetValidationRequestId("bulk-validation-2026-02-13")

	if _, httpResp, err := client.
		EmailAPI.
		RequestValidations(auth).
		ApiBulkRequest(*bulkReq).
		Execute(); err != nil {
		t.Fatalf("Failed to start bulk validations: %v (http=%+v)", err, httpResp)
	}
}

func TestGetEmailValidations(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		EmailAPI.
		GetValidations(auth).
		Size(10).
		Execute(); err != nil {
		t.Fatalf("Failed to get validations: %v (http=%+v)", err, httpResp)
	}
}

func TestGetEmailDeliveryReports(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		EmailAPI.
		GetEmailDeliveryReports(auth).
		Execute(); err != nil {
		t.Fatalf("Failed to get email delivery reports: %v (http=%+v)", err, httpResp)
	}
}

func TestGetEmailLogs(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		EmailAPI.
		GetEmailLogs(auth).
		Execute(); err != nil {
		t.Fatalf("Failed to get email logs: %v (http=%+v)", err, httpResp)
	}
}

func TestGetEmailTemplates(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		EmailAPI.
		GetEmailTemplates(auth).
		Execute(); err != nil {
		t.Fatalf("Failed to get email templates: %v (http=%+v)", err, httpResp)
	}
}

func TestGetScheduledEmails(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		EmailAPI.
		GetScheduledEmails(auth).
		Execute(); err != nil {
		t.Fatalf("Failed to get scheduled emails: %v (http=%+v)", err, httpResp)
	}
}
