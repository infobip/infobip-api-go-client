package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/rcs"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	RcsSendMessagesEndpoint          string = "/rcs/2/messages"
	RcsDeliveryReportsEndpoint       string = "/rcs/2/reports"
	RcsLogsEndpoint                  string = "/rcs/2/logs"
	RcsCapabilityCheckNotifyEndpoint string = "/rcs/2/capability-check/notify"
	RcsCapabilityCheckQueryEndpoint  string = "/rcs/2/capability-check/query"
)

type rcsMessageResponse struct {
	MessageId   string
	Destination string
}

// buildSendRcsMessagesResponse returns the common mock response body for SendRcsMessages tests.
func buildSendRcsMessagesResponse(bulkId string, messages ...rcsMessageResponse) string {
	type status struct {
		GroupId     int    `json:"groupId"`
		GroupName   string `json:"groupName"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	type message struct {
		MessageId   string `json:"messageId"`
		Status      status `json:"status"`
		Destination string `json:"destination"`
	}

	type response struct {
		BulkId   string    `json:"bulkId"`
		Messages []message `json:"messages"`
	}

	msgStatus := status{
		GroupId:     1,
		GroupName:   "PENDING",
		Id:          26,
		Name:        "PENDING_ACCEPTED",
		Description: "Message sent to next instance",
	}

	resp := response{BulkId: bulkId}
	for _, m := range messages {
		resp.Messages = append(resp.Messages, message{
			MessageId:   m.MessageId,
			Status:      msgStatus,
			Destination: m.Destination,
		})
	}

	encoded, _ := json.Marshal(resp)
	return string(encoded)
}

func TestSendRcsMessages_BasicTextMessage(t *testing.T) {
	// Test data from OpenAPI example: "Basic text message"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"
	givenText := "Some text"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"text": "Some text",
						"type": "TEXT"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	textContent := rcs.NewOutboundTextContent(givenText)
	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundTextContentAsOutboundMessageContent(textContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))

	msg := messages[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	status := msg.GetStatus()
	assert.Equal(t, int32(1), status.GetGroupId())
	assert.Equal(t, rcs.MESSAGEGENERALSTATUS_PENDING, status.GetGroupName())
	assert.Equal(t, int32(26), status.GetId())
	assert.Equal(t, "PENDING_ACCEPTED", status.GetName())
	assert.Equal(t, "Message sent to next instance", status.GetDescription())
}

func TestSendRcsMessages_TextMessageWithReplySuggestion(t *testing.T) {
	// Test data from OpenAPI example: "Single text message"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"
	givenText := "Some text"
	givenSuggestionText := "Example text"
	givenPostbackData := "Example postback data"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"suggestions": [
							{
								"postbackData": "Example postback data",
								"text": "Example text",
								"type": "REPLY"
							}
						],
						"text": "Some text",
						"type": "TEXT"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build request with reply suggestion
	replySuggestion := rcs.NewReplySuggestion(givenSuggestionText, givenPostbackData)
	suggestion := rcs.ReplySuggestionAsSuggestion(replySuggestion)

	textContent := rcs.NewOutboundTextContent(givenText)
	textContent.SetSuggestions([]rcs.Suggestion{suggestion})

	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundTextContentAsOutboundMessageContent(textContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))

	msg := messages[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenDestination, msg.GetDestination())

	status := msg.GetStatus()
	assert.Equal(t, int32(1), status.GetGroupId())
	assert.Equal(t, "PENDING_ACCEPTED", status.GetName())
}

func TestSendRcsMessages_FileMessage(t *testing.T) {
	// Test data from OpenAPI example: "File message"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"
	givenFileUrl := "https://www.example.com/video.mp4"
	givenThumbnailUrl := "https://www.example.com/thumbnail.jpg"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"file": {
							"url": "https://www.example.com/video.mp4"
						},
						"suggestions": [
							{
								"postbackData": "Example postback data",
								"text": "Example text",
								"type": "REPLY"
							}
						],
						"thumbnail": {
							"url": "https://www.example.com/thumbnail.jpg"
						},
						"type": "FILE"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build file content
	fileResource := rcs.NewResourceSchema(givenFileUrl)
	fileContent := rcs.NewOutboundFileContent(*fileResource)

	thumbnail := rcs.NewResource(givenThumbnailUrl)
	fileContent.SetThumbnail(*thumbnail)

	// Add reply suggestion
	replySuggestion := rcs.NewReplySuggestion("Example text", "Example postback data")
	fileContent.SetSuggestions([]rcs.Suggestion{rcs.ReplySuggestionAsSuggestion(replySuggestion)})

	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundFileContentAsOutboundMessageContent(fileContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, givenMessageId, messages[0].GetMessageId())
	assert.Equal(t, givenDestination, messages[0].GetDestination())
}

func TestSendRcsMessages_CardMessage(t *testing.T) {
	// Test data from OpenAPI example: "Card message"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"
	givenTitle := "Some title"
	givenDescription := "Some description"
	givenFileUrl := "https://www.example.com/video.mp4"
	givenThumbnailUrl := "https://www.example.com/thumbnail.jpg"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"alignment": "LEFT",
						"content": {
							"description": "Some description",
							"media": {
								"file": {
									"url": "https://www.example.com/video.mp4"
								},
								"height": "TALL",
								"thumbnail": {
									"url": "https://www.example.com/thumbnail.jpg"
								}
							},
							"suggestions": [
								{
									"application": "BROWSER",
									"postbackData": "Example postback data",
									"text": "Example text",
									"type": "OPEN_URL",
									"url": "https://www.example.com/"
								},
								{
									"postbackData": "Example postback data",
									"text": "Example text",
									"type": "REPLY"
								}
							],
							"title": "Some title"
						},
						"orientation": "HORIZONTAL",
						"suggestions": [
							{
								"postbackData": "Example postback data",
								"text": "Example text",
								"type": "REPLY"
							}
						],
						"type": "CARD"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build card content
	fileResource := rcs.NewCardResourceSchema(givenFileUrl)
	thumbnail := rcs.NewResource(givenThumbnailUrl)
	cardMedia := rcs.NewCardMedia(*fileResource, rcs.HEIGHT_TALL)
	cardMedia.SetThumbnail(*thumbnail)

	// Card content suggestions (inside the card)
	openUrlSuggestion := rcs.NewOpenUrlSuggestion("Example text", "Example postback data", "https://www.example.com/")
	openUrlSuggestion.Application = (*rcs.OpenUrlApplicationType)(ptr("BROWSER"))
	replySuggestion := rcs.NewReplySuggestion("Example text", "Example postback data")

	cardContentInner := rcs.NewCardContent()
	cardContentInner.SetTitle(givenTitle)
	cardContentInner.SetDescription(givenDescription)
	cardContentInner.SetMedia(*cardMedia)
	cardContentInner.SetSuggestions([]rcs.Suggestion{
		rcs.OpenUrlSuggestionAsSuggestion(openUrlSuggestion),
		rcs.ReplySuggestionAsSuggestion(replySuggestion),
	})

	// Build card outer content
	cardContent := rcs.NewOutboundCardContent(rcs.ORIENTATION_HORIZONTAL, rcs.ALIGNMENT_LEFT, *cardContentInner)

	// Outer suggestions
	outerReplySuggestion := rcs.NewReplySuggestion("Example text", "Example postback data")
	cardContent.SetSuggestions([]rcs.Suggestion{rcs.ReplySuggestionAsSuggestion(outerReplySuggestion)})

	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundCardContentAsOutboundMessageContent(cardContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, givenMessageId, messages[0].GetMessageId())
	assert.Equal(t, givenDestination, messages[0].GetDestination())
}

func TestSendRcsMessages_ProviderTemplateMessage(t *testing.T) {
	// Test data from OpenAPI example: "Provider template message"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"
	givenTemplateId := "ExampleTemplateId"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"parameters": [
							{
								"name": "param1",
								"value": "paramValue1"
							},
							{
								"name": "param2",
								"value": "paramValue2"
							}
						],
						"templateId": "ExampleTemplateId",
						"type": "PROVIDER_TEMPLATE"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build provider template content
	param1 := rcs.NewParameter("param1", "paramValue1")
	param2 := rcs.NewParameter("param2", "paramValue2")

	templateContent := rcs.NewOutboundProviderTemplateContent(givenTemplateId)
	templateContent.SetParameters([]rcs.Parameter{*param1, *param2})

	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundProviderTemplateContentAsOutboundMessageContent(templateContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, givenMessageId, messages[0].GetMessageId())
	assert.Equal(t, givenDestination, messages[0].GetDestination())
}

func TestSendRcsMessages_TextWithSmsFailover(t *testing.T) {
	// Test data from OpenAPI example: "Text message with SMS failover"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"
	givenText := "Some text"
	givenFailoverText := "Some failover text"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"suggestions": [
							{
								"postbackData": "Example postback data",
								"text": "Example text",
								"type": "REPLY"
							}
						],
						"text": "Some text",
						"type": "TEXT"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"options": {
						"campaignReferenceId": "Example campaign id",
						"deliveryTimeWindow": {
							"days": [
								"MONDAY",
								"TUESDAY"
							],
							"from": {
								"hour": 9,
								"minute": 0
							},
							"to": {
								"hour": 17,
								"minute": 0
							}
						},
						"platform": {
							"applicationId": "Example application id",
							"entityId": "Example entity id"
						},
						"smsFailover": {
							"sender": "DemoSender",
							"text": "Some failover text",
							"validityPeriod": {
								"amount": 1,
								"timeUnit": "HOURS"
							}
						},
						"validityPeriod": {
							"amount": 1,
							"timeUnit": "HOURS"
						}
					},
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build text content with suggestion
	replySuggestion := rcs.NewReplySuggestion("Example text", "Example postback data")
	textContent := rcs.NewOutboundTextContent(givenText)
	textContent.SetSuggestions([]rcs.Suggestion{rcs.ReplySuggestionAsSuggestion(replySuggestion)})

	// Build SMS failover
	smsFailover := rcs.NewDefaultSmsFailover(givenFailoverText)
	smsFailover.SetSender(givenSender)
	validityPeriod := rcs.NewValidityPeriod(1)
	validityPeriod.SetTimeUnit(rcs.VALIDITYPERIODTIMEUNIT_HOURS)
	smsFailover.SetValidityPeriod(*validityPeriod)

	// Build failover options with delivery time window
	options := rcs.NewFailoverOptions()
	options.SetSmsFailover(*smsFailover)
	options.SetCampaignReferenceId("Example campaign id")

	platform := rcs.NewPlatform()
	platform.SetEntityId("Example entity id")
	platform.SetApplicationId("Example application id")
	options.SetPlatform(*platform)

	msgValidityPeriod := rcs.NewValidityPeriod(1)
	msgValidityPeriod.SetTimeUnit(rcs.VALIDITYPERIODTIMEUNIT_HOURS)
	options.SetValidityPeriod(*msgValidityPeriod)

	deliveryTimeWindow := rcs.NewDeliveryTimeWindow([]rcs.DeliveryDay{rcs.DELIVERYDAY_MONDAY, rcs.DELIVERYDAY_TUESDAY})
	fromTime := rcs.NewDeliveryTime(9, 0)
	toTime := rcs.NewDeliveryTime(17, 0)
	deliveryTimeWindow.SetFrom(*fromTime)
	deliveryTimeWindow.SetTo(*toTime)
	options.SetDeliveryTimeWindow(*deliveryTimeWindow)

	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundTextContentAsOutboundMessageContent(textContent))
	message.SetOptions(*options)

	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, givenMessageId, messages[0].GetMessageId())
	assert.Equal(t, givenDestination, messages[0].GetDestination())
}

func TestSendRcsMessages_CarouselMessage(t *testing.T) {
	// Test data from OpenAPI example: "Carousel message"
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenDestination := "441134960001"
	givenSender := "DemoSender"

	givenResponse := buildSendRcsMessagesResponse(givenBulkId, rcsMessageResponse{MessageId: givenMessageId, Destination: givenDestination})

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"cardWidth": "MEDIUM",
						"contents": [
							{
								"description": "Some description",
								"media": {
									"file": {
										"url": "https://www.example.com/video.mp4"
									},
									"height": "MEDIUM",
									"thumbnail": {
										"url": "https://www.example.com/thumbnail.jpg"
									}
								},
								"suggestions": [
									{
										"application": "BROWSER",
										"postbackData": "Example postback data",
										"text": "Example text",
										"type": "OPEN_URL",
										"url": "https://www.example.com/"
									},
									{
										"postbackData": "Example postback data",
										"text": "Example text",
										"type": "REPLY"
									}
								],
								"title": "Some title"
							},
							{
								"description": "Another description",
								"media": {
									"file": {
										"url": "https://www.sample.com/video_2.mp4"
									},
									"height": "MEDIUM",
									"thumbnail": {
										"url": "https://www.sample.com/thumbnail_2.jpg"
									}
								},
								"suggestions": [
									{
										"application": "WEBVIEW",
										"postbackData": "Another example postback data",
										"text": "Another example text",
										"type": "OPEN_URL",
										"url": "https://www.sample.com/",
										"webviewViewMode": "FULL"
									},
									{
										"postbackData": "Another example postback data",
										"text": "Another example text",
										"type": "REPLY"
									}
								],
								"title": "Another title"
							}
						],
						"suggestions": [
							{
								"postbackData": "Example postback data",
								"text": "Example text",
								"type": "REPLY"
							}
						],
						"type": "CAROUSEL"
					},
					"destinations": [
						{
							"to": "441134960001"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build first card
	fileResource1 := rcs.NewCardResourceSchema("https://www.example.com/video.mp4")
	thumbnail1 := rcs.NewResource("https://www.example.com/thumbnail.jpg")
	cardMedia1 := rcs.NewCardMedia(*fileResource1, rcs.HEIGHT_MEDIUM)
	cardMedia1.SetThumbnail(*thumbnail1)

	openUrlSuggestion1 := rcs.NewOpenUrlSuggestion("Example text", "Example postback data", "https://www.example.com/")
	openUrlSuggestion1.Application = (*rcs.OpenUrlApplicationType)(ptr("BROWSER"))
	replySuggestion1 := rcs.NewReplySuggestion("Example text", "Example postback data")

	cardContent1 := rcs.NewCardContent()
	cardContent1.SetTitle("Some title")
	cardContent1.SetDescription("Some description")
	cardContent1.SetMedia(*cardMedia1)
	cardContent1.SetSuggestions([]rcs.Suggestion{
		rcs.OpenUrlSuggestionAsSuggestion(openUrlSuggestion1),
		rcs.ReplySuggestionAsSuggestion(replySuggestion1),
	})

	// Build second card
	fileResource2 := rcs.NewCardResourceSchema("https://www.sample.com/video_2.mp4")
	thumbnail2 := rcs.NewResource("https://www.sample.com/thumbnail_2.jpg")
	cardMedia2 := rcs.NewCardMedia(*fileResource2, rcs.HEIGHT_MEDIUM)
	cardMedia2.SetThumbnail(*thumbnail2)

	openUrlSuggestion2 := rcs.NewOpenUrlSuggestion("Another example text", "Another example postback data", "https://www.sample.com/")
	openUrlSuggestion2.Application = (*rcs.OpenUrlApplicationType)(ptr("WEBVIEW"))
	webviewMode := rcs.WEBVIEWVIEWMODETYPE_FULL
	openUrlSuggestion2.WebviewViewMode = &webviewMode
	replySuggestion2 := rcs.NewReplySuggestion("Another example text", "Another example postback data")

	cardContent2 := rcs.NewCardContent()
	cardContent2.SetTitle("Another title")
	cardContent2.SetDescription("Another description")
	cardContent2.SetMedia(*cardMedia2)
	cardContent2.SetSuggestions([]rcs.Suggestion{
		rcs.OpenUrlSuggestionAsSuggestion(openUrlSuggestion2),
		rcs.ReplySuggestionAsSuggestion(replySuggestion2),
	})

	// Build carousel content
	carouselContent := rcs.NewOutboundCarouselContent(rcs.WIDTH_MEDIUM, []rcs.CardContent{*cardContent1, *cardContent2})

	// Outer suggestions
	outerReplySuggestion := rcs.NewReplySuggestion("Example text", "Example postback data")
	carouselContent.SetSuggestions([]rcs.Suggestion{rcs.ReplySuggestionAsSuggestion(outerReplySuggestion)})

	destination := rcs.NewToDestination(givenDestination)
	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination}, rcs.OutboundCarouselContentAsOutboundMessageContent(carouselContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, givenMessageId, messages[0].GetMessageId())
	assert.Equal(t, givenDestination, messages[0].GetDestination())
}

func TestSendRcsMessages_MultipleDestinations(t *testing.T) {
	// Test sending to multiple destinations
	givenBulkId := "a28dd97c-2222-4fcf-99f1-0b557ed381da"
	givenMessageId1 := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenMessageId2 := "b38ee97d-2ffc-5gdf-00g2-1c668ed492eb"
	givenDestination1 := "441134960001"
	givenDestination2 := "441134960002"
	givenSender := "DemoSender"
	givenText := "Hello from Infobip!"

	givenResponse := buildSendRcsMessagesResponse(
		givenBulkId,
		rcsMessageResponse{MessageId: givenMessageId1, Destination: givenDestination1},
		rcsMessageResponse{MessageId: givenMessageId2, Destination: givenDestination2},
	)

	givenRequest := `
		{
			"messages": [
				{
					"content": {
						"text": "Hello from Infobip!",
						"type": "TEXT"
					},
					"destinations": [
						{
							"to": "441134960001"
						},
						{
							"to": "441134960002"
						}
					],
					"sender": "DemoSender"
				}
			]
		}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsSendMessagesEndpoint, givenResponse, 200)

	// Build request with multiple destinations
	textContent := rcs.NewOutboundTextContent(givenText)

	destination1 := rcs.NewToDestination(givenDestination1)
	destination2 := rcs.NewToDestination(givenDestination2)

	message := rcs.NewMessage(givenSender, []rcs.ToDestination{*destination1, *destination2}, rcs.OutboundTextContentAsOutboundMessageContent(textContent))
	request := rcs.NewMessageRequest([]rcs.Message{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		SendRcsMessages(context.Background()).
		MessageRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	messages := responseBody.GetMessages()
	assert.Equal(t, 2, len(messages))

	assert.Equal(t, givenMessageId1, messages[0].GetMessageId())
	assert.Equal(t, givenDestination1, messages[0].GetDestination())

	assert.Equal(t, givenMessageId2, messages[1].GetMessageId())
	assert.Equal(t, givenDestination2, messages[1].GetDestination())
}

func TestGetOutboundRcsMessageDeliveryReports(t *testing.T) {
	givenBulkId := "bulk-123-rcs"
	givenMessageId := "msg-456-rcs"
	givenTo := "385916242493"
	givenSentAt := "2021-08-25T16:00:00.000+0000"
	givenDoneAt := "2021-08-25T16:00:00.500+0000"
	givenPricePerMessage := 0.05
	givenCurrency := "EUR"

	givenResponse := fmt.Sprintf(`
		{
			"results": [
				{
					"bulkId": "%s",
					"messageId": "%s",
					"to": "%s",
					"sentAt": "%s",
					"doneAt": "%s",
					"messageCount": 1,
					"price": {
						"pricePerMessage": %g,
						"currency": "%s"
					},
					"status": {
						"groupId": 3,
						"groupName": "DELIVERED",
						"id": 5,
						"name": "DELIVERED_TO_HANDSET",
						"description": "Message delivered to handset"
					},
					"error": {
						"groupId": 0,
						"groupName": "OK",
						"id": 0,
						"name": "NO_ERROR",
						"description": "No error",
						"permanent": false
					}
				}
			]
		}`,
		givenBulkId,
		givenMessageId,
		givenTo,
		givenSentAt,
		givenDoneAt,
		givenPricePerMessage,
		givenCurrency,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", RcsDeliveryReportsEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		GetOutboundRcsMessageDeliveryReports(context.Background()).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)

	result := results[0]
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenTo, result.GetTo())

	price := result.GetPrice()
	assert.NotNil(t, price)
	assert.Equal(t, float32(givenPricePerMessage), price.GetPricePerMessage())
	assert.Equal(t, givenCurrency, price.GetCurrency())

	status := result.GetStatus()
	assert.Equal(t, int32(3), status.GetGroupId())
	assert.Equal(t, int32(5), status.GetId())
	assert.Equal(t, "DELIVERED_TO_HANDSET", status.GetName())

	errorInfo := result.GetError()
	assert.Equal(t, int32(0), errorInfo.GetGroupId())
	assert.Equal(t, "NO_ERROR", errorInfo.GetName())
}

func TestGetOutboundRcsMessageLogs(t *testing.T) {
	givenBulkId := "bulk-123-rcs"
	givenMessageId := "msg-456-rcs"
	givenDestination := "385916242493"
	givenSentAt := "2021-08-25T16:00:00.000+0000"
	givenDoneAt := "2021-08-25T16:00:00.500+0000"

	givenResponse := fmt.Sprintf(`
		{
			"results": [
				{
					"bulkId": "%s",
					"messageId": "%s",
					"destination": "%s",
					"sentAt": "%s",
					"doneAt": "%s",
					"messageCount": 1,
					"status": {
						"groupId": 3,
						"groupName": "DELIVERED",
						"id": 5,
						"name": "DELIVERED_TO_HANDSET",
						"description": "Message delivered to handset"
					},
					"error": {
						"groupId": 0,
						"groupName": "OK",
						"id": 0,
						"name": "NO_ERROR",
						"description": "No error",
						"permanent": false
					}
				}
			]
		}`,
		givenBulkId,
		givenMessageId,
		givenDestination,
		givenSentAt,
		givenDoneAt,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", RcsLogsEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		GetOutboundRcsMessageLogs(context.Background()).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetResults()
	assert.True(t, len(results) == 1)

	result := results[0]
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenDestination, result.GetDestination())

	status := result.GetStatus()
	assert.Equal(t, int32(3), status.GetGroupId())
	assert.Equal(t, int32(5), status.GetId())
	assert.Equal(t, "DELIVERED_TO_HANDSET", status.GetName())
}

func TestCapabilityCheckRcsDestinationsNotify(t *testing.T) {
	// Test data from OpenAPI example
	givenBulkId := "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8"
	givenMessageId1 := "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8"
	givenMessageId2 := "b5c3bdff-2d44-4f74-8a8e-3792fa57dfc3"
	givenPhoneNumber1 := "441134960001"
	givenPhoneNumber2 := "441134960002"
	givenSender := "DemoSender"

	givenResponse := fmt.Sprintf(`
		{
			"bulkId": "%s",
			"capabilityCheckRequestStates": [
				{
					"messageId": "%s",
					"phoneNumber": "%s",
					"status": "PENDING_ENROUTE"
				},
				{
					"messageId": "%s",
					"phoneNumber": "%s",
					"status": "PENDING_ENROUTE"
				}
			]
		}`,
		givenBulkId,
		givenMessageId1,
		givenPhoneNumber1,
		givenMessageId2,
		givenPhoneNumber2,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsCapabilityCheckNotifyEndpoint, givenResponse, 200)

	// Build request using SDK models
	request := rcs.NewCapabilityCheckAsyncRequest(givenSender, []string{givenPhoneNumber1, givenPhoneNumber2})
	request.SetNotifyUrl("http://example.com/notify")
	request.SetNotifyContentType("application/json")

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		CapabilityCheckRcsDestinationsNotify(context.Background()).
		CapabilityCheckAsyncRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())

	states := responseBody.GetCapabilityCheckRequestStates()
	assert.Equal(t, 2, len(states))

	assert.Equal(t, givenMessageId1, states[0].GetMessageId())
	assert.Equal(t, givenPhoneNumber1, states[0].GetPhoneNumber())
	assert.Equal(t, rcs.STATUSREASON_PENDING_ENROUTE, states[0].GetStatus())

	assert.Equal(t, givenMessageId2, states[1].GetMessageId())
	assert.Equal(t, givenPhoneNumber2, states[1].GetPhoneNumber())
	assert.Equal(t, rcs.STATUSREASON_PENDING_ENROUTE, states[1].GetStatus())
}

func TestCapabilityCheckRcsDestinationsNotify_WithPlatformOptions(t *testing.T) {
	// Test with platform options from OpenAPI example
	givenBulkId := "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8"
	givenSender := "DemoSender"
	givenPhoneNumber := "441134960001"
	givenEntityId := "Example entity id"
	givenApplicationId := "Example application id"

	givenResponse := fmt.Sprintf(`
		{
			"bulkId": "%s",
			"capabilityCheckRequestStates": [
				{
					"messageId": "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8",
					"phoneNumber": "%s",
					"status": "PENDING_ENROUTE"
				}
			]
		}`,
		givenBulkId,
		givenPhoneNumber,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsCapabilityCheckNotifyEndpoint, givenResponse, 200)

	// Build request with platform options
	platform := rcs.NewPlatform()
	platform.SetEntityId(givenEntityId)
	platform.SetApplicationId(givenApplicationId)

	options := rcs.NewCapabilityCheckOptions()
	options.SetPlatform(*platform)

	request := rcs.NewCapabilityCheckAsyncRequest(givenSender, []string{givenPhoneNumber})
	request.SetOptions(*options)

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		CapabilityCheckRcsDestinationsNotify(context.Background()).
		CapabilityCheckAsyncRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")
	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
}

func TestCapabilityCheckRcsDestinationsQuery(t *testing.T) {
	// Test data from OpenAPI example
	givenMessageId1 := "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8"
	givenMessageId2 := "b5c3bdff-2d44-4f74-8a8e-3792fa57dfc3"
	givenPhoneNumber1 := "441134960001"
	givenPhoneNumber2 := "441134960002"
	givenSender := "DemoSender"

	givenResponse := fmt.Sprintf(`
		{
			"capabilityCheckResults": [
				{
					"messageId": "%s",
					"phoneNumber": "%s",
					"code": "ENABLED"
				},
				{
					"messageId": "%s",
					"phoneNumber": "%s",
					"code": "UNREACHABLE"
				}
			]
		}`,
		givenMessageId1,
		givenPhoneNumber1,
		givenMessageId2,
		givenPhoneNumber2,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsCapabilityCheckQueryEndpoint, givenResponse, 200)

	// Build request using SDK models
	request := rcs.NewCapabilityCheckSyncRequest(givenSender, []string{givenPhoneNumber1, givenPhoneNumber2})

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		CapabilityCheckRcsDestinationsQuery(context.Background()).
		CapabilityCheckSyncRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetCapabilityCheckResults()
	assert.Equal(t, 2, len(results))

	// First result - ENABLED
	assert.Equal(t, givenMessageId1, results[0].GetMessageId())
	assert.Equal(t, givenPhoneNumber1, results[0].GetPhoneNumber())
	assert.Equal(t, rcs.CAPABILITYCHECKRESPONSECODE_ENABLED, results[0].GetCode())

	// Second result - UNREACHABLE
	assert.Equal(t, givenMessageId2, results[1].GetMessageId())
	assert.Equal(t, givenPhoneNumber2, results[1].GetPhoneNumber())
	assert.Equal(t, rcs.CAPABILITYCHECKRESPONSECODE_UNREACHABLE, results[1].GetCode())
}

func TestCapabilityCheckRcsDestinationsQuery_WithPlatformOptions(t *testing.T) {
	// Test with platform options from OpenAPI example
	givenMessageId := "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8"
	givenPhoneNumber := "441134960001"
	givenSender := "DemoSender"
	givenEntityId := "Example entity id"
	givenApplicationId := "Example application id"

	givenResponse := fmt.Sprintf(`
		{
			"capabilityCheckResults": [
				{
					"messageId": "%s",
					"phoneNumber": "%s",
					"code": "ENABLED"
				}
			],
			"options": {
				"platform": {
					"entityId": "%s",
					"applicationId": "%s"
				}
			}
		}`,
		givenMessageId,
		givenPhoneNumber,
		givenEntityId,
		givenApplicationId,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", RcsCapabilityCheckQueryEndpoint, givenResponse, 200)

	// Build request with platform options
	platform := rcs.NewPlatform()
	platform.SetEntityId(givenEntityId)
	platform.SetApplicationId(givenApplicationId)

	options := rcs.NewCapabilityCheckOptions()
	options.SetPlatform(*platform)

	request := rcs.NewCapabilityCheckSyncRequest(givenSender, []string{givenPhoneNumber})
	request.SetOptions(*options)

	// Execute the request
	responseBody, r, err := infobipClient.
		RcsAPI.
		CapabilityCheckRcsDestinationsQuery(context.Background()).
		CapabilityCheckSyncRequest(*request).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	results := responseBody.GetCapabilityCheckResults()
	assert.Equal(t, 1, len(results))
	assert.Equal(t, rcs.CAPABILITYCHECKRESPONSECODE_ENABLED, results[0].GetCode())

	// Verify options are returned
	returnedOptions := responseBody.GetOptions()
	returnedPlatform := returnedOptions.GetPlatform()
	assert.Equal(t, givenEntityId, returnedPlatform.GetEntityId())
	assert.Equal(t, givenApplicationId, returnedPlatform.GetApplicationId())
}

// ==============================================================================
// Webhook deserialization tests
// ==============================================================================

func TestReceiveRcsDeliveryReportsWebhook(t *testing.T) {
	givenBulkId := "bulk-123-rcs"
	givenMessageId := "msg-456-rcs"
	givenTo := "441134960001"
	givenSender := "DemoSender"
	givenSentAt := "2025-02-06T15:35:12.000+0000"
	givenDoneAt := "2025-02-06T15:35:14.000+0000"
	givenMessageCount := 1
	givenMccMnc := "21901"
	givenCallbackData := "callbackData"
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"
	givenCampaignReferenceId := "campaignRef"

	givenPayload := fmt.Sprintf(`
		{
			"results": [
				{
					"bulkId": "%s",
					"messageId": "%s",
					"to": "%s",
					"sender": "%s",
					"sentAt": "%s",
					"doneAt": "%s",
					"messageCount": %d,
					"mccMnc": "%s",
					"callbackData": "%s",
					"price": {
						"pricePerMessage": 0.01,
						"currency": "EUR"
					},
					"status": {
						"groupId": 3,
						"groupName": "DELIVERED",
						"id": 5,
						"name": "DELIVERED_TO_HANDSET",
						"description": "Message delivered to handset"
					},
					"error": {
						"groupId": 0,
						"groupName": "OK",
						"id": 0,
						"name": "NO_ERROR",
						"description": "No error",
						"permanent": false
					},
					"platform": {
						"entityId": "%s",
						"applicationId": "%s"
					},
					"campaignReferenceId": "%s"
				}
			]
		}`,
		givenBulkId,
		givenMessageId,
		givenTo,
		givenSender,
		givenSentAt,
		givenDoneAt,
		givenMessageCount,
		givenMccMnc,
		givenCallbackData,
		givenEntityId,
		givenApplicationId,
		givenCampaignReferenceId,
	)

	var responseBody rcs.DeliveryReports
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, responseBody, "Expected non-nil response body")
	assert.NotNil(t, responseBody.Results, "Expected non-nil results")
	assert.Equal(t, 1, len(responseBody.GetResults()))

	result := responseBody.GetResults()[0]
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenSender, result.GetSender())
	assert.Equal(t, int32(givenMessageCount), result.GetMessageCount())
	assert.Equal(t, givenMccMnc, result.GetMccMnc())
	assert.Equal(t, givenCallbackData, result.GetCallbackData())
	assert.Equal(t, givenCampaignReferenceId, result.GetCampaignReferenceId())

	price := result.GetPrice()
	assert.Equal(t, float32(0.01), price.GetPricePerMessage())
	assert.Equal(t, "EUR", price.GetCurrency())

	status := result.GetStatus()
	assert.Equal(t, int32(3), status.GetGroupId())
	assert.Equal(t, rcs.MESSAGEGENERALSTATUS_DELIVERED, status.GetGroupName())
	assert.Equal(t, int32(5), status.GetId())
	assert.Equal(t, "DELIVERED_TO_HANDSET", status.GetName())
	assert.Equal(t, "Message delivered to handset", status.GetDescription())

	errorInfo := result.GetError()
	assert.Equal(t, int32(0), errorInfo.GetGroupId())
	assert.Equal(t, rcs.MESSAGEERRORGROUP_OK, errorInfo.GetGroupName())
	assert.Equal(t, int32(0), errorInfo.GetId())
	assert.Equal(t, "NO_ERROR", errorInfo.GetName())
	assert.Equal(t, false, errorInfo.GetPermanent())

	platform := result.GetPlatform()
	assert.Equal(t, givenEntityId, platform.GetEntityId())
	assert.Equal(t, givenApplicationId, platform.GetApplicationId())

	// Manual construction to verify parity
	expectedPrice := rcs.NewMessagePrice()
	expectedPrice.SetPricePerMessage(0.01)
	expectedPrice.SetCurrency("EUR")

	expectedStatus := rcs.NewMessageStatus()
	expectedStatus.SetGroupId(3)
	expectedStatus.SetGroupName(rcs.MESSAGEGENERALSTATUS_DELIVERED)
	expectedStatus.SetId(5)
	expectedStatus.SetName("DELIVERED_TO_HANDSET")
	expectedStatus.SetDescription("Message delivered to handset")

	expectedError := rcs.NewMessageError()
	expectedError.SetGroupId(0)
	expectedError.SetGroupName(rcs.MESSAGEERRORGROUP_OK)
	expectedError.SetId(0)
	expectedError.SetName("NO_ERROR")
	expectedError.SetDescription("No error")
	expectedError.SetPermanent(false)

	expectedPlatform := rcs.NewPlatform()
	expectedPlatform.SetEntityId(givenEntityId)
	expectedPlatform.SetApplicationId(givenApplicationId)

	expectedReport := rcs.NewDeliveryReport()
	expectedReport.SetBulkId(givenBulkId)
	expectedReport.SetMessageId(givenMessageId)
	expectedReport.SetTo(givenTo)
	expectedReport.SetSender(givenSender)
	expectedReport.SetSentAt(ibTime(givenSentAt))
	expectedReport.SetDoneAt(ibTime(givenDoneAt))
	expectedReport.SetMessageCount(int32(givenMessageCount))
	expectedReport.SetMccMnc(givenMccMnc)
	expectedReport.SetCallbackData(givenCallbackData)
	expectedReport.SetPrice(*expectedPrice)
	expectedReport.SetStatus(*expectedStatus)
	expectedReport.SetError(*expectedError)
	expectedReport.SetPlatform(*expectedPlatform)
	expectedReport.SetCampaignReferenceId(givenCampaignReferenceId)

	expectedBody := rcs.NewDeliveryReports()
	expectedBody.SetResults([]rcs.DeliveryReport{*expectedReport})

	assert.EqualValues(t, expectedBody.GetResults(), responseBody.GetResults())
}

