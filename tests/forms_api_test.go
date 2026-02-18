package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/moments"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	FormsEndpoint     = "/forms/1/forms"
	FormByIdEndpoint  = "/forms/1/forms/%s"
	FormDataEndpoint  = "/forms/1/forms/%s/data"
	FormViewsEndpoint = "/forms/1/forms/%s/views"
)

func TestShouldGetAllForms(t *testing.T) {
	givenFormId1 := "f23f0f7c-9898-4feb-8f21-5afe2c29db7e"
	givenFormName1 := "Test form"
	givenOffset := int32(0)
	givenLimit := int32(25)
	givenTotal := int64(1)
	givenResubmitEnabled := true
	givenFormType := moments.FORMSTYPE_OPT_IN
	givenFormStatus := moments.FORMSSTATUS_ACTIVE
	givenCreatedAt := "2015-02-22T17:42:05.390+0100"
	givenCreatedAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAt)
	ibTimeGivenCreatedAt := infobip.Time{
		T: givenCreatedAtDateTime,
	}

	givenUpdatedAt := "2015-02-22T17:42:05.390+0100"
	givenUpdatedAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAt)
	ibTimeGivenUpdatedAt := infobip.Time{
		T: givenUpdatedAtDateTime,
	}

	givenComponent := moments.FORMSCOMPONENTTYPE_TEXT
	givenFieldId := "last_name"
	givenPersonField := ""
	givenLabel := ""
	givenIsHidden := true
	givenIsRequired := true
	givenPlaceholder := ""

	givenResponse := fmt.Sprintf(`{
        "forms": [
            {
                "id": "%s",
                "name": "%s",
                "elements": [
                    {
                        "component": "%s",
                        "fieldId": "%s",
                        "personField": "%s",
                        "label": "%s",
                        "isRequired": %t,
                        "isHidden": %t,
                        "placeholder": "%s"
                    }
                ],
                "createdAt": "%s",
                "updatedAt": "%s",
                "resubmitEnabled": %t,
                "formType": "%s",
                "formStatus": "%s"
            }
        ],
        "offset": %d,
        "limit": %d,
        "total": %d
    }`, givenFormId1, givenFormName1, givenComponent, givenFieldId, givenPersonField, givenLabel, givenIsRequired, givenIsHidden, givenPlaceholder, givenCreatedAt, givenUpdatedAt, givenResubmitEnabled, givenFormType, givenFormStatus, givenOffset, givenLimit, givenTotal)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", FormsEndpoint, givenResponse, 200)

	response, _, err := infobipClient.
		FormsAPI.
		GetForms(context.Background()).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenOffset, response.GetOffset())
	assert.Equal(t, givenLimit, response.GetLimit())
	assert.Equal(t, givenTotal, response.GetTotal())

	forms := response.GetForms()
	assert.Len(t, forms, 1)

	form := forms[0]
	assert.Equal(t, givenFormId1, form.GetId())
	assert.Equal(t, givenFormName1, form.GetName())
	assert.Equal(t, ibTimeGivenCreatedAt, form.GetCreatedAt())
	assert.Equal(t, ibTimeGivenUpdatedAt, form.GetUpdatedAt())
	assert.Equal(t, givenResubmitEnabled, form.GetResubmitEnabled())
	assert.Equal(t, givenFormType, form.GetFormType())
	assert.Equal(t, givenFormStatus, form.GetFormStatus())

	elements := form.GetElements()
	assert.Len(t, elements, 1)

	element := elements[0]
	assert.Equal(t, givenComponent, element.GetComponent())
	assert.Equal(t, givenFieldId, element.GetFieldId())
	assert.Equal(t, givenPersonField, element.GetPersonField())
	assert.Equal(t, givenLabel, element.GetLabel())
	assert.Equal(t, givenIsRequired, element.GetIsRequired())
	assert.Equal(t, givenIsHidden, element.GetIsHidden())
	assert.Equal(t, givenPlaceholder, element.GetPlaceholder())
}

