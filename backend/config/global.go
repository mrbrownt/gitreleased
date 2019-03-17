package config

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
)

// Global config required for application
type Global struct {
	Port      string
	BaseURL   string
	DB        *gorm.DB
	callCount int
}

var globalConf Global

// New global config, only call this once from main
func New() (gc Global, err error) {
	globalConf.callCount = globalConf.callCount + 1
	if globalConf.callCount != 1 {
		return gc, errors.New("conf.New() callend more than once")
	}

	globalConf.Port = getEnv("PORT", "3000")
	globalConf.BaseURL = getEnv("BASE_URL", "localhost:"+globalConf.Port)

	err = setupDB()
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
	if globalConf.callCount != 1 {
		return gc, errors.New("config.New() hasn't been called")
	}
	return globalConf, err
}
