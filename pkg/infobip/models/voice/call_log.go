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

// checks if the CallLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallLog{}

// CallLog struct for CallLog
type CallLog struct {
	// Unique call ID.
	CallId   *string
	Endpoint CallEndpoint
	// Caller ID.
	From *string
	// Callee ID.
	To        *string
	Direction *CallDirection
	State     *CallState
	// Date and time for when the call has been created.
	StartTime *Time
	// Date and time for when the call has been answered.
	AnswerTime *Time
	// Date and time for when the call has been finished.
	EndTime *Time
	// Parent call ID.
	ParentCallId     *string
	MachineDetection *MachineDetectionProperties
	// Ringing duration in seconds.
	RingDuration *int64
	// IDs of the calls configurations used during the call.
	CallsConfigurationIds []string
	Platform              *Platform
	// IDs of the conferences where the call was a participant.
	ConferenceIds []string
	// Call duration in seconds.
	Duration *int64
	// Indicates if camera was enabled during the call.
	HasCameraVideo *bool
	// Indicates if screen sharing was enabled during the call.
	HasScreenshareVideo *bool
	ErrorCode           *ErrorCodeInfo
	// Custom data.
	CustomData *map[string]string
	// Dialog ID.
	DialogId *string
	// Sender.
	Sender       *string
	HangupSource *HangupSource
}

type _CallLog CallLog

// NewCallLog instantiates a new CallLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCallLog(endpoint CallEndpoint) *CallLog {
	this := CallLog{}
	this.Endpoint = endpoint
	return &this
}

// NewCallLogWithDefaults instantiates a new CallLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallLogWithDefaults() *CallLog {
	this := CallLog{}

	return &this
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *CallLog) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *CallLog) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *CallLog) SetCallId(v string) {
	o.CallId = &v
}

// GetEndpoint returns the Endpoint field value
func (o *CallLog) GetEndpoint() CallEndpoint {
	if o == nil {
		var ret CallEndpoint
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CallLog) GetEndpointOk() (*CallEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CallLog) SetEndpoint(v CallEndpoint) {
	o.Endpoint = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CallLog) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CallLog) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CallLog) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CallLog) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CallLog) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CallLog) SetTo(v string) {
	o.To = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *CallLog) GetDirection() CallDirection {
	if o == nil || IsNil(o.Direction) {
		var ret CallDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetDirectionOk() (*CallDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *CallLog) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given CallDirection and assigns it to the Direction field.
func (o *CallLog) SetDirection(v CallDirection) {
	o.Direction = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CallLog) GetState() CallState {
	if o == nil || IsNil(o.State) {
		var ret CallState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetStateOk() (*CallState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CallLog) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CallState and assigns it to the State field.
func (o *CallLog) SetState(v CallState) {
	o.State = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CallLog) GetStartTime() Time {
	if o == nil || IsNil(o.StartTime) {
		var ret Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetStartTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CallLog) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given Time and assigns it to the StartTime field.
func (o *CallLog) SetStartTime(v Time) {
	o.StartTime = &v
}

// GetAnswerTime returns the AnswerTime field value if set, zero value otherwise.
func (o *CallLog) GetAnswerTime() Time {
	if o == nil || IsNil(o.AnswerTime) {
		var ret Time
		return ret
	}
	return *o.AnswerTime
}

// GetAnswerTimeOk returns a tuple with the AnswerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetAnswerTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.AnswerTime) {
		return nil, false
	}
	return o.AnswerTime, true
}

// HasAnswerTime returns a boolean if a field has been set.
func (o *CallLog) HasAnswerTime() bool {
	if o != nil && !IsNil(o.AnswerTime) {
		return true
	}

	return false
}

// SetAnswerTime gets a reference to the given Time and assigns it to the AnswerTime field.
func (o *CallLog) SetAnswerTime(v Time) {
	o.AnswerTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CallLog) GetEndTime() Time {
	if o == nil || IsNil(o.EndTime) {
		var ret Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetEndTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CallLog) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given Time and assigns it to the EndTime field.
func (o *CallLog) SetEndTime(v Time) {
	o.EndTime = &v
}

// GetParentCallId returns the ParentCallId field value if set, zero value otherwise.
func (o *CallLog) GetParentCallId() string {
	if o == nil || IsNil(o.ParentCallId) {
		var ret string
		return ret
	}
	return *o.ParentCallId
}

// GetParentCallIdOk returns a tuple with the ParentCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetParentCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCallId) {
		return nil, false
	}
	return o.ParentCallId, true
}

// HasParentCallId returns a boolean if a field has been set.
func (o *CallLog) HasParentCallId() bool {
	if o != nil && !IsNil(o.ParentCallId) {
		return true
	}

	return false
}

// SetParentCallId gets a reference to the given string and assigns it to the ParentCallId field.
func (o *CallLog) SetParentCallId(v string) {
	o.ParentCallId = &v
}

// GetMachineDetection returns the MachineDetection field value if set, zero value otherwise.
func (o *CallLog) GetMachineDetection() MachineDetectionProperties {
	if o == nil || IsNil(o.MachineDetection) {
		var ret MachineDetectionProperties
		return ret
	}
	return *o.MachineDetection
}

// GetMachineDetectionOk returns a tuple with the MachineDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetMachineDetectionOk() (*MachineDetectionProperties, bool) {
	if o == nil || IsNil(o.MachineDetection) {
		return nil, false
	}
	return o.MachineDetection, true
}

// HasMachineDetection returns a boolean if a field has been set.
func (o *CallLog) HasMachineDetection() bool {
	if o != nil && !IsNil(o.MachineDetection) {
		return true
	}

	return false
}

// SetMachineDetection gets a reference to the given MachineDetectionProperties and assigns it to the MachineDetection field.
func (o *CallLog) SetMachineDetection(v MachineDetectionProperties) {
	o.MachineDetection = &v
}

// GetRingDuration returns the RingDuration field value if set, zero value otherwise.
func (o *CallLog) GetRingDuration() int64 {
	if o == nil || IsNil(o.RingDuration) {
		var ret int64
		return ret
	}
	return *o.RingDuration
}

// GetRingDurationOk returns a tuple with the RingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetRingDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.RingDuration) {
		return nil, false
	}
	return o.RingDuration, true
}

