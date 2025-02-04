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

// CallRoutingRecordingType Recording type.
type CallRoutingRecordingType string

// List of CallRoutingRecordingType
const (
	CALLROUTINGRECORDINGTYPE_AUDIO           CallRoutingRecordingType = "AUDIO"
	CALLROUTINGRECORDINGTYPE_AUDIO_AND_VIDEO CallRoutingRecordingType = "AUDIO_AND_VIDEO"
)

// All allowed values of CallRoutingRecordingType enum
var AllowedCallRoutingRecordingTypeEnumValues = []CallRoutingRecordingType{
	"AUDIO",
	"AUDIO_AND_VIDEO",
}

func (v *CallRoutingRecordingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CallRoutingRecordingType(value)
	for _, existing := range AllowedCallRoutingRecordingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CallRoutingRecordingType", value)
}

// NewCallRoutingRecordingTypeFromValue returns a pointer to a valid CallRoutingRecordingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCallRoutingRecordingTypeFromValue(v string) (*CallRoutingRecordingType, error) {
	ev := CallRoutingRecordingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CallRoutingRecordingType: valid values are %v", v, AllowedCallRoutingRecordingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CallRoutingRecordingType) IsValid() bool {
	for _, existing := range AllowedCallRoutingRecordingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CallRoutingRecordingType value
func (v CallRoutingRecordingType) Ptr() *CallRoutingRecordingType {
	return &v
}

type NullableCallRoutingRecordingType struct {
	value *CallRoutingRecordingType
	isSet bool
}

func (v NullableCallRoutingRecordingType) Get() *CallRoutingRecordingType {
	return v.value
}

func (v *NullableCallRoutingRecordingType) Set(val *CallRoutingRecordingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRoutingRecordingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRoutingRecordingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRoutingRecordingType(val *CallRoutingRecordingType) *NullableCallRoutingRecordingType {
	return &NullableCallRoutingRecordingType{value: val, isSet: true}
}

func (v NullableCallRoutingRecordingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRoutingRecordingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
