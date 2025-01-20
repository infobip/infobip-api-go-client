/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voice

import (
	"encoding/json"
	"fmt"
)

// AudioCodec Audio codec.
type AudioCodec string

// List of AudioCodec
const (
	AUDIOCODEC_PCMU AudioCodec = "PCMU"
	AUDIOCODEC_PCMA AudioCodec = "PCMA"
	AUDIOCODEC_G729 AudioCodec = "G729"
)

// All allowed values of AudioCodec enum
var AllowedAudioCodecEnumValues = []AudioCodec{
	"PCMU",
	"PCMA",
	"G729",
}

func (v *AudioCodec) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AudioCodec(value)
	for _, existing := range AllowedAudioCodecEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AudioCodec", value)
}

// NewAudioCodecFromValue returns a pointer to a valid AudioCodec
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudioCodecFromValue(v string) (*AudioCodec, error) {
	ev := AudioCodec(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudioCodec: valid values are %v", v, AllowedAudioCodecEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudioCodec) IsValid() bool {
	for _, existing := range AllowedAudioCodecEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AudioCodec value
func (v AudioCodec) Ptr() *AudioCodec {
	return &v
}

type NullableAudioCodec struct {
	value *AudioCodec
	isSet bool
}

func (v NullableAudioCodec) Get() *AudioCodec {
	return v.value
}

func (v *NullableAudioCodec) Set(val *AudioCodec) {
	v.value = val
	v.isSet = true
}

func (v NullableAudioCodec) IsSet() bool {
	return v.isSet
}

func (v *NullableAudioCodec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudioCodec(val *AudioCodec) *NullableAudioCodec {
	return &NullableAudioCodec{value: val, isSet: true}
}

func (v NullableAudioCodec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudioCodec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