// HasRingDuration returns a boolean if a field has been set.
func (o *CallLog) HasRingDuration() bool {
	if o != nil && !IsNil(o.RingDuration) {
		return true
	}

	return false
}

// SetRingDuration gets a reference to the given int64 and assigns it to the RingDuration field.
func (o *CallLog) SetRingDuration(v int64) {
	o.RingDuration = &v
}

// GetCallsConfigurationIds returns the CallsConfigurationIds field value if set, zero value otherwise.
func (o *CallLog) GetCallsConfigurationIds() []string {
	if o == nil || IsNil(o.CallsConfigurationIds) {
		var ret []string
		return ret
	}
	return o.CallsConfigurationIds
}

// GetCallsConfigurationIdsOk returns a tuple with the CallsConfigurationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetCallsConfigurationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CallsConfigurationIds) {
		return nil, false
	}
	return o.CallsConfigurationIds, true
}

// HasCallsConfigurationIds returns a boolean if a field has been set.
func (o *CallLog) HasCallsConfigurationIds() bool {
	if o != nil && !IsNil(o.CallsConfigurationIds) {
		return true
	}

	return false
}

// SetCallsConfigurationIds gets a reference to the given []string and assigns it to the CallsConfigurationIds field.
func (o *CallLog) SetCallsConfigurationIds(v []string) {
	o.CallsConfigurationIds = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CallLog) GetPlatform() Platform {
	if o == nil || IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetPlatformOk() (*Platform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CallLog) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *CallLog) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetConferenceIds returns the ConferenceIds field value if set, zero value otherwise.
func (o *CallLog) GetConferenceIds() []string {
	if o == nil || IsNil(o.ConferenceIds) {
		var ret []string
		return ret
	}
	return o.ConferenceIds
}

// GetConferenceIdsOk returns a tuple with the ConferenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetConferenceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConferenceIds) {
		return nil, false
	}
	return o.ConferenceIds, true
}

// HasConferenceIds returns a boolean if a field has been set.
func (o *CallLog) HasConferenceIds() bool {
	if o != nil && !IsNil(o.ConferenceIds) {
		return true
	}

	return false
}

// SetConferenceIds gets a reference to the given []string and assigns it to the ConferenceIds field.
func (o *CallLog) SetConferenceIds(v []string) {
	o.ConferenceIds = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CallLog) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CallLog) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *CallLog) SetDuration(v int64) {
	o.Duration = &v
}

// GetHasCameraVideo returns the HasCameraVideo field value if set, zero value otherwise.
func (o *CallLog) GetHasCameraVideo() bool {
	if o == nil || IsNil(o.HasCameraVideo) {
		var ret bool
		return ret
	}
	return *o.HasCameraVideo
}

// GetHasCameraVideoOk returns a tuple with the HasCameraVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetHasCameraVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCameraVideo) {
		return nil, false
	}
	return o.HasCameraVideo, true
}

// HasHasCameraVideo returns a boolean if a field has been set.
func (o *CallLog) HasHasCameraVideo() bool {
	if o != nil && !IsNil(o.HasCameraVideo) {
		return true
	}

	return false
}

// SetHasCameraVideo gets a reference to the given bool and assigns it to the HasCameraVideo field.
func (o *CallLog) SetHasCameraVideo(v bool) {
	o.HasCameraVideo = &v
}

// GetHasScreenshareVideo returns the HasScreenshareVideo field value if set, zero value otherwise.
func (o *CallLog) GetHasScreenshareVideo() bool {
	if o == nil || IsNil(o.HasScreenshareVideo) {
		var ret bool
		return ret
	}
	return *o.HasScreenshareVideo
}

