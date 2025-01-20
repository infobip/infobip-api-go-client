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

// checks if the ClickToCallMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClickToCallMessageRequest{}

// ClickToCallMessageRequest struct for ClickToCallMessageRequest
type ClickToCallMessageRequest struct {
	// The ID which uniquely identifies the request.
	BulkId *string
	// Array of click to call messages to be sent.
	Messages []ClickToCallMessage
}

type _ClickToCallMessageRequest ClickToCallMessageRequest

// NewClickToCallMessageRequest instantiates a new ClickToCallMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewClickToCallMessageRequest(messages []ClickToCallMessage) *ClickToCallMessageRequest {
	this := ClickToCallMessageRequest{}
	this.Messages = messages
	return &this
}

// NewClickToCallMessageRequestWithDefaults instantiates a new ClickToCallMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickToCallMessageRequestWithDefaults() *ClickToCallMessageRequest {
	this := ClickToCallMessageRequest{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *ClickToCallMessageRequest) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickToCallMessageRequest) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *ClickToCallMessageRequest) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *ClickToCallMessageRequest) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessages returns the Messages field value
func (o *ClickToCallMessageRequest) GetMessages() []ClickToCallMessage {
	if o == nil {
		var ret []ClickToCallMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ClickToCallMessageRequest) GetMessagesOk() ([]ClickToCallMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ClickToCallMessageRequest) SetMessages(v []ClickToCallMessage) {
	o.Messages = v
}

func (o ClickToCallMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClickToCallMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullableClickToCallMessageRequest struct {
	value *ClickToCallMessageRequest
	isSet bool
}

func (v NullableClickToCallMessageRequest) Get() *ClickToCallMessageRequest {
	return v.value
}

func (v *NullableClickToCallMessageRequest) Set(val *ClickToCallMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClickToCallMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClickToCallMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickToCallMessageRequest(val *ClickToCallMessageRequest) *NullableClickToCallMessageRequest {
	return &NullableClickToCallMessageRequest{value: val, isSet: true}
}

func (v NullableClickToCallMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickToCallMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
