package main

import (
	"fmt"
	"log"

	"github.com/diegogomesaraujo/fund-manager-api/internal/config"
	"github.com/diegogomesaraujo/fund-manager-api/internal/db"
	"github.com/diegogomesaraujo/fund-manager-api/internal/server"
	"github.com/thatisuday/commando"
)

const version = "0.0.1"

func main() {
	commando.
		SetExecutableName("fund-manager-api").
		SetVersion(version).
		SetDescription("Fund Manager API")

	commando.
		Register("server").
		SetShortDescription("Run server mode.").
		AddFlag("config,c", "The path to config file", commando.String, config.ConfigFile).
		SetAction(serverCommand)

	commando.
		Register("encrypt-db-password").
		SetShortDescription("This command encrypt the password database.").
		AddArgument("data", "The data to must be encrypted", "").
		SetAction(encryptDbPasswordCommand)

	commando.
		Register("test-database").
		SetShortDescription("This command test the database connection configuration").
		AddFlag("config,c", "The path to config file", commando.String, config.ConfigFile).
		SetAction(testDbConnCommand)

	commando.Parse(nil)
}

func serverCommand(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	configFile, _ := flags["config"].GetString()

	config := config.Load(configFile)

	server.Start(&config)

}

func encryptDbPasswordCommand(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	encryptedValue := db.EncryptDbUserPassword(args["data"].Value)

	fmt.Println()
	fmt.Println("Encrypted password:", encryptedValue)
	fmt.Println()
}

func testDbConnCommand(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	configFile, _ := flags["config"].GetString()

	config := config.Load(configFile)

	log.Println("Opening database connection...")

	db.Open(&config)

	log.Println("Database connected with success.")
	log.Println("Closing database connection...")

	db.Close()

	log.Println("Database connection closed.")
}
