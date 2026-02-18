package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	ClickToCallSendEndpoint string = "/voice/ctc/1/send"
)

func TestSendClickToCallMessage(t *testing.T) {
	givenBulkId := "bulk-12345"
	givenMessageId := "msg-12345"
	givenFrom := "441134960000"
	givenTo := "441134960001"
	givenAudioUrl := "http://example.com/audio.wav"

	givenResponse := fmt.Sprintf(`
		{
			"bulkId": "%s",
			"messages": [
				{
					"messageId": "%s",
					"to": "%s",
					"status": {
						"groupId": 1,
						"groupName": "PENDING",
						"id": 26,
						"name": "PENDING_ACCEPTED",
						"description": "Message sent to next instance"
					}
				}
			]
		}`,
		givenBulkId,
		givenMessageId,
		givenTo,
	)

	givenRequest := fmt.Sprintf(`
		{
			"messages": [
				{
					"from": "%s",
					"destinationA": "%s",
					"destinationB": "%s",
					"audioFileUrl": "%s"
				}
			]
		}`,
		givenFrom,
		givenTo,
		givenTo,
		givenAudioUrl,
	)

	// Prepare the request
	message := voice.NewClickToCallMessage(givenTo, givenTo, givenFrom)
	message.SetAudioFileUrl(givenAudioUrl)
	request := voice.NewClickToCallMessageRequest([]voice.ClickToCallMessage{*message})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", ClickToCallSendEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		ClickToCallAPI.
		SendClickToCallMessage(context.Background()).
		ClickToCallMessageRequest(*request).Execute()

	// Ensure r and responseBody are not nil before dereferencing
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenBulkId, responseBody.GetBulkId())
	assert.True(t, len(responseBody.GetMessages()) == 1)

	msg := responseBody.GetMessages()[0]
	assert.Equal(t, givenMessageId, msg.GetMessageId())
	assert.Equal(t, givenTo, msg.GetTo())

	status := msg.GetStatus()
	assert.Equal(t, int32(1), status.GetGroupId())
	assert.Equal(t, "PENDING", string(status.GetGroupName()))
	assert.Equal(t, int32(26), status.GetId())
	assert.Equal(t, "PENDING_ACCEPTED", status.GetName())
	assert.Equal(t, "Message sent to next instance", status.GetDescription())
}
