package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestShouldApplicationTransfer(t *testing.T) {
	givenStatus := voice.ACTIONSTATUS_PENDING
	givenTimeout := int32(20)
	givenCustomDataField := "custom"
	givenCustomDataFieldValue := "data"
	givenEntityId := "entityId"
	givenDestinationApplicationId := "61c060db2675060027d8c7a6"
	givenDestinationCallsConfigurationId := "dc5942707c704551a00cd2ea"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenDestinationApplicationId,
	}
	givenCustomData := map[string]string{
		givenCustomDataField: givenCustomDataFieldValue,
	}

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
		"destinationCallsConfigurationId": "%s",
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "timeout": %d,
        "customData": {
            "%s": "%s"
        }
    }`, givenDestinationCallsConfigurationId, givenDestinationApplicationId, givenEntityId, givenTimeout, givenCustomDataField, givenCustomDataFieldValue)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/application-transfer", givenDestinationApplicationId), givenResponse, 200)

	request := voice.ApplicationTransferRequest{
		DestinationCallsConfigurationId: givenDestinationCallsConfigurationId,
		Platform:                        &givenPlatform,
		Timeout:                         &givenTimeout,
		CustomData:                      &givenCustomData,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ApplicationTransfer(context.Background(), givenDestinationApplicationId).ApplicationTransferRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldApplicationTransferAccept(t *testing.T) {
	givenCallId := "123"
	givenTransferId := "1234"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/application-transfer/%s/accept", givenCallId, givenTransferId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.ApplicationTransferAccept(context.Background(), givenCallId, givenTransferId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldApplicationTransferReject(t *testing.T) {
	givenCallId := "123"
	givenTransferId := "1234"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/application-transfer/%s/reject", givenCallId, givenTransferId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.ApplicationTransferReject(context.Background(), givenCallId, givenTransferId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldGetCalls(t *testing.T) {
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenEntityId := "entityId"
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId,
	}
	givenFrom := "44790123456"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenUserMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenParentCallId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenDetectionResult := voice.DETECTIONRESULT_HUMAN
	givenRingDuration := int32(3)
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenKey1 := "value1"
	givenKey2 := "value2"
	givenPage := int32(0)
	givenPageSize := int32(1)
	givenPageTotalPages := int32(0)
	givenPageTotalResults := int64(0)

	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "id": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "from": "%s",
                "to": "%s",
                "direction": "%s",
                "state": "%s",
                "media": {
                    "audio": {
                        "muted": %t,
                        "userMuted": %t,
                        "deaf": %t
                    },
                    "video": {
                        "camera": %t,
                        "screenShare": %t
                    }
                },
                "startTime": "%s",
                "answerTime": "%s",
                "endTime": "%s",
                "parentCallId": "%s",
                "machineDetection": {
                    "detectionResult": "%s"
                },
                "ringDuration": %d,
                "platform": {
                    "applicationId": "%s",
                    "entityId": "%s"
                },
                "conferenceId": "%s",
                "customData": {
                    "key1": "%s",
                    "key2": "%s"
                }
            }
        ],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenUserMuted, givenDeaf, givenCamera, givenScreenShare, givenStartTime, givenAnswerTime, givenEndTime, givenParentCallId, givenDetectionResult, givenRingDuration, givenApplicationId, givenEntityId, givenConferenceId, givenKey1, givenKey2, givenPage, givenPageSize, givenPageTotalPages, givenPageTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/calls", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCalls(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenPage, response.GetPaging().Page)
	assert.Equal(t, givenPageSize, response.GetPaging().Size)
	assert.Equal(t, givenPageTotalPages, response.GetPaging().TotalPages)
	assert.Equal(t, givenPageTotalResults, response.GetPaging().TotalResults)
	assert.Len(t, response.GetResults(), 1)

	result := response.GetResults()[0]
	assert.Equal(t, givenCallId, result.GetId())
	assert.Equal(t, givenType, result.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, result.GetFrom())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenDirection, result.GetDirection())
	assert.Equal(t, givenCallState, result.GetState())
	resultMedia := result.GetMedia()
	resultAudio := resultMedia.GetAudio()
	resultVideo := resultMedia.GetVideo()
	assert.Equal(t, givenMuted, resultAudio.GetMuted())
	assert.Equal(t, givenUserMuted, resultAudio.GetUserMuted())
	assert.Equal(t, givenDeaf, resultAudio.GetDeaf())
	assert.Equal(t, givenCamera, resultVideo.GetCamera())
	assert.Equal(t, givenScreenShare, resultVideo.GetScreenShare())
	assert.Equal(t, ibTimeStartTime, result.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, result.GetAnswerTime())
	assert.Equal(t, ibTimeEndTime, result.GetEndTime())
	assert.Equal(t, givenParentCallId, result.GetParentCallId())
	assert.Equal(t, &givenDetectionResult, result.GetMachineDetection().DetectionResult)
	assert.Equal(t, givenRingDuration, result.GetRingDuration())
	assert.Equal(t, givenPlatform, result.GetPlatform())
	assert.Equal(t, givenConferenceId, result.GetConferenceId())
	assert.Equal(t, givenKey1, result.GetCustomData()["key1"])
	assert.Equal(t, givenKey2, result.GetCustomData()["key2"])
}

func TestShouldCreateCall(t *testing.T) {
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenFrom := "44790123456"
	givenConnectTimeout := int32(0)
	givenRecordingType := voice.RECORDINGTYPE_AUDIO
	givenMaxDuration := int32(28800)
	givenEntityId := "entityId"
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId,
	}
	givenParentCallId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenCallsConfigurationId := "dc5942707c704551a00cd2ea"

	expectedRequest := fmt.Sprintf(`{
		"callsConfigurationId": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "from": "%s",
        "connectTimeout": %d,
        "recording": {
            "recordingType": "%s"
        },
        "maxDuration": %d,
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "parentCallId": "%s"
    }`, givenCallsConfigurationId, givenPhoneNumber, givenType, givenFrom, givenConnectTimeout, givenRecordingType, givenMaxDuration, givenApplicationId, givenEntityId, givenParentCallId)

	givenCallId := "string"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenUserMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenDetectionResult := voice.DETECTIONRESULT_HUMAN
	givenRingDuration := int32(3)
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenKey1 := "value1"
	givenKey2 := "value2"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "from": "%s",
        "to": "%s",
        "direction": "%s",
        "state": "%s",
        "media": {
            "audio": {
                "muted": %t,
                "userMuted": %t,
                "deaf": %t
            },
            "video": {
                "camera": %t,
                "screenShare": %t
            }
        },
        "startTime": "%s",
        "answerTime": "%s",
        "endTime": "%s",
        "parentCallId": "%s",
        "machineDetection": {
            "detectionResult": "%s"
        },
        "ringDuration": %d,
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "conferenceId": "%s",
        "customData": {
            "key1": "%s",
            "key2": "%s"
        }
    }`, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenUserMuted, givenDeaf, givenCamera, givenScreenShare, givenStartTime, givenAnswerTime, givenEndTime, givenParentCallId, givenDetectionResult, givenRingDuration, givenApplicationId, givenEntityId, givenConferenceId, givenKey1, givenKey2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/calls", givenResponse, 201)

	request := voice.CallRequest{
		CallsConfigurationId: givenCallsConfigurationId,
		Endpoint:             voice.PhoneEndpointAsCallEndpoint(voice.NewPhoneEndpoint(givenPhoneNumber)),
		From:                 givenFrom,
		ConnectTimeout:       &givenConnectTimeout,
		Recording: &voice.CallRecordingRequest{
			RecordingType: givenRecordingType,
		},
		MaxDuration:  &givenMaxDuration,
		Platform:     &givenPlatform,
		ParentCallId: &givenParentCallId,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateCall(context.Background()).CallRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenCallId, response.GetId())
	assert.Equal(t, givenType, response.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, response.GetFrom())
	assert.Equal(t, givenTo, response.GetTo())
	assert.Equal(t, givenDirection, response.GetDirection())
	assert.Equal(t, givenCallState, response.GetState())
	resultMedia := response.GetMedia()
	resultAudio := resultMedia.GetAudio()
	resultVideo := resultMedia.GetVideo()
	assert.Equal(t, givenMuted, resultAudio.GetMuted())
	assert.Equal(t, givenUserMuted, resultAudio.GetUserMuted())
	assert.Equal(t, givenDeaf, resultAudio.GetDeaf())
	assert.Equal(t, givenCamera, resultVideo.GetCamera())
	assert.Equal(t, givenScreenShare, resultVideo.GetScreenShare())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, response.GetAnswerTime())
	assert.Equal(t, ibTimeEndTime, response.GetEndTime())
	assert.Equal(t, givenParentCallId, response.GetParentCallId())
	assert.Equal(t, &givenDetectionResult, response.GetMachineDetection().DetectionResult)
	assert.Equal(t, givenRingDuration, response.GetRingDuration())
	assert.Equal(t, givenPlatform, response.GetPlatform())
	assert.Equal(t, givenConferenceId, response.GetConferenceId())
	assert.Equal(t, givenKey1, response.GetCustomData()["key1"])
	assert.Equal(t, givenKey2, response.GetCustomData()["key2"])
}

func TestShouldGetCall(t *testing.T) {
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}
	givenFrom := "44790123456"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenUserMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenParentCallId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenDetectionResult := voice.DETECTIONRESULT_HUMAN
	givenRingDuration := int32(3)
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenKey1 := "value1"
	givenKey2 := "value2"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "from": "%s",
        "to": "%s",
        "direction": "%s",
        "state": "%s",
        "media": {
            "audio": {
                "muted": %t,
                "userMuted": %t,
                "deaf": %t
            },
            "video": {
                "camera": %t,
                "screenShare": %t
            }
        },
        "startTime": "%s",
        "answerTime": "%s",
        "endTime": "%s",
        "parentCallId": "%s",
        "machineDetection": {
            "detectionResult": "%s"
        },
        "ringDuration": %d,
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "conferenceId": "%s",
        "customData": {
            "key1": "%s",
            "key2": "%s"
        }
    }`, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenUserMuted, givenDeaf, givenCamera, givenScreenShare, givenStartTime, givenAnswerTime, givenEndTime, givenParentCallId, givenDetectionResult, givenRingDuration, givenApplicationId, givenEntityId, givenConferenceId, givenKey1, givenKey2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/calls/%s", givenCallId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCall(context.Background(), givenCallId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenCallId, response.GetId())
	assert.Equal(t, givenType, response.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, response.GetFrom())
	assert.Equal(t, givenTo, response.GetTo())
	assert.Equal(t, givenDirection, response.GetDirection())
	assert.Equal(t, givenCallState, response.GetState())
	assert.Equal(t, givenMuted, response.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenUserMuted, response.GetMedia().Audio.GetUserMuted())
	assert.Equal(t, givenDeaf, response.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenCamera, response.GetMedia().Video.GetCamera())
	assert.Equal(t, givenScreenShare, response.GetMedia().Video.GetScreenShare())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, response.GetAnswerTime())
	assert.Equal(t, ibTimeEndTime, response.GetEndTime())
	assert.Equal(t, givenParentCallId, response.GetParentCallId())
	assert.Equal(t, &givenDetectionResult, response.GetMachineDetection().DetectionResult)
	assert.Equal(t, givenRingDuration, response.GetRingDuration())
	assert.Equal(t, givenPlatform, response.GetPlatform())
	assert.Equal(t, givenConferenceId, response.GetConferenceId())
	assert.Equal(t, givenKey1, response.GetCustomData()["key1"])
	assert.Equal(t, givenKey2, response.GetCustomData()["key2"])
}

func TestShouldGetCallsHistory(t *testing.T) {
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenFrom := "44790123456"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenUserMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenParentCallId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenDetectionResult := voice.DETECTIONRESULT_HUMAN
	givenDuration := int64(5)
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenHasCameraVideo := false
	givenHasScreenShareVideo := false
	givenErrorCodeId := int32(2)
	givenErrorCodeName := "NORMAL_HANGUP"
	givenErrorCodeDescription := "The call has ended with hangup initiated by caller, callee or API"
	givenKey1 := "value1"
	givenKey2 := "value2"
	givenPage := int32(0)
	givenPageSize := int32(1)
	givenPageTotalPages := int32(0)
	givenPageTotalResults := int64(0)
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "callId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "from": "%s",
                "to": "%s",
                "direction": "%s",
                "state": "%s",
                "media": {
                    "audio": {
                        "muted": %t,
                        "userMuted": %t,
                        "deaf": %t
                    },
                    "video": {
                        "camera": %t,
                        "screenShare": %t
                    }
                },
                "startTime": "%s",
                "answerTime": "%s",
                "endTime": "%s",
                "parentCallId": "%s",
                "machineDetection": {
                    "detectionResult": "%s"
                },
                "platform": {
                    "applicationId": "%s",
                    "entityId": "%s"
                },
                "conferenceIds": [
                    "%s"
                ],
                "duration": %d,
                "hasCameraVideo": %t,
                "hasScreenshareVideo": %t,
                "errorCode": {
                    "id": %d,
                    "name": "%s",
                    "description": "%s"
                },
                "customData": {
                    "key1": "%s",
                    "key2": "%s"
                }
            }
        ],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenUserMuted, givenDeaf, givenCamera, givenScreenShare, givenStartTime, givenAnswerTime, givenEndTime, givenParentCallId, givenDetectionResult, givenApplicationId, givenEntityId, givenConferenceId, givenDuration, givenHasCameraVideo, givenHasScreenShareVideo, givenErrorCodeId, givenErrorCodeName, givenErrorCodeDescription, givenKey1, givenKey2, givenPage, givenPageSize, givenPageTotalPages, givenPageTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/calls/history", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallsHistory(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenPage, response.GetPaging().Page)
	assert.Equal(t, givenPageSize, response.GetPaging().Size)
	assert.Equal(t, givenPageTotalPages, response.GetPaging().TotalPages)
	assert.Equal(t, givenPageTotalResults, response.GetPaging().TotalResults)
	assert.Len(t, response.GetResults(), 1)

	result := response.GetResults()[0]
	assert.Equal(t, givenCallId, result.GetCallId())
	assert.Equal(t, givenType, result.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, result.GetFrom())
	assert.Equal(t, givenTo, result.GetTo())
	assert.Equal(t, givenDirection, result.GetDirection())
	assert.Equal(t, givenCallState, result.GetState())
	assert.Equal(t, ibTimeStartTime, result.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, result.GetAnswerTime())
	assert.Equal(t, ibTimeEndTime, result.GetEndTime())
	assert.Equal(t, givenParentCallId, result.GetParentCallId())
	assert.Equal(t, &givenDetectionResult, result.GetMachineDetection().DetectionResult)
	assert.Equal(t, givenPlatform, result.GetPlatform())
	assert.Equal(t, givenConferenceId, result.GetConferenceIds()[0])
	assert.Equal(t, givenDuration, result.GetDuration())
	assert.Equal(t, givenHasCameraVideo, result.GetHasCameraVideo())
	assert.Equal(t, givenHasScreenShareVideo, result.GetHasScreenshareVideo())
	assert.Equal(t, &givenErrorCodeId, result.GetErrorCode().Id)
	assert.Equal(t, &givenErrorCodeName, result.GetErrorCode().Name)
	assert.Equal(t, &givenErrorCodeDescription, result.GetErrorCode().Description)
	assert.Equal(t, givenKey1, result.GetCustomData()["key1"])
	assert.Equal(t, givenKey2, result.GetCustomData()["key2"])
}

func TestShouldGetCallHistory(t *testing.T) {
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenFrom := "44790123456"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenUserMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenParentCallId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenDetectionResult := voice.DETECTIONRESULT_HUMAN
	givenDuration := int64(5)
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenHasCameraVideo := false
	givenHasScreenShareVideo := false
	givenErrorCodeId := int32(2)
	givenErrorCodeName := "NORMAL_HANGUP"
	givenErrorCodeDescription := "The call has ended with hangup initiated by caller, callee or API"
	givenKey1 := "value1"
	givenKey2 := "value2"
	givenDialogId := "dialogId"
	givenSender := "sender"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "callId": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "from": "%s",
        "to": "%s",
        "direction": "%s",
        "state": "%s",
        "media": {
            "audio": {
                "muted": %t,
                "userMuted": %t,
                "deaf": %t
            },
            "video": {
                "camera": %t,
                "screenShare": %t
            }
        },
        "startTime": "%s",
        "answerTime": "%s",
        "endTime": "%s",
        "parentCallId": "%s",
        "machineDetection": {
            "detectionResult": "%s"
        },
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "conferenceIds": [
            "%s"
        ],
        "duration": %d,
        "hasCameraVideo": %t,
        "hasScreenshareVideo": %t,
        "errorCode": {
            "id": %d,
            "name": "%s",
            "description": "%s"
        },
        "customData": {
            "key1": "%s",
            "key2": "%s"
        },
        "dialogId": "%s",
        "sender": "%s"
    }`, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenUserMuted, givenDeaf, givenCamera, givenScreenShare, givenStartTime, givenAnswerTime, givenEndTime, givenParentCallId, givenDetectionResult, givenApplicationId, givenEntityId, givenConferenceId, givenDuration, givenHasCameraVideo, givenHasScreenShareVideo, givenErrorCodeId, givenErrorCodeName, givenErrorCodeDescription, givenKey1, givenKey2, givenDialogId, givenSender)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/calls/%s/history", givenCallId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallHistory(context.Background(), givenCallId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenCallId, response.GetCallId())
	assert.Equal(t, givenType, response.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, response.GetFrom())
	assert.Equal(t, givenTo, response.GetTo())
	assert.Equal(t, givenDirection, response.GetDirection())
	assert.Equal(t, givenCallState, response.GetState())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, response.GetAnswerTime())
	assert.Equal(t, ibTimeEndTime, response.GetEndTime())
	assert.Equal(t, givenParentCallId, response.GetParentCallId())
	assert.Equal(t, &givenDetectionResult, response.GetMachineDetection().DetectionResult)
	assert.Equal(t, givenPlatform, response.GetPlatform())
	assert.Equal(t, givenConferenceId, response.GetConferenceIds()[0])
	assert.Equal(t, givenDuration, response.GetDuration())
	assert.Equal(t, givenHasCameraVideo, response.GetHasCameraVideo())
	assert.Equal(t, givenHasScreenShareVideo, response.GetHasScreenshareVideo())
	assert.Equal(t, &givenErrorCodeId, response.GetErrorCode().Id)
	assert.Equal(t, &givenErrorCodeName, response.GetErrorCode().Name)
	assert.Equal(t, &givenErrorCodeDescription, response.GetErrorCode().Description)
	assert.Equal(t, givenKey1, response.GetCustomData()["key1"])
	assert.Equal(t, givenKey2, response.GetCustomData()["key2"])
	assert.Equal(t, givenDialogId, response.GetDialogId())
	assert.Equal(t, givenSender, response.GetSender())
}

