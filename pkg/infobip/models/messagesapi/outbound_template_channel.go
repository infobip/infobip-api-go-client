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

// OutboundTemplateChannel Messaging channel used for sending a message.
type OutboundTemplateChannel string

// List of OutboundTemplateChannel
const (
	OUTBOUNDTEMPLATECHANNEL_APPLE_MB OutboundTemplateChannel = "APPLE_MB"
	OUTBOUNDTEMPLATECHANNEL_RCS      OutboundTemplateChannel = "RCS"
	OUTBOUNDTEMPLATECHANNEL_WHATSAPP OutboundTemplateChannel = "WHATSAPP"
)

// All allowed values of OutboundTemplateChannel enum
var AllowedOutboundTemplateChannelEnumValues = []OutboundTemplateChannel{
	"APPLE_MB",
	"RCS",
	"WHATSAPP",
}

func (v *OutboundTemplateChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutboundTemplateChannel(value)
	for _, existing := range AllowedOutboundTemplateChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutboundTemplateChannel", value)
}

// NewOutboundTemplateChannelFromValue returns a pointer to a valid OutboundTemplateChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutboundTemplateChannelFromValue(v string) (*OutboundTemplateChannel, error) {
	ev := OutboundTemplateChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutboundTemplateChannel: valid values are %v", v, AllowedOutboundTemplateChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutboundTemplateChannel) IsValid() bool {
	for _, existing := range AllowedOutboundTemplateChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutboundTemplateChannel value
func (v OutboundTemplateChannel) Ptr() *OutboundTemplateChannel {
	return &v
}

type NullableOutboundTemplateChannel struct {
	value *OutboundTemplateChannel
	isSet bool
}

func (v NullableOutboundTemplateChannel) Get() *OutboundTemplateChannel {
	return v.value
}

func (v *NullableOutboundTemplateChannel) Set(val *OutboundTemplateChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundTemplateChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundTemplateChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundTemplateChannel(val *OutboundTemplateChannel) *NullableOutboundTemplateChannel {
	return &NullableOutboundTemplateChannel{value: val, isSet: true}
}

func (v NullableOutboundTemplateChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundTemplateChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
