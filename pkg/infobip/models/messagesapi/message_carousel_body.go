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

// checks if the MessageCarouselBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCarouselBody{}

// MessageCarouselBody struct for MessageCarouselBody
type MessageCarouselBody struct {
	Type MessageBodyType
	// Carousel cards.
	Cards []MessageCarouselCard
}

type _MessageCarouselBody MessageCarouselBody

// NewMessageCarouselBody instantiates a new MessageCarouselBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCarouselBody(cards []MessageCarouselCard) *MessageCarouselBody {
	this := MessageCarouselBody{}
	this.Type = MESSAGEBODYTYPE_CAROUSEL
	this.Cards = cards
	return &this
}

// NewMessageCarouselBodyWithDefaults instantiates a new MessageCarouselBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCarouselBodyWithDefaults() *MessageCarouselBody {
	this := MessageCarouselBody{}
	return &this
}

// GetCards returns the Cards field value
func (o *MessageCarouselBody) GetCards() []MessageCarouselCard {
	if o == nil {
		var ret []MessageCarouselCard
		return ret
	}

	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value
// and a boolean to check if the value has been set.
func (o *MessageCarouselBody) GetCardsOk() ([]MessageCarouselCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cards, true
}

// SetCards sets field value
func (o *MessageCarouselBody) SetCards(v []MessageCarouselCard) {
	o.Cards = v
}

func (o MessageCarouselBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCarouselBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["cards"] = o.Cards
	return toSerialize, nil
}

type NullableMessageCarouselBody struct {
	value *MessageCarouselBody
	isSet bool
}

func (v NullableMessageCarouselBody) Get() *MessageCarouselBody {
	return v.value
}

func (v *NullableMessageCarouselBody) Set(val *MessageCarouselBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCarouselBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCarouselBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCarouselBody(val *MessageCarouselBody) *NullableMessageCarouselBody {
	return &NullableMessageCarouselBody{value: val, isSet: true}
}

func (v NullableMessageCarouselBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCarouselBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
