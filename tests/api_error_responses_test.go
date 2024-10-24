package infobip

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/sms"
)

func TestShouldProcessErrorResponses(t *testing.T) {
	testCases := []struct {
		code        int
		name        string
		description string
	}{
		{400, "BAD_REQUEST", "Bad request"},
		{500, "GENERAL_ERROR", "Something went wrong. Please contact support."},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc // capture range variable

			givenText := "This is a sample message"
			givenResponse := GetGivenResponse(tc.name, tc.description)

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			SetUpSuccessRequest("POST", SmsSendPreviewEndpoint, givenResponse, tc.code)

			smsPreviewRequest := sms.NewPreviewRequest(givenText)

			_, _, err := infobipClient.
				SmsAPI.
				PreviewSmsMessage(context.Background()).
				PreviewRequest(*smsPreviewRequest).
				Execute()

			if err != nil {
				apiErr, isApiErr := err.(*api.GenericOpenAPIError)
				if isApiErr {
					ibErr, isIbErr := apiErr.Model().(sms.ApiException)
					if isIbErr {
						assert.Equal(t, tc.name, ibErr.RequestError.ServiceException.GetMessageId())
						assert.Equal(t, tc.description, ibErr.RequestError.ServiceException.GetText())
					} else {
						t.Fatalf("Expected apiErr to be of type `infobip.SmsApiException`")
					}
				} else {
					t.Fatalf("Expected ibErr to be of type `infobip.GenericOpenAPIError`")
				}
			} else {
				t.Fatalf("Expected response err field not to be nil")
			}
		})
	}
}

func GetGivenResponse(errorName string, errorDescription string) string {
	return fmt.Sprintf(`
		{
			"requestError": {
				"serviceException": {
					"messageId": "%s",
					"text": "%s"
				}
			}
		}
	`,
		errorName,
		errorDescription,
	)
}
