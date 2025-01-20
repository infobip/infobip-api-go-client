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

// checks if the RecordOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordOptions{}

// RecordOptions struct for RecordOptions
type RecordOptions struct {
	// DTMF codes which should stop recording when input by the user.
	EscapeDigits *string
	// Flag indicating a beep sound should be played at the start of the recording.
	Beep *bool
	// The amount of silence to wait for before stopping the recording.
	MaxSilence *int32
	// The identifier for the recording. Identified recordings can be reused in Play from Recording. The parameter can be constructed using variables.
	Identifier *string
}

// NewRecordOptions instantiates a new RecordOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewRecordOptions() *RecordOptions {
	this := RecordOptions{}
	return &this
}

// NewRecordOptionsWithDefaults instantiates a new RecordOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordOptionsWithDefaults() *RecordOptions {
	this := RecordOptions{}

	return &this
}

// GetEscapeDigits returns the EscapeDigits field value if set, zero value otherwise.
func (o *RecordOptions) GetEscapeDigits() string {
	if o == nil || IsNil(o.EscapeDigits) {
		var ret string
		return ret
	}
	return *o.EscapeDigits
}

// GetEscapeDigitsOk returns a tuple with the EscapeDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordOptions) GetEscapeDigitsOk() (*string, bool) {
	if o == nil || IsNil(o.EscapeDigits) {
		return nil, false
	}
	return o.EscapeDigits, true
}

// HasEscapeDigits returns a boolean if a field has been set.
func (o *RecordOptions) HasEscapeDigits() bool {
	if o != nil && !IsNil(o.EscapeDigits) {
		return true
	}

	return false
}

// SetEscapeDigits gets a reference to the given string and assigns it to the EscapeDigits field.
func (o *RecordOptions) SetEscapeDigits(v string) {
	o.EscapeDigits = &v
}

// GetBeep returns the Beep field value if set, zero value otherwise.
func (o *RecordOptions) GetBeep() bool {
	if o == nil || IsNil(o.Beep) {
		var ret bool
		return ret
	}
	return *o.Beep
}

// GetBeepOk returns a tuple with the Beep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordOptions) GetBeepOk() (*bool, bool) {
	if o == nil || IsNil(o.Beep) {
		return nil, false
	}
	return o.Beep, true
}

// HasBeep returns a boolean if a field has been set.
func (o *RecordOptions) HasBeep() bool {
	if o != nil && !IsNil(o.Beep) {
		return true
	}

	return false
}

// SetBeep gets a reference to the given bool and assigns it to the Beep field.
func (o *RecordOptions) SetBeep(v bool) {
	o.Beep = &v
}

// GetMaxSilence returns the MaxSilence field value if set, zero value otherwise.
func (o *RecordOptions) GetMaxSilence() int32 {
	if o == nil || IsNil(o.MaxSilence) {
		var ret int32
		return ret
	}
	return *o.MaxSilence
}

// GetMaxSilenceOk returns a tuple with the MaxSilence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordOptions) GetMaxSilenceOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSilence) {
		return nil, false
	}
	return o.MaxSilence, true
}

// HasMaxSilence returns a boolean if a field has been set.
func (o *RecordOptions) HasMaxSilence() bool {
	if o != nil && !IsNil(o.MaxSilence) {
		return true
	}

	return false
}

// SetMaxSilence gets a reference to the given int32 and assigns it to the MaxSilence field.
func (o *RecordOptions) SetMaxSilence(v int32) {
	o.MaxSilence = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *RecordOptions) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordOptions) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *RecordOptions) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *RecordOptions) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o RecordOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EscapeDigits) {
		toSerialize["escapeDigits"] = o.EscapeDigits
	}
	if !IsNil(o.Beep) {
		toSerialize["beep"] = o.Beep
	}
	if !IsNil(o.MaxSilence) {
		toSerialize["maxSilence"] = o.MaxSilence
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	return toSerialize, nil
}

type NullableRecordOptions struct {
	value *RecordOptions
	isSet bool
}

func (v NullableRecordOptions) Get() *RecordOptions {
	return v.value
}

func (v *NullableRecordOptions) Set(val *RecordOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordOptions(val *RecordOptions) *NullableRecordOptions {
	return &NullableRecordOptions{value: val, isSet: true}
}

func (v NullableRecordOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
