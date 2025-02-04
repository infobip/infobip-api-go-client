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

// checks if the StaticSipTrunkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticSipTrunkRequest{}

// StaticSipTrunkRequest struct for StaticSipTrunkRequest
type StaticSipTrunkRequest struct {
	Type                      SipTrunkType
	Name                      string
	Location                  *SipTrunkLocation
	Tls                       *bool
	InternationalCallsAllowed *bool
	ChannelLimit              int32
	BillingPackage            BillingPackage
	// List of audio codecs supported by a SIP trunk.
	Codecs        []AudioCodec
	Dtmf          *DtmfType
	Fax           *FaxType
	NumberFormat  *NumberPresentationFormat
	Anonymization *AnonymizationType
	// List of SIP trunk source hosts. If empty, destination host list must not be empty. Source hosts can be sent in 2 formats: IP address without port or domain without port.
	SourceHosts []string
	// List of SIP trunk destination hosts. If empty, source host list must not be empty. Destination hosts can be sent in 3 formats: IP address with port, domain name with port or domain name without port. The port must fall in the range 1025-65535 or be 0 for SRV lookup.
	DestinationHosts []string
	Strategy         *SelectionStrategy
	SipOptions       *SipOptions
}

type _StaticSipTrunkRequest StaticSipTrunkRequest

// NewStaticSipTrunkRequest instantiates a new StaticSipTrunkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticSipTrunkRequest() *StaticSipTrunkRequest {
	this := StaticSipTrunkRequest{}
	this.Type = "STATIC"
	var dtmf DtmfType = DTMFTYPE_RFC2833
	this.Dtmf = &dtmf
	var fax FaxType = FAXTYPE_NONE
	this.Fax = &fax
	var anonymization AnonymizationType = ANONYMIZATIONTYPE_NONE
	this.Anonymization = &anonymization
	var strategy SelectionStrategy = SELECTIONSTRATEGY_FAILOVER
	this.Strategy = &strategy
	return &this
}

// NewStaticSipTrunkRequestWithDefaults instantiates a new StaticSipTrunkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticSipTrunkRequestWithDefaults() *StaticSipTrunkRequest {
	this := StaticSipTrunkRequest{}
	this.Type = "STATIC"
	var dtmf DtmfType = DTMFTYPE_RFC2833
	this.Dtmf = &dtmf
	var fax FaxType = FAXTYPE_NONE
	this.Fax = &fax
	var anonymization AnonymizationType = ANONYMIZATIONTYPE_NONE
	this.Anonymization = &anonymization
	var strategy SelectionStrategy = SELECTIONSTRATEGY_FAILOVER
	this.Strategy = &strategy
	return &this
}

// GetName returns the Name field value
func (o *StaticSipTrunkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StaticSipTrunkRequest) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetLocation() SipTrunkLocation {
	if o == nil || IsNil(o.Location) {
		var ret SipTrunkLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetLocationOk() (*SipTrunkLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SipTrunkLocation and assigns it to the Location field.
func (o *StaticSipTrunkRequest) SetLocation(v SipTrunkLocation) {
	o.Location = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetTls() bool {
	if o == nil || IsNil(o.Tls) {
		var ret bool
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given bool and assigns it to the Tls field.
func (o *StaticSipTrunkRequest) SetTls(v bool) {
	o.Tls = &v
}

// GetInternationalCallsAllowed returns the InternationalCallsAllowed field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetInternationalCallsAllowed() bool {
	if o == nil || IsNil(o.InternationalCallsAllowed) {
		var ret bool
		return ret
	}
	return *o.InternationalCallsAllowed
}

// GetInternationalCallsAllowedOk returns a tuple with the InternationalCallsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetInternationalCallsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.InternationalCallsAllowed) {
		return nil, false
	}
	return o.InternationalCallsAllowed, true
}

// HasInternationalCallsAllowed returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasInternationalCallsAllowed() bool {
	if o != nil && !IsNil(o.InternationalCallsAllowed) {
		return true
	}

	return false
}

// SetInternationalCallsAllowed gets a reference to the given bool and assigns it to the InternationalCallsAllowed field.
func (o *StaticSipTrunkRequest) SetInternationalCallsAllowed(v bool) {
	o.InternationalCallsAllowed = &v
}

// GetChannelLimit returns the ChannelLimit field value
func (o *StaticSipTrunkRequest) GetChannelLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChannelLimit
}

// GetChannelLimitOk returns a tuple with the ChannelLimit field value
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetChannelLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelLimit, true
}

