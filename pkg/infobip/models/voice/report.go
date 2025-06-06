/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voice

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the Report type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Report{}

// Report Arrays of delivery reports, one object for every message.
type Report struct {
	// The ID that uniquely identifies the bulk of messages.
	BulkId *string
	// The ID that uniquely identifies the message sent.
	MessageId *string
	// The sender ID which can be alphanumeric or numeric.
	From *string
	// Destination address of the voice message.
	To *string
	// Date and time when the voice message was initiated. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	SentAt *string
	// Mobile country and network codes.
	MccMnc *string
	// Custom data sent over to the notifyUrl.
	CallbackData *string
	VoiceCall    *VoiceData
	Price        *Price
	Status       *SingleMessageStatus
	Error        *VoiceError
}

// NewReport instantiates a new Report object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewReport() *Report {
	this := Report{}
	return &this
}

// NewReportWithDefaults instantiates a new Report object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportWithDefaults() *Report {
	this := Report{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *Report) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *Report) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *Report) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *Report) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *Report) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *Report) SetMessageId(v string) {
	o.MessageId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Report) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Report) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *Report) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *Report) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *Report) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *Report) SetTo(v string) {
	o.To = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *Report) GetSentAt() string {
	if o == nil || IsNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetSentAtOk() (*string, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *Report) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *Report) SetSentAt(v string) {
	o.SentAt = &v
}

// GetMccMnc returns the MccMnc field value if set, zero value otherwise.
func (o *Report) GetMccMnc() string {
	if o == nil || IsNil(o.MccMnc) {
		var ret string
		return ret
	}
	return *o.MccMnc
}

// GetMccMncOk returns a tuple with the MccMnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetMccMncOk() (*string, bool) {
	if o == nil || IsNil(o.MccMnc) {
		return nil, false
	}
	return o.MccMnc, true
}

// HasMccMnc returns a boolean if a field has been set.
func (o *Report) HasMccMnc() bool {
	if o != nil && !IsNil(o.MccMnc) {
		return true
	}

	return false
}

// SetMccMnc gets a reference to the given string and assigns it to the MccMnc field.
func (o *Report) SetMccMnc(v string) {
	o.MccMnc = &v
}

// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *Report) GetCallbackData() string {
	if o == nil || IsNil(o.CallbackData) {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetCallbackDataOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackData) {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *Report) HasCallbackData() bool {
	if o != nil && !IsNil(o.CallbackData) {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *Report) SetCallbackData(v string) {
	o.CallbackData = &v
}

// GetVoiceCall returns the VoiceCall field value if set, zero value otherwise.
func (o *Report) GetVoiceCall() VoiceData {
	if o == nil || IsNil(o.VoiceCall) {
		var ret VoiceData
		return ret
	}
	return *o.VoiceCall
}

// GetVoiceCallOk returns a tuple with the VoiceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetVoiceCallOk() (*VoiceData, bool) {
	if o == nil || IsNil(o.VoiceCall) {
		return nil, false
	}
	return o.VoiceCall, true
}

// HasVoiceCall returns a boolean if a field has been set.
func (o *Report) HasVoiceCall() bool {
	if o != nil && !IsNil(o.VoiceCall) {
		return true
	}

	return false
}

// SetVoiceCall gets a reference to the given VoiceData and assigns it to the VoiceCall field.
func (o *Report) SetVoiceCall(v VoiceData) {
	o.VoiceCall = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Report) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Report) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *Report) SetPrice(v Price) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Report) GetStatus() SingleMessageStatus {
	if o == nil || IsNil(o.Status) {
		var ret SingleMessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetStatusOk() (*SingleMessageStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Report) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SingleMessageStatus and assigns it to the Status field.
func (o *Report) SetStatus(v SingleMessageStatus) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Report) GetError() VoiceError {
	if o == nil || IsNil(o.Error) {
		var ret VoiceError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetErrorOk() (*VoiceError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Report) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given VoiceError and assigns it to the Error field.
func (o *Report) SetError(v VoiceError) {
	o.Error = &v
}

func (o Report) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Report) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.MccMnc) {
		toSerialize["mccMnc"] = o.MccMnc
	}
	if !IsNil(o.CallbackData) {
		toSerialize["callbackData"] = o.CallbackData
	}
	if !IsNil(o.VoiceCall) {
		toSerialize["voiceCall"] = o.VoiceCall
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
	return toSerialize, nil
}

type NullableReport struct {
	value *Report
	isSet bool
}

func (v NullableReport) Get() *Report {
	return v.value
}

func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

func (v NullableReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
