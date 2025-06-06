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

// checks if the MessageRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageRequestOptions{}

// MessageRequestOptions Options applicable to all messages in the request.
type MessageRequestOptions struct {
	Schedule *RequestSchedulingSettings
	Tracking *UrlOptions
	// Set to true to return smsCount in the response. Default is false. smsCount is the total count of SMS submitted in the request. SMS messages have a character limit and messages longer than that limit will be split into multiple SMS and reflected in the total count of SMS submitted.
	IncludeSmsCountInResponse *bool
	ConversionTracking        *Tracking
}

// NewMessageRequestOptions instantiates a new MessageRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMessageRequestOptions() *MessageRequestOptions {
	this := MessageRequestOptions{}
	var includeSmsCountInResponse bool = false
	this.IncludeSmsCountInResponse = &includeSmsCountInResponse
	return &this
}

// NewMessageRequestOptionsWithDefaults instantiates a new MessageRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageRequestOptionsWithDefaults() *MessageRequestOptions {
	this := MessageRequestOptions{}

	var includeSmsCountInResponse bool = false
	this.IncludeSmsCountInResponse = &includeSmsCountInResponse
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *MessageRequestOptions) GetSchedule() RequestSchedulingSettings {
	if o == nil || IsNil(o.Schedule) {
		var ret RequestSchedulingSettings
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRequestOptions) GetScheduleOk() (*RequestSchedulingSettings, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *MessageRequestOptions) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given RequestSchedulingSettings and assigns it to the Schedule field.
func (o *MessageRequestOptions) SetSchedule(v RequestSchedulingSettings) {
	o.Schedule = &v
}

// GetTracking returns the Tracking field value if set, zero value otherwise.
func (o *MessageRequestOptions) GetTracking() UrlOptions {
	if o == nil || IsNil(o.Tracking) {
		var ret UrlOptions
		return ret
	}
	return *o.Tracking
}

// GetTrackingOk returns a tuple with the Tracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRequestOptions) GetTrackingOk() (*UrlOptions, bool) {
	if o == nil || IsNil(o.Tracking) {
		return nil, false
	}
	return o.Tracking, true
}

// HasTracking returns a boolean if a field has been set.
func (o *MessageRequestOptions) HasTracking() bool {
	if o != nil && !IsNil(o.Tracking) {
		return true
	}

	return false
}

// SetTracking gets a reference to the given UrlOptions and assigns it to the Tracking field.
func (o *MessageRequestOptions) SetTracking(v UrlOptions) {
	o.Tracking = &v
}

// GetIncludeSmsCountInResponse returns the IncludeSmsCountInResponse field value if set, zero value otherwise.
func (o *MessageRequestOptions) GetIncludeSmsCountInResponse() bool {
	if o == nil || IsNil(o.IncludeSmsCountInResponse) {
		var ret bool
		return ret
	}
	return *o.IncludeSmsCountInResponse
}

// GetIncludeSmsCountInResponseOk returns a tuple with the IncludeSmsCountInResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRequestOptions) GetIncludeSmsCountInResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSmsCountInResponse) {
		return nil, false
	}
	return o.IncludeSmsCountInResponse, true
}

// HasIncludeSmsCountInResponse returns a boolean if a field has been set.
func (o *MessageRequestOptions) HasIncludeSmsCountInResponse() bool {
	if o != nil && !IsNil(o.IncludeSmsCountInResponse) {
		return true
	}

	return false
}

// SetIncludeSmsCountInResponse gets a reference to the given bool and assigns it to the IncludeSmsCountInResponse field.
func (o *MessageRequestOptions) SetIncludeSmsCountInResponse(v bool) {
	o.IncludeSmsCountInResponse = &v
}

// GetConversionTracking returns the ConversionTracking field value if set, zero value otherwise.
func (o *MessageRequestOptions) GetConversionTracking() Tracking {
	if o == nil || IsNil(o.ConversionTracking) {
		var ret Tracking
		return ret
	}
	return *o.ConversionTracking
}

// GetConversionTrackingOk returns a tuple with the ConversionTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRequestOptions) GetConversionTrackingOk() (*Tracking, bool) {
	if o == nil || IsNil(o.ConversionTracking) {
		return nil, false
	}
	return o.ConversionTracking, true
}

// HasConversionTracking returns a boolean if a field has been set.
func (o *MessageRequestOptions) HasConversionTracking() bool {
	if o != nil && !IsNil(o.ConversionTracking) {
		return true
	}

	return false
}

// SetConversionTracking gets a reference to the given Tracking and assigns it to the ConversionTracking field.
func (o *MessageRequestOptions) SetConversionTracking(v Tracking) {
	o.ConversionTracking = &v
}

func (o MessageRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Tracking) {
		toSerialize["tracking"] = o.Tracking
	}
	if !IsNil(o.IncludeSmsCountInResponse) {
		toSerialize["includeSmsCountInResponse"] = o.IncludeSmsCountInResponse
	}
	if !IsNil(o.ConversionTracking) {
		toSerialize["conversionTracking"] = o.ConversionTracking
	}
	return toSerialize, nil
}

type NullableMessageRequestOptions struct {
	value *MessageRequestOptions
	isSet bool
}

func (v NullableMessageRequestOptions) Get() *MessageRequestOptions {
	return v.value
}

func (v *NullableMessageRequestOptions) Set(val *MessageRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageRequestOptions(val *MessageRequestOptions) *NullableMessageRequestOptions {
	return &NullableMessageRequestOptions{value: val, isSet: true}
}

func (v NullableMessageRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
