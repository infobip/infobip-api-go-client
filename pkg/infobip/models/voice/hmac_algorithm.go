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

// HmacAlgorithm Hmac algorithm.
type HmacAlgorithm string

// List of HmacAlgorithm
const (
	HMACALGORITHM_MD5     HmacAlgorithm = "HMAC_MD5"
	HMACALGORITHM_SHA_1   HmacAlgorithm = "HMAC_SHA_1"
	HMACALGORITHM_SHA_224 HmacAlgorithm = "HMAC_SHA_224"
	HMACALGORITHM_SHA_256 HmacAlgorithm = "HMAC_SHA_256"
	HMACALGORITHM_SHA_384 HmacAlgorithm = "HMAC_SHA_384"
	HMACALGORITHM_SHA_512 HmacAlgorithm = "HMAC_SHA_512"
)

// All allowed values of HmacAlgorithm enum
var AllowedHmacAlgorithmEnumValues = []HmacAlgorithm{
	"HMAC_MD5",
	"HMAC_SHA_1",
	"HMAC_SHA_224",
	"HMAC_SHA_256",
	"HMAC_SHA_384",
	"HMAC_SHA_512",
}

func (v *HmacAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HmacAlgorithm(value)
	for _, existing := range AllowedHmacAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HmacAlgorithm", value)
}

// NewHmacAlgorithmFromValue returns a pointer to a valid HmacAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHmacAlgorithmFromValue(v string) (*HmacAlgorithm, error) {
	ev := HmacAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HmacAlgorithm: valid values are %v", v, AllowedHmacAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HmacAlgorithm) IsValid() bool {
	for _, existing := range AllowedHmacAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HmacAlgorithm value
func (v HmacAlgorithm) Ptr() *HmacAlgorithm {
	return &v
}

type NullableHmacAlgorithm struct {
	value *HmacAlgorithm
	isSet bool
}

func (v NullableHmacAlgorithm) Get() *HmacAlgorithm {
	return v.value
}

func (v *NullableHmacAlgorithm) Set(val *HmacAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableHmacAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableHmacAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHmacAlgorithm(val *HmacAlgorithm) *NullableHmacAlgorithm {
	return &NullableHmacAlgorithm{value: val, isSet: true}
}

func (v NullableHmacAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHmacAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
