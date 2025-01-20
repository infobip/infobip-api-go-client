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

// checks if the MoEventImageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoEventImageContent{}

// MoEventImageContent struct for MoEventImageContent
type MoEventImageContent struct {
	Type MoEventContentType
	// URL of the image.
	Url string
	// Image caption.
	Text *string
}

type _MoEventImageContent MoEventImageContent

// NewMoEventImageContent instantiates a new MoEventImageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoEventImageContent(url string) *MoEventImageContent {
	this := MoEventImageContent{}
	this.Type = "IMAGE"
	this.Url = url
	return &this
}

// NewMoEventImageContentWithDefaults instantiates a new MoEventImageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoEventImageContentWithDefaults() *MoEventImageContent {
	this := MoEventImageContent{}
	this.Type = "IMAGE"
	return &this
}

// GetUrl returns the Url field value
func (o *MoEventImageContent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MoEventImageContent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MoEventImageContent) SetUrl(v string) {
	o.Url = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MoEventImageContent) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoEventImageContent) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MoEventImageContent) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *MoEventImageContent) SetText(v string) {
	o.Text = &v
}

func (o MoEventImageContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoEventImageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableMoEventImageContent struct {
	value *MoEventImageContent
	isSet bool
}

func (v NullableMoEventImageContent) Get() *MoEventImageContent {
	return v.value
}

func (v *NullableMoEventImageContent) Set(val *MoEventImageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMoEventImageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMoEventImageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoEventImageContent(val *MoEventImageContent) *NullableMoEventImageContent {
	return &NullableMoEventImageContent{value: val, isSet: true}
}

func (v NullableMoEventImageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoEventImageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
