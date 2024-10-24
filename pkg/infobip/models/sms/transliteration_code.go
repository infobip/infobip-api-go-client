/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sms

import (
	"encoding/json"
	"fmt"
)

// TransliterationCode The transliteration of your sent message from one script to another. Transliteration is used to replace characters which are not recognized as part of your defaulted alphabet. Possible values: `TURKISH`, `GREEK`, `CYRILLIC`, `SERBIAN_CYRILLIC`, `BULGARIAN_CYRILLIC`, `CENTRAL_EUROPEAN`, `BALTIC`, `PORTUGUESE`, `COLOMBIAN`, and `NON_UNICODE`.
type TransliterationCode string

// List of TransliterationCode
const (
	TRANSLITERATIONCODE_NONE               TransliterationCode = "NONE"
	TRANSLITERATIONCODE_TURKISH            TransliterationCode = "TURKISH"
	TRANSLITERATIONCODE_GREEK              TransliterationCode = "GREEK"
	TRANSLITERATIONCODE_CYRILLIC           TransliterationCode = "CYRILLIC"
	TRANSLITERATIONCODE_SERBIAN_CYRILLIC   TransliterationCode = "SERBIAN_CYRILLIC"
	TRANSLITERATIONCODE_CENTRAL_EUROPEAN   TransliterationCode = "CENTRAL_EUROPEAN"
	TRANSLITERATIONCODE_BALTIC             TransliterationCode = "BALTIC"
	TRANSLITERATIONCODE_NON_UNICODE        TransliterationCode = "NON_UNICODE"
	TRANSLITERATIONCODE_PORTUGUESE         TransliterationCode = "PORTUGUESE"
	TRANSLITERATIONCODE_COLOMBIAN          TransliterationCode = "COLOMBIAN"
	TRANSLITERATIONCODE_BULGARIAN_CYRILLIC TransliterationCode = "BULGARIAN_CYRILLIC"
	TRANSLITERATIONCODE_ALL                TransliterationCode = "ALL"
)

// All allowed values of TransliterationCode enum
var AllowedTransliterationCodeEnumValues = []TransliterationCode{
	"NONE",
	"TURKISH",
	"GREEK",
	"CYRILLIC",
	"SERBIAN_CYRILLIC",
	"CENTRAL_EUROPEAN",
	"BALTIC",
	"NON_UNICODE",
	"PORTUGUESE",
	"COLOMBIAN",
	"BULGARIAN_CYRILLIC",
	"ALL",
}

func (v *TransliterationCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransliterationCode(value)
	for _, existing := range AllowedTransliterationCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransliterationCode", value)
}

// NewTransliterationCodeFromValue returns a pointer to a valid TransliterationCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransliterationCodeFromValue(v string) (*TransliterationCode, error) {
	ev := TransliterationCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransliterationCode: valid values are %v", v, AllowedTransliterationCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransliterationCode) IsValid() bool {
	for _, existing := range AllowedTransliterationCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransliterationCode value
func (v TransliterationCode) Ptr() *TransliterationCode {
	return &v
}

type NullableTransliterationCode struct {
	value *TransliterationCode
	isSet bool
}

func (v NullableTransliterationCode) Get() *TransliterationCode {
	return v.value
}

func (v *NullableTransliterationCode) Set(val *TransliterationCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransliterationCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransliterationCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransliterationCode(val *TransliterationCode) *NullableTransliterationCode {
	return &NullableTransliterationCode{value: val, isSet: true}
}

func (v NullableTransliterationCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransliterationCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}