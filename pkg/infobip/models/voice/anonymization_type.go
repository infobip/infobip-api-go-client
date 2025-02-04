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

// AnonymizationType Preferred presentation of anonymized calls.
type AnonymizationType string

// List of AnonymizationType
const (
	ANONYMIZATIONTYPE_NONE               AnonymizationType = "NONE"
	ANONYMIZATIONTYPE_REMOTE_PARTY_ID    AnonymizationType = "REMOTE_PARTY_ID"
	ANONYMIZATIONTYPE_ASSERTED_IDENTITY  AnonymizationType = "ASSERTED_IDENTITY"
	ANONYMIZATIONTYPE_PREFERRED_IDENTITY AnonymizationType = "PREFERRED_IDENTITY"
)

// All allowed values of AnonymizationType enum
var AllowedAnonymizationTypeEnumValues = []AnonymizationType{
	"NONE",
	"REMOTE_PARTY_ID",
	"ASSERTED_IDENTITY",
	"PREFERRED_IDENTITY",
}

func (v *AnonymizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnonymizationType(value)
	for _, existing := range AllowedAnonymizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnonymizationType", value)
}

// NewAnonymizationTypeFromValue returns a pointer to a valid AnonymizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnonymizationTypeFromValue(v string) (*AnonymizationType, error) {
	ev := AnonymizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnonymizationType: valid values are %v", v, AllowedAnonymizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnonymizationType) IsValid() bool {
	for _, existing := range AllowedAnonymizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnonymizationType value
func (v AnonymizationType) Ptr() *AnonymizationType {
	return &v
}

type NullableAnonymizationType struct {
	value *AnonymizationType
	isSet bool
}

func (v NullableAnonymizationType) Get() *AnonymizationType {
	return v.value
}

func (v *NullableAnonymizationType) Set(val *AnonymizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymizationType(val *AnonymizationType) *NullableAnonymizationType {
	return &NullableAnonymizationType{value: val, isSet: true}
}

func (v NullableAnonymizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
