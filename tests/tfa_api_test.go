package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	TfaApplicationsEndpoint       string = "/2fa/2/applications"
	TfaApplicationEndpoint        string = "/2fa/2/applications/%s"
	TfaMessageTemplatesEndpoint   string = "/2fa/2/applications/%s/messages"
	TfaMessageTemplateEndpoint    string = "/2fa/2/applications/%s/messages/%s"
	TfaEmailTemplatesEndpoint     string = "/2fa/2/applications/%s/email/messages"
	TfaEmailTemplateEndpoint      string = "/2fa/2/applications/%s/email/messages/%s"
	TfaSendPinOverSmsEndpoint     string = "/2fa/2/pin"
	TfaResendPinOverSmsEndpoint   string = "/2fa/2/pin/%s/resend"
	TfaSendPinOverVoiceEndpoint   string = "/2fa/2/pin/voice"
	TfaResendPinOverVoiceEndpoint string = "/2fa/2/pin/%s/resend/voice"
	TfaSendPinOverEmailEndpoint   string = "/2fa/2/pin/email"
	TfaResendPinOverEmailEndpoint string = "/2fa/2/pin/%s/resend/email"
	TfaVerifyPhoneNumberEndpoint  string = "/2fa/2/pin/%s/verify"
	TfaVerificationStatusEndpoint string = "/2fa/2/applications/%s/verifications"
)

func TestCreateTfaApplication(t *testing.T) {
	givenApplicationId := "1234567"
	givenRequestApplicationName := "2fa application name"
	givenResponseApplicationName := "Application name"
	givenEnabled := true
	givenPinAttempts := int32(5)
	givenAllowMultiplePinVerifications := true
	givenPinTimeToLive := "10m"
	givenVerifyPinLimit := "2/4s"
	givenSendPinPerApplicationLimit := "5000/12h"
	givenSendPinPerPhoneNumberLimit := "2/1d"

	givenResponse := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"name": "%s",
			"configuration": {
				"pinAttempts": %d,
				"allowMultiplePinVerifications": %t,
				"pinTimeToLive": "%s",
				"verifyPinLimit": "%s",
				"sendPinPerApplicationLimit": "%s",
				"sendPinPerPhoneNumberLimit": "%s"
			},
			"enabled": %t
		}`,
		givenApplicationId,
		givenResponseApplicationName,
		givenPinAttempts,
		givenAllowMultiplePinVerifications,
		givenPinTimeToLive,
		givenVerifyPinLimit,
		givenSendPinPerApplicationLimit,
		givenSendPinPerPhoneNumberLimit,
		givenEnabled,
	)

	givenRequest := fmt.Sprintf(`
		{
			"name": "%s",
			"enabled": %t,
			"configuration": {
				"pinAttempts": %d,
				"allowMultiplePinVerifications": %t,
				"pinTimeToLive": "%s",
				"verifyPinLimit": "%s",
				"sendPinPerApplicationLimit": "%s",
				"sendPinPerPhoneNumberLimit": "%s"
			}
		}`,
		givenRequestApplicationName,
		givenEnabled,
		givenPinAttempts,
		givenAllowMultiplePinVerifications,
		givenPinTimeToLive,
		givenVerifyPinLimit,
		givenSendPinPerApplicationLimit,
		givenSendPinPerPhoneNumberLimit,
	)

	// Prepare the request
	request := sms.NewApplicationRequest(givenRequestApplicationName)
	request.SetEnabled(givenEnabled)
	configuration := sms.NewApplicationConfiguration()
	configuration.SetPinAttempts(givenPinAttempts)
	configuration.SetAllowMultiplePinVerifications(givenAllowMultiplePinVerifications)
	configuration.SetPinTimeToLive(givenPinTimeToLive)
	configuration.SetVerifyPinLimit(givenVerifyPinLimit)
	configuration.SetSendPinPerApplicationLimit(givenSendPinPerApplicationLimit)
	configuration.SetSendPinPerPhoneNumberLimit(givenSendPinPerPhoneNumberLimit)
	request.SetConfiguration(*configuration)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", TfaApplicationsEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		CreateTfaApplication(context.Background()).
		ApplicationRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenResponseApplicationName, responseBody.GetName())
	assert.Equal(t, givenEnabled, responseBody.GetEnabled())
}

func TestGetTfaApplication(t *testing.T) {
	givenApplicationId := "1234567"
	givenApplicationName := "Application name"

	givenResponse := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"name": "%s",
			"configuration": {
				"pinAttempts": 5,
				"allowMultiplePinVerifications": true,
				"pinTimeToLive": "10m",
				"verifyPinLimit": "2/4s",
				"sendPinPerApplicationLimit": "5000/12h",
				"sendPinPerPhoneNumberLimit": "2/1d"
			},
			"enabled": true
		}`,
		givenApplicationId,
		givenApplicationName,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaApplicationEndpoint, givenApplicationId)
	SetUpSuccessRequest("GET", endpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		GetTfaApplication(context.Background(), givenApplicationId).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenApplicationName, responseBody.GetName())
	assert.True(t, responseBody.GetEnabled())
}