func TestReceiveRcsDeliveryReportsWebhook_WithChannel(t *testing.T) {
	givenPayload := `
		{
			"results": [
				{
					"bulkId": "bulk-123",
					"channel": "RCS",
					"messageId": "msg-456",
					"to": "441134960001",
					"sender": "DemoSender",
					"sentAt": "2025-02-06T15:35:12.000+0000",
					"doneAt": "2025-02-06T15:35:14.000+0000",
					"messageCount": 1,
					"status": {
						"groupId": 3,
						"groupName": "DELIVERED",
						"id": 5,
						"name": "DELIVERED_TO_HANDSET",
						"description": "Message delivered to handset"
					},
					"error": {
						"groupId": 0,
						"groupName": "OK",
						"id": 0,
						"name": "NO_ERROR",
						"description": "No error",
						"permanent": false
					}
				},
				{
					"bulkId": "bulk-123",
					"channel": "SMS",
					"messageId": "msg-789",
					"to": "441134960001",
					"sender": "DemoSender",
					"sentAt": "2025-02-06T15:36:12.000+0000",
					"doneAt": "2025-02-06T15:36:14.000+0000",
					"messageCount": 1,
					"status": {
						"groupId": 3,
						"groupName": "DELIVERED",
						"id": 5,
						"name": "DELIVERED_TO_HANDSET",
						"description": "Message delivered to handset"
					},
					"error": {
						"groupId": 0,
						"groupName": "OK",
						"id": 0,
						"name": "NO_ERROR",
						"description": "No error",
						"permanent": false
					}
				}
			]
		}`

	var responseBody rcs.WebhookDeliveryReports
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, 2, len(responseBody.GetResults()))

	// First result is RCS channel
	assert.Equal(t, "RCS", responseBody.GetResults()[0].GetChannel())
	assert.Equal(t, "msg-456", responseBody.GetResults()[0].GetMessageId())

	// Second result is SMS failover channel
	assert.Equal(t, "SMS", responseBody.GetResults()[1].GetChannel())
	assert.Equal(t, "msg-789", responseBody.GetResults()[1].GetMessageId())

	// Manual construction to match parsed objects
	baseStatus := rcs.NewMessageStatus()
	baseStatus.SetGroupId(3)
	baseStatus.SetGroupName(rcs.MESSAGEGENERALSTATUS_DELIVERED)
	baseStatus.SetId(5)
	baseStatus.SetName("DELIVERED_TO_HANDSET")
	baseStatus.SetDescription("Message delivered to handset")

	baseError := rcs.NewMessageError()
	baseError.SetGroupId(0)
	baseError.SetGroupName(rcs.MESSAGEERRORGROUP_OK)
	baseError.SetId(0)
	baseError.SetName("NO_ERROR")
	baseError.SetDescription("No error")
	baseError.SetPermanent(false)

	primary := rcs.NewWebhookDeliveryReport()
	primary.SetBulkId("bulk-123")
	primary.SetChannel("RCS")
	primary.SetMessageId("msg-456")
	primary.SetTo("441134960001")
	primary.SetSender("DemoSender")
	primary.SetSentAt(ibTime("2025-02-06T15:35:12.000+0000"))
	primary.SetDoneAt(ibTime("2025-02-06T15:35:14.000+0000"))
	primary.SetMessageCount(1)
	primary.SetStatus(*baseStatus)
	primary.SetError(*baseError)

	failover := rcs.NewWebhookDeliveryReport()
	failover.SetBulkId("bulk-123")
	failover.SetChannel("SMS")
	failover.SetMessageId("msg-789")
	failover.SetTo("441134960001")
	failover.SetSender("DemoSender")
	failover.SetSentAt(ibTime("2025-02-06T15:36:12.000+0000"))
	failover.SetDoneAt(ibTime("2025-02-06T15:36:14.000+0000"))
	failover.SetMessageCount(1)
	failover.SetStatus(*baseStatus)
	failover.SetError(*baseError)

	expected := rcs.NewWebhookDeliveryReports()
	expected.SetResults([]rcs.WebhookDeliveryReport{*primary, *failover})
	assert.EqualValues(t, expected.GetResults(), responseBody.GetResults())
}

