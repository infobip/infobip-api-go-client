package infobip

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetVoiceMaskingConfig(t *testing.T) {
	givenKey := "string"
	givenName := "string"
	givenCallbackUrl := "string"
	givenStatusUrl := "string"
	givenBackupCallbackUrl := "string"
	givenBackupStatusUrl := "string"
	givenDescription := "string"
	givenInsertDateTime := "2015-02-22T17:42:05.390+0100"
	givenInsertDateTimeOffset, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenInsertDateTime)
	ibTimeGivenInsertDateTime := infobip.Time{
		T: givenInsertDateTimeOffset,
	}

	givenUpdateDateTime := "2015-02-22T17:42:05.390+0100"
	givenUpdateDateTimeOffset, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenInsertDateTime)
	ibTimeGivenUpdateDateTime := infobip.Time{
		T: givenUpdateDateTimeOffset,
	}

	givenResponse := fmt.Sprintf(`[
        {
            "key": "%s",
            "name": "%s",
            "callbackUrl": "%s",
            "statusUrl": "%s",
            "backupCallbackUrl": "%s",
            "backupStatusUrl": "%s",
            "description": "%s",
            "insertDateTime": "%s",
            "updateDateTime": "%s"
        }
    ]`, givenKey, givenName, givenCallbackUrl, givenStatusUrl, givenBackupCallbackUrl, givenBackupStatusUrl, givenDescription, givenInsertDateTime, givenUpdateDateTime)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/voice/masking/2/config", givenResponse, 200)

	response, _, err := infobipClient.NumberMaskingAPI.GetNumberMaskingConfigurations(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response, 1)

	config := response[0]
	assert.Equal(t, givenKey, config.GetKey())
	assert.Equal(t, givenName, config.GetName())
	assert.Equal(t, givenCallbackUrl, config.GetCallbackUrl())
	assert.Equal(t, givenStatusUrl, config.GetStatusUrl())
	assert.Equal(t, givenBackupCallbackUrl, config.GetBackupCallbackUrl())
	assert.Equal(t, givenBackupStatusUrl, config.GetBackupStatusUrl())
	assert.Equal(t, givenDescription, config.GetDescription())
	assert.Equal(t, ibTimeGivenInsertDateTime, config.GetInsertDateTime())
	assert.Equal(t, ibTimeGivenUpdateDateTime, config.GetUpdateDateTime())
}

func TestShouldCreateNumberMaskingConfiguration(t *testing.T) {
	givenKey := "3FC0C9CB4AFAEAC67E8FC6BA3B1E044A"
	givenName := "UniqueConfigurationName"
	givenCallbackUrl := "http://xyz.com/1/callback"
	givenStatusUrl := "http://xyz.com/1/status"
	givenInsertDateTime := "2015-02-22T17:42:05.390+0100"
	givenInsertDateTimeOffset, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenInsertDateTime)
	ibTimeGivenInsertDateTime := infobip.Time{
		T: givenInsertDateTimeOffset,
	}

	givenUpdateDateTime := "2015-02-22T17:42:05.390+0100"
	givenUpdateDateTimeOffset, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenInsertDateTime)
	ibTimeGivenUpdateDateTime := infobip.Time{
		T: givenUpdateDateTimeOffset,
	}

	expectedRequest := fmt.Sprintf(`{
        "name": "%s",
        "callbackUrl": "%s",
        "statusUrl": "%s"
    }`, givenName, givenCallbackUrl, givenStatusUrl)

	givenResponse := fmt.Sprintf(`{
        "key": "%s",
        "name": "%s",
        "callbackUrl": "%s",
        "statusUrl": "%s",
        "insertDateTime": "%s",
        "updateDateTime": "%s"
    }`, givenKey, givenName, givenCallbackUrl, givenStatusUrl, givenInsertDateTime, givenUpdateDateTime)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/voice/masking/2/config", givenResponse, 200)

	request := voice.NumberMaskingSetupBody{
		Name:        givenName,
		CallbackUrl: givenCallbackUrl,
		StatusUrl:   &givenStatusUrl,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.NumberMaskingAPI.CreateNumberMaskingConfiguration(context.Background()).NumberMaskingSetupBody(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenKey, response.GetKey())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenCallbackUrl, response.GetCallbackUrl())
	assert.Equal(t, givenStatusUrl, response.GetStatusUrl())
	assert.Equal(t, ibTimeGivenInsertDateTime, response.GetInsertDateTime())
	assert.Equal(t, ibTimeGivenUpdateDateTime, response.GetUpdateDateTime())
}

