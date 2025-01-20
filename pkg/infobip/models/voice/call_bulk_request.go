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

// checks if the CallBulkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallBulkRequest{}

// CallBulkRequest struct for CallBulkRequest
type CallBulkRequest struct {
	// Unique ID of the bulk request. If it's not set, a unique ID will be assigned to the bulk request.
	BulkId *string
	// Calls Configuration ID.
	CallsConfigurationId string
	Platform             *Platform
	// Bulk item list.
	Items []BulkItem
}

type _CallBulkRequest CallBulkRequest

// NewCallBulkRequest instantiates a new CallBulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCallBulkRequest(callsConfigurationId string, items []BulkItem) *CallBulkRequest {
	this := CallBulkRequest{}
	this.CallsConfigurationId = callsConfigurationId
	this.Items = items
	return &this
}

// NewCallBulkRequestWithDefaults instantiates a new CallBulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallBulkRequestWithDefaults() *CallBulkRequest {
	this := CallBulkRequest{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *CallBulkRequest) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBulkRequest) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *CallBulkRequest) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *CallBulkRequest) SetBulkId(v string) {
	o.BulkId = &v
}

// GetCallsConfigurationId returns the CallsConfigurationId field value
func (o *CallBulkRequest) GetCallsConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallsConfigurationId
}

// GetCallsConfigurationIdOk returns a tuple with the CallsConfigurationId field value
// and a boolean to check if the value has been set.
func (o *CallBulkRequest) GetCallsConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallsConfigurationId, true
}

// SetCallsConfigurationId sets field value
func (o *CallBulkRequest) SetCallsConfigurationId(v string) {
	o.CallsConfigurationId = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CallBulkRequest) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBulkRequest) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CallBulkRequest) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *CallBulkRequest) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetItems returns the Items field value
func (o *CallBulkRequest) GetItems() []BulkItem {
	if o == nil {
		var ret []BulkItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CallBulkRequest) GetItemsOk() ([]BulkItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CallBulkRequest) SetItems(v []BulkItem) {
	o.Items = v
}

func (o CallBulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallBulkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	toSerialize["callsConfigurationId"] = o.CallsConfigurationId
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableCallBulkRequest struct {
	value *CallBulkRequest
	isSet bool
}

func (v NullableCallBulkRequest) Get() *CallBulkRequest {
	return v.value
}

func (v *NullableCallBulkRequest) Set(val *CallBulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCallBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCallBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallBulkRequest(val *CallBulkRequest) *NullableCallBulkRequest {
	return &NullableCallBulkRequest{value: val, isSet: true}
}

func (v NullableCallBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