func TestReceiveRcsSeenReportsWebhook(t *testing.T) {
	givenMessageId := "msg-456-rcs"
	givenFrom := "441134960001"
	givenTo := "DemoSender"
	givenSentAt := "2025-02-06T15:35:12.000+0000"
	givenSeenAt := "2025-02-06T15:36:00.000+0000"
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"
	givenCampaignReferenceId := "campaignRef"

	givenPayload := fmt.Sprintf(`
		{
			"results": [
				{
					"messageId": "%s",
					"from": "%s",
					"to": "%s",
					"sentAt": "%s",
					"seenAt": "%s",
					"entityId": "%s",
					"applicationId": "%s",
					"campaignReferenceId": "%s"
				}
			]
		}`,
		givenMessageId,
		givenFrom,
		givenTo,
		givenSentAt,
		givenSeenAt,
		givenEntityId,
		givenApplicationId,
		givenCampaignReferenceId,
	)

	var responseBody rcs.SeenReports
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, responseBody, "Expected non-nil response body")
	assert.Equal(t, 1, len(responseBody.GetResults()))

	result := responseBody.GetResults()[0]
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenFrom, result.GetFrom())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenEntityId, result.GetEntityId())
	assert.Equal(t, givenApplicationId, result.GetApplicationId())
	assert.Equal(t, givenCampaignReferenceId, result.GetCampaignReferenceId())

	expectedReport := rcs.NewSeenReport()
	expectedReport.SetMessageId(givenMessageId)
	expectedReport.SetFrom(givenFrom)
	expectedReport.SetTo(givenTo)
	expectedReport.SetSentAt(ibTime(givenSentAt))
	expectedReport.SetSeenAt(ibTime(givenSeenAt))
	expectedReport.SetEntityId(givenEntityId)
	expectedReport.SetApplicationId(givenApplicationId)
	expectedReport.SetCampaignReferenceId(givenCampaignReferenceId)

	expected := rcs.NewSeenReports()
	expected.SetResults([]rcs.SeenReport{*expectedReport})
	assert.EqualValues(t, expected.GetResults(), responseBody.GetResults())
}

