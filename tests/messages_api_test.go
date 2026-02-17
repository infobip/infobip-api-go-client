package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/messagesapi"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	SendMessagesApiMessages            string = "/messages-api/1/messages"
	SendMessagesApiEvents              string = "/messages-api/1/events"
	MessagesApiDeliveryReportsEndpoint string = "/messages-api/1/reports"
	MessagesApiInboundEndpoint         string = "/messages-api/1/inbound"
	MessagesApiValidateEndpoint        string = "/messages-api/1/messages/validate"
)

func TestSendMessagesApiMessage(t *testing.T) {
	givenBulkId := "1688025180464000013"
	givenMessageId := "1688025180464000014"
	givenDestination := "48600700800"

	givenResponse := fmt.Sprintf(`
		{
			"bulkId": "%s",
			"messages": [
			  {
				"messageId": "%s",
				"status": {
				  "groupId": %d,
				  "groupName": "%s",
				  "id": %d,
				  "name": "%s",
				  "description": "%s"
				},
				"destination": "%s"
			  }
			]
		}`,
		givenBulkId,
		givenMessageId,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		givenDestination,
	)

	givenChannel := "WHATSAPP"
	givenSender := "447491163862"
	givenTo := "123456789"
	givenText := "Sending you lots of otterly delightful vibes today!"
	givenType := "TEXT"

	givenRequest := fmt.Sprintf(`
		{
			"messages": [
			  {
				"channel": "%s",
				"sender": "%s",
				"destinations": [
				  {
					"to": "%s"
				  }
				],
				"content": {
				  "body": {
					"text": "%s",
					"type": "%s"
				  }
				},
				"options": {
				  "platform": {
					"entityId": "entityId",
					"applicationId": "applicationId"
				  },
				  "validityPeriod": {
					"amount": 10,
					"timeUnit": "MINUTES"
				  },
				  "adaptationMode": false,
				  "campaignReferenceId": "campaignReferenceId"
				},
				"webhooks": {
				  "delivery": {
					"url": "https://example.com/delivery-webhook"
				  },
				  "callbackData": "callbackData",
				  "seen": {
					"url": "https://example.com/seen-webhook"
				  }
				}
			  }
			]
		}`,
		givenChannel,
		givenSender,
		givenTo,
		givenText,
		givenType,
	)

	body := messagesapi.MessageBody{
		MessageTextBody: messagesapi.NewMessageTextBody(givenText),
	}

	destination := messagesapi.MessageDestination{
		ToDestination: messagesapi.NewToDestination(givenTo),
	}

	givenMessage := messagesapi.NewMessage(
		messagesapi.OUTBOUNDMESSAGECHANNEL_WHATSAPP,
		givenSender,
		[]messagesapi.MessageDestination{
			destination,
		},
		messagesapi.MessageContent{
			Body: body,
		},
	)

	platform := messagesapi.NewPlatform()
	platform.SetEntityId("entityId")
	platform.SetApplicationId("applicationId")

	validityPeriod := messagesapi.NewValidityPeriod(10)
	validityPeriod.SetTimeUnit(messagesapi.VALIDITYPERIODTIMEUNIT_MINUTES)

	options := messagesapi.NewMessageOptions()
	options.SetPlatform(*platform)
	options.SetValidityPeriod(*validityPeriod)
	options.SetAdaptationMode(false)
	options.SetCampaignReferenceId("campaignReferenceId")

	delivery := messagesapi.NewMessageDeliveryReporting()
	delivery.SetUrl("https://example.com/delivery-webhook")

	seen := messagesapi.NewSeenStatusReporting()
	seen.SetUrl("https://example.com/seen-webhook")

	webhooks := messagesapi.NewOttWebhooks()
	webhooks.SetDelivery(*delivery)
	webhooks.SetCallbackData("callbackData")
	webhooks.SetSeen(*seen)

	givenMessage.Options = options
	givenMessage.Webhooks = webhooks

	request := messagesapi.NewRequest([]messagesapi.BaseMessage{
		{Message: givenMessage},
	})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SendMessagesApiMessages, givenResponse, 200)

	responseBody, r, _ := infobipClient.
		MessagesAPI.
		SendMessagesApiMessage(context.Background()).
		Request(*request).Execute()

	// Ensure r and responseBody are not nil before dereferencing
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.True(t, len(responseBody.GetMessages()) == 1)

	msg := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	messageStatus := msg.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), messageStatus.GetGroupId())
	assert.Equal(t, messagesapi.MESSAGEGENERALSTATUS_PENDING, messageStatus.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), messageStatus.GetId())
	assert.Equal(t, PendingStatusName, messageStatus.GetName())
	assert.Equal(t, PendingStatusDescription, messageStatus.GetDescription())
}

