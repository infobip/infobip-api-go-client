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

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message An array of message objects of a single message or multiple messages sent under one bulk ID.
type Message struct {
	// The sender ID. It can be alphanumeric or numeric (e.g., `CompanyName`). Make sure you don't exceed [character limit](https://www.infobip.com/docs/sms/get-started#sender-names).
	Sender *string
	// An array of destination objects for where messages are being sent. A valid destination is required.
	Destinations []Destination
	Content      LogContent
	Options      *MessageOptions
	Webhooks     *Webhooks
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(destinations []Destination, content LogContent) *Message {
	this := Message{}
	this.Destinations = destinations
	this.Content = content
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *Message) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *Message) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *Message) SetSender(v string) {
	o.Sender = &v
}

// GetDestinations returns the Destinations field value
func (o *Message) GetDestinations() []Destination {
	if o == nil {
		var ret []Destination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *Message) GetDestinationsOk() ([]Destination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *Message) SetDestinations(v []Destination) {
	o.Destinations = v
}

// GetContent returns the Content field value
func (o *Message) GetContent() LogContent {
	if o == nil {
		var ret LogContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Message) GetContentOk() (*LogContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Message) SetContent(v LogContent) {
	o.Content = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Message) GetOptions() MessageOptions {
	if o == nil || IsNil(o.Options) {
		var ret MessageOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetOptionsOk() (*MessageOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Message) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given MessageOptions and assigns it to the Options field.
func (o *Message) SetOptions(v MessageOptions) {
	o.Options = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *Message) GetWebhooks() Webhooks {
	if o == nil || IsNil(o.Webhooks) {
		var ret Webhooks
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetWebhooksOk() (*Webhooks, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *Message) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given Webhooks and assigns it to the Webhooks field.
func (o *Message) SetWebhooks(v Webhooks) {
	o.Webhooks = &v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	toSerialize["destinations"] = o.Destinations
	toSerialize["content"] = o.Content
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	return toSerialize, nil
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
