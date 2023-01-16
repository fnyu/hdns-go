package field

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type testHelperResponse struct {
	Created *DateTime `json:"created"`
}

func TestDateTimeUnmarshaling(t *testing.T) {
	data := []byte(`{
		"created": "2023-01-13 16:26:40.086 +0000 UTC"
	}`)

	want := time.Date(2023, time.January, 13, 16, 26, 40, 0, time.UTC)

	var v testHelperResponse
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatal(err)
	}

	if v.Created == nil {
		t.Errorf("expected %T, got nil", &DateTime{})
	}

	if ok := cmp.Equal(want, *v.Created.Time); !ok {
		t.Errorf("expected %v, got %v", want, v.Created.Time)
	}
}
