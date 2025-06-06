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

// checks if the SearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResponse{}

// SearchResponse struct for SearchResponse
type SearchResponse struct {
	// Creation timestamp.
	CreateTime *Time
	// Scenario description.
	Description *string
	// Scenario key. It is used for launching IVR scenario.
	Id *string
	// Scenario name.
	Name *string
	// Array of IVR actions defining scenario. NOTE: Answering Machine Detection, Call Recording and Speech Recognition (used for Capture action) are add-on features. To enable these add-ons, please contact our [sales](https://www.infobip.com/contact) organisation.
	Script []ScriptInner
	// Update timestamp
	UpdateTime *Time
	// Last usage date. `null` for scenarios that are used last time before `2024-01-01`.
	LastUsageDate *Time
}

// NewSearchResponse instantiates a new SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSearchResponse() *SearchResponse {
	this := SearchResponse{}
	return &this
}

// NewSearchResponseWithDefaults instantiates a new SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseWithDefaults() *SearchResponse {
	this := SearchResponse{}

	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *SearchResponse) GetCreateTime() Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetCreateTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *SearchResponse) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given Time and assigns it to the CreateTime field.
func (o *SearchResponse) SetCreateTime(v Time) {
	o.CreateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchResponse) SetName(v string) {
	o.Name = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *SearchResponse) GetScript() []ScriptInner {
	if o == nil || IsNil(o.Script) {
		var ret []ScriptInner
		return ret
	}
	return o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetScriptOk() ([]ScriptInner, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *SearchResponse) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given []ScriptInner and assigns it to the Script field.
func (o *SearchResponse) SetScript(v []ScriptInner) {
	o.Script = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *SearchResponse) GetUpdateTime() Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetUpdateTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *SearchResponse) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given Time and assigns it to the UpdateTime field.
func (o *SearchResponse) SetUpdateTime(v Time) {
	o.UpdateTime = &v
}

// GetLastUsageDate returns the LastUsageDate field value if set, zero value otherwise.
func (o *SearchResponse) GetLastUsageDate() Time {
	if o == nil || IsNil(o.LastUsageDate) {
		var ret Time
		return ret
	}
	return *o.LastUsageDate
}

// GetLastUsageDateOk returns a tuple with the LastUsageDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetLastUsageDateOk() (*Time, bool) {
	if o == nil || IsNil(o.LastUsageDate) {
		return nil, false
	}
	return o.LastUsageDate, true
}

// HasLastUsageDate returns a boolean if a field has been set.
func (o *SearchResponse) HasLastUsageDate() bool {
	if o != nil && !IsNil(o.LastUsageDate) {
		return true
	}

	return false
}

// SetLastUsageDate gets a reference to the given Time and assigns it to the LastUsageDate field.
func (o *SearchResponse) SetLastUsageDate(v Time) {
	o.LastUsageDate = &v
}

func (o SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.LastUsageDate) {
		toSerialize["lastUsageDate"] = o.LastUsageDate
	}
	return toSerialize, nil
}

type NullableSearchResponse struct {
	value *SearchResponse
	isSet bool
}

func (v NullableSearchResponse) Get() *SearchResponse {
	return v.value
}

func (v *NullableSearchResponse) Set(val *SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponse(val *SearchResponse) *NullableSearchResponse {
	return &NullableSearchResponse{value: val, isSet: true}
}

func (v NullableSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
