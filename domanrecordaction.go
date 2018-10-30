package nettrigger

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

// DomainRecordActionBuilder constructs DNS actions from action specifications
// in the following forms:
//
//   dns.a     name zone ip    [ttl]
//   dns.cname name zone alias
func DomainRecordActionBuilder(spec ActionSpec) (Action, error) {
	switch spec.Type {
	case "dns.a":
		return DomainRecordAction{
			Domain: spec.Arg(1),
			Template: DomainRecordTemplate{
				Type: "A",
				Name: spec.Arg(0),
				Data: spec.Arg(2),
				TTL:  spec.Arg(3),
			},
		}.Apply, nil
	case "dns.ptr":
		return nil, errors.New("dns.ptr actions not yet supported")
	case "dns.cname":
		return DomainRecordAction{
			Domain: spec.Arg(1),
			Template: DomainRecordTemplate{
				Type: "CNAME",
				Name: spec.Arg(0),
				Data: spec.Arg(2),
				TTL:  spec.Arg(3),
			},
		}.Apply, nil
	default:
		return nil, nil
	}
}

// DomainRecordAction performs an action on a DNS record.
type DomainRecordAction struct {
	Domain   string
	Template DomainRecordTemplate
}

// Apply runs the domain record action.
func (action DomainRecordAction) Apply(ctx context.Context, env Environment, prov Providers) error {
	if prov.DNS == nil {
		return errors.New("unable to create DNS record: a DNS provider has not been configured")
	}

	record := DomainRecord{
		Type: action.Template.Type,
		Name: env.Expand(action.Template.Name),
		Data: env.Expand(action.Template.Data),
	}

	if ttl := env.Expand(action.Template.TTL); ttl != "" {
		value, err := strconv.ParseInt(ttl, 10, 32)
		if err != nil {
			return err
		}
		record.TTL = int(value)
	}

	if priority := env.Expand(action.Template.Priority); priority != "" {
		value, err := strconv.ParseInt(priority, 10, 32)
		if err != nil {
			return err
		}
		record.Priority = int(value)
	}

	if port := env.Expand(action.Template.Port); port != "" {
		value, err := strconv.ParseInt(port, 10, 32)
		if err != nil {
			return err
		}
		record.Port = int(value)
	}

	if weight := env.Expand(action.Template.Weight); weight != "" {
		value, err := strconv.ParseInt(weight, 10, 32)
		if err != nil {
			return err
		}
		record.Weight = int(value)
	}

	fmt.Printf("Register: %s: %+v\n", action.Domain, record)

	return prov.Register(ctx, env.Expand(action.Domain), record)
}
