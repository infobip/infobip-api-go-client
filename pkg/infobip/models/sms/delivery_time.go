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

// checks if the DeliveryTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryTime{}

// DeliveryTime The exact time of the day. The time is expressed in the UTC time zone.
type DeliveryTime struct {
	// Hour when the time window opens when used in the `from` property or closes when used in the `to` property.
	Hour int32
	// Minute when the time window opens when used in the `from` property or closes when used in the `to` property.
	Minute int32
}

type _DeliveryTime DeliveryTime

// NewDeliveryTime instantiates a new DeliveryTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryTime(hour int32, minute int32) *DeliveryTime {
	this := DeliveryTime{}
	this.Hour = hour
	this.Minute = minute
	return &this
}

// NewDeliveryTimeWithDefaults instantiates a new DeliveryTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryTimeWithDefaults() *DeliveryTime {
	this := DeliveryTime{}
	return &this
}

// GetHour returns the Hour field value
func (o *DeliveryTime) GetHour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hour
}

// GetHourOk returns a tuple with the Hour field value
// and a boolean to check if the value has been set.
func (o *DeliveryTime) GetHourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hour, true
}

// SetHour sets field value
func (o *DeliveryTime) SetHour(v int32) {
	o.Hour = v
}

// GetMinute returns the Minute field value
func (o *DeliveryTime) GetMinute() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value
// and a boolean to check if the value has been set.
func (o *DeliveryTime) GetMinuteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minute, true
}

// SetMinute sets field value
func (o *DeliveryTime) SetMinute(v int32) {
	o.Minute = v
}

func (o DeliveryTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hour"] = o.Hour
	toSerialize["minute"] = o.Minute
	return toSerialize, nil
}

type NullableDeliveryTime struct {
	value *DeliveryTime
	isSet bool
}

func (v NullableDeliveryTime) Get() *DeliveryTime {
	return v.value
}

func (v *NullableDeliveryTime) Set(val *DeliveryTime) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryTime) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryTime(val *DeliveryTime) *NullableDeliveryTime {
	return &NullableDeliveryTime{value: val, isSet: true}
}

func (v NullableDeliveryTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
