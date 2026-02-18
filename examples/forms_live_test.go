package examples

import (
	"fmt"
	"testing"
)

func TestGetFormById(t *testing.T) {
	client, auth := newClientAndAuth()

	apiResponse, httpResponse, err := client.
		FormsAPI.
		GetForm(auth, "<FORM_ID>").
		Execute()

	if err != nil {
		t.Fatalf("Failed to fetch form: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Id == "" {
		t.Fatalf("Invalid response: expected form id, got: %+v", apiResponse)
	}
}

func TestSubmitFormDataAdvanced(t *testing.T) {
	client, auth := newClientAndAuth()

	requestBody := map[string]interface{}{
		"firstName": "Grace",
		"email":     "grace@example.com",
		"consent":   true,
	}

	apiResponse, httpResponse, err := client.
		FormsAPI.
		SubmitFormData(auth, "<FORM_ID>").
		RequestBody(requestBody).
		IbSubmissionSource("landing-page").
		IbSubmissionFormCampaign("spring-campaign").
		Execute()

	if err != nil {
		t.Fatalf("Failed to submit form data: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Status == "" {
		t.Fatalf("Invalid response: expected status, got: %+v", apiResponse)
	}
}
