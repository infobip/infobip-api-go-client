package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/moments"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	CommunicationFlowParticipantsEndpoint = "/communication/1/flows/%d/participants"
	MomentsFlowParticipantsEndpoint       = "/moments/1/flows/%d/participants"
	MomentsFlowParticipantsReportEndpoint = "/moments/1/flows/%d/participants/report"
)

func TestShouldAddParticipantsToFlow(t *testing.T) {
	campaignId := int64(200000000000001)
	notifyUrl := "https://example.com"
	callbackData := "Callback Data"

	identifier1 := "370329180020364"
	type1 := moments.FLOWPERSONUNIQUEFIELDTYPE_FACEBOOK

	identifier2 := "test@infobip.com"
	type2 := moments.FLOWPERSONUNIQUEFIELDTYPE_EMAIL
	orderNumber2 := 1167873391

	identifier3 := "test2@infobip.com"
	type3 := moments.FLOWPERSONUNIQUEFIELDTYPE_FACEBOOK
	orderNumber3 := 1595299041
	externalId3 := "optional_external_person_id"
	contractExpiry3 := "2023-04-01"
	company3 := "Infobip"
	email3 := "test@infobip.com"

	givenRequest := fmt.Sprintf(`{
        "participants": [
            {
                "identifyBy": {
                    "identifier": "%s",
                    "type": "%s"
                }
            },
            {
                "identifyBy": {
                    "identifier": "%s",
                    "type": "%s"
                },
                "variables": {
                    "orderNumber": %d
                }
            },
            {
                "identifyBy": {
                    "identifier": "%s",
                    "type": "%s"
                },
                "variables": {
                    "orderNumber": %d
                },
                "person": {
                    "externalId": "%s",
                    "customAttributes": {
                        "Contract Expiry": "%s",
                        "Company": "%s"
                    },
                    "contactInformation": {
                        "email": [
                            {
                                "address": "%s"
                            }
                        ]
                    }
                }
            }
        ],
        "notifyUrl": "%s",
        "callbackData": "%s"
    }`, identifier1, type1, identifier2, type2, orderNumber2, identifier3, type3, orderNumber3, externalId3, contractExpiry3, company3, email3, notifyUrl, callbackData)

	participant1UniqueField := *moments.NewFlowPersonUniqueField(identifier1, moments.FLOWPERSONUNIQUEFIELDTYPE_FACEBOOK)
	participant1 := moments.NewFlowParticipant(participant1UniqueField)

	participant2UniqueField := *moments.NewFlowPersonUniqueField(identifier2, moments.FLOWPERSONUNIQUEFIELDTYPE_EMAIL)
	participant2 := moments.NewFlowParticipant(participant2UniqueField)
	participant2.SetVariables(map[string]interface{}{"orderNumber": orderNumber2})

	participant3UniqueField := *moments.NewFlowPersonUniqueField(identifier3, moments.FLOWPERSONUNIQUEFIELDTYPE_FACEBOOK)
	participant3 := moments.NewFlowParticipant(participant3UniqueField)
	participant3.SetVariables(map[string]interface{}{"orderNumber": orderNumber3})
	person := moments.NewFlowPerson()
	person.SetExternalId(externalId3)
	person.SetCustomAttributes(map[string]interface{}{"Contract Expiry": contractExpiry3, "Company": company3})
	contactInfo := moments.NewFlowPersonContacts()
	contactInfo.SetEmail([]moments.FlowEmailContact{{Address: &email3}})
	person.SetContactInformation(*contactInfo)
	participant3.SetPerson(*person)

	request := moments.NewFlowAddFlowParticipantsRequest([]moments.FlowParticipant{*participant1, *participant2, *participant3})
	request.SetNotifyUrl(notifyUrl)
	request.SetCallbackData(callbackData)

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, givenRequest, string(actualRequest))

	givenOperationId := "03f2d474-0508-46bf-9f3d-d8e2c28adaea"

	givenResponse := fmt.Sprintf(`{
        "operationId": "%s"
    }`, givenOperationId)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf(MomentsFlowParticipantsEndpoint, campaignId), givenResponse, 200)

	response, _, err := infobipClient.
		FlowAPI.
		AddFlowParticipants(context.Background(), campaignId).
		FlowAddFlowParticipantsRequest(*request).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenOperationId, response.GetOperationId())
}

func TestShouldGetParticipantsReport(t *testing.T) {
	campaignId := int64(200000000000001)
	givenOperationId := "03f2d474-0508-46bf-9f3d-d8e2c28adaea"
	givenCallbackData := "Callback Data"
	givenIdentifier1 := "test@infobip.com"
	givenIdentifier2 := "test2@infobip.com"
	givenType := moments.FLOWPERSONUNIQUEFIELDTYPE_EMAIL
	givenStatus1 := moments.FLOWADDFLOWPARTICIPANTSTATUS_ACCEPTED
	givenStatus2 := moments.FLOWADDFLOWPARTICIPANTSTATUS_REJECTED
	givenErrorReason := moments.FLOWERRORSTATUSREASON_INVALID_CONTACT

	givenResponse := fmt.Sprintf(`{
        "operationId": "%s",
        "campaignId": %d,
        "callbackData": "%s",
        "participants": [
            {
                "identifyBy": {
                    "identifier": "%s",
                    "type": "%s"
                },
                "status": "%s"
            },
            {
                "identifyBy": {
                    "identifier": "%s",
                    "type": "%s"
                },
                "status": "%s",
                "errorReason": "%s"
            }
        ]
    }`, givenOperationId, campaignId, givenCallbackData, givenIdentifier1, givenType, givenStatus1, givenIdentifier2, givenType, givenStatus2, givenErrorReason)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(MomentsFlowParticipantsReportEndpoint, campaignId), givenResponse, 200)

	response, _, err := infobipClient.
		FlowAPI.
		GetFlowParticipantsAddedReport(context.Background(), campaignId).
		OperationId(givenOperationId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenOperationId, response.GetOperationId())
	assert.Equal(t, campaignId, response.GetCampaignId())
	assert.Equal(t, givenCallbackData, response.GetCallbackData())

	participants := response.GetParticipants()
	assert.Len(t, participants, 2)

	participant1 := participants[0]
	assert.Equal(t, givenIdentifier1, participant1.GetIdentifyBy().Identifier)
	assert.Equal(t, givenType, participant1.GetIdentifyBy().Type)
	assert.Equal(t, givenStatus1, participant1.GetStatus())

	participant2 := participants[1]
	assert.Equal(t, givenIdentifier2, participant2.GetIdentifyBy().Identifier)
	assert.Equal(t, givenType, participant2.GetIdentifyBy().Type)
	assert.Equal(t, givenStatus2, participant2.GetStatus())
	assert.Equal(t, givenErrorReason, participant2.GetErrorReason())
}

func TestShouldRemoveParticipantFromFlow(t *testing.T) {
	campaignId := int64(200000000000001)

	expectedResponse := "{}"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("DELETE", fmt.Sprintf(CommunicationFlowParticipantsEndpoint, campaignId), expectedResponse, 200)

	_, err := infobipClient.
		FlowAPI.
		RemovePeopleFromFlow(context.Background(), campaignId).
		Execute()

	assert.Nil(t, err)
}
