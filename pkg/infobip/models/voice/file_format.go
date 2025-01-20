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

// FileFormat File format.
type FileFormat string

// List of FileFormat
const (
	FILEFORMAT_MP3 FileFormat = "MP3"
	FILEFORMAT_WAV FileFormat = "WAV"
	FILEFORMAT_MP4 FileFormat = "MP4"
)

// All allowed values of FileFormat enum
var AllowedFileFormatEnumValues = []FileFormat{
	"MP3",
	"WAV",
	"MP4",
}

func (v *FileFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileFormat(value)
	for _, existing := range AllowedFileFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileFormat", value)
}

// NewFileFormatFromValue returns a pointer to a valid FileFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileFormatFromValue(v string) (*FileFormat, error) {
	ev := FileFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileFormat: valid values are %v", v, AllowedFileFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileFormat) IsValid() bool {
	for _, existing := range AllowedFileFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileFormat value
func (v FileFormat) Ptr() *FileFormat {
	return &v
}

type NullableFileFormat struct {
	value *FileFormat
	isSet bool
}

func (v NullableFileFormat) Get() *FileFormat {
	return v.value
}

func (v *NullableFileFormat) Set(val *FileFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableFileFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableFileFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileFormat(val *FileFormat) *NullableFileFormat {
	return &NullableFileFormat{value: val, isSet: true}
}

func (v NullableFileFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
