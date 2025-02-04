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

// checks if the IvrMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IvrMessage{}

// IvrMessage Array of IVR messages to be sent, one object per every message.
type IvrMessage struct {
	// Scenario key.
	ScenarioId string
	// Numeric sender ID length should be between 3 and 14 characters.
	From *string
	// Array of message destination addresses. Maximum number of destination addresses is 20k.
	Destinations []Destination
	// The URL on your callback server on which the Delivery report will be sent.
	NotifyUrl *string
	// Preferred Delivery report content type. Can be `application/json` or `application/xml`.
	NotifyContentType *string
	// Specifies the version of the report model to be sent. Can be `1` ([deprecated version 1](#programmable-communications/voice/receive-voice-delivery-reports-deprecated)) or `2` ([current version 2](#programmable-communications/voice/receive-voice-delivery-reports)). The default is version 2.
	NotifyContentVersion *int32
	// Additional client's data that will be sent on the notifyUrl. The maximum value is 200 characters.
	CallbackData *string
	// The message validity period shown in minutes. When the period expires, it will not be allowed for the message to be sent. A validity period longer than 48h is not supported (in this case, it will be automatically set to 48h).
	ValidityPeriod *int32
	// Used for scheduled Voice messages (message not to be sent immediately, but at scheduled time).
	SendAt *Time
	Retry  *Retry
	// Ringing duration, unless there are no operator limitations. Default value is `45`. Note: There are no limitations on the Voice platform regarding this value, however, most of the operators have their own ring timeout limitations and it is advisable to keep the ringTimeout value up to `45` seconds.
	RingTimeout  *int32
	SendingSpeed *SendingSpeed
	// The parameters that should be passed to the scenario on execution.
	Parameters *map[string]string
	// Indicating period of time in seconds between end user answering the call and message starting to be played. Minimal value is `0` and maximum value is `10` seconds. Default value is `0`.
	Pause *int32
	// [Early access: Contact your account manager to enable the usage] Record the call and expose it to client as URL inside the delivery report. Can be `true` or `false`.
	Record             *bool
	DeliveryTimeWindow *DeliveryTimeWindow
	// Maximum possible duration of the call to be set, shown in seconds.
	CallTimeout *int32
}

type _IvrMessage IvrMessage

// NewIvrMessage instantiates a new IvrMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewIvrMessage(scenarioId string, destinations []Destination) *IvrMessage {
	this := IvrMessage{}
	this.ScenarioId = scenarioId
	this.Destinations = destinations
	return &this
}

// NewIvrMessageWithDefaults instantiates a new IvrMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIvrMessageWithDefaults() *IvrMessage {
	this := IvrMessage{}

	return &this
}

// GetScenarioId returns the ScenarioId field value
func (o *IvrMessage) GetScenarioId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScenarioId
}

// GetScenarioIdOk returns a tuple with the ScenarioId field value
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetScenarioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScenarioId, true
}

// SetScenarioId sets field value
func (o *IvrMessage) SetScenarioId(v string) {
	o.ScenarioId = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *IvrMessage) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *IvrMessage) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *IvrMessage) SetFrom(v string) {
	o.From = &v
}

// GetDestinations returns the Destinations field value
func (o *IvrMessage) GetDestinations() []Destination {
	if o == nil {
		var ret []Destination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetDestinationsOk() ([]Destination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *IvrMessage) SetDestinations(v []Destination) {
	o.Destinations = v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *IvrMessage) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetNotifyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *IvrMessage) HasNotifyUrl() bool {
	if o != nil && !IsNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *IvrMessage) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetNotifyContentType returns the NotifyContentType field value if set, zero value otherwise.
func (o *IvrMessage) GetNotifyContentType() string {
	if o == nil || IsNil(o.NotifyContentType) {
		var ret string
		return ret
	}
	return *o.NotifyContentType
}

// GetNotifyContentTypeOk returns a tuple with the NotifyContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetNotifyContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyContentType) {
		return nil, false
	}
	return o.NotifyContentType, true
}

// HasNotifyContentType returns a boolean if a field has been set.
func (o *IvrMessage) HasNotifyContentType() bool {
	if o != nil && !IsNil(o.NotifyContentType) {
		return true
	}

	return false
}

// SetNotifyContentType gets a reference to the given string and assigns it to the NotifyContentType field.
func (o *IvrMessage) SetNotifyContentType(v string) {
	o.NotifyContentType = &v
}

// GetNotifyContentVersion returns the NotifyContentVersion field value if set, zero value otherwise.
func (o *IvrMessage) GetNotifyContentVersion() int32 {
	if o == nil || IsNil(o.NotifyContentVersion) {
		var ret int32
		return ret
	}
	return *o.NotifyContentVersion
}

