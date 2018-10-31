package nettrigger

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type digitalOceanDNS struct {
	c *godo.Client
}

// NewDigitalOceanDNS returns a DigitalOcean DNS provider.
func NewDigitalOceanDNS(token string) DNS {
	tokenSource := &digitalOceanTokenSource{
		AccessToken: token,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	return digitalOceanDNS{
		c: godo.NewClient(oauthClient),
	}
}

func (dns digitalOceanDNS) Register(ctx context.Context, domain string, record DomainRecord) error {
	existingID, existing, err := dns.existing(ctx, domain, record)
	if err != nil {
		return err
	}

	data := godo.DomainRecordEditRequest{
		Type:     record.Type,
		Name:     record.Name,
		Data:     record.Data,
		Priority: record.Priority,
		Port:     record.Port,
		TTL:      record.TTL,
		Weight:   record.Weight,
	}

	if existing {
		_, _, err = dns.c.Domains.EditRecord(ctx, domain, existingID, &data)
		return err
	}

	_, _, err = dns.c.Domains.CreateRecord(ctx, domain, &data)
	return err
}

func (dns digitalOceanDNS) existing(ctx context.Context, domain string, record DomainRecord) (id int, ok bool, err error) {
	opt := &godo.ListOptions{
		PerPage: 200,
	}
	for {
		records, resp, err := dns.c.Domains.Records(ctx, domain, opt)
		if err != nil {
			return 0, false, err
		}

		for _, r := range records {
			if strings.EqualFold(r.Type, record.Type) && strings.EqualFold(r.Name, record.Name) {
				return r.ID, true, nil
			}
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return 0, false, err
		}

		opt.Page = page + 1
	}

	return 0, false, nil
}
