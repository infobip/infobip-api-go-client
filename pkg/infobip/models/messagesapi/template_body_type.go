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

// TemplateBodyType the model 'TemplateBodyType'
type TemplateBodyType string

// List of TemplateBodyType
const (
	TEMPLATEBODYTYPE_TEXT     TemplateBodyType = "TEXT"
	TEMPLATEBODYTYPE_CAROUSEL TemplateBodyType = "CAROUSEL"
)

// All allowed values of TemplateBodyType enum
var AllowedTemplateBodyTypeEnumValues = []TemplateBodyType{
	"TEXT",
	"CAROUSEL",
}

func (v *TemplateBodyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TemplateBodyType(value)
	for _, existing := range AllowedTemplateBodyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TemplateBodyType", value)
}

// NewTemplateBodyTypeFromValue returns a pointer to a valid TemplateBodyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTemplateBodyTypeFromValue(v string) (*TemplateBodyType, error) {
	ev := TemplateBodyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TemplateBodyType: valid values are %v", v, AllowedTemplateBodyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TemplateBodyType) IsValid() bool {
	for _, existing := range AllowedTemplateBodyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemplateBodyType value
func (v TemplateBodyType) Ptr() *TemplateBodyType {
	return &v
}

type NullableTemplateBodyType struct {
	value *TemplateBodyType
	isSet bool
}

func (v NullableTemplateBodyType) Get() *TemplateBodyType {
	return v.value
}

func (v *NullableTemplateBodyType) Set(val *TemplateBodyType) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateBodyType) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateBodyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateBodyType(val *TemplateBodyType) *NullableTemplateBodyType {
	return &NullableTemplateBodyType{value: val, isSet: true}
}

func (v NullableTemplateBodyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateBodyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
