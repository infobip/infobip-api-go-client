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

// checks if the BulkStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkStatusResponse{}

// BulkStatusResponse struct for BulkStatusResponse
type BulkStatusResponse struct {
	// Unique ID assigned to the request if messaging multiple recipients or sending multiple messages via a single API request.
	BulkId *string
	Status *BulkStatus
}

// NewBulkStatusResponse instantiates a new BulkStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkStatusResponse() *BulkStatusResponse {
	this := BulkStatusResponse{}
	return &this
}

// NewBulkStatusResponseWithDefaults instantiates a new BulkStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkStatusResponseWithDefaults() *BulkStatusResponse {
	this := BulkStatusResponse{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *BulkStatusResponse) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkStatusResponse) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *BulkStatusResponse) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *BulkStatusResponse) SetBulkId(v string) {
	o.BulkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkStatusResponse) GetStatus() BulkStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkStatusResponse) GetStatusOk() (*BulkStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkStatusResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkStatus and assigns it to the Status field.
func (o *BulkStatusResponse) SetStatus(v BulkStatus) {
	o.Status = &v
}

func (o BulkStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBulkStatusResponse struct {
	value *BulkStatusResponse
	isSet bool
}

func (v NullableBulkStatusResponse) Get() *BulkStatusResponse {
	return v.value
}

func (v *NullableBulkStatusResponse) Set(val *BulkStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkStatusResponse(val *BulkStatusResponse) *NullableBulkStatusResponse {
	return &NullableBulkStatusResponse{value: val, isSet: true}
}

func (v NullableBulkStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
