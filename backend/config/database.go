package config

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"

	"github.com/gobuffalo/envy"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jinzhu/gorm"

	// Allows migratations from files
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func setupDB() (err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envy.Get("DB_HOST", "db"),
		envy.Get("DB_PORT", "5432"),
		envy.Get("DB_USER", "postgres"),
		envy.Get("DB_PASS", "pass"),
		envy.Get("DB_NAME", "gitreleased"),
		envy.Get("DB_SSLMODE", "disable"),
	)

	dbConnType := ""
	if cloudsql := envy.Get("CLOUDSQL", ""); cloudsql != "" {
		dbConnType = "cloudsqlpostgres"
	} else {
		dbConnType = "postgres"
	}

	db, err := gorm.Open(dbConnType, dsn)
	if err != nil {
		return err
	}

	debug := os.Getenv("DEBUG")
	if debug != "" {
		db.LogMode(true)
	}

	err = migrateDB(db)
	if err != nil {
		return err
	}

	globalConf.DB = db
	return err
}

func migrateDB(db *gorm.DB) (err error) {
	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{
		DatabaseName: envy.Get("DB_NAME", "gitreleased"),
	})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}
