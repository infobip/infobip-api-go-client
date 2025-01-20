/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voice

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the StopPlayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopPlayRequest{}

// StopPlayRequest struct for StopPlayRequest
type StopPlayRequest struct {
	// Optional parameter to update a call's custom data.
	CustomData *map[string]string
}

// NewStopPlayRequest instantiates a new StopPlayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewStopPlayRequest() *StopPlayRequest {
	this := StopPlayRequest{}
	return &this
}

// NewStopPlayRequestWithDefaults instantiates a new StopPlayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopPlayRequestWithDefaults() *StopPlayRequest {
	this := StopPlayRequest{}

	return &this
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *StopPlayRequest) GetCustomData() map[string]string {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]string
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPlayRequest) GetCustomDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomData) {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *StopPlayRequest) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]string and assigns it to the CustomData field.
func (o *StopPlayRequest) SetCustomData(v map[string]string) {
	o.CustomData = &v
}

func (o StopPlayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopPlayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	return toSerialize, nil
}

type NullableStopPlayRequest struct {
	value *StopPlayRequest
	isSet bool
}

func (v NullableStopPlayRequest) Get() *StopPlayRequest {
	return v.value
}

func (v *NullableStopPlayRequest) Set(val *StopPlayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStopPlayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStopPlayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopPlayRequest(val *StopPlayRequest) *NullableStopPlayRequest {
	return &NullableStopPlayRequest{value: val, isSet: true}
}

func (v NullableStopPlayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopPlayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
