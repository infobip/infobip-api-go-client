package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
)

func TestSendClickToCallBasic(t *testing.T) {
	client, auth := newClientAndAuth()

	message := voice.NewClickToCallMessage("<DESTINATION_A>", "<DESTINATION_B>", "<FROM_NUMBER>")
	request := voice.NewClickToCallMessageRequest([]voice.ClickToCallMessage{*message})

	apiResponse, httpResponse, err := client.
		ClickToCallAPI.
		SendClickToCallMessage(auth).
		ClickToCallMessageRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send click-to-call message: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, got: %+v", apiResponse)
	}
}
