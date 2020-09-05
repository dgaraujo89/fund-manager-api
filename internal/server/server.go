package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/diegogomesaraujo/fund-manager-api/internal/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var server http.Server

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

// Start server to listen connections
func Start(config *config.Config) {
	port := strconv.FormatInt(config.Server.Port, 10)

	addressingListen := config.Server.Host + ":" + port

	router := configureRoutes()

	allowedOriginsCors := handlers.AllowedOrigins(config.Server.AllowOrigins)
	allowedMethodsCors := handlers.AllowedMethods(config.Server.AllowMethods)

	log.Printf("Server starting on %v...\n", addressingListen)

	server := &http.Server{
		Addr:    addressingListen,
		Handler: handlers.CORS(allowedOriginsCors, allowedMethodsCors)(router),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// // Shutdown
// func Shutdown() {
// 	context, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

// 	server.Shutdown()
// }
