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

// PlayContentType the model 'PlayContentType'
type PlayContentType string

// List of PlayContentType
const (
	PLAYCONTENTTYPE_FILE      PlayContentType = "FILE"
	PLAYCONTENTTYPE_URL       PlayContentType = "URL"
	PLAYCONTENTTYPE_RECORDING PlayContentType = "RECORDING"
	PLAYCONTENTTYPE_TEXT      PlayContentType = "TEXT"
)

// All allowed values of PlayContentType enum
var AllowedPlayContentTypeEnumValues = []PlayContentType{
	"FILE",
	"URL",
	"RECORDING",
	"TEXT",
}

func (v *PlayContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlayContentType(value)
	for _, existing := range AllowedPlayContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlayContentType", value)
}

// NewPlayContentTypeFromValue returns a pointer to a valid PlayContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlayContentTypeFromValue(v string) (*PlayContentType, error) {
	ev := PlayContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlayContentType: valid values are %v", v, AllowedPlayContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlayContentType) IsValid() bool {
	for _, existing := range AllowedPlayContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayContentType value
func (v PlayContentType) Ptr() *PlayContentType {
	return &v
}

type NullablePlayContentType struct {
	value *PlayContentType
	isSet bool
}

func (v NullablePlayContentType) Get() *PlayContentType {
	return v.value
}

func (v *NullablePlayContentType) Set(val *PlayContentType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayContentType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayContentType(val *PlayContentType) *NullablePlayContentType {
	return &NullablePlayContentType{value: val, isSet: true}
}

func (v NullablePlayContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
