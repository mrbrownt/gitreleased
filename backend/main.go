package main

import (
	"net/http"
	"os"
	"strings"

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

	//  Git mode must be set before gin.Default
	if gc.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(sentry.Recovery(raven.DefaultClient, false))

	api := router.Group("/api")
	handlers.UserHandler(api.Group("/user"))
	handlers.RepoHandler(api.Group("/repo"))

	auth := router.Group("/auth")
	handlers.AuthHandler(auth)

	if gc.Environment == "production" {
		router.Use(cachingHeaders())
		router.Use(redrectNaked())
		router.StaticFile("/", "./dist/index.html")
		router.StaticFile("/index.html", "./dist/index.html")
		router.StaticFile("/index.htm", "./dist/index.html")
		router.StaticFile("/favicon.ico", "./dist/favicon.ico")
		router.Static("/js", "./dist/js")
		router.Static("/css", "./dist/css")
		router.Static("/img", "./dist/img")
	}

	err := router.Run("0.0.0.0:" + gc.Port)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
	}
}

func redrectNaked() (middleware gin.HandlerFunc) {
	return func(c *gin.Context) {
		host := c.Request.Host
		if strings.HasPrefix(host, "gitreleased.app") {
			c.Redirect(http.StatusPermanentRedirect, "www.gitrelased.app")
		}

		c.Next()
	}
}

func cachingHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "max-age=31622400, public")
		c.Next()
	}
}
