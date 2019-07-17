package main

import (
	"fmt"
	"log"
	"os"

	"github.com/getsentry/raven-go"
	"github.com/gobuffalo/envy"
	"github.com/jinzhu/gorm"

	// Postgres and cloudql postgres drivers
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var environment string
var baseURL string

func config() {
	var err error
	environment = envy.Get("ENVIRONMENT", "development")
	baseURL = envy.Get("BASE_URL", "localhost")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envy.Get("DB_HOST", "localhost"),
		envy.Get("DB_PORT", "5432"),
		envy.Get("DB_USER", "postgres"),
		envy.Get("DB_PASS", ""),
		envy.Get("DB_NAME", "gitreleased"),
		envy.Get("DB_SSLMODE", "disable"),
	)

	dbConnType := ""
	if cloudsql := envy.Get("CLOUDSQL", ""); cloudsql != "" {
		dbConnType = "cloudsqlpostgres"
	} else {
		dbConnType = "postgres"
	}

	db, err = gorm.Open(dbConnType, dsn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	debug := os.Getenv("DEBUG")
	if debug != "" {
		db.LogMode(true)
	}

	sentryDSN := envy.Get("SENTRY_DSN", "")
	if sentryDSN != "" {
		raven.SetDSN(sentryDSN)
	}
}
