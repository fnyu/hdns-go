package hdns

import (
	"time"
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