func TestShouldConnectCalls(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId1 := "d6d6058c-5077-49f9-a030-2fc40e8ca195"
	givenCallId2 := "6539fcb4-4b2a-4ac9-a43a-d60807af29b0"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING

	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "participants": [
            {
                "callId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "state": "%s",
                "joinTime": "%s",
                "media": {
                    "audio": {
                        "muted": %t,
                        "deaf": %t
                    },
                    "video": {
                        "camera": %t,
                        "screenShare": %t
                    }
                }
            },
            {
                "callId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "state": "%s",
                "joinTime": "%s",
                "media": {
                    "audio": {
                        "muted": %t,
                        "deaf": %t
                    },
                    "video": {
                        "camera": %t,
                        "screenShare": %t
                    }
                }
            }
        ],
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        }
    }`, givenId, givenName, givenCallId1, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenCallId2, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId)

	expectedRequest := fmt.Sprintf(`{
        "callIds": [
            "%s",
            "%s"
        ]
    }`, givenCallId1, givenCallId2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/connect", givenResponse, 200)

	request := voice.ConnectRequest{
		CallIds: []string{givenCallId1, givenCallId2},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ConnectCalls(context.Background()).ConnectRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.NotNil(t, response.GetParticipants())
	assert.Len(t, response.GetParticipants(), 2)

	participant := response.GetParticipants()[0]
	assert.NotNil(t, participant.GetEndpoint())
	assert.Equal(t, givenType, participant.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant.GetState())
	assert.Equal(t, ibTimeJoinTime, participant.GetJoinTime())
	assert.NotNil(t, participant.GetMedia())
	assert.NotNil(t, participant.GetMedia().Audio)
	assert.Equal(t, givenAudioMuted, participant.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserDeaf, participant.GetMedia().Audio.GetDeaf())
	assert.NotNil(t, participant.GetMedia().Video)
	assert.Equal(t, givenVideoCamera, participant.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant.GetMedia().Video.GetScreenShare())

	participant2 := response.GetParticipants()[1]
	assert.NotNil(t, participant2.GetEndpoint())
	assert.Equal(t, givenType, participant2.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant2.GetState())
	assert.Equal(t, ibTimeJoinTime, participant2.GetJoinTime())
	assert.NotNil(t, participant2.GetMedia())
	assert.NotNil(t, participant2.GetMedia().Audio)
	assert.Equal(t, givenAudioMuted, participant2.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserDeaf, participant2.GetMedia().Audio.GetDeaf())
	assert.NotNil(t, participant2.GetMedia().Video)
	assert.Equal(t, givenVideoCamera, participant2.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant2.GetMedia().Video.GetScreenShare())

	assert.Equal(t, givenPlatform, response.GetPlatform())
}

func TestShouldConnectWithNewCall(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING
	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenFrom := "44790123456"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenUserMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenParentCallId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenDetectionResult := voice.DETECTIONRESULT_HUMAN
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenKey1 := "value1"
	givenKey2 := "value2"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "conference": {
            "id": "%s",
            "name": "%s",
            "participants": [
                {
                    "callId": "%s",
                    "endpoint": {
                        "phoneNumber": "%s",
                        "type": "%s"
                    },
                    "state": "%s",
                    "joinTime": "%s",
                    "media": {
                        "audio": {
                            "muted": %t,
                            "userMuted": %t,
                            "deaf": %t
                        },
                        "video": {
                            "camera": %t,
                            "screenShare": %t
                        }
                    }
                }
            ],
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            }
        },
        "call": {
            "id": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "from": "%s",
            "to": "%s",
            "direction": "%s",
            "state": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "userMuted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            },
            "startTime": "%s",
            "answerTime": "%s",
            "endTime": "%s",
            "parentCallId": "%s",
            "machineDetection": {
                "detectionResult": "%s"
            },
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "conferenceId": "%s",
            "customData": {
                "key1": "%s",
                "key2": "%s"
            }
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenUserMuted, givenDeaf, givenCamera, givenScreenShare, givenStartTime, givenAnswerTime, givenEndTime, givenParentCallId, givenDetectionResult, givenApplicationId, givenEntityId, givenConferenceId, givenKey1, givenKey2)

	givenConnectOnEarlyMedia := false

	expectedRequest := fmt.Sprintf(`{
        "callRequest": {
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "from": "%s"
        },
        "connectOnEarlyMedia": %t
    }`, givenPhoneNumber, givenType, givenFrom, givenConnectOnEarlyMedia)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/connect", givenId), givenResponse, 200)

	request := voice.ConnectWithNewCallRequest{
		CallRequest: &voice.ActionCallRequest{
			Endpoint: voice.PhoneEndpointAsCallEndpoint(voice.NewPhoneEndpoint(givenPhoneNumber)),
			From:     givenFrom,
		},
		ConnectOnEarlyMedia: &givenConnectOnEarlyMedia,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ConnectWithNewCall(context.Background(), givenId).ConnectWithNewCallRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)

	conference := response.GetConference()
	assert.NotNil(t, conference)
	assert.Equal(t, givenId, conference.GetId())
	assert.Equal(t, givenName, conference.GetName())

	participants := conference.GetParticipants()
	assert.NotNil(t, participants)
	assert.Len(t, participants, 1)

	participant := participants[0]
	assert.Equal(t, givenCallId, participant.GetCallId())
	assert.NotNil(t, participant.GetEndpoint())
	assert.Equal(t, givenType, participant.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant.GetState())
	assert.Equal(t, ibTimeJoinTime, participant.GetJoinTime())
	assert.NotNil(t, participant.GetMedia())
	assert.NotNil(t, participant.GetMedia().Audio)
	assert.Equal(t, givenAudioMuted, participant.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserMuted, participant.GetMedia().Audio.GetUserMuted())
	assert.Equal(t, givenAudioUserDeaf, participant.GetMedia().Audio.GetDeaf())
	assert.NotNil(t, participant.GetMedia().Video)
	assert.Equal(t, givenVideoCamera, participant.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant.GetMedia().Video.GetScreenShare())
	assert.Equal(t, givenPlatform, conference.GetPlatform())

	call := response.GetCall()
	assert.NotNil(t, call)
	assert.Equal(t, givenType, call.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, call.GetFrom())
	assert.Equal(t, givenTo, call.GetTo())
	assert.Equal(t, givenDirection, call.GetDirection())
	assert.Equal(t, givenCallState, call.GetState())
	assert.NotNil(t, call.GetMedia().Audio)
	assert.Equal(t, givenMuted, call.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenUserMuted, call.GetMedia().Audio.GetUserMuted())
	assert.Equal(t, givenDeaf, call.GetMedia().Audio.GetDeaf())
	assert.NotNil(t, call.GetMedia().Video)
	assert.Equal(t, givenCamera, call.GetMedia().Video.GetCamera())
	assert.Equal(t, givenScreenShare, call.GetMedia().Video.GetScreenShare())
	assert.Equal(t, ibTimeStartTime, call.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, call.GetAnswerTime())
	assert.Equal(t, ibTimeEndTime, call.GetEndTime())
	assert.Equal(t, givenParentCallId, call.GetParentCallId())
	assert.Equal(t, &givenDetectionResult, call.GetMachineDetection().DetectionResult)
	assert.Equal(t, givenPlatform, call.GetPlatform())
	assert.Equal(t, givenConferenceId, call.GetConferenceId())
	assert.NotNil(t, call.GetCustomData())
	assert.Equal(t, givenKey1, call.GetCustomData()["key1"])
	assert.Equal(t, givenKey2, call.GetCustomData()["key2"])
}

func TestShouldPreAnswerCall(t *testing.T) {
	givenCallId := "123"
	givenCustomData := map[string]string{"custom": "data"}
	givenCallPreRequest := voice.PreAnswerRequest{
		CustomData: &givenCustomData,
	}
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenRequest := fmt.Sprintf(`{
		"customData": {
			"custom": "data"
		}
	}`)

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/pre-answer", givenCallId), givenResponse, 200)

	actualRequest, _ := json.Marshal(givenCallPreRequest)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.PreAnswerCall(context.Background(), givenCallId).PreAnswerRequest(givenCallPreRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldAnswerCall(t *testing.T) {
	givenCallId := "123"
	givenRecordingType := voice.RECORDINGTYPE_AUDIO
	givenStatus := voice.ACTIONSTATUS_PENDING
	givenCustomDataField := "custom"
	givenCustomDataFieldValue := "data"
	givenCustomData := map[string]string{givenCustomDataField: givenCustomDataFieldValue}

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "recording": {
            "recordingType": "%s"
        },
        "customData": {
            "%s": "%s"
        }
    }`, givenRecordingType, givenCustomDataField, givenCustomDataFieldValue)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/answer", givenCallId), givenResponse, 200)

	request := voice.AnswerRequest{
		Recording: &voice.CallRecordingRequest{
			RecordingType: givenRecordingType,
		},
		CustomData: &givenCustomData,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.AnswerCall(context.Background(), givenCallId).AnswerRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldHangupCall(t *testing.T) {
	givenErrorCode := voice.ERRORCODE_NORMAL_HANGUP
	givenCallId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "44790123456"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenFrom := "44790123456"
	givenTo := "44790123456"
	givenDirection := voice.CALLDIRECTION_OUTBOUND
	givenState := voice.CALLSTATE_CALLING
	givenAudioMuted := false
	givenAudioDeaf := false
	givenVideoCamera := false
	givenVideoScreenShare := false
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenAnswerTime := "2021-08-24T15:00:00.000+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}
	givenRingDuration := int32(2)
	givenEntityId := "entityId"
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenKey2 := "value2"
	givenKey1 := "value1"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "from": "%s",
        "to": "%s",
        "direction": "%s",
        "state": "%s",
        "media": {
            "audio": {
                "muted": %t,
                "deaf": %t
            },
            "video": {
                "camera": %t,
                "screenShare": %t
            }
        },
        "startTime": "%s",
        "answerTime": "%s",
        "ringDuration": %d,
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "conferenceId": "%s",
        "customData": {
            "key2": "%s",
            "key1": "%s"
        }
    }`, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenState, givenAudioMuted, givenAudioDeaf, givenVideoCamera, givenVideoScreenShare, givenStartTime, givenAnswerTime, givenRingDuration, givenApplicationId, givenEntityId, givenConferenceId, givenKey2, givenKey1)

	expectedRequest := fmt.Sprintf(`{
        "errorCode": "%s"
    }`, givenErrorCode)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/hangup", givenCallId), givenResponse, 200)

	request := voice.HangupRequest{
		ErrorCode: &givenErrorCode,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.HangupCall(context.Background(), givenCallId).HangupRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenCallId, response.GetId())
	assert.Equal(t, givenFrom, response.GetFrom())
	assert.Equal(t, givenTo, response.GetTo())
	assert.Equal(t, givenDirection, response.GetDirection())
	assert.Equal(t, givenState, response.GetState())
	assert.Equal(t, givenAudioMuted, response.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioDeaf, response.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, response.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, response.GetMedia().Video.GetScreenShare())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, ibTimeAnswerTime, response.GetAnswerTime())
	assert.Equal(t, givenRingDuration, response.GetRingDuration())
	assert.Equal(t, givenPlatform, response.GetPlatform())
	assert.Equal(t, givenConferenceId, response.GetConferenceId())
	assert.Equal(t, givenKey2, response.GetCustomData()["key2"])
	assert.Equal(t, givenKey1, response.GetCustomData()["key1"])
}

func TestShouldCallPlayFile(t *testing.T) {
	givenCallId := "123"
	givenStatus := voice.ACTIONSTATUS_PENDING
	givenLoopCount := int32(0)
	givenContentType := voice.PLAYCONTENTTYPE_FILE
	givenFileId := "100"

	givenResponse := fmt.Sprintf(`{
      "status": "%s"
  }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
      "loopCount": %d,
      "content": {
          "fileId": "%s",
          "type": "%s"
      }
  }`, givenLoopCount, givenFileId, givenContentType)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/play", givenCallId), givenResponse, 200)

	filePlayContent := voice.NewFilePlayContent(givenFileId)

	request := voice.PlayRequest{
		LoopCount: &givenLoopCount,
		Content:   voice.FilePlayContentAsPlayContent(filePlayContent),
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallPlayFile(context.Background(), givenCallId).PlayRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCallStopPlayingFile(t *testing.T) {
	givenCallId := "123"
	givenCustomData := map[string]string{"custom": "data"}
	givenCallsStopPlayRequest := voice.StopPlayRequest{
		CustomData: &givenCustomData,
	}
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenRequest := fmt.Sprintf(`{
		"customData": {
			"custom": "data"
		}
	}`)

	actualRequest, _ := json.Marshal(givenCallsStopPlayRequest)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/stop-play", givenCallId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.CallStopPlayingFile(context.Background(), givenCallId).StopPlayRequest(givenCallsStopPlayRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCallSayText(t *testing.T) {
	givenCallId := "123"
	givenText := "string"
	givenLanguage := voice.LANGUAGE_AR
	givenSpeechRate := 0.5
	givenLoopCount := int32(0)
	givenCallsGender := voice.GENDER_FEMALE
	givenVoiceName := voice.VOICENAME_DARINA
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "text": "%s",
        "language": "%s",
        "speechRate": %f,
        "loopCount": %d,
        "preferences": {
            "voiceGender": "%s",
            "voiceName": "%s"
        }
    }`, givenText, givenLanguage, givenSpeechRate, givenLoopCount, givenCallsGender, givenVoiceName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/say", givenCallId), givenResponse, 200)

	request := voice.CallSayRequest{
		Text:       givenText,
		Language:   givenLanguage,
		SpeechRate: &givenSpeechRate,
		LoopCount:  &givenLoopCount,
		Preferences: &voice.VoicePreferences{
			VoiceGender: &givenCallsGender,
			VoiceName:   &givenVoiceName,
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallSayText(context.Background(), givenCallId).CallSayRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCallSendDtmf(t *testing.T) {
	givenCallId := "123"
	givenDtmf := "3"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "dtmf": "%s"
    }`, givenDtmf)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/send-dtmf", givenCallId), givenResponse, 200)

	request := voice.DtmfSendRequest{
		Dtmf: givenDtmf,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallSendDtmf(context.Background(), givenCallId).DtmfSendRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCallCaptureDtmf(t *testing.T) {
	givenCallId := "123"
	givenMaxLength := int32(0)
	givenTimeout := int32(0)
	givenTerminator := "string"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "maxLength": %d,
        "timeout": %d,
        "terminator": "%s"
    }`, givenMaxLength, givenTimeout, givenTerminator)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/capture/dtmf", givenCallId), givenResponse, 200)

	request := voice.DtmfCaptureRequest{
		MaxLength:  givenMaxLength,
		Timeout:    givenTimeout,
		Terminator: &givenTerminator,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallCaptureDtmf(context.Background(), givenCallId).DtmfCaptureRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCaptureSpeech(t *testing.T) {
	givenCallId := "123"
	givenStatus := voice.ACTIONSTATUS_PENDING
	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedLanguage := voice.TRANSCRIPTIONLANGUAGE_EN_GB
	expectedTimeout := int32(30)
	expectedMaxSilence := int32(3)
	expectedKeyPhrase1 := "phrase"
	expectedKeyPhrase2 := "word"
	expectedRequest := fmt.Sprintf(`{
        "language": "%s",
        "timeout": %d,
        "maxSilence": %d,
        "keyPhrases": [
            "%s",
            "%s"
        ]
    }`, expectedLanguage, expectedTimeout, expectedMaxSilence, expectedKeyPhrase1, expectedKeyPhrase2)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/capture/speech", givenCallId), givenResponse, 200)

	request := voice.SpeechCaptureRequest{
		Language:   expectedLanguage,
		Timeout:    expectedTimeout,
		MaxSilence: &expectedMaxSilence,
		KeyPhrases: []string{expectedKeyPhrase1, expectedKeyPhrase2},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallCaptureSpeech(context.Background(), givenCallId).SpeechCaptureRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCallStartRecording(t *testing.T) {
	givenCallId := "123"
	givenFilePrefix := "filePrefix"
	givenCustomDataField := "custom"
	givenCustomDataFieldValue := "data"
	givenCustomData := map[string]string{givenCustomDataField: givenCustomDataFieldValue}
	givenRecordingType := voice.RECORDINGTYPE_AUDIO
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "recording": {
            "recordingType": "%s",
            "customData": {
                "%s": "%s"
            },
            "filePrefix": "%s"
        }
    }`, givenRecordingType, givenCustomDataField, givenCustomDataFieldValue, givenFilePrefix)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/start-recording", givenCallId), givenResponse, 200)

	request := voice.RecordingStartRequest{
		Recording: &voice.RecordingRequest{
			RecordingType: givenRecordingType,
			CustomData:    &givenCustomData,
			FilePrefix:    &givenFilePrefix,
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallStartRecording(context.Background(), givenCallId).RecordingStartRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCallStopRecording(t *testing.T) {
	givenCallId := "123"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/stop-recording", givenCallId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.CallStopRecording(context.Background(), givenCallId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldStartMediaStream(t *testing.T) {
	givenCallId := "123"
	givenMediaStreamConfigId := "string"
	givenReplaceMedia := true
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "mediaStream": {
            "audioProperties": {
                "mediaStreamConfigId": "%s",
                "replaceMedia": %t
            }
        }
    }`, givenMediaStreamConfigId, givenReplaceMedia)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/start-media-stream", givenCallId), givenResponse, 200)

	givenMediaStreamProperties := &voice.MediaStreamAudioProperties{
		MediaStreamConfigId: givenMediaStreamConfigId,
		ReplaceMedia:        &givenReplaceMedia,
	}

	request := voice.StartMediaStreamRequest{
		MediaStream: voice.MediaStream{
			AudioProperties: givenMediaStreamProperties,
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.StartMediaStream(context.Background(), givenCallId).StartMediaStreamRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldStopMediaStream(t *testing.T) {
	givenCallId := "123"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/stop-media-stream", givenCallId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.StopMediaStream(context.Background(), givenCallId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldGetConferences(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING
	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenPage := int32(0)
	givenPageSize := int32(1)
	givenPageTotalPages := int32(0)
	givenPageTotalResults := int64(0)
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "results": [{
            "id": "%s",
            "name": "%s",
            "participants": [{
                "callId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "state": "%s",
                "joinTime": "%s",
                "media": {
                    "audio": {
                        "muted": %t,
                        "userMuted": %t,
                        "deaf": %t
                    },
                    "video": {
                        "camera": %t,
                        "screenShare": %t
                    }
                }
            }],
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            }
        }],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId, givenPage, givenPageSize, givenPageTotalPages, givenPageTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/conferences", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetConferences(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.Results[0].GetId())
	assert.Equal(t, givenName, response.Results[0].GetName())
	assert.Equal(t, givenCallId, response.Results[0].Participants[0].GetCallId())
	assert.Equal(t, givenType, response.Results[0].Participants[0].Endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenState, response.Results[0].Participants[0].GetState())
	assert.Equal(t, ibTimeJoinTime, response.Results[0].Participants[0].GetJoinTime())
	assert.Equal(t, givenAudioMuted, response.Results[0].Participants[0].Media.Audio.GetMuted())
	assert.Equal(t, givenAudioUserMuted, response.Results[0].Participants[0].Media.Audio.GetUserMuted())
	assert.Equal(t, givenAudioUserDeaf, response.Results[0].Participants[0].Media.Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, response.Results[0].Participants[0].Media.Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, response.Results[0].Participants[0].Media.Video.GetScreenShare())
	assert.Equal(t, givenPlatform, response.Results[0].GetPlatform())
	assert.Equal(t, givenPage, response.Paging.GetPage())
	assert.Equal(t, givenPageSize, response.Paging.GetSize())
	assert.Equal(t, givenPageTotalPages, response.Paging.GetTotalPages())
	assert.Equal(t, givenPageTotalResults, response.Paging.GetTotalResults())
}

func TestShouldCreateConference(t *testing.T) {
	givenRecordingType := voice.RECORDINGTYPE_AUDIO
	givenCompositionEnabled := false
	givenMaxDuration := int32(28800)
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING
	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenCustomDataField := "custom"
	givenCustomDataFieldValue := "data"
	givenCustomData := map[string]string{givenCustomDataField: givenCustomDataFieldValue}
	givenFilePrefix := "filePrefix"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}
	givenCallsConfigurationId := "dc5942707c704551a00cd2ea"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "participants": [{
            "callId": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "state": "%s",
            "joinTime": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "userMuted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            }
        }],
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId)

	expectedRequest := fmt.Sprintf(`{
        "name": "%s",
		"callsConfigurationId": "%s",
        "recording": {
            "recordingType": "%s",
            "conferenceComposition": {
                "enabled": %t
            },
            "customData": {
                "%s": "%s"
            },
            "filePrefix": "%s"
        },
        "maxDuration": %d,
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        }
    }`, givenName, givenCallsConfigurationId, givenRecordingType, givenCompositionEnabled, givenCustomDataField, givenCustomDataFieldValue, givenFilePrefix, givenMaxDuration, givenApplicationId, givenEntityId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/conferences", givenResponse, 201)

	request := voice.ConferenceRequest{
		Name:                 &givenName,
		CallsConfigurationId: givenCallsConfigurationId,
		Recording: &voice.ConferenceRecordingRequest{
			RecordingType: givenRecordingType,
			ConferenceComposition: &voice.ConferenceComposition{
				Enabled: &givenCompositionEnabled,
			},
			CustomData: &givenCustomData,
			FilePrefix: &givenFilePrefix,
		},
		MaxDuration: &givenMaxDuration,
		Platform:    &givenPlatform,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateConference(context.Background()).ConferenceRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenCallId, response.Participants[0].GetCallId())
	assert.Equal(t, givenType, response.Participants[0].Endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenState, response.Participants[0].GetState())
	assert.Equal(t, ibTimeJoinTime, response.Participants[0].GetJoinTime())
	assert.Equal(t, givenAudioMuted, response.Participants[0].Media.Audio.GetMuted())
	assert.Equal(t, givenAudioUserMuted, response.Participants[0].Media.Audio.GetUserMuted())
	assert.Equal(t, givenAudioUserDeaf, response.Participants[0].Media.Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, response.Participants[0].Media.Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, response.Participants[0].Media.Video.GetScreenShare())
	assert.Equal(t, givenPlatform, response.GetPlatform())
}

func TestShouldGetConference(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING
	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "participants": [{
            "callId": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "state": "%s",
            "joinTime": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "userMuted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            }
        }],
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/conferences/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetConference(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetParticipants()))
	participant := response.GetParticipants()[0]
	assert.Equal(t, givenCallId, participant.GetCallId())
	assert.Equal(t, givenType, participant.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant.GetState())
	assert.Equal(t, ibTimeJoinTime, participant.GetJoinTime())
	assert.Equal(t, givenAudioMuted, participant.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserMuted, participant.GetMedia().Audio.GetUserMuted())
	assert.Equal(t, givenAudioUserDeaf, participant.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, participant.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant.GetMedia().Video.GetScreenShare())
	assert.Equal(t, givenPlatform, response.GetPlatform())
}

func TestShouldGetConferencesHistory(t *testing.T) {
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenApplicationId := "61c060db2675060027d8c7a6"

	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenDuration := int64(55)

	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}

	givenLeaveTime := "2021-08-24T15:00:00.000+0000"
	givenLeaveTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenLeaveTime)
	ibTimeLeaveTime := infobip.Time{
		T: givenLeaveTimeDateTime,
	}
	givenCallId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenFileName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)

	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenFilesDuration := int64(3)

	givenFilesStartTime := "2021-08-24T15:00:00.000+0000"
	givenFilesStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenFilesStartTime)

	ibTimeFilesStartTime := infobip.Time{
		T: givenFilesStartTimeDateTime,
	}

	givenFilesEndTime := "2021-08-24T15:00:00.000+0000"
	givenFilesEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenFilesEndTime)
	ibTimeFilesEndTime := infobip.Time{
		T: givenFilesEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP
	givenStatus := voice.RECORDINGSTATUS_SUCCESSFUL
	givenReason := "string"
	givenPage := int32(0)
	givenPageSize := int32(1)
	givenPageTotalPages := int32(0)
	givenPageTotalResults := int64(0)
	givenErrorCodeId := int32(10000)
	givenErrorCodeName := "NORMAL_HANGUP"
	givenErrorCodeDescription := "The call has ended with hangup initiated by caller, callee or API"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}
	sftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED

	givenResponse := fmt.Sprintf(`{
        "results": [{
            "conferenceId": "%s",
            "name": "%s",
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "startTime": "%s",
            "endTime": "%s",
            "duration": %d,
            "sessions": [{
                "callId": "%s",
                "joinTime": "%s",
                "leaveTime": "%s"
            }],
            "recording": {
                "composedFiles": [{
                    "id": "%s",
                    "name": "%s",
                    "fileFormat": "%s",
                    "size": %d,
                    "sftpUploadStatus": "%s",
                    "creationTime": "%s",
                    "expirationTime": "%s",
                    "duration": %d,
                    "startTime": "%s",
                    "endTime": "%s"
                }],
                "callRecordings": [{
                    "callId": "%s",
                    "endpoint": {
                        "phoneNumber": "%s",
                        "type": "%s"
                    },
                    "direction": "%s",
                    "files": [{
                        "id": "%s",
                        "name": "%s",
                        "fileFormat": "%s",
                        "size": %d,
                        "sftpUploadStatus": "%s",
                        "creationTime": "%s",
                        "expirationTime": "%s",
                        "duration": %d,
                        "startTime": "%s",
                        "endTime": "%s",
                        "location": "%s"
                    }],
                    "status": "%s",
                    "reason": "%s"
                }]
            },
            "errorCode": {
                "id": %d,
                "name": "%s",
                "description": "%s"
            }
        }],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenConferenceId, givenName, givenApplicationId, givenEntityId, givenStartTime, givenEndTime, givenDuration, givenCallId, givenJoinTime, givenLeaveTime, givenFileId, givenFileName, givenFileFormat, givenSize, sftpUploadStatus, givenCreationTime, givenExpirationTime, givenFilesDuration, givenFilesStartTime, givenFilesEndTime, givenCallId, givenPhoneNumber, givenType, givenDirection, givenFileId, givenFileName, givenFileFormat, givenSize, sftpUploadStatus, givenCreationTime, givenExpirationTime, givenFilesDuration, givenFilesStartTime, givenFilesEndTime, givenLocation, givenStatus, givenReason, givenErrorCodeId, givenErrorCodeName, givenErrorCodeDescription, givenPage, givenPageSize, givenPageTotalPages, givenPageTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/conferences/history", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetConferencesHistory(context.Background()).Name(givenName).CallId(givenCallId).ApplicationId(givenApplicationId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.GetResults()))
	result := response.GetResults()[0]
	assert.Equal(t, givenConferenceId, result.GetConferenceId())
	assert.Equal(t, givenName, result.GetName())
	assert.Equal(t, givenPlatform, result.GetPlatform())
	assert.Equal(t, ibTimeStartTime, result.GetStartTime())
	assert.Equal(t, ibTimeEndTime, result.GetEndTime())
	assert.Equal(t, givenDuration, result.GetDuration())
	assert.Equal(t, 1, len(result.GetSessions()))
	session := result.GetSessions()[0]
	assert.Equal(t, givenCallId, session.GetCallId())
	assert.Equal(t, ibTimeJoinTime, session.GetJoinTime())
	assert.Equal(t, ibTimeLeaveTime, session.GetLeaveTime())
	recording := result.GetRecording()
	assert.Equal(t, 1, len(recording.GetComposedFiles()))
	file := recording.GetComposedFiles()[0]
	assert.Equal(t, givenFileId, file.GetId())
	assert.Equal(t, givenFileName, file.GetName())
	assert.Equal(t, givenFileFormat, file.GetFileFormat())
	assert.Equal(t, givenSize, file.GetSize())
	assert.Equal(t, sftpUploadStatus, file.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, file.GetCreationTime())
	assert.Equal(t, givenFilesDuration, file.GetDuration())
	assert.Equal(t, ibTimeFilesStartTime, file.GetStartTime())
	assert.Equal(t, ibTimeFilesEndTime, file.GetEndTime())
	assert.Equal(t, 1, len(recording.GetCallRecordings()))
	callRecording := recording.GetCallRecordings()[0]
	assert.Equal(t, givenCallId, callRecording.GetCallId())
	assert.Equal(t, givenType, callRecording.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenDirection, callRecording.GetDirection())
	assert.Equal(t, 1, len(callRecording.GetFiles()))
	callRecordingFile := callRecording.GetFiles()[0]
	assert.Equal(t, givenFileId, callRecordingFile.GetId())
	assert.Equal(t, givenFileName, callRecordingFile.GetName())
	assert.Equal(t, givenFileFormat, callRecordingFile.GetFileFormat())
	assert.Equal(t, givenSize, callRecordingFile.GetSize())
	assert.Equal(t, sftpUploadStatus, callRecordingFile.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, callRecordingFile.GetCreationTime())
	assert.Equal(t, givenFilesDuration, callRecordingFile.GetDuration())
	assert.Equal(t, ibTimeFilesStartTime, callRecordingFile.GetStartTime())
	assert.Equal(t, ibTimeFilesEndTime, callRecordingFile.GetEndTime())
	assert.Equal(t, givenLocation, callRecordingFile.GetLocation())
	assert.Equal(t, givenStatus, callRecording.GetStatus())
	assert.Equal(t, givenReason, callRecording.GetReason())
	errorCode := result.GetErrorCode()
	assert.Equal(t, givenErrorCodeId, errorCode.GetId())
	assert.Equal(t, givenErrorCodeName, errorCode.GetName())
	assert.Equal(t, givenErrorCodeDescription, errorCode.GetDescription())
	paging := response.GetPaging()
	assert.Equal(t, givenPage, paging.GetPage())
	assert.Equal(t, givenPageSize, paging.GetSize())
	assert.Equal(t, givenPageTotalPages, paging.GetTotalPages())
	assert.Equal(t, givenPageTotalResults, paging.GetTotalResults())
}

func TestShouldAddNewConferenceCall(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING
	givenTime := "2021-08-24T15:00:00.000+0000"
	givenTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenTime)
	ibTimeTime := infobip.Time{
		T: givenTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenFrom := "44790123456"
	givenTo := "44790987654"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenCallState := voice.CALLSTATE_CALLING
	givenMuted := true
	givenDeaf := true
	givenCamera := true
	givenScreenShare := true
	givenRingDuration := int32(3)
	givenConferenceId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenKey1 := "value1"
	givenKey2 := "value2"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}
	givenConnectOnEarlyMedia := false

	givenResponse := fmt.Sprintf(`{
        "conference": {
            "id": "%s",
            "name": "%s",
            "participants": [{
                "callId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "state": "%s",
                "joinTime": "%s",
                "media": {
                    "audio": {
                        "muted": %t,
                        "deaf": %t
                    },
                    "video": {
                        "camera": %t,
                        "screenShare": %t
                    }
                }
            }],
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            }
        },
        "call": {
            "id": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "from": "%s",
            "to": "%s",
            "direction": "%s",
            "state": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            },
            "startTime": "%s",
            "answerTime": "%s",
            "ringDuration": %d,
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "conferenceId": "%s",
            "customData": {
                "key1": "%s",
                "key2": "%s"
            }
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenTime, givenAudioMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId, givenCallId, givenPhoneNumber, givenType, givenFrom, givenTo, givenDirection, givenCallState, givenMuted, givenDeaf, givenCamera, givenScreenShare, givenTime, givenTime, givenRingDuration, givenApplicationId, givenEntityId, givenConferenceId, givenKey1, givenKey2)

	expectedRequest := fmt.Sprintf(`{
        "callRequest": {
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "from": "%s"
        },
        "connectOnEarlyMedia": %t
    }`, givenPhoneNumber, givenType, givenFrom, givenConnectOnEarlyMedia)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/call", givenId), givenResponse, 200)

	request := voice.AddNewCallRequest{
		CallRequest: voice.ActionCallRequest{
			Endpoint: voice.PhoneEndpointAsCallEndpoint(voice.NewPhoneEndpoint(givenPhoneNumber)),
			From:     givenFrom,
		},
		ConnectOnEarlyMedia: &givenConnectOnEarlyMedia,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.AddNewConferenceCall(context.Background(), givenConferenceId).AddNewCallRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, &givenId, response.GetConference().Id)
	assert.Equal(t, &givenName, response.GetConference().Name)
	assert.Equal(t, 1, len(response.GetConference().Participants))
	participant := response.Conference.GetParticipants()[0]
	assert.Equal(t, givenCallId, participant.GetCallId())
	assert.Equal(t, givenType, participant.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant.GetState())
	assert.Equal(t, ibTimeTime, participant.GetJoinTime())
	assert.Equal(t, givenAudioMuted, participant.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserDeaf, participant.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, participant.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant.GetMedia().Video.GetScreenShare())
	assert.Equal(t, &givenPlatform, response.GetConference().Platform)
	call := response.GetCall()
	assert.Equal(t, givenCallId, call.GetId())
	assert.Equal(t, givenType, call.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenFrom, call.GetFrom())
	assert.Equal(t, givenTo, call.GetTo())
	assert.Equal(t, givenDirection, call.GetDirection())
	assert.Equal(t, givenCallState, call.GetState())
	assert.Equal(t, givenAudioMuted, call.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserDeaf, call.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, call.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, call.GetMedia().Video.GetScreenShare())
	assert.Equal(t, ibTimeTime, call.GetStartTime())
	assert.Equal(t, ibTimeTime, call.GetAnswerTime())
	assert.Equal(t, givenRingDuration, call.GetRingDuration())
	assert.Equal(t, givenPlatform, call.GetPlatform())
	assert.Equal(t, givenConferenceId, call.GetConferenceId())
	assert.Equal(t, map[string]string{"key1": givenKey1, "key2": givenKey2}, call.GetCustomData())
}