// SetChannelLimit sets field value
func (o *StaticSipTrunkRequest) SetChannelLimit(v int32) {
	o.ChannelLimit = v
}

// GetBillingPackage returns the BillingPackage field value
func (o *StaticSipTrunkRequest) GetBillingPackage() BillingPackage {
	if o == nil {
		var ret BillingPackage
		return ret
	}

	return o.BillingPackage
}

// GetBillingPackageOk returns a tuple with the BillingPackage field value
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetBillingPackageOk() (*BillingPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingPackage, true
}

// SetBillingPackage sets field value
func (o *StaticSipTrunkRequest) SetBillingPackage(v BillingPackage) {
	o.BillingPackage = v
}

// GetCodecs returns the Codecs field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetCodecs() []AudioCodec {
	if o == nil || IsNil(o.Codecs) {
		var ret []AudioCodec
		return ret
	}
	return o.Codecs
}

// GetCodecsOk returns a tuple with the Codecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetCodecsOk() ([]AudioCodec, bool) {
	if o == nil || IsNil(o.Codecs) {
		return nil, false
	}
	return o.Codecs, true
}

// HasCodecs returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasCodecs() bool {
	if o != nil && !IsNil(o.Codecs) {
		return true
	}

	return false
}

// SetCodecs gets a reference to the given []AudioCodec and assigns it to the Codecs field.
func (o *StaticSipTrunkRequest) SetCodecs(v []AudioCodec) {
	o.Codecs = v
}

// GetDtmf returns the Dtmf field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetDtmf() DtmfType {
	if o == nil || IsNil(o.Dtmf) {
		var ret DtmfType
		return ret
	}
	return *o.Dtmf
}

// GetDtmfOk returns a tuple with the Dtmf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetDtmfOk() (*DtmfType, bool) {
	if o == nil || IsNil(o.Dtmf) {
		return nil, false
	}
	return o.Dtmf, true
}

// HasDtmf returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasDtmf() bool {
	if o != nil && !IsNil(o.Dtmf) {
		return true
	}

	return false
}

// SetDtmf gets a reference to the given DtmfType and assigns it to the Dtmf field.
func (o *StaticSipTrunkRequest) SetDtmf(v DtmfType) {
	o.Dtmf = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetFax() FaxType {
	if o == nil || IsNil(o.Fax) {
		var ret FaxType
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetFaxOk() (*FaxType, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given FaxType and assigns it to the Fax field.
func (o *StaticSipTrunkRequest) SetFax(v FaxType) {
	o.Fax = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetNumberFormat() NumberPresentationFormat {
	if o == nil || IsNil(o.NumberFormat) {
		var ret NumberPresentationFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetNumberFormatOk() (*NumberPresentationFormat, bool) {
	if o == nil || IsNil(o.NumberFormat) {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasNumberFormat() bool {
	if o != nil && !IsNil(o.NumberFormat) {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given NumberPresentationFormat and assigns it to the NumberFormat field.
func (o *StaticSipTrunkRequest) SetNumberFormat(v NumberPresentationFormat) {
	o.NumberFormat = &v
}

// GetAnonymization returns the Anonymization field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetAnonymization() AnonymizationType {
	if o == nil || IsNil(o.Anonymization) {
		var ret AnonymizationType
		return ret
	}
	return *o.Anonymization
}

// GetAnonymizationOk returns a tuple with the Anonymization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetAnonymizationOk() (*AnonymizationType, bool) {
	if o == nil || IsNil(o.Anonymization) {
		return nil, false
	}
	return o.Anonymization, true
}

// HasAnonymization returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasAnonymization() bool {
	if o != nil && !IsNil(o.Anonymization) {
		return true
	}

	return false
}

// SetAnonymization gets a reference to the given AnonymizationType and assigns it to the Anonymization field.
func (o *StaticSipTrunkRequest) SetAnonymization(v AnonymizationType) {
	o.Anonymization = &v
}

// GetSourceHosts returns the SourceHosts field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetSourceHosts() []string {
	if o == nil || IsNil(o.SourceHosts) {
		var ret []string
		return ret
	}
	return o.SourceHosts
}

// GetSourceHostsOk returns a tuple with the SourceHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetSourceHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceHosts) {
		return nil, false
	}
	return o.SourceHosts, true
}

// HasSourceHosts returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasSourceHosts() bool {
	if o != nil && !IsNil(o.SourceHosts) {
		return true
	}

	return false
}

// SetSourceHosts gets a reference to the given []string and assigns it to the SourceHosts field.
func (o *StaticSipTrunkRequest) SetSourceHosts(v []string) {
	o.SourceHosts = v
}

// GetDestinationHosts returns the DestinationHosts field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetDestinationHosts() []string {
	if o == nil || IsNil(o.DestinationHosts) {
		var ret []string
		return ret
	}
	return o.DestinationHosts
}

// GetDestinationHostsOk returns a tuple with the DestinationHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetDestinationHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.DestinationHosts) {
		return nil, false
	}
	return o.DestinationHosts, true
}

// HasDestinationHosts returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasDestinationHosts() bool {
	if o != nil && !IsNil(o.DestinationHosts) {
		return true
	}

	return false
}

