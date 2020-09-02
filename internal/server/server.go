package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func configureRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, routes := range routesList {
		for _, route := range routes {
			router.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}

	return router
}

// StartServer to listen connections
func StartServer(addressingListen string, allowedOrigins []string, allowedMethods []string) {
	router := configureRoutes()

	allowedOriginsCors := handlers.AllowedOrigins(allowedOrigins)
	allowedMethodsCors := handlers.AllowedMethods(allowedMethods)

	log.Printf("Server starting on %v...\n", addressingListen)

	err := http.ListenAndServe(addressingListen, handlers.CORS(allowedOriginsCors, allowedMethodsCors)(router))

	if err != nil {
		log.Fatal(err)
	}
}
