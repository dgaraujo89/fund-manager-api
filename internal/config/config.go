package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// ConfigFile default
const ConfigFile = "configs/config.toml"

// Load is a function read de configurations
func Load(filePath string) Config {
	log.Println("loading configs...")

	var config Config

	if filePath == ConfigFile {
		path, err := os.Getwd()

		if err != nil {
			log.Fatalln("An error has occurred when get the current path")
			panic(err)
		}

		filePath = path + "/" + ConfigFile
	}

	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}

	return config
}
