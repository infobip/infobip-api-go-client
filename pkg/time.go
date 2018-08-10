package infobip

import (
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