func TestSendMessagesApiEvent(t *testing.T) {

	givenBulkId := "1688025180464000013"
	givenMessageId := "1688025180464000014"
	givenDestination := "48600700800"

	givenResponse := fmt.Sprintf(`
		{
		  "bulkId": "%s",
		  "messages": [
		    {
		      "messageId": "%s",
		      "status": {
		        "groupId": %d,
		        "groupName": "%s",
		        "id": %d,
		        "name": "%s",
		        "description": "%s"
		      },
		      "destination": "%s"
		    }
		  ]
		}`,
		givenBulkId,
		givenMessageId,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		givenDestination,
	)

	givenChannel := "APPLE_MB"
	givenSender := "senderNumber"
	givenTo := "41793026727"
	givenEvent := "TYPING_STARTED"

	givenRequest := fmt.Sprintf(`
		{
			"events": [
			  {
			    "channel": "%s",
			    "sender": "%s",
			    "destinations": [
			      {
			        "to": "%s"
			      }
			    ],
			    "event": "%s"
			  }
			]
		}`,
		givenChannel,
		givenSender,
		givenTo,
		givenEvent,
	)

	destination := messagesapi.ToDestination{
		To: givenTo,
	}

	event := messagesapi.NewOutboundTypingStartedEvent(
		messagesapi.OUTBOUNDEVENTCHANNEL_APPLE_MB,
		givenSender,
		[]messagesapi.ToDestination{
			destination,
		},
	)

	request := messagesapi.EventRequest{
		Events: []messagesapi.OutboundEvent{
			{OutboundTypingStartedEvent: event},
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SendMessagesApiEvents, givenResponse, 200)

	// Execute the request
	responseBody, r, _ := infobipClient.
		MessagesAPI.
		SendMessagesApiEvents(context.Background()).
		EventRequest(request).Execute()

	// Ensure r and responseBody are not nil before dereferencing
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.True(t, len(responseBody.GetMessages()) == 1)

	msg := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	messageStatus := msg.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), messageStatus.GetGroupId())
	assert.Equal(t, messagesapi.MESSAGEGENERALSTATUS_PENDING, messageStatus.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), messageStatus.GetId())
	assert.Equal(t, PendingStatusName, messageStatus.GetName())
	assert.Equal(t, PendingStatusDescription, messageStatus.GetDescription())
}

func TestReceiveDeliveryReportsWebhook(t *testing.T) {

	expectedEvent := "DELIVERY"
	expectedChannel := messagesapi.InboundDlrChannel("WHATSAPP")
	expectedSender := "senderNumber"
	expectedDestination := "41793026727"
	expectedSentAt := "2024-02-06T14:18:29.797+0000"
	expectedDoneAt := "2024-02-06T17:18:29.797+0000"
	expectedBulkId := "1688025180464000013"
	expectedMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	expectedMessageCount := 1
	expectedEntityId := "my-entity-id"
	expectedApplicationId := "my-application-id"
	expectedCampaignReferenceId := "campaignRef"

	expectedRequest := fmt.Sprintf(`
		{
			"results": [
			  {
				"event": "%s",
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"sentAt": "%s",
				"doneAt": "%s",
				"bulkId": "%s",
				"messageId": "%s",
				"messageCount": %d,
				"status": {
				  "groupId": %d,
				  "groupName": "%s",
				  "id": %d,
				  "name": "%s",
				  "description": "%s"
				},
				"error": {
				  "groupId": %d,
				  "groupName": "%s",
				  "id": %d,
				  "name": "%s",
				  "description": "%s",
				  "permanent": %t
				},
				"platform": {
				  "entityId": "%s",
				  "applicationId": "%s"
				},
				"campaignReferenceId": "%s"
			  }
			]
		}`,
		expectedEvent,
		expectedChannel,
		expectedSender,
		expectedDestination,
		expectedSentAt,
		expectedDoneAt,
		expectedBulkId,
		expectedMessageId,
		expectedMessageCount,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		NoErrorStatusGroupId,
		NoErrorStatusGroupName,
		NoErrorStatusId,
		NoErrorStatusName,
		NoErrorStatusDescription,
		NoErrorPermanent,
		expectedEntityId,
		expectedApplicationId,
		expectedCampaignReferenceId,
	)

	var responseBody messagesapi.DeliveryReport

	err := json.Unmarshal([]byte(expectedRequest), &responseBody)

	// Ensure there is no error and responseBody is not nil before dereferencing
	assert.Nil(t, err, "Expected nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.True(t, len(responseBody.GetResults()) == 1)

	result := responseBody.GetResults()[0]
	assert.Equal(t, expectedEvent, result.GetEvent())
	assert.Equal(t, expectedChannel, result.GetChannel())
	assert.Equal(t, expectedSender, result.GetSender())
	assert.Equal(t, expectedDestination, result.GetDestination())
	assert.Equal(t, expectedSentAt, result.GetSentAt())
	assert.Equal(t, expectedDoneAt, result.GetDoneAt())
	assert.Equal(t, expectedBulkId, result.GetBulkId())
	assert.Equal(t, expectedMessageId, result.GetMessageId())
	assert.Equal(t, int32(expectedMessageCount), result.GetMessageCount())

	assert.NotNil(t, result.GetStatus(), "Expected non-nil status")

	status := result.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), status.GetGroupId())
	assert.Equal(t, string(PendingStatusGroupName), status.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), status.GetId())
	assert.Equal(t, PendingStatusName, status.GetName())
	assert.Equal(t, PendingStatusDescription, status.GetDescription())

	assert.NotNil(t, result.GetError(), "Expected non-nil error")

	error2 := result.GetError()
	assert.Equal(t, int32(NoErrorStatusGroupId), error2.GetGroupId())
	assert.Equal(t, string(NoErrorStatusGroupName), error2.GetGroupName())
	assert.Equal(t, int32(NoErrorStatusId), error2.GetId())
	assert.Equal(t, NoErrorStatusName, error2.GetName())
	assert.Equal(t, NoErrorStatusDescription, error2.GetDescription())
	assert.Equal(t, NoErrorPermanent, error2.GetPermanent())

	assert.NotNil(t, result.GetPlatform(), "Expected non-nil platform")

	platform := result.GetPlatform()
	assert.Equal(t, expectedEntityId, platform.GetEntityId())
	assert.Equal(t, expectedApplicationId, platform.GetApplicationId())

	assert.Equal(t, expectedCampaignReferenceId, result.GetCampaignReferenceId())
}

