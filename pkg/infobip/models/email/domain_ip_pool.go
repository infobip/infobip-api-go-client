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

// checks if the DomainIpPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainIpPool{}

// DomainIpPool IP pools assigned to the domain. Assigned IP pools with the lowest priority value have the highest sending precedence.
type DomainIpPool struct {
	// IP pool identifier.
	Id string
	// IP pool name.
	Name string
	// IP pool sending priority. Higher value will result in a lower sending precedence for the specified IP pool.
	Priority int32
	Ips      []IpResponse
}

type _DomainIpPool DomainIpPool

// NewDomainIpPool instantiates a new DomainIpPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewDomainIpPool(id string, name string, priority int32, ips []IpResponse) *DomainIpPool {
	this := DomainIpPool{}
	this.Id = id
	this.Name = name
	this.Priority = priority
	this.Ips = ips
	return &this
}

// NewDomainIpPoolWithDefaults instantiates a new DomainIpPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainIpPoolWithDefaults() *DomainIpPool {
	this := DomainIpPool{}

	return &this
}

// GetId returns the Id field value
func (o *DomainIpPool) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DomainIpPool) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DomainIpPool) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DomainIpPool) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DomainIpPool) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DomainIpPool) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value
func (o *DomainIpPool) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *DomainIpPool) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *DomainIpPool) SetPriority(v int32) {
	o.Priority = v
}

// GetIps returns the Ips field value
func (o *DomainIpPool) GetIps() []IpResponse {
	if o == nil {
		var ret []IpResponse
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *DomainIpPool) GetIpsOk() ([]IpResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *DomainIpPool) SetIps(v []IpResponse) {
	o.Ips = v
}

func (o DomainIpPool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainIpPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	toSerialize["ips"] = o.Ips
	return toSerialize, nil
}

type NullableDomainIpPool struct {
	value *DomainIpPool
	isSet bool
}

func (v NullableDomainIpPool) Get() *DomainIpPool {
	return v.value
}

func (v *NullableDomainIpPool) Set(val *DomainIpPool) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainIpPool) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainIpPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainIpPool(val *DomainIpPool) *NullableDomainIpPool {
	return &NullableDomainIpPool{value: val, isSet: true}
}

func (v NullableDomainIpPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainIpPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
