/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voice

import (
	"encoding/json"
	"fmt"

	"gopkg.in/validator.v2"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// ScriptInner - struct for ScriptInner
type ScriptInner struct {
	CallApi             *CallApi
	Capture             *Capture
	Collect             *Collect
	Dial                *Dial
	DialToConversations *DialToConversations
	DialToMany          *DialToMany
	DialToWebRTC        *DialToWebRTC
	ForEach             *ForEach
	GoTo                *GoTo
	Hangup              *Hangup
	IfThenElse          *IfThenElse
	MachineDetection    *MachineDetection
	Pause               *Pause
	Play                *Play
	PlayFromRecording   *PlayFromRecording
	Record              *Record
	RepeatUntil         *RepeatUntil
	RepeatWhile         *RepeatWhile
	Say                 *Say
	SendSms             *SendSms
	SetVariable         *SetVariable
	SwitchCase          *SwitchCase
	WhileDo             *WhileDo
}

// CallApiAsScriptInner is a convenience function that returns CallApi wrapped in ScriptInner
func CallApiAsScriptInner(v *CallApi) ScriptInner {
	return ScriptInner{
		CallApi: v,
	}
}

// CaptureAsScriptInner is a convenience function that returns Capture wrapped in ScriptInner
func CaptureAsScriptInner(v *Capture) ScriptInner {
	return ScriptInner{
		Capture: v,
	}
}

// CollectAsScriptInner is a convenience function that returns Collect wrapped in ScriptInner
func CollectAsScriptInner(v *Collect) ScriptInner {
	return ScriptInner{
		Collect: v,
	}
}

// DialAsScriptInner is a convenience function that returns Dial wrapped in ScriptInner
func DialAsScriptInner(v *Dial) ScriptInner {
	return ScriptInner{
		Dial: v,
	}
}

// DialToConversationsAsScriptInner is a convenience function that returns DialToConversations wrapped in ScriptInner
func DialToConversationsAsScriptInner(v *DialToConversations) ScriptInner {
	return ScriptInner{
		DialToConversations: v,
	}
}

// DialToManyAsScriptInner is a convenience function that returns DialToMany wrapped in ScriptInner
func DialToManyAsScriptInner(v *DialToMany) ScriptInner {
	return ScriptInner{
		DialToMany: v,
	}
}

// DialToWebRTCAsScriptInner is a convenience function that returns DialToWebRTC wrapped in ScriptInner
func DialToWebRTCAsScriptInner(v *DialToWebRTC) ScriptInner {
	return ScriptInner{
		DialToWebRTC: v,
	}
}

// ForEachAsScriptInner is a convenience function that returns ForEach wrapped in ScriptInner
func ForEachAsScriptInner(v *ForEach) ScriptInner {
	return ScriptInner{
		ForEach: v,
	}
}

// GoToAsScriptInner is a convenience function that returns GoTo wrapped in ScriptInner
func GoToAsScriptInner(v *GoTo) ScriptInner {
	return ScriptInner{
		GoTo: v,
	}
}

// HangupAsScriptInner is a convenience function that returns Hangup wrapped in ScriptInner
func HangupAsScriptInner(v *Hangup) ScriptInner {
	return ScriptInner{
		Hangup: v,
	}
}

// IfThenElseAsScriptInner is a convenience function that returns IfThenElse wrapped in ScriptInner
func IfThenElseAsScriptInner(v *IfThenElse) ScriptInner {
	return ScriptInner{
		IfThenElse: v,
	}
}

// MachineDetectionAsScriptInner is a convenience function that returns MachineDetection wrapped in ScriptInner
func MachineDetectionAsScriptInner(v *MachineDetection) ScriptInner {
	return ScriptInner{
		MachineDetection: v,
	}
}

// PauseAsScriptInner is a convenience function that returns Pause wrapped in ScriptInner
func PauseAsScriptInner(v *Pause) ScriptInner {
	return ScriptInner{
		Pause: v,
	}
}

// PlayAsScriptInner is a convenience function that returns Play wrapped in ScriptInner
func PlayAsScriptInner(v *Play) ScriptInner {
	return ScriptInner{
		Play: v,
	}
}

// PlayFromRecordingAsScriptInner is a convenience function that returns PlayFromRecording wrapped in ScriptInner
func PlayFromRecordingAsScriptInner(v *PlayFromRecording) ScriptInner {
	return ScriptInner{
		PlayFromRecording: v,
	}
}

// RecordAsScriptInner is a convenience function that returns Record wrapped in ScriptInner
func RecordAsScriptInner(v *Record) ScriptInner {
	return ScriptInner{
		Record: v,
	}
}

// RepeatUntilAsScriptInner is a convenience function that returns RepeatUntil wrapped in ScriptInner
func RepeatUntilAsScriptInner(v *RepeatUntil) ScriptInner {
	return ScriptInner{
		RepeatUntil: v,
	}
}

// RepeatWhileAsScriptInner is a convenience function that returns RepeatWhile wrapped in ScriptInner
func RepeatWhileAsScriptInner(v *RepeatWhile) ScriptInner {
	return ScriptInner{
		RepeatWhile: v,
	}
}

// SayAsScriptInner is a convenience function that returns Say wrapped in ScriptInner
func SayAsScriptInner(v *Say) ScriptInner {
	return ScriptInner{
		Say: v,
	}
}

// SendSmsAsScriptInner is a convenience function that returns SendSms wrapped in ScriptInner
func SendSmsAsScriptInner(v *SendSms) ScriptInner {
	return ScriptInner{
		SendSms: v,
	}
}

// SetVariableAsScriptInner is a convenience function that returns SetVariable wrapped in ScriptInner
func SetVariableAsScriptInner(v *SetVariable) ScriptInner {
	return ScriptInner{
		SetVariable: v,
	}
}

// SwitchCaseAsScriptInner is a convenience function that returns SwitchCase wrapped in ScriptInner
func SwitchCaseAsScriptInner(v *SwitchCase) ScriptInner {
	return ScriptInner{
		SwitchCase: v,
	}
}

// WhileDoAsScriptInner is a convenience function that returns WhileDo wrapped in ScriptInner
func WhileDoAsScriptInner(v *WhileDo) ScriptInner {
	return ScriptInner{
		WhileDo: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ScriptInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CallApi
	err = NewStrictDecoder(data).Decode(&dst.CallApi)
	if err == nil {
		jsonCallApi, _ := json.Marshal(dst.CallApi)
		if string(jsonCallApi) == "{}" { // empty struct
			dst.CallApi = nil
		} else {
			if err = validator.Validate(dst.CallApi); err != nil {
				dst.CallApi = nil
			} else {
				match++
			}
		}
	} else {
		dst.CallApi = nil
	}

	// try to unmarshal data into Capture
	err = NewStrictDecoder(data).Decode(&dst.Capture)
	if err == nil {
		jsonCapture, _ := json.Marshal(dst.Capture)
		if string(jsonCapture) == "{}" { // empty struct
			dst.Capture = nil
		} else {
			if err = validator.Validate(dst.Capture); err != nil {
				dst.Capture = nil
			} else {
				match++
			}
		}
	} else {
		dst.Capture = nil
	}

	// try to unmarshal data into Collect
	err = NewStrictDecoder(data).Decode(&dst.Collect)
	if err == nil {
		jsonCollect, _ := json.Marshal(dst.Collect)
		if string(jsonCollect) == "{}" { // empty struct
			dst.Collect = nil
		} else {
			if err = validator.Validate(dst.Collect); err != nil {
				dst.Collect = nil
			} else {
				match++
			}
		}
	} else {
		dst.Collect = nil
	}

	// try to unmarshal data into Dial
	err = NewStrictDecoder(data).Decode(&dst.Dial)
	if err == nil {
		jsonDial, _ := json.Marshal(dst.Dial)
		if string(jsonDial) == "{}" { // empty struct
			dst.Dial = nil
		} else {
			if err = validator.Validate(dst.Dial); err != nil {
				dst.Dial = nil
			} else {
				match++
			}
		}
	} else {
		dst.Dial = nil
	}

	// try to unmarshal data into DialToConversations
	err = NewStrictDecoder(data).Decode(&dst.DialToConversations)
	if err == nil {
		jsonDialToConversations, _ := json.Marshal(dst.DialToConversations)
		if string(jsonDialToConversations) == "{}" { // empty struct
			dst.DialToConversations = nil
		} else {
			if err = validator.Validate(dst.DialToConversations); err != nil {
				dst.DialToConversations = nil
			} else {
				match++
			}
		}
	} else {
		dst.DialToConversations = nil
	}

	// try to unmarshal data into DialToMany
	err = NewStrictDecoder(data).Decode(&dst.DialToMany)
	if err == nil {
		jsonDialToMany, _ := json.Marshal(dst.DialToMany)
		if string(jsonDialToMany) == "{}" { // empty struct
			dst.DialToMany = nil
		} else {
			if err = validator.Validate(dst.DialToMany); err != nil {
				dst.DialToMany = nil
			} else {
				match++
			}
		}
	} else {
		dst.DialToMany = nil
	}

	// try to unmarshal data into DialToWebRTC
	err = NewStrictDecoder(data).Decode(&dst.DialToWebRTC)
	if err == nil {
		jsonDialToWebRTC, _ := json.Marshal(dst.DialToWebRTC)
		if string(jsonDialToWebRTC) == "{}" { // empty struct
			dst.DialToWebRTC = nil
		} else {
			if err = validator.Validate(dst.DialToWebRTC); err != nil {
				dst.DialToWebRTC = nil
			} else {
				match++
			}
		}
	} else {
		dst.DialToWebRTC = nil
	}

	// try to unmarshal data into ForEach
	err = NewStrictDecoder(data).Decode(&dst.ForEach)
	if err == nil {
		jsonForEach, _ := json.Marshal(dst.ForEach)
		if string(jsonForEach) == "{}" { // empty struct
			dst.ForEach = nil
		} else {
			if err = validator.Validate(dst.ForEach); err != nil {
				dst.ForEach = nil
			} else {
				match++
			}
		}
	} else {
		dst.ForEach = nil
	}

	// try to unmarshal data into GoTo
	err = NewStrictDecoder(data).Decode(&dst.GoTo)
	if err == nil {
		jsonGoTo, _ := json.Marshal(dst.GoTo)
		if string(jsonGoTo) == "{}" { // empty struct
			dst.GoTo = nil
		} else {
			if err = validator.Validate(dst.GoTo); err != nil {
				dst.GoTo = nil
			} else {
				match++
			}
		}
	} else {
		dst.GoTo = nil
	}

	// try to unmarshal data into Hangup
	err = NewStrictDecoder(data).Decode(&dst.Hangup)
	if err == nil {
		jsonHangup, _ := json.Marshal(dst.Hangup)
		if string(jsonHangup) == "{}" { // empty struct
			dst.Hangup = nil
		} else {
			if err = validator.Validate(dst.Hangup); err != nil {
				dst.Hangup = nil
			} else {
				match++
			}
		}
	} else {
		dst.Hangup = nil
	}

	// try to unmarshal data into IfThenElse
	err = NewStrictDecoder(data).Decode(&dst.IfThenElse)
	if err == nil {
		jsonIfThenElse, _ := json.Marshal(dst.IfThenElse)
		if string(jsonIfThenElse) == "{}" { // empty struct
			dst.IfThenElse = nil
		} else {
			if err = validator.Validate(dst.IfThenElse); err != nil {
				dst.IfThenElse = nil
			} else {
				match++
			}
		}
	} else {
		dst.IfThenElse = nil
	}

	// try to unmarshal data into MachineDetection
	err = NewStrictDecoder(data).Decode(&dst.MachineDetection)
	if err == nil {
		jsonMachineDetection, _ := json.Marshal(dst.MachineDetection)
		if string(jsonMachineDetection) == "{}" { // empty struct
			dst.MachineDetection = nil
		} else {
			if err = validator.Validate(dst.MachineDetection); err != nil {
				dst.MachineDetection = nil
			} else {
				match++
			}
		}
	} else {
		dst.MachineDetection = nil
	}

	// try to unmarshal data into Pause
	err = NewStrictDecoder(data).Decode(&dst.Pause)
	if err == nil {
		jsonPause, _ := json.Marshal(dst.Pause)
		if string(jsonPause) == "{}" { // empty struct
			dst.Pause = nil
		} else {
			if err = validator.Validate(dst.Pause); err != nil {
				dst.Pause = nil
			} else {
				match++
			}
		}
	} else {
		dst.Pause = nil
	}

	// try to unmarshal data into Play
	err = NewStrictDecoder(data).Decode(&dst.Play)
	if err == nil {
		jsonPlay, _ := json.Marshal(dst.Play)
		if string(jsonPlay) == "{}" { // empty struct
			dst.Play = nil
		} else {
			if err = validator.Validate(dst.Play); err != nil {
				dst.Play = nil
			} else {
				match++
			}
		}
	} else {
		dst.Play = nil
	}

	// try to unmarshal data into PlayFromRecording
	err = NewStrictDecoder(data).Decode(&dst.PlayFromRecording)
	if err == nil {
		jsonPlayFromRecording, _ := json.Marshal(dst.PlayFromRecording)
		if string(jsonPlayFromRecording) == "{}" { // empty struct
			dst.PlayFromRecording = nil
		} else {
			if err = validator.Validate(dst.PlayFromRecording); err != nil {
				dst.PlayFromRecording = nil
			} else {
				match++
			}
		}
	} else {
		dst.PlayFromRecording = nil
	}

	// try to unmarshal data into Record
	err = NewStrictDecoder(data).Decode(&dst.Record)
	if err == nil {
		jsonRecord, _ := json.Marshal(dst.Record)
		if string(jsonRecord) == "{}" { // empty struct
			dst.Record = nil
		} else {
			if err = validator.Validate(dst.Record); err != nil {
				dst.Record = nil
			} else {
				match++
			}
		}
	} else {
		dst.Record = nil
	}

	// try to unmarshal data into RepeatUntil
	err = NewStrictDecoder(data).Decode(&dst.RepeatUntil)
	if err == nil {
		jsonRepeatUntil, _ := json.Marshal(dst.RepeatUntil)
		if string(jsonRepeatUntil) == "{}" { // empty struct
			dst.RepeatUntil = nil
		} else {
			if err = validator.Validate(dst.RepeatUntil); err != nil {
				dst.RepeatUntil = nil
			} else {
				match++
			}
		}
	} else {
		dst.RepeatUntil = nil
	}

	// try to unmarshal data into RepeatWhile
	err = NewStrictDecoder(data).Decode(&dst.RepeatWhile)
	if err == nil {
		jsonRepeatWhile, _ := json.Marshal(dst.RepeatWhile)
		if string(jsonRepeatWhile) == "{}" { // empty struct
			dst.RepeatWhile = nil
		} else {
			if err = validator.Validate(dst.RepeatWhile); err != nil {
				dst.RepeatWhile = nil
			} else {
				match++
			}
		}
	} else {
		dst.RepeatWhile = nil
	}

	// try to unmarshal data into Say
	err = NewStrictDecoder(data).Decode(&dst.Say)
	if err == nil {
		jsonSay, _ := json.Marshal(dst.Say)
		if string(jsonSay) == "{}" { // empty struct
			dst.Say = nil
		} else {
			if err = validator.Validate(dst.Say); err != nil {
				dst.Say = nil
			} else {
				match++
			}
		}
	} else {
		dst.Say = nil
	}

	// try to unmarshal data into SendSms
	err = NewStrictDecoder(data).Decode(&dst.SendSms)
	if err == nil {
		jsonSendSms, _ := json.Marshal(dst.SendSms)
		if string(jsonSendSms) == "{}" { // empty struct
			dst.SendSms = nil
		} else {
			if err = validator.Validate(dst.SendSms); err != nil {
				dst.SendSms = nil
			} else {
				match++
			}
		}
	} else {
		dst.SendSms = nil
	}

	// try to unmarshal data into SetVariable
	err = NewStrictDecoder(data).Decode(&dst.SetVariable)
	if err == nil {
		jsonSetVariable, _ := json.Marshal(dst.SetVariable)
		if string(jsonSetVariable) == "{}" { // empty struct
			dst.SetVariable = nil
		} else {
			if err = validator.Validate(dst.SetVariable); err != nil {
				dst.SetVariable = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetVariable = nil
	}

	// try to unmarshal data into SwitchCase
	err = NewStrictDecoder(data).Decode(&dst.SwitchCase)
	if err == nil {
		jsonSwitchCase, _ := json.Marshal(dst.SwitchCase)
		if string(jsonSwitchCase) == "{}" { // empty struct
			dst.SwitchCase = nil
		} else {
			if err = validator.Validate(dst.SwitchCase); err != nil {
				dst.SwitchCase = nil
			} else {
				match++
			}
		}
	} else {
		dst.SwitchCase = nil
	}

	// try to unmarshal data into WhileDo
	err = NewStrictDecoder(data).Decode(&dst.WhileDo)
	if err == nil {
		jsonWhileDo, _ := json.Marshal(dst.WhileDo)
		if string(jsonWhileDo) == "{}" { // empty struct
			dst.WhileDo = nil
		} else {
			if err = validator.Validate(dst.WhileDo); err != nil {
				dst.WhileDo = nil
			} else {
				match++
			}
		}
	} else {
		dst.WhileDo = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CallApi = nil
		dst.Capture = nil
		dst.Collect = nil
		dst.Dial = nil
		dst.DialToConversations = nil
		dst.DialToMany = nil
		dst.DialToWebRTC = nil
		dst.ForEach = nil
		dst.GoTo = nil
		dst.Hangup = nil
		dst.IfThenElse = nil
		dst.MachineDetection = nil
		dst.Pause = nil
		dst.Play = nil
		dst.PlayFromRecording = nil
		dst.Record = nil
		dst.RepeatUntil = nil
		dst.RepeatWhile = nil
		dst.Say = nil
		dst.SendSms = nil
		dst.SetVariable = nil
		dst.SwitchCase = nil
		dst.WhileDo = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ScriptInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ScriptInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ScriptInner) MarshalJSON() ([]byte, error) {
	if src.CallApi != nil {
		return json.Marshal(&src.CallApi)
	}

	if src.Capture != nil {
		return json.Marshal(&src.Capture)
	}

	if src.Collect != nil {
		return json.Marshal(&src.Collect)
	}

	if src.Dial != nil {
		return json.Marshal(&src.Dial)
	}

	if src.DialToConversations != nil {
		return json.Marshal(&src.DialToConversations)
	}

	if src.DialToMany != nil {
		return json.Marshal(&src.DialToMany)
	}

	if src.DialToWebRTC != nil {
		return json.Marshal(&src.DialToWebRTC)
	}

	if src.ForEach != nil {
		return json.Marshal(&src.ForEach)
	}

	if src.GoTo != nil {
		return json.Marshal(&src.GoTo)
	}

	if src.Hangup != nil {
		return json.Marshal(&src.Hangup)
	}

	if src.IfThenElse != nil {
		return json.Marshal(&src.IfThenElse)
	}

	if src.MachineDetection != nil {
		return json.Marshal(&src.MachineDetection)
	}

	if src.Pause != nil {
		return json.Marshal(&src.Pause)
	}

	if src.Play != nil {
		return json.Marshal(&src.Play)
	}

	if src.PlayFromRecording != nil {
		return json.Marshal(&src.PlayFromRecording)
	}

	if src.Record != nil {
		return json.Marshal(&src.Record)
	}

	if src.RepeatUntil != nil {
		return json.Marshal(&src.RepeatUntil)
	}

	if src.RepeatWhile != nil {
		return json.Marshal(&src.RepeatWhile)
	}

	if src.Say != nil {
		return json.Marshal(&src.Say)
	}

	if src.SendSms != nil {
		return json.Marshal(&src.SendSms)
	}

	if src.SetVariable != nil {
		return json.Marshal(&src.SetVariable)
	}

	if src.SwitchCase != nil {
		return json.Marshal(&src.SwitchCase)
	}

	if src.WhileDo != nil {
		return json.Marshal(&src.WhileDo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ScriptInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CallApi != nil {
		return obj.CallApi
	}

	if obj.Capture != nil {
		return obj.Capture
	}

	if obj.Collect != nil {
		return obj.Collect
	}

	if obj.Dial != nil {
		return obj.Dial
	}

	if obj.DialToConversations != nil {
		return obj.DialToConversations
	}

	if obj.DialToMany != nil {
		return obj.DialToMany
	}

	if obj.DialToWebRTC != nil {
		return obj.DialToWebRTC
	}

	if obj.ForEach != nil {
		return obj.ForEach
	}

	if obj.GoTo != nil {
		return obj.GoTo
	}

	if obj.Hangup != nil {
		return obj.Hangup
	}

	if obj.IfThenElse != nil {
		return obj.IfThenElse
	}

	if obj.MachineDetection != nil {
		return obj.MachineDetection
	}

	if obj.Pause != nil {
		return obj.Pause
	}

	if obj.Play != nil {
		return obj.Play
	}

	if obj.PlayFromRecording != nil {
		return obj.PlayFromRecording
	}

	if obj.Record != nil {
		return obj.Record
	}

	if obj.RepeatUntil != nil {
		return obj.RepeatUntil
	}

	if obj.RepeatWhile != nil {
		return obj.RepeatWhile
	}

	if obj.Say != nil {
		return obj.Say
	}

	if obj.SendSms != nil {
		return obj.SendSms
	}

	if obj.SetVariable != nil {
		return obj.SetVariable
	}

	if obj.SwitchCase != nil {
		return obj.SwitchCase
	}

	if obj.WhileDo != nil {
		return obj.WhileDo
	}

	// all schemas are nil
	return nil
}

type NullableScriptInner struct {
	value *ScriptInner
	isSet bool
}

func (v NullableScriptInner) Get() *ScriptInner {
	return v.value
}

func (v *NullableScriptInner) Set(val *ScriptInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptInner(val *ScriptInner) *NullableScriptInner {
	return &NullableScriptInner{value: val, isSet: true}
}

func (v NullableScriptInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
