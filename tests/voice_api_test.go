package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	VoiceTtsBulksEndpoint        = "/tts/3/bulks"
	VoiceTtsBulksStatusEndpoint  = "/tts/3/bulks/status"
	VoiceTtsLogsEndpoint         = "/tts/3/logs"
	VoiceTtsMultiEndpoint        = "/tts/3/multi"
	VoiceTtsSingleEndpoint       = "/tts/3/single"
	VoiceTtsVoiceByIdEndpoint    = "/tts/3/voices/%s"
	VoiceReportsEndpoint         = "/voice/1/reports"
	VoiceIvrFilesEndpoint        = "/voice/ivr/1/files"
	VoiceIvrUploadsEndpoint      = "/voice/ivr/1/uploads"
	VoiceIvrScenariosEndpoint    = "/voice/ivr/1/scenarios"
	VoiceIvrScenarioByIdEndpoint = "/voice/ivr/1/scenarios/%s"
)

func TestShouldGetSentBulks(t *testing.T) {
	givenBulkId := "string"
	givenSendAt := "2015-02-22T17:42:05.390+0100"
	givenSendAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAt)
	ibTimeSendAt := infobip.Time{
		T: givenSendAtDateTime,
	}

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "sendAt": "%s"
    }`, givenBulkId, givenSendAt)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", VoiceTtsBulksEndpoint, givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.GetSentBulks(context.Background()).BulkId(givenBulkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeSendAt, response.GetSendAt())
}

func TestShouldRescheduleSentBulk(t *testing.T) {
	givenBulkId := "123"
	givenSendAt := "2015-02-22T17:42:05.390+0100"
	givenSendAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAt)
	ibTimeSendAt := infobip.Time{
		T: givenSendAtDateTime,
	}

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "sendAt": "%s"
    }`, givenBulkId, givenSendAt)

	expectedRequest := fmt.Sprintf(`{
        "sendAt": "%s"
    }`, givenSendAt)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", VoiceTtsBulksEndpoint, givenResponse, 200)

	request := voice.BulkRequest{
		SendAt: ibTimeSendAt,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.RescheduleSentBulk(context.Background()).BulkId(givenBulkId).BulkRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeSendAt, response.GetSendAt())
}

