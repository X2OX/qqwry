package main

import (
	"log"
	"net/http"

	"go.x2ox.com/qqwry/api/dns"
	"go.x2ox.com/qqwry/api/qqwry"
)

func main() {
	http.HandleFunc("/api/qqwry", qqwry.Handler)
	http.HandleFunc("/api/dns", dns.Handler)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalf("http listen failed: %s", err)
	}
}
