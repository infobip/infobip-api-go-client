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

// checks if the GoTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoTo{}

// GoTo Go-To action is used to go back to some specified action that was already executed in the scenario and continue the execution from that point.
type GoTo struct {
	// The actionId of an action to which to jump
	GoTo    int32
	Options *GoToOptions
}

type _GoTo GoTo

// NewGoTo instantiates a new GoTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewGoTo(goTo int32) *GoTo {
	this := GoTo{}
	this.GoTo = goTo
	return &this
}

// NewGoToWithDefaults instantiates a new GoTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoToWithDefaults() *GoTo {
	this := GoTo{}

	return &this
}

// GetGoTo returns the GoTo field value
func (o *GoTo) GetGoTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GoTo
}

// GetGoToOk returns a tuple with the GoTo field value
// and a boolean to check if the value has been set.
func (o *GoTo) GetGoToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoTo, true
}

// SetGoTo sets field value
func (o *GoTo) SetGoTo(v int32) {
	o.GoTo = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *GoTo) GetOptions() GoToOptions {
	if o == nil || IsNil(o.Options) {
		var ret GoToOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoTo) GetOptionsOk() (*GoToOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GoTo) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given GoToOptions and assigns it to the Options field.
func (o *GoTo) SetOptions(v GoToOptions) {
	o.Options = &v
}

func (o GoTo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoTo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["goTo"] = o.GoTo
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableGoTo struct {
	value *GoTo
	isSet bool
}

func (v NullableGoTo) Get() *GoTo {
	return v.value
}

func (v *NullableGoTo) Set(val *GoTo) {
	v.value = val
	v.isSet = true
}

func (v NullableGoTo) IsSet() bool {
	return v.isSet
}

func (v *NullableGoTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoTo(val *GoTo) *NullableGoTo {
	return &NullableGoTo{value: val, isSet: true}
}

func (v NullableGoTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
