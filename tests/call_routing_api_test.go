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

func TestShouldGetCallRoutes(t *testing.T) {
	givenId := "f8fc8aca-786d-4943-9af2-e7ec01b5e80d"
	givenName := "SIP endpoint route"
	givenDestinationsValueUsername := "41793026834"
	givenDestinationsValueSipTrunkId := "string"
	givenDestinationsValueCustomHeadersString := "string"
	givenDestinationsValueType := voice.CALLROUTINGENDPOINTTYPE_SIP
	givenDestinationsConnectTimeout := int32(30)
	givenDestinationsRecordingRecordingType := voice.CALLROUTINGRECORDINGTYPE_AUDIO
	givenDestinationsRecordingRecordingCompositionEnabled := true
	givenDestinationsRecordingCustomDataString := "string"
	givenDestinationsRecordingFilePrefix := "string"
	givenDestinationsType := voice.CALLROUTINGDESTINATIONTYPE_ENDPOINT

	givenSecondId := "f8fc8aca-786d-4943-9af2-e7ec01b5e80d"
	givenSecondName := "Phone endpoint route"
	givenSecondDestinationsValuePhoneNumber := "41793026834"
	givenSecondDestinationsValueType := voice.CALLROUTINGENDPOINTTYPE_PHONE
	givenSecondDestinationsConnectTimeout := int32(30)
	givenSecondDestinationsRecordingRecordingType := voice.CALLROUTINGRECORDINGTYPE_AUDIO
	givenSecondDestinationsRecordingRecordingCompositionEnabled := true
	givenSecondDestinationsRecordingCustomDataString := "string"
	givenSecondDestinationsRecordingFilePrefix := "string"
	givenSecondDestinationsType := voice.CALLROUTINGDESTINATIONTYPE_ENDPOINT

	givenPagingPage := int32(0)
	givenPagingSize := int32(20)
	givenPagingTotalPages := int32(1)
	givenPagingTotalResults := int64(2)

	givenResponse := fmt.Sprintf(`{
        "results": [
            {
                "id": "%s",
                "name": "%s",
                "destinations": [
                    {
                        "value": {
                            "username": "%s",
                            "sipTrunkId": "%s",
                            "customHeaders": {
                                "string": "%s"
                            },
                            "type": "%s"
                        },
                        "connectTimeout": %d,
                        "recording": {
                            "recordingType": "%s",
                            "recordingComposition": {
                                "enabled": %t
                            },
                            "customData": {
                                "string": "%s"
                            },
                            "filePrefix": "%s"
                        },
                        "type": "%s"
                    }
                ]
            },
            {
                "id": "%s",
                "name": "%s",
                "destinations": [
                    {
                        "value": {
                            "phoneNumber": "%s",
                            "type": "%s"
                        },
                        "connectTimeout": %d,
                        "recording": {
                            "recordingType": "%s",
                            "recordingComposition": {
                                "enabled": %t
                            },
                            "customData": {
                                "string": "%s"
                            },
                            "filePrefix": "%s"
                        },
                        "type": "%s"
                    }
                ]
            }
        ],
        "paging": {
            "page": %d,
            "size": %d,
            "totalPages": %d,
            "totalResults": %d
        }
    }`, givenId, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType, givenSecondId, givenSecondName, givenSecondDestinationsValuePhoneNumber, givenSecondDestinationsValueType, givenSecondDestinationsConnectTimeout, givenSecondDestinationsRecordingRecordingType, givenSecondDestinationsRecordingRecordingCompositionEnabled, givenSecondDestinationsRecordingCustomDataString, givenSecondDestinationsRecordingFilePrefix, givenSecondDestinationsType, givenPagingPage, givenPagingSize, givenPagingTotalPages, givenPagingTotalResults)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/callrouting/1/routes", givenResponse, 200)

	response, _, err := infobipClient.
		CallRoutingAPI.
		GetCallRoutes(context.Background()).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 2, len(response.GetResults()))

	firstResult := response.GetResults()[0]
	assert.Equal(t, givenId, firstResult.GetId())
	assert.Equal(t, givenName, firstResult.GetName())
	assert.Equal(t, 1, len(firstResult.GetDestinations()))

	firstDestination := firstResult.GetDestinations()[0].CallRoutingEndpointDestination
	firstDestinationValue := firstDestination.GetValue().CallRoutingSipEndpoint
	assert.Equal(t, givenDestinationsValueUsername, firstDestinationValue.GetUsername())
	assert.Equal(t, givenDestinationsValueSipTrunkId, firstDestinationValue.GetSipTrunkId())
	assert.Equal(t, givenDestinationsValueCustomHeadersString, firstDestinationValue.GetCustomHeaders()["string"])
	assert.Equal(t, givenDestinationsValueType, firstDestinationValue.Type)
	assert.Equal(t, givenDestinationsConnectTimeout, firstDestination.GetConnectTimeout())
	assert.Equal(t, givenDestinationsRecordingRecordingType, firstDestination.GetRecording().RecordingType)
	assert.Equal(t, givenDestinationsRecordingRecordingCompositionEnabled, firstDestination.GetRecording().RecordingComposition.GetEnabled())
	assert.Equal(t, &givenDestinationsRecordingFilePrefix, firstDestination.GetRecording().FilePrefix)
	assert.Equal(t, givenDestinationsType, firstDestination.Type)

	secondResult := response.GetResults()[1]
	assert.Equal(t, givenSecondId, secondResult.GetId())
	assert.Equal(t, givenSecondName, secondResult.GetName())
	assert.Equal(t, 1, len(secondResult.GetDestinations()))
	secondDestination := secondResult.GetDestinations()[0].CallRoutingEndpointDestination
	secondDestinationValue := secondDestination.GetValue().CallRoutingPhoneEndpoint
	assert.Equal(t, givenSecondDestinationsValuePhoneNumber, secondDestinationValue.GetPhoneNumber())
	assert.Equal(t, givenSecondDestinationsValueType, secondDestinationValue.Type)
	assert.Equal(t, givenSecondDestinationsConnectTimeout, secondDestination.GetConnectTimeout())
	assert.Equal(t, givenSecondDestinationsRecordingRecordingType, secondDestination.GetRecording().RecordingType)
	assert.Equal(t, givenSecondDestinationsRecordingRecordingCompositionEnabled, secondDestination.GetRecording().RecordingComposition.GetEnabled())
	assert.Equal(t, &givenSecondDestinationsRecordingFilePrefix, secondDestination.GetRecording().FilePrefix)
	assert.Equal(t, givenSecondDestinationsType, secondDestination.Type)

	paging := response.GetPaging()
	assert.Equal(t, givenPagingPage, paging.GetPage())
	assert.Equal(t, givenPagingSize, paging.GetSize())
	assert.Equal(t, givenPagingTotalPages, paging.GetTotalPages())
	assert.Equal(t, givenPagingTotalResults, paging.GetTotalResults())
}

