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

// checks if the EventOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventOptions{}

// EventOptions Event options.
type EventOptions struct {
	Platform       *Platform
	ValidityPeriod *ValidityPeriod
}

// NewEventOptions instantiates a new EventOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewEventOptions() *EventOptions {
	this := EventOptions{}
	return &this
}

// NewEventOptionsWithDefaults instantiates a new EventOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventOptionsWithDefaults() *EventOptions {
	this := EventOptions{}

	return &this
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *EventOptions) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOptions) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *EventOptions) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *EventOptions) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *EventOptions) GetValidityPeriod() ValidityPeriod {
	if o == nil || IsNil(o.ValidityPeriod) {
		var ret ValidityPeriod
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventOptions) GetValidityPeriodOk() (*ValidityPeriod, bool) {
	if o == nil || IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *EventOptions) HasValidityPeriod() bool {
	if o != nil && !IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given ValidityPeriod and assigns it to the ValidityPeriod field.
func (o *EventOptions) SetValidityPeriod(v ValidityPeriod) {
	o.ValidityPeriod = &v
}

func (o EventOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	return toSerialize, nil
}

type NullableEventOptions struct {
	value *EventOptions
	isSet bool
}

func (v NullableEventOptions) Get() *EventOptions {
	return v.value
}

func (v *NullableEventOptions) Set(val *EventOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableEventOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableEventOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventOptions(val *EventOptions) *NullableEventOptions {
	return &NullableEventOptions{value: val, isSet: true}
}

func (v NullableEventOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