func TestShouldAddExistingConferenceCall(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING

	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}
	givenConnectOnEarlyMedia := false

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "participants": [{
            "callId": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "state": "%s",
            "joinTime": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "userMuted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            }
        }],
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId)

	expectedRequest := fmt.Sprintf(`{
        "connectOnEarlyMedia": %t
    }`, givenConnectOnEarlyMedia)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf("/calls/1/conferences/%s/call/%s", givenId, givenCallId), givenResponse, 200)

	request := voice.AddExistingCallRequest{
		ConnectOnEarlyMedia: &givenConnectOnEarlyMedia,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.AddExistingConferenceCall(context.Background(), givenId, givenCallId).AddExistingCallRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetParticipants()))
	participant := response.GetParticipants()[0]
	assert.Equal(t, givenCallId, participant.GetCallId())
	assert.Equal(t, givenType, participant.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant.GetState())
	assert.Equal(t, ibTimeJoinTime, participant.GetJoinTime())
	assert.Equal(t, givenAudioMuted, participant.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserMuted, participant.GetMedia().Audio.GetUserMuted())
	assert.Equal(t, givenAudioUserDeaf, participant.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, participant.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant.GetMedia().Video.GetScreenShare())
	assert.Equal(t, givenPlatform, response.GetPlatform())
}

