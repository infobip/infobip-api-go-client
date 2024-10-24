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

// checks if the InboundTypingStoppedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundTypingStoppedEvent{}

// InboundTypingStoppedEvent struct for InboundTypingStoppedEvent
type InboundTypingStoppedEvent struct {
	Event InboundEventType
	Channel InboundTypingIndicatorChannel
	// Identifier (usually number) of the party which sent the message.
	Sender string
	// Sender provided during the activation process.
	Destination string
	// Date and time when Infobip received the message.
	ReceivedAt Time
	// The ID that uniquely identifies the received message.
	MessageId string
	// Message ID of the MT message that this MO message is a response to.
	PairedMessageId *string
	// Value of the `callbackData` field from the MT message (if exists) or from the MO Action setup (if exists).
	CallbackData *string
	Platform     *Platform
}

type _InboundTypingStoppedEvent InboundTypingStoppedEvent

// NewInboundTypingStoppedEvent instantiates a new InboundTypingStoppedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundTypingStoppedEvent(channel InboundTypingIndicatorChannel, sender string, destination string, receivedAt Time, messageId string) *InboundTypingStoppedEvent {
	this := InboundTypingStoppedEvent{}
	this.Event = INBOUNDEVENTTYPE_TYPING_STOPPED
	this.Channel = channel
	this.Sender = sender
	this.Destination = destination
	this.ReceivedAt = receivedAt
	this.MessageId = messageId
	return &this
}

// NewInboundTypingStoppedEventWithDefaults instantiates a new InboundTypingStoppedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundTypingStoppedEventWithDefaults() *InboundTypingStoppedEvent {
	this := InboundTypingStoppedEvent{}
	return &this
}

// GetChannel returns the Channel field value
func (o *InboundTypingStoppedEvent) GetChannel() InboundTypingIndicatorChannel {
	if o == nil {
		var ret InboundTypingIndicatorChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetChannelOk() (*InboundTypingIndicatorChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *InboundTypingStoppedEvent) SetChannel(v InboundTypingIndicatorChannel) {
	o.Channel = v
}

// GetSender returns the Sender field value
func (o *InboundTypingStoppedEvent) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *InboundTypingStoppedEvent) SetSender(v string) {
	o.Sender = v
}

// GetDestination returns the Destination field value
func (o *InboundTypingStoppedEvent) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InboundTypingStoppedEvent) SetDestination(v string) {
	o.Destination = v
}

// GetReceivedAt returns the ReceivedAt field value
func (o *InboundTypingStoppedEvent) GetReceivedAt() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.ReceivedAt
}

// GetReceivedAtOk returns a tuple with the ReceivedAt field value
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetReceivedAtOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivedAt, true
}

// SetReceivedAt sets field value
func (o *InboundTypingStoppedEvent) SetReceivedAt(v Time) {
	o.ReceivedAt = v
}

// GetMessageId returns the MessageId field value
func (o *InboundTypingStoppedEvent) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *InboundTypingStoppedEvent) SetMessageId(v string) {
	o.MessageId = v
}

// GetPairedMessageId returns the PairedMessageId field value if set, zero value otherwise.
func (o *InboundTypingStoppedEvent) GetPairedMessageId() string {
	if o == nil || IsNil(o.PairedMessageId) {
		var ret string
		return ret
	}
	return *o.PairedMessageId
}

// GetPairedMessageIdOk returns a tuple with the PairedMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetPairedMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PairedMessageId) {
		return nil, false
	}
	return o.PairedMessageId, true
}

// HasPairedMessageId returns a boolean if a field has been set.
func (o *InboundTypingStoppedEvent) HasPairedMessageId() bool {
	if o != nil && !IsNil(o.PairedMessageId) {
		return true
	}

	return false
}

// SetPairedMessageId gets a reference to the given string and assigns it to the PairedMessageId field.
func (o *InboundTypingStoppedEvent) SetPairedMessageId(v string) {
	o.PairedMessageId = &v
}

// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *InboundTypingStoppedEvent) GetCallbackData() string {
	if o == nil || IsNil(o.CallbackData) {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetCallbackDataOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackData) {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *InboundTypingStoppedEvent) HasCallbackData() bool {
	if o != nil && !IsNil(o.CallbackData) {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *InboundTypingStoppedEvent) SetCallbackData(v string) {
	o.CallbackData = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *InboundTypingStoppedEvent) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundTypingStoppedEvent) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *InboundTypingStoppedEvent) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *InboundTypingStoppedEvent) SetPlatform(v Platform) {
	o.Platform = &v
}

func (o InboundTypingStoppedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundTypingStoppedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["channel"] = o.Channel
	toSerialize["sender"] = o.Sender
	toSerialize["destination"] = o.Destination
	toSerialize["receivedAt"] = o.ReceivedAt
	toSerialize["messageId"] = o.MessageId
	if !IsNil(o.PairedMessageId) {
		toSerialize["pairedMessageId"] = o.PairedMessageId
	}
	if !IsNil(o.CallbackData) {
		toSerialize["callbackData"] = o.CallbackData
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	return toSerialize, nil
}

type NullableInboundTypingStoppedEvent struct {
	value *InboundTypingStoppedEvent
	isSet bool
}

func (v NullableInboundTypingStoppedEvent) Get() *InboundTypingStoppedEvent {
	return v.value
}

func (v *NullableInboundTypingStoppedEvent) Set(val *InboundTypingStoppedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundTypingStoppedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundTypingStoppedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundTypingStoppedEvent(val *InboundTypingStoppedEvent) *NullableInboundTypingStoppedEvent {
	return &NullableInboundTypingStoppedEvent{value: val, isSet: true}
}

func (v NullableInboundTypingStoppedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundTypingStoppedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}