func TestGetTfaApplications(t *testing.T) {
	givenApplicationId1 := "0933F3BC087D2A617AC6DCB2EF5B8A61"
	givenApplicationName1 := "Test application BASIC 1"
	givenApplicationId2 := "5F04FACFAA4978F62FCAEBA97B37E90F"
	givenApplicationName2 := "Test application BASIC 2"
	givenApplicationId3 := "B450F966A8EF017180F148AF22C42642"
	givenApplicationName3 := "Test application BASIC 3"

	givenResponse := fmt.Sprintf(`
		[
			{
				"applicationId": "%s",
				"name": "%s",
				"configuration": {
					"pinAttempts": 10,
					"allowMultiplePinVerifications": true,
					"pinTimeToLive": "2h",
					"verifyPinLimit": "1/3s",
					"sendPinPerApplicationLimit": "10000/1d",
					"sendPinPerPhoneNumberLimit": "3/1d"
				},
				"enabled": true
			},
			{
				"applicationId": "%s",
				"name": "%s",
				"configuration": {
					"pinAttempts": 12,
					"allowMultiplePinVerifications": true,
					"pinTimeToLive": "10m",
					"verifyPinLimit": "2/1s",
					"sendPinPerApplicationLimit": "10000/1d",
					"sendPinPerPhoneNumberLimit": "5/1h"
				},
				"enabled": true
			},
			{
				"applicationId": "%s",
				"name": "%s",
				"configuration": {
					"pinAttempts": 15,
					"allowMultiplePinVerifications": true,
					"pinTimeToLive": "1h",
					"verifyPinLimit": "30/10s",
					"sendPinPerApplicationLimit": "10000/3d",
					"sendPinPerPhoneNumberLimit": "10/20m"
				},
				"enabled": true
			}
		]`,
		givenApplicationId1,
		givenApplicationName1,
		givenApplicationId2,
		givenApplicationName2,
		givenApplicationId3,
		givenApplicationName3,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("GET", TfaApplicationsEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		GetTfaApplications(context.Background()).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")
	assert.Len(t, responseBody, 3)

	app1 := responseBody[0]
	assert.Equal(t, givenApplicationId1, app1.GetApplicationId())
	assert.Equal(t, givenApplicationName1, app1.GetName())
	assert.True(t, app1.GetEnabled())

	app2 := responseBody[1]
	assert.Equal(t, givenApplicationId2, app2.GetApplicationId())
	assert.Equal(t, givenApplicationName2, app2.GetName())
	assert.True(t, app2.GetEnabled())

	app3 := responseBody[2]
	assert.Equal(t, givenApplicationId3, app3.GetApplicationId())
	assert.Equal(t, givenApplicationName3, app3.GetName())
	assert.True(t, app3.GetEnabled())
}

func TestUpdateTfaApplication(t *testing.T) {
	givenApplicationId := "1234567"
	givenRequestApplicationName := "2fa application name"
	givenResponseApplicationName := "Application name"
	givenEnabled := true
	givenPinAttempts := int32(5)
	givenAllowMultiplePinVerifications := true
	givenPinTimeToLive := "10m"
	givenVerifyPinLimit := "2/4s"
	givenSendPinPerApplicationLimit := "5000/12h"
	givenSendPinPerPhoneNumberLimit := "2/1d"

	givenResponse := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"name": "%s",
			"configuration": {
				"pinAttempts": %d,
				"allowMultiplePinVerifications": true,
				"pinTimeToLive": "%s",
				"verifyPinLimit": "%s",
				"sendPinPerApplicationLimit": "%s",
				"sendPinPerPhoneNumberLimit": "%s"
			},
			"enabled": %t
		}`,
		givenApplicationId,
		givenResponseApplicationName,
		givenPinAttempts,
		givenPinTimeToLive,
		givenVerifyPinLimit,
		givenSendPinPerApplicationLimit,
		givenSendPinPerPhoneNumberLimit,
		givenEnabled,
	)

	givenRequest := fmt.Sprintf(`
		{
			"name": "%s",
			"enabled": %t,
			"configuration": {
				"pinAttempts": %d,
				"allowMultiplePinVerifications": %t,
				"pinTimeToLive": "%s",
				"verifyPinLimit": "%s",
				"sendPinPerApplicationLimit": "%s",
				"sendPinPerPhoneNumberLimit": "%s"
			}
		}`,
		givenRequestApplicationName,
		givenEnabled,
		givenPinAttempts,
		givenAllowMultiplePinVerifications,
		givenPinTimeToLive,
		givenVerifyPinLimit,
		givenSendPinPerApplicationLimit,
		givenSendPinPerPhoneNumberLimit,
	)

	// Prepare the request
	request := sms.NewApplicationRequest(givenRequestApplicationName)
	request.SetEnabled(givenEnabled)
	configuration := sms.NewApplicationConfiguration()
	configuration.SetPinAttempts(givenPinAttempts)
	configuration.SetAllowMultiplePinVerifications(givenAllowMultiplePinVerifications)
	configuration.SetPinTimeToLive(givenPinTimeToLive)
	configuration.SetVerifyPinLimit(givenVerifyPinLimit)
	configuration.SetSendPinPerApplicationLimit(givenSendPinPerApplicationLimit)
	configuration.SetSendPinPerPhoneNumberLimit(givenSendPinPerPhoneNumberLimit)
	request.SetConfiguration(*configuration)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaApplicationEndpoint, givenApplicationId)
	SetUpSuccessRequest("PUT", endpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		UpdateTfaApplication(context.Background(), givenApplicationId).
		ApplicationRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenResponseApplicationName, responseBody.GetName())
	assert.Equal(t, givenEnabled, responseBody.GetEnabled())
}

func TestCreateTfaMessageTemplate(t *testing.T) {
	givenApplicationId := "19F2491982980DBC2C40E1597FD175AB"
	givenMessageId := "B628A31410306DDB2BFB0D9CF1F3A3FF"
	givenRequestPinType := sms.PINTYPE_NUMERIC
	givenResponsePinType := sms.PINTYPE_ALPHANUMERIC
	givenPinLength := int32(4)
	givenMessageText := "Your pin is {{pin}}"
	givenLanguage := sms.TFALANGUAGE_EN
	givenVoiceName := "Joanna"
	givenSenderId := "Infobip 2FA"
	givenRepeatDTMF := "1#"
	givenSpeechRate := float64(1.0)

	givenResponse := fmt.Sprintf(`
		{
			"pinType": "%s",
			"messageId": "%s",
			"applicationId": "%s",
			"pinLength": %d,
			"messageText": "%s",
			"pinPlaceholder": "{{pin}}",
			"senderId": "%s",
			"language": "%s",
			"voiceName": "%s",
			"repeatDTMF": "%s",
			"speechRate": %.1f
		}`,
		givenResponsePinType,
		givenMessageId,
		givenApplicationId,
		givenPinLength,
		givenMessageText,
		givenSenderId,
		givenLanguage,
		givenVoiceName,
		givenRepeatDTMF,
		givenSpeechRate,
	)

	givenRequest := fmt.Sprintf(`
		{
			"pinType": "%s",
			"pinLength": %d,
			"messageText": "%s",
			"language": "%s",
			"voiceName": "%s",
			"senderId": "%s",
			"repeatDTMF": "%s",
			"speechRate": %.1f
		}`,
		givenRequestPinType,
		givenPinLength,
		givenMessageText,
		givenLanguage,
		givenVoiceName,
		givenSenderId,
		givenRepeatDTMF,
		givenSpeechRate,
	)

	// Prepare the request
	request := sms.NewCreateMessageRequest(givenMessageText, givenRequestPinType)
	request.SetPinLength(givenPinLength)
	request.SetLanguage(givenLanguage)
	request.SetVoiceName(givenVoiceName)
	request.SetSenderId(givenSenderId)
	request.SetRepeatDTMF(givenRepeatDTMF)
	request.SetSpeechRate(givenSpeechRate)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaMessageTemplatesEndpoint, givenApplicationId)
	SetUpSuccessRequest("POST", endpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		CreateTfaMessageTemplate(context.Background(), givenApplicationId).
		CreateMessageRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenMessageId, responseBody.GetMessageId())
	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenResponsePinType, responseBody.GetPinType())
	assert.Equal(t, givenPinLength, responseBody.GetPinLength())
	assert.Equal(t, givenMessageText, responseBody.GetMessageText())
}

func TestGetTfaMessageTemplate(t *testing.T) {
	givenApplicationId := "32G5F37A635KJ8BHJ675435E3A6EA434"
	givenMessageId := "A89CE542F3F12"
	givenPinType := sms.PINTYPE_ALPHANUMERIC
	givenMessageText := "Your pin is {{pin}}"

	givenResponse := fmt.Sprintf(`
		{
			"pinType": "%s",
			"messageId": "%s",
			"applicationId": "%s",
			"pinLength": 4,
			"messageText": "%s",
			"pinPlaceholder": "{{pin}}",
			"senderId": "Infobip 2FA"
		}`,
		givenPinType,
		givenMessageId,
		givenApplicationId,
		givenMessageText,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaMessageTemplateEndpoint, givenApplicationId, givenMessageId)
	SetUpSuccessRequest("GET", endpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		GetTfaMessageTemplate(context.Background(), givenApplicationId, givenMessageId).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenMessageId, responseBody.GetMessageId())
	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
}

func TestCreateTfaEmailMessageTemplate(t *testing.T) {
	givenApplicationId := "HJ675435E3A6EA43432G5F37A635KJ8B"
	givenMessageId := "9C815F8AF3328"
	givenEmailTemplateId := int64(1234)
	givenFrom := "company@example.com"
	givenLandingPageId := "1_23456"
	givenPinLength := int32(4)
	givenRequestPinType := sms.PINTYPE_NUMERIC
	givenResponsePinType := sms.PINTYPE_ALPHANUMERIC

	givenResponse := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"emailTemplateId": %d,
			"from": "%s",
			"messageId": "%s",
			"pinLength": %d,
			"pinType": "%s",
			"landingPageId": "%s"
		}`,
		givenApplicationId,
		givenEmailTemplateId,
		givenFrom,
		givenMessageId,
		givenPinLength,
		givenResponsePinType,
		givenLandingPageId,
	)

	givenRequest := fmt.Sprintf(`
		{
			"emailTemplateId": %d,
			"from": "%s",
			"pinLength": %d,
			"pinType": "%s",
			"landingPageId": "%s"
		}`,
		givenEmailTemplateId,
		givenFrom,
		givenPinLength,
		givenRequestPinType,
		givenLandingPageId,
	)

	request := sms.NewCreateEmailMessageRequest(givenEmailTemplateId)
	request.SetFrom(givenFrom)
	request.SetPinLength(givenPinLength)
	request.SetPinType(givenRequestPinType)
	request.SetLandingPageId(givenLandingPageId)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaEmailTemplatesEndpoint, givenApplicationId)
	SetUpSuccessRequest("POST", endpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		TfaAPI.
		CreateTfaEmailMessageTemplate(context.Background(), givenApplicationId).
		CreateEmailMessageRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenMessageId, responseBody.GetMessageId())
	assert.Equal(t, givenEmailTemplateId, responseBody.GetEmailTemplateId())
	assert.Equal(t, givenFrom, responseBody.GetFrom())
	assert.Equal(t, givenPinLength, responseBody.GetPinLength())
	assert.Equal(t, givenResponsePinType, responseBody.GetPinType())
	assert.Equal(t, givenLandingPageId, responseBody.GetLandingPageId())
}