// GetNotifyContentVersionOk returns a tuple with the NotifyContentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetNotifyContentVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifyContentVersion) {
		return nil, false
	}
	return o.NotifyContentVersion, true
}

// HasNotifyContentVersion returns a boolean if a field has been set.
func (o *IvrMessage) HasNotifyContentVersion() bool {
	if o != nil && !IsNil(o.NotifyContentVersion) {
		return true
	}

	return false
}

// SetNotifyContentVersion gets a reference to the given int32 and assigns it to the NotifyContentVersion field.
func (o *IvrMessage) SetNotifyContentVersion(v int32) {
	o.NotifyContentVersion = &v
}

// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *IvrMessage) GetCallbackData() string {
	if o == nil || IsNil(o.CallbackData) {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetCallbackDataOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackData) {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *IvrMessage) HasCallbackData() bool {
	if o != nil && !IsNil(o.CallbackData) {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *IvrMessage) SetCallbackData(v string) {
	o.CallbackData = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *IvrMessage) GetValidityPeriod() int32 {
	if o == nil || IsNil(o.ValidityPeriod) {
		var ret int32
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetValidityPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *IvrMessage) HasValidityPeriod() bool {
	if o != nil && !IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given int32 and assigns it to the ValidityPeriod field.
func (o *IvrMessage) SetValidityPeriod(v int32) {
	o.ValidityPeriod = &v
}

// GetSendAt returns the SendAt field value if set, zero value otherwise.
func (o *IvrMessage) GetSendAt() Time {
	if o == nil || IsNil(o.SendAt) {
		var ret Time
		return ret
	}
	return *o.SendAt
}

// GetSendAtOk returns a tuple with the SendAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetSendAtOk() (*Time, bool) {
	if o == nil || IsNil(o.SendAt) {
		return nil, false
	}
	return o.SendAt, true
}

// HasSendAt returns a boolean if a field has been set.
func (o *IvrMessage) HasSendAt() bool {
	if o != nil && !IsNil(o.SendAt) {
		return true
	}

	return false
}

// SetSendAt gets a reference to the given Time and assigns it to the SendAt field.
func (o *IvrMessage) SetSendAt(v Time) {
	o.SendAt = &v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *IvrMessage) GetRetry() Retry {
	if o == nil || IsNil(o.Retry) {
		var ret Retry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetRetryOk() (*Retry, bool) {
	if o == nil || IsNil(o.Retry) {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *IvrMessage) HasRetry() bool {
	if o != nil && !IsNil(o.Retry) {
		return true
	}

	return false
}

// SetRetry gets a reference to the given Retry and assigns it to the Retry field.
func (o *IvrMessage) SetRetry(v Retry) {
	o.Retry = &v
}

// GetRingTimeout returns the RingTimeout field value if set, zero value otherwise.
func (o *IvrMessage) GetRingTimeout() int32 {
	if o == nil || IsNil(o.RingTimeout) {
		var ret int32
		return ret
	}
	return *o.RingTimeout
}

// GetRingTimeoutOk returns a tuple with the RingTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetRingTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.RingTimeout) {
		return nil, false
	}
	return o.RingTimeout, true
}

// HasRingTimeout returns a boolean if a field has been set.
func (o *IvrMessage) HasRingTimeout() bool {
	if o != nil && !IsNil(o.RingTimeout) {
		return true
	}

	return false
}

// SetRingTimeout gets a reference to the given int32 and assigns it to the RingTimeout field.
func (o *IvrMessage) SetRingTimeout(v int32) {
	o.RingTimeout = &v
}

// GetSendingSpeed returns the SendingSpeed field value if set, zero value otherwise.
func (o *IvrMessage) GetSendingSpeed() SendingSpeed {
	if o == nil || IsNil(o.SendingSpeed) {
		var ret SendingSpeed
		return ret
	}
	return *o.SendingSpeed
}

// GetSendingSpeedOk returns a tuple with the SendingSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetSendingSpeedOk() (*SendingSpeed, bool) {
	if o == nil || IsNil(o.SendingSpeed) {
		return nil, false
	}
	return o.SendingSpeed, true
}

// HasSendingSpeed returns a boolean if a field has been set.
func (o *IvrMessage) HasSendingSpeed() bool {
	if o != nil && !IsNil(o.SendingSpeed) {
		return true
	}

	return false
}

// SetSendingSpeed gets a reference to the given SendingSpeed and assigns it to the SendingSpeed field.
func (o *IvrMessage) SetSendingSpeed(v SendingSpeed) {
	o.SendingSpeed = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *IvrMessage) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IvrMessage) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *IvrMessage) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetPause returns the Pause field value if set, zero value otherwise.
func (o *IvrMessage) GetPause() int32 {
	if o == nil || IsNil(o.Pause) {
		var ret int32
		return ret
	}
	return *o.Pause
}

