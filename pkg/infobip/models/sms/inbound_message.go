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

// checks if the InboundMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundMessage{}

// InboundMessage struct for InboundMessage
type InboundMessage struct {
	// Application id linked to the message. For more details, see our [documentation](https://www.infobip.com/docs/cpaas-x/application-and-entity-management).
	ApplicationId *string
	// Custom callback data sent over the notifyUrl.
	CallbackData *string
	// ID that allows you to track, analyze, and show an aggregated overview and the performance of individual campaigns per sending channel.
	CampaignReferenceId *string
	// Content of the message without a keyword (if a keyword was sent).
	CleanText *string
	// Entity id linked to the message. For more details, see our [documentation](https://www.infobip.com/docs/cpaas-x/application-and-entity-management).
	EntityId *string
	// Sender ID that can be alphanumeric or numeric.
	From *string
	// Keyword extracted from the message content.
	Keyword *string
	// Unique message ID.
	MessageId *string
	Price     *Price
	// Indicates when the Infobip platform received the message. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	ReceivedAt *Time
	// The number of characters within a message
	SmsCount *int32
	// Full content of the message.
	Text *string
	// The destination address of the message.
	To *string
}

// NewInboundMessage instantiates a new InboundMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundMessage() *InboundMessage {
	this := InboundMessage{}
	return &this
}

// NewInboundMessageWithDefaults instantiates a new InboundMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundMessageWithDefaults() *InboundMessage {
	this := InboundMessage{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *InboundMessage) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *InboundMessage) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *InboundMessage) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *InboundMessage) GetCallbackData() string {
	if o == nil || IsNil(o.CallbackData) {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetCallbackDataOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackData) {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *InboundMessage) HasCallbackData() bool {
	if o != nil && !IsNil(o.CallbackData) {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *InboundMessage) SetCallbackData(v string) {
	o.CallbackData = &v
}

// GetCampaignReferenceId returns the CampaignReferenceId field value if set, zero value otherwise.
func (o *InboundMessage) GetCampaignReferenceId() string {
	if o == nil || IsNil(o.CampaignReferenceId) {
		var ret string
		return ret
	}
	return *o.CampaignReferenceId
}

// GetCampaignReferenceIdOk returns a tuple with the CampaignReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetCampaignReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignReferenceId) {
		return nil, false
	}
	return o.CampaignReferenceId, true
}

// HasCampaignReferenceId returns a boolean if a field has been set.
func (o *InboundMessage) HasCampaignReferenceId() bool {
	if o != nil && !IsNil(o.CampaignReferenceId) {
		return true
	}

	return false
}

// SetCampaignReferenceId gets a reference to the given string and assigns it to the CampaignReferenceId field.
func (o *InboundMessage) SetCampaignReferenceId(v string) {
	o.CampaignReferenceId = &v
}

// GetCleanText returns the CleanText field value if set, zero value otherwise.
func (o *InboundMessage) GetCleanText() string {
	if o == nil || IsNil(o.CleanText) {
		var ret string
		return ret
	}
	return *o.CleanText
}

// GetCleanTextOk returns a tuple with the CleanText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetCleanTextOk() (*string, bool) {
	if o == nil || IsNil(o.CleanText) {
		return nil, false
	}
	return o.CleanText, true
}

// HasCleanText returns a boolean if a field has been set.
func (o *InboundMessage) HasCleanText() bool {
	if o != nil && !IsNil(o.CleanText) {
		return true
	}

	return false
}

// SetCleanText gets a reference to the given string and assigns it to the CleanText field.
func (o *InboundMessage) SetCleanText(v string) {
	o.CleanText = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *InboundMessage) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *InboundMessage) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *InboundMessage) SetEntityId(v string) {
	o.EntityId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *InboundMessage) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *InboundMessage) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *InboundMessage) SetFrom(v string) {
	o.From = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *InboundMessage) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *InboundMessage) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *InboundMessage) SetKeyword(v string) {
	o.Keyword = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *InboundMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *InboundMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *InboundMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InboundMessage) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InboundMessage) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *InboundMessage) SetPrice(v Price) {
	o.Price = &v
}

// GetReceivedAt returns the ReceivedAt field value if set, zero value otherwise.
func (o *InboundMessage) GetReceivedAt() Time {
	if o == nil || IsNil(o.ReceivedAt) {
		var ret Time
		return ret
	}
	return *o.ReceivedAt
}

// GetReceivedAtOk returns a tuple with the ReceivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetReceivedAtOk() (*Time, bool) {
	if o == nil || IsNil(o.ReceivedAt) {
		return nil, false
	}
	return o.ReceivedAt, true
}

// HasReceivedAt returns a boolean if a field has been set.
func (o *InboundMessage) HasReceivedAt() bool {
	if o != nil && !IsNil(o.ReceivedAt) {
		return true
	}

	return false
}

// SetReceivedAt gets a reference to the given Time and assigns it to the ReceivedAt field.
func (o *InboundMessage) SetReceivedAt(v Time) {
	o.ReceivedAt = &v
}

// GetSmsCount returns the SmsCount field value if set, zero value otherwise.
func (o *InboundMessage) GetSmsCount() int32 {
	if o == nil || IsNil(o.SmsCount) {
		var ret int32
		return ret
	}
	return *o.SmsCount
}

// GetSmsCountOk returns a tuple with the SmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetSmsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SmsCount) {
		return nil, false
	}
	return o.SmsCount, true
}

// HasSmsCount returns a boolean if a field has been set.
func (o *InboundMessage) HasSmsCount() bool {
	if o != nil && !IsNil(o.SmsCount) {
		return true
	}

	return false
}

// SetSmsCount gets a reference to the given int32 and assigns it to the SmsCount field.
func (o *InboundMessage) SetSmsCount(v int32) {
	o.SmsCount = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *InboundMessage) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InboundMessage) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *InboundMessage) SetText(v string) {
	o.Text = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *InboundMessage) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessage) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *InboundMessage) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *InboundMessage) SetTo(v string) {
	o.To = &v
}

func (o InboundMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.CallbackData) {
		toSerialize["callbackData"] = o.CallbackData
	}
	if !IsNil(o.CampaignReferenceId) {
		toSerialize["campaignReferenceId"] = o.CampaignReferenceId
	}
	if !IsNil(o.CleanText) {
		toSerialize["cleanText"] = o.CleanText
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ReceivedAt) {
		toSerialize["receivedAt"] = o.ReceivedAt
	}
	if !IsNil(o.SmsCount) {
		toSerialize["smsCount"] = o.SmsCount
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableInboundMessage struct {
	value *InboundMessage
	isSet bool
}

func (v NullableInboundMessage) Get() *InboundMessage {
	return v.value
}

func (v *NullableInboundMessage) Set(val *InboundMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundMessage(val *InboundMessage) *NullableInboundMessage {
	return &NullableInboundMessage{value: val, isSet: true}
}

func (v NullableInboundMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