func TestShouldRemoveConferenceCall(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenCallId := "string"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/conferences/%s/call/%s", givenId, givenCallId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.RemoveConferenceCall(context.Background(), givenId, givenCallId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldHangupConference(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenName := "Example conference"
	givenCallId := "string"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenState := voice.PARTICIPANTSTATE_JOINING
	givenJoinTime := "2021-08-24T15:00:00.000+0000"
	givenJoinTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenJoinTime)
	ibTimeJoinTime := infobip.Time{
		T: givenJoinTimeDateTime,
	}
	givenAudioMuted := true
	givenAudioUserMuted := true
	givenAudioUserDeaf := true
	givenVideoCamera := true
	givenVideoScreenShare := true
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		ApplicationId: &givenApplicationId,
		EntityId:      &givenEntityId,
	}

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "participants": [{
            "callId": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "state": "%s",
            "joinTime": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "userMuted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            }
        }],
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        }
    }`, givenId, givenName, givenCallId, givenPhoneNumber, givenType, givenState, givenJoinTime, givenAudioMuted, givenAudioUserMuted, givenAudioUserDeaf, givenVideoCamera, givenVideoScreenShare, givenApplicationId, givenEntityId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/hangup", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.HangupConference(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetParticipants()))
	participant := response.GetParticipants()[0]
	assert.Equal(t, givenCallId, participant.GetCallId())
	assert.Equal(t, givenType, participant.GetEndpoint().PhoneEndpoint.Type)
	assert.Equal(t, givenState, participant.GetState())
	assert.Equal(t, ibTimeJoinTime, participant.GetJoinTime())
	assert.Equal(t, givenAudioMuted, participant.GetMedia().Audio.GetMuted())
	assert.Equal(t, givenAudioUserMuted, participant.GetMedia().Audio.GetUserMuted())
	assert.Equal(t, givenAudioUserDeaf, participant.GetMedia().Audio.GetDeaf())
	assert.Equal(t, givenVideoCamera, participant.GetMedia().Video.GetCamera())
	assert.Equal(t, givenVideoScreenShare, participant.GetMedia().Video.GetScreenShare())
	assert.Equal(t, givenPlatform, response.GetPlatform())
}

func TestShouldConferencePlayFile(t *testing.T) {
	givenId := "123"
	givenLoopCount := int32(1)
	givenContentType := voice.PLAYCONTENTTYPE_FILE
	givenFileId := "100"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "loopCount": %d,
        "content": {
            "type": "%s",
            "fileId": "%s"
        }
    }`, givenLoopCount, givenContentType, givenFileId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/play", givenId), givenResponse, 200)

	request := voice.ConferencePlayRequest{
		LoopCount: &givenLoopCount,
		Content:   voice.FilePlayContentAsPlayContent(voice.NewFilePlayContent(givenFileId)),
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ConferencePlayFile(context.Background(), givenId).ConferencePlayRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldConferenceSayText(t *testing.T) {
	givenId := "123"
	givenText := "string"
	givenLanguage := voice.LANGUAGE_AR
	givenSpeechRate := 0.5
	givenLoopCount := int32(0)
	givenCallsGender := voice.GENDER_FEMALE
	givenVoiceName := voice.VOICENAME_ABIGAIL
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "text": "%s",
        "language": "%s",
        "speechRate": %f,
        "loopCount": %d,
        "preferences": {
            "voiceGender": "%s",
            "voiceName": "%s"
        }
    }`, givenText, givenLanguage, givenSpeechRate, givenLoopCount, givenCallsGender, givenVoiceName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/say", givenId), givenResponse, 200)

	request := voice.SayRequest{
		Text:       givenText,
		Language:   givenLanguage,
		SpeechRate: &givenSpeechRate,
		LoopCount:  &givenLoopCount,
		Preferences: &voice.VoicePreferences{
			VoiceGender: &givenCallsGender,
			VoiceName:   &givenVoiceName,
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ConferenceSayText(context.Background(), givenId).SayRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldConferenceStopPlayingFile(t *testing.T) {
	givenId := "123"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/stop-play", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.ConferenceStopPlayingFile(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldConferenceStartRecording(t *testing.T) {
	givenId := "123"
	givenRecordingType := voice.RECORDINGTYPE_AUDIO
	givenEnabled := false
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/start-recording", givenId), givenResponse, 200)

	request := voice.ConferenceRecordingRequest{
		RecordingType: givenRecordingType,
		ConferenceComposition: &voice.ConferenceComposition{
			Enabled: &givenEnabled,
		},
	}

	response, _, err := infobipClient.CallsAPI.ConferenceStartRecording(context.Background(), givenId).ConferenceRecordingRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldConferenceStopRecording(t *testing.T) {
	givenId := "123"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/stop-recording", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.ConferenceStopRecording(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldConferenceBroadcastWebRtcText(t *testing.T) {
	givenStatus := voice.ACTIONSTATUS_PENDING
	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedText := "This meeting will end in 5 minutes."
	expectedRequest := fmt.Sprintf(`{
        "text": "%s"
    }`, expectedText)

	givenConferenceId := "123"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/conferences/%s/broadcast-webrtc-text", givenConferenceId), givenResponse, 200)

	request := voice.ConferenceBroadcastWebrtcTextRequest{
		Text: &expectedText,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ConferenceBroadcastWebrtcText(context.Background(), givenConferenceId).ConferenceBroadcastWebrtcTextRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldGetCallsFiles(t *testing.T) {
	givenId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenCreationMethod := voice.CREATIONMETHOD_RECORDED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenExpirationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenExpirationTime)
	ibTimeExpirationTime := infobip.Time{
		T: givenExpirationTimeDateTime,
	}

	givenDuration := int64(3)
	givenPage := int32(0)
	givenPageSize := int32(1)
	givenPageTotalPages := int32(0)
	givenPageTotalResults := int64(0)

	givenResponse := fmt.Sprintf(`{
        "results": [{
            "id": "%s",
            "name": "%s",
            "fileFormat": "%s",
            "size": %d,
            "creationMethod": "%s",
            "creationTime": "%s",
            "expirationTime": "%s",
            "duration": %d
        }],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenId, givenName, givenFileFormat, givenSize, givenCreationMethod, givenCreationTime, givenExpirationTime, givenDuration, givenPage, givenPageSize, givenPageTotalPages, givenPageTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/files", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallsFiles(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.GetResults()))
	result := response.GetResults()[0]
	assert.Equal(t, givenId, result.GetId())
	assert.Equal(t, givenName, result.GetName())
	assert.Equal(t, givenFileFormat, result.GetFileFormat())
	assert.Equal(t, givenSize, result.GetSize())
	assert.Equal(t, givenCreationMethod, result.GetCreationMethod())
	assert.Equal(t, ibTimeCreationTime, result.GetCreationTime())
	assert.Equal(t, ibTimeExpirationTime, result.GetExpirationTime())
	assert.Equal(t, givenDuration, result.GetDuration())
	paging := response.GetPaging()
	assert.Equal(t, givenPage, paging.GetPage())
	assert.Equal(t, givenPageSize, paging.GetSize())
	assert.Equal(t, givenPageTotalPages, paging.GetTotalPages())
	assert.Equal(t, givenPageTotalResults, paging.GetTotalResults())
}

func TestShouldUploadCallsAudioFile(t *testing.T) {
	givenId := "string"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenCreationMethod := voice.CREATIONMETHOD_RECORDED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenExpirationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenExpirationTime)
	ibTimeExpirationTime := infobip.Time{
		T: givenExpirationTimeDateTime,
	}

	givenDuration := int64(3)

	givenAttachmentText := "Test file text"
	tempFile, err := ioutil.TempFile("", "attachment.txt")
	assert.Nil(t, err)
	defer os.Remove(tempFile.Name())
	_, err = tempFile.WriteString(givenAttachmentText)
	assert.Nil(t, err)

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "fileFormat": "%s",
        "size": %d,
        "creationMethod": "%s",
        "creationTime": "%s",
        "expirationTime": "%s",
        "duration": %d
    }`, givenId, givenName, givenFileFormat, givenSize, givenCreationMethod, givenCreationTime, givenExpirationTime, givenDuration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/files", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.UploadCallsAudioFile(context.Background()).File(tempFile).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenFileFormat, response.GetFileFormat())
	assert.Equal(t, givenSize, response.GetSize())
	assert.Equal(t, givenCreationMethod, response.GetCreationMethod())
	assert.Equal(t, ibTimeCreationTime, response.GetCreationTime())
	assert.Equal(t, ibTimeExpirationTime, response.GetExpirationTime())
	assert.Equal(t, givenDuration, response.GetDuration())
}

func TestShouldGetCallsFile(t *testing.T) {
	givenId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenCreationMethod := voice.CREATIONMETHOD_RECORDED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenExpirationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenExpirationTime)
	ibTimeExpirationTime := infobip.Time{
		T: givenExpirationTimeDateTime,
	}

	givenDuration := int64(3)

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "fileFormat": "%s",
        "size": %d,
        "creationMethod": "%s",
        "creationTime": "%s",
        "expirationTime": "%s",
        "duration": %d
    }`, givenId, givenName, givenFileFormat, givenSize, givenCreationMethod, givenCreationTime, givenExpirationTime, givenDuration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/files/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallsFile(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenFileFormat, response.GetFileFormat())
	assert.Equal(t, givenSize, response.GetSize())
	assert.Equal(t, givenCreationMethod, response.GetCreationMethod())
	assert.Equal(t, ibTimeCreationTime, response.GetCreationTime())
	assert.Equal(t, ibTimeExpirationTime, response.GetExpirationTime())
	assert.Equal(t, givenDuration, response.GetDuration())
}

