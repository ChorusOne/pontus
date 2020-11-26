package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv configures environment variables from a provided path or default location
func LoadEnv() {
	var configFileFlag = flag.String("config", "../config/config.env", "Pontus config.env file location.")
	err := godotenv.Load(*configFileFlag)
	if err != nil {
		log.Fatalln("Failed to open Pontus config file.")
	}
}
