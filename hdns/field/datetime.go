package field

import (
	"encoding/json"
	"time"
)

// DateTime is a custom field for datetime values that are returned
// from the Hetzner DNS API. There is a discrepancy between the format
// of the datetime fields that are actually returned and the docs.
type DateTime struct {
	Time *time.Time
	Ns   int
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v == "" {
		return nil
	}

	t, err := time.Parse("2006-01-02 15:04:05.999 -0700 MST", v)
	if err != nil {
		return err
	}

	// Due to a bizzarre date time format in the responses
	// We have to strip the nanoseconds data from the Time
	// field. The value will be preserved in the struct
	// and restored in the marshaling process
	d.Ns = t.Nanosecond()
	t = t.Add(-1 * time.Duration(t.Nanosecond()))
	d.Time = &t

	return nil
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	var s string
	if d.Time != nil {
		s = d.Time.Add(time.Nanosecond * time.Duration(d.Ns)).String()
	} else {
		s = ""
	}

	return json.Marshal(s)
}
