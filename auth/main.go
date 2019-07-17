package main

import (
	"fmt"
	"os"

	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func main() {
	envy.Load()
	config()
	setupGoth()

	router := gin.Default()

	if sentryDSN := envy.Get("SENTRY_DSN", ""); sentryDSN != "" {
		router.Use(sentry.Recovery(raven.DefaultClient, false))
	}

	authHandler(router.Group("/auth"))

	if environment == "production" {
		gin.SetMode("release")
	}

	port := envy.Get("PORT", "8082")
	router.Run("0.0.0.0:" + port)
}

func setupGoth() {

	// This stupid thing is added just so goth will work, uses a temp cookie
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	var secure = false
	if environment == "production" {
		secure = true
	}

	store.Options(sessions.Options{
		Secure:   secure,
		HttpOnly: true,
		Path:     "/",
		Domain:   baseURL,
	})
	gothic.Store = store

	callbackURL := fmt.Sprintf("http://%s:8080/auth/callback/github", baseURL)

	// Auth providers
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			callbackURL,
		),
	)
}
