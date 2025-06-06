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

// checks if the LogsReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsReport{}

// LogsReport Array of voice message logs, one object per each message.
type LogsReport struct {
	// The ID that uniquely identifies the bulk of messages.
	BulkId *string
	// The ID that uniquely identifies the message sent.
	MessageId *string
	// Destination address of the voice message.
	To *string
	// The sender ID which can be alphanumeric or numeric.
	From *string
	// Content of the voice message that was sent.
	Text *string
	// Date and time when the voice message was initiated. Has the following format: yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	SentAt *Time
	// Date and time when the Infobip services finished processing the voice message (i.e. delivered to the destination, delivered to the destination network, etc.).
	DoneAt *Time
	// Voice message duration in seconds.
	Duration *int32
	// Mobile country and network codes.
	MccMnc *string
	Price  *Price
	Status *SingleMessageStatus
	Error  *VoiceError
}

// NewLogsReport instantiates a new LogsReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewLogsReport() *LogsReport {
	this := LogsReport{}
	return &this
}

// NewLogsReportWithDefaults instantiates a new LogsReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsReportWithDefaults() *LogsReport {
	this := LogsReport{}

	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *LogsReport) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *LogsReport) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *LogsReport) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *LogsReport) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *LogsReport) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *LogsReport) SetMessageId(v string) {
	o.MessageId = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *LogsReport) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LogsReport) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *LogsReport) SetTo(v string) {
	o.To = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *LogsReport) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *LogsReport) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *LogsReport) SetFrom(v string) {
	o.From = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *LogsReport) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *LogsReport) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *LogsReport) SetText(v string) {
	o.Text = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *LogsReport) GetSentAt() Time {
	if o == nil || IsNil(o.SentAt) {
		var ret Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetSentAtOk() (*Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *LogsReport) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given Time and assigns it to the SentAt field.
func (o *LogsReport) SetSentAt(v Time) {
	o.SentAt = &v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise.
func (o *LogsReport) GetDoneAt() Time {
	if o == nil || IsNil(o.DoneAt) {
		var ret Time
		return ret
	}
	return *o.DoneAt
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetDoneAtOk() (*Time, bool) {
	if o == nil || IsNil(o.DoneAt) {
		return nil, false
	}
	return o.DoneAt, true
}

// HasDoneAt returns a boolean if a field has been set.
func (o *LogsReport) HasDoneAt() bool {
	if o != nil && !IsNil(o.DoneAt) {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given Time and assigns it to the DoneAt field.
func (o *LogsReport) SetDoneAt(v Time) {
	o.DoneAt = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LogsReport) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LogsReport) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *LogsReport) SetDuration(v int32) {
	o.Duration = &v
}

// GetMccMnc returns the MccMnc field value if set, zero value otherwise.
func (o *LogsReport) GetMccMnc() string {
	if o == nil || IsNil(o.MccMnc) {
		var ret string
		return ret
	}
	return *o.MccMnc
}

// GetMccMncOk returns a tuple with the MccMnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetMccMncOk() (*string, bool) {
	if o == nil || IsNil(o.MccMnc) {
		return nil, false
	}
	return o.MccMnc, true
}

// HasMccMnc returns a boolean if a field has been set.
func (o *LogsReport) HasMccMnc() bool {
	if o != nil && !IsNil(o.MccMnc) {
		return true
	}

	return false
}

// SetMccMnc gets a reference to the given string and assigns it to the MccMnc field.
func (o *LogsReport) SetMccMnc(v string) {
	o.MccMnc = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *LogsReport) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *LogsReport) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *LogsReport) SetPrice(v Price) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogsReport) GetStatus() SingleMessageStatus {
	if o == nil || IsNil(o.Status) {
		var ret SingleMessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetStatusOk() (*SingleMessageStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogsReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SingleMessageStatus and assigns it to the Status field.
func (o *LogsReport) SetStatus(v SingleMessageStatus) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LogsReport) GetError() VoiceError {
	if o == nil || IsNil(o.Error) {
		var ret VoiceError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsReport) GetErrorOk() (*VoiceError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LogsReport) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given VoiceError and assigns it to the Error field.
func (o *LogsReport) SetError(v VoiceError) {
	o.Error = &v
}

func (o LogsReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.DoneAt) {
		toSerialize["doneAt"] = o.DoneAt
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.MccMnc) {
		toSerialize["mccMnc"] = o.MccMnc
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

type NullableLogsReport struct {
	value *LogsReport
	isSet bool
}

func (v NullableLogsReport) Get() *LogsReport {
	return v.value
}

func (v *NullableLogsReport) Set(val *LogsReport) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsReport) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsReport(val *LogsReport) *NullableLogsReport {
	return &NullableLogsReport{value: val, isSet: true}
}

func (v NullableLogsReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
