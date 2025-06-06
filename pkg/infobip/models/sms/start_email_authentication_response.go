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

// checks if the StartEmailAuthenticationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartEmailAuthenticationResponse{}

// StartEmailAuthenticationResponse struct for StartEmailAuthenticationResponse
type StartEmailAuthenticationResponse struct {
	EmailStatus *EmailStatus
	// Sent PIN code ID.
	PinId *string
	// Email address to which the 2FA message has been sent. Example: john.smith@example.com.
	To *string
}

// NewStartEmailAuthenticationResponse instantiates a new StartEmailAuthenticationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewStartEmailAuthenticationResponse() *StartEmailAuthenticationResponse {
	this := StartEmailAuthenticationResponse{}
	return &this
}

// NewStartEmailAuthenticationResponseWithDefaults instantiates a new StartEmailAuthenticationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartEmailAuthenticationResponseWithDefaults() *StartEmailAuthenticationResponse {
	this := StartEmailAuthenticationResponse{}

	return &this
}

// GetEmailStatus returns the EmailStatus field value if set, zero value otherwise.
func (o *StartEmailAuthenticationResponse) GetEmailStatus() EmailStatus {
	if o == nil || IsNil(o.EmailStatus) {
		var ret EmailStatus
		return ret
	}
	return *o.EmailStatus
}

// GetEmailStatusOk returns a tuple with the EmailStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartEmailAuthenticationResponse) GetEmailStatusOk() (*EmailStatus, bool) {
	if o == nil || IsNil(o.EmailStatus) {
		return nil, false
	}
	return o.EmailStatus, true
}

// HasEmailStatus returns a boolean if a field has been set.
func (o *StartEmailAuthenticationResponse) HasEmailStatus() bool {
	if o != nil && !IsNil(o.EmailStatus) {
		return true
	}

	return false
}

// SetEmailStatus gets a reference to the given EmailStatus and assigns it to the EmailStatus field.
func (o *StartEmailAuthenticationResponse) SetEmailStatus(v EmailStatus) {
	o.EmailStatus = &v
}

// GetPinId returns the PinId field value if set, zero value otherwise.
func (o *StartEmailAuthenticationResponse) GetPinId() string {
	if o == nil || IsNil(o.PinId) {
		var ret string
		return ret
	}
	return *o.PinId
}

// GetPinIdOk returns a tuple with the PinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartEmailAuthenticationResponse) GetPinIdOk() (*string, bool) {
	if o == nil || IsNil(o.PinId) {
		return nil, false
	}
	return o.PinId, true
}

// HasPinId returns a boolean if a field has been set.
func (o *StartEmailAuthenticationResponse) HasPinId() bool {
	if o != nil && !IsNil(o.PinId) {
		return true
	}

	return false
}

// SetPinId gets a reference to the given string and assigns it to the PinId field.
func (o *StartEmailAuthenticationResponse) SetPinId(v string) {
	o.PinId = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *StartEmailAuthenticationResponse) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartEmailAuthenticationResponse) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *StartEmailAuthenticationResponse) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *StartEmailAuthenticationResponse) SetTo(v string) {
	o.To = &v
}

func (o StartEmailAuthenticationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartEmailAuthenticationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailStatus) {
		toSerialize["emailStatus"] = o.EmailStatus
	}
	if !IsNil(o.PinId) {
		toSerialize["pinId"] = o.PinId
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableStartEmailAuthenticationResponse struct {
	value *StartEmailAuthenticationResponse
	isSet bool
}

func (v NullableStartEmailAuthenticationResponse) Get() *StartEmailAuthenticationResponse {
	return v.value
}

func (v *NullableStartEmailAuthenticationResponse) Set(val *StartEmailAuthenticationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartEmailAuthenticationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartEmailAuthenticationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartEmailAuthenticationResponse(val *StartEmailAuthenticationResponse) *NullableStartEmailAuthenticationResponse {
	return &NullableStartEmailAuthenticationResponse{value: val, isSet: true}
}

func (v NullableStartEmailAuthenticationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartEmailAuthenticationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
