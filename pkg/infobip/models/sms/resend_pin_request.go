/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sms

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the ResendPinRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResendPinRequest{}

// ResendPinRequest struct for ResendPinRequest
type ResendPinRequest struct {
	// Key value pairs that will be replaced during message sending. Placeholder keys should NOT contain curly brackets and should NOT contain a `pin` placeholder. Valid example: `\"placeholders\":{\"firstName\":\"John\"}`
	Placeholders *map[string]string
}

// NewResendPinRequest instantiates a new ResendPinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResendPinRequest() *ResendPinRequest {
	this := ResendPinRequest{}
	return &this
}

// NewResendPinRequestWithDefaults instantiates a new ResendPinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResendPinRequestWithDefaults() *ResendPinRequest {
	this := ResendPinRequest{}
	return &this
}

// GetPlaceholders returns the Placeholders field value if set, zero value otherwise.
func (o *ResendPinRequest) GetPlaceholders() map[string]string {
	if o == nil || IsNil(o.Placeholders) {
		var ret map[string]string
		return ret
	}
	return *o.Placeholders
}

// GetPlaceholdersOk returns a tuple with the Placeholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResendPinRequest) GetPlaceholdersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Placeholders) {
		return nil, false
	}
	return o.Placeholders, true
}

// HasPlaceholders returns a boolean if a field has been set.
func (o *ResendPinRequest) HasPlaceholders() bool {
	if o != nil && !IsNil(o.Placeholders) {
		return true
	}

	return false
}

// SetPlaceholders gets a reference to the given map[string]string and assigns it to the Placeholders field.
func (o *ResendPinRequest) SetPlaceholders(v map[string]string) {
	o.Placeholders = &v
}

func (o ResendPinRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResendPinRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Placeholders) {
		toSerialize["placeholders"] = o.Placeholders
	}
	return toSerialize, nil
}

type NullableResendPinRequest struct {
	value *ResendPinRequest
	isSet bool
}

func (v NullableResendPinRequest) Get() *ResendPinRequest {
	return v.value
}

func (v *NullableResendPinRequest) Set(val *ResendPinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResendPinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResendPinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResendPinRequest(val *ResendPinRequest) *NullableResendPinRequest {
	return &NullableResendPinRequest{value: val, isSet: true}
}

func (v NullableResendPinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResendPinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}