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

// checks if the SipTrunkServiceAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SipTrunkServiceAddress{}

// SipTrunkServiceAddress struct for SipTrunkServiceAddress
type SipTrunkServiceAddress struct {
	// SIP trunk service address ID.
	Id *string
	// SIP trunk service address name.
	Name *string
	// SIP trunk service address street.
	Street *string
	// SIP trunk service address city.
	City *string
	// SIP trunk service address post code.
	PostCode *string
	// SIP trunk service address suite.
	Suite   *string
	Country *Country
	Region  *Region
}

// NewSipTrunkServiceAddress instantiates a new SipTrunkServiceAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewSipTrunkServiceAddress() *SipTrunkServiceAddress {
	this := SipTrunkServiceAddress{}
	return &this
}

// NewSipTrunkServiceAddressWithDefaults instantiates a new SipTrunkServiceAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSipTrunkServiceAddressWithDefaults() *SipTrunkServiceAddress {
	this := SipTrunkServiceAddress{}

	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SipTrunkServiceAddress) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SipTrunkServiceAddress) SetName(v string) {
	o.Name = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *SipTrunkServiceAddress) SetStreet(v string) {
	o.Street = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SipTrunkServiceAddress) SetCity(v string) {
	o.City = &v
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetPostCode() string {
	if o == nil || IsNil(o.PostCode) {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetPostCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostCode) {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasPostCode() bool {
	if o != nil && !IsNil(o.PostCode) {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *SipTrunkServiceAddress) SetPostCode(v string) {
	o.PostCode = &v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetSuite() string {
	if o == nil || IsNil(o.Suite) {
		var ret string
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetSuiteOk() (*string, bool) {
	if o == nil || IsNil(o.Suite) {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasSuite() bool {
	if o != nil && !IsNil(o.Suite) {
		return true
	}

	return false
}

// SetSuite gets a reference to the given string and assigns it to the Suite field.
func (o *SipTrunkServiceAddress) SetSuite(v string) {
	o.Suite = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetCountry() Country {
	if o == nil || IsNil(o.Country) {
		var ret Country
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetCountryOk() (*Country, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given Country and assigns it to the Country field.
func (o *SipTrunkServiceAddress) SetCountry(v Country) {
	o.Country = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SipTrunkServiceAddress) GetRegion() Region {
	if o == nil || IsNil(o.Region) {
		var ret Region
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipTrunkServiceAddress) GetRegionOk() (*Region, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SipTrunkServiceAddress) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given Region and assigns it to the Region field.
func (o *SipTrunkServiceAddress) SetRegion(v Region) {
	o.Region = &v
}

func (o SipTrunkServiceAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SipTrunkServiceAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.PostCode) {
		toSerialize["postCode"] = o.PostCode
	}
	if !IsNil(o.Suite) {
		toSerialize["suite"] = o.Suite
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableSipTrunkServiceAddress struct {
	value *SipTrunkServiceAddress
	isSet bool
}

func (v NullableSipTrunkServiceAddress) Get() *SipTrunkServiceAddress {
	return v.value
}

func (v *NullableSipTrunkServiceAddress) Set(val *SipTrunkServiceAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableSipTrunkServiceAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableSipTrunkServiceAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipTrunkServiceAddress(val *SipTrunkServiceAddress) *NullableSipTrunkServiceAddress {
	return &NullableSipTrunkServiceAddress{value: val, isSet: true}
}

func (v NullableSipTrunkServiceAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipTrunkServiceAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
