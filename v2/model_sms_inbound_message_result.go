/*
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * Contact: support@infobip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infobip

import (
	"encoding/json"
)

// SmsInboundMessageResult struct for SmsInboundMessageResult
type SmsInboundMessageResult struct {
	// The number of messages returned in the `results` array.
	MessageCount *int32 `json:"messageCount,omitempty"`
	// The number of messages that have not been pulled in.
	PendingMessageCount *int32 `json:"pendingMessageCount,omitempty"`
	// An array of result objects.
	Results *[]SmsInboundMessage `json:"results,omitempty"`
}

// NewSmsInboundMessageResult instantiates a new SmsInboundMessageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsInboundMessageResult() *SmsInboundMessageResult {
	this := SmsInboundMessageResult{}
	return &this
}

// NewSmsInboundMessageResultWithDefaults instantiates a new SmsInboundMessageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsInboundMessageResultWithDefaults() *SmsInboundMessageResult {
	this := SmsInboundMessageResult{}
	return &this
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *SmsInboundMessageResult) GetMessageCount() int32 {
	if o == nil || o.MessageCount == nil {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsInboundMessageResult) GetMessageCountOk() (*int32, bool) {
	if o == nil || o.MessageCount == nil {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *SmsInboundMessageResult) HasMessageCount() bool {
	if o != nil && o.MessageCount != nil {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *SmsInboundMessageResult) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetPendingMessageCount returns the PendingMessageCount field value if set, zero value otherwise.
func (o *SmsInboundMessageResult) GetPendingMessageCount() int32 {
	if o == nil || o.PendingMessageCount == nil {
		var ret int32
		return ret
	}
	return *o.PendingMessageCount
}

// GetPendingMessageCountOk returns a tuple with the PendingMessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsInboundMessageResult) GetPendingMessageCountOk() (*int32, bool) {
	if o == nil || o.PendingMessageCount == nil {
		return nil, false
	}
	return o.PendingMessageCount, true
}

// HasPendingMessageCount returns a boolean if a field has been set.
func (o *SmsInboundMessageResult) HasPendingMessageCount() bool {
	if o != nil && o.PendingMessageCount != nil {
		return true
	}

	return false
}

// SetPendingMessageCount gets a reference to the given int32 and assigns it to the PendingMessageCount field.
func (o *SmsInboundMessageResult) SetPendingMessageCount(v int32) {
	o.PendingMessageCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SmsInboundMessageResult) GetResults() []SmsInboundMessage {
	if o == nil || o.Results == nil {
		var ret []SmsInboundMessage
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsInboundMessageResult) GetResultsOk() (*[]SmsInboundMessage, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SmsInboundMessageResult) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SmsInboundMessage and assigns it to the Results field.
func (o *SmsInboundMessageResult) SetResults(v []SmsInboundMessage) {
	o.Results = &v
}

func (o SmsInboundMessageResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageCount != nil {
		toSerialize["messageCount"] = o.MessageCount
	}
	if o.PendingMessageCount != nil {
		toSerialize["pendingMessageCount"] = o.PendingMessageCount
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSmsInboundMessageResult struct {
	value *SmsInboundMessageResult
	isSet bool
}

func (v NullableSmsInboundMessageResult) Get() *SmsInboundMessageResult {
	return v.value
}

func (v *NullableSmsInboundMessageResult) Set(val *SmsInboundMessageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsInboundMessageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsInboundMessageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsInboundMessageResult(val *SmsInboundMessageResult) *NullableSmsInboundMessageResult {
	return &NullableSmsInboundMessageResult{value: val, isSet: true}
}

func (v NullableSmsInboundMessageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsInboundMessageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
