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

// checks if the ValidationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationResponse{}

// ValidationResponse struct for ValidationResponse
type ValidationResponse struct {
	// Email address of the recipient.
	To *string
	// Represents status of recipient email address.
	ValidMailbox *string
	// Represents syntax of recipient email address.
	ValidSyntax *bool
	// Denotes catch all status of recipient email address.
	CatchAll *bool
	// Suggests alternate addresses that maybe valid.
	DidYouMean *string
	Disposable *bool
	RoleBased  *bool
	// Reason is provided when validMailbox status is unknown. 1. INBOX_FULL - The user quota exceeded / The user inbox is full / The user doesn't accept any more requests.  2. UNEXPECTED_FAILURE - The mail Server returned a temporary error. 3. THROTTLED - The mail server is not allowing us momentarily because of too many requests. 4. TIMED_OUT - The Mail Server took a longer time to respond / there was a delay in the network. 5. TEMP_REJECTION - Mail server temporarily rejected. 6. UNABLE_TO_CONNECT - Unable to connect to the Mail Server.
	Reason *string
	// Is provided when validMailbox is 'unknown' or 'false' and lists reasons clarifying why validMailbox has that status.
	DetailedReasons *string
	// Returns one of the following values: 'High', 'Medium', 'Low' or 'Unknown'. High risk addresses have very high chances of bouncing (and potentially damaging the sender's reputation), whereas low risk addresses have very low chances of bouncing and damaging the sender's reputation.
	Risk *string
}

// NewValidationResponse instantiates a new ValidationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed

func NewValidationResponse() *ValidationResponse {
	this := ValidationResponse{}
	return &this
}

// NewValidationResponseWithDefaults instantiates a new ValidationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationResponseWithDefaults() *ValidationResponse {
	this := ValidationResponse{}

	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ValidationResponse) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ValidationResponse) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *ValidationResponse) SetTo(v string) {
	o.To = &v
}

// GetValidMailbox returns the ValidMailbox field value if set, zero value otherwise.
func (o *ValidationResponse) GetValidMailbox() string {
	if o == nil || IsNil(o.ValidMailbox) {
		var ret string
		return ret
	}
	return *o.ValidMailbox
}

// GetValidMailboxOk returns a tuple with the ValidMailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetValidMailboxOk() (*string, bool) {
	if o == nil || IsNil(o.ValidMailbox) {
		return nil, false
	}
	return o.ValidMailbox, true
}

// HasValidMailbox returns a boolean if a field has been set.
func (o *ValidationResponse) HasValidMailbox() bool {
	if o != nil && !IsNil(o.ValidMailbox) {
		return true
	}

	return false
}

// SetValidMailbox gets a reference to the given string and assigns it to the ValidMailbox field.
func (o *ValidationResponse) SetValidMailbox(v string) {
	o.ValidMailbox = &v
}

// GetValidSyntax returns the ValidSyntax field value if set, zero value otherwise.
func (o *ValidationResponse) GetValidSyntax() bool {
	if o == nil || IsNil(o.ValidSyntax) {
		var ret bool
		return ret
	}
	return *o.ValidSyntax
}

// GetValidSyntaxOk returns a tuple with the ValidSyntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetValidSyntaxOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidSyntax) {
		return nil, false
	}
	return o.ValidSyntax, true
}

// HasValidSyntax returns a boolean if a field has been set.
func (o *ValidationResponse) HasValidSyntax() bool {
	if o != nil && !IsNil(o.ValidSyntax) {
		return true
	}

	return false
}

// SetValidSyntax gets a reference to the given bool and assigns it to the ValidSyntax field.
func (o *ValidationResponse) SetValidSyntax(v bool) {
	o.ValidSyntax = &v
}

// GetCatchAll returns the CatchAll field value if set, zero value otherwise.
func (o *ValidationResponse) GetCatchAll() bool {
	if o == nil || IsNil(o.CatchAll) {
		var ret bool
		return ret
	}
	return *o.CatchAll
}

// GetCatchAllOk returns a tuple with the CatchAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetCatchAllOk() (*bool, bool) {
	if o == nil || IsNil(o.CatchAll) {
		return nil, false
	}
	return o.CatchAll, true
}

// HasCatchAll returns a boolean if a field has been set.
func (o *ValidationResponse) HasCatchAll() bool {
	if o != nil && !IsNil(o.CatchAll) {
		return true
	}

	return false
}

// SetCatchAll gets a reference to the given bool and assigns it to the CatchAll field.
func (o *ValidationResponse) SetCatchAll(v bool) {
	o.CatchAll = &v
}

// GetDidYouMean returns the DidYouMean field value if set, zero value otherwise.
func (o *ValidationResponse) GetDidYouMean() string {
	if o == nil || IsNil(o.DidYouMean) {
		var ret string
		return ret
	}
	return *o.DidYouMean
}

