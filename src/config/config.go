package config

import (
	"flag"
	"os"
)

type Config struct {
	ValidAPIKey string
}

var config Config

func InitConfig() {
	flag.StringVar(&config.ValidAPIKey, "ValidAPIKey", os.Getenv("API_KEY"), "API key for authentication")
	flag.Parse()
}

func GetApiKey() string {
	return config.ValidAPIKey
}
