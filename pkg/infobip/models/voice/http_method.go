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

// HttpMethod Http method
type HttpMethod string

// List of HttpMethod
const (
	HTTPMETHOD_GET    HttpMethod = "GET"
	HTTPMETHOD_POST   HttpMethod = "POST"
	HTTPMETHOD_PUT    HttpMethod = "PUT"
	HTTPMETHOD_DELETE HttpMethod = "DELETE"
	HTTPMETHOD_PATCH  HttpMethod = "PATCH"
)

// All allowed values of HttpMethod enum
var AllowedHttpMethodEnumValues = []HttpMethod{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"PATCH",
}

func (v *HttpMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HttpMethod(value)
	for _, existing := range AllowedHttpMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HttpMethod", value)
}

// NewHttpMethodFromValue returns a pointer to a valid HttpMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHttpMethodFromValue(v string) (*HttpMethod, error) {
	ev := HttpMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HttpMethod: valid values are %v", v, AllowedHttpMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HttpMethod) IsValid() bool {
	for _, existing := range AllowedHttpMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HttpMethod value
func (v HttpMethod) Ptr() *HttpMethod {
	return &v
}

type NullableHttpMethod struct {
	value *HttpMethod
	isSet bool
}

func (v NullableHttpMethod) Get() *HttpMethod {
	return v.value
}

func (v *NullableHttpMethod) Set(val *HttpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpMethod(val *HttpMethod) *NullableHttpMethod {
	return &NullableHttpMethod{value: val, isSet: true}
}

func (v NullableHttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
