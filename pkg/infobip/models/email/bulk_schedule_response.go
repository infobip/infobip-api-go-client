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

// checks if the BulkScheduleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkScheduleResponse{}

// BulkScheduleResponse struct for BulkScheduleResponse
type BulkScheduleResponse struct {
	ExternalBulkId *string
	Bulks          []BulkInfo
}

// NewBulkScheduleResponse instantiates a new BulkScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewBulkScheduleResponse() *BulkScheduleResponse {
	this := BulkScheduleResponse{}
	return &this
}

// NewBulkScheduleResponseWithDefaults instantiates a new BulkScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkScheduleResponseWithDefaults() *BulkScheduleResponse {
	this := BulkScheduleResponse{}

	return &this
}

// GetExternalBulkId returns the ExternalBulkId field value if set, zero value otherwise.
func (o *BulkScheduleResponse) GetExternalBulkId() string {
	if o == nil || IsNil(o.ExternalBulkId) {
		var ret string
		return ret
	}
	return *o.ExternalBulkId
}

// GetExternalBulkIdOk returns a tuple with the ExternalBulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkScheduleResponse) GetExternalBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalBulkId) {
		return nil, false
	}
	return o.ExternalBulkId, true
}

// HasExternalBulkId returns a boolean if a field has been set.
func (o *BulkScheduleResponse) HasExternalBulkId() bool {
	if o != nil && !IsNil(o.ExternalBulkId) {
		return true
	}

	return false
}

// SetExternalBulkId gets a reference to the given string and assigns it to the ExternalBulkId field.
func (o *BulkScheduleResponse) SetExternalBulkId(v string) {
	o.ExternalBulkId = &v
}

// GetBulks returns the Bulks field value if set, zero value otherwise.
func (o *BulkScheduleResponse) GetBulks() []BulkInfo {
	if o == nil || IsNil(o.Bulks) {
		var ret []BulkInfo
		return ret
	}
	return o.Bulks
}

// GetBulksOk returns a tuple with the Bulks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkScheduleResponse) GetBulksOk() ([]BulkInfo, bool) {
	if o == nil || IsNil(o.Bulks) {
		return nil, false
	}
	return o.Bulks, true
}

// HasBulks returns a boolean if a field has been set.
func (o *BulkScheduleResponse) HasBulks() bool {
	if o != nil && !IsNil(o.Bulks) {
		return true
	}

	return false
}

// SetBulks gets a reference to the given []BulkInfo and assigns it to the Bulks field.
func (o *BulkScheduleResponse) SetBulks(v []BulkInfo) {
	o.Bulks = v
}

func (o BulkScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalBulkId) {
		toSerialize["externalBulkId"] = o.ExternalBulkId
	}
	if !IsNil(o.Bulks) {
		toSerialize["bulks"] = o.Bulks
	}
	return toSerialize, nil
}

type NullableBulkScheduleResponse struct {
	value *BulkScheduleResponse
	isSet bool
}

func (v NullableBulkScheduleResponse) Get() *BulkScheduleResponse {
	return v.value
}

func (v *NullableBulkScheduleResponse) Set(val *BulkScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkScheduleResponse(val *BulkScheduleResponse) *NullableBulkScheduleResponse {
	return &NullableBulkScheduleResponse{value: val, isSet: true}
}

func (v NullableBulkScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
