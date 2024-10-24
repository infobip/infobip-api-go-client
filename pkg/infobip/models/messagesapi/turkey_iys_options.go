/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package messagesapi

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the TurkeyIysOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TurkeyIysOptions{}

// TurkeyIysOptions IYS regulations specific parameters required for sending promotional SMS to phone numbers registered in Turkey.
type TurkeyIysOptions struct {
	// Brand code is an ID of the company based on a company VAT number. If not provided in request, default value is used from your Infobip account.
	BrandCode     *int32
	RecipientType RecipientType
}

type _TurkeyIysOptions TurkeyIysOptions

// NewTurkeyIysOptions instantiates a new TurkeyIysOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTurkeyIysOptions(recipientType RecipientType) *TurkeyIysOptions {
	this := TurkeyIysOptions{}
	this.RecipientType = recipientType
	return &this
}

// NewTurkeyIysOptionsWithDefaults instantiates a new TurkeyIysOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTurkeyIysOptionsWithDefaults() *TurkeyIysOptions {
	this := TurkeyIysOptions{}
	return &this
}

// GetBrandCode returns the BrandCode field value if set, zero value otherwise.
func (o *TurkeyIysOptions) GetBrandCode() int32 {
	if o == nil || IsNil(o.BrandCode) {
		var ret int32
		return ret
	}
	return *o.BrandCode
}

// GetBrandCodeOk returns a tuple with the BrandCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TurkeyIysOptions) GetBrandCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.BrandCode) {
		return nil, false
	}
	return o.BrandCode, true
}

// HasBrandCode returns a boolean if a field has been set.
func (o *TurkeyIysOptions) HasBrandCode() bool {
	if o != nil && !IsNil(o.BrandCode) {
		return true
	}

	return false
}

// SetBrandCode gets a reference to the given int32 and assigns it to the BrandCode field.
func (o *TurkeyIysOptions) SetBrandCode(v int32) {
	o.BrandCode = &v
}

// GetRecipientType returns the RecipientType field value
func (o *TurkeyIysOptions) GetRecipientType() RecipientType {
	if o == nil {
		var ret RecipientType
		return ret
	}

	return o.RecipientType
}

// GetRecipientTypeOk returns a tuple with the RecipientType field value
// and a boolean to check if the value has been set.
func (o *TurkeyIysOptions) GetRecipientTypeOk() (*RecipientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientType, true
}

// SetRecipientType sets field value
func (o *TurkeyIysOptions) SetRecipientType(v RecipientType) {
	o.RecipientType = v
}

func (o TurkeyIysOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TurkeyIysOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandCode) {
		toSerialize["brandCode"] = o.BrandCode
	}
	toSerialize["recipientType"] = o.RecipientType
	return toSerialize, nil
}

type NullableTurkeyIysOptions struct {
	value *TurkeyIysOptions
	isSet bool
}

func (v NullableTurkeyIysOptions) Get() *TurkeyIysOptions {
	return v.value
}

func (v *NullableTurkeyIysOptions) Set(val *TurkeyIysOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTurkeyIysOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTurkeyIysOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTurkeyIysOptions(val *TurkeyIysOptions) *NullableTurkeyIysOptions {
	return &NullableTurkeyIysOptions{value: val, isSet: true}
}

func (v NullableTurkeyIysOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTurkeyIysOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
