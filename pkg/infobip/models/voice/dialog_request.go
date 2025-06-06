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

// checks if the DialogRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DialogRequest{}

// DialogRequest struct for DialogRequest
type DialogRequest struct {
	// Unique parent call ID.
	ParentCallId     string
	ChildCallRequest *DialogCallRequest
	Recording        *DialogRecordingRequest
	// Max duration in seconds.
	MaxDuration        *int32
	PropagationOptions *DialogPropagationOptions
}

type _DialogRequest DialogRequest

// NewDialogRequest instantiates a new DialogRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDialogRequest(parentCallId string) *DialogRequest {
	this := DialogRequest{}
	this.ParentCallId = parentCallId
	var maxDuration int32 = 28800
	this.MaxDuration = &maxDuration
	return &this
}

// NewDialogRequestWithDefaults instantiates a new DialogRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDialogRequestWithDefaults() *DialogRequest {
	this := DialogRequest{}

	var maxDuration int32 = 28800
	this.MaxDuration = &maxDuration
	return &this
}

// GetParentCallId returns the ParentCallId field value
func (o *DialogRequest) GetParentCallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentCallId
}

// GetParentCallIdOk returns a tuple with the ParentCallId field value
// and a boolean to check if the value has been set.
func (o *DialogRequest) GetParentCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentCallId, true
}

// SetParentCallId sets field value
func (o *DialogRequest) SetParentCallId(v string) {
	o.ParentCallId = v
}

// GetChildCallRequest returns the ChildCallRequest field value if set, zero value otherwise.
func (o *DialogRequest) GetChildCallRequest() DialogCallRequest {
	if o == nil || IsNil(o.ChildCallRequest) {
		var ret DialogCallRequest
		return ret
	}
	return *o.ChildCallRequest
}

// GetChildCallRequestOk returns a tuple with the ChildCallRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogRequest) GetChildCallRequestOk() (*DialogCallRequest, bool) {
	if o == nil || IsNil(o.ChildCallRequest) {
		return nil, false
	}
	return o.ChildCallRequest, true
}

// HasChildCallRequest returns a boolean if a field has been set.
func (o *DialogRequest) HasChildCallRequest() bool {
	if o != nil && !IsNil(o.ChildCallRequest) {
		return true
	}

	return false
}

// SetChildCallRequest gets a reference to the given DialogCallRequest and assigns it to the ChildCallRequest field.
func (o *DialogRequest) SetChildCallRequest(v DialogCallRequest) {
	o.ChildCallRequest = &v
}

// GetRecording returns the Recording field value if set, zero value otherwise.
func (o *DialogRequest) GetRecording() DialogRecordingRequest {
	if o == nil || IsNil(o.Recording) {
		var ret DialogRecordingRequest
		return ret
	}
	return *o.Recording
}

// GetRecordingOk returns a tuple with the Recording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogRequest) GetRecordingOk() (*DialogRecordingRequest, bool) {
	if o == nil || IsNil(o.Recording) {
		return nil, false
	}
	return o.Recording, true
}

// HasRecording returns a boolean if a field has been set.
func (o *DialogRequest) HasRecording() bool {
	if o != nil && !IsNil(o.Recording) {
		return true
	}

	return false
}

// SetRecording gets a reference to the given DialogRecordingRequest and assigns it to the Recording field.
func (o *DialogRequest) SetRecording(v DialogRecordingRequest) {
	o.Recording = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *DialogRequest) GetMaxDuration() int32 {
	if o == nil || IsNil(o.MaxDuration) {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogRequest) GetMaxDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDuration) {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *DialogRequest) HasMaxDuration() bool {
	if o != nil && !IsNil(o.MaxDuration) {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
func (o *DialogRequest) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetPropagationOptions returns the PropagationOptions field value if set, zero value otherwise.
func (o *DialogRequest) GetPropagationOptions() DialogPropagationOptions {
	if o == nil || IsNil(o.PropagationOptions) {
		var ret DialogPropagationOptions
		return ret
	}
	return *o.PropagationOptions
}

// GetPropagationOptionsOk returns a tuple with the PropagationOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogRequest) GetPropagationOptionsOk() (*DialogPropagationOptions, bool) {
	if o == nil || IsNil(o.PropagationOptions) {
		return nil, false
	}
	return o.PropagationOptions, true
}

// HasPropagationOptions returns a boolean if a field has been set.
func (o *DialogRequest) HasPropagationOptions() bool {
	if o != nil && !IsNil(o.PropagationOptions) {
		return true
	}

	return false
}

// SetPropagationOptions gets a reference to the given DialogPropagationOptions and assigns it to the PropagationOptions field.
func (o *DialogRequest) SetPropagationOptions(v DialogPropagationOptions) {
	o.PropagationOptions = &v
}

func (o DialogRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DialogRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parentCallId"] = o.ParentCallId
	if !IsNil(o.ChildCallRequest) {
		toSerialize["childCallRequest"] = o.ChildCallRequest
	}
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

type NullableDialogRequest struct {
	value *DialogRequest
	isSet bool
}

func (v NullableDialogRequest) Get() *DialogRequest {
	return v.value
}

func (v *NullableDialogRequest) Set(val *DialogRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDialogRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDialogRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDialogRequest(val *DialogRequest) *NullableDialogRequest {
	return &NullableDialogRequest{value: val, isSet: true}
}

func (v NullableDialogRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDialogRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
