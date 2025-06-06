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

// checks if the SeenReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeenReport{}

// SeenReport struct for SeenReport
type SeenReport struct {
	// Collection of reports, one per every message.
	Results []SeenResult
}

type _SeenReport SeenReport

// NewSeenReport instantiates a new SeenReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSeenReport(results []SeenResult) *SeenReport {
	this := SeenReport{}
	this.Results = results
	return &this
}

// NewSeenReportWithDefaults instantiates a new SeenReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeenReportWithDefaults() *SeenReport {
	this := SeenReport{}

	return &this
}

// GetResults returns the Results field value
func (o *SeenReport) GetResults() []SeenResult {
	if o == nil {
		var ret []SeenResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *SeenReport) GetResultsOk() ([]SeenResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *SeenReport) SetResults(v []SeenResult) {
	o.Results = v
}

func (o SeenReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeenReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableSeenReport struct {
	value *SeenReport
	isSet bool
}

func (v NullableSeenReport) Get() *SeenReport {
	return v.value
}

func (v *NullableSeenReport) Set(val *SeenReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSeenReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSeenReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeenReport(val *SeenReport) *NullableSeenReport {
	return &NullableSeenReport{value: val, isSet: true}
}

func (v NullableSeenReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeenReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
