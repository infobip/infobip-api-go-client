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

// checks if the TemplateFailover type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateFailover{}

// TemplateFailover Configuration of a single failover step
type TemplateFailover struct {
	Channel OutboundTemplateChannel
	// Sender for channel specified above.
	Sender         string
	Template       Template
	Content        *TemplateMessageContent
	ValidityPeriod *ValidityPeriod
}

type _TemplateFailover TemplateFailover

// NewTemplateFailover instantiates a new TemplateFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFailover(channel OutboundTemplateChannel, sender string, template Template) *TemplateFailover {
	this := TemplateFailover{}
	this.Channel = channel
	this.Sender = sender
	this.Template = template
	return &this
}

// NewTemplateFailoverWithDefaults instantiates a new TemplateFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFailoverWithDefaults() *TemplateFailover {
	this := TemplateFailover{}
	return &this
}

// GetChannel returns the Channel field value
func (o *TemplateFailover) GetChannel() OutboundTemplateChannel {
	if o == nil {
		var ret OutboundTemplateChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *TemplateFailover) GetChannelOk() (*OutboundTemplateChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *TemplateFailover) SetChannel(v OutboundTemplateChannel) {
	o.Channel = v
}

// GetSender returns the Sender field value
func (o *TemplateFailover) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *TemplateFailover) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *TemplateFailover) SetSender(v string) {
	o.Sender = v
}

// GetTemplate returns the Template field value
func (o *TemplateFailover) GetTemplate() Template {
	if o == nil {
		var ret Template
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *TemplateFailover) GetTemplateOk() (*Template, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *TemplateFailover) SetTemplate(v Template) {
	o.Template = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TemplateFailover) GetContent() TemplateMessageContent {
	if o == nil || IsNil(o.Content) {
		var ret TemplateMessageContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFailover) GetContentOk() (*TemplateMessageContent, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TemplateFailover) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given TemplateMessageContent and assigns it to the Content field.
func (o *TemplateFailover) SetContent(v TemplateMessageContent) {
	o.Content = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *TemplateFailover) GetValidityPeriod() ValidityPeriod {
	if o == nil || IsNil(o.ValidityPeriod) {
		var ret ValidityPeriod
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFailover) GetValidityPeriodOk() (*ValidityPeriod, bool) {
	if o == nil || IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *TemplateFailover) HasValidityPeriod() bool {
	if o != nil && !IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given ValidityPeriod and assigns it to the ValidityPeriod field.
func (o *TemplateFailover) SetValidityPeriod(v ValidityPeriod) {
	o.ValidityPeriod = &v
}

func (o TemplateFailover) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateFailover) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["sender"] = o.Sender
	toSerialize["template"] = o.Template
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	return toSerialize, nil
}

type NullableTemplateFailover struct {
	value *TemplateFailover
	isSet bool
}

func (v NullableTemplateFailover) Get() *TemplateFailover {
	return v.value
}

func (v *NullableTemplateFailover) Set(val *TemplateFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFailover(val *TemplateFailover) *NullableTemplateFailover {
	return &NullableTemplateFailover{value: val, isSet: true}
}

func (v NullableTemplateFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
