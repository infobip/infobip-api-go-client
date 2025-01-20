package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
)

func TestSendSMS(t *testing.T) {
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

	destinations := []sms.Destination{
		{To: "<DESTINATION>"},
	}

	content := sms.MessageContent{
		TextMessageContent: sms.NewTextMessageContent("Congratulations on sending your first message with GO library."),
	}

	givenMessage := sms.NewMessage(destinations, content)
	givenMessage.SetSender("<SENDER>")

	request := sms.NewRequestEnvelope([]sms.Message{
		*givenMessage,
	})

	// Send the SMS message
	apiResponse, httpResponse, err := infobipClient.
		SmsAPI.
		SendSmsMessages(auth).
		RequestEnvelope(*request).
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to send SMS: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}
