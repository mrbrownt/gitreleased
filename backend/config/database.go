package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func setupDB() (err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "26257"),
		getEnv("DB_USER", "gitreleased"),
		getEnv("DB_NAME", "migrations"),
	)

	db, err := gorm.Open("postgres", dsn)

	debug := os.Getenv("DEBUG")
	if debug != "" {
		db.LogMode(true)
	}

	globalConf.DB = db
	return err
}
