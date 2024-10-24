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

// checks if the MessageStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageStatus{}

// MessageStatus Indicates the [status](https://www.infobip.com/docs/essentials/response-status-and-error-codes#api-status-codes) of the message and how to recover from an error should there be any.
type MessageStatus struct {
	// Status group ID.
	GroupId   *int32
	GroupName *MessageGeneralStatus
	// Status code ID.
	Id *int32
	// [Status code](https://www.infobip.com/docs/essentials/response-status-and-error-codes#api-status-codes) name.
	Name *string
	// Human-readable description of the status.
	Description *string
	// Action to take to recover from the error.
	Action *string
}

// NewMessageStatus instantiates a new MessageStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageStatus() *MessageStatus {
	this := MessageStatus{}
	return &this
}

// NewMessageStatusWithDefaults instantiates a new MessageStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageStatusWithDefaults() *MessageStatus {
	this := MessageStatus{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *MessageStatus) GetGroupId() int32 {
	if o == nil || IsNil(o.GroupId) {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatus) GetGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MessageStatus) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *MessageStatus) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *MessageStatus) GetGroupName() MessageGeneralStatus {
	if o == nil || IsNil(o.GroupName) {
		var ret MessageGeneralStatus
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatus) GetGroupNameOk() (*MessageGeneralStatus, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *MessageStatus) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given MessageGeneralStatus and assigns it to the GroupName field.
func (o *MessageStatus) SetGroupName(v MessageGeneralStatus) {
	o.GroupName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageStatus) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatus) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageStatus) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MessageStatus) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageStatus) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatus) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageStatus) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageStatus) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessageStatus) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageStatus) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessageStatus) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *MessageStatus) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageStatus) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *MessageStatus) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *MessageStatus) SetAction(v string) {
	o.Action = &v
}

func (o MessageStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableMessageStatus struct {
	value *MessageStatus
	isSet bool
}

func (v NullableMessageStatus) Get() *MessageStatus {
	return v.value
}

func (v *NullableMessageStatus) Set(val *MessageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageStatus(val *MessageStatus) *NullableMessageStatus {
	return &NullableMessageStatus{value: val, isSet: true}
}

func (v NullableMessageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