func TestShouldGetSentBulksStatus(t *testing.T) {
	givenBulkId := "string"
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "status": "%s"
    }`, givenBulkId, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", VoiceTtsBulksStatusEndpoint, givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.GetSentBulksStatus(context.Background()).BulkId(givenBulkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldManageSentBulksStatus(t *testing.T) {
	givenBulkId := "string"
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
      "bulkId": "%s",
      "status": "%s"
  }`, givenBulkId, givenStatus)

	expectedRequest := fmt.Sprintf(`{
      "status": "%s"
  }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", VoiceTtsBulksStatusEndpoint, givenResponse, 200)

	request := voice.BulkStatusRequest{
		Status: givenStatus,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.ManageSentBulksStatus(context.Background()).BulkId(givenBulkId).BulkStatusRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldSendSingleVoiceTts(t *testing.T) {
	givenBulkId := "4fda521a-c680-470d-b134-83d468f7ac80"
	givenTo := "41793026727"
	givenStatusGroupId := int32(1)
	givenStatusGroupName := "PENDING"
	givenStatusId := int32(26)
	givenStatusName := "PENDING_ACCEPTED"
	givenStatusDescription := "Message accepted, pending for delivery."
	givenMessageId := "2250be2d4219-3af1-78856-aabe-1362af1edfd2"

	givenText := "Test Voice message."
	givenLanguage := "en"
	givenName := "Joanna"
	givenGender := "female"
	givenFrom := "442032864231"

	givenResponse := fmt.Sprintf(`{
       "bulkId": "%s",
       "messages": [
           {
               "to": "%s",
               "status": {
                   "groupId": %d,
                   "groupName": "%s",
                   "id": %d,
                   "name": "%s",
                   "description": "%s"
               },
               "messageId": "%s"
           }
       ]
   }`, givenBulkId, givenTo, givenStatusGroupId, givenStatusGroupName, givenStatusId, givenStatusName, givenStatusDescription, givenMessageId)

	expectedRequest := fmt.Sprintf(`{
       "text": "%s",
       "language": "%s",
       "voice": {
           "name": "%s",
           "gender": "%s"
       },
       "from": "%s",
       "to": "%s"
   }`, givenText, givenLanguage, givenName, givenGender, givenFrom, givenTo)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", VoiceTtsSingleEndpoint, givenResponse, 200)

	givenVoice := voice.Voice{
		Name:   &givenName,
		Gender: &givenGender,
	}

	request := voice.SingleRequest{
		Text:     &givenText,
		Language: &givenLanguage,
		Voice:    &givenVoice,
		From:     givenFrom,
		To:       givenTo,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.SendSingleVoiceTts(context.Background()).SingleRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Len(t, response.GetMessages(), 1)

	message := response.GetMessages()[0]
	assert.Equal(t, givenTo, message.GetTo())
	status := message.GetStatus()
	assert.Equal(t, givenStatusGroupId, status.GetGroupId())
	assert.Equal(t, givenStatusGroupName, status.GetGroupName())
	assert.Equal(t, givenStatusId, status.GetId())
	assert.Equal(t, givenStatusName, status.GetName())
	assert.Equal(t, givenStatusDescription, status.GetDescription())
	assert.Equal(t, givenMessageId, message.GetMessageId())
}

func TestShouldSendMultipleVoiceTts(t *testing.T) {
	givenBulkId := "4fda521a-c680-470d-b134-83d468f7ac80"
	givenTo := "41793026727"
	givenStatusGroupId := int32(1)
	givenStatusGroupName := "PENDING"
	givenStatusId := int32(26)
	givenStatusName := "PENDING_ACCEPTED"
	givenStatusDescription := "Message accepted, pending for delivery."
	givenMessageId := "2250be2d4219-3af1-78856-aabe-1362af1edfd2"

	givenAudioFileUrl := "https://www.example.com/media.mp3"
	givenReqFrom := "41793026700"
	givenReqTo1 := "41793026727"
	givenReqTo2 := "41793026731"
	givenText := "Hello world!"
	givenLanguage := "en"
	givenName := "Joanna"
	givenGender := "female"
	givenTo1 := "41793026785"

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "messages": [
            {
                "to": "%s",
                "status": {
                    "groupId": %d,
                    "groupName": "%s",
                    "id": %d,
                    "name": "%s",
                    "description": "%s"
                },
                "messageId": "%s"
            }
        ]
    }`, givenBulkId, givenTo, givenStatusGroupId, givenStatusGroupName, givenStatusId, givenStatusName, givenStatusDescription, givenMessageId)

	expectedRequest := fmt.Sprintf(`{
        "messages": [
            {
                "audioFileUrl": "%s",
                "from": "%s",
                "to": [
                    "%s",
                    "%s"
                ]
            },
            {
                "from": "%s",
                "to": [
                    "%s"
                ],
                "language": "%s",
                "text": "%s",
                "voice": {
                    "gender": "%s",
                    "name": "%s"
                }
            }
        ]
    }`, givenAudioFileUrl, givenReqFrom, givenReqTo1, givenReqTo2, givenReqFrom, givenTo1, givenLanguage, givenText, givenGender, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", VoiceTtsMultiEndpoint, givenResponse, 200)

	request := voice.MultiRequest{
		Messages: []voice.MultiMessage{
			{
				AudioFileUrl: &givenAudioFileUrl,
				From:         &givenReqFrom,
				To:           []string{givenReqTo1, givenReqTo2},
			},
			{
				Text:     &givenText,
				Language: &givenLanguage,
				Voice: &voice.Voice{
					Name:   &givenName,
					Gender: &givenGender,
				},
				From: &givenReqFrom,
				To:   []string{givenTo1},
			},
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.SendMultipleVoiceTts(context.Background()).MultiRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Len(t, response.GetMessages(), 1)

	message := response.GetMessages()[0]
	assert.Equal(t, givenTo, message.GetTo())
	status := message.GetStatus()
	assert.Equal(t, givenStatusGroupId, status.GetGroupId())
	assert.Equal(t, givenStatusGroupName, status.GetGroupName())
	assert.Equal(t, givenStatusId, status.GetId())
	assert.Equal(t, givenStatusName, status.GetName())
	assert.Equal(t, givenStatusDescription, status.GetDescription())
	assert.Equal(t, givenMessageId, message.GetMessageId())
}

func TestShouldGetVoices(t *testing.T) {
	givenLanguage := "en"
	givenName1 := "Benjamin"
	givenGender1 := "male"
	givenName2 := "Ivy"
	givenGender2 := "female"
	givenName3 := "Joanna"
	givenGender3 := "female"
	givenName4 := "Joey"
	givenGender4 := "male"

	givenResponse := fmt.Sprintf(`{
        "voices": [
            {
                "name": "%s",
                "gender": "%s"
            },
            {
                "name": "%s",
                "gender": "%s"
            },
            {
                "name": "%s",
                "gender": "%s"
            },
            {
                "name": "%s",
                "gender": "%s"
            }
        ]
    }`, givenName1, givenGender1, givenName2, givenGender2, givenName3, givenGender3, givenName4, givenGender4)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(VoiceTtsVoiceByIdEndpoint, givenLanguage), givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.GetVoices(context.Background(), givenLanguage).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.GetVoices(), 4)

	voices := response.GetVoices()
	assert.Equal(t, givenName1, voices[0].GetName())
	assert.Equal(t, givenGender1, voices[0].GetGender())
	assert.Equal(t, givenName2, voices[1].GetName())
	assert.Equal(t, givenGender2, voices[1].GetGender())
	assert.Equal(t, givenName3, voices[2].GetName())
	assert.Equal(t, givenGender3, voices[2].GetGender())
	assert.Equal(t, givenName4, voices[3].GetName())
	assert.Equal(t, givenGender4, voices[3].GetGender())
}

func TestShouldGetVoiceDeliveryReports(t *testing.T) {
	givenBulkId := "8c20f086-d82b-48cc-b2b3-3ca5f7aca9fb"
	givenMessageId := "ff4804ef-6ab6-4abd-984d-ab3b1387e852"
	givenFrom := "385333444"
	givenTo := "385981178"
	givenSentAt := "2018-06-25T13:38:14.730+0000"
	givenMccMnc := "21901"
	givenCallbackData := "DLR callback data"
	givenFeature := "Voice-message"
	givenStartTime := "2018-06-25T13:38:15.000+0000"
	givenAnswerTime := "2018-06-25T13:38:25.000+0000"
	givenEndTime := "2018-06-25T13:38:28.316+0000"
	givenDuration := int32(10)
	givenChargedDuration := int32(30)
	givenFileDuration := 19.3
	givenScenarioId := "333"
	givenScenarioName := "Scenario name"
	givenPricePerSecond := float32(0.01)
	givenCurrency := "EUR"
	givenStatusGroupId := int32(3)
	givenStatusGroupName := "DELIVERED"
	givenStatusId := int32(5)
	givenStatusName := "DELIVERED_TO_HANDSET"
	givenStatusDescription := "Message delivered to handset"
	givenErrorGroupId := int32(0)
	givenErrorGroupName := "OK"
	givenErrorId := int32(5000)
	givenErrorName := "VOICE_ANSWERED"
	givenErrorDescription := "Call answered by human"
	givenErrorPermanent := true

	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "bulkId": "%s",
                "messageId": "%s",
                "from": "%s",
                "to": "%s",
                "sentAt": "%s",
                "mccMnc": "%s",
                "callbackData": "%s",
                "voiceCall": {
                    "feature": "%s",
                    "startTime": "%s",
                    "answerTime": "%s",
                    "endTime": "%s",
                    "duration": %d,
                    "chargedDuration": %d,
                    "fileDuration": %.1f,
                    "ivr": {
                        "scenarioId": "%s",
                        "scenarioName": "%s"
                    }
                },
                "price": {
                    "pricePerSecond": %.2f,
                    "currency": "%s"
                },
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
    }`, givenBulkId, givenMessageId, givenFrom, givenTo, givenSentAt, givenMccMnc, givenCallbackData, givenFeature, givenStartTime, givenAnswerTime, givenEndTime, givenDuration, givenChargedDuration, givenFileDuration, givenScenarioId, givenScenarioName, givenPricePerSecond, givenCurrency, givenStatusGroupId, givenStatusGroupName, givenStatusId, givenStatusName, givenStatusDescription, givenErrorGroupId, givenErrorGroupName, givenErrorId, givenErrorName, givenErrorDescription, givenErrorPermanent)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", VoiceReportsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.GetVoiceDeliveryReports(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.GetResults(), 1)

	result := response.GetResults()[0]
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenFrom, result.GetFrom())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenSentAt, result.GetSentAt())
	assert.Equal(t, givenMccMnc, result.GetMccMnc())
	assert.Equal(t, givenCallbackData, result.GetCallbackData())

	voiceCall := result.GetVoiceCall()
	assert.Equal(t, givenFeature, voiceCall.GetFeature())
	assert.Equal(t, givenStartTime, voiceCall.GetStartTime())
	assert.Equal(t, givenAnswerTime, voiceCall.GetAnswerTime())
	assert.Equal(t, givenEndTime, voiceCall.GetEndTime())
	assert.Equal(t, givenDuration, voiceCall.GetDuration())
	assert.Equal(t, givenChargedDuration, voiceCall.GetChargedDuration())
	assert.Equal(t, givenFileDuration, voiceCall.GetFileDuration())

	ivr := voiceCall.GetIvr()
	assert.Equal(t, givenScenarioId, ivr.GetScenarioId())
	assert.Equal(t, givenScenarioName, ivr.GetScenarioName())

	price := result.GetPrice()
	assert.Equal(t, givenPricePerSecond, price.GetPricePerSecond())
	assert.Equal(t, givenCurrency, price.GetCurrency())

	status := result.GetStatus()
	assert.Equal(t, givenStatusGroupId, status.GetGroupId())
	assert.Equal(t, givenStatusGroupName, status.GetGroupName())
	assert.Equal(t, givenStatusId, status.GetId())
	assert.Equal(t, givenStatusName, status.GetName())
	assert.Equal(t, givenStatusDescription, status.GetDescription())

	error := result.GetError()
	assert.Equal(t, givenErrorGroupId, error.GetGroupId())
	assert.Equal(t, givenErrorGroupName, error.GetGroupName())
	assert.Equal(t, givenErrorId, error.GetId())
	assert.Equal(t, givenErrorName, error.GetName())
	assert.Equal(t, givenErrorDescription, error.GetDescription())
	assert.Equal(t, givenErrorPermanent, error.GetPermanent())
}

func TestShouldGetVoiceLogs(t *testing.T) {
	givenBulkId := "06479ba3-5977-47f6-9346-fee0369bc76b"
	givenMessageId := "1f21d8d7-f306-4f53-9f6e-eddfce9849ea"
	givenTo := "41793026727"
	givenFrom := "41793026700"
	givenText := "Test voice message."
	givenSendAt := "2015-02-22T17:42:05.390+0100"
	givenSendAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAt)
	ibTimeSendAt := infobip.Time{
		T: givenSendAtDateTime,
	}
	givenDoneAt := "2015-02-22T17:42:05.390+0100"
	givenDoneAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSendAt)
	ibTimeDoneAt := infobip.Time{
		T: givenDoneAtDateTime,
	}
	givenDuration := int32(10)
	givenMccMnc := "22801"
	givenPricePerSecond := float32(0.01)
	givenCurrency := "EUR"
	givenStatusGroupId := int32(3)
	givenStatusGroupName := "DELIVERED"
	givenStatusId := int32(5)
	givenStatusName := "DELIVERED_TO_HANDSET"
	givenStatusDescription := "Message delivered to handset"
	givenErrorGroupId := int32(0)
	givenErrorGroupName := "OK"
	givenErrorId := int32(5003)
	givenErrorName := "EC_VOICE_NO_ANSWER"
	givenErrorDescription := "User was notified, but did not answer call"
	givenErrorPermanent := true

	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "bulkId": "%s",
                "messageId": "%s",
                "to": "%s",
                "from": "%s",
                "text": "%s",
                "sentAt": "%s",
                "doneAt": "%s",
                "duration": %d,
                "mccMnc": "%s",
                "price": {
                    "pricePerSecond": %.2f,
                    "currency": "%s"
                },
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
    }`, givenBulkId, givenMessageId, givenTo, givenFrom, givenText, givenSendAt, givenDoneAt, givenDuration, givenMccMnc, givenPricePerSecond, givenCurrency, givenStatusGroupId, givenStatusGroupName, givenStatusId, givenStatusName, givenStatusDescription, givenErrorGroupId, givenErrorGroupName, givenErrorId, givenErrorName, givenErrorDescription, givenErrorPermanent)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", VoiceTtsLogsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.GetSentVoiceLogs(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.GetResults(), 1)

	result := response.GetResults()[0]
	assert.Equal(t, givenBulkId, result.GetBulkId())
	assert.Equal(t, givenMessageId, result.GetMessageId())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenFrom, result.GetFrom())
	assert.Equal(t, givenText, result.GetText())
	assert.Equal(t, ibTimeSendAt, result.GetSentAt())
	assert.Equal(t, ibTimeDoneAt, result.GetDoneAt())
	assert.Equal(t, givenDuration, result.GetDuration())
	assert.Equal(t, givenMccMnc, result.GetMccMnc())

	price := result.GetPrice()
	assert.Equal(t, givenPricePerSecond, price.GetPricePerSecond())
	assert.Equal(t, givenCurrency, price.GetCurrency())

	status := result.GetStatus()
	assert.Equal(t, givenStatusGroupId, status.GetGroupId())
	assert.Equal(t, givenStatusGroupName, status.GetGroupName())
	assert.Equal(t, givenStatusId, status.GetId())
	assert.Equal(t, givenStatusName, status.GetName())
	assert.Equal(t, givenStatusDescription, status.GetDescription())

	error := result.GetError()
	assert.Equal(t, givenErrorGroupId, error.GetGroupId())
	assert.Equal(t, givenErrorGroupName, error.GetGroupName())
	assert.Equal(t, givenErrorId, error.GetId())
	assert.Equal(t, givenErrorName, error.GetName())
	assert.Equal(t, givenErrorDescription, error.GetDescription())
	assert.Equal(t, givenErrorPermanent, error.GetPermanent())
}

