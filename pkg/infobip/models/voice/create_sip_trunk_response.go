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

// CreateSipTrunkResponse struct for CreateSipTrunkResponse
type CreateSipTrunkResponse struct {
	CreateProviderSipTrunkResponse   *CreateProviderSipTrunkResponse
	CreateRegisteredSipTrunkResponse *CreateRegisteredSipTrunkResponse
	CreateStaticSipTrunkResponse     *CreateStaticSipTrunkResponse
}

// CreateProviderSipTrunkResponseAsCreateSipTrunkResponse is a convenience function that returns CreateProviderSipTrunkResponse wrapped in CreateSipTrunkResponse
func CreateProviderSipTrunkResponseAsCreateSipTrunkResponse(v *CreateProviderSipTrunkResponse) CreateSipTrunkResponse {
	return CreateSipTrunkResponse{
		CreateProviderSipTrunkResponse: v,
	}
}

// CreateRegisteredSipTrunkResponseAsCreateSipTrunkResponse is a convenience function that returns CreateRegisteredSipTrunkResponse wrapped in CreateSipTrunkResponse
func CreateRegisteredSipTrunkResponseAsCreateSipTrunkResponse(v *CreateRegisteredSipTrunkResponse) CreateSipTrunkResponse {
	return CreateSipTrunkResponse{
		CreateRegisteredSipTrunkResponse: v,
	}
}

// CreateStaticSipTrunkResponseAsCreateSipTrunkResponse is a convenience function that returns CreateStaticSipTrunkResponse wrapped in CreateSipTrunkResponse
func CreateStaticSipTrunkResponseAsCreateSipTrunkResponse(v *CreateStaticSipTrunkResponse) CreateSipTrunkResponse {
	return CreateSipTrunkResponse{
		CreateStaticSipTrunkResponse: v,
	}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreateSipTrunkResponse) UnmarshalJSON(data []byte) error {
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err := json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'PROVIDER'
	if jsonDict["type"] == "PROVIDER" {
		// try to unmarshal JSON data into CreateProviderSipTrunkResponse
		err = json.Unmarshal(data, &dst.CreateProviderSipTrunkResponse)
		if err == nil {
			jsonCreateProviderSipTrunkResponse, _ := json.Marshal(dst.CreateProviderSipTrunkResponse)
			if string(jsonCreateProviderSipTrunkResponse) == "{}" { // empty struct
				dst.CreateProviderSipTrunkResponse = nil
			} else {
				return nil // data stored in dst.CreateProviderSipTrunkResponse, return on the first match
			}
		} else {
			dst.CreateProviderSipTrunkResponse = nil
		}
	}
	// check if the discriminator value is 'REGISTERED'
	if jsonDict["type"] == "REGISTERED" {
		// try to unmarshal JSON data into CreateRegisteredSipTrunkResponse
		err = json.Unmarshal(data, &dst.CreateRegisteredSipTrunkResponse)
		if err == nil {
			jsonCreateRegisteredSipTrunkResponse, _ := json.Marshal(dst.CreateRegisteredSipTrunkResponse)
			if string(jsonCreateRegisteredSipTrunkResponse) == "{}" { // empty struct
				dst.CreateRegisteredSipTrunkResponse = nil
			} else {
				return nil // data stored in dst.CreateRegisteredSipTrunkResponse, return on the first match
			}
		} else {
			dst.CreateRegisteredSipTrunkResponse = nil
		}
	}
	// check if the discriminator value is 'STATIC'
	if jsonDict["type"] == "STATIC" {
		// try to unmarshal JSON data into CreateStaticSipTrunkResponse
		err = json.Unmarshal(data, &dst.CreateStaticSipTrunkResponse)
		if err == nil {
			jsonCreateStaticSipTrunkResponse, _ := json.Marshal(dst.CreateStaticSipTrunkResponse)
			if string(jsonCreateStaticSipTrunkResponse) == "{}" { // empty struct
				dst.CreateStaticSipTrunkResponse = nil
			} else {
				return nil // data stored in dst.CreateStaticSipTrunkResponse, return on the first match
			}
		} else {
			dst.CreateStaticSipTrunkResponse = nil
		}
	}
	return fmt.Errorf("Data failed to match schemas in anyOf(CreateSipTrunkResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateSipTrunkResponse) MarshalJSON() ([]byte, error) {
	if src.CreateProviderSipTrunkResponse != nil {
		return json.Marshal(&src.CreateProviderSipTrunkResponse)
	}
	if src.CreateRegisteredSipTrunkResponse != nil {
		return json.Marshal(&src.CreateRegisteredSipTrunkResponse)
	}
	if src.CreateStaticSipTrunkResponse != nil {
		return json.Marshal(&src.CreateStaticSipTrunkResponse)
	}
	return nil, nil // no data in anyOf schemas
}

// Get the actual instance
func (obj *CreateSipTrunkResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateProviderSipTrunkResponse != nil {
		return obj.CreateProviderSipTrunkResponse
	}
	if obj.CreateRegisteredSipTrunkResponse != nil {
		return obj.CreateRegisteredSipTrunkResponse
	}
	if obj.CreateStaticSipTrunkResponse != nil {
		return obj.CreateStaticSipTrunkResponse
	}
	// all schemas are nil
	return nil
}

type NullableCreateSipTrunkResponse struct {
	value *CreateSipTrunkResponse
	isSet bool
}

func (v NullableCreateSipTrunkResponse) Get() *CreateSipTrunkResponse {
	return v.value
}

func (v *NullableCreateSipTrunkResponse) Set(val *CreateSipTrunkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSipTrunkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSipTrunkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSipTrunkResponse(val *CreateSipTrunkResponse) *NullableCreateSipTrunkResponse {
	return &NullableCreateSipTrunkResponse{value: val, isSet: true}
}

func (v NullableCreateSipTrunkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSipTrunkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
