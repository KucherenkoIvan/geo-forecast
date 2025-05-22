package http

import (
	"fmt"
	"geoforecast/api/controllers"
	"geoforecast/internal/config"
	"geoforecast/web/templates"
	"github.com/a-h/templ"
	"log"
	"net/http"
	"strings"
)

var (
	handlers            map[string]http.HandlerFunc
	registered_prefixes []string
)

func getRouteKey(method string, prefix string) string {
	return fmt.Sprintf("%s %s", method, prefix)
}

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header["Authorization"]
		if len(auth_header) != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		auth_content := strings.Split(auth_header[0], " ")

		if len(auth_content) != 2 || auth_content[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := auth_content[1]
		if !tokenAuth(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}

func tokenAuth(token string) bool {
	auth := false

	for i := 0; i < len(config.Values.ACCEPT_KEYS); i++ {
		if strings.Compare(config.Values.ACCEPT_KEYS[i], token) == 0 {
			auth = true
		}
	}

	return auth
}

func route(method string, prefix string, handler http.HandlerFunc) {
	handlers[getRouteKey(method, prefix)] = handler

	is_registered := false
	for i := 0; i < len(registered_prefixes); i++ {
		is_registered = is_registered || (prefix == registered_prefixes[i])
	}

	if !is_registered {
		log.Printf("[HTTP] Mapped %s %s", method, prefix)
		http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
			key := getRouteKey(r.Method, prefix)
			handler := handlers[key]
			if handler == nil {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			} else {
				log.Printf("[HTTP] Request on %s %s", method, prefix)
				handler(w, r)
				log.Printf("[HTTP] Handled %s %s", method, prefix)
			}
		})
	}

	registered_prefixes = append(registered_prefixes, prefix)
}

func Start() {
	handlers = make(map[string]http.HandlerFunc)
	registered_prefixes = []string{}

	route("GET", "/app/info", controllers.AppInfo)
	route("POST", "/api/position_log", withAuth(controllers.PositionLog))
	route("GET", "/api/tracks", withAuth(controllers.TracksList))
	route("GET", "/api/track", withAuth(controllers.Track))

	route("GET", "/static", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("alkwdjlawjdawldkjalwkdjwlka")
		fmt.Println(r.URL)

	})

	route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.Contains(r.URL.Path, "/static/") {
			http.FileServer(http.Dir("../web")).ServeHTTP(w, r)
			return
		}

		templ.Handler(templates.IndexPage()).ServeHTTP(w, r)
	})

	var error error
	for i := 0; i < config.Values.RESTART_ATTEMPTS; i++ {
		address := fmt.Sprintf("%s%d", ":", config.Values.LISTEN_PORT+i)
		log.Printf("[HTTP] Starting http server on %s...", address)
		error = http.ListenAndServe(address, nil)
		log.Printf("Error %s", error)
	}

	log.Fatal(error)
}