func TestShouldSearchVoiceIvrRecordedFiles(t *testing.T) {
	givenMessageId1 := "453e161a-fe4f-4f3c-80c0-ab520de9a969"
	givenFrom1 := "442032864231"
	givenTo1 := "38712345678"
	givenScenarioId1 := "C9CE33CF130511D8E333C1260BABA309"
	givenGroupId1 := "#/script/1"
	givenUrl1 := "/voice/ivr/1/files/3C67336FA555A606C85FA9637906A6AB98436B7AFC65D857A416F6521D39F8F0E1D3D2469FF580D8968D3DD89A2DB561"
	givenRecordedAt1 := "2021-08-24T15:00:00.000+0000"
	givenRecordedAt1DateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenRecordedAt1)
	ibTimeRecordedAt1 := infobip.Time{
		T: givenRecordedAt1DateTime,
	}

	givenRecordedAt2 := "2021-08-24T15:00:00.000+0000"
	givenRecordedAt2DateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenRecordedAt2)
	ibTimeRecordedAt2 := infobip.Time{
		T: givenRecordedAt2DateTime,
	}

	givenMessageId2 := "05b2859d-85c6-4068-9347-2e563b5c9cf4"
	givenFrom2 := "442032864231"
	givenTo2 := "38712345678"
	givenScenarioId2 := "4A6177C9B92039306F1F091708851A2E"
	givenGroupId2 := "#/script/1"
	givenUrl2 := "/voice/ivr/1/files/305DE72BA11D81D1BAED75BFC46706761580BDEC2218C22628447FD3814E7913D3058E4ECBFD6F55C80E976235EEB111"

	givenResponse := fmt.Sprintf(`{
      "files": [
          {
              "messageId": "%s",
              "from": "%s",
              "to": "%s",
              "scenarioId": "%s",
              "groupId": "%s",
              "url": "%s",
              "recordedAt": "%s"
          },
          {
              "messageId": "%s",
              "from": "%s",
              "to": "%s",
              "scenarioId": "%s",
              "groupId": "%s",
              "url": "%s",
              "recordedAt": "%s"
          }
      ]
  }`, givenMessageId1, givenFrom1, givenTo1, givenScenarioId1, givenGroupId1, givenUrl1, givenRecordedAt1, givenMessageId2, givenFrom2, givenTo2, givenScenarioId2, givenGroupId2, givenUrl2, givenRecordedAt2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", VoiceIvrFilesEndpoint, givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.SearchVoiceIvrRecordedFiles(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.GetFiles(), 2)

	file1 := response.GetFiles()[0]
	assert.Equal(t, givenMessageId1, file1.GetMessageId())
	assert.Equal(t, givenFrom1, file1.GetFrom())
	assert.Equal(t, givenTo1, file1.GetTo())
	assert.Equal(t, givenScenarioId1, file1.GetScenarioId())
	assert.Equal(t, givenGroupId1, file1.GetGroupId())
	assert.Equal(t, givenUrl1, file1.GetUrl())
	assert.Equal(t, ibTimeRecordedAt1, file1.GetRecordedAt())

	file2 := response.GetFiles()[1]
	assert.Equal(t, givenMessageId2, file2.GetMessageId())
	assert.Equal(t, givenFrom2, file2.GetFrom())
	assert.Equal(t, givenTo2, file2.GetTo())
	assert.Equal(t, givenScenarioId2, file2.GetScenarioId())
	assert.Equal(t, givenGroupId2, file2.GetGroupId())
	assert.Equal(t, givenUrl2, file2.GetUrl())
	assert.Equal(t, ibTimeRecordedAt2, file2.GetRecordedAt())
}

