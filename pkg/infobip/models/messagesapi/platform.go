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

// checks if the Platform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Platform{}

// Platform Platform options. For more details, see [documentation](https://www.infobip.com/docs/cpaas-x/application-and-entity-management).
type Platform struct {
	// Used when specifying an entity in outbound send requests. It is also returned in notification events. For detailed usage, refer to the [documentation](https://www.infobip.com/docs/cpaas-x/application-and-entity-management).
	EntityId *string
	// Used when specifying an application in outbound send requests. It is also returned in notification events. For detailed usage, refer to the [documentation](https://www.infobip.com/docs/cpaas-x/application-and-entity-management).
	ApplicationId *string
}

// NewPlatform instantiates a new Platform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatform() *Platform {
	this := Platform{}
	return &this
}

// NewPlatformWithDefaults instantiates a new Platform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformWithDefaults() *Platform {
	this := Platform{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *Platform) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *Platform) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *Platform) SetEntityId(v string) {
	o.EntityId = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *Platform) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *Platform) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *Platform) SetApplicationId(v string) {
	o.ApplicationId = &v
}

func (o Platform) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Platform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	return toSerialize, nil
}

type NullablePlatform struct {
	value *Platform
	isSet bool
}

func (v NullablePlatform) Get() *Platform {
	return v.value
}

func (v *NullablePlatform) Set(val *Platform) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatform(val *Platform) *NullablePlatform {
	return &NullablePlatform{value: val, isSet: true}
}

func (v NullablePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