func TestReceiveRcsTypingIndicatorWebhook(t *testing.T) {
	givenSender := "441134960001"
	givenTo := "DemoSender"
	givenReceivedAt := "2025-02-06T15:35:12.123+0000"
	givenEntityId := "my-entity-id"
	givenApplicationId := "my-application-id"

	givenPayload := fmt.Sprintf(`
		{
			"results": [
				{
					"sender": "%s",
					"to": "%s",
					"receivedAt": "%s",
					"event": {
						"type": "TYPING_INDICATOR"
					},
					"entityId": "%s",
					"applicationId": "%s"
				}
			],
			"eventCount": 1,
			"pendingEventCount": 0
		}`,
		givenSender,
		givenTo,
		givenReceivedAt,
		givenEntityId,
		givenApplicationId,
	)

	var responseBody rcs.IsTypingEvents
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, responseBody, "Expected non-nil response body")
	assert.Equal(t, int32(1), responseBody.GetEventCount())
	assert.Equal(t, int32(0), responseBody.GetPendingEventCount())
	assert.Equal(t, 1, len(responseBody.GetResults()))

	result := responseBody.GetResults()[0]
	assert.Equal(t, givenSender, result.GetSender())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenReceivedAt, result.GetReceivedAt())
	assert.Equal(t, givenEntityId, result.GetEntityId())
	assert.Equal(t, givenApplicationId, result.GetApplicationId())

	event := result.GetEvent()
	assert.Equal(t, rcs.ISTYPINGEVENTTYPE_TYPING_INDICATOR, event.GetType())

	expectedEventContent := rcs.NewIsTypingEventContent(rcs.ISTYPINGEVENTTYPE_TYPING_INDICATOR)
	expectedEvent := rcs.NewIsTypingEvent(givenSender, givenTo, givenReceivedAt, *expectedEventContent)
	expectedEvent.SetEntityId(givenEntityId)
	expectedEvent.SetApplicationId(givenApplicationId)

	expectedEvents := rcs.NewIsTypingEvents([]rcs.IsTypingEvent{*expectedEvent}, 1, 0)

	assert.EqualValues(t, expectedEvents.GetResults(), responseBody.GetResults())
	assert.Equal(t, expectedEvents.GetEventCount(), responseBody.GetEventCount())
	assert.Equal(t, expectedEvents.GetPendingEventCount(), responseBody.GetPendingEventCount())
}

