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

// checks if the RecordingFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingFile{}

// RecordingFile struct for RecordingFile
type RecordingFile struct {
	// File ID.
	Id *string
	// File name.
	Name       string
	FileFormat FileFormat
	// File size in bytes.
	Size *int64
	// File creation time.
	CreationTime *Time
	// File duration in seconds.
	Duration *int64
	// Date and time when the recording started.
	StartTime *Time
	// Date and time when the recording ended.
	EndTime          *Time
	Location         *RecordingFileLocation
	SftpUploadStatus *SftpUploadStatus
	// Custom data.
	CustomData *map[string]string
}

type _RecordingFile RecordingFile

// NewRecordingFile instantiates a new RecordingFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewRecordingFile(name string, fileFormat FileFormat) *RecordingFile {
	this := RecordingFile{}
	this.Name = name
	this.FileFormat = fileFormat
	return &this
}

// NewRecordingFileWithDefaults instantiates a new RecordingFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingFileWithDefaults() *RecordingFile {
	this := RecordingFile{}

	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecordingFile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecordingFile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RecordingFile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RecordingFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RecordingFile) SetName(v string) {
	o.Name = v
}

// GetFileFormat returns the FileFormat field value
func (o *RecordingFile) GetFileFormat() FileFormat {
	if o == nil {
		var ret FileFormat
		return ret
	}

	return o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetFileFormatOk() (*FileFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileFormat, true
}

// SetFileFormat sets field value
func (o *RecordingFile) SetFileFormat(v FileFormat) {
	o.FileFormat = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RecordingFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RecordingFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *RecordingFile) SetSize(v int64) {
	o.Size = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *RecordingFile) GetCreationTime() Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetCreationTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *RecordingFile) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given Time and assigns it to the CreationTime field.
func (o *RecordingFile) SetCreationTime(v Time) {
	o.CreationTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *RecordingFile) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *RecordingFile) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *RecordingFile) SetDuration(v int64) {
	o.Duration = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RecordingFile) GetStartTime() Time {
	if o == nil || IsNil(o.StartTime) {
		var ret Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetStartTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RecordingFile) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given Time and assigns it to the StartTime field.
func (o *RecordingFile) SetStartTime(v Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RecordingFile) GetEndTime() Time {
	if o == nil || IsNil(o.EndTime) {
		var ret Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetEndTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RecordingFile) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given Time and assigns it to the EndTime field.
func (o *RecordingFile) SetEndTime(v Time) {
	o.EndTime = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *RecordingFile) GetLocation() RecordingFileLocation {
	if o == nil || IsNil(o.Location) {
		var ret RecordingFileLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetLocationOk() (*RecordingFileLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *RecordingFile) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given RecordingFileLocation and assigns it to the Location field.
func (o *RecordingFile) SetLocation(v RecordingFileLocation) {
	o.Location = &v
}

// GetSftpUploadStatus returns the SftpUploadStatus field value if set, zero value otherwise.
func (o *RecordingFile) GetSftpUploadStatus() SftpUploadStatus {
	if o == nil || IsNil(o.SftpUploadStatus) {
		var ret SftpUploadStatus
		return ret
	}
	return *o.SftpUploadStatus
}

// GetSftpUploadStatusOk returns a tuple with the SftpUploadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetSftpUploadStatusOk() (*SftpUploadStatus, bool) {
	if o == nil || IsNil(o.SftpUploadStatus) {
		return nil, false
	}
	return o.SftpUploadStatus, true
}

// HasSftpUploadStatus returns a boolean if a field has been set.
func (o *RecordingFile) HasSftpUploadStatus() bool {
	if o != nil && !IsNil(o.SftpUploadStatus) {
		return true
	}

	return false
}

// SetSftpUploadStatus gets a reference to the given SftpUploadStatus and assigns it to the SftpUploadStatus field.
func (o *RecordingFile) SetSftpUploadStatus(v SftpUploadStatus) {
	o.SftpUploadStatus = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *RecordingFile) GetCustomData() map[string]string {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]string
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFile) GetCustomDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomData) {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *RecordingFile) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]string and assigns it to the CustomData field.
func (o *RecordingFile) SetCustomData(v map[string]string) {
	o.CustomData = &v
}

func (o RecordingFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["fileFormat"] = o.FileFormat
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.SftpUploadStatus) {
		toSerialize["sftpUploadStatus"] = o.SftpUploadStatus
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	return toSerialize, nil
}

type NullableRecordingFile struct {
	value *RecordingFile
	isSet bool
}

func (v NullableRecordingFile) Get() *RecordingFile {
	return v.value
}

func (v *NullableRecordingFile) Set(val *RecordingFile) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingFile) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingFile(val *RecordingFile) *NullableRecordingFile {
	return &NullableRecordingFile{value: val, isSet: true}
}

func (v NullableRecordingFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
