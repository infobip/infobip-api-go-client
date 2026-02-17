package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
)

func TestCreateNumberMaskingConfigBasic(t *testing.T) {
	client, auth := newClientAndAuth()

	body := voice.NewNumberMaskingSetupBody("ride-hailing", "https://example.com/masking/callback")
	body.SetStatusUrl("https://example.com/masking/status")

	apiResponse, httpResponse, err := client.
		NumberMaskingAPI.
		CreateNumberMaskingConfiguration(auth).
		NumberMaskingSetupBody(*body).
		Execute()

	if err != nil {
		t.Fatalf("Failed to create number masking configuration: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Key == nil {
		t.Fatalf("Invalid response: expected configuration key, got: %+v", apiResponse)
	}
}

func TestUploadMaskingAudioAndCredentials(t *testing.T) {
	client, auth := newClientAndAuth()

	uploadBody := voice.NumberMaskingUploadBody{}
	uploadBody.SetUrl("https://example.com/audio/welcome.mp3")

	uploadResp, uploadHttpResp, err := client.
		NumberMaskingAPI.
		UploadAudioFiles(auth).
		NumberMaskingUploadBody(uploadBody).
		Execute()

	if err != nil {
		t.Fatalf("Failed to upload masking audio: %v", err)
	}

	fmt.Printf("Upload Response: %+v\n", uploadResp)
	fmt.Printf("Upload HTTP Response Details: %+v\n", uploadHttpResp)

	creds := voice.NewNumberMaskingCredentialsBody("<MASKING_API_ID>", "<MASKING_API_KEY>")

	credResp, credHttpResp, err := client.
		NumberMaskingAPI.
		UpdateNumberMaskingCredentials(auth).
		NumberMaskingCredentialsBody(*creds).
		Execute()

	if err != nil {
		t.Fatalf("Failed to update masking credentials: %v", err)
	}

	fmt.Printf("Credentials Response: %+v\n", credResp)
	fmt.Printf("Credentials HTTP Response Details: %+v\n", credHttpResp)

	if credResp == nil || credResp.Key == nil {
		t.Fatalf("Invalid response: expected credentials, got: %+v", credResp)
	}
}