func TestReceiveRcsCapabilityCheckResultWebhook(t *testing.T) {
	givenBulkId := "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8"
	givenMessageId := "a28dd97c-1ffb-4fcf-99f1-0b557ed381da"
	givenPhoneNumber := "441134960001"

	givenPayload := fmt.Sprintf(`
		{
			"bulkId": "%s",
			"messageId": "%s",
			"phoneNumber": "%s",
			"code": "ENABLED"
		}`,
		givenBulkId,
		givenMessageId,
		givenPhoneNumber,
	)

	var result rcs.CapabilityCheckResult
	err := json.Unmarshal([]byte(givenPayload), &result)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenPhoneNumber, result.GetPhoneNumber())
	assert.Equal(t, rcs.CAPABILITYCHECKRESPONSECODE_ENABLED, result.GetCode())

	expected := rcs.NewCapabilityCheckResult(givenBulkId, givenMessageId, givenPhoneNumber, rcs.CAPABILITYCHECKRESPONSECODE_ENABLED)
	assert.EqualValues(t, *expected, result)
}

func TestReceiveRcsCapabilityCheckResultWebhook_Unreachable(t *testing.T) {
	givenPayload := `
		{
			"bulkId": "d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8",
			"messageId": "a28dd97c-1ffb-4fcf-99f1-0b557ed381da",
			"phoneNumber": "441134960002",
			"code": "UNREACHABLE"
		}`

	var result rcs.CapabilityCheckResult
	err := json.Unmarshal([]byte(givenPayload), &result)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, "441134960002", result.GetPhoneNumber())
	assert.Equal(t, rcs.CAPABILITYCHECKRESPONSECODE_UNREACHABLE, result.GetCode())

	expected := rcs.NewCapabilityCheckResult("d5c3bdff-2d44-4f74-8a8e-3792fa57dfc8", "a28dd97c-1ffb-4fcf-99f1-0b557ed381da", "441134960002", rcs.CAPABILITYCHECKRESPONSECODE_UNREACHABLE)
	assert.EqualValues(t, *expected, result)
}

