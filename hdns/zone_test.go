package hdns

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/fnyu/hdns-go/hdns/field"
	"github.com/fnyu/hdns-go/hdns/schema"
)

func TestZoneClientGetByID(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	env.Mux.HandleFunc("/zones/cho7asai6vao7Aig7viH4l", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(schema.ZoneResponse{
			Zone: schema.Zone{
				ID:       "cho7asai6vao7Aig7viH4l",
				Created:  &field.DateTime{},
				Modified: &field.DateTime{},
				Verified: &field.DateTime{},
			},
		})
	})

	ctx := context.Background()
	zone, _, err := env.Client.Zone.GetByID(ctx, "cho7asai6vao7Aig7viH4l")
	if err != nil {
		t.Fatal(err)
	}

	if zone == nil {
		t.Fatal("zone not found")
	}

	if zone.ID != "cho7asai6vao7Aig7viH4l" {
		t.Errorf("unexpected zone ID: %v", zone.ID)
	}
}
