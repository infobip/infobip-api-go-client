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

// checks if the SwitchCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchCase{}

// SwitchCase Switch-Case action takes a saved variable (obtained via either a Collect, Call API action or scenario parameters) and performs flow control based on it. It works similarly to the switch case block, using case-sensitive comparison.
type SwitchCase struct {
	// Name of the variable whose value to inspect.
	Switch string
	Case   CaseObject
	// User-defined ID of an action that can be used with go-to action.
	ActionId *int32
}

type _SwitchCase SwitchCase

// NewSwitchCase instantiates a new SwitchCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSwitchCase(switch_ string, case_ CaseObject) *SwitchCase {
	this := SwitchCase{}
	this.Switch = switch_
	this.Case = case_
	return &this
}

// NewSwitchCaseWithDefaults instantiates a new SwitchCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchCaseWithDefaults() *SwitchCase {
	this := SwitchCase{}

	return &this
}

// GetSwitch returns the Switch field value
func (o *SwitchCase) GetSwitch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value
// and a boolean to check if the value has been set.
func (o *SwitchCase) GetSwitchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Switch, true
}

// SetSwitch sets field value
func (o *SwitchCase) SetSwitch(v string) {
	o.Switch = v
}

// GetCase returns the Case field value
func (o *SwitchCase) GetCase() CaseObject {
	if o == nil {
		var ret CaseObject
		return ret
	}

	return o.Case
}

// GetCaseOk returns a tuple with the Case field value
// and a boolean to check if the value has been set.
func (o *SwitchCase) GetCaseOk() (*CaseObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Case, true
}

// SetCase sets field value
func (o *SwitchCase) SetCase(v CaseObject) {
	o.Case = v
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *SwitchCase) GetActionId() int32 {
	if o == nil || IsNil(o.ActionId) {
		var ret int32
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchCase) GetActionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *SwitchCase) HasActionId() bool {
	if o != nil && !IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given int32 and assigns it to the ActionId field.
func (o *SwitchCase) SetActionId(v int32) {
	o.ActionId = &v
}

func (o SwitchCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["switch"] = o.Switch
	toSerialize["case"] = o.Case
	if !IsNil(o.ActionId) {
		toSerialize["actionId"] = o.ActionId
	}
	return toSerialize, nil
}

type NullableSwitchCase struct {
	value *SwitchCase
	isSet bool
}

func (v NullableSwitchCase) Get() *SwitchCase {
	return v.value
}

func (v *NullableSwitchCase) Set(val *SwitchCase) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchCase) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchCase(val *SwitchCase) *NullableSwitchCase {
	return &NullableSwitchCase{value: val, isSet: true}
}

func (v NullableSwitchCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