func TestReceiveRcsInboundMessageWebhook(t *testing.T) {
	givenSender := "441134960001"
	givenTo := "DemoSender"
	givenIntegrationType := "RCS"
	givenReceivedAt := "2025-02-06T15:35:12.000+0000"
	givenMessageId := "msg-inbound-123"
	givenText := "Hello from user"
	givenPricePerMessage := 0.0
	givenCurrency := "EUR"

	givenPayload := fmt.Sprintf(`
		{
			"results": [
				{
					"sender": "%s",
					"to": "%s",
					"integrationType": "%s",
					"receivedAt": "%s",
					"messageId": "%s",
					"message": {
						"type": "TEXT",
						"text": "%s"
					},
					"price": {
						"pricePerMessage": %.1f,
						"currency": "%s"
					}
				}
			],
			"messageCount": 1,
			"pendingMessageCount": 0
		}`,
		givenSender,
		givenTo,
		givenIntegrationType,
		givenReceivedAt,
		givenMessageId,
		givenText,
		givenPricePerMessage,
		givenCurrency,
	)

	var responseBody rcs.InboundMessages
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.NoError(t, err, "Expected successful unmarshal for inbound message webhook")
	assert.NotNil(t, responseBody.GetResults())
	assert.Equal(t, int32(1), responseBody.GetMessageCount())
	assert.Equal(t, int32(0), responseBody.GetPendingMessageCount())

	result := responseBody.GetResults()[0]
	assert.Equal(t, givenSender, result.GetSender())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenIntegrationType, result.GetIntegrationType())
	assert.Equal(t, givenMessageId, result.GetMessageId())

	price := result.GetPrice()
	assert.Equal(t, float32(givenPricePerMessage), price.GetPricePerMessage())
	assert.Equal(t, givenCurrency, price.GetCurrency())

	// Validate discriminator selected text content
	messageContent := result.Message
	textContent := messageContent.InboundTextContent
	if assert.NotNil(t, textContent, "Expected TEXT content to be selected") {
		assert.Equal(t, rcs.INBOUNDMESSAGECONTENTTYPE_TEXT, textContent.Type)
		assert.Equal(t, givenText, textContent.GetText())
	}

	expectedPrice := rcs.NewMessagePrice()
	expectedPrice.SetPricePerMessage(float32(givenPricePerMessage))
	expectedPrice.SetCurrency(givenCurrency)

	expectedText := rcs.NewInboundTextContent(givenText)
	expectedMessage := rcs.NewInboundMessage(
		givenSender,
		givenTo,
		givenIntegrationType,
		ibTime(givenReceivedAt),
		givenMessageId,
		rcs.InboundTextContentAsInboundMessageContent(expectedText),
		*expectedPrice,
	)

	expected := rcs.NewInboundMessages([]rcs.InboundMessage{*expectedMessage}, 1, 0)
	assert.EqualValues(t, expected.GetResults(), responseBody.GetResults())
}

