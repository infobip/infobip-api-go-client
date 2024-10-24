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

// checks if the TemplateCarouselBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateCarouselBody{}

// TemplateCarouselBody struct for TemplateCarouselBody
type TemplateCarouselBody struct {
	Type TemplateBodyType
	// Carousel cards.
	Cards []TemplateCarouselCard
}

type _TemplateCarouselBody TemplateCarouselBody

// NewTemplateCarouselBody instantiates a new TemplateCarouselBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCarouselBody(cards []TemplateCarouselCard) *TemplateCarouselBody {
	this := TemplateCarouselBody{}
	this.Type = TEMPLATEBODYTYPE_CAROUSEL
	this.Cards = cards
	return &this
}

// NewTemplateCarouselBodyWithDefaults instantiates a new TemplateCarouselBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCarouselBodyWithDefaults() *TemplateCarouselBody {
	this := TemplateCarouselBody{}
	return &this
}

// GetCards returns the Cards field value
func (o *TemplateCarouselBody) GetCards() []TemplateCarouselCard {
	if o == nil {
		var ret []TemplateCarouselCard
		return ret
	}

	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value
// and a boolean to check if the value has been set.
func (o *TemplateCarouselBody) GetCardsOk() ([]TemplateCarouselCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cards, true
}

// SetCards sets field value
func (o *TemplateCarouselBody) SetCards(v []TemplateCarouselCard) {
	o.Cards = v
}

func (o TemplateCarouselBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateCarouselBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["cards"] = o.Cards
	return toSerialize, nil
}

type NullableTemplateCarouselBody struct {
	value *TemplateCarouselBody
	isSet bool
}

func (v NullableTemplateCarouselBody) Get() *TemplateCarouselBody {
	return v.value
}

func (v *NullableTemplateCarouselBody) Set(val *TemplateCarouselBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCarouselBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCarouselBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCarouselBody(val *TemplateCarouselBody) *NullableTemplateCarouselBody {
	return &NullableTemplateCarouselBody{value: val, isSet: true}
}

func (v NullableTemplateCarouselBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCarouselBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}