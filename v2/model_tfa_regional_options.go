/*
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * Contact: support@infobip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infobip

import (
	"encoding/json"
)

// TfaRegionalOptions Region-specific parameters, often imposed by local laws. Use this, if country or region that you are sending a message to requires additional information.
type TfaRegionalOptions struct {
	IndiaDlt *TfaIndiaDltOptions `json:"indiaDlt,omitempty"`
}

// NewTfaRegionalOptions instantiates a new TfaRegionalOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaRegionalOptions() *TfaRegionalOptions {
	this := TfaRegionalOptions{}
	return &this
}

// NewTfaRegionalOptionsWithDefaults instantiates a new TfaRegionalOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaRegionalOptionsWithDefaults() *TfaRegionalOptions {
	this := TfaRegionalOptions{}
	return &this
}

// GetIndiaDlt returns the IndiaDlt field value if set, zero value otherwise.
func (o *TfaRegionalOptions) GetIndiaDlt() TfaIndiaDltOptions {
	if o == nil || o.IndiaDlt == nil {
		var ret TfaIndiaDltOptions
		return ret
	}
	return *o.IndiaDlt
}

// GetIndiaDltOk returns a tuple with the IndiaDlt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaRegionalOptions) GetIndiaDltOk() (*TfaIndiaDltOptions, bool) {
	if o == nil || o.IndiaDlt == nil {
		return nil, false
	}
	return o.IndiaDlt, true
}

// HasIndiaDlt returns a boolean if a field has been set.
func (o *TfaRegionalOptions) HasIndiaDlt() bool {
	if o != nil && o.IndiaDlt != nil {
		return true
	}

	return false
}

// SetIndiaDlt gets a reference to the given TfaIndiaDltOptions and assigns it to the IndiaDlt field.
func (o *TfaRegionalOptions) SetIndiaDlt(v TfaIndiaDltOptions) {
	o.IndiaDlt = &v
}

func (o TfaRegionalOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IndiaDlt != nil {
		toSerialize["indiaDlt"] = o.IndiaDlt
	}
	return json.Marshal(toSerialize)
}

type NullableTfaRegionalOptions struct {
	value *TfaRegionalOptions
	isSet bool
}

func (v NullableTfaRegionalOptions) Get() *TfaRegionalOptions {
	return v.value
}

func (v *NullableTfaRegionalOptions) Set(val *TfaRegionalOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaRegionalOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaRegionalOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaRegionalOptions(val *TfaRegionalOptions) *NullableTfaRegionalOptions {
	return &NullableTfaRegionalOptions{value: val, isSet: true}
}

func (v NullableTfaRegionalOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaRegionalOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
