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

// checks if the MessageTextBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageTextBody{}

// MessageTextBody struct for MessageTextBody
type MessageTextBody struct {
	Type MessageBodyType
	// Text to be sent.
	Text string
}

type _MessageTextBody MessageTextBody

// NewMessageTextBody instantiates a new MessageTextBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageTextBody(text string) *MessageTextBody {
	this := MessageTextBody{}
	this.Type = MESSAGEBODYTYPE_TEXT
	this.Text = text
	return &this
}

// NewMessageTextBodyWithDefaults instantiates a new MessageTextBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageTextBodyWithDefaults() *MessageTextBody {
	this := MessageTextBody{}
	return &this
}

// GetText returns the Text field value
func (o *MessageTextBody) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MessageTextBody) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MessageTextBody) SetText(v string) {
	o.Text = v
}

func (o MessageTextBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageTextBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableMessageTextBody struct {
	value *MessageTextBody
	isSet bool
}

func (v NullableMessageTextBody) Get() *MessageTextBody {
	return v.value
}

func (v *NullableMessageTextBody) Set(val *MessageTextBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageTextBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageTextBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageTextBody(val *MessageTextBody) *NullableMessageTextBody {
	return &NullableMessageTextBody{value: val, isSet: true}
}

func (v NullableMessageTextBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageTextBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
