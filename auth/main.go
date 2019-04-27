package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/jinzhu/gorm"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"

	// Postgres and cloudql postgres drivers
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	err := envy.Load()
	if err != nil {
		log.Println("Not using .env")
	}
}

// DB is bad and should be removed ASAP
var DB *gorm.DB

func main() {
	gc, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	DB = gc.DB

	router := gin.Default()
	authHandler(router.Group(""))

	if gc.Environment == "production" {
		gin.SetMode("release")
	}

	port := envy.Get("PORT", "8082")
	router.Run("0.0.0.0:" + port)
}
