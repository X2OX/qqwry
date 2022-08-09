package dnsV2

import (
	"encoding/json"
	"net/http"

	"github.com/miekg/dns"
	"go.x2ox.com/qqwry/dot"
	"go.x2ox.com/utils/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cors.CORS(w, r, Info)
}

func Info(w http.ResponseWriter, r *http.Request) {
	var (
		query  = r.URL.Query()
		domain = query.Get("name")
	)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	_ = json.NewEncoder(w).Encode(parseDomainIP(domain))
}

func parseDomainIP(domain string) []string {
	var arr []string

	p := dot.NewParser(domain, "A")
	if p != nil {
		result := p.Parse("")
		for _, v := range result.Msg.Answer {
			if r, ok := v.(*dns.A); ok {
				arr = append(arr, r.A.String())
			}
		}
	}

	p = dot.NewParser(domain, "AAAA")
	if p != nil {
		result := p.Parse("")
		for _, v := range result.Msg.Answer {
			if r, ok := v.(*dns.AAAA); ok {
				arr = append(arr, r.AAAA.String())
			}
		}
	}
	if len(arr) != 0 {
		return arr
	}

	p = dot.NewParser(domain, "CNAME")
	if p != nil {
		result := p.Parse("")
		for _, v := range result.Msg.Answer {
			if r, ok := v.(*dns.CNAME); ok {
				arr = append(arr, parseDomainIP(r.Target)...)
			}
		}
	}

	return arr
}
