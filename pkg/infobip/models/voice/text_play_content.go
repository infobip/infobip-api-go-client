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

// checks if the TextPlayContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextPlayContent{}

// TextPlayContent struct for TextPlayContent
type TextPlayContent struct {
	Type        PlayContentType
	Text        string
	Language    Language
	SpeechRate  *float64
	Preferences *VoicePreferences
}

type _TextPlayContent TextPlayContent

// NewTextPlayContent instantiates a new TextPlayContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextPlayContent(text string, language Language) *TextPlayContent {
	this := TextPlayContent{}
	this.Type = "TEXT"
	this.Text = text
	this.Language = language
	return &this
}

// NewTextPlayContentWithDefaults instantiates a new TextPlayContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextPlayContentWithDefaults() *TextPlayContent {
	this := TextPlayContent{}
	this.Type = "TEXT"
	return &this
}

// GetText returns the Text field value
func (o *TextPlayContent) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TextPlayContent) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TextPlayContent) SetText(v string) {
	o.Text = v
}

// GetLanguage returns the Language field value
func (o *TextPlayContent) GetLanguage() Language {
	if o == nil {
		var ret Language
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *TextPlayContent) GetLanguageOk() (*Language, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *TextPlayContent) SetLanguage(v Language) {
	o.Language = v
}

// GetSpeechRate returns the SpeechRate field value if set, zero value otherwise.
func (o *TextPlayContent) GetSpeechRate() float64 {
	if o == nil || IsNil(o.SpeechRate) {
		var ret float64
		return ret
	}
	return *o.SpeechRate
}

// GetSpeechRateOk returns a tuple with the SpeechRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPlayContent) GetSpeechRateOk() (*float64, bool) {
	if o == nil || IsNil(o.SpeechRate) {
		return nil, false
	}
	return o.SpeechRate, true
}

// HasSpeechRate returns a boolean if a field has been set.
func (o *TextPlayContent) HasSpeechRate() bool {
	if o != nil && !IsNil(o.SpeechRate) {
		return true
	}

	return false
}

// SetSpeechRate gets a reference to the given float64 and assigns it to the SpeechRate field.
func (o *TextPlayContent) SetSpeechRate(v float64) {
	o.SpeechRate = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *TextPlayContent) GetPreferences() VoicePreferences {
	if o == nil || IsNil(o.Preferences) {
		var ret VoicePreferences
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextPlayContent) GetPreferencesOk() (*VoicePreferences, bool) {
	if o == nil || IsNil(o.Preferences) {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *TextPlayContent) HasPreferences() bool {
	if o != nil && !IsNil(o.Preferences) {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given VoicePreferences and assigns it to the Preferences field.
func (o *TextPlayContent) SetPreferences(v VoicePreferences) {
	o.Preferences = &v
}

func (o TextPlayContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextPlayContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["text"] = o.Text
	toSerialize["language"] = o.Language
	if !IsNil(o.SpeechRate) {
		toSerialize["speechRate"] = o.SpeechRate
	}
	if !IsNil(o.Preferences) {
		toSerialize["preferences"] = o.Preferences
	}
	return toSerialize, nil
}

type NullableTextPlayContent struct {
	value *TextPlayContent
	isSet bool
}

func (v NullableTextPlayContent) Get() *TextPlayContent {
	return v.value
}

func (v *NullableTextPlayContent) Set(val *TextPlayContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTextPlayContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTextPlayContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextPlayContent(val *TextPlayContent) *NullableTextPlayContent {
	return &NullableTextPlayContent{value: val, isSet: true}
}

func (v NullableTextPlayContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextPlayContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
