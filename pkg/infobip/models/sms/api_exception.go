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

// checks if the ApiException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiException{}

// ApiException struct for ApiException
type ApiException struct {
	RequestError *ApiRequestError
}

// NewApiException instantiates a new ApiException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiException() *ApiException {
	this := ApiException{}
	return &this
}

// NewApiExceptionWithDefaults instantiates a new ApiException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiExceptionWithDefaults() *ApiException {
	this := ApiException{}
	return &this
}

// GetRequestError returns the RequestError field value if set, zero value otherwise.
func (o *ApiException) GetRequestError() ApiRequestError {
	if o == nil || IsNil(o.RequestError) {
		var ret ApiRequestError
		return ret
	}
	return *o.RequestError
}

// GetRequestErrorOk returns a tuple with the RequestError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiException) GetRequestErrorOk() (*ApiRequestError, bool) {
	if o == nil || IsNil(o.RequestError) {
		return nil, false
	}
	return o.RequestError, true
}

// HasRequestError returns a boolean if a field has been set.
func (o *ApiException) HasRequestError() bool {
	if o != nil && !IsNil(o.RequestError) {
		return true
	}

	return false
}

// SetRequestError gets a reference to the given ApiRequestError and assigns it to the RequestError field.
func (o *ApiException) SetRequestError(v ApiRequestError) {
	o.RequestError = &v
}

func (o ApiException) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestError) {
		toSerialize["requestError"] = o.RequestError
	}
	return toSerialize, nil
}

type NullableApiException struct {
	value *ApiException
	isSet bool
}

func (v NullableApiException) Get() *ApiException {
	return v.value
}

func (v *NullableApiException) Set(val *ApiException) {
	v.value = val
	v.isSet = true
}

func (v NullableApiException) IsSet() bool {
	return v.isSet
}

func (v *NullableApiException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiException(val *ApiException) *NullableApiException {
	return &NullableApiException{value: val, isSet: true}
}

func (v NullableApiException) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
