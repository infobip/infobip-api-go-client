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

// checks if the StartAuthenticationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartAuthenticationResponse{}

// StartAuthenticationResponse struct for StartAuthenticationResponse
type StartAuthenticationResponse struct {
	// Call status, e.g. `PENDING_ACCEPTED`.
	CallStatus *string
	// Status of sent [Number Lookup](https://www.infobip.com/docs/number-lookup). Number Lookup status can have one of the following values: `NC_DESTINATION_UNKNOWN`, `NC_DESTINATION_REACHABLE`, `NC_DESTINATION_NOT_REACHABLE`, `NC_NOT_CONFIGURED`. Contact your Account Manager, if you get the `NC_NOT_CONFIGURED` status. SMS will not be sent only if Number Lookup status is `NC_NOT_REACHABLE`.
	NcStatus *string
	// Sent PIN code ID.
	PinId *string
	// Sent SMS status. Can have one of the following values: `MESSAGE_SENT`, `MESSAGE_NOT_SENT`.
	SmsStatus *string
	// Phone number to which the 2FA message will be sent. Example: `41793026727`.
	To *string
}

// NewStartAuthenticationResponse instantiates a new StartAuthenticationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewStartAuthenticationResponse() *StartAuthenticationResponse {
	this := StartAuthenticationResponse{}
	return &this
}

// NewStartAuthenticationResponseWithDefaults instantiates a new StartAuthenticationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartAuthenticationResponseWithDefaults() *StartAuthenticationResponse {
	this := StartAuthenticationResponse{}

	return &this
}

// GetCallStatus returns the CallStatus field value if set, zero value otherwise.
func (o *StartAuthenticationResponse) GetCallStatus() string {
	if o == nil || IsNil(o.CallStatus) {
		var ret string
		return ret
	}
	return *o.CallStatus
}

// GetCallStatusOk returns a tuple with the CallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAuthenticationResponse) GetCallStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CallStatus) {
		return nil, false
	}
	return o.CallStatus, true
}

// HasCallStatus returns a boolean if a field has been set.
func (o *StartAuthenticationResponse) HasCallStatus() bool {
	if o != nil && !IsNil(o.CallStatus) {
		return true
	}

	return false
}

// SetCallStatus gets a reference to the given string and assigns it to the CallStatus field.
func (o *StartAuthenticationResponse) SetCallStatus(v string) {
	o.CallStatus = &v
}

// GetNcStatus returns the NcStatus field value if set, zero value otherwise.
func (o *StartAuthenticationResponse) GetNcStatus() string {
	if o == nil || IsNil(o.NcStatus) {
		var ret string
		return ret
	}
	return *o.NcStatus
}

// GetNcStatusOk returns a tuple with the NcStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAuthenticationResponse) GetNcStatusOk() (*string, bool) {
	if o == nil || IsNil(o.NcStatus) {
		return nil, false
	}
	return o.NcStatus, true
}

// HasNcStatus returns a boolean if a field has been set.
func (o *StartAuthenticationResponse) HasNcStatus() bool {
	if o != nil && !IsNil(o.NcStatus) {
		return true
	}

	return false
}

// SetNcStatus gets a reference to the given string and assigns it to the NcStatus field.
func (o *StartAuthenticationResponse) SetNcStatus(v string) {
	o.NcStatus = &v
}

// GetPinId returns the PinId field value if set, zero value otherwise.
func (o *StartAuthenticationResponse) GetPinId() string {
	if o == nil || IsNil(o.PinId) {
		var ret string
		return ret
	}
	return *o.PinId
}

// GetPinIdOk returns a tuple with the PinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAuthenticationResponse) GetPinIdOk() (*string, bool) {
	if o == nil || IsNil(o.PinId) {
		return nil, false
	}
	return o.PinId, true
}

// HasPinId returns a boolean if a field has been set.
func (o *StartAuthenticationResponse) HasPinId() bool {
	if o != nil && !IsNil(o.PinId) {
		return true
	}

	return false
}

// SetPinId gets a reference to the given string and assigns it to the PinId field.
func (o *StartAuthenticationResponse) SetPinId(v string) {
	o.PinId = &v
}

// GetSmsStatus returns the SmsStatus field value if set, zero value otherwise.
func (o *StartAuthenticationResponse) GetSmsStatus() string {
	if o == nil || IsNil(o.SmsStatus) {
		var ret string
		return ret
	}
	return *o.SmsStatus
}

// GetSmsStatusOk returns a tuple with the SmsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAuthenticationResponse) GetSmsStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SmsStatus) {
		return nil, false
	}
	return o.SmsStatus, true
}

// HasSmsStatus returns a boolean if a field has been set.
func (o *StartAuthenticationResponse) HasSmsStatus() bool {
	if o != nil && !IsNil(o.SmsStatus) {
		return true
	}

	return false
}

// SetSmsStatus gets a reference to the given string and assigns it to the SmsStatus field.
func (o *StartAuthenticationResponse) SetSmsStatus(v string) {
	o.SmsStatus = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *StartAuthenticationResponse) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartAuthenticationResponse) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *StartAuthenticationResponse) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *StartAuthenticationResponse) SetTo(v string) {
	o.To = &v
}

func (o StartAuthenticationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartAuthenticationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallStatus) {
		toSerialize["callStatus"] = o.CallStatus
	}
	if !IsNil(o.NcStatus) {
		toSerialize["ncStatus"] = o.NcStatus
	}
	if !IsNil(o.PinId) {
		toSerialize["pinId"] = o.PinId
	}
	if !IsNil(o.SmsStatus) {
		toSerialize["smsStatus"] = o.SmsStatus
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableStartAuthenticationResponse struct {
	value *StartAuthenticationResponse
	isSet bool
}

func (v NullableStartAuthenticationResponse) Get() *StartAuthenticationResponse {
	return v.value
}

func (v *NullableStartAuthenticationResponse) Set(val *StartAuthenticationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartAuthenticationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartAuthenticationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartAuthenticationResponse(val *StartAuthenticationResponse) *NullableStartAuthenticationResponse {
	return &NullableStartAuthenticationResponse{value: val, isSet: true}
}

func (v NullableStartAuthenticationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartAuthenticationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
