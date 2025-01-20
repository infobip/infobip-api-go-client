/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voice

import (
	"encoding/json"
	"fmt"
)

// CallbackResponse struct for CallbackResponse
type CallbackResponse struct {
	AudioCallbackResponse       *AudioCallbackResponse
	CaptureDtmfCallbackResponse *CaptureDtmfCallbackResponse
	DialCallbackResponse        *DialCallbackResponse
}

// AudioCallbackResponseAsCallbackResponse is a convenience function that returns AudioCallbackResponse wrapped in CallbackResponse
func AudioCallbackResponseAsCallbackResponse(v *AudioCallbackResponse) CallbackResponse {
	return CallbackResponse{
		AudioCallbackResponse: v,
	}
}

// CaptureDtmfCallbackResponseAsCallbackResponse is a convenience function that returns CaptureDtmfCallbackResponse wrapped in CallbackResponse
func CaptureDtmfCallbackResponseAsCallbackResponse(v *CaptureDtmfCallbackResponse) CallbackResponse {
	return CallbackResponse{
		CaptureDtmfCallbackResponse: v,
	}
}

// DialCallbackResponseAsCallbackResponse is a convenience function that returns DialCallbackResponse wrapped in CallbackResponse
func DialCallbackResponseAsCallbackResponse(v *DialCallbackResponse) CallbackResponse {
	return CallbackResponse{
		DialCallbackResponse: v,
	}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CallbackResponse) UnmarshalJSON(data []byte) error {
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err := json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'audio'
	if jsonDict["command"] == "audio" {
		// try to unmarshal JSON data into AudioCallbackResponse
		err = json.Unmarshal(data, &dst.AudioCallbackResponse)
		if err == nil {
			jsonAudioCallbackResponse, _ := json.Marshal(dst.AudioCallbackResponse)
			if string(jsonAudioCallbackResponse) == "{}" { // empty struct
				dst.AudioCallbackResponse = nil
			} else {
				return nil // data stored in dst.AudioCallbackResponse, return on the first match
			}
		} else {
			dst.AudioCallbackResponse = nil
		}
	}
	// check if the discriminator value is 'captureDtmf'
	if jsonDict["command"] == "captureDtmf" {
		// try to unmarshal JSON data into CaptureDtmfCallbackResponse
		err = json.Unmarshal(data, &dst.CaptureDtmfCallbackResponse)
		if err == nil {
			jsonCaptureDtmfCallbackResponse, _ := json.Marshal(dst.CaptureDtmfCallbackResponse)
			if string(jsonCaptureDtmfCallbackResponse) == "{}" { // empty struct
				dst.CaptureDtmfCallbackResponse = nil
			} else {
				return nil // data stored in dst.CaptureDtmfCallbackResponse, return on the first match
			}
		} else {
			dst.CaptureDtmfCallbackResponse = nil
		}
	}
	// check if the discriminator value is 'dial'
	if jsonDict["command"] == "dial" {
		// try to unmarshal JSON data into DialCallbackResponse
		err = json.Unmarshal(data, &dst.DialCallbackResponse)
		if err == nil {
			jsonDialCallbackResponse, _ := json.Marshal(dst.DialCallbackResponse)
			if string(jsonDialCallbackResponse) == "{}" { // empty struct
				dst.DialCallbackResponse = nil
			} else {
				return nil // data stored in dst.DialCallbackResponse, return on the first match
			}
		} else {
			dst.DialCallbackResponse = nil
		}
	}
	return fmt.Errorf("Data failed to match schemas in anyOf(CallbackResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CallbackResponse) MarshalJSON() ([]byte, error) {
	if src.AudioCallbackResponse != nil {
		return json.Marshal(&src.AudioCallbackResponse)
	}
	if src.CaptureDtmfCallbackResponse != nil {
		return json.Marshal(&src.CaptureDtmfCallbackResponse)
	}
	if src.DialCallbackResponse != nil {
		return json.Marshal(&src.DialCallbackResponse)
	}
	return nil, nil // no data in anyOf schemas
}

// Get the actual instance
func (obj *CallbackResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AudioCallbackResponse != nil {
		return obj.AudioCallbackResponse
	}
	if obj.CaptureDtmfCallbackResponse != nil {
		return obj.CaptureDtmfCallbackResponse
	}
	if obj.DialCallbackResponse != nil {
		return obj.DialCallbackResponse
	}
	// all schemas are nil
	return nil
}

type NullableCallbackResponse struct {
	value *CallbackResponse
	isSet bool
}

func (v NullableCallbackResponse) Get() *CallbackResponse {
	return v.value
}

func (v *NullableCallbackResponse) Set(val *CallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallbackResponse(val *CallbackResponse) *NullableCallbackResponse {
	return &NullableCallbackResponse{value: val, isSet: true}
}

func (v NullableCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
