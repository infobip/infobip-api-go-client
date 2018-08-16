package infobip

import (
	"errors"
	"strings"
	"time"
)

type Time struct {
	T time.Time
}

const inputFormat = "2006-01-02T15:04:05.000-0700"
const outputFormat = "2006-01-02T15:04:05.000-0700"

func (ibtime Time) MarshalJSON() ([]byte, error) {
	s := time.Time(ibtime.T).Format(outputFormat)
	return []byte(`"` + s + `"`), nil
}

func (ibtime *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	t, err := time.Parse(inputFormat, s)
	if err != nil {
		return err
	}
	*ibtime = Time{t}
	return nil
}

type TimeUnit int

const (
	TU_NANOSECONDS TimeUnit = iota
	TU_MICROSECONDS
	TU_MILLISECONDS
	TU_SECONDS
	TU_MINUTES
	TU_HOURS
	TU_DAYS
)

var TimeUnitValues = []TimeUnit{
	TU_NANOSECONDS, TU_MICROSECONDS, TU_MILLISECONDS, TU_SECONDS, TU_MINUTES, TU_HOURS, TU_DAYS,
}

func (val TimeUnit) String() string {
	names := []string{
		"NANOSECONDS",
		"MICROSECONDS",
		"MILLISECONDS",
		"SECONDS",
		"MINUTES",
		"HOURS",
		"DAYS",
	}
	if val < TU_NANOSECONDS || val > TU_DAYS {
		return "Unknown"
	}
	return names[val]
}

func (val TimeUnit) MarshalJSON() ([]byte, error) {
	return []byte(`"` + val.String() + `"`), nil
}

func (val *TimeUnit) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	switch s {
	case "NANOSECONDS":
		*val = TU_NANOSECONDS
	case "MICROSECONDS":
		*val = TU_MICROSECONDS
	case "MILLISECONDS":
		*val = TU_MILLISECONDS
	case "SECONDS":
		*val = TU_SECONDS
	case "MINUTES":
		*val = TU_MINUTES
	case "HOURS":
		*val = TU_HOURS
	case "DAYS":
		*val = TU_DAYS
	default:
		return errors.New("Unknown TimeUnit value " + s)
	}
	return nil
}
