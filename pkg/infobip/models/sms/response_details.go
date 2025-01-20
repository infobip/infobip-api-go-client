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

// checks if the ResponseDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDetails{}

// ResponseDetails An array of message objects of a single message or multiple messages sent under one bulk ID.
type ResponseDetails struct {
	// Unique message ID. If not provided, it will be auto-generated and returned in the API response.
	MessageId *string
	Status    *MessageStatus
	// The destination address of the message, i.e., its recipient.
	Destination *string
	Details     *MessageResponseDetails
}

// NewResponseDetails instantiates a new ResponseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewResponseDetails() *ResponseDetails {
	this := ResponseDetails{}
	return &this
}

// NewResponseDetailsWithDefaults instantiates a new ResponseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDetailsWithDefaults() *ResponseDetails {
	this := ResponseDetails{}

	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ResponseDetails) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDetails) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ResponseDetails) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ResponseDetails) SetMessageId(v string) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseDetails) GetStatus() MessageStatus {
	if o == nil || IsNil(o.Status) {
		var ret MessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDetails) GetStatusOk() (*MessageStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MessageStatus and assigns it to the Status field.
func (o *ResponseDetails) SetStatus(v MessageStatus) {
	o.Status = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ResponseDetails) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDetails) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ResponseDetails) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *ResponseDetails) SetDestination(v string) {
	o.Destination = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ResponseDetails) GetDetails() MessageResponseDetails {
	if o == nil || IsNil(o.Details) {
		var ret MessageResponseDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDetails) GetDetailsOk() (*MessageResponseDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ResponseDetails) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given MessageResponseDetails and assigns it to the Details field.
func (o *ResponseDetails) SetDetails(v MessageResponseDetails) {
	o.Details = &v
}

func (o ResponseDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableResponseDetails struct {
	value *ResponseDetails
	isSet bool
}

func (v NullableResponseDetails) Get() *ResponseDetails {
	return v.value
}

func (v *NullableResponseDetails) Set(val *ResponseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDetails(val *ResponseDetails) *NullableResponseDetails {
	return &NullableResponseDetails{value: val, isSet: true}
}

func (v NullableResponseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
