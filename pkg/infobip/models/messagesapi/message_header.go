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

// MessageHeader Header of the message.
type MessageHeader struct {
	MessageTextHeader *MessageTextHeader
}

// MessageTextHeaderAsMessageHeader is a convenience function that returns MessageTextHeader wrapped in MessageHeader
func MessageTextHeaderAsMessageHeader(v *MessageTextHeader) MessageHeader {
	return MessageHeader{
		MessageTextHeader: v,
	}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MessageHeader) UnmarshalJSON(data []byte) error {
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err := json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'TEXT'
	if jsonDict["type"] == "TEXT" {
		// try to unmarshal JSON data into MessageTextHeader
		err = json.Unmarshal(data, &dst.MessageTextHeader)
		if err == nil {
			jsonMessageTextHeader, _ := json.Marshal(dst.MessageTextHeader)
			if string(jsonMessageTextHeader) == "{}" { // empty struct
				dst.MessageTextHeader = nil
			} else {
				return nil // data stored in dst.MessageTextHeader, return on the first match
			}
		} else {
			dst.MessageTextHeader = nil
		}
	}
	return fmt.Errorf("Data failed to match schemas in anyOf(MessageHeader)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageHeader) MarshalJSON() ([]byte, error) {
	if src.MessageTextHeader != nil {
		return json.Marshal(&src.MessageTextHeader)
	}
	return nil, nil // no data in anyOf schemas
}

// Get the actual instance
func (obj *MessageHeader) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MessageTextHeader != nil {
		return obj.MessageTextHeader
	}
	// all schemas are nil
	return nil
}

type NullableMessageHeader struct {
	value *MessageHeader
	isSet bool
}

func (v NullableMessageHeader) Get() *MessageHeader {
	return v.value
}

func (v *NullableMessageHeader) Set(val *MessageHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageHeader(val *MessageHeader) *NullableMessageHeader {
	return &NullableMessageHeader{value: val, isSet: true}
}

func (v NullableMessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
