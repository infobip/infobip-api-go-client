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

// checks if the BulkStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkStatus{}

// BulkStatus struct for BulkStatus
type BulkStatus struct {
	// Unique ID of the bulk request.
	BulkId *string
	// Bulk request handling start time.
	StartTime *Time
	Status    *Status
}

// NewBulkStatus instantiates a new BulkStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewBulkStatus() *BulkStatus {
	this := BulkStatus{}
	return &this
}

// NewBulkStatusWithDefaults instantiates a new BulkStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkStatusWithDefaults() *BulkStatus {
	this := BulkStatus{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *BulkStatus) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkStatus) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *BulkStatus) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *BulkStatus) SetBulkId(v string) {
	o.BulkId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *BulkStatus) GetStartTime() Time {
	if o == nil || IsNil(o.StartTime) {
		var ret Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkStatus) GetStartTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *BulkStatus) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given Time and assigns it to the StartTime field.
func (o *BulkStatus) SetStartTime(v Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkStatus) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkStatus) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *BulkStatus) SetStatus(v Status) {
	o.Status = &v
}

func (o BulkStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBulkStatus struct {
	value *BulkStatus
	isSet bool
}

func (v NullableBulkStatus) Get() *BulkStatus {
	return v.value
}

func (v *NullableBulkStatus) Set(val *BulkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkStatus(val *BulkStatus) *NullableBulkStatus {
	return &NullableBulkStatus{value: val, isSet: true}
}

func (v NullableBulkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
