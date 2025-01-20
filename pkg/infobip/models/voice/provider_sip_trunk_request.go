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

// checks if the ProviderSipTrunkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderSipTrunkRequest{}

// ProviderSipTrunkRequest struct for ProviderSipTrunkRequest
type ProviderSipTrunkRequest struct {
	Type                      SipTrunkType
	Name                      string
	Location                  *SipTrunkLocation
	Tls                       *bool
	InternationalCallsAllowed *bool
	ChannelLimit              int32
	BillingPackage            BillingPackage
	Provider                  Provider
}

type _ProviderSipTrunkRequest ProviderSipTrunkRequest

// NewProviderSipTrunkRequest instantiates a new ProviderSipTrunkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderSipTrunkRequest(provider Provider) *ProviderSipTrunkRequest {
	this := ProviderSipTrunkRequest{}
	this.Type = "PROVIDER"
	this.Provider = provider
	return &this
}

// NewProviderSipTrunkRequestWithDefaults instantiates a new ProviderSipTrunkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderSipTrunkRequestWithDefaults() *ProviderSipTrunkRequest {
	this := ProviderSipTrunkRequest{}
	this.Type = "PROVIDER"
	return &this
}

// GetName returns the Name field value
func (o *ProviderSipTrunkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProviderSipTrunkRequest) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ProviderSipTrunkRequest) GetLocation() SipTrunkLocation {
	if o == nil || IsNil(o.Location) {
		var ret SipTrunkLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetLocationOk() (*SipTrunkLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ProviderSipTrunkRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SipTrunkLocation and assigns it to the Location field.
func (o *ProviderSipTrunkRequest) SetLocation(v SipTrunkLocation) {
	o.Location = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ProviderSipTrunkRequest) GetTls() bool {
	if o == nil || IsNil(o.Tls) {
		var ret bool
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ProviderSipTrunkRequest) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given bool and assigns it to the Tls field.
func (o *ProviderSipTrunkRequest) SetTls(v bool) {
	o.Tls = &v
}

// GetInternationalCallsAllowed returns the InternationalCallsAllowed field value if set, zero value otherwise.
func (o *ProviderSipTrunkRequest) GetInternationalCallsAllowed() bool {
	if o == nil || IsNil(o.InternationalCallsAllowed) {
		var ret bool
		return ret
	}
	return *o.InternationalCallsAllowed
}

// GetInternationalCallsAllowedOk returns a tuple with the InternationalCallsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetInternationalCallsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.InternationalCallsAllowed) {
		return nil, false
	}
	return o.InternationalCallsAllowed, true
}

// HasInternationalCallsAllowed returns a boolean if a field has been set.
func (o *ProviderSipTrunkRequest) HasInternationalCallsAllowed() bool {
	if o != nil && !IsNil(o.InternationalCallsAllowed) {
		return true
	}

	return false
}

// SetInternationalCallsAllowed gets a reference to the given bool and assigns it to the InternationalCallsAllowed field.
func (o *ProviderSipTrunkRequest) SetInternationalCallsAllowed(v bool) {
	o.InternationalCallsAllowed = &v
}

// GetChannelLimit returns the ChannelLimit field value
func (o *ProviderSipTrunkRequest) GetChannelLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChannelLimit
}

// GetChannelLimitOk returns a tuple with the ChannelLimit field value
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetChannelLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelLimit, true
}

// SetChannelLimit sets field value
func (o *ProviderSipTrunkRequest) SetChannelLimit(v int32) {
	o.ChannelLimit = v
}

// GetBillingPackage returns the BillingPackage field value
func (o *ProviderSipTrunkRequest) GetBillingPackage() BillingPackage {
	if o == nil {
		var ret BillingPackage
		return ret
	}

	return o.BillingPackage
}

// GetBillingPackageOk returns a tuple with the BillingPackage field value
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetBillingPackageOk() (*BillingPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingPackage, true
}

// SetBillingPackage sets field value
func (o *ProviderSipTrunkRequest) SetBillingPackage(v BillingPackage) {
	o.BillingPackage = v
}

// GetProvider returns the Provider field value
func (o *ProviderSipTrunkRequest) GetProvider() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkRequest) GetProviderOk() (*Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ProviderSipTrunkRequest) SetProvider(v Provider) {
	o.Provider = v
}

func (o ProviderSipTrunkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderSipTrunkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	if !IsNil(o.InternationalCallsAllowed) {
		toSerialize["internationalCallsAllowed"] = o.InternationalCallsAllowed
	}
	toSerialize["channelLimit"] = o.ChannelLimit
	toSerialize["billingPackage"] = o.BillingPackage
	toSerialize["provider"] = o.Provider
	return toSerialize, nil
}

type NullableProviderSipTrunkRequest struct {
	value *ProviderSipTrunkRequest
	isSet bool
}

func (v NullableProviderSipTrunkRequest) Get() *ProviderSipTrunkRequest {
	return v.value
}

func (v *NullableProviderSipTrunkRequest) Set(val *ProviderSipTrunkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderSipTrunkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderSipTrunkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderSipTrunkRequest(val *ProviderSipTrunkRequest) *NullableProviderSipTrunkRequest {
	return &NullableProviderSipTrunkRequest{value: val, isSet: true}
}

func (v NullableProviderSipTrunkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderSipTrunkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
