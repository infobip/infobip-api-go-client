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

// checks if the MessageRichLinkBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageRichLinkBody{}

// MessageRichLinkBody struct for MessageRichLinkBody
type MessageRichLinkBody struct {
	Type MessageBodyType
	// URL to be redirected to.
	RedirectUrl string
	// Text to be sent.
	Text string
	// URL of the rich link. It can be image or video. When you provide it with `isVideo=true` then it will be mapped as a video, otherwise - as an image.
	Url string
	// Indicates if provided `url` is a video.
	IsVideo *bool
	// URL of the thumbnail image. If you add this, then we use as thumbnail image for video you provided in `url`. Applicable only when `isVideo=true`.
	ThumbnailUrl *string
	// Button label text.
	ButtonText  *string
	CardOptions *MessageCardOptions
}

type _MessageRichLinkBody MessageRichLinkBody

// NewMessageRichLinkBody instantiates a new MessageRichLinkBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageRichLinkBody(redirectUrl string, text string, url string) *MessageRichLinkBody {
	this := MessageRichLinkBody{}
	this.Type = "RICH_LINK"
	this.RedirectUrl = redirectUrl
	this.Text = text
	this.Url = url
	var isVideo bool = false
	this.IsVideo = &isVideo
	return &this
}

// NewMessageRichLinkBodyWithDefaults instantiates a new MessageRichLinkBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageRichLinkBodyWithDefaults() *MessageRichLinkBody {
	this := MessageRichLinkBody{}
	this.Type = "RICH_LINK"
	var isVideo bool = false
	this.IsVideo = &isVideo
	return &this
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *MessageRichLinkBody) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *MessageRichLinkBody) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// GetText returns the Text field value
func (o *MessageRichLinkBody) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MessageRichLinkBody) SetText(v string) {
	o.Text = v
}

// GetUrl returns the Url field value
func (o *MessageRichLinkBody) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MessageRichLinkBody) SetUrl(v string) {
	o.Url = v
}

// GetIsVideo returns the IsVideo field value if set, zero value otherwise.
func (o *MessageRichLinkBody) GetIsVideo() bool {
	if o == nil || IsNil(o.IsVideo) {
		var ret bool
		return ret
	}
	return *o.IsVideo
}

// GetIsVideoOk returns a tuple with the IsVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetIsVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVideo) {
		return nil, false
	}
	return o.IsVideo, true
}

// HasIsVideo returns a boolean if a field has been set.
func (o *MessageRichLinkBody) HasIsVideo() bool {
	if o != nil && !IsNil(o.IsVideo) {
		return true
	}

	return false
}

// SetIsVideo gets a reference to the given bool and assigns it to the IsVideo field.
func (o *MessageRichLinkBody) SetIsVideo(v bool) {
	o.IsVideo = &v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise.
func (o *MessageRichLinkBody) GetThumbnailUrl() string {
	if o == nil || IsNil(o.ThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailUrl) {
		return nil, false
	}
	return o.ThumbnailUrl, true
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *MessageRichLinkBody) HasThumbnailUrl() bool {
	if o != nil && !IsNil(o.ThumbnailUrl) {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given string and assigns it to the ThumbnailUrl field.
func (o *MessageRichLinkBody) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = &v
}

// GetButtonText returns the ButtonText field value if set, zero value otherwise.
func (o *MessageRichLinkBody) GetButtonText() string {
	if o == nil || IsNil(o.ButtonText) {
		var ret string
		return ret
	}
	return *o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetButtonTextOk() (*string, bool) {
	if o == nil || IsNil(o.ButtonText) {
		return nil, false
	}
	return o.ButtonText, true
}

// HasButtonText returns a boolean if a field has been set.
func (o *MessageRichLinkBody) HasButtonText() bool {
	if o != nil && !IsNil(o.ButtonText) {
		return true
	}

	return false
}

// SetButtonText gets a reference to the given string and assigns it to the ButtonText field.
func (o *MessageRichLinkBody) SetButtonText(v string) {
	o.ButtonText = &v
}

// GetCardOptions returns the CardOptions field value if set, zero value otherwise.
func (o *MessageRichLinkBody) GetCardOptions() MessageCardOptions {
	if o == nil || IsNil(o.CardOptions) {
		var ret MessageCardOptions
		return ret
	}
	return *o.CardOptions
}

// GetCardOptionsOk returns a tuple with the CardOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageRichLinkBody) GetCardOptionsOk() (*MessageCardOptions, bool) {
	if o == nil || IsNil(o.CardOptions) {
		return nil, false
	}
	return o.CardOptions, true
}

// HasCardOptions returns a boolean if a field has been set.
func (o *MessageRichLinkBody) HasCardOptions() bool {
	if o != nil && !IsNil(o.CardOptions) {
		return true
	}

	return false
}

// SetCardOptions gets a reference to the given MessageCardOptions and assigns it to the CardOptions field.
func (o *MessageRichLinkBody) SetCardOptions(v MessageCardOptions) {
	o.CardOptions = &v
}

func (o MessageRichLinkBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageRichLinkBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["redirectUrl"] = o.RedirectUrl
	toSerialize["text"] = o.Text
	toSerialize["url"] = o.Url
	if !IsNil(o.IsVideo) {
		toSerialize["isVideo"] = o.IsVideo
	}
	if !IsNil(o.ThumbnailUrl) {
		toSerialize["thumbnailUrl"] = o.ThumbnailUrl
	}
	if !IsNil(o.ButtonText) {
		toSerialize["buttonText"] = o.ButtonText
	}
	if !IsNil(o.CardOptions) {
		toSerialize["cardOptions"] = o.CardOptions
	}
	return toSerialize, nil
}

type NullableMessageRichLinkBody struct {
	value *MessageRichLinkBody
	isSet bool
}

func (v NullableMessageRichLinkBody) Get() *MessageRichLinkBody {
	return v.value
}

func (v *NullableMessageRichLinkBody) Set(val *MessageRichLinkBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageRichLinkBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageRichLinkBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageRichLinkBody(val *MessageRichLinkBody) *NullableMessageRichLinkBody {
	return &NullableMessageRichLinkBody{value: val, isSet: true}
}

func (v NullableMessageRichLinkBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageRichLinkBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