func TestShouldPostCallRoute(t *testing.T) {
	givenId := "f8fc8aca-786d-4943-9af2-e7ec01b5e80d"
	givenName := "SIP endpoint route"
	givenDestinationsValueUsername := "41793026834"
	givenDestinationsValueSipTrunkId := "string"
	givenDestinationsValueCustomHeadersString := "string"
	givenDestinationsValueType := voice.CALLROUTINGENDPOINTTYPE_SIP
	givenDestinationsConnectTimeout := int32(30)
	givenDestinationsRecordingRecordingType := voice.CALLROUTINGRECORDINGTYPE_AUDIO
	givenDestinationsRecordingRecordingCompositionEnabled := true
	givenDestinationsRecordingCustomDataString := "string"
	givenDestinationsRecordingFilePrefix := "string"
	givenDestinationsType := voice.CALLROUTINGDESTINATIONTYPE_ENDPOINT

	expectedRequest := fmt.Sprintf(`{
       "name": "%s",
       "destinations": [
           {
               "value": {
                   "username": "%s",
                   "sipTrunkId": "%s",
                   "customHeaders": {
                       "string": "%s"
                   },
                   "type": "%s"
               },
               "connectTimeout": %d,
               "recording": {
                   "recordingType": "%s",
                   "recordingComposition": {
                       "enabled": %t
                   },
                   "customData": {
                       "string": "%s"
                   },
                   "filePrefix": "%s"
               },
               "type": "%s"
           }
       ]
   }`, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType)

	sipEndpoint := voice.NewCallRoutingSipEndpoint(givenDestinationsValueSipTrunkId)
	sipEndpoint.SetUsername(givenDestinationsValueUsername)
	sipEndpoint.SetCustomHeaders(map[string]string{"string": givenDestinationsValueCustomHeadersString})
	callRoutingEndpoint := voice.CallRoutingSipEndpointAsCallRoutingEndpoint(sipEndpoint)
	destination := voice.NewCallRoutingEndpointDestination(callRoutingEndpoint)

	destination.SetConnectTimeout(givenDestinationsConnectTimeout)
	recording := voice.NewCallRoutingRecording(givenDestinationsRecordingRecordingType)
	recordingComposition := voice.NewCallRoutingRecordingComposition(givenDestinationsRecordingRecordingCompositionEnabled)
	recording.SetRecordingComposition(*recordingComposition)
	recording.SetCustomData(map[string]string{"string": givenDestinationsRecordingCustomDataString})
	recording.SetFilePrefix(givenDestinationsRecordingFilePrefix)
	destination.SetRecording(*recording)
	destinations := []voice.CallRoutingDestination{voice.CallRoutingEndpointDestinationAsCallRoutingDestination(destination)}
	request := voice.NewCallRoutingRouteRequest(givenName, destinations)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "destinations": [
           {
               "value": {
                   "username": "%s",
                   "sipTrunkId": "%s",
                   "customHeaders": {
                       "string": "%s"
                   },
                   "type": "%s"
               },
               "connectTimeout": %d,
               "recording": {
                   "recordingType": "%s",
                   "recordingComposition": {
                       "enabled": %t
                   },
                   "customData": {
                       "string": "%s"
                   },
                   "filePrefix": "%s"
               },
               "type": "%s"
           }
       ]
   }`, givenId, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/callrouting/1/routes", givenResponse, 200)

	response, _, err := infobipClient.
		CallRoutingAPI.
		CreateCallRoute(context.Background()).
		CallRoutingRouteRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetDestinations()))
	destinationResponse := response.GetDestinations()[0].CallRoutingEndpointDestination
	destinationResponseValue := destinationResponse.GetValue().CallRoutingSipEndpoint
	assert.Equal(t, givenDestinationsValueUsername, destinationResponseValue.GetUsername())
	assert.Equal(t, givenDestinationsValueSipTrunkId, destinationResponseValue.GetSipTrunkId())
	assert.Equal(t, givenDestinationsValueCustomHeadersString, destinationResponseValue.GetCustomHeaders()["string"])
	assert.Equal(t, givenDestinationsValueType, destinationResponseValue.Type)
	assert.Equal(t, givenDestinationsConnectTimeout, destinationResponse.GetConnectTimeout())
	assert.Equal(t, givenDestinationsRecordingRecordingType, destinationResponse.GetRecording().RecordingType)
	assert.Equal(t, givenDestinationsRecordingRecordingCompositionEnabled, destinationResponse.GetRecording().RecordingComposition.GetEnabled())
	assert.Equal(t, &givenDestinationsRecordingFilePrefix, destinationResponse.GetRecording().FilePrefix)
	assert.Equal(t, givenDestinationsType, destinationResponse.Type)
}

func TestShouldGetCallRoute(t *testing.T) {
	givenId := "f8fc8aca-786d-4943-9af2-e7ec01b5e80d"
	givenName := "SIP endpoint route"
	givenDestinationsValueUsername := "41793026834"
	givenDestinationsValueSipTrunkId := "string"
	givenDestinationsValueCustomHeadersString := "string"
	givenDestinationsValueType := voice.CALLROUTINGENDPOINTTYPE_SIP
	givenDestinationsConnectTimeout := int32(30)
	givenDestinationsRecordingRecordingType := voice.CALLROUTINGRECORDINGTYPE_AUDIO
	givenDestinationsRecordingRecordingCompositionEnabled := true
	givenDestinationsRecordingCustomDataString := "string"
	givenDestinationsRecordingFilePrefix := "string"
	givenDestinationsType := voice.CALLROUTINGDESTINATIONTYPE_ENDPOINT

	givenResponse := fmt.Sprintf(`{
       "id": "%s",
       "name": "%s",
       "destinations": [
           {
               "value": {
                   "username": "%s",
                   "sipTrunkId": "%s",
                   "customHeaders": {
                       "string": "%s"
                   },
                   "type": "%s"
               },
               "connectTimeout": %d,
               "recording": {
                   "recordingType": "%s",
                   "recordingComposition": {
                       "enabled": %t
                   },
                   "customData": {
                       "string": "%s"
                   },
                   "filePrefix": "%s"
               },
               "type": "%s"
           }
       ]
   }`, givenId, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf("/callrouting/1/routes/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.
		CallRoutingAPI.
		GetCallRoute(context.Background(), givenId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetDestinations()))
	destinationResponse := response.GetDestinations()[0].CallRoutingEndpointDestination
	destinationResponseValue := destinationResponse.GetValue().CallRoutingSipEndpoint
	assert.Equal(t, givenDestinationsValueUsername, destinationResponseValue.GetUsername())
	assert.Equal(t, givenDestinationsValueSipTrunkId, destinationResponseValue.GetSipTrunkId())
	assert.Equal(t, givenDestinationsValueCustomHeadersString, destinationResponseValue.GetCustomHeaders()["string"])
	assert.Equal(t, givenDestinationsValueType, destinationResponseValue.Type)
	assert.Equal(t, givenDestinationsConnectTimeout, destinationResponse.GetConnectTimeout())
	assert.Equal(t, givenDestinationsRecordingRecordingType, destinationResponse.GetRecording().RecordingType)
	assert.Equal(t, givenDestinationsRecordingRecordingCompositionEnabled, destinationResponse.GetRecording().RecordingComposition.GetEnabled())
	assert.Equal(t, &givenDestinationsRecordingFilePrefix, destinationResponse.GetRecording().FilePrefix)
	assert.Equal(t, givenDestinationsType, destinationResponse.Type)
}

func TestShouldPutCallRoute(t *testing.T) {
	givenId := "f8fc8aca-786d-4943-9af2-e7ec01b5e80d"
	givenName := "SIP endpoint route"
	givenDestinationsValueUsername := "41793026834"
	givenDestinationsValueSipTrunkId := "string"
	givenDestinationsValueCustomHeadersString := "string"
	givenDestinationsValueType := voice.CALLROUTINGENDPOINTTYPE_SIP
	givenDestinationsConnectTimeout := int32(30)
	givenDestinationsRecordingRecordingType := voice.CALLROUTINGRECORDINGTYPE_AUDIO
	givenDestinationsRecordingRecordingCompositionEnabled := true
	givenDestinationsRecordingCustomDataString := "string"
	givenDestinationsRecordingFilePrefix := "string"
	givenDestinationsType := voice.CALLROUTINGDESTINATIONTYPE_ENDPOINT

	expectedRequest := fmt.Sprintf(`{
        "name": "%s",
        "destinations": [
            {
                "value": {
                    "username": "%s",
                    "sipTrunkId": "%s",
                    "customHeaders": {
                        "string": "%s"
                    },
                    "type": "%s"
                },
                "connectTimeout": %d,
                "recording": {
                    "recordingType": "%s",
                    "recordingComposition": {
                        "enabled": %t
                    },
                    "customData": {
                        "string": "%s"
                    },
                    "filePrefix": "%s"
                },
                "type": "%s"
            }
        ]
    }`, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType)

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "destinations": [
            {
                "value": {
                    "username": "%s",
                    "sipTrunkId": "%s",
                    "customHeaders": {
                        "string": "%s"
                    },
                    "type": "%s"
                },
                "connectTimeout": %d,
                "recording": {
                    "recordingType": "%s",
                    "recordingComposition": {
                        "enabled": %t
                    },
                    "customData": {
                        "string": "%s"
                    },
                    "filePrefix": "%s"
                },
                "type": "%s"
            }
        ]
    }`, givenId, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf("/callrouting/1/routes/%s", givenId), givenResponse, 200)

	sipEndpoint := voice.NewCallRoutingSipEndpoint(givenDestinationsValueSipTrunkId)
	sipEndpoint.SetUsername(givenDestinationsValueUsername)
	sipEndpoint.SetCustomHeaders(map[string]string{"string": givenDestinationsValueCustomHeadersString})
	callRoutingEndpoint := voice.CallRoutingSipEndpointAsCallRoutingEndpoint(sipEndpoint)
	destination := voice.NewCallRoutingEndpointDestination(callRoutingEndpoint)
	destination.SetConnectTimeout(givenDestinationsConnectTimeout)
	recording := voice.NewCallRoutingRecording(givenDestinationsRecordingRecordingType)
	recordingComposition := voice.NewCallRoutingRecordingComposition(givenDestinationsRecordingRecordingCompositionEnabled)
	recording.SetRecordingComposition(*recordingComposition)
	recording.SetCustomData(map[string]string{"string": givenDestinationsRecordingCustomDataString})
	recording.SetFilePrefix(givenDestinationsRecordingFilePrefix)
	destination.SetRecording(*recording)
	destinations := []voice.CallRoutingDestination{voice.CallRoutingEndpointDestinationAsCallRoutingDestination(destination)}
	request := voice.NewCallRoutingRouteRequest(givenName, destinations)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.
		CallRoutingAPI.
		UpdateCallRoute(context.Background(), givenId).
		CallRoutingRouteRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetDestinations()))
	destinationResponse := response.GetDestinations()[0].CallRoutingEndpointDestination
	destinationResponseValue := destinationResponse.GetValue().CallRoutingSipEndpoint
	assert.Equal(t, givenDestinationsValueUsername, destinationResponseValue.GetUsername())
	assert.Equal(t, givenDestinationsValueSipTrunkId, destinationResponseValue.GetSipTrunkId())
	assert.Equal(t, givenDestinationsValueCustomHeadersString, destinationResponseValue.GetCustomHeaders()["string"])
	assert.Equal(t, givenDestinationsValueType, destinationResponseValue.Type)
	assert.Equal(t, givenDestinationsConnectTimeout, destinationResponse.GetConnectTimeout())
	assert.Equal(t, givenDestinationsRecordingRecordingType, destinationResponse.GetRecording().RecordingType)
	assert.Equal(t, givenDestinationsRecordingRecordingCompositionEnabled, destinationResponse.GetRecording().RecordingComposition.GetEnabled())
	assert.Equal(t, &givenDestinationsRecordingFilePrefix, destinationResponse.GetRecording().FilePrefix)
	assert.Equal(t, givenDestinationsType, destinationResponse.Type)
}

