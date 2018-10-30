package nettrigger

// DomainRecordTemplate is a template used by DNS actions. Its fields may
// contain environment variable expressions.
type DomainRecordTemplate struct {
	Type     string
	Name     string
	Data     string
	TTL      string
	Priority string
	Port     string
	Weight   string
}
