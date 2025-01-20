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

// MessageButton List of buttons of the message.
type MessageButton struct {
	MessageAddCalendarEventButton *MessageAddCalendarEventButton
	MessageOpenUrlButton          *MessageOpenUrlButton
	MessageReplyButton            *MessageReplyButton
	MessageRequestLocationButton  *MessageRequestLocationButton
}

// MessageAddCalendarEventButtonAsMessageButton is a convenience function that returns MessageAddCalendarEventButton wrapped in MessageButton
func MessageAddCalendarEventButtonAsMessageButton(v *MessageAddCalendarEventButton) MessageButton {
	return MessageButton{
		MessageAddCalendarEventButton: v,
	}
}

// MessageOpenUrlButtonAsMessageButton is a convenience function that returns MessageOpenUrlButton wrapped in MessageButton
func MessageOpenUrlButtonAsMessageButton(v *MessageOpenUrlButton) MessageButton {
	return MessageButton{
		MessageOpenUrlButton: v,
	}
}

// MessageReplyButtonAsMessageButton is a convenience function that returns MessageReplyButton wrapped in MessageButton
func MessageReplyButtonAsMessageButton(v *MessageReplyButton) MessageButton {
	return MessageButton{
		MessageReplyButton: v,
	}
}

// MessageRequestLocationButtonAsMessageButton is a convenience function that returns MessageRequestLocationButton wrapped in MessageButton
func MessageRequestLocationButtonAsMessageButton(v *MessageRequestLocationButton) MessageButton {
	return MessageButton{
		MessageRequestLocationButton: v,
	}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MessageButton) UnmarshalJSON(data []byte) error {
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err := json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'ADD_CALENDAR_EVENT'
	if jsonDict["type"] == "ADD_CALENDAR_EVENT" {
		// try to unmarshal JSON data into MessageAddCalendarEventButton
		err = json.Unmarshal(data, &dst.MessageAddCalendarEventButton)
		if err == nil {
			jsonMessageAddCalendarEventButton, _ := json.Marshal(dst.MessageAddCalendarEventButton)
			if string(jsonMessageAddCalendarEventButton) == "{}" { // empty struct
				dst.MessageAddCalendarEventButton = nil
			} else {
				return nil // data stored in dst.MessageAddCalendarEventButton, return on the first match
			}
		} else {
			dst.MessageAddCalendarEventButton = nil
		}
	}
	// check if the discriminator value is 'OPEN_URL'
	if jsonDict["type"] == "OPEN_URL" {
		// try to unmarshal JSON data into MessageOpenUrlButton
		err = json.Unmarshal(data, &dst.MessageOpenUrlButton)
		if err == nil {
			jsonMessageOpenUrlButton, _ := json.Marshal(dst.MessageOpenUrlButton)
			if string(jsonMessageOpenUrlButton) == "{}" { // empty struct
				dst.MessageOpenUrlButton = nil
			} else {
				return nil // data stored in dst.MessageOpenUrlButton, return on the first match
			}
		} else {
			dst.MessageOpenUrlButton = nil
		}
	}
	// check if the discriminator value is 'REPLY'
	if jsonDict["type"] == "REPLY" {
		// try to unmarshal JSON data into MessageReplyButton
		err = json.Unmarshal(data, &dst.MessageReplyButton)
		if err == nil {
			jsonMessageReplyButton, _ := json.Marshal(dst.MessageReplyButton)
			if string(jsonMessageReplyButton) == "{}" { // empty struct
				dst.MessageReplyButton = nil
			} else {
				return nil // data stored in dst.MessageReplyButton, return on the first match
			}
		} else {
			dst.MessageReplyButton = nil
		}
	}
	// check if the discriminator value is 'REQUEST_LOCATION'
	if jsonDict["type"] == "REQUEST_LOCATION" {
		// try to unmarshal JSON data into MessageRequestLocationButton
		err = json.Unmarshal(data, &dst.MessageRequestLocationButton)
		if err == nil {
			jsonMessageRequestLocationButton, _ := json.Marshal(dst.MessageRequestLocationButton)
			if string(jsonMessageRequestLocationButton) == "{}" { // empty struct
				dst.MessageRequestLocationButton = nil
			} else {
				return nil // data stored in dst.MessageRequestLocationButton, return on the first match
			}
		} else {
			dst.MessageRequestLocationButton = nil
		}
	}
	return fmt.Errorf("Data failed to match schemas in anyOf(MessageButton)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageButton) MarshalJSON() ([]byte, error) {
	if src.MessageAddCalendarEventButton != nil {
		return json.Marshal(&src.MessageAddCalendarEventButton)
	}
	if src.MessageOpenUrlButton != nil {
		return json.Marshal(&src.MessageOpenUrlButton)
	}
	if src.MessageReplyButton != nil {
		return json.Marshal(&src.MessageReplyButton)
	}
	if src.MessageRequestLocationButton != nil {
		return json.Marshal(&src.MessageRequestLocationButton)
	}
	return nil, nil // no data in anyOf schemas
}

// Get the actual instance
func (obj *MessageButton) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MessageAddCalendarEventButton != nil {
		return obj.MessageAddCalendarEventButton
	}
	if obj.MessageOpenUrlButton != nil {
		return obj.MessageOpenUrlButton
	}
	if obj.MessageReplyButton != nil {
		return obj.MessageReplyButton
	}
	if obj.MessageRequestLocationButton != nil {
		return obj.MessageRequestLocationButton
	}
	// all schemas are nil
	return nil
}

type NullableMessageButton struct {
	value *MessageButton
	isSet bool
}

func (v NullableMessageButton) Get() *MessageButton {
	return v.value
}

func (v *NullableMessageButton) Set(val *MessageButton) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageButton) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageButton(val *MessageButton) *NullableMessageButton {
	return &NullableMessageButton{value: val, isSet: true}
}

func (v NullableMessageButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
