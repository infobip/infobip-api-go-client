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

// checks if the TemplateOpenUrlButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateOpenUrlButton{}

// TemplateOpenUrlButton struct for TemplateOpenUrlButton
type TemplateOpenUrlButton struct {
	Type TemplateButtonType
	// Extension of a URL defined in the registered template.
	Suffix string
}

type _TemplateOpenUrlButton TemplateOpenUrlButton

// NewTemplateOpenUrlButton instantiates a new TemplateOpenUrlButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateOpenUrlButton(suffix string) *TemplateOpenUrlButton {
	this := TemplateOpenUrlButton{}
	this.Type = TEMPLATEBUTTONTYPE_OPEN_URL
	this.Suffix = suffix
	return &this
}

// NewTemplateOpenUrlButtonWithDefaults instantiates a new TemplateOpenUrlButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateOpenUrlButtonWithDefaults() *TemplateOpenUrlButton {
	this := TemplateOpenUrlButton{}
	return &this
}

// GetSuffix returns the Suffix field value
func (o *TemplateOpenUrlButton) GetSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value
// and a boolean to check if the value has been set.
func (o *TemplateOpenUrlButton) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suffix, true
}

// SetSuffix sets field value
func (o *TemplateOpenUrlButton) SetSuffix(v string) {
	o.Suffix = v
}

func (o TemplateOpenUrlButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateOpenUrlButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["suffix"] = o.Suffix
	return toSerialize, nil
}

type NullableTemplateOpenUrlButton struct {
	value *TemplateOpenUrlButton
	isSet bool
}

func (v NullableTemplateOpenUrlButton) Get() *TemplateOpenUrlButton {
	return v.value
}

func (v *NullableTemplateOpenUrlButton) Set(val *TemplateOpenUrlButton) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateOpenUrlButton) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateOpenUrlButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateOpenUrlButton(val *TemplateOpenUrlButton) *NullableTemplateOpenUrlButton {
	return &NullableTemplateOpenUrlButton{value: val, isSet: true}
}

func (v NullableTemplateOpenUrlButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateOpenUrlButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}