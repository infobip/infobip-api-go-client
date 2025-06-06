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

// checks if the LaunchScenarioRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LaunchScenarioRequest{}

// LaunchScenarioRequest struct for LaunchScenarioRequest
type LaunchScenarioRequest struct {
	// The ID which uniquely identifies the request.
	BulkId *string
	// Array of IVR messages to be sent, one object per every message.
	Messages []IvrMessage
}

type _LaunchScenarioRequest LaunchScenarioRequest

// NewLaunchScenarioRequest instantiates a new LaunchScenarioRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewLaunchScenarioRequest(messages []IvrMessage) *LaunchScenarioRequest {
	this := LaunchScenarioRequest{}
	this.Messages = messages
	return &this
}

// NewLaunchScenarioRequestWithDefaults instantiates a new LaunchScenarioRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchScenarioRequestWithDefaults() *LaunchScenarioRequest {
	this := LaunchScenarioRequest{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *LaunchScenarioRequest) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchScenarioRequest) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *LaunchScenarioRequest) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *LaunchScenarioRequest) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessages returns the Messages field value
func (o *LaunchScenarioRequest) GetMessages() []IvrMessage {
	if o == nil {
		var ret []IvrMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *LaunchScenarioRequest) GetMessagesOk() ([]IvrMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *LaunchScenarioRequest) SetMessages(v []IvrMessage) {
	o.Messages = v
}

func (o LaunchScenarioRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LaunchScenarioRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullableLaunchScenarioRequest struct {
	value *LaunchScenarioRequest
	isSet bool
}

func (v NullableLaunchScenarioRequest) Get() *LaunchScenarioRequest {
	return v.value
}

func (v *NullableLaunchScenarioRequest) Set(val *LaunchScenarioRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchScenarioRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchScenarioRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchScenarioRequest(val *LaunchScenarioRequest) *NullableLaunchScenarioRequest {
	return &NullableLaunchScenarioRequest{value: val, isSet: true}
}

func (v NullableLaunchScenarioRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchScenarioRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
