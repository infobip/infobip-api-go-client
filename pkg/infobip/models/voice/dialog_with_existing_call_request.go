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

// checks if the DialogWithExistingCallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DialogWithExistingCallRequest{}

// DialogWithExistingCallRequest struct for DialogWithExistingCallRequest
type DialogWithExistingCallRequest struct {
	Recording *DialogRecordingRequest
	// Max duration in seconds.
	MaxDuration        *int32
	PropagationOptions *DialogPropagationOptions
}

// NewDialogWithExistingCallRequest instantiates a new DialogWithExistingCallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDialogWithExistingCallRequest() *DialogWithExistingCallRequest {
	this := DialogWithExistingCallRequest{}
	var maxDuration int32 = 28800
	this.MaxDuration = &maxDuration
	return &this
}

// NewDialogWithExistingCallRequestWithDefaults instantiates a new DialogWithExistingCallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDialogWithExistingCallRequestWithDefaults() *DialogWithExistingCallRequest {
	this := DialogWithExistingCallRequest{}

	var maxDuration int32 = 28800
	this.MaxDuration = &maxDuration
	return &this
}

// GetRecording returns the Recording field value if set, zero value otherwise.
func (o *DialogWithExistingCallRequest) GetRecording() DialogRecordingRequest {
	if o == nil || IsNil(o.Recording) {
		var ret DialogRecordingRequest
		return ret
	}
	return *o.Recording
}

// GetRecordingOk returns a tuple with the Recording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogWithExistingCallRequest) GetRecordingOk() (*DialogRecordingRequest, bool) {
	if o == nil || IsNil(o.Recording) {
		return nil, false
	}
	return o.Recording, true
}

// HasRecording returns a boolean if a field has been set.
func (o *DialogWithExistingCallRequest) HasRecording() bool {
	if o != nil && !IsNil(o.Recording) {
		return true
	}

	return false
}

// SetRecording gets a reference to the given DialogRecordingRequest and assigns it to the Recording field.
func (o *DialogWithExistingCallRequest) SetRecording(v DialogRecordingRequest) {
	o.Recording = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *DialogWithExistingCallRequest) GetMaxDuration() int32 {
	if o == nil || IsNil(o.MaxDuration) {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogWithExistingCallRequest) GetMaxDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDuration) {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *DialogWithExistingCallRequest) HasMaxDuration() bool {
	if o != nil && !IsNil(o.MaxDuration) {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
func (o *DialogWithExistingCallRequest) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetPropagationOptions returns the PropagationOptions field value if set, zero value otherwise.
func (o *DialogWithExistingCallRequest) GetPropagationOptions() DialogPropagationOptions {
	if o == nil || IsNil(o.PropagationOptions) {
		var ret DialogPropagationOptions
		return ret
	}
	return *o.PropagationOptions
}

// GetPropagationOptionsOk returns a tuple with the PropagationOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogWithExistingCallRequest) GetPropagationOptionsOk() (*DialogPropagationOptions, bool) {
	if o == nil || IsNil(o.PropagationOptions) {
		return nil, false
	}
	return o.PropagationOptions, true
}

// HasPropagationOptions returns a boolean if a field has been set.
func (o *DialogWithExistingCallRequest) HasPropagationOptions() bool {
	if o != nil && !IsNil(o.PropagationOptions) {
		return true
	}

	return false
}

// SetPropagationOptions gets a reference to the given DialogPropagationOptions and assigns it to the PropagationOptions field.
func (o *DialogWithExistingCallRequest) SetPropagationOptions(v DialogPropagationOptions) {
	o.PropagationOptions = &v
}

func (o DialogWithExistingCallRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DialogWithExistingCallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recording) {
		toSerialize["recording"] = o.Recording
	}
	if !IsNil(o.MaxDuration) {
		toSerialize["maxDuration"] = o.MaxDuration
	}
	if !IsNil(o.PropagationOptions) {
		toSerialize["propagationOptions"] = o.PropagationOptions
	}
	return toSerialize, nil
}

type NullableDialogWithExistingCallRequest struct {
	value *DialogWithExistingCallRequest
	isSet bool
}

func (v NullableDialogWithExistingCallRequest) Get() *DialogWithExistingCallRequest {
	return v.value
}

func (v *NullableDialogWithExistingCallRequest) Set(val *DialogWithExistingCallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDialogWithExistingCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDialogWithExistingCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDialogWithExistingCallRequest(val *DialogWithExistingCallRequest) *NullableDialogWithExistingCallRequest {
	return &NullableDialogWithExistingCallRequest{value: val, isSet: true}
}

func (v NullableDialogWithExistingCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDialogWithExistingCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
