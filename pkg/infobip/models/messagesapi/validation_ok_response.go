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

// checks if the ValidationOkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationOkResponse{}

// ValidationOkResponse struct for ValidationOkResponse
type ValidationOkResponse struct {
	// A detailed description.
	Description string
	// An action that should be taken.
	Action string
	// List of violations that may be omitted, but is recommended to address.
	SkippableViolations []ApiErrorViolation
}

type _ValidationOkResponse ValidationOkResponse

// NewValidationOkResponse instantiates a new ValidationOkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewValidationOkResponse(description string, action string) *ValidationOkResponse {
	this := ValidationOkResponse{}
	this.Description = description
	this.Action = action
	return &this
}

// NewValidationOkResponseWithDefaults instantiates a new ValidationOkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationOkResponseWithDefaults() *ValidationOkResponse {
	this := ValidationOkResponse{}

	return &this
}

// GetDescription returns the Description field value
func (o *ValidationOkResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ValidationOkResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ValidationOkResponse) SetDescription(v string) {
	o.Description = v
}

// GetAction returns the Action field value
func (o *ValidationOkResponse) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ValidationOkResponse) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ValidationOkResponse) SetAction(v string) {
	o.Action = v
}

// GetSkippableViolations returns the SkippableViolations field value if set, zero value otherwise.
func (o *ValidationOkResponse) GetSkippableViolations() []ApiErrorViolation {
	if o == nil || IsNil(o.SkippableViolations) {
		var ret []ApiErrorViolation
		return ret
	}
	return o.SkippableViolations
}

// GetSkippableViolationsOk returns a tuple with the SkippableViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationOkResponse) GetSkippableViolationsOk() ([]ApiErrorViolation, bool) {
	if o == nil || IsNil(o.SkippableViolations) {
		return nil, false
	}
	return o.SkippableViolations, true
}

// HasSkippableViolations returns a boolean if a field has been set.
func (o *ValidationOkResponse) HasSkippableViolations() bool {
	if o != nil && !IsNil(o.SkippableViolations) {
		return true
	}

	return false
}

// SetSkippableViolations gets a reference to the given []ApiErrorViolation and assigns it to the SkippableViolations field.
func (o *ValidationOkResponse) SetSkippableViolations(v []ApiErrorViolation) {
	o.SkippableViolations = v
}

func (o ValidationOkResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationOkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["action"] = o.Action
	if !IsNil(o.SkippableViolations) {
		toSerialize["skippableViolations"] = o.SkippableViolations
	}
	return toSerialize, nil
}

type NullableValidationOkResponse struct {
	value *ValidationOkResponse
	isSet bool
}

func (v NullableValidationOkResponse) Get() *ValidationOkResponse {
	return v.value
}

func (v *NullableValidationOkResponse) Set(val *ValidationOkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationOkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationOkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationOkResponse(val *ValidationOkResponse) *NullableValidationOkResponse {
	return &NullableValidationOkResponse{value: val, isSet: true}
}

func (v NullableValidationOkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationOkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
