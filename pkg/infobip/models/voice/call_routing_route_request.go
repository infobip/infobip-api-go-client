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

// checks if the CallRoutingRouteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRoutingRouteRequest{}

// CallRoutingRouteRequest Route request object.
type CallRoutingRouteRequest struct {
	// Route name.
	Name string
	// List of criteria used to match a route. For a route to match, any criterion should be met.
	Criteria []CallRoutingSearchCriteria
	// List of destinations. First destination in the list is the first one to be executed. Subsequent destinations are executed only if the previous one fails.
	Destinations []CallRoutingDestination
}

type _CallRoutingRouteRequest CallRoutingRouteRequest

// NewCallRoutingRouteRequest instantiates a new CallRoutingRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCallRoutingRouteRequest(name string, destinations []CallRoutingDestination) *CallRoutingRouteRequest {
	this := CallRoutingRouteRequest{}
	this.Name = name
	this.Destinations = destinations
	return &this
}

// NewCallRoutingRouteRequestWithDefaults instantiates a new CallRoutingRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRoutingRouteRequestWithDefaults() *CallRoutingRouteRequest {
	this := CallRoutingRouteRequest{}

	return &this
}

// GetName returns the Name field value
func (o *CallRoutingRouteRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CallRoutingRouteRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CallRoutingRouteRequest) SetName(v string) {
	o.Name = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *CallRoutingRouteRequest) GetCriteria() []CallRoutingSearchCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret []CallRoutingSearchCriteria
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingRouteRequest) GetCriteriaOk() ([]CallRoutingSearchCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *CallRoutingRouteRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []CallRoutingSearchCriteria and assigns it to the Criteria field.
func (o *CallRoutingRouteRequest) SetCriteria(v []CallRoutingSearchCriteria) {
	o.Criteria = v
}

// GetDestinations returns the Destinations field value
func (o *CallRoutingRouteRequest) GetDestinations() []CallRoutingDestination {
	if o == nil {
		var ret []CallRoutingDestination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *CallRoutingRouteRequest) GetDestinationsOk() ([]CallRoutingDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *CallRoutingRouteRequest) SetDestinations(v []CallRoutingDestination) {
	o.Destinations = v
}

func (o CallRoutingRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRoutingRouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	toSerialize["destinations"] = o.Destinations
	return toSerialize, nil
}

type NullableCallRoutingRouteRequest struct {
	value *CallRoutingRouteRequest
	isSet bool
}

func (v NullableCallRoutingRouteRequest) Get() *CallRoutingRouteRequest {
	return v.value
}

func (v *NullableCallRoutingRouteRequest) Set(val *CallRoutingRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRoutingRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRoutingRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRoutingRouteRequest(val *CallRoutingRouteRequest) *NullableCallRoutingRouteRequest {
	return &NullableCallRoutingRouteRequest{value: val, isSet: true}
}

func (v NullableCallRoutingRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRoutingRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
