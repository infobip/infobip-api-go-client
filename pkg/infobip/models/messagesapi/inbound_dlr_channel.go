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

// InboundDlrChannel Messaging channel used for sending an event.
type InboundDlrChannel string

// List of InboundDlrChannel
const (
	INBOUNDDLRCHANNEL_SMS          InboundDlrChannel = "SMS"
	INBOUNDDLRCHANNEL_MMS          InboundDlrChannel = "MMS"
	INBOUNDDLRCHANNEL_WHATSAPP     InboundDlrChannel = "WHATSAPP"
	INBOUNDDLRCHANNEL_VIBER_BM     InboundDlrChannel = "VIBER_BM"
	INBOUNDDLRCHANNEL_VIBER_BOT    InboundDlrChannel = "VIBER_BOT"
	INBOUNDDLRCHANNEL_RCS          InboundDlrChannel = "RCS"
	INBOUNDDLRCHANNEL_APPLE_MB     InboundDlrChannel = "APPLE_MB"
	INBOUNDDLRCHANNEL_INSTAGRAM_DM InboundDlrChannel = "INSTAGRAM_DM"
	INBOUNDDLRCHANNEL_LINE_ON      InboundDlrChannel = "LINE_ON"
	INBOUNDDLRCHANNEL_MESSENGER    InboundDlrChannel = "MESSENGER"
)

// All allowed values of InboundDlrChannel enum
var AllowedInboundDlrChannelEnumValues = []InboundDlrChannel{
	"SMS",
	"MMS",
	"WHATSAPP",
	"VIBER_BM",
	"VIBER_BOT",
	"RCS",
	"APPLE_MB",
	"INSTAGRAM_DM",
	"LINE_ON",
	"MESSENGER",
}

func (v *InboundDlrChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InboundDlrChannel(value)
	for _, existing := range AllowedInboundDlrChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InboundDlrChannel", value)
}

// NewInboundDlrChannelFromValue returns a pointer to a valid InboundDlrChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInboundDlrChannelFromValue(v string) (*InboundDlrChannel, error) {
	ev := InboundDlrChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InboundDlrChannel: valid values are %v", v, AllowedInboundDlrChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InboundDlrChannel) IsValid() bool {
	for _, existing := range AllowedInboundDlrChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboundDlrChannel value
func (v InboundDlrChannel) Ptr() *InboundDlrChannel {
	return &v
}

type NullableInboundDlrChannel struct {
	value *InboundDlrChannel
	isSet bool
}

func (v NullableInboundDlrChannel) Get() *InboundDlrChannel {
	return v.value
}

func (v *NullableInboundDlrChannel) Set(val *InboundDlrChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundDlrChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundDlrChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundDlrChannel(val *InboundDlrChannel) *NullableInboundDlrChannel {
	return &NullableInboundDlrChannel{value: val, isSet: true}
}

func (v NullableInboundDlrChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundDlrChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
