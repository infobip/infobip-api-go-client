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

// checks if the TemplateFlowButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateFlowButton{}

// TemplateFlowButton struct for TemplateFlowButton
type TemplateFlowButton struct {
	Type TemplateButtonType
	// Flow token.
	Token *string
	// Message action payload data. JSON object with the data payload for the first screen
	Data map[string]map[string]interface{}
}

type _TemplateFlowButton TemplateFlowButton

// NewTemplateFlowButton instantiates a new TemplateFlowButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFlowButton() *TemplateFlowButton {
	this := TemplateFlowButton{}
	this.Type = "FLOW"
	return &this
}

// NewTemplateFlowButtonWithDefaults instantiates a new TemplateFlowButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFlowButtonWithDefaults() *TemplateFlowButton {
	this := TemplateFlowButton{}
	this.Type = "FLOW"
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TemplateFlowButton) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFlowButton) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TemplateFlowButton) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TemplateFlowButton) SetToken(v string) {
	o.Token = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TemplateFlowButton) GetData() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFlowButton) GetDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TemplateFlowButton) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *TemplateFlowButton) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

func (o TemplateFlowButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateFlowButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTemplateFlowButton struct {
	value *TemplateFlowButton
	isSet bool
}

func (v NullableTemplateFlowButton) Get() *TemplateFlowButton {
	return v.value
}

func (v *NullableTemplateFlowButton) Set(val *TemplateFlowButton) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFlowButton) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFlowButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFlowButton(val *TemplateFlowButton) *NullableTemplateFlowButton {
	return &NullableTemplateFlowButton{value: val, isSet: true}
}

func (v NullableTemplateFlowButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFlowButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
