package hdns

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// This is the base URL of the API.
	Endpoint = "https://dns.hetzner.com/api/v1"

	// This string will be used in User-Agent HTTP header.
	UserAgent = "hdnsapi/" + Version
)

// Client is a client for the Hetzner DNS API.
type Client struct {
	endpoint   string
	token      string
	httpClient *http.Client
	userAgent  string

	Zone ZoneClient
}

// A ClientOption is used to configure a Client.
type ClientOption func(*Client)

// WithEndpoint configures a Client to use the specified API endpoint.
func WithEndpoint(endpoint string) ClientOption {
	return func(client *Client) {
		client.endpoint = strings.TrimRight(endpoint, "/")
	}
}

// WithToken configures a Client to use the specified token for authentication.
func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.token = token
	}
}

// WithHTTPClient configures a Client to perform HTTP requests with httpClient.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

// NewClient creates a new client.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		endpoint:   Endpoint,
		httpClient: &http.Client{},
		userAgent:  UserAgent,
	}

	for _, option := range options {
		option(client)
	}

	client.Zone = ZoneClient{client: client}

	return client
}

// NewRequest creates an HTTP request against the API. The returned request
// is assigned with ctx and has all necessary headers set (auth, user agent, etc.).
func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	url := c.endpoint + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)

	if c.token != "" {
		req.Header.Set("Auth-API-Token", c.token)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req = req.WithContext(ctx)
	return req, nil
}

// Do performs an HTTP request against the API.
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	var err error

	res, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return res, err
	}

	return res, nil
}

type ListOpts struct {
	Page    int
	PerPage int
}

func (o ListOpts) values() url.Values {
	vals := url.Values{}
	if o.Page > 0 {
		vals.Add("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		vals.Add("per_page", strconv.Itoa(o.PerPage))
	}
	return vals
}
