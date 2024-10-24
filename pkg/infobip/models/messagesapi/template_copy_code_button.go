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

// checks if the TemplateCopyCodeButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateCopyCodeButton{}

// TemplateCopyCodeButton struct for TemplateCopyCodeButton
type TemplateCopyCodeButton struct {
	Type TemplateButtonType
	// Coupon code.
	Code string
}

type _TemplateCopyCodeButton TemplateCopyCodeButton

// NewTemplateCopyCodeButton instantiates a new TemplateCopyCodeButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCopyCodeButton(code string) *TemplateCopyCodeButton {
	this := TemplateCopyCodeButton{}
	this.Type = TEMPLATEBUTTONTYPE_COPY_CODE
	this.Code = code
	return &this
}

// NewTemplateCopyCodeButtonWithDefaults instantiates a new TemplateCopyCodeButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCopyCodeButtonWithDefaults() *TemplateCopyCodeButton {
	this := TemplateCopyCodeButton{}
	return &this
}

// GetCode returns the Code field value
func (o *TemplateCopyCodeButton) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TemplateCopyCodeButton) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TemplateCopyCodeButton) SetCode(v string) {
	o.Code = v
}

func (o TemplateCopyCodeButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateCopyCodeButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

type NullableTemplateCopyCodeButton struct {
	value *TemplateCopyCodeButton
	isSet bool
}

func (v NullableTemplateCopyCodeButton) Get() *TemplateCopyCodeButton {
	return v.value
}

func (v *NullableTemplateCopyCodeButton) Set(val *TemplateCopyCodeButton) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCopyCodeButton) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCopyCodeButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCopyCodeButton(val *TemplateCopyCodeButton) *NullableTemplateCopyCodeButton {
	return &NullableTemplateCopyCodeButton{value: val, isSet: true}
}

func (v NullableTemplateCopyCodeButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCopyCodeButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
