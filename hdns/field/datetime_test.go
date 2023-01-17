package field

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

type testHelperResponse struct {
	Created  *DateTime `json:"created"`
	Modified *DateTime `json:"modified"`
}

func TestDateTimeUnmarshaling(t *testing.T) {
	data := []byte(`{
		"created": "2023-01-13 16:26:40.086 +0000 UTC",
		"modified": "2023-01-13 16:32:22.171 +0000 UTC"
	}`)

	var v testHelperResponse
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatal(err)
	}

	if v.Created == nil {
		t.Errorf("expected %T, got nil", &DateTime{})
	}

	var want time.Time

	want = time.Date(2023, time.January, 13, 16, 26, 40, int(86*time.Millisecond), time.UTC)
	if !v.Created.Time.Equal(want) {
		t.Errorf("expected %v, got %v", want, v.Modified.Time)
	}

	want = time.Date(2023, time.January, 13, 16, 32, 22, int(171*time.Millisecond), time.UTC)
	if !v.Modified.Time.Equal(want) {
		t.Errorf("expected %v, got %v", want, v.Modified.Time)
	}

}

func TestDateTimeMarshaling(t *testing.T) {
	created := time.Date(2023, time.January, 13, 16, 26, 40, int(86*time.Millisecond), time.UTC)
	modified := time.Date(2023, time.January, 13, 16, 32, 22, int(171*time.Millisecond), time.UTC)
	v := testHelperResponse{
		Created:  &DateTime{Time: &created},
		Modified: &DateTime{Time: &modified},
	}

	data, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	want := []byte(`{"created":"2023-01-13 16:26:40.086 +0000 UTC","modified":"2023-01-13 16:32:22.171 +0000 UTC"}`)

	if !bytes.Equal(want, data) {
		t.Errorf("expected %v, got %v", want, data)
	}
}
