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

// checks if the ConferenceRecordingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceRecordingRequest{}

// ConferenceRecordingRequest struct for ConferenceRecordingRequest
type ConferenceRecordingRequest struct {
	RecordingType         RecordingType
	ConferenceComposition *ConferenceComposition
	// Custom data.
	CustomData *map[string]string
	// Custom name for the recording's zip file. Applicable only when SFTP server is enabled on [Voice settings page](https://portal.infobip.com/apps/voice/recording/settings). Using the same `filePrefix` will override the files on the SFTP server. For recording without composition, `callId` and `fileId` will be appended to the `filePrefix` value.
	FilePrefix *string
}

type _ConferenceRecordingRequest ConferenceRecordingRequest

// NewConferenceRecordingRequest instantiates a new ConferenceRecordingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewConferenceRecordingRequest(recordingType RecordingType) *ConferenceRecordingRequest {
	this := ConferenceRecordingRequest{}
	this.RecordingType = recordingType
	return &this
}

// NewConferenceRecordingRequestWithDefaults instantiates a new ConferenceRecordingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceRecordingRequestWithDefaults() *ConferenceRecordingRequest {
	this := ConferenceRecordingRequest{}

	return &this
}

// GetRecordingType returns the RecordingType field value
func (o *ConferenceRecordingRequest) GetRecordingType() RecordingType {
	if o == nil {
		var ret RecordingType
		return ret
	}

	return o.RecordingType
}

// GetRecordingTypeOk returns a tuple with the RecordingType field value
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingRequest) GetRecordingTypeOk() (*RecordingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordingType, true
}

// SetRecordingType sets field value
func (o *ConferenceRecordingRequest) SetRecordingType(v RecordingType) {
	o.RecordingType = v
}

// GetConferenceComposition returns the ConferenceComposition field value if set, zero value otherwise.
func (o *ConferenceRecordingRequest) GetConferenceComposition() ConferenceComposition {
	if o == nil || IsNil(o.ConferenceComposition) {
		var ret ConferenceComposition
		return ret
	}
	return *o.ConferenceComposition
}

// GetConferenceCompositionOk returns a tuple with the ConferenceComposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingRequest) GetConferenceCompositionOk() (*ConferenceComposition, bool) {
	if o == nil || IsNil(o.ConferenceComposition) {
		return nil, false
	}
	return o.ConferenceComposition, true
}

// HasConferenceComposition returns a boolean if a field has been set.
func (o *ConferenceRecordingRequest) HasConferenceComposition() bool {
	if o != nil && !IsNil(o.ConferenceComposition) {
		return true
	}

	return false
}

// SetConferenceComposition gets a reference to the given ConferenceComposition and assigns it to the ConferenceComposition field.
func (o *ConferenceRecordingRequest) SetConferenceComposition(v ConferenceComposition) {
	o.ConferenceComposition = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *ConferenceRecordingRequest) GetCustomData() map[string]string {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]string
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingRequest) GetCustomDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomData) {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *ConferenceRecordingRequest) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]string and assigns it to the CustomData field.
func (o *ConferenceRecordingRequest) SetCustomData(v map[string]string) {
	o.CustomData = &v
}

// GetFilePrefix returns the FilePrefix field value if set, zero value otherwise.
func (o *ConferenceRecordingRequest) GetFilePrefix() string {
	if o == nil || IsNil(o.FilePrefix) {
		var ret string
		return ret
	}
	return *o.FilePrefix
}

// GetFilePrefixOk returns a tuple with the FilePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRecordingRequest) GetFilePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.FilePrefix) {
		return nil, false
	}
	return o.FilePrefix, true
}

// HasFilePrefix returns a boolean if a field has been set.
func (o *ConferenceRecordingRequest) HasFilePrefix() bool {
	if o != nil && !IsNil(o.FilePrefix) {
		return true
	}

	return false
}

// SetFilePrefix gets a reference to the given string and assigns it to the FilePrefix field.
func (o *ConferenceRecordingRequest) SetFilePrefix(v string) {
	o.FilePrefix = &v
}

func (o ConferenceRecordingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceRecordingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recordingType"] = o.RecordingType
	if !IsNil(o.ConferenceComposition) {
		toSerialize["conferenceComposition"] = o.ConferenceComposition
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.FilePrefix) {
		toSerialize["filePrefix"] = o.FilePrefix
	}
	return toSerialize, nil
}

type NullableConferenceRecordingRequest struct {
	value *ConferenceRecordingRequest
	isSet bool
}

func (v NullableConferenceRecordingRequest) Get() *ConferenceRecordingRequest {
	return v.value
}

func (v *NullableConferenceRecordingRequest) Set(val *ConferenceRecordingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceRecordingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceRecordingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceRecordingRequest(val *ConferenceRecordingRequest) *NullableConferenceRecordingRequest {
	return &NullableConferenceRecordingRequest{value: val, isSet: true}
}

func (v NullableConferenceRecordingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceRecordingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