// GetDidYouMeanOk returns a tuple with the DidYouMean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetDidYouMeanOk() (*string, bool) {
	if o == nil || IsNil(o.DidYouMean) {
		return nil, false
	}
	return o.DidYouMean, true
}

// HasDidYouMean returns a boolean if a field has been set.
func (o *ValidationResponse) HasDidYouMean() bool {
	if o != nil && !IsNil(o.DidYouMean) {
		return true
	}

	return false
}

// SetDidYouMean gets a reference to the given string and assigns it to the DidYouMean field.
func (o *ValidationResponse) SetDidYouMean(v string) {
	o.DidYouMean = &v
}

// GetDisposable returns the Disposable field value if set, zero value otherwise.
func (o *ValidationResponse) GetDisposable() bool {
	if o == nil || IsNil(o.Disposable) {
		var ret bool
		return ret
	}
	return *o.Disposable
}

// GetDisposableOk returns a tuple with the Disposable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetDisposableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disposable) {
		return nil, false
	}
	return o.Disposable, true
}

// HasDisposable returns a boolean if a field has been set.
func (o *ValidationResponse) HasDisposable() bool {
	if o != nil && !IsNil(o.Disposable) {
		return true
	}

	return false
}

// SetDisposable gets a reference to the given bool and assigns it to the Disposable field.
func (o *ValidationResponse) SetDisposable(v bool) {
	o.Disposable = &v
}

// GetRoleBased returns the RoleBased field value if set, zero value otherwise.
func (o *ValidationResponse) GetRoleBased() bool {
	if o == nil || IsNil(o.RoleBased) {
		var ret bool
		return ret
	}
	return *o.RoleBased
}

// GetRoleBasedOk returns a tuple with the RoleBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetRoleBasedOk() (*bool, bool) {
	if o == nil || IsNil(o.RoleBased) {
		return nil, false
	}
	return o.RoleBased, true
}

// HasRoleBased returns a boolean if a field has been set.
func (o *ValidationResponse) HasRoleBased() bool {
	if o != nil && !IsNil(o.RoleBased) {
		return true
	}

	return false
}

// SetRoleBased gets a reference to the given bool and assigns it to the RoleBased field.
func (o *ValidationResponse) SetRoleBased(v bool) {
	o.RoleBased = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ValidationResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ValidationResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ValidationResponse) SetReason(v string) {
	o.Reason = &v
}

// GetDetailedReasons returns the DetailedReasons field value if set, zero value otherwise.
func (o *ValidationResponse) GetDetailedReasons() string {
	if o == nil || IsNil(o.DetailedReasons) {
		var ret string
		return ret
	}
	return *o.DetailedReasons
}

// GetDetailedReasonsOk returns a tuple with the DetailedReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetDetailedReasonsOk() (*string, bool) {
	if o == nil || IsNil(o.DetailedReasons) {
		return nil, false
	}
	return o.DetailedReasons, true
}

// HasDetailedReasons returns a boolean if a field has been set.
func (o *ValidationResponse) HasDetailedReasons() bool {
	if o != nil && !IsNil(o.DetailedReasons) {
		return true
	}

	return false
}

// SetDetailedReasons gets a reference to the given string and assigns it to the DetailedReasons field.
func (o *ValidationResponse) SetDetailedReasons(v string) {
	o.DetailedReasons = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *ValidationResponse) GetRisk() string {
	if o == nil || IsNil(o.Risk) {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResponse) GetRiskOk() (*string, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *ValidationResponse) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *ValidationResponse) SetRisk(v string) {
	o.Risk = &v
}

func (o ValidationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.ValidMailbox) {
		toSerialize["validMailbox"] = o.ValidMailbox
	}
	if !IsNil(o.ValidSyntax) {
		toSerialize["validSyntax"] = o.ValidSyntax
	}
	if !IsNil(o.CatchAll) {
		toSerialize["catchAll"] = o.CatchAll
	}
	if !IsNil(o.DidYouMean) {
		toSerialize["didYouMean"] = o.DidYouMean
	}
	if !IsNil(o.Disposable) {
		toSerialize["disposable"] = o.Disposable
	}
	if !IsNil(o.RoleBased) {
		toSerialize["roleBased"] = o.RoleBased
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.DetailedReasons) {
		toSerialize["detailedReasons"] = o.DetailedReasons
	}
	if !IsNil(o.Risk) {
		toSerialize["risk"] = o.Risk
	}
	return toSerialize, nil
}

type NullableValidationResponse struct {
	value *ValidationResponse
	isSet bool
}

func (v NullableValidationResponse) Get() *ValidationResponse {
	return v.value
}

func (v *NullableValidationResponse) Set(val *ValidationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationResponse(val *ValidationResponse) *NullableValidationResponse {
	return &NullableValidationResponse{value: val, isSet: true}
}

func (v NullableValidationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
