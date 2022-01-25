package dot

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/miekg/dns"
)

type Parser struct {
	Name   string // 域名
	RRType uint16 // 解析类型
}

func NewParser(name, rrType string) *Parser {
	p := &Parser{}
	if p.Name = parseDomain(name); p.Name == "" {
		return nil
	}
	if p.RRType = parseType(rrType); p.RRType == 0 {
		return nil
	}
	return p
}

func (p *Parser) Parse(provider string) *Result {
	var (
		r   = &Result{provider: provider, msg: new(dns.Msg)}
		msg = new(dns.Msg)
		c   = dns.Client{Net: "tcp-tls"}
	)
	if _, ok := providers[provider]; !ok {
		r.provider = "Google"
	}
	msg.SetQuestion(p.Name, p.RRType)

	r.msg, r.rtt, r.err = c.Exchange(msg, providers[r.provider]+":853")
	return r
}

type Result struct {
	provider string
	msg      *dns.Msg
	rtt      time.Duration
	err      error
}

func (d *Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.msg)
}

func (d *Result) String() string {
	if d.err != nil {
		return fmt.Sprintf(outputTpl, d.provider, d.rtt, d.err.Error())
	}
	return fmt.Sprintf(outputTpl, d.provider, d.rtt, d.msg)
}

func (d *Result) IP() string {
	for _, rr := range d.msg.Answer {
		switch rr := rr.(type) {
		case *dns.A:
			return rr.A.String()
		}
	}
	return ""
}

var providers = map[string]string{
	"Cloudflare": "cloudflare-dns.com",
	"AdGuard":    "dns-unfiltered.adguard.com",
	"Quad9":      "dns10.quad9.net",
	"Google":     "dns.google",
	"Switch":     "dns.switch.ch",
	"AliDNS":     "dns.alidns.com",
	"IIJ":        "public.dns.iij.jp",
	"DnsPod":     "dot.pub",
	"DeCloudUs":  "dns.decloudus.com",
	"DnsSb":      "185.222.222.222",
}
