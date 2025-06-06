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

// MessageGeneralStatus Status group name that describes which category the status code belongs to, i.e., [PENDING](https://www.infobip.com/docs/essentials/response-status-and-error-codes#pending-general-status-codes), [UNDELIVERABLE](https://www.infobip.com/docs/essentials/response-status-and-error-codes#undeliverable-general-status-codes), [DELIVERED](https://www.infobip.com/docs/essentials/response-status-and-error-codes#delivered-general-status-codes), [EXPIRED](https://www.infobip.com/docs/essentials/response-status-and-error-codes#expired-general-status-codes), [REJECTED](https://www.infobip.com/docs/essentials/response-status-and-error-codes#rejected-general-status-codes).
type MessageGeneralStatus string

// List of MessageGeneralStatus
const (
	MESSAGEGENERALSTATUS_ACCEPTED      MessageGeneralStatus = "ACCEPTED"
	MESSAGEGENERALSTATUS_PENDING       MessageGeneralStatus = "PENDING"
	MESSAGEGENERALSTATUS_UNDELIVERABLE MessageGeneralStatus = "UNDELIVERABLE"
	MESSAGEGENERALSTATUS_DELIVERED     MessageGeneralStatus = "DELIVERED"
	MESSAGEGENERALSTATUS_EXPIRED       MessageGeneralStatus = "EXPIRED"
	MESSAGEGENERALSTATUS_REJECTED      MessageGeneralStatus = "REJECTED"
)

// All allowed values of MessageGeneralStatus enum
var AllowedMessageGeneralStatusEnumValues = []MessageGeneralStatus{
	"ACCEPTED",
	"PENDING",
	"UNDELIVERABLE",
	"DELIVERED",
	"EXPIRED",
	"REJECTED",
}

func (v *MessageGeneralStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageGeneralStatus(value)
	for _, existing := range AllowedMessageGeneralStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageGeneralStatus", value)
}

// NewMessageGeneralStatusFromValue returns a pointer to a valid MessageGeneralStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageGeneralStatusFromValue(v string) (*MessageGeneralStatus, error) {
	ev := MessageGeneralStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageGeneralStatus: valid values are %v", v, AllowedMessageGeneralStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageGeneralStatus) IsValid() bool {
	for _, existing := range AllowedMessageGeneralStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageGeneralStatus value
func (v MessageGeneralStatus) Ptr() *MessageGeneralStatus {
	return &v
}

type NullableMessageGeneralStatus struct {
	value *MessageGeneralStatus
	isSet bool
}

func (v NullableMessageGeneralStatus) Get() *MessageGeneralStatus {
	return v.value
}

func (v *NullableMessageGeneralStatus) Set(val *MessageGeneralStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageGeneralStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageGeneralStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageGeneralStatus(val *MessageGeneralStatus) *NullableMessageGeneralStatus {
	return &NullableMessageGeneralStatus{value: val, isSet: true}
}

func (v NullableMessageGeneralStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageGeneralStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
