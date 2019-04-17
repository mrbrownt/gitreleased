package main

import (
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Debugln("Not using .env")
	}

	setupGoth()
}

func main() {
	router := gin.Default()
	authHandler(router.Group(""))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	router.Run("0.0.0.0:" + port)
}

func setupGoth() {
	// This stupid thing is added just so goth will work, uses a temp cookie
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	store.Options(sessions.Options{
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		// Domain:   "gitreleased.app",
	})
	gothic.Store = store

	// Auth providers
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			"https://auth.gitreleased.app/callback/github",
		),
	)
}
