package models

import (
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// GetDB so we don't have to pass DB down the chain
func GetDB() *gorm.DB {
	return db
}

// SetupDB for useage through out application
func SetupDB() (newDB *gorm.DB, err error) {
	db, err = gorm.Open(
		"mysql",
		// TODO: paramaterize this
		"root:root@tcp(127.0.0.1:3316)/gomigrate?charset=utf8&parseTime=True",
	)

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
