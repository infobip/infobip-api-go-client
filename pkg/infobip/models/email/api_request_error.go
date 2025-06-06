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

// checks if the ApiRequestError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRequestError{}

// ApiRequestError struct for ApiRequestError
type ApiRequestError struct {
	ServiceException *ApiRequestErrorDetails
}

// NewApiRequestError instantiates a new ApiRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewApiRequestError() *ApiRequestError {
	this := ApiRequestError{}
	return &this
}

// NewApiRequestErrorWithDefaults instantiates a new ApiRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRequestErrorWithDefaults() *ApiRequestError {
	this := ApiRequestError{}

	return &this
}

// GetServiceException returns the ServiceException field value if set, zero value otherwise.
func (o *ApiRequestError) GetServiceException() ApiRequestErrorDetails {
	if o == nil || IsNil(o.ServiceException) {
		var ret ApiRequestErrorDetails
		return ret
	}
	return *o.ServiceException
}

// GetServiceExceptionOk returns a tuple with the ServiceException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequestError) GetServiceExceptionOk() (*ApiRequestErrorDetails, bool) {
	if o == nil || IsNil(o.ServiceException) {
		return nil, false
	}
	return o.ServiceException, true
}

// HasServiceException returns a boolean if a field has been set.
func (o *ApiRequestError) HasServiceException() bool {
	if o != nil && !IsNil(o.ServiceException) {
		return true
	}

	return false
}

// SetServiceException gets a reference to the given ApiRequestErrorDetails and assigns it to the ServiceException field.
func (o *ApiRequestError) SetServiceException(v ApiRequestErrorDetails) {
	o.ServiceException = &v
}

func (o ApiRequestError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRequestError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceException) {
		toSerialize["serviceException"] = o.ServiceException
	}
	return toSerialize, nil
}

type NullableApiRequestError struct {
	value *ApiRequestError
	isSet bool
}

func (v NullableApiRequestError) Get() *ApiRequestError {
	return v.value
}

func (v *NullableApiRequestError) Set(val *ApiRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRequestError(val *ApiRequestError) *NullableApiRequestError {
	return &NullableApiRequestError{value: val, isSet: true}
}

func (v NullableApiRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
