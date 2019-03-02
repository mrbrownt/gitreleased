package config

import (
	"errors"
	"os"
)

// Global config required for application
type Global struct {
	BaseURL string
	Port    string
}

var globalConf Global
var callCount int

// New global config, only call this once from main
func New() (gc Global, err error) {
	callCount = callCount + 1
	if callCount != 1 {
		return gc, errors.New("conf.New() callend more than once")
	}

	globalConf.Port = os.Getenv("PORT")
	if globalConf.Port == "" {
		globalConf.Port = "3000"
	}

	globalConf.BaseURL = os.Getenv("BASE_URL")
	if globalConf.BaseURL == "" {
		globalConf.BaseURL = "localhost:" + globalConf.Port
	}

	return globalConf, err
}

// Get global config after it's been setup
func Get() (gc Global, err error) {
	if callCount != 1 {
		return gc, errors.New("config.New() hasn't been called")
	}
	return globalConf, err
}
