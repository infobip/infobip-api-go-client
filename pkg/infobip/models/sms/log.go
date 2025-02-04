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

// checks if the Log type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Log{}

// Log An array of message log results, one object per each message log entry.
type Log struct {
	// The sender ID which can be alphanumeric or numeric.
	Sender *string
	// Message destination address.
	Destination *string
	// Unique ID assigned to the request if messaging multiple recipients or sending multiple messages via a single API request.
	BulkId *string
	// Unique message ID for which a log is requested.
	MessageId *string
	// Date and time when the message was sent. Has the following format: yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	SentAt *Time
	// Date and time when the Infobip services finished processing the message (i.e., delivered to the destination, network, etc.). Has the following format: yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	DoneAt *Time
	// The number of messages content was split to.
	MessageCount *int32
	Price        *MessagePrice
	Status       *MessageStatus
	Error        *MessageError
	Platform     *Platform
	Content      *MessageContent
	// ID of a campaign that was sent in the message.
	CampaignReferenceId *string
	// Mobile country and network codes.
	MccMnc *string
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewLog() *Log {
	this := Log{}
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}

	return &this
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *Log) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *Log) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *Log) SetSender(v string) {
	o.Sender = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Log) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Log) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *Log) SetDestination(v string) {
	o.Destination = &v
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *Log) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *Log) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *Log) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *Log) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *Log) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *Log) SetMessageId(v string) {
	o.MessageId = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *Log) GetSentAt() Time {
	if o == nil || IsNil(o.SentAt) {
		var ret Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetSentAtOk() (*Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *Log) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given Time and assigns it to the SentAt field.
func (o *Log) SetSentAt(v Time) {
	o.SentAt = &v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise.
func (o *Log) GetDoneAt() Time {
	if o == nil || IsNil(o.DoneAt) {
		var ret Time
		return ret
	}
	return *o.DoneAt
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetDoneAtOk() (*Time, bool) {
	if o == nil || IsNil(o.DoneAt) {
		return nil, false
	}
	return o.DoneAt, true
}

// HasDoneAt returns a boolean if a field has been set.
func (o *Log) HasDoneAt() bool {
	if o != nil && !IsNil(o.DoneAt) {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given Time and assigns it to the DoneAt field.
func (o *Log) SetDoneAt(v Time) {
	o.DoneAt = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *Log) GetMessageCount() int32 {
	if o == nil || IsNil(o.MessageCount) {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageCount) {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *Log) HasMessageCount() bool {
	if o != nil && !IsNil(o.MessageCount) {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *Log) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Log) GetPrice() MessagePrice {
	if o == nil || IsNil(o.Price) {
		var ret MessagePrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetPriceOk() (*MessagePrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Log) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given MessagePrice and assigns it to the Price field.
func (o *Log) SetPrice(v MessagePrice) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Log) GetStatus() MessageStatus {
	if o == nil || IsNil(o.Status) {
		var ret MessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetStatusOk() (*MessageStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Log) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MessageStatus and assigns it to the Status field.
func (o *Log) SetStatus(v MessageStatus) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Log) GetError() MessageError {
	if o == nil || IsNil(o.Error) {
		var ret MessageError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetErrorOk() (*MessageError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Log) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MessageError and assigns it to the Error field.
func (o *Log) SetError(v MessageError) {
	o.Error = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *Log) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *Log) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *Log) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Log) GetContent() MessageContent {
	if o == nil || IsNil(o.Content) {
		var ret MessageContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetContentOk() (*MessageContent, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Log) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given MessageContent and assigns it to the Content field.
func (o *Log) SetContent(v MessageContent) {
	o.Content = &v
}

// GetCampaignReferenceId returns the CampaignReferenceId field value if set, zero value otherwise.
func (o *Log) GetCampaignReferenceId() string {
	if o == nil || IsNil(o.CampaignReferenceId) {
		var ret string
		return ret
	}
	return *o.CampaignReferenceId
}

// GetCampaignReferenceIdOk returns a tuple with the CampaignReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetCampaignReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignReferenceId) {
		return nil, false
	}
	return o.CampaignReferenceId, true
}

// HasCampaignReferenceId returns a boolean if a field has been set.
func (o *Log) HasCampaignReferenceId() bool {
	if o != nil && !IsNil(o.CampaignReferenceId) {
		return true
	}

	return false
}

// SetCampaignReferenceId gets a reference to the given string and assigns it to the CampaignReferenceId field.
func (o *Log) SetCampaignReferenceId(v string) {
	o.CampaignReferenceId = &v
}

// GetMccMnc returns the MccMnc field value if set, zero value otherwise.
func (o *Log) GetMccMnc() string {
	if o == nil || IsNil(o.MccMnc) {
		var ret string
		return ret
	}
	return *o.MccMnc
}

// GetMccMncOk returns a tuple with the MccMnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetMccMncOk() (*string, bool) {
	if o == nil || IsNil(o.MccMnc) {
		return nil, false
	}
	return o.MccMnc, true
}

// HasMccMnc returns a boolean if a field has been set.
func (o *Log) HasMccMnc() bool {
	if o != nil && !IsNil(o.MccMnc) {
		return true
	}

	return false
}

// SetMccMnc gets a reference to the given string and assigns it to the MccMnc field.
func (o *Log) SetMccMnc(v string) {
	o.MccMnc = &v
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Log) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.DoneAt) {
		toSerialize["doneAt"] = o.DoneAt
	}
	if !IsNil(o.MessageCount) {
		toSerialize["messageCount"] = o.MessageCount
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CampaignReferenceId) {
		toSerialize["campaignReferenceId"] = o.CampaignReferenceId
	}
	if !IsNil(o.MccMnc) {
		toSerialize["mccMnc"] = o.MccMnc
	}
	return toSerialize, nil
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
