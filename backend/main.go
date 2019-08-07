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

	if gc.Environment == "production" {
		router.Use(productionHeaders)
		router.Use(redrectNaked)
		router.Use(sentry.Recovery(raven.DefaultClient, false))
	}

	api := router.Group("/api")
	handlers.UserHandler(api.Group("/user"))
	handlers.RepoHandler(api.Group("/repo"))

	auth := router.Group("/auth")
	handlers.AuthHandler(auth)

	if gc.Environment == "production" {
		// Index
		router.StaticFile("/", "./dist/index.html")
		router.StaticFile("/index.html", "./dist/index.html")
		router.StaticFile("/index.htm", "./dist/index.html")

		// Assets that should be cached
		router.Use(cachingHeaders)
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

func redrectNaked(c *gin.Context) {
	host := c.Request.Host
	proto := c.Request.Header.Get("X-Forwarded-Proto")
	if strings.HasPrefix(host, "gitreleased.app") || proto != "https" {
		c.Redirect(http.StatusPermanentRedirect, "https://www.gitrelased.app")
		c.Abort()
	}
	c.Next()
}

func cachingHeaders(c *gin.Context) {
	c.Header("Cache-Control", "max-age=31622400, public")
	c.Next()
}

func productionHeaders(c *gin.Context) {
	c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	c.Next()
}
