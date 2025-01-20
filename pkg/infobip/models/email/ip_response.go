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

// checks if the IpResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpResponse{}

// IpResponse Dedicated IPs assigned to the IP pool.
type IpResponse struct {
	// Dedicated IP identifier.
	Id string
	// Dedicated IP address.
	Ip string
}

type _IpResponse IpResponse

// NewIpResponse instantiates a new IpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewIpResponse(id string, ip string) *IpResponse {
	this := IpResponse{}
	this.Id = id
	this.Ip = ip
	return &this
}

// NewIpResponseWithDefaults instantiates a new IpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpResponseWithDefaults() *IpResponse {
	this := IpResponse{}

	return &this
}

// GetId returns the Id field value
func (o *IpResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IpResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IpResponse) SetId(v string) {
	o.Id = v
}

// GetIp returns the Ip field value
func (o *IpResponse) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *IpResponse) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *IpResponse) SetIp(v string) {
	o.Ip = v
}

func (o IpResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

type NullableIpResponse struct {
	value *IpResponse
	isSet bool
}

func (v NullableIpResponse) Get() *IpResponse {
	return v.value
}

func (v *NullableIpResponse) Set(val *IpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpResponse(val *IpResponse) *NullableIpResponse {
	return &NullableIpResponse{value: val, isSet: true}
}

func (v NullableIpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
