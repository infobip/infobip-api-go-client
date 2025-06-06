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

// MessageCardOptionsOrientationEnum Orientation of the card.
type MessageCardOptionsOrientationEnum string

// List of MessageCardOptionsOrientationEnum
const (
	MESSAGECARDOPTIONSORIENTATIONENUM_HORIZONTAL MessageCardOptionsOrientationEnum = "HORIZONTAL"
	MESSAGECARDOPTIONSORIENTATIONENUM_VERTICAL   MessageCardOptionsOrientationEnum = "VERTICAL"
)

// All allowed values of MessageCardOptionsOrientationEnum enum
var AllowedMessageCardOptionsOrientationEnumEnumValues = []MessageCardOptionsOrientationEnum{
	"HORIZONTAL",
	"VERTICAL",
}

func (v *MessageCardOptionsOrientationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageCardOptionsOrientationEnum(value)
	for _, existing := range AllowedMessageCardOptionsOrientationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageCardOptionsOrientationEnum", value)
}

// NewMessageCardOptionsOrientationEnumFromValue returns a pointer to a valid MessageCardOptionsOrientationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageCardOptionsOrientationEnumFromValue(v string) (*MessageCardOptionsOrientationEnum, error) {
	ev := MessageCardOptionsOrientationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageCardOptionsOrientationEnum: valid values are %v", v, AllowedMessageCardOptionsOrientationEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageCardOptionsOrientationEnum) IsValid() bool {
	for _, existing := range AllowedMessageCardOptionsOrientationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageCardOptionsOrientationEnum value
func (v MessageCardOptionsOrientationEnum) Ptr() *MessageCardOptionsOrientationEnum {
	return &v
}

type NullableMessageCardOptionsOrientationEnum struct {
	value *MessageCardOptionsOrientationEnum
	isSet bool
}

func (v NullableMessageCardOptionsOrientationEnum) Get() *MessageCardOptionsOrientationEnum {
	return v.value
}

func (v *NullableMessageCardOptionsOrientationEnum) Set(val *MessageCardOptionsOrientationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCardOptionsOrientationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCardOptionsOrientationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCardOptionsOrientationEnum(val *MessageCardOptionsOrientationEnum) *NullableMessageCardOptionsOrientationEnum {
	return &NullableMessageCardOptionsOrientationEnum{value: val, isSet: true}
}

func (v NullableMessageCardOptionsOrientationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCardOptionsOrientationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
