package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
)

var db *gorm.DB

// GetDB so we don't have to pass DB down the chain
func GetDB() *gorm.DB {
	return db
}

// SetupDB for useage through out application
func SetupDB() (newDB *gorm.DB, err error) {
	gc, err := config.Get()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		gc.DBHost,
		gc.DBPort,
		gc.DBUser,
		gc.DBName,
	)

	db, err = gorm.Open("postgres", dsn)

	debug := os.Getenv("DEBUG")
	if debug != "" {
		db.LogMode(true)
	}

	return db, err
}

// AutoMigrate database, eventually replace with proper migrations
func AutoMigrate() {
	db.AutoMigrate(
		&User{},
		&Release{},
		&Repository{},
	)
}
