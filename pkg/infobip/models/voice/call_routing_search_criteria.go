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

// checks if the CallRoutingSearchCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRoutingSearchCriteria{}

// CallRoutingSearchCriteria List of criteria that should match route. For a route to match, any criterion should be met.
type CallRoutingSearchCriteria struct {
	// Phone number or regular expression pattern representing the phone number. Applies to all search criteria types except `WEBRTC`. Examples: `41793026727`, `41793(.+)`.
	To    *string
	Value *CallRoutingCriteria
}

// NewCallRoutingSearchCriteria instantiates a new CallRoutingSearchCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCallRoutingSearchCriteria() *CallRoutingSearchCriteria {
	this := CallRoutingSearchCriteria{}
	return &this
}

// NewCallRoutingSearchCriteriaWithDefaults instantiates a new CallRoutingSearchCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRoutingSearchCriteriaWithDefaults() *CallRoutingSearchCriteria {
	this := CallRoutingSearchCriteria{}

	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CallRoutingSearchCriteria) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingSearchCriteria) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CallRoutingSearchCriteria) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CallRoutingSearchCriteria) SetTo(v string) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CallRoutingSearchCriteria) GetValue() CallRoutingCriteria {
	if o == nil || IsNil(o.Value) {
		var ret CallRoutingCriteria
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingSearchCriteria) GetValueOk() (*CallRoutingCriteria, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CallRoutingSearchCriteria) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CallRoutingCriteria and assigns it to the Value field.
func (o *CallRoutingSearchCriteria) SetValue(v CallRoutingCriteria) {
	o.Value = &v
}

func (o CallRoutingSearchCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRoutingSearchCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCallRoutingSearchCriteria struct {
	value *CallRoutingSearchCriteria
	isSet bool
}

func (v NullableCallRoutingSearchCriteria) Get() *CallRoutingSearchCriteria {
	return v.value
}

func (v *NullableCallRoutingSearchCriteria) Set(val *CallRoutingSearchCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRoutingSearchCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRoutingSearchCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRoutingSearchCriteria(val *CallRoutingSearchCriteria) *NullableCallRoutingSearchCriteria {
	return &NullableCallRoutingSearchCriteria{value: val, isSet: true}
}

func (v NullableCallRoutingSearchCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRoutingSearchCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
