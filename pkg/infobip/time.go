/*
Infobip Client API Libraries OpenAPI Specification

OpenAPI specification containing public endpoints supported in client API libraries.

Contact: support@infobip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infobip

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	T time.Time
}

const INFOBIP_TIME_FORMAT = "2006-01-02T15:04:05.000-0700"

var SUPPORTED_DATE_FORMATS = []string{
	INFOBIP_TIME_FORMAT,       // Default Infobip format
	"2006-01-02",              // yyyy-MM-dd
	"2006-01-02T15:04:05.000", // yyyy-MM-ddTHH:mm:ss.SSS (no timezone)
	"2006-01-02T15:04:05Z",    // yyyy-MM-ddTHH:mm:ssZ (UTC timezone)
}

func TimeNow() Time {
	s := time.Now().Format(INFOBIP_TIME_FORMAT)
	t, _ := time.Parse(INFOBIP_TIME_FORMAT, s)
	return Time{t}
}

func (ibtime Time) MarshalJSON() ([]byte, error) {
	s := time.Time(ibtime.T).Format(INFOBIP_TIME_FORMAT)
	return []byte(`"` + s + `"`), nil
}

func (ibtime *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	var parsedTime time.Time
	var err error
	for _, format := range SUPPORTED_DATE_FORMATS {
		parsedTime, err = time.Parse(format, s)
		if err == nil {
			*ibtime = Time{parsedTime}
			return nil
		}
	}
	return fmt.Errorf("unsupported date format: %s", s)
}

func (ibtime Time) String() string {
	return ibtime.T.Format(INFOBIP_TIME_FORMAT)
}
