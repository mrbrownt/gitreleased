package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
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
	config := newConfig()

	db, err := gorm.Open("postgres", config.postgresConnection)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db

	setupGoth(config)

	router := gin.Default()
	authHandler(router.Group(""))

	if config.environment == "production" {
		gin.SetMode("release")
	}

	port := envy.Get("PORT", "8082")
	router.Run("0.0.0.0:" + port)
}

func setupGoth(config config) {

	// This stupid thing is added just so goth will work, uses a temp cookie
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	var secure = false
	if config.environment == "production" {
		secure = true
	}

	store.Options(sessions.Options{
		Secure:   secure,
		HttpOnly: true,
		Path:     "/",
		Domain:   config.fullURL,
	})
	gothic.Store = store

	// Auth providers
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			config.callbackURL,
		),
	)
}

type config struct {
	environment        string
	baseURL            string
	fullURL            string
	callbackURL        string
	postgresConnection string
}

func newConfig() (c config) {
	c.environment = os.Getenv("ENVIRONMENT")
	if c.environment == "" {
		c.environment = "development"
	}

	c.baseURL = os.Getenv("BASE_URL")
	if c.baseURL == "" {
		c.baseURL = "localhost"
		c.fullURL = c.baseURL
		c.callbackURL = "http://localhost:8082/callback/github"
	} else {
		c.fullURL = fmt.Sprintf("auth.%s", c.baseURL)
		c.callbackURL = fmt.Sprintf("http://%s/callback/github", c.fullURL)
	}

	c.postgresConnection = fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s sslmode=%s",
		envy.Get("DB_NAME", "gitreleased"),
		envy.Get("DB_USER", "postgres"),
		envy.Get("DB_PASS", ""),
		envy.Get("DB_HOST", "localhost"),
		envy.Get("DB_SSLMODE", "disable"),
	)

	return c
}
