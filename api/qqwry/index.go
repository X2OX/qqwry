package qqwry

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"go.x2ox.com/qqwry/data"
	"go.x2ox.com/utils/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cors.CORS(w, r, QQwry)
}

func QQwry(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	ip := query.Get("ip")
	if ip == "" {
		ip = getRemoteAddr(r)
	}

	cz := data.New().Find(ip)

	_type := strings.ToLower(query.Get("type"))
	switch {
	case cz.Error != nil:
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		_, _ = w.Write([]byte(cz.Error.Error()))
	case _type == "json":
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		_ = json.NewEncoder(w).Encode(cz)
	default:
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		_, _ = w.Write([]byte(cz.Country + " " + cz.Area))
	}
}

func getRemoteAddr(r *http.Request) string {
	if ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]); ip != "" {
		return ip
	}
	if ip := strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