func TestShouldDeleteCallRoute(t *testing.T) {
	givenId := "f8fc8aca-786d-4943-9af2-e7ec01b5e80d"
	givenName := "SIP endpoint route"
	givenDestinationsValueUsername := "41793026834"
	givenDestinationsValueSipTrunkId := "string"
	givenDestinationsValueCustomHeadersString := "string"
	givenDestinationsValueType := voice.CALLROUTINGENDPOINTTYPE_SIP
	givenDestinationsConnectTimeout := int32(30)
	givenDestinationsRecordingRecordingType := voice.CALLROUTINGRECORDINGTYPE_AUDIO
	givenDestinationsRecordingRecordingCompositionEnabled := true
	givenDestinationsRecordingCustomDataString := "string"
	givenDestinationsRecordingFilePrefix := "string"
	givenDestinationsType := voice.CALLROUTINGDESTINATIONTYPE_ENDPOINT

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "destinations": [
            {
                "value": {
                    "username": "%s",
                    "sipTrunkId": "%s",
                    "customHeaders": {
                        "string": "%s"
                    },
                    "type": "%s"
                },
                "connectTimeout": %d,
                "recording": {
                    "recordingType": "%s",
                    "recordingComposition": {
                        "enabled": %t
                    },
                    "customData": {
                        "string": "%s"
                    },
                    "filePrefix": "%s"
                },
                "type": "%s"
            }
        ]
    }`, givenId, givenName, givenDestinationsValueUsername, givenDestinationsValueSipTrunkId, givenDestinationsValueCustomHeadersString, givenDestinationsValueType, givenDestinationsConnectTimeout, givenDestinationsRecordingRecordingType, givenDestinationsRecordingRecordingCompositionEnabled, givenDestinationsRecordingCustomDataString, givenDestinationsRecordingFilePrefix, givenDestinationsType)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf("/callrouting/1/routes/%s", givenId), givenResponse, 200)

	response, _, err := infobipClient.
		CallRoutingAPI.
		DeleteCallRoute(context.Background(), givenId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenId, response.GetId())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, 1, len(response.GetDestinations()))
	destinationResponse := response.GetDestinations()[0].CallRoutingEndpointDestination
	destinationResponseValue := destinationResponse.GetValue().CallRoutingSipEndpoint
	assert.Equal(t, givenDestinationsValueUsername, destinationResponseValue.GetUsername())
	assert.Equal(t, givenDestinationsValueSipTrunkId, destinationResponseValue.GetSipTrunkId())
	assert.Equal(t, givenDestinationsValueCustomHeadersString, destinationResponseValue.GetCustomHeaders()["string"])
	assert.Equal(t, givenDestinationsValueType, destinationResponseValue.Type)
	assert.Equal(t, givenDestinationsConnectTimeout, destinationResponse.GetConnectTimeout())
	assert.Equal(t, givenDestinationsRecordingRecordingType, destinationResponse.GetRecording().RecordingType)
	assert.Equal(t, givenDestinationsRecordingRecordingCompositionEnabled, destinationResponse.GetRecording().RecordingComposition.GetEnabled())
	assert.Equal(t, &givenDestinationsRecordingFilePrefix, destinationResponse.GetRecording().FilePrefix)
	assert.Equal(t, givenDestinationsType, destinationResponse.Type)
}
