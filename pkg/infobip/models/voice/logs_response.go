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

// checks if the LogsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsResponse{}

// LogsResponse struct for LogsResponse
type LogsResponse struct {
	// Array of voice message logs, one object per each message.
	Results []LogsReport
}

// NewLogsResponse instantiates a new LogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewLogsResponse() *LogsResponse {
	this := LogsResponse{}
	return &this
}

// NewLogsResponseWithDefaults instantiates a new LogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsResponseWithDefaults() *LogsResponse {
	this := LogsResponse{}

	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *LogsResponse) GetResults() []LogsReport {
	if o == nil || IsNil(o.Results) {
		var ret []LogsReport
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsResponse) GetResultsOk() ([]LogsReport, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *LogsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []LogsReport and assigns it to the Results field.
func (o *LogsResponse) SetResults(v []LogsReport) {
	o.Results = v
}

func (o LogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableLogsResponse struct {
	value *LogsResponse
	isSet bool
}

func (v NullableLogsResponse) Get() *LogsResponse {
	return v.value
}

func (v *NullableLogsResponse) Set(val *LogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsResponse(val *LogsResponse) *NullableLogsResponse {
	return &NullableLogsResponse{value: val, isSet: true}
}

func (v NullableLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
