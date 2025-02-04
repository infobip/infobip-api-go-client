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

// checks if the RecordedIvrFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordedIvrFile{}

// RecordedIvrFile Array of recorded files metadata, one for each recorded file.
type RecordedIvrFile struct {
	// The ID that uniquely identifies the sent message.
	MessageId *string
	// Numeric sender ID.
	From *string
	// Destination address.
	To *string
	// Scenario key.
	ScenarioId *string
	// Differentiates recordings made by separate Record actions.
	GroupId *string
	// Relative URL path to the recorded file. To download the audio, just perform a GET request using the relative URL of a specific file. The returned audio data is encoded as PCM 16bit 8kHz WAVE audio. The files are available on Infobip servers for 2 months.
	Url *string
	// The time the recording took place.
	RecordedAt *Time
}

// NewRecordedIvrFile instantiates a new RecordedIvrFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewRecordedIvrFile() *RecordedIvrFile {
	this := RecordedIvrFile{}
	return &this
}

// NewRecordedIvrFileWithDefaults instantiates a new RecordedIvrFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordedIvrFileWithDefaults() *RecordedIvrFile {
	this := RecordedIvrFile{}

	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *RecordedIvrFile) SetMessageId(v string) {
	o.MessageId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *RecordedIvrFile) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *RecordedIvrFile) SetTo(v string) {
	o.To = &v
}

// GetScenarioId returns the ScenarioId field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetScenarioId() string {
	if o == nil || IsNil(o.ScenarioId) {
		var ret string
		return ret
	}
	return *o.ScenarioId
}

// GetScenarioIdOk returns a tuple with the ScenarioId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetScenarioIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScenarioId) {
		return nil, false
	}
	return o.ScenarioId, true
}

// HasScenarioId returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasScenarioId() bool {
	if o != nil && !IsNil(o.ScenarioId) {
		return true
	}

	return false
}

// SetScenarioId gets a reference to the given string and assigns it to the ScenarioId field.
func (o *RecordedIvrFile) SetScenarioId(v string) {
	o.ScenarioId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *RecordedIvrFile) SetGroupId(v string) {
	o.GroupId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RecordedIvrFile) SetUrl(v string) {
	o.Url = &v
}

// GetRecordedAt returns the RecordedAt field value if set, zero value otherwise.
func (o *RecordedIvrFile) GetRecordedAt() Time {
	if o == nil || IsNil(o.RecordedAt) {
		var ret Time
		return ret
	}
	return *o.RecordedAt
}

// GetRecordedAtOk returns a tuple with the RecordedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordedIvrFile) GetRecordedAtOk() (*Time, bool) {
	if o == nil || IsNil(o.RecordedAt) {
		return nil, false
	}
	return o.RecordedAt, true
}

// HasRecordedAt returns a boolean if a field has been set.
func (o *RecordedIvrFile) HasRecordedAt() bool {
	if o != nil && !IsNil(o.RecordedAt) {
		return true
	}

	return false
}

// SetRecordedAt gets a reference to the given Time and assigns it to the RecordedAt field.
func (o *RecordedIvrFile) SetRecordedAt(v Time) {
	o.RecordedAt = &v
}

func (o RecordedIvrFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordedIvrFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["messageId"] = o.MessageId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.ScenarioId) {
		toSerialize["scenarioId"] = o.ScenarioId
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.RecordedAt) {
		toSerialize["recordedAt"] = o.RecordedAt
	}
	return toSerialize, nil
}

type NullableRecordedIvrFile struct {
	value *RecordedIvrFile
	isSet bool
}

func (v NullableRecordedIvrFile) Get() *RecordedIvrFile {
	return v.value
}

func (v *NullableRecordedIvrFile) Set(val *RecordedIvrFile) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordedIvrFile) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordedIvrFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordedIvrFile(val *RecordedIvrFile) *NullableRecordedIvrFile {
	return &NullableRecordedIvrFile{value: val, isSet: true}
}

func (v NullableRecordedIvrFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordedIvrFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
