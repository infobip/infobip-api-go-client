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

// checks if the MessageError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageError{}

// MessageError Indicates whether an [error](https://www.infobip.com/docs/essentials/response-status-and-error-codes#error-codes) occurred during the query execution.
type MessageError struct {
	// Error group ID.
	GroupId   *int32
	GroupName *MessageErrorGroup
	// Error ID.
	Id *int32
	// [Error name](https://www.infobip.com/docs/essentials/response-status-and-error-codes#error-codes).
	Name *string
	// Human-readable description of the error.
	Description *string
	// Indicates whether the error is recoverable or not.
	Permanent *bool
}

// NewMessageError instantiates a new MessageError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageError() *MessageError {
	this := MessageError{}
	return &this
}

// NewMessageErrorWithDefaults instantiates a new MessageError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageErrorWithDefaults() *MessageError {
	this := MessageError{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *MessageError) GetGroupId() int32 {
	if o == nil || IsNil(o.GroupId) {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageError) GetGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MessageError) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *MessageError) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *MessageError) GetGroupName() MessageErrorGroup {
	if o == nil || IsNil(o.GroupName) {
		var ret MessageErrorGroup
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageError) GetGroupNameOk() (*MessageErrorGroup, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *MessageError) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given MessageErrorGroup and assigns it to the GroupName field.
func (o *MessageError) SetGroupName(v MessageErrorGroup) {
	o.GroupName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageError) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageError) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MessageError) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageError) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageError) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageError) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageError) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessageError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessageError) SetDescription(v string) {
	o.Description = &v
}

// GetPermanent returns the Permanent field value if set, zero value otherwise.
func (o *MessageError) GetPermanent() bool {
	if o == nil || IsNil(o.Permanent) {
		var ret bool
		return ret
	}
	return *o.Permanent
}

// GetPermanentOk returns a tuple with the Permanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageError) GetPermanentOk() (*bool, bool) {
	if o == nil || IsNil(o.Permanent) {
		return nil, false
	}
	return o.Permanent, true
}

// HasPermanent returns a boolean if a field has been set.
func (o *MessageError) HasPermanent() bool {
	if o != nil && !IsNil(o.Permanent) {
		return true
	}

	return false
}

// SetPermanent gets a reference to the given bool and assigns it to the Permanent field.
func (o *MessageError) SetPermanent(v bool) {
	o.Permanent = &v
}

func (o MessageError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Permanent) {
		toSerialize["permanent"] = o.Permanent
	}
	return toSerialize, nil
}

type NullableMessageError struct {
	value *MessageError
	isSet bool
}

func (v NullableMessageError) Get() *MessageError {
	return v.value
}

func (v *NullableMessageError) Set(val *MessageError) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageError) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageError(val *MessageError) *NullableMessageError {
	return &NullableMessageError{value: val, isSet: true}
}

func (v NullableMessageError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
