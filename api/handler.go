package http

import (
	"fmt"
	"geoforecast/api/controllers"
	"geoforecast/internal/config"
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
		http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
			key := getRouteKey(r.Method, prefix)
			handler := handlers[key]
			if handler == nil {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			} else {
				handler(w, r)
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

	var error error
	for i := 0; i < config.Values.RESTART_ATTEMPTS; i++ {
		address := fmt.Sprintf("%s%d", "localhost:", config.Values.LISTEN_PORT+i)
		log.Printf("Starting http server on %s...", address)
		error = http.ListenAndServe(address, nil)
		log.Printf("Error %s", error)
	}

	log.Fatal(error)
}
