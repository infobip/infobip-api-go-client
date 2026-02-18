package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/messagesapi"
)

func TestSendMessagesApiMessage(t *testing.T) {
	infobipClient, auth := newClientAndAuth()

	channel := messagesapi.OUTBOUNDMESSAGECHANNEL_SMS
	sender := "<SENDER>"
	content := messagesapi.MessageContent{
		Body: messagesapi.MessageBody{
			MessageTextBody: messagesapi.NewMessageTextBody("Congratulations on sending your first message with GO library."),
		},
	}
	destinations := []messagesapi.MessageDestination{
		{
			ToDestination: messagesapi.NewToDestination("<DESTINATION>"),
		},
	}

	givenMessage := messagesapi.NewMessage(
		channel,
		sender,
		destinations,
		content,
	)

	request := messagesapi.NewRequest([]messagesapi.RequestMessagesInner{
		{Message: givenMessage},
	})

	// Send the Messages Api message
	apiResponse, httpResponse, err := infobipClient.
		MessagesAPI.
		SendMessagesApiMessage(auth).
		Request(*request).
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to send message: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

// Sends a scheduled Messages API SMS with URL tracking and campaign metadata.
func TestSendMessagesApiScheduledWithTracking(t *testing.T) {
	client, auth := newClientAndAuth()

	body := messagesapi.MessageTextBodyAsMessageBody(messagesapi.NewMessageTextBody("Sending you lots of otterly delightful vibes today!"))
	content := messagesapi.NewMessageContent(body)
	dest := messagesapi.ToDestinationAsMessageDestination(messagesapi.NewToDestination("<DESTINATION>"))

	givenMessage := messagesapi.NewMessage(
		messagesapi.OUTBOUNDMESSAGECHANNEL_SMS,
		"<SENDER>",
		[]messagesapi.MessageDestination{dest},
		*content,
	)

	givenMessage.SetOptions(messagesapi.MessageOptions{
		ValidityPeriod:      messagesapi.NewValidityPeriod(10),
		CampaignReferenceId: infobip.PtrString("campaign-spring"),
	})

	req := messagesapi.NewRequest([]messagesapi.RequestMessagesInner{
		{Message: givenMessage},
	})

	sendAtTime := time.Now().Add(2 * time.Minute)
	req.SetOptions(messagesapi.DefaultMessageRequestOptions{
		Schedule: &messagesapi.RequestSchedulingSettings{
			BulkId: infobip.PtrString("bulk-messagesapi-001"),
			SendAt: &infobip.Time{T: sendAtTime},
		},
		Tracking: &messagesapi.UrlOptions{
			ShortenUrl:     infobip.PtrBool(true),
			TrackClicks:    infobip.PtrBool(true),
			TrackingUrl:    infobip.PtrString("https://example.com/track-url"),
			RemoveProtocol: infobip.PtrBool(true),
		},
	})

	apiResponse, httpResponse, err := client.
		MessagesAPI.
		SendMessagesApiMessage(auth).
		Request(*req).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send scheduled Messages API message: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}
