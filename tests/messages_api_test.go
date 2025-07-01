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
	SendMessagesApiMessages string = "/messages-api/1/messages"
	SendMessagesApiEvents   string = "/messages-api/1/events"
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

	givenChannel := "SMS"
	givenSender := "447491163443"
	givenTo := "111111111"
	givenText := "May the Force be with you."
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
		messagesapi.OUTBOUNDMESSAGECHANNEL_SMS,
		givenSender,
		[]messagesapi.MessageDestination{
			destination,
		},
		messagesapi.MessageContent{
			Body: body,
		},
	)

	request := messagesapi.NewRequest([]messagesapi.RequestMessagesInner{
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
				"callbackData": {
					"key1": "value1", 
					"key2": "value2"
				},
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
				"callbackData": {},
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
	givenTemplateName := "episodeV"
	givenLanguage := "en_GB"

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
					"1": "Luke",
					"2": "Skywalker",
					"type": "TEXT"
				  },
				  "buttons": [
					{
					  "postbackData": "I am the father",
					  "type": "QUICK_REPLY"
					},
					{
					  "postbackData": "I am not the father",
					  "type": "QUICK_REPLY"
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
	templateTextBody.AdditionalProperties["1"] = "Luke"
	templateTextBody.AdditionalProperties["2"] = "Skywalker"

	body := messagesapi.TemplateBody{
		TemplateTextBody: templateTextBody,
	}

	destination := messagesapi.MessageDestination{
		ToDestination: messagesapi.NewToDestination(givenTo),
	}

	givenButtonPostbackData1 := "I am the father"
	givenButtonPostbackData2 := "I am not the father"

	givenButton1 := messagesapi.NewTemplateQuickReplyButton(givenButtonPostbackData1)
	givenTemplateButton1 := messagesapi.TemplateButton{TemplateQuickReplyButton: givenButton1}
	givenButton2 := messagesapi.NewTemplateQuickReplyButton(givenButtonPostbackData2)
	givenTemplateButton2 := messagesapi.TemplateButton{TemplateQuickReplyButton: givenButton2}

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

	request := messagesapi.NewRequest([]messagesapi.RequestMessagesInner{
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
