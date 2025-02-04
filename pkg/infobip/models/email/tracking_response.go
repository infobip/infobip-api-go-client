/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the TrackingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingResponse{}

// TrackingResponse Tracking details of the domain.
type TrackingResponse struct {
	// Indicates whether tracking of clicks is enabled.
	Clicks *bool
	// Indicates whether tracking of opens is enabled.
	Opens *bool
	// Indicates whether tracking of unsubscribes is enabled.
	Unsubscribe *bool
}

// NewTrackingResponse instantiates a new TrackingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewTrackingResponse() *TrackingResponse {
	this := TrackingResponse{}
	return &this
}

// NewTrackingResponseWithDefaults instantiates a new TrackingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingResponseWithDefaults() *TrackingResponse {
	this := TrackingResponse{}

	return &this
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *TrackingResponse) GetClicks() bool {
	if o == nil || IsNil(o.Clicks) {
		var ret bool
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingResponse) GetClicksOk() (*bool, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *TrackingResponse) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given bool and assigns it to the Clicks field.
func (o *TrackingResponse) SetClicks(v bool) {
	o.Clicks = &v
}

// GetOpens returns the Opens field value if set, zero value otherwise.
func (o *TrackingResponse) GetOpens() bool {
	if o == nil || IsNil(o.Opens) {
		var ret bool
		return ret
	}
	return *o.Opens
}

// GetOpensOk returns a tuple with the Opens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingResponse) GetOpensOk() (*bool, bool) {
	if o == nil || IsNil(o.Opens) {
		return nil, false
	}
	return o.Opens, true
}

// HasOpens returns a boolean if a field has been set.
func (o *TrackingResponse) HasOpens() bool {
	if o != nil && !IsNil(o.Opens) {
		return true
	}

	return false
}

// SetOpens gets a reference to the given bool and assigns it to the Opens field.
func (o *TrackingResponse) SetOpens(v bool) {
	o.Opens = &v
}

// GetUnsubscribe returns the Unsubscribe field value if set, zero value otherwise.
func (o *TrackingResponse) GetUnsubscribe() bool {
	if o == nil || IsNil(o.Unsubscribe) {
		var ret bool
		return ret
	}
	return *o.Unsubscribe
}

// GetUnsubscribeOk returns a tuple with the Unsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingResponse) GetUnsubscribeOk() (*bool, bool) {
	if o == nil || IsNil(o.Unsubscribe) {
		return nil, false
	}
	return o.Unsubscribe, true
}

// HasUnsubscribe returns a boolean if a field has been set.
func (o *TrackingResponse) HasUnsubscribe() bool {
	if o != nil && !IsNil(o.Unsubscribe) {
		return true
	}

	return false
}

// SetUnsubscribe gets a reference to the given bool and assigns it to the Unsubscribe field.
func (o *TrackingResponse) SetUnsubscribe(v bool) {
	o.Unsubscribe = &v
}

func (o TrackingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.Opens) {
		toSerialize["opens"] = o.Opens
	}
	if !IsNil(o.Unsubscribe) {
		toSerialize["unsubscribe"] = o.Unsubscribe
	}
	return toSerialize, nil
}

type NullableTrackingResponse struct {
	value *TrackingResponse
	isSet bool
}

func (v NullableTrackingResponse) Get() *TrackingResponse {
	return v.value
}

func (v *NullableTrackingResponse) Set(val *TrackingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingResponse(val *TrackingResponse) *NullableTrackingResponse {
	return &NullableTrackingResponse{value: val, isSet: true}
}

func (v NullableTrackingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
