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

// checks if the VoicePreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoicePreferences{}

// VoicePreferences Voice preferences.
type VoicePreferences struct {
	VoiceGender *Gender
	VoiceName   *VoiceName
}

// NewVoicePreferences instantiates a new VoicePreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewVoicePreferences() *VoicePreferences {
	this := VoicePreferences{}
	return &this
}

// NewVoicePreferencesWithDefaults instantiates a new VoicePreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoicePreferencesWithDefaults() *VoicePreferences {
	this := VoicePreferences{}

	return &this
}

// GetVoiceGender returns the VoiceGender field value if set, zero value otherwise.
func (o *VoicePreferences) GetVoiceGender() Gender {
	if o == nil || IsNil(o.VoiceGender) {
		var ret Gender
		return ret
	}
	return *o.VoiceGender
}

// GetVoiceGenderOk returns a tuple with the VoiceGender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicePreferences) GetVoiceGenderOk() (*Gender, bool) {
	if o == nil || IsNil(o.VoiceGender) {
		return nil, false
	}
	return o.VoiceGender, true
}

// HasVoiceGender returns a boolean if a field has been set.
func (o *VoicePreferences) HasVoiceGender() bool {
	if o != nil && !IsNil(o.VoiceGender) {
		return true
	}

	return false
}

// SetVoiceGender gets a reference to the given Gender and assigns it to the VoiceGender field.
func (o *VoicePreferences) SetVoiceGender(v Gender) {
	o.VoiceGender = &v
}

// GetVoiceName returns the VoiceName field value if set, zero value otherwise.
func (o *VoicePreferences) GetVoiceName() VoiceName {
	if o == nil || IsNil(o.VoiceName) {
		var ret VoiceName
		return ret
	}
	return *o.VoiceName
}

// GetVoiceNameOk returns a tuple with the VoiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicePreferences) GetVoiceNameOk() (*VoiceName, bool) {
	if o == nil || IsNil(o.VoiceName) {
		return nil, false
	}
	return o.VoiceName, true
}

// HasVoiceName returns a boolean if a field has been set.
func (o *VoicePreferences) HasVoiceName() bool {
	if o != nil && !IsNil(o.VoiceName) {
		return true
	}

	return false
}

// SetVoiceName gets a reference to the given VoiceName and assigns it to the VoiceName field.
func (o *VoicePreferences) SetVoiceName(v VoiceName) {
	o.VoiceName = &v
}

func (o VoicePreferences) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoicePreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoiceGender) {
		toSerialize["voiceGender"] = o.VoiceGender
	}
	if !IsNil(o.VoiceName) {
		toSerialize["voiceName"] = o.VoiceName
	}
	return toSerialize, nil
}

type NullableVoicePreferences struct {
	value *VoicePreferences
	isSet bool
}

func (v NullableVoicePreferences) Get() *VoicePreferences {
	return v.value
}

func (v *NullableVoicePreferences) Set(val *VoicePreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableVoicePreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableVoicePreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoicePreferences(val *VoicePreferences) *NullableVoicePreferences {
	return &NullableVoicePreferences{value: val, isSet: true}
}

func (v NullableVoicePreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoicePreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
