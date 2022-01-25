package dot

import (
	"strings"

	"github.com/miekg/dns"
	"golang.org/x/net/idna"
)

const outputTpl = ";; X2OX DoT Query [%s] - %s\n;;\n%s"

func parseDomain(domain string) string {
	var err error
	if domain, err = idna.New(
		idna.MapForLookup(), idna.Transitional(true), idna.StrictDomainName(false),
	).ToASCII(strings.TrimSpace(domain)); err != nil || len(domain) <= 2 {
		return ""
	}

	if domain[len(domain)-1] != '.' {
		domain += "."
	}

	return domain
}

func parseType(_type string) uint16 {
	t, ok := parseTypeMap[strings.ToTitle(_type)]
	if ok {
		return t
	}
	return 0
}

var parseTypeMap = map[string]uint16{
	"A": dns.TypeA, "NS": dns.TypeNS, "MD": dns.TypeMD, "MF": dns.TypeMF, "CNAME": dns.TypeCNAME,
	"SOA": dns.TypeSOA, "MB": dns.TypeMB, "MG": dns.TypeMG, "MR": dns.TypeMR, "NULL": dns.TypeNULL,
	"PTR": dns.TypePTR, "HINFO": dns.TypeHINFO, "MINFO": dns.TypeMINFO, "MX": dns.TypeMX,
	"TXT": dns.TypeTXT, "RP": dns.TypeRP, "AFSDB": dns.TypeAFSDB, "X25": dns.TypeX25,
	"ISDN": dns.TypeISDN, "RT": dns.TypeRT, "NSAPPTR": dns.TypeNSAPPTR, "SIG": dns.TypeSIG,
	"KEY": dns.TypeKEY, "PX": dns.TypePX, "GPOS": dns.TypeGPOS, "AAAA": dns.TypeAAAA, "LOC": dns.TypeLOC,
	"NXT": dns.TypeNXT, "EID": dns.TypeEID, "NIMLOC": dns.TypeNIMLOC, "SRV": dns.TypeSRV,
	"ATMA": dns.TypeATMA, "NAPTR": dns.TypeNAPTR, "KX": dns.TypeKX, "CERT": dns.TypeCERT,
	"DNAME": dns.TypeDNAME, "OPT": dns.TypeOPT, "APL": dns.TypeAPL, "DS": dns.TypeDS, "SSHFP": dns.TypeSSHFP,
	"RRSIG": dns.TypeRRSIG, "NSEC": dns.TypeNSEC, "DNSKEY": dns.TypeDNSKEY, "DHCID": dns.TypeDHCID,
	"NSEC3": dns.TypeNSEC3, "NSEC3PARAM": dns.TypeNSEC3PARAM, "TLSA": dns.TypeTLSA, "SMIMEA": dns.TypeSMIMEA,
	"HIP": dns.TypeHIP, "NINFO": dns.TypeNINFO, "RKEY": dns.TypeRKEY, "TALINK": dns.TypeTALINK,
	"CDS": dns.TypeCDS, "CDNSKEY": dns.TypeCDNSKEY, "OPENPGPKEY": dns.TypeOPENPGPKEY, "CSYNC": dns.TypeCSYNC,
	"ZONEMD": dns.TypeZONEMD, "SVCB": dns.TypeSVCB, "HTTPS": dns.TypeHTTPS, "SPF": dns.TypeSPF,
	"UINFO": dns.TypeUINFO, "UID": dns.TypeUID, "GID": dns.TypeGID, "UNSPEC": dns.TypeUNSPEC,
	"NID": dns.TypeNID, "L32": dns.TypeL32, "L64": dns.TypeL64, "LP": dns.TypeLP, "EUI48": dns.TypeEUI48,
	"EUI64": dns.TypeEUI64, "URI": dns.TypeURI, "CAA": dns.TypeCAA, "AVC": dns.TypeAVC, "TKEY": dns.TypeTKEY,
	"TSIG": dns.TypeTSIG, "IXFR": dns.TypeIXFR, "AXFR": dns.TypeAXFR, "MAILB": dns.TypeMAILB,
	"MAILA": dns.TypeMAILA, "ANY": dns.TypeANY, "TA": dns.TypeTA, "DLV": dns.TypeDLV, "Reserved": dns.TypeReserved,
}
