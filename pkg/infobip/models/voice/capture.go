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

// checks if the Capture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Capture{}

// Capture Capture action performs speech recognition of a specified duration. The recognized text can be used in other actions of the scenario.
type Capture struct {
	// Variable name. If speech recognition matches one of the specified keyPhrases, a variable with this name will be set to match the keyphrase. Otherwise, this variable will be an empty string. Together with this variable, an implicit variable that contains the full text of the captured speech will be created. The name of this variable is constructed by adding _Full suffix to the variable name.
	Capture string
	// Number of seconds used for capturing speech or digits failover. Minimum value is 1 and maximum value is 30. Can be overriden with speechOptions.maxSilence.
	Timeout       int32
	SpeechOptions SpeechOptions
	DtmfOptions   *DtmfOptions
	SendToReports *SendToReports
	// User-defined ID of an action that can be used with go-to action.
	ActionId *int32
}

type _Capture Capture

// NewCapture instantiates a new Capture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCapture(capture string, timeout int32, speechOptions SpeechOptions) *Capture {
	this := Capture{}
	this.Capture = capture
	this.Timeout = timeout
	this.SpeechOptions = speechOptions
	return &this
}

// NewCaptureWithDefaults instantiates a new Capture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureWithDefaults() *Capture {
	this := Capture{}

	return &this
}

// GetCapture returns the Capture field value
func (o *Capture) GetCapture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capture
}

// GetCaptureOk returns a tuple with the Capture field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCaptureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capture, true
}

// SetCapture sets field value
func (o *Capture) SetCapture(v string) {
	o.Capture = v
}

// GetTimeout returns the Timeout field value
func (o *Capture) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *Capture) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *Capture) SetTimeout(v int32) {
	o.Timeout = v
}

// GetSpeechOptions returns the SpeechOptions field value
func (o *Capture) GetSpeechOptions() SpeechOptions {
	if o == nil {
		var ret SpeechOptions
		return ret
	}

	return o.SpeechOptions
}

// GetSpeechOptionsOk returns a tuple with the SpeechOptions field value
// and a boolean to check if the value has been set.
func (o *Capture) GetSpeechOptionsOk() (*SpeechOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpeechOptions, true
}

// SetSpeechOptions sets field value
func (o *Capture) SetSpeechOptions(v SpeechOptions) {
	o.SpeechOptions = v
}

// GetDtmfOptions returns the DtmfOptions field value if set, zero value otherwise.
func (o *Capture) GetDtmfOptions() DtmfOptions {
	if o == nil || IsNil(o.DtmfOptions) {
		var ret DtmfOptions
		return ret
	}
	return *o.DtmfOptions
}

// GetDtmfOptionsOk returns a tuple with the DtmfOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capture) GetDtmfOptionsOk() (*DtmfOptions, bool) {
	if o == nil || IsNil(o.DtmfOptions) {
		return nil, false
	}
	return o.DtmfOptions, true
}

// HasDtmfOptions returns a boolean if a field has been set.
func (o *Capture) HasDtmfOptions() bool {
	if o != nil && !IsNil(o.DtmfOptions) {
		return true
	}

	return false
}

// SetDtmfOptions gets a reference to the given DtmfOptions and assigns it to the DtmfOptions field.
func (o *Capture) SetDtmfOptions(v DtmfOptions) {
	o.DtmfOptions = &v
}

// GetSendToReports returns the SendToReports field value if set, zero value otherwise.
func (o *Capture) GetSendToReports() SendToReports {
	if o == nil || IsNil(o.SendToReports) {
		var ret SendToReports
		return ret
	}
	return *o.SendToReports
}

// GetSendToReportsOk returns a tuple with the SendToReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capture) GetSendToReportsOk() (*SendToReports, bool) {
	if o == nil || IsNil(o.SendToReports) {
		return nil, false
	}
	return o.SendToReports, true
}

// HasSendToReports returns a boolean if a field has been set.
func (o *Capture) HasSendToReports() bool {
	if o != nil && !IsNil(o.SendToReports) {
		return true
	}

	return false
}

// SetSendToReports gets a reference to the given SendToReports and assigns it to the SendToReports field.
func (o *Capture) SetSendToReports(v SendToReports) {
	o.SendToReports = &v
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *Capture) GetActionId() int32 {
	if o == nil || IsNil(o.ActionId) {
		var ret int32
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capture) GetActionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *Capture) HasActionId() bool {
	if o != nil && !IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given int32 and assigns it to the ActionId field.
func (o *Capture) SetActionId(v int32) {
	o.ActionId = &v
}

func (o Capture) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Capture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capture"] = o.Capture
	toSerialize["timeout"] = o.Timeout
	toSerialize["speechOptions"] = o.SpeechOptions
	if !IsNil(o.DtmfOptions) {
		toSerialize["dtmfOptions"] = o.DtmfOptions
	}
	if !IsNil(o.SendToReports) {
		toSerialize["sendToReports"] = o.SendToReports
	}
	if !IsNil(o.ActionId) {
		toSerialize["actionId"] = o.ActionId
	}
	return toSerialize, nil
}

type NullableCapture struct {
	value *Capture
	isSet bool
}

func (v NullableCapture) Get() *Capture {
	return v.value
}

func (v *NullableCapture) Set(val *Capture) {
	v.value = val
	v.isSet = true
}

func (v NullableCapture) IsSet() bool {
	return v.isSet
}

func (v *NullableCapture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapture(val *Capture) *NullableCapture {
	return &NullableCapture{value: val, isSet: true}
}

func (v NullableCapture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
