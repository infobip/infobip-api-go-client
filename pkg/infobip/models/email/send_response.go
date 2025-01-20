/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the SendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendResponse{}

// SendResponse struct for SendResponse
type SendResponse struct {
	// The ID that uniquely identifies a list of message responses.
	BulkId *string
	// List of message response details.
	Messages []ResponseDetails
}

// NewSendResponse instantiates a new SendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSendResponse() *SendResponse {
	this := SendResponse{}
	return &this
}

// NewSendResponseWithDefaults instantiates a new SendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendResponseWithDefaults() *SendResponse {
	this := SendResponse{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *SendResponse) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResponse) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *SendResponse) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *SendResponse) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *SendResponse) GetMessages() []ResponseDetails {
	if o == nil || IsNil(o.Messages) {
		var ret []ResponseDetails
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendResponse) GetMessagesOk() ([]ResponseDetails, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *SendResponse) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ResponseDetails and assigns it to the Messages field.
func (o *SendResponse) SetMessages(v []ResponseDetails) {
	o.Messages = v
}

func (o SendResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableSendResponse struct {
	value *SendResponse
	isSet bool
}

func (v NullableSendResponse) Get() *SendResponse {
	return v.value
}

func (v *NullableSendResponse) Set(val *SendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendResponse(val *SendResponse) *NullableSendResponse {
	return &NullableSendResponse{value: val, isSet: true}
}

func (v NullableSendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
