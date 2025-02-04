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

// checks if the DialogLogResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DialogLogResponse{}

// DialogLogResponse struct for DialogLogResponse
type DialogLogResponse struct {
	// Unique dialog ID.
	DialogId *string
	// Calls Configuration ID.
	CallsConfigurationId *string
	Platform             *Platform
	State                *DialogState
	// Date and time for when the dialog has been created.
	StartTime *Time
	// Date and time for when the dialog has been established.
	EstablishTime *Time
	// Date and time for when the dialog has been finished.
	EndTime *Time
	// Unique parent call ID.
	ParentCallId *string
	// Unique child call ID.
	ChildCallId *string
	// Dialog duration in seconds.
	Duration  *int64
	Recording *DialogRecordingLog
	ErrorCode *ErrorCodeInfo
}

// NewDialogLogResponse instantiates a new DialogLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDialogLogResponse() *DialogLogResponse {
	this := DialogLogResponse{}
	return &this
}

// NewDialogLogResponseWithDefaults instantiates a new DialogLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDialogLogResponseWithDefaults() *DialogLogResponse {
	this := DialogLogResponse{}

	return &this
}

// GetDialogId returns the DialogId field value if set, zero value otherwise.
func (o *DialogLogResponse) GetDialogId() string {
	if o == nil || IsNil(o.DialogId) {
		var ret string
		return ret
	}
	return *o.DialogId
}

// GetDialogIdOk returns a tuple with the DialogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetDialogIdOk() (*string, bool) {
	if o == nil || IsNil(o.DialogId) {
		return nil, false
	}
	return o.DialogId, true
}

// HasDialogId returns a boolean if a field has been set.
func (o *DialogLogResponse) HasDialogId() bool {
	if o != nil && !IsNil(o.DialogId) {
		return true
	}

	return false
}

// SetDialogId gets a reference to the given string and assigns it to the DialogId field.
func (o *DialogLogResponse) SetDialogId(v string) {
	o.DialogId = &v
}

// GetCallsConfigurationId returns the CallsConfigurationId field value if set, zero value otherwise.
func (o *DialogLogResponse) GetCallsConfigurationId() string {
	if o == nil || IsNil(o.CallsConfigurationId) {
		var ret string
		return ret
	}
	return *o.CallsConfigurationId
}

// GetCallsConfigurationIdOk returns a tuple with the CallsConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetCallsConfigurationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallsConfigurationId) {
		return nil, false
	}
	return o.CallsConfigurationId, true
}

// HasCallsConfigurationId returns a boolean if a field has been set.
func (o *DialogLogResponse) HasCallsConfigurationId() bool {
	if o != nil && !IsNil(o.CallsConfigurationId) {
		return true
	}

	return false
}

// SetCallsConfigurationId gets a reference to the given string and assigns it to the CallsConfigurationId field.
func (o *DialogLogResponse) SetCallsConfigurationId(v string) {
	o.CallsConfigurationId = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *DialogLogResponse) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DialogLogResponse) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *DialogLogResponse) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DialogLogResponse) GetState() DialogState {
	if o == nil || IsNil(o.State) {
		var ret DialogState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetStateOk() (*DialogState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DialogLogResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given DialogState and assigns it to the State field.
func (o *DialogLogResponse) SetState(v DialogState) {
	o.State = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DialogLogResponse) GetStartTime() Time {
	if o == nil || IsNil(o.StartTime) {
		var ret Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetStartTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DialogLogResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given Time and assigns it to the StartTime field.
func (o *DialogLogResponse) SetStartTime(v Time) {
	o.StartTime = &v
}

// GetEstablishTime returns the EstablishTime field value if set, zero value otherwise.
func (o *DialogLogResponse) GetEstablishTime() Time {
	if o == nil || IsNil(o.EstablishTime) {
		var ret Time
		return ret
	}
	return *o.EstablishTime
}

// GetEstablishTimeOk returns a tuple with the EstablishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetEstablishTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.EstablishTime) {
		return nil, false
	}
	return o.EstablishTime, true
}

// HasEstablishTime returns a boolean if a field has been set.
func (o *DialogLogResponse) HasEstablishTime() bool {
	if o != nil && !IsNil(o.EstablishTime) {
		return true
	}

	return false
}

// SetEstablishTime gets a reference to the given Time and assigns it to the EstablishTime field.
func (o *DialogLogResponse) SetEstablishTime(v Time) {
	o.EstablishTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DialogLogResponse) GetEndTime() Time {
	if o == nil || IsNil(o.EndTime) {
		var ret Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetEndTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DialogLogResponse) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given Time and assigns it to the EndTime field.
func (o *DialogLogResponse) SetEndTime(v Time) {
	o.EndTime = &v
}

// GetParentCallId returns the ParentCallId field value if set, zero value otherwise.
func (o *DialogLogResponse) GetParentCallId() string {
	if o == nil || IsNil(o.ParentCallId) {
		var ret string
		return ret
	}
	return *o.ParentCallId
}

// GetParentCallIdOk returns a tuple with the ParentCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetParentCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCallId) {
		return nil, false
	}
	return o.ParentCallId, true
}

// HasParentCallId returns a boolean if a field has been set.
func (o *DialogLogResponse) HasParentCallId() bool {
	if o != nil && !IsNil(o.ParentCallId) {
		return true
	}

	return false
}

// SetParentCallId gets a reference to the given string and assigns it to the ParentCallId field.
func (o *DialogLogResponse) SetParentCallId(v string) {
	o.ParentCallId = &v
}

// GetChildCallId returns the ChildCallId field value if set, zero value otherwise.
func (o *DialogLogResponse) GetChildCallId() string {
	if o == nil || IsNil(o.ChildCallId) {
		var ret string
		return ret
	}
	return *o.ChildCallId
}

// GetChildCallIdOk returns a tuple with the ChildCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetChildCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChildCallId) {
		return nil, false
	}
	return o.ChildCallId, true
}

