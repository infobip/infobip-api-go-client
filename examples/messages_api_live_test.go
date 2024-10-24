package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/messagesapi"
)

func TestSendMessagesApiMessage(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

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
