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

// TagIdentifierType Defines which identifier will be provided in the tags field. Possible values are ID and NAME. If not provided, default value of NAME is used.
type TagIdentifierType string

// List of TagIdentifierType
const (
	TAGIDENTIFIERTYPE_ID   TagIdentifierType = "ID"
	TAGIDENTIFIERTYPE_NAME TagIdentifierType = "NAME"
)

// All allowed values of TagIdentifierType enum
var AllowedTagIdentifierTypeEnumValues = []TagIdentifierType{
	"ID",
	"NAME",
}

func (v *TagIdentifierType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TagIdentifierType(value)
	for _, existing := range AllowedTagIdentifierTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TagIdentifierType", value)
}

// NewTagIdentifierTypeFromValue returns a pointer to a valid TagIdentifierType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTagIdentifierTypeFromValue(v string) (*TagIdentifierType, error) {
	ev := TagIdentifierType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TagIdentifierType: valid values are %v", v, AllowedTagIdentifierTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TagIdentifierType) IsValid() bool {
	for _, existing := range AllowedTagIdentifierTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagIdentifierType value
func (v TagIdentifierType) Ptr() *TagIdentifierType {
	return &v
}

type NullableTagIdentifierType struct {
	value *TagIdentifierType
	isSet bool
}

func (v NullableTagIdentifierType) Get() *TagIdentifierType {
	return v.value
}

func (v *NullableTagIdentifierType) Set(val *TagIdentifierType) {
	v.value = val
	v.isSet = true
}

func (v NullableTagIdentifierType) IsSet() bool {
	return v.isSet
}

func (v *NullableTagIdentifierType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagIdentifierType(val *TagIdentifierType) *NullableTagIdentifierType {
	return &NullableTagIdentifierType{value: val, isSet: true}
}

func (v NullableTagIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagIdentifierType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