func TestShouldGetVoiceIvrUploadedFiles(t *testing.T) {
	givenId1 := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName1 := "example-file.wav"

	givenId2 := "328fcfda-d145-541e-0g37-9g2b8g1h3e14"
	givenName2 := "another-file.mp3"

	givenResponse := fmt.Sprintf(`{
      "files": [
          {
              "id": "%s",
              "name": "%s"
          },
          {
              "id": "%s",
              "name": "%s"
          }
      ]
  }`, givenId1, givenName1, givenId2, givenName2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", VoiceIvrUploadsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.IvrUploadGetFiles(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestShouldUploadVoiceIvrAudioFile(t *testing.T) {
	givenId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "uploaded-audio.wav"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}
	givenExpirationTime := "2021-08-25T15:00:00.000+0000"
	givenExpirationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenExpirationTime)
	ibTimeExpirationTime := infobip.Time{
		T: givenExpirationTimeDateTime,
	}
	givenDuration := int32(3)

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "fileFormat": "%s",
       "size": %d,
       "creationTime": "%s",
       "expirationTime": "%s",
       "duration": %d
   }`, givenId, givenName, givenFileFormat, givenSize, givenCreationTime, givenExpirationTime, givenDuration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", VoiceIvrUploadsEndpoint, givenResponse, 200)

	tempFile, err := os.CreateTemp("", "ivr-audio-*.wav")
	assert.Nil(t, err)
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	response, _, err := infobipClient.VoiceAPI.IvrUploadAudioFile(context.Background()).File(tempFile).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, string(givenFileFormat), response.GetFileFormat())
	assert.Equal(t, givenSize, response.GetSize())
	assert.Equal(t, ibTimeCreationTime, response.GetCreationTime())
	assert.Equal(t, ibTimeExpirationTime, response.GetExpirationTime())
	assert.Equal(t, int64(givenDuration), response.GetDuration())
}

func TestShouldGetVoiceIvrScenarioById(t *testing.T) {
	givenId := "E83E787CF2613450157ADA3476171E3F"
	givenName := "scenario"
	givenDescription := "Description"
	givenCreateTime := "2021-08-24T15:00:00.000+0000"
	givenCreateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreateTime)
	ibTimeCreateTime := infobip.Time{
		T: givenCreateTimeDateTime,
	}
	givenUpdateTime := "2021-08-24T15:00:00.000+0000"
	givenUpdateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreateTime)
	ibTimeUpdateTime := infobip.Time{
		T: givenUpdateTimeDateTime,
	}

	givenLastUsageDate := "2021-08-28"

	givenVoiceIvrResponse := voice.UpdateScenarioResponse{
		Id:            &givenId,
		Name:          &givenName,
		Description:   &givenDescription,
		CreateTime:    &ibTimeCreateTime,
		UpdateTime:    &ibTimeUpdateTime,
		LastUsageDate: &givenLastUsageDate,
	}

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "description": "%s",
       "createTime": "%s",
       "updateTime": "%s",
	   "lastUsageDate": "%s"
   }`, givenId, givenName, givenDescription, givenCreateTime, givenUpdateTime, givenLastUsageDate)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(VoiceIvrScenarioByIdEndpoint, givenId), givenResponse, 200)

	response, _, err := infobipClient.VoiceAPI.GetAVoiceIvrScenario(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.Equal(t, givenLastUsageDate, response.GetLastUsageDate())
	assert.Equal(t, &givenVoiceIvrResponse, response)
}

