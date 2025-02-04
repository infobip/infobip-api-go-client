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

// checks if the ExtendedSipTrunkStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendedSipTrunkStatusResponse{}

// ExtendedSipTrunkStatusResponse struct for ExtendedSipTrunkStatusResponse
type ExtendedSipTrunkStatusResponse struct {
	AdminStatus        *SipTrunkAdminStatus
	ActionStatus       *SipTrunkActionStatusResponse
	RegistrationStatus *SipTrunkRegistrationStatus
	// Number of active calls.
	ActiveCalls *int32
}

// NewExtendedSipTrunkStatusResponse instantiates a new ExtendedSipTrunkStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewExtendedSipTrunkStatusResponse() *ExtendedSipTrunkStatusResponse {
	this := ExtendedSipTrunkStatusResponse{}
	return &this
}

// NewExtendedSipTrunkStatusResponseWithDefaults instantiates a new ExtendedSipTrunkStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedSipTrunkStatusResponseWithDefaults() *ExtendedSipTrunkStatusResponse {
	this := ExtendedSipTrunkStatusResponse{}

	return &this
}

// GetAdminStatus returns the AdminStatus field value if set, zero value otherwise.
func (o *ExtendedSipTrunkStatusResponse) GetAdminStatus() SipTrunkAdminStatus {
	if o == nil || IsNil(o.AdminStatus) {
		var ret SipTrunkAdminStatus
		return ret
	}
	return *o.AdminStatus
}

// GetAdminStatusOk returns a tuple with the AdminStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedSipTrunkStatusResponse) GetAdminStatusOk() (*SipTrunkAdminStatus, bool) {
	if o == nil || IsNil(o.AdminStatus) {
		return nil, false
	}
	return o.AdminStatus, true
}

// HasAdminStatus returns a boolean if a field has been set.
func (o *ExtendedSipTrunkStatusResponse) HasAdminStatus() bool {
	if o != nil && !IsNil(o.AdminStatus) {
		return true
	}

	return false
}

// SetAdminStatus gets a reference to the given SipTrunkAdminStatus and assigns it to the AdminStatus field.
func (o *ExtendedSipTrunkStatusResponse) SetAdminStatus(v SipTrunkAdminStatus) {
	o.AdminStatus = &v
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *ExtendedSipTrunkStatusResponse) GetActionStatus() SipTrunkActionStatusResponse {
	if o == nil || IsNil(o.ActionStatus) {
		var ret SipTrunkActionStatusResponse
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedSipTrunkStatusResponse) GetActionStatusOk() (*SipTrunkActionStatusResponse, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *ExtendedSipTrunkStatusResponse) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given SipTrunkActionStatusResponse and assigns it to the ActionStatus field.
func (o *ExtendedSipTrunkStatusResponse) SetActionStatus(v SipTrunkActionStatusResponse) {
	o.ActionStatus = &v
}

// GetRegistrationStatus returns the RegistrationStatus field value if set, zero value otherwise.
func (o *ExtendedSipTrunkStatusResponse) GetRegistrationStatus() SipTrunkRegistrationStatus {
	if o == nil || IsNil(o.RegistrationStatus) {
		var ret SipTrunkRegistrationStatus
		return ret
	}
	return *o.RegistrationStatus
}

// GetRegistrationStatusOk returns a tuple with the RegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedSipTrunkStatusResponse) GetRegistrationStatusOk() (*SipTrunkRegistrationStatus, bool) {
	if o == nil || IsNil(o.RegistrationStatus) {
		return nil, false
	}
	return o.RegistrationStatus, true
}

// HasRegistrationStatus returns a boolean if a field has been set.
func (o *ExtendedSipTrunkStatusResponse) HasRegistrationStatus() bool {
	if o != nil && !IsNil(o.RegistrationStatus) {
		return true
	}

	return false
}

// SetRegistrationStatus gets a reference to the given SipTrunkRegistrationStatus and assigns it to the RegistrationStatus field.
func (o *ExtendedSipTrunkStatusResponse) SetRegistrationStatus(v SipTrunkRegistrationStatus) {
	o.RegistrationStatus = &v
}

// GetActiveCalls returns the ActiveCalls field value if set, zero value otherwise.
func (o *ExtendedSipTrunkStatusResponse) GetActiveCalls() int32 {
	if o == nil || IsNil(o.ActiveCalls) {
		var ret int32
		return ret
	}
	return *o.ActiveCalls
}

// GetActiveCallsOk returns a tuple with the ActiveCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedSipTrunkStatusResponse) GetActiveCallsOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveCalls) {
		return nil, false
	}
	return o.ActiveCalls, true
}

// HasActiveCalls returns a boolean if a field has been set.
func (o *ExtendedSipTrunkStatusResponse) HasActiveCalls() bool {
	if o != nil && !IsNil(o.ActiveCalls) {
		return true
	}

	return false
}

// SetActiveCalls gets a reference to the given int32 and assigns it to the ActiveCalls field.
func (o *ExtendedSipTrunkStatusResponse) SetActiveCalls(v int32) {
	o.ActiveCalls = &v
}

func (o ExtendedSipTrunkStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedSipTrunkStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminStatus) {
		toSerialize["adminStatus"] = o.AdminStatus
	}
	if !IsNil(o.ActionStatus) {
		toSerialize["actionStatus"] = o.ActionStatus
	}
	if !IsNil(o.RegistrationStatus) {
		toSerialize["registrationStatus"] = o.RegistrationStatus
	}
	if !IsNil(o.ActiveCalls) {
		toSerialize["activeCalls"] = o.ActiveCalls
	}
	return toSerialize, nil
}

type NullableExtendedSipTrunkStatusResponse struct {
	value *ExtendedSipTrunkStatusResponse
	isSet bool
}

func (v NullableExtendedSipTrunkStatusResponse) Get() *ExtendedSipTrunkStatusResponse {
	return v.value
}

func (v *NullableExtendedSipTrunkStatusResponse) Set(val *ExtendedSipTrunkStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedSipTrunkStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedSipTrunkStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedSipTrunkStatusResponse(val *ExtendedSipTrunkStatusResponse) *NullableExtendedSipTrunkStatusResponse {
	return &NullableExtendedSipTrunkStatusResponse{value: val, isSet: true}
}

func (v NullableExtendedSipTrunkStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedSipTrunkStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
