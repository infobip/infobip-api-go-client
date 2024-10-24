/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package messagesapi

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the MessageReplyButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReplyButton{}

// MessageReplyButton struct for MessageReplyButton
type MessageReplyButton struct {
	Type MessageButtonType
	// Text to be displayed on the button.
	Text string
	// Custom data that will be sent to you when user reply to a message.
	PostbackData string
}

type _MessageReplyButton MessageReplyButton

// NewMessageReplyButton instantiates a new MessageReplyButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReplyButton(text string, postbackData string) *MessageReplyButton {
	this := MessageReplyButton{}
	this.Type = MESSAGEBUTTONTYPE_REPLY
	this.Text = text
	this.PostbackData = postbackData
	return &this
}

// NewMessageReplyButtonWithDefaults instantiates a new MessageReplyButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReplyButtonWithDefaults() *MessageReplyButton {
	this := MessageReplyButton{}
	return &this
}

// GetText returns the Text field value
func (o *MessageReplyButton) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MessageReplyButton) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MessageReplyButton) SetText(v string) {
	o.Text = v
}

// GetPostbackData returns the PostbackData field value
func (o *MessageReplyButton) GetPostbackData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostbackData
}

// GetPostbackDataOk returns a tuple with the PostbackData field value
// and a boolean to check if the value has been set.
func (o *MessageReplyButton) GetPostbackDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostbackData, true
}

// SetPostbackData sets field value
func (o *MessageReplyButton) SetPostbackData(v string) {
	o.PostbackData = v
}

func (o MessageReplyButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReplyButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text"] = o.Text
	toSerialize["postbackData"] = o.PostbackData
	return toSerialize, nil
}

type NullableMessageReplyButton struct {
	value *MessageReplyButton
	isSet bool
}

func (v NullableMessageReplyButton) Get() *MessageReplyButton {
	return v.value
}

func (v *NullableMessageReplyButton) Set(val *MessageReplyButton) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReplyButton) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReplyButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReplyButton(val *MessageReplyButton) *NullableMessageReplyButton {
	return &NullableMessageReplyButton{value: val, isSet: true}
}

func (v NullableMessageReplyButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReplyButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