func TestUpdateTfaEmailMessageTemplate(t *testing.T) {
	givenApplicationId := "HJ675435E3A6EA43432G5F37A635KJ8B"
	givenMessageId := "9C815F8AF3328"
	givenEmailTemplateId := int64(1234)
	givenFrom := "company@example.com"
	givenLandingPageId := "1_23456"
	givenPinLength := int32(4)
	givenRequestPinType := sms.PINTYPE_NUMERIC
	givenResponsePinType := sms.PINTYPE_ALPHANUMERIC

	givenResponse := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"emailTemplateId": %d,
			"from": "%s",
			"messageId": "%s",
			"pinLength": %d,
			"pinType": "%s",
			"landingPageId": "%s"
		}`,
		givenApplicationId,
		givenEmailTemplateId,
		givenFrom,
		givenMessageId,
		givenPinLength,
		givenResponsePinType,
		givenLandingPageId,
	)

	givenRequest := fmt.Sprintf(`
		{
			"emailTemplateId": %d,
			"from": "%s",
			"pinLength": %d,
			"pinType": "%s",
			"landingPageId": "%s"
		}`,
		givenEmailTemplateId,
		givenFrom,
		givenPinLength,
		givenRequestPinType,
		givenLandingPageId,
	)

	request := sms.NewUpdateEmailMessageRequest()
	request.SetEmailTemplateId(givenEmailTemplateId)
	request.SetFrom(givenFrom)
	request.SetPinLength(givenPinLength)
	request.SetPinType(givenRequestPinType)
	request.SetLandingPageId(givenLandingPageId)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaEmailTemplateEndpoint, givenApplicationId, givenMessageId)
	SetUpSuccessRequest("PUT", endpoint, givenResponse, 200)

	responseBody, r, err := infobipClient.
		TfaAPI.
		UpdateTfaEmailMessageTemplate(context.Background(), givenApplicationId, givenMessageId).
		UpdateEmailMessageRequest(*request).
		Execute()

	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenApplicationId, responseBody.GetApplicationId())
	assert.Equal(t, givenMessageId, responseBody.GetMessageId())
	assert.Equal(t, givenEmailTemplateId, responseBody.GetEmailTemplateId())
	assert.Equal(t, givenFrom, responseBody.GetFrom())
	assert.Equal(t, givenPinLength, responseBody.GetPinLength())
	assert.Equal(t, givenResponsePinType, responseBody.GetPinType())
	assert.Equal(t, givenLandingPageId, responseBody.GetLandingPageId())
}

func TestSendTfaPinCodeOverSms(t *testing.T) {
	givenApplicationId := "1234567"
	givenMessageId := "7654321"
	givenTo := "41793026727"
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"
	givenFrom := "Sender 1"
	givenFirstName := "John"

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"to": "%s",
			"ncStatus": "NC_DESTINATION_REACHABLE",
			"smsStatus": "MESSAGE_SENT"
		}`,
		givenPinId,
		givenTo,
	)

	givenRequest := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"messageId": "%s",
			"to": "%s",
			"trackDelivery": false,
			"from": "%s",
			"placeholders": {
				"firstName": "%s"
			}
		}`,
		givenApplicationId,
		givenMessageId,
		givenTo,
		givenFrom,
		givenFirstName,
	)

	// Prepare the request
	request := sms.NewStartAuthenticationRequest(givenApplicationId, givenMessageId, givenTo)
	request.SetFrom(givenFrom)
	request.SetPlaceholders(map[string]string{"firstName": givenFirstName})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", TfaSendPinOverSmsEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		SendTfaPinCodeOverSms(context.Background()).
		StartAuthenticationRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenPinId, responseBody.GetPinId())
	assert.Equal(t, givenTo, responseBody.GetTo())
}

func TestResendTfaPinCodeOverSms(t *testing.T) {
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"
	givenTo := "41793026727"
	givenFirstName := "John"

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"to": "%s",
			"ncStatus": "NC_DESTINATION_REACHABLE",
			"smsStatus": "MESSAGE_SENT"
		}`,
		givenPinId,
		givenTo,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaResendPinOverSmsEndpoint, givenPinId)
	SetUpSuccessRequest("POST", endpoint, givenResponse, 200)

	// Build request body
	resendRequest := sms.NewResendPinRequest()
	resendRequest.SetPlaceholders(map[string]string{"firstName": givenFirstName})

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		ResendTfaPinCodeOverSms(context.Background(), givenPinId).
		ResendPinRequest(*resendRequest).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenPinId, responseBody.GetPinId())
	assert.Equal(t, givenTo, responseBody.GetTo())
}