func TestReceiveSeenReportsWebhook(t *testing.T) {
	expectedEvent := "SEEN"
	expectedChannel := messagesapi.InboundSeenChannel("WHATSAPP")
	expectedSender := "senderNumber"
	expectedDestination := "41793026727"
	expectedSentAt := "2024-02-06T14:18:29.797+0000"
	expectedSeenAt := "2024-02-06T17:18:29.797+0000"
	expectedBulkId := "1688025180464000013"
	expectedMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	expectedCallbackData := "reference-callback-data-from-outbound-message"
	expectedEntityId := "my-entity-id"
	expectedApplicationId := "my-application-id"
	expectedCampaignReferenceId := "campaignRef"

	expectedRequest := fmt.Sprintf(`
		{
			"results": [
			  {
				"event": "%s",
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"sentAt": "%s",
				"seenAt": "%s",
				"bulkId": "%s",
				"messageId": "%s",
				"callbackData": "%s",
				"platform": {
				  "entityId": "%s",
				  "applicationId": "%s"
				},
				"campaignReferenceId": "%s"
			  }
			]
		}`,
		expectedEvent,
		expectedChannel,
		expectedSender,
		expectedDestination,
		expectedSentAt,
		expectedSeenAt,
		expectedBulkId,
		expectedMessageId,
		expectedCallbackData,
		expectedEntityId,
		expectedApplicationId,
		expectedCampaignReferenceId,
	)

	var responseBody messagesapi.SeenReport

	err := json.Unmarshal([]byte(expectedRequest), &responseBody)

	// Ensure there is no error and responseBody is not nil before dereferencing
	assert.Nil(t, err, "Expected nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.True(t, len(responseBody.GetResults()) == 1)

	result := responseBody.GetResults()[0]
	assert.Equal(t, expectedEvent, result.GetEvent())
	assert.Equal(t, expectedChannel, result.GetChannel())
	assert.Equal(t, expectedSender, result.GetSender())
	assert.Equal(t, expectedDestination, result.GetDestination())
	assert.Equal(t, expectedSentAt, result.GetSentAt())
	assert.Equal(t, expectedSeenAt, result.GetSeenAt())
	assert.Equal(t, expectedBulkId, result.GetBulkId())
	assert.Equal(t, expectedMessageId, result.GetMessageId())

	assert.NotNil(t, result.GetPlatform(), "Expected non-nil platform")

	platform := result.GetPlatform()
	assert.Equal(t, expectedEntityId, platform.GetEntityId())
	assert.Equal(t, expectedApplicationId, platform.GetApplicationId())

	assert.Equal(t, expectedCampaignReferenceId, result.GetCampaignReferenceId())
}

func TestReceiveIncomingMessagesWebhook(t *testing.T) {
	expectedChannel := messagesapi.InboundMoEventChannel("SMS")
	expectedSender := "48123234567"
	expectedDestination := "41793026727"
	expectedText := "Text message 123"
	expectedCleanText := "Text message"
	expectedType := "TEXT"
	expectedReceivedAt := "2024-02-06T14:18:29.797+0000"
	expectedMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	expectedEntityId := "my-entity-id"
	expectedApplicationId := "my-application-id"
	expectedEvent := messagesapi.InboundEventType("MO")

	expectedRequest := fmt.Sprintf(`
		{
			"results": [
			  {
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
				  {
					"text": "%s",
					"cleanText": "%s",
					"type": "%s"
				  }
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"platform": {
				  "entityId": "%s",
				  "applicationId": "%s"
				},
				"event": "%s"
			  }
			]
		}`,
		expectedChannel,
		expectedSender,
		expectedDestination,
		expectedText,
		expectedCleanText,
		expectedType,
		expectedReceivedAt,
		expectedMessageId,
		expectedEntityId,
		expectedApplicationId,
		expectedEvent,
	)

	var responseBody messagesapi.IncomingMessage

	err := json.Unmarshal([]byte(expectedRequest), &responseBody)

	assert.Nil(t, err, "Expected nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.True(t, len(responseBody.GetResults()) == 1)

	result := responseBody.GetResults()[0]
	assert.Equal(t, expectedChannel, result.MoEvent.GetChannel())
	assert.Equal(t, expectedSender, result.MoEvent.GetSender())
	assert.Equal(t, expectedDestination, result.MoEvent.GetDestination())

	assert.Equal(t, expectedReceivedAt, result.MoEvent.GetReceivedAt().String())
	assert.Equal(t, expectedMessageId, result.MoEvent.GetMessageId())

	platform := result.MoEvent.GetPlatform()
	assert.Equal(t, expectedEntityId, platform.GetEntityId())
	assert.Equal(t, expectedApplicationId, platform.GetApplicationId())
}

func TestReceiveIncomingMessagesWebhookWithCallBackDataAsString(t *testing.T) {
	expectedChannel := messagesapi.InboundMoEventChannel("SMS")
	expectedSender := "48123234567"
	expectedDestination := "41793026727"
	expectedText := "Text message 123"
	expectedCleanText := "Text message"
	expectedType := "TEXT"
	expectedReceivedAt := "2024-02-06T14:18:29.797+0000"
	expectedMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	expectedEntityId := "my-entity-id"
	expectedApplicationId := "my-application-id"
	expectedEvent := messagesapi.InboundEventType("MO")
	expectedCallBackData := "callbackData"

	expectedRequest := fmt.Sprintf(`
		{
			"results": [
			  {
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
				  {
					"text": "%s",
					"cleanText": "%s",
					"type": "%s"
				  }
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"callbackData": "%s",
				"platform": {
				  "entityId": "%s",
				  "applicationId": "%s"
				},
				"event": "%s"
			  }
			]
		}`,
		expectedChannel,
		expectedSender,
		expectedDestination,
		expectedText,
		expectedCleanText,
		expectedType,
		expectedReceivedAt,
		expectedMessageId,
		expectedCallBackData,
		expectedEntityId,
		expectedApplicationId,
		expectedEvent,
	)

	var responseBody messagesapi.IncomingMessage

	err := json.Unmarshal([]byte(expectedRequest), &responseBody)

	// Ensure there is no error and responseBody is not nil before dereferencing
	assert.Nil(t, err, "Expected nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.True(t, len(responseBody.GetResults()) == 1)

	result := responseBody.GetResults()[0]
	assert.Equal(t, expectedChannel, result.MoEvent.GetChannel())
	assert.Equal(t, expectedSender, result.MoEvent.GetSender())
	assert.Equal(t, expectedDestination, result.MoEvent.GetDestination())

	assert.Equal(t, expectedReceivedAt, result.MoEvent.GetReceivedAt().String())
	assert.Equal(t, expectedMessageId, result.MoEvent.GetMessageId())
	assert.Equal(t, expectedCallBackData, result.MoEvent.GetCallbackData())

	platform := result.MoEvent.GetPlatform()
	assert.Equal(t, expectedEntityId, platform.GetEntityId())
	assert.Equal(t, expectedApplicationId, platform.GetApplicationId())
}

func TestReceiveIncomingMessagesWebhookWithCallBackDataAsObject(t *testing.T) {
	expectedChannel := messagesapi.InboundMoEventChannel("SMS")
	expectedSender := "48123234567"
	expectedDestination := "41793026727"
	expectedText := "Text message 123"
	expectedCleanText := "Text message"
	expectedType := "TEXT"
	expectedReceivedAt := "2024-02-06T14:18:29.797+0000"
	expectedMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	expectedEntityId := "my-entity-id"
	expectedApplicationId := "my-application-id"
	expectedEvent := messagesapi.InboundEventType("MO")
	expectedCallBackData := "{\"key1\":\"value1\",\"key2\":\"value2\"}"

	expectedRequest := fmt.Sprintf(`
		{
			"results": [
			  {
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
				  {
					"text": "%s",
					"cleanText": "%s",
					"type": "%s"
				  }
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"callbackData": "{\"key1\":\"value1\",\"key2\":\"value2\"}",
				"platform": {
				  "entityId": "%s",
				  "applicationId": "%s"
				},
				"event": "%s"
			  }
			]
		}`,
		expectedChannel,
		expectedSender,
		expectedDestination,
		expectedText,
		expectedCleanText,
		expectedType,
		expectedReceivedAt,
		expectedMessageId,
		expectedEntityId,
		expectedApplicationId,
		expectedEvent,
	)

	var responseBody messagesapi.IncomingMessage

	err := json.Unmarshal([]byte(expectedRequest), &responseBody)

	assert.Nil(t, err, "Expected nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.True(t, len(responseBody.GetResults()) == 1)

	result := responseBody.GetResults()[0]
	assert.Equal(t, expectedChannel, result.MoEvent.GetChannel())
	assert.Equal(t, expectedSender, result.MoEvent.GetSender())
	assert.Equal(t, expectedDestination, result.MoEvent.GetDestination())

	assert.Equal(t, expectedReceivedAt, result.MoEvent.GetReceivedAt().String())
	assert.Equal(t, expectedMessageId, result.MoEvent.GetMessageId())
	assert.Equal(t, expectedCallBackData, result.MoEvent.GetCallbackData())

	platform := result.MoEvent.GetPlatform()
	assert.Equal(t, expectedEntityId, platform.GetEntityId())
	assert.Equal(t, expectedApplicationId, platform.GetApplicationId())
}

func TestReceiveIncomingMessagesWebhookWithCallBackDataAsEmptyObject(t *testing.T) {
	expectedChannel := messagesapi.InboundMoEventChannel("SMS")
	expectedSender := "48123234567"
	expectedDestination := "41793026727"
	expectedText := "Text message 123"
	expectedCleanText := "Text message"
	expectedType := "TEXT"
	expectedReceivedAt := "2024-02-06T14:18:29.797+0000"
	expectedMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	expectedEntityId := "my-entity-id"
	expectedApplicationId := "my-application-id"
	expectedEvent := messagesapi.InboundEventType("MO")
	expectedCallBackData := "{}"

	expectedRequest := fmt.Sprintf(`
		{
			"results": [
			  {
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
				  {
					"text": "%s",
					"cleanText": "%s",
					"type": "%s"
				  }
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"callbackData": "{}",
				"platform": {
				  "entityId": "%s",
				  "applicationId": "%s"
				},
				"event": "%s"
			  }
			]
		}`,
		expectedChannel,
		expectedSender,
		expectedDestination,
		expectedText,
		expectedCleanText,
		expectedType,
		expectedReceivedAt,
		expectedMessageId,
		expectedEntityId,
		expectedApplicationId,
		expectedEvent,
	)

	var responseBody messagesapi.IncomingMessage

	err := json.Unmarshal([]byte(expectedRequest), &responseBody)

	assert.Nil(t, err, "Expected nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.True(t, len(responseBody.GetResults()) == 1)

	result := responseBody.GetResults()[0]
	assert.Equal(t, expectedChannel, result.MoEvent.GetChannel())
	assert.Equal(t, expectedSender, result.MoEvent.GetSender())
	assert.Equal(t, expectedDestination, result.MoEvent.GetDestination())

	assert.Equal(t, expectedReceivedAt, result.MoEvent.GetReceivedAt().String())
	assert.Equal(t, expectedMessageId, result.MoEvent.GetMessageId())
	assert.Equal(t, expectedCallBackData, result.MoEvent.GetCallbackData())

	platform := result.MoEvent.GetPlatform()
	assert.Equal(t, expectedEntityId, platform.GetEntityId())
	assert.Equal(t, expectedApplicationId, platform.GetApplicationId())
}

func TestSendMessagesApiMessage_WhatsAppTemplate(t *testing.T) {
	givenBulkId := "1688025180464000013"
	givenMessageId := "1688025180464000014"
	givenDestination := "11111111"

	givenResponse := fmt.Sprintf(`
		{
			"bulkId": "%s",
			"messages": [
			  {
				"messageId": "%s",
				"status": {
				  "groupId": %d,
				  "groupName": "%s",
				  "id": %d,
				  "name": "%s",
				  "description": "%s"
				},
				"destination": "%s"
			  }
			]
		}`,
		givenBulkId,
		givenMessageId,
		PendingStatusGroupId,
		PendingStatusGroupName,
		PendingStatusId,
		PendingStatusName,
		PendingStatusDescription,
		givenDestination,
	)

	givenChannel := "WHATSAPP"
	givenSender := "447860099299"
	givenTo := "11111111"
	givenTemplateName := "templateName"
	givenLanguage := "en"

	givenRequest := fmt.Sprintf(`
		{
			"messages": [
			  {
				"channel": "%s",
				"sender": "%s",
				"destinations": [
				  {
					"to": "%s"
				  }
				],
				"template": {
				  "templateName": "%s",
				  "language": "%s"
				},
				"content": {
				  "body": {
					"1": "Infobip",
					"2": "5:30 AM",
					"type": "TEXT"
				  },
				  "buttons": [
					{
					  "suffix": "example1",
					  "type": "OPEN_URL"
					},
					{
					  "suffix": "example2",
					  "type": "OPEN_URL"
					}
				  ]
				}
			  }
			]
		}`,
		givenChannel,
		givenSender,
		givenTo,
		givenTemplateName,
		givenLanguage,
	)

	templateTextBody := messagesapi.NewTemplateTextBody()
	templateTextBody.AdditionalProperties["1"] = "Infobip"
	templateTextBody.AdditionalProperties["2"] = "5:30 AM"

	body := messagesapi.TemplateBody{
		TemplateTextBody: templateTextBody,
	}

	destination := messagesapi.MessageDestination{
		ToDestination: messagesapi.NewToDestination(givenTo),
	}

	givenButton1 := messagesapi.NewTemplateOpenUrlButton("example1")
	givenTemplateButton1 := messagesapi.TemplateButton{TemplateOpenUrlButton: givenButton1}
	givenButton2 := messagesapi.NewTemplateOpenUrlButton("example2")
	givenTemplateButton2 := messagesapi.TemplateButton{TemplateOpenUrlButton: givenButton2}

	givenTemplate := messagesapi.NewTemplate(givenTemplateName)
	givenTemplate.Language = &givenLanguage

	givenTemplateMessage := messagesapi.NewTemplateMessage(
		messagesapi.OUTBOUNDTEMPLATECHANNEL_WHATSAPP,
		givenSender,
		[]messagesapi.MessageDestination{destination},
		*givenTemplate,
	)

	givenMessageContent := messagesapi.TemplateMessageContent{
		Body:    &body,
		Buttons: []messagesapi.TemplateButton{givenTemplateButton1, givenTemplateButton2},
	}

	givenTemplateMessage.Content = &givenMessageContent

	request := messagesapi.NewRequest([]messagesapi.BaseMessage{
		{TemplateMessage: givenTemplateMessage},
	})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mock HTTP
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", SendMessagesApiMessages, givenResponse, 200)

	responseBody, r, _ := infobipClient.
		MessagesAPI.
		SendMessagesApiMessage(context.Background()).
		Request(*request).Execute()

	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.Len(t, responseBody.GetMessages(), 1)

	msg := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	messageStatus := msg.GetStatus()
	assert.Equal(t, int32(PendingStatusGroupId), messageStatus.GetGroupId())
	assert.Equal(t, messagesapi.MESSAGEGENERALSTATUS_PENDING, messageStatus.GetGroupName())
	assert.Equal(t, int32(PendingStatusId), messageStatus.GetId())
	assert.Equal(t, PendingStatusName, messageStatus.GetName())
	assert.Equal(t, PendingStatusDescription, messageStatus.GetDescription())
}

func TestGetMessagesApiDeliveryReports(t *testing.T) {
	givenBulkId := "1688025180464000013"
	givenMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	givenSender := "senderNumber"
	givenDestination := "41793026727"
	givenSentAt := "2024-02-06T14:18:29.797+0000"
	givenDoneAt := "2024-02-06T17:18:29.797+0000"
	givenMessageCount := 1
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"
	givenCampaignReferenceId := "campaignRef"

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"event": "DELIVERY",
				"channel": "WHATSAPP",
				"sender": "%s",
				"destination": "%s",
				"sentAt": "%s",
				"doneAt": "%s",
				"bulkId": "%s",
				"messageId": "%s",
				"messageCount": %d,
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				},
				"error": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s",
					"permanent": %t
				},
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				},
				"campaignReferenceId": "%s"
			}
		]
	}`,
		givenSender,
		givenDestination,
		givenSentAt,
		givenDoneAt,
		givenBulkId,
		givenMessageId,
		givenMessageCount,
		StatusGroupId, StatusGroupName, StatusId, StatusName, StatusDescription,
		NoErrorStatusGroupId, NoErrorStatusGroupName, NoErrorStatusId, NoErrorStatusName, NoErrorStatusDescription, NoErrorPermanent,
		givenEntityId,
		givenApplicationId,
		givenCampaignReferenceId,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", MessagesApiDeliveryReportsEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		GetMessagesApiDeliveryReports(context.Background()).
		BulkId(givenBulkId).
		MessageId(givenMessageId).
		Limit(1).
		EntityId(givenEntityId).
		ApplicationId(givenApplicationId).
		CampaignReferenceId(givenCampaignReferenceId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)

	result := results[0]
	assert.Equal(t, "DELIVERY", result.GetEvent())
	assert.Equal(t, givenSender, result.GetSender())
	assert.Equal(t, givenDestination, result.GetDestination())
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, int32(givenMessageCount), result.GetMessageCount())
	assert.Equal(t, givenCampaignReferenceId, result.GetCampaignReferenceId())

	status := result.GetStatus()
	assert.Equal(t, int32(StatusGroupId), status.GetGroupId())
	assert.Equal(t, StatusName, status.GetName())

	errorInfo := result.GetError()
	assert.Equal(t, int32(NoErrorStatusGroupId), errorInfo.GetGroupId())
	assert.Equal(t, NoErrorStatusName, errorInfo.GetName())
	assert.Equal(t, NoErrorPermanent, errorInfo.GetPermanent())

	platform := result.GetPlatform()
	assert.Equal(t, givenEntityId, platform.GetEntityId())
	assert.Equal(t, givenApplicationId, platform.GetApplicationId())
}

func TestGetMessagesApiDeliveryReportsWithChannel(t *testing.T) {
	givenBulkId := "bulk-123"
	givenMessageId := "msg-456"
	givenDestination := "41793026727"

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"event": "DELIVERY",
				"channel": "SMS",
				"destination": "%s",
				"bulkId": "%s",
				"messageId": "%s",
				"messageCount": 1,
				"status": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s"
				},
				"error": {
					"groupId": %d,
					"groupName": "%s",
					"id": %d,
					"name": "%s",
					"description": "%s",
					"permanent": %t
				}
			}
		]
	}`,
		givenDestination,
		givenBulkId,
		givenMessageId,
		StatusGroupId, StatusGroupName, StatusId, StatusName, StatusDescription,
		NoErrorStatusGroupId, NoErrorStatusGroupName, NoErrorStatusId, NoErrorStatusName, NoErrorStatusDescription, NoErrorPermanent,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", MessagesApiDeliveryReportsEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		GetMessagesApiDeliveryReports(context.Background()).
		Channel(messagesapi.CHANNEL_SMS).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)

	result := results[0]
	assert.Equal(t, "SMS", string(result.GetChannel()))
}

