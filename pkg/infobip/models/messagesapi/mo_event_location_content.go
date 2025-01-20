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

// checks if the MoEventLocationContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoEventLocationContent{}

// MoEventLocationContent struct for MoEventLocationContent
type MoEventLocationContent struct {
	Type MoEventContentType
	// Latitude of the location.
	Latitude *float64
	// Longitude of the location.
	Longitude *float64
	// URL of the location.
	Url *string
}

type _MoEventLocationContent MoEventLocationContent

// NewMoEventLocationContent instantiates a new MoEventLocationContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoEventLocationContent() *MoEventLocationContent {
	this := MoEventLocationContent{}
	this.Type = "LOCATION"
	return &this
}

// NewMoEventLocationContentWithDefaults instantiates a new MoEventLocationContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoEventLocationContentWithDefaults() *MoEventLocationContent {
	this := MoEventLocationContent{}
	this.Type = "LOCATION"
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *MoEventLocationContent) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude) {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoEventLocationContent) GetLatitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MoEventLocationContent) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *MoEventLocationContent) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *MoEventLocationContent) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude) {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoEventLocationContent) GetLongitudeOk() (*float64, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MoEventLocationContent) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *MoEventLocationContent) SetLongitude(v float64) {
	o.Longitude = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MoEventLocationContent) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoEventLocationContent) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MoEventLocationContent) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MoEventLocationContent) SetUrl(v string) {
	o.Url = &v
}

func (o MoEventLocationContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoEventLocationContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableMoEventLocationContent struct {
	value *MoEventLocationContent
	isSet bool
}

func (v NullableMoEventLocationContent) Get() *MoEventLocationContent {
	return v.value
}

func (v *NullableMoEventLocationContent) Set(val *MoEventLocationContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMoEventLocationContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMoEventLocationContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoEventLocationContent(val *MoEventLocationContent) *NullableMoEventLocationContent {
	return &NullableMoEventLocationContent{value: val, isSet: true}
}

func (v NullableMoEventLocationContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoEventLocationContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
