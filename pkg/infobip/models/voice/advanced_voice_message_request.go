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

// checks if the AdvancedVoiceMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedVoiceMessageRequest{}

// AdvancedVoiceMessageRequest struct for AdvancedVoiceMessageRequest
type AdvancedVoiceMessageRequest struct {
	// The ID which uniquely identifies the request.
	BulkId *string
	// Array of messages to be sent, one object per every message
	Messages []AdvancedVoiceMessage
}

type _AdvancedVoiceMessageRequest AdvancedVoiceMessageRequest

// NewAdvancedVoiceMessageRequest instantiates a new AdvancedVoiceMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewAdvancedVoiceMessageRequest(messages []AdvancedVoiceMessage) *AdvancedVoiceMessageRequest {
	this := AdvancedVoiceMessageRequest{}
	this.Messages = messages
	return &this
}

// NewAdvancedVoiceMessageRequestWithDefaults instantiates a new AdvancedVoiceMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedVoiceMessageRequestWithDefaults() *AdvancedVoiceMessageRequest {
	this := AdvancedVoiceMessageRequest{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *AdvancedVoiceMessageRequest) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedVoiceMessageRequest) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *AdvancedVoiceMessageRequest) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *AdvancedVoiceMessageRequest) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessages returns the Messages field value
func (o *AdvancedVoiceMessageRequest) GetMessages() []AdvancedVoiceMessage {
	if o == nil {
		var ret []AdvancedVoiceMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *AdvancedVoiceMessageRequest) GetMessagesOk() ([]AdvancedVoiceMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *AdvancedVoiceMessageRequest) SetMessages(v []AdvancedVoiceMessage) {
	o.Messages = v
}

func (o AdvancedVoiceMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedVoiceMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullableAdvancedVoiceMessageRequest struct {
	value *AdvancedVoiceMessageRequest
	isSet bool
}

func (v NullableAdvancedVoiceMessageRequest) Get() *AdvancedVoiceMessageRequest {
	return v.value
}

func (v *NullableAdvancedVoiceMessageRequest) Set(val *AdvancedVoiceMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedVoiceMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedVoiceMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedVoiceMessageRequest(val *AdvancedVoiceMessageRequest) *NullableAdvancedVoiceMessageRequest {
	return &NullableAdvancedVoiceMessageRequest{value: val, isSet: true}
}

func (v NullableAdvancedVoiceMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedVoiceMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
