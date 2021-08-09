package main

import (
	"log"
	"net/http"

	"go.x2ox.com/qqwry/api/qqwry"
)

func main() {
	http.HandleFunc("/api/qqwry", qqwry.Handler)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalf("http listen failed: %s", err)
	}
}
