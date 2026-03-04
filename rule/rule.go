package rule

type Type int

const (
	DomainSuffix Type = iota
	DomainKeyword
	GEOIP
	IPCIDR
	FINAL
)

type Rule interface {
	Type() Type
}
