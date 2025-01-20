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

// checks if the DialogSayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DialogSayRequest{}

// DialogSayRequest struct for DialogSayRequest
type DialogSayRequest struct {
	// Text to read.
	Text     string
	Language Language
	// Speech rate. Must be within `[0.5 - 2.0]` range, default value is `1`.
	SpeechRate *float64
	// Number of times to read the text.
	LoopCount   *int32
	Preferences *VoicePreferences
}

type _DialogSayRequest DialogSayRequest

// NewDialogSayRequest instantiates a new DialogSayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDialogSayRequest(text string, language Language) *DialogSayRequest {
	this := DialogSayRequest{}
	this.Text = text
	this.Language = language
	return &this
}

// NewDialogSayRequestWithDefaults instantiates a new DialogSayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDialogSayRequestWithDefaults() *DialogSayRequest {
	this := DialogSayRequest{}

	return &this
}

// GetText returns the Text field value
func (o *DialogSayRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *DialogSayRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *DialogSayRequest) SetText(v string) {
	o.Text = v
}

// GetLanguage returns the Language field value
func (o *DialogSayRequest) GetLanguage() Language {
	if o == nil {
		var ret Language
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *DialogSayRequest) GetLanguageOk() (*Language, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *DialogSayRequest) SetLanguage(v Language) {
	o.Language = v
}

// GetSpeechRate returns the SpeechRate field value if set, zero value otherwise.
func (o *DialogSayRequest) GetSpeechRate() float64 {
	if o == nil || IsNil(o.SpeechRate) {
		var ret float64
		return ret
	}
	return *o.SpeechRate
}

// GetSpeechRateOk returns a tuple with the SpeechRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogSayRequest) GetSpeechRateOk() (*float64, bool) {
	if o == nil || IsNil(o.SpeechRate) {
		return nil, false
	}
	return o.SpeechRate, true
}

// HasSpeechRate returns a boolean if a field has been set.
func (o *DialogSayRequest) HasSpeechRate() bool {
	if o != nil && !IsNil(o.SpeechRate) {
		return true
	}

	return false
}

// SetSpeechRate gets a reference to the given float64 and assigns it to the SpeechRate field.
func (o *DialogSayRequest) SetSpeechRate(v float64) {
	o.SpeechRate = &v
}

// GetLoopCount returns the LoopCount field value if set, zero value otherwise.
func (o *DialogSayRequest) GetLoopCount() int32 {
	if o == nil || IsNil(o.LoopCount) {
		var ret int32
		return ret
	}
	return *o.LoopCount
}

// GetLoopCountOk returns a tuple with the LoopCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogSayRequest) GetLoopCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LoopCount) {
		return nil, false
	}
	return o.LoopCount, true
}

// HasLoopCount returns a boolean if a field has been set.
func (o *DialogSayRequest) HasLoopCount() bool {
	if o != nil && !IsNil(o.LoopCount) {
		return true
	}

	return false
}

// SetLoopCount gets a reference to the given int32 and assigns it to the LoopCount field.
func (o *DialogSayRequest) SetLoopCount(v int32) {
	o.LoopCount = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *DialogSayRequest) GetPreferences() VoicePreferences {
	if o == nil || IsNil(o.Preferences) {
		var ret VoicePreferences
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogSayRequest) GetPreferencesOk() (*VoicePreferences, bool) {
	if o == nil || IsNil(o.Preferences) {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *DialogSayRequest) HasPreferences() bool {
	if o != nil && !IsNil(o.Preferences) {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given VoicePreferences and assigns it to the Preferences field.
func (o *DialogSayRequest) SetPreferences(v VoicePreferences) {
	o.Preferences = &v
}

func (o DialogSayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DialogSayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["language"] = o.Language
	if !IsNil(o.SpeechRate) {
		toSerialize["speechRate"] = o.SpeechRate
	}
	if !IsNil(o.LoopCount) {
		toSerialize["loopCount"] = o.LoopCount
	}
	if !IsNil(o.Preferences) {
		toSerialize["preferences"] = o.Preferences
	}
	return toSerialize, nil
}

type NullableDialogSayRequest struct {
	value *DialogSayRequest
	isSet bool
}

func (v NullableDialogSayRequest) Get() *DialogSayRequest {
	return v.value
}

func (v *NullableDialogSayRequest) Set(val *DialogSayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDialogSayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDialogSayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDialogSayRequest(val *DialogSayRequest) *NullableDialogSayRequest {
	return &NullableDialogSayRequest{value: val, isSet: true}
}

func (v NullableDialogSayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDialogSayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
