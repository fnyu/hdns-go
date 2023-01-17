package hdns

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/fnyu/hdns-go/hdns/schema"
)

func TestZoneFromSchema(t *testing.T) {
	data := []byte(`{
    "id": "cho7asai6vao7Aig7viH4l",
    "created": "2023-01-13 16:26:40.086 +0000 UTC",
    "modified": "2023-01-13 16:32:22.171 +0000 UTC",
    "legacy_dns_host": "",
    "legacy_ns": [
      "ns1.hetzner.cloud",
			"ns2.hetzner.cloud"
    ],
    "name": "string",
    "ns": [
      "string"
    ],
    "owner": "",
    "paused": true,
    "permission": "string",
    "project": "string",
    "registrar": "string",
    "status": "verified",
    "ttl": 0,
    "verified": "",
    "records_count": 0,
    "is_secondary_dns": true,
    "txt_verification": {
      "name": "string",
      "token": "string"
    }
}`)

	var s schema.Zone
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatal(err)
	}

	zone := ZoneFromSchema(s)

	if zone.ID != "cho7asai6vao7Aig7viH4l" {
		t.Errorf("unexpected ID: %v", zone.ID)
	}

	if !zone.Created.Equal(time.Date(2023, time.January, 13, 16, 26, 40, int(86*time.Millisecond), time.UTC)) {
		t.Errorf("unexpected Created: %v", zone.Created)
	}
}
