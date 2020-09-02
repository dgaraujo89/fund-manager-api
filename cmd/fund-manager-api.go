package main

import (
	"os"

	"github.com/diegogomesaraujo/fund-manager-api/internal/server"
)

const version = "0.0.1"

func main() {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "8080"
	}

	allowedOrigins := []string{"*"}
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	server.StartServer(host+":"+port, allowedOrigins, allowedMethods)

	os.Exit(0)
}