func TestShouldCreateVoiceIvrScenario(t *testing.T) {
	givenId := "E83E787CF2613450157ADA3476171E3F"
	givenName := "scenario"
	givenDescription := ""
	givenCreateTime := "2015-02-22T17:42:05.390+0100"
	givenCreateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreateTime)
	ibTimeCreateTime := infobip.Time{
		T: givenCreateTimeDateTime,
	}
	givenUpdateTime := "2015-02-22T17:42:05.390+0100"
	givenUpdateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenUpdateTime)
	ibTimeUpdateTime := infobip.Time{
		T: givenUpdateTimeDateTime,
	}
	givenDial := "dial"
	givenScript := fmt.Sprintf(`[{"dial": "%s"}]`, givenDial)

	givenRequest := fmt.Sprintf(`{
       "name": "%s",
       "description": "%s",
       "script": [
           {
               "dial": "%s"
           }
       ]
   }`, givenName, givenDescription, givenDial)

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "description": "%s",
       "createTime": "%s",
       "updateTime": "%s",
       "script": [
			{
               "dial": "%s"
           }
       ]
   }`, givenId, givenName, givenDescription, givenCreateTime, givenUpdateTime, givenDial)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", VoiceIvrScenariosEndpoint, givenResponse, 200)

	request := voice.UpdateScenarioRequest{
		Name:        givenName,
		Description: &givenDescription,
		Script:      givenScript,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.CreateAVoiceIvrScenario(context.Background()).UpdateScenarioRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenDescription, response.GetDescription())
	assert.Equal(t, ibTimeCreateTime, response.GetCreateTime())
	assert.Equal(t, ibTimeUpdateTime, response.GetUpdateTime())
	assert.NotNil(t, response.GetScript())
}

func TestShouldCreateVoiceIvrScenarios(t *testing.T) {
	givenId := "E83E787CF2613450157ADA3476171E3F"
	givenName := "Capture speech or digit"
	givenDescription := "Capture speech or digit"
	givenCreateTime := "2023-09-14T15:13:36.735+0000"
	givenCreateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreateTime)
	ibTimeCreateTime := infobip.Time{
		T: givenCreateTimeDateTime,
	}

	givenUpdateTime := "2015-02-22T17:42:05.390+0100"
	givenUpdateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenUpdateTime)
	ibTimeUpdateTime := infobip.Time{
		T: givenUpdateTimeDateTime,
	}
	givenScript := `[
            {
                "say": "Say discount or press 1 to get discount. Say exit or press 0 to exit."
            },
            {
                "capture": "myVar",
                "timeout": 5,
                "speechOptions": {
                    "language": "en-US",
                    "maxSilence": 2,
                    "keyPhrases": [
                        "discount",
                        "exit"
                    ]
                },
                "dtmfOptions": {
                    "maxInputLength": 1
                }
            },
            {
                "if": "${myVar == 'discount' || myVar == '1'}",
                "then": [
                    {
                        "say": "You will get discount"
                    }
                ],
                "else": [
                    {
                        "if": "${myVar == 'exit' || myVar == '0'}",
                        "then": [
                            {
                                "say": "Goodbye"
                            }
                        ],
                        "else": [
                            {
                                "say": "I did not understand"
                            }
                        ]
                    }
                ]
            },
            "hangup"
        ]`
	givenDial := "dial"

	givenRequest := `{
        "name": "Capture speech or digit",
        "description": "Capture speech or digit",
        "script": [
            {
                "say": "Say discount or press 1 to get discount. Say exit or press 0 to exit."
            },
            {
                "capture": "myVar",
                "timeout": 5,
                "speechOptions": {
                    "language": "en-US",
                    "maxSilence": 2,
                    "keyPhrases": [
                        "discount",
                        "exit"
                    ]
                },
                "dtmfOptions": {
                    "maxInputLength": 1
                }
            },
            {
                "if": "${myVar == 'discount' || myVar == '1'}",
                "then": [
                    {
                        "say": "You will get discount"
                    }
                ],
                "else": [
                    {
                        "if": "${myVar == 'exit' || myVar == '0'}",
                        "then": [
                            {
                                "say": "Goodbye"
                            }
                        ],
                        "else": [
                            {
                                "say": "I did not understand"
                            }
                        ]
                    }
                ]
            },
            "hangup"
        ]
    }`

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "description": "%s",
       "createTime": "%s",
       "updateTime": "%s",
       "script": [
			{
               "dial": "%s"
           }
       ]
   }`, givenId, givenName, givenDescription, givenCreateTime, givenUpdateTime, givenDial)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", VoiceIvrScenariosEndpoint, givenResponse, 200)

	request := voice.UpdateScenarioRequest{
		Name:        givenName,
		Description: &givenDescription,
		Script:      givenScript,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.CreateAVoiceIvrScenario(context.Background()).UpdateScenarioRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenDescription, response.GetDescription())
	assert.Equal(t, ibTimeCreateTime, response.GetCreateTime())
	assert.Equal(t, ibTimeUpdateTime, response.GetUpdateTime())
	assert.NotNil(t, response.GetScript())
}

