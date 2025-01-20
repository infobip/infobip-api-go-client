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

// checks if the ApiReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReport{}

// ApiReport struct for ApiReport
type ApiReport struct {
	// The Application ID sent in the email request.
	ApplicationId *string
	// The Entity ID sent in the email request.
	EntityId *string
	// The ID that uniquely identifies bulks of request.
	BulkId *string
	// The ID that uniquely identifies the sent email request.
	MessageId *string
	// The recipient email address.
	To *string
	// Tells when the email was initiated. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	SentAt *Time
	// Tells when the email request was processed by Infobip.
	DoneAt *Time
	// Email request count.
	MessageCount *int32
	Price        *Price
	Status       *Status
	Error        *Error
}

// NewApiReport instantiates a new ApiReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewApiReport() *ApiReport {
	this := ApiReport{}
	return &this
}

// NewApiReportWithDefaults instantiates a new ApiReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportWithDefaults() *ApiReport {
	this := ApiReport{}

	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiReport) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiReport) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiReport) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ApiReport) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ApiReport) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ApiReport) SetEntityId(v string) {
	o.EntityId = &v
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *ApiReport) GetBulkId() string {
	if o == nil || IsNil(o.BulkId) {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetBulkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulkId) {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *ApiReport) HasBulkId() bool {
	if o != nil && !IsNil(o.BulkId) {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *ApiReport) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ApiReport) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ApiReport) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ApiReport) SetMessageId(v string) {
	o.MessageId = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ApiReport) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ApiReport) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *ApiReport) SetTo(v string) {
	o.To = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *ApiReport) GetSentAt() Time {
	if o == nil || IsNil(o.SentAt) {
		var ret Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetSentAtOk() (*Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *ApiReport) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given Time and assigns it to the SentAt field.
func (o *ApiReport) SetSentAt(v Time) {
	o.SentAt = &v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise.
func (o *ApiReport) GetDoneAt() Time {
	if o == nil || IsNil(o.DoneAt) {
		var ret Time
		return ret
	}
	return *o.DoneAt
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetDoneAtOk() (*Time, bool) {
	if o == nil || IsNil(o.DoneAt) {
		return nil, false
	}
	return o.DoneAt, true
}

// HasDoneAt returns a boolean if a field has been set.
func (o *ApiReport) HasDoneAt() bool {
	if o != nil && !IsNil(o.DoneAt) {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given Time and assigns it to the DoneAt field.
func (o *ApiReport) SetDoneAt(v Time) {
	o.DoneAt = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *ApiReport) GetMessageCount() int32 {
	if o == nil || IsNil(o.MessageCount) {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageCount) {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *ApiReport) HasMessageCount() bool {
	if o != nil && !IsNil(o.MessageCount) {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *ApiReport) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ApiReport) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ApiReport) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *ApiReport) SetPrice(v Price) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiReport) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *ApiReport) SetStatus(v Status) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApiReport) GetError() Error {
	if o == nil || IsNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReport) GetErrorOk() (*Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiReport) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *ApiReport) SetError(v Error) {
	o.Error = &v
}

func (o ApiReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.BulkId) {
		toSerialize["bulkId"] = o.BulkId
	}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
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
	return toSerialize, nil
}

type NullableApiReport struct {
	value *ApiReport
	isSet bool
}

func (v NullableApiReport) Get() *ApiReport {
	return v.value
}

func (v *NullableApiReport) Set(val *ApiReport) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReport) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReport(val *ApiReport) *NullableApiReport {
	return &NullableApiReport{value: val, isSet: true}
}

func (v NullableApiReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
