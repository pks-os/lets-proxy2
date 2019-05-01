//nolint:golint
package cert_manager

import "go.uber.org/zap"

type certNameType string

func (n certNameType) String() string {
	return string(n)
}

func certNameFromDomain(domain DomainName) certNameType {
	return certNameType(domain)
}

func domainNamesFromCertificateName(name certNameType) []DomainName {
	return []DomainName{DomainName(name), DomainName("www." + name)}
}

func logCetName(certName certNameType) zap.Field {
	return zap.String("cert_name", string(certName))
}