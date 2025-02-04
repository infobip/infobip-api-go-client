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

// checks if the MessageCarouselCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCarouselCard{}

// MessageCarouselCard Carousel cards.
type MessageCarouselCard struct {
	Body MessageCarouselCardBody
	// List of buttons of a card.
	Buttons []MessageButton
}

type _MessageCarouselCard MessageCarouselCard

// NewMessageCarouselCard instantiates a new MessageCarouselCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMessageCarouselCard(body MessageCarouselCardBody) *MessageCarouselCard {
	this := MessageCarouselCard{}
	this.Body = body
	return &this
}

// NewMessageCarouselCardWithDefaults instantiates a new MessageCarouselCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCarouselCardWithDefaults() *MessageCarouselCard {
	this := MessageCarouselCard{}

	return &this
}

// GetBody returns the Body field value
func (o *MessageCarouselCard) GetBody() MessageCarouselCardBody {
	if o == nil {
		var ret MessageCarouselCardBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *MessageCarouselCard) GetBodyOk() (*MessageCarouselCardBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *MessageCarouselCard) SetBody(v MessageCarouselCardBody) {
	o.Body = v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *MessageCarouselCard) GetButtons() []MessageButton {
	if o == nil || IsNil(o.Buttons) {
		var ret []MessageButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageCarouselCard) GetButtonsOk() ([]MessageButton, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *MessageCarouselCard) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []MessageButton and assigns it to the Buttons field.
func (o *MessageCarouselCard) SetButtons(v []MessageButton) {
	o.Buttons = v
}

func (o MessageCarouselCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCarouselCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}
	return toSerialize, nil
}

type NullableMessageCarouselCard struct {
	value *MessageCarouselCard
	isSet bool
}

func (v NullableMessageCarouselCard) Get() *MessageCarouselCard {
	return v.value
}

func (v *NullableMessageCarouselCard) Set(val *MessageCarouselCard) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCarouselCard) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCarouselCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCarouselCard(val *MessageCarouselCard) *NullableMessageCarouselCard {
	return &NullableMessageCarouselCard{value: val, isSet: true}
}

func (v NullableMessageCarouselCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCarouselCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