func TestSendTfaPinCodeOverVoice(t *testing.T) {
	givenApplicationId := "1234567"
	givenMessageId := "7654321"
	givenTo := "41793026727"
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"
	givenFrom := "Sender 1"
	givenFirstName := "John"

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"to": "%s",
			"callStatus": "PENDING_ACCEPTED"
		}`,
		givenPinId,
		givenTo,
	)

	givenRequest := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"messageId": "%s",
			"to": "%s",
			"trackDelivery": false,
			"from": "%s",
			"placeholders": {
				"firstName": "%s"
			}
		}`,
		givenApplicationId,
		givenMessageId,
		givenTo,
		givenFrom,
		givenFirstName,
	)

	// Prepare the request
	request := sms.NewStartAuthenticationRequest(givenApplicationId, givenMessageId, givenTo)
	request.SetFrom(givenFrom)
	request.SetPlaceholders(map[string]string{"firstName": givenFirstName})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", TfaSendPinOverVoiceEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		SendTfaPinCodeOverVoice(context.Background()).
		StartAuthenticationRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenPinId, responseBody.GetPinId())
	assert.Equal(t, givenTo, responseBody.GetTo())
}

func TestResendTfaPinCodeOverVoice(t *testing.T) {
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"
	givenTo := "41793026727"
	givenFirstName := "John"

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"to": "%s",
			"callStatus": "PENDING_ACCEPTED"
		}`,
		givenPinId,
		givenTo,
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaResendPinOverVoiceEndpoint, givenPinId)
	SetUpSuccessRequest("POST", endpoint, givenResponse, 200)

	resendRequest := sms.NewResendPinRequest()
	resendRequest.SetPlaceholders(map[string]string{"firstName": givenFirstName})

	responseBody, r, err := infobipClient.
		TfaAPI.
		ResendTfaPinCodeOverVoice(context.Background(), givenPinId).
		ResendPinRequest(*resendRequest).
		Execute()

	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenPinId, responseBody.GetPinId())
	assert.Equal(t, givenTo, responseBody.GetTo())
}

func TestVerifyTfaPhoneNumber(t *testing.T) {
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"
	givenPin := "1598"
	givenVerified := true
	givenAttemptsRemaining := 0

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"msisdn": "41793026727",
			"verified": %t,
			"attemptsRemaining": %d
		}`,
		givenPinId,
		givenVerified,
		givenAttemptsRemaining,
	)

	givenRequest := fmt.Sprintf(`
		{
			"pin": "%s"
		}`,
		givenPin,
	)

	// Prepare the request
	request := sms.NewVerifyPinRequest(givenPin)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaVerifyPhoneNumberEndpoint, givenPinId)
	SetUpSuccessRequest("POST", endpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		VerifyTfaPhoneNumber(context.Background(), givenPinId).
		VerifyPinRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenVerified, responseBody.GetVerified())
	assert.Equal(t, int32(givenAttemptsRemaining), responseBody.GetAttemptsRemaining())
}