func TestShouldGetFormById(t *testing.T) {
	givenFormId := "f23f0f7c-9898-4feb-8f21-5afe2c29db7e"
	givenFormName := "Test form"
	givenResubmitEnabled := true
	givenFormType := moments.FORMSTYPE_OPT_IN
	givenFormStatus := moments.FORMSSTATUS_ACTIVE

	givenCreatedAt := "2015-02-22T17:42:05.390+0100"
	givenCreatedAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAt)
	ibTimeGivenCreatedAt := infobip.Time{
		T: givenCreatedAtDateTime,
	}

	givenUpdatedAt := "2015-02-22T17:42:05.390+0100"
	givenUpdatedAtDateTime, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenCreatedAt)
	ibTimeGivenUpdatedAt := infobip.Time{
		T: givenUpdatedAtDateTime,
	}

	givenComponent := moments.FORMSCOMPONENTTYPE_TEXT
	givenFieldId := "last_name"
	givenPersonField := ""
	givenLabel := ""
	givenIsHidden := true
	givenIsRequired := true
	givenPlaceholder := ""

	givenResponse := fmt.Sprintf(`{
        "id": "%s",
        "name": "%s",
        "elements": [
            {
                "component": "%s",
                "fieldId": "%s",
                "personField": "%s",
                "label": "%s",
                "isRequired": %t,
                "isHidden": %t,
                "placeholder": "%s"
            }
        ],
        "createdAt": "%s",
        "updatedAt": "%s",
        "resubmitEnabled": %t,
        "formType": "%s",
        "formStatus": "%s"
    }`, givenFormId, givenFormName, givenComponent, givenFieldId, givenPersonField, givenLabel, givenIsRequired, givenIsHidden, givenPlaceholder, givenCreatedAt, givenUpdatedAt, givenResubmitEnabled, givenFormType, givenFormStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", fmt.Sprintf(FormByIdEndpoint, givenFormId), givenResponse, 200)

	response, _, err := infobipClient.
		FormsAPI.
		GetForm(context.Background(), givenFormId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenFormId, response.GetId())
	assert.Equal(t, givenFormName, response.GetName())
	assert.Equal(t, ibTimeGivenCreatedAt, response.GetCreatedAt())
	assert.Equal(t, ibTimeGivenUpdatedAt, response.GetUpdatedAt())
	assert.Equal(t, givenResubmitEnabled, response.GetResubmitEnabled())
	assert.Equal(t, givenFormType, response.GetFormType())
	assert.Equal(t, givenFormStatus, response.GetFormStatus())

	elements := response.GetElements()
	assert.Len(t, elements, 1)

	element := elements[0]
	assert.Equal(t, givenComponent, element.GetComponent())
	assert.Equal(t, givenFieldId, element.GetFieldId())
	assert.Equal(t, givenPersonField, element.GetPersonField())
	assert.Equal(t, givenLabel, element.GetLabel())
	assert.Equal(t, givenIsRequired, element.GetIsRequired())
	assert.Equal(t, givenIsHidden, element.GetIsHidden())
	assert.Equal(t, givenPlaceholder, element.GetPlaceholder())
}

func TestShouldIncrementFormViewCount(t *testing.T) {
	formId := "12345"
	givenStatus := "OK"

	givenResponse := fmt.Sprintf(`{
        "status": "%s"
    }`, givenStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf(FormViewsEndpoint, formId), givenResponse, 200)

	response, _, err := infobipClient.
		FormsAPI.
		IncrementViewCount(context.Background(), formId).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}

func TestShouldSubmitFormData(t *testing.T) {
	formId := "12345"
	givenStatus := "OK"

	givenResponse := fmt.Sprintf(`{
       "status": "%s"
   }`, givenStatus)

	givenNumber := 26
	givenBoolean := true

	requestBody := fmt.Sprintf(`{
       "number": %d,
       "boolean": %t
   }`, givenNumber, givenBoolean)

	formDataRequest := map[string]interface{}{
		"number":  givenNumber,
		"boolean": givenBoolean,
	}

	actualRequest, _ := json.Marshal(formDataRequest)
	ValidateExpectedRequestBodiesMatches(t, requestBody, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", fmt.Sprintf(FormDataEndpoint, formId), givenResponse, 200)

	response, _, err := infobipClient.
		FormsAPI.
		SubmitFormData(context.Background(), formId).
		RequestBody(formDataRequest).
		Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenStatus, response.GetStatus())
}
