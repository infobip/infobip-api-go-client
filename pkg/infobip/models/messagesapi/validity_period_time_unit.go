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

// ValidityPeriodTimeUnit Message validity period time unit.
type ValidityPeriodTimeUnit string

// List of ValidityPeriodTimeUnit
const (
	VALIDITYPERIODTIMEUNIT_SECONDS ValidityPeriodTimeUnit = "SECONDS"
	VALIDITYPERIODTIMEUNIT_MINUTES ValidityPeriodTimeUnit = "MINUTES"
	VALIDITYPERIODTIMEUNIT_HOURS   ValidityPeriodTimeUnit = "HOURS"
)

// All allowed values of ValidityPeriodTimeUnit enum
var AllowedValidityPeriodTimeUnitEnumValues = []ValidityPeriodTimeUnit{
	"SECONDS",
	"MINUTES",
	"HOURS",
}

func (v *ValidityPeriodTimeUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ValidityPeriodTimeUnit(value)
	for _, existing := range AllowedValidityPeriodTimeUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ValidityPeriodTimeUnit", value)
}

// NewValidityPeriodTimeUnitFromValue returns a pointer to a valid ValidityPeriodTimeUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewValidityPeriodTimeUnitFromValue(v string) (*ValidityPeriodTimeUnit, error) {
	ev := ValidityPeriodTimeUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ValidityPeriodTimeUnit: valid values are %v", v, AllowedValidityPeriodTimeUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ValidityPeriodTimeUnit) IsValid() bool {
	for _, existing := range AllowedValidityPeriodTimeUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidityPeriodTimeUnit value
func (v ValidityPeriodTimeUnit) Ptr() *ValidityPeriodTimeUnit {
	return &v
}

type NullableValidityPeriodTimeUnit struct {
	value *ValidityPeriodTimeUnit
	isSet bool
}

func (v NullableValidityPeriodTimeUnit) Get() *ValidityPeriodTimeUnit {
	return v.value
}

func (v *NullableValidityPeriodTimeUnit) Set(val *ValidityPeriodTimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableValidityPeriodTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableValidityPeriodTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidityPeriodTimeUnit(val *ValidityPeriodTimeUnit) *NullableValidityPeriodTimeUnit {
	return &NullableValidityPeriodTimeUnit{value: val, isSet: true}
}

func (v NullableValidityPeriodTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidityPeriodTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}