func TestShouldDeleteCallsFile(t *testing.T) {
	givenId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenCreationMethod := voice.CREATIONMETHOD_RECORDED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenExpirationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenExpirationTime)
	ibTimeExpirationTime := infobip.Time{
		T: givenExpirationTimeDateTime,
	}
	givenDuration := int64(3)

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "fileFormat": "%s",
        "size": %d,
        "creationMethod": "%s",
        "creationTime": "%s",
        "expirationTime": "%s",
        "duration": %d
    }`, givenId, givenName, givenFileFormat, givenSize, givenCreationMethod, givenCreationTime, givenExpirationTime, givenDuration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/files/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DeleteCallsFile(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenFileFormat, response.GetFileFormat())
	assert.Equal(t, givenSize, response.GetSize())
	assert.Equal(t, givenCreationMethod, response.GetCreationMethod())
	assert.Equal(t, ibTimeCreationTime, response.GetCreationTime())
	assert.Equal(t, ibTimeExpirationTime, response.GetExpirationTime())
	assert.Equal(t, givenDuration, response.GetDuration())
}

func TestShouldGetCallRecordings(t *testing.T) {
	givenId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenSftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenDuration := int64(3)

	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP
	givenStatus := voice.RECORDINGSTATUS_SUCCESSFUL
	givenReason := "string"

	givenResponse := fmt.Sprintf(`{
        "callId": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "direction": "%s",
        "files": [{
            "id": "%s",
            "name": "%s",
            "fileFormat": "%s",
            "size": %d,
            "sftpUploadStatus": "%s",
            "creationTime": "%s",
            "expirationTime": "%s",
            "duration": %d,
            "startTime": "%s",
            "endTime": "%s",
            "location": "%s"
        }],
        "status": "%s",
        "reason": "%s"
    }`, givenId, givenPhoneNumber, givenType, givenDirection, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation, givenStatus, givenReason)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/recordings/calls/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallRecordings(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetCallId())
	endpoint := response.GetEndpoint()
	assert.Equal(t, givenType, endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenDirection, response.GetDirection())
	files := response.GetFiles()
	assert.Equal(t, 1, len(files))
	file := files[0]
	assert.Equal(t, givenFileId, file.GetId())
	assert.Equal(t, givenName, file.GetName())
	assert.Equal(t, givenFileFormat, file.GetFileFormat())
	assert.Equal(t, givenSize, file.GetSize())
	assert.Equal(t, givenSftpUploadStatus, file.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, file.GetCreationTime())
	assert.Equal(t, givenDuration, file.GetDuration())
	assert.Equal(t, ibTimeStartTime, file.GetStartTime())
	assert.Equal(t, ibTimeEndTime, file.GetEndTime())
	assert.Equal(t, givenLocation, file.GetLocation())
	assert.Equal(t, givenStatus, response.GetStatus())
	assert.Equal(t, givenReason, response.GetReason())
}

func TestShouldDeleteCallRecordings(t *testing.T) {
	givenId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenSftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenDuration := int64(3)
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP
	givenStatus := voice.RECORDINGSTATUS_SUCCESSFUL
	givenReason := "string"

	givenResponse := fmt.Sprintf(`{
        "callId": "%s",
        "endpoint": {
            "phoneNumber": "%s",
            "type": "%s"
        },
        "direction": "%s",
        "files": [{
            "id": "%s",
            "name": "%s",
            "fileFormat": "%s",
            "size": %d,
            "sftpUploadStatus": "%s",
            "creationTime": "%s",
            "expirationTime": "%s",
            "duration": %d,
            "startTime": "%s",
            "endTime": "%s",
            "location": "%s"
        }],
        "status": "%s",
        "reason": "%s"
    }`, givenId, givenPhoneNumber, givenType, givenDirection, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation, givenStatus, givenReason)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/recordings/calls/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DeleteCallRecordings(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetCallId())
	endpoint := response.GetEndpoint()
	assert.Equal(t, givenType, endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenDirection, response.GetDirection())
	files := response.GetFiles()
	assert.Equal(t, 1, len(files))
	file := files[0]
	assert.Equal(t, givenFileId, file.GetId())
	assert.Equal(t, givenName, file.GetName())
	assert.Equal(t, givenFileFormat, file.GetFileFormat())
	assert.Equal(t, givenSize, file.GetSize())
	assert.Equal(t, givenSftpUploadStatus, file.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, file.GetCreationTime())
	assert.Equal(t, givenDuration, file.GetDuration())
	assert.Equal(t, ibTimeStartTime, file.GetStartTime())
	assert.Equal(t, ibTimeEndTime, file.GetEndTime())
	assert.Equal(t, givenLocation, file.GetLocation())
	assert.Equal(t, givenStatus, response.GetStatus())
	assert.Equal(t, givenReason, response.GetReason())
}

func TestShouldGetConferencesRecordings(t *testing.T) {
	givenConferenceId := "string"
	givenConferenceName := "string"
	givenApplicationId := "string"
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId,
	}
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenDuration := int64(3)
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP
	givenId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenStatus := voice.RECORDINGSTATUS_SUCCESSFUL
	givenReason := "string"
	givenPage := int32(0)
	givenPageSize := int32(1)
	givenPageTotalPages := int32(0)
	givenPageTotalResults := int64(0)
	givenSftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED

	givenResponse := fmt.Sprintf(`{
        "results": [{
            "conferenceId": "%s",
            "conferenceName": "%s",
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "composedFiles": [{
                "id": "%s",
                "name": "%s",
                "fileFormat": "%s",
                "size": %d,
                "sftpUploadStatus": "%s",
                "creationTime": "%s",
                "expirationTime": "%s",
                "duration": %d,
                "startTime": "%s",
                "endTime": "%s",
                "location": "%s"
            }],
            "callRecordings": [{
                "callId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                },
                "direction": "%s",
                "files": [{
                    "id": "%s",
                    "name": "%s",
                    "fileFormat": "%s",
                    "size": %d,
                    "sftpUploadStatus": "%s",
                    "creationTime": "%s",
                    "expirationTime": "%s",
                    "duration": %d,
                    "startTime": "%s",
                    "endTime": "%s",
                    "location": "%s"
                }],
                "status": "%s",
                "reason": "%s"
            }]
        }],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenConferenceId, givenConferenceName, givenApplicationId, givenEntityId, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation, givenId, givenPhoneNumber, givenType, givenDirection, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation, givenStatus, givenReason, givenPage, givenPageSize, givenPageTotalPages, givenPageTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/recordings/conferences", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetConferencesRecordings(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.GetResults()))
	result := response.GetResults()[0]
	assert.Equal(t, givenConferenceId, result.GetConferenceId())
	assert.Equal(t, givenConferenceName, result.GetConferenceName())
	assert.Equal(t, givenPlatform, result.GetPlatform())
	composedFiles := result.GetComposedFiles()
	assert.Equal(t, 1, len(composedFiles))
	file := composedFiles[0]
	assert.Equal(t, givenFileId, file.GetId())
	assert.Equal(t, givenName, file.GetName())
	assert.Equal(t, givenFileFormat, file.GetFileFormat())
	assert.Equal(t, givenSize, file.GetSize())
	assert.Equal(t, givenSftpUploadStatus, file.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, file.GetCreationTime())
	assert.Equal(t, givenDuration, file.GetDuration())
	assert.Equal(t, ibTimeStartTime, file.GetStartTime())
	assert.Equal(t, ibTimeEndTime, file.GetEndTime())
	assert.Equal(t, givenLocation, file.GetLocation())
	callRecordings := result.GetCallRecordings()
	assert.Equal(t, 1, len(callRecordings))
	callRecording := callRecordings[0]
	assert.Equal(t, givenId, callRecording.GetCallId())
	endpoint := callRecording.GetEndpoint()
	assert.Equal(t, givenType, endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenDirection, callRecording.GetDirection())
	callRecordingsFiles := callRecording.GetFiles()
	assert.Equal(t, 1, len(callRecordingsFiles))
	callRecordingFile := callRecordingsFiles[0]
	assert.Equal(t, givenFileId, callRecordingFile.GetId())
	assert.Equal(t, givenName, callRecordingFile.GetName())
	assert.Equal(t, givenFileFormat, callRecordingFile.GetFileFormat())
	assert.Equal(t, givenSize, callRecordingFile.GetSize())
	assert.Equal(t, givenSftpUploadStatus, callRecordingFile.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, callRecordingFile.GetCreationTime())
	assert.Equal(t, givenDuration, callRecordingFile.GetDuration())
	assert.Equal(t, ibTimeStartTime, callRecordingFile.GetStartTime())
	assert.Equal(t, ibTimeEndTime, callRecordingFile.GetEndTime())
	assert.Equal(t, givenLocation, callRecordingFile.GetLocation())
	assert.Equal(t, givenStatus, callRecording.GetStatus())
	assert.Equal(t, givenReason, callRecording.GetReason())
	paging := response.GetPaging()
	assert.Equal(t, givenPage, paging.GetPage())
	assert.Equal(t, givenPageSize, paging.GetSize())
	assert.Equal(t, givenPageTotalPages, paging.GetTotalPages())
	assert.Equal(t, givenPageTotalResults, paging.GetTotalResults())
}

func TestShouldGetConferenceRecordings(t *testing.T) {
	givenConferenceId := "string"
	givenConferenceName := "string"
	givenApplicationId := "string"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId,
	}
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenSftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenDuration := int64(3)
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP
	givenId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenStatus := voice.RECORDINGSTATUS_SUCCESSFUL
	givenReason := "string"

	givenResponse := fmt.Sprintf(`{
        "conferenceId": "%s",
        "conferenceName": "%s",
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "composedFiles": [{
            "id": "%s",
            "name": "%s",
            "fileFormat": "%s",
            "size": %d,
            "creationTime": "%s",
            "duration": %d,
            "startTime": "%s",
            "endTime": "%s"
        }],
        "callRecordings": [{
            "callId": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "direction": "%s",
            "files": [{
                "id": "%s",
                "name": "%s",
                "fileFormat": "%s",
                "size": %d,
                "sftpUploadStatus": "%s",
                "creationTime": "%s",
                "expirationTime": "%s",
                "duration": %d,
                "startTime": "%s",
                "endTime": "%s",
                "location": "%s"
            }],
            "status": "%s",
            "reason": "%s"
        }]
    }`, givenConferenceId, givenConferenceName, givenApplicationId, givenEntityId, givenFileId, givenName, givenFileFormat, givenSize, givenCreationTime, givenDuration, givenStartTime, givenEndTime, givenId, givenPhoneNumber, givenType, givenDirection, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation, givenStatus, givenReason)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/recordings/conferences/%s", givenConferenceId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetConferenceRecordings(context.Background(), givenConferenceId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenConferenceId, response.GetConferenceId())
	assert.Equal(t, givenConferenceName, response.GetConferenceName())
	assert.Equal(t, givenPlatform, response.GetPlatform())
	composedFiles := response.GetComposedFiles()
	assert.Equal(t, 1, len(composedFiles))
	file := composedFiles[0]
	assert.Equal(t, givenFileId, file.GetId())
	assert.Equal(t, givenName, file.GetName())
	assert.Equal(t, givenFileFormat, file.GetFileFormat())
	assert.Equal(t, givenSize, file.GetSize())
	assert.Equal(t, ibTimeCreationTime, file.GetCreationTime())
	assert.Equal(t, givenDuration, file.GetDuration())
	assert.Equal(t, ibTimeStartTime, file.GetStartTime())
	assert.Equal(t, ibTimeEndTime, file.GetEndTime())
	callRecordings := response.GetCallRecordings()
	assert.Equal(t, 1, len(callRecordings))
	callRecording := callRecordings[0]
	assert.Equal(t, givenId, callRecording.GetCallId())
	endpoint := callRecording.GetEndpoint()
	assert.Equal(t, givenType, endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenDirection, callRecording.GetDirection())
	callRecordingsFiles := callRecording.GetFiles()
	assert.Equal(t, 1, len(callRecordingsFiles))
	callRecordingFile := callRecordingsFiles[0]
	assert.Equal(t, givenFileId, callRecordingFile.GetId())
	assert.Equal(t, givenName, callRecordingFile.GetName())
	assert.Equal(t, givenFileFormat, callRecordingFile.GetFileFormat())
	assert.Equal(t, givenSize, callRecordingFile.GetSize())
	assert.Equal(t, givenSftpUploadStatus, callRecordingFile.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, callRecordingFile.GetCreationTime())
	assert.Equal(t, givenDuration, callRecordingFile.GetDuration())
	assert.Equal(t, ibTimeStartTime, callRecordingFile.GetStartTime())
	assert.Equal(t, ibTimeEndTime, callRecordingFile.GetEndTime())
	assert.Equal(t, givenLocation, callRecordingFile.GetLocation())
	assert.Equal(t, givenStatus, callRecording.GetStatus())
	assert.Equal(t, givenReason, callRecording.GetReason())
}

func TestShouldDeleteConferenceRecordings(t *testing.T) {
	givenConferenceId := "string"
	givenConferenceName := "string"
	givenApplicationId := "string"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId,
	}
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenSftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenDuration := int64(3)
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP
	givenId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenPhoneNumber := "41792030000"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenStatus := voice.RECORDINGSTATUS_SUCCESSFUL
	givenReason := "string"

	givenResponse := fmt.Sprintf(`{
        "conferenceId": "%s",
        "conferenceName": "%s",
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "composedFiles": [{
            "id": "%s",
            "name": "%s",
            "fileFormat": "%s",
            "size": %d,
            "creationTime": "%s",
            "duration": %d,
            "startTime": "%s",
            "endTime": "%s"
        }],
        "callRecordings": [{
            "callId": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "direction": "%s",
            "files": [{
                "id": "%s",
                "name": "%s",
                "fileFormat": "%s",
                "size": %d,
                "sftpUploadStatus": "%s",
                "creationTime": "%s",
                "expirationTime": "%s",
                "duration": %d,
                "startTime": "%s",
                "endTime": "%s",
                "location": "%s"
            }],
            "status": "%s",
            "reason": "%s"
        }]
    }`, givenConferenceId, givenConferenceName, givenApplicationId, givenEntityId, givenFileId, givenName, givenFileFormat, givenSize, givenCreationTime, givenDuration, givenStartTime, givenEndTime, givenId, givenPhoneNumber, givenType, givenDirection, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation, givenStatus, givenReason)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/recordings/conferences/%s", givenConferenceId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DeleteConferenceRecordings(context.Background(), givenConferenceId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenConferenceId, response.GetConferenceId())
	assert.Equal(t, givenConferenceName, response.GetConferenceName())
	assert.Equal(t, givenPlatform, response.GetPlatform())
	composedFiles := response.GetComposedFiles()
	assert.Equal(t, 1, len(composedFiles))
	file := composedFiles[0]
	assert.Equal(t, givenFileId, file.GetId())
	assert.Equal(t, givenName, file.GetName())
	assert.Equal(t, givenFileFormat, file.GetFileFormat())
	assert.Equal(t, givenSize, file.GetSize())
	assert.Equal(t, ibTimeCreationTime, file.GetCreationTime())
	assert.Equal(t, givenDuration, file.GetDuration())
	assert.Equal(t, ibTimeStartTime, file.GetStartTime())
	assert.Equal(t, ibTimeEndTime, file.GetEndTime())
	callRecordings := response.GetCallRecordings()
	assert.Equal(t, 1, len(callRecordings))
	callRecording := callRecordings[0]
	assert.Equal(t, givenId, callRecording.GetCallId())
	endpoint := callRecording.GetEndpoint()
	assert.Equal(t, givenType, endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenDirection, callRecording.GetDirection())
	callRecordingsFiles := callRecording.GetFiles()
	assert.Equal(t, 1, len(callRecordingsFiles))
	callRecordingFile := callRecordingsFiles[0]
	assert.Equal(t, givenFileId, callRecordingFile.GetId())
	assert.Equal(t, givenName, callRecordingFile.GetName())
	assert.Equal(t, givenFileFormat, callRecordingFile.GetFileFormat())
	assert.Equal(t, givenSize, callRecordingFile.GetSize())
	assert.Equal(t, givenSftpUploadStatus, callRecordingFile.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, callRecordingFile.GetCreationTime())
	assert.Equal(t, givenDuration, callRecordingFile.GetDuration())
	assert.Equal(t, ibTimeStartTime, callRecordingFile.GetStartTime())
	assert.Equal(t, ibTimeEndTime, callRecordingFile.GetEndTime())
	assert.Equal(t, givenLocation, callRecordingFile.GetLocation())
	assert.Equal(t, givenStatus, callRecording.GetStatus())
	assert.Equal(t, givenReason, callRecording.GetReason())
}

