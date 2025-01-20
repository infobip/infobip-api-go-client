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

// checks if the ProviderSipTrunkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderSipTrunkResponse{}

// ProviderSipTrunkResponse struct for ProviderSipTrunkResponse
type ProviderSipTrunkResponse struct {
	Id                        *string
	Type                      SipTrunkType
	Name                      *string
	Location                  *SipTrunkLocation
	Tls                       *bool
	Codecs                    []AudioCodec
	Dtmf                      *DtmfType
	Fax                       *FaxType
	NumberFormat              *NumberPresentationFormat
	InternationalCallsAllowed *bool
	ChannelLimit              *int32
	Anonymization             *AnonymizationType
	BillingPackage            *BillingPackage
	SbcHosts                  *SbcHosts
	SipOptions                *SipOptions
	Provider                  *Provider
}

// NewProviderSipTrunkResponse instantiates a new ProviderSipTrunkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderSipTrunkResponse() *ProviderSipTrunkResponse {
	this := ProviderSipTrunkResponse{}
	this.Type = "PROVIDER"
	return &this
}

// NewProviderSipTrunkResponseWithDefaults instantiates a new ProviderSipTrunkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderSipTrunkResponseWithDefaults() *ProviderSipTrunkResponse {
	this := ProviderSipTrunkResponse{}
	this.Type = "PROVIDER"
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProviderSipTrunkResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProviderSipTrunkResponse) SetName(v string) {
	o.Name = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetLocation() SipTrunkLocation {
	if o == nil || IsNil(o.Location) {
		var ret SipTrunkLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetLocationOk() (*SipTrunkLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SipTrunkLocation and assigns it to the Location field.
func (o *ProviderSipTrunkResponse) SetLocation(v SipTrunkLocation) {
	o.Location = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetTls() bool {
	if o == nil || IsNil(o.Tls) {
		var ret bool
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given bool and assigns it to the Tls field.
func (o *ProviderSipTrunkResponse) SetTls(v bool) {
	o.Tls = &v
}

// GetCodecs returns the Codecs field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetCodecs() []AudioCodec {
	if o == nil || IsNil(o.Codecs) {
		var ret []AudioCodec
		return ret
	}
	return o.Codecs
}

// GetCodecsOk returns a tuple with the Codecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetCodecsOk() ([]AudioCodec, bool) {
	if o == nil || IsNil(o.Codecs) {
		return nil, false
	}
	return o.Codecs, true
}

// HasCodecs returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasCodecs() bool {
	if o != nil && !IsNil(o.Codecs) {
		return true
	}

	return false
}

// SetCodecs gets a reference to the given []AudioCodec and assigns it to the Codecs field.
func (o *ProviderSipTrunkResponse) SetCodecs(v []AudioCodec) {
	o.Codecs = v
}

// GetDtmf returns the Dtmf field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetDtmf() DtmfType {
	if o == nil || IsNil(o.Dtmf) {
		var ret DtmfType
		return ret
	}
	return *o.Dtmf
}

// GetDtmfOk returns a tuple with the Dtmf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetDtmfOk() (*DtmfType, bool) {
	if o == nil || IsNil(o.Dtmf) {
		return nil, false
	}
	return o.Dtmf, true
}

// HasDtmf returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasDtmf() bool {
	if o != nil && !IsNil(o.Dtmf) {
		return true
	}

	return false
}

// SetDtmf gets a reference to the given DtmfType and assigns it to the Dtmf field.
func (o *ProviderSipTrunkResponse) SetDtmf(v DtmfType) {
	o.Dtmf = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetFax() FaxType {
	if o == nil || IsNil(o.Fax) {
		var ret FaxType
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetFaxOk() (*FaxType, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given FaxType and assigns it to the Fax field.
func (o *ProviderSipTrunkResponse) SetFax(v FaxType) {
	o.Fax = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetNumberFormat() NumberPresentationFormat {
	if o == nil || IsNil(o.NumberFormat) {
		var ret NumberPresentationFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetNumberFormatOk() (*NumberPresentationFormat, bool) {
	if o == nil || IsNil(o.NumberFormat) {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasNumberFormat() bool {
	if o != nil && !IsNil(o.NumberFormat) {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given NumberPresentationFormat and assigns it to the NumberFormat field.
func (o *ProviderSipTrunkResponse) SetNumberFormat(v NumberPresentationFormat) {
	o.NumberFormat = &v
}

// GetInternationalCallsAllowed returns the InternationalCallsAllowed field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetInternationalCallsAllowed() bool {
	if o == nil || IsNil(o.InternationalCallsAllowed) {
		var ret bool
		return ret
	}
	return *o.InternationalCallsAllowed
}

// GetInternationalCallsAllowedOk returns a tuple with the InternationalCallsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetInternationalCallsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.InternationalCallsAllowed) {
		return nil, false
	}
	return o.InternationalCallsAllowed, true
}

// HasInternationalCallsAllowed returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasInternationalCallsAllowed() bool {
	if o != nil && !IsNil(o.InternationalCallsAllowed) {
		return true
	}

	return false
}

// SetInternationalCallsAllowed gets a reference to the given bool and assigns it to the InternationalCallsAllowed field.
func (o *ProviderSipTrunkResponse) SetInternationalCallsAllowed(v bool) {
	o.InternationalCallsAllowed = &v
}

// GetChannelLimit returns the ChannelLimit field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetChannelLimit() int32 {
	if o == nil || IsNil(o.ChannelLimit) {
		var ret int32
		return ret
	}
	return *o.ChannelLimit
}

// GetChannelLimitOk returns a tuple with the ChannelLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetChannelLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelLimit) {
		return nil, false
	}
	return o.ChannelLimit, true
}

// HasChannelLimit returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasChannelLimit() bool {
	if o != nil && !IsNil(o.ChannelLimit) {
		return true
	}

	return false
}

// SetChannelLimit gets a reference to the given int32 and assigns it to the ChannelLimit field.
func (o *ProviderSipTrunkResponse) SetChannelLimit(v int32) {
	o.ChannelLimit = &v
}

// GetAnonymization returns the Anonymization field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetAnonymization() AnonymizationType {
	if o == nil || IsNil(o.Anonymization) {
		var ret AnonymizationType
		return ret
	}
	return *o.Anonymization
}

// GetAnonymizationOk returns a tuple with the Anonymization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetAnonymizationOk() (*AnonymizationType, bool) {
	if o == nil || IsNil(o.Anonymization) {
		return nil, false
	}
	return o.Anonymization, true
}

// HasAnonymization returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasAnonymization() bool {
	if o != nil && !IsNil(o.Anonymization) {
		return true
	}

	return false
}

// SetAnonymization gets a reference to the given AnonymizationType and assigns it to the Anonymization field.
func (o *ProviderSipTrunkResponse) SetAnonymization(v AnonymizationType) {
	o.Anonymization = &v
}

// GetBillingPackage returns the BillingPackage field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetBillingPackage() BillingPackage {
	if o == nil || IsNil(o.BillingPackage) {
		var ret BillingPackage
		return ret
	}
	return *o.BillingPackage
}

