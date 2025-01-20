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

// checks if the CallRoutingUrlDestinationHttpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRoutingUrlDestinationHttpRequest{}

// CallRoutingUrlDestinationHttpRequest struct for CallRoutingUrlDestinationHttpRequest
type CallRoutingUrlDestinationHttpRequest struct {
	// Identifier of the application that originated the call.
	ApplicationId *string
	// Identifier of the route that is used to process the call.
	RouteId *string
	// Identifier of the call that is being processed.
	CallId *string
	// Phone number from which the call originated from.
	From *string
	// Destination phone number of the call.
	To *string
	// Timestamp representing start time of the call.
	StartTime *Time
}

// NewCallRoutingUrlDestinationHttpRequest instantiates a new CallRoutingUrlDestinationHttpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewCallRoutingUrlDestinationHttpRequest() *CallRoutingUrlDestinationHttpRequest {
	this := CallRoutingUrlDestinationHttpRequest{}
	var applicationId string = "CALL_ROUTING"
	this.ApplicationId = &applicationId
	return &this
}

// NewCallRoutingUrlDestinationHttpRequestWithDefaults instantiates a new CallRoutingUrlDestinationHttpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRoutingUrlDestinationHttpRequestWithDefaults() *CallRoutingUrlDestinationHttpRequest {
	this := CallRoutingUrlDestinationHttpRequest{}

	var applicationId string = "CALL_ROUTING"
	this.ApplicationId = &applicationId
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *CallRoutingUrlDestinationHttpRequest) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingUrlDestinationHttpRequest) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *CallRoutingUrlDestinationHttpRequest) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *CallRoutingUrlDestinationHttpRequest) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *CallRoutingUrlDestinationHttpRequest) GetRouteId() string {
	if o == nil || IsNil(o.RouteId) {
		var ret string
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingUrlDestinationHttpRequest) GetRouteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RouteId) {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *CallRoutingUrlDestinationHttpRequest) HasRouteId() bool {
	if o != nil && !IsNil(o.RouteId) {
		return true
	}

	return false
}

// SetRouteId gets a reference to the given string and assigns it to the RouteId field.
func (o *CallRoutingUrlDestinationHttpRequest) SetRouteId(v string) {
	o.RouteId = &v
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *CallRoutingUrlDestinationHttpRequest) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingUrlDestinationHttpRequest) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *CallRoutingUrlDestinationHttpRequest) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *CallRoutingUrlDestinationHttpRequest) SetCallId(v string) {
	o.CallId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CallRoutingUrlDestinationHttpRequest) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingUrlDestinationHttpRequest) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CallRoutingUrlDestinationHttpRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CallRoutingUrlDestinationHttpRequest) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CallRoutingUrlDestinationHttpRequest) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingUrlDestinationHttpRequest) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CallRoutingUrlDestinationHttpRequest) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CallRoutingUrlDestinationHttpRequest) SetTo(v string) {
	o.To = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CallRoutingUrlDestinationHttpRequest) GetStartTime() Time {
	if o == nil || IsNil(o.StartTime) {
		var ret Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRoutingUrlDestinationHttpRequest) GetStartTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CallRoutingUrlDestinationHttpRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given Time and assigns it to the StartTime field.
func (o *CallRoutingUrlDestinationHttpRequest) SetStartTime(v Time) {
	o.StartTime = &v
}

func (o CallRoutingUrlDestinationHttpRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRoutingUrlDestinationHttpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.RouteId) {
		toSerialize["routeId"] = o.RouteId
	}
	if !IsNil(o.CallId) {
		toSerialize["callId"] = o.CallId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableCallRoutingUrlDestinationHttpRequest struct {
	value *CallRoutingUrlDestinationHttpRequest
	isSet bool
}

func (v NullableCallRoutingUrlDestinationHttpRequest) Get() *CallRoutingUrlDestinationHttpRequest {
	return v.value
}

func (v *NullableCallRoutingUrlDestinationHttpRequest) Set(val *CallRoutingUrlDestinationHttpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRoutingUrlDestinationHttpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRoutingUrlDestinationHttpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRoutingUrlDestinationHttpRequest(val *CallRoutingUrlDestinationHttpRequest) *NullableCallRoutingUrlDestinationHttpRequest {
	return &NullableCallRoutingUrlDestinationHttpRequest{value: val, isSet: true}
}

func (v NullableCallRoutingUrlDestinationHttpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRoutingUrlDestinationHttpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
