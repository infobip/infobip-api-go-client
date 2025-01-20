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

// checks if the TemplateMessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateMessageContent{}

// TemplateMessageContent Content of the message.
type TemplateMessageContent struct {
	Header *TemplateHeader
	Body   *TemplateBody
	// List of buttons of a template message.
	Buttons []TemplateButton
	// Expiration time of a limited-time offer template message. Has the following format: yyyy-MM-dd'T'HH:mm:ss.SSSZ.
	ExpirationTime *string
}

// NewTemplateMessageContent instantiates a new TemplateMessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewTemplateMessageContent() *TemplateMessageContent {
	this := TemplateMessageContent{}
	return &this
}

// NewTemplateMessageContentWithDefaults instantiates a new TemplateMessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateMessageContentWithDefaults() *TemplateMessageContent {
	this := TemplateMessageContent{}

	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *TemplateMessageContent) GetHeader() TemplateHeader {
	if o == nil || IsNil(o.Header) {
		var ret TemplateHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessageContent) GetHeaderOk() (*TemplateHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *TemplateMessageContent) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given TemplateHeader and assigns it to the Header field.
func (o *TemplateMessageContent) SetHeader(v TemplateHeader) {
	o.Header = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *TemplateMessageContent) GetBody() TemplateBody {
	if o == nil || IsNil(o.Body) {
		var ret TemplateBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessageContent) GetBodyOk() (*TemplateBody, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *TemplateMessageContent) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given TemplateBody and assigns it to the Body field.
func (o *TemplateMessageContent) SetBody(v TemplateBody) {
	o.Body = &v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *TemplateMessageContent) GetButtons() []TemplateButton {
	if o == nil || IsNil(o.Buttons) {
		var ret []TemplateButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessageContent) GetButtonsOk() ([]TemplateButton, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *TemplateMessageContent) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []TemplateButton and assigns it to the Buttons field.
func (o *TemplateMessageContent) SetButtons(v []TemplateButton) {
	o.Buttons = v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *TemplateMessageContent) GetExpirationTime() string {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret string
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateMessageContent) GetExpirationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *TemplateMessageContent) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given string and assigns it to the ExpirationTime field.
func (o *TemplateMessageContent) SetExpirationTime(v string) {
	o.ExpirationTime = &v
}

func (o TemplateMessageContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateMessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	return toSerialize, nil
}

type NullableTemplateMessageContent struct {
	value *TemplateMessageContent
	isSet bool
}

func (v NullableTemplateMessageContent) Get() *TemplateMessageContent {
	return v.value
}

func (v *NullableTemplateMessageContent) Set(val *TemplateMessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateMessageContent(val *TemplateMessageContent) *NullableTemplateMessageContent {
	return &NullableTemplateMessageContent{value: val, isSet: true}
}

func (v NullableTemplateMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
