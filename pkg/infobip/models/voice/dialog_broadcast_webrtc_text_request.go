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

// checks if the DialogBroadcastWebrtcTextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DialogBroadcastWebrtcTextRequest{}

// DialogBroadcastWebrtcTextRequest struct for DialogBroadcastWebrtcTextRequest
type DialogBroadcastWebrtcTextRequest struct {
	// Text to broadcast.
	Text *string
}

// NewDialogBroadcastWebrtcTextRequest instantiates a new DialogBroadcastWebrtcTextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDialogBroadcastWebrtcTextRequest() *DialogBroadcastWebrtcTextRequest {
	this := DialogBroadcastWebrtcTextRequest{}
	return &this
}

// NewDialogBroadcastWebrtcTextRequestWithDefaults instantiates a new DialogBroadcastWebrtcTextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDialogBroadcastWebrtcTextRequestWithDefaults() *DialogBroadcastWebrtcTextRequest {
	this := DialogBroadcastWebrtcTextRequest{}

	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *DialogBroadcastWebrtcTextRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogBroadcastWebrtcTextRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *DialogBroadcastWebrtcTextRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *DialogBroadcastWebrtcTextRequest) SetText(v string) {
	o.Text = &v
}

func (o DialogBroadcastWebrtcTextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DialogBroadcastWebrtcTextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableDialogBroadcastWebrtcTextRequest struct {
	value *DialogBroadcastWebrtcTextRequest
	isSet bool
}

func (v NullableDialogBroadcastWebrtcTextRequest) Get() *DialogBroadcastWebrtcTextRequest {
	return v.value
}

func (v *NullableDialogBroadcastWebrtcTextRequest) Set(val *DialogBroadcastWebrtcTextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDialogBroadcastWebrtcTextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDialogBroadcastWebrtcTextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDialogBroadcastWebrtcTextRequest(val *DialogBroadcastWebrtcTextRequest) *NullableDialogBroadcastWebrtcTextRequest {
	return &NullableDialogBroadcastWebrtcTextRequest{value: val, isSet: true}
}

func (v NullableDialogBroadcastWebrtcTextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDialogBroadcastWebrtcTextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
