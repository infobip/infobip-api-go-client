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

// checks if the DialogPropagationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DialogPropagationOptions{}

// DialogPropagationOptions struct for DialogPropagationOptions
type DialogPropagationOptions struct {
	// Flag indicating if the parent call should be terminated when the child call ends.
	ChildCallHangup *bool
	// Flag indicating if a child call's ringing should be propagated to the parent call. The parent call must be `INBOUND`. Cannot be `true` when `ringbackGeneration` is enabled.
	ChildCallRinging   *bool
	RingbackGeneration *RingbackGeneration
}

// NewDialogPropagationOptions instantiates a new DialogPropagationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDialogPropagationOptions() *DialogPropagationOptions {
	this := DialogPropagationOptions{}
	var childCallHangup bool = true
	this.ChildCallHangup = &childCallHangup
	var childCallRinging bool = false
	this.ChildCallRinging = &childCallRinging
	return &this
}

// NewDialogPropagationOptionsWithDefaults instantiates a new DialogPropagationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDialogPropagationOptionsWithDefaults() *DialogPropagationOptions {
	this := DialogPropagationOptions{}

	var childCallHangup bool = true
	this.ChildCallHangup = &childCallHangup
	var childCallRinging bool = false
	this.ChildCallRinging = &childCallRinging
	return &this
}

// GetChildCallHangup returns the ChildCallHangup field value if set, zero value otherwise.
func (o *DialogPropagationOptions) GetChildCallHangup() bool {
	if o == nil || IsNil(o.ChildCallHangup) {
		var ret bool
		return ret
	}
	return *o.ChildCallHangup
}

// GetChildCallHangupOk returns a tuple with the ChildCallHangup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogPropagationOptions) GetChildCallHangupOk() (*bool, bool) {
	if o == nil || IsNil(o.ChildCallHangup) {
		return nil, false
	}
	return o.ChildCallHangup, true
}

// HasChildCallHangup returns a boolean if a field has been set.
func (o *DialogPropagationOptions) HasChildCallHangup() bool {
	if o != nil && !IsNil(o.ChildCallHangup) {
		return true
	}

	return false
}

// SetChildCallHangup gets a reference to the given bool and assigns it to the ChildCallHangup field.
func (o *DialogPropagationOptions) SetChildCallHangup(v bool) {
	o.ChildCallHangup = &v
}

// GetChildCallRinging returns the ChildCallRinging field value if set, zero value otherwise.
func (o *DialogPropagationOptions) GetChildCallRinging() bool {
	if o == nil || IsNil(o.ChildCallRinging) {
		var ret bool
		return ret
	}
	return *o.ChildCallRinging
}

// GetChildCallRingingOk returns a tuple with the ChildCallRinging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogPropagationOptions) GetChildCallRingingOk() (*bool, bool) {
	if o == nil || IsNil(o.ChildCallRinging) {
		return nil, false
	}
	return o.ChildCallRinging, true
}

// HasChildCallRinging returns a boolean if a field has been set.
func (o *DialogPropagationOptions) HasChildCallRinging() bool {
	if o != nil && !IsNil(o.ChildCallRinging) {
		return true
	}

	return false
}

// SetChildCallRinging gets a reference to the given bool and assigns it to the ChildCallRinging field.
func (o *DialogPropagationOptions) SetChildCallRinging(v bool) {
	o.ChildCallRinging = &v
}

// GetRingbackGeneration returns the RingbackGeneration field value if set, zero value otherwise.
func (o *DialogPropagationOptions) GetRingbackGeneration() RingbackGeneration {
	if o == nil || IsNil(o.RingbackGeneration) {
		var ret RingbackGeneration
		return ret
	}
	return *o.RingbackGeneration
}

// GetRingbackGenerationOk returns a tuple with the RingbackGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogPropagationOptions) GetRingbackGenerationOk() (*RingbackGeneration, bool) {
	if o == nil || IsNil(o.RingbackGeneration) {
		return nil, false
	}
	return o.RingbackGeneration, true
}

// HasRingbackGeneration returns a boolean if a field has been set.
func (o *DialogPropagationOptions) HasRingbackGeneration() bool {
	if o != nil && !IsNil(o.RingbackGeneration) {
		return true
	}

	return false
}

// SetRingbackGeneration gets a reference to the given RingbackGeneration and assigns it to the RingbackGeneration field.
func (o *DialogPropagationOptions) SetRingbackGeneration(v RingbackGeneration) {
	o.RingbackGeneration = &v
}

func (o DialogPropagationOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DialogPropagationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChildCallHangup) {
		toSerialize["childCallHangup"] = o.ChildCallHangup
	}
	if !IsNil(o.ChildCallRinging) {
		toSerialize["childCallRinging"] = o.ChildCallRinging
	}
	if !IsNil(o.RingbackGeneration) {
		toSerialize["ringbackGeneration"] = o.RingbackGeneration
	}
	return toSerialize, nil
}

type NullableDialogPropagationOptions struct {
	value *DialogPropagationOptions
	isSet bool
}

func (v NullableDialogPropagationOptions) Get() *DialogPropagationOptions {
	return v.value
}

func (v *NullableDialogPropagationOptions) Set(val *DialogPropagationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDialogPropagationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDialogPropagationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDialogPropagationOptions(val *DialogPropagationOptions) *NullableDialogPropagationOptions {
	return &NullableDialogPropagationOptions{value: val, isSet: true}
}

func (v NullableDialogPropagationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDialogPropagationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
