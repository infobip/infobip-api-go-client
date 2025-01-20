/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the DeliveryReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryReport{}

// DeliveryReport struct for DeliveryReport
type DeliveryReport struct {
	// The ID that uniquely identifies a list of email messages. This is either defined by user in the request or auto generated.
	BulkId *string
	Price  *Price
	Status *Status
	Error  *Error
	// The ID that uniquely identifies the email sent to the recipient.
	MessageId *string
	// Delivery date and time.
	DoneAt *Time
	// The number of emails sent.
	SmsCount *int32
	// Send date and time. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	SentAt *Time
	// Contains the link to the HTML sent to recipient. This will be present only if the _view in browser_ feature is used in the email. (Please note that this feature is not activated automatically for Email traffic sent over API. If you would like to utilize it please reach out to your Infobip person of contact.)
	BrowserLink *string
	// The IP address that was used to send out the email.
	SendingIp *string
	// Callback data sent through `callbackData` field in fully featured email.
	CallbackData *string
	// Destination email address.
	To *string
}

// NewDeliveryReport instantiates a new DeliveryReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDeliveryReport() *DeliveryReport {
	this := DeliveryReport{}
	return &this
}

// NewDeliveryReportWithDefaults instantiates a new DeliveryReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryReportWithDefaults() *DeliveryReport {
	this := DeliveryReport{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *DeliveryReport) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *DeliveryReport) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *DeliveryReport) SetBulkId(v string) {
	o.BulkId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *DeliveryReport) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *DeliveryReport) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *DeliveryReport) SetPrice(v Price) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeliveryReport) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeliveryReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *DeliveryReport) SetStatus(v Status) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DeliveryReport) GetError() Error {
	if o == nil || IsNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetErrorOk() (*Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DeliveryReport) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *DeliveryReport) SetError(v Error) {
	o.Error = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *DeliveryReport) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *DeliveryReport) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *DeliveryReport) SetMessageId(v string) {
	o.MessageId = &v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise.
func (o *DeliveryReport) GetDoneAt() Time {
	if o == nil || IsNil(o.DoneAt) {
		var ret Time
		return ret
	}
	return *o.DoneAt
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetDoneAtOk() (*Time, bool) {
	if o == nil || IsNil(o.DoneAt) {
		return nil, false
	}
	return o.DoneAt, true
}

// HasDoneAt returns a boolean if a field has been set.
func (o *DeliveryReport) HasDoneAt() bool {
	if o != nil && !IsNil(o.DoneAt) {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given Time and assigns it to the DoneAt field.
func (o *DeliveryReport) SetDoneAt(v Time) {
	o.DoneAt = &v
}

// GetSmsCount returns the SmsCount field value if set, zero value otherwise.
func (o *DeliveryReport) GetSmsCount() int32 {
	if o == nil || IsNil(o.SmsCount) {
		var ret int32
		return ret
	}
	return *o.SmsCount
}

// GetSmsCountOk returns a tuple with the SmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetSmsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SmsCount) {
		return nil, false
	}
	return o.SmsCount, true
}

// HasSmsCount returns a boolean if a field has been set.
func (o *DeliveryReport) HasSmsCount() bool {
	if o != nil && !IsNil(o.SmsCount) {
		return true
	}

	return false
}

// SetSmsCount gets a reference to the given int32 and assigns it to the SmsCount field.
func (o *DeliveryReport) SetSmsCount(v int32) {
	o.SmsCount = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *DeliveryReport) GetSentAt() Time {
	if o == nil || IsNil(o.SentAt) {
		var ret Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetSentAtOk() (*Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *DeliveryReport) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given Time and assigns it to the SentAt field.
func (o *DeliveryReport) SetSentAt(v Time) {
	o.SentAt = &v
}

// GetBrowserLink returns the BrowserLink field value if set, zero value otherwise.
func (o *DeliveryReport) GetBrowserLink() string {
	if o == nil || IsNil(o.BrowserLink) {
		var ret string
		return ret
	}
	return *o.BrowserLink
}

// GetBrowserLinkOk returns a tuple with the BrowserLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetBrowserLinkOk() (*string, bool) {
	if o == nil || IsNil(o.BrowserLink) {
		return nil, false
	}
	return o.BrowserLink, true
}

// HasBrowserLink returns a boolean if a field has been set.
func (o *DeliveryReport) HasBrowserLink() bool {
	if o != nil && !IsNil(o.BrowserLink) {
		return true
	}

	return false
}

// SetBrowserLink gets a reference to the given string and assigns it to the BrowserLink field.
func (o *DeliveryReport) SetBrowserLink(v string) {
	o.BrowserLink = &v
}

// GetSendingIp returns the SendingIp field value if set, zero value otherwise.
func (o *DeliveryReport) GetSendingIp() string {
	if o == nil || IsNil(o.SendingIp) {
		var ret string
		return ret
	}
	return *o.SendingIp
}

// GetSendingIpOk returns a tuple with the SendingIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetSendingIpOk() (*string, bool) {
	if o == nil || IsNil(o.SendingIp) {
		return nil, false
	}
	return o.SendingIp, true
}

// HasSendingIp returns a boolean if a field has been set.
func (o *DeliveryReport) HasSendingIp() bool {
	if o != nil && !IsNil(o.SendingIp) {
		return true
	}

	return false
}

// SetSendingIp gets a reference to the given string and assigns it to the SendingIp field.
func (o *DeliveryReport) SetSendingIp(v string) {
	o.SendingIp = &v
}

// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *DeliveryReport) GetCallbackData() string {
	if o == nil || IsNil(o.CallbackData) {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetCallbackDataOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackData) {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *DeliveryReport) HasCallbackData() bool {
	if o != nil && !IsNil(o.CallbackData) {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *DeliveryReport) SetCallbackData(v string) {
	o.CallbackData = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *DeliveryReport) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReport) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *DeliveryReport) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *DeliveryReport) SetTo(v string) {
	o.To = &v
}

func (o DeliveryReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
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
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.DoneAt) {
		toSerialize["doneAt"] = o.DoneAt
	}
	if !IsNil(o.SmsCount) {
		toSerialize["smsCount"] = o.SmsCount
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.BrowserLink) {
		toSerialize["browserLink"] = o.BrowserLink
	}
	if !IsNil(o.SendingIp) {
		toSerialize["sendingIp"] = o.SendingIp
	}
	if !IsNil(o.CallbackData) {
		toSerialize["callbackData"] = o.CallbackData
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableDeliveryReport struct {
	value *DeliveryReport
	isSet bool
}

func (v NullableDeliveryReport) Get() *DeliveryReport {
	return v.value
}

func (v *NullableDeliveryReport) Set(val *DeliveryReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryReport(val *DeliveryReport) *NullableDeliveryReport {
	return &NullableDeliveryReport{value: val, isSet: true}
}

func (v NullableDeliveryReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
