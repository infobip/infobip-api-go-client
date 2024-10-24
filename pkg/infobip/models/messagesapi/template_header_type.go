/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package messagesapi

import (
	"encoding/json"
	"fmt"
)

// TemplateHeaderType the model 'TemplateHeaderType'
type TemplateHeaderType string

// List of TemplateHeaderType
const (
	TEMPLATEHEADERTYPE_TEXT     TemplateHeaderType = "TEXT"
	TEMPLATEHEADERTYPE_DOCUMENT TemplateHeaderType = "DOCUMENT"
	TEMPLATEHEADERTYPE_IMAGE    TemplateHeaderType = "IMAGE"
	TEMPLATEHEADERTYPE_VIDEO    TemplateHeaderType = "VIDEO"
	TEMPLATEHEADERTYPE_LOCATION TemplateHeaderType = "LOCATION"
)

// All allowed values of TemplateHeaderType enum
var AllowedTemplateHeaderTypeEnumValues = []TemplateHeaderType{
	"TEXT",
	"DOCUMENT",
	"IMAGE",
	"VIDEO",
	"LOCATION",
}

func (v *TemplateHeaderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TemplateHeaderType(value)
	for _, existing := range AllowedTemplateHeaderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TemplateHeaderType", value)
}

// NewTemplateHeaderTypeFromValue returns a pointer to a valid TemplateHeaderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTemplateHeaderTypeFromValue(v string) (*TemplateHeaderType, error) {
	ev := TemplateHeaderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TemplateHeaderType: valid values are %v", v, AllowedTemplateHeaderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TemplateHeaderType) IsValid() bool {
	for _, existing := range AllowedTemplateHeaderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemplateHeaderType value
func (v TemplateHeaderType) Ptr() *TemplateHeaderType {
	return &v
}

type NullableTemplateHeaderType struct {
	value *TemplateHeaderType
	isSet bool
}

func (v NullableTemplateHeaderType) Get() *TemplateHeaderType {
	return v.value
}

func (v *NullableTemplateHeaderType) Set(val *TemplateHeaderType) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateHeaderType) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateHeaderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateHeaderType(val *TemplateHeaderType) *NullableTemplateHeaderType {
	return &NullableTemplateHeaderType{value: val, isSet: true}
}

func (v NullableTemplateHeaderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateHeaderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
