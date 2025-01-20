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

// checks if the SynthesisVoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SynthesisVoiceResponse{}

// SynthesisVoiceResponse Array of voices belonging to the specified language.
type SynthesisVoiceResponse struct {
	// Name of the voice. Example: `Joanna`
	Name *string
	// Gender of the voice. Can be `male` or `female`.
	Gender *string
	// Name of the supplier for text to speech synthesis.
	Supplier *string
	// Indicates if SSML is supported.
	SsmlSupported *bool
	// Indicates whether voice is default voice for a given language. If voice is not chosen for the language, then default voice will be used.
	IsDefault *bool
	// Indicates whether voice is neural. Using neural voice will generate additional cost.
	IsNeural *bool
}

// NewSynthesisVoiceResponse instantiates a new SynthesisVoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSynthesisVoiceResponse() *SynthesisVoiceResponse {
	this := SynthesisVoiceResponse{}
	return &this
}

// NewSynthesisVoiceResponseWithDefaults instantiates a new SynthesisVoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynthesisVoiceResponseWithDefaults() *SynthesisVoiceResponse {
	this := SynthesisVoiceResponse{}

	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SynthesisVoiceResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynthesisVoiceResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SynthesisVoiceResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SynthesisVoiceResponse) SetName(v string) {
	o.Name = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *SynthesisVoiceResponse) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynthesisVoiceResponse) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *SynthesisVoiceResponse) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *SynthesisVoiceResponse) SetGender(v string) {
	o.Gender = &v
}

// GetSupplier returns the Supplier field value if set, zero value otherwise.
func (o *SynthesisVoiceResponse) GetSupplier() string {
	if o == nil || IsNil(o.Supplier) {
		var ret string
		return ret
	}
	return *o.Supplier
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynthesisVoiceResponse) GetSupplierOk() (*string, bool) {
	if o == nil || IsNil(o.Supplier) {
		return nil, false
	}
	return o.Supplier, true
}

// HasSupplier returns a boolean if a field has been set.
func (o *SynthesisVoiceResponse) HasSupplier() bool {
	if o != nil && !IsNil(o.Supplier) {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given string and assigns it to the Supplier field.
func (o *SynthesisVoiceResponse) SetSupplier(v string) {
	o.Supplier = &v
}

// GetSsmlSupported returns the SsmlSupported field value if set, zero value otherwise.
func (o *SynthesisVoiceResponse) GetSsmlSupported() bool {
	if o == nil || IsNil(o.SsmlSupported) {
		var ret bool
		return ret
	}
	return *o.SsmlSupported
}

// GetSsmlSupportedOk returns a tuple with the SsmlSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynthesisVoiceResponse) GetSsmlSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.SsmlSupported) {
		return nil, false
	}
	return o.SsmlSupported, true
}

// HasSsmlSupported returns a boolean if a field has been set.
func (o *SynthesisVoiceResponse) HasSsmlSupported() bool {
	if o != nil && !IsNil(o.SsmlSupported) {
		return true
	}

	return false
}

// SetSsmlSupported gets a reference to the given bool and assigns it to the SsmlSupported field.
func (o *SynthesisVoiceResponse) SetSsmlSupported(v bool) {
	o.SsmlSupported = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *SynthesisVoiceResponse) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynthesisVoiceResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *SynthesisVoiceResponse) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *SynthesisVoiceResponse) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsNeural returns the IsNeural field value if set, zero value otherwise.
func (o *SynthesisVoiceResponse) GetIsNeural() bool {
	if o == nil || IsNil(o.IsNeural) {
		var ret bool
		return ret
	}
	return *o.IsNeural
}

// GetIsNeuralOk returns a tuple with the IsNeural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynthesisVoiceResponse) GetIsNeuralOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNeural) {
		return nil, false
	}
	return o.IsNeural, true
}

// HasIsNeural returns a boolean if a field has been set.
func (o *SynthesisVoiceResponse) HasIsNeural() bool {
	if o != nil && !IsNil(o.IsNeural) {
		return true
	}

	return false
}

// SetIsNeural gets a reference to the given bool and assigns it to the IsNeural field.
func (o *SynthesisVoiceResponse) SetIsNeural(v bool) {
	o.IsNeural = &v
}

func (o SynthesisVoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SynthesisVoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Supplier) {
		toSerialize["supplier"] = o.Supplier
	}
	if !IsNil(o.SsmlSupported) {
		toSerialize["ssmlSupported"] = o.SsmlSupported
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.IsNeural) {
		toSerialize["isNeural"] = o.IsNeural
	}
	return toSerialize, nil
}

type NullableSynthesisVoiceResponse struct {
	value *SynthesisVoiceResponse
	isSet bool
}

func (v NullableSynthesisVoiceResponse) Get() *SynthesisVoiceResponse {
	return v.value
}

func (v *NullableSynthesisVoiceResponse) Set(val *SynthesisVoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSynthesisVoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSynthesisVoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynthesisVoiceResponse(val *SynthesisVoiceResponse) *NullableSynthesisVoiceResponse {
	return &NullableSynthesisVoiceResponse{value: val, isSet: true}
}

func (v NullableSynthesisVoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynthesisVoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
