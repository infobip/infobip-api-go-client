package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/moments"
)

// Adds multiple participants to a Moments flow with variables and person details.
// Replace <USER_EMAIL> and <PHONE_NUMBER> with values accepted by your account before execution.
func TestAddFlowParticipants(t *testing.T) {
	client, auth := newClientAndAuth()

	participantEmail := moments.NewFlowParticipant(
		*moments.NewFlowPersonUniqueField("<USER_EMAIL>", moments.FLOWPERSONUNIQUEFIELDTYPE_EMAIL),
	)
	participantEmail.SetVariables(map[string]interface{}{"orderNumber": 1167873391})

	participantPhone := moments.NewFlowParticipant(
		*moments.NewFlowPersonUniqueField("<PHONE_NUMBER>", moments.FLOWPERSONUNIQUEFIELDTYPE_PHONE),
	)
	participantPhone.SetPerson(moments.FlowPerson{
		ExternalId: infobip.PtrString("external-id-001"),
		ContactInformation: &moments.FlowPersonContacts{
			Phone: []moments.FlowPhoneContact{
				{Number: infobip.PtrString("<PHONE_NUMBER>")},
			},
		},
	})

	request := moments.NewFlowAddFlowParticipantsRequest([]moments.FlowParticipant{
		*participantEmail,
		*participantPhone,
	})
	request.SetNotifyUrl("https://example.com/moments/notify")
	request.SetCallbackData("flow-participants-callback")

	apiResponse, httpResponse, err := client.
		FlowAPI.
		AddFlowParticipants(auth, 200000000000001).
		FlowAddFlowParticipantsRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to add flow participants: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.OperationId == "" {
		t.Fatalf("Invalid response: expected operation id, but got: %+v", apiResponse)
	}
}

// Add a single email participant with minimal variables.
func TestAddSingleFlowParticipant(t *testing.T) {
	client, auth := newClientAndAuth()

	participant := moments.NewFlowParticipant(
		*moments.NewFlowPersonUniqueField("<USER_EMAIL>", moments.FLOWPERSONUNIQUEFIELDTYPE_EMAIL),
	)
	participant.SetVariables(map[string]interface{}{"firstName": "Ava"})

	request := moments.NewFlowAddFlowParticipantsRequest([]moments.FlowParticipant{*participant})

	apiResponse, httpResponse, err := client.
		FlowAPI.
		AddFlowParticipants(auth, 200000000000001).
		FlowAddFlowParticipantsRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to add single flow participant: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.OperationId == "" {
		t.Fatalf("Invalid response: expected operation id, but got: %+v", apiResponse)
	}
}
