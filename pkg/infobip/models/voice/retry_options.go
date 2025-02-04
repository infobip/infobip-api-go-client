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

// checks if the RetryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetryOptions{}

// RetryOptions Used to determine whether to retry the delivery of a bulk call if the initial attempt fails. Additional retries will occur based on the schedule defined by the _minWaitPeriod_ and _maxWaitPeriod_ parameters, as well as the platform's internal retry logic. If _minWaitPeriod_ differs from _maxWaitPeriod_, the delivery will be retried according to the following schedule: after 1 minute, 2 minutes, 5 minutes, 10 minutes, 20 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, 8 hours, 16 hours, and 24 hours, or until _maxWaitPeriod_ is reached. Once the retry attempt for _maxWaitPeriod_ is reached, _maxWaitPeriod_ will be used for all subsequent retries. If _minWaitPeriod_ and _maxWaitPeriod_ are defined as equal values, the period between retries will be equal to this value. Bulk call delivery will be retried until successful delivery, call validity expiration, or reaching the _maxAttempts_ value.
type RetryOptions struct {
	// Defines the minimum waiting time (in minutes) after the previous failed attempt before trying to establish the call again. Supported values are 1 minute, 2 minutes, 5 minutes, 10 minutes, 20 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, 8 hours, 16 hours, and 24 hours. If entered a value that is not in the list but is smaller than 24 hours, the next bigger value from the list will be used. If a value larger than 24 hours is entered, 24 hours will be used.
	MinWaitPeriod *int32
	// Defines the maximum waiting time (in minutes) after the previous failed attempt before trying to establish the call again. Supported values are 1 minute, 2 minutes, 5 minutes, 10 minutes, 20 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, 8 hours, 16 hours, and 24 hours. If entered a value that is not in the list but is smaller than 24 hours, the next bigger value from the list will be used. If a value larger than 24 hours is entered, 24 hours will be used.
	MaxWaitPeriod *int32
	// Defines the maximum number of retry attempts. The maximum value is `4`. If a value higher than `4` is entered, it will be set to `4`.
	MaxAttempts *int32
}

// NewRetryOptions instantiates a new RetryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewRetryOptions() *RetryOptions {
	this := RetryOptions{}
	return &this
}

// NewRetryOptionsWithDefaults instantiates a new RetryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryOptionsWithDefaults() *RetryOptions {
	this := RetryOptions{}

	return &this
}

// GetMinWaitPeriod returns the MinWaitPeriod field value if set, zero value otherwise.
func (o *RetryOptions) GetMinWaitPeriod() int32 {
	if o == nil || IsNil(o.MinWaitPeriod) {
		var ret int32
		return ret
	}
	return *o.MinWaitPeriod
}

// GetMinWaitPeriodOk returns a tuple with the MinWaitPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryOptions) GetMinWaitPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.MinWaitPeriod) {
		return nil, false
	}
	return o.MinWaitPeriod, true
}

// HasMinWaitPeriod returns a boolean if a field has been set.
func (o *RetryOptions) HasMinWaitPeriod() bool {
	if o != nil && !IsNil(o.MinWaitPeriod) {
		return true
	}

	return false
}

// SetMinWaitPeriod gets a reference to the given int32 and assigns it to the MinWaitPeriod field.
func (o *RetryOptions) SetMinWaitPeriod(v int32) {
	o.MinWaitPeriod = &v
}

// GetMaxWaitPeriod returns the MaxWaitPeriod field value if set, zero value otherwise.
func (o *RetryOptions) GetMaxWaitPeriod() int32 {
	if o == nil || IsNil(o.MaxWaitPeriod) {
		var ret int32
		return ret
	}
	return *o.MaxWaitPeriod
}

// GetMaxWaitPeriodOk returns a tuple with the MaxWaitPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryOptions) GetMaxWaitPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxWaitPeriod) {
		return nil, false
	}
	return o.MaxWaitPeriod, true
}

// HasMaxWaitPeriod returns a boolean if a field has been set.
func (o *RetryOptions) HasMaxWaitPeriod() bool {
	if o != nil && !IsNil(o.MaxWaitPeriod) {
		return true
	}

	return false
}

// SetMaxWaitPeriod gets a reference to the given int32 and assigns it to the MaxWaitPeriod field.
func (o *RetryOptions) SetMaxWaitPeriod(v int32) {
	o.MaxWaitPeriod = &v
}

// GetMaxAttempts returns the MaxAttempts field value if set, zero value otherwise.
func (o *RetryOptions) GetMaxAttempts() int32 {
	if o == nil || IsNil(o.MaxAttempts) {
		var ret int32
		return ret
	}
	return *o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryOptions) GetMaxAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAttempts) {
		return nil, false
	}
	return o.MaxAttempts, true
}

// HasMaxAttempts returns a boolean if a field has been set.
func (o *RetryOptions) HasMaxAttempts() bool {
	if o != nil && !IsNil(o.MaxAttempts) {
		return true
	}

	return false
}

// SetMaxAttempts gets a reference to the given int32 and assigns it to the MaxAttempts field.
func (o *RetryOptions) SetMaxAttempts(v int32) {
	o.MaxAttempts = &v
}

func (o RetryOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinWaitPeriod) {
		toSerialize["minWaitPeriod"] = o.MinWaitPeriod
	}
	if !IsNil(o.MaxWaitPeriod) {
		toSerialize["maxWaitPeriod"] = o.MaxWaitPeriod
	}
	if !IsNil(o.MaxAttempts) {
		toSerialize["maxAttempts"] = o.MaxAttempts
	}
	return toSerialize, nil
}

type NullableRetryOptions struct {
	value *RetryOptions
	isSet bool
}

func (v NullableRetryOptions) Get() *RetryOptions {
	return v.value
}

func (v *NullableRetryOptions) Set(val *RetryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryOptions(val *RetryOptions) *NullableRetryOptions {
	return &NullableRetryOptions{value: val, isSet: true}
}

func (v NullableRetryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
