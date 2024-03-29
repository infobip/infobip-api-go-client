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

// ConversionResponse struct for ConversionResponse
type ConversionResponse struct {
	// Process key assigned to account ID.
	ProcessKey *string `json:"processKey,omitempty"`
}

// NewConversionResponse instantiates a new ConversionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversionResponse() *ConversionResponse {
	this := ConversionResponse{}
	return &this
}

// NewConversionResponseWithDefaults instantiates a new ConversionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversionResponseWithDefaults() *ConversionResponse {
	this := ConversionResponse{}
	return &this
}

// GetProcessKey returns the ProcessKey field value if set, zero value otherwise.
func (o *ConversionResponse) GetProcessKey() string {
	if o == nil || o.ProcessKey == nil {
		var ret string
		return ret
	}
	return *o.ProcessKey
}

// GetProcessKeyOk returns a tuple with the ProcessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversionResponse) GetProcessKeyOk() (*string, bool) {
	if o == nil || o.ProcessKey == nil {
		return nil, false
	}
	return o.ProcessKey, true
}

// HasProcessKey returns a boolean if a field has been set.
func (o *ConversionResponse) HasProcessKey() bool {
	if o != nil && o.ProcessKey != nil {
		return true
	}

	return false
}

// SetProcessKey gets a reference to the given string and assigns it to the ProcessKey field.
func (o *ConversionResponse) SetProcessKey(v string) {
	o.ProcessKey = &v
}

func (o ConversionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProcessKey != nil {
		toSerialize["processKey"] = o.ProcessKey
	}
	return json.Marshal(toSerialize)
}

type NullableConversionResponse struct {
	value *ConversionResponse
	isSet bool
}

func (v NullableConversionResponse) Get() *ConversionResponse {
	return v.value
}

func (v *NullableConversionResponse) Set(val *ConversionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConversionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConversionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversionResponse(val *ConversionResponse) *NullableConversionResponse {
	return &NullableConversionResponse{value: val, isSet: true}
}

func (v NullableConversionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
