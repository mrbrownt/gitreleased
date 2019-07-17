package main

import (
	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
	"gitlab.com/mrbrownt/gitreleased.app/backend/handlers"

	// Postgres and cloudql postgres drivers
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	_ = godotenv.Load()

	err := raven.SetDSN(os.Getenv("SENTRY_DSN"))
	if err != nil {
		logrus.Fatalln(err)
	}
}

func main() {
	gc := config.Get()
	defer gc.DB.Close()

	router := gin.Default()

	router.Use(sentry.Recovery(raven.DefaultClient, false))

	api := router.Group("/api")

	handlers.UserHandler(api.Group("/user"))
	handlers.RepoHandler(api.Group("/repo"))

	auth := router.Group("/auth")
	handlers.AuthHandler(auth)

	err := router.Run("0.0.0.0:" + gc.Port)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
}
