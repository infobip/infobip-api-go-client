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

// checks if the RecipientInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientInfo{}

// RecipientInfo Recipient information such as device type, OS, device name.
type RecipientInfo struct {
	// The type of device used by the recipient to do the user action.
	DeviceType *string
	// The type OS present in the device used by the recipient.
	Os *string
	// Device name of the action originating device.
	DeviceName *string
}

// NewRecipientInfo instantiates a new RecipientInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewRecipientInfo() *RecipientInfo {
	this := RecipientInfo{}
	return &this
}

// NewRecipientInfoWithDefaults instantiates a new RecipientInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientInfoWithDefaults() *RecipientInfo {
	this := RecipientInfo{}

	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *RecipientInfo) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *RecipientInfo) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *RecipientInfo) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *RecipientInfo) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *RecipientInfo) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *RecipientInfo) SetOs(v string) {
	o.Os = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *RecipientInfo) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientInfo) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *RecipientInfo) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *RecipientInfo) SetDeviceName(v string) {
	o.DeviceName = &v
}

func (o RecipientInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	return toSerialize, nil
}

type NullableRecipientInfo struct {
	value *RecipientInfo
	isSet bool
}

func (v NullableRecipientInfo) Get() *RecipientInfo {
	return v.value
}

func (v *NullableRecipientInfo) Set(val *RecipientInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientInfo(val *RecipientInfo) *NullableRecipientInfo {
	return &NullableRecipientInfo{value: val, isSet: true}
}

func (v NullableRecipientInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
