package sms

//This is a generated file and is not intended for modification!

import (
	"encoding/json"
	"testing"
)

var (
	//reference importted packages just in case auto-generated code doesn't
	_ = json.Compact
	_ = testing.RunTests
)

func TestDeliveryDay(t *testing.T) {
	values := DeliveryDayValues
	for i := 0; i < len(values); i++ {
		j, err := json.Marshal(values[i])
		if err != nil {
			t.Fatalf("Expected to be able to marshal %v, but was error %v", values[i], err)
		}
		var actual DeliveryDay
		json.Unmarshal(j, &actual)
		if actual != values[i] {
			t.Errorf("Expected %v to be equal to %v", actual, values[i])
		}
	}
}