// GetBillingPackageOk returns a tuple with the BillingPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetBillingPackageOk() (*BillingPackage, bool) {
	if o == nil || IsNil(o.BillingPackage) {
		return nil, false
	}
	return o.BillingPackage, true
}

// HasBillingPackage returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasBillingPackage() bool {
	if o != nil && !IsNil(o.BillingPackage) {
		return true
	}

	return false
}

// SetBillingPackage gets a reference to the given BillingPackage and assigns it to the BillingPackage field.
func (o *ProviderSipTrunkResponse) SetBillingPackage(v BillingPackage) {
	o.BillingPackage = &v
}

// GetSbcHosts returns the SbcHosts field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetSbcHosts() SbcHosts {
	if o == nil || IsNil(o.SbcHosts) {
		var ret SbcHosts
		return ret
	}
	return *o.SbcHosts
}

// GetSbcHostsOk returns a tuple with the SbcHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetSbcHostsOk() (*SbcHosts, bool) {
	if o == nil || IsNil(o.SbcHosts) {
		return nil, false
	}
	return o.SbcHosts, true
}

// HasSbcHosts returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasSbcHosts() bool {
	if o != nil && !IsNil(o.SbcHosts) {
		return true
	}

	return false
}

// SetSbcHosts gets a reference to the given SbcHosts and assigns it to the SbcHosts field.
func (o *ProviderSipTrunkResponse) SetSbcHosts(v SbcHosts) {
	o.SbcHosts = &v
}

// GetSipOptions returns the SipOptions field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetSipOptions() SipOptions {
	if o == nil || IsNil(o.SipOptions) {
		var ret SipOptions
		return ret
	}
	return *o.SipOptions
}

// GetSipOptionsOk returns a tuple with the SipOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetSipOptionsOk() (*SipOptions, bool) {
	if o == nil || IsNil(o.SipOptions) {
		return nil, false
	}
	return o.SipOptions, true
}

// HasSipOptions returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasSipOptions() bool {
	if o != nil && !IsNil(o.SipOptions) {
		return true
	}

	return false
}

// SetSipOptions gets a reference to the given SipOptions and assigns it to the SipOptions field.
func (o *ProviderSipTrunkResponse) SetSipOptions(v SipOptions) {
	o.SipOptions = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ProviderSipTrunkResponse) GetProvider() Provider {
	if o == nil || IsNil(o.Provider) {
		var ret Provider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderSipTrunkResponse) GetProviderOk() (*Provider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ProviderSipTrunkResponse) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given Provider and assigns it to the Provider field.
func (o *ProviderSipTrunkResponse) SetProvider(v Provider) {
	o.Provider = &v
}

func (o ProviderSipTrunkResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderSipTrunkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	if !IsNil(o.Codecs) {
		toSerialize["codecs"] = o.Codecs
	}
	if !IsNil(o.Dtmf) {
		toSerialize["dtmf"] = o.Dtmf
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.NumberFormat) {
		toSerialize["numberFormat"] = o.NumberFormat
	}
	if !IsNil(o.InternationalCallsAllowed) {
		toSerialize["internationalCallsAllowed"] = o.InternationalCallsAllowed
	}
	if !IsNil(o.ChannelLimit) {
		toSerialize["channelLimit"] = o.ChannelLimit
	}
	if !IsNil(o.Anonymization) {
		toSerialize["anonymization"] = o.Anonymization
	}
	if !IsNil(o.BillingPackage) {
		toSerialize["billingPackage"] = o.BillingPackage
	}
	if !IsNil(o.SbcHosts) {
		toSerialize["sbcHosts"] = o.SbcHosts
	}
	if !IsNil(o.SipOptions) {
		toSerialize["sipOptions"] = o.SipOptions
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableProviderSipTrunkResponse struct {
	value *ProviderSipTrunkResponse
	isSet bool
}

func (v NullableProviderSipTrunkResponse) Get() *ProviderSipTrunkResponse {
	return v.value
}

func (v *NullableProviderSipTrunkResponse) Set(val *ProviderSipTrunkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderSipTrunkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderSipTrunkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderSipTrunkResponse(val *ProviderSipTrunkResponse) *NullableProviderSipTrunkResponse {
	return &NullableProviderSipTrunkResponse{value: val, isSet: true}
}

func (v NullableProviderSipTrunkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderSipTrunkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
