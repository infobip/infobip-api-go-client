/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sms

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the MessagePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagePrice{}

// MessagePrice Sent message price.
type MessagePrice struct {
	// Price per one message.
	PricePerMessage *float32
	// The currency in which the price is expressed.
	Currency *string
}

// NewMessagePrice instantiates a new MessagePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMessagePrice() *MessagePrice {
	this := MessagePrice{}
	return &this
}

// NewMessagePriceWithDefaults instantiates a new MessagePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagePriceWithDefaults() *MessagePrice {
	this := MessagePrice{}

	return &this
}

// GetPricePerMessage returns the PricePerMessage field value if set, zero value otherwise.
func (o *MessagePrice) GetPricePerMessage() float32 {
	if o == nil || IsNil(o.PricePerMessage) {
		var ret float32
		return ret
	}
	return *o.PricePerMessage
}

// GetPricePerMessageOk returns a tuple with the PricePerMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePrice) GetPricePerMessageOk() (*float32, bool) {
	if o == nil || IsNil(o.PricePerMessage) {
		return nil, false
	}
	return o.PricePerMessage, true
}

// HasPricePerMessage returns a boolean if a field has been set.
func (o *MessagePrice) HasPricePerMessage() bool {
	if o != nil && !IsNil(o.PricePerMessage) {
		return true
	}

	return false
}

// SetPricePerMessage gets a reference to the given float32 and assigns it to the PricePerMessage field.
func (o *MessagePrice) SetPricePerMessage(v float32) {
	o.PricePerMessage = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MessagePrice) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagePrice) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MessagePrice) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MessagePrice) SetCurrency(v string) {
	o.Currency = &v
}

func (o MessagePrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PricePerMessage) {
		toSerialize["pricePerMessage"] = o.PricePerMessage
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableMessagePrice struct {
	value *MessagePrice
	isSet bool
}

func (v NullableMessagePrice) Get() *MessagePrice {
	return v.value
}

func (v *NullableMessagePrice) Set(val *MessagePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagePrice(val *MessagePrice) *NullableMessagePrice {
	return &NullableMessagePrice{value: val, isSet: true}
}

func (v NullableMessagePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
