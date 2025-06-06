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

// DeliveryDay Allowed day.
type DeliveryDay string

// List of DeliveryDay
const (
	DELIVERYDAY_MONDAY    DeliveryDay = "MONDAY"
	DELIVERYDAY_TUESDAY   DeliveryDay = "TUESDAY"
	DELIVERYDAY_WEDNESDAY DeliveryDay = "WEDNESDAY"
	DELIVERYDAY_THURSDAY  DeliveryDay = "THURSDAY"
	DELIVERYDAY_FRIDAY    DeliveryDay = "FRIDAY"
	DELIVERYDAY_SATURDAY  DeliveryDay = "SATURDAY"
	DELIVERYDAY_SUNDAY    DeliveryDay = "SUNDAY"
)

// All allowed values of DeliveryDay enum
var AllowedDeliveryDayEnumValues = []DeliveryDay{
	"MONDAY",
	"TUESDAY",
	"WEDNESDAY",
	"THURSDAY",
	"FRIDAY",
	"SATURDAY",
	"SUNDAY",
}

func (v *DeliveryDay) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryDay(value)
	for _, existing := range AllowedDeliveryDayEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryDay", value)
}

// NewDeliveryDayFromValue returns a pointer to a valid DeliveryDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryDayFromValue(v string) (*DeliveryDay, error) {
	ev := DeliveryDay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryDay: valid values are %v", v, AllowedDeliveryDayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryDay) IsValid() bool {
	for _, existing := range AllowedDeliveryDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliveryDay value
func (v DeliveryDay) Ptr() *DeliveryDay {
	return &v
}

type NullableDeliveryDay struct {
	value *DeliveryDay
	isSet bool
}

func (v NullableDeliveryDay) Get() *DeliveryDay {
	return v.value
}

func (v *NullableDeliveryDay) Set(val *DeliveryDay) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryDay) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryDay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryDay(val *DeliveryDay) *NullableDeliveryDay {
	return &NullableDeliveryDay{value: val, isSet: true}
}

func (v NullableDeliveryDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryDay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