// SetDestinationHosts gets a reference to the given []string and assigns it to the DestinationHosts field.
func (o *StaticSipTrunkRequest) SetDestinationHosts(v []string) {
	o.DestinationHosts = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetStrategy() SelectionStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret SelectionStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetStrategyOk() (*SelectionStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given SelectionStrategy and assigns it to the Strategy field.
func (o *StaticSipTrunkRequest) SetStrategy(v SelectionStrategy) {
	o.Strategy = &v
}

// GetSipOptions returns the SipOptions field value if set, zero value otherwise.
func (o *StaticSipTrunkRequest) GetSipOptions() SipOptions {
	if o == nil || IsNil(o.SipOptions) {
		var ret SipOptions
		return ret
	}
	return *o.SipOptions
}

// GetSipOptionsOk returns a tuple with the SipOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticSipTrunkRequest) GetSipOptionsOk() (*SipOptions, bool) {
	if o == nil || IsNil(o.SipOptions) {
		return nil, false
	}
	return o.SipOptions, true
}

// HasSipOptions returns a boolean if a field has been set.
func (o *StaticSipTrunkRequest) HasSipOptions() bool {
	if o != nil && !IsNil(o.SipOptions) {
		return true
	}

	return false
}

// SetSipOptions gets a reference to the given SipOptions and assigns it to the SipOptions field.
func (o *StaticSipTrunkRequest) SetSipOptions(v SipOptions) {
	o.SipOptions = &v
}

func (o StaticSipTrunkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticSipTrunkRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Anonymization) {
		toSerialize["anonymization"] = o.Anonymization
	}
	if !IsNil(o.SourceHosts) {
		toSerialize["sourceHosts"] = o.SourceHosts
	}
	if !IsNil(o.DestinationHosts) {
		toSerialize["destinationHosts"] = o.DestinationHosts
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.SipOptions) {
		toSerialize["sipOptions"] = o.SipOptions
	}
	return toSerialize, nil
}

type NullableStaticSipTrunkRequest struct {
	value *StaticSipTrunkRequest
	isSet bool
}

func (v NullableStaticSipTrunkRequest) Get() *StaticSipTrunkRequest {
	return v.value
}

func (v *NullableStaticSipTrunkRequest) Set(val *StaticSipTrunkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticSipTrunkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticSipTrunkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticSipTrunkRequest(val *StaticSipTrunkRequest) *NullableStaticSipTrunkRequest {
	return &NullableStaticSipTrunkRequest{value: val, isSet: true}
}

func (v NullableStaticSipTrunkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticSipTrunkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
