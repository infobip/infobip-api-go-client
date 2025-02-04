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

// RecordingType Recording type.
type RecordingType string

// List of RecordingType
const (
	RECORDINGTYPE_AUDIO           RecordingType = "AUDIO"
	RECORDINGTYPE_AUDIO_AND_VIDEO RecordingType = "AUDIO_AND_VIDEO"
)

// All allowed values of RecordingType enum
var AllowedRecordingTypeEnumValues = []RecordingType{
	"AUDIO",
	"AUDIO_AND_VIDEO",
}

func (v *RecordingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecordingType(value)
	for _, existing := range AllowedRecordingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordingType", value)
}

// NewRecordingTypeFromValue returns a pointer to a valid RecordingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecordingTypeFromValue(v string) (*RecordingType, error) {
	ev := RecordingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordingType: valid values are %v", v, AllowedRecordingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecordingType) IsValid() bool {
	for _, existing := range AllowedRecordingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecordingType value
func (v RecordingType) Ptr() *RecordingType {
	return &v
}

type NullableRecordingType struct {
	value *RecordingType
	isSet bool
}

func (v NullableRecordingType) Get() *RecordingType {
	return v.value
}

func (v *NullableRecordingType) Set(val *RecordingType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingType(val *RecordingType) *NullableRecordingType {
	return &NullableRecordingType{value: val, isSet: true}
}

func (v NullableRecordingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
