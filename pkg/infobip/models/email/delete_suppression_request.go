/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the DeleteSuppressionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSuppressionRequest{}

// DeleteSuppressionRequest Suppressions request.
type DeleteSuppressionRequest struct {
	// Email addresses to delete from the suppression list. Number of destinations cannot exceed 10,000.
	Suppressions []DeleteSuppression
}

type _DeleteSuppressionRequest DeleteSuppressionRequest

// NewDeleteSuppressionRequest instantiates a new DeleteSuppressionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDeleteSuppressionRequest(suppressions []DeleteSuppression) *DeleteSuppressionRequest {
	this := DeleteSuppressionRequest{}
	this.Suppressions = suppressions
	return &this
}

// NewDeleteSuppressionRequestWithDefaults instantiates a new DeleteSuppressionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSuppressionRequestWithDefaults() *DeleteSuppressionRequest {
	this := DeleteSuppressionRequest{}

	return &this
}

// GetSuppressions returns the Suppressions field value
func (o *DeleteSuppressionRequest) GetSuppressions() []DeleteSuppression {
	if o == nil {
		var ret []DeleteSuppression
		return ret
	}

	return o.Suppressions
}

// GetSuppressionsOk returns a tuple with the Suppressions field value
// and a boolean to check if the value has been set.
func (o *DeleteSuppressionRequest) GetSuppressionsOk() ([]DeleteSuppression, bool) {
	if o == nil {
		return nil, false
	}
	return o.Suppressions, true
}

// SetSuppressions sets field value
func (o *DeleteSuppressionRequest) SetSuppressions(v []DeleteSuppression) {
	o.Suppressions = v
}

func (o DeleteSuppressionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSuppressionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suppressions"] = o.Suppressions
	return toSerialize, nil
}

type NullableDeleteSuppressionRequest struct {
	value *DeleteSuppressionRequest
	isSet bool
}

func (v NullableDeleteSuppressionRequest) Get() *DeleteSuppressionRequest {
	return v.value
}

func (v *NullableDeleteSuppressionRequest) Set(val *DeleteSuppressionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSuppressionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSuppressionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSuppressionRequest(val *DeleteSuppressionRequest) *NullableDeleteSuppressionRequest {
	return &NullableDeleteSuppressionRequest{value: val, isSet: true}
}

func (v NullableDeleteSuppressionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSuppressionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
