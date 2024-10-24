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

// checks if the Tracking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tracking{}

// Tracking Allows you to set up tracking parameters to track conversion metrics. For more details on SMS Conversion, see: [Track Conversion](https://www.infobip.com/docs/sms/api#track-conversion).
type Tracking struct {
	// Indicates if a message has to be tracked for conversion rates. Default \"false\".
	UseConversionTracking *bool
	// Sets a custom conversion type naming convention, e.g. `ONE_TIME_PIN` or `SOCIAL_INVITES`.
	ConversionTrackingName *string
}

// NewTracking instantiates a new Tracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTracking() *Tracking {
	this := Tracking{}
	return &this
}

// NewTrackingWithDefaults instantiates a new Tracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingWithDefaults() *Tracking {
	this := Tracking{}
	return &this
}

// GetUseConversionTracking returns the UseConversionTracking field value if set, zero value otherwise.
func (o *Tracking) GetUseConversionTracking() bool {
	if o == nil || IsNil(o.UseConversionTracking) {
		var ret bool
		return ret
	}
	return *o.UseConversionTracking
}

// GetUseConversionTrackingOk returns a tuple with the UseConversionTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetUseConversionTrackingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseConversionTracking) {
		return nil, false
	}
	return o.UseConversionTracking, true
}

// HasUseConversionTracking returns a boolean if a field has been set.
func (o *Tracking) HasUseConversionTracking() bool {
	if o != nil && !IsNil(o.UseConversionTracking) {
		return true
	}

	return false
}

// SetUseConversionTracking gets a reference to the given bool and assigns it to the UseConversionTracking field.
func (o *Tracking) SetUseConversionTracking(v bool) {
	o.UseConversionTracking = &v
}

// GetConversionTrackingName returns the ConversionTrackingName field value if set, zero value otherwise.
func (o *Tracking) GetConversionTrackingName() string {
	if o == nil || IsNil(o.ConversionTrackingName) {
		var ret string
		return ret
	}
	return *o.ConversionTrackingName
}

// GetConversionTrackingNameOk returns a tuple with the ConversionTrackingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tracking) GetConversionTrackingNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConversionTrackingName) {
		return nil, false
	}
	return o.ConversionTrackingName, true
}

// HasConversionTrackingName returns a boolean if a field has been set.
func (o *Tracking) HasConversionTrackingName() bool {
	if o != nil && !IsNil(o.ConversionTrackingName) {
		return true
	}

	return false
}

// SetConversionTrackingName gets a reference to the given string and assigns it to the ConversionTrackingName field.
func (o *Tracking) SetConversionTrackingName(v string) {
	o.ConversionTrackingName = &v
}

func (o Tracking) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseConversionTracking) {
		toSerialize["useConversionTracking"] = o.UseConversionTracking
	}
	if !IsNil(o.ConversionTrackingName) {
		toSerialize["conversionTrackingName"] = o.ConversionTrackingName
	}
	return toSerialize, nil
}

type NullableTracking struct {
	value *Tracking
	isSet bool
}

func (v NullableTracking) Get() *Tracking {
	return v.value
}

func (v *NullableTracking) Set(val *Tracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTracking(val *Tracking) *NullableTracking {
	return &NullableTracking{value: val, isSet: true}
}

func (v NullableTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}