package config

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/jinzhu/gorm"
)

// Global config required for application
type Global struct {
	Port        string
	BaseURL     string
	DB          *gorm.DB
	Environment string
	initialized bool
}

var globalConf Global

// Get global config, only call this once from main
func Get() (gc Global) {
	if globalConf.initialized {
		return globalConf
	}

	globalConf.Port = envy.Get("PORT", "3000")
	globalConf.BaseURL = envy.Get("BASE_URL", "localhost:8080")
	globalConf.Environment = envy.Get("ENVIRONMENT", "development")

	err := setupDB()
	if err != nil {
		log.Fatalln(err)
	}

	setupGoth()

	globalConf.initialized = true
	return globalConf
}
