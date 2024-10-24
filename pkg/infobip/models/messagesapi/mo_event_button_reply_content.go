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

// checks if the MoEventButtonReplyContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoEventButtonReplyContent{}

// MoEventButtonReplyContent struct for MoEventButtonReplyContent
type MoEventButtonReplyContent struct {
	Type MoEventContentType
	// Title of the selected button.
	Text string
	// Payload of the selected button.
	Payload *string
	// Identifier of the selected button.
	Id *string
}

type _MoEventButtonReplyContent MoEventButtonReplyContent

// NewMoEventButtonReplyContent instantiates a new MoEventButtonReplyContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoEventButtonReplyContent(text string) *MoEventButtonReplyContent {
	this := MoEventButtonReplyContent{}
	this.Type = MOEVENTCONTENTTYPE_BUTTON_REPLY
	this.Text = text
	return &this
}

// NewMoEventButtonReplyContentWithDefaults instantiates a new MoEventButtonReplyContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoEventButtonReplyContentWithDefaults() *MoEventButtonReplyContent {
	this := MoEventButtonReplyContent{}
	return &this
}

// GetText returns the Text field value
func (o *MoEventButtonReplyContent) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MoEventButtonReplyContent) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MoEventButtonReplyContent) SetText(v string) {
	o.Text = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *MoEventButtonReplyContent) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoEventButtonReplyContent) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *MoEventButtonReplyContent) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *MoEventButtonReplyContent) SetPayload(v string) {
	o.Payload = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MoEventButtonReplyContent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoEventButtonReplyContent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MoEventButtonReplyContent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MoEventButtonReplyContent) SetId(v string) {
	o.Id = &v
}

func (o MoEventButtonReplyContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoEventButtonReplyContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text"] = o.Text
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableMoEventButtonReplyContent struct {
	value *MoEventButtonReplyContent
	isSet bool
}

func (v NullableMoEventButtonReplyContent) Get() *MoEventButtonReplyContent {
	return v.value
}

func (v *NullableMoEventButtonReplyContent) Set(val *MoEventButtonReplyContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMoEventButtonReplyContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMoEventButtonReplyContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoEventButtonReplyContent(val *MoEventButtonReplyContent) *NullableMoEventButtonReplyContent {
	return &NullableMoEventButtonReplyContent{value: val, isSet: true}
}

func (v NullableMoEventButtonReplyContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoEventButtonReplyContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}