func TestShouldUpdateVoiceIvrScenario(t *testing.T) {
	givenId := "E83E787CF2613450157ADA3476171E3F"
	givenName := "Call API"
	givenDescription := "Perform a POST request to provided URL with headers and payload."
	givenRequestUrl := "https://requestb.in/12345"
	givenRequestMethod := "POST"
	givenRequestHeaderKey := "content-type"
	givenRequestHeaderValue := "application/json"
	givenRequestBodyPayload := "payload"
	givenCreateTime := "2015-02-22T17:42:05.390+0100"
	givenCreateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreateTime)
	ibTimeCreateTime := infobip.Time{
		T: givenCreateTimeDateTime,
	}
	givenUpdateTime := "2015-02-22T17:42:05.390+0100"
	givenUpdateTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenUpdateTime)
	ibTimeUpdateTime := infobip.Time{
		T: givenUpdateTimeDateTime,
	}
	givenScript := fmt.Sprintf(`[{"request":"%s","options":{"method":"%s","headers":{"%s":"%s"},"body":"%s"}}]`, givenRequestUrl, givenRequestMethod, givenRequestHeaderKey, givenRequestHeaderValue, givenRequestBodyPayload)

	givenRequest := fmt.Sprintf(`{
       "name": "%s",
       "description": "%s",
       "script": [
           {
               "request": "%s",
               "options": {
                   "method": "%s",
                   "headers": {
                       "%s": "%s"
                   },
                   "body": "%s"
               }
           }
       ]
   }`, givenName, givenDescription, givenRequestUrl, givenRequestMethod, givenRequestHeaderKey, givenRequestHeaderValue, givenRequestBodyPayload)

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "description": "%s",
       "createTime": "%s",
       "updateTime": "%s",
       "script": [
           {
               "request": "%s",
               "options": {
                   "method": "%s",
                   "headers": {
                       "%s": "%s"
                   },
                   "body": "%s"
               }
           }
       ]
   }`, givenId, givenName, givenDescription, givenCreateTime, givenUpdateTime, givenRequestUrl, givenRequestMethod, givenRequestHeaderKey, givenRequestHeaderValue, givenRequestBodyPayload)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf(VoiceIvrScenarioByIdEndpoint, givenId), givenResponse, 200)

	request := voice.UpdateScenarioRequest{
		Name:        givenName,
		Description: &givenDescription,
		Script:      givenScript,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	response, _, err := infobipClient.VoiceAPI.UpdateVoiceIvrScenario(context.Background(), givenId).UpdateScenarioRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenDescription, response.GetDescription())
	assert.Equal(t, ibTimeCreateTime, response.GetCreateTime())
	assert.Equal(t, ibTimeUpdateTime, response.GetUpdateTime())
	assert.Equal(t, givenScript, response.GetScript())
}

func TestShouldDeleteVoiceIvrScenario(t *testing.T) {
	givenId := "E83E787CF2613450157ADA3476171E3F"
	givenResponse := "{}"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf(VoiceIvrScenarioByIdEndpoint, givenId), givenResponse, 200)

	_, err := infobipClient.VoiceAPI.DeleteAVoiceIvrScenario(context.Background(), givenId).Execute()

	assert.Nil(t, err)
}

func TestShouldDownloadVoiceIvrRecordedFile(t *testing.T) {
	givenFileId := "3C67336FA555A606C85FA9637906A6AB98436B7AFC65D857A416F6521D39F8F0E1D3D2469FF580D8968D3DD89A2DB561"
	givenBinaryData := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpGetRequestWithFileContent(fmt.Sprintf("/voice/ivr/1/files/%s", givenFileId), givenBinaryData, 200)

	file, _, err := infobipClient.VoiceAPI.DownloadVoiceIvrRecordedFile(context.Background(), givenFileId).Execute()

	assert.Nil(t, err, "Expected no error while executing the request")
	assert.NotNil(t, file, "Expected a valid file in the response")

	content, err := io.ReadAll(file)
	assert.Nil(t, err, "Expected no error while reading the file content")
	assert.Equal(t, givenBinaryData, content, "File content does not match the expected binary data")
}

// ==============================================================================
// Webhook deserialization tests
// ==============================================================================

func TestReceiveNumberMaskingCallbackWebhook(t *testing.T) {
	givenFrom := "441134960001"
	givenTo := "441134960002"
	givenCorrelationId := "corr-123"
	givenNmCorrelationId := "nm-corr-456"

	givenPayload := fmt.Sprintf(`
		{
			"from": "%s",
			"to": "%s",
			"correlationId": "%s",
			"nmCorrelationId": "%s"
		}`,
		givenFrom,
		givenTo,
		givenCorrelationId,
		givenNmCorrelationId,
	)

	var requestBody voice.NumberMaskingCallbackRequest
	err := json.Unmarshal([]byte(givenPayload), &requestBody)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, givenFrom, requestBody.GetFrom())
	assert.Equal(t, givenTo, requestBody.GetTo())
	assert.Equal(t, givenCorrelationId, requestBody.GetCorrelationId())
	assert.Equal(t, givenNmCorrelationId, requestBody.GetNmCorrelationId())
}

func TestReceiveNumberMaskingCallbackWebhook_WithDtmf(t *testing.T) {
	givenPayload := `
		{
			"from": "441134960001",
			"to": "441134960002",
			"correlationId": "corr-123",
			"nmCorrelationId": "nm-corr-456",
			"dtmfCaptured": true
		}`

	var requestBody voice.NumberMaskingCallbackRequest
	err := json.Unmarshal([]byte(givenPayload), &requestBody)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, "441134960001", requestBody.GetFrom())
	assert.Equal(t, true, requestBody.GetDtmfCaptured())
}

func TestReceiveMaskedCallReportWebhook(t *testing.T) {
	givenAction := "dial"
	givenFrom := "441134960001"
	givenTo := "441134960002"
	givenTransferTo := "441134960003"
	givenDuration := 45
	givenStatus := "answered"
	givenNmCorrelationId := "nm-corr-456"
	givenCorrelationId := "corr-123"

	givenPayload := fmt.Sprintf(`
		{
			"action": "%s",
			"from": "%s",
			"to": "%s",
			"transferTo": "%s",
			"duration": %d,
			"status": "%s",
			"nmCorrelationId": "%s",
			"correlationId": "%s",
			"inboundDuration": 50,
			"calculatedDuration": 60
		}`,
		givenAction,
		givenFrom,
		givenTo,
		givenTransferTo,
		givenDuration,
		givenStatus,
		givenNmCorrelationId,
		givenCorrelationId,
	)

	var requestBody voice.NumberMaskingStatusRequest
	err := json.Unmarshal([]byte(givenPayload), &requestBody)

	assert.Nil(t, err, "Expected nil error")
	assert.Equal(t, givenAction, requestBody.GetAction())
	assert.Equal(t, givenFrom, requestBody.GetFrom())
	assert.Equal(t, givenTo, requestBody.GetTo())
	assert.Equal(t, givenTransferTo, requestBody.GetTransferTo())
	assert.Equal(t, int32(givenDuration), requestBody.GetDuration())
	assert.Equal(t, givenStatus, requestBody.GetStatus())
	assert.Equal(t, givenNmCorrelationId, requestBody.GetNmCorrelationId())
	assert.Equal(t, givenCorrelationId, requestBody.GetCorrelationId())
	assert.Equal(t, int32(50), requestBody.GetInboundDuration())
	assert.Equal(t, int32(60), requestBody.GetCalculatedDuration())
}
