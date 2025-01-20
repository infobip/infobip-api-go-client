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

// BulkEndpointType the model 'BulkEndpointType'
type BulkEndpointType string

// List of BulkEndpointType
const (
	BULKENDPOINTTYPE_PHONE BulkEndpointType = "PHONE"
)

// All allowed values of BulkEndpointType enum
var AllowedBulkEndpointTypeEnumValues = []BulkEndpointType{
	"PHONE",
}

func (v *BulkEndpointType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkEndpointType(value)
	for _, existing := range AllowedBulkEndpointTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkEndpointType", value)
}

// NewBulkEndpointTypeFromValue returns a pointer to a valid BulkEndpointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkEndpointTypeFromValue(v string) (*BulkEndpointType, error) {
	ev := BulkEndpointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkEndpointType: valid values are %v", v, AllowedBulkEndpointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkEndpointType) IsValid() bool {
	for _, existing := range AllowedBulkEndpointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BulkEndpointType value
func (v BulkEndpointType) Ptr() *BulkEndpointType {
	return &v
}

type NullableBulkEndpointType struct {
	value *BulkEndpointType
	isSet bool
}

func (v NullableBulkEndpointType) Get() *BulkEndpointType {
	return v.value
}

func (v *NullableBulkEndpointType) Set(val *BulkEndpointType) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkEndpointType) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkEndpointType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkEndpointType(val *BulkEndpointType) *NullableBulkEndpointType {
	return &NullableBulkEndpointType{value: val, isSet: true}
}

func (v NullableBulkEndpointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkEndpointType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
