package schema

import "github.com/fnyu/hdns-go/hdns/field"

// Zone represents single Zone object returned by Hetzner DNS API
type Zone struct {
	ID              string              `json:"id"`
	Created         *field.DateTime     `json:"created"`
	Modified        *field.DateTime     `json:"modified"`
	LegacyDnsHost   string              `json:"legacy_dns_host"`
	LegacyNs        []string            `json:"legacy_ns"`
	Name            string              `json:"name"`
	Ns              []string            `json:"ns"`
	Owner           string              `json:"owner"`
	Paused          bool                `json:"paused"`
	Permission      string              `json:"permission"`
	Project         string              `json:"project"`
	Registrar       string              `json:"registrar"`
	Status          string              `json:"status"`
	Ttl             int                 `json:"ttl"`
	Verified        *field.DateTime     `json:"verified"`
	RecordsCount    int                 `json:"records_count"`
	IsSecondaryDns  bool                `json:"is_secondary_dns"`
	TxtVerification ZoneTxtVerification `json:"txt_verification"`
}

type ZoneResponse struct {
	Zone Zone `json:"zone"`
}

type ZonesResponse struct {
	Zones []Zone `json:"zones"`
}

type ZoneTxtVerification struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ZoneCreateRequest struct {
	Name string `json:"name"`
	Ttl  int    `json:"ttl"`
}

type ZoneUpdateRequest struct {
	Name string `json:"name"`
	Ttl  int    `json:"ttl"`
}
