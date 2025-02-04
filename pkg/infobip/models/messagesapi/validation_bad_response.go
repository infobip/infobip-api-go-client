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

// checks if the ValidationBadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationBadResponse{}

// ValidationBadResponse struct for ValidationBadResponse
type ValidationBadResponse struct {
	// A detailed description of an error.
	Description string
	// An action that should be taken to recover from the error.
	Action string
	// List of violations that caused the error.
	Violations []ApiErrorViolation
	// List of violations that may be omitted, but is recommended to address.
	SkippableViolations []ApiErrorViolation
}

type _ValidationBadResponse ValidationBadResponse

// NewValidationBadResponse instantiates a new ValidationBadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewValidationBadResponse(description string, action string, violations []ApiErrorViolation) *ValidationBadResponse {
	this := ValidationBadResponse{}
	this.Description = description
	this.Action = action
	this.Violations = violations
	return &this
}

// NewValidationBadResponseWithDefaults instantiates a new ValidationBadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationBadResponseWithDefaults() *ValidationBadResponse {
	this := ValidationBadResponse{}

	return &this
}

// GetDescription returns the Description field value
func (o *ValidationBadResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ValidationBadResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ValidationBadResponse) SetDescription(v string) {
	o.Description = v
}

// GetAction returns the Action field value
func (o *ValidationBadResponse) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ValidationBadResponse) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ValidationBadResponse) SetAction(v string) {
	o.Action = v
}

// GetViolations returns the Violations field value
func (o *ValidationBadResponse) GetViolations() []ApiErrorViolation {
	if o == nil {
		var ret []ApiErrorViolation
		return ret
	}

	return o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value
// and a boolean to check if the value has been set.
func (o *ValidationBadResponse) GetViolationsOk() ([]ApiErrorViolation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Violations, true
}

// SetViolations sets field value
func (o *ValidationBadResponse) SetViolations(v []ApiErrorViolation) {
	o.Violations = v
}

// GetSkippableViolations returns the SkippableViolations field value if set, zero value otherwise.
func (o *ValidationBadResponse) GetSkippableViolations() []ApiErrorViolation {
	if o == nil || IsNil(o.SkippableViolations) {
		var ret []ApiErrorViolation
		return ret
	}
	return o.SkippableViolations
}

// GetSkippableViolationsOk returns a tuple with the SkippableViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationBadResponse) GetSkippableViolationsOk() ([]ApiErrorViolation, bool) {
	if o == nil || IsNil(o.SkippableViolations) {
		return nil, false
	}
	return o.SkippableViolations, true
}

// HasSkippableViolations returns a boolean if a field has been set.
func (o *ValidationBadResponse) HasSkippableViolations() bool {
	if o != nil && !IsNil(o.SkippableViolations) {
		return true
	}

	return false
}

// SetSkippableViolations gets a reference to the given []ApiErrorViolation and assigns it to the SkippableViolations field.
func (o *ValidationBadResponse) SetSkippableViolations(v []ApiErrorViolation) {
	o.SkippableViolations = v
}

func (o ValidationBadResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationBadResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["action"] = o.Action
	toSerialize["violations"] = o.Violations
	if !IsNil(o.SkippableViolations) {
		toSerialize["skippableViolations"] = o.SkippableViolations
	}
	return toSerialize, nil
}

type NullableValidationBadResponse struct {
	value *ValidationBadResponse
	isSet bool
}

func (v NullableValidationBadResponse) Get() *ValidationBadResponse {
	return v.value
}

func (v *NullableValidationBadResponse) Set(val *ValidationBadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationBadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationBadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationBadResponse(val *ValidationBadResponse) *NullableValidationBadResponse {
	return &NullableValidationBadResponse{value: val, isSet: true}
}

func (v NullableValidationBadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationBadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
