package field

import (
	"encoding/json"
	"time"
)

// DateTime is a custom field for datetime values that are returned
// from the Hetzner DNS API. There is a discrepancy between the format
// of the datetime fields that are actually returned and the docs.
// I have filled a request to changed that, but for now we have to live with it.
type DateTime struct {
	Time *time.Time
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

	d.Time = &t

	return nil
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	var s string
	if d.Time != nil {
		s = d.Time.String()
	} else {
		s = ""
	}

	return json.Marshal(s)
}
