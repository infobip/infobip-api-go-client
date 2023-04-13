/*
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * Contact: support@infobip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infobip

import (
	"encoding/json"
)

// SmsBinaryContent struct for SmsBinaryContent
type SmsBinaryContent struct {
	// Binary content data coding. The default value is (`0`) for GSM7. Example: (`8`) for  Unicode data.
	DataCoding *int32 `json:"dataCoding,omitempty"`
	// “Esm_class” parameter. Indicate special message attributes associated with the SMS. Default value is (`0`).
	EsmClass *int32 `json:"esmClass,omitempty"`
	// Hexadecimal string. This is the representation of your binary data. Two hex digits represent one byte. They should be separated by the space character (Example: `0f c2 4a bf 34 13 ba`).
	Hex string `json:"hex"`
}

// NewSmsBinaryContent instantiates a new SmsBinaryContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsBinaryContent(hex string) *SmsBinaryContent {
	this := SmsBinaryContent{}
	this.Hex = hex
	return &this
}

// NewSmsBinaryContentWithDefaults instantiates a new SmsBinaryContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsBinaryContentWithDefaults() *SmsBinaryContent {
	this := SmsBinaryContent{}
	return &this
}

// GetDataCoding returns the DataCoding field value if set, zero value otherwise.
func (o *SmsBinaryContent) GetDataCoding() int32 {
	if o == nil || o.DataCoding == nil {
		var ret int32
		return ret
	}
	return *o.DataCoding
}

// GetDataCodingOk returns a tuple with the DataCoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsBinaryContent) GetDataCodingOk() (*int32, bool) {
	if o == nil || o.DataCoding == nil {
		return nil, false
	}
	return o.DataCoding, true
}

// HasDataCoding returns a boolean if a field has been set.
func (o *SmsBinaryContent) HasDataCoding() bool {
	if o != nil && o.DataCoding != nil {
		return true
	}

	return false
}

// SetDataCoding gets a reference to the given int32 and assigns it to the DataCoding field.
func (o *SmsBinaryContent) SetDataCoding(v int32) {
	o.DataCoding = &v
}

// GetEsmClass returns the EsmClass field value if set, zero value otherwise.
func (o *SmsBinaryContent) GetEsmClass() int32 {
	if o == nil || o.EsmClass == nil {
		var ret int32
		return ret
	}
	return *o.EsmClass
}

// GetEsmClassOk returns a tuple with the EsmClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsBinaryContent) GetEsmClassOk() (*int32, bool) {
	if o == nil || o.EsmClass == nil {
		return nil, false
	}
	return o.EsmClass, true
}

// HasEsmClass returns a boolean if a field has been set.
func (o *SmsBinaryContent) HasEsmClass() bool {
	if o != nil && o.EsmClass != nil {
		return true
	}

	return false
}

// SetEsmClass gets a reference to the given int32 and assigns it to the EsmClass field.
func (o *SmsBinaryContent) SetEsmClass(v int32) {
	o.EsmClass = &v
}

// GetHex returns the Hex field value
func (o *SmsBinaryContent) GetHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hex
}

// GetHexOk returns a tuple with the Hex field value
// and a boolean to check if the value has been set.
func (o *SmsBinaryContent) GetHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hex, true
}

// SetHex sets field value
func (o *SmsBinaryContent) SetHex(v string) {
	o.Hex = v
}

func (o SmsBinaryContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataCoding != nil {
		toSerialize["dataCoding"] = o.DataCoding
	}
	if o.EsmClass != nil {
		toSerialize["esmClass"] = o.EsmClass
	}
	if true {
		toSerialize["hex"] = o.Hex
	}
	return json.Marshal(toSerialize)
}

type NullableSmsBinaryContent struct {
	value *SmsBinaryContent
	isSet bool
}

func (v NullableSmsBinaryContent) Get() *SmsBinaryContent {
	return v.value
}

func (v *NullableSmsBinaryContent) Set(val *SmsBinaryContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsBinaryContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsBinaryContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsBinaryContent(val *SmsBinaryContent) *NullableSmsBinaryContent {
	return &NullableSmsBinaryContent{value: val, isSet: true}
}

func (v NullableSmsBinaryContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsBinaryContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
