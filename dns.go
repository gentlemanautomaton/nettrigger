package nettrigger

import "context"

// DNS is a DNS provider interface.
type DNS interface {
	Register(ctx context.Context, domain string, record DomainRecord) error
}
