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

// checks if the BulkCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkCall{}

// BulkCall Bulk call list.
type BulkCall struct {
	// Calls Configuration ID.
	CallsConfigurationId *string
	Platform             *Platform
	// Unique call ID.
	CallId *string
	// Client-defined call ID.
	ExternalId *string
	// Caller identifier.
	From     *string
	Endpoint *BulkEndpoint
	Status   *ActionStatus
	// Failure reason in human-readable format.
	Reason *string
}

// NewBulkCall instantiates a new BulkCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewBulkCall() *BulkCall {
	this := BulkCall{}
	return &this
}

// NewBulkCallWithDefaults instantiates a new BulkCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCallWithDefaults() *BulkCall {
	this := BulkCall{}

	return &this
}

// GetCallsConfigurationId returns the CallsConfigurationId field value if set, zero value otherwise.
func (o *BulkCall) GetCallsConfigurationId() string {
	if o == nil || IsNil(o.CallsConfigurationId) {
		var ret string
		return ret
	}
	return *o.CallsConfigurationId
}

// GetCallsConfigurationIdOk returns a tuple with the CallsConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetCallsConfigurationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallsConfigurationId) {
		return nil, false
	}
	return o.CallsConfigurationId, true
}

// HasCallsConfigurationId returns a boolean if a field has been set.
func (o *BulkCall) HasCallsConfigurationId() bool {
	if o != nil && !IsNil(o.CallsConfigurationId) {
		return true
	}

	return false
}

// SetCallsConfigurationId gets a reference to the given string and assigns it to the CallsConfigurationId field.
func (o *BulkCall) SetCallsConfigurationId(v string) {
	o.CallsConfigurationId = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *BulkCall) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *BulkCall) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *BulkCall) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *BulkCall) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *BulkCall) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *BulkCall) SetCallId(v string) {
	o.CallId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *BulkCall) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *BulkCall) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *BulkCall) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BulkCall) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BulkCall) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *BulkCall) SetFrom(v string) {
	o.From = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *BulkCall) GetEndpoint() BulkEndpoint {
	if o == nil || IsNil(o.Endpoint) {
		var ret BulkEndpoint
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetEndpointOk() (*BulkEndpoint, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *BulkCall) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given BulkEndpoint and assigns it to the Endpoint field.
func (o *BulkCall) SetEndpoint(v BulkEndpoint) {
	o.Endpoint = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkCall) GetStatus() ActionStatus {
	if o == nil || IsNil(o.Status) {
		var ret ActionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetStatusOk() (*ActionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkCall) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ActionStatus and assigns it to the Status field.
func (o *BulkCall) SetStatus(v ActionStatus) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BulkCall) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCall) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BulkCall) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BulkCall) SetReason(v string) {
	o.Reason = &v
}

func (o BulkCall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallsConfigurationId) {
		toSerialize["callsConfigurationId"] = o.CallsConfigurationId
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.CallId) {
		toSerialize["callId"] = o.CallId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableBulkCall struct {
	value *BulkCall
	isSet bool
}

func (v NullableBulkCall) Get() *BulkCall {
	return v.value
}

func (v *NullableBulkCall) Set(val *BulkCall) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCall) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCall(val *BulkCall) *NullableBulkCall {
	return &NullableBulkCall{value: val, isSet: true}
}

func (v NullableBulkCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
