package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
)

func TestSendSingleVoiceTTS(t *testing.T) {
	client, auth := newClientAndAuth()

	request := voice.NewSingleRequest("<FROM_NUMBER>", "<TO_NUMBER>")
	request.SetLanguage("en")
	request.SetText("Test Voice message using the Go SDK with custom voice and language.")
	request.SetVoice(voice.Voice{
		Name:   infobip.PtrString("Joanna"),
		Gender: infobip.PtrString("female"),
	})

	apiResponse, httpResponse, err := client.
		VoiceAPI.
		SendSingleVoiceTts(auth).
		SingleRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send TTS message: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestSendMultipleVoiceTTSAdvanced(t *testing.T) {
	client, auth := newClientAndAuth()

	first := voice.NewMultiMessage([]string{"<TO_NUMBER_1>"})
	first.SetFrom("<FROM_NUMBER>")
	first.SetText("<speak><break time=\"500ms\"/>Your pickup is arriving soon.</speak>")
	first.SetLanguage("en")

	second := voice.NewMultiMessage([]string{"<TO_NUMBER_2>"})
	second.SetFrom("<FROM_NUMBER>")
	second.SetText("<speak>Please press 1 to confirm.</speak>")
	second.SetLanguage("en")

	request := voice.NewMultiRequest([]voice.MultiMessage{*first, *second})

	apiResponse, httpResponse, err := client.
		VoiceAPI.
		SendMultipleVoiceTts(auth).
		MultiRequest(*request).
		Execute()

	if err != nil {
		t.Fatalf("Failed to send advanced multiple TTS: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}
