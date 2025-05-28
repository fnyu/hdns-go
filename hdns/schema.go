package hdns

import "github.com/fnyu/hdns-go/hdns/schema"

// This file provides converter functions to convert models in the
// schema package to models in the hdns package and viceversa.
// This behavior was inspired by https://github.com/hetznercloud/hcloud-go

// ZoneFromSchema converts a schema.Zone to Zone.
func ZoneFromSchema(s schema.Zone) *Zone {
	zone := &Zone{
		ID:              s.ID,
		Created:         s.Created.Time,
		Modified:        s.Modified.Time,
		LegacyDnsHost:   s.LegacyDnsHost,
		LegacyNs:        s.LegacyNs,
		Name:            s.Name,
		Ns:              s.Ns,
		Owner:           s.Owner,
		Paused:          s.Paused,
		Permission:      s.Permission,
		Project:         s.Project,
		Registrar:       s.Registrar,
		Status:          ZoneStatus(s.Status),
		Ttl:             s.Ttl,
		Verified:        s.Verified.Time,
		RecordsCount:    s.RecordsCount,
		IsSecondaryDns:  s.IsSecondaryDns,
		TxtVerification: ZoneTxtVerification(s.TxtVerification),
	}

	return zone
}

// ZonesFromSchema converts a schema.Zones to []Zone
func ZonesFromSchema(s []schema.Zone) []*Zone {
	zones := make([]*Zone, len(s))
	for i, zone := range s {
		zones[i] = ZoneFromSchema(zone)
	}

	return zones
}
