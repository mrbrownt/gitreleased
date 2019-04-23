package config

import (
	"errors"

	"github.com/gobuffalo/envy"
	"github.com/jinzhu/gorm"
)

// Global config required for application
type Global struct {
	Port        string
	BaseURL     string
	DB          *gorm.DB
	Environment string
	callCount   int
}

var globalConf Global

// New global config, only call this once from main
func New() (gc Global, err error) {
	globalConf.callCount = globalConf.callCount + 1
	if globalConf.callCount != 1 {
		return gc, errors.New("conf.New() callend more than once")
	}

	globalConf.Port = envy.Get("PORT", "3000")
	globalConf.BaseURL = envy.Get("BASE_URL", "localhost:"+globalConf.Port)
	globalConf.Environment = envy.Get("ENVIRONMENT", "development")

	err = setupDB()
	return globalConf, err
}

// Get global config after it's been setup
func Get() (gc Global, err error) {
	if globalConf.callCount != 1 {
		return gc, errors.New("config.New() hasn't been called")
	}
	return globalConf, err
}