func TestShouldDownloadCallsRecording(t *testing.T) {
	givenId := "givenId"
	givenResponse := "10101010"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/recordings/files/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DownloadRecordingFile(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestShouldDeleteCallsRecording(t *testing.T) {
	givenFileId := "218eceba-c044-430d-9f26-8f1a7f0g2d03"
	givenName := "Example file"
	givenFileFormat := voice.FILEFORMAT_WAV
	givenSize := int64(292190)
	givenSftpUploadStatus := voice.SFTPUPLOADSTATUS_UPLOADED
	givenCreationTime := "2021-08-24T15:00:00.000+0000"
	givenCreationTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreationTime)
	ibTimeCreationTime := infobip.Time{
		T: givenCreationTimeDateTime,
	}

	givenExpirationTime := "2021-08-24T15:00:00.000+0000"
	givenDuration := int64(3)
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEndTime := "2021-08-24T15:00:00.000+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenLocation := voice.RECORDINGFILELOCATION_SFTP

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "fileFormat": "%s",
        "size": %d,
        "sftpUploadStatus": "%s",
        "creationTime": "%s",
        "expirationTime": "%s",
        "duration": %d,
        "startTime": "%s",
        "endTime": "%s",
        "location": "%s"
    }`, givenFileId, givenName, givenFileFormat, givenSize, givenSftpUploadStatus, givenCreationTime, givenExpirationTime, givenDuration, givenStartTime, givenEndTime, givenLocation)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/recordings/files/%s", givenFileId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DeleteRecordingFile(context.Background(), givenFileId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenFileId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenFileFormat, response.GetFileFormat())
	assert.Equal(t, givenSize, response.GetSize())
	assert.Equal(t, givenSftpUploadStatus, response.GetSftpUploadStatus())
	assert.Equal(t, ibTimeCreationTime, response.GetCreationTime())
	assert.Equal(t, givenDuration, response.GetDuration())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, ibTimeEndTime, response.GetEndTime())
	assert.Equal(t, givenLocation, response.GetLocation())
}

func TestShouldComposeConferenceRecording(t *testing.T) {
	givenConferenceId := "string"
	givenDeleteCallRecordings := true
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "deleteCallRecordings": %t
    }`, givenDeleteCallRecordings)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/recordings/conferences/%s/compose", givenConferenceId), givenResponse, 200)

	request := voice.OnDemandComposition{
		DeleteCallRecordings: &givenDeleteCallRecordings,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.ComposeConferenceRecording(context.Background(), givenConferenceId).OnDemandComposition(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCreateBulk(t *testing.T) {
	givenBulkId := "46ab0413-448f-4153-ada9-b68b14242dc3"
	givenApplicationId1 := "61c060db2675060027d8c7a6"
	givenCallId1 := "7672f60c-3418-40d9-8d6c-d2bac7e9a2b7"
	givenExternalId1 := "7672f60c-3418-40d9-8d6c-d2bac7e9a2b7"
	givenFrom1 := "41793026834"
	givenPhoneNumber1 := "41792036727"
	givenType1 := voice.CALLENDPOINTTYPE_PHONE
	givenStatus1 := voice.ACTIONSTATUS_COMPLETED
	givenReason1 := "Created"
	givenEntityId := "entityId"
	givenRecordingType := "AUDIO"
	givenMachineDetectionEnabled := true
	givenMaxDuration := int32(28000)
	givenConnectTimeout := int32(20000)
	givenMaxCalls := int32(10)
	givenTimeUnit := voice.TIMEUNIT_MINUTES
	givenValidityPeriod := int32(60)
	givenMinWaitPeriod := int32(5)
	givenMaxWaitPeriod := int32(10)
	givenMaxAttempts := int32(5)
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}
	givenDays := "MONDAY"
	givenHourFrom := int32(9)
	givenMinuteFrom := int32(0)
	givenHourTo := int32(17)
	givenMinuteTo := int32(0)
	givenKey2 := "value2"
	givenKey1 := "value1"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId1,
	}
	givenCallsConfigurationId := "dc5942707c704551a00cd2ea"

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "calls": [{
			"callsConfigurationId": "%s",
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "callId": "%s",
            "externalId": "%s",
            "from": "%s",
            "endpoint": {
                "phoneNumber": "%s",
                "type": "%s"
            },
            "status": "%s",
            "reason": "%s"
        }]
    }`, givenBulkId, givenCallsConfigurationId, givenApplicationId1, givenEntityId, givenCallId1, givenExternalId1, givenFrom1, givenPhoneNumber1, givenType1, givenStatus1, givenReason1)

	expectedRequest := fmt.Sprintf(`{
		"callsConfigurationId": "%s",
        "bulkId": "%s",
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "items": [{
            "from": "%s",
            "callRequests": [{
                "externalId": "%s",
                "endpoint": {
                    "phoneNumber": "%s",
                    "type": "%s"
                }
            }],
            "recording": {
                "recordingType": "%s"
            },
            "machineDetection": {
                "enabled": %t
            },
            "maxDuration": %d,
            "connectTimeout": %d,
            "callRate": {
                "maxCalls": %d,
                "timeUnit": "%s"
            },
            "validityPeriod": %d,
            "retryOptions": {
                "minWaitPeriod": %d,
                "maxWaitPeriod": %d,
                "maxAttempts": %d
            },
            "schedulingOptions": {
                "startTime": "%s",
                "callingTimeWindow": {
                    "from": {
                        "hour": %d,
                        "minute": %d
                    },
                    "to": {
                        "hour": %d,
                        "minute": %d
                    },
                    "days": ["%s"]
                }
            },
            "customData": {
                "key2": "%s",
                "key1": "%s"
            }
        }]
    }`, givenCallsConfigurationId, givenBulkId, givenApplicationId1, givenEntityId, givenFrom1, givenExternalId1, givenPhoneNumber1, givenType1, givenRecordingType, givenMachineDetectionEnabled, givenMaxDuration, givenConnectTimeout, givenMaxCalls, givenTimeUnit, givenValidityPeriod, givenMinWaitPeriod, givenMaxWaitPeriod, givenMaxAttempts, givenStartTime, givenHourFrom, givenMinuteFrom, givenHourTo, givenMinuteTo, givenDays, givenKey2, givenKey1)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/bulks", givenResponse, 201)

	request := voice.CallBulkRequest{
		CallsConfigurationId: givenCallsConfigurationId,
		BulkId:               &givenBulkId,
		Platform:             &givenPlatform,
		Items: []voice.BulkItem{{
			From: givenFrom1,
			CallRequests: []voice.BulkCallRequest{{
				Endpoint:   voice.BulkPhoneEndpointAsBulkEndpoint(voice.NewBulkPhoneEndpoint(givenPhoneNumber1)),
				ExternalId: &givenExternalId1,
			}},
			Recording: &voice.CallRecordingRequest{
				RecordingType: voice.RECORDINGTYPE_AUDIO,
			},
			MachineDetection: &voice.MachineDetectionRequest{
				Enabled: givenMachineDetectionEnabled,
			},
			MaxDuration:    &givenMaxDuration,
			ConnectTimeout: &givenConnectTimeout,
			CallRate: &voice.CallRate{
				MaxCalls: &givenMaxCalls,
				TimeUnit: &givenTimeUnit,
			},
			ValidityPeriod: &givenValidityPeriod,
			RetryOptions: &voice.RetryOptions{
				MinWaitPeriod: &givenMinWaitPeriod,
				MaxWaitPeriod: &givenMaxWaitPeriod,
				MaxAttempts:   &givenMaxAttempts,
			},
			SchedulingOptions: &voice.SchedulingOptions{
				StartTime: &ibTimeStartTime,
				CallingTimeWindow: &voice.DeliveryTimeWindow{
					From: &voice.DeliveryTime{
						Hour:   givenHourFrom,
						Minute: givenMinuteFrom,
					},
					To: &voice.DeliveryTime{
						Hour:   givenHourTo,
						Minute: givenMinuteTo,
					},
					Days: []voice.DeliveryDay{voice.DELIVERYDAY_MONDAY},
				},
			},
			CustomData: &map[string]string{
				"key2": givenKey2,
				"key1": givenKey1,
			},
		}},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateBulk(context.Background()).CallBulkRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	calls := response.GetCalls()
	assert.Equal(t, 1, len(calls))
	call1 := calls[0]
	assert.Equal(t, givenPlatform, call1.GetPlatform())
	assert.Equal(t, givenCallId1, call1.GetCallId())
	assert.Equal(t, givenExternalId1, call1.GetExternalId())
	assert.Equal(t, givenFrom1, call1.GetFrom())
	assert.Equal(t, givenStatus1, call1.GetStatus())
	assert.Equal(t, givenReason1, call1.GetReason())
	assert.NotNil(t, call1.GetEndpoint())
}

func TestShouldGetBulkStatus(t *testing.T) {
	givenBulkId := "46ab0413-448f-4153-ada9-b68b14242dc3"
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "startTime": "%s",
        "status": "%s"
    }`, givenBulkId, givenStartTime, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/bulks/%s", givenBulkId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetBulkStatus(context.Background(), givenBulkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldRescheduleBulk(t *testing.T) {
	givenBulkId := "46ab0413-448f-4153-ada9-b68b14242dc3"
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "startTime": "%s",
        "status": "%s"
    }`, givenBulkId, givenStartTime, givenStatus)

	expectedRequest := fmt.Sprintf(`{
        "startTime": "%s"
    }`, givenStartTime)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/bulks/%s/reschedule", givenBulkId), givenResponse, 200)

	request := voice.RescheduleRequest{
		StartTime: ibTimeStartTime,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.RescheduleBulk(context.Background(), givenBulkId).RescheduleRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldPauseBulk(t *testing.T) {
	givenBulkId := "46ab0413-448f-4153-ada9-b68b14242dc3"
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "startTime": "%s",
        "status": "%s"
    }`, givenBulkId, givenStartTime, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/bulks/%s/pause", givenBulkId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.PauseBulk(context.Background(), givenBulkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldResumeBulk(t *testing.T) {
	givenBulkId := "46ab0413-448f-4153-ada9-b68b14242dc3"
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "startTime": "%s",
        "status": "%s"
    }`, givenBulkId, givenStartTime, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/bulks/%s/resume", givenBulkId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.ResumeBulk(context.Background(), givenBulkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCancelBulk(t *testing.T) {
	givenBulkId := "46ab0413-448f-4153-ada9-b68b14242dc3"
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}
	givenStatus := voice.STATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "bulkId": "%s",
        "startTime": "%s",
        "status": "%s"
    }`, givenBulkId, givenStartTime, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/bulks/%s/cancel", givenBulkId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.CancelBulk(context.Background(), givenBulkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenBulkId, response.GetBulkId())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldPostDialogCall(t *testing.T) {
	givenId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenApplicationId := "61c060db2675060027d8c7a6"
	givenState := voice.DIALOGSTATE_ESTABLISHED
	givenStartTime := "2021-08-24T15:00:00.000+0000"
	givenStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenStartTime)
	ibTimeStartTime := infobip.Time{
		T: givenStartTimeDateTime,
	}

	givenEstablishTime := "2022-01-01T00:00:02.100+0000"
	givenEstablishTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEstablishTime)
	ibTimeEstablishTime := infobip.Time{
		T: givenEstablishTimeDateTime,
	}
	givenEndTime := "2022-01-01T00:01:00.100+0000"
	givenEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenEndTime)
	ibTimeEndTime := infobip.Time{
		T: givenEndTimeDateTime,
	}
	givenSecondId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenType := voice.CALLENDPOINTTYPE_PHONE
	givenPhoneNumber := "44790123456"
	givenDirection := voice.CALLDIRECTION_INBOUND
	givenSecondState := voice.CALLSTATE_ESTABLISHED
	givenMuted := false
	givenDeaf := false
	givenCamera := false
	givenScreenShare := false
	givenSecondStartTime := "2022-01-01T00:00:00.100+0000"
	givenSecondStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSecondStartTime)
	ibTimeSecondStartTime := infobip.Time{
		T: givenSecondStartTimeDateTime,
	}

	givenAnswerTime := "2022-01-01T00:00:02.100+0000"
	givenAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenAnswerTime)
	ibTimeAnswerTime := infobip.Time{
		T: givenAnswerTimeDateTime,
	}
	givenSecondEndTime := "2022-01-01T00:01:00.100+0000"
	givenSecondEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSecondEndTime)
	ibTimeSecondEndTime := infobip.Time{
		T: givenSecondEndTimeDateTime,
	}

	givenRingDuration := int32(2)
	givenDialogId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenThirdId := "3ad8805e-d401-424e-9b03-e02a2016a5e2"
	givenSecondType := voice.CALLENDPOINTTYPE_PHONE
	givenSecondPhoneNumber := "44790987654"
	givenSecondDirection := voice.CALLDIRECTION_OUTBOUND
	givenThirdState := voice.CALLSTATE_ESTABLISHED
	givenSecondMuted := false
	givenSecondDeaf := false
	givenSecondCamera := false
	givenThirdScreenShare := false
	givenThirdStartTime := "2022-01-01T00:00:00.100+0000"
	givenThirdStartTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenThirdStartTime)
	ibTimeThirdStartTime := infobip.Time{
		T: givenThirdStartTimeDateTime,
	}
	givenSecondAnswerTime := "2022-01-01T00:00:02.100+0000"
	givenSecondAnswerTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenSecondAnswerTime)
	ibTimeSecondAnswerTime := infobip.Time{
		T: givenSecondAnswerTimeDateTime,
	}
	givenThirdEndTime := "2022-01-01T00:01:00.100+0000"
	givenThirdEndTimeDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenThirdEndTime)
	ibTimeThirdEndTime := infobip.Time{
		T: givenThirdEndTimeDateTime,
	}
	givenSecondRingDuration := int32(2)
	givenSecondDialogId := "034e622a-cc7e-456d-8a10-0ba43b11aa5e"
	givenEntityId := "entityId"
	givenPlatform := voice.Platform{
		EntityId:      &givenEntityId,
		ApplicationId: &givenApplicationId,
	}

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "platform": {
            "applicationId": "%s",
            "entityId": "%s"
        },
        "state": "%s",
        "startTime": "%s",
        "establishTime": "%s",
        "endTime": "%s",
        "parentCall": {
            "id": "%s",
            "endpoint": {
                "type": "%s",
                "phoneNumber": "%s"
            },
            "direction": "%s",
            "state": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            },
            "startTime": "%s",
            "answerTime": "%s",
            "endTime": "%s",
            "ringDuration": %d,
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "dialogId": "%s"
        },
        "childCall": {
            "id": "%s",
            "endpoint": {
                "type": "%s",
                "phoneNumber": "%s"
            },
            "direction": "%s",
            "state": "%s",
            "media": {
                "audio": {
                    "muted": %t,
                    "deaf": %t
                },
                "video": {
                    "camera": %t,
                    "screenShare": %t
                }
            },
            "startTime": "%s",
            "answerTime": "%s",
            "endTime": "%s",
            "ringDuration": %d,
            "platform": {
                "applicationId": "%s",
                "entityId": "%s"
            },
            "dialogId": "%s"
        }
    }`, givenId, givenApplicationId, givenEntityId, givenState, givenStartTime, givenEstablishTime, givenEndTime, givenSecondId, givenType, givenPhoneNumber, givenDirection, givenSecondState, givenMuted, givenDeaf, givenCamera, givenScreenShare, givenSecondStartTime, givenAnswerTime, givenSecondEndTime, givenRingDuration, givenApplicationId, givenEntityId, givenDialogId, givenThirdId, givenSecondType, givenSecondPhoneNumber, givenSecondDirection, givenThirdState, givenSecondMuted, givenSecondDeaf, givenSecondCamera, givenThirdScreenShare, givenThirdStartTime, givenSecondAnswerTime, givenThirdEndTime, givenSecondRingDuration, givenApplicationId, givenEntityId, givenSecondDialogId)

	givenParentCallId := "d8d84155-3831-43fb-91c9-bb897149a79d"
	givenRequestType := voice.CALLENDPOINTTYPE_PHONE
	givenRequestPhoneNumber := "44790987654"
	givenFrom := "44790123456"
	givenConnectTimeout := int32(60)
	givenEnabled := true
	givenKey2 := "value2"
	givenKey1 := "value1"
	givenRecordingType := voice.RECORDINGTYPE_AUDIO
	givenRequestEnabled := false
	givenMaxDuration := int32(3600)
	givenChildCallHangup := false
	givenCallRinging := false
	givenFilePrefix := "filePrefix"
	givenCustomDataField := "custom"
	givenCustomDataFieldValue := "data"
	givenCustomData := map[string]string{givenCustomDataField: givenCustomDataFieldValue}

	expectedRequest := fmt.Sprintf(`{
        "parentCallId": "%s",
        "childCallRequest": {
            "endpoint": {
                "type": "%s",
                "phoneNumber": "%s"
            },
            "from": "%s",
            "connectTimeout": %d,
            "machineDetection": {
                "enabled": %t
            },
            "customData": {
                "key2": "%s",
                "key1": "%s"
            }
        },
        "recording": {
            "recordingType": "%s",
            "dialogComposition": {
                "enabled": %t
            },
            "customData": {
                "%s": "%s"
            },
            "filePrefix": "%s"
        },
        "maxDuration": %d,
        "propagationOptions": {
            "childCallHangup": %t,
            "childCallRinging": %t
        }
    }`, givenParentCallId, givenRequestType, givenRequestPhoneNumber, givenFrom, givenConnectTimeout, givenEnabled, givenKey2, givenKey1, givenRecordingType, givenRequestEnabled, givenCustomDataField, givenCustomDataFieldValue, givenFilePrefix, givenMaxDuration, givenChildCallHangup, givenCallRinging)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/dialogs", givenResponse, 200)

	givenEndpoint := voice.PhoneEndpointAsCallEndpoint(voice.NewPhoneEndpoint(givenRequestPhoneNumber))

	request := voice.DialogRequest{
		ParentCallId: givenParentCallId,
		ChildCallRequest: &voice.DialogCallRequest{
			Endpoint:       &givenEndpoint,
			From:           givenFrom,
			ConnectTimeout: &givenConnectTimeout,
			MachineDetection: &voice.MachineDetectionRequest{
				Enabled: givenEnabled,
			},
			CustomData: &map[string]string{
				"key2": givenKey2,
				"key1": givenKey1,
			},
		},
		Recording: &voice.DialogRecordingRequest{
			RecordingType: givenRecordingType,
			DialogComposition: &voice.DialogRecordingComposition{
				Enabled: &givenRequestEnabled,
			},
			CustomData: &givenCustomData,
			FilePrefix: &givenFilePrefix,
		},
		MaxDuration: &givenMaxDuration,
		PropagationOptions: &voice.DialogPropagationOptions{
			ChildCallHangup:  &givenChildCallHangup,
			ChildCallRinging: &givenCallRinging,
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateDialog(context.Background()).DialogRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenPlatform, response.GetPlatform())
	assert.Equal(t, givenState, response.GetState())
	assert.Equal(t, ibTimeStartTime, response.GetStartTime())
	assert.Equal(t, ibTimeEstablishTime, response.GetEstablishTime())
	assert.Equal(t, ibTimeEndTime, response.GetEndTime())
	assert.NotNil(t, response.GetParentCall())
	assert.Equal(t, &givenSecondId, response.GetParentCall().Id)
	assert.NotNil(t, response.GetParentCall().Endpoint)
	assert.Equal(t, givenType, response.GetParentCall().Endpoint.PhoneEndpoint.Type)
	assert.Equal(t, givenPhoneNumber, response.GetParentCall().Endpoint.PhoneEndpoint.GetPhoneNumber())
	assert.Equal(t, &givenDirection, response.GetParentCall().Direction)
	assert.Equal(t, &givenSecondState, response.GetParentCall().State)
	assert.Equal(t, &voice.MediaProperties{
		Audio: &voice.AudioMediaProperties{
			Muted: &givenMuted,
			Deaf:  &givenDeaf,
		},
		Video: &voice.VideoMediaProperties{
			Camera:      &givenCamera,
			ScreenShare: &givenScreenShare,
		},
	}, response.GetParentCall().Media)
	assert.Equal(t, &ibTimeSecondStartTime, response.GetParentCall().StartTime)
	assert.Equal(t, &ibTimeAnswerTime, response.GetParentCall().AnswerTime)
	assert.Equal(t, &ibTimeSecondEndTime, response.GetParentCall().EndTime)
	assert.Equal(t, &givenRingDuration, response.GetParentCall().RingDuration)
	assert.Equal(t, &givenPlatform, response.GetParentCall().Platform)
	assert.Equal(t, &givenDialogId, response.GetParentCall().DialogId)
	assert.NotNil(t, response.GetChildCall())
	assert.Equal(t, &givenThirdId, response.GetChildCall().Id)
	assert.Equal(t, givenSecondPhoneNumber, response.GetChildCall().Endpoint.PhoneEndpoint.GetPhoneNumber())
	assert.Equal(t, &givenSecondDirection, response.GetChildCall().Direction)
	assert.Equal(t, &givenThirdState, response.GetChildCall().State)
	assert.Equal(t, &voice.MediaProperties{
		Audio: &voice.AudioMediaProperties{
			Muted: &givenSecondMuted,
			Deaf:  &givenSecondDeaf,
		},
		Video: &voice.VideoMediaProperties{
			Camera:      &givenSecondCamera,
			ScreenShare: &givenThirdScreenShare,
		},
	}, response.GetChildCall().Media)
	assert.Equal(t, &ibTimeThirdStartTime, response.GetChildCall().StartTime)
	assert.Equal(t, &ibTimeSecondAnswerTime, response.GetChildCall().AnswerTime)
	assert.Equal(t, &ibTimeThirdEndTime, response.GetChildCall().EndTime)
	assert.Equal(t, &givenSecondRingDuration, response.GetChildCall().RingDuration)
	assert.Equal(t, &givenPlatform, response.GetChildCall().Platform)
	assert.Equal(t, &givenSecondDialogId, response.GetChildCall().DialogId)
}

func TestShouldDialogBroadcastText(t *testing.T) {
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	expectedText := "This dialog will end in 5 minutes."
	expectedRequest := fmt.Sprintf(`{
        "text": "%s"
    }`, expectedText)

	givenDialogId := "123"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/dialogs/%s/broadcast-webrtc-text", givenDialogId), givenResponse, 200)

	callsDialogBroadcastWebrtcTextRequest := voice.DialogBroadcastWebrtcTextRequest{
		Text: &expectedText,
	}

	actualRequest, _ := json.Marshal(callsDialogBroadcastWebrtcTextRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.DialogBroadcastWebrtcText(context.Background(), givenDialogId).DialogBroadcastWebrtcTextRequest(callsDialogBroadcastWebrtcTextRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldCreateSipTrunk(t *testing.T) {
	givenId := "string"
	givenName := "string"
	givenLocation := voice.SIPTRUNKLOCATION_SAO_PAULO
	givenTls := true
	givenCodecs := voice.AUDIOCODEC_PCMU
	givenDtmf := voice.DTMFTYPE_RFC2833
	givenFax := voice.FAXTYPE_NONE
	givenNumberFormat := voice.NUMBERPRESENTATIONFORMAT_E164
	givenInternationalCallsAllowed := true
	givenChannelLimit := int32(0)
	givenAnonymization := voice.ANONYMIZATIONTYPE_NONE
	givenPackageType := voice.BILLINGPACKAGETYPE_METERED
	givenAddressId := "string"
	givenPrimary := "string"
	givenBackup := "string"
	givenType := voice.SIPTRUNKTYPE_STATIC
	givenSourceHosts := "string"
	givenDestinationHosts := "string"
	givenStrategy := voice.SELECTIONSTRATEGY_FAILOVER
	givenEnabled := false

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "location": "%s",
       "tls": %t,
       "codecs": ["%s"],
       "dtmf": "%s",
       "fax": "%s",
       "numberFormat": "%s",
       "internationalCallsAllowed": %t,
       "channelLimit": %d,
       "anonymization": "%s",
       "billingPackage": {
           "packageType": "%s",
           "addressId": "%s"
       },
       "sbcHosts": {
           "primary": ["%s"],
           "backup": ["%s"]
       },
       "type": "%s",
       "sourceHosts": ["%s"],
       "destinationHosts": ["%s"],
       "strategy": "%s",
       "sipOptions": {
           "enabled": %t
       }
   }`, givenId, givenName, givenLocation, givenTls, givenCodecs, givenDtmf, givenFax, givenNumberFormat, givenInternationalCallsAllowed, givenChannelLimit, givenAnonymization, givenPackageType, givenAddressId, givenPrimary, givenBackup, givenType, givenSourceHosts, givenDestinationHosts, givenStrategy, givenEnabled)

	expectedRequest := fmt.Sprintf(`{
       "type": "%s",
       "name": "%s",
       "location": "%s",
       "tls": %t,
       "codecs": ["%s"],
       "dtmf": "%s",
       "fax": "%s",
       "numberFormat": "%s",
       "internationalCallsAllowed": %t,
       "channelLimit": %d,
       "anonymization": "%s",
       "billingPackage": {
           "packageType": "%s",
           "addressId": "%s"
       },
       "sourceHosts": ["%s"],
       "destinationHosts": ["%s"],
       "strategy": "%s",
       "sipOptions": {
           "enabled": %t
       }
   }`, givenType, givenName, givenLocation, givenTls, givenCodecs, givenDtmf, givenFax, givenNumberFormat, givenInternationalCallsAllowed, givenChannelLimit, givenAnonymization, givenPackageType, givenAddressId, givenSourceHosts, givenDestinationHosts, givenStrategy, givenEnabled)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/sip-trunks", givenResponse, 202)

	callsStaticSipTrunkRequest := voice.NewStaticSipTrunkRequest()
	callsStaticSipTrunkRequest.SourceHosts = []string{givenSourceHosts}
	callsStaticSipTrunkRequest.DestinationHosts = []string{givenDestinationHosts}
	callsStaticSipTrunkRequest.Strategy = &givenStrategy
	callsStaticSipTrunkRequest.SipOptions = &voice.SipOptions{Enabled: &givenEnabled}
	callsStaticSipTrunkRequest.Codecs = []voice.AudioCodec{givenCodecs}
	callsStaticSipTrunkRequest.Anonymization = &givenAnonymization
	callsStaticSipTrunkRequest.Dtmf = &givenDtmf
	callsStaticSipTrunkRequest.Fax = &givenFax
	callsStaticSipTrunkRequest.NumberFormat = &givenNumberFormat
	callsStaticSipTrunkRequest.InternationalCallsAllowed = &givenInternationalCallsAllowed
	callsStaticSipTrunkRequest.ChannelLimit = givenChannelLimit
	callsStaticSipTrunkRequest.BillingPackage = voice.BillingPackage{PackageType: givenPackageType, AddressId: &givenAddressId}
	callsStaticSipTrunkRequest.Name = givenName
	callsStaticSipTrunkRequest.Location = &givenLocation
	callsStaticSipTrunkRequest.Tls = &givenTls

	callsSipTrunkRequest := voice.StaticSipTrunkRequestAsSipTrunkRequest(callsStaticSipTrunkRequest)

	actualRequest, _ := json.Marshal(callsSipTrunkRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateSipTrunk(context.Background()).SipTrunkRequest(callsSipTrunkRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.CreateStaticSipTrunkResponse.GetId())
	assert.Equal(t, givenName, response.CreateStaticSipTrunkResponse.GetName())
	assert.Equal(t, givenLocation, response.CreateStaticSipTrunkResponse.GetLocation())
	assert.Equal(t, givenTls, response.CreateStaticSipTrunkResponse.GetTls())
	assert.Equal(t, 1, len(response.CreateStaticSipTrunkResponse.GetCodecs()))
	assert.Equal(t, givenCodecs, response.CreateStaticSipTrunkResponse.GetCodecs()[0])
	assert.Equal(t, givenDtmf, response.CreateStaticSipTrunkResponse.GetDtmf())
	assert.Equal(t, givenFax, response.CreateStaticSipTrunkResponse.GetFax())
	assert.Equal(t, givenNumberFormat, response.CreateStaticSipTrunkResponse.GetNumberFormat())
	assert.Equal(t, givenInternationalCallsAllowed, response.CreateStaticSipTrunkResponse.GetInternationalCallsAllowed())
	assert.Equal(t, givenChannelLimit, response.CreateStaticSipTrunkResponse.GetChannelLimit())
	assert.Equal(t, givenAnonymization, response.CreateStaticSipTrunkResponse.GetAnonymization())
	assert.NotNil(t, response.CreateStaticSipTrunkResponse.GetBillingPackage())
	assert.Equal(t, givenPackageType, response.CreateStaticSipTrunkResponse.GetBillingPackage().PackageType)
	assert.Equal(t, &givenAddressId, response.CreateStaticSipTrunkResponse.GetBillingPackage().AddressId)
	assert.Equal(t, givenPrimary, response.CreateStaticSipTrunkResponse.GetSbcHosts().Primary[0])
	assert.Equal(t, givenBackup, response.CreateStaticSipTrunkResponse.GetSbcHosts().Backup[0])
	assert.Equal(t, givenType, response.CreateStaticSipTrunkResponse.Type)
	assert.Equal(t, givenSourceHosts, response.CreateStaticSipTrunkResponse.GetSourceHosts()[0])
	assert.Equal(t, givenDestinationHosts, response.CreateStaticSipTrunkResponse.GetDestinationHosts()[0])
	assert.Equal(t, givenStrategy, response.CreateStaticSipTrunkResponse.GetStrategy())
	assert.Equal(t, &givenEnabled, response.CreateStaticSipTrunkResponse.GetSipOptions().Enabled)
}

