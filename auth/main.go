package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
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

	setupGoth(gc)

	router := gin.Default()
	authHandler(router.Group(""))

	if gc.Environment == "production" {
		gin.SetMode("release")
	}

	port := envy.Get("PORT", "8082")
	router.Run("0.0.0.0:" + port)
}

func setupGoth(config config.Global) {

	// This stupid thing is added just so goth will work, uses a temp cookie
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	var secure = false
	if config.Environment == "production" {
		secure = true
	}

	store.Options(sessions.Options{
		Secure:   secure,
		HttpOnly: true,
		Path:     "/",
		Domain:   config.BaseURL,
	})
	gothic.Store = store

	callbackURL := fmt.Sprintf("%s/callback/github", config.BaseURL)

	// Auth providers
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			callbackURL,
		),
	)
}
