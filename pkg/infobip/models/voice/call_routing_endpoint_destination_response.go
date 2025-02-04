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

// checks if the CallRoutingEndpointDestinationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRoutingEndpointDestinationResponse{}

// CallRoutingEndpointDestinationResponse struct for CallRoutingEndpointDestinationResponse
type CallRoutingEndpointDestinationResponse struct {
	Type  CallRoutingDestinationType
	Value CallRoutingEndpoint
	// Time to wait, in seconds, to establish a call toward the destination endpoint. The call will be terminated if it is not answered within the specified time.
	ConnectTimeout *int32
	Recording      *CallRoutingRecording
}

type _CallRoutingEndpointDestinationResponse CallRoutingEndpointDestinationResponse

// NewCallRoutingEndpointDestinationResponse instantiates a new CallRoutingEndpointDestinationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallRoutingEndpointDestinationResponse(value CallRoutingEndpoint) *CallRoutingEndpointDestinationResponse {
	this := CallRoutingEndpointDestinationResponse{}
	this.Type = "ENDPOINT"
	this.Value = value
	return &this
}

// NewCallRoutingEndpointDestinationResponseWithDefaults instantiates a new CallRoutingEndpointDestinationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRoutingEndpointDestinationResponseWithDefaults() *CallRoutingEndpointDestinationResponse {
	this := CallRoutingEndpointDestinationResponse{}
	this.Type = "ENDPOINT"
	return &this
}

// GetValue returns the Value field value
func (o *CallRoutingEndpointDestinationResponse) GetValue() CallRoutingEndpoint {
	if o == nil {
		var ret CallRoutingEndpoint
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CallRoutingEndpointDestinationResponse) GetValueOk() (*CallRoutingEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CallRoutingEndpointDestinationResponse) SetValue(v CallRoutingEndpoint) {
	o.Value = v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *CallRoutingEndpointDestinationResponse) GetConnectTimeout() int32 {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret int32
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingEndpointDestinationResponse) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *CallRoutingEndpointDestinationResponse) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given int32 and assigns it to the ConnectTimeout field.
func (o *CallRoutingEndpointDestinationResponse) SetConnectTimeout(v int32) {
	o.ConnectTimeout = &v
}

// GetRecording returns the Recording field value if set, zero value otherwise.
func (o *CallRoutingEndpointDestinationResponse) GetRecording() CallRoutingRecording {
	if o == nil || IsNil(o.Recording) {
		var ret CallRoutingRecording
		return ret
	}
	return *o.Recording
}

// GetRecordingOk returns a tuple with the Recording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingEndpointDestinationResponse) GetRecordingOk() (*CallRoutingRecording, bool) {
	if o == nil || IsNil(o.Recording) {
		return nil, false
	}
	return o.Recording, true
}

// HasRecording returns a boolean if a field has been set.
func (o *CallRoutingEndpointDestinationResponse) HasRecording() bool {
	if o != nil && !IsNil(o.Recording) {
		return true
	}

	return false
}

// SetRecording gets a reference to the given CallRoutingRecording and assigns it to the Recording field.
func (o *CallRoutingEndpointDestinationResponse) SetRecording(v CallRoutingRecording) {
	o.Recording = &v
}

func (o CallRoutingEndpointDestinationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRoutingEndpointDestinationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.Recording) {
		toSerialize["recording"] = o.Recording
	}
	return toSerialize, nil
}

type NullableCallRoutingEndpointDestinationResponse struct {
	value *CallRoutingEndpointDestinationResponse
	isSet bool
}

func (v NullableCallRoutingEndpointDestinationResponse) Get() *CallRoutingEndpointDestinationResponse {
	return v.value
}

func (v *NullableCallRoutingEndpointDestinationResponse) Set(val *CallRoutingEndpointDestinationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRoutingEndpointDestinationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRoutingEndpointDestinationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRoutingEndpointDestinationResponse(val *CallRoutingEndpointDestinationResponse) *NullableCallRoutingEndpointDestinationResponse {
	return &NullableCallRoutingEndpointDestinationResponse{value: val, isSet: true}
}

func (v NullableCallRoutingEndpointDestinationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRoutingEndpointDestinationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