func TestShouldGetSipTrunk(t *testing.T) {
	givenId := "string"
	givenName := "string"
	givenLocation := voice.SIPTRUNKLOCATION_SAO_PAULO
	givenTls := true
	givenCodecs := voice.AUDIOCODEC_PCMU
	givenDtmf := voice.DTMFTYPE_RFC2833
	givenFax := voice.FAXTYPE_NONE
	givenNumberFormat := voice.NUMBERPRESENTATIONFORMAT_E164
	givenInternationalCallsAllowed := true
	givenChannelLimit := int32(0)
	givenAnonymization := voice.ANONYMIZATIONTYPE_NONE
	givenPackageType := voice.BILLINGPACKAGETYPE_METERED
	givenAddressId := "string"
	givenPrimary := "string"
	givenBackup := "string"
	givenType := voice.SIPTRUNKTYPE_STATIC
	givenSourceHosts := "string"
	givenDestinationHosts := "string"
	givenStrategy := voice.SELECTIONSTRATEGY_FAILOVER
	givenEnabled := false

	givenResponse := fmt.Sprintf(`{
      "id": "%s",
      "name": "%s",
      "location": "%s",
      "tls": %t,
      "codecs": ["%s"],
      "dtmf": "%s",
      "fax": "%s",
      "numberFormat": "%s",
      "internationalCallsAllowed": %t,
      "channelLimit": %d,
      "anonymization": "%s",
      "billingPackage": {
          "packageType": "%s",
          "addressId": "%s"
      },
      "sbcHosts": {
          "primary": ["%s"],
          "backup": ["%s"]
      },
      "type": "%s",
      "sourceHosts": ["%s"],
      "destinationHosts": ["%s"],
      "strategy": "%s",
      "sipOptions": {
          "enabled": %t
      }
  }`, givenId, givenName, givenLocation, givenTls, givenCodecs, givenDtmf, givenFax, givenNumberFormat, givenInternationalCallsAllowed, givenChannelLimit, givenAnonymization, givenPackageType, givenAddressId, givenPrimary, givenBackup, givenType, givenSourceHosts, givenDestinationHosts, givenStrategy, givenEnabled)

	givenSipTrunkId := "123"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/sip-trunks/%s", givenSipTrunkId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetSipTrunk(context.Background(), givenSipTrunkId).Execute()
	staticResponse := response.StaticSipTrunkResponse

	assert.Nil(t, err)
	assert.NotNil(t, staticResponse)
	assert.Equal(t, givenId, staticResponse.GetId())
	assert.Equal(t, givenName, staticResponse.GetName())
	assert.Equal(t, givenLocation, staticResponse.GetLocation())
	assert.Equal(t, givenTls, staticResponse.GetTls())
	assert.Equal(t, 1, len(staticResponse.GetCodecs()))
	assert.Equal(t, givenCodecs, staticResponse.GetCodecs()[0])
	assert.Equal(t, givenDtmf, staticResponse.GetDtmf())
	assert.Equal(t, givenFax, staticResponse.GetFax())
	assert.Equal(t, givenNumberFormat, staticResponse.GetNumberFormat())
	assert.Equal(t, givenInternationalCallsAllowed, staticResponse.GetInternationalCallsAllowed())
	assert.Equal(t, givenChannelLimit, staticResponse.GetChannelLimit())
	assert.Equal(t, givenAnonymization, staticResponse.GetAnonymization())
	assert.NotNil(t, staticResponse.GetBillingPackage())
	assert.Equal(t, givenPackageType, staticResponse.GetBillingPackage().PackageType)
	assert.Equal(t, &givenAddressId, staticResponse.GetBillingPackage().AddressId)
	assert.Equal(t, givenPrimary, staticResponse.GetSbcHosts().Primary[0])
	assert.Equal(t, givenBackup, staticResponse.GetSbcHosts().Backup[0])
	assert.Equal(t, givenType, staticResponse.Type)
	assert.Equal(t, givenSourceHosts, staticResponse.GetSourceHosts()[0])
	assert.Equal(t, givenDestinationHosts, staticResponse.GetDestinationHosts()[0])
	assert.Equal(t, givenStrategy, staticResponse.GetStrategy())
	assert.Equal(t, &givenEnabled, staticResponse.GetSipOptions().Enabled)
}

