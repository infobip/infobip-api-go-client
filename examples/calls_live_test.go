package examples

import (
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/models/voice"
)

func TestConnectCallsSimple(t *testing.T) {
	client, auth := newClientAndAuth()

	connectReq := voice.NewConnectRequest([]string{"<CALL_ID_A>", "<CALL_ID_B>"})

	apiResponse, httpResponse, err := client.
		CallsAPI.
		ConnectCalls(auth).
		ConnectRequest(*connectReq).
		Execute()

	if err != nil {
		t.Fatalf("Failed to connect calls: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Id == nil {
		t.Fatalf("Invalid response: expected conference id, got: %+v", apiResponse)
	}
}

func TestCallSayTextAdvanced(t *testing.T) {
	client, auth := newClientAndAuth()

	sayReq := voice.NewCallSayRequest("<speak>Welcome to our priority line.</speak>", voice.LANGUAGE_EN)
	sayReq.SetSpeechRate(1.1)
	sayReq.SetLoopCount(2)
	voiceName := voice.VOICENAME_JOANNA
	voiceGender := voice.GENDER_FEMALE
	sayReq.SetPreferences(voice.VoicePreferences{VoiceName: &voiceName, VoiceGender: &voiceGender})
	sayReq.SetCustomData(map[string]string{"segment": "vip", "correlationId": "call-2026-02-11"})

	apiResponse, httpResponse, err := client.
		CallsAPI.
		CallSayText(auth, "<ACTIVE_CALL_ID>").
		CallSayRequest(*sayReq).
		Execute()

	if err != nil {
		t.Fatalf("Failed to say text into call: %v", err)
	}

	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	if apiResponse == nil || apiResponse.Status == nil {
		t.Fatalf("Invalid response: expected status, got: %+v", apiResponse)
	}
}