// GetPauseOk returns a tuple with the Pause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetPauseOk() (*int32, bool) {
	if o == nil || IsNil(o.Pause) {
		return nil, false
	}
	return o.Pause, true
}

// HasPause returns a boolean if a field has been set.
func (o *IvrMessage) HasPause() bool {
	if o != nil && !IsNil(o.Pause) {
		return true
	}

	return false
}

// SetPause gets a reference to the given int32 and assigns it to the Pause field.
func (o *IvrMessage) SetPause(v int32) {
	o.Pause = &v
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *IvrMessage) GetRecord() bool {
	if o == nil || IsNil(o.Record) {
		var ret bool
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetRecordOk() (*bool, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *IvrMessage) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given bool and assigns it to the Record field.
func (o *IvrMessage) SetRecord(v bool) {
	o.Record = &v
}

// GetDeliveryTimeWindow returns the DeliveryTimeWindow field value if set, zero value otherwise.
func (o *IvrMessage) GetDeliveryTimeWindow() DeliveryTimeWindow {
	if o == nil || IsNil(o.DeliveryTimeWindow) {
		var ret DeliveryTimeWindow
		return ret
	}
	return *o.DeliveryTimeWindow
}

// GetDeliveryTimeWindowOk returns a tuple with the DeliveryTimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetDeliveryTimeWindowOk() (*DeliveryTimeWindow, bool) {
	if o == nil || IsNil(o.DeliveryTimeWindow) {
		return nil, false
	}
	return o.DeliveryTimeWindow, true
}

// HasDeliveryTimeWindow returns a boolean if a field has been set.
func (o *IvrMessage) HasDeliveryTimeWindow() bool {
	if o != nil && !IsNil(o.DeliveryTimeWindow) {
		return true
	}

	return false
}

// SetDeliveryTimeWindow gets a reference to the given DeliveryTimeWindow and assigns it to the DeliveryTimeWindow field.
func (o *IvrMessage) SetDeliveryTimeWindow(v DeliveryTimeWindow) {
	o.DeliveryTimeWindow = &v
}

// GetCallTimeout returns the CallTimeout field value if set, zero value otherwise.
func (o *IvrMessage) GetCallTimeout() int32 {
	if o == nil || IsNil(o.CallTimeout) {
		var ret int32
		return ret
	}
	return *o.CallTimeout
}

// GetCallTimeoutOk returns a tuple with the CallTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IvrMessage) GetCallTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.CallTimeout) {
		return nil, false
	}
	return o.CallTimeout, true
}

// HasCallTimeout returns a boolean if a field has been set.
func (o *IvrMessage) HasCallTimeout() bool {
	if o != nil && !IsNil(o.CallTimeout) {
		return true
	}

	return false
}

// SetCallTimeout gets a reference to the given int32 and assigns it to the CallTimeout field.
func (o *IvrMessage) SetCallTimeout(v int32) {
	o.CallTimeout = &v
}

func (o IvrMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IvrMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scenarioId"] = o.ScenarioId
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	toSerialize["destinations"] = o.Destinations
	if !IsNil(o.NotifyUrl) {
		toSerialize["notifyUrl"] = o.NotifyUrl
	}
	if !IsNil(o.NotifyContentType) {
		toSerialize["notifyContentType"] = o.NotifyContentType
	}
	if !IsNil(o.NotifyContentVersion) {
		toSerialize["notifyContentVersion"] = o.NotifyContentVersion
	}
	if !IsNil(o.CallbackData) {
		toSerialize["callbackData"] = o.CallbackData
	}
	if !IsNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}
	if !IsNil(o.SendAt) {
		toSerialize["sendAt"] = o.SendAt
	}
	if !IsNil(o.Retry) {
		toSerialize["retry"] = o.Retry
	}
	if !IsNil(o.RingTimeout) {
		toSerialize["ringTimeout"] = o.RingTimeout
	}
	if !IsNil(o.SendingSpeed) {
		toSerialize["sendingSpeed"] = o.SendingSpeed
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Pause) {
		toSerialize["pause"] = o.Pause
	}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	if !IsNil(o.DeliveryTimeWindow) {
		toSerialize["deliveryTimeWindow"] = o.DeliveryTimeWindow
	}
	if !IsNil(o.CallTimeout) {
		toSerialize["callTimeout"] = o.CallTimeout
	}
	return toSerialize, nil
}

type NullableIvrMessage struct {
	value *IvrMessage
	isSet bool
}

func (v NullableIvrMessage) Get() *IvrMessage {
	return v.value
}

func (v *NullableIvrMessage) Set(val *IvrMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableIvrMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableIvrMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIvrMessage(val *IvrMessage) *NullableIvrMessage {
	return &NullableIvrMessage{value: val, isSet: true}
}

func (v NullableIvrMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIvrMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
