/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the SuppressionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuppressionInfo{}

// SuppressionInfo Suppression get response.
type SuppressionInfo struct {
	// Name of the requested domain.
	DomainName string
	// Email address that is suppressed.
	EmailAddress string
	// Type of suppression.
	Type string
	// Date and time when email address was suppressed.
	CreatedDate Time
	// Reason for suppression.
	Reason string
}

type _SuppressionInfo SuppressionInfo

// NewSuppressionInfo instantiates a new SuppressionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSuppressionInfo(domainName string, emailAddress string, type_ string, createdDate Time, reason string) *SuppressionInfo {
	this := SuppressionInfo{}
	this.DomainName = domainName
	this.EmailAddress = emailAddress
	this.Type = type_
	this.CreatedDate = createdDate
	this.Reason = reason
	return &this
}

// NewSuppressionInfoWithDefaults instantiates a new SuppressionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuppressionInfoWithDefaults() *SuppressionInfo {
	this := SuppressionInfo{}

	return &this
}

// GetDomainName returns the DomainName field value
func (o *SuppressionInfo) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *SuppressionInfo) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *SuppressionInfo) SetDomainName(v string) {
	o.DomainName = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *SuppressionInfo) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *SuppressionInfo) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *SuppressionInfo) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetType returns the Type field value
func (o *SuppressionInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SuppressionInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SuppressionInfo) SetType(v string) {
	o.Type = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *SuppressionInfo) GetCreatedDate() Time {
	if o == nil {
		var ret Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *SuppressionInfo) GetCreatedDateOk() (*Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *SuppressionInfo) SetCreatedDate(v Time) {
	o.CreatedDate = v
}

// GetReason returns the Reason field value
func (o *SuppressionInfo) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SuppressionInfo) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SuppressionInfo) SetReason(v string) {
	o.Reason = v
}

func (o SuppressionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuppressionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domainName"] = o.DomainName
	toSerialize["emailAddress"] = o.EmailAddress
	toSerialize["type"] = o.Type
	toSerialize["createdDate"] = o.CreatedDate
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

type NullableSuppressionInfo struct {
	value *SuppressionInfo
	isSet bool
}

func (v NullableSuppressionInfo) Get() *SuppressionInfo {
	return v.value
}

func (v *NullableSuppressionInfo) Set(val *SuppressionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSuppressionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSuppressionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuppressionInfo(val *SuppressionInfo) *NullableSuppressionInfo {
	return &NullableSuppressionInfo{value: val, isSet: true}
}

func (v NullableSuppressionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuppressionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
