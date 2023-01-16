package hdns

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testEnv struct {
	Server *httptest.Server
	Mux    *http.ServeMux
	Client *Client
}

func (env *testEnv) Teardown() {
	env.Server.Close()
	env.Server = nil
	env.Mux = nil
	env.Client = nil
}

func newTestEnv() testEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(
		WithEndpoint(server.URL),
		WithToken("token"),
	)
	return testEnv{
		Server: server,
		Mux:    mux,
		Client: client,
	}
}

func TestClientEndpointTrailingSlashesRemoved(t *testing.T) {
	client := NewClient(WithEndpoint("http://api/v1.0/////"))
	if strings.HasSuffix(client.endpoint, "/") {
		t.Fatalf("endpoint has trailing slashes: %q", client.endpoint)
	}
}