func TestGetMessagesApiInboundMessages_TextMessage(t *testing.T) {
	givenChannel := "SMS"
	givenSender := "48123234567"
	givenDestination := "48123098765"
	givenText := "Text message 123"
	givenCleanText := "Text message"
	givenReceivedAt := "2020-02-06T14:18:29.797+0000"
	givenMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	givenMessageCount := 1
	givenPendingMessageCount := 0
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
					{
						"text": "%s",
						"cleanText": "%s",
						"type": "TEXT"
					}
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"messageCount": %d,
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				},
				"event": "MO"
			}
		],
		"messageCount": %d,
		"pendingMessageCount": %d
	}`,
		givenChannel,
		givenSender,
		givenDestination,
		givenText,
		givenCleanText,
		givenReceivedAt,
		givenMessageId,
		givenMessageCount,
		givenEntityId,
		givenApplicationId,
		givenMessageCount,
		givenPendingMessageCount,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", MessagesApiInboundEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		GetMessagesApiInboundMessages(context.Background()).
		Channel(messagesapi.INBOUNDMOGETENDPOINTCHANNEL_SMS).
		Limit(10).
		EntityId(givenEntityId).
		ApplicationId(givenApplicationId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, int32(givenMessageCount), responseBody.GetMessageCount())
	assert.Equal(t, int32(givenPendingMessageCount), responseBody.GetPendingMessageCount())

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)

	result := results[0]
	assert.NotNil(t, result.MoEvent)
	assert.Equal(t, messagesapi.InboundMoEventChannel(givenChannel), result.MoEvent.GetChannel())
	assert.Equal(t, givenSender, result.MoEvent.GetSender())
	assert.Equal(t, givenDestination, result.MoEvent.GetDestination())
	assert.Equal(t, givenMessageId, result.MoEvent.GetMessageId())

	platform := result.MoEvent.GetPlatform()
	assert.Equal(t, givenEntityId, platform.GetEntityId())
	assert.Equal(t, givenApplicationId, platform.GetApplicationId())
}

func TestGetMessagesApiInboundMessages_TextMessageWithKeyword(t *testing.T) {
	givenChannel := "SMS"
	givenSender := "48123234567"
	givenDestination := "48123098765"
	givenText := "KWRDText message 123"
	givenCleanText := "Text message"
	givenKeyword := "KWRD"
	givenReceivedAt := "2020-02-06T14:18:29.797+0000"
	givenMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	givenMessageCount := 1
	givenPendingMessageCount := 0
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
					{
						"text": "%s",
						"cleanText": "%s",
						"keyword": "%s",
						"type": "TEXT"
					}
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"messageCount": %d,
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				},
				"event": "MO"
			}
		],
		"messageCount": %d,
		"pendingMessageCount": %d
	}`,
		givenChannel,
		givenSender,
		givenDestination,
		givenText,
		givenCleanText,
		givenKeyword,
		givenReceivedAt,
		givenMessageId,
		givenMessageCount,
		givenEntityId,
		givenApplicationId,
		givenMessageCount,
		givenPendingMessageCount,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", MessagesApiInboundEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		GetMessagesApiInboundMessages(context.Background()).
		Channel(messagesapi.INBOUNDMOGETENDPOINTCHANNEL_SMS).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)
}

