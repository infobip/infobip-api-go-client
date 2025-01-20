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

// checks if the MessageDeliveryReporting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDeliveryReporting{}

// MessageDeliveryReporting Provides options for configuring the delivery report behavior.
type MessageDeliveryReporting struct {
	// The URL on your call back server where a delivery report will be sent. If your URL becomes unavailable then the [retry cycle](https://www.infobip.com/docs/sms/api#notify-url) uses the following formula: `1min + (1min * retryNumber * retryNumber)`.
	Url *string
	// The real-time intermediate delivery report containing GSM error codes, messages status, pricing, network and country codes, etc., which will be sent on your callback server. Defaults to `false`.
	IntermediateReport *bool
}

// NewMessageDeliveryReporting instantiates a new MessageDeliveryReporting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMessageDeliveryReporting() *MessageDeliveryReporting {
	this := MessageDeliveryReporting{}
	return &this
}

// NewMessageDeliveryReportingWithDefaults instantiates a new MessageDeliveryReporting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDeliveryReportingWithDefaults() *MessageDeliveryReporting {
	this := MessageDeliveryReporting{}

	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MessageDeliveryReporting) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeliveryReporting) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MessageDeliveryReporting) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MessageDeliveryReporting) SetUrl(v string) {
	o.Url = &v
}

// GetIntermediateReport returns the IntermediateReport field value if set, zero value otherwise.
func (o *MessageDeliveryReporting) GetIntermediateReport() bool {
	if o == nil || IsNil(o.IntermediateReport) {
		var ret bool
		return ret
	}
	return *o.IntermediateReport
}

// GetIntermediateReportOk returns a tuple with the IntermediateReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageDeliveryReporting) GetIntermediateReportOk() (*bool, bool) {
	if o == nil || IsNil(o.IntermediateReport) {
		return nil, false
	}
	return o.IntermediateReport, true
}

// HasIntermediateReport returns a boolean if a field has been set.
func (o *MessageDeliveryReporting) HasIntermediateReport() bool {
	if o != nil && !IsNil(o.IntermediateReport) {
		return true
	}

	return false
}

// SetIntermediateReport gets a reference to the given bool and assigns it to the IntermediateReport field.
func (o *MessageDeliveryReporting) SetIntermediateReport(v bool) {
	o.IntermediateReport = &v
}

func (o MessageDeliveryReporting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDeliveryReporting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.IntermediateReport) {
		toSerialize["intermediateReport"] = o.IntermediateReport
	}
	return toSerialize, nil
}

type NullableMessageDeliveryReporting struct {
	value *MessageDeliveryReporting
	isSet bool
}

func (v NullableMessageDeliveryReporting) Get() *MessageDeliveryReporting {
	return v.value
}

func (v *NullableMessageDeliveryReporting) Set(val *MessageDeliveryReporting) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDeliveryReporting) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDeliveryReporting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDeliveryReporting(val *MessageDeliveryReporting) *NullableMessageDeliveryReporting {
	return &NullableMessageDeliveryReporting{value: val, isSet: true}
}

func (v NullableMessageDeliveryReporting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDeliveryReporting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
