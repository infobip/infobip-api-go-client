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

// checks if the MessageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageContent{}

// MessageContent Content of the message.
type MessageContent struct {
	Header *MessageHeader
	Body   MessageBody
	// List of buttons of the message.
	Buttons          []MessageButton
	ConfirmationBody *MessageConfirmationBody
	Footer           *MessageFooter
}

type _MessageContent MessageContent

// NewMessageContent instantiates a new MessageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMessageContent(body MessageBody) *MessageContent {
	this := MessageContent{}
	this.Body = body
	return &this
}

// NewMessageContentWithDefaults instantiates a new MessageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageContentWithDefaults() *MessageContent {
	this := MessageContent{}

	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *MessageContent) GetHeader() MessageHeader {
	if o == nil || IsNil(o.Header) {
		var ret MessageHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageContent) GetHeaderOk() (*MessageHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *MessageContent) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given MessageHeader and assigns it to the Header field.
func (o *MessageContent) SetHeader(v MessageHeader) {
	o.Header = &v
}

// GetBody returns the Body field value
func (o *MessageContent) GetBody() MessageBody {
	if o == nil {
		var ret MessageBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *MessageContent) GetBodyOk() (*MessageBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *MessageContent) SetBody(v MessageBody) {
	o.Body = v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *MessageContent) GetButtons() []MessageButton {
	if o == nil || IsNil(o.Buttons) {
		var ret []MessageButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageContent) GetButtonsOk() ([]MessageButton, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *MessageContent) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []MessageButton and assigns it to the Buttons field.
func (o *MessageContent) SetButtons(v []MessageButton) {
	o.Buttons = v
}

// GetConfirmationBody returns the ConfirmationBody field value if set, zero value otherwise.
func (o *MessageContent) GetConfirmationBody() MessageConfirmationBody {
	if o == nil || IsNil(o.ConfirmationBody) {
		var ret MessageConfirmationBody
		return ret
	}
	return *o.ConfirmationBody
}

// GetConfirmationBodyOk returns a tuple with the ConfirmationBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageContent) GetConfirmationBodyOk() (*MessageConfirmationBody, bool) {
	if o == nil || IsNil(o.ConfirmationBody) {
		return nil, false
	}
	return o.ConfirmationBody, true
}

// HasConfirmationBody returns a boolean if a field has been set.
func (o *MessageContent) HasConfirmationBody() bool {
	if o != nil && !IsNil(o.ConfirmationBody) {
		return true
	}

	return false
}

// SetConfirmationBody gets a reference to the given MessageConfirmationBody and assigns it to the ConfirmationBody field.
func (o *MessageContent) SetConfirmationBody(v MessageConfirmationBody) {
	o.ConfirmationBody = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *MessageContent) GetFooter() MessageFooter {
	if o == nil || IsNil(o.Footer) {
		var ret MessageFooter
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageContent) GetFooterOk() (*MessageFooter, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *MessageContent) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given MessageFooter and assigns it to the Footer field.
func (o *MessageContent) SetFooter(v MessageFooter) {
	o.Footer = &v
}

func (o MessageContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	toSerialize["body"] = o.Body
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}
	if !IsNil(o.ConfirmationBody) {
		toSerialize["confirmationBody"] = o.ConfirmationBody
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	return toSerialize, nil
}

type NullableMessageContent struct {
	value *MessageContent
	isSet bool
}

func (v NullableMessageContent) Get() *MessageContent {
	return v.value
}

func (v *NullableMessageContent) Set(val *MessageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageContent(val *MessageContent) *NullableMessageContent {
	return &NullableMessageContent{value: val, isSet: true}
}

func (v NullableMessageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