func TestShouldCreateSipTrunkServiceAddress(t *testing.T) {
	givenId := "abc-def-ghi"
	givenName := "Location address name"
	givenStreet := "Location address street"
	givenCity := "My city"
	givenPostCode := "71000"
	givenSuite := "1030"
	givenCountryName := "Croatia"
	givenCountryCode := "HRV"
	givenRegionName := "Zagreb County"
	givenRegionCode := "HR-01"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "street": "%s",
        "city": "%s",
        "postCode": "%s",
        "suite": "%s",
        "country": {
            "name": "%s",
            "code": "%s"
        },
        "region": {
            "name": "%s",
            "code": "%s"
        }
    }`, givenId, givenName, givenStreet, givenCity, givenPostCode, givenSuite, givenCountryName, givenCountryCode, givenRegionName, givenRegionCode)

	expectedRequest := fmt.Sprintf(`{
        "name": "%s",
        "street": "%s",
        "city": "%s",
        "postCode": "%s",
        "suite": "%s",
        "countryCode": "%s",
        "countryRegionCode": "%s"
    }`, givenName, givenStreet, givenCity, givenPostCode, givenSuite, givenCountryCode, givenRegionCode)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/sip-trunks/service-addresses", givenResponse, 201)

	callsSipTrunkServiceAddressRequest := voice.SipTrunkServiceAddressRequest{
		Name:              givenName,
		Street:            givenStreet,
		City:              givenCity,
		PostCode:          givenPostCode,
		Suite:             &givenSuite,
		CountryCode:       givenCountryCode,
		CountryRegionCode: &givenRegionCode,
	}

	actualRequest, _ := json.Marshal(callsSipTrunkServiceAddressRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateSipTrunkServiceAddress(context.Background()).SipTrunkServiceAddressRequest(callsSipTrunkServiceAddressRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenStreet, response.GetStreet())
	assert.Equal(t, givenCity, response.GetCity())
	assert.Equal(t, givenPostCode, response.GetPostCode())
	assert.Equal(t, givenSuite, response.GetSuite())
	assert.NotNil(t, response.GetCountry())
	assert.Equal(t, &givenCountryName, response.GetCountry().Name)
	assert.Equal(t, &givenCountryCode, response.GetCountry().Code)
	assert.NotNil(t, response.GetRegion())
	assert.Equal(t, &givenRegionName, response.GetRegion().Name)
	assert.Equal(t, &givenRegionCode, response.GetRegion().Code)
}

func TestShouldGetSipTrunkServiceAddress(t *testing.T) {
	givenId := "abc-def-ghi"
	givenName := "Location address name"
	givenStreet := "Location address street"
	givenCity := "My city"
	givenPostCode := "71000"
	givenSuite := "1030"
	givenCountryName := "Croatia"
	givenCountryCode := "HRV"
	givenRegionName := "Zagreb County"
	givenRegionCode := "HR-01"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "street": "%s",
        "city": "%s",
        "postCode": "%s",
        "suite": "%s",
        "country": {
            "name": "%s",
            "code": "%s"
        },
        "region": {
            "name": "%s",
            "code": "%s"
        }
    }`, givenId, givenName, givenStreet, givenCity, givenPostCode, givenSuite, givenCountryName, givenCountryCode, givenRegionName, givenRegionCode)

	givenSipTrunkServiceAddressId := "123"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/sip-trunks/service-addresses/%s", givenSipTrunkServiceAddressId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetSipTrunkServiceAddress(context.Background(), givenSipTrunkServiceAddressId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenStreet, response.GetStreet())
	assert.Equal(t, givenCity, response.GetCity())
	assert.Equal(t, givenPostCode, response.GetPostCode())
	assert.Equal(t, givenSuite, response.GetSuite())
	assert.NotNil(t, response.GetCountry())
	assert.Equal(t, &givenCountryName, response.GetCountry().Name)
	assert.Equal(t, &givenCountryCode, response.GetCountry().Code)
	assert.NotNil(t, response.GetRegion())
	assert.Equal(t, &givenRegionName, response.GetRegion().Name)
	assert.Equal(t, &givenRegionCode, response.GetRegion().Code)
}

func TestShouldUpdateSipTrunkServiceAddress(t *testing.T) {
	givenId := "abc-def-ghi"
	givenName := "Location address name"
	givenStreet := "Location address street"
	givenCity := "My city"
	givenPostCode := "71000"
	givenSuite := "1030"
	givenCountryName := "Croatia"
	givenCountryCode := "HRV"
	givenRegionName := "Zagreb County"
	givenRegionCode := "HR-01"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "street": "%s",
        "city": "%s",
        "postCode": "%s",
        "suite": "%s",
        "country": {
            "name": "%s",
            "code": "%s"
        },
        "region": {
            "name": "%s",
            "code": "%s"
        }
    }`, givenId, givenName, givenStreet, givenCity, givenPostCode, givenSuite, givenCountryName, givenCountryCode, givenRegionName, givenRegionCode)

	expectedRequest := fmt.Sprintf(`{
        "name": "%s",
        "street": "%s",
        "city": "%s",
        "postCode": "%s",
        "suite": "%s",
        "countryCode": "%s",
        "countryRegionCode": "%s"
    }`, givenName, givenStreet, givenCity, givenPostCode, givenSuite, givenCountryCode, givenRegionCode)

	givenSipTrunkServiceAddressId := "123"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf("/calls/1/sip-trunks/service-addresses/%s", givenSipTrunkServiceAddressId), givenResponse, 200)

	callsSipTrunkServiceAddressRequest := voice.SipTrunkServiceAddressRequest{
		Name:              givenName,
		Street:            givenStreet,
		City:              givenCity,
		PostCode:          givenPostCode,
		Suite:             &givenSuite,
		CountryCode:       givenCountryCode,
		CountryRegionCode: &givenRegionCode,
	}

	actualRequest, _ := json.Marshal(callsSipTrunkServiceAddressRequest)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.UpdateSipTrunkServiceAddress(context.Background(), givenSipTrunkServiceAddressId).SipTrunkServiceAddressRequest(callsSipTrunkServiceAddressRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenStreet, response.GetStreet())
	assert.Equal(t, givenCity, response.GetCity())
	assert.Equal(t, givenPostCode, response.GetPostCode())
	assert.Equal(t, givenSuite, response.GetSuite())
	assert.NotNil(t, response.GetCountry())
	assert.Equal(t, &givenCountryName, response.GetCountry().Name)
	assert.Equal(t, &givenCountryCode, response.GetCountry().Code)
	assert.NotNil(t, response.GetRegion())
	assert.Equal(t, &givenRegionName, response.GetRegion().Name)
	assert.Equal(t, &givenRegionCode, response.GetRegion().Code)
}

func TestShouldDeleteSipTrunkServiceAddress(t *testing.T) {
	givenId := "abc-def-ghi"
	givenName := "Location address name"
	givenStreet := "Location address street"
	givenCity := "My city"
	givenPostCode := "71000"
	givenSuite := "1030"
	givenCountryName := "Croatia"
	givenCountryCode := "HRV"
	givenRegionName := "Zagreb County"
	givenRegionCode := "HR-01"

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "street": "%s",
        "city": "%s",
        "postCode": "%s",
        "suite": "%s",
        "country": {
            "name": "%s",
            "code": "%s"
        },
        "region": {
            "name": "%s",
            "code": "%s"
        }
    }`, givenId, givenName, givenStreet, givenCity, givenPostCode, givenSuite, givenCountryName, givenCountryCode, givenRegionName, givenRegionCode)

	givenSipTrunkServiceAddressId := "123"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/sip-trunks/service-addresses/%s", givenSipTrunkServiceAddressId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DeleteSipTrunkServiceAddress(context.Background(), givenSipTrunkServiceAddressId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenStreet, response.GetStreet())
	assert.Equal(t, givenCity, response.GetCity())
	assert.Equal(t, givenPostCode, response.GetPostCode())
	assert.Equal(t, givenSuite, response.GetSuite())
	assert.NotNil(t, response.GetCountry())
	assert.Equal(t, &givenCountryName, response.GetCountry().Name)
	assert.Equal(t, &givenCountryCode, response.GetCountry().Code)
	assert.NotNil(t, response.GetRegion())
	assert.Equal(t, &givenRegionName, response.GetRegion().Name)
	assert.Equal(t, &givenRegionCode, response.GetRegion().Code)
}

func TestShouldGetCountries(t *testing.T) {
	givenName1 := "New Zealand"
	givenCode1 := "NZL"
	givenName2 := "Fiji"
	givenCode2 := "FJI"
	givenName3 := "Guadeloupe"
	givenCode3 := "GLP"

	givenResponse := fmt.Sprintf(`{
        "countries": [
            {"name": "%s", "code": "%s"},
            {"name": "%s", "code": "%s"},
            {"name": "%s", "code": "%s"}
        ]
    }`, givenName1, givenCode1, givenName2, givenCode2, givenName3, givenCode3)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/sip-trunks/service-addresses/countries", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCountries(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 3, len(response.GetCountries()))
	assert.Equal(t, givenName1, response.GetCountries()[0].GetName())
	assert.Equal(t, givenCode1, response.GetCountries()[0].GetCode())
	assert.Equal(t, givenName2, response.GetCountries()[1].GetName())
	assert.Equal(t, givenCode2, response.GetCountries()[1].GetCode())
	assert.Equal(t, givenName3, response.GetCountries()[2].GetName())
	assert.Equal(t, givenCode3, response.GetCountries()[2].GetCode())
}

func TestShouldGetRegions(t *testing.T) {
	givenName1 := "Dubrovnik-Neretva County"
	givenCode1 := "HR-19"
	givenCountryCode1 := "HRV"
	givenName2 := "Međimurje County"
	givenCode2 := "HR-20"
	givenCountryCode2 := "HRV"
	givenName3 := "City of Zagreb"
	givenCode3 := "HR-21"
	givenCountryCode3 := "HRV"

	givenResponse := fmt.Sprintf(`{
        "regions": [
            {"name": "%s", "code": "%s", "countryCode": "%s"},
            {"name": "%s", "code": "%s", "countryCode": "%s"},
            {"name": "%s", "code": "%s", "countryCode": "%s"}
        ]
    }`, givenName1, givenCode1, givenCountryCode1, givenName2, givenCode2, givenCountryCode2, givenName3, givenCode3, givenCountryCode3)

	givenCountryCode := "HRV"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/sip-trunks/service-addresses/countries/regions?countryCode=%s", givenCountryCode), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetRegions(context.Background()).CountryCode(givenCountryCode).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 3, len(response.GetRegions()))
	assert.Equal(t, givenName1, response.GetRegions()[0].GetName())
	assert.Equal(t, givenCode1, response.GetRegions()[0].GetCode())
	assert.Equal(t, givenCountryCode1, response.GetRegions()[0].GetCountryCode())
	assert.Equal(t, givenName2, response.GetRegions()[1].GetName())
	assert.Equal(t, givenCode2, response.GetRegions()[1].GetCode())
	assert.Equal(t, givenCountryCode2, response.GetRegions()[1].GetCountryCode())
	assert.Equal(t, givenName3, response.GetRegions()[2].GetName())
	assert.Equal(t, givenCode3, response.GetRegions()[2].GetCode())
	assert.Equal(t, givenCountryCode3, response.GetRegions()[2].GetCountryCode())
}

func TestShouldResetSipTrunkPassword(t *testing.T) {
	sipTrunkId := "426c8402-691c-11ee-8c99-0242ac120002"
	givenUsername := "426c8402-691c-11ee-8c99-0242ac120002"
	givenPassword := "fkZ1921tM87"

	givenResponse := fmt.Sprintf(`{
        "username": "%s",
        "password": "%s"
    }`, givenUsername, givenPassword)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/sip-trunks/%s/reset-password", sipTrunkId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.ResetSipTrunkPassword(context.Background(), sipTrunkId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenUsername, response.GetUsername())
	assert.Equal(t, givenPassword, response.GetPassword())
}

func TestShouldStartTranscription(t *testing.T) {
	callId := "12345"
	givenLanguage := voice.TRANSCRIPTIONLANGUAGE_ES_AR
	givenSendInterimResults := false
	givenStatus := voice.ACTIONSTATUS_PENDING

	expectedRequest := fmt.Sprintf(`{
        "transcription": {
            "language": "%s",
            "sendInterimResults": %t
        }
    }`, givenLanguage, givenSendInterimResults)

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/start-transcription", callId), givenResponse, 200)

	request := voice.StartTranscriptionRequest{
		Transcription: &voice.Transcription{
			Language:           givenLanguage,
			SendInterimResults: &givenSendInterimResults,
		},
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CallStartTranscription(context.Background(), callId).StartTranscriptionRequest(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldStopTranscription(t *testing.T) {
	callId := "12345"
	givenStatus := voice.ACTIONSTATUS_PENDING

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf("/calls/1/calls/%s/stop-transcription", callId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.CallStopTranscription(context.Background(), callId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldGetCallsConfigurations(t *testing.T) {
	givenId := "63467c6e2885a5389ba11d80"
	givenName := "Calls configuration"
	givenPage := int32(0)
	givenSize := int32(1)
	givenTotalPages := int32(1)
	givenTotalResults := int64(1)

	givenResponse := fmt.Sprintf(`{
		"results": [
			{
				"id": "%s",
				"name": "%s"
			}
		],
		"paging": {
			"page": %d,
			"size": %d,
			"totalPages": %d,
			"totalResults": %d
		}
	}`, givenId, givenName, givenPage, givenSize, givenTotalPages, givenTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/calls/1/configurations", givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallsConfigurations(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.GetResults()))
	result := response.GetResults()[0]
	assert.Equal(t, givenId, result.GetId())
	assert.Equal(t, givenName, result.GetName())
	paging := response.GetPaging()
	assert.Equal(t, givenPage, paging.GetPage())
	assert.Equal(t, givenSize, paging.GetSize())
	assert.Equal(t, givenTotalPages, paging.GetTotalPages())
	assert.Equal(t, givenTotalResults, paging.GetTotalResults())
}

func TestShouldCreateCallsConfiguration(t *testing.T) {
	givenId := "63467c6e2885a5389ba11d80"
	givenName := "Calls configuration"

	givenRequest := voice.CallsConfigurationCreateRequest{
		Id:   &givenId,
		Name: &givenName,
	}

	givenResponse := fmt.Sprintf(`{
		"id": "%s",
		"name": "%s"
	}`, givenId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/calls/1/configurations", givenResponse, 201)

	actualRequest, _ := json.Marshal(givenRequest)
	ValidateExpectedRequestBodiesMatches(t, givenResponse, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.CreateCallsConfiguration(context.Background()).CallsConfigurationCreateRequest(givenRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
}

func TestShouldGetCallsConfiguration(t *testing.T) {
	givenId := "63467c6e2885a5389ba11d80"
	givenName := "Calls configuration"

	givenResponse := fmt.Sprintf(`{
		"id": "%s",
		"name": "%s"
	}`, givenId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/calls/1/configurations/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.GetCallsConfiguration(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
}

func TestShouldUpdateCallsConfiguration(t *testing.T) {
	givenId := "63467c6e2885a5389ba11d80"
	givenName := "Calls configuration"

	givenRequest := voice.CallsConfigurationUpdateRequest{
		Name: &givenName,
	}

	givenResponse := fmt.Sprintf(`{
		"id": "%s",
		"name": "%s"
	}`, givenId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf("/calls/1/configurations/%s", givenId), givenResponse, 200)

	actualRequest, _ := json.Marshal(givenRequest)
	expectedRequest := fmt.Sprintf(`{
		"name": "%s"
	}`, givenName)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.CallsAPI.UpdateCallsConfiguration(context.Background(), givenId).CallsConfigurationUpdateRequest(givenRequest).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
}

func TestShouldDeleteCallsConfiguration(t *testing.T) {
	givenId := "63467c6e2885a5389ba11d80"
	givenName := "Calls configuration"

	givenResponse := fmt.Sprintf(`{
		"id": "%s",
		"name": "%s"
	}`, givenId, givenName)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/calls/1/configurations/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.CallsAPI.DeleteCallsConfiguration(context.Background(), givenId).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
}

func TestShouldDownloadRecordingFile(t *testing.T) {
	givenFileId := "3C67336FA555A606C85FA9637906A6AB98436B7AFC65D857A416F6521D39F8F0E1D3D2469FF580D8968D3DD89A2DB561"
	givenBinaryData := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	SetUpGetRequestOctet(fmt.Sprintf("/calls/1/recordings/files/%s", givenFileId), givenBinaryData, 200)

	file, _, err := infobipClient.CallsAPI.DownloadRecordingFile(context.Background(), givenFileId).Execute()

	assert.Nil(t, err, "Expected no error while executing the request")
	assert.NotNil(t, file, "Expected a valid file in the response")

	defer file.Close()
	content, err := io.ReadAll(file)
	assert.Nil(t, err, "Expected no error while reading the file content")
	assert.Equal(t, givenBinaryData, content, "File content does not match the expected binary data")
}
