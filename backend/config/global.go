package config

import (
	"errors"
	"os"
)

// Global config required for application
type Global struct {
	BaseURL string
	Port    string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

var globalConf Global
var callCount int

// New global config, only call this once from main
func New() (gc Global, err error) {
	callCount = callCount + 1
	if callCount != 1 {
		return gc, errors.New("conf.New() callend more than once")
	}

	globalConf.Port = getEnv("PORT", "3000")
	globalConf.BaseURL = getEnv("BASE_URL", "localhost:"+globalConf.Port)
	globalConf.DBHost = getEnv("DB_HOST", "127.0.0.1")
	globalConf.DBPort = getEnv("DB_PORT", "26257")
	globalConf.DBUser = getEnv("DB_USER", "gitreleased")
	globalConf.DBName = getEnv("DB_NAME", "gitreleased")

	return globalConf, err
}

func getEnv(key, def string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = def
	}
	return value
}

// Get global config after it's been setup
func Get() (gc Global, err error) {
	if callCount != 1 {
		return gc, errors.New("config.New() hasn't been called")
	}
	return globalConf, err
}
