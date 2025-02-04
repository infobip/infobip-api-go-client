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

// checks if the ConferenceComposition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceComposition{}

// ConferenceComposition struct for ConferenceComposition
type ConferenceComposition struct {
	// Indicates whether to create a single recording of all participants. If set to `True`, all participants are merged into a single audio or video file. Otherwise, all participants will have their own individual audio or video file.
	Enabled *bool
}

// NewConferenceComposition instantiates a new ConferenceComposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewConferenceComposition() *ConferenceComposition {
	this := ConferenceComposition{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewConferenceCompositionWithDefaults instantiates a new ConferenceComposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceCompositionWithDefaults() *ConferenceComposition {
	this := ConferenceComposition{}

	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConferenceComposition) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceComposition) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConferenceComposition) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConferenceComposition) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ConferenceComposition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceComposition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableConferenceComposition struct {
	value *ConferenceComposition
	isSet bool
}

func (v NullableConferenceComposition) Get() *ConferenceComposition {
	return v.value
}

func (v *NullableConferenceComposition) Set(val *ConferenceComposition) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceComposition) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceComposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceComposition(val *ConferenceComposition) *NullableConferenceComposition {
	return &NullableConferenceComposition{value: val, isSet: true}
}

func (v NullableConferenceComposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceComposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
