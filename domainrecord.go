package nettrigger

// DomainRecord is a domain record used by DNS providers.
type DomainRecord struct {
	Type     string
	Name     string
	Data     string
	TTL      int
	Priority int
	Port     int
	Weight   int
}
