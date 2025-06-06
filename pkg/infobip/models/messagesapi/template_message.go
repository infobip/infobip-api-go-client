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

// checks if the TemplateMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateMessage{}

// TemplateMessage Represents a single template message.
type TemplateMessage struct {
	Channel OutboundTemplateChannel
	// The sender ID. It can be alphanumeric or numeric (e.g., `CompanyName`). Make sure you don't exceed [character limit](https://www.infobip.com/docs/sms/get-started#sender-names).
	Sender string
	// Array of destination objects for where messages are being sent. A valid destination is required.
	Destinations []MessageDestination
	Template     Template
	Content      *TemplateMessageContent
	Options      *MessageOptions
	Webhooks     *OttWebhooks
	// Provides options for configuring a message failover. When message fails it will be sent over channels in order specified in an array. Make sure to provide correct sender and destinations specified as `Channels Destination` for each channel.
	Failover []BaseFailover
}

type _TemplateMessage TemplateMessage

// NewTemplateMessage instantiates a new TemplateMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewTemplateMessage(channel OutboundTemplateChannel, sender string, destinations []MessageDestination, template Template) *TemplateMessage {
	this := TemplateMessage{}
	this.Channel = channel
	this.Sender = sender
	this.Destinations = destinations
	this.Template = template
	return &this
}

// NewTemplateMessageWithDefaults instantiates a new TemplateMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateMessageWithDefaults() *TemplateMessage {
	this := TemplateMessage{}

	return &this
}

// GetChannel returns the Channel field value
func (o *TemplateMessage) GetChannel() OutboundTemplateChannel {
	if o == nil {
		var ret OutboundTemplateChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetChannelOk() (*OutboundTemplateChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *TemplateMessage) SetChannel(v OutboundTemplateChannel) {
	o.Channel = v
}

// GetSender returns the Sender field value
func (o *TemplateMessage) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *TemplateMessage) SetSender(v string) {
	o.Sender = v
}

// GetDestinations returns the Destinations field value
func (o *TemplateMessage) GetDestinations() []MessageDestination {
	if o == nil {
		var ret []MessageDestination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetDestinationsOk() ([]MessageDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *TemplateMessage) SetDestinations(v []MessageDestination) {
	o.Destinations = v
}

// GetTemplate returns the Template field value
func (o *TemplateMessage) GetTemplate() Template {
	if o == nil {
		var ret Template
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetTemplateOk() (*Template, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *TemplateMessage) SetTemplate(v Template) {
	o.Template = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TemplateMessage) GetContent() TemplateMessageContent {
	if o == nil || IsNil(o.Content) {
		var ret TemplateMessageContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetContentOk() (*TemplateMessageContent, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TemplateMessage) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given TemplateMessageContent and assigns it to the Content field.
func (o *TemplateMessage) SetContent(v TemplateMessageContent) {
	o.Content = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TemplateMessage) GetOptions() MessageOptions {
	if o == nil || IsNil(o.Options) {
		var ret MessageOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetOptionsOk() (*MessageOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TemplateMessage) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given MessageOptions and assigns it to the Options field.
func (o *TemplateMessage) SetOptions(v MessageOptions) {
	o.Options = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *TemplateMessage) GetWebhooks() OttWebhooks {
	if o == nil || IsNil(o.Webhooks) {
		var ret OttWebhooks
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetWebhooksOk() (*OttWebhooks, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *TemplateMessage) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given OttWebhooks and assigns it to the Webhooks field.
func (o *TemplateMessage) SetWebhooks(v OttWebhooks) {
	o.Webhooks = &v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *TemplateMessage) GetFailover() []BaseFailover {
	if o == nil || IsNil(o.Failover) {
		var ret []BaseFailover
		return ret
	}
	return o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessage) GetFailoverOk() ([]BaseFailover, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *TemplateMessage) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given []BaseFailover and assigns it to the Failover field.
func (o *TemplateMessage) SetFailover(v []BaseFailover) {
	o.Failover = v
}

func (o TemplateMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["sender"] = o.Sender
	toSerialize["destinations"] = o.Destinations
	toSerialize["template"] = o.Template
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	if !IsNil(o.Failover) {
		toSerialize["failover"] = o.Failover
	}
	return toSerialize, nil
}

type NullableTemplateMessage struct {
	value *TemplateMessage
	isSet bool
}

func (v NullableTemplateMessage) Get() *TemplateMessage {
	return v.value
}

func (v *NullableTemplateMessage) Set(val *TemplateMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateMessage(val *TemplateMessage) *NullableTemplateMessage {
	return &NullableTemplateMessage{value: val, isSet: true}
}

func (v NullableTemplateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
