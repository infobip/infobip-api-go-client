package infobip

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

func TestSupportedSpecifiedTimeFormat(t *testing.T) {
	testCases := []struct {
		time string
	}{
		{`"2035-08-18T12:08:42.777+0000"`},
		{`"2035-08-18T13:08:42.777+0100"`},
		{`"2035-08-18T11:08:42.777-0100"`},
		{`"2035-08-18T17:08:42.777+0500"`},
		{`"2035-08-18T07:08:42.777-0500"`},
		{`"2035-08-18T13:38:42.777+0130"`},
		{`"2035-08-18T10:38:42.777-0130"`},
		{`"2035-08-18T17:38:42.777+0530"`},
		{`"2035-08-18T06:38:42.777-0530"`},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.time, func(t *testing.T) {
			var tm infobip.Time
			err := json.Unmarshal([]byte(tc.time), &tm)
			if err != nil {
				t.Fatalf("Expected json string `%s` to be unmarshalable, but encountered an error: %s", tc.time, err.Error())
			}
			u := tm.T.Unix()
			var e int64 = 2071051722
			if e != u {
				t.Errorf("Expected unmarshaled json string `%s` to have Unix timestamp %d, but was %d", tc.time, e, u)
			}
		})
	}
}

func TestSupportedSpecifiedTimeJsonFormat(t *testing.T) {
	testCases := []struct {
		timeString string
		time       infobip.Time
	}{
		{`"2035-08-18T12:08:42.777+0000"`, infobip.Time{T: time.Unix(2071051722, 777000000).UTC()}},
		{`"2035-08-18T13:08:42.777+0100"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("+1.0", 3600))}},
		{`"2035-08-18T11:08:42.777-0100"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("-1.0", -3600))}},
		{`"2035-08-18T17:08:42.777+0500"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("+1.0", 18000))}},
		{`"2035-08-18T07:08:42.777-0500"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("-1.0", -18000))}},
		{`"2035-08-18T13:38:42.777+0130"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("+1.5", 5400))}},
		{`"2035-08-18T10:38:42.777-0130"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("-1.5", -5400))}},
		{`"2035-08-18T17:38:42.777+0530"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("+5.5", 19800))}},
		{`"2035-08-18T06:38:42.777-0530"`, infobip.Time{T: time.Unix(2071051722, 777000000).In(time.FixedZone("-5.5", -19800))}},
	}
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.timeString, func(t *testing.T) {
			j, err := json.Marshal(tc.time)
			if err != nil {
				t.Fatalf("Expected time to be marshalable, but encountered an error: %s", err.Error())
			}
			s := string(j)
			e := tc.timeString
			if s != e {
				t.Errorf("Expected time to marshal into `%s`, but was `%s`", e, s)
			}
		})
	}
}
