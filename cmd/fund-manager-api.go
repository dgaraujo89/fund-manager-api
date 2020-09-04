package main

import (
	"fmt"

	"github.com/diegogomesaraujo/fund-manager-api/internal/config"
	"github.com/diegogomesaraujo/fund-manager-api/internal/crypt"
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

	commando.Parse([]string{"help"})
}

func serverCommand(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	configFile, _ := flags["config"].GetString()

	config := config.Load(configFile)

	server.Start(&config)

}

func encryptDbPasswordCommand(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	value := args["data"].Value

	encryptedValue := crypt.EncryptAsString([]byte(value), "database")

	fmt.Println()
	fmt.Println("Encrypted password:", encryptedValue)
	fmt.Println()
}
