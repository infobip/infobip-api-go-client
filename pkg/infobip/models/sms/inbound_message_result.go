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

// checks if the InboundMessageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundMessageResult{}

// InboundMessageResult struct for InboundMessageResult
type InboundMessageResult struct {
	// The number of messages returned in the `results` array.
	MessageCount *int32
	// The number of messages that have not been pulled in.
	PendingMessageCount *int32
	// An array of result objects.
	Results []InboundMessage
}

// NewInboundMessageResult instantiates a new InboundMessageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewInboundMessageResult() *InboundMessageResult {
	this := InboundMessageResult{}
	return &this
}

// NewInboundMessageResultWithDefaults instantiates a new InboundMessageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundMessageResultWithDefaults() *InboundMessageResult {
	this := InboundMessageResult{}

	return &this
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *InboundMessageResult) GetMessageCount() int32 {
	if o == nil || IsNil(o.MessageCount) {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessageResult) GetMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageCount) {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *InboundMessageResult) HasMessageCount() bool {
	if o != nil && !IsNil(o.MessageCount) {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *InboundMessageResult) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetPendingMessageCount returns the PendingMessageCount field value if set, zero value otherwise.
func (o *InboundMessageResult) GetPendingMessageCount() int32 {
	if o == nil || IsNil(o.PendingMessageCount) {
		var ret int32
		return ret
	}
	return *o.PendingMessageCount
}

// GetPendingMessageCountOk returns a tuple with the PendingMessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessageResult) GetPendingMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PendingMessageCount) {
		return nil, false
	}
	return o.PendingMessageCount, true
}

// HasPendingMessageCount returns a boolean if a field has been set.
func (o *InboundMessageResult) HasPendingMessageCount() bool {
	if o != nil && !IsNil(o.PendingMessageCount) {
		return true
	}

	return false
}

// SetPendingMessageCount gets a reference to the given int32 and assigns it to the PendingMessageCount field.
func (o *InboundMessageResult) SetPendingMessageCount(v int32) {
	o.PendingMessageCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InboundMessageResult) GetResults() []InboundMessage {
	if o == nil || IsNil(o.Results) {
		var ret []InboundMessage
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessageResult) GetResultsOk() ([]InboundMessage, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InboundMessageResult) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InboundMessage and assigns it to the Results field.
func (o *InboundMessageResult) SetResults(v []InboundMessage) {
	o.Results = v
}

func (o InboundMessageResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundMessageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageCount) {
		toSerialize["messageCount"] = o.MessageCount
	}
	if !IsNil(o.PendingMessageCount) {
		toSerialize["pendingMessageCount"] = o.PendingMessageCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableInboundMessageResult struct {
	value *InboundMessageResult
	isSet bool
}

func (v NullableInboundMessageResult) Get() *InboundMessageResult {
	return v.value
}

func (v *NullableInboundMessageResult) Set(val *InboundMessageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundMessageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundMessageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundMessageResult(val *InboundMessageResult) *NullableInboundMessageResult {
	return &NullableInboundMessageResult{value: val, isSet: true}
}

func (v NullableInboundMessageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundMessageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