func TestGetMessagesApiInboundMessages_ImageWithCaption(t *testing.T) {
	givenChannel := "WHATSAPP"
	givenSender := "48123234567"
	givenDestination := "48123098765"
	givenUrl := "http://my.domain/image.jpg"
	givenText := "Image caption"
	givenReceivedAt := "2020-02-06T14:18:29.797+0000"
	givenMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	givenMessageCount := 1
	givenPendingMessageCount := 0
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
					{
						"url": "%s",
						"text": "%s",
						"type": "IMAGE"
					}
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"messageCount": %d,
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				},
				"event": "MO"
			}
		],
		"messageCount": %d,
		"pendingMessageCount": %d
	}`,
		givenChannel,
		givenSender,
		givenDestination,
		givenUrl,
		givenText,
		givenReceivedAt,
		givenMessageId,
		givenMessageCount,
		givenEntityId,
		givenApplicationId,
		givenMessageCount,
		givenPendingMessageCount,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", MessagesApiInboundEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		GetMessagesApiInboundMessages(context.Background()).
		Channel(messagesapi.INBOUNDMOGETENDPOINTCHANNEL_WHATSAPP).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)
}

func TestGetMessagesApiInboundMessages_LocationMessage(t *testing.T) {
	givenChannel := "WHATSAPP"
	givenSender := "48123234567"
	givenDestination := "48123098765"
	givenLatitude := float64(16)
	givenLongitude := float64(18)
	givenUrl := "http://my.domain/media/my-location"
	givenReceivedAt := "2020-02-06T14:18:29.797+0000"
	givenMessageId := "ABEGVUGWh3gEAgo-sLTvmQCS5kwjhsy"
	givenMessageCount := 1
	givenPendingMessageCount := 0
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"channel": "%s",
				"sender": "%s",
				"destination": "%s",
				"content": [
					{
						"latitude": %g,
						"longitude": %g,
						"url": "%s",
						"type": "LOCATION"
					}
				],
				"receivedAt": "%s",
				"messageId": "%s",
				"messageCount": %d,
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				},
				"event": "MO"
			}
		],
		"messageCount": %d,
		"pendingMessageCount": %d
	}`,
		givenChannel,
		givenSender,
		givenDestination,
		givenLatitude,
		givenLongitude,
		givenUrl,
		givenReceivedAt,
		givenMessageId,
		givenMessageCount,
		givenEntityId,
		givenApplicationId,
		givenMessageCount,
		givenPendingMessageCount,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", MessagesApiInboundEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		GetMessagesApiInboundMessages(context.Background()).
		Channel(messagesapi.INBOUNDMOGETENDPOINTCHANNEL_WHATSAPP).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)
}