// HasChildCallId returns a boolean if a field has been set.
func (o *DialogLogResponse) HasChildCallId() bool {
	if o != nil && !IsNil(o.ChildCallId) {
		return true
	}

	return false
}

// SetChildCallId gets a reference to the given string and assigns it to the ChildCallId field.
func (o *DialogLogResponse) SetChildCallId(v string) {
	o.ChildCallId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DialogLogResponse) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DialogLogResponse) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *DialogLogResponse) SetDuration(v int64) {
	o.Duration = &v
}

// GetRecording returns the Recording field value if set, zero value otherwise.
func (o *DialogLogResponse) GetRecording() DialogRecordingLog {
	if o == nil || IsNil(o.Recording) {
		var ret DialogRecordingLog
		return ret
	}
	return *o.Recording
}

// GetRecordingOk returns a tuple with the Recording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetRecordingOk() (*DialogRecordingLog, bool) {
	if o == nil || IsNil(o.Recording) {
		return nil, false
	}
	return o.Recording, true
}

// HasRecording returns a boolean if a field has been set.
func (o *DialogLogResponse) HasRecording() bool {
	if o != nil && !IsNil(o.Recording) {
		return true
	}

	return false
}

// SetRecording gets a reference to the given DialogRecordingLog and assigns it to the Recording field.
func (o *DialogLogResponse) SetRecording(v DialogRecordingLog) {
	o.Recording = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *DialogLogResponse) GetErrorCode() ErrorCodeInfo {
	if o == nil || IsNil(o.ErrorCode) {
		var ret ErrorCodeInfo
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DialogLogResponse) GetErrorCodeOk() (*ErrorCodeInfo, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *DialogLogResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given ErrorCodeInfo and assigns it to the ErrorCode field.
func (o *DialogLogResponse) SetErrorCode(v ErrorCodeInfo) {
	o.ErrorCode = &v
}

func (o DialogLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DialogLogResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DialogId) {
		toSerialize["dialogId"] = o.DialogId
	}
	if !IsNil(o.CallsConfigurationId) {
		toSerialize["callsConfigurationId"] = o.CallsConfigurationId
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EstablishTime) {
		toSerialize["establishTime"] = o.EstablishTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.ParentCallId) {
		toSerialize["parentCallId"] = o.ParentCallId
	}
	if !IsNil(o.ChildCallId) {
		toSerialize["childCallId"] = o.ChildCallId
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Recording) {
		toSerialize["recording"] = o.Recording
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableDialogLogResponse struct {
	value *DialogLogResponse
	isSet bool
}

func (v NullableDialogLogResponse) Get() *DialogLogResponse {
	return v.value
}

func (v *NullableDialogLogResponse) Set(val *DialogLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDialogLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDialogLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDialogLogResponse(val *DialogLogResponse) *NullableDialogLogResponse {
	return &NullableDialogLogResponse{value: val, isSet: true}
}

func (v NullableDialogLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDialogLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
