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

// checks if the Retry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Retry{}

// Retry Used to define if the delivery of the Voice messages should be retried in case the first try doesn't succeed. Additional retries will be made according to the schedule defined by _minPeriod_ and _maxPeriod_ parameters and platform's internal retry logic. If the _minPeriod_ differs _maxPeriod_, delivery will be retried in the following manner: after 1 min, 2 min, 5 min, 10 min, 20 min, 30 min, 1 hour, 2 hours, 4 hours, 8 hours, 16 hours, 24 hours or until maxPeriod is reached.  If the retry attempt for the _maxPeriod_ is reached, the _maxPeriod_ will be used for all subsequent retries. If the _minPeriod_ and the _maxPeriod_ are defined as equal values, the period of time between retries will be equal to this value. Message delivery will be retried until the successful delivery or message validity or _maxCount_ value is reached.
type Retry struct {
	// Specify the maximum number of retry attempts. Maximum value of the maxCount is `4`. If the value is higher than `4`, it will be set to `4`.
	MaxCount int32
	// Defines the maximum waiting time (in minutes) after the previous failed attempt to try to deliver the message again. Supported values are 1 min, 2 min, 5 min, 10 min, 20 min, 30 min, 1 hour, 2 hours, 4 hours, 8 hours, 16 hours, 24 hours. If entered a value that isn't from the previous list but that's smaller than 24 hours, it is used the next bigger value from the list. If entered value that is bigger than 24 hours than it is used 24 hours.
	MaxPeriod int32
	// Defines the minimal waiting time (in minutes) after the previous failed attempt to try to deliver the message again. Supported values are 1 min, 2 min, 5 min, 10 min, 20 min, 30 min, 1 hour, 2 hours, 4 hours, 8 hours, 16 hours, 24 hours. If entered a value that isn't from the previous list but that's smaller than 24 hours, it is used the next bigger value from the list. If entered value that is bigger than 24 hours than it is used 24 hours.
	MinPeriod int32
}

type _Retry Retry

// NewRetry instantiates a new Retry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewRetry(maxCount int32, maxPeriod int32, minPeriod int32) *Retry {
	this := Retry{}
	this.MaxCount = maxCount
	this.MaxPeriod = maxPeriod
	this.MinPeriod = minPeriod
	return &this
}

// NewRetryWithDefaults instantiates a new Retry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryWithDefaults() *Retry {
	this := Retry{}

	return &this
}

// GetMaxCount returns the MaxCount field value
func (o *Retry) GetMaxCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value
// and a boolean to check if the value has been set.
func (o *Retry) GetMaxCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCount, true
}

// SetMaxCount sets field value
func (o *Retry) SetMaxCount(v int32) {
	o.MaxCount = v
}

// GetMaxPeriod returns the MaxPeriod field value
func (o *Retry) GetMaxPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPeriod
}

// GetMaxPeriodOk returns a tuple with the MaxPeriod field value
// and a boolean to check if the value has been set.
func (o *Retry) GetMaxPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPeriod, true
}

// SetMaxPeriod sets field value
func (o *Retry) SetMaxPeriod(v int32) {
	o.MaxPeriod = v
}

// GetMinPeriod returns the MinPeriod field value
func (o *Retry) GetMinPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinPeriod
}

// GetMinPeriodOk returns a tuple with the MinPeriod field value
// and a boolean to check if the value has been set.
func (o *Retry) GetMinPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPeriod, true
}

// SetMinPeriod sets field value
func (o *Retry) SetMinPeriod(v int32) {
	o.MinPeriod = v
}

func (o Retry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Retry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxCount"] = o.MaxCount
	toSerialize["maxPeriod"] = o.MaxPeriod
	toSerialize["minPeriod"] = o.MinPeriod
	return toSerialize, nil
}

type NullableRetry struct {
	value *Retry
	isSet bool
}

func (v NullableRetry) Get() *Retry {
	return v.value
}

func (v *NullableRetry) Set(val *Retry) {
	v.value = val
	v.isSet = true
}

func (v NullableRetry) IsSet() bool {
	return v.isSet
}

func (v *NullableRetry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetry(val *Retry) *NullableRetry {
	return &NullableRetry{value: val, isSet: true}
}

func (v NullableRetry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
