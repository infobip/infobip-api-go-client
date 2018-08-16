package infobip

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMarshalingUTC(t *testing.T) {
	tm := Time{T: time.Unix(1533933076, 0).UTC()}
	j, err := json.Marshal(tm)
	if err != nil {
		t.Fatalf("Expected time to be marshalable, but encountered an error: %s", err.Error())
	}
	s := string(j)
	e := `"2018-08-10T20:31:16.000+0000"`
	if s != e {
		t.Errorf("Expected time to marshal into `%s`, but was `%s`", e, s)
	}
}

func TestMarshalingZoned(t *testing.T) {
	tm := Time{T: time.Unix(1533933076, 0).In(time.FixedZone("+5.5", 19800))}
	j, err := json.Marshal(tm)
	if err != nil {
		t.Fatalf("Expected time %+v to be marshalable, but encountered an error: %s", tm, err.Error())
	}
	s := string(j)
	e := `"2018-08-11T02:01:16.000+0530"`
	if s != e {
		t.Errorf("Expected time %+v to marshal into `%s`, but was `%s`", tm, e, s)
	}
}

func TestUnmarshalingUTC(t *testing.T) {
	s := `"2018-08-10T20:31:16.000+0000"`
	var tm Time
	err := json.Unmarshal([]byte(s), &tm)
	if err != nil {
		t.Fatalf("Expected json string `%s` to be unmarshalable, but encountered an error: %s", s, err.Error())
	}
	u := tm.T.Unix()
	var e int64 = 1533933076
	if e != u {
		t.Errorf("Expected unmarshaled json string `%s` to have Unix timestamp %d, but was %d", s, e, u)
	}
}

func TestUnmarshalingZoned(t *testing.T) {
	s := `"2018-08-11T02:01:16.000+0530"`
	var tm Time
	err := json.Unmarshal([]byte(s), &tm)
	if err != nil {
		t.Fatalf("Expected json string `%s` to be unmarshalable, but encountered an error: %s", s, err.Error())
	}
	u := tm.T.Unix()
	var e int64 = 1533933076
	if e != u {
		t.Errorf("Expected unmarshaled json string `%s` to have Unix timestamp %d, but was %d", s, e, u)
	}
}

func TestTimeUnit(t *testing.T) {
	values := TimeUnitValues
	for i := 0; i < len(values); i++ {
		j, err := json.Marshal(values[i])
		if err != nil {
			t.Fatalf("Expected to be able to marshal %v, but was error %v", values[i], err)
		}
		var actual TimeUnit
		json.Unmarshal(j, &actual)
		if actual != values[i] {
			t.Errorf("Expected %v to be equal to %v", actual, values[i])
		}
	}
}
