package main

import (
	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
	"gitlab.com/mrbrownt/gitreleased.app/backend/handlers"
)

func init() {
	_ = godotenv.Load()

	err := raven.SetDSN(os.Getenv("SENTRY_DSN"))
	if err != nil {
		logrus.Fatalln(err)
	}
}

func main() {
	gc, err := config.New()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer gc.DB.Close()

	err = config.Goth()
	if err != nil {
		logrus.Fatalln(err)
	}

	router := gin.Default()

	router.Use(sentry.Recovery(raven.DefaultClient, false))

	api := router.Group("/api")

	handlers.AuthHandler(router.Group("/auth"))
	handlers.UserHandler(api.Group("/user"))
	handlers.RepoHandler(api.Group("/repo"))

	err = router.Run("0.0.0.0:" + gc.Port)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
}
