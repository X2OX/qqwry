package dns

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.x2ox.com/qqwry/dot"
	"go.x2ox.com/utils/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cors.CORS(w, r, Info)
}

func Info(w http.ResponseWriter, r *http.Request) {
	var (
		query    = r.URL.Query()
		domain   = query.Get("name")
		rrType   = query.Get("type")
		provider = query.Get("provider")
		output   = query.Get("output")
	)

	p := dot.NewParser(domain, rrType)
	if p == nil {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		_, _ = w.Write([]byte("wrong parameter"))
		return
	}

	result := p.Parse(provider)

	_type := strings.ToLower(output)
	switch {
	case _type == "json":
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		_ = json.NewEncoder(w).Encode(result)
	case _type == "ip":
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		_, _ = w.Write([]byte(result.IP()))
	default:
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		_, _ = w.Write([]byte(result.String()))
	}
}
