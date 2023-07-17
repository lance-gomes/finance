package config

import (
	"os"
)

type Configuration struct {
	Username    string
	Password    string
	ClientID    string
	Secret      string
	Environment string
	Countries string
	Products string
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func New() *Configuration {
	return &Configuration{
		Username: getEnv("username", ""),
		Password: getEnv("password", ""),
		ClientID: getEnv("client_id", ""),
		Secret:   getEnv("secret", ""),
		Environment:   getEnv("environment", ""),
		Countries: getEnv("countries", ""),
		Products: getEnv("products", ""),
	}
}
