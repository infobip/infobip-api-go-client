/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moments

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the FlowCommonPushContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowCommonPushContact{}

// FlowCommonPushContact A list of person's web push destinations.
type FlowCommonPushContact struct {
	// Application Id on which the user is subscribed.
	ApplicationId *string
	// Push registration ID.
	RegistrationId *string
	// Unique user ID for a person.
	AdditionalData map[string]interface{}
	// System data collected from the user's profile.
	SystemData map[string]interface{}
}

// NewFlowCommonPushContact instantiates a new FlowCommonPushContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewFlowCommonPushContact() *FlowCommonPushContact {
	this := FlowCommonPushContact{}
	return &this
}

// NewFlowCommonPushContactWithDefaults instantiates a new FlowCommonPushContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowCommonPushContactWithDefaults() *FlowCommonPushContact {
	this := FlowCommonPushContact{}

	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *FlowCommonPushContact) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCommonPushContact) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *FlowCommonPushContact) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *FlowCommonPushContact) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetRegistrationId returns the RegistrationId field value if set, zero value otherwise.
func (o *FlowCommonPushContact) GetRegistrationId() string {
	if o == nil || IsNil(o.RegistrationId) {
		var ret string
		return ret
	}
	return *o.RegistrationId
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCommonPushContact) GetRegistrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationId) {
		return nil, false
	}
	return o.RegistrationId, true
}

// HasRegistrationId returns a boolean if a field has been set.
func (o *FlowCommonPushContact) HasRegistrationId() bool {
	if o != nil && !IsNil(o.RegistrationId) {
		return true
	}

	return false
}

// SetRegistrationId gets a reference to the given string and assigns it to the RegistrationId field.
func (o *FlowCommonPushContact) SetRegistrationId(v string) {
	o.RegistrationId = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *FlowCommonPushContact) GetAdditionalData() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalData) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCommonPushContact) GetAdditionalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *FlowCommonPushContact) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *FlowCommonPushContact) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = v
}

// GetSystemData returns the SystemData field value if set, zero value otherwise.
func (o *FlowCommonPushContact) GetSystemData() map[string]interface{} {
	if o == nil || IsNil(o.SystemData) {
		var ret map[string]interface{}
		return ret
	}
	return o.SystemData
}

// GetSystemDataOk returns a tuple with the SystemData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowCommonPushContact) GetSystemDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SystemData) {
		return map[string]interface{}{}, false
	}
	return o.SystemData, true
}

// HasSystemData returns a boolean if a field has been set.
func (o *FlowCommonPushContact) HasSystemData() bool {
	if o != nil && !IsNil(o.SystemData) {
		return true
	}

	return false
}

// SetSystemData gets a reference to the given map[string]interface{} and assigns it to the SystemData field.
func (o *FlowCommonPushContact) SetSystemData(v map[string]interface{}) {
	o.SystemData = v
}

func (o FlowCommonPushContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowCommonPushContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.RegistrationId) {
		toSerialize["registrationId"] = o.RegistrationId
	}
	if !IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !IsNil(o.SystemData) {
		toSerialize["systemData"] = o.SystemData
	}
	return toSerialize, nil
}

type NullableFlowCommonPushContact struct {
	value *FlowCommonPushContact
	isSet bool
}

func (v NullableFlowCommonPushContact) Get() *FlowCommonPushContact {
	return v.value
}

func (v *NullableFlowCommonPushContact) Set(val *FlowCommonPushContact) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowCommonPushContact) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowCommonPushContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowCommonPushContact(val *FlowCommonPushContact) *NullableFlowCommonPushContact {
	return &NullableFlowCommonPushContact{value: val, isSet: true}
}

func (v NullableFlowCommonPushContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowCommonPushContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
