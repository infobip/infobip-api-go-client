package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/rcs"
)

func TestSendRcsMessage_WithSuggestions(t *testing.T) {
	client, auth := newClientAndAuth()

	destinations := []rcs.ToDestination{
		*rcs.NewToDestination("<DESTINATION>"),
	}

	reply := rcs.ReplySuggestionAsSuggestion(rcs.NewReplySuggestion("Yes", "yes"))
	openURL := rcs.OpenUrlSuggestionAsSuggestion(rcs.NewOpenUrlSuggestion("Learn more", "learn-more", "https://example.com/rcs"))

	content := rcs.NewOutboundTextContent("RCS with suggestions and tracking.")
	content.Suggestions = []rcs.Suggestion{reply, openURL}

	message := rcs.NewMessage(
		"<SENDER>",
		destinations,
		rcs.OutboundTextContentAsOutboundMessageContent(content),
	)

	request := rcs.NewMessageRequest([]rcs.Message{*message})
	request.SetOptions(rcs.MessageRequestOptions{
		Tracking: &rcs.UrlOptions{
			ShortenUrl:     infobip.PtrBool(true),
			TrackClicks:    infobip.PtrBool(true),
			TrackingUrl:    infobip.PtrString("https://example.com/rcs-tracking"),
			RemoveProtocol: infobip.PtrBool(false),
		},
	})

	apiResponse, httpResponse, err := client.
		RcsAPI.
		SendRcsMessages(auth).
		MessageRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send RCS message: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestSendRcsText_Basic(t *testing.T) {
	client, auth := newClientAndAuth()

	destinations := []rcs.ToDestination{
		*rcs.NewToDestination("<DESTINATION>"),
	}

	content := rcs.NewOutboundTextContent("Hello from RCS!")

	message := rcs.NewMessage(
		"<SENDER>",
		destinations,
		rcs.OutboundTextContentAsOutboundMessageContent(content),
	)

	request := rcs.NewMessageRequest([]rcs.Message{*message})

	apiResponse, httpResponse, err := client.
		RcsAPI.
		SendRcsMessages(auth).
		MessageRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send basic RCS message: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestGetRcsDeliveryReports(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		RcsAPI.
		GetOutboundRcsMessageDeliveryReports(auth).
		Execute(); err != nil {
		t.Fatalf("Failed to fetch RCS delivery reports: %v (http=%+v)", err, httpResp)
	}
}

func TestGetRcsMessageLogs(t *testing.T) {
	client, auth := newClientAndAuth()

	if _, httpResp, err := client.
		RcsAPI.
		GetOutboundRcsMessageLogs(auth).
		Execute(); err != nil {
		t.Fatalf("Failed to fetch RCS logs: %v (http=%+v)", err, httpResp)
	}
}

func TestRcsCapabilityCheckSyncAndAsync(t *testing.T) {
	client, auth := newClientAndAuth()

	syncReq := rcs.NewCapabilityCheckSyncRequest("<SENDER>", []string{"<DESTINATION>"})
	if _, httpResp, err := client.
		RcsAPI.
		CapabilityCheckRcsDestinationsQuery(auth).
		CapabilityCheckSyncRequest(*syncReq).
		Execute(); err != nil {
		t.Fatalf("Sync capability check failed: %v (http=%+v)", err, httpResp)
	}

	asyncReq := rcs.NewCapabilityCheckAsyncRequest("<SENDER>", []string{"<DESTINATION>"})
	if _, httpResp, err := client.
		RcsAPI.
		CapabilityCheckRcsDestinationsNotify(auth).
		CapabilityCheckAsyncRequest(*asyncReq).
		Execute(); err != nil {
		t.Fatalf("Async capability check failed: %v (http=%+v)", err, httpResp)
	}
}

func TestSendRcsTypingEvent(t *testing.T) {
	client, auth := newClientAndAuth()

	dest := []rcs.ToDestination{*rcs.NewToDestination("<DESTINATION>")}
	typingContent := rcs.NewOutboundEventTypingIndicatorContent()
	content := rcs.OutboundEventTypingIndicatorContentAsOutboundEventContent(
		typingContent,
	)
	event := rcs.NewEvent("<SENDER>", dest, content)
	eventReq := rcs.NewEventRequest([]rcs.Event{*event})

	if _, httpResp, err := client.
		RcsAPI.
		SendRcsEvents(auth).
		EventRequest(*eventReq).
		Execute(); err != nil {
		t.Fatalf("Failed to send RCS typing event: %v (http=%+v)", err, httpResp)
	}
}
