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

// checks if the CarouselTemplateOpenUrlButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarouselTemplateOpenUrlButton{}

// CarouselTemplateOpenUrlButton struct for CarouselTemplateOpenUrlButton
type CarouselTemplateOpenUrlButton struct {
	Type CarouselTemplateButtonType
	// Extension of a URL defined in the registered template.
	Suffix string
}

type _CarouselTemplateOpenUrlButton CarouselTemplateOpenUrlButton

// NewCarouselTemplateOpenUrlButton instantiates a new CarouselTemplateOpenUrlButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarouselTemplateOpenUrlButton(suffix string) *CarouselTemplateOpenUrlButton {
	this := CarouselTemplateOpenUrlButton{}
	this.Type = "OPEN_URL"
	this.Suffix = suffix
	return &this
}

// NewCarouselTemplateOpenUrlButtonWithDefaults instantiates a new CarouselTemplateOpenUrlButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarouselTemplateOpenUrlButtonWithDefaults() *CarouselTemplateOpenUrlButton {
	this := CarouselTemplateOpenUrlButton{}
	this.Type = "OPEN_URL"
	return &this
}

// GetSuffix returns the Suffix field value
func (o *CarouselTemplateOpenUrlButton) GetSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value
// and a boolean to check if the value has been set.
func (o *CarouselTemplateOpenUrlButton) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suffix, true
}

// SetSuffix sets field value
func (o *CarouselTemplateOpenUrlButton) SetSuffix(v string) {
	o.Suffix = v
}

func (o CarouselTemplateOpenUrlButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CarouselTemplateOpenUrlButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["suffix"] = o.Suffix
	return toSerialize, nil
}

type NullableCarouselTemplateOpenUrlButton struct {
	value *CarouselTemplateOpenUrlButton
	isSet bool
}

func (v NullableCarouselTemplateOpenUrlButton) Get() *CarouselTemplateOpenUrlButton {
	return v.value
}

func (v *NullableCarouselTemplateOpenUrlButton) Set(val *CarouselTemplateOpenUrlButton) {
	v.value = val
	v.isSet = true
}

func (v NullableCarouselTemplateOpenUrlButton) IsSet() bool {
	return v.isSet
}

func (v *NullableCarouselTemplateOpenUrlButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarouselTemplateOpenUrlButton(val *CarouselTemplateOpenUrlButton) *NullableCarouselTemplateOpenUrlButton {
	return &NullableCarouselTemplateOpenUrlButton{value: val, isSet: true}
}

func (v NullableCarouselTemplateOpenUrlButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarouselTemplateOpenUrlButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
