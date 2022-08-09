package main

import (
	"log"
	"net/http"

	"go.x2ox.com/qqwry/api/dns"
	dnsV2 "go.x2ox.com/qqwry/api/dns/v2"
	"go.x2ox.com/qqwry/api/info"
	"go.x2ox.com/qqwry/api/qqwry"
	"go.x2ox.com/qqwry/api/whois"
)

func main() {
	http.HandleFunc("/api/qqwry", qqwry.Handler)
	http.HandleFunc("/api/dns", dns.Handler)
	http.HandleFunc("/api/dns/v2", dnsV2.Handler)
	http.HandleFunc("/api/info", info.Handler)
	http.HandleFunc("/api/whois", whois.Handler)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalf("http listen failed: %s", err)
	}
}
