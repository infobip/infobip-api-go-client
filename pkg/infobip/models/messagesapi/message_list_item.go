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

// checks if the MessageListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageListItem{}

// MessageListItem List Items.
type MessageListItem struct {
	// Item ID (supported only by WhatsApp), has to be unique for each item.
	Id *string
	// Item Text.
	Text string
	// Item Description.
	Description *string
	// Item's Image URL.
	ImageUrl *string
}

type _MessageListItem MessageListItem

// NewMessageListItem instantiates a new MessageListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMessageListItem(text string) *MessageListItem {
	this := MessageListItem{}
	this.Text = text
	return &this
}

// NewMessageListItemWithDefaults instantiates a new MessageListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageListItemWithDefaults() *MessageListItem {
	this := MessageListItem{}

	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageListItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageListItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageListItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MessageListItem) SetId(v string) {
	o.Id = &v
}

// GetText returns the Text field value
func (o *MessageListItem) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *MessageListItem) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *MessageListItem) SetText(v string) {
	o.Text = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MessageListItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MessageListItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MessageListItem) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *MessageListItem) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageListItem) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *MessageListItem) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *MessageListItem) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o MessageListItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["text"] = o.Text
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	return toSerialize, nil
}

type NullableMessageListItem struct {
	value *MessageListItem
	isSet bool
}

func (v NullableMessageListItem) Get() *MessageListItem {
	return v.value
}

func (v *NullableMessageListItem) Set(val *MessageListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageListItem(val *MessageListItem) *NullableMessageListItem {
	return &NullableMessageListItem{value: val, isSet: true}
}

func (v NullableMessageListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