func TestGetTfaVerificationStatus(t *testing.T) {
	givenApplicationId := "16A8B5FE2BCD6CA716A2D780CB3F3390"
	givenMsisdn := "41793026727"
	givenSentAtFirst := int64(1418364246)
	givenVerifiedAtFirst := int64(1418364366)

	givenResponse := fmt.Sprintf(`
		{
			"verifications": [
				{
					"msisdn": "%s",
					"verified": true,
					"verifiedAt": %d,
					"sentAt": %d
				},
				{
					"msisdn": "41793026746",
					"verified": false,
					"verifiedAt": 1418364226,
					"sentAt": 1418333246
				}
			]
		}`,
		givenMsisdn,
		givenVerifiedAtFirst,
		givenSentAtFirst,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaVerificationStatusEndpoint, givenApplicationId)
	SetUpSuccessRequest("GET", endpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		GetTfaVerificationStatus(context.Background(), givenApplicationId).
		Msisdn(givenMsisdn).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	verifications := responseBody.GetVerifications()
	assert.True(t, len(verifications) == 2)

	verification := verifications[0]
	assert.Equal(t, givenMsisdn, verification.GetMsisdn())
	assert.True(t, verification.GetVerified())
	assert.Equal(t, givenVerifiedAtFirst, verification.GetVerifiedAt())
	secondVerification := verifications[1]
	assert.False(t, secondVerification.GetVerified())
}

func TestSend2faPinCodeOverEmail(t *testing.T) {
	givenApplicationId := "1234567"
	givenMessageId := "7654321"
	givenTo := "john.smith@example.com"
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"to": "%s",
			"emailStatus": {
				"name": "PENDING_ACCEPTED",
				"description": "Message accepted, pending for delivery."
			}
		}`,
		givenPinId,
		givenTo,
	)

	givenRequest := fmt.Sprintf(`
		{
			"applicationId": "%s",
			"messageId": "%s",
			"to": "%s",
			"placeholders": {
				"firstName": "John"
			},
			"landingPagePlaceholders": {
				"name": "John",
				"surname": "Smith"
			}
		}`,
		givenApplicationId,
		givenMessageId,
		givenTo,
	)

	// Prepare the request
	request := sms.NewStartEmailAuthenticationRequest(givenApplicationId, givenMessageId, givenTo)
	request.SetPlaceholders(map[string]string{"firstName": "John"})
	request.SetLandingPagePlaceholders(map[string]string{"name": "John", "surname": "Smith"})

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpSuccessRequest("POST", TfaSendPinOverEmailEndpoint, givenResponse, 200)

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		Send2faPinCodeOverEmail(context.Background()).
		StartEmailAuthenticationRequest(*request).Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	responseHeaders := *r
	ValidateExpectedRequestHeadersMatches(t, responseHeaders.Header)

	assert.Equal(t, givenPinId, responseBody.GetPinId())
	assert.Equal(t, givenTo, responseBody.GetTo())
}

func TestResend2faPinCodeOverEmail(t *testing.T) {
	givenPinId := "9C817C6F8AF3D48F9FE553282AFA2B67"
	givenTo := "john.smith@example.com"

	givenResponse := fmt.Sprintf(`
		{
			"pinId": "%s",
			"to": "%s",
			"emailStatus": {
				"name": "PENDING_ACCEPTED",
				"description": "Message accepted, pending for delivery."
			}
		}`,
		givenPinId,
		givenTo,
	)

	// Mocking HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := fmt.Sprintf(TfaResendPinOverEmailEndpoint, givenPinId)
	SetUpSuccessRequest("POST", endpoint, givenResponse, 200)

	// Build request body
	resendRequest := sms.NewResendPinRequestViaEmail()

	// Execute the request
	responseBody, r, err := infobipClient.
		TfaAPI.
		Resend2faPinCodeOverEmail(context.Background(), givenPinId).
		ResendPinRequestViaEmail(*resendRequest).
		Execute()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, r, "Expected non-nil response")
	assert.NotNil(t, responseBody, "Expected non-nil response body")

	assert.Equal(t, givenPinId, responseBody.GetPinId())
	assert.Equal(t, givenTo, responseBody.GetTo())
}
