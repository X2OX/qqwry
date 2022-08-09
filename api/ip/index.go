package ip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.x2ox.com/utils/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cors.CORS(w, r, func(w http.ResponseWriter, r *http.Request) {
		gi := NewGoIP(r)

		switch strings.ToLower(r.URL.Query().Get("type")) {
		case "json":
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			_ = json.NewEncoder(w).Encode(gi)
		default:
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
			_, _ = w.Write([]byte(gi.String()))
		}
	})
}

type GeoIP struct {
	IP       string `json:"ip"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
}

func NewGoIP(r *http.Request) GeoIP {
	return GeoIP{
		IP:       r.Header.Get("X-Real-Ip"),
		Country:  r.Header.Get("X-Vercel-Ip-Country"),
		Province: r.Header.Get("X-Vercel-Ip-Country-Region"),
		City:     r.Header.Get("X-Vercel-Ip-City"),
	}
}

func (g GeoIP) String() string {
	return fmt.Sprintf("%s\n%s %s %s", g.IP, g.Country, g.Province, g.City)
}
