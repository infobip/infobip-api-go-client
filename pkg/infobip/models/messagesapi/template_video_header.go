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

// checks if the TemplateVideoHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateVideoHeader{}

// TemplateVideoHeader struct for TemplateVideoHeader
type TemplateVideoHeader struct {
	Type TemplateHeaderType
	// URL of a video.
	Url string
}

type _TemplateVideoHeader TemplateVideoHeader

// NewTemplateVideoHeader instantiates a new TemplateVideoHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateVideoHeader(url string) *TemplateVideoHeader {
	this := TemplateVideoHeader{}
	this.Type = "VIDEO"
	this.Url = url
	return &this
}

// NewTemplateVideoHeaderWithDefaults instantiates a new TemplateVideoHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateVideoHeaderWithDefaults() *TemplateVideoHeader {
	this := TemplateVideoHeader{}
	this.Type = "VIDEO"
	return &this
}

// GetUrl returns the Url field value
func (o *TemplateVideoHeader) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TemplateVideoHeader) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TemplateVideoHeader) SetUrl(v string) {
	o.Url = v
}

func (o TemplateVideoHeader) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateVideoHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableTemplateVideoHeader struct {
	value *TemplateVideoHeader
	isSet bool
}

func (v NullableTemplateVideoHeader) Get() *TemplateVideoHeader {
	return v.value
}

func (v *NullableTemplateVideoHeader) Set(val *TemplateVideoHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateVideoHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateVideoHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateVideoHeader(val *TemplateVideoHeader) *NullableTemplateVideoHeader {
	return &NullableTemplateVideoHeader{value: val, isSet: true}
}

func (v NullableTemplateVideoHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateVideoHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