func TestValidateMessagesApiMessage_Success(t *testing.T) {
	givenChannel := "SMS"
	givenSender := "447491163862"
	givenTo := "123456789"
	givenText := "Sending you lots of otterly delightful vibes today!"

	givenRequest := fmt.Sprintf(`{
		"messages": [
			{
				"channel": "%s",
				"sender": "%s",
				"destinations": [
					{
						"to": "%s"
					}
				],
				"content": {
					"body": {
						"text": "%s",
						"type": "TEXT"
					}
				}
			}
		]
	}`,
		givenChannel,
		givenSender,
		givenTo,
		givenText,
	)

	givenResponse := `{
		"description": "Request can be sent through '/messages' endpoint and should be accepted by our platform.",
		"action": "No action is required, but it is recommended to check and address any violations.",
		"skippableViolations": [
			{
				"property": "messages[0].metadata",
				"violation": "Unknown property"
			}
		]
	}`

	body := messagesapi.MessageBody{
		MessageTextBody: messagesapi.NewMessageTextBody(givenText),
	}

	destination := messagesapi.MessageDestination{
		ToDestination: messagesapi.NewToDestination(givenTo),
	}

	givenMessage := messagesapi.NewMessage(
		messagesapi.OUTBOUNDMESSAGECHANNEL_SMS,
		givenSender,
		[]messagesapi.MessageDestination{
			destination,
		},
		messagesapi.MessageContent{
			Body: body,
		},
	)

	request := messagesapi.NewRequest([]messagesapi.BaseMessage{
		{Message: givenMessage},
	})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", MessagesApiValidateEndpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		MessagesAPI.
		ValidateMessagesApiMessage(context.Background()).
		Request(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, "Request can be sent through '/messages' endpoint and should be accepted by our platform.", responseBody.GetDescription())
	assert.Equal(t, "No action is required, but it is recommended to check and address any violations.", responseBody.GetAction())
	skippableViolations := responseBody.GetSkippableViolations()
	assert.Len(t, skippableViolations, 1)
	assert.Equal(t, "messages[0].metadata", skippableViolations[0].GetProperty())
	assert.Equal(t, "Unknown property", skippableViolations[0].GetViolation())
}

