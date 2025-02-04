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

// checks if the MultiMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiMessage{}

// MultiMessage Array of messages to be sent, one per every message.
type MultiMessage struct {
	// An audio file can be delivered as a voice message to the recipients. An audio file must be uploaded online, so that the existing URL can be available for file download. Size of the audio file must be below 4 MB. Supported formats of the provided file are aac, aiff, m4a, mp2, mp3, mp4 (audio only), ogg, wav and wma. Our platform needs to have permission to make GET and HEAD HTTP requests on the provided URL. Standard http ports (like 80, 8080, etc.) are advised.
	AudioFileUrl *string
	// Numeric sender ID in E.164 standard format (Example: 41793026727). This is caller ID that will be presented to the end user where applicable.
	From *string
	// If the message is in text format, the language in which the message is written must be defined for correct pronunciation. More about Text-to-speech functionality and supported TTS languages can be found [here](https://www.infobip.com/docs/voice-and-video/outbound-calls#text-to-speech-voice-over-broadcast). If not set, default language is `English [en]`. If voice is not set, then default voice for that specific language is used. In the case of English language, the voice is `[Joanna]`.
	Language *string
	// Text of the message that will be sent. Message text can be up to 1400 characters long and cannot contain only punctuation. SSML (_Speech Synthesis Markup Language_) is supported and can be used to fully customize pronunciation of the provided text.
	Text *string
	// Phone number of the recipient. Phone number must be written in E.164 standard format (Example: 41793026727). Maximum number of phone numbers listed is 20k.
	To    []string
	Voice *Voice
}

type _MultiMessage MultiMessage

// NewMultiMessage instantiates a new MultiMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMultiMessage(to []string) *MultiMessage {
	this := MultiMessage{}
	this.To = to
	return &this
}

// NewMultiMessageWithDefaults instantiates a new MultiMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiMessageWithDefaults() *MultiMessage {
	this := MultiMessage{}

	return &this
}

// GetAudioFileUrl returns the AudioFileUrl field value if set, zero value otherwise.
func (o *MultiMessage) GetAudioFileUrl() string {
	if o == nil || IsNil(o.AudioFileUrl) {
		var ret string
		return ret
	}
	return *o.AudioFileUrl
}

// GetAudioFileUrlOk returns a tuple with the AudioFileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiMessage) GetAudioFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AudioFileUrl) {
		return nil, false
	}
	return o.AudioFileUrl, true
}

// HasAudioFileUrl returns a boolean if a field has been set.
func (o *MultiMessage) HasAudioFileUrl() bool {
	if o != nil && !IsNil(o.AudioFileUrl) {
		return true
	}

	return false
}

// SetAudioFileUrl gets a reference to the given string and assigns it to the AudioFileUrl field.
func (o *MultiMessage) SetAudioFileUrl(v string) {
	o.AudioFileUrl = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MultiMessage) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiMessage) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MultiMessage) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *MultiMessage) SetFrom(v string) {
	o.From = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MultiMessage) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiMessage) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MultiMessage) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MultiMessage) SetLanguage(v string) {
	o.Language = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *MultiMessage) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiMessage) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MultiMessage) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *MultiMessage) SetText(v string) {
	o.Text = &v
}

// GetTo returns the To field value
func (o *MultiMessage) GetTo() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *MultiMessage) GetToOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *MultiMessage) SetTo(v []string) {
	o.To = v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *MultiMessage) GetVoice() Voice {
	if o == nil || IsNil(o.Voice) {
		var ret Voice
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiMessage) GetVoiceOk() (*Voice, bool) {
	if o == nil || IsNil(o.Voice) {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *MultiMessage) HasVoice() bool {
	if o != nil && !IsNil(o.Voice) {
		return true
	}

	return false
}

// SetVoice gets a reference to the given Voice and assigns it to the Voice field.
func (o *MultiMessage) SetVoice(v Voice) {
	o.Voice = &v
}

func (o MultiMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioFileUrl) {
		toSerialize["audioFileUrl"] = o.AudioFileUrl
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	toSerialize["to"] = o.To
	if !IsNil(o.Voice) {
		toSerialize["voice"] = o.Voice
	}
	return toSerialize, nil
}

type NullableMultiMessage struct {
	value *MultiMessage
	isSet bool
}

func (v NullableMultiMessage) Get() *MultiMessage {
	return v.value
}

func (v *NullableMultiMessage) Set(val *MultiMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiMessage(val *MultiMessage) *NullableMultiMessage {
	return &NullableMultiMessage{value: val, isSet: true}
}

func (v NullableMultiMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
