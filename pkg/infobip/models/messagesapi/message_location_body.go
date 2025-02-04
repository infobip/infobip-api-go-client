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

// checks if the MessageLocationBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageLocationBody{}

// MessageLocationBody struct for MessageLocationBody
type MessageLocationBody struct {
	Type MessageBodyType
	// Latitude of a location.
	Latitude float64
	// Longitude of a location.
	Longitude float64
	// Location name.
	Name *string
	// Location address.
	Address *string
}

type _MessageLocationBody MessageLocationBody

// NewMessageLocationBody instantiates a new MessageLocationBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageLocationBody(latitude float64, longitude float64) *MessageLocationBody {
	this := MessageLocationBody{}
	this.Type = "LOCATION"
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewMessageLocationBodyWithDefaults instantiates a new MessageLocationBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageLocationBodyWithDefaults() *MessageLocationBody {
	this := MessageLocationBody{}
	this.Type = "LOCATION"
	return &this
}

// GetLatitude returns the Latitude field value
func (o *MessageLocationBody) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *MessageLocationBody) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *MessageLocationBody) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *MessageLocationBody) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *MessageLocationBody) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *MessageLocationBody) SetLongitude(v float64) {
	o.Longitude = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageLocationBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageLocationBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageLocationBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageLocationBody) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MessageLocationBody) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageLocationBody) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MessageLocationBody) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MessageLocationBody) SetAddress(v string) {
	o.Address = &v
}

func (o MessageLocationBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageLocationBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableMessageLocationBody struct {
	value *MessageLocationBody
	isSet bool
}

func (v NullableMessageLocationBody) Get() *MessageLocationBody {
	return v.value
}

func (v *NullableMessageLocationBody) Set(val *MessageLocationBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageLocationBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageLocationBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageLocationBody(val *MessageLocationBody) *NullableMessageLocationBody {
	return &NullableMessageLocationBody{value: val, isSet: true}
}

func (v NullableMessageLocationBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageLocationBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