func TestValidateMessagesApiMessage_WithSkippableViolations(t *testing.T) {
	givenResponse := `{
		"description": "Request can be sent through '/messages' endpoint and should be accepted by our platform.",
		"action": "No action is required, but it is recommended to check and address any violations.",
		"skippableViolations": [
			{
				"property": "messages[0].metadata",
				"violation": "Unknown property"
			}
		]
	}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", MessagesApiValidateEndpoint, givenResponse, 200)

	body := messagesapi.MessageBody{
		MessageTextBody: messagesapi.NewMessageTextBody("Test message"),
	}

	destination := messagesapi.MessageDestination{
		ToDestination: messagesapi.NewToDestination("123456789"),
	}

	givenMessage := messagesapi.NewMessage(
		messagesapi.OUTBOUNDMESSAGECHANNEL_SMS,
		"447491163862",
		[]messagesapi.MessageDestination{destination},
		messagesapi.MessageContent{Body: body},
	)

	request := messagesapi.NewRequest([]messagesapi.BaseMessage{
		{Message: givenMessage},
	})

	responseBody, r, err := infobipClient.
		MessagesAPI.
		ValidateMessagesApiMessage(context.Background()).
		Request(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	skippableViolations := responseBody.GetSkippableViolations()
	assert.True(t, len(skippableViolations) == 1)
	assert.Equal(t, "messages[0].metadata", skippableViolations[0].GetProperty())
	assert.Equal(t, "Unknown property", skippableViolations[0].GetViolation())
}
