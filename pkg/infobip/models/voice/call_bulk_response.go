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

// checks if the CallBulkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallBulkResponse{}

// CallBulkResponse struct for CallBulkResponse
type CallBulkResponse struct {
	// Unique ID of the bulk request.
	BulkId *string
	// Bulk call list.
	Calls []BulkCall
}

// NewCallBulkResponse instantiates a new CallBulkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCallBulkResponse() *CallBulkResponse {
	this := CallBulkResponse{}
	return &this
}

// NewCallBulkResponseWithDefaults instantiates a new CallBulkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallBulkResponseWithDefaults() *CallBulkResponse {
	this := CallBulkResponse{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *CallBulkResponse) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBulkResponse) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *CallBulkResponse) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *CallBulkResponse) SetBulkId(v string) {
	o.BulkId = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *CallBulkResponse) GetCalls() []BulkCall {
	if o == nil || IsNil(o.Calls) {
		var ret []BulkCall
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBulkResponse) GetCallsOk() ([]BulkCall, bool) {
	if o == nil || IsNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *CallBulkResponse) HasCalls() bool {
	if o != nil && !IsNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given []BulkCall and assigns it to the Calls field.
func (o *CallBulkResponse) SetCalls(v []BulkCall) {
	o.Calls = v
}

func (o CallBulkResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallBulkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	return toSerialize, nil
}

type NullableCallBulkResponse struct {
	value *CallBulkResponse
	isSet bool
}

func (v NullableCallBulkResponse) Get() *CallBulkResponse {
	return v.value
}

func (v *NullableCallBulkResponse) Set(val *CallBulkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallBulkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallBulkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallBulkResponse(val *CallBulkResponse) *NullableCallBulkResponse {
	return &NullableCallBulkResponse{value: val, isSet: true}
}

func (v NullableCallBulkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallBulkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
