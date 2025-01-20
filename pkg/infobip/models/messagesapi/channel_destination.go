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

// checks if the ChannelDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelDestination{}

// ChannelDestination Array of substitute destinations distinguished by a `channel` they belong to. Only one substitute destination per `channel` is permitted.
type ChannelDestination struct {
	Channel OutboundMessageChannel
	// The destination address of the message associated with given channel. It can be alphanumeric or numeric.
	To string
	// The ID that uniquely identifies the message sent.
	MessageId *string
}

type _ChannelDestination ChannelDestination

// NewChannelDestination instantiates a new ChannelDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewChannelDestination(channel OutboundMessageChannel, to string) *ChannelDestination {
	this := ChannelDestination{}
	this.Channel = channel
	this.To = to
	return &this
}

// NewChannelDestinationWithDefaults instantiates a new ChannelDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelDestinationWithDefaults() *ChannelDestination {
	this := ChannelDestination{}

	return &this
}

// GetChannel returns the Channel field value
func (o *ChannelDestination) GetChannel() OutboundMessageChannel {
	if o == nil {
		var ret OutboundMessageChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ChannelDestination) GetChannelOk() (*OutboundMessageChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ChannelDestination) SetChannel(v OutboundMessageChannel) {
	o.Channel = v
}

// GetTo returns the To field value
func (o *ChannelDestination) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ChannelDestination) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *ChannelDestination) SetTo(v string) {
	o.To = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ChannelDestination) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDestination) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ChannelDestination) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ChannelDestination) SetMessageId(v string) {
	o.MessageId = &v
}

func (o ChannelDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["to"] = o.To
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	return toSerialize, nil
}

type NullableChannelDestination struct {
	value *ChannelDestination
	isSet bool
}

func (v NullableChannelDestination) Get() *ChannelDestination {
	return v.value
}

func (v *NullableChannelDestination) Set(val *ChannelDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelDestination(val *ChannelDestination) *NullableChannelDestination {
	return &NullableChannelDestination{value: val, isSet: true}
}

func (v NullableChannelDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
