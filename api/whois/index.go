package whois

import (
	"net/http"

	"go.x2ox.com/utils/cors"
	"go.x2ox.com/whois"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cors.CORS(w, r, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		d, err := whois.Parse(r.URL.Query().Get("domain"))
		if err == nil {
			_, _ = w.Write([]byte(d.Query()))
		} else {
			_, _ = w.Write([]byte(err.Error()))
		}
	})
}