func TestReceiveRcsInboundMessageWebhook_FileAndSuggestionContent(t *testing.T) {
	fileUrl := "https://www.example.com/voice-note.amr"
	fileName := "voice-note.amr"
	fileContentType := "audio/AMR"
	fileSize := int64(2048)

	suggestionText := "Yes"
	suggestionPostback := "CONFIRM_YES"

	givenPayload := fmt.Sprintf(`
		{
			"results": [
				{
					"sender": "441134960001",
					"to": "DemoSender",
					"integrationType": "RCS",
					"receivedAt": "2025-02-06T15:35:12.000+0000",
					"messageId": "msg-file-123",
					"message": {
						"type": "FILE",
						"url": "%s",
						"name": "%s",
						"contentType": "%s",
						"size": %d
					},
					"price": {
						"pricePerMessage": 0.0,
						"currency": "EUR"
					}
				},
				{
					"sender": "441134960002",
					"to": "DemoSender",
					"integrationType": "RCS",
					"receivedAt": "2025-02-06T15:36:12.000+0000",
					"messageId": "msg-suggestion-456",
					"message": {
						"type": "SUGGESTION",
						"text": "%s",
						"postbackData": "%s"
					},
					"price": {
						"pricePerMessage": 0.0,
						"currency": "EUR"
					}
				}
			],
			"messageCount": 2,
			"pendingMessageCount": 0
		}`,
		fileUrl,
		fileName,
		fileContentType,
		fileSize,
		suggestionText,
		suggestionPostback,
	)

	var responseBody rcs.InboundMessages
	err := json.Unmarshal([]byte(givenPayload), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, int32(2), responseBody.GetMessageCount())

	fileResult := responseBody.GetResults()[0]
	fileContent := fileResult.Message.InboundFileContent
	if assert.NotNil(t, fileContent) {
		assert.Equal(t, rcs.INBOUNDMESSAGECONTENTTYPE_FILE, fileContent.Type)
		assert.Equal(t, fileUrl, fileContent.GetUrl())
		assert.Equal(t, fileName, fileContent.GetName())
		assert.Equal(t, fileContentType, fileContent.GetContentType())
		assert.Equal(t, fileSize, fileContent.GetSize())
	}

	suggestionResult := responseBody.GetResults()[1]
	suggestionContent := suggestionResult.Message.InboundSuggestionContent
	if assert.NotNil(t, suggestionContent) {
		assert.Equal(t, rcs.INBOUNDMESSAGECONTENTTYPE_SUGGESTION, suggestionContent.Type)
		assert.Equal(t, suggestionText, suggestionContent.GetText())
		assert.Equal(t, suggestionPostback, suggestionContent.GetPostbackData())
	}

	// Manual instantiation to validate conversion
	priceZero := rcs.NewMessagePrice()
	priceZero.SetPricePerMessage(0.0)
	priceZero.SetCurrency("EUR")

	fileContentModel := rcs.NewInboundFileContent(fileUrl, fileName, fileContentType, fileSize)
	fileMessage := rcs.NewInboundMessage(
		"441134960001",
		"DemoSender",
		"RCS",
		ibTime("2025-02-06T15:35:12.000+0000"),
		"msg-file-123",
		rcs.InboundFileContentAsInboundMessageContent(fileContentModel),
		*priceZero,
	)

	suggestionContentModel := rcs.NewInboundSuggestionContent(suggestionText, suggestionPostback)
	suggestionMessage := rcs.NewInboundMessage(
		"441134960002",
		"DemoSender",
		"RCS",
		ibTime("2025-02-06T15:36:12.000+0000"),
		"msg-suggestion-456",
		rcs.InboundSuggestionContentAsInboundMessageContent(suggestionContentModel),
		*priceZero,
	)

	expected := rcs.NewInboundMessages(
		[]rcs.InboundMessage{*fileMessage, *suggestionMessage},
		2,
		0,
	)
	assert.EqualValues(t, expected.GetResults(), responseBody.GetResults())
}

