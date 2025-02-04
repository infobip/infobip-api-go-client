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

// checks if the GetVoicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVoicesResponse{}

// GetVoicesResponse struct for GetVoicesResponse
type GetVoicesResponse struct {
	// Array of voices belonging to the specified language.
	Voices []SynthesisVoiceResponse
}

// NewGetVoicesResponse instantiates a new GetVoicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewGetVoicesResponse() *GetVoicesResponse {
	this := GetVoicesResponse{}
	return &this
}

// NewGetVoicesResponseWithDefaults instantiates a new GetVoicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVoicesResponseWithDefaults() *GetVoicesResponse {
	this := GetVoicesResponse{}

	return &this
}

// GetVoices returns the Voices field value if set, zero value otherwise.
func (o *GetVoicesResponse) GetVoices() []SynthesisVoiceResponse {
	if o == nil || IsNil(o.Voices) {
		var ret []SynthesisVoiceResponse
		return ret
	}
	return o.Voices
}

// GetVoicesOk returns a tuple with the Voices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVoicesResponse) GetVoicesOk() ([]SynthesisVoiceResponse, bool) {
	if o == nil || IsNil(o.Voices) {
		return nil, false
	}
	return o.Voices, true
}

// HasVoices returns a boolean if a field has been set.
func (o *GetVoicesResponse) HasVoices() bool {
	if o != nil && !IsNil(o.Voices) {
		return true
	}

	return false
}

// SetVoices gets a reference to the given []SynthesisVoiceResponse and assigns it to the Voices field.
func (o *GetVoicesResponse) SetVoices(v []SynthesisVoiceResponse) {
	o.Voices = v
}

func (o GetVoicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVoicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Voices) {
		toSerialize["voices"] = o.Voices
	}
	return toSerialize, nil
}

type NullableGetVoicesResponse struct {
	value *GetVoicesResponse
	isSet bool
}

func (v NullableGetVoicesResponse) Get() *GetVoicesResponse {
	return v.value
}

func (v *NullableGetVoicesResponse) Set(val *GetVoicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVoicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVoicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVoicesResponse(val *GetVoicesResponse) *NullableGetVoicesResponse {
	return &NullableGetVoicesResponse{value: val, isSet: true}
}

func (v NullableGetVoicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVoicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
