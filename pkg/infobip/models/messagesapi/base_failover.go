/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package messagesapi

import (
	"encoding/json"
	"fmt"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"gopkg.in/validator.v2"
)

// BaseFailover - Provides options for configuring a message failover. When message fails it will be sent over channels in order specified in an array. Make sure to provide correct sender and destinations specified as `Channels Destination` for each channel.
type BaseFailover struct {
	Failover         *Failover
	TemplateFailover *TemplateFailover
}

// FailoverAsBaseFailover is a convenience function that returns Failover wrapped in BaseFailover
func FailoverAsBaseFailover(v *Failover) BaseFailover {
	return BaseFailover{
		Failover: v,
	}
}

// TemplateFailoverAsBaseFailover is a convenience function that returns TemplateFailover wrapped in BaseFailover
func TemplateFailoverAsBaseFailover(v *TemplateFailover) BaseFailover {
	return BaseFailover{
		TemplateFailover: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BaseFailover) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Failover
	err = NewStrictDecoder(data).Decode(&dst.Failover)
	if err == nil {
		jsonFailover, _ := json.Marshal(dst.Failover)
		if string(jsonFailover) == "{}" { // empty struct
			dst.Failover = nil
		} else {
			if err = validator.Validate(dst.Failover); err != nil {
				dst.Failover = nil
			} else {
				match++
			}
		}
	} else {
		dst.Failover = nil
	}

	// try to unmarshal data into TemplateFailover
	err = NewStrictDecoder(data).Decode(&dst.TemplateFailover)
	if err == nil {
		jsonTemplateFailover, _ := json.Marshal(dst.TemplateFailover)
		if string(jsonTemplateFailover) == "{}" { // empty struct
			dst.TemplateFailover = nil
		} else {
			if err = validator.Validate(dst.TemplateFailover); err != nil {
				dst.TemplateFailover = nil
			} else {
				match++
			}
		}
	} else {
		dst.TemplateFailover = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Failover = nil
		dst.TemplateFailover = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BaseFailover)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BaseFailover)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BaseFailover) MarshalJSON() ([]byte, error) {
	if src.Failover != nil {
		return json.Marshal(&src.Failover)
	}

	if src.TemplateFailover != nil {
		return json.Marshal(&src.TemplateFailover)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BaseFailover) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Failover != nil {
		return obj.Failover
	}

	if obj.TemplateFailover != nil {
		return obj.TemplateFailover
	}

	// all schemas are nil
	return nil
}

type NullableBaseFailover struct {
	value *BaseFailover
	isSet bool
}

func (v NullableBaseFailover) Get() *BaseFailover {
	return v.value
}

func (v *NullableBaseFailover) Set(val *BaseFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFailover(val *BaseFailover) *NullableBaseFailover {
	return &NullableBaseFailover{value: val, isSet: true}
}

func (v NullableBaseFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
