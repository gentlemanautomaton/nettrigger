package nettrigger

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2/google"
	dns "google.golang.org/api/dns/v1"
)

type googleDNS struct {
	project string
	client  *http.Client
}

// NewGoogleDNS returns a Google DNS provider.
func NewGoogleDNS(project string) (DNS, error) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, dns.NdevClouddnsReadwriteScope)
	if err != nil {
		return nil, err
	}

	return googleDNS{
		project: project,
		client:  c,
	}, nil
}

func (p googleDNS) Register(ctx context.Context, domain string, record DomainRecord) error {
	dnsService, err := dns.New(p.client)
	if err != nil {
		return err
	}

	recordName := record.Name + "." + domain + "."
	zoneName, err := p.zoneName(ctx, domain, dnsService)
	if err != nil {
		return err
	}

	existingRecords, err := p.existingRecords(ctx, zoneName, recordName, dnsService)
	if err != nil {
		return err
	}

	change := dns.Change{
		Additions: []*dns.ResourceRecordSet{
			&dns.ResourceRecordSet{
				Name:    record.Name + "." + domain + ".",
				Type:    record.Type,
				Rrdatas: []string{record.Data},
				Ttl:     int64(record.TTL),
			},
		},
		Deletions: existingRecords,
	}

	call := dnsService.Changes.Create(p.project, zoneName, &change)
	call.Context(ctx)
	_, err = call.Do()

	return err
}

func (p googleDNS) zoneName(ctx context.Context, domain string, s *dns.Service) (zoneName string, err error) {
	zones, err := s.ManagedZones.List(p.project).Do()
	if err != nil {
		return "", err
	}

	for _, zone := range zones.ManagedZones {
		if strings.TrimSuffix(zone.DnsName, ".") == domain {
			return zone.Name, nil
		}
	}

	return "", fmt.Errorf("couldn't find zone for \"%s\" domain within \"%s\" project", domain, p.project)
}

func (p googleDNS) existingRecords(ctx context.Context, zoneName, recordName string, s *dns.Service) (records []*dns.ResourceRecordSet, err error) {
	var pageToken string

	for {
		call := s.ResourceRecordSets.List(p.project, zoneName)
		call.Context(ctx)
		call.Name(recordName)
		call.MaxResults(200)
		if pageToken != "" {
			call.PageToken(pageToken)
		}

		result, err := call.Do()
		if err != nil {
			return nil, err
		}
		records = append(records, result.Rrsets...)

		if result.NextPageToken == "" {
			return records, nil
		}
		pageToken = result.NextPageToken
	}
}