func TestReceiveRcsInboundEventWebhook(t *testing.T) {
	givenPayload := `
			{
				"results": [
					{
						"sender": "441134960001",
						"to": "DemoSender",
						"integrationType": "RCS",
						"receivedAt": "2025-02-06T15:35:12.000+0000",
						"messageId": "msg-event-123",
						"message": {
							"type": "SUGGESTION",
							"text": "Yes",
							"postbackData": "pb-yes"
						},
						"price": {
							"pricePerMessage": 0.0,
							"currency": "EUR"
						}
					}
				],
				"messageCount": 1,
				"pendingMessageCount": 0
			}`

	var responseBody rcs.InboundEvents
	err := json.Unmarshal([]byte(givenPayload), &responseBody)

	assert.NoError(t, err)
	if assert.Equal(t, 1, len(responseBody.GetResults())) {
		event := responseBody.GetResults()[0]
		assert.Equal(t, "441134960001", event.GetSender())
		assert.Equal(t, "DemoSender", event.GetTo())
		assert.Equal(t, "RCS", event.GetIntegrationType())
		assert.Equal(t, ibTime("2025-02-06T15:35:12.000+0000"), event.GetReceivedAt())
		assert.Equal(t, "msg-event-123", event.GetMessageId())

		price := event.GetPrice()
		assert.InDelta(t, 0.0, price.GetPricePerMessage(), 0.0001)
		assert.Equal(t, "EUR", price.GetCurrency())

		msg := event.GetMessage()
		// The discriminator should map to suggestion content.
		if assert.NotNil(t, msg.InboundSuggestionContent) {
			assert.Equal(t, "SUGGESTION", string(msg.InboundSuggestionContent.Type))
			assert.Equal(t, "Yes", msg.InboundSuggestionContent.GetText())
			assert.Equal(t, "pb-yes", msg.InboundSuggestionContent.GetPostbackData())
		}
	}
}

// Helper function for string pointer
func ptr(s string) *string {
	return &s
}

// ibTime converts a string in INFOBIP_TIME_FORMAT to infobip.Time for manual model building.
func ibTime(value string) infobip.Time {
	t, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, value)
	return infobip.Time{t}
}
