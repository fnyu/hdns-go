package hdns

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/fnyu/hdns-go/hdns/schema"
)

type Zone struct {
	ID              string
	Created         *time.Time
	Modified        *time.Time
	LegacyDnsHost   string
	LegacyNs        []string
	Name            string
	Ns              []string
	Owner           string
	Paused          bool
	Permission      string
	Project         string
	Registrar       string
	Status          ZoneStatus
	Ttl             int
	Verified        *time.Time
	RecordsCount    int
	IsSecondaryDns  bool
	TxtVerification ZoneTxtVerification
}

type ZoneTxtVerification struct {
	Name  string
	Token string
}

type ZoneStatus string

const (
	ZoneStatusVerified ZoneStatus = "verified"

	ZoneStatusPending ZoneStatus = "pending"

	ZoneStatusFailed ZoneStatus = "failed"
)

// ZoneClient is a client for the zones API.
type ZoneClient struct {
	client *Client
}

// GetByID retrieves a zone by its ID. If the zone does not exist, nil is returned.
func (c *ZoneClient) GetByID(ctx context.Context, id string) (*Zone, *http.Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/zones/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.ZoneResponse
	res, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, Error([]int{
			http.StatusUnauthorized,
			http.StatusForbidden,
			http.StatusNotFound,
			http.StatusNotAcceptable,
		}, res.StatusCode)
	}

	return ZoneFromSchema(body.Zone), res, nil
}

// ZoneListOpts specifies options for listing zones.
type ZoneListOpts struct {
	ListOpts
	Name       string
	SearchName string
}

func (o ZoneListOpts) values() url.Values {
	vals := o.ListOpts.values()
	if len(o.Name) > 0 {
		vals.Add("name", o.Name)
	}
	if len(o.SearchName) > 0 {
		vals.Add("search_name", o.SearchName)
	}

	return vals
}

// List returns a list of zones for a specific page and filters.
//
// Please note that filters specified in opts are not taken into account
// when their value corresponds to their zero value or when they are empty.
func (c *ZoneClient) List(ctx context.Context, opts ZoneListOpts) ([]*Zone, *http.Response, error) {
	path := "/zones?" + opts.values().Encode()
	req, err := c.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var body schema.ZonesResponse
	res, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res, Error([]int{
			http.StatusUnauthorized,
			http.StatusNotAcceptable,
		}, res.StatusCode)
	}

	zones := make([]*Zone, 0, len(body.Zones))
	for _, z := range body.Zones {
		zones = append(zones, ZoneFromSchema(z))
	}

	return zones, res, nil
}

type ZoneCreateOpts struct {
	Name string
	Ttl  int
}

func (o ZoneCreateOpts) Validate() error {
	if o.Name == "" {
		return errors.New("missing name")
	}
	if o.Ttl < 0 {
		return errors.New("ttl must be a positive integer")
	}

	return nil
}

// Create creates zone with specified name and default TTL
func (c *ZoneClient) Create(ctx context.Context, opts ZoneCreateOpts) (*Zone, *http.Response, error) {
	if err := opts.Validate(); err != nil {
		return nil, nil, err
	}

	reqData := schema.ZoneCreateRequest{
		Name: opts.Name,
		Ttl:  opts.Ttl,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(ctx, http.MethodPost, "/zones", bytes.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}

	var body schema.ZoneResponse
	res, err := c.client.Do(req, &body)
	if err != nil {
		return nil, res, err
	}

	return ZoneFromSchema(body.Zone), res, nil
}

type ZoneUpdateOpts struct{}

func (c *ZoneClient) Update(ctx context.Context, zone *Zone, opts ZoneUpdateOpts) (*Zone, *http.Response, error) {
	return nil, nil, nil
}

func (c *ZoneClient) Delete(ctx context.Context, zone *Zone) (*http.Response, error) {
	return nil, nil
}
