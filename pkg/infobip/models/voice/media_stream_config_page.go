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

// checks if the MediaStreamConfigPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaStreamConfigPage{}

// MediaStreamConfigPage struct for MediaStreamConfigPage
type MediaStreamConfigPage struct {
	// The list of the results for this page.
	Results []MediaStreamConfigResponse
	Paging  *PageInfo
}

// NewMediaStreamConfigPage instantiates a new MediaStreamConfigPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewMediaStreamConfigPage() *MediaStreamConfigPage {
	this := MediaStreamConfigPage{}
	return &this
}

// NewMediaStreamConfigPageWithDefaults instantiates a new MediaStreamConfigPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaStreamConfigPageWithDefaults() *MediaStreamConfigPage {
	this := MediaStreamConfigPage{}

	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MediaStreamConfigPage) GetResults() []MediaStreamConfigResponse {
	if o == nil || IsNil(o.Results) {
		var ret []MediaStreamConfigResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamConfigPage) GetResultsOk() ([]MediaStreamConfigResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MediaStreamConfigPage) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MediaStreamConfigResponse and assigns it to the Results field.
func (o *MediaStreamConfigPage) SetResults(v []MediaStreamConfigResponse) {
	o.Results = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *MediaStreamConfigPage) GetPaging() PageInfo {
	if o == nil || IsNil(o.Paging) {
		var ret PageInfo
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamConfigPage) GetPagingOk() (*PageInfo, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *MediaStreamConfigPage) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given PageInfo and assigns it to the Paging field.
func (o *MediaStreamConfigPage) SetPaging(v PageInfo) {
	o.Paging = &v
}

func (o MediaStreamConfigPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaStreamConfigPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableMediaStreamConfigPage struct {
	value *MediaStreamConfigPage
	isSet bool
}

func (v NullableMediaStreamConfigPage) Get() *MediaStreamConfigPage {
	return v.value
}

func (v *NullableMediaStreamConfigPage) Set(val *MediaStreamConfigPage) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamConfigPage) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamConfigPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamConfigPage(val *MediaStreamConfigPage) *NullableMediaStreamConfigPage {
	return &NullableMediaStreamConfigPage{value: val, isSet: true}
}

func (v NullableMediaStreamConfigPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamConfigPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
