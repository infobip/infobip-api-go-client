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

// checks if the TemplateQuickReplyButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateQuickReplyButton{}

// TemplateQuickReplyButton struct for TemplateQuickReplyButton
type TemplateQuickReplyButton struct {
	Type TemplateButtonType
	// Custom client data that will be included in a user's response.
	PostbackData string
}

type _TemplateQuickReplyButton TemplateQuickReplyButton

// NewTemplateQuickReplyButton instantiates a new TemplateQuickReplyButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateQuickReplyButton(postbackData string) *TemplateQuickReplyButton {
	this := TemplateQuickReplyButton{}
	this.Type = TEMPLATEBUTTONTYPE_QUICK_REPLY
	this.PostbackData = postbackData
	return &this
}

// NewTemplateQuickReplyButtonWithDefaults instantiates a new TemplateQuickReplyButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateQuickReplyButtonWithDefaults() *TemplateQuickReplyButton {
	this := TemplateQuickReplyButton{}
	return &this
}

// GetPostbackData returns the PostbackData field value
func (o *TemplateQuickReplyButton) GetPostbackData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostbackData
}

// GetPostbackDataOk returns a tuple with the PostbackData field value
// and a boolean to check if the value has been set.
func (o *TemplateQuickReplyButton) GetPostbackDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostbackData, true
}

// SetPostbackData sets field value
func (o *TemplateQuickReplyButton) SetPostbackData(v string) {
	o.PostbackData = v
}

func (o TemplateQuickReplyButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateQuickReplyButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["postbackData"] = o.PostbackData
	return toSerialize, nil
}

type NullableTemplateQuickReplyButton struct {
	value *TemplateQuickReplyButton
	isSet bool
}

func (v NullableTemplateQuickReplyButton) Get() *TemplateQuickReplyButton {
	return v.value
}

func (v *NullableTemplateQuickReplyButton) Set(val *TemplateQuickReplyButton) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateQuickReplyButton) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateQuickReplyButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateQuickReplyButton(val *TemplateQuickReplyButton) *NullableTemplateQuickReplyButton {
	return &NullableTemplateQuickReplyButton{value: val, isSet: true}
}

func (v NullableTemplateQuickReplyButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateQuickReplyButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}