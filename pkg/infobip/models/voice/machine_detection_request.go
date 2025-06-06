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

// checks if the MachineDetectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineDetectionRequest{}

// MachineDetectionRequest struct for MachineDetectionRequest
type MachineDetectionRequest struct {
	// Indicates whether to enable answering machine detection. If set to true, answering machine detection will generate an event indicating if the call was answered by a human or a machine.
	Enabled bool
	// Indicates maximum duration for detecting the end of the message when a answering machine is detected. If set to 0, no end of message detection will be done. Expressed in seconds.
	MessageDetectionTimeout *int32
}

type _MachineDetectionRequest MachineDetectionRequest

// NewMachineDetectionRequest instantiates a new MachineDetectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMachineDetectionRequest(enabled bool) *MachineDetectionRequest {
	this := MachineDetectionRequest{}
	this.Enabled = enabled
	return &this
}

// NewMachineDetectionRequestWithDefaults instantiates a new MachineDetectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineDetectionRequestWithDefaults() *MachineDetectionRequest {
	this := MachineDetectionRequest{}

	return &this
}

// GetEnabled returns the Enabled field value
func (o *MachineDetectionRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MachineDetectionRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MachineDetectionRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMessageDetectionTimeout returns the MessageDetectionTimeout field value if set, zero value otherwise.
func (o *MachineDetectionRequest) GetMessageDetectionTimeout() int32 {
	if o == nil || IsNil(o.MessageDetectionTimeout) {
		var ret int32
		return ret
	}
	return *o.MessageDetectionTimeout
}

// GetMessageDetectionTimeoutOk returns a tuple with the MessageDetectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDetectionRequest) GetMessageDetectionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageDetectionTimeout) {
		return nil, false
	}
	return o.MessageDetectionTimeout, true
}

// HasMessageDetectionTimeout returns a boolean if a field has been set.
func (o *MachineDetectionRequest) HasMessageDetectionTimeout() bool {
	if o != nil && !IsNil(o.MessageDetectionTimeout) {
		return true
	}

	return false
}

// SetMessageDetectionTimeout gets a reference to the given int32 and assigns it to the MessageDetectionTimeout field.
func (o *MachineDetectionRequest) SetMessageDetectionTimeout(v int32) {
	o.MessageDetectionTimeout = &v
}

func (o MachineDetectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineDetectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.MessageDetectionTimeout) {
		toSerialize["messageDetectionTimeout"] = o.MessageDetectionTimeout
	}
	return toSerialize, nil
}

type NullableMachineDetectionRequest struct {
	value *MachineDetectionRequest
	isSet bool
}

func (v NullableMachineDetectionRequest) Get() *MachineDetectionRequest {
	return v.value
}

func (v *NullableMachineDetectionRequest) Set(val *MachineDetectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineDetectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineDetectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineDetectionRequest(val *MachineDetectionRequest) *NullableMachineDetectionRequest {
	return &NullableMachineDetectionRequest{value: val, isSet: true}
}

func (v NullableMachineDetectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineDetectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