// GetHasScreenshareVideoOk returns a tuple with the HasScreenshareVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetHasScreenshareVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.HasScreenshareVideo) {
		return nil, false
	}
	return o.HasScreenshareVideo, true
}

// HasHasScreenshareVideo returns a boolean if a field has been set.
func (o *CallLog) HasHasScreenshareVideo() bool {
	if o != nil && !IsNil(o.HasScreenshareVideo) {
		return true
	}

	return false
}

// SetHasScreenshareVideo gets a reference to the given bool and assigns it to the HasScreenshareVideo field.
func (o *CallLog) SetHasScreenshareVideo(v bool) {
	o.HasScreenshareVideo = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CallLog) GetErrorCode() ErrorCodeInfo {
	if o == nil || IsNil(o.ErrorCode) {
		var ret ErrorCodeInfo
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetErrorCodeOk() (*ErrorCodeInfo, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CallLog) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given ErrorCodeInfo and assigns it to the ErrorCode field.
func (o *CallLog) SetErrorCode(v ErrorCodeInfo) {
	o.ErrorCode = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *CallLog) GetCustomData() map[string]string {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]string
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetCustomDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomData) {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *CallLog) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]string and assigns it to the CustomData field.
func (o *CallLog) SetCustomData(v map[string]string) {
	o.CustomData = &v
}

// GetDialogId returns the DialogId field value if set, zero value otherwise.
func (o *CallLog) GetDialogId() string {
	if o == nil || IsNil(o.DialogId) {
		var ret string
		return ret
	}
	return *o.DialogId
}

// GetDialogIdOk returns a tuple with the DialogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetDialogIdOk() (*string, bool) {
	if o == nil || IsNil(o.DialogId) {
		return nil, false
	}
	return o.DialogId, true
}

// HasDialogId returns a boolean if a field has been set.
func (o *CallLog) HasDialogId() bool {
	if o != nil && !IsNil(o.DialogId) {
		return true
	}

	return false
}

// SetDialogId gets a reference to the given string and assigns it to the DialogId field.
func (o *CallLog) SetDialogId(v string) {
	o.DialogId = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *CallLog) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *CallLog) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
func (o *CallLog) SetSender(v string) {
	o.Sender = &v
}

// GetHangupSource returns the HangupSource field value if set, zero value otherwise.
func (o *CallLog) GetHangupSource() HangupSource {
	if o == nil || IsNil(o.HangupSource) {
		var ret HangupSource
		return ret
	}
	return *o.HangupSource
}

// GetHangupSourceOk returns a tuple with the HangupSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLog) GetHangupSourceOk() (*HangupSource, bool) {
	if o == nil || IsNil(o.HangupSource) {
		return nil, false
	}
	return o.HangupSource, true
}

// HasHangupSource returns a boolean if a field has been set.
func (o *CallLog) HasHangupSource() bool {
	if o != nil && !IsNil(o.HangupSource) {
		return true
	}

	return false
}

// SetHangupSource gets a reference to the given HangupSource and assigns it to the HangupSource field.
func (o *CallLog) SetHangupSource(v HangupSource) {
	o.HangupSource = &v
}

func (o CallLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallId) {
		toSerialize["callId"] = o.CallId
	}
	toSerialize["endpoint"] = o.Endpoint
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.AnswerTime) {
		toSerialize["answerTime"] = o.AnswerTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.ParentCallId) {
		toSerialize["parentCallId"] = o.ParentCallId
	}
	if !IsNil(o.MachineDetection) {
		toSerialize["machineDetection"] = o.MachineDetection
	}
	if !IsNil(o.RingDuration) {
		toSerialize["ringDuration"] = o.RingDuration
	}
	if !IsNil(o.CallsConfigurationIds) {
		toSerialize["callsConfigurationIds"] = o.CallsConfigurationIds
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ConferenceIds) {
		toSerialize["conferenceIds"] = o.ConferenceIds
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.HasCameraVideo) {
		toSerialize["hasCameraVideo"] = o.HasCameraVideo
	}
	if !IsNil(o.HasScreenshareVideo) {
		toSerialize["hasScreenshareVideo"] = o.HasScreenshareVideo
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.DialogId) {
		toSerialize["dialogId"] = o.DialogId
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.HangupSource) {
		toSerialize["hangupSource"] = o.HangupSource
	}
	return toSerialize, nil
}

type NullableCallLog struct {
	value *CallLog
	isSet bool
}

func (v NullableCallLog) Get() *CallLog {
	return v.value
}

func (v *NullableCallLog) Set(val *CallLog) {
	v.value = val
	v.isSet = true
}

func (v NullableCallLog) IsSet() bool {
	return v.isSet
}

func (v *NullableCallLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallLog(val *CallLog) *NullableCallLog {
	return &NullableCallLog{value: val, isSet: true}
}

func (v NullableCallLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
