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

// SelectionStrategy Strategy for SIP trunk host selection.
type SelectionStrategy string

// List of SelectionStrategy
const (
	SELECTIONSTRATEGY_FAILOVER    SelectionStrategy = "FAILOVER"
	SELECTIONSTRATEGY_ROUND_ROBIN SelectionStrategy = "ROUND_ROBIN"
)

// All allowed values of SelectionStrategy enum
var AllowedSelectionStrategyEnumValues = []SelectionStrategy{
	"FAILOVER",
	"ROUND_ROBIN",
}

func (v *SelectionStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelectionStrategy(value)
	for _, existing := range AllowedSelectionStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelectionStrategy", value)
}

// NewSelectionStrategyFromValue returns a pointer to a valid SelectionStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelectionStrategyFromValue(v string) (*SelectionStrategy, error) {
	ev := SelectionStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelectionStrategy: valid values are %v", v, AllowedSelectionStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelectionStrategy) IsValid() bool {
	for _, existing := range AllowedSelectionStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelectionStrategy value
func (v SelectionStrategy) Ptr() *SelectionStrategy {
	return &v
}

type NullableSelectionStrategy struct {
	value *SelectionStrategy
	isSet bool
}

func (v NullableSelectionStrategy) Get() *SelectionStrategy {
	return v.value
}

func (v *NullableSelectionStrategy) Set(val *SelectionStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectionStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectionStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectionStrategy(val *SelectionStrategy) *NullableSelectionStrategy {
	return &NullableSelectionStrategy{value: val, isSet: true}
}

func (v NullableSelectionStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectionStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
