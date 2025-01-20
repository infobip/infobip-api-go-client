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

// checks if the VerifyPinResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyPinResponse{}

// VerifyPinResponse struct for VerifyPinResponse
type VerifyPinResponse struct {
	// Number of remaining PIN attempts.
	AttemptsRemaining *int32
	// Phone number (`MSISDN`) to which the 2FA message was sent.
	Msisdn *string
	// Indicates whether an error has occurred during PIN verification.
	PinError *string
	// Sent PIN code ID.
	PinId *string
	// Indicates whether the phone number (`MSISDN`) was successfully verified.
	Verified *bool
}

// NewVerifyPinResponse instantiates a new VerifyPinResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewVerifyPinResponse() *VerifyPinResponse {
	this := VerifyPinResponse{}
	return &this
}

// NewVerifyPinResponseWithDefaults instantiates a new VerifyPinResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPinResponseWithDefaults() *VerifyPinResponse {
	this := VerifyPinResponse{}

	return &this
}

// GetAttemptsRemaining returns the AttemptsRemaining field value if set, zero value otherwise.
func (o *VerifyPinResponse) GetAttemptsRemaining() int32 {
	if o == nil || IsNil(o.AttemptsRemaining) {
		var ret int32
		return ret
	}
	return *o.AttemptsRemaining
}

// GetAttemptsRemainingOk returns a tuple with the AttemptsRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPinResponse) GetAttemptsRemainingOk() (*int32, bool) {
	if o == nil || IsNil(o.AttemptsRemaining) {
		return nil, false
	}
	return o.AttemptsRemaining, true
}

// HasAttemptsRemaining returns a boolean if a field has been set.
func (o *VerifyPinResponse) HasAttemptsRemaining() bool {
	if o != nil && !IsNil(o.AttemptsRemaining) {
		return true
	}

	return false
}

// SetAttemptsRemaining gets a reference to the given int32 and assigns it to the AttemptsRemaining field.
func (o *VerifyPinResponse) SetAttemptsRemaining(v int32) {
	o.AttemptsRemaining = &v
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *VerifyPinResponse) GetMsisdn() string {
	if o == nil || IsNil(o.Msisdn) {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPinResponse) GetMsisdnOk() (*string, bool) {
	if o == nil || IsNil(o.Msisdn) {
		return nil, false
	}
	return o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *VerifyPinResponse) HasMsisdn() bool {
	if o != nil && !IsNil(o.Msisdn) {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *VerifyPinResponse) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetPinError returns the PinError field value if set, zero value otherwise.
func (o *VerifyPinResponse) GetPinError() string {
	if o == nil || IsNil(o.PinError) {
		var ret string
		return ret
	}
	return *o.PinError
}

// GetPinErrorOk returns a tuple with the PinError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPinResponse) GetPinErrorOk() (*string, bool) {
	if o == nil || IsNil(o.PinError) {
		return nil, false
	}
	return o.PinError, true
}

// HasPinError returns a boolean if a field has been set.
func (o *VerifyPinResponse) HasPinError() bool {
	if o != nil && !IsNil(o.PinError) {
		return true
	}

	return false
}

// SetPinError gets a reference to the given string and assigns it to the PinError field.
func (o *VerifyPinResponse) SetPinError(v string) {
	o.PinError = &v
}

// GetPinId returns the PinId field value if set, zero value otherwise.
func (o *VerifyPinResponse) GetPinId() string {
	if o == nil || IsNil(o.PinId) {
		var ret string
		return ret
	}
	return *o.PinId
}

// GetPinIdOk returns a tuple with the PinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPinResponse) GetPinIdOk() (*string, bool) {
	if o == nil || IsNil(o.PinId) {
		return nil, false
	}
	return o.PinId, true
}

// HasPinId returns a boolean if a field has been set.
func (o *VerifyPinResponse) HasPinId() bool {
	if o != nil && !IsNil(o.PinId) {
		return true
	}

	return false
}

// SetPinId gets a reference to the given string and assigns it to the PinId field.
func (o *VerifyPinResponse) SetPinId(v string) {
	o.PinId = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *VerifyPinResponse) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPinResponse) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *VerifyPinResponse) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *VerifyPinResponse) SetVerified(v bool) {
	o.Verified = &v
}

func (o VerifyPinResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyPinResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttemptsRemaining) {
		toSerialize["attemptsRemaining"] = o.AttemptsRemaining
	}
	if !IsNil(o.Msisdn) {
		toSerialize["msisdn"] = o.Msisdn
	}
	if !IsNil(o.PinError) {
		toSerialize["pinError"] = o.PinError
	}
	if !IsNil(o.PinId) {
		toSerialize["pinId"] = o.PinId
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	return toSerialize, nil
}

type NullableVerifyPinResponse struct {
	value *VerifyPinResponse
	isSet bool
}

func (v NullableVerifyPinResponse) Get() *VerifyPinResponse {
	return v.value
}

func (v *NullableVerifyPinResponse) Set(val *VerifyPinResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPinResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPinResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPinResponse(val *VerifyPinResponse) *NullableVerifyPinResponse {
	return &NullableVerifyPinResponse{value: val, isSet: true}
}

func (v NullableVerifyPinResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPinResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
