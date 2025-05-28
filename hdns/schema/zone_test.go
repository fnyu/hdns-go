package schema

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestZoneCreateUpdateRequest(t *testing.T) {
	testCases := []struct {
		name string
		in   ZoneCreateRequest
		out  []byte
	}{
		{
			name: "empty create request",
			in:   ZoneCreateRequest{},
			out:  []byte(`{"name":"","ttl":0}`),
		},
		{
			name: "create request with ttl set",
			in:   ZoneCreateRequest{Ttl: 3600},
			out:  []byte(`{"name":"","ttl":3600}`),
		},
		{
			name: "create request with name only",
			in:   ZoneCreateRequest{Name: "example.com"},
			out:  []byte(`{"name":"example.com","ttl":0}`),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			data, err := json.Marshal(testCase.in)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(data, testCase.out) {
				t.Fatalf("output %s does not match %s", data, testCase.out)
			}
		})
	}
}
