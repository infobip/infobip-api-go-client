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

// checks if the TextMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextMessageContent{}

// TextMessageContent struct for TextMessageContent
type TextMessageContent struct {
	// Content of the message being sent.
	Text            string
	Transliteration *TransliterationCode
	Language        *LanguageV3
}

type _TextMessageContent TextMessageContent

// NewTextMessageContent instantiates a new TextMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextMessageContent(text string) *TextMessageContent {
	this := TextMessageContent{}
	this.Text = text
	return &this
}

// NewTextMessageContentWithDefaults instantiates a new TextMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextMessageContentWithDefaults() *TextMessageContent {
	this := TextMessageContent{}
	return &this
}

// GetText returns the Text field value
func (o *TextMessageContent) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TextMessageContent) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TextMessageContent) SetText(v string) {
	o.Text = v
}

// GetTransliteration returns the Transliteration field value if set, zero value otherwise.
func (o *TextMessageContent) GetTransliteration() TransliterationCode {
	if o == nil || IsNil(o.Transliteration) {
		var ret TransliterationCode
		return ret
	}
	return *o.Transliteration
}

// GetTransliterationOk returns a tuple with the Transliteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextMessageContent) GetTransliterationOk() (*TransliterationCode, bool) {
	if o == nil || IsNil(o.Transliteration) {
		return nil, false
	}
	return o.Transliteration, true
}

// HasTransliteration returns a boolean if a field has been set.
func (o *TextMessageContent) HasTransliteration() bool {
	if o != nil && !IsNil(o.Transliteration) {
		return true
	}

	return false
}

// SetTransliteration gets a reference to the given TransliterationCode and assigns it to the Transliteration field.
func (o *TextMessageContent) SetTransliteration(v TransliterationCode) {
	o.Transliteration = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TextMessageContent) GetLanguage() LanguageV3 {
	if o == nil || IsNil(o.Language) {
		var ret LanguageV3
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextMessageContent) GetLanguageOk() (*LanguageV3, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TextMessageContent) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given LanguageV3 and assigns it to the Language field.
func (o *TextMessageContent) SetLanguage(v LanguageV3) {
	o.Language = &v
}

func (o TextMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.Transliteration) {
		toSerialize["transliteration"] = o.Transliteration
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableTextMessageContent struct {
	value *TextMessageContent
	isSet bool
}

func (v NullableTextMessageContent) Get() *TextMessageContent {
	return v.value
}

func (v *NullableTextMessageContent) Set(val *TextMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTextMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTextMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextMessageContent(val *TextMessageContent) *NullableTextMessageContent {
	return &NullableTextMessageContent{value: val, isSet: true}
}

func (v NullableTextMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}