func TestShouldUpdateNumberMaskingConfiguration(t *testing.T) {
	givenKey := "3FC0C9CB4AFAEAC67E8FC6BA3B1E044A"
	givenName := "UniqueConfigurationName"
	givenCallbackUrl := "http://example.com/1/callback"
	givenStatusUrl := "http://example.com/1/status"
	givenInsertDateTime := "2015-02-22T17:42:05.390+0100"
	givenInsertDateTimeOffset, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenInsertDateTime)
	ibTimeGivenInsertDateTime := infobip.Time{
		T: givenInsertDateTimeOffset,
	}

	givenUpdateDateTime := "2015-02-22T17:42:05.390+0100"
	givenUpdateDateTimeOffset, _ := time.Parse(infobip.INFOBIP_TIME_FORMAT, givenInsertDateTime)
	ibTimeGivenUpdateDateTime := infobip.Time{
		T: givenUpdateDateTimeOffset,
	}

	expectedRequest := fmt.Sprintf(`{
        "name": "%s",
        "callbackUrl": "%s",
        "statusUrl": "%s"
    }`, givenName, givenCallbackUrl, givenStatusUrl)

	givenResponse := fmt.Sprintf(`{
        "key": "%s",
        "name": "%s",
        "callbackUrl": "%s",
        "statusUrl": "%s",
        "insertDateTime": "%s",
        "updateDateTime": "%s"
    }`, givenKey, givenName, givenCallbackUrl, givenStatusUrl, givenInsertDateTime, givenUpdateDateTime)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", fmt.Sprintf("/voice/masking/2/config/%s", givenKey), givenResponse, 200)

	request := voice.NumberMaskingSetupBody{
		Name:        givenName,
		CallbackUrl: givenCallbackUrl,
		StatusUrl:   &givenStatusUrl,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	response, _, err := infobipClient.NumberMaskingAPI.UpdateNumberMaskingConfiguration(context.Background(), givenKey).NumberMaskingSetupBody(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenKey, response.GetKey())
	assert.Equal(t, givenName, response.GetName())
	assert.Equal(t, givenCallbackUrl, response.GetCallbackUrl())
	assert.Equal(t, givenStatusUrl, response.GetStatusUrl())
	assert.Equal(t, ibTimeGivenInsertDateTime, response.GetInsertDateTime())
	assert.Equal(t, ibTimeGivenUpdateDateTime, response.GetUpdateDateTime())
}

func TestShouldUploadAudioFile(t *testing.T) {
	givenUrl := "http://www.example.com/audio.wav"
	givenFileId := "cb702ae4-f356-4efd-b2dd-7a667b570af5"

	expectedRequest := fmt.Sprintf(`{
        "url": "%s"
    }`, givenUrl)

	givenResponse := fmt.Sprintf(`{
        "fileId": "%s"
    }`, givenFileId)

	request := voice.NumberMaskingUploadBody{
		Url: &givenUrl,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/voice/masking/1/upload", givenResponse, 200)

	response, _, err := infobipClient.NumberMaskingAPI.UploadAudioFiles(context.Background()).NumberMaskingUploadBody(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenFileId, response.GetFileId())
}

func TestShouldGetNumberMaskingCredentials(t *testing.T) {
	givenApiId := "55ddccad2df62a4b615b7e3c472b2ab6"
	givenKey := "5da086b6a8e4424993646b8699c333ca"

	givenResponse := fmt.Sprintf(`{
        "apiId": "%s",
        "key": "%s"
    }`, givenApiId, givenKey)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("GET", "/voice/masking/2/credentials", givenResponse, 200)

	response, _, err := infobipClient.NumberMaskingAPI.GetNumberMaskingCredentials(context.Background()).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenApiId, response.GetApiId())
	assert.Equal(t, givenKey, response.GetKey())
}

func TestShouldCreateNumberMaskingCredentials(t *testing.T) {
	givenApiId := "55ddccad2df62a4b615b7e3c472b2ab6"
	givenKey := "5da086b6a8e4424993646b8699c333ca"

	expectedRequest := fmt.Sprintf(`{
        "apiId": "%s",
        "key": "%s"
    }`, givenApiId, givenKey)

	givenResponse := fmt.Sprintf(`{
        "apiId": "%s",
        "key": "%s"
    }`, givenApiId, givenKey)

	request := voice.NumberMaskingCredentialsBody{
		ApiId: givenApiId,
		Key:   givenKey,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("POST", "/voice/masking/2/credentials", givenResponse, 200)

	response, _, err := infobipClient.NumberMaskingAPI.CreateNumberMaskingCredentials(context.Background()).NumberMaskingCredentialsBody(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenApiId, response.GetApiId())
	assert.Equal(t, givenKey, response.GetKey())
}

func TestShouldUpdateNumberMaskingCredentials(t *testing.T) {
	givenApiId := "55ddccad2df62a4b615b7e3c472b2ab6"
	givenKey := "5da086b6a8e4424993646b8699c333ca"

	expectedRequest := fmt.Sprintf(`{
        "apiId": "%s",
        "key": "%s"
    }`, givenApiId, givenKey)

	givenResponse := fmt.Sprintf(`{
        "apiId": "%s",
        "key": "%s"
    }`, givenApiId, givenKey)

	request := voice.NumberMaskingCredentialsBody{
		ApiId: givenApiId,
		Key:   givenKey,
	}

	actualRequest, _ := json.Marshal(request)
	ValidateExpectedRequestBodiesMatches(t, expectedRequest, string(actualRequest))

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	SetUpSuccessRequest("PUT", "/voice/masking/2/credentials", givenResponse, 200)

	response, _, err := infobipClient.NumberMaskingAPI.UpdateNumberMaskingCredentials(context.Background()).NumberMaskingCredentialsBody(request).Execute()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, givenApiId, response.GetApiId())
	assert.Equal(t, givenKey, response.GetKey())
}
