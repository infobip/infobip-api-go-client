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

// checks if the ApplicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResponse{}

// ApplicationResponse struct for ApplicationResponse
type ApplicationResponse struct {
	// The ID of the application that represents your service, e.g. 2FA for login, 2FA for changing the password, etc.
	ApplicationId *string
	Configuration *ApplicationConfiguration
	// Indicates whether the created application is enabled.
	Enabled *bool
	// 2FA application name.
	Name *string
}

// NewApplicationResponse instantiates a new ApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewApplicationResponse() *ApplicationResponse {
	this := ApplicationResponse{}
	return &this
}

// NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResponseWithDefaults() *ApplicationResponse {
	this := ApplicationResponse{}

	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApplicationResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApplicationResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApplicationResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ApplicationResponse) GetConfiguration() ApplicationConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret ApplicationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetConfigurationOk() (*ApplicationConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ApplicationResponse) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given ApplicationConfiguration and assigns it to the Configuration field.
func (o *ApplicationResponse) SetConfiguration(v ApplicationConfiguration) {
	o.Configuration = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplicationResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplicationResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplicationResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationResponse) SetName(v string) {
	o.Name = &v
}

func (o ApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApplicationResponse struct {
	value *ApplicationResponse
	isSet bool
}

func (v NullableApplicationResponse) Get() *ApplicationResponse {
	return v.value
}

func (v *NullableApplicationResponse) Set(val *ApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResponse(val *ApplicationResponse) *NullableApplicationResponse {
	return &NullableApplicationResponse{value: val, isSet: true}
}

func (v NullableApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
